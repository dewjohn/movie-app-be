package admin

import (
	"github.com/gin-gonic/gin"
	"movie-app/response"
	"movie-app/service/admin"
)

func CountComment(ctx *gin.Context) {
	res := admin.CountCommentService()
	response.HandleResponse(ctx, res)
}

func CountReply(ctx *gin.Context) {
	res := admin.CountReplyService()
	response.HandleResponse(ctx, res)
}

func CountUser(ctx *gin.Context) {
	res := admin.CountUserService()
	response.HandleResponse(ctx, res)
}

func CountMovie(ctx *gin.Context) {
	res := admin.CountMovieService()
	response.HandleResponse(ctx, res)
}
