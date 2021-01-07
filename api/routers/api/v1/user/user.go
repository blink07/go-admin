package user

import (
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"go-admin/api/utils/app"
	"go-admin/api/utils/e"
	"net/http"
)

type UserForm struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email string `json:"email"`
	Mobile string `json:"mobile"`
	IsActive int8 `json:"is_active"`
	Address string `json:"address"`
	RoleId string `json:"role_id"`
}

func Register(c *gin.Context) {
	appG := app.Gin{c}
	var userForm UserForm
	err := c.Bind(&userForm)
	if err != nil {
		appG.Response(http.StatusOK, e.INVALID_PARAMS, nil)
	}
	valid := validation.Validation{}
	valid.Required(userForm.Username,"username").Message("用户名不能为空")
	valid.MinSize(userForm.Username, 6, "username").Message("用户名长度不能小于6位")
	valid.MaxSize(userForm.Username,20, "username").Message("用户名长度不能超过20位")

	valid.Required(userForm.Password, "password").Message("密码不能为空")
	valid.MinSize(userForm.Password, 6, "password").Message("密码长度不能小于6位")
	valid.MaxSize(userForm.Password,20, "password").Message("密码长度不能超过20位")



	return
}
