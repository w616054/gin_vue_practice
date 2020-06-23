package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"math/rand"
	"time"
)

type User struct {
	gorm.Model
	Name string 		`gorm:"type:varchar(20);not null"`
	Password string		`gorm:"type:varchar(100);not null"`
	Telephone string	`gorm:"type:varchar(11);not null"`
}

func main() {
	db := InitDB()
	defer db.Close()

	r := gin.Default()
	r.POST("/api/user/register", func(ctx *gin.Context) {
		// 接受用户数据
		name := ctx.PostForm("name")
		telephone := ctx.PostForm("telephone")
		password := ctx.PostForm("password")

		// 数据验证
		if len(telephone) != 11 {
			ctx.JSON(400, gin.H{"msg": "手机号格式错误"})
			return
		}

		if len(password) < 6 {
			ctx.JSON(400, gin.H{"msg": "密码最短6位",})
			return
		}

		if len(name) == 0 {
			name = RandomString(6)
		}

		// 手机号已经存在
		if isTelePhoneExist(db, telephone) {
			ctx.JSON(400, gin.H{"msg": "手机号已经存在"})
			return
		}

		var tmpUser = User{Name:name, Password: password, Telephone: telephone}
		// 创建用户
		db.Create(&tmpUser)
		ctx.JSON(200, gin.H{
			"message": "注册成功11111",
		})
	})
	r.Run(":8081") // listen and serve on 0.0.0.0:8080
}

func  InitDB() *gorm.DB {
	driverName := "mysql"
	host := "192.168.3.102"
	port := "3306"
	dbPasspword := "root"
	dbUser := "root"
	dbName := "gin_vue"
	charset := "utf8"
	args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true",
		dbUser,
		dbPasspword,
		host,
		port,
		dbName,
		charset,)
	db, err := gorm.Open(driverName, args)
	if err != nil {
		panic("faild to connect mysql, err:" + err.Error())
	}else{
		db.AutoMigrate(&User{})
		return db
	}
}

func RandomString(n int) string {
	var letters = []byte("asdnaksdnasdnknowqieithjnasdnasdnaksd")
	result := make([]byte, n)

	rand.Seed(time.Now().Unix())
	for i := range result {
		result[i] = letters[rand.Intn(len(letters))]
	}

	return string(result)
}

func isTelePhoneExist(db *gorm.DB, telephone string)  bool {
	var user User
	db.Where("telephone = ?", telephone).First(&user)
	if user.ID != 0 {
		return true
	}
	return false
}