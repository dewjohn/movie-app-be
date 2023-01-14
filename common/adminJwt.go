package common

import (
	"github.com/golang-jwt/jwt/v4"
	"movie-app/model"
	"time"
)

var adminJwtKey = []byte("server.jwt_secret")

type AdminClaims struct {
	AdminId uint
	jwt.RegisteredClaims
}

// 发放管理员 token
func ReleaseAdminToken(admin model.Admin) (string, error) {
	expirationTime := time.Now().Add(7 * 24 * time.Hour)
	adminClaims := &AdminClaims{
		AdminId: admin.ID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "movie-app by john",
			Subject:   "admin_token",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, adminClaims)
	tokenString, err := token.SignedString(adminJwtKey)

	if err != nil {
		return "", err
	}
	return tokenString, nil

}

func ParseAdminToken(tokenString string) (*jwt.Token, *AdminClaims, error) {
	adminClaims := &AdminClaims{}
	token, err := jwt.ParseWithClaims(tokenString, adminClaims, func(token *jwt.Token) (i interface{}, e error) {
		return adminJwtKey, nil
	})
	return token, adminClaims, err
}
