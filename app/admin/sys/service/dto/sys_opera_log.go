package dto

import (
	"time"

	"go-admin/core/dto"
)

type SysOperLogQueryReq struct {
	dto.Pagination `search:"-"`
	Title          string `form:"title" search:"type:contains;column:title;table:admin_sys_oper_log" comment:"操作模块"`
	Method         string `form:"method" search:"type:contains;column:method;table:admin_sys_oper_log" comment:"函数"`
	RequestMethod  string `form:"requestMethod" search:"type:contains;column:request_method;table:admin_sys_oper_log" comment:"请求方式"`
	OperUrl        string `form:"operUrl" search:"type:contains;column:oper_url;table:admin_sys_oper_log" comment:"访问地址"`
	OperIp         string `form:"operIp" search:"type:exact;column:oper_ip;table:admin_sys_oper_log" comment:"客户端ip"`
	Status         string `form:"status" search:"type:exact;column:status;table:admin_sys_oper_log" comment:"状态"`
	BeginCreatedAt string `form:"beginCreatedAt" search:"type:gte;column:created_at;table:admin_sys_oper_log" comment:"创建时间"`
	EndCreatedAt   string `form:"endCreatedAt" search:"type:lte;column:created_at;table:admin_sys_oper_log" comment:"创建时间"`
	SysOperLogOrder
}

type SysOperLogOrder struct {
	CreatedAtOrder string `search:"type:order;column:created_at;table:admin_sys_oper_log" form:"createdAtOrder"`
}

func (m *SysOperLogQueryReq) GetNeedSearch() interface{} {
	return *m
}

type SysOperLogInsertReq struct {
	Id            int64      `uri:"id" json:"-" `
	Title         string     `json:"title"`
	BusinessType  string     `json:"businessType"`
	BusinessTypes string     `json:"businessTypes"`
	Method        string     `json:"method"`
	RequestMethod string     `json:"requestMethod"`
	OperatorType  string     `json:"operatorType"`
	OperName      string     `json:"operName"`
	DeptName      string     `json:"deptName"`
	OperUrl       string     `json:"operUrl"`
	OperIp        string     `json:"operIp"`
	OperLocation  string     `json:"operLocation"`
	OperParam     string     `json:"operParam"`
	Status        string     `json:"status"`
	OperTime      *time.Time `json:"operTime"`
	JsonResult    string     `json:"jsonResult"`
	Remark        string     `json:"remark"`
	LatencyTime   string     `json:"latencyTime"`
	UserAgent     string     `json:"userAgent"`
	CurrUserId    int64      `json:"-" comment:""`
}

type SysOperLogGetReq struct {
	Id int64 `uri:"id" json:"-"`
}

// SysOperLogDeleteReq 功能删除请求参数
type SysOperLogDeleteReq struct {
	Ids []int64 `json:"ids"`
}
