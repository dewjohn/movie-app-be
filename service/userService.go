package service

import (
	"github.com/gin-gonic/gin"
	"log"
	"movie-app/common"
	"movie-app/dto"
	"movie-app/model"
	"movie-app/response"
	"movie-app/utils"
	"net/http"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func RegisterService(requestUser dto.RegisterDto) response.ResponseStruct {
	DB := common.GetDB()
	res := response.ResponseStruct{
		HttpStatus: http.StatusOK,
		Code:       http.StatusOK,
		Data:       nil,
		Msg:        response.OK,
	}
	// 判断手机号是否存在
	if utils.IsTelephoneExist(DB, requestUser.Telephone) {
		res.HttpStatus = http.StatusUnprocessableEntity
		res.Code = 422
		res.Msg = response.PhoneRegistered
		return res
	}

	// 加密密码
	hasedPassword, err := bcrypt.GenerateFromPassword([]byte(requestUser.Password), bcrypt.DefaultCost)
	if err != nil {
		res.HttpStatus = http.StatusInternalServerError
		res.Code = 500
		res.Msg = response.SystemError
		return res
	}

	// 创建用户
	newUser := model.User{
		Name:      requestUser.Name,
		Email:     requestUser.Email,
		Telephone: requestUser.Telephone,
		Password:  string(hasedPassword),
	}
	DB.Create(&newUser)
	// 成功返回
	return res
}

func LoginService(requestUser dto.LoginDto) response.ResponseStruct {
	DB := common.GetDB()
	res := response.ResponseStruct{
		HttpStatus: http.StatusOK,
		Code:       http.StatusOK,
		Data:       nil,
		Msg:        response.OK,
	}
	// 判断手机号是否存在
	var user model.User
	DB.Where("telephone = ?", requestUser.Telephone).First(&user)
	if user.ID == 0 {
		res.HttpStatus = http.StatusUnprocessableEntity
		res.Code = 422
		res.Msg = response.UserNoExit
		return res
	}
	// 判断密码是否正确
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(requestUser.Password)); err != nil {
		res.HttpStatus = http.StatusBadRequest
		res.Code = 400
		res.Msg = response.PasswordError
		return res
	}
	// 发放token
	token, err := common.ReleaseToken(user)
	if err != nil {
		res.HttpStatus = http.StatusInternalServerError
		res.Code = 500
		res.Msg = response.SystemError
		log.Printf("token generate error: %v", err)
		return res
	}

	res.Data = gin.H{"token": token}
	return res
}

func UserModifyService(requestUser dto.UserModifyDto, userId interface{}, tBirthday time.Time) response.ResponseStruct {
	res := response.ResponseStruct{
		HttpStatus: http.StatusOK,
		Code:       http.StatusOK,
		Data:       nil,
		Msg:        response.OK,
	}
	DB := common.GetDB()
	err := DB.Model(&model.User{}).Where("id = ?", userId).Updates(map[string]interface{}{
		"name":     requestUser.Name,
		"gender":   requestUser.Gender,
		"birthday": tBirthday,
		"sign":     requestUser.Sign,
	}).Error
	if err != nil {
		res.HttpStatus = http.StatusInternalServerError
		res.Code = 500
		res.Msg = response.SystemError
	}
	return res
}

// 用户修改密码
func UserModifyPasswordService(requestUser dto.UserModifyPasswordDto, user model.User) response.ResponseStruct {
	res := response.ResponseStruct{
		HttpStatus: http.StatusOK,
		Code:       http.StatusOK,
		Data:       nil,
		Msg:        response.OK,
	}

	// 验证旧密码
	isRight := utils.ComparePasswords(user.Password, []byte(requestUser.OldPassword))
	if !isRight {
		res.HttpStatus = http.StatusBadRequest
		res.Code = 422
		res.Msg = response.OldPasswordError
		return res
	}

	//更新密码
	DB := common.GetDB()
	// 加密密码
	hasedPassword, err1 := bcrypt.GenerateFromPassword([]byte(requestUser.Password), bcrypt.DefaultCost)
	if err1 != nil {
		res.HttpStatus = http.StatusInternalServerError
		res.Code = 500
		res.Msg = response.SystemError
		return res
	}
	err := DB.Model(&user).Update("password", hasedPassword).Error
	if err != nil {
		res.HttpStatus = http.StatusBadRequest
		res.Code = 500
		res.Msg = response.SystemError
		return res
	}
	return res
}
