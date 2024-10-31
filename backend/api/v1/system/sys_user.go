package system

import (
	"gin-pro/global"
	"gin-pro/model/common/response"
	"gin-pro/model/system/request"
	"gin-pro/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
)

type UserApi struct {
}

// GetUserLists
// @Tags      User
// @Summary   获取用户列表
// @Security  ApiKeyAuth
// @Produce   application/json
// @Param     data  query       request.SysUserQuery        true  "query"
// @Success   200   {object}  response.Response{data=response.PageResult,msg=string}   "列表获取成功"
// @Router    /user/getLists [get]
func (u *UserApi) GetUserLists(c *gin.Context) {
	var query request.SysUserQuery
	err := c.ShouldBindQuery(&query)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(query, utils.PageVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	lists, total, err := userService.GetUserLists(query)
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

// UpdateUser
// @Tags      User
// @Summary   更新用户信息
// @Security  ApiKeyAuth
// @Produce   application/json
// @Param     data body      request.SysUserQuery        true  "body"
// @Success   200   {object}  response.Response{data=response.PageResult,msg=string}   "用户信息更新成功"
// @Router    /user/update [post]
func (u *UserApi) UpdateUser(c *gin.Context) {

}

// GetUseDetail
// @Tags      User
// @Summary   获取用户信息
// @Security  ApiKeyAuth
// @Produce   application/json
// @Param     data query  string   true  "用户id"
// @Success   200   {object}  response.Response{data=response.PageResult,msg=string}   "用户信息更新成功"
// @Router    /user/detail/:id [get]
func (u *UserApi) GetUseDetail(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response.FailWithMessage("参数获取失败", c)
		global.GVA_LOG.Error("参数获取失败", zap.Error(err))
		return
	}
	user, err := userService.GetUseDetail(id)
	if err != nil {
		global.GVA_LOG.Error("获取失败", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}
	response.OkWithDetailed(user, "获取成功", c)
}

func (u *UserApi) DeleteUser(c *gin.Context) {}
