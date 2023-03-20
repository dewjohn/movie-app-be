package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"movie-app/dto"
	"movie-app/response"
	"movie-app/service"
	"strconv"
)

func FilterMovie(ctx *gin.Context) {
	var request dto.FilterMovieDto
	request.Page, _ = strconv.Atoi(ctx.Query("page"))
	request.PageSize, _ = strconv.Atoi(ctx.Query("page_size"))
	request.Category = ctx.Query("category")
	request.Types = ctx.Query("types")

	if request.Page <= 0 || request.PageSize <= 0 {
		response.Fail(ctx, nil, response.PageError)
		return
	}
	if request.PageSize >= 30 {
		response.Fail(ctx, nil, response.RequestTooMany)
		return
	}

	tips := fmt.Sprintf("电影%s不能为空", request.Category)
	if len(request.Types) == 0 {
		response.CheckFail(ctx, nil, tips)
		return
	}
	res := service.FilterService(request)
	response.HandleResponse(ctx, res)
}
