package system

import (
	v1 "gin-pro/api/v1"
	"github.com/gin-gonic/gin"
)

type InitRouter struct {
}

func (s *InitRouter) InitInitRouter(Router *gin.RouterGroup) {
	baseRouter := Router.Group("init")
	dbApi := v1.ApiGroupApp.SystemApiGroup.DBApi
	{
		baseRouter.POST("initdb", dbApi.InitDB)
		baseRouter.POST("checkdb", dbApi.CheckDB)
	}

}
