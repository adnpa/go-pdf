package main

import (
	"fmt"
	"github.com/adnpa/gpdf/app/pdf/service"
	conf "github.com/adnpa/gpdf/config"
	"github.com/adnpa/gpdf/proto/pb"
	"github.com/go-micro/plugins/v4/registry/etcd"
	"go-micro.dev/v4"
	"go-micro.dev/v4/registry"
)

func main() {
	cfg := conf.Cfg.EtcdConfig
	etcdReg := etcd.NewRegistry(
		registry.Addrs(fmt.Sprintf("%s:%d", cfg.Host, cfg.Port)),
	)
	// 得到一个微服务实例
	microService := micro.NewService(
		micro.Name("rpcPdfService"), // 微服务名字
		micro.Address(conf.Cfg.PdfSrvAddress),
		micro.Registry(etcdReg), // etcd注册件
	)
	// 结构命令行参数，初始化
	microService.Init()
	// 服务注册
	_ = pb.RegisterPdfServiceHandler(microService.Server(), service.GetPdfSrv())
	// 启动微服务
	_ = microService.Run()
}
