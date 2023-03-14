package controller

import (
	"github.com/gin-gonic/gin"
	"movie-app/response"
	"movie-app/service"
	"strings"
)

func SerchMovie(ctx *gin.Context) {
	search := ctx.Query("keywords")
	if len(search) == 0 {
		response.CheckFail(ctx, nil, response.SearchNotEmpty)
		return
	}
	keywords := "%" + strings.Replace(search, " ", "%", -1) + "%"
	res := service.SearchMovieService(keywords)
	response.HandleResponse(ctx, res)
}
