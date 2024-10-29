package request

import (
	"gin-pro/model/common/request"
	"gin-pro/model/system"
)

type SysApiOrder struct {
	ID          string `json:"id"`
	Path        string `json:"path"`
	BgGroup     string `json:"bgGroup"`
	Description string `json:"description"`
	Method      string `json:"method"`
}
type SysApiQuery struct {
	request.PageQuery
	system.SysApi
	Order SysApiOrder `json:"order" gorm:"comment:排序查询"` //排序查询
}
