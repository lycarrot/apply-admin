package service

import (
	"gin-pro/service/example"
	"gin-pro/service/system"
)

type ServiceGroup struct {
	ExampleServiceGroup example.ExampleGroup
	SystemServiceGroup  system.ServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)
