package user

import (
	"github.com/dgrijalva/jwt-go"
	"go-admin/api/middlewares/JWT"
	"go-admin/api/models"
	h "go-admin/api/utils/hash"
	"time"
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

func (user *UserService) Login() (map[string]interface{}, error) {

	bytes := h.Encryption([]byte(user.Password))
	roleId, err := models.Login(user.Username, bytes)
	if err != nil {
		return nil, err
	}
	token, err := tokenNext(roleId, user.Username)
	if err != nil {
		return nil, err
	}
	return map[string]interface{}{"token":token}, nil
}

// 登录以后签发jwt
func tokenNext(roleId int, username string) (string, error)  {
	key := JWT.NewJWT()
	claims := JWT.CustomClaims{
		Username: username,
		RoleId: roleId,
		StandardClaims:jwt.StandardClaims{
			NotBefore: int64(time.Now().Unix()-1000),
			ExpiresAt: int64(time.Now().Unix() + 60 * 60 * 24 * 7),
			Issuer: "blink07",
		},
	}
	token, err := key.CreateToken(claims)
	// TODO 还未将token保存
	if err != nil{
		return "", err
	}
	return token, nil
}


func (user UserService)UserInfo() (*models.User, error){

	userInfo, err := models.UserInfo(user.Id)
	if err != nil {
		return nil, err
	}
	return userInfo, nil
}


func (user UserService)UserList() ([]*models.UserListForm, error){
	userList,err := models.UserList(user.PageNum)
	if err!=nil {
		return nil, err
	}
	return userList, nil
}