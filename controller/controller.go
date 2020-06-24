package controller

import (
	"fmt"
	"gin_vue_practice/common"
	"gin_vue_practice/dto"
	"gin_vue_practice/model"
	"gin_vue_practice/response"
	"gin_vue_practice/utils"
	"github.com/gin-gonic/gin"
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
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "手机号格式错误")
		return
	}

	if len(password) < 6 {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "密码最短6位")
		return
	}

	if len(name) == 0 {
		name = utils.RandomString(6)
	}

	// 手机号已经存在
	if utils.IsTelePhoneExist(DB, telephone) {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "手机号已经存在")
		return
	}

	hasedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		response.Response(ctx, http.StatusInternalServerError, 500, nil, "内部错误")
	}

	// 创建用户
	var tmpUser = model.User{Name: name, Password: string(hasedPassword), Telephone: telephone}
	DB.Create(&tmpUser)
	response.Success(ctx, nil, "注册成功")
}

func Login(ctx *gin.Context) {
	var user model.User
	DB := common.GetDB()
	telephone := ctx.PostForm("telephone")
	password := ctx.PostForm("password")

	// 数据验证
	if len(telephone) != 11 {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "手机号格式错误")
		return
	}

	if len(password) < 6 {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "密码最短6位")
		return
	}

	// 判断手机号是否存在
	DB.Where("telephone = ?", telephone).First(&user)
	if user.ID == 0 {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "该手机号不存在")
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "密码错误")
		return
	}

	// 发功token
	token, err := common.RleaseToken(user)
	if err != nil {
		response.Response(ctx, http.StatusInternalServerError, 500, nil, "发放jwt错误")
		return
	}

	response.Success(ctx, gin.H{"token": token}, "登陆成功" )
}

func Info(ctx *gin.Context)  {
	user, _ := ctx.Get("user")
	fmt.Println(user)

	response.Success(ctx, gin.H{"code": 200, "data": gin.H{"user": dto.ToUserDto(user.(model.User))}}, "" )
}