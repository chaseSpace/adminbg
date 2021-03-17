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
	v1OnlyAuth := engine.Group("/web/v1").Use(mw.IfAuthenticated)
	v1OnlyAuth.POST("/SignOut", handler.SignOut) // same as logout

	/*
		v1AuthAPI sub-router hold handle-funcs(APIs) that can only be requested by authenticated users,
		and the users who can call these APIs.
	*/
	v1AuthAPI := engine.Group("/web/v1").
		Use(mw.IfAuthenticated).
		Use(mw.IfCanCallThisAPI(mw.SimpleRole_Comm))

	{
		v1AuthAPI.POST("/NewUser", handler.NewUser)
		v1AuthAPI.POST("/UpdateUser", handler.UpdateUser)
		v1AuthAPI.GET("/QueryUser", handler.QueryUser)
		v1AuthAPI.GET("/GetUserList", handler.GetUserList)

		v1AuthAPI.POST("/NewUserGroup", handler.NewUserGroup)

		v1AuthAPI.POST("/NewMenu", handler.NewMenu)
		v1AuthAPI.GET("/GetMenuList", handler.GetMenuList)
		v1AuthAPI.POST("/UpdateMenu", handler.UpdateMenu)
		v1AuthAPI.DELETE("/DeleteMenus", handler.DeleteMenus)
		v1AuthAPI.POST("/NewFunction", handler.NewFunction)
		v1AuthAPI.POST("/UpdateFunction", handler.UpdateFunction)

		v1AuthAPI.GET("/GetAPIList", handler.GetAPIList)
		v1AuthAPI.POST("/UpdateFuncAndAPIBindInfo", handler.UpdateFuncAndAPIBindInfo)
		v1AuthAPI.POST("/NewAPI", handler.NewAPI)
		v1AuthAPI.POST("/UpdateAPI", handler.UpdateAPI)
	}

	// v1AuthOnlySuperAdmin sub-router hold handle-funcs(APIs) that can only be requested by Administrator(the only one).
	v1AuthOnlySuperAdmin := engine.Group("/web/v1").
		Use(mw.IfAuthenticated).
		Use(mw.IfCanCallThisAPI(mw.SimpleRole_SuperAdmin))
	{
		_ = v1AuthOnlySuperAdmin
	}
}
