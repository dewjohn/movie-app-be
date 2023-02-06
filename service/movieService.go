package service

import (
	"github.com/gin-gonic/gin"
	"movie-app/common"
	"movie-app/dto"
	"movie-app/model"
	"movie-app/response"
	"movie-app/vo"
	"net/http"
	"time"
)

func UploadVideoInfoService(video dto.VideoDto, adminId interface{}, tReleaseTime time.Time) response.ResponseStruct {
	res := response.ResponseStruct{
		HttpStatus: http.StatusOK,
		Code:       200,
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
	}
	DB := common.GetDB()
	DB.Create(&newVideo)
	res.Data = gin.H{"vid": newVideo.ID}
	return res
}

func ModifyVideoInfoService(video dto.ModifyVideoDto, tReleaseTime time.Time) response.ResponseStruct {
	res := response.ResponseStruct{
		HttpStatus: http.StatusOK,
		Code:       200,
		Data:       nil,
		Msg:        response.OK,
	}

	DB := common.GetDB()
	err := DB.Model(&model.Movie{}).Where("id = ?", video.Vid).Updates(
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
		}).Error
	if err != nil {
		res.HttpStatus = http.StatusInternalServerError
		res.Code = 500
		res.Msg = response.SystemError
	}
	return res
}

func GetMovieListService(query dto.GetMovieListDto) response.ResponseStruct {
	res := response.ResponseStruct{
		HttpStatus: http.StatusOK,
		Code:       200,
		Data:       nil,
		Msg:        response.OK,
	}
	DB := common.GetDB()
	var total int64 // 记录总数
	var movie []vo.SearchMovieVo
	Pagination := DB.Limit(query.PageSize).Offset((query.Page - 1) * query.PageSize)

	Pagination.Model(&model.Movie{}).Select("id, title, cover").Scan(&movie).Count(&total)

	res.Data = gin.H{"count": total, "movies": movie}

	return res
}
