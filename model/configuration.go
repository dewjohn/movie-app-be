package model

import "gorm.io/gorm"

type Configuration struct {
	gorm.Model
	Key   string `gorm:"type:varchar(50);not null;index"`
	Value string `gorm:"type:varchar(255);not null"`
}
