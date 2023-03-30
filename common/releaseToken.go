package common

import (
	"github.com/golang-jwt/jwt/v4"
	"github.com/spf13/viper"
	"movie-app/model"
	"time"
)

/*
* 发放 admin access token
 */
func ReleaseAdminAccessToken(admin model.Admin) (string, error) {
	accessJwtKey := []byte(viper.GetString("server.access_jwt_secret"))

	expirationTime := time.Now().Add(1 * time.Minute) // 过期时间 1 分钟

	AccessClaims := &AdminClaims{
		AdminId:   admin.ID,
		Authority: admin.Authority,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "john",
			Subject:   AccessTypeToken,
		},
	}
	return GenerateAdminToken(accessJwtKey, AccessClaims)
}

/*
* 发放 user access token
 */
func ReleaseUserAccessToken(user model.User) (string, error) {
	accessJwtKey := []byte(viper.GetString("server.access_jwt_secret"))

	expirationTime := time.Now().Add(1 * time.Minute) // 过期时间 1 分钟

	AccessClaims := &UserClaims{
		UserId: user.ID,
		State:  user.State,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "john",
			Subject:   AccessTypeToken,
		},
	}
	return GenerateUserToken(accessJwtKey, AccessClaims)
}

/*
* 发放 admin refresh token
 */
func ReleaseAdminRefreshToken(admin model.Admin) (string, error) {
	refreshJwtKey := []byte(viper.GetString("server.refresh_jwt_secret"))
	// token过期时间
	expirationTime := time.Now().Add(1 * 24 * time.Hour) // 14天有效

	refreshClaims := &AdminClaims{
		AdminId:   admin.ID,
		Authority: admin.Authority,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "john",
			Subject:   RefreshTypeToken,
		},
	}

	return GenerateAdminToken(refreshJwtKey, refreshClaims)
}

/*
* 发放 user refresh token
 */
func ReleaseUserRefreshToken(user model.User) (string, error) {
	refreshJwtKey := []byte(viper.GetString("server.refresh_jwt_secret"))
	// token过期时间
	expirationTime := time.Now().Add(1 * 24 * time.Hour) // 14天有效

	refreshClaims := &UserClaims{
		UserId: user.ID,
		State:  user.State,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "john",
			Subject:   RefreshTypeToken,
		},
	}

	return GenerateUserToken(refreshJwtKey, refreshClaims)
}

/*
* 生成 admin token 字符串
 */
func GenerateAdminToken(key []byte, claims *AdminClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(key)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

/*
* 生成 user token 字符串
 */
func GenerateUserToken(key []byte, claims *UserClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(key)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
