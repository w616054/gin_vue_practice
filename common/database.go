package common

import (
	"fmt"
	"gin_vue_practice/model"
	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

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
		db.AutoMigrate(&model.User{})
		DB = db
		return DB
	}
}

func GetDB() *gorm.DB{
	return DB
}
