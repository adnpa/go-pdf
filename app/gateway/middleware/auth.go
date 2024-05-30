package middleware

import (
	"fmt"
	"github.com/adnpa/gpdf/app/gateway/handler"
	"github.com/adnpa/gpdf/enums"
	"github.com/adnpa/gpdf/pkg/utils"
	"github.com/gin-gonic/gin"
)

// JWTAuthMiddleware 基于jwt的授权中间件
func JWTAuthMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		aToken, _ := handler.ParseJwtHeader(c)
		// 处理
		claims, err := utils.ParseToken(aToken)
		if err != nil {
			fmt.Println(err)
			handler.ResponseError(c, enums.CodeInvalidToken)
			c.Abort()
			return
		}
		// 将当前请求的userID信息保存到请求的上下文c上
		c.Set(handler.ContextUserIDKey, claims.UserID)
		c.Next() // 后续的处理函数可以用过c.Get(ContextUserIDKey)来获取当前请求的用户信息
	}
}
