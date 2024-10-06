package system

import (
	"gin-pro/middleware"
	"github.com/gin-gonic/gin"
)

type MenuRouter struct {
}

func (m *MenuRouter) InitMenuRouter(Router *gin.RouterGroup) {
	router := Router.Group("menu").Use(middleware.OperationRecordHandler())
	routerWithoutRecord := Router.Group("menu")
	{
		router.POST("create", menuApi.CreateMenuItem) //新增菜单
		router.POST("del", menuApi.DelMenuItem)       //删除菜单
		router.POST("update", menuApi.UpdateMenuItem) //更新菜单
	}
	{
		routerWithoutRecord.GET("getMenuList", menuApi.GetMenuLists) //获取菜单信息
	}
}
