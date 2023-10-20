package lang

import (
	"go-admin/core/lang"
)

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

var (
	MsgInfo = map[int]string{
		PluginsAnnouncementTitleEmptyCode:   "公告标题不得为空",
		PluginsAnnouncementContentEmptyCode: "公告内容不得为空",
		PluginsAnnouncementStatusEmptyCode:  "公告状态不得为空",
		PluginsAnnouncementNumCode:          "阅读次数不得小于0",
		PluginsAnnouncementTitleHasUsedCode: "公告标题已存在，请重新输入",

		PluginsCategoryNameCode:        "分类名称不得为空",
		PluginsCategoryNameHasUsedCode: "该分类名称已被使用",
		PluginsCategoryNotFoundCode:    "分类不存在，请检查",

		PluginsArticleNameCode:        "文章名称不得为空",
		PluginsArticleNameHasUsedCode: "文章名称已被使用",
		PluginsArticleContentCode:     "文章内容不得为空",
		PluginsArticleCatIdEmptyCode:  "文章分类编号错误",
	}
)

// Init 初始化
func Init() {
	for k, v := range MsgInfo {
		if lang.MsgInfo[k] == "" {
			lang.MsgInfo[k] = v
		} else {
			panic("Your plugin lang code %d is used by system or other plugins,please check")
		}
	}
}
