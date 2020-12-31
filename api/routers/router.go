package routers

import (
	"github.com/gin-gonic/gin"
	"go-admin/api/middlewares/log"
)

func InitRouter() *gin.Engine {
	r:=gin.New()

	r.Use(log.Logger())

	return r
}
