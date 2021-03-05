package handler

import (
	"adminbg/cproto"
	"adminbg/pkg/common"
	"github.com/gin-gonic/gin"
)

func NewAPI(c *gin.Context) {
	var req cproto.NewAPIReq
	common.MustExtractReqParams(c, &req)

	rsp, err := NewAPILogic(&req)
	common.SetRsp(c, err, rsp)
}

func UpdateAPI(c *gin.Context) {
	var req cproto.UpdateAPIReq
	common.MustExtractReqParams(c, &req)

	rsp, err := UpdateAPILogic(&req)
	common.SetRsp(c, err, rsp)
}
