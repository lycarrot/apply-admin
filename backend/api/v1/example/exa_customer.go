package example

import (
	"gin-pro/model/common/response"
	"github.com/gin-gonic/gin"
)

type CustomerApi struct {
}

// CreateExaCustomer
// @Tags      ExaCustomer
// @Summary   创建客户
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      example.ExaCustomer            true  "客户用户名, 客户手机号码"
// @Success   200   {object}  response.Response{msg=string}  "创建客户"
// @Router    /customer/customer [post]
func (e *CustomerApi) CreateExaCustomer(c *gin.Context) {
	response.OkWithMessage("创建成功", c)
}

// DeleteExaCustomer
// @Tags      ExaCustomer
// @Summary   删除客户
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      example.ExaCustomer            true  "客户ID"
// @Success   200   {object}  response.Response{msg=string}  "删除客户"
// @Router    /customer/customer [delete]
func (e *CustomerApi) DeleteExaCustomer(c *gin.Context) {
	response.OkWithMessage("删除成功", c)
}

// UpdateExaCustomer
// @Tags      ExaCustomer
// @Summary   更新客户信息
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      example.ExaCustomer            true  "客户ID, 客户信息"
// @Success   200   {object}  response.Response{msg=string}  "更新客户信息"
// @Router    /customer/customer [put]
func (e *CustomerApi) UpdateExaCustomer(c *gin.Context) {
	response.OkWithMessage("更新成功", c)
}

// GetExaCustomer
// @Tags      ExaCustomer
// @Summary   获取单一客户信息
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  query     example.ExaCustomer                                                true  "客户ID"
// @Success   200   {object}  response.Response{msg=string}  "更新客户信息"
// @Router    /customer/customer [get]
func (e *CustomerApi) GetExaCustomer(c *gin.Context) {
	response.OkWithMessage("更新成功", c)
}

// GetExaCustomerList
// @Tags      ExaCustomer
// @Summary   分页获取权限客户列表
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// / @Success   200   {object}  response.Response{msg=string}  "更新客户信息"
// @Router    /customer/customerList [get]
func (e *CustomerApi) GetExaCustomerList(c *gin.Context) {
	response.OkWithMessage("更新成功", c)
}
