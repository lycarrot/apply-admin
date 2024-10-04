package system

import "github.com/gin-gonic/gin"

type CasbinRouter struct {
}

func (c *CasbinRouter) InitCasbinRouter(Router *gin.Engine) {
	router := Router.Group("casbin")
	{
		router.POST("update", casbinApi.UpdateCasbin)
	}
}
