package system

import (
	"gin-pro/global"
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"go.uber.org/zap"
	"gorm.io/gorm"
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
	syncedCachedEnforcer *casbin.SyncedCachedEnforcer
	once                 sync.Once
)

// 初始化casbin访问策略控制
func (c *CasbinService) Casbin() *casbin.SyncedCachedEnforcer {
	//整个生命周期只执行一次方法
	once.Do(func() {
		a, err := gormadapter.NewAdapterByDB(global.GVA_DB)
		if err != nil {
			zap.L().Error("适配数据库失败请检查casbin表是否为InnoDB引擎!", zap.Error(err))
			return
		}
		//casbin控制模型
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
		m = r.sub == p.sub && keyMatch2(r.obj,p.obj) && r.act == p.act
		`
		m, err := model.NewModelFromString(text)
		if err != nil {
			zap.L().Error("字符串加载模型失败!", zap.Error(err))
			return
		}
		//创建了一个同步的缓存执行器
		syncedCachedEnforcer, _ = casbin.NewSyncedCachedEnforcer(m, a)
		//设置缓存的过期时间，如果策略未更改，执行器将使用缓存的结果。
		syncedCachedEnforcer.SetExpireTime(60 * 60)
		//将存储在适配器中的策略加载到执行器中，以便进行访问控制检查。
		_ = syncedCachedEnforcer.LoadPolicy()
	})
	return syncedCachedEnforcer
}
