package models

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"time"
)

type SysUser struct {
	Id        int64      `gorm:"primaryKey;autoIncrement;comment:编码"  json:"id"`
	Username  string     `json:"username" gorm:"size:64;comment:用户名"`
	Password  string     `json:"-" gorm:"size:128;comment:密码"`
	NickName  string     `json:"nickName" gorm:"size:128;comment:昵称"`
	Phone     string     `json:"phone" gorm:"size:11;comment:手机号"`
	RoleId    int        `json:"roleId" gorm:"size:20;comment:角色ID"`
	Salt      string     `json:"-" gorm:"size:255;comment:加盐"`
	Avatar    string     `json:"avatar" gorm:"size:255;comment:头像"`
	Sex       string     `json:"sex" gorm:"size:255;comment:性别"`
	Email     string     `json:"email" gorm:"size:128;comment:邮箱"`
	DeptId    int        `json:"deptId" gorm:"size:20;comment:部门"`
	PostId    int        `json:"postId" gorm:"size:20;comment:岗位"`
	Remark    string     `json:"remark" gorm:"size:255;comment:备注"`
	Status    string     `json:"status" gorm:"size:4;comment:状态"`
	Dept      *SysDept   `json:"dept"`
	Post      *SysPost   `json:"post"`
	Role      *SysRole   `json:"role"`
	CreateBy  int64      `json:"createBy" gorm:"index;comment:创建者"`
	UpdateBy  int64      `json:"updateBy" gorm:"index;comment:更新者"`
	CreatedAt *time.Time `json:"createdAt" gorm:"comment:创建时间"`
	UpdatedAt *time.Time `json:"updatedAt" gorm:"comment:最后更新时间"`
}

func (SysUser) TableName() string {
	return "sys_user"
}

// 加密
func (e *SysUser) Encrypt() (err error) {
	if e.Password == "" {
		return
	}

	var hash []byte
	if hash, err = bcrypt.GenerateFromPassword([]byte(e.Password), bcrypt.DefaultCost); err != nil {
		return
	} else {
		e.Password = string(hash)
		return
	}
}

func (e *SysUser) BeforeCreate(_ *gorm.DB) error {
	return e.Encrypt()
}

func (e *SysUser) BeforeUpdate(_ *gorm.DB) error {
	var err error
	if e.Password != "" {
		err = e.Encrypt()
	}
	return err
}

func (e *SysUser) AfterFind(_ *gorm.DB) error {
	/*e.DeptIds = []int{e.Id}
	e.PostIds = []int{e.PostId}
	e.RoleIds = []int{e.Id}*/
	return nil
}
