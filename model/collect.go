package model

import "gorm.io/gorm"

type Collect struct {
	gorm.Model
	Uid uint `gorm:"not null;index"`
	Vid uint `gorm:"not not"`
}
