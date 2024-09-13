package v1

import (
	"gin-pro/api/v1/example"
	"gin-pro/api/v1/system"
)

type ApiGroup struct {
	ExampleApiGroup example.ApiGroup
	SystemApiGroup  system.ApiGroup
}

var ApiGroupApp = new(ApiGroup)
