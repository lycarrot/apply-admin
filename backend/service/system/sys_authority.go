package system

import (
	"errors"
	"gin-pro/global"
	"gin-pro/model/system"
	systemReq "gin-pro/model/system/request"
	"gorm.io/gorm"
	"strconv"
)

type AuthorityService struct {
}

var ErrRoleExistence = errors.New("存在相同角色id")

var AuthorityServiceApp = new(AuthorityService)

func (c *AuthorityService) CreateAuthority(auth system.SysAuthority) (authority system.SysAuthority, err error) {
	if err = global.GVA_DB.Where("authority_id=?", auth.AuthorityId).First(&system.SysAuthority{}).Error; !errors.Is(err, gorm.ErrRecordNotFound) {
		return auth, ErrRoleExistence
	}
	//开启事务
	e := global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&auth).Error; err != nil {
			return err
		}
		auth.SysBaseMenus = systemReq.DefaultMenu()
		//关联菜单数据
		if err = tx.Model(&auth).Association("SysBaseMenus").Replace(&auth.SysBaseMenus); err != nil {
			return err
		}
		casbinInfos := systemReq.DefaultCasbin()
		authorityId := strconv.Itoa(int(auth.AuthorityId))
		rules := [][]string{}
		for _, v := range casbinInfos {
			rules = append(rules, []string{authorityId, v.Path, v.Method})
		}
		return CasbinServiceApp.AddPolicies(tx, rules)
	})
	return auth, e
}
