package admin

import (
	"movie-app/common"
	"movie-app/model"
	"movie-app/response"
	"net/http"
)

func DeleteMovieVideoService(vid uint) response.ResponseStruct {
	res := response.ResponseStruct{
		HttpStatus: http.StatusOK,
		Code:       200,
		Data:       nil,
		Msg:        response.OK,
	}
	DB := common.GetDB()
	DB.Where("id = ?", vid).Delete(&model.Movie{})
	return res
}
