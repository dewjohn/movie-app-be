package service

import (
	"movie-app/common"
	"movie-app/model"
	"movie-app/response"
	"net/http"
)

func UploadAvatarService(localFileName string, objectName string, uid uint) response.ResponseStruct {
	res := response.ResponseStruct{
		HttpStatus: http.StatusOK,
		Code:       200,
		Data:       nil,
		Msg:        response.OK,
	}
	url := "/api/" + objectName
	DB := common.GetDB()
	DB.Model(model.User{}).Where("id = ?", uid).Update("avatar", url)
	return res
}
