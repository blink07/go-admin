package role

import (
	"github.com/gin-gonic/gin"
	"go-admin/api/utils/app"
	"github.com/unknwon/com"
	"github.com/astaxie/beego/validation"
	"go-admin/api/utils/e"
	"net/http"
)

type RoleForm struct {

}

func RoleInfo(c *gin.Context) {
	appG := app.Gin{c}

	id := com.StrTo(c.Param("id")).MustInt()

	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("ID必须大于1")

	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusOK, e.INVALID_PARAMS, nil)
		return
	}


}
