package user

import (
	"encoding/json"
	"github.com/dgrijalva/jwt-go"
	"go-admin/api/middlewares/JWT"
	"go-admin/api/models"
	"go-admin/api/myredis"
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

// 获取用户详细信息
func (user UserService)UserInfo() (*models.User, error){
	var myUser *models.User

	cacheUser := UserCache{Id: user.Id}
	key := cacheUser.GetUserKey()
	// 如果缓存中有则取缓存中的数据
	if myredis.Exists(key) {
		data, err := myredis.Get(key)
		if err != nil {
			return nil, err
		}else {
			json.Unmarshal(data, &myUser)
			return myUser, nil
		}
	}
	// 如果缓存中没有则取数据库内的数据
	userInfo, err := models.UserInfo(user.Id)
	if err != nil {
		return nil, err
	}
	// 从数据库取出数据后放入缓存以便后面进来的请求使用
	myredis.Set(key, userInfo, 3600)
	return userInfo, nil
}


func (user UserService)UserList() ([]*models.UserListForm, error){

	var usersCacheList []*models.UserListForm

	usersCache := UserCache{
		Id: user.Id,
		RoleId: user.RoleId,
		Is_Active: user.IsActive,
		PageSize: user.PageSize,
		PageNum: user.PageNum,
	}
	userKeys := usersCache.GetUsersKeys()
	// 检查缓存中是否有数据，有则在缓存中取数据
	if myredis.Exists(userKeys) {
		println(">>>>>>>>>>>>>>")
		data, err := myredis.Get(userKeys)
		if err != nil {
			return nil, err
		}
		err = json.Unmarshal(data, &usersCacheList)
		if err!=nil {
			return nil ,err
		}
		return usersCacheList, nil
	}
	// 缓存中没有数据则从数据库中取
	userList,err := models.UserList(user.PageNum)
	if err!=nil {
		return nil, err
	}
	// 取出后将数据放入缓存
	err = myredis.Set(userKeys, userList, 3600)
	if err !=nil{
		return nil, err
	}
	return userList, nil
}