package model

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	Vid      uint   `gorm:"not not;index"`              // 电影id
	Content  string `gorm:"type:varchar(255);not null"` // 评论内容
	Uid      uint   `gorm:"not null"`                   // 评论人id
	ParentId uint   `gorm:"default:0"`                  // 0为评论，其他为回复
}
