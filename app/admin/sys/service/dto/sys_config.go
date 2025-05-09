package dto

import (
	"go-admin/core/dto"
)

// SysConfigQueryReq 列表或者搜索使用结构体
type SysConfigQueryReq struct {
	dto.Pagination `search:"-"`
	ConfigName     string `form:"configName" search:"type:contains;column:config_name;table:admin_sys_config"`
	ConfigKey      string `form:"configKey" search:"type:contains;column:config_key;table:admin_sys_config"`
	ConfigType     string `form:"configType" search:"type:exact;column:config_type;table:admin_sys_config"`
	IsFrontend     string `form:"isFrontend" search:"type:exact;column:is_frontend;table:admin_sys_config"`
	BeginCreatedAt string `form:"beginCreatedAt" search:"type:gte;column:created_at;table:admin_sys_config" comment:"创建时间"`
	EndCreatedAt   string `form:"endCreatedAt" search:"type:lte;column:created_at;table:admin_sys_config" comment:"创建时间"`
	SysConfigOrder
}

type SysConfigOrder struct {
	IdOrder         string `search:"type:order;column:id;table:admin_sys_config" form:"idOrder"`
	ConfigNameOrder string `search:"type:order;column:config_name;table:admin_sys_config" form:"configNameOrder"`
	ConfigKeyOrder  string `search:"type:order;column:config_key;table:admin_sys_config" form:"configKeyOrder"`
	ConfigTypeOrder string `search:"type:order;column:config_type;table:admin_sys_config" form:"configTypeOrder"`
	CreatedAtOrder  string `search:"type:order;column:created_at;table:admin_sys_config" form:"createdAtOrder"`
}

func (m *SysConfigQueryReq) GetNeedSearch() interface{} {
	return *m
}

// SysConfigInsertReq 增、改使用的结构体
type SysConfigInsertReq struct {
	ConfigName  string `json:"configName" comment:""`
	ConfigKey   string `uri:"configKey" json:"configKey" comment:""`
	ConfigValue string `json:"configValue" comment:""`
	ConfigType  string `json:"configType" comment:""`
	IsFrontend  string `json:"isFrontend"`
	Remark      string `json:"remark" comment:""`
	CurrUserId  int64  `json:"-" comment:""`
}

type SysConfigUpdateReq struct {
	Id          int64  `uri:"id" comment:"编码"` // 编码
	ConfigName  string `json:"configName" comment:""`
	ConfigKey   string `uri:"configKey" json:"configKey" comment:""`
	ConfigValue string `json:"configValue" comment:""`
	ConfigType  string `json:"configType" comment:""`
	IsFrontend  string `json:"isFrontend"`
	Remark      string `json:"remark" comment:""`
	CurrUserId  int64  `json:"-" comment:""`
}

// SysConfigByKeyReq 根据Key获取配置
type SysConfigByKeyReq struct {
	ConfigKey string `uri:"configKey" search:"type:contains;column:config_key;table:admin_sys_config"`
}

func (m *SysConfigByKeyReq) GetNeedSearch() interface{} {
	return *m
}

type SysConfigByKeyResp struct {
	ConfigKey   string `json:"configKey" comment:""`
	ConfigValue string `json:"configValue" comment:""`
}

type SysConfigGetReq struct {
	Id int64 `uri:"id" json:"-"`
}

// SysConfigDeleteReq 获取单个或者删除的结构体
type SysConfigDeleteReq struct {
	Ids []int64 `json:"ids"`
}
