package handler

import (
	"adminbg/cproto"
	"adminbg/pkg/common"
	"github.com/gin-gonic/gin"
)

func NewUserGroup(c *gin.Context) {
	var req cproto.NewUserGroupReq
	common.MustExtractReqParams(c, &req)

	rsp, err := NewUserGroupLogic(&req)
	common.SetRsp(c, err, rsp)
}

func UpdateUserGroup(c *gin.Context) {
	var req cproto.UpdateUserGroupReq
	common.MustExtractReqParams(c, &req)

	rsp, err := UpdateUserGroupLogic(&req)
	common.SetRsp(c, err, rsp)
}

func QueryUserGroup(c *gin.Context) {
	var req cproto.QueryUserGroupReq
	common.MustExtractReqParams(c, &req)

	rsp, err := QueryUserGroupLogic(&req)
	common.SetRsp(c, err, rsp)
}
