package role

import (
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"go-admin/api/service"
	"strconv"
	"time"

	"github.com/tealeg/xlsx"
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


func ExportRoleExcel(c *gin.Context) {
	appG := app.Gin{c}
	role := roles.Role{}
	roleList,err := role.RoleList()
	if err != nil {
		appG.Response(http.StatusOK, e.ERROR, err.Error())
		return
	}

	file := xlsx.NewFile()
	sheet, err := file.AddSheet("角色信息")
	if err != nil {
		appG.Response(http.StatusOK, e.ERROR, err.Error())
		return
	}

	titles := []string{"ID", "RoleName", "Description", "CreatedOn", "ModifiedOn", "DeletedOn"}
	row := sheet.AddRow()

	var cell *xlsx.Cell
	for _, title := range titles {
		cell = row.AddCell()
		cell.Value = title
	}

	for _, v := range roleList {
		values := []string{
			strconv.Itoa(v.ID),
			v.RoleName,
			v.Description,
			strconv.Itoa(v.CreatedOn),
			strconv.Itoa(v.ModifiedOn),
			strconv.Itoa(v.DeletedOn),
		}
		row = sheet.AddRow()
		for _,value := range values {
			cell = row.AddCell()
			cell.Value = value
		}
	}

	time := strconv.Itoa(int(time.Now().Unix()))
	filename := "role-" + time + ".xlsx"
	fullPathUrl := service.GetExcelFullUrl(filename)

	err = file.Save(fullPathUrl)
	if err != nil {
		appG.Response(http.StatusOK, e.ERROR, err.Error())
		return
	}
	appG.Response(http.StatusOK, e.SUCCESS, map[string]string{
		"export_save_url":service.GetExcelFullUrl(filename),
		"export_url": service.GetExcelFullPath()+filename,
	})
	return

}

// excel 导入角色
func ImportExcelRole(c *gin.Context){
	appG := app.Gin{c}

	file, _, err := c.Request.FormFile("file")
	if err != nil {
		appG.Response(http.StatusOK, e.ERROR, nil)
		return
	}
	roleService := roles.Role{}

	err = roleService.ImportRole(file)
	if err != nil {
		appG.Response(http.StatusOK, e.ERROR, err.Error())
		return
	}
	appG.Response(http.StatusOK, e.SUCCESS, nil)
	return
}