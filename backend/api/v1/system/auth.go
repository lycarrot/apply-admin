package system

import (
	"gin-pro/global"
	"gin-pro/model/common/response"
	"gin-pro/model/system"
	systemReq "gin-pro/model/system/request"
	systemRes "gin-pro/model/system/response"
	"gin-pro/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"time"
)

type AuthApi struct{}

// Login
// @Tags Auth
// @Summary  用户登录
// @Produce   application/json
// @Param    data  body      systemReq.Login                                             true  "用户名, 密码, 验证码"
// @Success  200   {object}  response.Response{data=systemRes.LoginResponse,msg=string}  "返回包括用户信息,token,过期时间"
// @Router   /auth/login [post]
func (a *AuthApi) Login(c *gin.Context) {
	var l systemReq.Login
	err := c.ShouldBindJSON(&l)
	key := c.ClientIP()
	if err != nil {
		response.FailWithMessage(err.Error(), c)
	}
	//err =
	//判断验证码是否开启
	openCaptcha := global.GVA_CONFIG.Captcha.OpenCaptcha
	openCaptchaTimeOut := global.GVA_CONFIG.Captcha.OpenCaptchaTimeOut
	v, ok := global.BlockCache.Get(key)
	if !ok {
		global.BlockCache.Set(key, 1, time.Second*time.Duration(openCaptchaTimeOut))
	}
	var oc bool = openCaptcha == 0 || openCaptcha < interfaceToInt(v)
	if !oc || (l.CaptchaId != "" && l.Captcha != "" && store.Verify(l.CaptchaId, l.Captcha, true)) {
		u := &system.SysUser{Username: l.Username, Password: l.Password}
		user, err := userService.Login(u)
		if err != nil {
			global.GVA_LOG.Error("登陆失败! 用户名不存在或者密码错误!", zap.Error(err))
			global.BlockCache.Increment(key, 1)
			response.FailWithMessage("用户名不存在或者密码错误", c)
			return
		}
		if user.Enable != 1 {
			global.GVA_LOG.Error("登陆失败! 用户被禁止登录!", zap.Error(err))
			global.BlockCache.Increment(key, 1)
			response.FailWithMessage("用户被禁止登", c)
			return
		}
		a.TokenNext(c, *user)
		return
	}
	global.BlockCache.Increment(key, 1)
	response.FailWithMessage("验证码错误", c)
}

// TokenNext 登录以后签发jwt
func (a *AuthApi) TokenNext(c *gin.Context, user system.SysUser) {
	j := utils.JWT{SigningKey: []byte(global.GVA_CONFIG.Jwt.SigningKey)}
	claims := j.CreateClaims(systemReq.BaseClaims{
		UUID:        user.UUID,
		ID:          user.ID,
		NickName:    user.NickName,
		Username:    user.Username,
		AuthorityId: user.AuthorityId,
	})
	token, err := j.CreateToken(claims)
	if err != nil {
		global.GVA_LOG.Error("获取token失败", zap.Error(err))
		response.FailWithMessage("获取token失败", c)
	}
	if !global.GVA_CONFIG.System.UseMultipoint {
		utils.SetToken(c, token, int(claims.RegisteredClaims.ExpiresAt.Unix()-time.Now().Unix()))
		response.OkWithDetailed(systemRes.LoginResponse{
			User:      user,
			Token:     token,
			ExpiresAt: claims.RegisteredClaims.ExpiresAt.Unix() * 1000,
		}, "登录成功", c)
		return
	}

}
