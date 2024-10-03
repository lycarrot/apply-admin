package system

import (
	"errors"
	"fmt"
	"gin-pro/global"
	"gin-pro/model/system"
	"gin-pro/utils"
	"github.com/gofrs/uuid/v5"
	"gorm.io/gorm"
)

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

func (s *UserService) Register(u system.SysUser) (userInter system.SysUser, err error) {
	var user system.SysUser
	if !errors.Is(global.GVA_DB.Where("username=?", u.Username).First(&user).Error, gorm.ErrRecordNotFound) {
		return userInter, errors.New("用户已注册")
	}

	u.Password = utils.BcryptHash(u.Password)
	u.UUID = uuid.Must(uuid.NewV4())

	err = global.GVA_DB.Create(&u).Error
	return u, err
}
