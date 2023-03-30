package common

import (
	"github.com/golang-jwt/jwt/v4"
	"github.com/spf13/viper"
	"movie-app/model"
	"regexp"
)

type AdminClaims struct {
	AdminId   uint
	Authority int
	jwt.RegisteredClaims
}

func ReleaseAdminToken(admin model.Admin) (string, string, error) {
	refreshToken, err := ReleaseAdminRefreshToken(admin)
	if err != nil {
		return "", "", err
	}
	accessToken, err := ReleaseAdminAccessToken(admin)
	if err != nil {
		return "", "", err
	}
	return refreshToken, accessToken, nil
}

/**
* 解析 admin token
 */
func ParseAdminToken(tokenString, tokenType string) (*jwt.Token, *AdminClaims, error, bool) {
	var jwtKey []byte
	if tokenType == AccessTypeToken {
		jwtKey = []byte(viper.GetString("server.access_jwt_secret"))
	} else {
		jwtKey = []byte(viper.GetString("server.refresh_jwt_secret"))
	}

	claims := &AdminClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	isExpired := false
	// 判断token是否过期
	if err != nil {
		reg := regexp.MustCompile(`token is expired`)
		if reg.MatchString(err.Error()) {
			isExpired = true
		}
	}
	return token, claims, err, isExpired
}
