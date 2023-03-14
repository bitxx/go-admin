package dto

import (
	"go-admin/common/dto"
	"time"
)

type MsgCodeQueryReq struct {
	dto.Pagination `search:"-"`
	UserId         int64  `form:"userId"  search:"type:exact;column:user_id;table:plugins_msg_code" comment:"用户编号"`
	CodeType       string `form:"codeType"  search:"type:exact;column:code_type;table:plugins_msg_code" comment:"验证码类型 1-邮箱；2-短信"`
	Status         string `form:"status"  search:"type:exact;column:status;table:plugins_msg_code" comment:"验证码状态 1-发送成功 2-发送失败"`
	BeginTime      string `form:"beginTime" search:"type:gte;column:created_at;table:plugins_msg_code" comment:"创建时间"`
	EndTime        string `form:"endTime" search:"type:lte;column:created_at;table:plugins_msg_code" comment:"创建时间"`
	MsgCodeOrder
}

type MsgCodeOrder struct {
	IdOrder        int64      `form:"idOrder"  search:"type:order;column:id;table:plugins_msg_code"`
	UserIdOrder    int64      `form:"userIdOrder"  search:"type:order;column:user_id;table:plugins_msg_code"`
	CodeOrder      string     `form:"codeOrder"  search:"type:order;column:code;table:plugins_msg_code"`
	CodeTypeOrder  string     `form:"codeTypeOrder"  search:"type:order;column:code_type;table:plugins_msg_code"`
	RemarkOrder    string     `form:"remarkOrder"  search:"type:order;column:remark;table:plugins_msg_code"`
	StatusOrder    string     `form:"statusOrder"  search:"type:order;column:status;table:plugins_msg_code"`
	CreateByOrder  int64      `form:"createByOrder"  search:"type:order;column:create_by;table:plugins_msg_code"`
	UpdateByOrder  int64      `form:"updateByOrder"  search:"type:order;column:update_by;table:plugins_msg_code"`
	CreatedAtOrder *time.Time `form:"createdAtOrder"  search:"type:order;column:created_at;table:plugins_msg_code"`
	UpdatedAtOrder *time.Time `form:"updatedAtOrder"  search:"type:order;column:updated_at;table:plugins_msg_code"`
}

func (m *MsgCodeQueryReq) GetNeedSearch() interface{} {
	return *m
}

// MsgCodeGetReq 功能获取请求参数
type MsgCodeGetReq struct {
	Id int64 `uri:"id"`
}
