package dto

import (
	"go-admin/core/dto"
)

type SysUserQueryReq struct {
	dto.Pagination `search:"-"`
	Id             int     `form:"userId" search:"type:exact;column:id;table:admin_sys_user" comment:"用户ID"`
	Username       string  `form:"username" search:"type:contains;column:username;table:admin_sys_user" comment:"用户名"`
	NickName       string  `form:"-" search:"type:contains;column:nick_name;table:admin_sys_user" comment:"昵称"`
	Phone          string  `form:"phone" search:"type:exact;column:phone;table:admin_sys_user" comment:"手机号"`
	RoleId         int64   `form:"-" search:"type:exact;column:role_id;table:admin_sys_user" comment:"角色ID"`
	RoleIds        []int64 `form:"-" search:"type:in;column:role_id;table:admin_sys_user" comment:"角色ID集合"`
	Sex            string  `form:"-" search:"type:exact;column:sex;table:admin_sys_user" comment:"性别"`
	Email          string  `form:"email" search:"type:exact;column:email;table:admin_sys_user" comment:"邮箱"`
	PostId         int64   `form:"postId" search:"type:exact;column:post_id;table:admin_sys_user" comment:"岗位"`
	DeptId         int64   `form:"deptId" search:"type:exact;column:dept_id;table:admin_sys_user" comment:"部门"`
	Status         string  `form:"status" search:"type:exact;column:status;table:admin_sys_user" comment:"状态"`
	DeptJoin       `search:"type:inner;on:id:dept_id;table:admin_sys_user;join:admin_sys_dept"`
	SysUserOrder
}

type SysUserOrder struct {
	IdOrder        string `search:"type:order;column:id;table:admin_sys_user" form:"userIdOrder"`
	UsernameOrder  string `search:"type:order;column:username;table:admin_sys_user" form:"usernameOrder"`
	StatusOrder    string `search:"type:order;column:status;table:admin_sys_user" form:"statusOrder"`
	CreatedAtOrder string `search:"type:order;column:created_at;table:admin_sys_user" form:"createdAtOrder"`
}

type DeptJoin struct {
	Id string `search:"type:contains;column:dept_path;table:admin_sys_dept" form:"id"`
}

func (m *SysUserQueryReq) GetNeedSearch() interface{} {
	return *m
}

type ResetSysUserPwdReq struct {
	UserId     int64  `json:"userId" comment:"用户ID" binding:"required"` // 用户ID
	Password   string `json:"password" comment:"密码" binding:"required"`
	CurrUserId int64  `json:"-" comment:""`
}

type SysUserAvatarUpdateReq struct {
	Avatar     string `json:"avatar" comment:"头像"`
	CurrUserId int64  `json:"-" comment:""`
}

type SysUserStatusUpdateReq struct {
	UserId     int64  `json:"userId" comment:"用户ID"` // 用户ID
	Status     string `json:"status" comment:"状态"`
	CurrUserId int64  `json:"-" comment:""`
}

type SysUserInsertReq struct {
	UserId     int64  `json:"userId" comment:"用户ID"` // 用户ID
	Username   string `json:"username" comment:"用户名"`
	Password   string `json:"password" comment:"密码"`
	NickName   string `json:"nickName" comment:"昵称"`
	Phone      string `json:"phone" comment:"手机号"`
	RoleId     int    `json:"roleId" comment:"角色ID"`
	Avatar     string `json:"avatar" comment:"头像"`
	Sex        string `json:"sex" comment:"性别"`
	Email      string `json:"email" comment:"邮箱"`
	DeptId     int    `json:"deptId" comment:"部门"`
	PostId     int    `json:"postId" comment:"岗位"`
	Remark     string `json:"remark" comment:"备注"`
	Status     string `json:"status" comment:"状态"`
	CurrUserId int64  `json:"-" comment:""`
}

type SysUserUpdateReq struct {
	Id         int64  `uri:"id" json:"-" comment:"用户编号"`
	Username   string `json:"username" comment:"用户名"`
	NickName   string `json:"nickName" comment:"昵称"`
	Phone      string `json:"phone" comment:"手机号"`
	RoleId     int    `json:"roleId" comment:"角色ID"`
	Avatar     string `json:"avatar" comment:"头像"`
	Sex        string `json:"sex" comment:"性别"`
	Email      string `json:"email" comment:"邮箱"`
	DeptId     int    `json:"deptId" comment:"部门"`
	PostId     int    `json:"postId" comment:"岗位"`
	Remark     string `json:"remark" comment:"备注"`
	Status     string `json:"status" comment:"状态"sss`
	CurrUserId int64  `json:"-" comment:""`
	//Password   string `json:"password" comment:""`
}

type SysUserPhoneUpdateReq struct {
	CurrUserId int64  `json:"-" comment:""`
	Phone      string `json:"phone" comment:"手机号"`
}

type SysUserNickNameUpdateReq struct {
	CurrUserId int64  `json:"-" comment:""`
	NickName   string `json:"nickName" comment:"昵称"`
}

type SysUserUpdateEmailReq struct {
	CurrUserId int64  `json:"-" comment:""`
	Email      string `json:"email" comment:"邮箱号"`
}

type SysUserGetReq struct {
	Id         int64 `uri:"id" json:"-" comment:"用户编号"`
	CurrUserId int64 `json:"-" comment:""`
}

type SysUserResp struct {
	Id          int64    `json:"id" comment:"用户ID"` // 用户ID
	Username    string   `json:"username"`
	Avatar      string   `json:"avatar"`
	Phone       string   `json:"phone"`
	Sex         string   `json:"sex"`
	Email       string   `json:"email"`
	DeptName    string   `json:"deptName"`
	RoleName    string   `json:"roleName"`
	Permissions []string `json:"permissions"`
	RoleKyes    []string `json:"roleKeys"`
	CreatedAt   string   `json:"createdAt"`
}

type UpdateSysUserPwdReq struct {
	OldPassword string `json:"oldPassword" comment:"旧密码"`
	NewPassword string `json:"newPassword" comment:"新密码"`
	CurrUserId  int64  `json:"-" comment:""`
}

type LoginReq struct {
	Username string `form:"username" json:"username"`
	Password string `form:"password" json:"password"`
	Code     string `form:"code" json:"code"`
	UUID     string `form:"uuid" json:"uuid"`
}

// SysUserDeleteReq 功能删除请求参数
type SysUserDeleteReq struct {
	Ids []int64 `json:"ids"`
}
