package crud

import (
	"adminbg/cerror"
	"adminbg/cproto"
	"adminbg/pkg/g"
	"adminbg/pkg/model"
	"adminbg/util/_gorm"
	"fmt"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// Temp struct is used by crud pkg
type UserIdentity struct {
	Uid             int32
	AccountId       string
	ContainsDeleted bool
}

func CheckUserPassport(accountId string, plainPwd string) (*model.User, error) {
	// It use mysql's SHA1 function compute encrypted password here.
	sql := fmt.Sprintf(`
		SELECT *
		FROM %s
		WHERE account_id = ?
			AND encrypted_pwd = SHA1(CONCAT(?, (
				SELECT salt
				FROM %s
				WHERE account_id = ?
				LIMIT 1
			)));
		`, TN.User, TN.User)
	row := model.User{}
	err := g.MySQL.Raw(sql, accountId, plainPwd, accountId).Scan(&row).Error
	return &row, cerror.WrapMysqlErr(err)
}

func GetUserByUid(uid int32) (*model.User, error) {
	row := model.User{}
	err := g.MySQL.Take(&row, "uid=?", uid).Error
	return &row, cerror.WrapMysqlErr(err)
}

func GetUserByAccountId(accountId string) (*model.User, error) {
	row := model.User{}
	err := g.MySQL.Take(&row, "account_id=?", accountId).Error
	return &row, cerror.WrapMysqlErr(err)
}

// Get a valid user's basic info(contains roleID,groupID)
func GetUserBase(uid int32) (*cproto.User, error) {
	sql := fmt.Sprintf(`
		SELECT a.*, b.group_id, c.role_id
		FROM %s a
				 LEFT JOIN %s b ON a.uid = b.uid
				 LEFT JOIN %s c ON b.group_id = c.group_id
		WHERE a.uid = ?`, TN.User, TN.UserGroupRef, TN.UserGroup)
	row := new(model.User)
	exec := g.MySQL.Raw(sql, uid).Scan(row)
	if _gorm.IsDBErr(exec.Error) {
		return nil, cerror.WrapMysqlErr(exec.Error)
	}
	if exec.RowsAffected != 1 {
		return nil, errors.New(fmt.Sprintf("uid:%d not found", uid))
	}
	return row.Proto(), nil
}

func InsertUser(entity *model.UserBase) error {
	// All validation should be completed in outside.
	insertUser := fmt.Sprintf(`
		INSERT INTO %s (account_id, encrypted_pwd, salt, nick_name, phone
								 , email, sex, remark, status)
		SELECT ?, SHA1(CONCAT(?, ?)), ?
			 , ?, ?, ?, ?
			 , ?, ?
		WHERE NOT EXISTS (
				SELECT 1
				FROM %s
				WHERE account_id = ?
			)`, TN.User, TN.User)
	insertUGRef := fmt.Sprintf(`
		INSERT INTO %s (uid, group_id)
		SELECT ?, ?
		WHERE EXISTS (
			SELECT 1
			FROM %s
			WHERE group_id = ?
		)`, TN.UserGroupRef, TN.UserGroup)
	err := g.MySQL.Transaction(func(tx *gorm.DB) error {
		exec := tx.Exec(insertUser, entity.AccountId, entity.EncryptedPwd, entity.Salt, entity.Salt, entity.NickName,
			entity.Phone, entity.Email, entity.Sex, entity.Remark, entity.Status, entity.AccountId)
		if exec.Error != nil {
			return cerror.WrapMysqlErr(exec.Error)
		}
		// We don't have to insert row to table `YOUR_TABLE_PREFIX_user_group_ref` here.
		if exec.RowsAffected != 1 {
			return errors.Wrap(cerror.ErrParams, "account_id exists")
		}
		user := new(model.User)
		err := tx.Take(user, "account_id=?", entity.AccountId).Error
		if _gorm.IsDBErr(err) {
			return cerror.WrapMysqlErr(err)
		}
		exec = tx.Exec(insertUGRef, user.Uid, entity.GroupId, entity.GroupId)
		if exec.Error != nil {
			return cerror.WrapMysqlErr(err)
		}
		if exec.RowsAffected != 1 {
			return errors.Wrap(cerror.ErrParams, "group_id not exists")
		}
		return nil
	})
	return err
}

func DeleteUser(userIdt UserIdentity) (bool, error) {
	var exec = g.MySQL
	if userIdt.Uid != 0 {
		exec = exec.Delete(new(model.User), "uid=?", userIdt.Uid)
	} else if userIdt.AccountId != "" {
		exec = g.MySQL.Delete(new(model.User), "account_id=?", userIdt.AccountId)
	} else {
		return false, cerror.ErrParams
	}
	return exec.RowsAffected == 1, cerror.WrapMysqlErr(exec.Error)
}

// Update user info, that is not permitted to modify account_id by default.
func UpdateUserBase(uid int32, alter *cproto.User, containsDeleted bool) error {
	sql := fmt.Sprintf(`UPDATE %s
		SET nick_name = ?, phone = ?, email = ?, sex = ?, status = ?, remark = ?, 
			encrypted_pwd = if(? = '', encrypted_pwd, SHA1(CONCAT(?, salt)))
		WHERE`, TN.User)

	expr := clause.Expr{
		SQL: sql,
		Vars: []interface{}{alter.Name, alter.Phone, alter.Email, alter.Sex, alter.Status,
			alter.Remark, alter.Pwd, alter.Pwd},
	}

	expr.SQL += "uid=?"
	expr.Vars = append(expr.Vars, uid)

	if !containsDeleted {
		expr.SQL += " AND deleted_at IS NULL"
	}
	ret := g.MySQL.Exec(expr.SQL, expr.Vars...)
	if ret.RowsAffected == 0 {
		return cerror.ErrNothingUpdated
	}
	return cerror.WrapMysqlErr(ret.Error)
}

func UpdateUserGroupRef(uid int32, groupId int16, delete bool) error {
	var sql string
	if delete {
		sql = `delete from %s where uid=? and group_id=?`
	} else {
		sql = `insert into %s (uid, group_id) values (?,?) on duplicate key update group_id = group_id`
	}
	exec := g.MySQL.Exec(fmt.Sprintf(sql, TN.UserGroupRef), uid, groupId)
	if exec.Error != nil {
		return exec.Error
	}
	if exec.RowsAffected != 1 {
		if delete {
			return cerror.ErrNothingDeleted
		}
		return cerror.ErrYouHaveDoneIt
	}
	return nil
}

// Used by only administrator
func GetUserList(pageNum, pageSize uint16, orderByParams ...OrderByOption) ([]*model.User, int64, error) {
	var total int64
	err := g.MySQL.Model(new(model.User)).Count(&total).Error
	if err != nil {
		return nil, 0, cerror.WrapMysqlErr(err)
	}
	if total <= int64((pageNum-1)*pageSize) {
		return nil, total, nil
	}
	list := make([]*model.User, 0)
	offset := (pageNum - 1) * pageSize

	order := "created_at desc"
	if len(orderByParams) > 0 {
		order = GenOrderByClause(orderByParams...)
	}
	err = g.MySQL.Order(order).Offset(int(offset)).Limit(int(pageSize)).Find(&list).Error
	if err != nil {
		return nil, 0, cerror.WrapMysqlErr(err)
	}
	return list, total, nil
}
