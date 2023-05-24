package admin

import (
	"github.com/gin-gonic/gin"
	"movie-app/common"
	"movie-app/dto"
	"movie-app/model"
	"movie-app/response"
	"movie-app/service/admin"
	"movie-app/utils"
	"net/http"
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
	adminId, _ := ctx.Get("adminId")
	res := admin.AdminInfoService(adminId.(uint))
	response.HandleResponse(ctx, res)
}

// 增加管理员
func AddAdmin(ctx *gin.Context) {
	var requestAdmin = dto.AddAdminDto{}
	err := ctx.Bind(&requestAdmin)
	if err != nil {
		response.Response(ctx, http.StatusBadRequest, 400, nil, response.RequestError)
		return
	}
	adminAuthorization, _ := ctx.Get("adminAuthorization")
	if adminAuthorization.(int) < 2000 {
		response.CheckFail(ctx, nil, response.NoAuthorization)
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
	if authority != common.Admin && authority != common.Auditor {
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

/**
* 通过 refreshtoken 刷新 accesstoken
 */
func GetAdminAccessToken(ctx *gin.Context) {
	adminId, exist := ctx.Get("adminId")

	if exist {
		DB := common.GetDB()
		var admin model.Admin
		if err := DB.First(&admin, adminId).Error; err != nil {
			response.Fail(ctx, nil, "请求错误")
			return
		}
		token, _ := common.ReleaseAdminAccessToken(admin)
		response.Success(ctx, gin.H{"accessToken": token}, response.OK)
		return
	}
	response.Fail(ctx, nil, "请求错误")
}
