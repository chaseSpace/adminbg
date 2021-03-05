package mw

import (
	"adminbg/log"
	"adminbg/pkg/common"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

func Recovery(c *gin.Context) {
	defer func() {
		r := recover()
		if r == nil {
			return
		}
		err, ok := r.(error)
		if !ok {
			PANIC := fmt.Sprintf("unknown panic -> %v", r)
			log.Errorf("[gin-middleware: Recovery] %s", PANIC)
			common.SetRsp(c, errors.New(PANIC))
			return
		}
		log.Warnf("[gin-middleware: Recovery] recovered err:%v", err)
		common.SetRsp(c, err)
	}()

	c.Next()
}
