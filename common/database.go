package common

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"movie-app/model"
)

var DB *gorm.DB

func InitDB() *gorm.DB {
	dsn := "root:Zdw11.11.11@tcp(127.0.0.1:3306)/movie-app?charset=utf8mb4&parseTime=True&loc=Local"
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
