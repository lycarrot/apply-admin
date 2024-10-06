package system

import (
	"errors"
	"gin-pro/global"
	"gin-pro/model/system/request"
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	gormadapter "github.com/casbin/gorm-adapter/v3"

	"go.uber.org/zap"
	"gorm.io/gorm"
	"strconv"
	"sync"
)

type CasbinService struct{}

var CasbinServiceApp = new(CasbinService)

func (c *CasbinService) AddPolicies(db *gorm.DB, rules [][]string) error {
	var casbinRules []gormadapter.CasbinRule
	for i := range rules {
		casbinRules = append(casbinRules, gormadapter.CasbinRule{
			Ptype: "p",
			V0:    rules[i][0],
			V1:    rules[i][1],
			V2:    rules[i][2],
		})
	}
	return db.Create(&casbinRules).Error
}

func (c *CasbinService) FreshCasbin() (err error) {
	e := c.Casbin()
	err = e.LoadPolicy()
	return err
}

var (
	once                 sync.Once
	syncedCachedEnforcer *casbin.SyncedCachedEnforcer
)

func (casbinService *CasbinService) Casbin() *casbin.SyncedCachedEnforcer {
	once.Do(func() {
		a, err := gormadapter.NewAdapterByDB(global.GVA_DB)
		if err != nil {
			zap.L().Error("适配数据库失败，请检查 casbin 表是否为 InnoDB 引擎！", zap.Error(err))
			return
		}

		text := `
		[request_definition]
		r = sub, obj, act
		
		[policy_definition]
		p = sub, obj, act
		
		[role_definition]
		g = _, _
		
		[policy_effect]
		e = some(where (p.eft == allow))
		
		[matchers]
		m = r.sub == p.sub && keyMatch2(r.obj, p.obj) && r.act == p.act
		`
		m, err := model.NewModelFromString(text)
		if err != nil {
			zap.L().Error("字符串加载模型失败！", zap.Error(err))
			return
		}

		syncedCachedEnforcer, err = casbin.NewSyncedCachedEnforcer(m, a)
		if err != nil {
			zap.L().Error("创建 SyncedCachedEnforcer 失败！", zap.Error(err))
			return
		}
		syncedCachedEnforcer.SetExpireTime(60 * 60)
		_ = syncedCachedEnforcer.LoadPolicy()
	})

	if syncedCachedEnforcer == nil {
		zap.L().Error("SyncedCachedEnforcer 未正确初始化")
	}
	return syncedCachedEnforcer
}

// @function: UpdateCasbin
// @description: 更新casbin权限
// @param: authorityId string, casbinInfos []request.CasbinInfo
// @return: error
func (c *CasbinService) UpdateCasbin(authorityId uint, casbinInfo []request.CasbinInfo) error {
	casLists := map[string]bool{}
	casRules := [][]string{}
	newAuthorityId := strconv.Itoa(int(authorityId))
	for _, v := range casbinInfo {
		key := newAuthorityId + v.Path + v.Method
		_, ok := casLists[key]
		if !ok {
			casLists[key] = true
			casRules = append(casRules, []string{newAuthorityId, v.Path, v.Method})
		}
	}
	if len(casRules) == 0 {
		return nil
	}
	e := c.Casbin()
	success, _ := e.AddPolicies(casRules)
	if !success {
		return errors.New("存在相同api,添加失败")
	}
	return nil
}
