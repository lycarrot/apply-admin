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
// @Param	   data  query   request.SysOperationQuery	true "query"
// @success   200  {object}	 response.Response{data=response.PageResult,msg=string}  "分页获取操作历史列表"
// @Router    /record/lists [get]
func (o *OperationRecordApi) GetRecordLists(c *gin.Context) {
	var query request.SysOperationQuery
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
