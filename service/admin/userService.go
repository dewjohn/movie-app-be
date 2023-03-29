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
		Code:       200,
		Data:       nil,
		Msg:        response.OK,
	}

	var total int64
	var user []dto.UserInfoToAdminDto

	DB := common.GetDB()

	Pagination := DB.Limit(pageSize).Offset((page - 1) * pageSize)

	Pagination.Model(&model.User{}).
		Select("avatar, name, email, telephone, gender, birthday, sign, state").
		Order("created_at desc").Scan(&user).Count(&total)

	res.Data = gin.H{"count": total, "user": user}
	return res
}

func ChangeUserStateService(adminId interface{}, stateDto dto.UserStateDto) response.ResponseStruct {
	DB := common.GetDB()
	res := response.ResponseStruct{
		HttpStatus: http.StatusOK,
		Code:       http.StatusOK,
		Data:       nil,
		Msg:        response.OK,
	}

	var admin model.Admin
	DB.Where("id = ?", adminId).First(&admin)
	if admin.ID == 0 || admin.Authority < 2000 {
		res.HttpStatus = http.StatusUnprocessableEntity
		res.Code = 422
		res.Msg = "无该管理员或者权限不足"
		return res
	}
	err := DB.Model(model.User{}).Where("id = ?", stateDto.Uid).Update("state", stateDto.State).Error
	if err != nil {
		res.HttpStatus = http.StatusInternalServerError
		res.Code = 500
		res.Msg = response.SystemError
	}
	return res
}