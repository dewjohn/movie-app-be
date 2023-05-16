package routes

import (
	"github.com/gin-gonic/gin"
	"movie-app/controller"
)

func GetCatetoryRoutes(route *gin.RouterGroup) {
	category := route.Group("/category")
	{
		category.GET("/type", controller.GetTypeCategory)         // 类型
		category.GET("/origin", controller.GetOriginCategory)     // 地区
		category.GET("/language", controller.GetLanguageCategory) // 语言
		category.GET("/release", controller.GetReleaseCategory)   // 年代
	}
}
