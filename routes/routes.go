package routes

import (
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"movie-app/controller"
	"movie-app/controller/admin"
	"movie-app/middleWare"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CollectRoute(r *gin.Engine) *gin.Engine {
	r.Use(middleWare.CorsMiddleWare())
	v1 := r.Group("/api/v1")
	{
		GetUserRoutes(v1)
		GetAdminRoutes(v1)
		GetMovieRoutes(v1)

		// 评论
		v1.GET("comment/get", controller.GetComment)
		v1.GET("comment/reply", controller.GetReplyDetails)
		comment := v1.Group("/comment")
		comment.Use(middleWare.UserAuthMiddleWare())
		{
			comment.POST("", controller.Comment)                  // 评论
			comment.POST("/delete", controller.DeleteComment)     // 删除评论
			comment.POST("/reply", controller.Reply)              // 回复
			comment.POST("/reply/delete", controller.DeleteReply) // 删除回复
		}
		// 用户文件上传相关
		userFile := v1.Group("/upload")
		userFile.Use(middleWare.UserAuthMiddleWare())
		{
			userFile.POST("/avatar", controller.UploadAvatar) // 上传用户头像
		}

		// 管理员文件上传相关
		adminFile := v1.Group("/upload")
		adminFile.Use(middleWare.AdminAuthMiddleWare())
		{
			adminFile.POST("/cover", admin.UploadMovieCover)
			adminFile.POST("/video", admin.UploadMovieVideo)
		}
		//获取静态文件
		r.StaticFS("/api/avatar", http.Dir("./files/avatar"))
		r.StaticFS("/api/cover", http.Dir("./files/cover"))
		r.StaticFS("/api/movie", http.Dir("./files/movie"))

		// swagger
		r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}
	return r
}
