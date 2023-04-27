package controller

import (
	"github.com/gin-gonic/gin"
	"movie-app/response"
	"movie-app/service"
	"movie-app/utils"
	"os"
	"path"
	"strconv"
	"time"
)

func UploadAvatar(ctx *gin.Context) {
	avatar, err := ctx.FormFile("avatar")
	if err != nil {
		response.Fail(ctx, nil, response.FailUploadImage)
		return
	}
	suffix := path.Ext(avatar.Filename) // 获取文件的扩展名
	if suffix != ".jpg" && suffix != ".png" && suffix != ".webp" && suffix != ".jpeg" {
		response.CheckFail(ctx, nil, response.ImageTypeError)
	}
	avatar.Filename = utils.RandomString(3) + strconv.FormatInt(time.Now().UnixNano(), 10) + suffix // 重定义头像命名
	// 如果不存在avatar文件夹创建
	if _, err := os.Stat("./files/avatar"); os.IsNotExist(err) {
		err := os.Mkdir("./files/avatar", os.ModePerm)
		if err != nil {
			return
		}
	}
	// 保存文件
	dst := path.Join("./files/avatar", avatar.Filename)
	errSave := ctx.SaveUploadedFile(avatar, dst)
	if errSave != nil {
		response.Fail(ctx, nil, response.SaveImageError)
		return
	}
	// 获取文件属性
	fileInfo, err := os.Stat("./files/avatar/" + avatar.Filename)
	//大小限制到5M
	if fileInfo == nil || fileInfo.Size() > 1024*1024*5 || err != nil {
		response.CheckFail(ctx, nil, response.ImageTypeError)
		return
	}
	uid, _ := ctx.Get("userId")
	// 拼接上传图片的路径信息
	localFileName := "./file/avatar/" + avatar.Filename
	objectName := "avatar/" + avatar.Filename

	res := service.UploadAvatarService(localFileName, objectName, uid.(uint))
	response.HandleResponse(ctx, res)
}
