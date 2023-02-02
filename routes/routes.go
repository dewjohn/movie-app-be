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
