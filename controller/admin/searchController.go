package admin

import (
	"github.com/gin-gonic/gin"
	"movie-app/response"
	"movie-app/service/admin"
	"strings"
)

func SearchUser(ctx *gin.Context) {
	search := ctx.Query("uname")
	if len(search) == 0 {
		response.CheckFail(ctx, nil, response.SearchNotEmpty)
		return
	}
	uname := "%" + strings.Replace(search, " ", "%", -1) + "%"
	res := admin.SearchUserService(uname)
	response.HandleResponse(ctx, res)
}
