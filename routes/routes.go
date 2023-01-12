package routes

import (
	"github.com/gin-gonic/gin"
	"movie-app/controller"
)

func CollectRoute(r *gin.Engine) *gin.Engine {
	r.POST("/api/auth/register", controller.Register)
	return r
}
