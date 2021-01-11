package user

import (
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"go-admin/api/service/user"
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
	RoleId int `json:"role_id"`
}

func Register(c *gin.Context) {
	appG := app.Gin{c}
	var userForm UserForm
	err := c.Bind(&userForm)
	println(">>>>>>>>>>>>",userForm.RoleId)
	if err != nil {
		appG.Response(http.StatusOK, e.INVALID_PARAMS, nil)
		return
	}
	valid := validation.Validation{}
	valid.Required(userForm.Username,"username").Message("用户名不能为空")
	valid.MinSize(userForm.Username, 6, "username").Message("用户名长度不能小于6位")
	valid.MaxSize(userForm.Username,20, "username").Message("用户名长度不能超过20位")

	valid.Required(userForm.Password, "password").Message("密码不能为空")
	valid.MinSize(userForm.Password, 6, "password").Message("密码长度不能小于6位")
	valid.MaxSize(userForm.Password,20, "password").Message("密码长度不能超过20位")
	if userForm.Email != ""{
		valid.Email(userForm.Email, "email").Message("邮箱格式不正确")
	}
	valid.Required(userForm.RoleId, "role_id").Message("角色ID不能为空")
	valid.Min(userForm.RoleId, 1, "role_id").Message("角色ID不能小于1")

	if valid.HasErrors() {
		errStr := app.MarkErrors(valid.Errors)
		appG.Response(http.StatusOK, e.INVALID_PARAMS, errStr)
		return
	}

	user := user.UserService{Username:userForm.Username, Password: userForm.Password, Email: userForm.Email, Mobile: userForm.Mobile, IsActive: userForm.IsActive, RoleId: userForm.RoleId}
	err = user.UserRegister()
	if err != nil {
		appG.Response(http.StatusOK, e.ERROR, err.Error())
		return
	}
	appG.Response(http.StatusOK, e.SUCCESS, "注册成功~")
	return
}

// 用户登录
func Login(c *gin.Context)  {
	appG := app.Gin{c}
	var userForm UserForm
	err := c.Bind(&userForm)
	if err != nil {
		appG.Response(http.StatusOK, e.INVALID_PARAMS, "参数传入错误")
		return
	}
	valid := validation.Validation{}
	valid.Required(userForm.Username, "username").Message("请输入密码")
	valid.MinSize(userForm.Username, 6, "username").Message("用户名输入错误")
	valid.MaxSize(userForm.Username, 20, "username").Message("用户名输入错误")
	valid.Required(userForm.Password, "password").Message("请输入密码")
	valid.MinSize(userForm.Password, 6, "password").Message("密码输入错误")
	valid.MaxSize(userForm.Password, 20, "password").Message("密码输入错误")

	user := user.UserService{Username: userForm.Username, Password: userForm.Password}
	data, err := user.Login()
	if err != nil {
		appG.Response(http.StatusOK, e.USERNAME_OR_PASSWORD, err.Error())
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, data)
}

func UserInfo(c *gin.Context) {
	appG := app.Gin{c}

	id := com.StrTo(c.Param("id")).MustInt()

	valid := validation.Validation{}
	valid.Required(id, "id")
	valid.Min(id, 1, "id").Message("用户ID必须大于1")
	if valid.HasErrors() {
		//println(valid.Errors)
		appG.Response(http.StatusOK, e.INVALID_PARAMS, "参数传入错误")
		return
	}

	userService := user.UserService{Id:id}
	userInfo,err := userService.UserInfo()
	if err!=nil {
		appG.Response(http.StatusOK, e.ERROR, err.Error())
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, userInfo)
	return

}