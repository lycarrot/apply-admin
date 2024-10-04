package system

import (
	"gin-pro/global"
	"gin-pro/model/common/response"
	"gin-pro/model/system/request"
	"gin-pro/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type CasbinApi struct{}

// UpdateCasbin
// @Tags      Casbin
// @Summary   更新角色api权限
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      request.CasbinInReceive        true  "权限id, 权限模型列表"
// @Success   200   {object}  response.Response{msg=string}  "更新角色api权限"
// @Router    /casbin/update [post]
func (c *CasbinApi) UpdateCasbin(g *gin.Context) {
	var r request.CasbinInReceive
	err := g.ShouldBindJSON(&r)
	if err != nil {
		response.FailWithMessage(err.Error(), g)
		return
	}
	err = utils.Verify(r, utils.CasbinVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), g)
		return
	}
	//authorityId := utils.GetUserAuthorityId(g)
	err = casbinService.UpdateCasbin(r.AuthorityId, r.CasbinInfos)
	if err != nil {
		response.FailWithMessage("添加失败", g)
		global.GVA_LOG.Error("添加失败", zap.Error(err))
		return
	}
	response.OkWithMessage("添加成功", g)
}
