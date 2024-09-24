package system

import "github.com/gin-gonic/gin"

type UserRouter struct{}

func (u *UserRouter) InitUserRouter(Router *gin.RouterGroup) {}
