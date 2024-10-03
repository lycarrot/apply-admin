package system

import "github.com/gin-gonic/gin"

type ApiRouter struct {
}

func (a *ApiRouter) InitApiRouter(Router *gin.RouterGroup) {
	router := Router.Group("/api")
	{
		router.POST("/create", apiApi.CreateApi)
	}
}
