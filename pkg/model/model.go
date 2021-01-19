package model

import (
	"adminbg/cerror"
	"adminbg/cproto"
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

type UserGroupRef struct {
	gorm.Model
	Uid     int32
	GroupId int32
}

func (*UserGroupRef) TableName() string {
	return TablePrefix + "user_group_ref"
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

type RoleMfRef struct {
	gorm.Model
	RoleId int32
	MfId   int32
}

func (*RoleMfRef) TableName() string {
	return TablePrefix + "role_mf_ref"
}

type MenuAndFunction struct {
	BaseModel
	MfId     int32
	MfName   string
	Path     string
	ParentId int32
	Level    int8
	Type     cproto.MfType
}

func (*MenuAndFunction) TableName() string {
	return TablePrefix + "menu_and_function"
}

type MfApiRef struct {
	gorm.Model
	MfId  int32
	ApiId int32
}

func (*MfApiRef) TableName() string {
	return TablePrefix + "mf_api_ref"
}

type Api struct {
	BaseModel
	ApiId    int32
	Identity string
	MfId     int32
}

func (*Api) TableName() string {
	return TablePrefix + "api"
}

type OperationLog struct {
	OpId       int32
	Type       string
	OpUid      int32
	OpUsername string
	Remark     string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

func (*OperationLog) TableName() string {
	return TablePrefix + "operation_log"
}

func MustInit(db *gorm.DB) {
	// We should execute `DDL.sql` to migrate tables' schema, not in runtime.
}
