package admin

import (
	"github.com/gin-gonic/gin"
	"movie-app/dto"
	"movie-app/response"
	"movie-app/service"
	"movie-app/service/admin"
	"time"
)

func UploadVideoInfo(ctx *gin.Context) {
	var video dto.VideoDto
	err := ctx.Bind(&video)
	if err != nil {
		response.Fail(ctx, nil, "请求错误")
	}
	Title := video.Title
	Cover := video.Cover // 通过 api/v1/upload/cover 接口返回 cover string
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
	if len(Title) == 0 {
		response.CheckFail(ctx, nil, "标题不能为空")
		return
	}
	if len(Cover) == 0 {
		response.CheckFail(ctx, nil, "视频封面不能为空")
		return
	}
	// 判断日期
	tReleaseTime, err := time.Parse("2006-01-02", ReleaseTime)
	if err != nil {
		response.CheckFail(ctx, nil, "请输入正确日期")
		return
	}
	if SheetLength == 0 {
		response.CheckFail(ctx, nil, "视频片长不能为空")
		return
	}
	if len(Origin) == 0 {
		response.CheckFail(ctx, nil, "视频产地不能为空")
		return
	}
	if len(Type) == 0 {
		response.CheckFail(ctx, nil, "视频类型不能为空")
		return
	}
	if len(Director) == 0 {
		response.CheckFail(ctx, nil, "请输入导演信息")
		return
	}
	if len(Screenwriter) == 0 {
		response.CheckFail(ctx, nil, "请输入编剧信息")
		return
	}
	if len(Actors) == 0 {
		response.CheckFail(ctx, nil, "请输入演员信息")
		return
	}
	if len(Language) == 0 {
		response.CheckFail(ctx, nil, "请输入电影语言类型")
		return
	}
	if len(Introduction) == 0 {
		response.CheckFail(ctx, nil, "视频简介不能为空")
		return
	}

	res := service.UploadVideoInfoService(video, adminId, tReleaseTime)
	response.HandleResponse(ctx, res)
}

// 修改视频信息
func ModifyVideoInfo(ctx *gin.Context) {
	var video = dto.ModifyVideoDto{}
	err := ctx.Bind(&video)
	if err != nil {
		response.Fail(ctx, nil, "请求错误")
	}
	Vid := video.Vid
	Title := video.Title
	Cover := video.Cover // 通过 api/v1/upload/cover 接口返回 cover string
	ReleaseTime := video.ReleaseTime
	SheetLength := video.SheetLength // 占位符，v2引入自动计算上传视频后获取视频的长度
	Origin := video.Origin
	Type := video.Type
	Director := video.Director
	Screenwriter := video.Screenwriter
	Actors := video.Actors
	Language := video.Language
	Introduction := video.Introduction
	// 验证数据
	if Vid == 0 {
		response.CheckFail(ctx, nil, "视频ID不能为空")
		return
	}
	if len(Title) == 0 {
		response.CheckFail(ctx, nil, "标题不能为空")
		return
	}
	if len(Cover) == 0 {
		response.CheckFail(ctx, nil, "视频封面不能为空")
		return
	}
	// 判断日期
	tReleaseTime, err := time.Parse("2006-01-02", ReleaseTime)
	if err != nil {
		response.CheckFail(ctx, nil, "请输入正确日期")
		return
	}
	if SheetLength == 0 {
		response.CheckFail(ctx, nil, "视频片长不能为空")
		return
	}
	if len(Origin) == 0 {
		response.CheckFail(ctx, nil, "视频产地不能为空")
		return
	}
	if len(Type) == 0 {
		response.CheckFail(ctx, nil, "视频类型不能为空")
		return
	}
	if len(Director) == 0 {
		response.CheckFail(ctx, nil, "请输入导演信息")
		return
	}
	if len(Screenwriter) == 0 {
		response.CheckFail(ctx, nil, "请输入编剧信息")
		return
	}
	if len(Actors) == 0 {
		response.CheckFail(ctx, nil, "请输入演员信息")
		return
	}
	if len(Language) == 0 {
		response.CheckFail(ctx, nil, "请输入电影语言类型")
		return
	}
	if len(Introduction) == 0 {
		response.CheckFail(ctx, nil, "视频简介不能为空")
		return
	}
	res := service.ModifyVideoInfoService(video, tReleaseTime)
	response.HandleResponse(ctx, res)
}

// 删除视频信息
func DeleteMovieVideo(ctx *gin.Context) {
	var video dto.VideoIdDto
	err := ctx.Bind(&video)
	if err != nil {
		response.Fail(ctx, nil, "请求错误")
		return
	}
	id := video.Id
	if id == 0 {
		response.CheckFail(ctx, nil, "视频不存在")
		return
	}
	res := admin.DeleteMovieVideoService(id)
	response.HandleResponse(ctx, res)
}