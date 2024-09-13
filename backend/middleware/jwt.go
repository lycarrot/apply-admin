package middleware

import (
	"errors"
	"gin-pro/global"
	"gin-pro/model/common/response"
	"gin-pro/model/system"
	"gin-pro/service"
	"gin-pro/utils"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"go.uber.org/zap"
	"strconv"
	"time"
)

var JwtService = service.ServiceGroupApp.SystemServiceGroup.JwtService

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		//这里jwt鉴权取头部信息 x-token 登录时回返回token信息 这里前端需要把token存储到cookie或者本地localStorage中 不过需要跟后端协商过期时间 可以约定刷新令牌或者重新登录
		token := utils.GetToken(c)
		if token == "" {
			response.FailWithDetailed(gin.H{"reload": true}, "未登录或非法访问", c)
			c.Abort()
			return
		}
		if JwtService.IsBlacklist(token) {
			response.FailWithDetailed(gin.H{"reload": true}, "您的帐户异地登陆或令牌失效", c)
			utils.ClearToken(c)
			c.Abort()
			return
		}
		j := utils.NewJWT()
		claims, err := j.ParseToken(token)
		if err != nil {
			if errors.Is(err, utils.TokenExpired) {
				response.FailWithDetailed(gin.H{"repload": true}, "授权已过期", c)
				utils.ClearToken(c)
				c.Abort()
				return
			}
			response.FailWithDetailed(gin.H{"repload": true}, err.Error(), c)
			utils.ClearToken(c)
			c.Abort()
			return
		}
		// 已登录用户被管理员禁用 需要使该用户的jwt失效 此处比较消耗性能 如果需要 请自行打开
		// 用户被删除的逻辑 需要优化 此处比较消耗性能 如果需要 请自行打开

		//if user, err := userService.FindUserByUuid(claims.UUID.String()); err != nil || user.Enable == 2 {
		//	_ = jwtService.JsonInBlacklist(system.JwtBlacklist{Jwt: token})
		//	response.FailWithDetailed(gin.H{"reload": true}, err.Error(), c)
		//	c.Abort()
		//}

		c.Set("claims", claims)
		c.Next()
		//比较 JWT token 中声明的过期时间（claims.ExpiresAt）和当前时间之间的差值是否小于某个预定义的缓冲时间（claims.BufferTime）。
		if claims.ExpiresAt.Unix()-time.Now().Unix() < claims.BufferTime {
			dr, _ := utils.ParseDuration(global.GVA_CONFIG.Jwt.ExpiresTime)
			claims.ExpiresAt = jwt.NewNumericDate(time.Now().Add(dr))
			newToken, _ := j.CreateTokenByOldToken(token, *claims)
			newClaims, _ := j.ParseToken(newToken)
			c.Header("new-token", newToken)
			c.Header("new-expires-at", strconv.FormatInt(newClaims.ExpiresAt.Unix(), 10))
			utils.SetToken(c, newToken, int(dr.Seconds()))
			if global.GVA_CONFIG.System.UseMultipoint {
				RedisJwtToken, err := JwtService.GetRedisJWT(newClaims.Username)
				if err != nil {
					global.GVA_LOG.Error("get redis jwt failed", zap.Error(err))
				} else {
					// 当之前的取成功时才进行拉黑操作
					_ = JwtService.JsonInBlacklist(system.JwtBlacklist{Jwt: RedisJwtToken})
				}
				// 无论如何都要记录当前的活跃状态
				_ = JwtService.SetRedisJWT(newToken, newClaims.Username)
			}
		}
	}

}
