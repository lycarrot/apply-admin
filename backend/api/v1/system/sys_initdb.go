package system

import (
	"gin-pro/global"
	"gin-pro/model/common/response"
	"gin-pro/model/system/request"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type DBApi struct {
}

// InitDB
// @Tags     InitDB
// @Summary  初始化数据库
// @Produce  application/json
// @Param data body request.InitDB  true "初始化数据库参数"
// @Success  200   {object}  response.Response{data=string}  "初始化用户数据库"
// @Router   /init/initdb [post]
func (i *DBApi) InitDB(c *gin.Context) {
	if global.GVA_DB != nil {
		global.GVA_LOG.Error("已存在数据库配置!")
		response.FailWithMessage("已存在数据库配置", c)
		return
	}
	var dbInfo request.InitDB
	if err := c.ShouldBindJSON(&dbInfo); err != nil {
		global.GVA_LOG.Error("参数校验不通过!", zap.Error(err))
		response.FailWithMessage("参数校验不通过", c)
		return
	}
	response.OkWithMessage("自动创建数据库成功", c)
}

// CheckDB
// @Tags     CheckDB
// @Summary  查询用户数据库存在
// @Produce  application/json
// @Success  200  {object}  response.Response{data=map[string]interface{},msg=string}  "查询用户数据库存在"
// @Router   /init/checkdb [post]
func (i *DBApi) CheckDB(c *gin.Context) {
	var (
		message  = "前往初始化数据库"
		needInit = true
	)
	if global.GVA_DB != nil {
		message = "数据库无需初始化"
		needInit = false
	}
	global.GVA_LOG.Info(message)
	response.OkWithDetailed(gin.H{"needInit": needInit}, "message", c)

}
