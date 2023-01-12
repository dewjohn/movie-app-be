package main

import (
	"github.com/gin-gonic/gin"
	"movie-app/common"
	"movie-app/routes"
)

func main() {
	db := common.InitDB()
	_ = db
	r := gin.Default()
	r = routes.CollectRoute(r)
	panic(r.Run()) // listen and  serve on 0.0.0.0:8080
}
