package system

import (
	"context"
	"database/sql"
	"errors"
	"gin-pro/global"
	"gin-pro/model/system/request"
	"gorm.io/gorm"
	"sort"
)

const (
	Mysql           = "mysql"
	Pgsql           = "pgsql"
	Sqlite          = "sqlite"
	Mssql           = "mssql"
	InitSuccess     = "\n[%v] --> 初始数据成功!\n"
	InitDataExist   = "\n[%v]-->%v的数据库已存在!\n"
	InitDataFailed  = "\n[%v] --> %v 初始数据失败! \nerr: %+v\n"
	InitDataSuccess = "\n[%v] --> %v 初始数据成功!\n"
)

var (
	ErrMissingDBContext        = errors.New("missing db in context")
	ErrMissingDependentContext = errors.New("missing dependent value in context")
	ErrDBTypeMismatch          = errors.New("db type mismatch")
)

type SubInitializer interface {
	InitializerName() string
	MigrateTable(ctx context.Context) (next context.Context, err error)
	InitializeData(ctx context.Context) (next context.Context, err error)
	TableCreated(ctx context.Context) bool
	DataInserted(ctx context.Context) bool
}

type TypedDBInitHandler interface {
	EnsureDB(ctx context.Context, conf *request.InitDB) (context.Context, error)
	WriteConfig(ctx context.Context) error
	InitTables(ctx context.Context, inits initSlice) error
	InitData(ctx context.Context, inits initSlice) error
}

type orderedInitializer struct {
	order int
	SubInitializer
}

type initSlice []*orderedInitializer

var (
	initializers initSlice
	cache        map[string]*orderedInitializer
)

type InitDBService struct {
}

func (s *InitDBService) InitDB(conf request.InitDB) (err error) {
	ctx := context.TODO()
	if len(initializers) == 0 {
		return errors.New("无可用初始化过程，请检查初始化是否已执行完成")
	}
	sort.Sort(&initializers)
	var initHandler TypedDBInitHandler
	switch conf.DBType {
	case "mysql":
		initHandler = NewMysqlInitHandler()
		ctx = context.WithValue(ctx, "dbtype", "mysql")
	default:
		initHandler = NewMysqlInitHandler()
		ctx = context.WithValue(ctx, "dbtype", "mysql")
	}
	ctx, err = initHandler.EnsureDB(ctx, &conf)
	if err != nil {
		return err
	}
	db := ctx.Value("db").(*gorm.DB)
	global.GVA_DB = db
	if err = initHandler.InitTables(ctx, initializers); err != nil {
		return err
	}
	if err = initHandler.InitData(ctx, initializers); err != nil {
		return err
	}
	if err = initHandler.WriteConfig(ctx); err != nil {
		return err
	}
	if err = initHandler.WriteConfig(ctx); err != nil {
		return err
	}
	initializers = initSlice{}
	cache = map[string]*orderedInitializer{}
	return nil
}

// 创建数据库连接和创建数据库
func createDataBase(dsn string, driver string, createSql string) error {
	db, err := sql.Open(driver, dsn)
	if err != nil {
		return err
	}
	defer func(db *sql.DB) {
		err = db.Close()
	}(db)
	if err = db.Ping(); err != nil {
		return err
	}
	_, err = db.Exec(createSql)
	return err
}

func createTables(ctx context.Context, inits initSlice) error {
	next, cancel := context.WithCancel(ctx)
	defer func(c func()) {
		c()
	}(cancel)
	for _, init := range inits {
		if init.TableCreated(next) {
			continue
		}
		if n, err := init.MigrateTable(next); err != nil {
			return err
		} else {
			next = n
		}
	}
	return nil
}

func (a initSlice) Len() int {
	return len(a)
}

func (a initSlice) Less(i, j int) bool {
	return a[i].order < a[j].order
}
func (a initSlice) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}
