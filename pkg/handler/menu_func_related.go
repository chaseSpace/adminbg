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
