package model

import (
	"time"

	"gorm.io/gorm"
)

type Movie struct {
	gorm.Model
	AdminId      uint      `gorm:"not null;index"` // 上传者ID
	Title        string    `gorm:"type:varchar(50);not null;index"`
	Cover        string    `gorm:"size:255;not null"`
	Videos       string    `gorm:"size:255;"`            // 先用字符串视频链接，下一版本引入本地视频
	ReleaseTime  time.Time `gorm:"default:'1970-01-01'"` // 上映时间
	SheetLength  int       `gorm:"default:0"`            // 片长
	Origin       string    `gorm:"varchar(50);"`         // 产地
	Type         string    `gorm:"varchar(50);"`         // 种类
	Director     string    `gorm:"varchar(50);"`         // 导演
	Screenwriter string    `gorm:"varchar(50);"`         // 编剧
	Actors       string    `gorm:"varchar(50);"`         // 演员
	Language     string    `gorm:"varchar(50);"`         // 语言
	Introduction string    `gorm:"size:255;"`            // 简介
	Score        float64   `gorm:"default:0"`            // 评分
}
