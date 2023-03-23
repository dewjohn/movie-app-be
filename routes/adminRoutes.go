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
			adminAuth.GET("info", admin2.AdminInfo)                       // 获取管理员信息
			adminAuth.POST("/addAdmin", admin2.AddAdmin)                  // 新增管理员
			adminAuth.POST("/movie/upload/intro", admin2.UploadVideoInfo) // 1. 上传视频信息 2. 拿到Vid传给 upload/video接口
			adminAuth.PUT("/movie/modify", admin2.ModifyMovieInfo)        // 管理员修改电影信息
			adminAuth.PUT("/movie/delete/intro", admin2.DeleteMovieVideo) // 管理员删除电影信息
			adminAuth.PUT("/movie/delete/video", admin2.DeleteResource)   // 管理员删除电影视频
			adminAuth.GET("/movie/get", admin2.GetMovieDataList)          // 管理员获取电影所有信息
			adminAuth.GET("/movie/id", admin2.GetMovieByVid)              // 管理员获取指定Id电影所有信息
		}
	}
}
