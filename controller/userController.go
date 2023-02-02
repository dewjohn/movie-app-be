package controller

import (
	"movie-app/dto"
	"movie-app/model"
	"movie-app/response"
	"movie-app/service"
	"movie-app/utils"
	"movie-app/vo"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// @Summary 用户注册
// @Produce  json
// @Tags User
// @Param name body string true "姓名"
// @Param email body string true "邮箱"
// @Param telephone body string true "电话"
// @Param password body string true "密码"
// @Router /api/v1/user/register [post]
func Register(ctx *gin.Context) {
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
	if len(password) < 6 {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "密码至少6位")
		return
	}
	// 如果名称没有传，随机字符串
	if len(name) == 0 {
		name = utils.RandomString(10)
	}
	res := service.RegisterService(requestUser)
	response.HandleResponse(ctx, res)
}

// 登陆
func Login(ctx *gin.Context) {
	// 获取参数
	var requestUser = dto.LoginDto{}
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
	if len(password) < 6 {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "密码至少为6位")
		return
	}
	res := service.LoginService(requestUser)
	response.HandleResponse(ctx, res)
}

// 用户获取个人信息
func UserInfo(ctx *gin.Context) {
	user, _ := ctx.Get("user")
	response.Response(ctx, http.StatusOK, 200, gin.H{"user": vo.ToUserVo(user.(model.User))}, "获取个人信息成功")
}

// 用户修改个人信息
func UserModify(ctx *gin.Context) {
	var requestUser = dto.UserModifyDto{}
	err := ctx.Bind(&requestUser)

	name := requestUser.Name
	birthday := requestUser.Birthday

	// 判断昵称
	if len(name) == 0 {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "昵称不能为空")
		return
	}
	// 判断日期
	tBirthday, err := time.Parse("2006-01-02", birthday)
	if err != nil {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "日期错误")
		return
	}

	// 获取上下文的 userId
	userId, _ := ctx.Get("userId")
	res := service.UserModifyService(requestUser, userId, tBirthday)
	response.HandleResponse(ctx, res)
}

// 用户修改密码
func UserModifyPassword(ctx *gin.Context) {
	var requestUser = dto.UserModifyPasswordDto{}
	err := ctx.Bind(&requestUser)
	if err != nil {
		response.Fail(ctx, nil, "请求错误")
		return
	}
	// 判断密码不能为空
	if len(requestUser.Password) < 6 {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "密码至少6位")
		return
	}
	// 获取上下文的 user
	user, _ := ctx.Get("user")
	modelUser := user.(model.User)

	res := service.UserModifyPasswordService(requestUser, modelUser)
	response.HandleResponse(ctx, res)
}
