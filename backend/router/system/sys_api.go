package system

import (
	"gin-pro/middleware"
	"github.com/gin-gonic/gin"
)

type ApiRouter struct {
}

func (a *ApiRouter) InitApiRouter(Router *gin.RouterGroup) {
	router := Router.Group("/api").Use(middleware.OperationRecordHandler())
	routerNotRecord := Router.Group("/api")
	{
		router.POST("/create", apiApi.CreateApi)        //创建api
		router.PUT("/update", apiApi.UpdateApi)         //更新api
		router.GET("/detail/:id", apiApi.GetApiDetail)  //获取api详情
		router.DELETE("/batchDel", apiApi.BatchDelApis) //批量删除api

	}
	{
		routerNotRecord.POST("/getLists", apiApi.GetApiLists) //获取api列表
	}
}
