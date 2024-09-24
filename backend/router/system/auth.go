package system

import (
	"github.com/gin-gonic/gin"
)

type AuthRouter struct{}

func (u *AuthRouter) InitAuthRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	AuthRouter := Router.Group("auth")
	{
		AuthRouter.POST("login", authApi.Login)
		AuthRouter.POST("captcha", authApi.Captcha)
	}

	return AuthRouter
}
