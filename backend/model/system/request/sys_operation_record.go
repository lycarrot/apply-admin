package request

import (
	"gin-pro/model/common/request"
	"gin-pro/model/system"
)

type SysOperationQuery struct {
	system.SysOperationRecord
	request.PageQuery
}
