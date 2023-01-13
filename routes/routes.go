package routes

import (
	"movie-app/middleWare"

	"github.com/gin-gonic/gin"
)

func CollectRoute(r *gin.Engine) *gin.Engine {
	r.Use(middleWare.CorsMiddleWare())
	v1 := r.Group("/api/v1")
	{
		GetUserRoutes(v1)
		GetMovieRoutes(v1)
	}
	return r
}
