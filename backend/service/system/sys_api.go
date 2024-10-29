package system

import (
	"errors"
	"gin-pro/global"
	"gin-pro/model/system"
	"gin-pro/model/system/request"
	"gorm.io/gorm"
	"reflect"
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
	if query.BgGroup != "" {
		db = db.Where("`group` = ?", query.BgGroup)
	}

	err = db.Count(&total).Error
	if err != nil {
		return apiLists, total, err
	}
	db = db.Limit(limit).Offset(offset)
	orderVal := reflect.ValueOf(query.Order)
	orderTyp := orderVal.Type()
	orderMap := map[string]bool{
		"Id":          true,
		"Path":        true,
		"BgGroup":     true,
		"Description": true,
		"Method":      true,
	}

	orderStr := ""
	l := orderVal.NumField()
	isFirst := true
	for i := 0; i < l; i++ {
		val := orderVal.Field(i).String()
		name := orderTyp.Field(i).Name
		if orderMap[name] && (val == "desc" || val == "asc") {
			item := strings.ToLower(name) + " " + val
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
