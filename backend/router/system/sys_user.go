package system

import (
	"gin-pro/middleware"
	"github.com/gin-gonic/gin"
)

type UserRouter struct{}

func (u *UserRouter) InitUserRouter(Router *gin.RouterGroup) {
	router := Router.Group("user").Use(middleware.OperationRecordHandler())
	routerNotRecord := Router.Group("user")
	{
		router.GET("update", userApi.UpdateUser)
		router.GET("detail/:id", userApi.GetUseDetail)
	}
	{
		routerNotRecord.GET("getLists", userApi.GetUserLists)
	}
}
