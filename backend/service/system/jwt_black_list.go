package system

import (
	"context"
	"gin-pro/global"
	"gin-pro/model/system"
	"gin-pro/utils"
	"go.uber.org/zap"
)

type JwtService struct {
}

// @description: 拉黑jwt
func (j *JwtService) JsonInBlacklist(jwtList system.JwtBlacklist) (err error) {
	err = global.GVA_DB.Create(&jwtList).Error
	if err != nil {
		return
	}
	global.BlockCache.SetDefault(jwtList.Jwt, struct {
	}{})
	return
}

// @description: 判断JWT是否在黑名单内部
func (j *JwtService) IsBlacklist(jwt string) bool {
	_, ok := global.BlockCache.Get(jwt)
	return ok
}

// @description:  jwt存入redis并设置过期时间
func (j *JwtService) SetRedisJWT(jwt string, userName string) (err error) {
	dr, err := utils.ParseDuration(global.GVA_CONFIG.Jwt.ExpiresTime)
	if err != nil {
		return err
	}
	timer := dr
	err = global.GVA_REDIS.Set(context.Background(), userName, jwt, timer).Err()
	return
}

// @description: 从redis取jwt
func (j *JwtService) GetRedisJWT(userName string) (redisJWT string, err error) {
	redisJWT, err = global.GVA_REDIS.Get(context.Background(), userName).Result()
	return redisJWT, err
}

func LoadAll() {
	var data []string
	err := global.GVA_DB.Model(&system.JwtBlacklist{}).Select("jwt").Find(&data).Error
	if err != nil {
		global.GVA_LOG.Error("加载数据库jwt黑名单失败!", zap.Error(err))
	}
	for i := 0; i < len(data); i++ {
		global.BlockCache.SetDefault(data[i], struct {
		}{})
	}
}
