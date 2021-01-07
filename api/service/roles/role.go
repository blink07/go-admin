package roles

import "go-admin/api/models"

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
