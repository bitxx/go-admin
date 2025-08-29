package service

import (
	"errors"
	"fmt"
	"github.com/xuri/excelize/v2"
	adminService "go-admin/app/admin/sys/service"
	"go-admin/app/app/user/models"
	"go-admin/app/app/user/service/dto"
	"go-admin/config/base/constant"
	baseLang "go-admin/config/base/lang"
	"go-admin/core/config"
	cDto "go-admin/core/dto"
	"go-admin/core/dto/service"
	"go-admin/core/global"
	"go-admin/core/lang"
	"go-admin/core/middleware"
	"go-admin/core/utils/encrypt"
	"go-admin/core/utils/idgen"
	"go-admin/core/utils/strutils"
	"strconv"
	"strings"

	"gorm.io/gorm"
	"time"
)

type User struct {
	service.Service
}

// NewUserService app-实例化用户管理
func NewUserService(s *service.Service) *User {
	var srv = new(User)
	srv.Orm = s.Orm
	srv.Log = s.Log
	return srv
}

// GetPage app-获取用户管理分页列表
func (e *User) GetPage(c *dto.UserQueryReq, p *middleware.DataPermission) ([]models.User, int64, int, error) {
	var data models.User
	var list []models.User
	var count int64
	if c.Mobile != "" {
		c.Mobile, _ = encrypt.AesEncrypt(c.Mobile, []byte(config.AuthConfig.Secret))
	}

	if c.Email != "" {
		c.Email, _ = encrypt.AesEncrypt(c.Email, []byte(config.AuthConfig.Secret))
	}

	//上级邀请码查询
	if c.ParentRefCode != "" {
		m := &models.User{}
		queryUserCondition := &dto.UserQueryReq{}
		queryUserCondition.RefCode = c.ParentRefCode
		m, respCode, err := e.QueryOne(queryUserCondition, p)
		if err != nil {
			return nil, 0, respCode, err
		}
		c.ParentId = m.ParentId
	}

	err := e.Orm.Preload("UserLevel").Order("created_at desc").Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
			middleware.Permission(data.TableName(), p),
		).Find(&list).Limit(-1).Offset(-1).Count(&count).Error
	if err != nil {
		return nil, 0, baseLang.DataQueryLogCode, lang.MsgLogErrf(e.Log, e.Lang, baseLang.DataQueryCode, baseLang.DataQueryLogCode, err)
	}

	//只做局部缓存
	userCache := map[int64]*models.User{}

	for index, user := range list {
		parentId := user.ParentId
		if parentId > 0 {
			//获取邀请人信息
			parentUser := userCache[parentId]
			if parentUser == nil {
				var err error
				parentUser, _, err = e.Get(parentId, p)
				if err == nil {
					userCache[parentId] = parentUser
				}
			}
			if parentUser != nil {
				list[index].ParentRefCode = parentUser.RefCode
				list[index].ParentUserName = parentUser.UserName
			}

		}

		mobile, err := encrypt.AesDecrypt(user.Mobile, []byte(config.AuthConfig.Secret))
		if err == nil {
			list[index].Mobile = strutils.HidePartStr(mobile, 3)
		}

		email, err := encrypt.AesDecrypt(user.Email, []byte(config.AuthConfig.Secret))
		if err == nil {
			list[index].Email = strutils.HidePartStr(email, 5)
		}
	}
	return list, count, baseLang.SuccessCode, nil
}

// Get app-获取用户管理详情
func (e *User) Get(id int64, p *middleware.DataPermission) (*models.User, int, error) {
	if id <= 0 {
		return nil, baseLang.ParamErrCode, lang.MsgErr(baseLang.ParamErrCode, e.Lang)
	}
	data := &models.User{}
	err := e.Orm.Scopes(
		middleware.Permission(data.TableName(), p),
	).First(data, id).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, baseLang.DataQueryLogCode, lang.MsgLogErrf(e.Log, e.Lang, baseLang.DataQueryCode, baseLang.DataQueryLogCode, err)
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, baseLang.DataNotFoundCode, lang.MsgErr(baseLang.DataNotFoundCode, e.Lang)
	}
	if data.Mobile != "" {
		data.Mobile, _ = encrypt.AesDecrypt(data.Mobile, []byte(config.AuthConfig.Secret))
	}
	if data.Email != "" {
		data.Email, _ = encrypt.AesDecrypt(data.Email, []byte(config.AuthConfig.Secret))
	}
	return data, baseLang.SuccessCode, nil
}

// QueryOne app-获取用户管理一条记录
func (e *User) QueryOne(queryCondition *dto.UserQueryReq, p *middleware.DataPermission) (*models.User, int, error) {
	data := &models.User{}
	err := e.Orm.Scopes(
		cDto.MakeCondition(queryCondition.GetNeedSearch()),
		middleware.Permission(data.TableName(), p),
	).First(data).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, baseLang.DataQueryLogCode, lang.MsgLogErrf(e.Log, e.Lang, baseLang.DataQueryCode, baseLang.DataQueryLogCode, err)
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, baseLang.DataNotFoundCode, lang.MsgErr(baseLang.DataNotFoundCode, e.Lang)
	}
	return data, baseLang.SuccessCode, nil
}

// QueryOne app-获取用户在树同一个层最大排序号
func (e *User) queryMaxTreeSort(queryCondition *dto.UserQueryReq) (int64, int, error) {
	maxSort := int64(0)
	err := e.Orm.Scopes(
		cDto.MakeCondition(queryCondition.GetNeedSearch()),
	).Model(&models.User{}).Select("max(tree_sort)").Scan(&maxSort).Error
	if err != nil {
		maxSort = 0
	}
	return maxSort, baseLang.SuccessCode, nil
}

// Count admin-获取用户管理数据总数
func (e *User) Count(queryCondition *dto.UserQueryReq) (int64, int, error) {
	var err error
	var count int64
	err = e.Orm.Model(&models.User{}).
		Scopes(
			cDto.MakeCondition(queryCondition.GetNeedSearch()),
		).Limit(-1).Offset(-1).Count(&count).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return 0, baseLang.DataQueryLogCode, lang.MsgLogErrf(e.Log, e.Lang, baseLang.DataQueryCode, baseLang.DataQueryLogCode, err)
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return 0, baseLang.DataNotFoundCode, lang.MsgErr(baseLang.DataNotFoundCode, e.Lang)
	}
	return count, baseLang.SuccessCode, nil
}

// Insert app-新增用户管理
func (e *User) Insert(c *dto.UserInsertReq) (int, error) {
	if c.CurrUserId <= 0 {
		return baseLang.ParamErrCode, lang.MsgErr(baseLang.ParamErrCode, e.Lang)
	}
	if c.Emails == "" && c.Mobiles == "" {
		return baseLang.UserEmailOrMobileNeedCode, lang.MsgErr(baseLang.UserEmailOrMobileNeedCode, e.Lang)
	}

	var mobiles, emails []string
	if c.Emails != "" {
		emails = strings.Split(c.Emails, ",")
		for _, email := range emails {
			if !strutils.IsEmail(email) {
				return baseLang.UserEmailFormatErrCode, lang.MsgErr(baseLang.UserEmailFormatErrCode, e.Lang)
			}
		}
	}

	if c.Mobiles != "" {
		if c.MobileTitle == "" {
			return baseLang.UserMobileNeedTitleCode, lang.MsgErr(baseLang.UserMobileNeedTitleCode, e.Lang)
		}
		mobiles = strings.Split(c.Mobiles, ",")
		for _, mobile := range mobiles {
			if !strutils.IsMobile(mobile) {
				return baseLang.UserMobileFormatErrCode, lang.MsgErr(baseLang.UserMobileFormatErrCode, e.Lang)
			}
		}
	}
	var refUser *models.User
	if c.RefCode != "" {
		queryUserCondition := dto.UserQueryReq{}
		queryUserCondition.RefCode = c.RefCode
		queryUserCondition.Status = global.SysStatusOk
		u, respCode, err := e.QueryOne(&queryUserCondition, nil)
		if err != nil && respCode != baseLang.DataNotFoundCode {
			return respCode, err
		}
		if respCode == baseLang.SuccessCode {
			refUser = u
		}
	}

	//事务 开启-关闭
	var err error
	respCode := baseLang.SuccessCode
	e.Orm = e.Orm.Begin()
	defer func() {
		if err != nil {
			e.Orm.Rollback()
		} else {
			e.Orm.Commit()
		}
	}()

	for _, mobile := range mobiles {
		mobile, err = encrypt.AesEncrypt(mobile, []byte(config.AuthConfig.Secret))
		if err != nil {
			return baseLang.UserMobileEncryptErrLogCode, lang.MsgLogErrf(e.Log, e.Lang, baseLang.UserMobileEncryptErrCode, baseLang.UserMobileEncryptErrLogCode, err)
		}
		respCode, err = e.register(constant.AccountMobileType, "", mobile, c.MobileTitle, refUser)
		if err != nil {
			return respCode, err
		}
	}
	for _, email := range emails {
		email, err = encrypt.AesEncrypt(email, []byte(config.AuthConfig.Secret))
		if err != nil {
			return baseLang.UserEmailEncryptErrLogCode, lang.MsgLogErrf(e.Log, e.Lang, baseLang.UserEmailEncryptErrCode, baseLang.UserEmailEncryptErrLogCode, err)
		}
		respCode, err = e.register(constant.AccountEmailType, email, "", "", refUser)
		if err != nil {
			return respCode, err
		}
	}

	return baseLang.SuccessCode, nil
}

// register app-内部方法，注册用户管理
func (e *User) register(registerType, email, mobile, mobileTitle string, refUser *models.User) (int, error) {

	//验证手机号和邮箱是否重复
	queryUserCondition := dto.UserQueryReq{}
	if registerType == constant.AccountMobileType {
		if (mobile != "" && mobileTitle == "") || (mobile == "" && mobileTitle != "") || (mobile == "" && mobileTitle == "") {
			return baseLang.UserMobileNeedTitleCode, lang.MsgErr(baseLang.UserMobileNeedTitleCode, e.Lang)
		}
		queryUserCondition.Mobile = mobile
		queryUserCondition.MobileTitle = mobileTitle
	}
	if registerType == constant.AccountEmailType {
		queryUserCondition.Email = email
	}

	count, respCode, err := e.Count(&queryUserCondition)
	if err != nil && respCode != baseLang.DataNotFoundCode {
		return respCode, err
	}
	if count > 0 {
		account := ""
		if email != "" {
			account, _ = encrypt.AesDecrypt(email, []byte(config.AuthConfig.Secret))
		} else if mobile != "" {
			account, _ = encrypt.AesDecrypt(mobile, []byte(config.AuthConfig.Secret))
		}
		return baseLang.UserAccountExistLogCode, lang.MsgErrf(baseLang.UserAccountExistLogCode, e.Lang, account)
	}

	return e.insertRegisterInfo(registerType, email, mobile, mobileTitle, refUser)
}

// insertRegisterInfo app-内部方法，注册用户管理
func (e *User) insertRegisterInfo(registerType, emial, mobile, mobileTitle string, refCode *models.User) (int, error) {

	//插入用户数据
	uid, respCode, err := e.insertMemUser(registerType, emial, mobile, mobileTitle, refCode)
	if err != nil {
		return respCode, err
	}

	//插入用户配置
	userConfService := NewUserConfService(&e.Service)
	userConfReq := dto.UserConfInsertReq{}
	userConfReq.UserId = uid
	userConfReq.CurrUserId = uid
	userConfReq.CanLogin = global.SysStatusOk
	_, respCode, err = userConfService.Insert(&userConfReq)
	if err != nil {
		return respCode, err
	}

	// 若有别的表需要加，则在这里追加
	// todo 添加用户配置

	return baseLang.SuccessCode, nil
}

// insertMemUser app-内部方法，注册用户管理
func (e *User) insertMemUser(registerType, email, mobile, mobileTitle string, refUser *models.User) (int64, int, error) {
	user := models.User{}
	if registerType == constant.AccountMobileType {
		user.MobileTitle = mobileTitle
	}

	//确保推荐码不重复
	refCode := idgen.InviteId()
	queryUserCondition := dto.UserQueryReq{}
	queryUserCondition.RefCode = refCode
	count, respCode, err := e.Count(&queryUserCondition)
	if err != nil && respCode != baseLang.DataNotFoundCode {
		return 0, respCode, err
	}
	if count > 0 {
		return 0, baseLang.UserRefCodeErrLogCode, lang.MsgLogErrf(e.Log, e.Lang, baseLang.UserRegisterErrCode, baseLang.UserRefCodeErrLogCode, err)
	}

	//插入数据
	sysConfService := adminService.NewSysConfigService(&e.Service)
	defaultAvatar, respCode, err := sysConfService.GetWithKeyStr("admin_sys_user_default_avatar")
	if err != nil {
		return 0, respCode, err
	}
	refUserId := int64(0)
	refUserParentIds := ""
	refUserTreeSort := int64(0)
	refUserTreeSorts := ""
	user.TrueName = "- -"
	if refUser != nil {
		refUserId = refUser.Id
		refUserParentIds = refUser.ParentIds
		//refUserTreeSort = refUser.TreeSort
		refUserTreeSorts = refUser.TreeSorts
		queryUser := dto.UserQueryReq{}
		queryUser.ParentId = refUser.Id
		refUserTreeSort, respCode, err = e.queryMaxTreeSort(&queryUser)
		if err != nil && respCode != baseLang.DataNotFoundCode {
			return 0, respCode, err
		}
	}
	user.ParentId = refUserId
	user.ParentIds = refUserParentIds + strconv.FormatInt(refUserId, 10) + ","
	user.TreeSort = refUserTreeSort + 1
	user.TreeSorts = refUserTreeSorts + strconv.FormatInt(user.TreeSort, 10) + ","
	user.TreeLeaf = global.SysStatusOk
	user.TreeLevel = int64(strings.Count(user.ParentIds, ","))
	user.RefCode = refCode
	user.UserName = "- -"
	if registerType == constant.AccountEmailType {
		user.Email = email
	}
	if registerType == constant.AccountMobileType {
		user.Mobile = mobile
	}
	user.Avatar = defaultAvatar

	now := time.Now()
	user.CreatedAt = &now
	user.UpdatedAt = &now
	user.UpdateBy = user.Id
	user.CreateBy = user.Id
	user.Status = global.SysStatusOk
	err = e.Orm.Create(&user).Error
	if err != nil {
		return 0, baseLang.DataInsertLogCode, lang.MsgLogErrf(e.Log, e.Lang, baseLang.DataInsertCode, baseLang.DataInsertLogCode, err)
	}

	//如果原先上级为叶子节点，则将上级变为非叶子节点
	if refUser != nil && refUser.TreeLeaf == global.SysStatusOk {
		updateUser := map[string]interface{}{}
		updateUser["tree_leaf"] = global.SysStatusNotOk
		updateUser["updated_at"] = time.Now()
		updateUser["update_by"] = user.Id
		err = e.Orm.Model(&models.User{}).Where("id=?", refUser.Id).Updates(&updateUser).Error
		if err != nil {
			return 0, baseLang.DataInsertLogCode, lang.MsgLogErrf(e.Log, e.Lang, baseLang.DataInsertCode, baseLang.DataInsertLogCode, err)
		}
	}

	return user.Id, baseLang.SuccessCode, nil
}

// Update app-更新用户管理
func (e *User) Update(c *dto.UserUpdateReq, p *middleware.DataPermission) (bool, int, error) {
	if c.Id <= 0 || c.CurrUserId <= 0 {
		return false, baseLang.ParamErrCode, lang.MsgErr(baseLang.ParamErrCode, e.Lang)
	}
	data, respCode, err := e.Get(c.Id, p)
	if err != nil {
		return false, respCode, err
	}
	if (c.Mobile != "" && c.MobileTitle == "") || (c.Mobile == "" && c.MobileTitle != "") {
		return false, baseLang.UserMobileNeedTitleCode, lang.MsgErr(baseLang.UserMobileNeedTitleCode, e.Lang)
	}

	mobile := ""
	if c.Mobile != "" {
		mobile, err = encrypt.AesEncrypt(strings.TrimSpace(c.Mobile), []byte(config.AuthConfig.Secret))
		if err != nil {
			return false, baseLang.UserMobileEncryptErrLogCode, lang.MsgLogErrf(e.Log, e.Lang, baseLang.UserMobileEncryptErrCode, baseLang.UserMobileEncryptErrLogCode, err)
		}
	}

	//邮箱加密
	email := ""
	if c.Email != "" {
		email, err = encrypt.AesEncrypt(c.Email, []byte(config.AuthConfig.Secret))
		if err != nil {
			return false, baseLang.UserEmailEncryptErrLogCode, lang.MsgLogErrf(e.Log, e.Lang, baseLang.UserEmailEncryptErrCode, baseLang.UserEmailEncryptErrLogCode, err)
		}
	}

	//验证手机号是否重复
	if mobile != "" {
		queryMobileUserCondition := dto.UserQueryReq{}
		queryMobileUserCondition.Mobile = mobile
		mobileUser, respCode, err := e.QueryOne(&queryMobileUserCondition, nil)
		if err != nil && respCode != baseLang.DataNotFoundCode {
			return false, respCode, err
		}
		if respCode == baseLang.SuccessCode && mobileUser.Id != c.Id && mobile == mobileUser.Mobile {
			return false, baseLang.UserMobileExistLogCode, lang.MsgErr(baseLang.UserMobileExistLogCode, e.Lang)
		}
	}

	//验证邮箱是否重复
	if email != "" {
		queryEmailUserCondition := dto.UserQueryReq{}
		queryEmailUserCondition.Email = email
		emailUserUser, respCode, err := e.QueryOne(&queryEmailUserCondition, nil)
		if err != nil && respCode != baseLang.DataNotFoundCode {
			return false, respCode, err
		}
		if respCode == baseLang.SuccessCode && emailUserUser.Id != c.Id && email == emailUserUser.Email {
			return false, baseLang.UserEmailExistLogCode, lang.MsgErr(baseLang.UserEmailExistLogCode, e.Lang)
		}
	}

	//最小化变更改动过的数据
	updates := map[string]interface{}{}
	if c.LevelId > 0 && data.LevelId != c.LevelId {
		updates["level_id"] = c.LevelId
	}
	if c.UserName != "" && data.UserName != c.UserName {
		updates["user_name"] = c.UserName
	}
	if c.TrueName != "" && data.TrueName != c.TrueName {
		updates["true_name"] = c.TrueName
	}
	if c.Email != "" && data.Email != c.Email {
		updates["email"] = email
	}
	if c.MobileTitle != "" && data.MobileTitle != c.MobileTitle {
		updates["mobile_title"] = c.MobileTitle
	}
	if c.Mobile != "" && data.Mobile != c.Mobile {
		updates["mobile"] = mobile
	}
	if len(updates) > 0 {
		updates["updated_at"] = time.Now()
		updates["update_by"] = c.CurrUserId
		err = e.Orm.Model(&data).Where("id=?", data.Id).Updates(&updates).Error
		if err != nil {
			return false, baseLang.DataUpdateLogCode, lang.MsgLogErrf(e.Log, e.Lang, baseLang.DataUpdateCode, baseLang.DataUpdateLogCode, err)
		}
		return true, baseLang.SuccessCode, nil
	}
	return false, baseLang.SuccessCode, nil
}

// UpdateStatus app-更新用户管理状态
func (e *User) UpdateStatus(c *dto.UserStatusUpdateReq, p *middleware.DataPermission) (bool, int, error) {
	if c.CurrUserId <= 0 || c.Id <= 0 {
		return false, baseLang.ParamErrCode, lang.MsgErr(baseLang.ParamErrCode, e.Lang)
	}
	if c.Status == "" {
		return false, baseLang.UserStatusEmptyCode, lang.MsgErr(baseLang.UserStatusEmptyCode, e.Lang)
	}
	var err error
	u, respCode, err := e.Get(c.Id, p)
	if err != nil {
		return false, respCode, err
	}

	updates := map[string]interface{}{}
	if c.Status != "" && u.Status != c.Status {
		updates["status"] = c.Status
	}

	if len(updates) > 0 {
		updates["update_by"] = c.CurrUserId
		updates["updated_at"] = time.Now()
		err = e.Orm.Model(&models.User{}).Where("id=?", c.Id).Updates(updates).Error
		if err != nil {
			return false, baseLang.DataUpdateLogCode, lang.MsgLogErrf(e.Log, e.Lang, baseLang.DataUpdateCode, baseLang.DataUpdateLogCode, err)
		}
		return true, baseLang.SuccessCode, nil
	}
	return false, baseLang.SuccessCode, nil
}

// Export app-导出用户管理
func (e *User) Export(list []models.User) ([]byte, error) {
	sheetName := "User"
	xlsx := excelize.NewFile()
	defer xlsx.Close()
	no, _ := xlsx.NewSheet(sheetName)
	_ = xlsx.SetColWidth(sheetName, "A", "L", 25)
	_ = xlsx.SetSheetRow(sheetName, "A1", &[]interface{}{
		"编号", "状态"})
	dictService := adminService.NewSysDictDataService(&e.Service)
	for i, item := range list {
		axis := fmt.Sprintf("A%d", i+2)
		status := dictService.GetLabel("admin_sys_status", item.Status)

		//按标签对应输入数据
		_ = xlsx.SetSheetRow(sheetName, axis, &[]interface{}{
			item.Id, status,
		})
	}
	xlsx.SetActiveSheet(no)
	data, _ := xlsx.WriteToBuffer()
	return data.Bytes(), nil
}

// GetSummaries app-统计用户管理数据
func (e *User) GetSummaries(c *dto.UserQueryReq, p *middleware.DataPermission) (*models.User, int, error) {
	var err error
	var data models.User

	result := models.User{}
	err = e.Orm.Model(&models.User{}).
		Select("sum(app_user.money) as money").
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			middleware.Permission(data.TableName(), p),
		).Find(&result).Limit(-1).Offset(-1).Error

	if err != nil {
		return nil, baseLang.DataQueryLogCode, lang.MsgLogErrf(e.Log, e.Lang, baseLang.DataQueryCode, baseLang.DataQueryLogCode, err)
	}
	return &result, baseLang.SuccessCode, nil
}
