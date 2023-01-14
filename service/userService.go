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

	"golang.org/x/crypto/bcrypt"
)

func RegisterService(requestUser dto.RegisterDto) response.ResponseStruct {
	DB := common.GetDB()
	res := response.ResponseStruct{
		HttpStatus: http.StatusOK,
		Code:       http.StatusOK,
		Data:       nil,
		Msg:        "注册成功",
	}
	// 判断手机号是否存在
	if utils.IsTelephoneExist(DB, requestUser.Telephone) {
		res.HttpStatus = http.StatusUnprocessableEntity
		res.Code = 422
		res.Msg = "用户已存在"
		return res
	}

	// 加密密码
	hasedPassword, err := bcrypt.GenerateFromPassword([]byte(requestUser.Password), bcrypt.DefaultCost)
	if err != nil {
		res.HttpStatus = http.StatusInternalServerError
		res.Code = 500
		res.Msg = "加密错误"
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
		Msg:        "登陆成功",
	}
	// 判断手机号是否存在
	var user model.User
	DB.Where("telephone = ?", requestUser.Telephone).First(&user)
	if user.ID == 0 {
		res.HttpStatus = http.StatusUnprocessableEntity
		res.Code = 422
		res.Msg = "用户不存在"
		return res
	}
	// 判断密码是否正确
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(requestUser.Password)); err != nil {
		res.HttpStatus = http.StatusBadRequest
		res.Code = 400
		res.Msg = "密码错误"
		return res
	}
	// 发放token
	token, err := common.ReleaseToken(user)
	if err != nil {
		res.HttpStatus = http.StatusInternalServerError
		res.Code = 500
		res.Msg = "系统异常"
		log.Printf("token generate error: %v", err)
		return res
	}

	res.Data = gin.H{"token": token}
	return res
}

func UserModifyService(requestUser dto.UserModifyDto, userId interface{}) response.ResponseStruct {
	res := response.ResponseStruct{
		HttpStatus: http.StatusOK,
		Code:       http.StatusOK,
		Data:       nil,
		Msg:        "修改个人信息成功",
	}
	DB := common.GetDB()
	err := DB.Model(model.User{}).Where("id = ?", userId).Updates(map[string]interface{}{
		"name":     requestUser.Name,
		"avatar":   requestUser.Avatar,
		"gender":   requestUser.Gender,
		"birthday": requestUser.Birthday,
		"sign":     requestUser.Sign,
	}).Error
	if err != nil {
		res.HttpStatus = http.StatusInternalServerError
		res.Code = 500
		res.Msg = "修改失败"
	}
	return res
}

// 用户修改密码
func UserModifyPasswordService(requestUser dto.UserModifyPasswordDto, user model.User) response.ResponseStruct {
	res := response.ResponseStruct{
		HttpStatus: http.StatusOK,
		Code:       http.StatusOK,
		Data:       nil,
		Msg:        "修改成功",
	}

	// 验证旧密码
	isRight := utils.ComparePasswords(user.Password, []byte(requestUser.OldPassword))
	if !isRight {
		res.HttpStatus = http.StatusBadRequest
		res.Code = 422
		res.Msg = "旧密码错误"
		return res
	}

	//更新密码
	DB := common.GetDB()
	// 加密密码
	hasedPassword, err1 := bcrypt.GenerateFromPassword([]byte(requestUser.Password), bcrypt.DefaultCost)
	if err1 != nil {
		res.HttpStatus = http.StatusInternalServerError
		res.Code = 500
		res.Msg = "加密错误"
		return res
	}
	err := DB.Model(&user).Update("password", hasedPassword).Error
	if err != nil {
		res.HttpStatus = http.StatusBadRequest
		res.Code = 500
		res.Msg = "修改密码失败"
		return res
	}
	return res
}
