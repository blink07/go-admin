package routers

import (
	"github.com/gin-gonic/gin"
	"go-admin/api/middlewares/log"
)
//var logru = logrus.New()

func InitRouter() *gin.Engine {
	r:=gin.New()

	// 加载日志中间件
	r.Use(log.Logger())

	//看官方注释文档 ,Recovery 中间件会恢复(recovers) 任何恐慌(panics) 如果存在恐慌，中间件将会写入500。这个中间件还是很必要的，因为当你程序里有些异常情况你没考虑到的时候，程序就退出了，服务就停止了，所以是必要的。
	// 总的来说，程序崩溃时，还是会返回500
	r.Use(gin.Recovery())

	apiv1 := r.Group("/api/v1")
	apiv1.GET("/ping", func(context *gin.Context) {
		log.Info("BBBBBBBBBBBBBBB")
		context.JSON(200, gin.H{
			"message":"pong",
		})
	})
	return r
}
