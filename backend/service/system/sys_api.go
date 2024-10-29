package system

import (
	"errors"
	"gin-pro/global"
	"gin-pro/model/system"
	"gin-pro/model/system/request"
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

// @function: GetApiLists
// @description: 获取api列表
// @param: query request.SysApiQuery
// @return: lists []system.SysApi, total int64, err error
func (a *ApiService) GetApiLists(query request.SysApiQuery) (lists []system.SysApi, total int64, err error) {
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
