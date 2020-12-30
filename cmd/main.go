package main

import (
	"fmt"
	"go-admin/api/routers"
	"syscall"

	//"github.com/gin-gonic/gin"
	//"go-admin/api/middlewares/log"
	"go-admin/api/models"
	"go-admin/conf/settings"
	l "log"
	"github.com/fvbock/endless"
)

var logger *l.Logger
func main() {
	settings.Setup()
	models.SetUp()

	//r := gin.Default()
	//r.Use(log.Logger())
	//r.GET("/ping", func(c *gin.Context) {
	//	logger.Println("AAAAAAAAAAAA")
	//	c.JSON(200, gin.H{
	//		"message":"pong",
	//	})
	//})
	//r.Run()

	endless.DefaultReadTimeOut = settings.ServerSetting.ReadTimeout
	endless.DefaultWriteTimeOut = settings.ServerSetting.WriteTimeout
	endless.DefaultMaxHeaderBytes = 1 << 20
	endPoint := fmt.Sprintf(":%d", settings.ServerSetting.HttpPort)

	server := endless.NewServer(endPoint, routers.InitRouter())

	server.BeforeBegin = func(add string) {
		logger.Printf("Actual pid is %d", syscall.Getpid())

	}

	err := server.ListenAndServe()
	if err != nil {
		logger.Printf("Server err: %v", err)
	}
}
