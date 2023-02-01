package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Response(ctx *gin.Context, httpStatus int, code int, data gin.H, msg string) {
	ctx.JSON(httpStatus, gin.H{"code": code, "data": data, "msg": msg})
}

func Success(ctx *gin.Context, data gin.H, msg string) {
	Response(ctx, http.StatusOK, 200, data, msg)
}

func Fail(ctx *gin.Context, data gin.H, msg string) {
	Response(ctx, http.StatusBadRequest, 400, data, msg)
}

func CheckFail(ctx *gin.Context, data gin.H, msg string) {
	Response(ctx, http.StatusUnprocessableEntity, 422, data, msg)
}

type ResponseStruct struct {
	HttpStatus int    //http状态
	Code       int    //状态码
	Data       gin.H  //数据
	Msg        string //信息
}

// 统一处理返回信息
func HandleResponse(ctx *gin.Context, res ResponseStruct) {
	ctx.JSON(res.HttpStatus, gin.H{"code": res.Code, "data": res.Data, "msg": res.Msg})
}
