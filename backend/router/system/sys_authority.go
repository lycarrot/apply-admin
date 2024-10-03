package system

import "github.com/gin-gonic/gin"

type AuthorityRouter struct {
}

func (a *AuthorityRouter) InitAuthorityRouter(Router *gin.RouterGroup) {
	router := Router.Group("authority")
	routerWithout := Router.Group("authority")
	{
		router.POST("create", authorityApi.CreateAuthority) //创建角色
	}
	{
		routerWithout.POST("lists", authorityApi.GetAuthorityList) //获取角色列表
	}
}
