package controller

import (
	"github.com/gin-gonic/gin"
	"movie-app/response"
	"movie-app/service"
	"movie-app/utils"
)

func Collect(ctx *gin.Context) {
	vid := utils.StringToInt(ctx.Query("vid"))
	if vid <= 0 {
		response.CheckFail(ctx, nil, "电影不存在")
		return
	}
	uid, _ := ctx.Get("userId")
	res := service.CollectService(vid, uid)
	response.HandleResponse(ctx, res)
}

func DeleteCollect(ctx *gin.Context) {
	vid := utils.StringToInt(ctx.Query("vid"))
	if vid <= 0 {
		response.CheckFail(ctx, nil, "电影不存在")
		return
	}
	uid, _ := ctx.Get("userId")
	res := service.DeleteCollectService(vid, uid)
	response.HandleResponse(ctx, res)
}

func GetCollectList(ctx *gin.Context) {
	page := utils.StringToInt(ctx.Query("page"))
	pageSize := utils.StringToInt(ctx.Query("page_size"))
	uid, _ := ctx.Get("userId")
	res := service.GetCollectService(page, pageSize, uid)
	response.HandleResponse(ctx, res)
}

func IsCollected(ctx *gin.Context) {
	vid := utils.StringToInt(ctx.Query("vid"))
	uid, _ := ctx.Get("userId")
	res := service.IsCollectedService(vid, uid)
	response.HandleResponse(ctx, res)
}
