package model

import (
	"adminbg/cerror"
	"adminbg/cproto"
	"adminbg/util"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"time"
)

/*
  	It's not recommend to depend struct's gorm tag, we defined all tables struct at `/pkg/model/DDL.sql`,
  if you need to modify them, just directly modify `DDL.sql`.
*/

type BaseModel struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

type User struct {
	BaseModel
	UserBase
}

func (*User) TableName() string {
	return TablePrefix + "user"
}

type UserBase struct {
	Uid          int32
	AccountId    string
	EncryptedPwd string
	Salt         string
	NickName     string
	Phone        string
	Email        string
	Sex          cproto.SexTyp
	Remark       string
	Status       cproto.UserStatusTyp
	GroupId      int16
}

func (u *UserBase) Check() error {
	if u.Uid == 0 {
		return errors.Wrap(cerror.ErrParams, "invalid uid")
	}
	if !u.Sex.IsValid() {
		return errors.Wrap(cerror.ErrParams, "invalid sex")
	}
	if !u.Status.IsValid() {
		return errors.Wrap(cerror.ErrParams, "invalid status")
	}
	return nil
}

type UserGroup struct {
	BaseModel
	GroupId   int16
	GroupName string
	RoleId    int16
}

func (*UserGroup) TableName() string {
	return TablePrefix + "user_group"
}

type Role struct {
	BaseModel
	RoleId   int16
	RoleName string
}

func (*Role) TableName() string {
	return TablePrefix + "role"
}

func MustInit(db *gorm.DB) {
	err := db.AutoMigrate(new(User), new(UserGroup), new(Role))
	util.PanicIfErr(err, nil)
}
