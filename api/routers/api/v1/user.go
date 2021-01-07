package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"github.com/astaxie/beego/validation"
)

func GetUserInfo(c *gin.Context)  {
	id := com.StrTo(c.Param("id")).MustInt()
	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("ID必须大于1")
	if valid.HasErrors() {

	}
}
