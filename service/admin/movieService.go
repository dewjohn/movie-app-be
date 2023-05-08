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
		Code:       response.SuccessCode,
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
		Code:       response.SuccessCode,
		Data:       nil,
		Msg:        response.OK,
	}
	DB := common.GetDB()
	var total int64 // 记录总数
	var movie []model.Movie
	DB.Model(&model.Movie{}).Count(&total)
	Pagination := DB.Limit(query.PageSize).Offset((query.Page - 1) * query.PageSize)

	Pagination.Model(&model.Movie{}).Order("created_at desc").Scan(&movie)
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
		Code:       response.SuccessCode,
		Data:       nil,
		Msg:        response.OK,
	}
	var movie model.Movie
	DB := common.GetDB()
	DB.Model(&model.Movie{}).Where("id = ?", vid).First(&movie)
	if movie.ID == 0 {
		res.HttpStatus = http.StatusUnprocessableEntity
		res.Code = response.CheckFailCode
		res.Msg = response.MovieNotExit
		return res
	}
	// 获取当前视频的resource
	resource := service.GetVideoResource(DB, uint(vid))
	movie.Videos = resource
	res.Data = gin.H{"movie": vo.ToAdminMovieById(movie)}
	return res
}

func UploadVideoByUrlService(request dto.UploadVideoByUrlDto) response.ResponseStruct {
	res := response.ResponseStruct{
		HttpStatus: http.StatusOK,
		Code:       2000,
		Data:       nil,
		Msg:        response.OK,
	}
	var movie model.Movie
	vid := request.Vid
	url := request.Url
	DB := common.GetDB()
	DB.Where("id = ?", vid).First(&movie)
	if movie.ID == 0 {
		res.HttpStatus = http.StatusUnprocessableEntity
		res.Code = response.CheckFailCode
		res.Msg = response.MovieNotExit
	}
	tx := DB.Begin()
	var err error
	var newResource model.Resource
	newResource.Vid = vid
	newResource.Title = "外链视频"
	newResource.Original = url

	// 创建新的资源
	if err = tx.Model(&model.Resource{}).Create(&newResource).Error; err != nil {
		tx.Rollback()
		res.HttpStatus = http.StatusInternalServerError
		res.Code = response.ServerErrorCode
		res.Msg = response.FailUploadFile
		return res
	}
	tx.Commit()
	return res
}
