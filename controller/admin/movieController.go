package admin

import (
	"github.com/gin-gonic/gin"
	"movie-app/dto"
	"movie-app/response"
	"movie-app/service"
	"time"
)

func UploadVideoInfo(ctx *gin.Context) {
	var video dto.VideoDto
	err := ctx.Bind(&video)
	if err != nil {
		response.Fail(ctx, nil, "请求错误")
	}
	title := video.Title
	cover := video.Cover
	ReleaseTime := video.ReleaseTime
	SheetLength := video.SheetLength // 占位符，v2引入自动计算上传视频后获取视频的长度
	Origin := video.Origin
	Type := video.Type
	Director := video.Director
	Screenwriter := video.Screenwriter
	Actors := video.Actors
	Language := video.Language
	Introduction := video.Introduction
	adminId, _ := ctx.Get("adminId") // 在上下文中拿到用户id

	// 验证数据
	if len(title) == 0 {
		response.CheckFail(ctx, nil, "标题不能为空")
	}
	if len(cover) == 0 {
		response.CheckFail(ctx, nil, "视频封面不能为空")
	}
	// 判断日期
	tReleaseTime, err := time.Parse("2006-01-02", ReleaseTime)
	if err != nil {
		response.CheckFail(ctx, nil, "请输入正确日期")
	}
	if SheetLength == 0 {
		response.CheckFail(ctx, nil, "视频片长不能为空")
	}
	if len(Origin) == 0 {
		response.CheckFail(ctx, nil, "视频产地不能为空")
	}
	if len(Type) == 0 {
		response.CheckFail(ctx, nil, "视频类型不能为空")
	}
	if len(Director) == 0 {
		response.CheckFail(ctx, nil, "请输入导演信息")
	}
	if len(Screenwriter) == 0 {
		response.CheckFail(ctx, nil, "请输入编剧信息")
	}
	if len(Actors) == 0 {
		response.CheckFail(ctx, nil, "请输入演员信息")
	}
	if len(Language) == 0 {
		response.CheckFail(ctx, nil, "请输入电影语言类型")
	}
	if len(Introduction) == 0 {
		response.CheckFail(ctx, nil, "视频简介不能为空")
	}
	res := service.UploadVideoInfoService(video, adminId, tReleaseTime)
	response.HandleResponse(ctx, res)
}
