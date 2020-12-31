package routers

import (
	"github.com/gin-gonic/gin"
	"go-admin/api/middlewares/log"
)
//var logru = logrus.New()

func InitRouter() *gin.Engine {
	r:=gin.New()

	r.Use(log.Logger())
	//r.Use(log.WinLoggerHandler())

	apiv1 := r.Group("/api/v1")
	apiv1.GET("/ping", func(context *gin.Context) {
		//println("AAAAAAAAAA")
		//logrus.Debugf("BBBBBBBBBBBBB")
		log.Info("BBBBBBBBBBBBBBB")
		context.JSON(200, gin.H{
			"message":"pong",
		})
	})
	return r
}
