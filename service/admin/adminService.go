package admin

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"log"
	"movie-app/common"
	"movie-app/dto"
	"movie-app/model"
	"movie-app/response"
	"movie-app/utils"
	"movie-app/vo"
	"net/http"
)

func AdminLoginService(requestAdmin dto.AdminLoginDto) response.ResponseStruct {
	DB := common.GetDB()
	res := response.ResponseStruct{
		HttpStatus: http.StatusOK,
		Code:       response.SuccessCode,
		Data:       nil,
		Msg:        response.OK,
	}
	// 判断手机号是否存在
	var admin model.Admin
	DB.Where("telephone = ?", requestAdmin.Telephone).First(&admin)
	if admin.ID == 0 {
		res.HttpStatus = http.StatusUnprocessableEntity
		res.Code = response.CheckFailCode
		res.Msg = response.UserNoExit
		return res
	}
	// 判断密码是否正确
	if err := bcrypt.CompareHashAndPassword([]byte(admin.Password), []byte(requestAdmin.Password)); err != nil {
		res.HttpStatus = http.StatusUnprocessableEntity
		res.Code = response.CheckFailCode
		res.Msg = response.PasswordError
		return res
	}
	// 发放 token
	refreshToken, accessToken, err := common.ReleaseAdminToken(admin)
	if err != nil {
		res.HttpStatus = http.StatusInternalServerError
		res.Code = response.ServerErrorCode
		res.Msg = response.SystemError
		log.Printf("admin_token generate error: %v", err)
		return res
	}
	res.Data = gin.H{"refreshToken": refreshToken, "accessToken": accessToken}
	return res
}

// 增加管理员
func AddAdminService(requestAdmin dto.AddAdminDto) response.ResponseStruct {
	DB := common.GetDB()
	res := response.ResponseStruct{
		HttpStatus: http.StatusOK,
		Code:       response.SuccessCode,
		Data:       nil,
		Msg:        response.OK,
	}
	// 判断手机号是否存在
	if utils.IsAdminTelephoneExit(DB, requestAdmin.Telephone) {
		res.HttpStatus = http.StatusUnprocessableEntity
		res.Code = response.CheckFailCode
		res.Msg = response.PhoneRegistered
		return res
	}
	// 加密密码
	hasedPassword, err := bcrypt.GenerateFromPassword([]byte(requestAdmin.Password), bcrypt.DefaultCost)
	if err != nil {
		res.HttpStatus = http.StatusInternalServerError
		res.Code = response.ServerErrorCode
		res.Msg = response.SystemError
		return res
	}
	// 新增管理员
	newAdmin := model.Admin{
		Name:      requestAdmin.Name,
		Email:     requestAdmin.Email,
		Telephone: requestAdmin.Telephone,
		Password:  string(hasedPassword),
		Authority: requestAdmin.Authority,
	}
	DB.Create(&newAdmin)
	return res
}

func AdminInfoService(adminId uint) response.ResponseStruct {
	res := response.ResponseStruct{
		HttpStatus: http.StatusOK,
		Code:       response.SuccessCode,
		Data:       nil,
		Msg:        response.OK,
	}
	DB := common.GetDB()

	var admin model.Admin
	if err := DB.Where("id = ? ", adminId).First(&admin).Error; err != nil {
		res.HttpStatus = http.StatusUnprocessableEntity
		res.Code = response.CheckFailCode
		res.Msg = response.SystemError
		return res
	}
	res.Data = gin.H{"admin": vo.ToAdminVo(admin)}
	return res
}
