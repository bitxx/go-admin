package lang

import "go-admin/core/lang"

const (
	PluginsAnnouncementTitleEmptyCode   = 20000
	PluginsAnnouncementContentEmptyCode = 20001
	PluginsAnnouncementStatusEmptyCode  = 20002
	PluginsAnnouncementNumCode          = 20003
	PluginsAnnouncementTitleHasUsedCode = 20004

	PluginsCategoryNameCode        = 20100
	PluginsCategoryNameHasUsedCode = 20101
	PluginsCategoryNotFoundCode    = 20102

	PluginsArticleNameCode        = 20200
	PluginsArticleNameHasUsedCode = 20201
	PluginsArticleContentCode     = 20202
	PluginsArticleCatIdEmptyCode  = 20203
)

func init() {
	if lang.MsgInfo == nil {
		return
	}
	lang.MsgInfo[PluginsAnnouncementTitleEmptyCode] = "公告标题不得为空"
	lang.MsgInfo[PluginsAnnouncementContentEmptyCode] = "公告内容不得为空"
	lang.MsgInfo[PluginsAnnouncementStatusEmptyCode] = "公告状态不得为空"
	lang.MsgInfo[PluginsAnnouncementNumCode] = "阅读次数不得小于0"
	lang.MsgInfo[PluginsAnnouncementTitleHasUsedCode] = "公告标题已存在，请重新输入"

	lang.MsgInfo[PluginsCategoryNameCode] = "分类名称不得为空"
	lang.MsgInfo[PluginsCategoryNameHasUsedCode] = "该分类名称已被使用"
	lang.MsgInfo[PluginsCategoryNotFoundCode] = "分类不存在，请检查"

	lang.MsgInfo[PluginsArticleNameCode] = "文章名称不得为空"
	lang.MsgInfo[PluginsArticleNameHasUsedCode] = "文章名称已被使用"
	lang.MsgInfo[PluginsArticleContentCode] = "文章内容不得为空"
	lang.MsgInfo[PluginsArticleCatIdEmptyCode] = "文章分类编号错误"
}
