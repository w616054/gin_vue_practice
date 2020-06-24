package middleware

import (
	"gin_vue_practice/common"
	"gin_vue_practice/model"
	"gin_vue_practice/response"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 获取authorization header
		tokenString := ctx.GetHeader("Authorization")

		// 验证token 格式
		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer") {
			response.Response(ctx, http.StatusUnauthorized, 401, nil, "未认证的请求")
			ctx.Abort()
			return
		}
		tokenString = tokenString[7:]

		token, claims, err := common.ParseToken(tokenString)
		if err != nil || ! token.Valid {
			response.Response(ctx, http.StatusUnauthorized, 401, nil, "未认证的请求")
			ctx.Abort()
			return
		}

		// 获取claim 中的userId
		userId := claims.UserId

		DB := common.GetDB()
		user := model.User{}
		DB.Where("id = ?", userId).First(&user) // or DB.First(&user, userId)
		if user.ID == 0 {  // 用户不存在
			response.Response(ctx, http.StatusUnauthorized, 401, nil, "未认证的请求")
			ctx.Abort()
			return
		}
		// 用户存在将用户信息写入上下文
		ctx.Set("user", user)
		ctx.Next()
	}
}