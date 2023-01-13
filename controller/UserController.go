package controller

import (
	"log"
	"movie-app/common"
	"movie-app/dto"
	"movie-app/model"
	"movie-app/response"
	"movie-app/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Register(ctx *gin.Context) {
	DB := common.GetDB()
	// 获取参数
	var requestUser = dto.RegisterDto{}
	err := ctx.Bind(&requestUser)
	if err != nil {
		return
	}
	name := requestUser.Name
	email := requestUser.Email
	telephone := requestUser.Telephone
	password := requestUser.Password
	// 验证数据
	if len(telephone) != 11 {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "手机号必须为11位")
		return
	}
	if !utils.VerifyEmailFormat(email) {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "邮箱格式错误")
		return
	}
	if len(password) != 6 {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "密码至少为6位")
		return
	}
	// 如果名称没有传，随机字符串
	if len(name) == 0 {
		name = utils.RandomString(10)
	}

	// 判断手机号是否存在
	if utils.IsTelephoneExist(DB, telephone) {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "用户已经存在")
		return
	}

	// 创建用户
	hasedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		response.Response(ctx, http.StatusInternalServerError, 500, nil, "加密错误")
		return
	}
	newUser := model.User{
		Name:      name,
		Telephone: telephone,
		Password:  string(hasedPassword),
	}
	DB.Create(&newUser)

	// 发放token
	token, err := common.ReleaseToken(newUser)
	if err != nil {
		response.Response(ctx, http.StatusInternalServerError, 500, nil, "系统异常")
		log.Printf("token generate error: %v", err)
		return
	}
	// 返回结果
	//ctx.JSON(200, gin.H{
	//	"code": 200,
	//	"msg":  "注册成功",
	//})
	response.Success(ctx, gin.H{"token": token}, "注册成功")
}

func Login(ctx *gin.Context) {
	DB := common.GetDB()
	// 获取参数
	var requestUser = model.User{}
	err := ctx.Bind(&requestUser)
	if err != nil {
		return
	}
	telephone := requestUser.Telephone
	password := requestUser.Password
	// 数据验证
	if len(telephone) != 11 {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "手机号必须为11位")
		return
	}
	if len(password) != 6 {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "密码至少为6位")
		return
	}
	// 判断手机号是否存在
	var user model.User
	DB.Where("telephone = ?", telephone).First(&user)
	if user.ID == 0 {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "用户不存在")
		return
	}
	// 判断密码是否正确
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		response.Response(ctx, http.StatusBadRequest, 400, nil, "密码错误")
		return
	}
	// 发放token
	token, err := common.ReleaseToken(user)
	if err != nil {
		response.Response(ctx, http.StatusInternalServerError, 500, nil, "系统异常")
		log.Printf("token generate error: %v", err)
		return
	}

	// 返回结果
	//ctx.JSON(200, gin.H{
	//	"code": 200,
	//	"data": gin.H{"token": token},
	//	"msg":  "登陆成功",
	//})
	response.Success(ctx, gin.H{"token": token}, "登陆成功")
}

func Info(ctx *gin.Context) {
	user, _ := ctx.Get("user")
	response.Response(ctx, http.StatusOK, 200, gin.H{"user": dto.ToUserDto(user.(model.User))}, "获取个人信息成功")
}
