package service

import (
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	adminService "go-admin/app/admin/service"
	"go-admin/app/app/user/constant"
	uLang "go-admin/app/app/user/lang"
	"go-admin/app/app/user/models"
	"go-admin/app/app/user/service/dto"
	"go-admin/common/core/service"
	cDto "go-admin/common/dto"
	"go-admin/common/global"
	"go-admin/common/middleware"
	"go-admin/common/utils/encrypt"
	"go-admin/common/utils/idgen"
	"go-admin/common/utils/strutils"
	"strconv"
	"strings"

	"go-admin/config/lang"
	"gorm.io/gorm"
	"time"
)

type User struct {
	service.Service
}

// NewUserService
// @Description: 实例化User
// @param s
// @return *User
func NewUserService(s *service.Service) *User {
	var srv = new(User)
	srv.Orm = s.Orm
	srv.Log = s.Log
	return srv
}

// GetPage
// @Description: 获取User列表
// @receiver e
// @param c
// @param p
// @return []models.User
// @return int64
// @return int
// @return error
func (e *User) GetPage(c *dto.UserQueryReq, p *middleware.DataPermission) ([]models.User, int64, int, error) {
	var data models.User
	var list []models.User
	var count int64
	if c.Mobile != "" {
		c.Mobile, _ = encrypt.AesEncryptDefault(c.Mobile)
	}

	if c.Email != "" {
		c.Email, _ = encrypt.AesEncryptDefault(c.Email)
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
		return nil, 0, lang.DataQueryLogCode, lang.MsgLogErrf(e.Log, e.Lang, lang.DataQueryCode, lang.DataQueryLogCode, err)
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
			list[index].ParentRefCode = parentUser.RefCode
			list[index].ParentUserName = parentUser.UserName
		}

		mobile, err := encrypt.AesDecryptDefault(user.Mobile)
		if err == nil {
			list[index].Mobile = strutils.HidePartStr(mobile, 3)
		}

		email, err := encrypt.AesDecryptDefault(user.Email)
		if err == nil {
			list[index].Email = strutils.HidePartStr(email, 5)
		}
	}
	return list, count, lang.SuccessCode, nil
}

// Get
// @Description: 获取User对象
// @receiver e
// @param id 编号
// @param p
// @return *models.User
// @return int
// @return error
func (e *User) Get(id int64, p *middleware.DataPermission) (*models.User, int, error) {
	if id <= 0 {
		return nil, lang.ParamErrCode, lang.MsgErr(lang.ParamErrCode, e.Lang)
	}
	data := &models.User{}
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

// QueryOne
// @Description: 通过自定义条件获取User一条记录
// @receiver e
// @param queryCondition 条件
// @return *models.User
// @return error
func (e *User) QueryOne(queryCondition *dto.UserQueryReq, p *middleware.DataPermission) (*models.User, int, error) {
	data := &models.User{}
	err := e.Orm.Scopes(
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

func (e *User) queryMaxTreeSort(queryCondition *dto.UserQueryReq) (int64, int, error) {
	maxSort := int64(0)
	err := e.Orm.Scopes(
		cDto.MakeCondition(queryCondition.GetNeedSearch()),
	).Model(&models.User{}).Select("max(tree_sort)").Scan(&maxSort).Error
	if err != nil {
		maxSort = 0
	}
	return maxSort, lang.SuccessCode, nil
}

// Count
//
//	@Description: 获取条数
//	@receiver e
//	@param c
//	@return int64
//	@return int
//	@return error
func (e *User) Count(queryCondition *dto.UserQueryReq) (int64, int, error) {
	var err error
	var count int64
	err = e.Orm.Model(&models.User{}).
		Scopes(
			cDto.MakeCondition(queryCondition.GetNeedSearch()),
		).Limit(-1).Offset(-1).Count(&count).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return 0, lang.DataQueryLogCode, lang.MsgLogErrf(e.Log, e.Lang, lang.DataQueryCode, lang.DataQueryLogCode, err)
	}
	if err == gorm.ErrRecordNotFound {
		return 0, lang.DataNotFoundCode, lang.MsgErr(lang.DataNotFoundCode, e.Lang)
	}
	return count, lang.SuccessCode, nil
}

// Insert
// @Description: 创建User对象
// @receiver e
// @param c
// @return int64 插入数据的主键
// @return int
// @return error
func (e *User) Insert(c *dto.UserInsertReq) (int, error) {
	if c.CurrUserId <= 0 {
		return lang.ParamErrCode, lang.MsgErr(lang.ParamErrCode, e.Lang)
	}
	if c.Emails == "" && c.Mobiles == "" {
		return uLang.AppUserEmailOrMobileNeedCode, lang.MsgErr(uLang.AppUserEmailOrMobileNeedCode, e.Lang)
	}

	var mobiles, emails []string
	if c.Emails != "" {
		emails = strings.Split(c.Emails, ",")
		for _, email := range emails {
			fmt.Println(!strutils.IsEmail(email))
			if !strutils.IsEmail(email) {
				return uLang.AppUserEmailFormatErrCode, lang.MsgErr(uLang.AppUserEmailFormatErrCode, e.Lang)
			}
		}
	}

	if c.Mobiles != "" {
		if c.MobileTitle == "" {
			return uLang.AppUserMobileNeedTitleCode, lang.MsgErr(uLang.AppUserMobileNeedTitleCode, e.Lang)
		}
		mobiles = strings.Split(c.Mobiles, ",")
		for _, mobile := range mobiles {
			if !strutils.IsMobile(mobile) {
				return uLang.AppUserMobileFormatErrCode, lang.MsgErr(uLang.AppUserMobileFormatErrCode, e.Lang)
			}
		}
	}
	var refUser *models.User
	if c.RefCode != "" {
		queryUserCondition := dto.UserQueryReq{}
		queryUserCondition.RefCode = c.RefCode
		queryUserCondition.Status = global.SysStatusOk
		u, respCode, err := e.QueryOne(&queryUserCondition, nil)
		if err != nil && respCode != lang.DataNotFoundCode {
			return respCode, err
		}
		if respCode == lang.SuccessCode {
			refUser = u
		}
	}

	//事务 开启-关闭
	var err error
	respCode := lang.SuccessCode
	e.Orm = e.Orm.Begin()
	defer func() {
		if err != nil {
			e.Orm.Rollback()
		} else {
			e.Orm.Commit()
		}
	}()

	for _, mobile := range mobiles {
		mobile, err = encrypt.AesEncryptDefault(mobile)
		if err != nil {
			return uLang.AppUserMobileEncryptErrLogCode, lang.MsgLogErrf(e.Log, e.Lang, uLang.AppUserMobileEncryptErrCode, uLang.AppUserMobileEncryptErrLogCode, err)
		}
		respCode, err = e.register(constant.AccountMobileType, "", mobile, c.MobileTitle, refUser)
		if err != nil {
			return respCode, err
		}
	}
	for _, email := range emails {
		email, err = encrypt.AesEncryptDefault(email)
		if err != nil {
			return uLang.AppUserEmailEncryptErrLogCode, lang.MsgLogErrf(e.Log, e.Lang, uLang.AppUserEmailEncryptErrCode, uLang.AppUserEmailEncryptErrLogCode, err)
		}
		respCode, err = e.register(constant.AccountEmailType, email, "", "", refUser)
		if err != nil {
			return respCode, err
		}
	}

	return lang.SuccessCode, nil
}

func (e *User) register(registerType, email, mobile, mobileTitle string, refUser *models.User) (int, error) {

	//验证手机号和邮箱是否重复
	queryUserCondition := dto.UserQueryReq{}
	if registerType == constant.AccountMobileType {
		if (mobile != "" && mobileTitle == "") || (mobile == "" && mobileTitle != "") || (mobile == "" && mobileTitle == "") {
			return uLang.AppUserMobileNeedTitleCode, lang.MsgErr(uLang.AppUserMobileNeedTitleCode, e.Lang)
		}
		queryUserCondition.Mobile = mobile
		queryUserCondition.MobileTitle = mobileTitle
	}
	if registerType == constant.AccountEmailType {
		queryUserCondition.Email = email
	}

	count, respCode, err := e.Count(&queryUserCondition)
	if err != nil && respCode != lang.DataNotFoundCode {
		return respCode, err
	}
	if count > 0 {
		account := ""
		if email != "" {
			account, _ = encrypt.AesDecryptDefault(email)
		} else if mobile != "" {
			account, _ = encrypt.AesDecryptDefault(mobile)
		}
		return uLang.AppUserAccountExistLogCode, lang.MsgErrf(uLang.AppUserAccountExistLogCode, e.Lang, account)
	}

	return e.insertRegisterInfo(registerType, email, mobile, mobileTitle, refUser)
}

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

	return lang.SuccessCode, nil
}

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
	if err != nil && respCode != lang.DataNotFoundCode {
		return 0, respCode, err
	}
	if count > 0 {
		return 0, uLang.AppUserRefCodeErrLogCode, lang.MsgLogErrf(e.Log, e.Lang, uLang.AppUserRegisterErrCode, uLang.AppUserRefCodeErrLogCode, err)
	}

	//插入数据
	sysConfService := adminService.NewSysConfigService(&e.Service)
	defaultAvatar, respCode, err := sysConfService.GetWithKeyStr("app_user_default_avatar")
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
		if err != nil && respCode != lang.DataNotFoundCode {
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
		return 0, lang.DataInsertLogCode, lang.MsgLogErrf(e.Log, e.Lang, lang.DataInsertCode, lang.DataInsertLogCode, err)
	}

	//如果原先上级为叶子节点，则将上级变为非叶子节点
	if refUser != nil && refUser.TreeLeaf == global.SysStatusOk {
		updateUser := map[string]interface{}{}
		updateUser["tree_leaf"] = global.SysStatusNotOk
		updateUser["updated_at"] = time.Now()
		updateUser["update_by"] = user.Id
		err = e.Orm.Model(&models.User{}).Where("id=?", refUser.Id).Updates(&updateUser).Error
		if err != nil {
			return 0, lang.DataInsertLogCode, lang.MsgLogErrf(e.Log, e.Lang, lang.DataInsertCode, lang.DataInsertLogCode, err)
		}
	}

	return user.Id, lang.SuccessCode, nil
}

// Update
// @Description: 修改User对象
// @receiver e
// @param c
// @param p
// @return bool 是否有数据更新
// @return error
func (e *User) Update(c *dto.UserUpdateReq, p *middleware.DataPermission) (bool, int, error) {
	if c.Id <= 0 || c.CurrUserId <= 0 {
		return false, lang.ParamErrCode, lang.MsgErr(lang.ParamErrCode, e.Lang)
	}
	data, respCode, err := e.Get(c.Id, p)
	if err != nil {
		return false, respCode, err
	}
	if (c.Mobile != "" && c.MobileTitle == "") || (c.Mobile == "" && c.MobileTitle != "") {
		return false, uLang.AppUserMobileNeedTitleCode, lang.MsgErr(uLang.AppUserMobileNeedTitleCode, e.Lang)
	}

	mobile := ""
	if c.Mobile != "" {
		mobile, err = encrypt.AesEncryptDefault(strings.TrimSpace(c.Mobile))
		if err != nil {
			return false, uLang.AppUserMobileEncryptErrLogCode, lang.MsgLogErrf(e.Log, e.Lang, uLang.AppUserMobileEncryptErrCode, uLang.AppUserMobileEncryptErrLogCode, err)
		}
	}

	//邮箱加密
	email := ""
	if c.Email != "" {
		email, err = encrypt.AesEncryptDefault(c.Email)
		if err != nil {
			return false, uLang.AppUserEmailEncryptErrLogCode, lang.MsgLogErrf(e.Log, e.Lang, uLang.AppUserEmailEncryptErrCode, uLang.AppUserEmailEncryptErrLogCode, err)
		}
	}

	//验证手机号是否重复
	if mobile != "" {
		queryMobileUserCondition := dto.UserQueryReq{}
		queryMobileUserCondition.Mobile = mobile
		mobileUser, respCode, err := e.QueryOne(&queryMobileUserCondition, nil)
		if err != nil && respCode != lang.DataNotFoundCode {
			return false, respCode, err
		}
		if respCode == lang.SuccessCode && mobileUser.Id != c.Id && mobile == mobileUser.Mobile {
			return false, uLang.AppUserMobileExistLogCode, lang.MsgErr(uLang.AppUserMobileExistLogCode, e.Lang)
		}
	}

	//验证邮箱是否重复
	if email != "" {
		queryEmailUserCondition := dto.UserQueryReq{}
		queryEmailUserCondition.Email = email
		emailUserUser, respCode, err := e.QueryOne(&queryEmailUserCondition, nil)
		if err != nil && respCode != lang.DataNotFoundCode {
			return false, respCode, err
		}
		if respCode == lang.SuccessCode && emailUserUser.Id != c.Id && email == emailUserUser.Email {
			return false, uLang.AppUserEmailExistLogCode, lang.MsgErr(uLang.AppUserEmailExistLogCode, e.Lang)
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
	if email != "" && data.Email != email {
		updates["email"] = email
	}
	if c.MobileTitle != "" && data.MobileTitle != c.MobileTitle {
		updates["mobile_title"] = c.MobileTitle
	}
	if mobile != "" && data.Mobile != mobile {
		updates["mobile"] = mobile
	}
	if len(updates) > 0 {
		updates["updated_at"] = time.Now()
		updates["update_by"] = c.CurrUserId
		err = e.Orm.Model(&data).Where("id=?", data.Id).Updates(&updates).Error
		if err != nil {
			return false, lang.DataUpdateLogCode, lang.MsgLogErrf(e.Log, e.Lang, lang.DataUpdateCode, lang.DataUpdateLogCode, err)
		}
		return true, lang.SuccessCode, nil
	}
	return false, lang.SuccessCode, nil
}

// UpdateStatus 更新用户状态
func (e *User) UpdateStatus(c *dto.UserStatusUpdateReq, p *middleware.DataPermission) (bool, int, error) {
	if c.CurrUserId <= 0 || c.Id <= 0 {
		return false, lang.ParamErrCode, lang.MsgErr(lang.ParamErrCode, e.Lang)
	}
	if c.Status == "" {
		return false, uLang.AppUserStatusEmptyCode, lang.MsgErr(uLang.AppUserStatusEmptyCode, e.Lang)
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
			return false, lang.DataUpdateLogCode, lang.MsgLogErrf(e.Log, e.Lang, lang.DataUpdateCode, lang.DataUpdateLogCode, err)
		}
		return true, lang.SuccessCode, nil
	}
	return false, lang.SuccessCode, nil
}

// GetExcel
// @Description: GetExcel 导出User excel数据
// @receiver e
// @param list
// @return []byte
// @return int
// @return error
func (e *User) GetExcel(list []models.User) ([]byte, error) {
	sheetName := "User"
	xlsx := excelize.NewFile()
	no := xlsx.NewSheet(sheetName)
	xlsx.SetColWidth(sheetName, "A", "L", 25)
	xlsx.SetSheetRow(sheetName, "A1", &[]interface{}{
		"编号", "状态"})
	dictService := adminService.NewSysDictDataService(&e.Service)
	for i, item := range list {
		axis := fmt.Sprintf("A%d", i+2)
		status := dictService.GetLabel("sys_status", item.Status)

		//按标签对应输入数据
		xlsx.SetSheetRow(sheetName, axis, &[]interface{}{
			item.Id, status,
		})
	}
	xlsx.SetActiveSheet(no)
	data, _ := xlsx.WriteToBuffer()
	return data.Bytes(), nil
}
