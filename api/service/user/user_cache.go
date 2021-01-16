package user

import (
	"strconv"
	"strings"
)

type UserCache struct {
	Id int
	RoleId int
	Is_Active int8

	PageNum int
	PageSize int
}

const USER_CACHE = "USER"


// 拼接User的Key值
func (u *UserCache) GetUserKey() string {
	return USER_CACHE +"_" + strconv.Itoa(u.Id)
}

// 拼接Users的key值
func (u *UserCache) GetUsersKeys() string {
	keys := []string{
		USER_CACHE,
		"LIST",
	}
	if u.Id >0 {
		keys = append(keys, strconv.Itoa(u.Id))
	}

	if u.RoleId >0 {
		keys = append(keys, strconv.Itoa(u.RoleId))
	}

	if u.Is_Active > 0 {
		keys = append(keys, strconv.Itoa(int(u.Is_Active)))
	}

	if u.PageNum >0 {
		keys = append(keys, strconv.Itoa(u.PageNum))
	}

	if u.PageSize > 0 {
		keys = append(keys, strconv.Itoa(u.PageSize))
	}

	return strings.Join(keys, "_")

}
