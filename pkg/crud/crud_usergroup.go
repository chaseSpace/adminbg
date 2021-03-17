package crud

import (
	"adminbg/cerror"
	"adminbg/pkg/g"
	"fmt"
	"github.com/pkg/errors"
)

func InsertUserGroup(groupName string, roleId int16) error {
	sql := fmt.Sprintf(`
		INSERT INTO %s (group_name, role_id)
		SELECT ?, ?
		WHERE EXISTS (
			SELECT 1
			FROM %s
			WHERE role_id = ?
		)`, TN.UserGroup, TN.Role)
	exec := g.MySQL.Exec(sql, groupName, roleId, roleId)
	if exec.Error != nil {
		return cerror.WrapMysqlErr(exec.Error)
	}
	if exec.RowsAffected != 1 {
		return errors.Wrap(cerror.ErrParams, fmt.Sprintf("roleId:%d not exist", roleId))
	}
	return nil
}
