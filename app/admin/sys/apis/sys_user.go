package apis

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"go-admin/app/admin/sys/constant"
	sysLang "go-admin/app/admin/sys/lang"
	"go-admin/app/admin/sys/service"
	"go-admin/app/admin/sys/service/dto"
	"go-admin/core/config"
	"go-admin/core/dto/api"
	_ "go-admin/core/dto/response"
	"go-admin/core/global"
	"go-admin/core/lang"
	"go-admin/core/middleware"
	"go-admin/core/middleware/auth"
	"go-admin/core/middleware/auth/authdto"
	"go-admin/core/utils/captchautils"
	"go-admin/core/utils/fileutils"
	"go-admin/core/utils/idgen"
)

type SysUser struct {
	api.Api
}

// GetPage sys-获取系统用户管理分页列表
func (e SysUser) GetPage(c *gin.Context) {
	s := service.SysUser{}
	req := dto.SysUserQueryReq{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Error(lang.DataDecodeCode, lang.MsgLogErrf(e.Logger, e.Lang, lang.DataDecodeCode, lang.DataDecodeLogCode, err).Error())
		return
	}
	p := middleware.GetPermissionFromContext(c)
	list, count, respCode, err := s.GetPage(&req, p)
	if err != nil {
		e.Error(respCode, err.Error())
		return
	}
	e.PageOK(list, nil, count, req.GetPageIndex(), req.GetPageSize(), lang.MsgByCode(lang.SuccessCode, e.Lang))
}

// Get sys-获取系统用户管理详情
func (e SysUser) Get(c *gin.Context) {
	s := service.SysUser{}
	req := dto.SysUserGetReq{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Error(lang.DataDecodeCode, lang.MsgLogErrf(e.Logger, e.Lang, lang.DataDecodeCode, lang.DataDecodeLogCode, err).Error())
		return
	}
	p := middleware.GetPermissionFromContext(c)
	result, respCode, err := s.Get(req.Id, p)
	if err != nil {
		e.Error(respCode, err.Error())
		return
	}
	e.OK(result, lang.MsgByCode(lang.SuccessCode, e.Lang))
}

// Insert sys-新增系统用户管理
func (e SysUser) Insert(c *gin.Context) {
	s := service.SysUser{}
	req := dto.SysUserInsertReq{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Error(lang.DataDecodeCode, lang.MsgLogErrf(e.Logger, e.Lang, lang.DataDecodeCode, lang.DataDecodeLogCode, err).Error())
		return
	}
	uid, rCode, err := auth.Auth.GetUserId(c)
	if err != nil {
		e.Error(rCode, err.Error())
		return
	}
	req.CurrUserId = uid
	id, respCode, err := s.Insert(&req)
	if err != nil {
		e.Error(respCode, err.Error())
		return
	}
	e.OK(id, lang.MsgByCode(lang.SuccessCode, e.Lang))
}

// Update sys-更新系统用户管理
func (e SysUser) Update(c *gin.Context) {
	s := service.SysUser{}
	req := dto.SysUserUpdateReq{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Error(lang.DataDecodeCode, lang.MsgLogErrf(e.Logger, e.Lang, lang.DataDecodeCode, lang.DataDecodeLogCode, err).Error())
		return
	}
	p := middleware.GetPermissionFromContext(c)
	uid, rCode, err := auth.Auth.GetUserId(c)
	if err != nil {
		e.Error(rCode, err.Error())
		return
	}
	req.CurrUserId = uid
	b, respCode, err := s.Update(&req, p)
	if err != nil {
		e.Error(respCode, err.Error())
		return
	}
	if !b {
		e.OK(nil, lang.MsgByCode(lang.DataNotUpdateCode, e.Lang))
		return
	}
	e.OK(nil, lang.MsgByCode(lang.SuccessCode, e.Lang))
}

// Delete sys-删除系统用户管理
func (e SysUser) Delete(c *gin.Context) {
	s := service.SysUser{}
	req := dto.SysUserDeleteReq{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Error(lang.DataDecodeCode, lang.MsgLogErrf(e.Logger, e.Lang, lang.DataDecodeCode, lang.DataDecodeLogCode, err).Error())
		return
	}

	p := middleware.GetPermissionFromContext(c)
	respCode, err := s.Delete(req.Ids, p)
	if err != nil {
		e.Error(respCode, err.Error())
		return
	}
	e.OK(nil, lang.MsgByCode(lang.SuccessCode, e.Lang))
}

// UpdateStatus sys-更新系统用户状态
func (e SysUser) UpdateStatus(c *gin.Context) {
	s := service.SysUser{}
	req := dto.SysUserStatusUpdateReq{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req, binding.JSON, nil).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Error(lang.DataDecodeCode, lang.MsgLogErrf(e.Logger, e.Lang, lang.DataDecodeCode, lang.DataDecodeLogCode, err).Error())
		return
	}

	uid, rCode, err := auth.Auth.GetUserId(c)
	if err != nil {
		e.Error(rCode, err.Error())
		return
	}
	req.CurrUserId = uid

	//数据权限检查
	p := middleware.GetPermissionFromContext(c)

	b, respCode, err := s.UpdateStatus(&req, p)
	if err != nil {
		e.Error(respCode, err.Error())
		return
	}
	if !b {
		e.OK(nil, lang.MsgByCode(lang.DataNotUpdateCode, e.Lang))
		return
	}
	e.OK(nil, lang.MsgByCode(lang.SuccessCode, e.Lang))
}

// ResetPwd sys-重置系统用户密码
func (e SysUser) ResetPwd(c *gin.Context) {
	s := service.SysUser{}
	req := dto.ResetSysUserPwdReq{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Error(lang.DataDecodeCode, lang.MsgLogErrf(e.Logger, e.Lang, lang.DataDecodeCode, lang.DataDecodeLogCode, err).Error())
		return
	}

	uid, rCode, err := auth.Auth.GetUserId(c)
	if err != nil {
		e.Error(rCode, err.Error())
		return
	}
	req.CurrUserId = uid

	//数据权限检查
	p := middleware.GetPermissionFromContext(c)

	b, respCode, err := s.ResetPwd(&req, p)
	if err != nil {
		e.Error(respCode, err.Error())
		return
	}
	if !b {
		e.OK(nil, lang.MsgByCode(lang.DataNotUpdateCode, e.Lang))
		return
	}
	e.OK(nil, lang.MsgByCode(lang.SuccessCode, e.Lang))
}

// UpdateProfileAvatar sys-更新系统登录用户头像
func (e SysUser) UpdateProfileAvatar(c *gin.Context) {
	s := service.SysUser{}
	req := dto.SysUserAvatarUpdateReq{}
	err := e.MakeContext(c).
		MakeOrm().
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Error(lang.DataDecodeCode, lang.MsgLogErrf(e.Logger, e.Lang, lang.DataDecodeCode, lang.DataDecodeLogCode, err).Error())
		return
	}

	form, _ := c.MultipartForm()
	files := form.File["avatar"]
	guid := idgen.UUID()
	reqPath := config.ApplicationConfig.FileRootPath + "admin/avatar/"
	err = fileutils.IsNotExistMkDir(reqPath)
	if err != nil {
		e.Error(sysLang.SysUseAvatarUploadErrLogCode, lang.MsgLogErrf(e.Logger, e.Lang, sysLang.SysUseAvatarUploadErrCode, sysLang.SysUseAvatarUploadErrLogCode, err).Error())
		/*err = fileutil.CreateDirAll(reqPath)
		if err != nil {
			e.Error(sysLang.SysUseAvatarUploadErrLogCode, lang.MsgLogErrf(e.Logger, e.Lang, sysLang.SysUseAvatarUploadErrCode, sysLang.SysUseAvatarUploadErrLogCode, err).Error())
			return
		}*/
	}
	filPath := reqPath + guid + ".jpg"
	for _, file := range files {
		// 上传文件至指定目录
		err = c.SaveUploadedFile(file, filPath)
		if err != nil {
			e.Error(sysLang.SysUseAvatarUploadErrLogCode, lang.MsgLogErrf(e.Logger, e.Lang, sysLang.SysUseAvatarUploadErrCode, sysLang.SysUseAvatarUploadErrLogCode, err).Error())
			return
		}
	}
	// 数据权限检查
	req.Avatar = global.RouteRootPath + "/" + filPath

	uid, rCode, err := auth.Auth.GetUserId(c)
	if err != nil {
		e.Error(rCode, err.Error())
		return
	}
	req.CurrUserId = uid

	p := middleware.GetPermissionFromContext(c)
	b, respCode, err := s.UpdateProfileAvatar(&req, p)
	if err != nil {
		e.Error(respCode, err.Error())
		return
	}
	if !b {
		e.OK(nil, lang.MsgByCode(lang.DataNotUpdateCode, e.Lang))
		return
	}
	e.OK(req.Avatar, lang.MsgByCode(lang.SuccessCode, e.Lang))
}

// UpdateProfilePwd sys-更新系统登录用户密码
func (e SysUser) UpdateProfilePwd(c *gin.Context) {
	s := service.SysUser{}
	req := dto.UpdateSysUserPwdReq{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Error(lang.DataDecodeCode, lang.MsgLogErrf(e.Logger, e.Lang, lang.DataDecodeCode, lang.DataDecodeLogCode, err).Error())
		return
	}

	// 数据权限检查
	p := middleware.GetPermissionFromContext(c)
	uid, rCode, err := auth.Auth.GetUserId(c)
	if err != nil {
		e.Error(rCode, err.Error())
		return
	}
	req.CurrUserId = uid
	b, respCode, err := s.UpdateProfilePwd(req, p)
	if err != nil {
		e.Error(respCode, err.Error())
		return
	}
	if !b {
		e.OK(nil, lang.MsgByCode(lang.DataNotUpdateCode, e.Lang))
		return
	}
	e.OK(nil, lang.MsgByCode(lang.SuccessCode, e.Lang))
}

// GetProfile sys-获取系统登录用户信息
func (e SysUser) GetProfile(c *gin.Context) {
	s := service.SysUser{}
	err := e.MakeContext(c).
		MakeOrm().
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Error(lang.DataDecodeCode, lang.MsgLogErrf(e.Logger, e.Lang, lang.DataDecodeCode, lang.DataDecodeLogCode, err).Error())
		return
	}

	uid, rCode, err := auth.Auth.GetUserId(c)
	if err != nil {
		e.Error(rCode, err.Error())
		return
	}

	user, respCode, err := s.GetProfile(uid)
	if err != nil {
		e.Error(respCode, err.Error())
		return
	}
	e.OK(user, lang.MsgByCode(lang.SuccessCode, e.Lang))
}

// UpdateProfile sys-更新系统登录用户信息
// 当前登录用户才能更新自己的信息
// 受限的子账户登录时，为了数据安全，不能让用户通过Update方法/接口来修改自己账户
func (e SysUser) UpdateProfile(c *gin.Context) {
	s := service.SysUser{}
	req := dto.SysUserUpdateReq{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Error(lang.DataDecodeCode, lang.MsgLogErrf(e.Logger, e.Lang, lang.DataDecodeCode, lang.DataDecodeLogCode, err).Error())
		return
	}
	uid, rCode, err := auth.Auth.GetUserId(c)
	if err != nil {
		e.Error(rCode, err.Error())
		return
	}
	req.CurrUserId = uid
	req.Id = uid
	b, respCode, err := s.UpdateProfile(&req)
	if err != nil {
		e.Error(respCode, err.Error())
		return
	}
	if !b {
		e.OK(nil, lang.MsgByCode(lang.DataNotUpdateCode, e.Lang))
		return
	}
	e.OK(nil, lang.MsgByCode(lang.SuccessCode, e.Lang))
}

// Login sys-登录系统
func (e SysUser) Login(c *gin.Context) {
	req := dto.LoginReq{}
	s := service.SysUser{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Error(lang.DataDecodeCode, lang.MsgLogErrf(e.Logger, e.Lang, lang.DataDecodeCode, lang.DataDecodeLogCode, err).Error())
		return
	}

	if req.Code == "" || req.Password == "" || req.Username == "" {
		e.Error(lang.ParamErrCode, lang.MsgByCode(lang.ParamErrCode, e.Lang))
		return
	}

	if config.ApplicationConfig.Mode != "dev" {
		if !captchautils.Verify(req.UUID, req.Code, true) {
			e.Error(sysLang.SysUseCapErrLogCode, lang.MsgByCode(sysLang.SysUseCapErrLogCode, e.Lang))
			return
		}
	}

	userResp, respCode, err := s.LoginVerify(&req)
	if err != nil {
		e.Error(respCode, err.Error())
		return
	}

	c.Set(authdto.LoginUserId, userResp.Id)
	c.Set(authdto.UserName, userResp.Username)
	c.Set(authdto.RoleId, userResp.Role.Id)
	c.Set(authdto.RoleKey, userResp.Role.RoleKey)
	c.Set(authdto.DeptId, userResp.Dept.Id)
	c.Set(authdto.DataScope, userResp.Role.DataScope)
	c.Set(authdto.UserInfo, userResp)
	auth.Auth.Login(c)
	s.LoginLogToDB(c, constant.UserLoginStatus, lang.MsgByCode(sysLang.SysUseLoginOpCode, e.Lang), userResp.Id)
}

// LogOut sys-退出系统
func (e SysUser) LogOut(c *gin.Context) {
	s := new(service.SysUser)
	err := e.MakeContext(c).
		MakeOrm().
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Error(lang.DataDecodeCode, lang.MsgLogErrf(e.Logger, e.Lang, lang.DataDecodeCode, lang.DataDecodeLogCode, err).Error())
		return
	}
	uid, rCode, err := auth.Auth.GetUserId(c)
	if err != nil {
		e.Error(rCode, err.Error())
		return
	}
	s.LoginLogToDB(c, constant.UserLogoutStatus, lang.MsgByCode(sysLang.SysUseLoginOpCode, e.Lang), uid)

	e.OK(nil, lang.MsgByCode(sysLang.SysUseLogoutSuccessCode, e.Lang))
}

// GenCaptcha sys-获取图形验证码
func (e SysUser) GenCaptcha(c *gin.Context) {
	err := e.MakeContext(c).Errors
	if err != nil {
		e.Error(lang.DataDecodeCode, lang.MsgLogErrf(e.Logger, e.Lang, lang.DataDecodeCode, lang.DataDecodeLogCode, err).Error())
		return
	}
	id, b64s, err := captchautils.DriverDigitFunc()
	if err != nil {
		e.Error(sysLang.SysUseGenCaptchaErrLogCode, lang.MsgLogErrf(e.Logger, e.Lang, sysLang.SysUseGenCaptchaErrCode, sysLang.SysUseGenCaptchaErrLogCode, err).Error())
		return
	}
	resp := map[string]string{
		"data": b64s,
		"id":   id,
	}
	e.OK(resp, lang.MsgByCode(lang.SuccessCode, e.Lang))
}
