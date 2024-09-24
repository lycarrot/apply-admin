package core

import (
	"fmt"
	"gin-pro/global"
	"gin-pro/initialize"
	"gin-pro/service/system"
	"go.uber.org/zap"
)

type server interface {
	ListenAndServe() error
}

func RunWindowsServer() {
	if global.GVA_CONFIG.System.UseMultipoint || global.GVA_CONFIG.System.UseRedis {
		initialize.Redis()
	}

	//if global.GVA_CONFIG.System.UseMongo {
	//}

	if global.GVA_DB != nil {
		system.LoadAll()
	}

	//初始化路由
	Router := initialize.Routers()
	Router.Static("/form-generator", "./resource/page")
	address := fmt.Sprintf(":%d", global.GVA_CONFIG.System.Addr)
	s := initServer(address, Router)
	global.GVA_LOG.Info("server run success on", zap.String("address", address))

	fmt.Print(`
	欢迎使用 gin-apply-admin
`, address)

	global.GVA_LOG.Error(s.ListenAndServe().Error())
}
