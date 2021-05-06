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
