package system

import (
	"github.com/gin-gonic/gin"
)

type AuthRouter struct{}

func (u *AuthRouter) InitAuthRouter(Router *gin.RouterGroup) {
	router := Router.Group("auth")
	routerWithout := Router.Group("auth")
	{
		router.POST("admin/register", authApi.Register)
	}
	{
		routerWithout.POST("login", authApi.Login)
		routerWithout.POST("captcha", authApi.Captcha)
	}

}
