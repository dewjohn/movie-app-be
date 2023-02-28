package model

import "gorm.io/gorm"

type Reply struct {
	gorm.Model
	Cid       uint   `gorm:"not null;index"`             // 评论的id
	Content   string `gorm:"type:varchar(255);not null"` // 回复内容
	Uid       uint   `gorm:"not null"`                   // 用户
	ReplyUid  uint   // 回复的人的uid
	ReplyName string `gorm:"type:varchar(20)"` // 回复的人的昵称
}
