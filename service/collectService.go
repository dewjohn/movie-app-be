package service

import (
	"github.com/gin-gonic/gin"
	"movie-app/common"
	"movie-app/dto"
	"movie-app/model"
	"movie-app/response"
	"movie-app/utils"
	"net/http"
)

func CollectService(vid int, uid interface{}) response.ResponseStruct {
	res := response.ResponseStruct{
		HttpStatus: http.StatusOK,
		Code:       200,
		Data:       nil,
		Msg:        response.OK,
	}
	DB := common.GetDB()
	if utils.IsCollected(DB, vid, uid.(uint)) {
		res.HttpStatus = http.StatusUnprocessableEntity
		res.Code = 400
		res.Msg = "禁止重复收藏"
		return res
	}
	newCollect := model.Collect{
		Vid: vid,
		Uid: uid.(uint),
	}
	DB.Create(&newCollect)
	res.Data = gin.H{"id": newCollect.ID}
	return res
}

func DeleteCollectService(vid int, uid interface{}) response.ResponseStruct {
	res := response.ResponseStruct{
		HttpStatus: http.StatusOK,
		Code:       400,
		Data:       nil,
		Msg:        response.RequestError,
	}
	DB := common.GetDB()
	var collect model.Collect
	DB.Where("vid = ? and uid = ?", vid, uid).First(&collect)
	if collect.ID != 0 {
		if collect.Uid == uid {
			DB.Where("vid = ? and uid = ?", vid, uid).Delete(&collect)
			res.HttpStatus = http.StatusOK
			res.Code = 200
			res.Msg = response.OK
		}
	}
	return res
}

func GetCollectService(page, pageSize int, uid interface{}) response.ResponseStruct {

	var total int64
	var collects []dto.CollectResDto

	DB := common.GetDB()
	DB.Model(model.Collect{}).Where("uid = ?", uid).Count(&total)
	sql := "select vid from collects where deleted_at is null and uid = ? limit ? offset ?"
	DB.Raw(sql, uid, pageSize, (page-1)*pageSize).Scan(&collects)
	sqlMovie := "select title, cover, release_time, score from movies where deleted_at is null and id = ?"
	for i := 0; i < len(collects); i++ {
		DB.Raw(sqlMovie, collects[i].Vid).Scan(&collects[i].MovieInfo)
	}
	return response.ResponseStruct{
		HttpStatus: http.StatusOK,
		Code:       http.StatusOK,
		Data:       gin.H{"count": total, "collections": collects},
		Msg:        response.OK,
	}
}

func IsCollectedService(vid int, uid interface{}) response.ResponseStruct {
	res := response.ResponseStruct{
		HttpStatus: http.StatusOK,
		Code:       http.StatusOK,
		Data:       nil,
		Msg:        response.OK,
	}
	var collect model.Collect
	DB := common.GetDB()
	DB.Where("vid = ? and uid = ?", vid, uid).First(&collect)
	if collect.ID != 0 {
		res.Data = gin.H{"isCollected": true}
		return res
	}
	res.Data = gin.H{"isCollected": false}
	return res
}
