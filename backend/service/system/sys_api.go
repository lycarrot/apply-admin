package system

import (
	"errors"
	"gin-pro/global"
	"gin-pro/model/common/request"
	"gin-pro/model/system"
	sysRes "gin-pro/model/system/request"
	"gorm.io/gorm"
	"strings"
)

type ApiService struct {
}

var ApiServiceApp = new(ApiService)

// @function: CreateApi
// @description: 新增api
// @param: api model.SysApi
// @return: err error
func (a *ApiService) CreateApi(api system.SysApi) (err error) {
	if !errors.Is(global.GVA_DB.Where("path = ? AND method = ?", api.Path, api.Method).First(&system.SysApi{}).Error, gorm.ErrRecordNotFound) {
		return errors.New("当前api对应方法已存在")
	}
	return global.GVA_DB.Create(&api).Error
}

// @function: GetDetail
// @description: 根据id获取api详情
// @param: id int
// @return: err error
func (a *ApiService) GetDetail(id int) (api system.SysApi, err error) {
	err = global.GVA_DB.First(&api, "id = ?", id).Error
	return api, err
}

// @function: UpdateApi
// @description: 更新api信息
// @param: api system.SysApi)
// @return: err error
func (a *ApiService) UpdateApi(api system.SysApi) (err error) {
	var item system.SysApi
	err = global.GVA_DB.First(&item, "id = ?", api.Id).Error
	if err != nil {
		return err
	}
	if item.Path != api.Path || item.Method != api.Method {
		var duplicateItem system.SysApi
		dupErr := global.GVA_DB.First(&duplicateItem, "path = ? AND method = ?", api.Path, api.Method).Error
		if dupErr != nil {
			if !errors.Is(dupErr, gorm.ErrRecordNotFound) {
				return dupErr
			}
		} else {
			if duplicateItem.Id != api.Id {
				return errors.New("当前api路径已存在")
			}
		}
	}
	err = CasbinServiceApp.UpdateCasbinApi(item.Path, item.Method, api.Path, api.Method)
	if err != nil {
		return err
	}
	return global.GVA_DB.Save(&item).Error

}

// @function: BatchDel
// @description: 批量删除api
// @param: ids request.IdsQuery
// @return: err error
func (a *ApiService) BatchDel(ids request.IdsQuery) (err error) {
	return global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		var apis []system.SysApi
		err = tx.Find(&apis, "id in ?", ids.Ids).Error
		if err != nil {
			return err
		}
		err = tx.Delete(&[]system.SysApi{}, "id in ?", ids.Ids).Error
		if err != nil {
			return err
		}
		for _, api := range apis {
			CasbinServiceApp.ClearCasbin(1, api.Path, api.Method)
		}
		return err
	})
}

// @function: GetApiLists
// @description: 获取api列表
// @param: query sysRes.SysApiQuery
// @return: lists []system.SysApi, total int64, err error
func (a *ApiService) GetApiLists(query sysRes.SysApiQuery) (lists []system.SysApi, total int64, err error) {
	var limit = query.PageSize
	var offset = (query.Page - 1) * limit
	var apiLists []system.SysApi
	db := global.GVA_DB.Model(&system.SysApi{})
	if query.Path != "" {
		db = db.Where("path LIKE ?", "%"+query.Path+"%")
	}
	if query.Description != "" {
		db = db.Where("description LIKE ?", "%"+query.Description+"%")
	}
	if query.Method != "" {
		db = db.Where("method = ?", query.Method)
	}
	if query.Category != "" {
		db = db.Where("category = ?", query.Category)
	}

	err = db.Count(&total).Error
	if err != nil {
		return apiLists, total, err
	}
	db = db.Limit(limit).Offset(offset)
	orderMap := map[string]bool{
		"id":          true,
		"path":        true,
		"category":    true,
		"description": true,
		"method":      true,
	}
	orderStr := ""
	isFirst := true
	for _, v := range query.Order {
		if orderMap[v.Field] && (v.Value == "desc" || v.Value == "asc") {
			item := strings.ToLower(v.Field) + " " + v.Value
			if isFirst == false {
				item = ", " + item
			}
			isFirst = false
			orderStr += item
		}
	}
	if orderStr != "" {
		db = db.Order(orderStr)
	}
	err = db.Find(&apiLists).Error
	return apiLists, total, err
}
