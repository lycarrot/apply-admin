package system

import v1 "gin-pro/api/v1"

type RouterGroup struct {
	AuthRouter
	UserRouter
}

var (
	authApi = v1.ApiGroupApp.SystemApiGroup.AuthApi
)
