package JWT

import "github.com/dgrijalva/jwt-go"

type CustomClaims struct {
	username string
	RoleId int
	jwt.StandardClaims
}



