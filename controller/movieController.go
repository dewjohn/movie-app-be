package controller

import (
	"github.com/gin-gonic/gin"
	"movie-app/dto"
	"movie-app/response"
	"movie-app/service"
	"strconv"
)

func GetMovieList(ctx *gin.Context) {
	var request dto.GetMovieListDto
	request.Page, _ = strconv.Atoi(ctx.Query("page"))
	request.PageSize, _ = strconv.Atoi(ctx.Query("page_size"))
	if request.Page <= 0 || request.PageSize <= 0 {
		response.Fail(ctx, nil, response.PageError)
		return
	}
	if request.PageSize >= 30 {
		response.Fail(ctx, nil, response.RequestTooMany)
		return
	}
	res := service.GetMovieListService(request)
	response.HandleResponse(ctx, res)
}

func GetHighScoreMovie(ctx *gin.Context) {
	var request dto.GetMovieListDto
	request.Page, _ = strconv.Atoi(ctx.Query("page"))
	request.PageSize, _ = strconv.Atoi(ctx.Query("page_size"))
	if request.Page <= 0 || request.PageSize <= 0 {
		response.Fail(ctx, nil, response.PageError)
		return
	}
	if request.PageSize >= 30 {
		response.Fail(ctx, nil, response.RequestTooMany)
		return
	}
	res := service.GetHighScoreMovieService(request)
	response.HandleResponse(ctx, res)
}

func GetMovieByID(ctx *gin.Context) {
	vid, _ := strconv.Atoi(ctx.Query("vid"))
	if vid == 0 {
		response.CheckFail(ctx, nil, response.MovieNotExit)
		return
	}
	res := service.GetMovieByIdService(vid)
	response.HandleResponse(ctx, res)
}
