package admin

import (
	"github.com/gin-gonic/gin"
	"movie-app/common"
	"movie-app/dto"
	"movie-app/model"
	"movie-app/response"
	"net/http"
)

func UploadCoverService(objectName string) response.ResponseStruct {
	res := response.ResponseStruct{
		HttpStatus: http.StatusOK,
		Code:       200,
		Data:       nil,
		Msg:        response.OK,
	}
	url := "/api/" + objectName
	res.Data = gin.H{"url": url}
	return res
}

func UploadVideoService(urls dto.ResDto, vid int) response.ResponseStruct {
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
		res.Msg = "视频不存在"
	}
	tx := DB.Begin()
	var err error
	var newResource model.Resource
	newResource.Vid = uint(vid)
	newResource.Original = urls.Original // 原始分辨率

	// 创建新的资源
	if err = tx.Model(&model.Resource{}).Create(&newResource).Error; err != nil {
		tx.Rollback()
		res.HttpStatus = http.StatusBadRequest
		res.Code = 400
		res.Msg = "文件上传失败"
		return res
	}
	tx.Commit()
	return res
}
