package request

import (
	"gin-pro/model/common/request"
	"gin-pro/model/system"
)

type SysOperationRecordSearch struct {
	system.SysOperationRecord
	request.PageInfo
}
