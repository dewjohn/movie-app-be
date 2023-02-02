package main

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	_ "github.com/swaggo/files"
	_ "github.com/swaggo/gin-swagger"
	"movie-app/common"
	"movie-app/docs"
	"movie-app/routes"
	"os"
)

func main() {
	docs.SwaggerInfo.Title = "movie app"
	docs.SwaggerInfo.Version = "1.0"
	InitConfig()
	db := common.InitDB()
	_ = db
	r := gin.Default()
	r = routes.CollectRoute(r)
	port := viper.GetString("server.port")
	if port != "" {
		panic(r.Run(":" + port))
	}
	panic(r.Run())
}
func InitConfig() {
	workDir, _ := os.Getwd()
	viper.SetConfigName("application")
	viper.SetConfigType("yml")
	viper.AddConfigPath(workDir + "/config")
	err := viper.ReadInConfig()
	if err != nil {
		panic("")
	}
}
