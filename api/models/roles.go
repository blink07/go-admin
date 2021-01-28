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


// 查询所有角色
func RoleList() ([]*Role, error){
	var roleList []*Role
	err := db.Find(&roleList).Error
	if err != nil {
		return nil, err
	}
	return roleList, nil
}


// 硬删除角色
func DeleteRole() bool {
	println(111111)
	var role Role
	err := db.Unscoped().Where("deleted_on != ? ", 0).Delete(&role).Error
	if err!=nil {
		println("AAAAAAAAAA"+err.Error())
		return false
	}
	println(22222)
	return true
}

