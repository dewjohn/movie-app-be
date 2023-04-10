package admin

import (
	"github.com/gin-gonic/gin"
	"movie-app/common"
	"movie-app/model"
	"movie-app/response"
	"net/http"
)

func CountCommentService() response.ResponseStruct {
	res := response.ResponseStruct{
		HttpStatus: http.StatusOK,
		Code:       response.SuccessCode,
		Data:       nil,
		Msg:        response.OK,
	}
	DB := common.GetDB()
	var count int64
	DB.Model(&model.Comment{}).Count(&count)
	res.Data = gin.H{"count": count}
	return res
}

func CountReplyService() response.ResponseStruct {
	res := response.ResponseStruct{
		HttpStatus: http.StatusOK,
		Code:       response.SuccessCode,
		Data:       nil,
		Msg:        response.OK,
	}
	DB := common.GetDB()
	var count int64
	DB.Model(&model.Comment{}).Where("deleted_at is null and parent_id > 0").Count(&count)
	res.Data = gin.H{"count": count}
	return res
}

func CountUserService() response.ResponseStruct {
	res := response.ResponseStruct{
		HttpStatus: http.StatusOK,
		Code:       response.SuccessCode,
		Data:       nil,
		Msg:        response.OK,
	}
	DB := common.GetDB()
	var count int64
	DB.Model(&model.User{}).Where("deleted_at is null").Count(&count)
	res.Data = gin.H{"count": count}
	return res
}

func CountMovieService() response.ResponseStruct {
	res := response.ResponseStruct{
		HttpStatus: http.StatusOK,
		Code:       response.SuccessCode,
		Data:       nil,
		Msg:        response.OK,
	}
	DB := common.GetDB()
	var count int64
	DB.Model(&model.Movie{}).Where("deleted_at is null").Count(&count)
	res.Data = gin.H{"count": count}
	return res
}
