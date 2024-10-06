package system

import "gin-pro/service"

type ApiGroup struct {
	AuthApi
	AuthorityApi
	MenuApi
	ApiApi
	CasbinApi
	OperationRecordApi
}

var (
	userService            = service.ServiceGroupApp.SystemServiceGroup.UserService
	authorityService       = service.ServiceGroupApp.SystemServiceGroup.AuthorityService
	casbinService          = service.ServiceGroupApp.SystemServiceGroup.CasbinService
	menuService            = service.ServiceGroupApp.SystemServiceGroup.MenuService
	apiService             = service.ServiceGroupApp.SystemServiceGroup.ApiService
	operationRecordService = service.ServiceGroupApp.SystemServiceGroup.OperationRecordService
)
