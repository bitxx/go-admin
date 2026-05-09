package service

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/mssola/user_agent"
	"go-admin/internal/app/admin/sys/models"
	"go-admin/internal/app/admin/sys/service/dto"
	clang "go-admin/internal/common/lang"
	"go-admin/pkg/config"
	"go-admin/pkg/dto/service"
	"go-admin/pkg/global"
	"go-admin/pkg/lang"
	"go-admin/pkg/middleware"
	"go-admin/pkg/middleware/auth/jwtauth"
	"go-admin/pkg/runtime"
	"go-admin/pkg/utils/dateutils"
	"go-admin/pkg/utils/iputils"
	"go-admin/pkg/utils/strutils"
	"gorm.io/gorm"
	"strconv"
	"time"

	cDto "go-admin/pkg/dto"
)

type SysUser struct {
	service.Service
}

// NewSysUserService admin-实例化用户管理
func NewSysUserService(s *service.Service) *SysUser {
	var srv = new(SysUser)
	srv.Orm = s.Orm
	srv.Log = s.Log
	return srv
}

// GetPage admin-获取系统用户管理分页列表
func (e *SysUser) GetPage(c *dto.SysUserQueryReq, p *middleware.DataPermission) ([]models.SysUser, int64, int, error) {
	var list []models.SysUser
	var data models.SysUser
	var count int64

	err := e.Orm.Model(&data).Preload("Dept").Preload("Role").Preload("Post").Order("created_at desc").
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
			middleware.Permission(data.TableName(), p),
		).Find(&list).Limit(-1).Offset(-1).Count(&count).Error
	if err != nil {
		return nil, 0, clang.DataQueryLogCode, lang.MsgLogErrf(e.Log, e.Lang, clang.DataQueryCode, clang.DataQueryLogCode, err)
	}
	return list, count, clang.SuccessCode, nil
}

// Get admin-获取系统用户管理详情
func (e *SysUser) Get(id int64, p *middleware.DataPermission) (*models.SysUser, int, error) {
	if id <= 0 {
		return nil, clang.ParamErrCode, lang.MsgErr(clang.ParamErrCode, e.Lang)
	}
	data := &models.SysUser{}
	err := e.Orm.Scopes(
		middleware.Permission(data.TableName(), p),
	).First(data, id).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, clang.DataQueryLogCode, lang.MsgLogErrf(e.Log, e.Lang, clang.DataQueryCode, clang.DataQueryLogCode, err)
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, clang.DataNotFoundCode, lang.MsgErr(clang.DataNotFoundCode, e.Lang)
	}
	return data, clang.SuccessCode, nil
}

// QueryOne admin-获取系统用户管理一条记录
func (e *SysUser) QueryOne(queryCondition *dto.SysUserQueryReq, p *middleware.DataPermission) (*models.SysUser, int, error) {
	data := &models.SysUser{}
	err := e.Orm.Model(&models.SysUser{}).
		Scopes(
			cDto.MakeCondition(queryCondition.GetNeedSearch()),
			middleware.Permission(data.TableName(), p),
		).First(data).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, clang.DataQueryLogCode, lang.MsgLogErrf(e.Log, e.Lang, clang.DataQueryCode, clang.DataQueryLogCode, err)
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, clang.DataNotFoundCode, lang.MsgErr(clang.DataNotFoundCode, e.Lang)
	}
	return data, clang.SuccessCode, nil
}

// Count admin-获取系统用户管理数据总数
func (e *SysUser) Count(c *dto.SysUserQueryReq) (int64, int, error) {
	var err error
	var count int64
	err = e.Orm.Model(&models.SysUser{}).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
		).Limit(-1).Offset(-1).
		Count(&count).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return 0, clang.DataQueryLogCode, lang.MsgLogErrf(e.Log, e.Lang, clang.DataQueryCode, clang.DataQueryLogCode, err)
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return 0, clang.DataNotFoundCode, lang.MsgErr(clang.DataNotFoundCode, e.Lang)
	}
	return count, clang.SuccessCode, nil
}

// Insert admin-新增系统用户管理
func (e *SysUser) Insert(c *dto.SysUserInsertReq) (int64, int, error) {
	if c.CurrUserId <= 0 {
		return 0, clang.ParamErrCode, lang.MsgErr(clang.ParamErrCode, e.Lang)
	}
	if c.Username == "" {
		return 0, clang.SysUserNameEmptyCode, lang.MsgErr(clang.SysUserNameEmptyCode, e.Lang)
	}
	if c.NickName == "" {
		return 0, clang.SysNickNameEmptyCode, lang.MsgErr(clang.SysNickNameEmptyCode, e.Lang)
	}
	if c.Phone == "" {
		return 0, clang.SysUserPhoneEmptyCode, lang.MsgErr(clang.SysUserPhoneEmptyCode, e.Lang)
	}
	if c.Email == "" {
		return 0, clang.SysUserEmailEmptyCode, lang.MsgErr(clang.SysUserEmailEmptyCode, e.Lang)
	}
	if c.DeptId <= 0 {
		return 0, clang.SysUserDeptEmptyCode, lang.MsgErr(clang.SysUserDeptEmptyCode, e.Lang)
	}
	if c.Password == "" {
		return 0, clang.SysUserPwdEmptyCode, lang.MsgErr(clang.SysUserPwdEmptyCode, e.Lang)
	}

	if c.Username != "" {
		query := dto.SysUserQueryReq{}
		query.Username = c.Username
		count, respCode, err := e.Count(&query)
		if err != nil && respCode != clang.DataNotFoundCode {
			return 0, respCode, err
		}
		if count > 0 {
			return 0, clang.SysUserNameExistCode, lang.MsgErr(clang.SysUserNameExistCode, e.Lang)
		}
	}
	if c.NickName != "" {
		query := dto.SysUserQueryReq{}
		query.NickName = c.NickName
		count, respCode, err := e.Count(&query)
		if err != nil && respCode != clang.DataNotFoundCode {
			return 0, respCode, err
		}
		if count > 0 {
			return 0, clang.SysUserNickNameExistCode, lang.MsgErr(clang.SysUserNickNameExistCode, e.Lang)
		}
	}
	if c.Phone != "" {
		query := dto.SysUserQueryReq{}
		query.Phone = c.Phone
		count, respCode, err := e.Count(&query)
		if err != nil && respCode != clang.DataNotFoundCode {
			return 0, respCode, err
		}
		if count > 0 {
			return 0, clang.SysUserPhoneExistCode, lang.MsgErr(clang.SysUserPhoneExistCode, e.Lang)
		}
	}
	if c.Email != "" {
		query := dto.SysUserQueryReq{}
		query.Email = c.Email
		count, respCode, err := e.Count(&query)
		if err != nil && respCode != clang.DataNotFoundCode {
			return 0, respCode, err
		}
		if count > 0 {
			return 0, clang.SysUserEmailExistCode, lang.MsgErr(clang.SysUserEmailExistCode, e.Lang)
		}
	}

	if c.Avatar == "" {
		sysConfService := NewSysConfigService(&e.Service)
		defaultAvatar, respCode, err := sysConfService.GetWithKeyStr("admin_sys_user_default_avatar")
		if err != nil {
			return 0, respCode, err
		}
		c.Avatar = defaultAvatar
	}

	// insert data
	now := time.Now()
	data := models.SysUser{}
	data.Username = c.Username
	data.Password = c.Password
	data.NickName = c.NickName
	data.Phone = c.Phone
	data.RoleId = c.RoleId
	data.Avatar = c.Avatar
	data.Sex = c.Sex
	data.Email = c.Email
	data.DeptId = c.DeptId
	data.PostId = c.PostId
	data.Status = c.Status
	data.Remark = c.Remark
	data.CreateBy = c.CurrUserId
	data.UpdateBy = c.CurrUserId
	data.CreatedAt = &now
	data.UpdatedAt = &now
	err := e.Orm.Create(&data).Error
	if err != nil {
		return 0, clang.DataInsertLogCode, lang.MsgLogErrf(e.Log, e.Lang, clang.DataInsertCode, clang.DataInsertLogCode, err)
	}
	return data.Id, clang.SuccessCode, nil
}

// Update admin-更新系统用户管理
func (e *SysUser) Update(c *dto.SysUserUpdateReq, p *middleware.DataPermission) (bool, int, error) {
	if c.Id <= 0 || c.CurrUserId <= 0 {
		return false, clang.ParamErrCode, lang.MsgErr(clang.ParamErrCode, e.Lang)
	}
	if c.Username == "" {
		return false, clang.SysUserNameEmptyCode, lang.MsgErr(clang.SysUserNameEmptyCode, e.Lang)
	}
	//if c.NickName == "" {
	//	return false, cLang.SysNickNameEmptyCode, lang.MsgErr(cLang.SysNickNameEmptyCode, e.Lang)
	//}
	if c.Phone == "" {
		return false, clang.SysUserPhoneEmptyCode, lang.MsgErr(clang.SysUserPhoneEmptyCode, e.Lang)
	}
	if c.Email == "" {
		return false, clang.SysUserEmailEmptyCode, lang.MsgErr(clang.SysUserEmailEmptyCode, e.Lang)
	}
	/*	if c.DeptId <= 0 {
		return false, cLang.SysUserDeptEmptyCode, lang.MsgErr(cLang.SysUserDeptEmptyCode, e.Lang)
	}*/

	authChange := false //检查影响用户登录认证的字段是否发生变化，若发生变化，则需要强制该用户退出登录

	data, respCode, err := e.Get(c.Id, p)
	if err != nil {
		return false, respCode, err
	}

	updates := map[string]interface{}{}
	if c.Username != "" && data.Username != c.Username {
		req := dto.SysUserQueryReq{}
		req.Username = c.Username
		resp, respCode, err := e.QueryOne(&req, p)
		if err != nil && respCode != clang.DataNotFoundCode {
			return false, respCode, err
		}
		if respCode == clang.SuccessCode && resp.Id != data.Id {
			return false, clang.SysUserNameExistCode, lang.MsgErr(clang.SysUserNameExistCode, e.Lang)
		}
		updates["username"] = c.Username
	}
	if c.NickName != "" && data.NickName != c.NickName {
		req := dto.SysUserQueryReq{}
		req.NickName = c.NickName
		resp, respCode, err := e.QueryOne(&req, p)
		if err != nil && respCode != clang.DataNotFoundCode {
			return false, respCode, err
		}
		if respCode == clang.SuccessCode && resp.Id != data.Id {
			return false, clang.SysUserNickNameExistCode, lang.MsgErr(clang.SysUserNickNameExistCode, e.Lang)
		}
		updates["nick_name"] = c.NickName
	}
	if c.Phone != "" && data.Phone != c.Phone {
		req := dto.SysUserQueryReq{}
		req.Phone = c.Phone
		resp, respCode, err := e.QueryOne(&req, p)
		if err != nil && respCode != clang.DataNotFoundCode {
			return false, respCode, err
		}
		if respCode == clang.SuccessCode && resp.Id != data.Id {
			return false, clang.SysUserPhoneExistCode, lang.MsgErr(clang.SysUserPhoneExistCode, e.Lang)
		}
		updates["phone"] = c.Phone
	}
	if c.RoleId > 0 && data.RoleId != c.RoleId {
		updates["role_id"] = c.RoleId
		authChange = true //角色切换，需要重新登录
	}
	if c.Avatar != "" && data.Avatar != c.Avatar {
		updates["avatar"] = c.Avatar
	}
	if c.Sex != "" && data.Sex != c.Sex {
		updates["sex"] = c.Sex
	}
	if c.Email != "" && data.Email != c.Email {
		if !strutils.VerifyEmailFormat(c.Email) {
			return false, clang.SysUserEmailFormatErrCode, lang.MsgErr(clang.SysUserEmailFormatErrCode, e.Lang)
		}
		req := dto.SysUserQueryReq{}
		req.Email = c.Email
		resp, respCode, err := e.QueryOne(&req, p)
		if err != nil && respCode != clang.DataNotFoundCode {
			return false, respCode, err
		}
		if respCode == clang.SuccessCode && resp.Id != data.Id {
			return false, clang.SysUserEmailExistCode, lang.MsgErr(clang.SysUserEmailExistCode, e.Lang)
		}
		updates["email"] = c.Email
	}
	if c.DeptId > 0 && data.DeptId != c.DeptId {
		updates["dept_id"] = c.DeptId
	}
	if c.PostId > 0 && data.PostId != c.PostId {
		updates["post_id"] = c.PostId
	}
	if c.Status != "" && data.Status != c.Status {
		updates["status"] = c.Status
	}
	if c.Remark != "" && data.Remark != c.Remark {
		updates["remark"] = c.Remark
	}
	if len(updates) > 0 {
		updates["update_by"] = c.CurrUserId
		updates["updated_at"] = time.Now()
		err = e.Orm.Model(&data).Where("id=?", data.Id).Updates(&updates).Error
		if err != nil {
			return false, clang.DataUpdateLogCode, lang.MsgLogErrf(e.Log, e.Lang, clang.DataUpdateCode, clang.DataUpdateLogCode, err)
		}

		if authChange {
			sysRoleService := NewSysRoleService(&e.Service)
			role, _, _ := sysRoleService.Get(c.RoleId, nil)
			if role != nil {
				//设置变更的角色到内存，后续该用户操作时，会强制该用户退出
				runtime.RuntimeConfig.GetCacheAdapter().Set(
					jwtauth.JwtRolePrefix,
					strconv.FormatInt(c.Id, 10),
					role.RoleKey,
					config.AuthConfig.MaxRefresh,
				)
			}

		}
		return true, clang.SuccessCode, nil
	}
	return false, clang.SuccessCode, nil
}

// UpdateStatus admin-更新系统用户状态
func (e *SysUser) UpdateStatus(c *dto.SysUserStatusUpdateReq, p *middleware.DataPermission) (bool, int, error) {
	if c.CurrUserId <= 0 || c.UserId <= 0 {
		return false, clang.ParamErrCode, lang.MsgErr(clang.ParamErrCode, e.Lang)
	}
	if c.Status == "" {
		return false, clang.SysUserStatusEmptyCode, lang.MsgErr(clang.SysUserStatusEmptyCode, e.Lang)
	}
	var err error
	u, respCode, err := e.Get(c.UserId, p)
	if err != nil {
		return false, respCode, err
	}

	updates := map[string]interface{}{}
	if c.Status != "" && u.Avatar != c.Status {
		updates["status"] = c.Status
	}

	if len(updates) > 0 {
		updates["update_by"] = c.CurrUserId
		updates["updated_at"] = time.Now()
		err = e.Orm.Model(&models.SysUser{}).Where("id=?", c.UserId).Updates(updates).Error
		if err != nil {
			return false, clang.DataUpdateLogCode, lang.MsgLogErrf(e.Log, e.Lang, clang.DataUpdateCode, clang.DataUpdateLogCode, err)
		}
		return true, clang.SuccessCode, nil
	}
	return false, clang.SuccessCode, nil
}

// ResetPwd admin-重置系统用户密码
func (e *SysUser) ResetPwd(c *dto.ResetSysUserPwdReq, p *middleware.DataPermission) (bool, int, error) {
	if c.CurrUserId <= 0 || c.UserId <= 0 {
		return false, clang.ParamErrCode, lang.MsgErr(clang.ParamErrCode, e.Lang)
	}

	var err error
	u, respCode, err := e.Get(c.UserId, p)
	if err != nil {
		return false, respCode, err
	}

	if u.Password != c.Password {
		now := time.Now()
		err = e.Orm.Where("id=?", c.UserId).Updates(&models.SysUser{
			Password:  c.Password,
			UpdatedAt: &now,
			UpdateBy:  c.CurrUserId,
		}).Error
		if err != nil {
			return false, clang.DataUpdateLogCode, lang.MsgLogErrf(e.Log, e.Lang, clang.DataUpdateCode, clang.DataUpdateLogCode, err)
		}
		return true, clang.SuccessCode, nil
	}
	return false, clang.SuccessCode, nil
}

// Delete admin-删除系统用户管理
func (e *SysUser) Delete(ids []int64, p *middleware.DataPermission) (int, error) {
	if len(ids) <= 0 {
		return clang.ParamErrCode, lang.MsgErr(clang.ParamErrCode, e.Lang)
	}

	//find if have admin account,not allow delete
	for _, id := range ids {
		u, respCode, err := e.Get(id, p)
		if err != nil {
			return respCode, err
		}
		if u.Username == global.RoleKeyAdmin {
			return clang.SysAdminUserNotAllowDeleteErrCode, lang.MsgErr(clang.SysAdminUserNotAllowDeleteErrCode, e.Lang)
		}
	}
	var err error
	var data models.SysUser
	err = e.Orm.Scopes(
		middleware.Permission(data.TableName(), p),
	).Delete(&data, ids).Error
	if err != nil {
		return clang.DataDeleteLogCode, lang.MsgLogErrf(e.Log, e.Lang, clang.DataDeleteCode, clang.DataDeleteLogCode, err)
	}
	return clang.SuccessCode, nil
}

// GetProfile admin-获取系统登录用户信息
func (e *SysUser) GetProfile(userId int64) (*dto.SysUserResp, int, error) {
	if userId <= 0 {
		return nil, clang.ParamErrCode, lang.MsgErr(clang.ParamErrCode, e.Lang)
	}
	user := &models.SysUser{}
	err := e.Orm.Preload("Dept").Preload("Post").Preload("Role").First(user, userId).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, clang.DataQueryLogCode, lang.MsgLogErrf(e.Log, e.Lang, clang.DataQueryCode, clang.DataQueryLogCode, err)
	}
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, clang.SysUserNoExistCode, lang.MsgErr(clang.SysUserNoExistCode, e.Lang)
	}

	if user.Role.RoleKey == "" {
		return nil, clang.SysUserNoRoleErrCode, lang.MsgErr(clang.SysUserNoRoleErrCode, e.Lang)
	}

	respUser := &dto.SysUserResp{}
	respUser.Id = user.Id
	respUser.Email = user.Email
	respUser.Phone = user.Phone
	respUser.Username = user.Username
	respUser.Avatar = user.Avatar
	respUser.CreatedAt = dateutils.ConvertToStrByPrt(user.CreatedAt, -1)
	respUser.Sex = user.Sex
	respUser.DeptName = user.Dept.DeptName
	respUser.RoleName = user.Role.RoleName

	if user.Role.RoleKey == global.RoleKeyAdmin {
		respUser.Permissions = []string{"*:*:*"}
	} else {
		roleService := NewSysRoleService(&e.Service)
		list, _, _ := roleService.GetPermissionsByRoleId(int64(user.RoleId))
		respUser.Permissions = list
	}
	respUser.RoleKyes = []string{user.Role.RoleKey}
	return respUser, clang.SuccessCode, nil
}

// UpdateProfile admin-更新系统登录用户信息
func (e *SysUser) UpdateProfile(c *dto.SysUserUpdateReq) (bool, int, error) {
	if c.Id <= 0 || c.CurrUserId <= 0 {
		return false, clang.ParamErrCode, lang.MsgErr(clang.ParamErrCode, e.Lang)
	}
	if c.Username == "" {
		return false, clang.SysUserNameEmptyCode, lang.MsgErr(clang.SysUserNameEmptyCode, e.Lang)
	}
	if c.Phone == "" {
		return false, clang.SysUserPhoneEmptyCode, lang.MsgErr(clang.SysUserPhoneEmptyCode, e.Lang)
	}
	if c.Email == "" {
		return false, clang.SysUserEmailEmptyCode, lang.MsgErr(clang.SysUserEmailEmptyCode, e.Lang)
	}

	data, respCode, err := e.Get(c.Id, nil)
	if err != nil {
		return false, respCode, err
	}

	updates := map[string]interface{}{}
	if c.Sex != "" && data.Sex != c.Sex {
		updates["sex"] = c.Sex
	}
	if c.Username != "" && data.Username != c.Username {
		req := dto.SysUserQueryReq{}
		req.Username = c.Username
		resp, respCode, err := e.QueryOne(&req, nil)
		if err != nil && respCode != clang.DataNotFoundCode {
			return false, respCode, err
		}
		if respCode == clang.SuccessCode && resp.Id != data.Id {
			return false, clang.SysUserNameExistCode, lang.MsgErr(clang.SysUserNameExistCode, e.Lang)
		}
		updates["username"] = c.Username
	}
	if c.Phone != "" && data.Phone != c.Phone {
		req := dto.SysUserQueryReq{}
		req.Phone = c.Phone
		resp, respCode, err := e.QueryOne(&req, nil)
		if err != nil && respCode != clang.DataNotFoundCode {
			return false, respCode, err
		}
		if respCode == clang.SuccessCode && resp.Id != data.Id {
			return false, clang.SysUserPhoneExistCode, lang.MsgErr(clang.SysUserPhoneExistCode, e.Lang)
		}
		updates["phone"] = c.Phone
	}
	if c.Email != "" && data.Email != c.Email {
		if !strutils.VerifyEmailFormat(c.Email) {
			return false, clang.SysUserEmailFormatErrCode, lang.MsgErr(clang.SysUserEmailFormatErrCode, e.Lang)
		}
		req := dto.SysUserQueryReq{}
		req.Email = c.Email
		resp, respCode, err := e.QueryOne(&req, nil)
		if err != nil && respCode != clang.DataNotFoundCode {
			return false, respCode, err
		}
		if respCode == clang.SuccessCode && resp.Id != data.Id {
			return false, clang.SysUserEmailExistCode, lang.MsgErr(clang.SysUserEmailExistCode, e.Lang)
		}
		updates["email"] = c.Email
	}
	if len(updates) > 0 {
		updates["update_by"] = c.CurrUserId
		updates["updated_at"] = time.Now()
		err = e.Orm.Model(&data).Where("id=?", data.Id).Updates(&updates).Error
		if err != nil {
			return false, clang.DataUpdateLogCode, lang.MsgLogErrf(e.Log, e.Lang, clang.DataUpdateCode, clang.DataUpdateLogCode, err)
		}
		return true, clang.SuccessCode, nil
	}
	return false, clang.SuccessCode, nil
}

// LoginVerify admin-登录验证
func (e *SysUser) LoginVerify(login *dto.LoginReq) (*models.SysUser, int, error) {
	user := &models.SysUser{}
	status := []string{global.SysStatusOk}
	if login.Username == global.RoleKeyAdmin {
		status = []string{global.SysStatusOk, global.SysStatusNotOk}
	}
	err := e.Orm.Preload("Dept").Preload("Post").Preload("Role").Where("username = ? and status in (?)", login.Username, status).First(user).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, clang.DataQueryLogCode, lang.MsgLogErrf(e.Log, e.Lang, clang.DataQueryCode, clang.DataQueryLogCode, err)
	}
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, clang.SysUserNoExistCode, lang.MsgErr(clang.SysUserNoExistCode, e.Lang)
	}
	if !strutils.CompareHashAndPassword(user.Password, login.Password) {
		return nil, clang.SysUserPwdErrCode, lang.MsgErr(clang.SysUserPwdErrCode, e.Lang)
	}
	return user, clang.SuccessCode, nil
}

// UpdateProfileAvatar admin-更新系统登录用户头像
func (e *SysUser) UpdateProfileAvatar(c *dto.SysUserAvatarUpdateReq, p *middleware.DataPermission) (bool, int, error) {
	if c.CurrUserId <= 0 {
		return false, clang.ParamErrCode, lang.MsgErr(clang.ParamErrCode, e.Lang)
	}
	var err error
	u, respCode, err := e.Get(c.CurrUserId, p)
	if err != nil {
		return false, respCode, err
	}

	updates := map[string]interface{}{}
	if c.Avatar != "" && u.Avatar != c.Avatar {
		updates["avatar"] = c.Avatar
	}

	if len(updates) > 0 {
		updates["update_by"] = c.CurrUserId
		updates["updated_at"] = time.Now()
		err = e.Orm.Model(&models.SysUser{}).Where("id=?", c.CurrUserId).Updates(updates).Error
		if err != nil {
			return false, clang.DataUpdateLogCode, lang.MsgLogErrf(e.Log, e.Lang, clang.DataUpdateCode, clang.DataUpdateLogCode, err)
		}
		return true, clang.SuccessCode, nil
	}
	return false, clang.SuccessCode, nil
}

// UpdateProfilePwd admin-更新系统登录用户密码
func (e *SysUser) UpdateProfilePwd(c dto.UpdateSysUserPwdReq, p *middleware.DataPermission) (bool, int, error) {
	if c.CurrUserId <= 0 {
		return false, clang.ParamErrCode, lang.MsgErr(clang.ParamErrCode, e.Lang)
	}
	if c.NewPassword == "" {
		return false, clang.SysUserNewPwdEmptyCode, lang.MsgErr(clang.SysUserNewPwdEmptyCode, e.Lang)
	}
	var err error
	u, respCode, err := e.Get(c.CurrUserId, p)
	if err != nil {
		return false, respCode, err
	}

	if !strutils.CompareHashAndPassword(u.Password, c.OldPassword) {
		return false, clang.SysUserPwdErrCode, lang.MsgErr(clang.SysUserPwdErrCode, e.Lang)
	}

	if !strutils.CompareHashAndPassword(u.Password, c.NewPassword) {
		now := time.Now()
		u.Password = c.NewPassword
		u.UpdateBy = c.CurrUserId
		u.UpdatedAt = &now
		err = e.Orm.Where("id=?", c.CurrUserId).Updates(&models.SysUser{
			Password:  c.NewPassword,
			UpdatedAt: &now,
			UpdateBy:  c.CurrUserId,
		}).Error
		if err != nil {
			return false, clang.DataUpdateLogCode, lang.MsgLogErrf(e.Log, e.Lang, clang.DataUpdateCode, clang.DataUpdateLogCode, err)
		}
		return true, clang.SuccessCode, nil
	}
	return false, clang.SuccessCode, nil
}

// LoginLogToDB admin-登录日志记录到数据库
func (e *SysUser) LoginLogToDB(c *gin.Context, status string, msg string, userId int64) {
	if !config.LoggerConfig.EnabledDB {
		return
	}
	l := make(map[string]interface{})

	ua := user_agent.New(c.Request.UserAgent())
	l["ipaddr"] = iputils.GetClientIP(c)
	//用于定位ip所在城市
	l["loginLocation"] = iputils.GetLocation(iputils.GetClientIP(c), config.ApplicationConfig.AmpKey)
	l["loginTime"] = strutils.GetCurrentTime()
	l["status"] = status
	l["agent"] = c.Request.UserAgent()
	browserName, browserVersion := ua.Browser()
	l["browser"] = browserName + " " + browserVersion
	l["os"] = ua.OS()
	l["platform"] = ua.Platform()
	l["userId"] = userId
	l["remark"] = msg

	q := runtime.RuntimeConfig.GetMemoryQueue(c.Request.Host)
	message, err := runtime.RuntimeConfig.GetStreamMessage("", global.LoginLog, l)
	if err != nil {
		e.Log.Errorf("SysUserService LoginLogToDB error:%s", err)
		//日志报错错误，不中断请求
	} else {
		err = q.Append(message)
		if e.Log != nil {
			e.Log.Errorf("SysUserService LoginLogToDB error:%s", err)
		}
	}
}
