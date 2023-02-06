package routes

import (
	"movie-app/controller"

	"github.com/gin-gonic/gin"
)

func GetMovieRoutes(route *gin.RouterGroup) {
	movie := route.Group("movie")
	{
		movie.GET("/list/get", controller.GetMovieList) // 获取视频列表
		movie.GET("/search", controller.SerchMovie)     // 搜索视频
		movie.GET("/get", controller.GetMovieByID)      // 通过视频id获取电影
	}
}
