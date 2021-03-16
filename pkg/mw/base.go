package mw

import (
	"adminbg/log"
	"adminbg/pkg/common"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"runtime/debug"
)

func Recovery(c *gin.Context) {
	logHEAD := "[gin-mw: Recovery panic]"
	defer func() {
		r := recover()
		if r == nil {
			return
		}
		debug.PrintStack()
		err, ok := r.(error)
		if !ok {
			PANIC := fmt.Sprintf("unknown panic -> %v", r)
			log.Errorf("%s unknown panic -> %v", logHEAD, PANIC)
			common.SetRsp(c, errors.New(PANIC))
			return
		}
		log.Warnf("%s err:%v", logHEAD, err)
		common.SetRsp(c, err)
	}()

	c.Next()
}
