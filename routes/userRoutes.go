package routes

import (
	"github.com/gin-gonic/gin"
	"movie-app/controller"
	"movie-app/middleWare"
)

func GetUserRoutes(route *gin.RouterGroup) {
	user := route.Group("/user")
	{
		user.POST("/register", controller.Register)
		user.POST("/login", controller.Login)

		//需要用户登录
		userAuth := user.Group("/")
		userAuth.Use(middleWare.AuthMiddleWare())
		{

			userAuth.GET("/info", controller.UserInfo)
			userAuth.POST("/info/modify", controller.UserModify)
			userAuth.POST("/info/password", controller.UserModifyPassword)
		}
	}
}
