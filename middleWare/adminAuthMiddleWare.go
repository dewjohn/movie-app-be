package middleWare

import (
	"github.com/gin-gonic/gin"
	"movie-app/common"
	"movie-app/model"
	"movie-app/response"
	"net/http"
	"strings"
)

func AdminAuthMiddleWare() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		res := response.ResponseStruct{
			HttpStatus: http.StatusUnauthorized,
			Code:       4011,
			Data:       nil,
			Msg:        response.Unauthorized,
		}

		// 获取 Authorization header
		tokenString := ctx.GetHeader("Authorization")
		parseSuccess, claims, isExpired := ParseAdminTokenString(tokenString, common.AccessTypeToken)
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
		adminId := claims.AdminId
		var admin model.Admin
		DB := common.GetDB()
		if err := DB.First(&admin, adminId).Error; err != nil || admin.ID == 0 {
			response.HandleResponse(ctx, res)
			ctx.Abort()
			return
		}
		ctx.Set("adminId", adminId)
		ctx.Set("adminAuthorization", admin.Authority)
		ctx.Next()
	}
}

func RefreshAdminTokenMiddleWare() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		res := response.ResponseStruct{
			HttpStatus: http.StatusUnauthorized,
			Code:       4011,
			Data:       nil,
			Msg:        response.Unauthorized,
		}
		// 获取 Authorization header
		tokenString := ctx.GetHeader("Authorization")
		parseSuccess, claims, isExpired := ParseAdminTokenString(tokenString, common.RefreshTypeToken)
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
		// 验证管理员是否存在
		adminId := claims.AdminId
		var admin model.Admin
		DB := common.GetDB()
		if err := DB.First(&admin, adminId).Error; err != nil || admin.ID == 0 {
			response.HandleResponse(ctx, res)
			ctx.Abort()
			return
		}
		ctx.Set("adminId", admin.ID)
		ctx.Set("adminAuthorization", admin.Authority)
		ctx.Next()
	}
}

func ParseAdminTokenString(tokenString, tokenType string) (bool, *common.AdminClaims, bool) {
	if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer") {
		return false, nil, false
	}
	tokenString = tokenString[7:]
	token, claims, err, isExpired := common.ParseAdminToken(tokenString, tokenType)
	if err != nil || !token.Valid {
		return false, nil, isExpired
	}
	if tokenType != claims.Subject {
		return false, nil, isExpired
	}
	return true, claims, isExpired
}
