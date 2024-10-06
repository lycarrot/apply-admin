package system

import (
	"gin-pro/middleware"
	"github.com/gin-gonic/gin"
)

type CasbinRouter struct {
}

func (c *CasbinRouter) InitCasbinRouter(Router *gin.RouterGroup) {
	router := Router.Group("casbin").Use(middleware.OperationRecordHandler())
	{
		router.POST("update", casbinApi.UpdateCasbin)
	}
}
