package handler

import (
	"adminbg/cproto"
	"adminbg/pkg/common"
	"github.com/gin-gonic/gin"
)

func NewMenu(c *gin.Context) {
	var req cproto.NewMenuReq
	common.MustExtractReqParams(c, &req)

	rsp, err := NewMenuLogic(&req)
	common.SetRsp(c, err, rsp)
}

func UpdateMenu(c *gin.Context) {
	var req cproto.UpdateMenuReq
	common.MustExtractReqParams(c, &req)

	rsp, err := UpdateMenuLogic(&req)
	common.SetRsp(c, err, rsp)
}

func GetMenuList(c *gin.Context) {
	var req cproto.GetMenuListReq
	common.MustExtractReqParams(c, &req)

	rsp, err := GetMenuListLogic(&req)
	common.SetRsp(c, err, rsp)
}

func DeleteMenus(c *gin.Context) {
	var req cproto.DeleteMenusReq
	common.MustExtractReqParams(c, &req)

	rsp, err := DeleteMenusLogic(&req)
	common.SetRsp(c, err, rsp)
}

func NewFunction(c *gin.Context) {
	var req cproto.NewFunctionReq
	common.MustExtractReqParams(c, &req)

	rsp, err := NewFunctionLogic(&req)
	common.SetRsp(c, err, rsp)
}

func UpdateFunction(c *gin.Context) {
	var req cproto.UpdateFunctionReq
	common.MustExtractReqParams(c, &req)

	rsp, err := UpdateFunctionLogic(&req)
	common.SetRsp(c, err, rsp)
}

func GetAPIList(c *gin.Context) {
	var req cproto.GetAPIListReq
	common.MustExtractReqParams(c, &req)

	rsp, err := GetAPIListLogic(&req)
	common.SetRsp(c, err, rsp)
}

func UpdateFuncAndAPIBindInfo(c *gin.Context) {
	var req cproto.UpdateFuncAndAPIBindInfoReq
	common.MustExtractReqParams(c, &req)

	rsp, err := UpdateFuncAndAPIBindInfoLogic(&req)
	common.SetRsp(c, err, rsp)
}
