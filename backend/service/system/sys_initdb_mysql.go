package system

import (
	"context"
	"errors"
	"fmt"
	"gin-pro/config"
	"gin-pro/global"
	"gin-pro/model/system/request"
	"gin-pro/utils"
	"github.com/gofrs/uuid/v5"
	"github.com/gookit/color"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"path/filepath"
)

type MysqlInitHandler struct {
}

func NewMysqlInitHandler() *MysqlInitHandler {
	return &MysqlInitHandler{}
}

// 创建数据库并初始化 mysql
func (h MysqlInitHandler) EnsureDB(ctx context.Context, conf *request.InitDB) (next context.Context, err error) {
	if s, ok := ctx.Value("dbtype").(string); !ok || s != "mysql" {
		return ctx, ErrDBTypeMismatch
	}
	c := conf.ToMysqlConfig()
	next = context.WithValue(ctx, "config", c)
	if c.Dbname == "" {
		return ctx, nil
	}
	dsn := conf.MysqlEmptyDsn()
	createSql := fmt.Sprintf("CREATE DATABASE IF NOT EXISTS `%s` DEFAULT CHARACTER SET utf8mb4 DEFAULT COLLATE utf8mb4_general_ci;", c.Dbname)
	if err = createDataBase(dsn, "myslq", createSql); err != nil {
		return nil, err
	}
	var db *gorm.DB
	if db, err = gorm.Open(mysql.New(mysql.Config{
		DSN:                       c.Dsn(),
		DefaultStringSize:         256,
		SkipInitializeWithVersion: true,
	}), &gorm.Config{DisableForeignKeyConstraintWhenMigrating: true}); err != nil {
		return ctx, err
	}
	global.GVA_CONFIG.AutoCode.Root, _ = filepath.Abs("..")
	next = context.WithValue(next, "db", db)
	return next, err
}

// mysql回写配置
func (h MysqlInitHandler) WriteConfig(ctx context.Context) error {
	c, ok := ctx.Value("config").(config.Mysql)
	if !ok {
		return errors.New("mysql config invalid")
	}
	global.GVA_CONFIG.System.DbType = "mysql"
	global.GVA_CONFIG.Mysql = c
	global.GVA_CONFIG.Jwt.SigningKey = uuid.Must(uuid.NewV4()).String()
	cs := utils.StructToMap(global.GVA_CONFIG)
	for k, v := range cs {
		global.GVA_VP.Set(k, v)
	}
	return global.GVA_VP.WriteConfig()
}
func (h MysqlInitHandler) InitTables(ctx context.Context, inits initSlice) error {
	return createTables(ctx, inits)

}
func (h MysqlInitHandler) InitData(ctx context.Context, inits initSlice) error {
	next, cancel := context.WithCancel(ctx)
	defer func(c func()) { c() }(cancel)
	for _, init := range inits {
		if init.DataInserted(next) {
			color.Info.Printf(InitDataExist, Mysql, init.InitializerName())
			continue
		}
		if n, err := init.InitializeData(next); err != nil {
			color.Info.Printf(InitDataFailed, Mysql, init.InitializerName(), err)
			return err
		} else {
			next = n
			color.Info.Printf(InitDataSuccess, Mysql, init.InitializerName())
		}
	}
	color.Info.Printf(InitSuccess, Mysql)
	return nil
}
