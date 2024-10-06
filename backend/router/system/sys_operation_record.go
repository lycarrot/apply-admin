package system

import "github.com/gin-gonic/gin"

type OperationRecordRouter struct{}

func (o *OperationRecordRouter) InitOperationRecordRouter(Router *gin.RouterGroup) {
	router := Router.Group("/record")
	{
		router.GET("getLists", operationRecordApi.GetLists)
	}
}
