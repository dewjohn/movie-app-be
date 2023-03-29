package middleWare

import (
	"github.com/gin-gonic/gin"
	"movie-app/common"
	"movie-app/model"
	"net/http"
	"strings"
)

func UserAuthMiddleWare() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 获取 header
		tokenString := ctx.GetHeader("Authorization")

		// 验证token格式
		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer ") {
			ctx.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "请先登录"})
			ctx.Abort()
			return
		}

		tokenString = tokenString[7:]

		token, claims, err := common.ParseToken(tokenString)
		if err != nil || !token.Valid {
			ctx.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "请先登录"})
			ctx.Abort()
			return
		}

		// token 通过验证，获取claims中的id
		userId := claims.UserId
		DB := common.GetDB()
		var user model.User
		DB.First(&user, userId)

		// 用户
		if user.ID == 0 {
			ctx.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "请先登录"})
			ctx.Abort()
			return
		}
		// 用户存在,将user信息写入上下文
		ctx.Set("userId", user.ID)
		ctx.Next()
	}
}
