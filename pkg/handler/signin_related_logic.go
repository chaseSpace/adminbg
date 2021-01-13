package handler

import (
	"adminbg/cerror"
	"adminbg/cproto"
	"adminbg/pkg/crud"
	"adminbg/util"
	"encoding/base64"
	"github.com/pkg/errors"
)

func SignInLogic(req *cproto.SignInReq) (*cproto.SignInRsp, error) {
	plainPwd, err := base64.StdEncoding.DecodeString(req.Pwd)
	if err != nil {
		return nil, errors.Wrap(cerror.ErrParams, "invalid params") // don't return clear err info
	}
	UserEntity, err := crud.CheckUserPassport(req.UserName, string(plainPwd))
	if err != nil {
		return nil, err
	}
	if UserEntity.Uid == 0 {
		return nil, cerror.ErrIncorrectInfoProvided
	}

	claims := util.AdminBgClaims
	claims.Set(util.JwtKey_UID, UserEntity.Uid)
	token, err := util.BgJWT.GenToken(claims)
	if err != nil {
		return nil, err
	}
	rsp := &cproto.SignInRsp{
		Token: string(token),
	}
	return rsp, nil
}

func SignOutLogic(req *cproto.SignOutReq) (*cproto.SignOutRsp, error) {
	// Doing some sign-out operations, but note:
	// In backend, we can't make jwt token expired manually, except that we store it in the cache.
	// With that, we lose the advantage of jwt.
	// However, we can clear the cache within browser.
	return &cproto.SignOutRsp{}, nil
}
