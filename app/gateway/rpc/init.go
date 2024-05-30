package rpc

import (
	"github.com/adnpa/gpdf/app/gateway/wrappers"
	"github.com/adnpa/gpdf/proto/pb"
	"go-micro.dev/v4"
)

var (
	UserService pb.UserService
	PdfService  pb.PdfService
)

func InitRPC() {
	// 用户
	userMicroService := micro.NewService(
		micro.Name("userService.client"),
		micro.WrapClient(wrappers.NewUserWrapper),
	)
	// 用户服务调用实例
	UserService = pb.NewUserService("rpcUserService", userMicroService.Client())
	// pdf工具服务
	pdfService := micro.NewService(
		micro.Name("taskService.client"),
		micro.WrapClient(wrappers.NewPdfWrapper),
	)
	PdfService = pb.NewPdfService("rpcPdfService", pdfService.Client())
}
