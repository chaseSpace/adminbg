package crud

import (
	"adminbg/cerror"
	"adminbg/pkg/g"
	"adminbg/pkg/model"
	"gorm.io/gorm"
)

func QueryRole(roleId int16) (*model.Role, error) {
	row := new(model.Role)
	err := g.MySQL.First(row, "role_id=?", roleId).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return row, cerror.WrapMysqlErr(err)
}
