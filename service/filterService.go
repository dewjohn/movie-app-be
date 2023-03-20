package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"movie-app/common"
	"movie-app/dto"
	"movie-app/model"
	"movie-app/response"
	"movie-app/vo"
	"net/http"
)

func FilterService(request dto.FilterMovieDto) response.ResponseStruct {
	res := response.ResponseStruct{
		HttpStatus: http.StatusOK,
		Code:       200,
		Data:       nil,
		Msg:        response.OK,
	}
	DB := common.GetDB()
	var total int64 // 记录总数
	var movie []vo.SearchMovieVo
	Pagination := DB.Limit(request.PageSize).Offset((request.Page - 1) * request.PageSize)

	search := fmt.Sprintf("%s like ?", request.Category)

	Pagination.Model(&model.Movie{}).
		Select("id, title, cover, release_time, score").
		Where(search, request.Types).Scan(&movie).Count(&total)

	res.Data = gin.H{"count": total, "movies": movie}

	return res
}