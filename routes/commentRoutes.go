package routes

import (
	"github.com/gin-gonic/gin"
	"movie-app/controller"
	"movie-app/middleWare"
)

func GetCommentRoutes(route *gin.RouterGroup) {
	comment := route.Group("/comment")
	{
		comment.GET("/list", controller.GetCommentList)        // 获取评论列表
		comment.GET("/reply/list", controller.GetReplyDetails) // 获取回复列表
		userAuth := comment.Group("/")
		userAuth.Use(middleWare.UserAuthMiddleWare())
		{
			userAuth.POST("/add", controller.Comment)          // 写评论
			userAuth.POST("/delete", controller.DeleteComment) // 删除评论
		}
	}
}
