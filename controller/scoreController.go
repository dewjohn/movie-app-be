package controller

import (
	"github.com/gin-gonic/gin"
	"movie-app/dto"
	"movie-app/response"
	"movie-app/service"
	"movie-app/utils"
)

func ReviewMovieScore(ctx *gin.Context) {
	var request = dto.ScoreDto{}
	err := ctx.Bind(&request)
	if err != nil {
		response.Fail(ctx, nil, response.RequestError)
		return
	}

	vid := request.Vid
	grade := request.Grade
	uid, _ := ctx.Get("userId")

	if vid == 0 {
		response.CheckFail(ctx, nil, response.MovieNotExit)
		return
	}
	if grade > 10.0 || grade < 0 {
		response.CheckFail(ctx, nil, "分数范围为1-10分")
		return
	}
	res := service.ReviewMovieScoreService(request, uid)
	response.HandleResponse(ctx, res)
}

func GetMovieScoreAvg(ctx *gin.Context) {
	vid := utils.StringToInt(ctx.Query("vid"))
	res := service.GetMovieScoreAvgService(vid)
	response.HandleResponse(ctx, res)
}
