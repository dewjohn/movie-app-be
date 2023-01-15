package utils

import (
	"gorm.io/gorm"
	"movie-app/model"
)

func IsAdminTelephoneExit(db *gorm.DB, telephone string) bool {
	var admin = model.Admin{}
	db.Where("telephone = ?", telephone).First(&admin)
	if admin.ID != 0 {
		return true
	}
	return false
}
