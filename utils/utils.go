package utils

import (
	"gin_vue_practice/model"
	"github.com/jinzhu/gorm"
	"math/rand"
	"time"
)

// 随机字符串
func RandomString(n int) string {
	var letters = []byte("asdnaksdnasdnknowqieithjnasdnasdnaksd")
	result := make([]byte, n)

	rand.Seed(time.Now().Unix())
	for i := range result {
		result[i] = letters[rand.Intn(len(letters))]
	}

	return string(result)
}

// 手机号是否存在
func IsTelePhoneExist(db *gorm.DB, telephone string)  bool {
	var user model.User
	db.Where("telephone = ?", telephone).First(&user)
	if user.ID != 0 {
		return true
	}
	return false
}


