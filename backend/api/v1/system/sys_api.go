package system

import (
	"gin-pro/global"
	"gin-pro/model/common/request"
	"gin-pro/model/common/response"
	"gin-pro/model/system"
	sysReq "gin-pro/model/system/request"
	"gin-pro/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
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
	err := c.ShouldBindJSON(&r)
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
		global.GVA_LOG.Error("创建失败", zap.Error(err))
		response.FailWithMessage("创建失败", c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

// UpdateApi
// @Tags Api
// @Summary 更新api
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      system.SysApi       true  "api路径, api中文描述, api组, 方法"
// @Success    200   {object}  response.Response{msg=string}  "更新api"
// @Router  /api/update [put]
func (a *ApiApi) UpdateApi(c *gin.Context) {
	var r system.SysApi
	err := c.ShouldBindJSON(&r)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(r, utils.ApiVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = apiService.UpdateApi(r)
	if err != nil {
		global.GVA_LOG.Error("更新失败", zap.Error(err))
		response.FailWithMessage("更新失败", c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// BatchDelApis
// @Tags Api
// @Summary 删除api
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body    request.IdsQuery  true  "Ids 集合"
// @Success    200   {object}  response.Response{msg=string}  "批量删除api"
// @Router  /api/batchDel [delete]
func (a *ApiApi) BatchDelApis(c *gin.Context) {
	var ids request.IdsQuery
	err := c.ShouldBindJSON(&ids)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = apiService.BatchDel(ids)
	if err != nil {
		response.FailWithMessage("删除失败", c)
		global.GVA_LOG.Error("删除失败", zap.Error(err))
		return
	}
	response.OkWithMessage("删除成功", c)

}

// GetApiDetail
// @Tags Api
// @Summary 创建api
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     id  query   string    true  "api id"
// @Success    200   {object}  response.Response{data=system.SysApi}  "获取api详情"
// @Router  /api/detail/:id [get]
func (a *ApiApi) GetApiDetail(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response.FailWithMessage("参数获取失败", c)
		global.GVA_LOG.Error("参数获取失败", zap.Error(err))
		return
	}
	api, err := apiService.GetDetail(id)
	if err != nil {
		response.FailWithMessage("获取失败", c)
		global.GVA_LOG.Error("获取失败", zap.Error(err))
		return
	}
	response.OkWithDetailed(api, "获取成功", c)
}

// GetApiLists
// @Tags Api
// @Summary 获取api列表
// @Security  ApiKeyAuth
// @Produce   application/json
// @Param     data  body  sysReq.SysApiQuery      true  "api路径, api中文描述, api组, 方法"
// @success   200  {object}	 response.Response{data=response.PageResult,msg=string}  "分页获取列表"
// @Router  /api/getLists [post]
func (a *ApiApi) GetApiLists(c *gin.Context) {
	var query sysReq.SysApiQuery
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
