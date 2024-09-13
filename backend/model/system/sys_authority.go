package system

import "time"

type SysAuthority struct {
	CreatedAt time.Time //创建时间
	UpdatedAt time.Time //更新时间
	//时间字段可以为空*time.Time
	DeletedAt       *time.Time      `sql:"index"`
	AuthorityId     uint            `json:"authorityId" gorm:"primary_key;comment:角色ID;size:90"`
	AuthorityName   string          `json:"authorityName" gorm:"comment:角色名"` // 角色名
	ParentId        *uint           `json:"parentId" gorm:"comment:父角色ID"`    // 父角色ID
	DataAuthorityId []*SysAuthority `json:"dataAuthorityId" gorm:"many2many:sys_data_authority_id"`
	Children        []SysAuthority  `json:"children" gorm:"-"`
	Users           []SysUser       `json:"-" gorm:"many2many:sys_user_authority;"`
	DefaultRouter   string          `json:"defaultRouter" gorm:"comment:默认菜单;default:dashboard"` // 默认菜单(默认dashboard)
}

func (SysAuthority) TableName() string {
	return "sys_authorities"
}
