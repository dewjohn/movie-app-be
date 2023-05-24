package admin

import (
	"github.com/gin-gonic/gin"
	"movie-app/common"
	"movie-app/dto"
	"movie-app/model"
	"movie-app/response"
	"net/http"
)

func GetUserService(page, pageSize int) response.ResponseStruct {
	res := response.ResponseStruct{
		HttpStatus: http.StatusOK,
		Code:       response.SuccessCode,
		Data:       nil,
		Msg:        response.OK,
	}

	var total int64
	var user []dto.UserInfoToAdminDto

	DB := common.GetDB()

	DB.Model(&model.User{}).Count(&total)

	Pagination := DB.Limit(pageSize).Offset((page - 1) * pageSize)

	Pagination.Model(&model.User{}).
		Select("id, avatar, name, email, telephone, gender, birthday, sign, state").
		Order("created_at desc").Scan(&user)

	res.Data = gin.H{"count": total, "user": user}
	return res
}

func ChangeUserStateService(stateDto dto.UserStateDto) response.ResponseStruct {
	DB := common.GetDB()
	res := response.ResponseStruct{
		HttpStatus: http.StatusOK,
		Code:       response.SuccessCode,
		Data:       nil,
		Msg:        response.OK,
	}
	err := DB.Model(model.User{}).Where("id = ?", stateDto.Uid).Update("state", stateDto.State).Error
	if err != nil {
		res.HttpStatus = http.StatusInternalServerError
		res.Code = response.ServerErrorCode
		res.Msg = response.SystemError
	}
	return res
}
