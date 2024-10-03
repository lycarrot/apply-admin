package system

import (
	"gin-pro/global"
	"gin-pro/model/common/response"
	"gin-pro/model/system"
	systemRes "gin-pro/model/system/response"
	"gin-pro/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type AuthorityApi struct {
}

// CreateAuthority
// @Tags      Authority
// @Summary   创建角色
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param	  data body system.SysAuthority        true  "权限id, 权限名, 父角色id"
// @Success   200   {object}  response.Response{data=systemRes.SysAuthorityResponse,msg=string}  "创建角色,返回包括系统角色详情"
// @Router    /authority/create [post]
func (a *AuthorityApi) CreateAuthority(c *gin.Context) {
	var authority, authBack system.SysAuthority
	var err error
	if err = c.ShouldBindJSON(&authority); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err = utils.Verify(authority, utils.AuthorityVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if *authority.ParentId == 0 && global.GVA_CONFIG.System.UseStrictAuth {
		authority.ParentId = utils.Pointer(utils.GetUserAuthorityId(c))
	}

	if authBack, err = authorityService.CreateAuthority(authority); err != nil {
		global.GVA_LOG.Error("创建失败！!", zap.Error(err))
		response.FailWithMessage("创建失败"+err.Error(), c)
		return
	}
	//刷新casbin规则
	err = casbinService.FreshCasbin()

	if err != nil {
		global.GVA_LOG.Error("创建成功，权限刷新失败。", zap.Error(err))
		response.FailWithMessage("创建成功，权限刷新失败。"+err.Error(), c)
		return
	}

	response.OkWithDetailed(systemRes.SysAuthorityResponse{Authority: authBack}, "创建成功", c)

}

func (a *AuthorityApi) GetAuthorityList(c *gin.Context) {}
