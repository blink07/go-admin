package main

import (
	"github.com/gin-gonic/gin"
	"go-admin/api/models"
	"go-admin/conf/settings"
)
func main() {
	settings.Setup()
	models.SetUp()

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message":"pong",
		})
	})
	r.Run()
}
