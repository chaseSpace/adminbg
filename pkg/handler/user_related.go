package handler

import (
	"adminbg/cproto"
	"adminbg/pkg/common"
	"github.com/gin-gonic/gin"
)

func SignIn(c *gin.Context) {
	var req cproto.SignInReq
	common.MustExtractReqParams(c, &req)

	rsp, err := SignInLogic(&req)
	common.SetRsp(c, err, rsp)
}

func SignOut(c *gin.Context) {
	var req cproto.SignOutReq
	common.MustExtractReqParams(c, &req)

	rsp, err := SignOutLogic(&req)
	common.SetRsp(c, err, rsp)
}

func NewUser(c *gin.Context) {
	var req cproto.NewUserReq
	common.MustExtractReqParams(c, &req)

	rsp, err := NewUserLogic(&req)
	common.SetRsp(c, err, rsp)
}

func UpdateUser(c *gin.Context) {
	var req cproto.UpdateUserReq
	common.MustExtractReqParams(c, &req)

	rsp, err := UpdateUserLogic(&req)
	common.SetRsp(c, err, rsp)
}

//func NewUser(c *gin.Context) {
//	var req cproto.NewUserReq
//	common.MustExtractReqParams(c, &req)
//
//	rsp, err := NewUserLogic(&req)
//	common.SetRsp(c, err, rsp)
//}
