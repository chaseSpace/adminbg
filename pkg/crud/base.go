package crud

import "adminbg/pkg/model"

/*
 CRUD: create, retrieve, update, delete operations for DB
*/

// TN is all tables' name holder, just for convenience
var TN = struct {
	User      string
	UserGroup string
	Role      string
}{
	User:      new(model.User).TableName(),
	UserGroup: new(model.UserGroup).TableName(),
	Role:      new(model.Role).TableName(),
}
