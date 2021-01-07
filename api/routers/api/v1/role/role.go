package role

import (
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	//"github.com/go-playground/validator/v10"
	"github.com/unknwon/com"
	"go-admin/api/service/roles"
	"go-admin/api/utils/app"
	"go-admin/api/utils/e"
	"net/http"
)

type RoleForm struct {
	RoleName string `json:"role_name"`
	Description string `json:"description"`
}


// 增加角色
func AddRole(c *gin.Context) {
	appG := app.Gin{c}
	var rf RoleForm
	err := c.Bind(&rf)
	if err!= nil {
		appG.Response(http.StatusOK, e.INVALID_PARAMS, nil)
		return
	}
	valid := validation.Validation{}
	valid.MinSize(rf.RoleName, 2, "role_name").Message("角色名称长度不能少于3")
	valid.MaxSize(rf.RoleName, 20,"role_name").Message("角色名称长度不能超过20")
	if valid.HasErrors() {
		validErr := app.MarkErrors(valid.Errors)
		appG.Response(http.StatusOK, e.INVALID_PARAMS, validErr)
		return
	}

	roleService := roles.Role{RoleName: rf.RoleName, Description: rf.Description}
	err = roleService.AddRole()
	if err != nil {
		appG.Response(http.StatusOK, e.DATA_INSERT_INOT_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)


}

// 查看某个角色信息
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

	role := roles.Role{ID:id}

	r, err := role.RoleInfo()
	if err !=nil {
		appG.Response(http.StatusOK,e.ERROR, nil)
	}

	appG.Response(http.StatusOK, e.SUCCESS, r)

}
