package system

import (
	"gin-pro/global"
	"gin-pro/model/common/response"
	"gin-pro/model/system/request"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type OperationRecordApi struct {
}

// GetLists
// @Tags      OperationRecord
// @Summary   获取操作历史记录
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param	   data  query   request.SysOperationRecordSearch	true "路由path, 父菜单ID, 路由name, 对应前端文件路径, 排序标记"
// @success   200  {object}	 response.Response{data=response.PageResult,msg=string}  "分页获取操作历史列表"
// @Router    /record/lists [get]
func (o *OperationRecordApi) GetLists(c *gin.Context) {
	var query request.SysOperationRecordSearch
	err := c.ShouldBindQuery(&query)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	lists, total, err := operationRecordService.GetOperationRecordLists(query)
	if err != nil {
		global.GVA_LOG.Error("查询失败", zap.Error(err))
		response.FailWithMessage("查询失败", c)
		return
	}
	response.OkWithDetailed(response.PageResult{
		Lists:    lists,
		Total:    total,
		Page:     query.Page,
		PageSize: query.PageSize,
	}, "获取成功", c)

}
