package crud

import (
	"adminbg/pkg/model"
	"gorm.io/gorm"
)

/*
 CRUD: create, retrieve, update, delete operations for DB
*/

// TN is all tables' name holder, for convenience
var TN = struct {
	User            string
	UserGroup       string
	UserGroupRef    string
	Role            string
	MenuAndFunction string
	RoleMfRef       string
	Api             string
	MfApiRef        string
}{
	User:            new(model.User).TableName(),
	UserGroup:       new(model.UserGroup).TableName(),
	UserGroupRef:    new(model.UserGroupRef).TableName(),
	Role:            new(model.Role).TableName(),
	MenuAndFunction: new(model.MenuAndFunction).TableName(),
	RoleMfRef:       new(model.RoleMfRef).TableName(),
	Api:             new(model.Api).TableName(),
	MfApiRef:        new(model.MfApiRef).TableName(),
}

type commTyp struct {
	Id int32
}

func LastInsertId(tx *gorm.DB) (int32, error) {
	row := commTyp{}
	err := tx.Raw("SELECT LAST_INSERT_ID() id").Scan(&row).Error
	return row.Id, err
}
