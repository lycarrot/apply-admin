package request

import (
	"gin-pro/global"
	"gin-pro/model/common/request"
	"gin-pro/model/system"
)

type SysMenuQuery struct {
	request.PageQuery
}

func DefaultMenu() []system.SysBaseMenu {
	return []system.SysBaseMenu{{
		GVA_MODEL: global.GVA_MODEL{Id: 1},
		ParentId:  0,
		Path:      "index",
		Name:      "index",
		Component: "index.tsx",
		Sort:      1,
		Meta: system.Meta{
			Title: "首页",
			Icon:  "index",
		},
	},
	}
}
