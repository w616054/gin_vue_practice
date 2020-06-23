package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"gin_vue_practice/common"
)



func main() {
	db := common.InitDB()
	defer db.Close()

	r := gin.Default()
	CollectRouter(r)
	r.Run(":8081") // listen and serve on 0.0.0.0:8080
}



