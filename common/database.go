package common

import (
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"movie-app/model"
)

var DB *gorm.DB

func InitDB() *gorm.DB {
	user := viper.GetString("datasource.username")
	password := viper.GetString("datasource.password")
	database := viper.GetString("datasource.database")
	charset := viper.GetString("datasource.charset")
	log.Printf(user, password, database, charset)
	dsn := user + ":" + password + "@tcp(127.0.0.1:3306)/" + database + "?charset=" + charset + "&parseTime=True&loc=Local"
	//dsn := "root:Zdw11.11.11@tcp(127.0.0.1:3306)/movie-app?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("fail to connect database, err:" + err.Error())
	}
	// 自动新建表
	db.AutoMigrate(&model.User{})
	DB = db
	return db
}

func GetDB() *gorm.DB {
	return DB
}
