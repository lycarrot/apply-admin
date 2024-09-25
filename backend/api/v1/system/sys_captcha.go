package system

import (
	"fmt"
	"gin-pro/global"
	"gin-pro/model/common/response"
	systemRes "gin-pro/model/system/response"
	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
	"go.uber.org/zap"
	"time"
)

var store = base64Captcha.DefaultMemStore

// Captcha
// @Tags Auth
// @Summary 生成验证码
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Success 200 {object} response.Response{data=systemRes.SysCaptchaResponse,msg=string} "生成验证码,返回包括随机数id,base64,验证码长度,是否开启验证码"
// @Router    /auth/captcha [post]
func (a *AuthApi) Captcha(c *gin.Context) {
	openCaptcha := global.GVA_CONFIG.Captcha.OpenCaptcha
	openCaptchaTimeOut := global.GVA_CONFIG.Captcha.OpenCaptchaTimeOut
	key := c.ClientIP()
	v, ok := global.BlockCache.Get(key)
	if !ok {
		global.BlockCache.Set(key, 1, time.Second*time.Duration(openCaptchaTimeOut))
	}

	var oc bool
	if openCaptcha == 0 || openCaptcha < interfaceToInt(v) {
		oc = true
	}
	//用于创建数字验证码驱动器
	dirver := base64Captcha.NewDriverDigit(global.GVA_CONFIG.Captcha.ImgHeight, global.GVA_CONFIG.Captcha.ImgWidth, global.GVA_CONFIG.Captcha.KeyLong, 0.7, 80)
	//用于创建验证码实例
	cp := base64Captcha.NewCaptcha(dirver, store)
	id, b64s, _, err := cp.Generate()
	if err != nil {
		global.GVA_LOG.Error("验证码获取失败", zap.Error(err))
		response.FailWithMessage("获取验证码失败", c)
	}
	response.OkWithDetailed(systemRes.SysCaptchaResponse{
		CaptchaId:     id,
		PicPath:       b64s,
		CaptchaLength: global.GVA_CONFIG.Captcha.KeyLong,
		OpenCaptcha:   oc,
	}, "获取验证码成功", c)

}

// 将传入的接口类型的值转换为int类型。如果传入的值是int类型，则直接返回该值；否则返回0。
func interfaceToInt(v interface{}) (i int) {
	switch v := v.(type) {
	case int:
		fmt.Println(v)
		i = v
	default:
		i = 0
	}
	return
}
