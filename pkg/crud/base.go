package crud

import "adminbg/pkg/model"

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
