package admin

import (
	"github.com/gin-gonic/gin"
	"movie-app/dto"
	"movie-app/model"
	"movie-app/response"
	"movie-app/service/admin"
	"movie-app/utils"
	"movie-app/vo"
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
		response.Response(ctx, http.StatusBadRequest, 400, nil, response.RequestError)
		return
	}
	telephone := requestAdmin.Telephone
	password := requestAdmin.Password

	// 验证数据
	if len(telephone) != 11 {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, response.PhoneNumberError)
		return
	}
	if len(password) < 6 {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, response.PasswordNumberError)
		return
	}
	res := admin.AdminLoginService(requestAdmin)
	response.HandleResponse(ctx, res)
}

func AdminInfo(ctx *gin.Context) {
	adminInfo, _ := ctx.Get("admin")
	response.Response(ctx, http.StatusOK, 200, gin.H{"admin": vo.ToAdminVo(adminInfo.(model.Admin))}, response.OK)
}

// 增加管理员
func AddAdmin(ctx *gin.Context) {
	var requestAdmin = dto.AddAdminDto{}
	err := ctx.Bind(&requestAdmin)
	if err != nil {
		response.Response(ctx, http.StatusBadRequest, 400, nil, response.RequestError)
		return
	}
	name := requestAdmin.Name
	email := requestAdmin.Email
	telephone := requestAdmin.Telephone
	password := requestAdmin.Password
	authority := requestAdmin.Authority

	// 验证数据
	if len(telephone) != 11 {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, response.PhoneNumberError)
		return
	}
	if !utils.VerifyEmailFormat(email) {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, response.MailTypeError)
		return
	}
	if len(password) < 6 {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, response.PasswordNumberError)
		return
	}
	if authority != Admin && authority != Auditor {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, response.AuthorityError)
		return
	}
	// 如果名称没有传，随机字符串
	if len(name) == 0 {
		name = utils.RandomString(10)
	}
	res := admin.AddAdminService(requestAdmin)
	response.HandleResponse(ctx, res)
}
