package user

import (
	"go-admin/api/models"
	h "go-admin/api/utils/hash"
)

type UserService struct {
	Id int
	Username string
	Password string
	Email string
	Mobile string
	Address string
	IsActive int8
	RoleId int

	PageSize int
	PageNum int
}



func (user *UserService)UserRegister() error {
	println("aaa" + string(user.IsActive))
	// 1.对password加密
	bytes := h.Encryption([]byte(user.Password))

	err := models.AddUser(
		map[string]interface{}{
			"username":user.Username,
			"password": bytes,
			"email": user.Email,
			"mobile": user.Mobile,
			"is_active":user.IsActive,
			"address":user.Address,
			"role_id":user.RoleId,
		})
	if err != nil {
		return err
	}
	// 2.加密后存储
	return nil
}

func (user *UserService) Login(username string, password string) (map[string]interface{}, error) {

	bytes := h.Encryption([]byte(password))
	err := models.Login(username, password)
	if err != nil {
		return nil, err
	}

	return nil, nil
}