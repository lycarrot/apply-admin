package system

import (
	"errors"
	"gin-pro/global"
	"gin-pro/model/system"
	"gorm.io/gorm"
)

type ApiService struct {
}

var ApiServiceApp = new(ApiService)

// @function: CreateApi
// @description: 新增api
// @param: api model.SysApi
// @return: err error
func (a *ApiService) CreateApi(api system.SysApi) (err error) {
	if !errors.Is(global.GVA_DB.Where("Path = ? AND Method =", api.Path, api.Method).First(&system.SysApi{}).Error, gorm.ErrRecordNotFound) {
		return errors.New("当前api对应方法已存在")
	}
	return global.GVA_DB.Create(&api).Error
}
