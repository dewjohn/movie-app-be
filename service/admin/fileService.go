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

func UploadCoverService(objectName string) response.ResponseStruct {
	res := response.ResponseStruct{
		HttpStatus: http.StatusOK,
		Code:       200,
		Data:       nil,
		Msg:        response.OK,
	}
	url := utils.GetUrl() + objectName
	res.Data = gin.H{"url": url}
	return res
}

func UploadVideoService(urls dto.ResDto, vid int, videoTitle string) response.ResponseStruct {
	res := response.ResponseStruct{
		HttpStatus: http.StatusOK,
		Code:       200,
		Data:       nil,
		Msg:        response.OK,
	}
	var movie model.Movie
	DB := common.GetDB()
	DB.Where("id = ?", vid).First(&movie)
	if movie.ID == 0 {
		res.HttpStatus = http.StatusBadRequest
		res.Code = 400
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
		res.HttpStatus = http.StatusBadRequest
		res.Code = 400
		res.Msg = response.FailUploadFile
		return res
	}
	tx.Commit()
	return res
}

func DeleteResourceService(uuid uuid.UUID) response.ResponseStruct {
	res := response.ResponseStruct{
		HttpStatus: http.StatusOK,
		Code:       200,
		Data:       nil,
		Msg:        response.OK,
	}
	DB := common.GetDB()
	DB.Where("uuid = ?", uuid).Delete(&model.Resource{})
	return res
}
