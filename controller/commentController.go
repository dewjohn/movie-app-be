package controller

import (
	"github.com/gin-gonic/gin"
	"movie-app/dto"
	"movie-app/response"
	"movie-app/service"
	"strconv"
)

func Comment(ctx *gin.Context) {
	var comment dto.CommentDto
	err := ctx.Bind(&comment)
	if err != nil {
		response.Fail(ctx, nil, "请求错误")
		return
	}
	content := comment.Content
	uid, _ := ctx.Get("userId")

	if len(content) == 0 {
		response.CheckFail(ctx, nil, "评论或回复内容不能为空")
		return
	}
	res := service.CommentService(comment, uid)
	response.HandleResponse(ctx, res)
}

func GetComment(ctx *gin.Context) {
	page, _ := strconv.Atoi(ctx.Query("page"))
	pageSize, _ := strconv.Atoi(ctx.Query("page_size"))
	vid, _ := strconv.Atoi(ctx.Query("vid"))

	if page <= 0 || pageSize <= 0 {
		response.CheckFail(ctx, nil, "页码或请求数量错误")
		return
	}
	if pageSize > 15 {
		response.CheckFail(ctx, nil, "请求数量过多")
		return
	}

	res := service.GetCommentService(page, pageSize, vid)
	response.HandleResponse(ctx, res)
}

func GetReplyDetails(ctx *gin.Context) {
	//获取分页信息
	page, _ := strconv.Atoi(ctx.Query("page"))
	pageSize, _ := strconv.Atoi(ctx.Query("page_size"))
	cid, _ := strconv.Atoi(ctx.Query("cid"))
	if cid <= 0 {
		response.CheckFail(ctx, nil, "参数错误")
		return
	}
	if page <= 0 || pageSize <= 0 {
		response.CheckFail(ctx, nil, "页码或请求数量错误")
		return
	}
	if pageSize > 15 {
		response.CheckFail(ctx, nil, "请求数量过多")
		return
	}

	res := service.GetReplyDetailsV2Service(cid, page, pageSize)
	response.HandleResponse(ctx, res)
}

func Reply(ctx *gin.Context) {
	var reply dto.ReplyDto
	err := ctx.Bind(&reply)
	if err != nil {
		response.CheckFail(ctx, nil, "请求错误")
		return
	}
	cid := reply.Cid
	content := reply.Content
	uid, _ := ctx.Get("userId")

	if cid == 0 {
		response.CheckFail(ctx, nil, "评论不存在或已被删除")
		return
	}
	if len(content) == 0 {
		response.CheckFail(ctx, nil, "评论或回复内容不能为空")
		return
	}
	res := service.ReplyService(reply, uid)
	response.HandleResponse(ctx, res)
}

func DeleteComment(ctx *gin.Context) {
	var request dto.CommentIdDto
	err := ctx.Bind(&request)
	if err != nil {
		response.CheckFail(ctx, nil, "请求错误")
		return
	}
	id := request.ID
	uid, _ := ctx.Get("userId")
	res := service.DeleteCommentService(id, uid)
	response.HandleResponse(ctx, res)
}

func DeleteReply(ctx *gin.Context) {
	var request dto.ReplyIdDto
	err := ctx.Bind(&request)
	if err != nil {
		response.CheckFail(ctx, nil, "请求错误")
		return
	}
	id := request.ID
	uid, _ := ctx.Get("userId")
	res := service.DeleteReplyService(id, uid)
	response.HandleResponse(ctx, res)
}
