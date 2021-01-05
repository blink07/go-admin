package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-admin/api/routers"
	"net/http"
	//"github.com/gin-gonic/gin"
	//"go-admin/api/middlewares/log"
	"go-admin/api/models"
	"go-admin/conf/settings"
)

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
	// windows下endless不兼容
	//endless.DefaultReadTimeOut = settings.ServerSetting.ReadTimeout
	//endless.DefaultWriteTimeOut = settings.ServerSetting.WriteTimeout
	//endless.DefaultMaxHeaderBytes = 1 << 20
	//endPoint := fmt.Sprintf(":%d", settings.ServerSetting.HttpPort)
	//
	//server := endless.NewServer(endPoint, routers.InitRouter())
	//
	//server.BeforeBegin = func(add string) {
	//	logger.Printf("Actual pid is %d", syscall.Getpid())
	//
	//}
	//err := server.ListenAndServe()
	//if err != nil {
	//	logger.Printf("Server err: %v", err)
	//}
	gin.SetMode(settings.ServerSetting.RunMode)
	routersInit := routers.InitRouter()
	readTimeout := settings.ServerSetting.ReadTimeout
	writeTimeout := settings.ServerSetting.WriteTimeout
	endPort := fmt.Sprintf(":%d", settings.ServerSetting.HttpPort)
	server := &http.Server{
		Addr: endPort,
		Handler: routersInit,
		ReadTimeout: readTimeout,
		WriteTimeout: writeTimeout,
	}
	server.ListenAndServe()
}
