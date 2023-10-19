package service

import (
	"github.com/gin-gonic/gin"
	"github.com/mssola/user_agent"
	sysLang "go-admin/app/admin/lang"
	"go-admin/app/admin/models"
	"go-admin/app/admin/service/dto"
	"go-admin/common/core"
	"go-admin/common/core/pkg"
	"go-admin/common/core/service"
	"go-admin/common/global"
	"go-admin/common/middleware"
	"go-admin/common/utils/dateUtils"
	"go-admin/common/utils/strutils"
	"go-admin/config/config"
	"go-admin/config/lang"
	"gorm.io/gorm"
	"time"

	cDto "go-admin/common/dto"
)

type SysUser struct {
	service.Service
}

func NewSysUserService(s *service.Service) *SysUser {
	var srv = new(SysUser)
	srv.Orm = s.Orm
	srv.Log = s.Log
	return srv
}

// GetPage 获取SysUser列表
func (e *SysUser) GetPage(c *dto.SysUserQueryReq, p *middleware.DataPermission) ([]models.SysUser, int64, int, error) {
	var list []models.SysUser
	var data models.SysUser
	var count int64

	err := e.Orm.Model(&data).Preload("Dept").Preload("Role").Order("created_at desc").
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
			middleware.Permission(data.TableName(), p),
		).Find(&list).Limit(-1).Offset(-1).Count(&count).Error
	if err != nil {
		return nil, 0, lang.DataQueryLogCode, lang.MsgLogErrf(e.Log, e.Lang, lang.DataQueryCode, lang.DataQueryLogCode, err)
	}
	return list, count, lang.SuccessCode, nil
}

// Get 获取SysUser对象
func (e *SysUser) Get(id int64, p *middleware.DataPermission) (*models.SysUser, int, error) {
	if id <= 0 {
		return nil, lang.ParamErrCode, lang.MsgErr(lang.ParamErrCode, e.Lang)
	}
	data := &models.SysUser{}
	err := e.Orm.Scopes(
		middleware.Permission(data.TableName(), p),
	).First(data, id).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, lang.DataQueryLogCode, lang.MsgLogErrf(e.Log, e.Lang, lang.DataQueryCode, lang.DataQueryLogCode, err)
	}
	if err == gorm.ErrRecordNotFound {
		return nil, lang.DataNotFoundCode, lang.MsgErr(lang.DataNotFoundCode, e.Lang)
	}
	return data, lang.SuccessCode, nil
}

// QueryOne 通过自定义条件获取一条记录
func (e *SysUser) QueryOne(queryCondition *dto.SysUserQueryReq, p *middleware.DataPermission) (*models.SysUser, int, error) {
	data := &models.SysUser{}
	err := e.Orm.Model(&models.SysUser{}).
		Scopes(
			cDto.MakeCondition(queryCondition.GetNeedSearch()),
			middleware.Permission(data.TableName(), p),
		).First(data).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, lang.DataQueryLogCode, lang.MsgLogErrf(e.Log, e.Lang, lang.DataQueryCode, lang.DataQueryLogCode, err)
	}
	if err == gorm.ErrRecordNotFound {
		return nil, lang.DataNotFoundCode, lang.MsgErr(lang.DataNotFoundCode, e.Lang)
	}
	return data, lang.SuccessCode, nil
}

// Count 获取条数
func (e *SysUser) Count(c *dto.SysUserQueryReq) (int64, int, error) {
	var err error
	var count int64
	err = e.Orm.Model(&models.SysUser{}).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
		).Limit(-1).Offset(-1).
		Count(&count).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return 0, lang.DataQueryLogCode, lang.MsgLogErrf(e.Log, e.Lang, lang.DataQueryCode, lang.DataQueryLogCode, err)
	}
	if err == gorm.ErrRecordNotFound {
		return 0, lang.DataNotFoundCode, lang.MsgErr(lang.DataNotFoundCode, e.Lang)
	}
	return count, lang.SuccessCode, nil
}

// Insert 创建SysUser对象
func (e *SysUser) Insert(c *dto.SysUserInsertReq) (int64, int, error) {
	if c.CurrUserId <= 0 {
		return 0, lang.ParamErrCode, lang.MsgErr(lang.ParamErrCode, e.Lang)
	}
	if c.Username == "" {
		return 0, sysLang.SysUserNameEmptyCode, lang.MsgErr(sysLang.SysUserNameEmptyCode, e.Lang)
	}
	if c.NickName == "" {
		return 0, sysLang.SysNickNameEmptyCode, lang.MsgErr(sysLang.SysNickNameEmptyCode, e.Lang)
	}
	if c.Phone == "" {
		return 0, sysLang.SysUserPhoneEmptyCode, lang.MsgErr(sysLang.SysUserPhoneEmptyCode, e.Lang)
	}
	if c.Email == "" {
		return 0, sysLang.SysUserEmailEmptyCode, lang.MsgErr(sysLang.SysUserEmailEmptyCode, e.Lang)
	}
	if c.DeptId <= 0 {
		return 0, sysLang.SysUserDeptEmptyCode, lang.MsgErr(sysLang.SysUserDeptEmptyCode, e.Lang)
	}
	if c.Password == "" {
		return 0, sysLang.SysUserPwdEmptyCode, lang.MsgErr(sysLang.SysUserPwdEmptyCode, e.Lang)
	}

	if c.Username != "" {
		query := dto.SysUserQueryReq{}
		query.Username = c.Username
		count, respCode, err := e.Count(&query)
		if err != nil && respCode != lang.DataNotFoundCode {
			return 0, respCode, err
		}
		if count > 0 {
			return 0, sysLang.SysUserNameExistCode, lang.MsgErr(sysLang.SysUserNameExistCode, e.Lang)
		}
	}
	if c.NickName != "" {
		query := dto.SysUserQueryReq{}
		query.NickName = c.NickName
		count, respCode, err := e.Count(&query)
		if err != nil && respCode != lang.DataNotFoundCode {
			return 0, respCode, err
		}
		if count > 0 {
			return 0, sysLang.SysUserNickNameExistCode, lang.MsgErr(sysLang.SysUserNickNameExistCode, e.Lang)
		}
	}
	if c.Phone != "" {
		query := dto.SysUserQueryReq{}
		query.Phone = c.Phone
		count, respCode, err := e.Count(&query)
		if err != nil && respCode != lang.DataNotFoundCode {
			return 0, respCode, err
		}
		if count > 0 {
			return 0, sysLang.SysUserPhoneExistCode, lang.MsgErr(sysLang.SysUserPhoneExistCode, e.Lang)
		}
	}
	if c.Email != "" {
		query := dto.SysUserQueryReq{}
		query.Email = c.Email
		count, respCode, err := e.Count(&query)
		if err != nil && respCode != lang.DataNotFoundCode {
			return 0, respCode, err
		}
		if count > 0 {
			return 0, sysLang.SysUserEmailExistCode, lang.MsgErr(sysLang.SysUserEmailExistCode, e.Lang)
		}
	}

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
		return 0, lang.DataInsertLogCode, lang.MsgLogErrf(e.Log, e.Lang, lang.DataInsertCode, lang.DataInsertLogCode, err)
	}
	return data.Id, lang.SuccessCode, nil
}

// Update 修改SysUser对象
func (e *SysUser) Update(c *dto.SysUserUpdateReq, p *middleware.DataPermission) (bool, int, error) {
	if c.Id <= 0 || c.CurrUserId <= 0 {
		return false, lang.ParamErrCode, lang.MsgErr(lang.ParamErrCode, e.Lang)
	}
	if c.Username == "" {
		return false, sysLang.SysUserNameEmptyCode, lang.MsgErr(sysLang.SysUserNameEmptyCode, e.Lang)
	}
	//if c.NickName == "" {
	//	return false, sysLang.SysNickNameEmptyCode, lang.MsgErr(sysLang.SysNickNameEmptyCode, e.Lang)
	//}
	if c.Phone == "" {
		return false, sysLang.SysUserPhoneEmptyCode, lang.MsgErr(sysLang.SysUserPhoneEmptyCode, e.Lang)
	}
	if c.Email == "" {
		return false, sysLang.SysUserEmailEmptyCode, lang.MsgErr(sysLang.SysUserEmailEmptyCode, e.Lang)
	}
	/*	if c.DeptId <= 0 {
		return false, sysLang.SysUserDeptEmptyCode, lang.MsgErr(sysLang.SysUserDeptEmptyCode, e.Lang)
	}*/

	data, respCode, err := e.Get(c.Id, p)
	if err != nil {
		return false, respCode, err
	}

	updates := map[string]interface{}{}
	if c.Username != "" && data.Username != c.Username {
		req := dto.SysUserQueryReq{}
		req.Username = c.Username
		resp, respCode, err := e.QueryOne(&req, p)
		if err != nil && respCode != lang.DataNotFoundCode {
			return false, respCode, err
		}
		if respCode == lang.SuccessCode && resp.Id != data.Id {
			return false, sysLang.SysUserNameExistCode, lang.MsgErr(sysLang.SysUserNameExistCode, e.Lang)
		}
		updates["username"] = c.Username
	}
	if c.NickName != "" && data.NickName != c.NickName {
		req := dto.SysUserQueryReq{}
		req.NickName = c.NickName
		resp, respCode, err := e.QueryOne(&req, p)
		if err != nil && respCode != lang.DataNotFoundCode {
			return false, respCode, err
		}
		if respCode == lang.SuccessCode && resp.Id != data.Id {
			return false, sysLang.SysUserNickNameExistCode, lang.MsgErr(sysLang.SysUserNickNameExistCode, e.Lang)
		}
		updates["nick_name"] = c.NickName
	}
	if c.Phone != "" && data.Phone != c.Phone {
		req := dto.SysUserQueryReq{}
		req.Phone = c.Phone
		resp, respCode, err := e.QueryOne(&req, p)
		if err != nil && respCode != lang.DataNotFoundCode {
			return false, respCode, err
		}
		if respCode == lang.SuccessCode && resp.Id != data.Id {
			return false, sysLang.SysUserPhoneExistCode, lang.MsgErr(sysLang.SysUserPhoneExistCode, e.Lang)
		}
		updates["phone"] = c.Phone
	}
	if c.RoleId > 0 && data.RoleId != c.RoleId {
		updates["role_id"] = c.RoleId
	}
	if c.Avatar != "" && data.Avatar != c.Avatar {
		updates["avatar"] = c.Avatar
	}
	if c.Sex != "" && data.Sex != c.Sex {
		updates["sex"] = c.Sex
	}
	if c.Email != "" && data.Email != c.Email {
		if !strutils.VerifyEmailFormat(c.Email) {
			return false, sysLang.SysUserEmailFormatErrCode, lang.MsgErr(sysLang.SysUserEmailFormatErrCode, e.Lang)
		}
		req := dto.SysUserQueryReq{}
		req.Email = c.Email
		resp, respCode, err := e.QueryOne(&req, p)
		if err != nil && respCode != lang.DataNotFoundCode {
			return false, respCode, err
		}
		if respCode == lang.SuccessCode && resp.Id != data.Id {
			return false, sysLang.SysUserEmailExistCode, lang.MsgErr(sysLang.SysUserEmailExistCode, e.Lang)
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
			return false, lang.DataUpdateLogCode, lang.MsgLogErrf(e.Log, e.Lang, lang.DataUpdateCode, lang.DataUpdateLogCode, err)
		}
		return true, lang.SuccessCode, nil
	}
	return false, lang.SuccessCode, nil
}

// UpdateStatus 更新用户状态
func (e *SysUser) UpdateStatus(c *dto.SysUserStatusUpdateReq, p *middleware.DataPermission) (bool, int, error) {
	if c.CurrUserId <= 0 || c.UserId <= 0 {
		return false, lang.ParamErrCode, lang.MsgErr(lang.ParamErrCode, e.Lang)
	}
	if c.Status == "" {
		return false, sysLang.SysUserStatusEmptyCode, lang.MsgErr(sysLang.SysUserStatusEmptyCode, e.Lang)
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
			return false, lang.DataUpdateLogCode, lang.MsgLogErrf(e.Log, e.Lang, lang.DataUpdateCode, lang.DataUpdateLogCode, err)
		}
		return true, lang.SuccessCode, nil
	}
	return false, lang.SuccessCode, nil
}

// ResetPwd 重置用户密码
func (e *SysUser) ResetPwd(c *dto.ResetSysUserPwdReq, p *middleware.DataPermission) (bool, int, error) {
	if c.CurrUserId <= 0 || c.UserId <= 0 {
		return false, lang.ParamErrCode, lang.MsgErr(lang.ParamErrCode, e.Lang)
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
			return false, lang.DataUpdateLogCode, lang.MsgLogErrf(e.Log, e.Lang, lang.DataUpdateCode, lang.DataUpdateLogCode, err)
		}
		return true, lang.SuccessCode, nil
	}
	return false, lang.SuccessCode, nil
}

// Remove 删除SysUser
func (e *SysUser) Remove(ids []int64, p *middleware.DataPermission) (int, error) {
	if len(ids) <= 0 {
		return lang.ParamErrCode, lang.MsgErr(lang.ParamErrCode, e.Lang)
	}
	var err error
	var data models.SysUser
	err = e.Orm.Scopes(
		middleware.Permission(data.TableName(), p),
	).Delete(&data, ids).Error
	if err != nil {
		return lang.DataDeleteLogCode, lang.MsgLogErrf(e.Log, e.Lang, lang.DataDeleteCode, lang.DataDeleteLogCode, err)
	}
	return lang.SuccessCode, nil
}

func (e *SysUser) GetProfile(userId int64) (*dto.SysUserResp, int, error) {
	if userId <= 0 {
		return nil, lang.ParamErrCode, lang.MsgErr(lang.ParamErrCode, e.Lang)
	}
	user := &models.SysUser{}
	err := e.Orm.Preload("Dept").Preload("Post").Preload("Role").First(user, userId).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, lang.DataQueryLogCode, lang.MsgLogErrf(e.Log, e.Lang, lang.DataQueryCode, lang.DataQueryLogCode, err)
	}
	if err != nil && err == gorm.ErrRecordNotFound {
		return nil, sysLang.SysUserNoExistCode, lang.MsgErr(sysLang.SysUserNoExistCode, e.Lang)
	}
	respUser := &dto.SysUserResp{}
	respUser.Id = user.Id
	respUser.Email = user.Email
	//respUser.PostId = user.PostId
	//respUser.DeptId = user.DeptId
	respUser.Phone = user.Phone
	respUser.Username = user.Username
	respUser.CreatedAt = dateUtils.ConvertToStrByPrt(user.CreatedAt, -1)
	respUser.Sex = user.Sex
	respUser.Dept = *user.Dept
	respUser.Role = *user.Role
	return respUser, lang.SuccessCode, nil
}

/*func (e *SysUser) GetProfile(c *dto.SysUserById, user *models.SysUser, roles *[]models.SysRole, posts *[]models.SysPost) error {
	err := e.Orm.Preload("Dept").First(user, c.GetId()).Error
	if err != nil {
		return err
	}
	err = e.Orm.Find(roles, user.RoleId).Error
	if err != nil {
		return err
	}
	err = e.Orm.Find(posts, user.PostIds).Error
	if err != nil {
		return err
	}

	return nil
}*/

func (e *SysUser) GetUser(login *dto.LoginReq) (*models.SysUser, int, error) {
	user := &models.SysUser{}
	err := e.Orm.Preload("Dept").Preload("Post").Preload("Role").Where("username = ?  and status = ?", login.Username, global.SysStatusOk).First(user).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, lang.DataQueryLogCode, lang.MsgLogErrf(e.Log, e.Lang, lang.DataQueryCode, lang.DataQueryLogCode, err)
	}
	if err != nil && err == gorm.ErrRecordNotFound {
		return nil, sysLang.SysUserNoExistCode, lang.MsgErr(sysLang.SysUserNoExistCode, e.Lang)
	}
	if !strutils.CompareHashAndPassword(user.Password, login.Password) {
		return nil, sysLang.SysUserPwdErrCode, lang.MsgErr(sysLang.SysUserPwdErrCode, e.Lang)
	}
	return user, lang.SuccessCode, nil
}

// UpdateSelfPhone 修改手机号
func (e *SysUser) UpdateSelfPhone(c *dto.SysUserPhoneUpdateReq) (bool, int, error) {
	if c.CurrUserId <= 0 {
		return false, lang.ParamErrCode, lang.MsgErr(lang.ParamErrCode, e.Lang)
	}
	var err error
	u, respCode, err := e.Get(c.CurrUserId, nil)
	if err != nil {
		return false, respCode, err
	}
	updates := map[string]interface{}{}
	if c.Phone != "" && u.Phone != c.Phone {
		req := dto.SysUserQueryReq{}
		req.Phone = c.Phone
		resp, respCode, err := e.QueryOne(&req, nil)
		if err != nil && respCode != lang.DataNotFoundCode {
			return false, respCode, err
		}
		if respCode == lang.SuccessCode && resp.Id != c.CurrUserId {
			return false, sysLang.SysUserPhoneExistCode, lang.MsgErr(sysLang.SysUserPhoneExistCode, e.Lang)
		}
		updates["phone"] = c.Phone
	}

	if len(updates) > 0 {
		updates["update_by"] = c.CurrUserId
		updates["updated_at"] = time.Now()
		err = e.Orm.Model(&models.SysUser{}).Where("id=?", c.CurrUserId).Updates(updates).Error
		if err != nil {
			return false, lang.DataUpdateLogCode, lang.MsgLogErrf(e.Log, e.Lang, lang.DataUpdateCode, lang.DataUpdateLogCode, err)
		}
		return true, lang.SuccessCode, nil
	}
	return false, lang.SuccessCode, nil
}

// UpdateSelfNickName 更新昵称
func (e *SysUser) UpdateSelfNickName(c *dto.SysUserNickNameUpdateReq) (bool, int, error) {
	if c.CurrUserId <= 0 {
		return false, lang.ParamErrCode, lang.MsgErr(lang.ParamErrCode, e.Lang)
	}
	if c.NickName == "" {
		return false, sysLang.SysNickNameEmptyCode, lang.MsgErr(sysLang.SysNickNameEmptyCode, e.Lang)
	}
	var err error
	u, respCode, err := e.Get(c.CurrUserId, nil)
	if err != nil {
		return false, respCode, err
	}

	updates := map[string]interface{}{}
	if c.NickName != "" && u.NickName != c.NickName {
		req := dto.SysUserQueryReq{}
		req.NickName = c.NickName
		resp, respCode, err := e.QueryOne(&req, nil)
		if err != nil && respCode != lang.DataNotFoundCode {
			return false, respCode, err
		}
		if respCode == lang.SuccessCode && resp.Id != c.CurrUserId {
			return false, sysLang.SysUserPhoneExistCode, lang.MsgErr(sysLang.SysUserPhoneExistCode, e.Lang)
		}
		updates["nick_name"] = c.NickName
	}

	if len(updates) > 0 {
		updates["update_by"] = c.CurrUserId
		updates["updated_at"] = time.Now()
		err = e.Orm.Model(&models.SysUser{}).Where("id=?", c.CurrUserId).Updates(updates).Error
		if err != nil {
			return false, lang.DataUpdateLogCode, lang.MsgLogErrf(e.Log, e.Lang, lang.DataUpdateCode, lang.DataUpdateLogCode, err)
		}
		return true, lang.SuccessCode, nil
	}
	return false, lang.SuccessCode, nil
}

// UpdateSelfEmail 修改邮箱号
func (e *SysUser) UpdateSelfEmail(c *dto.SysUserUpdateEmailReq) (bool, int, error) {
	if c.CurrUserId <= 0 {
		return false, lang.ParamErrCode, lang.MsgErr(lang.ParamErrCode, e.Lang)
	}
	if !strutils.VerifyEmailFormat(c.Email) {
		return false, sysLang.SysUserEmailFormatErrCode, lang.MsgErr(sysLang.SysUserEmailFormatErrCode, e.Lang)
	}
	var err error
	u, respCode, err := e.Get(c.CurrUserId, nil)
	if err != nil {
		return false, respCode, err
	}

	updates := map[string]interface{}{}
	if c.Email != "" && u.Email != c.Email {
		req := dto.SysUserQueryReq{}
		req.Email = c.Email
		resp, respCode, err := e.QueryOne(&req, nil)
		if err != nil && respCode != lang.DataNotFoundCode {
			return false, respCode, err
		}
		if respCode == lang.SuccessCode && resp.Id != c.CurrUserId {
			return false, sysLang.SysUserPhoneExistCode, lang.MsgErr(sysLang.SysUserPhoneExistCode, e.Lang)
		}
		updates["email"] = c.Email
	}

	if len(updates) > 0 {
		updates["update_by"] = c.CurrUserId
		updates["updated_at"] = time.Now()
		err = e.Orm.Model(&models.SysUser{}).Where("id=?", c.CurrUserId).Updates(updates).Error
		if err != nil {
			return false, lang.DataUpdateLogCode, lang.MsgLogErrf(e.Log, e.Lang, lang.DataUpdateCode, lang.DataUpdateLogCode, err)
		}
		return true, lang.SuccessCode, nil
	}
	return false, lang.SuccessCode, nil
}

// UpdateAvatar 更新用户头像
func (e *SysUser) UpdateAvatar(c *dto.SysUserAvatarUpdateReq, p *middleware.DataPermission) (bool, int, error) {
	if c.CurrUserId <= 0 {
		return false, lang.ParamErrCode, lang.MsgErr(lang.ParamErrCode, e.Lang)
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
			return false, lang.DataUpdateLogCode, lang.MsgLogErrf(e.Log, e.Lang, lang.DataUpdateCode, lang.DataUpdateLogCode, err)
		}
		return true, lang.SuccessCode, nil
	}
	return false, lang.SuccessCode, nil
}

// UpdatePwd 修改SysUser对象密码
func (e *SysUser) UpdatePwd(c dto.UpdateSysUserPwdReq, p *middleware.DataPermission) (bool, int, error) {
	if c.CurrUserId <= 0 {
		return false, lang.ParamErrCode, lang.MsgErr(lang.ParamErrCode, e.Lang)
	}
	if c.NewPassword == "" {
		return false, sysLang.SysUserNewPwdEmptyCode, lang.MsgErr(sysLang.SysUserNewPwdEmptyCode, e.Lang)
	}
	var err error
	u, respCode, err := e.Get(c.CurrUserId, p)
	if err != nil {
		return false, respCode, err
	}

	if !strutils.CompareHashAndPassword(u.Password, c.OldPassword) {
		return false, sysLang.SysUserPwdErrCode, lang.MsgErr(sysLang.SysUserPwdErrCode, e.Lang)
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
			return false, lang.DataUpdateLogCode, lang.MsgLogErrf(e.Log, e.Lang, lang.DataUpdateCode, lang.DataUpdateLogCode, err)
		}
		return true, lang.SuccessCode, nil
	}
	return false, lang.SuccessCode, nil
}

// LoginLogToDB Write log to database
func (e *SysUser) LoginLogToDB(c *gin.Context, status string, msg string, userId int64) {
	if !config.LoggerConfig.EnabledDB {
		return
	}
	l := make(map[string]interface{})

	ua := user_agent.New(c.Request.UserAgent())
	l["ipaddr"] = pkg.GetClientIP(c)
	//用于定位ip所在城市
	//fmt.Println("gaConfig.ExtConfig.AMap.Key", config.ApplicationConfig.AmpKey)
	l["loginLocation"] = pkg.GetLocation(pkg.GetClientIP(c), config.ApplicationConfig.AmpKey)
	l["loginTime"] = pkg.GetCurrentTime()
	l["status"] = status
	l["agent"] = c.Request.UserAgent()
	browserName, browserVersion := ua.Browser()
	l["browser"] = browserName + " " + browserVersion
	l["os"] = ua.OS()
	l["platform"] = ua.Platform()
	l["userId"] = userId
	l["remark"] = msg

	q := core.Runtime.GetMemoryQueue(c.Request.Host)
	message, err := core.Runtime.GetStreamMessage("", global.LoginLog, l)
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
