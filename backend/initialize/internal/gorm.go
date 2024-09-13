package internal

import (
	"gin-pro/global"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"os"
	"time"
)

type _gorm struct {
}

var Gorm = new(_gorm)

// 空接口
type DBBASE interface {
	GetLogMode() string
}

func (g *_gorm) Config(prefix string, singular bool) *gorm.Config {
	config := &gorm.Config{
		//定义数据库表名和字段名的命名规则。
		NamingStrategy: schema.NamingStrategy{
			//表前缀
			TablePrefix: prefix,
			//指定是否使用单数形式的表名。默认情况下，GORM 会使用复数形式的表名，如果设置为 true，则会使用单数形式的表名。
			SingularTable: singular,
		},
		//在迁移表结构时禁用外键约束。
		DisableForeignKeyConstraintWhenMigrating: true,
	}
	//配置 GORM 的日志记录器
	_default := logger.New(log.New(os.Stdout, "\r\n", log.LstdFlags), logger.Config{
		//设置慢查询的阈值时间为 200 毫秒。如果一次查询的执行时间超过了这个阈值，会被记录为慢查询。
		SlowThreshold: 200 * time.Millisecond,
		//设置日志级别为 logger.Warn，表示只记录警告级别及以上的日志消息。
		LogLevel: logger.Warn,
		Colorful: true,
	})

	var logMode DBBASE
	switch global.GVA_CONFIG.System.DbType {
	case "mysql":
		logMode = &global.GVA_CONFIG.Mysql
	default:
		logMode = &global.GVA_CONFIG.Mysql
	}
	switch logMode.GetLogMode() {
	//静默模式，即不记录任何日志。
	case "silent", "Silent":
		config.Logger = _default.LogMode(logger.Silent)
		//错误模式，记录错误级别及以上的日志消息。
	case "error", "Error":
		config.Logger = _default.LogMode(logger.Error)
		//警告模式，记录警告级别及以上的日志消息。
	case "warn", "Warn":
		config.Logger = _default.LogMode(logger.Warn)
		//信息模式，记录所有日志消息，包括 SQL 语句和执行时间。
	case "info", "Info":
		config.Logger = _default.LogMode(logger.Info)
	default:
		config.Logger = _default.LogMode(logger.Silent)
	}
	return config
}
