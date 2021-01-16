package myredis

import (
	"encoding/json"
	"github.com/gomodule/redigo/redis"
	"go-admin/conf/settings"
	"time"
)

var RedisConn *redis.Pool


// 链接redis并初始化
func Setup() error {
	RedisConn = &redis.Pool{
		MaxIdle: settings.RedisSettings.MaxIdle,
		MaxActive: settings.RedisSettings.MaxActive,
		IdleTimeout: settings.RedisSettings.IdleTimeout,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", settings.RedisSettings.Host)
			if err !=nil {
				return nil, err
			}
			if settings.RedisSettings.Password != "" {
				if _,err := c.Do("AUTH", settings.RedisSettings.Password); err != nil {
					c.Close()
					return nil, err
				}
			}
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("/PING")
			return err
		},
	}
	println("redis connected ...")
	return nil
}


// 柴数据到redis
func Set(key string, data interface{}, time int) error  {
	conn := RedisConn.Get()
	defer conn.Close()

	value, err := json.Marshal(data)
	if err != nil {
		return err
	}
	_, err = conn.Do("SET", key, value)
	if err != nil {
		return err
	}
	// 设置过期时间
	_,err = conn.Do("EXPIRE", key, time)
	if err != nil {
		return err
	}
	return nil
}


// 从redis取数据
func Get(key string) ([]byte, error) {
	conn := RedisConn.Get()
	defer conn.Close()
	value, err := redis.Bytes(conn.Do("GET", key))

	if err !=nil {
		return nil, err
	}
	return value, nil
}

// 检查key值是否存在
func Exists(key string) bool {

	conn := RedisConn.Get()

	defer conn.Close()

	exists, err := redis.Bool(conn.Do("EXISTS", key))
	if err != nil {
		return false
	}
	return exists
}


// 删除数据
func Delete(key string) (bool, error) {
	conn := RedisConn.Get()

	defer conn.Close()
	return redis.Bool(conn.Do("DEL", key))

}

// 模糊数据
func LikeDeletes(key string) error {
	conn := RedisConn.Get()

	defer conn.Close()

	keys, err := redis.Strings(conn.Do("KEYS", "*" + key + "*"))
	if err != nil {
		return err
	}
	for _, key := range keys {
		_, err = Delete(key)
		if err != nil {
			return err
		}
	}
	return nil
}