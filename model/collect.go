package model

import "gorm.io/gorm"

type Collect struct {
	gorm.Model
	Uid uint `gorm:"not null;index"`
	Vid int  `gorm:"not not"`
}
