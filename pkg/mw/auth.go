package mw

import (
	"adminbg/cerror"
	"adminbg/log"
	"adminbg/pkg/common"
	"adminbg/pkg/crud"
	"adminbg/util"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"gopkg.in/jose.v1/jws"
)

// IfAuthenticated assets if authenticated, then set user-binding info to context if true.
func IfAuthenticated(c *gin.Context) {
	JWT, err := jws.ParseJWTFromRequest(c.Request)
	if err != nil {
		log.Infof("[%s] -- jws.ParseJWTFromRequest err:%v", util.GetRunningFuncName("/"), err)
		common.SetRsp(c, errors.Wrap(cerror.ErrUnauthorized, "empty token"))
		return
	}
	claims, err := util.BgJWT.Verify(JWT)
	if err != nil {
		log.Debugf("[%s] -- auth failed, err:%v", util.GetRunningFuncName("/"), err)
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

	log.Debugf("[%s] -- auth passed, uid:%.f", util.GetRunningFuncName("/"), uid)
}

type SimpleRole uint8

const (
	SimpleRole_Comm       SimpleRole = iota
	SimpleRole_SuperAdmin SimpleRole = 1
)

/*
TODO:
	This mw func will execute two db operations, read five tables in total,
	may be an optimizable point.
*/
func IfCanCallThisAPI(sr SimpleRole) gin.HandlerFunc {
	return func(c *gin.Context) {
		v, _ := c.Get(common.GinCtxKey_UID)
		uid, ok := v.(int32)
		if !ok {
			common.SetRsp(c, errors.Wrap(cerror.ErrUnauthorized, "mw"))
			return
		}
		// Super admin user skips check below.
		if yes, _ := common.IsSuperAdmin(uid); yes {
			return
		}
		path := c.Request.URL.Path // e.g. /web/v1/NewUser

		if sr == SimpleRole_SuperAdmin {
			log.Warnf("[%s] [!!!request rejected] uid:%d try to request API:[%s] only super admin is permitted",
				util.GetRunningFuncName("/"), uid, path)
			common.SetRsp(c, errors.Wrap(cerror.ErrNotAllowed, "mw"))
			return
		}

		groups, err := crud.GetUserGroup(uid)
		if err != nil {
			common.SetRsp(c, errors.Wrap(err, "mw"))
			return
		}
		if len(groups) == 0 {
			log.Warnf("[%s] uid:%d is not bound to any group(Abnormal data)", util.GetRunningFuncName("/"), uid)
			common.SetRsp(c, errors.Wrap(cerror.ErrNotAllowed, "mw"))
			return
		}
		var roleIDs []int16
		for _, gp := range groups {
			roleIDs = append(roleIDs, gp.RoleId)
		}
		api, err := crud.GetAPIByRoleIds(path, roleIDs)
		if err != nil {
			common.SetRsp(c, errors.Wrap(err, "mw"))
			return
		}
		if api.ApiId < 1 {
			common.SetRsp(c, errors.Wrap(cerror.ErrNotAllowed, "mw"))
			return
		}
	}
}
