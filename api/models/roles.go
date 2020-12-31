package models

import "go-admin/api/middlewares/log"

type Role struct {
	Model
	
	RoleName string `json:"role_name"`
	Description string `json:"description"`
}


func AddRole(role map[string]interface{}) error {

	err := db.Create(&Role{
		RoleName: role["role_name"].(string),
		Description: role["role_desciption"].(string),
	}).Error

	if err!= nil {
		log.Error("Role Created Faile, err: %v", err)
		return err
	}

	return nil
}