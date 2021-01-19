package router

import (
	"adminbg/pkg/handler"
	"adminbg/pkg/mw"
	"github.com/gin-gonic/gin"
)

func Init(engine *gin.Engine) {

	v1Free := engine.Group("/web/v1")
	v1Free.POST("/SignIn", handler.SignIn) // same as login

	/*
		v1OnlyAuth sub-router hold handle-funcs(APIs) that can only be requested by authenticated users.
	*/
	v1OnlyAuth := v1Free.Use(mw.IfAuthenticated)
	v1OnlyAuth.POST("/SignOut", handler.SignOut) // same as logout

	/*
		v1AuthContainsApi sub-router hold handle-funcs(APIs) that can only be requested by authenticated users,
		and the users who can call these APIs.
	*/
	v1AuthContainsApi := v1OnlyAuth.Use(mw.IfCanCallThisAPI)
	v1AuthContainsApi.POST("/NewUser", handler.NewUser)
	v1AuthContainsApi.POST("/ModifyUser", handler.ModifyUser)
}
