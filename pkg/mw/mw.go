package mw

import (
	"adminbg/cerror"
	"adminbg/log"
	"adminbg/pkg/common"
	"adminbg/util"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"gopkg.in/jose.v1/jws"
)

// AssertAuthenticated assets if authenticated, then set user-binding info to context if true.
func AssertAuthenticated(c *gin.Context) {
	JWT, err := jws.ParseJWTFromRequest(c.Request)
	if JWT == nil {
		log.Infof("[gin-middleware: AssertAuthenticated] -- jws.ParseJWTFromRequest err:%v", err)
		common.SetRsp(c, errors.Wrap(cerror.ErrUnauthorized, "empty token"))
		return
	}
	claims, err := util.BgJWT.Verify(JWT)
	if err != nil {
		log.Debugf("[gin-middleware: AssertAuthenticated] -- auth failed, err:%v", err)
		common.SetRsp(c, errors.Wrap(cerror.ErrUnauthorized, err.Error()))
		return
	}
	uid := claims.Get(util.JwtKey_UID)
	// UID was converted to float64 from int32 by JWT.
	if uid, _ := uid.(float64); uid == 0 {
		common.SetRsp(c, errors.Wrap(cerror.ErrUnauthorized, "invalid UID from parsed token"))
		return
	}
	c.Set(common.GinCtxKey_UID, int32(uid.(float64)))

	log.Debugf("[gin-middleware: AssertAuthenticated] -- auth passed, uid:%.f", uid)
}

func Recovery(c *gin.Context) {
	defer func() {
		recovered := recover()
		if recovered == nil {
			return
		}
		err, ok := recovered.(error)
		if !ok {
			PANIC := fmt.Sprintf("unknown panic -> %v", recovered)
			log.Errorf("[gin-middleware: Recovery] %s", PANIC)
			common.SetRsp(c, errors.New(PANIC))
			return
		}
		log.Warnf("[gin-middleware: Recovery] recovered err:%v", err)
		common.SetRsp(c, err)
	}()

	c.Next()
}
