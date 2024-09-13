package system

import (
	v1 "gin-pro/api/v1"
	"github.com/gin-gonic/gin"
)

type BaseRouter struct {
}

func (s *BaseRouter) InitBaseRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	baseRouter := Router.Group("base")
	baseApi := v1.ApiGroupApp.SystemApiGroup.BaseApi
	{
		baseRouter.POST("captcha", baseApi.Captcha)
		baseRouter.POST("login", baseApi.Login)
	}
	return baseRouter
}
