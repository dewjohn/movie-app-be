package admin

import (
	"github.com/gin-gonic/gin"
	"movie-app/common"
	"movie-app/dto"
	"movie-app/model"
	"movie-app/response"
	"movie-app/service"
	"movie-app/vo"
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

func GetMovieDataListService(query dto.GetMovieListDto) response.ResponseStruct {
	res := response.ResponseStruct{
		HttpStatus: http.StatusOK,
		Code:       200,
		Data:       nil,
		Msg:        response.OK,
	}
	DB := common.GetDB()
	var total int64 // 记录总数
	var movie []model.Movie
	Pagination := DB.Limit(query.PageSize).Offset((query.Page - 1) * query.PageSize)

	Pagination.Model(&model.Movie{}).Order("created_at desc").Scan(&movie).Count(&total)
	// 获取当前视频的resource
	for i := 0; i < len(movie); i++ {
		resource := service.GetVideoResource(DB, uint(movie[i].ID))
		movie[i].Videos = vo.ToResource(resource)
	}
	res.Data = gin.H{"count": total, "movies": vo.ToAdminMovie(movie)}

	return res
}

func GetMovieByVidService(vid int) response.ResponseStruct {
	res := response.ResponseStruct{
		HttpStatus: http.StatusOK,
		Code:       http.StatusOK,
		Data:       nil,
		Msg:        response.OK,
	}
	var movie model.Movie
	DB := common.GetDB()
	DB.Model(&model.Movie{}).Where("id = ?", vid).First(&movie)
	if movie.ID == 0 {
		res.HttpStatus = http.StatusBadRequest
		res.Code = http.StatusBadRequest
		res.Msg = response.MovieNotExit
		return res
	}
	// 获取当前视频的resource
	resource := service.GetVideoResource(DB, uint(vid))
	movie.Videos = resource
	res.Data = gin.H{"movie": vo.ToAdminMovieById(movie)}
	return res
}
