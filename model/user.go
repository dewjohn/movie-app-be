package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Avatar    string    `gorm:"size:255;"`
	Name      string    `gorm:"type:varchar(20); not null"`
	Email     string    `gorm:"varchar(20);not null;index"`
	Telephone string    `gorm:"type:varchar(11); not null;unique"`
	Password  string    `gorm:"size:225;not null"`
	Gender    int       `gorm:"default:0"`
	Birthday  time.Time `gorm:"default:'1970-01-01'"`
	Sign      string    `gorm:"varchar(50);default:'这个人很懒，什么都没有留下'"`
}
