package system

import "gin-pro/service"

type ApiGroup struct {
	BaseApi
	DBApi
	SystemApiApi
}

var (
	userService = service.ServiceGroupApp.SystemServiceGroup.UserService
)
