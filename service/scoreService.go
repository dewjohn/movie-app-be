package service

import (
	"github.com/gin-gonic/gin"
	"movie-app/common"
	"movie-app/dto"
	"movie-app/model"
	"movie-app/response"
	"movie-app/utils"
	"net/http"
)

func ReviewMovieScoreService(score dto.ScoreDto, uid interface{}) response.ResponseStruct {
	res := response.ResponseStruct{
		HttpStatus: http.StatusOK,
		Code:       http.StatusOK,
		Data:       nil,
		Msg:        response.OK,
	}

	DB := common.GetDB()

	// 判断对该部电影是否评论过
	if utils.IsReviewedMovie(DB, score.Vid, uid) {
		res.HttpStatus = http.StatusUnprocessableEntity
		res.Code = 422
		res.Msg = response.ReviewScoreExit
		return res
	}
	newReviewScore := model.Score{
		Vid:   score.Vid,
		Grade: score.Grade,
		Uid:   uid.(uint),
	}
	DB.Create(&newReviewScore)
	res.Data = gin.H{"id": newReviewScore.ID}
	return res
}

func GetMovieScoreAvgService(vid int) response.ResponseStruct {
	res := response.ResponseStruct{
		HttpStatus: http.StatusOK,
		Code:       http.StatusOK,
		Data:       nil,
		Msg:        response.OK,
	}
	var average float64
	var reviewers int64
	DB := common.GetDB()
	DB.Model(model.Score{}).Select("avg(grade)").Where("vid = ?", vid).Find(&average).Count(&reviewers)
	res.Data = gin.H{
		"average":   average,
		"reviewers": reviewers,
	}
	return res
}
