package app

import (
	"github.com/gin-gonic/gin"
	"go-admin/api/utils/e"
)

type Gin struct {
	C *gin.Context
}

func (g *Gin) Response(httpCode, errCode int, data interface{}) {
	g.C.JSON(httpCode, gin.H{
		"code":errCode,
		"data":data,
		"message":e.GetErrMsg(errCode),
	})
}