package routes

import (
	"movie-app/controller"
	"movie-app/middleWare"

	"github.com/gin-gonic/gin"
)

func CollectRoute(r *gin.Engine) *gin.Engine {
	r.Use(middleWare.CorsMiddleWare())
	r.POST("/api/auth/register", controller.Register)
	r.POST("/api/auth/login", controller.Login)
	r.GET("api/auth/info", middleWare.AuthMiddleWare(), controller.Info)
	return r
}
