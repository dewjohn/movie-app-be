package service

import (
	"github.com/gin-gonic/gin"
	"movie-app/common"
	"movie-app/model"
	"movie-app/response"
	"movie-app/vo"
	"net/http"
)

func SearchMovieService(keywords string) response.ResponseStruct {
	var movies []vo.SearchMovieVo
	DB := common.GetDB()
	DB = DB.Limit(30)
	DB.Model(model.Movie{}).Select("id,title,cover").Where("title like ?", keywords).Scan(&movies)
	return response.ResponseStruct{
		HttpStatus: http.StatusOK,
		Code:       http.StatusOK,
		Data:       gin.H{"movies": movies},
		Msg:        response.OK,
	}
}
