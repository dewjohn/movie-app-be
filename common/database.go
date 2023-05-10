package common

import (
	"fmt"
	"movie-app/model"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() *gorm.DB {
	user := viper.GetString("datasource.username")
	password := viper.GetString("datasource.password")
	host := viper.GetString("datasource.host")
	port := viper.GetString("datasource.port")
	database := viper.GetString("datasource.database")
	charset := viper.GetString("datasource.charset")
	// dsn := user + ":" + password + "@tcp(127.0.0.1:3306)/" + database + "?charset=" + charset + "&parseTime=True&loc=Local"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true",
		user,
		password,
		host,
		port,
		database,
		charset)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("fail to connect database, err:" + err.Error())
	}

	db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci")
	// 自动新建表 AutoMigrate 用于自动迁移您的 schema
	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.Admin{})
	db.AutoMigrate(&model.Movie{})
	db.AutoMigrate(&model.Resource{})
	db.AutoMigrate(&model.Comment{})
	db.AutoMigrate(&model.Score{})
	db.AutoMigrate(&model.Collect{})
	db.AutoMigrate(model.Configuration{})
	DB = db
	return db
}

func GetDB() *gorm.DB {
	return DB
}
