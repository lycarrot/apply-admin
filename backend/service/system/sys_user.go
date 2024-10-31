package system

import (
	"gin-pro/global"
	"gin-pro/model/system"
	"gin-pro/model/system/request"
)

type UserService struct {
}

// @function: GetUserLists
// @description: 获取用户列表
// @param: query request.SysUserQuery
// @return: err error
func (u *UserService) GetUserLists(query request.SysUserQuery) (lists []system.SysUser, total int64, err error) {
	var limit = query.PageSize
	var offset = (query.Page - 1) * query.PageSize
	var db = global.GVA_DB.Model(&system.SysUser{})
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Offset(offset).Limit(limit).Find(&lists).Error
	return lists, total, err
}

// @function: GetUseDetail
// @description:查询用户信息
// @param: id int
// @return: err error
func (u *UserService) GetUseDetail(id int) (user system.SysUser, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&user).Error
	return user, err
}
