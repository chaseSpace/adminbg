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

type User struct {
	Name      string        `json:"name" binding:"required"`
	Uid       int32         `json:"uid"`        // only be used at query action
	AccountId string        `json:"account_id"` // if is update action, `account_id` is unused field
	Pwd       string        `json:"pwd"`        // if is query action, `pwd` is unused field
	Phone     string        `json:"phone"`
	Email     string        `json:"email"`
	Sex       SexTyp        `json:"sex" binding:"required"`    // string type, search `SexTyp` at this file
	Status    UserStatusTyp `json:"status" binding:"required"` // string type, search `UserStatusTyp` at this file
	RoleId    int16         `json:"role_id"`
	GroupId   int16         `json:"group_id"`
	Remark    string        `json:"remark"`
	CreatedAt string        `json:"created_at"` // YYYY-MM-dd HH:mm:SS; only be used at query action
	UpdatedAt string        `json:"updated_at"` // YYYY-MM-dd HH:mm:SS; only be used at query action
}

// POST /web/v1/NewUser
type NewUserReq struct {
	User
}

type NewUserRsp struct {
}

// Define sex type just as a mysql field's enum copy, we should
// take mysql table schema as the root(same with below).
type SexTyp string

const (
	Man        SexTyp = "MAN"
	Woman      SexTyp = "WOMAN"
	SexUnknown SexTyp = "UNKNOWN"
)

func (sex SexTyp) IsValid() bool {
	switch sex {
	case Man, Woman, SexUnknown:
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
type UpdateUserReq struct {
	Delete bool  `json:"delete"` // if true, server will delete user for this uid, other params would be ignored.
	Uid    int32 `json:"uid" binding:"required"`
	User
}

type UpdateUserRsp struct{}

// GET /web/v1/QueryUser
type QueryUserReq struct {
	Uid int32 `form:"uid" binding:"required"`
}

type QueryUserRsp struct {
	*User // *XXX means this field is pointer type, it might be null.
}

// GET /web/v1/GetUserList
type GetUserListReq struct {
	PageNum  uint16 `form:"pn"`
	PageSize uint16 `form:"ps"`
}

type GetUserListRsp struct {
	List  []*User `json:"list"` // Order by CreatedAt by default.
	Total int64   `json:"total"`
}
