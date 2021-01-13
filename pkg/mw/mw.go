package mw

import (
	"adminbg/cerror"
	"adminbg/log"
	"adminbg/pkg/common"
	"adminbg/util"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

// AssertAuthenticated assets if authenticated, if true, then set important info to context.
func AssertAuthenticated(c *gin.Context) {
	token := c.GetHeader("token")
	if token == "" {
		common.SetRsp(c, cerror.ErrUnauthorized, common.HttpRsp{Tip: "empty token"})
		return
	}
	claims, err := util.BgJWT.Verify([]byte(token))
	if err != nil {
		log.Debugf("[gin-middleware: AssertAuthenticated] -- auth failed, err:%v", err)
		common.SetRsp(c, cerror.ErrUnauthorized, common.HttpRsp{Tip: err.Error()})
		return
	}
	uid := claims.Get(util.JwtKey_UID)
	c.Set(common.GinCtxKey_UID, uid)

	log.Debugf("[gin-middleware: AssertAuthenticated] -- auth passed, uid:%d", uid)
	//c.Next()
}

func RecoverIfNeed(c *gin.Context) {
	defer func() {
		recovered := recover()
		if recovered == nil {
			return
		}
		err, ok := recovered.(error)
		if !ok {
			PANIC := fmt.Sprintf("unknown panic -> %v", recovered)
			log.Errorf("[gin-middleware: RecoverIfNeed] %s\n", PANIC)
			common.SetRsp(c, errors.New(PANIC))
			return
		}
		log.Warnf("[gin-middleware: RecoverIfNeed] recovered err:%v", err)
		common.SetRsp(c, err)
	}()

	c.Next()
}
