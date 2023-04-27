package admin

import (
	"github.com/gin-gonic/gin"
	"movie-app/common"
	"movie-app/dto"
	"movie-app/model"
	"movie-app/response"
	"net/http"
)

func SearchUserService(uname string) response.ResponseStruct {
	var user []dto.UserInfoToAdminDto
	DB := common.GetDB()
	DB = DB.Limit(30)
	var total int64
	DB.Model(model.User{}).Select("id, avatar, name, email, telephone, gender, birthday, sign, state").Where("name like ?", uname).Scan(&user).Count(&total)
	return response.ResponseStruct{
		HttpStatus: http.StatusOK,
		Code:       response.SuccessCode,
		Data:       gin.H{"user": user, "count": total},
		Msg:        response.OK,
	}
}
