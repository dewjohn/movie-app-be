package routes

import (
	"movie-app/controller"
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

		// 文件上传相关
		file := v1.Group("/upload")
		file.Use(middleWare.UserAuthMiddleWare())
		{
			// 上传用户头像
			file.POST("/avatar", controller.UploadAvatar)
		}
		//获取静态文件
		r.StaticFS("/api/avatar", http.Dir("./files/avatar"))
	}
	return r
}
