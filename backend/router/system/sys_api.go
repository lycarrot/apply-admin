package system

import (
	"gin-pro/middleware"
	"github.com/gin-gonic/gin"
)

type ApiRouter struct {
}

func (a *ApiRouter) InitApiRouter(Router *gin.RouterGroup) {
	router := Router.Group("/api").Use(middleware.OperationRecordHandler())
	{
		router.POST("/create", apiApi.CreateApi)
	}
}
