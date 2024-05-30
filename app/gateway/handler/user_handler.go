package handler

import (
	"github.com/adnpa/gpdf/app/gateway/rpc"
	"github.com/adnpa/gpdf/enums"
	"github.com/adnpa/gpdf/pkg/utils"
	"github.com/adnpa/gpdf/proto/pb"
	"github.com/gin-gonic/gin"
	"strings"
)

// SignupHandler UserRegisterHandler 用户注册
func SignupHandler(ctx *gin.Context) {
	var req pb.SignUpReq

	if err := ctx.ShouldBind(&req); err != nil {
		ResponseError(ctx, enums.CodeInvalidParams)
		return
	}

	userResp, err := rpc.Signup(ctx, &req)
	if err != nil {
		ResponseError(ctx, enums.CodeServerBusy)
		return
	}

	ResponseSuccess(ctx, userResp)
}

// LoginHandler UserLoginHandler 用户登录
func LoginHandler(ctx *gin.Context) {
	var req pb.LoginReq
	if err := ctx.Bind(&req); err != nil {
		ResponseError(ctx, enums.CodeInvalidParams)
		return
	}

	user, err := rpc.Login(ctx, &req)
	if err != nil {
		ResponseError(ctx, enums.CodeServerBusy)
		return
	}

	aToken, _, err := utils.GenToken(uint64(user.User.Id), user.User.UserName)
	if err != nil {
		ResponseError(ctx, enums.CodeServerBusy)
	}

	ResponseSuccess(ctx, &enums.TokenData{
		Token: aToken,
		User:  user.User,
	})
}

// ParseJwtHeader 辅助函数
func ParseJwtHeader(c *gin.Context) (aToken, rToken string) {
	//rt := c.Query("refresh_token")
	// 客户端携带Token有三种方式 1.放在请求头 2.放在请求体 3.放在URI
	// 这里假设Token放在Header的 Authorization 中，并使用 Bearer 开头
	// 这里的具体实现方式要依据你的实际业务情况决定
	authHeader := c.Request.Header.Get("Authorization")
	if authHeader == "" {
		ResponseErrorWithMsg(c, enums.CodeInvalidToken, "请求头缺少Auth Token")
		c.Abort()
		return
	}
	//	空格分割aToken和rToken
	parts := strings.SplitN(authHeader, " ", 2)
	if !(len(parts) == 2 && parts[0] == "Bearer") {
		ResponseErrorWithMsg(c, enums.CodeInvalidToken, "Token格式不对")
		c.Abort()
		return
	}
	aToken = parts[1]
	return
}
