package router

import (
	"github.com/gin-gonic/gin"
	_ "go-micro.dev/v4"
)

func NewRouter() *gin.Engine {
	r := gin.New()

	userControllerV1 := new(v1.User)

	// 路由分组
	v1 := r.Group("/v1")
	{
		v1.Use(gin.Logger(), gin.Recovery())
		userV1 := v1.Group("/user")
		{
			userV1.POST("/login", userControllerV1.Login)
		}
	}

	return r
}
