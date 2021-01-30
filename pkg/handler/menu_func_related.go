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
