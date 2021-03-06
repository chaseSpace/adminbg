package handler

import (
	"adminbg/cproto"
	"adminbg/pkg/common"
	"github.com/gin-gonic/gin"
)

func NewRole(c *gin.Context) {
	var req cproto.NewRoleReq
	common.MustExtractReqParams(c, &req)

	rsp, err := NewRoleLogic(&req)
	common.SetRsp(c, err, rsp)
}

func UpdateRole(c *gin.Context) {
	var req cproto.UpdateRoleReq
	common.MustExtractReqParams(c, &req)

	rsp, err := UpdateRoleLogic(&req)
	common.SetRsp(c, err, rsp)
}

func QueryRole(c *gin.Context) {
	var req cproto.QueryRoleReq
	common.MustExtractReqParams(c, &req)

	rsp, err := QueryRoleLogic(&req)
	common.SetRsp(c, err, rsp)
}

func GetRoleList(c *gin.Context) {
	var req cproto.GetRoleListReq
	common.MustExtractReqParams(c, &req)

	rsp, err := GetRoleListLogic(&req)
	common.SetRsp(c, err, rsp)
}
