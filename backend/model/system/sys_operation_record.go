package system

import (
	"gin-pro/global"
	"time"
)

type SysOperationRecord struct {
	global.GVA_MODEL
	Ip           string        `json:"ip" form:"ip" gorm:"comment:请求ip"`                              // 请求ip
	Method       string        `json:"method" form:"method" gorm:"comment:请求方法"`                      // 请求方法
	Path         string        `json:"path" form:"path" gorm:"comment:请求路径"`                          // 请求路径
	Status       int           `json:"status" form:"status" gorm:"comment:请求状态"`                      // 请求状态
	Latency      time.Duration `json:"latency" form:"latency" gorm:"comment:延迟" swaggertype:"string"` // 延迟
	Agent        string        `json:"agent" form:"agent" gorm:"type:text;comment:代理"`                // 代理
	ErrorMessage string        `json:"error_message" form:"error_message" gorm:"comment:错误信息"`        // 错误信息
	Body         string        `json:"body" form:"body" gorm:"type:text;comment:请求Body"`              // 请求Body
	Resp         string        `json:"resp" form:"resp" gorm:"type:text;comment:响应Body"`              // 响应Body
	UserID       int           `json:"user_id" form:"user_id" gorm:"comment:用户id"`                    // 用户id
	User         SysUser       `json:"user"`
}
