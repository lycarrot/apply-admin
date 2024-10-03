package request

import (
	"gin-pro/global"
	"gin-pro/model/system"
)

func DefaultMenu() []system.SysBaseMenu {
	return []system.SysBaseMenu{{
		GVA_MODEL: global.GVA_MODEL{ID: 1},
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
