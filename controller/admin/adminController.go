package admin

import (
	"github.com/gin-gonic/gin"
	"movie-app/dto"
	"movie-app/response"
	"movie-app/service"
	"movie-app/utils"
	"net/http"
)

const (
	SuperAdmin = 3000
	Admin      = 2000
	Auditor    = 1000
)

// 管理员登陆
func AdminLogin(ctx *gin.Context) {
	var requestAdmin = dto.AdminLoginDto{}
	err := ctx.Bind(&requestAdmin)
	if err != nil {
		response.Response(ctx, http.StatusBadRequest, 400, nil, "请求失败")
		return
	}
	telephone := requestAdmin.Telephone
	password := requestAdmin.Password

	// 验证数据
	if len(telephone) != 11 {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "手机号必须为11位")
		return
	}
	if len(password) < 6 {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "密码至少6位")
		return
	}
	res := service.AdminLoginService(requestAdmin)
	response.HandleResponse(ctx, res)
}

// 增加管理员
func AddAdmin(ctx *gin.Context) {
	var requestAdmin = dto.AddAdminDto{}
	err := ctx.Bind(&requestAdmin)
	if err != nil {
		response.Response(ctx, http.StatusBadRequest, 400, nil, "请求失败")
		return
	}
	name := requestAdmin.Name
	email := requestAdmin.Email
	telephone := requestAdmin.Telephone
	password := requestAdmin.Password
	authority := requestAdmin.Authority

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
	if authority != Admin && authority != Auditor {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "权限选择错误")
		return
	}
	// 如果名称没有传，随机字符串
	if len(name) == 0 {
		name = utils.RandomString(10)
	}
	res := service.AddAdminService(requestAdmin)
	response.HandleResponse(ctx, res)
}
