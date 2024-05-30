package router

import (
	"github.com/adnpa/gpdf/app/gateway/handler"
	"github.com/adnpa/gpdf/app/gateway/middleware"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func SetUpRouter(r *gin.Engine, mode string) {
	if mode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode)
	}

	r.Use(middleware.Cors())
	store := cookie.NewStore([]byte("something-very-secret"))
	r.Use(sessions.Sessions("mysession", store))

	//注册登录
	v1 := r.Group("/api/v1")
	{
		v1.POST("/signup", handler.SignupHandler)
		v1.POST("/login", handler.LoginHandler)

		//授权中间件
		v1.Use(middleware.JWTAuthMiddleware())
		{
			v1.POST("/split", handler.SplitHandler)               //切分
			v1.POST("/merge", handler.MergeHandler)               //合并
			v1.POST("/addwatermark", handler.AddWatermarkHandler) //加水印
		}
	}

	//性能测试
	//pprof.Register(v1)
	// 404
	//r.NoRoute(func(c *gin.Context) {
	//	log.Println("没有找到页面")
	//	c.JSON(http.StatusOK, gin.H{
	//		"msg": "404",
	//	})
	//})
}
