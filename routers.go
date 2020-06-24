package main

import (
	"gin_vue_practice/controller"
	"github.com/gin-gonic/gin"
)

func CollectRouter(r *gin.Engine) *gin.Engine {
	r.POST("/api/user/register", controller.Register)
	r.POST("/api/user/login", controller.Login)
	return  r
}
