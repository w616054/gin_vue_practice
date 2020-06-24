package middleware

import (
	"gin_vue_practice/common"
	"gin_vue_practice/model"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strings"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 获取authorization header
		tokenString := ctx.GetHeader("Authorization")

		// 验证token 格式
		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer") {
			ctx.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "未认证的请求"})
			ctx.Abort()
			return
		}
		tokenString = tokenString[7:]

		token, claims, err := common.ParseToken(tokenString)
		if err != nil || ! token.Valid {
			log.Println(err)
			log.Println(token.Valid)

			ctx.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg":"未认证"})
			ctx.Abort()
			return
		}

		// 获取claim 中的userId
		userId := claims.UserId

		DB := common.GetDB()
		user := &model.User{}
		DB.Where("id = ?", userId).First(&user) // or DB.First(&user, userId)
		if user.ID == 0 {  // 用户不存在
			ctx.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "未认证"})
			ctx.Abort()
			return
		}
		// 用户存在将用户信息写入上下文
		ctx.Set("user", user)
		ctx.Next()
	}
}