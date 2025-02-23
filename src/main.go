package main

import (
	userPb "github.com/sunshine-walker-93/phoenix_apis/protobuf3.pb/user_info_manage"
	goMicro "go-micro.dev/v5"

	_ "github.com/sunshine-walker-93/phoenix_account/src/gmysql"
	_ "github.com/sunshine-walker-93/phoenix_account/src/gredis"

	"github.com/micro/plugins/v5/registry/etcd"
	"github.com/sunshine-walker-93/phoenix_account/src/config"
	"github.com/sunshine-walker-93/phoenix_account/src/log"
	"github.com/sunshine-walker-93/phoenix_account/src/service"
	"go-micro.dev/v5/registry"
)

func main() {
	etcdReg := etcd.NewRegistry(
		registry.Addrs(config.ReferGlobalConfig().ServerSetting.RegisterAddress),
	)
	// 初始化服务
	srv := goMicro.NewService(
		goMicro.Name(config.ReferGlobalConfig().ServerSetting.RegisterServerName),
		goMicro.Version(config.ReferGlobalConfig().ServerSetting.RegisterServerVersion),
		goMicro.Registry(etcdReg),
	)

	err := userPb.RegisterUserServiceHandler(srv.Server(), &service.UserService{})
	if err != nil {
		log.Error("RegisterUserServiceHandler failed, err:%s", err)
		return
	}

	err = srv.Run()
	if err != nil {
		log.Error("run failed, err:%s", err)
	}
}
