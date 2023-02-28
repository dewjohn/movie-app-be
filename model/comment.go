package model

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	Vid        uint   `gorm:"not not;index"`              // 电影id
	Content    string `gorm:"type:varchar(255);not null"` // 评论内容
	Uid        uint   `gorm:"not null"`                   // 评论人id
	ReplyCount int    `gorm:"default:0"`                  // 回复数
}
