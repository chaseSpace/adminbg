package crud

import (
	"adminbg/cerror"
	"adminbg/cproto"
	"adminbg/pkg/g"
	"adminbg/pkg/model"
	"adminbg/util/_gorm"
	"fmt"
	"github.com/pkg/errors"
	"strings"
)

func InsertUserGroup(group *cproto.Group) error {
	sql := fmt.Sprintf(`
		INSERT INTO %s (group_name, role_id)
		SELECT ?, ?
		WHERE EXISTS (
			SELECT 1
			FROM %s
			WHERE role_id = ?
		)`, TN.UserGroup, TN.Role)
	exec := g.MySQL.Exec(sql, group.GroupName, group.RoleId, group.RoleId)

	if exec.Error != nil {
		if strings.Contains(exec.Error.Error(), "Duplicate") {
			return errors.New("group name exists")
		}
		return cerror.WrapMysqlErr(exec.Error)
	}
	if exec.RowsAffected == 0 {
		return errors.Wrap(cerror.ErrParams, fmt.Sprintf("role_id:%d not exist", group.RoleId))
	}
	return nil
}

func DelUserGroup(gid int16) error {
	exec := g.MySQL.Take(&model.User{}, "group_id=?", gid)
	if _gorm.IsDBErr(exec.Error) {
		return exec.Error
	}
	if exec.RowsAffected > 0 {
		return errors.Errorf("this group is being used, can't be delete")
	}
	exec = g.MySQL.Delete(&model.UserGroup{}, "group_id=?", gid)
	return cerror.WrapDeleteDBOneRecordErr(exec)
}

func UpdateUserGroup(group *cproto.Group) error {
	role, err := QueryRole(group.RoleId)
	if err != nil {
		return err
	}
	if role == nil {
		return errors.Wrap(cerror.ErrParams, fmt.Sprintf("role_id:%d not exist", group.RoleId))
	}
	sql := fmt.Sprintf(`
		UPDATE %s
		SET group_name = ?, role_id = ?
		WHERE group_id = ?`, TN.UserGroup)

	exec := g.MySQL.Exec(sql, group.GroupName, group.RoleId, group.GroupId)
	if exec.Error != nil {
		if strings.Contains(exec.Error.Error(), "Duplicate") {
			return errors.New("group name exists")
		}
		return cerror.WrapMysqlErr(exec.Error)
	}
	if exec.RowsAffected == 0 {
		return cerror.ErrNothingUpdated
	}
	return nil
}

func QueryUserGroup(gid int32) (*cproto.Group, error) {
	sql := fmt.Sprintf("SELECT * FROM %s WHERE deleted_at is null and group_id=?", TN.UserGroup)
	row := new(model.UserGroup)
	err := g.MySQL.Raw(sql, gid).Scan(row).Error
	if _gorm.IsDBErr(err) {
		return nil, err
	}
	if row.GroupId < 1 {
		return nil, fmt.Errorf("group_id:%d not found", gid)
	}
	return row.Proto(), nil
}

// Used by only administrator
func GetUserGroupList(pageNum, pageSize uint16, orderByParams ...OrderByOption) ([]*model.UserGroup, int64, error) {
	var total int64
	err := g.MySQL.Model(new(model.UserGroup)).Count(&total).Error
	if err != nil {
		return nil, 0, cerror.WrapMysqlErr(err)
	}
	if total <= int64((pageNum-1)*pageSize) {
		return nil, total, nil
	}
	list := make([]*model.UserGroup, 0)
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
