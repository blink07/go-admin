package JWT

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"go-admin/api/utils/app"
	"go-admin/api/utils/e"
	"go-admin/conf/settings"
	"net/http"
	"time"
)

type JWT struct {
	SigningKey []byte
}

type CustomClaims struct {
	Username string
	RoleId int
	jwt.StandardClaims
}

var (
	TokenExpired     error = errors.New("Token is expired")
	TokenNotValidYet error = errors.New("Token not active yet")
	TokenMalformed   error = errors.New("That's not even a token")
	TokenInvalid     error = errors.New("Couldn't handle this token:")
)

func NewJWT() *JWT {
	return &JWT{
		SigningKey: []byte(settings.ServerSetting.SigningKey),
	}
}

// 创建一个token
func (j *JWT)CreateToken(claims CustomClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,claims)
	return token.SignedString(j.SigningKey)
}


// 解析token
func (j JWT)ParseToken(tokenString string) (*CustomClaims, error)  {
	token, err := jwt.ParseWithClaims(tokenString,&CustomClaims{}, func(token *jwt.Token) (i interface{}, err error) {
		return j.SigningKey, nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed!=0{
				return nil, TokenMalformed
			}else if ve.Errors&jwt.ValidationErrorExpired !=0 {
				return nil, TokenExpired
			}else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, TokenNotValidYet
			} else {
				return nil, TokenInvalid
			}
		}
	}

	if token !=nil {
		if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
			return claims, nil
		}

		return nil, TokenInvalid
	}else {
		return nil, TokenInvalid
	}
}

// 更新token
func (j JWT)RefreshToken(tokenString string) (string, error)  {
	jwt.TimeFunc = func() time.Time {
		return time.Unix(0,0)
	}
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	})
	if err != nil {
		return "", err
	}
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		jwt.TimeFunc = time.Now
		claims.StandardClaims.ExpiresAt = time.Now().Add(1 * time.Hour).Unix()
		return j.CreateToken(*claims)
	}
	return "", TokenInvalid
}

func JWTAuth() gin.HandlerFunc {
	return func(context *gin.Context) {
		appG := app.Gin{context}
		path := context.Request.RequestURI
		println(">>>>>>",path)
		token := context.Request.Header.Get("token")
		println(token)
		j := NewJWT()
		claims,err := j.ParseToken(token)
		if err != nil {
			if err == TokenExpired {
				appG.Response(http.StatusOK, e.TOKEN_NOT_VALID, nil)
				return
			}
			appG.Response(http.StatusOK, e.TOKEN_NOT_VALID, nil)
			context.Abort()
		}
		context.Set("claims",claims)
		context.Next()
	}
}