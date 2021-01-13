package models

import (
	"github.com/pkg/errors"
	"go-admin/api/middlewares/log"
	"go-admin/conf/settings"
)

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
	var u User
	record := db.Where("username=?", user["username"]).First(&u).RecordNotFound()

	if !record {
		return errors.New("用户名已经存在~")
	}

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


func Login(username string, password string) (role_id int, err error) {
	var u User
	record := db.Where("username=? AND password=?", username, password).First(&u).RecordNotFound()

	if record {
		return 0, errors.New("用户名或密码错误~")
	}
	return u.RoleID, nil
}

func UserInfo(id int) (*User, error) {
	var u User
	err := db.Where("id=?", id).First(&u).Error

	if err != nil{
		return nil, err
	}

	var r Role
	err = db.Where("id=?", u.RoleID).First(&r).Error
	if err!= nil {
		return nil,err
	}
	u.Role = r

	return &u, nil
}

type UserListForm struct {
	Username string
	Email string
	Mobile string
	RoleId int
	RoleName string
}

func UserList(pageNum int) ([]*UserListForm, error) {
	//var user []*User
	var userList []*UserListForm
	err := db.Table("admin_user").Select("admin_user.username, admin_user.email, admin_user.mobile, admin_user.role_id, admin_role.role_name").Joins("left join admin_role on admin_user.role_id=admin_role.id").Offset(pageNum).Limit(settings.AppSetting.PageSize).Scan(&userList).Error

	if err!=nil{
		return nil, err
	}

	return userList, nil
}