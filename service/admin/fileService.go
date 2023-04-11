package admin

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"movie-app/common"
	"movie-app/dto"
	"movie-app/model"
	"movie-app/response"
	"movie-app/utils"
	"net/http"
)

func UploadCoverService(objectName string, vid int) response.ResponseStruct {
	res := response.ResponseStruct{
		HttpStatus: http.StatusOK,
		Code:       response.SuccessCode,
		Data:       nil,
		Msg:        response.OK,
	}
	url := utils.GetUrl() + objectName
	if vid != 0 {
		var movie model.Movie
		DB := common.GetDB()
		DB.Where("id = ?", vid).First(&movie)
		if movie.ID == 0 {
			res.HttpStatus = http.StatusUnprocessableEntity
			res.Code = response.CheckFailCode
			res.Msg = response.MovieNotExit
			return res
		}
		err := DB.Model(&model.Movie{}).Where("id = ?", vid).Update("cover", url).Error
		if err != nil {
			res.HttpStatus = http.StatusInternalServerError
			res.Code = 500
			res.Msg = response.SystemError
		}
	}
	res.Data = gin.H{"url": url}
	return res
}

func UploadVideoService(urls dto.ResDto, vid int, videoTitle string) response.ResponseStruct {
	res := response.ResponseStruct{
		HttpStatus: http.StatusOK,
		Code:       2000,
		Data:       nil,
		Msg:        response.OK,
	}
	var movie model.Movie
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
	newResource.Vid = uint(vid)
	newResource.Title = videoTitle
	newResource.Original = urls.Original // 原始分辨率

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

func DeleteResourceService(uuid uuid.UUID) response.ResponseStruct {
	res := response.ResponseStruct{
		HttpStatus: http.StatusOK,
		Code:       response.SuccessCode,
		Data:       nil,
		Msg:        response.OK,
	}
	DB := common.GetDB()
	DB.Where("uuid = ?", uuid).Delete(&model.Resource{})
	return res
}
