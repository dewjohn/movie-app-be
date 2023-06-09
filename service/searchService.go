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
	var total int64
	DB.Model(model.Movie{}).Select("id,title,cover").Where("concat(title, origin, type, director, actors, language) like ?", keywords).Scan(&movies).Count(&total)
	return response.ResponseStruct{
		HttpStatus: http.StatusOK,
		Code:       response.SuccessCode,
		Data:       gin.H{"movies": movies, "count": total},
		Msg:        response.OK,
	}
}
