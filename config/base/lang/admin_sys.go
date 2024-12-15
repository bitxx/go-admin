package lang

import (
	"go-admin/core/lang"
)

const (
	// 字典
	SysDictTypeNameEmptyCode   = 10000
	SysDictTypeTypeEmptyCode   = 10001
	SysDictTypeTypeExistCode   = 10002
	SysDictTypeTypeHasUsedCode = 10003
	SysDictDataLabelEmptyCode  = 10004
	SysDictDataValueEmptyCode  = 10005
	SysDictDataSortEmptyCode   = 10006
	SysDictDataValueExistCode  = 10007

	// 配置管理
	SysConfNameEmptyCode       = 10100
	SysConfKeyEmptyCode        = 10101
	SysConfValueEmptyCode      = 10102
	SysConfTypeEmptyCode       = 10103
	SysConfIsFrontendEmptyCode = 10104
	SysConfKeyExistCode        = 10105
	SysConfGetErrLogCode       = 10106
	SysConfGetErrCode          = 10107

	//部门管理
	SysDeptParentIdEmptyCode   = 10200
	SysDeptNameEmptyCode       = 10201
	SysDeptLeaderEmptyCode     = 10202
	SysDeptNameExistCode       = 10203
	SysDeptChildExistNoDelCode = 10204
	SysDeptParentSelfCode      = 10205

	//角色管理
	SysRoleNameEmptyCode         = 10301
	SysRoleStatusEmptyCode       = 10302
	SysRoleKeyEmptyCode          = 10303
	SysRoleSortEmptyCode         = 10304
	SysRoleKeyExistCode          = 10305
	SysRoleAdminNoOpCode         = 10306
	SysRoleUserExistNoDeleteCode = 10307

	//岗位管理
	SysPostNameEmptyCode   = 10400
	SysPostCodeEmptyCode   = 10401
	SysPostSortEmptyCode   = 10402
	SysPostStatusEmptyCode = 10403
	SysPostNameExistCode   = 10404

	//用户管理
	SysUserNameEmptyCode              = 10500
	SysNickNameEmptyCode              = 10501
	SysUserPhoneEmptyCode             = 10502
	SysUserEmailEmptyCode             = 10503
	SysUserDeptEmptyCode              = 10504
	SysUserPwdEmptyCode               = 10505
	SysUserNameExistCode              = 10506
	SysUserNickNameExistCode          = 10507
	SysUserPhoneExistCode             = 10508
	SysUserEmailExistCode             = 10509
	SysUserEmailFormatErrCode         = 10510
	SysUserStatusEmptyCode            = 10511
	SysUserNewPwdEmptyCode            = 10512
	SysUserPwdErrCode                 = 10513
	SysUserNoExistCode                = 10514
	SysUserAvatarErrCode              = 10515
	SysUserAvatarErrLogCode           = 10516
	SysUseCapErrLogCode               = 10517
	SysUseLogoutSuccessCode           = 10518
	SysUseLoginOpCode                 = 10519
	SysUseLogoutOpCode                = 10520
	SysUseGenCaptchaErrCode           = 10521
	SysUseGenCaptchaErrLogCode        = 10522
	SysUseAvatarUploadErrCode         = 10523
	SysUseAvatarUploadErrLogCode      = 10524
	SysAdminUserNotAllowDeleteErrCode = 10525
	SysUserNoRoleErrCode              = 10526

	//菜单
	SysMenuParentIdEmptyCode = 10600
	SysMenuTitleEmptyCode    = 10601
	SysMenuTypeEmptyCode     = 10602
	SysMenuSortEmptyCode     = 10603
	SysMenuHasChildCode      = 10604
	SysNoRoleMenuCode        = 10605

	//gen表
	SysGenTableSelectCode               = 10700
	SysGenTableInsertExistCode          = 10701
	SysGenTableImportErrLogCode         = 10702
	SysGenTableImportErrCode            = 10703
	SysGenTemplateModelReadErrCode      = 10704
	SysGenTemplateModelReadLogErrCode   = 10705
	SysGenTemplateApiReadErrCode        = 10706
	SysGenTemplateApiReadLogErrCode     = 10707
	SysGenTemplateJsReadErrCode         = 10708
	SysGenTemplateJsReadLogErrCode      = 10709
	SysGenTemplateVueReadErrCode        = 10710
	SysGenTemplateVueReadLogErrCode     = 10711
	SysGenTemplateRouterReadErrCode     = 10712
	SysGenTemplateRouterReadLogErrCode  = 10713
	SysGenTemplateDtoReadErrCode        = 10714
	SysGenTemplateDtoReadLogErrCode     = 10715
	SysGenTemplateServiceReadErrCode    = 10716
	SysGenTemplateServiceReadLogErrCode = 10717
	SysGenCreatePathLogErrCode          = 10718
	SysGenCreatePathErrCode             = 10719
	SysGenTemplateModelDecodeErrCode    = 10720
	SysGenTemplateModelDecodeLogErrCode = 10721
	SysGenFrontTypeErrCode              = 10722

	//API
	SysApiGetApiMqLogErrCode    = 10800
	SysApiAppendApiMqLogErrCode = 10801
	SysApiIsSyncErrCode         = 10802
	SysApiDirGetLogErrCode      = 10803
	SysApiDirGetErrCode         = 10804
	SysApiParseLogErrCode       = 10805
	SysApiParseErrCode          = 10806
)

func init() {
	if lang.MsgInfo == nil {
		return
	}
	//字典
	lang.MsgInfo[SysDictTypeNameEmptyCode] = "字典名称不得为空"
	lang.MsgInfo[SysDictTypeTypeEmptyCode] = "字典类型不得为空"
	lang.MsgInfo[SysDictTypeTypeExistCode] = "该字典类型已存在"
	lang.MsgInfo[SysDictTypeTypeHasUsedCode] = "该字典已被使用"
	lang.MsgInfo[SysDictDataLabelEmptyCode] = "标签不得为空"
	lang.MsgInfo[SysDictDataValueEmptyCode] = "键值不得为空不得为空"
	lang.MsgInfo[SysDictDataSortEmptyCode] = "排序值不得为空"
	lang.MsgInfo[SysDictDataValueExistCode] = "该键值已存在"

	//配置
	lang.MsgInfo[SysConfNameEmptyCode] = "配置名称不得为空"
	lang.MsgInfo[SysConfKeyEmptyCode] = "键名不得为空"
	lang.MsgInfo[SysConfValueEmptyCode] = "键值不得为空"
	lang.MsgInfo[SysConfTypeEmptyCode] = "配置类型不得为空"
	lang.MsgInfo[SysConfIsFrontendEmptyCode] = "前置不得为空"
	lang.MsgInfo[SysConfKeyExistCode] = "该配置已存在"
	lang.MsgInfo[SysConfGetErrLogCode] = "数据获取失败：%s"
	lang.MsgInfo[SysConfGetErrCode] = "数据获取失败"

	//部门
	lang.MsgInfo[SysDeptParentIdEmptyCode] = "上级部门选择异常"
	lang.MsgInfo[SysDeptNameEmptyCode] = "部门名称不得为空"
	lang.MsgInfo[SysDeptLeaderEmptyCode] = "负责人不得为空"
	lang.MsgInfo[SysDeptNameExistCode] = "该部门名称已存在"
	lang.MsgInfo[SysDeptChildExistNoDelCode] = "该部门有下级部门，不可删除"
	lang.MsgInfo[SysDeptParentSelfCode] = "不可将自己设置为上级部门"

	//角色
	lang.MsgInfo[SysRoleNameEmptyCode] = "角色名称不得为空"
	lang.MsgInfo[SysRoleStatusEmptyCode] = "角色状态不得为空"
	lang.MsgInfo[SysRoleKeyEmptyCode] = "角色标识不得为空"
	lang.MsgInfo[SysRoleSortEmptyCode] = "角色排序不得为空"
	lang.MsgInfo[SysRoleKeyExistCode] = "该角色类型已存在"
	lang.MsgInfo[SysRoleAdminNoOpCode] = "超级管理员不支持该操作"
	lang.MsgInfo[SysRoleUserExistNoDeleteCode] = "由用户存在，不可删除"

	//岗位管理
	lang.MsgInfo[SysPostNameEmptyCode] = "岗位名称不得为空"
	lang.MsgInfo[SysPostCodeEmptyCode] = "岗位编码不得为空"
	lang.MsgInfo[SysPostSortEmptyCode] = "岗位排序不得为空"
	lang.MsgInfo[SysPostStatusEmptyCode] = "岗位状态不得为空"
	lang.MsgInfo[SysPostNameExistCode] = "岗位名称已存在"

	//用户管理
	lang.MsgInfo[SysUserNameEmptyCode] = "用户名不得为空"
	lang.MsgInfo[SysNickNameEmptyCode] = "用户昵称不得为空"
	lang.MsgInfo[SysUserPhoneEmptyCode] = "用户联系电话不得为空"
	lang.MsgInfo[SysUserEmailEmptyCode] = "用户邮箱不得为空"
	lang.MsgInfo[SysUserDeptEmptyCode] = "用户部门不得为空"
	lang.MsgInfo[SysUserPwdEmptyCode] = "用户密码不得为空"
	lang.MsgInfo[SysUserNameExistCode] = "用户名已存在"
	lang.MsgInfo[SysUserNickNameExistCode] = "用户昵称已存在"
	lang.MsgInfo[SysUserPhoneExistCode] = "用户联系电话已存在"
	lang.MsgInfo[SysUserEmailExistCode] = "用户邮箱已存在"
	lang.MsgInfo[SysUserEmailFormatErrCode] = "用户邮箱格式错误"
	lang.MsgInfo[SysUserStatusEmptyCode] = "用户状态不得为空"
	lang.MsgInfo[SysUserNewPwdEmptyCode] = "新密码不得为空"
	lang.MsgInfo[SysUserPwdErrCode] = "密码错误"
	lang.MsgInfo[SysUserNoExistCode] = "账户不存在"
	lang.MsgInfo[SysUserAvatarErrCode] = "头像上传失败"
	lang.MsgInfo[SysUserAvatarErrLogCode] = "头像上传失败：%s"
	lang.MsgInfo[SysUseCapErrLogCode] = "验证码错误"
	lang.MsgInfo[SysUseLogoutSuccessCode] = "退出成功"
	lang.MsgInfo[SysUseLoginOpCode] = "登录操作"
	lang.MsgInfo[SysUseLogoutOpCode] = "退出操作"
	lang.MsgInfo[SysUseGenCaptchaErrCode] = "验证码获取失败"
	lang.MsgInfo[SysUseGenCaptchaErrLogCode] = "验证码获取失败：%s"
	lang.MsgInfo[SysUseAvatarUploadErrCode] = "头像上传失败"
	lang.MsgInfo[SysUseAvatarUploadErrLogCode] = "头像上传失败：%s"
	lang.MsgInfo[SysAdminUserNotAllowDeleteErrCode] = "该账户禁止删除"
	lang.MsgInfo[SysUserNoRoleErrCode] = "该账户尚未分配角色"

	//菜单
	lang.MsgInfo[SysMenuParentIdEmptyCode] = "上级菜单不得为空"
	lang.MsgInfo[SysMenuTitleEmptyCode] = "菜单标题不得为空"
	lang.MsgInfo[SysMenuTypeEmptyCode] = "菜单类型不得为空"
	lang.MsgInfo[SysMenuSortEmptyCode] = "菜单排序不得为空"
	lang.MsgInfo[SysMenuHasChildCode] = "该菜单有下级，不可删除"
	lang.MsgInfo[SysNoRoleMenuCode] = "该账户无授权功能可用，请联系管理员"

	//表
	lang.MsgInfo[SysGenTableSelectCode] = "请选择表"
	lang.MsgInfo[SysGenTableInsertExistCode] = "导入失败，不可表重复导入相同的表"
	lang.MsgInfo[SysGenTableImportErrLogCode] = "表数据导入失败：%s"
	lang.MsgInfo[SysGenTableImportErrCode] = "表数据导入失败"
	lang.MsgInfo[SysGenTemplateModelReadErrCode] = "model模版读取失败"
	lang.MsgInfo[SysGenTemplateModelReadLogErrCode] = "model模版读取失败：%s"
	lang.MsgInfo[SysGenTemplateApiReadErrCode] = "api模版读取失败"
	lang.MsgInfo[SysGenTemplateApiReadLogErrCode] = "api模版读取失败：%s"
	lang.MsgInfo[SysGenTemplateJsReadErrCode] = "js模版读取失败"
	lang.MsgInfo[SysGenTemplateJsReadLogErrCode] = "js模版读取失败：%s"
	lang.MsgInfo[SysGenTemplateVueReadErrCode] = "vue模版读取失败"
	lang.MsgInfo[SysGenTemplateVueReadLogErrCode] = "vue模版读取失败：%s"
	lang.MsgInfo[SysGenTemplateRouterReadErrCode] = "router模版读取失败"
	lang.MsgInfo[SysGenTemplateRouterReadLogErrCode] = "router模版读取失败：%s"
	lang.MsgInfo[SysGenTemplateDtoReadErrCode] = "dto模版读取失败"
	lang.MsgInfo[SysGenTemplateDtoReadLogErrCode] = "dto模版读取失败：%s"
	lang.MsgInfo[SysGenTemplateServiceReadErrCode] = "service模版读取失败"
	lang.MsgInfo[SysGenTemplateServiceReadLogErrCode] = "service模版读取失败：%s"
	lang.MsgInfo[SysGenCreatePathLogErrCode] = "创建目录失败：%s"
	lang.MsgInfo[SysGenCreatePathErrCode] = "创建目录失败"
	lang.MsgInfo[SysGenTemplateModelDecodeErrCode] = "模板解析异常"
	lang.MsgInfo[SysGenTemplateModelDecodeLogErrCode] = "模板解析异常：%s"
	lang.MsgInfo[SysGenFrontTypeErrCode] = "前端类型配置异常"
	lang.MsgInfo[SysApiGetApiMqLogErrCode] = "获取接口数据队列异常：s%"
	lang.MsgInfo[SysApiAppendApiMqLogErrCode] = "接口数据队列添加异常]= s%"
	lang.MsgInfo[SysApiIsSyncErrCode] = "接口数据正在同步中，请稍后"
	lang.MsgInfo[SysApiDirGetLogErrCode] = "接口文件路径获取异常：%s"
	lang.MsgInfo[SysApiDirGetErrCode] = "接口文件路径获取异常"
	lang.MsgInfo[SysApiParseLogErrCode] = "api文件解析异常：%s"
	lang.MsgInfo[SysApiParseErrCode] = "api文件解析异常"
}
