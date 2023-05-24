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
		admin.GET("/token/refresh", middleWare.RefreshAdminTokenMiddleWare(), admin2.GetAdminAccessToken) // 刷新 accesstoken
		// 需要管理员登陆
		adminAuth := admin.Group("/")
		adminAuth.Use(middleWare.AdminAuthMiddleWare())
		{
			adminAuth.GET("info", admin2.AdminInfo)                                           // 获取管理员信息
			adminAuth.POST("/addAdmin", admin2.AddAdmin)                                      // 新增管理员
			adminAuth.GET("/get", admin2.GetAdmin)                                            // 获取管理员列表
			adminAuth.POST("/auth", admin2.ChangeAdminAuthorization)                          // 修改管理员权限
			adminAuth.POST("/movie/upload/intro", admin2.UploadVideoInfo)                     // 1. 上传视频信息 2. 拿到Vid传给 upload/video接口
			adminAuth.PUT("/movie/modify", admin2.ModifyMovieInfo)                            // 管理员修改电影信息
			adminAuth.PUT("/movie/delete", admin2.DeleteMovie)                                // 管理员删除电影信息
			adminAuth.PUT("/movie/delete/video", admin2.DeleteResource)                       // 管理员删除电影视频
			adminAuth.GET("/movie/get", admin2.GetMovieDataList)                              // 管理员获取电影所有信息
			adminAuth.GET("/movie/id", admin2.GetMovieByVid)                                  // 管理员获取指定Id电影所有信息
			adminAuth.GET("/user/get", admin2.GetUser)                                        // 管理员获取所有用户信息
			adminAuth.GET("/user/search", admin2.SearchUser)                                  // 管理员搜索用户
			adminAuth.PUT("/user/state", admin2.ChangeUserState)                              // 管理员修改用户状态
			adminAuth.GET("/statistics/all", admin2.StatisticsAllData)                        // 统计整站数据 			// 统计电影数量
			adminAuth.GET("/statistics/recentMovie", admin2.StatisticsUploadMovieRecentMonth) // 统计近三个月上传电影数
			adminAuth.GET("/statistics/releaseMovie", admin2.StatisticsMovieRelease)          // 统计电影年代
			adminAuth.GET("/statistics/typeMovie", admin2.StatisticsMovieType)                // 统计电影类型
		}
	}
}
