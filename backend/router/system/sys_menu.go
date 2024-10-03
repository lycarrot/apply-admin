package system

import "github.com/gin-gonic/gin"

type MenuRouter struct {
}

func (m *MenuRouter) InitMenuRouter(Router *gin.RouterGroup) {
	router := Router.Group("menu")
	routerWithout := Router.Group("menu")
	{
		router.POST("addItem", menuApi.AddMenuItem)   //新增菜单
		router.POST("del", menuApi.DelMenuItem)       //删除菜单
		router.POST("update", menuApi.UpdateMenuItem) //更新菜单
	}
	{
		routerWithout.GET("getMenuList", menuApi.GetMenuLists) //获取菜单信息
	}
}
