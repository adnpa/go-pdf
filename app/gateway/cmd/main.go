package main

import (
	"fmt"
	"github.com/adnpa/gpdf/app/gateway/router"
	"github.com/adnpa/gpdf/app/gateway/rpc"
	conf "github.com/adnpa/gpdf/config"
	"github.com/adnpa/gpdf/pkg/cache"
	"github.com/gin-gonic/gin"
	"time"

	_ "github.com/adnpa/gpdf/pkg/logger"
	"github.com/go-micro/plugins/v4/registry/etcd"
	"go-micro.dev/v4/registry"
	"go-micro.dev/v4/web"
)

func main() {
	rpc.InitRPC()
	cache.InitCache()
	//log.InitLog()
	etcdReg := etcd.NewRegistry(
		registry.Addrs(fmt.Sprintf("%s:%d", conf.Cfg.EtcdConfig.Host, conf.Cfg.EtcdConfig.Port)),
	)

	r := gin.Default()
	router.SetUpRouter(r, conf.Cfg.Mode)

	// 创建微服务实例，使用gin暴露http接口并注册到etcd
	server := web.NewService(
		web.Name("httpService"),
		web.Address("127.0.0.1:4000"),
		// 将服务调用实例使用gin处理
		web.Handler(r),
		web.Registry(etcdReg),
		web.RegisterTTL(time.Second*30),
		web.RegisterInterval(time.Second*15),
		web.Metadata(map[string]string{"protocol": "http"}),
	)
	// 接收命令行参数
	_ = server.Init()
	_ = server.Run()
}
