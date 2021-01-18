package router

import (
	"adminbg/pkg/handler"
	"adminbg/pkg/mw"
	"github.com/gin-gonic/gin"
)

func Init(engine *gin.Engine) {

	webv1 := engine.Group("/web/v1")
	webv1.POST("/SignIn", handler.SignIn) // same as login

	/*
		webv1Auth sub-router holds APIs that could only be requested by authenticated users.
	*/
	webv1Auth := webv1.Use(mw.AssertAuthenticated, mw.AssertCanCallThisAPI)
	webv1Auth.POST("/SignOut", handler.SignOut) // same as logout
	webv1Auth.POST("/NewUser", handler.NewUser)
	webv1Auth.POST("/ModifyUser", handler.ModifyUser)
}
