package admin

import (
	"github.com/gin-gonic/gin"
	"movie-app/common"
	"movie-app/dto"
	"movie-app/model"
	"movie-app/response"
	"net/http"
	"time"
)

func StatisticsAllDataService() response.ResponseStruct {
	res := response.ResponseStruct{
		HttpStatus: http.StatusOK,
		Code:       response.SuccessCode,
		Data:       nil,
		Msg:        response.OK,
	}
	DB := common.GetDB()
	var countMovie int64
	var countUser int64
	var countReply int64
	var countReview int64
	DB.Model(&model.Comment{}).Where("deleted_at is null and parent_id = 0").Count(&countReview)
	DB.Model(&model.Comment{}).Where("deleted_at is null and parent_id > 0").Count(&countReply)
	DB.Model(&model.User{}).Where("deleted_at is null").Count(&countUser)
	DB.Model(&model.Movie{}).Where("deleted_at is null").Count(&countMovie)
	results := dto.StatisticsAllDataDtoResultDto{
		Type:  []string{"电影数", "用户数", "影评数", "讨论数"},
		Value: []int64{countMovie, countUser, countReview, countReply},
	}
	res.Data = gin.H{"count": results}
	return res
}

func StatisticsUploadMovieRecentMonthService() response.ResponseStruct {
	res := response.ResponseStruct{
		HttpStatus: http.StatusOK,
		Code:       response.SuccessCode,
		Data:       nil,
		Msg:        response.OK,
	}
	DB := common.GetDB()
	now := time.Now()
	var lastMonthCount int64 = 0
	var last2MonthCount int64 = 0
	var curMonthCount int64 = 0
	lastMonth := now.AddDate(0, -1, 0)
	last2Month := now.AddDate(0, -2, 0)

	// 获取上上个月第一天
	firstDayOfLastLastMonth := time.Date(last2Month.Year(), last2Month.Month(), 1, 0, 0, 0, 0, time.Local)

	// 获取上上个月最后一天
	lastDayOfLastLastMonth := firstDayOfLastLastMonth.AddDate(0, 1, -1)

	// 获取上个月第一天
	firstDayOfLastMonth := time.Date(lastMonth.Year(), lastMonth.Month(), 1, 0, 0, 0, 0, time.Local)
	// 获取上个月最后一天
	lastDayOfLastMonth := firstDayOfLastMonth.AddDate(0, 1, -1)

	// 本月的开始时间
	thisMonthStart := time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, now.Location())
	// 本月结束时间
	thisMonthEnd := thisMonthStart.AddDate(0, 1, -1)

	results := dto.StatisticsMovieDto{
		Month: []time.Month{last2Month.Month(), lastMonth.Month(), thisMonthStart.Month()},
		Value: []int64{},
	}
	err := DB.Model(&model.Movie{}).
		Where("created_at >= ? AND created_at < ?", firstDayOfLastLastMonth.Format("2006-01-02"), lastDayOfLastLastMonth.Format("2006-01-02")).
		Count(&last2MonthCount).Error
	results.Value = append(results.Value, last2MonthCount)
	if err != nil {
		results.Value = append(results.Value, last2MonthCount)
	}

	err = DB.Model(&model.Movie{}).
		Where("created_at >= ? AND created_at < ?", firstDayOfLastMonth.Format("2006-01-02"), lastDayOfLastMonth.Format("2006-01-02")).
		Count(&lastMonthCount).Error
	results.Value = append(results.Value, lastMonthCount)
	if err != nil {
		results.Value = append(results.Value, lastMonthCount)
	}

	err = DB.Model(&model.Movie{}).
		Where("created_at >= ? AND created_at < ?", thisMonthStart.Format("2006-01-02"), thisMonthEnd.Format("2006-01-02")).
		Count(&curMonthCount).Error
	results.Value = append(results.Value, curMonthCount)
	if err != nil {
		results.Value = append(results.Value, curMonthCount)
	}

	res.Data = gin.H{"count": results}
	return res
}
