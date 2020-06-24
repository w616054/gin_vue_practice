package controller

import (
	"gin_vue_practice/utils"
	"github.com/gin-gonic/gin"
	"gin_vue_practice/common"
	"gin_vue_practice/model"
	"golang.org/x/crypto/bcrypt"
	"net/http"
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
		ctx.JSON(400, gin.H{"msg": "密码最短6位"})
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

	hasedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"code": 500, "msg": "内部错误"})
	}

	// 创建用户
	var tmpUser = model.User{Name: name, Password: string(hasedPassword), Telephone: telephone}
	DB.Create(&tmpUser)
	ctx.JSON(200, gin.H{
		"message": "注册成功",
	})
}

func Login(ctx *gin.Context) {
	var user model.User
	DB := common.GetDB()
	telephone := ctx.PostForm("telephone")
	password := ctx.PostForm("password")

	// 数据验证
	if len(telephone) != 11 {
		ctx.JSON(400, gin.H{"msg": "手机号格式错误"})
		return
	}

	if len(password) < 6 {
		ctx.JSON(400, gin.H{"msg": "密码最短6位"})
		return
	}

	// 判断手机号是否存在
	DB.Where("telephone = ?", telephone).First(&user)
	if user.ID == 0 {
		ctx.JSON(400, gin.H{"msg": "该手机号不存在"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		ctx.JSON(400, gin.H{"msg": "密码错误"})
		return
	}

	// 发功token
	token := "11dasd"

	ctx.JSON(200, gin.H{
		"code": 200,
		"msg":  "登陆成功",
		"data": gin.H{"token": token},
	})
}
