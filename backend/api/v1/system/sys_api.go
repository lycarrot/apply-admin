package system

import (
	"gin-pro/global"
	"gin-pro/model/common/response"
	"gin-pro/model/system"
	"gin-pro/model/system/request"
	"gin-pro/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type ApiApi struct {
}

// CreateApi
// @Tags Api
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

// GetApiLists
// @Tags Api
// @Summary 获取api列表
// @Security  ApiKeyAuth
// @Produce   application/json
// @Param     data  body  request.SysApiQuery      true  "api路径, api中文描述, api组, 方法"
// @success   200  {object}	 response.Response{data=response.PageResult,msg=string}  "分页获取列表"
// @Router  /api/getLists [post]
func (a *ApiApi) GetApiLists(c *gin.Context) {
	var query request.SysApiQuery
	err := c.ShouldBindJSON(&query)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(query, utils.PageVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	lists, total, err := apiService.GetApiLists(query)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}
	response.OkWithDetailed(response.PageResult{
		Lists:    lists,
		Page:     query.Page,
		PageSize: query.PageSize,
		Total:    total,
	}, "获取成功", c)
}
