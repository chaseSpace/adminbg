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
	v1AuthContainsApi.POST("/UpdateUser", handler.UpdateUser)
	v1AuthContainsApi.GET("/GetUser", handler.GetUser)

	v1AuthContainsApi.POST("/NewMenu", handler.NewMenu)
	v1AuthContainsApi.GET("/GetMenuList", handler.GetMenuList)
	v1AuthContainsApi.POST("/UpdateMenu", handler.UpdateMenu)
	v1AuthContainsApi.DELETE("/DeleteMenus", handler.DeleteMenus)
	v1AuthContainsApi.POST("/NewFunction", handler.NewFunction)
	v1AuthContainsApi.POST("/UpdateFunction", handler.UpdateFunction)
	v1AuthContainsApi.GET("/GetAPIList", handler.GetAPIList)
	v1AuthContainsApi.POST("/UpdateFuncAndAPIBindInfo", handler.UpdateFuncAndAPIBindInfo)

	v1AuthContainsApi.POST("/NewAPI", handler.NewAPI)
	v1AuthContainsApi.POST("/UpdateAPI", handler.UpdateAPI)
}
