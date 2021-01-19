package routers

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go-admin/api/middlewares/JWT"
	"go-admin/api/middlewares/log"
	v1 "go-admin/api/routers/api/v1"
	"go-admin/api/routers/api/v1/role"
	"go-admin/api/routers/api/v1/user"
	"go-admin/api/service"
	_ "go-admin/docs" //没有用到也要注册进来，不然读不到swagger文件
	"net/http"
	_ "net/http"
)
//var logru = logrus.New()

func InitRouter() *gin.Engine {
	r:=gin.New()

	// 加载日志中间件
	r.Use(log.Logger())

	//看官方注释文档 ,Recovery 中间件会恢复(recovers) 任何恐慌(panics) 如果存在恐慌，中间件将会写入500。这个中间件还是很必要的，因为当你程序里有些异常情况你没考虑到的时候，程序就退出了，服务就停止了，所以是必要的。
	// 总的来说，程序崩溃时，还是会返回500
	r.Use(gin.Recovery())

	// 访问项目中的静态文件
	r.StaticFS("/upload/files/images", http.Dir(service.GetImagePath()))

	// 加载SWagger
	url := ginSwagger.URL("http://localhost:8081/swagger/doc.json")
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	r.POST("/upload", v1.ImageUpload)
	apiv1 := r.Group("/api/v1")

	// 用户模块注册和登录，不认证
	apiv1.POST("/user/register", user.Register)
	apiv1.POST("/user/login", user.Login)

	// 定义认证中间件
	apiv1.Use(JWT.JWTAuth())

	apiv1.GET("/ping", func(context *gin.Context) {
		log.Info("BBBBBBBBBBBBBBB")
		context.JSON(200, gin.H{
			"message":"pong",
		})
	})
	// 角色模块
	apiv1.POST("/role", role.AddRole)
	apiv1.GET("/role/:id", role.RoleInfo)

	// 用户模块
	apiv1.GET("/user/:id", user.UserInfo)
	apiv1.GET("/userList", user.UserList)

	return r
}
