package core

import "github.com/gin-gonic/gin"

var Server *gin.Engine

func init() {
	Server = gin.New()
}

var Tail = "--"

func RunServer() {
	Server.GET("/", func(c *gin.Context) {
		c.String(200, Tail)
	})
	Server.Run(":" + sillyGirl.Get("port", "8080"))
}
