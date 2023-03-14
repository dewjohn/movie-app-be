package controller

import (
	"github.com/gin-gonic/gin"
	"movie-app/dto"
	"movie-app/response"
	"movie-app/service"
	"movie-app/utils"
)

func Comment(ctx *gin.Context) {
	var comment dto.CommentDto
	err := ctx.Bind(&comment)
	if err != nil {
		response.Fail(ctx, nil, response.RequestError)
		return
	}
	content := comment.Content
	uid, _ := ctx.Get("userId")

	if len(content) == 0 {
		response.CheckFail(ctx, nil, response.CommentOrReplyError)
		return
	}
	res := service.CommentService(comment, uid)
	response.HandleResponse(ctx, res)
}

func GetCommentAndReply(ctx *gin.Context) {
	page := utils.StringToInt(ctx.Query("page"))
	pageSize := utils.StringToInt(ctx.Query("page_size"))
	vid := utils.StringToInt(ctx.Query("vid"))
	replyCount := utils.StringToInt(ctx.DefaultQuery("reply", "0"))

	if page <= 0 || pageSize <= 0 || replyCount > 15 || pageSize > 15 {
		response.CheckFail(ctx, nil, "页码或请求数量错误")
		return
	}

	res := service.GetCommentService(vid, replyCount, page, pageSize)
	response.HandleResponse(ctx, res)
}

func GetCommentList(ctx *gin.Context) {
	//获取分页信息
	page := utils.StringToInt(ctx.Query("page"))
	pageSize := utils.StringToInt(ctx.Query("page_size"))
	vid := utils.StringToInt(ctx.Query("vid"))

	if page <= 0 || pageSize <= 0 || pageSize > 15 {
		response.CheckFail(ctx, nil, response.PageError)
		return
	}

	res := service.GetCommentListService(vid, page, pageSize)
	response.HandleResponse(ctx, res)
}

func GetReplyDetails(ctx *gin.Context) {
	cid := utils.StringToInt(ctx.Query("cid"))
	page := utils.StringToInt(ctx.Query("page"))
	pageSize := utils.StringToInt(ctx.Query("page_size"))
	offset := utils.StringToInt(ctx.DefaultQuery("offset", "0"))

	if cid <= 0 {
		response.CheckFail(ctx, nil, response.ParameterError)
		return
	}
	if page <= 0 || pageSize <= 0 || pageSize > 15 {
		response.CheckFail(ctx, nil, response.PageError)
		return
	}

	res := service.GetReplyDetailsService(cid, offset, page, pageSize)
	response.HandleResponse(ctx, res)
}

func DeleteComment(ctx *gin.Context) {
	var request dto.CommentIdDto
	err := ctx.Bind(&request)
	if err != nil {
		response.CheckFail(ctx, nil, response.RequestError)
		return
	}
	id := request.ID
	uid, _ := ctx.Get("userId")
	res := service.DeleteCommentService(id, uid)
	response.HandleResponse(ctx, res)
}
