package service

import (
	"context"
	"errors"
	"github.com/adnpa/gpdf/app/user/model"
	"github.com/adnpa/gpdf/app/user/query"
	"github.com/adnpa/gpdf/proto/pb"
	"github.com/adnpa/gpdf/types"
	"gorm.io/gorm"
	"sync"
)

type UserSrv struct{}

var UserSrvIns *UserSrv
var UserSrvOnce sync.Once

func GetUserSrv() *UserSrv {
	UserSrvOnce.Do(func() {
		UserSrvIns = &UserSrv{}
	})
	return UserSrvIns
}

func (u *UserSrv) Login(ctx context.Context, req *pb.LoginReq, resp *pb.LoginResp) (err error) {
	resp.Code = types.SUCCESS
	user, err := query.NewUserDao(ctx).FindUserByUserName(req.Name)
	if err != nil {
		resp.Code = types.ERROR
		return
	}

	if !user.CheckPassword(req.Password) {
		resp.Code = types.InvalidParams
		return
	}

	resp.User = BuildUser(user)
	return
}

func (u *UserSrv) Signup(ctx context.Context, req *pb.SignUpReq, resp *pb.SignUpResp) (err error) {
	if req.Password != req.PasswordConfirm {
		err = errors.New("两次密码输入不一致")
		return
	}
	resp.Code = types.SUCCESS
	_, err = query.NewUserDao(ctx).FindUserByUserName(req.Name)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) { // 如果不存在就继续下去
			// ...continue
		} else {
			resp.Code = types.ERROR
			return
		}
	}
	user := &model.User{
		Name: req.Name,
	}
	// 加密密码
	if err = user.SetPassword(req.Password); err != nil {
		resp.Code = types.ERROR
		return
	}
	if err = query.NewUserDao(ctx).CreateUser(user); err != nil {
		resp.Code = types.ERROR
		return
	}

	resp.User = BuildUser(user)
	return
}

func BuildUser(item *model.User) *pb.UserModel {
	userModel := pb.UserModel{
		Id:        uint32(item.ID),
		UserName:  item.Name,
		CreatedAt: item.CreatedAt.Unix(),
		UpdatedAt: item.UpdatedAt.Unix(),
	}
	return &userModel
}
