package handler

import (
	"adminbg/cerror"
	"adminbg/cproto"
	"adminbg/pkg/crud"
	"adminbg/pkg/model"
	"adminbg/util"
	"encoding/base64"
	"fmt"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

func SignInLogic(req *cproto.SignInReq) (*cproto.SignInRsp, error) {
	plainPwd, err := base64.StdEncoding.DecodeString(req.Pwd)
	if err != nil {
		return nil, errors.Wrap(cerror.ErrParams, "invalid params") // No way to return clear err info
	}
	UserEntity, err := crud.CheckUserPassport(req.AccountId, string(plainPwd))
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

func NewUserLogic(req *cproto.NewUserReq) (*cproto.NewUserRsp, error) {
	err := (&model.UserBase{
		Uid:    999,
		Sex:    req.Sex,
		Status: req.Status,
	}).Check()
	if err != nil {
		return nil, err
	}
	if req.Pwd == "" {
		return nil, errors.Wrap(cerror.ErrParams, "password is required")
	}
	plainPwd, err := base64.StdEncoding.DecodeString(req.Pwd)
	if err != nil {
		return nil, errors.Wrap(cerror.ErrParams, "invalid params `pwd`, decode err") // No way to return clear err info
	}
	if req.AccountId == "" {
		return nil, errors.Wrap(cerror.ErrParams, "`account_id` is required")
	}
	userBase := &model.UserBase{
		AccountId:    req.AccountId,
		EncryptedPwd: string(plainPwd),
		Salt:         uuid.New().String()[24:],
		NickName:     req.Name,
		Phone:        req.Phone,
		Email:        req.Email,
		Sex:          req.Sex,
		Remark:       req.Remark,
		Status:       req.Status,
	}
	err = crud.InsertUser(userBase)
	return new(cproto.NewUserRsp), err
}

func UpdateUserLogic(req *cproto.UpdateUserReq) (*cproto.UpdateUserRsp, error) {
	rsp := new(cproto.UpdateUserRsp)

	if req.Delete {
		foundAndDeleted, err := crud.DeleteUser(crud.UserIdentity{Uid: req.Uid})
		if err != nil {
			return nil, err
		}
		if !foundAndDeleted {
			return nil, errors.Wrap(cerror.ErrParams, fmt.Sprintf("uid:%d not found", req.Uid))
		}
		return rsp, nil
	}

	// update
	if req.Pwd != "" {
		byteS, err := base64.StdEncoding.DecodeString(req.Pwd)
		if err != nil {
			return nil, errors.Wrap(cerror.ErrParams, "invalid pwd")
		}
		req.Pwd = string(byteS) // change to plain text
	}
	err := (&model.UserBase{
		Uid:    999, // feel free, the key is below fields.
		Sex:    req.Sex,
		Status: req.Status,
	}).Check()
	if err != nil {
		return nil, err
	}
	err = crud.UpdateUser(crud.UserIdentity{Uid: req.Uid, ContainsDeleted: true}, req)
	return rsp, err
}

func QueryUserLogic(req *cproto.QueryUserReq) (*cproto.QueryUserRsp, error) {
	userBase, err := crud.GetUserBase(req.Uid)
	if err != nil {
		return nil, err
	}
	rsp := &cproto.QueryUserRsp{User: userBase}
	return rsp, nil
}

func GetUserListLogic(req *cproto.GetUserListReq) (*cproto.GetUserListRsp, error) {
	if err := util.CheckSplitPageParams(req.PageNum, req.PageSize); err != nil {
		return nil, err
	}
	list, total, err := crud.GetUserList(req.PageNum, req.PageSize, crud.CreatedAtAsc, crud.UpdatedAtAsc)
	if err != nil {
		return nil, err
	}
	rsp := &cproto.GetUserListRsp{Total: total}
	for _, item := range list {
		rsp.List = append(rsp.List, item.Proto())
	}
	return rsp, nil
}
