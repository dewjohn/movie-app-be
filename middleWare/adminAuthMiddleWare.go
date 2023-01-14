package middleWare

import (
	"github.com/gin-gonic/gin"
	"movie-app/common"
	"movie-app/model"
	"net/http"
	"strings"
)

func AdminAuthMiddleWare() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 获取 header
		tokenString := ctx.GetHeader("Authorization")
		// 验证 token 格式
		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer ") {
			ctx.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "请先登录"})
			ctx.Abort()
			return
		}

		tokenString = tokenString[7:]
		adminToken, adminClaims, err := common.ParseAdminToken(tokenString)
		if err != nil || !adminToken.Valid {
			ctx.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "请先登录"})
			ctx.Abort()
			return
		}

		// token 通过验证，获取 adminClaims 中的 adminId
		adminId := adminClaims.AdminId
		DB := common.GetDB()
		var admin model.Admin
		DB.First(&admin, adminId)

		// 查无此管理员
		if admin.ID == 0 {
			ctx.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "请先登录"})
			ctx.Abort()
			return
		}
		// 查询成功
		ctx.Set("adminId", admin.ID)
		ctx.Set("admin", admin)
		ctx.Next()
	}
}
