package admin

import (
	"github.com/gin-gonic/gin"
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

const (
	Normal = 100
	ShutUp = 200
	Banned = 300
)

// 更改用户状态
func ChangeUserState(ctx *gin.Context) {
	var userState dto.UserStateDto
	err := ctx.Bind(&userState)
	if err != nil {
		response.Fail(ctx, nil, response.RequestError)
		return
	}
	if userState.State != Normal || userState.State != Banned || userState.State != ShutUp {
		response.CheckFail(ctx, nil, "状态码错误")
		return
	}
	adminId, _ := ctx.Get("adminId")
	res := admin.ChangeUserStateService(adminId, userState)
	response.HandleResponse(ctx, res)
}
