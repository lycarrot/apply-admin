package request

import (
	"fmt"
	"gin-pro/config"
)

type InitDB struct {
	DBType   string `json:"dbType"`
	Host     string `json:"host"`                      // 服务器地址
	Port     string `json:"port"`                      // 数据库连接端口
	UserName string `json:"userName"`                  // 数据库用户名
	Password string `json:"password"`                  // 数据库密码
	DBName   string `json:"dbName" binding:"required"` // 数据库名
	DBPath   string `json:"dbPath"`                    // sqlite数据库文件路径
}

func (i *InitDB) ToMysqlConfig() config.Mysql {
	return config.Mysql{
		GeneralDB: config.GeneralDB{
			Path:         i.Host,
			Port:         i.Port,
			UserName:     i.UserName,
			Password:     i.Password,
			Dbname:       i.DBName,
			MaxIdleConns: 10,
			MaxOpenConns: 100,
			LogMode:      "error",
			Config:       "charset=utf8mb4&parseTime=True&loc=Local",
		},
	}
}

func (i *InitDB) MysqlEmptyDsn() string {
	if i.Host == "" {
		i.Host = "127.0.0.1"
	}
	if i.Port == "" {
		i.Port = "3306"
	}
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/", i.UserName, i.Password, i.Host, i.Port)
}
