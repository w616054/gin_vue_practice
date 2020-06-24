package controller

import (
	"gin_vue_practice/utils"
	"github.com/gin-gonic/gin"
	"gin_vue_practice/common"
	"gin_vue_practice/model"

)

func Register(ctx *gin.Context) {
	DB := common.GetDB()
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
		name = utils.RandomString(6)
	}

	// 手机号已经存在
	if utils.IsTelePhoneExist(DB, telephone) {
		ctx.JSON(400, gin.H{"msg": "手机号已经存在"})
		return
	}

	var tmpUser = model.User{Name:name, Password: password, Telephone: telephone}
	// 创建用户
	DB.Create(&tmpUser)
	ctx.JSON(200, gin.H{
		"message": "注册成功11111",
	})
}