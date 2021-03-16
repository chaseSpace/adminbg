package crud

import (
	"adminbg/pkg/model"
	"fmt"
	"gorm.io/gorm"
)

/*
 CRUD: create, retrieve, update, delete operations for DB
*/

/*
TN is all tables' name holder, for convenience and beauty.
  For sometimes, we need coding a raw SQL at working, at this time,
we shouldn't directly write table name within code, that's a tight-coupling way.
Here we created a struct variable `TN` to hold all tables' name.
*/
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

type OrderByOption int

const (
	CreatedAtAsc OrderByOption = iota
	CreatedAtDesc
	UpdatedAtAsc
	UpdatedAtDesc
)

var OrderByMap = map[OrderByOption]string{
	CreatedAtAsc:  "created_at asc",
	CreatedAtDesc: "created_at desc",
	UpdatedAtAsc:  "updated_at asc",
	UpdatedAtDesc: "updated_at desc",
}

func GenOrderByClause(option ...OrderByOption) string {
	c := ""
	for _, o := range option {
		if c != "" {
			c += fmt.Sprintf(",%s", OrderByMap[o])
			continue
		}
		c = OrderByMap[o]
	}
	return c
}
