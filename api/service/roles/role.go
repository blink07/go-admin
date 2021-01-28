package roles

import (
	"github.com/360EntSecGroup-Skylar/excelize"
	"go-admin/api/models"
	"io"
)

type Role struct {
	ID int
	RoleName string
	Description string
	CreatedBy string
	ModifiedBy string

	PageNum int
	PageSize int
}

func (r *Role) AddRole() error {
	err := models.AddRole(map[string]interface{}{"role_name":r.RoleName, "description":r.Description})

	if err!= nil {
		return err
	}
	return nil
}

func (r *Role) RoleInfo() (*models.Role, error) {
	role, err := models.RoleInfo(r.ID)

	if err != nil {
		return nil, err
	}
	return role, nil
}


func (r *Role) RoleList() ([]*models.Role, error) {
	roleList, err := models.RoleList()
	if err!=nil {
		return nil, err
	}
	return roleList, err
}

// excel导入role
func (r *Role)ImportRole(read io.Reader) error {
	xls, err := excelize.OpenReader(read)
	if err!= nil {
		return err
	}

	rows := xls.GetRows("角色信息")
	for irow, row := range rows {
		if irow >0 {
			var data []string
			for _, cell := range row{
				data = append(data,cell)
			}
			models.AddRole(map[string]interface{}{"role_name":data[1], "description":data[2]})
		}
	}
	return nil
}