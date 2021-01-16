package cproto

// POST /web/v1/SignIn
type SignInReq struct {
	AccountId string `json:"account_id" binding:"required"`
	Pwd       string `json:"pwd" binding:"required"` // base64 encoded
}

type SignInRsp struct {
	Token string `json:"token"`
}

// POST /web/v1/SignOut
type SignOutReq struct{}

type SignOutRsp struct{}

// POST /web/v1/NewUser
type NewUserReq struct {
	Name      string        `json:"name" binding:"required"`
	AccountId string        `json:"account_id" binding:"required"`
	Pwd       string        `json:"pwd" binding:"required"`
	Phone     string        `json:"phone"`
	Email     string        `json:"email"`
	Sex       SexTyp        `json:"sex" binding:"required"`    // string type, search `SexTyp` at this file
	Status    UserStatusTyp `json:"status" binding:"required"` // string type, search `UserStatusTyp` at this file
	RoleId    int16         `json:"role_id"`
	GroupId   int16         `json:"group_id"`
	Remark    string        `json:"remark"`
}

type NewUserRsp struct {
}

// Define sex type just as a mysql field's enum copy, we should
// take mysql table schema as the root(same with below).
type SexTyp string

const (
	Man     SexTyp = "MAN"
	Woman   SexTyp = "WOMAN"
	Unknown SexTyp = "UNKNOWN"
)

func (sex SexTyp) IsValid() bool {
	switch sex {
	case Man, Woman, Unknown:
		return true
	}
	return false
}

type UserStatusTyp string

const (
	Normal  UserStatusTyp = "NORMAL"
	Suspend UserStatusTyp = "SUSPEND"
)

func (sta UserStatusTyp) IsValid() bool {
	switch sta {
	case Normal, Suspend:
		return true
	}
	return false
}

// POST /web/v1/ModifyUser
type ModifyUserReq struct {
	Delete  bool          `json:"delete"` // if true, server will delete user for this uid, other params would be ignored.
	Uid     int32         `json:"uid" binding:"required"`
	Name    string        `json:"name" binding:"required"`
	Pwd     string        `json:"pwd"` // if empty, it would be ignored.
	Phone   string        `json:"phone"`
	Email   string        `json:"email"`
	Sex     SexTyp        `json:"sex" binding:"required"`    // string type, search `SexTyp` at this file
	Status  UserStatusTyp `json:"status" binding:"required"` // string type, search `UserStatusTyp` at this file
	RoleId  int16         `json:"role_id"`
	GroupId int16         `json:"group_id"`
	Remark  string        `json:"remark"`
}

type ModifyUserRsp struct{}
