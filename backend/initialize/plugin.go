package initialize

import (
	"fmt"
	"gin-pro/utils/plugin"
	"github.com/gin-gonic/gin"
)

func PluginInit(group *gin.RouterGroup, Plugin ...plugin.Plugin) {
	for i := range Plugin {
		PluginGroup := group.Group(Plugin[i].RouterPath())
		Plugin[i].Register(PluginGroup)
	}
}
func InstallPlugin(Router *gin.Engine) {
	//用于创建一个新的路由组
	PublicGroup := Router.Group("")
	fmt.Println("无鉴权插件安装==》", PublicGroup)
	PrivateGroup := Router.Group("")
	fmt.Println("鉴权插件安装==》", PrivateGroup)
	//PrivateGroup.Use(middleware.JWTAuth())
	//PluginInit(PrivateGroup, email)
}
