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
