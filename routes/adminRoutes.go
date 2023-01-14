package routes

import (
	"github.com/gin-gonic/gin"
	admin2 "movie-app/controller/admin"
	"movie-app/middleWare"
)

func GetAdminRoutes(route *gin.RouterGroup) {
	admin := route.Group("/admin")
	{
		admin.POST("/login", admin2.AdminLogin)

		// 需要管理员登陆
		adminAuth := admin.Group("/")
		adminAuth.Use(middleWare.AdminAuthMiddleWare())
		{

		}
	}
}
