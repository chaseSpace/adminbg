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
