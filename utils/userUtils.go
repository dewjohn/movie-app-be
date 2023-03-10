package utils

import (
	"golang.org/x/crypto/bcrypt"
	"log"
	"math/rand"
	"movie-app/model"
	"regexp"
	"time"

	"gorm.io/gorm"
)

// RandomString 返回 n 位随机字符
func RandomString(n int) string {
	letters := []byte("qwertyuiop[]';lkjhgfdsazxcvbnm,./?><" +
		"';")
	result := make([]byte, n)
	rand.Seed(time.Now().Unix())
	for i := range result {
		result[i] = letters[rand.Intn(len(letters))]
	}
	return string(result)
}

// VerifyEmailFormat 验证邮箱
func VerifyEmailFormat(email string) bool {
	pattern := `\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*` //匹配电子邮箱
	reg := regexp.MustCompile(pattern)
	return reg.MatchString(email)
}

// IsTelephoneExist 判断手机号是否存在
func IsTelephoneExist(db *gorm.DB, telephone string) bool {
	var user model.User
	db.Where("telephone = ?", telephone).First(&user)
	if user.ID != 0 {
		return true
	}
	return false
}

// ComparePasswords 验证加密后的密码
func ComparePasswords(hashedPwd string, plainPwd []byte) bool {
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPwd)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}

// 判断该用户对这部电影是否评论过
func IsReviewedMovie(db *gorm.DB, vid uint, uid interface{}) bool {
	var reviewer model.Score
	db.Where("vid = ? and uid = ?", vid, uid).First(&reviewer)
	if reviewer.ID != 0 {
		return true
	}
	return false
}
