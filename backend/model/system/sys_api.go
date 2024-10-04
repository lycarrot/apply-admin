package system

import "gin-pro/global"

type SysApi struct {
	global.GVA_MODEL
	Path        string `json:"path" gorm:"comment:路径"`                 //路径
	Group       string `json:"group" gorm:"comment:分组"`                //分组
	Description string `json:"description" gorm:"comment:描述"`          //描述
	Method      string `json:"method" gorm:"comment:方法;default:POST;"` //方法:创建POST(默认)|查看GET|更新PUT|删除DELETE
}

func (SysApi) TableName() string {
	return "sys_apis"

}
