package common

import (
	"fmt"
	"gin_vue_practice/model"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
)

var DB *gorm.DB

func  InitDB() *gorm.DB {
	driverName := viper.GetString("datasource.diverName")
	host := viper.GetString("datasource.host")
	port := viper.GetString("datasource.port")
	dbUser := viper.GetString("datasource.username")
	dbPasspword := viper.GetString("datasource.password")
	dbName := viper.GetString("datasource.database")
	charset := viper.GetString("datasource.charset")
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
