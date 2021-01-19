package crud

import (
	"adminbg/cerror"
	"adminbg/pkg/g"
	"adminbg/pkg/model"
	"fmt"
)

func GetUserGroup(uid int32) ([]*model.UserGroup, error) {
	var row []*model.UserGroup
	sql := fmt.Sprintf(`
		SELECT *
		FROM %s
		WHERE group_id IN (
			SELECT group_id
			FROM %s
			WHERE uid = ?
		)
	`, TN.UserGroup, TN.UserGroupRef)
	err := g.MySQL.Raw(sql, uid).Scan(&row).Error
	return row, cerror.WrapMysqlErr(err)
}

func GetAPIByRoleId(apiPath string, roleId []int16) (*model.Api, error) {
	row := model.Api{}
	sql := fmt.Sprintf(`
		SELECT *
		FROM %s
		WHERE identity = ?
		  AND api_id IN (
			SELECT api_id
			FROM %s
			WHERE mf_id IN (
				SELECT mf_id
				FROM %s
				WHERE role_id in ?
			)
		)
	`, TN.Api, TN.MfApiRef, TN.RoleMfRef)
	err := g.MySQL.Raw(sql, apiPath, roleId).Scan(&row).Error
	return &row, cerror.WrapMysqlErr(err)
}
