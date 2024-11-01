package system

import v1 "gin-pro/api/v1"

type RouterGroup struct {
	AuthRouter
	UserRouter
	MenuRouter
	AuthorityRouter
	ApiRouter
	CasbinRouter
	OperationRecordRouter
}

var (
	authApi            = v1.ApiGroupApp.SystemApiGroup.AuthApi
	authorityApi       = v1.ApiGroupApp.SystemApiGroup.AuthorityApi
	menuApi            = v1.ApiGroupApp.SystemApiGroup.MenuApi
	apiApi             = v1.ApiGroupApp.SystemApiGroup.ApiApi
	casbinApi          = v1.ApiGroupApp.SystemApiGroup.CasbinApi
	operationRecordApi = v1.ApiGroupApp.SystemApiGroup.OperationRecordApi
	userApi            = v1.ApiGroupApp.SystemApiGroup.UserApi
)
