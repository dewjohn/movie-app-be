package service

import (
	"github.com/gin-gonic/gin"
	"movie-app/common"
	"movie-app/model"
	"movie-app/response"
	"movie-app/utils"
	"net/http"
	"regexp"
)

func GetTypeCategoryService() response.ResponseStruct {
	res := response.ResponseStruct{
		HttpStatus: http.StatusOK,
		Code:       response.SuccessCode,
		Data:       nil,
		Msg:        response.OK,
	}
	DB := common.GetDB()
	var movies []model.Movie
	re := regexp.MustCompile(`([\p{Han}]+)(?: \/ )?`) // 定义正则表达式，匹配中文字符和斜杠，最后一个中文后面可能没有斜杠
	var result []string
	DB.Find(&movies)
	for _, movie := range movies {
		matches := re.FindAllStringSubmatch(movie.Type, -1) // 在字符串中查找所有匹配的子串
		if len(matches) > 0 {
			chineseList := make([]string, len(matches))
			for i, match := range matches {
				chineseList[i] = match[1]
				result = append(result, chineseList[i])
			}
		}
	}
	// 去重
	var resultTags = utils.StringArrayUnique(result)
	res.Data = gin.H{"tags": resultTags}
	return res
}

func GetOriginCategoryService() response.ResponseStruct {
	res := response.ResponseStruct{
		HttpStatus: http.StatusOK,
		Code:       response.SuccessCode,
		Data:       nil,
		Msg:        response.OK,
	}
	DB := common.GetDB()
	var movies []model.Movie
	re := regexp.MustCompile(`([\p{Han}]+)(?: \/ )?`) // 定义正则表达式，匹配中文字符和斜杠，最后一个中文后面可能没有斜杠
	var result []string
	DB.Find(&movies)
	for _, movie := range movies {
		matches := re.FindAllStringSubmatch(movie.Origin, -1) // 在字符串中查找所有匹配的子串
		if len(matches) > 0 {
			chineseList := make([]string, len(matches))
			for i, match := range matches {
				chineseList[i] = match[1]
				result = append(result, chineseList[i])
			}
		}
	}
	// 去重
	var resultTags = utils.StringArrayUnique(result)
	res.Data = gin.H{"tags": resultTags}
	return res
}

func GetLanguageCategoryService() response.ResponseStruct {
	res := response.ResponseStruct{
		HttpStatus: http.StatusOK,
		Code:       response.SuccessCode,
		Data:       nil,
		Msg:        response.OK,
	}
	DB := common.GetDB()
	var movies []model.Movie
	re := regexp.MustCompile(`([\p{Han}]+)(?: \/ )?`) // 定义正则表达式，匹配中文字符和斜杠，最后一个中文后面可能没有斜杠
	var result []string
	DB.Find(&movies)
	for _, movie := range movies {
		matches := re.FindAllStringSubmatch(movie.Language, -1) // 在字符串中查找所有匹配的子串
		if len(matches) > 0 {
			chineseList := make([]string, len(matches))
			for i, match := range matches {
				chineseList[i] = match[1]
				result = append(result, chineseList[i])
			}
		}
	}
	// 去重
	var resultTags = utils.StringArrayUnique(result)
	res.Data = gin.H{"tags": resultTags}
	return res
}

func GetReleaseCategoryService() response.ResponseStruct {
	res := response.ResponseStruct{
		HttpStatus: http.StatusOK,
		Code:       response.SuccessCode,
		Data:       nil,
		Msg:        response.OK,
	}
	DB := common.GetDB()
	var movies []model.Movie
	var result []int
	DB.Find(&movies)
	for _, movie := range movies {
		result = append(result, movie.ReleaseTime.Year())
	}
	result = utils.IntArrayUnique(result)
	res.Data = gin.H{"tags": result}
	return res
}
