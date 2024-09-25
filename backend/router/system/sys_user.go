package system

import "github.com/gin-gonic/gin"

type UserRouter struct{}

func (u *UserRouter) InitUserRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	userRouter := Router.Group("user")
	{

	}
	return userRouter
}
