package utils

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"movie-app/response"
)

func verifyMovieInfo(ctx *gin.Context, Title string, Cover string, SheetLength int, Origin string, Type string, Director string, Screenwriter string, Actors string, Language string, Introduction string) {
	// 验证数据
	if len(Title) == 0 {
		response.CheckFail(ctx, nil, "标题不能为空")
		return
	}
	if len(Cover) == 0 {
		response.CheckFail(ctx, nil, "视频封面不能为空")
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
}

func GetUrl() string {
	return "http://localhost:" + viper.GetString("server.port") + "/api/"
}
