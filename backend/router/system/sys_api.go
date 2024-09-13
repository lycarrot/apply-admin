package system

import (
	"github.com/gin-gonic/gin"
)

type ApiRouter struct {
}

func (s *ApiRouter) InitApiRouter(Router *gin.RouterGroup, RouterPub *gin.RouterGroup) {
	//apiRouter := Router.Group("api")
	//apiRouterWithoutRecord := Router.Group("api")
	//apiPublicRouterWithoutRecord := RouterPub.Group("api")
	//apiRouterApi := v1.ApiGroupApp.SystemApiGroup.SystemApiApi
}
