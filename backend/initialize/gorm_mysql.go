package initialize

import (
	"gin-pro/global"
	"gin-pro/initialize/internal"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// GormMysql 初始化Mysql数据库
func GormMysql() *gorm.DB {
	m := global.GVA_CONFIG.Mysql
	if m.Dbname == "" {
		return nil
	}
	mysqlConfig := mysql.Config{
		//数据库连接字符串，包含用户名、密码、主机、端口和数据库名称等信息。
		DSN: m.Dsn(),
		//置默认的字符串字段长度,默认255
		DefaultStringSize: 255,
		//控制 GORM 是否跳过与 MySQL 版本相关的初始化
		SkipInitializeWithVersion: false,
	}
	if db, err := gorm.Open(mysql.New(mysqlConfig), internal.Gorm.Config(m.Prefix, m.Singular)); err != nil {
		return nil
	} else {
		db.InstanceSet("gorm:table_options", "ENGINE="+m.Engine)
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(m.MaxIdleConns)
		sqlDB.SetMaxOpenConns(m.MaxOpenConns)
		return db
	}
}

//// 初始化Mysql数据库用过传入配置
//func GormMysqlByConfig(m config.Mysql) *gorm.DB {
//	if m.Dbname == "" {
//		return nil
//	}
//	mysqlConfig := mysql.Config{
//		DSN: m.Dsn(),
//		//置默认的字符串字段长度,默认255
//		DefaultStringSize: 255,
//		//控制 GORM 是否跳过与 MySQL 版本相关的初始化
//		SkipInitializeWithVersion: false,
//	}
//	if db, err := gorm.Open(mysql.New(mysqlConfig), internal.Gorm.Config(m.Prefix, m.Singular)); err != nil {
//		panic(err)
//	} else {
//		db.InstanceSet("gorm:table_options", "ENGINE=InnoDB")
//		sqlDB, _ := db.DB()
//		sqlDB.SetMaxIdleConns(m.MaxIdleConns)
//		sqlDB.SetMaxOpenConns(m.MaxOpenConns)
//		return db
//	}
//}
