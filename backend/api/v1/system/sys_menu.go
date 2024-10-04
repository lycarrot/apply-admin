package system

import (
	"gin-pro/global"
	"gin-pro/model/common/response"
	"gin-pro/model/system"
	"gin-pro/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type MenuApi struct{}

// CreateMenuItem
// @Tags      Menu
// @Summary   新增菜单
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param	  data body system.SysBaseMenu	true "路由path, 父菜单ID, 路由name, 对应前端文件路径, 排序标记"
// @success   200  {object}	 response.Response{msg=string}  "新增菜单"
// @Router    /menu/create [post]
func (m *MenuApi) CreateMenuItem(c *gin.Context) {
	var menu system.SysBaseMenu
	err := c.ShouldBindJSON(&menu)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(menu, utils.MenuVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(menu, utils.MenuMetaVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = menuService.AddMenuItem(menu)
	if err != nil {
		global.GVA_LOG.Error("添加失败!", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithMessage("添加成功", c)

}

func (m *MenuApi) DelMenuItem(c *gin.Context) {

}

func (m *MenuApi) UpdateMenuItem(c *gin.Context) {

}
func (m *MenuApi) GetMenuLists(c *gin.Context) {

}
