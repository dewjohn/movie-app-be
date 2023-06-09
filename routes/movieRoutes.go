package routes

import (
	"github.com/gin-gonic/gin"
	"movie-app/controller"
	"movie-app/middleWare"
)

func GetMovieRoutes(route *gin.RouterGroup) {
	movie := route.Group("/movie")
	{
		movie.GET("/get", controller.GetMovieList)           // 获取视频列表
		movie.GET("/high/get", controller.GetHighScoreMovie) // 获取高分电影
		movie.GET("/search", controller.SerchMovie)          // 搜索视频
		movie.GET("/filter", controller.FilterMovie)         // 分类搜索
		movie.GET("/id", controller.GetMovieByID)            // 通过视频id获取电影
		movie.GET("/score/get", controller.GetMovieScoreAvg)
		movieAuth := movie.Group("/")
		movieAuth.Use(middleWare.UserAuthMiddleWare())
		{
			movieAuth.POST("/score/review", controller.ReviewMovieScore)
		}
	}
}
