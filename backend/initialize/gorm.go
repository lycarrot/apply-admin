package initialize

import (
	"gin-pro/global"
	"gin-pro/model/example"
	"gin-pro/model/system"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"os"
)

func Gorm() *gorm.DB {
	return GormMysql()
}

func RegisterTables() {
	db := global.GVA_DB
	//确保数据库中的表结构与 GORM 模型定义中的结构保持同步。
	err := db.AutoMigrate(
		system.SysUser{},
		system.SysAuthority{},
		system.SysApi{},
		system.SysBaseMenu{},
		system.JwtBlacklist{},
		example.ExaCustomer{},
	)
	if err != nil {
		global.GVA_LOG.Error("Register database error", zap.Error(err))
		os.Exit(0)
	}
	global.GVA_LOG.Info("Register database success")
}
