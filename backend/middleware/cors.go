package middleware

import (
	"gin-pro/config"
	"gin-pro/global"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		origin := c.Request.Header.Get("Origin")
		c.Header("Access-Control-Allow-Origin", origin)
		c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token,X-Token,X-User-Id")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS,DELETE,PUT")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type, New-Token, New-Expires-At")
		c.Header("Access-Control-Allow-Credentials", "true")
		if method == "OPTIONS" {
			//在处理请求时，立即返回状态码为 204，而不继续执行后续的处理逻辑。
			c.AbortWithStatus(http.StatusNoContent)
		}
		c.Next()
	}
}

func CorsByRules() gin.HandlerFunc {
	if global.GVA_CONFIG.Cors.Mode == "allow-all" {
		return Cors()
	}
	return func(c *gin.Context) {
		whitelist := checkCors(c.GetHeader("origin"))
		// 通过检查, 添加请求头
		if whitelist != nil {
			c.Header("Access-Control-Allow-Origin", whitelist.AllowOrigin)
			c.Header("Access-Control-Allow-Headers", whitelist.AllowHeaders)
			c.Header("Access-Control-Allow-Methods", whitelist.AllowMethods)
			c.Header("Access-Control-Expose-Headers", whitelist.ExposeHeaders)
			if whitelist.AllowCredentials {
				c.Header("Access-Control-Allow-Credentials", "true")
			}
		}
		// 严格白名单模式且未通过检查，直接拒绝处理请求
		if whitelist == nil && global.GVA_CONFIG.Cors.Mode == "strict-whitelist" && !(c.Request.Method == "GET" && c.Request.URL.Path == "/health") {
			c.AbortWithStatus(http.StatusForbidden)
		} else {
			// 非严格白名单模式，无论是否通过检查均放行所有 OPTIONS 方法
			if c.Request.Method == http.MethodOptions {
				c.AbortWithStatus(http.StatusNoContent)
			}
		}
		c.Next()

	}
}

func checkCors(currentOrigin string) *config.CORSWhitelist {
	for _, whitelist := range global.GVA_CONFIG.Cors.Whitelist {
		if currentOrigin == whitelist.AllowOrigin {
			return &whitelist
		}
	}
	return nil
}
