package model

import (
	"adminbg/cerror"
	"adminbg/cproto"
	"adminbg/util"
	"fmt"
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
	Uid          int32 `gorm:"primary_key"` // Set gorm tag could write back this field when insert
	AccountId    string
	EncryptedPwd string
	Salt         string
	NickName     string
	Phone        string
	Email        string
	Sex          cproto.SexTyp
	Remark       string
	Status       cproto.UserStatusTyp
	RoleId       int16 `gorm:"->"` // read-only field
	GroupId      int16 `gorm:"->"` // read-only field
}

func (u *User) Proto() *cproto.User {
	return &cproto.User{
		Name:      u.NickName,
		AccountId: u.AccountId,
		Uid:       u.Uid,
		//Pwd:       ,
		Phone:     u.Phone,
		Email:     u.Email,
		Sex:       u.Sex,
		Status:    u.Status,
		RoleId:    u.RoleId,
		GroupId:   u.GroupId,
		Remark:    u.Remark,
		CreatedAt: u.CreatedAt.Format(util.TimeLayout),
		UpdatedAt: u.UpdatedAt.Format(util.TimeLayout),
	}
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
	Id        int32
	Uid       int32
	GroupId   int32
	CreatedAt time.Time
	UpdatedAt time.Time
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
	Id        int32
	RoleId    int32
	MfId      int32
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (*RoleMfRef) TableName() string {
	return TablePrefix + "role_mf_ref"
}

type MenuAndFunction struct {
	BaseModel
	MfId        int32 `gorm:"primary_key"` // Set gorm tag could write back this field when insert
	MfName      string
	Path        string
	ParentId    int32
	Level       int8
	Type        cproto.MfType
	MenuRoute   string
	MenuDisplay cproto.MenuDisplay
	SortNum     uint16
}

func (*MenuAndFunction) TableName() string {
	return TablePrefix + "menu_and_function"
}
func (r *MenuAndFunction) Check() error {
	switch r.Type {
	case cproto.Menu:
		if !(r.Level > 0 && r.Level <= MaxMenuLevel) {
			return errors.New("invalid menu level")
		}
		if r.Level != 1 && r.ParentId == MenuRootId {
			return errors.New("parent_id could be 100 only when level is 1")
		}
		routeUTF8Len := len([]rune(r.MenuRoute))
		min, max := 1, 100
		if !(routeUTF8Len >= 1 && routeUTF8Len <= 100) {
			return fmt.Errorf("invalid menu route length, valid range is [%d,%d]", min, max)
		}
	case cproto.Function:
		if r.Level != 0 {
			return errors.New("function's level must be 0")
		}
		if r.ParentId == MenuRootId {
			return errors.New("function's parent_id can't be 100(this is root menu ID)")
		}
	}
	min, max := 1, 50
	nameUTF8Len := len([]rune(r.MfName))
	if !(nameUTF8Len > 0 && nameUTF8Len <= 50) {
		return fmt.Errorf("invalid name length, valid range is [%d,%d]", min, max)
	}
	if err := r.Type.Check(); err != nil {
		return err
	}
	if err := r.MenuDisplay.Check(); err != nil {
		return err
	}
	return nil
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
	Remark   string
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
