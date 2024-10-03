package system

import (
	"gin-pro/global"
	"gin-pro/model/common/response"
	"gin-pro/model/system"
	"gin-pro/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type ApiApi struct {
}

// CreateApi
// @Tags ApiApi
// @Summary 创建api
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      system.SysApi       true  "api路径, api中文描述, api组, 方法"
// @Success    200   {object}  response.Response{msg=string}  "创建基础api"
// @Router  /api/create [post]
func (a *ApiApi) CreateApi(c *gin.Context) {
	var r system.SysApi
	err := c.ShouldBind(&r)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(r, utils.ApiVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = apiService.CreateApi(r)
	if err != nil {
		global.GVA_LOG.Error("api创建失败", zap.Error(err))
		response.FailWithMessage("api创建失败", c)
		return
	}
	response.OkWithMessage("api创建成功", c)
}
