package initialize

import (
	"gin-pro/global"
	"gin-pro/utils"
	"github.com/songzhibin97/gkit/cache/local_cache"
)

func OtherInit() {
	dr, err := utils.ParseDuration(global.GVA_CONFIG.Jwt.ExpiresTime)
	if err != nil {
		panic(err)
	}
	_, err = utils.ParseDuration(global.GVA_CONFIG.Jwt.BufferTime)
	if err != nil {
		panic(err)
	}
	//本地缓存时长设置
	global.BlockCache = local_cache.NewCache(
		local_cache.SetDefaultExpire(dr))
}
