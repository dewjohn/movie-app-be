package model

import (
	"gorm.io/gorm"
)

// authority取值1000为审核员，取值2000为管理员

type Admin struct {
	gorm.Model
	Name      string `gorm:"type:varchar(20); not null"`
	Email     string `gorm:"varchar(20);not null;index"`
	Telephone string `gorm:"type:varchar(11); not null;unique"`
	Password  string `gorm:"size:225;not null"`
	Authority int    `gorm:"not null"`
}
