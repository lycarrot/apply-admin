package system

import (
	"errors"
	"fmt"
	"gin-pro/global"
	"gin-pro/model/system"
	"gin-pro/utils"
)

type UserService struct {
}

func (s *UserService) Login(u *system.SysUser) (userInter *system.SysUser, err error) {
	if nil == global.GVA_DB {
		return nil, fmt.Errorf("db  not init ")
	}
	var user system.SysUser
	err = global.GVA_DB.Where("username=?", u.Username).Preload("Authorities").Preload("Authority").First(&user).Error
	if err == nil {
		if ok := utils.BcryptCheck(u.Password, user.Password); !ok {
			return nil, errors.New("密码错误")
		}
	}
	return &user, err
}
