package router

import (
	"adminbg/pkg/handler"
	"adminbg/pkg/mw"
	"github.com/gin-gonic/gin"
)

func Init(engine *gin.Engine) {

	webV1 := engine.Group("/web/v1")
	webV1.POST("/sign_in", handler.SignIn) // same as login

	/*
		webV1AuthPass sub-router holds APIs that could only be requested by authenticated users.
	*/
	webV1AuthPass := webV1.Use(mw.AssertAuthenticated)
	webV1AuthPass.POST("/sign_out", handler.SignOut) // same as logout
}
