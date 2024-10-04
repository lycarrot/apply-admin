package system

import (
	"github.com/gin-gonic/gin"
)

type AuthRouter struct{}

func (u *AuthRouter) InitAuthRouter(Router *gin.RouterGroup, RouterPub *gin.RouterGroup) {
	router := Router.Group("auth")
	routerPub := RouterPub.Group("auth")
	{
		router.POST("/admin/register", authApi.Register)
	}
	{
		routerPub.POST("login", authApi.Login)
		routerPub.GET("captcha", authApi.Captcha)
	}

}
