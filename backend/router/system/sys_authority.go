package system

import (
	"gin-pro/middleware"
	"github.com/gin-gonic/gin"
)

type AuthorityRouter struct {
}

func (a *AuthorityRouter) InitAuthorityRouter(Router *gin.RouterGroup) {
	router := Router.Group("authority").Use(middleware.OperationRecordHandler())
	routerNotRecord := Router.Group("authority")
	{
		router.POST("create", authorityApi.CreateAuthority) //创建角色
	}
	{
		routerNotRecord.POST("getLists", authorityApi.GetAuthorityList) //获取角色列表
	}
}
