package admin

import (
	"github.com/gin-gonic/gin"
	"movie-app/common"
	"movie-app/dto"
	"movie-app/response"
	"movie-app/service/admin"
	"movie-app/utils"
)

func GetUser(ctx *gin.Context) {
	page := utils.StringToInt(ctx.Query("page"))
	pageSize := utils.StringToInt(ctx.Query("page_size"))
	if page <= 0 || pageSize <= 0 {
		response.Fail(ctx, nil, response.PageError)
		return
	}
	if pageSize >= 30 {
		response.Fail(ctx, nil, response.RequestTooMany)
		return
	}
	res := admin.GetUserService(page, pageSize)
	response.HandleResponse(ctx, res)
}

// 更改用户状态
func ChangeUserState(ctx *gin.Context) {
	var userState dto.UserStateDto
	err := ctx.Bind(&userState)
	if err != nil {
		response.Fail(ctx, nil, response.RequestError)
		return
	}
	if userState.State != common.Normal && userState.State != common.Banned && userState.State != common.ShutUp {
		response.CheckFail(ctx, nil, "状态码错误")
		return
	}
	adminAuth, _ := ctx.Get("adminAuthorization")
	if adminAuth.(uint) < 2000 {
		response.CheckFail(ctx, nil, "权限不足")
		return
	}
	res := admin.ChangeUserStateService(userState)
	response.HandleResponse(ctx, res)
}
