package main

import (
	"github.com/gin-gonic/gin"
	"go-admin/conf/settings"
)
func main() {
	settings.Setup()

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message":"pong",
		})
	})
	r.Run()
}
