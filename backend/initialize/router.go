package initialize

import (
	"gin-pro/docs"
	"gin-pro/global"
	"gin-pro/middleware"
	"gin-pro/router"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
	"os"
)

// 实现了 http.FileSystem 接口
type justFilesFilesystem struct {
	fs http.FileSystem
}

func (fs justFilesFilesystem) Open(name string) (http.File, error) {
	f, err := fs.fs.Open(name)
	if err != nil {
		return nil, err
	}
	stat, err := f.Stat()
	if stat.IsDir() {
		return nil, os.ErrPermission
	}
	return f, nil
}

func Routers() *gin.Engine {
	if global.GVA_CONFIG.System.Env == "public" {
		gin.SetMode(gin.ReleaseMode)
	}
	Router := gin.New()
	//在处理每一个请求时，都会先执行这个中间件，确保请求处理过程中的 panic 能够被正确地恢复。
	Router.Use(gin.Recovery())
	if global.GVA_CONFIG.System.Env != "public" {
		//记录请求的详细信息，并输出到标准输出。
		Router.Use(gin.Logger())
	}

	SystemRouter := router.RouterGroupApp.System
	ExampleRouter := router.RouterGroupApp.Example
	//静态文件配置
	Router.StaticFS(global.GVA_CONFIG.Local.StorePath, justFilesFilesystem{http.Dir(global.GVA_CONFIG.Local.StorePath)})
	//跨域配置
	Router.Use(middleware.Cors())
	Router.Use(middleware.CorsByRules())

	docs.SwaggerInfo.BasePath = global.GVA_CONFIG.System.RouterPrefix
	Router.GET(global.GVA_CONFIG.System.RouterPrefix+"/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	global.GVA_LOG.Info("register swagger handler")

	PublicGroup := Router.Group(global.GVA_CONFIG.System.RouterPrefix)
	PrivateGroup := Router.Group(global.GVA_CONFIG.System.RouterPrefix)

	{
		PublicGroup.GET("/health", func(c *gin.Context) {
			c.JSON(http.StatusOK, "ok")
		})
	}
	{
		// 注册基础功能路由 不做鉴权
		SystemRouter.InitAuthRouter(PublicGroup)
	}
	//PrivateGroup := Router.Group(global.GVA_CONFIG.System.RouterPrefix)
	//middleware.JWTAuth()).Use(middleware.CasbinHandler()
	PrivateGroup.Use()
	{
		ExampleRouter.InitCustomerRouter(PrivateGroup)
	}
	global.GVA_LOG.Info("router register success")
	return Router
}
