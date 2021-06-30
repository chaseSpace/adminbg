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
		v1AuthAPI sub-router hold handle-functions(APIs) that can only be requested by authenticated users,
		and the users who can call these APIs.
	*/
	v1AuthAPI := engine.Group("/web/v1").
		Use(mw.IfAuthenticated).
		Use(mw.IfCanCallThisAPI(mw.SimpleRole_Comm))
	{
		_ = v1AuthAPI
	}

	// v1AuthOnlySuperAdmin sub-router hold handle-functions(APIs) that can only be requested by super admin user(only one).
	v1AuthOnlySuperAdmin := engine.Group("/web/v1").
		Use(mw.IfAuthenticated).
		Use(mw.IfCanCallThisAPI(mw.SimpleRole_SuperAdmin))
	{
		v1AuthOnlySuperAdmin.POST("/NewUser", handler.NewUser)
		v1AuthOnlySuperAdmin.POST("/UpdateUser", handler.UpdateUser)
		v1AuthOnlySuperAdmin.GET("/QueryUser", handler.QueryUser)
		v1AuthOnlySuperAdmin.GET("/GetUserList", handler.GetUserList)

		v1AuthOnlySuperAdmin.POST("/NewUserGroup", handler.NewUserGroup)
		v1AuthOnlySuperAdmin.POST("/UpdateUserGroup", handler.UpdateUserGroup)
		v1AuthOnlySuperAdmin.GET("/QueryUserGroup", handler.QueryUserGroup)
		v1AuthOnlySuperAdmin.GET("/GetUserGroupList", handler.GetUserGroupList)

		v1AuthOnlySuperAdmin.POST("/NewRole", handler.NewRole)
		v1AuthOnlySuperAdmin.POST("/UpdateRole", handler.UpdateRole)
		v1AuthOnlySuperAdmin.GET("/QueryRole", handler.QueryRole)
		v1AuthOnlySuperAdmin.GET("/GetRoleList", handler.GetRoleList)

		v1AuthOnlySuperAdmin.POST("/NewMenu", handler.NewMenu)
		v1AuthOnlySuperAdmin.GET("/GetMenuList", handler.GetMenuList)
		v1AuthOnlySuperAdmin.POST("/UpdateMenu", handler.UpdateMenu)
		v1AuthOnlySuperAdmin.DELETE("/DeleteMenus", handler.DeleteMenus)
		v1AuthOnlySuperAdmin.POST("/NewFunction", handler.NewFunction)
		v1AuthOnlySuperAdmin.POST("/UpdateFunction", handler.UpdateFunction)

		v1AuthOnlySuperAdmin.GET("/GetAPIList", handler.GetAPIList)
		v1AuthOnlySuperAdmin.POST("/UpdateFuncAndAPIBindInfo", handler.UpdateFuncAndAPIBindInfo)
		v1AuthOnlySuperAdmin.POST("/NewAPI", handler.NewAPI)
		v1AuthOnlySuperAdmin.POST("/UpdateAPI", handler.UpdateAPI)
	}

}
