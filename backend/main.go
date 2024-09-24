package main

import (
	"gin-pro/core"
	"gin-pro/global"
	"gin-pro/initialize"
)

//go:generate go env -w GO111MODULE=on
//go:generate go env -w GOPROXY=https://goproxy.cn,direct
//go:generate go mod tidy
//go:generate go mod download

// @title gin-apply-admin API
// @version 1.0
// @description  API
// @host localhost:8080
// @securityDefinitions.apikey  ApiKeyAuth
// @in                          header
// @name                        x-token
// @BasePath /v1

func main() {
	//将配置文件的参数读取到全局挂载
	global.GVA_VP = core.Viper()
	// 初始化本地缓存对象
	initialize.OtherInit()
	global.GVA_LOG = core.Zap()
	//连接数据库
	global.GVA_DB = initialize.Gorm()
	if global.GVA_DB != nil {
		//初始化数据库表
		initialize.RegisterTables()
		db, _ := global.GVA_DB.DB()
		defer db.Close()
	}

	core.RunWindowsServer()
}
