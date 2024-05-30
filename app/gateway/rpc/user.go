package rpc

import (
	"context"
	"github.com/adnpa/gpdf/enums"
	"github.com/adnpa/gpdf/proto/pb"
)

// Login UserLogin 用户登陆
func Login(ctx context.Context, req *pb.LoginReq) (resp *pb.LoginResp, err error) {
	resp, err = UserService.Login(ctx, req)
	if err != nil {
		return
	}

	if resp.Code != enums.SUCCESS {

		return
	}

	return
}

// Signup UserRegister 用户注册
func Signup(ctx context.Context, req *pb.SignUpReq) (resp *pb.SignUpResp, err error) {
	resp, err = UserService.Signup(ctx, req)
	if err != nil {
		return
	}
	return
}
