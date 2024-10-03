package system

import (
	"errors"
	"gin-pro/global"
	"gin-pro/model/system"
	"gorm.io/gorm"
)

type MenuService struct {
}

var MenuServiceApp = new(MenuService)

// @function:AddMenuItem
// @description: 添加路由菜单
// @param: menu system.SysBaseMenu
// @return: error
func (m *MenuService) AddMenuItem(menu system.SysBaseMenu) error {
	if !errors.Is(global.GVA_DB.Where("name = ?", menu.Name).First(&system.SysBaseMenu{}).Error, gorm.ErrRecordNotFound) {
		return errors.New("当前菜单名称已存在")
	}
	return global.GVA_DB.Create(&menu).Error
}
