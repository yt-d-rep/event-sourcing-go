package server

import (
	"github.com/gin-gonic/gin"

	usercommand "yt-d-rep/github.com/event-sourcing-go/command/user_command"
)

func Serve() {
	router := gin.Default()
	route(router)
	router.Run(":8080")
}

func route(router *gin.Engine) {
	router.POST("/user", usercommand.Wire().Create)
}
