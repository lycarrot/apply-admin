package middleware

import (
	"fmt"
	"gin-pro/global"
	"gin-pro/model/common/response"
	"gin-pro/service"
	"gin-pro/utils"
	"github.com/gin-gonic/gin"
	"strconv"
	"strings"
)

// CasbinHandler 拦截器
func CasbinHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		claims, _ := utils.GetClaims(c)
		var casbinSerice = service.ServiceGroupApp.SystemServiceGroup.CasbinService
		e := casbinSerice.Casbin()
		sub := strconv.Itoa(int(claims.AuthorityId))
		obj := strings.TrimPrefix(c.Request.URL.Path, global.GVA_CONFIG.System.RouterPrefix)
		act := c.Request.Method
		success, err := e.Enforce(sub, obj, act)
		fmt.Print(err)
		if !success {
			response.FailWithDetailed(gin.H{}, "权限不足", c)
			c.Abort()
		}
		c.Next()
	}
}
