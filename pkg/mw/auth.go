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
		log.Infof("[gin-middleware: IfAuthenticated] -- jws.ParseJWTFromRequest err:%v", err)
		common.SetRsp(c, errors.Wrap(cerror.ErrUnauthorized, "empty token"))
		return
	}
	claims, err := util.BgJWT.Verify(JWT)
	if err != nil {
		log.Debugf("[gin-middleware: IfAuthenticated] -- auth failed, err:%v", err)
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

	log.Debugf("[gin-middleware: IfAuthenticated] -- auth passed, uid:%.f", uid)
}

/*
TODO:
	This mw func will execute two db operations, read five tables in total,
	may be an optimizable point.
*/
func IfCanCallThisAPI(c *gin.Context) {
	_uid, _ := c.Get(common.GinCtxKey_UID)
	uid := _uid.(int32)
	// Super-admin user skips check below.
	if yes, _ := common.IsSuperAdmin(uid); yes {
		return
	}
	path := c.Request.URL.Path // e.g. /web/v1/NewUser
	groups, err := crud.GetUserGroup(uid)
	if err != nil {
		common.SetRsp(c, errors.Wrap(err, "mw"))
		return
	}
	if len(groups) == 0 {
		log.Warnf("[gin-middleware: IfCanCallThisAPI] uid:%d is not bound to any group(Abnormal data)", uid)
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
	if api.ApiId == 0 {
		common.SetRsp(c, errors.Wrap(cerror.ErrNotAllowed, "mw"))
		return
	}
}
