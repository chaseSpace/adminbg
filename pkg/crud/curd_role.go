package crud

import (
	"adminbg/cerror"
	"adminbg/cproto"
	"adminbg/pkg/g"
	"adminbg/pkg/model"
	"adminbg/util/_gorm"
	"github.com/pkg/errors"
)

func NewRole(params *cproto.Role) error {
	ent := &model.Role{
		RoleName: params.RoleName,
	}
	err := g.MySQL.Create(ent).Error
	return err
}

func DeleteRole(roleId int16) error {
	exec := g.MySQL.Take(&model.UserGroup{}, "role_id=?", roleId)
	if _gorm.IsDBErr(exec.Error) {
		return exec.Error
	}
	if exec.RowsAffected > 0 {
		return errors.Errorf("this role is being used, can't be delete")
	}
	exec = g.MySQL.Delete(new(model.Role), "role_id=?", roleId)
	return cerror.WrapDeleteDBOneRecordErr(exec)
}

func UpdateRole(params *cproto.Role) error {
	ent := &model.Role{
		RoleId:   params.RoleId,
		RoleName: params.RoleName,
	}
	exec := g.MySQL.Where("role_id=?", params.RoleId).Select([]string{"role_name"}).Updates(ent)
	return cerror.WrapUpdateDBOneRecordErr(exec)
}
