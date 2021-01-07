package models

import (
	"github.com/jinzhu/gorm"
	"go-admin/api/middlewares/log"
)

type Role struct {
	Model
	
	RoleName string `json:"role_name"`
	Description string `json:"description"`
}


func AddRole(role map[string]interface{}) error {

	err := db.Create(&Role{
		RoleName: role["role_name"].(string),
		Description: role["description"].(string),
	}).Error

	if err!= nil {
		log.Error("Role Created Faile, err: %v", err)
		return err
	}

	return nil
}


// 获取角色信息
func RoleInfo(id int) (*Role, error) {
	var role Role
	err := db.Where("id = ? AND deleted_on = ? ", id, 0).First(&role).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return &role, nil
}