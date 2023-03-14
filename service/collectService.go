package service

import (
	"github.com/gin-gonic/gin"
	"movie-app/common"
	"movie-app/dto"
	"movie-app/model"
	"movie-app/response"
	"net/http"
)

func CollectService(collect dto.CollectDto, uid interface{}) response.ResponseStruct {
	res := response.ResponseStruct{
		HttpStatus: http.StatusOK,
		Code:       200,
		Data:       nil,
		Msg:        response.OK,
	}
	DB := common.GetDB()
	newCollect := model.Collect{
		Vid: collect.Vid,
		Uid: uid.(uint),
	}
	DB.Create(&newCollect)
	res.Data = gin.H{"id": newCollect.ID}
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
