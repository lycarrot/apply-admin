package system

import (
	"gin-pro/global"
	"gin-pro/model/system"
	"gin-pro/model/system/request"
)

type OperationRecordService struct{}

var OperationRecordServiceApp = new(OperationRecordService)

// @function: CreateSysOperationRecord
// @description: 创建记录
// @param: sysOperationRecord model.SysOperationRecord
// @return: err error
func (o *OperationRecordService) CreateOperationRecord(sysOperationRecord system.SysOperationRecord) error {
	return global.GVA_DB.Create(&sysOperationRecord).Error
}

// @function: GetOperationRecordLists
// @description: 获取记录列表
// @param: query request.SysOperationRecordSearch
// @return: err error
func (o *OperationRecordService) GetOperationRecordLists(query request.SysOperationRecordSearch) (lists []system.SysOperationRecord, total int64, err error) {
	var limit = query.PageSize
	var offset = (query.Page - 1) * limit
	var db = global.GVA_DB.Model(&system.SysOperationRecord{})
	if query.Method != "" {
		db = db.Where("method = ?", query.Method)
	}
	if query.Path != "" {
		db = db.Where("path LIKE ?", "%"+query.Path+"%")
	}
	if query.Status != 0 {
		db = db.Where("status = ?", query.Status)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Order("id desc").Offset(offset).Limit(limit).Preload("User").Find(&lists).Error
	return lists, total, err
}
