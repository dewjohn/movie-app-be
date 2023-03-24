package admin

import (
	"github.com/gin-gonic/gin"
	"movie-app/dto"
	"movie-app/response"
	"movie-app/service/admin"
	"movie-app/utils"
	"os"
	"path"
	"strconv"
	"time"
)

func UploadMovieCover(ctx *gin.Context) {
	cover, err := ctx.FormFile("cover")
	if err != nil {
		response.Fail(ctx, nil, "封面上传失败")
	}
	suffix := path.Ext(cover.Filename) // 获取文件的扩展名
	if suffix != ".jpg" && suffix != ".png" && suffix != ".webp" && suffix != ".jpeg" {
		response.CheckFail(ctx, nil, "图片不符合要求")
	}
	cover.Filename = utils.RandomString(3) + strconv.FormatInt(time.Now().UnixNano(), 10) + suffix // 重定义封面命名
	// 如果不存在cover文件夹创建
	if _, err := os.Stat("./files/cover"); os.IsNotExist(err) {
		err := os.Mkdir("./files/cover", os.ModePerm)
		if err != nil {
			return
		}
	}
	// 保存文件
	dst := path.Join("./files/cover", cover.Filename)
	errSave := ctx.SaveUploadedFile(cover, dst)
	if errSave != nil {
		response.Fail(ctx, nil, "图片保存失败")
		return
	}
	// 获取文件属性
	fileInfo, err := os.Stat("./files/cover/" + cover.Filename)
	//大小限制到5M
	if fileInfo == nil || fileInfo.Size() > 1024*1024*5 || err != nil {
		response.CheckFail(ctx, nil, "图片不符合要求")
		return
	}
	// 拼接上传图片的路径信息
	objectName := "cover/" + cover.Filename
	res := admin.UploadCoverService(objectName)
	response.HandleResponse(ctx, res)
}

func UploadMovieVideo(ctx *gin.Context) {
	vid, _ := strconv.Atoi(ctx.PostForm("vid")) // 从上传视频信息返回值中拿到生成的vid
	if vid < 0 {
		response.Fail(ctx, nil, "参数错误")
		return
	}
	video, err := ctx.FormFile("video")
	if err != nil {
		response.Fail(ctx, nil, "文件上传失败")
		return
	}
	suffix := path.Ext(video.Filename)                              // 视频后缀
	videoTitle := path.Base(video.Filename)                         // 视频名
	videoTitlePrefix := videoTitle[0 : len(videoTitle)-len(suffix)] // 视频名前缀

	if suffix != ".mp4" {
		response.CheckFail(ctx, nil, "文件格式不符合要求")
		return
	}
	// 生成自定义文件名
	reVideoName := utils.RandomString(3) + strconv.FormatInt(time.Now().UnixNano(), 10)
	video.Filename = reVideoName + suffix

	// 如果没有指定文件夹创建
	if _, err := os.Stat("./files/movie"); os.IsNotExist(err) {
		err := os.Mkdir("./files/movie", os.ModePerm)
		if err != nil {
			return
		}
	}
	// 保存
	//errSave := ctx.SaveUploadedFile(video, "./files/movie")
	dst := path.Join("./files/movie", video.Filename)
	errSave := ctx.SaveUploadedFile(video, dst)
	if errSave != nil {
		response.Fail(ctx, nil, "文件保存失败")
		return
	}

	fileInfo, err := os.Stat("./files/movie/" + video.Filename)
	if fileInfo == nil || fileInfo.Size() > 1024*1024*500 || err != nil {
		response.CheckFail(ctx, nil, "文件大小不符合要求")
		return
	}

	objectName := "movie/" + video.Filename
	var urls dto.ResDto
	urls.Original = utils.GetUrl() + "/api/" + objectName

	res := admin.UploadVideoService(urls, vid, videoTitlePrefix)
	response.HandleResponse(ctx, res)
}

func DeleteResource(ctx *gin.Context) {
	var id dto.UUID
	err := ctx.Bind(&id)
	if err != nil {
		response.Fail(ctx, nil, "请求错误")
		return
	}
	res := admin.DeleteResourceService(id.UUID)
	response.HandleResponse(ctx, res)
}
