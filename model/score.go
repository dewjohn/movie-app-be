package model

import "gorm.io/gorm"

type Score struct {
	gorm.Model
	Vid   uint    `gorm:"not not;index"`
	Grade float64 `gorm:"default:0"`
	Uid   uint    `gorm:"not null"`
}
