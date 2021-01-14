package router

import (
	"adminbg/pkg/handler"
	"adminbg/pkg/mw"
	"github.com/gin-gonic/gin"
)

func Init(engine *gin.Engine) {

	webv1 := engine.Group("/web/v1")
	webv1.POST("/sign_in", handler.SignIn) // same as login

	/*
		webv1Auth sub-router holds APIs that could only be requested by authenticated users.
	*/
	webv1Auth := webv1.Use(mw.AssertAuthenticated)
	webv1Auth.POST("/sign_out", handler.SignOut) // same as logout
}
