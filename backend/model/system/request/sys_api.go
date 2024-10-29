package request

import (
	"gin-pro/model/common/request"
	"gin-pro/model/system"
)

type OrderItem struct {
	Field string `json:"field"`
	Value string `json:"value"`
}

//	type SysApiOrder struct {
//		Id          string `json:"id"`
//		Path        string `json:"path"`
//		Category    string `json:"category"`
//		Description string `json:"description"`
//		Method      string `json:"method"`
//	}
type SysApiQuery struct {
	request.PageQuery
	system.SysApi
	Order []OrderItem `json:"order" gorm:"comment:排序查询"` //排序查询
}
