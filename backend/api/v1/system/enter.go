package system

import "gin-pro/service"

type ApiGroup struct {
	AuthApi
}

var (
	userService = service.ServiceGroupApp.SystemServiceGroup.UserService
)
