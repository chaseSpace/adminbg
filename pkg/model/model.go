package model

import (
	"adminbg/util"
	"gorm.io/gorm"
)

/*
  	It's not recommend to depend struct's gorm tag, we defined all tables struct at `/pkg/model/DDL.sql`,
  if you need to modify them, just directly modify `DDL.sql`.
*/

type User struct {
	Uid          int32
	EncryptedPwd string
	UserName     string
	GroupId      int16
}

func (*User) TableName() string {
	return "adminbg_user"
}

type UserGroup struct {
	GroupId   int16
	GroupName string
	RoleId    int16
}

func (*UserGroup) TableName() string {
	return "adminbg_user_group"
}

type Role struct {
	RoleId   int16
	RoleName string
}

func (*Role) TableName() string {
	return "adminbg_role"
}

func MustInit(db *gorm.DB) {
	err := db.AutoMigrate(new(User), new(UserGroup), new(Role))
	util.PanicIfErr(err, nil)
}
