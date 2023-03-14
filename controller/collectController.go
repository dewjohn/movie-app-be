package controller

import (
	"github.com/gin-gonic/gin"
	"movie-app/dto"
	"movie-app/response"
	"movie-app/service"
	"movie-app/utils"
)

func Collect(ctx *gin.Context) {
	var collect dto.CollectDto
	err := ctx.Bind(&collect)
	if err != nil {
		response.Fail(ctx, nil, response.RequestError)
		return
	}

	vid := collect.Vid
	if vid <= 0 {
		response.CheckFail(ctx, nil, "电影不存在")
		return
	}
	uid, _ := ctx.Get("userId")
	res := service.CollectService(collect, uid)
	response.HandleResponse(ctx, res)
}

func GetCollectList(ctx *gin.Context) {
	page := utils.StringToInt(ctx.Query("page"))
	pageSize := utils.StringToInt(ctx.Query("page_size"))
	uid, _ := ctx.Get("userId")
	res := service.GetCollectService(page, pageSize, uid)
	response.HandleResponse(ctx, res)
}
