package controller

import (
	"github.com/gin-gonic/gin"
	"movie-app/response"
	"movie-app/service"
)

func GetTypeCategory(ctx *gin.Context) {
	res := service.GetTypeCategoryService()
	response.HandleResponse(ctx, res)
}

func GetOriginCategory(ctx *gin.Context) {
	res := service.GetOriginCategoryService()
	response.HandleResponse(ctx, res)
}

func GetLanguageCategory(ctx *gin.Context) {
	res := service.GetLanguageCategoryService()
	response.HandleResponse(ctx, res)
}

func GetReleaseCategory(ctx *gin.Context) {
	res := service.GetReleaseCategoryService()
	response.HandleResponse(ctx, res)
}
