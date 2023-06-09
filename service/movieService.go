package service

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"movie-app/common"
	"movie-app/dto"
	"movie-app/model"
	"movie-app/response"
	"movie-app/vo"
	"net/http"
	"time"
)

func UploadVideoInfoService(video dto.MovieDto, adminId interface{}, tReleaseTime time.Time) response.ResponseStruct {
	res := response.ResponseStruct{
		HttpStatus: http.StatusOK,
		Code:       response.SuccessCode,
		Data:       nil,
		Msg:        response.OK,
	}

	newVideo := model.Movie{
		AdminId:      adminId.(uint),
		Title:        video.Title,
		Cover:        video.Cover,
		ReleaseTime:  tReleaseTime,
		SheetLength:  video.SheetLength,
		Origin:       video.Origin,
		Type:         video.Type,
		Director:     video.Director,
		Screenwriter: video.Screenwriter,
		Actors:       video.Actors,
		Language:     video.Language,
		Introduction: video.Introduction,
		Score:        video.Score,
	}
	DB := common.GetDB()
	DB.Create(&newVideo)
	res.Data = gin.H{"vid": newVideo.ID}
	return res
}

func ModifyMovieInfoService(vid int, video dto.ModifyMovieDto, tReleaseTime time.Time) response.ResponseStruct {
	res := response.ResponseStruct{
		HttpStatus: http.StatusOK,
		Code:       response.SuccessCode,
		Data:       nil,
		Msg:        response.OK,
	}

	DB := common.GetDB()
	var movie model.Movie
	DB.Model(&model.Movie{}).Where("id = ?", vid).First(&movie)
	if movie.ID == 0 {
		res.HttpStatus = http.StatusBadRequest
		res.Code = response.CheckFailCode
		res.Msg = response.MovieNotExit
		return res
	}
	if err := DB.Model(&model.Movie{}).Where("id = ?", vid).Updates(
		map[string]interface{}{
			"title":        video.Title,
			"cover":        video.Cover,
			"release_time": tReleaseTime,
			"sheet_length": video.SheetLength,
			"origin":       video.Origin,
			"type":         video.Type,
			"director":     video.Director,
			"screenwriter": video.Screenwriter,
			"actors":       video.Actors,
			"language":     video.Language,
			"introduction": video.Introduction,
			"score":        video.Score,
		}).Error; err != nil {
		res.HttpStatus = http.StatusInternalServerError
		res.Code = response.ServerErrorCode
		res.Msg = response.SystemError
	}
	return res
}

func GetMovieListService(query dto.GetMovieListDto) response.ResponseStruct {
	res := response.ResponseStruct{
		HttpStatus: http.StatusOK,
		Code:       response.SuccessCode,
		Data:       nil,
		Msg:        response.OK,
	}
	DB := common.GetDB()
	var total int64 // 记录总数
	var movie []vo.SearchMovieVo
	Pagination := DB.Limit(query.PageSize).Offset((query.Page - 1) * query.PageSize)

	Pagination.Model(&model.Movie{}).Select("id, title, cover, release_time, score").Order("created_at desc").Scan(&movie).Count(&total)

	res.Data = gin.H{"count": total, "movies": movie}

	return res
}

func GetMovieByIdService(vid int) response.ResponseStruct {
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
		res.HttpStatus = http.StatusBadRequest
		res.Code = response.CheckFailCode
		res.Msg = response.MovieNotExit
		return res
	}
	// 获取当前视频的resource
	resource := GetVideoResource(DB, uint(vid))
	res.Data = gin.H{"movie": vo.ToVideo(movie, resource)}
	return res
}

func GetHighScoreMovieService(query dto.GetMovieListDto) response.ResponseStruct {
	res := response.ResponseStruct{
		HttpStatus: http.StatusOK,
		Code:       response.SuccessCode,
		Data:       nil,
		Msg:        response.OK,
	}
	DB := common.GetDB()
	var total int64 // 记录总数
	var movies []vo.SearchMovieVo
	Pagination := DB.Limit(query.PageSize).Offset((query.Page - 1) * query.PageSize)
	Pagination.Model(&model.Movie{}).Select("id, title, cover, release_time, score").Where("score > 9.0").Scan(&movies).Count(&total)

	res.Data = gin.H{
		"count":  total,
		"movies": movies,
	}
	return res
}

func IsMovieExit(db *gorm.DB, Vid uint) bool {
	var movie model.Movie
	db.First(&movie, Vid)
	if movie.ID != 0 {
		return true
	}
	return false
}

func GetVideoResource(db *gorm.DB, vid uint) []model.Resource {
	var resource []model.Resource
	db.Model(&model.Resource{}).Where("vid = ?", vid).Find(&resource)
	return resource
}
