package middleWare

import (
	"github.com/gin-gonic/gin"
	"movie-app/common"
	"movie-app/model"
	"movie-app/response"
	"net/http"
	"strings"
)

func UserAuthMiddleWare() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		res := response.ResponseStruct{
			HttpStatus: http.StatusUnauthorized,
			Code:       4011,
			Data:       nil,
			Msg:        response.Unauthorized,
		}

		// 获取 Authorization header
		tokenString := ctx.GetHeader("Authorization")
		parseSuccess, claims, isExpired := ParseUserTokenString(tokenString, common.AccessTypeToken)
		if !parseSuccess {
			if isExpired {
				// token过期了
				res.Msg = response.TokenExpired
				res.Code = 4010
			}
			response.HandleResponse(ctx, res)
			ctx.Abort()
			return
		}
		// 验证用户或者管理员是否存在
		userId := claims.UserId
		var user model.User
		DB := common.GetDB()
		if err := DB.First(&user, userId).Error; err != nil || user.ID == 0 {
			response.HandleResponse(ctx, res)
			ctx.Abort()
			return
		}
		ctx.Set("userId", user.ID)
		ctx.Set("userState", user.State)
		ctx.Next()

	}
}

func RefreshUserTokenMiddleWare() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		res := response.ResponseStruct{
			HttpStatus: http.StatusUnauthorized,
			Code:       4011,
			Data:       nil,
			Msg:        response.Unauthorized,
		}
		// 获取 Authorization header
		tokenString := ctx.GetHeader("Authorization")
		parseSuccess, claims, isExpired := ParseUserTokenString(tokenString, common.RefreshTypeToken)
		if !parseSuccess {
			if isExpired {
				// token过期了
				res.Msg = response.TokenExpired
				res.Code = 4010
			}
			response.HandleResponse(ctx, res)
			ctx.Abort()
			return
		}
		// 验证用户是否存在
		userId := claims.UserId
		var user model.User
		DB := common.GetDB()
		if err := DB.First(&user, userId).Error; err != nil || user.ID == 0 {
			response.HandleResponse(ctx, res)
			ctx.Abort()
			return
		}
		ctx.Set("userId", user.ID)
		ctx.Set("state", user.State)
		ctx.Next()
	}
}

func ParseUserTokenString(tokenString, tokenType string) (bool, *common.UserClaims, bool) {
	if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer") {
		return false, nil, false
	}
	tokenString = tokenString[7:]
	token, claims, err, isExpired := common.ParseUserToken(tokenString, tokenType)
	if err != nil || !token.Valid {
		return false, nil, isExpired
	}
	if tokenType != claims.Subject {
		return false, nil, isExpired
	}
	return true, claims, isExpired
}
