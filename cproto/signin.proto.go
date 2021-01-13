package cproto

// POST /web/v1/SignIn
type SignInReq struct {
	UserName string `json:"user_name" binding:"required"`
	Pwd      string `json:"pwd" binding:"required"` // base64 encoded
}

type SignInRsp struct {
	Token string `json:"token"`
}

// POST /web/v1/SignOut
type SignOutReq struct{}

type SignOutRsp struct{}
