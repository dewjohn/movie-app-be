package service

import (
	"movie-app/common"
	"movie-app/dto"
	"movie-app/model"
	"movie-app/response"
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
	return res
}
