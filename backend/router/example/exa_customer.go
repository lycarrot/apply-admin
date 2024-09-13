package example

import (
	v1 "gin-pro/api/v1"
	"github.com/gin-gonic/gin"
)

type CustomerRouter struct {
}

func (e *CustomerRouter) InitCustomerRouter(Router *gin.RouterGroup) {
	customerRouter := Router.Group("customer")
	customerRouterWithoutRecord := Router.Group("customer")
	exaCustomerApi := v1.ApiGroupApp.ExampleApiGroup.CustomerApi
	{
		customerRouter.POST("customer", exaCustomerApi.CreateExaCustomer)
		customerRouter.PUT("customer", exaCustomerApi.UpdateExaCustomer)
		customerRouter.DELETE("customer", exaCustomerApi.DeleteExaCustomer)
	}
	{
		customerRouterWithoutRecord.GET("customer", exaCustomerApi.GetExaCustomer)
		customerRouterWithoutRecord.GET("customerList", exaCustomerApi.GetExaCustomerList)
	}
}
