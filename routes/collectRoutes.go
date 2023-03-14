package routes

import (
	"github.com/gin-gonic/gin"
	"movie-app/controller"
	"movie-app/middleWare"
)

func GetCollectRoutes(route *gin.RouterGroup) {
	collectAuth := route.Group("/collect")
	collectAuth.Use(middleWare.UserAuthMiddleWare())
	{
		collectAuth.POST("/add", controller.Collect)
		collectAuth.GET("/get", controller.GetCollectList)
	}
}
