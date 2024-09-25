package system

import (
	"github.com/gin-gonic/gin"
)

type AuthRouter struct{}

func (u *AuthRouter) InitAuthRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	authRouter := Router.Group("auth")
	{
		authRouter.POST("login", authApi.Login)
		authRouter.POST("captcha", authApi.Captcha)
		authRouter.POST("admin_register", authApi.Register)
	}

	return authRouter
}
