package admin

import (
	"github.com/gin-gonic/gin"
	"movie-app/response"
	"movie-app/service/admin"
)

func StatisticsAllData(ctx *gin.Context) {
	res := admin.StatisticsAllDataService()
	response.HandleResponse(ctx, res)
}

func StatisticsUploadMovieRecentMonth(ctx *gin.Context) {
	res := admin.StatisticsUploadMovieRecentMonthService()
	response.HandleResponse(ctx, res)
}
