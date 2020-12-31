package models

import "go-admin/api/middlewares/log"

type User struct {
	Model
	Username string `json:"username"`
	Password string `json:"password"`
	Email string `json:"email"`
	Mobile string `json:"mobile"`
	IsActive int8 `json:"is_active"`
	Address string `json:"address"`

	RoleID int `json:"role_id" gorm:"index"`
	Role Role `json:"role"`
}



func AddUser(user map[string]interface{}) error {

	err := db.Create(&User{
		Username: user["username"].(string),
		Password: user["password"].(string),
		Email: user["email"].(string),
		Mobile: user["mobile"].(string),
		IsActive: user["is_active"].(int8),
		Address: user["address"].(string),

		RoleID: user["role_id"].(int),
	}).Error
	if err != nil {
		log.Info("Create User err:%v", err)
		return err
	}
	return nil
}