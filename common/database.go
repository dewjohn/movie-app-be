package common

import (
	"movie-app/model"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() *gorm.DB {
	user := viper.GetString("datasource.username")
	password := viper.GetString("datasource.password")
	database := viper.GetString("datasource.database")
	charset := viper.GetString("datasource.charset")
	dsn := user + ":" + password + "@tcp(127.0.0.1:3306)/" + database + "?charset=" + charset + "&parseTime=True&loc=Local"
	//dsn := "root:Zdw11.11.11@tcp(127.0.0.1:3306)/movie-app?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("fail to connect database, err:" + err.Error())
	}

	// 自动新建表 AutoMigrate 用于自动迁移您的 schema，保持您的 schema 是最新的。
	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.Admin{})
	db.AutoMigrate(&model.Movie{})
	DB = db
	return db
}

func GetDB() *gorm.DB {
	return DB
}
