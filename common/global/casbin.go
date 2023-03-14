package global

import (
	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	"go-admin/common/core"
	"go-admin/common/core/api"
)

func LoadPolicy(c *gin.Context) (*casbin.SyncedEnforcer, error) {
	log := api.GetRequestLogger(c)
	if err := core.Runtime.GetCasbinKey(c.Request.Host).LoadPolicy(); err == nil {
		return core.Runtime.GetCasbinKey(c.Request.Host), err
	} else {
		log.Errorf("casbin rbac_model or policy init error, %s ", err.Error())
		return nil, err
	}
}
