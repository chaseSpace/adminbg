package crud

import (
	"adminbg/cerror"
	"adminbg/cproto"
	"adminbg/pkg/g"
	"adminbg/pkg/model"
	"fmt"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

/*
	For sometimes, we need coding a raw SQL at working, at this time,
we shouldn't directly write table name within code, that's a tight-coupling way.
Here we created a struct variable `TN` to control all tables' name, for convenience,
and beauty.
*/

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

func InsertUser(entity *model.UserBase) error {
	// Any validation should be completed in outside.
	sql := fmt.Sprintf(`
		INSERT INTO %s (account_id, encrypted_pwd, salt, nick_name, phone
								 , email, sex, remark, status)
		SELECT ?, SHA1(CONCAT(?, ?)), ?
			 , ?, ?, ?, ?
			 , ?, ?
		WHERE NOT EXISTS (
				SELECT 1
				FROM %s
				WHERE account_id = ?
			);
		`, TN.User, TN.User)
	ret := g.MySQL.Exec(sql, entity.AccountId, entity.EncryptedPwd, entity.Salt, entity.Salt, entity.NickName, entity.Phone, entity.Email,
		entity.Sex, entity.Remark, entity.Status, entity.AccountId)
	// todo: insert into user_group_ref
	if ret.RowsAffected == 0 {
		return errors.Wrap(cerror.ErrParams, "account_id exists")
	}
	return ret.Error
}

func DeleteUser(userIdt UserIdentity) (bool, error) {
	var ret *gorm.DB
	if userIdt.Uid != 0 {
		ret = g.MySQL.Delete(new(model.User), "uid=?", userIdt.Uid)
	} else if userIdt.AccountId != "" {
		ret = g.MySQL.Delete(new(model.User), "account_id=?", userIdt.AccountId)
	} else {
		return false, nil
	}
	return ret.RowsAffected == 1, ret.Error
}

// Update user info, that is not permitted to modify account_id by default.
func UpdateUser(userIdt UserIdentity, alter *cproto.UpdateUserReq) error {
	sql := `
		UPDATE adminbg_user
		SET nick_name = ?, phone = ?, email = ?, sex = ?, status = ?, group_id = ?, remark = ?, 
			encrypted_pwd = if(? = '', encrypted_pwd, SHA1(CONCAT(?, salt)))
		WHERE 
	`

	expr := clause.Expr{
		SQL: sql,
		Vars: []interface{}{alter.Name, alter.Phone, alter.Email, alter.Sex, alter.Status, alter.GroupId,
			alter.Remark, alter.Pwd, alter.Pwd},
	}

	if userIdt.Uid != 0 {
		expr.SQL += "uid=?"
		expr.Vars = append(expr.Vars, userIdt.Uid)
	} else if userIdt.AccountId != "" {
		expr.SQL += "account_id=?"
		expr.Vars = append(expr.Vars, userIdt.AccountId)
	} else {
		return nil
	}

	if !userIdt.ContainsDeleted {
		expr.SQL += " AND deleted_at IS NULL"
	}
	ret := g.MySQL.Exec(expr.SQL, expr.Vars...)
	if ret.RowsAffected == 0 {
		return cerror.ErrNothingUpdated
	}
	return ret.Error
}
