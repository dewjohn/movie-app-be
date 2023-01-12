package main

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"movie-app/common"
	"movie-app/routes"
	"os"
)

func main() {
	InitConfig()
	db := common.InitDB()
	_ = db
	r := gin.Default()
	r = routes.CollectRoute(r)
	port := viper.GetString("server.port")
	if port != "" {
		panic(r.Run(":" + port))
	}
	panic(r.Run()) // listen and  serve on 0.0.0.0:8080
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
