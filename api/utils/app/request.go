package app

import (
	"github.com/astaxie/beego/validation"
	"go-admin/api/middlewares/log"
)

// 在同时校验多个参数时对不合格的参数逐个提示
func MarkErrors(errors []*validation.Error) string {
	for _,err := range errors {
		log.Info(err.Key, err.Message)
		return err.Message
	}
	return ""
}