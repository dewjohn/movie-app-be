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
			adminAuth.POST("/addAdmin", admin2.AddAdmin)                  // 新增管理员
			adminAuth.POST("/movie/upload/info", admin2.UploadVideoInfo)  // 1. 上传视频信息 2. 拿到Vid传给 upload/video接口
			adminAuth.POST("/movie/modify/info", admin2.ModifyVideoInfo)  // 管理员修改电影信息
			adminAuth.POST("/movie/delete/info", admin2.DeleteMovieVideo) // 管理员删除电影信息
			adminAuth.POST("/movie/delete/video", admin2.DeleteResource)  // 管理员删除电影视频
		}
	}
}
