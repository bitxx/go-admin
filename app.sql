/*
 Navicat Premium Dump SQL

 Source Server         : sqlpub
 Source Server Type    : MySQL
 Source Server Version : 80040 (8.0.40)
 Source Host           : mysql.sqlpub.com:3306
 Source Schema         : bitxxadmin

 Target Server Type    : MySQL
 Target Server Version : 80040 (8.0.40)
 File Encoding         : 65001

 Date: 13/12/2024 19:27:38
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for admin_sys_api
-- ----------------------------
DROP TABLE IF EXISTS `admin_sys_api`;
CREATE TABLE `admin_sys_api` (
  `id` int NOT NULL AUTO_INCREMENT COMMENT '主键编码',
  `description` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '功能描述',
  `path` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '地址',
  `api_type` varchar(16) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '接口类型',
  `method` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '请求类型',
  `created_at` datetime DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime DEFAULT NULL COMMENT '最后更新时间',
  `create_by` int DEFAULT NULL COMMENT '创建者',
  `update_by` int DEFAULT NULL COMMENT '更新者',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;

-- ----------------------------
-- Records of admin_sys_api
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for admin_sys_casbin_rule
-- ----------------------------
DROP TABLE IF EXISTS `admin_sys_casbin_rule`;
CREATE TABLE `admin_sys_casbin_rule` (
  `p_type` varchar(100) DEFAULT NULL,
  `v0` varchar(100) DEFAULT NULL,
  `v1` varchar(100) DEFAULT NULL,
  `v2` varchar(100) DEFAULT NULL,
  `v3` varchar(100) DEFAULT NULL,
  `v4` varchar(100) DEFAULT NULL,
  `v5` varchar(100) DEFAULT NULL,
  UNIQUE KEY `idx_admin_sys_casbin_rule` (`p_type`,`v0`,`v1`,`v2`,`v3`,`v4`,`v5`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Records of admin_sys_casbin_rule
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for admin_sys_config
-- ----------------------------
DROP TABLE IF EXISTS `admin_sys_config`;
CREATE TABLE `admin_sys_config` (
  `id` int NOT NULL AUTO_INCREMENT COMMENT '主键编码',
  `config_name` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT 'ConfigName',
  `config_key` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT 'ConfigKey',
  `config_value` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT 'ConfigValue',
  `config_type` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT 'ConfigType',
  `is_frontend` char(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '是否前台',
  `remark` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT 'Remark',
  `create_by` int DEFAULT NULL COMMENT '创建者',
  `update_by` int DEFAULT NULL COMMENT '更新者',
  `created_at` datetime DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime DEFAULT NULL COMMENT '最后更新时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=13 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;

-- ----------------------------
-- Records of admin_sys_config
-- ----------------------------
BEGIN;
INSERT INTO `admin_sys_config` (`id`, `config_name`, `config_key`, `config_value`, `config_type`, `is_frontend`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (1, '管理-皮肤样式', 'admin_sys_index_skinName', 'skin-green', '1', '1', '主框架页-默认皮肤样式名称:蓝色 skin-blue、绿色 skin-green、紫色 skin-purple、红色 skin-red、黄色 skin-yellow', 1, 1, '2021-05-13 19:56:38', '2023-03-11 23:16:02');
INSERT INTO `admin_sys_config` (`id`, `config_name`, `config_key`, `config_value`, `config_type`, `is_frontend`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (2, '管理-初始密码', 'admin_sys_user_initPassword', '123456', '1', '1', '用户管理-账号初始密码:123456', 1, 1, '2021-05-13 19:56:38', '2023-03-11 23:16:10');
INSERT INTO `admin_sys_config` (`id`, `config_name`, `config_key`, `config_value`, `config_type`, `is_frontend`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (3, '管理-侧栏主题', 'admin_sys_index_sideTheme', 'theme-dark', '1', '1', '主框架页-侧边栏主题:深色主题theme-dark，浅色主题theme-light', 1, 1, '2021-05-13 19:56:38', '2023-03-11 23:16:06');
INSERT INTO `admin_sys_config` (`id`, `config_name`, `config_key`, `config_value`, `config_type`, `is_frontend`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (4, '管理-系统名称', 'admin_sys_app_name', 'go-admin后台管理系统', '1', '1', '', 1, 1, '2021-03-17 08:52:06', '2023-03-11 23:16:19');
INSERT INTO `admin_sys_config` (`id`, `config_name`, `config_key`, `config_value`, `config_type`, `is_frontend`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (5, '管理-系统logo', 'admin_sys_app_logo', 'http://www.bitxx.top/images/my_head-touch-icon-next.png', '1', '1', '', 1, 1, '2021-03-17 08:53:19', '2023-03-11 23:16:15');
INSERT INTO `admin_sys_config` (`id`, `config_name`, `config_key`, `config_value`, `config_type`, `is_frontend`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (6, '管理-单次excel导出数据量', 'admin_sys_max_export_size', '10000', '1', '1', '', 0, 1, '2021-07-28 16:53:48', '2023-03-11 23:15:56');
INSERT INTO `admin_sys_config` (`id`, `config_name`, `config_key`, `config_value`, `config_type`, `is_frontend`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (7, '插件-文件管理-App OSS Bucket', 'plugin_filemgr_app_oss_bucket', '请自行配置', '2', '2', '', 0, 1, '2021-08-13 14:36:23', '2023-03-11 23:14:45');
INSERT INTO `admin_sys_config` (`id`, `config_name`, `config_key`, `config_value`, `config_type`, `is_frontend`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (8, '插件-文件管理-App OSS AccessKeyId', 'plugin_filemgr_app_oss_access_key_id', '请自行配置', '2', '2', '', 0, 1, '2021-08-13 14:37:15', '2023-03-11 23:14:41');
INSERT INTO `admin_sys_config` (`id`, `config_name`, `config_key`, `config_value`, `config_type`, `is_frontend`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (9, '插件-文件管理-App OSS AccessKeySecret', 'plugin_filemgr_app_oss_access_key_secret', '请自行配置', '2', '2', '', 0, 1, '2021-08-13 14:38:00', '2023-03-11 23:14:33');
INSERT INTO `admin_sys_config` (`id`, `config_name`, `config_key`, `config_value`, `config_type`, `is_frontend`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (10, '插件-文件管理-App OSS Endpoint', 'plugin_filemgr_app_oss_endpoint', '请自行配置', '2', '2', '', 0, 1, '2021-08-13 14:38:50', '2023-03-11 23:14:28');
INSERT INTO `admin_sys_config` (`id`, `config_name`, `config_key`, `config_value`, `config_type`, `is_frontend`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (11, '插件-文件管理-App OSS 根目录', 'plugin_filemgr_app_oss_root_path', 'testfile/', '2', '2', '', 0, 1, '2021-08-13 14:39:31', '2023-03-11 23:14:22');
INSERT INTO `admin_sys_config` (`id`, `config_name`, `config_key`, `config_value`, `config_type`, `is_frontend`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (12, '管理-用户-默认头像', 'admin_sys_user_default_avatar', 'http://www.bitxx.top/images/my_head-touch-icon-next.png', '3', '2', '', 1, 1, '2023-03-10 18:07:03', '2023-03-10 18:07:03');
COMMIT;

-- ----------------------------
-- Table structure for admin_sys_dept
-- ----------------------------
DROP TABLE IF EXISTS `admin_sys_dept`;
CREATE TABLE `admin_sys_dept` (
  `id` int NOT NULL AUTO_INCREMENT,
  `parent_id` int DEFAULT NULL,
  `parent_ids` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL,
  `dept_name` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL,
  `sort` int DEFAULT NULL,
  `leader` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL,
  `phone` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL,
  `email` varchar(80) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL,
  `status` tinyint DEFAULT NULL,
  `create_by` int DEFAULT NULL COMMENT '创建者',
  `update_by` int DEFAULT NULL COMMENT '更新者',
  `created_at` datetime DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime DEFAULT NULL COMMENT '最后更新时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;

-- ----------------------------
-- Records of admin_sys_dept
-- ----------------------------
BEGIN;
INSERT INTO `admin_sys_dept` (`id`, `parent_id`, `parent_ids`, `dept_name`, `sort`, `leader`, `phone`, `email`, `status`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (1, 0, '0,', 'Admin', 0, 'admin', '', '', 1, 1, 1, '2021-05-13 19:56:38', '2022-05-14 11:20:25');
INSERT INTO `admin_sys_dept` (`id`, `parent_id`, `parent_ids`, `dept_name`, `sort`, `leader`, `phone`, `email`, `status`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (2, 1, '0,1,', '研发部', 1, 'admin', '', '', 1, 1, 1, '2021-05-13 19:56:38', '2023-03-04 13:17:45');
INSERT INTO `admin_sys_dept` (`id`, `parent_id`, `parent_ids`, `dept_name`, `sort`, `leader`, `phone`, `email`, `status`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (3, 1, '0,1,', '运维部', 1, 'admin', '', '', 1, 1, 1, '2021-05-13 19:56:38', '2023-03-04 13:17:45');
INSERT INTO `admin_sys_dept` (`id`, `parent_id`, `parent_ids`, `dept_name`, `sort`, `leader`, `phone`, `email`, `status`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (4, 1, '0,1,', '客服部', 0, 'admin', '', '', 1, 1, 1, '2021-05-13 19:56:38', '2022-05-14 11:20:50');
INSERT INTO `admin_sys_dept` (`id`, `parent_id`, `parent_ids`, `dept_name`, `sort`, `leader`, `phone`, `email`, `status`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (5, 1, '0,1,', '人力资源', 3, 'admin', '', '', 1, 1, 1, '2021-05-13 19:56:38', '2022-05-14 11:20:53');
INSERT INTO `admin_sys_dept` (`id`, `parent_id`, `parent_ids`, `dept_name`, `sort`, `leader`, `phone`, `email`, `status`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (6, 1, '0,1,', '市场', 10, 'admin', '', '', 1, 1, 1, '2021-12-02 10:13:38', '2021-12-02 10:13:38');
COMMIT;

-- ----------------------------
-- Table structure for admin_sys_dict_data
-- ----------------------------
DROP TABLE IF EXISTS `admin_sys_dict_data`;
CREATE TABLE `admin_sys_dict_data` (
  `id` int NOT NULL AUTO_INCREMENT,
  `dict_sort` int DEFAULT NULL,
  `dict_label` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL,
  `dict_value` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL,
  `dict_type` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL,
  `css_class` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL,
  `list_class` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL,
  `is_default` varchar(8) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL,
  `status` char(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL,
  `default` varchar(8) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL,
  `remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL,
  `create_by` int DEFAULT NULL COMMENT '创建者',
  `update_by` int DEFAULT NULL COMMENT '更新者',
  `created_at` datetime DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime DEFAULT NULL COMMENT '最后更新时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=99 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;

-- ----------------------------
-- Records of admin_sys_dict_data
-- ----------------------------
BEGIN;
INSERT INTO `admin_sys_dict_data` (`id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `default`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (1, 0, '正常', '2', 'admin_sys_normal_disable', '', '', '', '0', '', '系统正常', 1, 1, '2021-05-13 19:56:38', '2022-04-25 00:42:38');
INSERT INTO `admin_sys_dict_data` (`id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `default`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (2, 0, '停用', '1', 'admin_sys_normal_disable', '', '', '', '0', '', '系统停用', 1, 1, '2021-05-13 19:56:38', '2021-05-13 19:56:38');
INSERT INTO `admin_sys_dict_data` (`id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `default`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (3, 0, '男', '1', 'admin_sys_user_sex', '', '', '', '0', '', '性别男', 1, 1, '2021-05-13 19:56:38', '2021-05-13 19:56:38');
INSERT INTO `admin_sys_dict_data` (`id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `default`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (4, 0, '女', '2', 'admin_sys_user_sex', '', '', '', '0', '', '性别女', 1, 1, '2021-05-13 19:56:38', '2021-05-13 19:56:38');
INSERT INTO `admin_sys_dict_data` (`id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `default`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (5, 0, '未知', '3', 'admin_sys_user_sex', '', '', '', '0', '', '性别未知', 1, 1, '2021-05-13 19:56:38', '2023-03-05 12:03:33');
INSERT INTO `admin_sys_dict_data` (`id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `default`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (6, 0, '显示', '2', 'admin_sys_menu_show_hide', '', '', '', '0', '', '显示菜单', 1, 1, '2021-05-13 19:56:38', '2021-05-13 19:56:38');
INSERT INTO `admin_sys_dict_data` (`id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `default`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (7, 0, '隐藏', '1', 'admin_sys_menu_show_hide', '', '', '', '0', '', '隐藏菜单', 1, 1, '2021-05-13 19:56:38', '2021-05-13 19:56:38');
INSERT INTO `admin_sys_dict_data` (`id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `default`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (8, 0, '是', '1', 'admin_sys_yes_no', '', '', '', '0', '', '系统默认是', 1, 1, '2021-05-13 19:56:38', '2021-05-13 19:56:38');
INSERT INTO `admin_sys_dict_data` (`id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `default`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (9, 0, '否', '2', 'admin_sys_yes_no', '', '', '', '0', '', '系统默认否', 1, 1, '2021-05-13 19:56:38', '2021-05-13 19:56:38');
INSERT INTO `admin_sys_dict_data` (`id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `default`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (10, 0, '通知', '1', 'admin_sys_notice_type', '', '', '', '0', '', '通知', 1, 1, '2021-05-13 19:56:38', '2021-05-13 19:56:38');
INSERT INTO `admin_sys_dict_data` (`id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `default`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (11, 0, '公告', '2', 'admin_sys_notice_type', '', '', '', '0', '', '公告', 1, 1, '2021-05-13 19:56:38', '2021-05-13 19:56:38');
INSERT INTO `admin_sys_dict_data` (`id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `default`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (12, 0, '正常', '2', 'admin_sys_common_status', '', '', '', '0', '', '正常状态', 1, 1, '2021-05-13 19:56:38', '2021-05-13 19:56:38');
INSERT INTO `admin_sys_dict_data` (`id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `default`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (13, 0, '关闭', '1', 'admin_sys_common_status', '', '', '', '0', '', '关闭状态', 1, 1, '2021-05-13 19:56:38', '2021-05-13 19:56:38');
INSERT INTO `admin_sys_dict_data` (`id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `default`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (14, 0, '新增', '1', 'admin_sys_oper_type', '', '', '', '0', '', '新增操作', 1, 1, '2021-05-13 19:56:38', '2021-05-13 19:56:38');
INSERT INTO `admin_sys_dict_data` (`id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `default`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (15, 0, '修改', '2', 'admin_sys_oper_type', '', '', '', '0', '', '修改操作', 1, 1, '2021-05-13 19:56:38', '2021-05-13 19:56:38');
INSERT INTO `admin_sys_dict_data` (`id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `default`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (16, 0, '删除', '3', 'admin_sys_oper_type', '', '', '', '0', '', '删除操作', 1, 1, '2021-05-13 19:56:38', '2021-05-13 19:56:38');
INSERT INTO `admin_sys_dict_data` (`id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `default`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (17, 0, '授权', '4', 'admin_sys_oper_type', '', '', '', '0', '', '授权操作', 1, 1, '2021-05-13 19:56:38', '2021-05-13 19:56:38');
INSERT INTO `admin_sys_dict_data` (`id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `default`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (18, 0, '导出', '5', 'admin_sys_oper_type', '', '', '', '0', '', '导出操作', 1, 1, '2021-05-13 19:56:38', '2021-05-13 19:56:38');
INSERT INTO `admin_sys_dict_data` (`id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `default`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (19, 0, '导入', '6', 'admin_sys_oper_type', '', '', '', '0', '', '导入操作', 1, 1, '2021-05-13 19:56:38', '2021-05-13 19:56:38');
INSERT INTO `admin_sys_dict_data` (`id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `default`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (20, 0, '强退', '7', 'admin_sys_oper_type', '', '', '', '0', '', '强退操作', 1, 1, '2021-05-13 19:56:38', '2021-05-13 19:56:38');
INSERT INTO `admin_sys_dict_data` (`id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `default`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (21, 0, '生成代码', '8', 'admin_sys_oper_type', '', '', '', '0', '', '生成操作', 1, 1, '2021-05-13 19:56:38', '2021-05-13 19:56:38');
INSERT INTO `admin_sys_dict_data` (`id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `default`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (22, 0, '清空数据', '9', 'admin_sys_oper_type', '', '', '', '0', '', '清空操作', 1, 1, '2021-05-13 19:56:38', '2021-05-13 19:56:38');
INSERT INTO `admin_sys_dict_data` (`id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `default`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (23, 0, '成功', '1', 'admin_sys_notice_status', '', '', '', '0', '', '成功状态', 1, 1, '2021-05-13 19:56:38', '2021-05-13 19:56:38');
INSERT INTO `admin_sys_dict_data` (`id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `default`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (24, 0, '失败', '2', 'admin_sys_notice_status', '', '', '', '0', '', '失败状态', 1, 1, '2021-05-13 19:56:38', '2021-05-13 19:56:38');
INSERT INTO `admin_sys_dict_data` (`id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `default`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (25, 0, '登录', '10', 'admin_sys_oper_type', '', '', '', '0', '', '登录操作', 1, 1, '2021-05-13 19:56:38', '2021-05-13 19:56:38');
INSERT INTO `admin_sys_dict_data` (`id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `default`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (26, 0, '退出', '11', 'admin_sys_oper_type', '', '', '', '0', '', '', 1, 1, '2021-05-13 19:56:38', '2021-05-13 19:56:38');
INSERT INTO `admin_sys_dict_data` (`id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `default`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (27, 0, '获取验证码', '12', 'admin_sys_oper_type', '', '', '', '0', '', '获取验证码', 1, 1, '2021-05-13 19:56:38', '2021-05-13 19:56:38');
INSERT INTO `admin_sys_dict_data` (`id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `default`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (28, 0, '正常', '1', 'admin_sys_status', '', '', '', '0', '', '', 0, 0, '2021-07-09 11:40:01', '2021-07-09 11:40:01');
INSERT INTO `admin_sys_dict_data` (`id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `default`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (29, 0, '停用', '2', 'admin_sys_status', '', '', '', '0', '', '', 0, 0, '2021-07-09 11:40:14', '2021-07-09 11:40:14');
INSERT INTO `admin_sys_dict_data` (`id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `default`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (30, 0, '安卓', '1', 'plugin_filemgr_app_platform', '', '', '', '2', '', '', 0, 0, '2021-08-13 13:35:39', '2021-08-13 13:35:39');
INSERT INTO `admin_sys_dict_data` (`id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `default`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (31, 0, 'IOS', '2', 'plugin_filemgr_app_platform', '', '', '', '2', '', '', 0, 0, '2021-08-13 13:35:51', '2021-08-13 13:35:51');
INSERT INTO `admin_sys_dict_data` (`id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `default`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (32, 0, '类型1', '1', 'plugin_filemgr_app_type', '', '', '', '2', '', '', 0, 0, '2021-08-13 13:37:07', '2021-08-13 13:37:07');
INSERT INTO `admin_sys_dict_data` (`id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `default`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (33, 0, '类型2', '2', 'plugin_filemgr_app_type', '', '', '', '2', '', '', 0, 0, '2021-08-13 13:37:19', '2021-08-13 13:37:19');
INSERT INTO `admin_sys_dict_data` (`id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `default`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (34, 0, '类型3', '3', 'plugin_filemgr_app_type', '', '', '', '2', '', '', 0, 0, '2021-08-13 13:37:39', '2021-08-13 13:37:39');
INSERT INTO `admin_sys_dict_data` (`id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `default`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (35, 0, '本地', '1', 'plugin_filemgr_app_download_type', '', '', '', '2', '', '', 0, 0, '2021-08-13 14:02:44', '2021-08-13 14:02:44');
INSERT INTO `admin_sys_dict_data` (`id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `default`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (36, 0, '外链', '2', 'plugin_filemgr_app_download_type', '', '', '', '2', '', '', 0, 0, '2021-08-13 14:02:44', '2021-08-13 14:02:44');
INSERT INTO `admin_sys_dict_data` (`id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `default`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (37, 0, 'OSS', '3', 'plugin_filemgr_app_download_type', '', '', '', '2', '', '', 0, 0, '2021-08-13 14:02:33', '2021-08-13 14:02:33');
INSERT INTO `admin_sys_dict_data` (`id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `default`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (38, 0, '已发布', '2', 'plugin_filemgr_app_publish_status', '', '', '', '2', '', '', 0, 0, '2021-12-09 12:42:47', '2021-12-09 12:42:47');
INSERT INTO `admin_sys_dict_data` (`id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `default`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (39, 0, '待发布', '1', 'plugin_filemgr_app_publish_status', '', '', '', '2', '', '', 0, 0, '2021-12-09 12:42:54', '2021-12-09 12:42:54');
INSERT INTO `admin_sys_dict_data` (`id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `default`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (40, 0, '插件', '2', 'admin_sys_api_type', '', '', '', '0', '', '', 1, 1, '2022-04-25 23:58:24', '2023-03-01 21:45:53');
INSERT INTO `admin_sys_dict_data` (`id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `default`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (41, 0, '管理', '1', 'admin_sys_api_type', '', '', '', '0', '', '', 1, 1, '2022-04-25 23:58:41', '2023-03-01 21:45:41');
INSERT INTO `admin_sys_dict_data` (`id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `default`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (42, 0, 'GET', 'GET', 'admin_sys_api_method', '', '', '', '0', '', '', 1, 1, '2022-04-26 00:03:26', '2022-04-26 00:03:26');
INSERT INTO `admin_sys_dict_data` (`id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `default`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (43, 0, 'POST', 'POST', 'admin_sys_api_method', '', '', '', '0', '', '', 1, 1, '2022-04-26 00:03:40', '2022-04-26 00:03:40');
INSERT INTO `admin_sys_dict_data` (`id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `default`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (44, 0, 'DELETE', 'DELETE', 'admin_sys_api_method', '', '', '', '0', '', '', 1, 1, '2022-04-26 00:03:49', '2022-04-26 00:03:49');
INSERT INTO `admin_sys_dict_data` (`id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `default`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (45, 0, 'PUT', 'PUT', 'admin_sys_api_method', '', '', '', '0', '', '', 1, 1, '2022-04-26 00:04:06', '2022-04-26 00:04:06');
INSERT INTO `admin_sys_dict_data` (`id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `default`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (46, 0, 'HEAD', 'HEAD', 'admin_sys_api_method', '', '', '', '0', '', '', 1, 1, '2022-04-26 00:07:02', '2022-04-26 00:07:02');
INSERT INTO `admin_sys_dict_data` (`id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `default`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (47, 0, '系统内置', '1', 'admin_sys_config_type', '', '', '', '1', '', '', 1, 1, '2023-03-01 11:05:23', '2023-03-01 11:05:23');
INSERT INTO `admin_sys_dict_data` (`id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `default`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (48, 0, '插件', '2', 'admin_sys_config_type', '', '', '', '1', '', '', 1, 1, '2023-03-01 11:05:32', '2023-03-01 11:05:32');
INSERT INTO `admin_sys_dict_data` (`id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `default`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (49, 0, '应用', '3', 'admin_sys_config_type', '', '', '', '1', '', '', 1, 1, '2023-03-01 11:05:42', '2023-03-01 11:05:42');
INSERT INTO `admin_sys_dict_data` (`id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `default`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (50, 0, '展示', '1', 'admin_sys_config_is_frontend', '', '', '', '1', '', '', 1, 1, '2023-03-01 11:07:49', '2023-03-01 11:07:49');
INSERT INTO `admin_sys_dict_data` (`id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `default`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (51, 0, '隐藏', '2', 'admin_sys_config_is_frontend', '', '', '', '1', '', '', 1, 1, '2023-03-01 11:07:56', '2023-03-01 11:07:56');
INSERT INTO `admin_sys_dict_data` (`id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `default`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (52, 0, '登录', '1', 'admin_sys_loginlog_status', '', '', '', '1', '', '', 1, 1, '2023-03-01 14:43:04', '2023-03-01 14:43:04');
INSERT INTO `admin_sys_dict_data` (`id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `default`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (53, 0, '退出', '2', 'admin_sys_loginlog_status', '', '', '', '1', '', '', 1, 1, '2023-03-01 14:43:10', '2023-03-01 14:43:10');
INSERT INTO `admin_sys_dict_data` (`id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `default`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (54, 0, '应用', '3', 'admin_sys_api_type', '', '', '', '1', '', '', 1, 1, '2023-03-01 21:46:01', '2023-03-01 21:46:01');
INSERT INTO `admin_sys_dict_data` (`id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `default`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (55, 0, '全部数据权限', '1', 'admin_sys_role_data_scope', '', '', '', '1', '', '', 1, 1, '2023-03-04 13:29:36', '2023-03-04 13:29:36');
INSERT INTO `admin_sys_dict_data` (`id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `default`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (56, 0, '自定数据权限', '2', 'admin_sys_role_data_scope', '', '', '', '1', '', '', 1, 1, '2023-03-04 13:29:43', '2023-03-04 13:29:43');
INSERT INTO `admin_sys_dict_data` (`id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `default`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (57, 0, '本部门数据权限', '3', 'admin_sys_role_data_scope', '', '', '', '1', '', '', 1, 1, '2023-03-04 13:29:49', '2023-03-04 13:29:49');
INSERT INTO `admin_sys_dict_data` (`id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `default`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (58, 0, '本部门及以下数据权限', '4', 'admin_sys_role_data_scope', '', '', '', '1', '', '', 1, 1, '2023-03-04 13:29:56', '2023-03-04 13:29:56');
INSERT INTO `admin_sys_dict_data` (`id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `default`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (59, 0, '仅本人数据权限', '5', 'admin_sys_role_data_scope', '', '', '', '1', '', '', 1, 1, '2023-03-04 13:30:04', '2023-03-04 13:30:04');
INSERT INTO `admin_sys_dict_data` (`id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `default`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (60, 0, 'int64', 'int64', 'admin_sys_gen_go_type', '', '', '', '1', '', '', 1, 1, '2023-03-07 10:08:26', '2023-03-07 10:08:26');
INSERT INTO `admin_sys_dict_data` (`id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `default`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (61, 0, 'int', 'int', 'admin_sys_gen_go_type', '', '', '', '1', '', '', 1, 1, '2023-03-07 10:12:42', '2023-03-07 10:12:42');
INSERT INTO `admin_sys_dict_data` (`id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `default`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (62, 0, 'string', 'string', 'admin_sys_gen_go_type', '', '', '', '1', '', '', 1, 1, '2023-03-07 10:13:05', '2023-03-07 10:13:05');
INSERT INTO `admin_sys_dict_data` (`id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `default`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (63, 0, 'decimal', 'decimal.Decimal', 'admin_sys_gen_go_type', '', '', '', '1', '', '', 1, 1, '2023-03-07 10:13:16', '2023-03-07 10:13:29');
INSERT INTO `admin_sys_dict_data` (`id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `default`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (64, 0, 'time', '*time.Time', 'admin_sys_gen_go_type', '', '', '', '1', '', '', 1, 1, '2023-03-07 10:13:43', '2023-03-07 10:13:43');
INSERT INTO `admin_sys_dict_data` (`id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `default`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (65, 0, '=', 'EQ', 'admin_sys_gen_query_type', '', '', '', '1', '', '', 1, 1, '2023-03-07 10:20:53', '2023-03-07 10:20:53');
INSERT INTO `admin_sys_dict_data` (`id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `default`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (66, 0, '!=', 'NE', 'admin_sys_gen_query_type', '', '', '', '1', '', '', 1, 1, '2023-03-07 10:21:06', '2023-03-07 10:21:06');
INSERT INTO `admin_sys_dict_data` (`id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `default`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (67, 0, '>', 'GT', 'admin_sys_gen_query_type', '', '', '', '1', '', '', 1, 1, '2023-03-07 10:21:20', '2023-03-07 10:21:20');
INSERT INTO `admin_sys_dict_data` (`id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `default`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (68, 0, '>=', 'GTE', 'admin_sys_gen_query_type', '', '', '', '1', '', '', 1, 1, '2023-03-07 10:21:33', '2023-03-07 10:21:33');
INSERT INTO `admin_sys_dict_data` (`id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `default`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (69, 0, '<', 'LT', 'admin_sys_gen_query_type', '', '', '', '1', '', '', 1, 1, '2023-03-07 10:21:45', '2023-03-07 10:21:45');
INSERT INTO `admin_sys_dict_data` (`id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `default`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (70, 0, '<=', 'LTE', 'admin_sys_gen_query_type', '', '', '', '1', '', '', 1, 1, '2023-03-07 10:21:57', '2023-03-07 10:21:57');
INSERT INTO `admin_sys_dict_data` (`id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `default`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (71, 0, 'LIKE', 'LIKE', 'admin_sys_gen_query_type', '', '', '', '1', '', '', 1, 1, '2023-03-07 10:22:08', '2023-03-07 10:22:08');
INSERT INTO `admin_sys_dict_data` (`id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `default`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (72, 0, '文本框', 'input', 'admin_sys_gen_html_type', '', '', '', '1', '', '', 1, 1, '2023-03-07 10:23:39', '2023-03-07 10:23:39');
INSERT INTO `admin_sys_dict_data` (`id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `default`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (73, 0, '下拉框', 'select', 'admin_sys_gen_html_type', '', '', '', '1', '', '', 1, 1, '2023-03-07 10:23:49', '2023-03-07 10:23:49');
INSERT INTO `admin_sys_dict_data` (`id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `default`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (74, 0, '单选框', 'radio', 'admin_sys_gen_html_type', '', '', '', '1', '', '', 1, 1, '2023-03-07 10:23:59', '2023-03-07 10:23:59');
INSERT INTO `admin_sys_dict_data` (`id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `default`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (75, 0, '文本域', 'textarea', 'admin_sys_gen_html_type', '', '', '', '1', '', '', 1, 1, '2023-03-07 10:24:08', '2023-03-07 10:24:08');
INSERT INTO `admin_sys_dict_data` (`id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `default`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (76, 0, '目录', '1', 'admin_sys_menu_type', '', '', '', '1', '', '', 1, 1, '2023-03-08 10:42:00', '2023-03-08 10:42:14');
INSERT INTO `admin_sys_dict_data` (`id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `default`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (77, 0, '菜单', '2', 'admin_sys_menu_type', '', '', '', '1', '', '', 1, 1, '2023-03-08 10:42:10', '2023-03-08 10:42:10');
INSERT INTO `admin_sys_dict_data` (`id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `default`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (78, 0, '按钮', '3', 'admin_sys_menu_type', '', '', '', '1', '', '', 1, 1, '2023-03-08 10:42:22', '2023-03-08 10:42:22');
INSERT INTO `admin_sys_dict_data` (`id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `default`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (79, 0, '类型1', '1', 'app_user_level_type', '', '', '', '1', '', '', 1, 1, '2023-03-08 11:55:57', '2023-03-08 11:55:57');
INSERT INTO `admin_sys_dict_data` (`id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `default`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (80, 0, '类型2', '2', 'app_user_level_type', '', '', '', '1', '', '', 1, 1, '2023-03-08 11:56:02', '2023-03-08 11:56:02');
INSERT INTO `admin_sys_dict_data` (`id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `default`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (81, 0, '数字文本框', 'numInput', 'admin_sys_gen_html_type', '', '', '', '1', '', '', 1, 1, '2023-03-09 20:12:33', '2023-03-09 20:12:33');
INSERT INTO `admin_sys_dict_data` (`id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `default`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (82, 0, 'CNY', '1', 'app_money_type', '', '', '', '1', '', '', 1, 1, '2023-03-09 20:24:26', '2023-03-09 20:24:26');
INSERT INTO `admin_sys_dict_data` (`id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `default`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (83, 0, '类型1', '1', 'app_account_change_type', '', '', '', '1', '', '', 1, 1, '2023-03-09 20:27:45', '2023-03-09 20:27:45');
INSERT INTO `admin_sys_dict_data` (`id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `default`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (84, 0, '允许用户登录', '1', 'app_user_action_type', '', '', '', '1', '', '', 1, 1, '2023-03-11 14:08:01', '2023-03-11 14:08:01');
INSERT INTO `admin_sys_dict_data` (`id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `default`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (85, 0, '禁止用户登录', '2', 'app_user_action_type', '', '', '', '1', '', '', 1, 1, '2023-03-11 14:08:10', '2023-03-11 14:08:10');
INSERT INTO `admin_sys_dict_data` (`id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `default`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (86, 0, '后台用户', '1', 'app_user_by_type', '', '', '', '1', '', '', 1, 1, '2023-03-11 14:14:41', '2023-03-11 14:14:41');
INSERT INTO `admin_sys_dict_data` (`id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `default`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (87, 0, '前台用户', '2', 'app_user_by_type', '', '', '', '1', '', '', 1, 1, '2023-03-11 14:14:59', '2023-03-11 14:14:59');
INSERT INTO `admin_sys_dict_data` (`id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `default`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (88, 0, '发送成功', '1', 'plugin_msg_sendstatus', '', '', '', '1', '', '', 1, 1, '2023-09-26 10:42:22', '2023-09-26 10:42:22');
INSERT INTO `admin_sys_dict_data` (`id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `default`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (89, 0, '发送失败', '2', 'plugin_msg_sendstatus', '', '', '', '1', '', '', 1, 1, '2023-09-26 10:42:31', '2023-09-26 10:42:31');
INSERT INTO `admin_sys_dict_data` (`id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `default`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (90, 0, '邮箱', '1', 'plugin_msg_code_type', '', '', '', '1', '', '', 1, 1, '2023-09-26 10:42:58', '2023-09-26 10:42:58');
INSERT INTO `admin_sys_dict_data` (`id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `default`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (91, 0, '短信', '2', 'plugin_msg_code_type', '', '', '', '1', '', '', 1, 1, '2023-09-26 10:43:04', '2023-09-26 10:43:04');
INSERT INTO `admin_sys_dict_data` (`id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `default`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (92, 0, '待发布', '1', 'plugin_filemgr_publish_status', '', '', '', '1', '', '', 1, 1, '2024-12-01 23:20:36', '2024-12-01 23:20:36');
INSERT INTO `admin_sys_dict_data` (`id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `default`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (93, 0, '已发布', '2', 'plugin_filemgr_publish_status', '', '', '', '1', '', '', 1, 1, '2024-12-01 23:20:45', '2024-12-01 23:20:45');
INSERT INTO `admin_sys_dict_data` (`id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `default`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (94, 0, '数字文本框', 'numInput', 'admin_sys_gen_html_type', '', '', '', '1', '', '', 1, 1, '2023-03-07 10:23:39', '2023-03-07 10:23:39');
INSERT INTO `admin_sys_dict_data` (`id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `default`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (95, 0, '未同步', '0', 'admin_sys_api_sync_status', '', '', '', '1', '', '启动程序初始化值', 1, 1, '2024-12-13 00:30:24', '2024-12-13 00:30:24');
INSERT INTO `admin_sys_dict_data` (`id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `default`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (96, 0, '上次同步成功', '1', 'admin_sys_api_sync_status', '', '', '', '1', '', '每次同步正常完毕都是代表上次同步成功', 1, 1, '2024-12-13 00:30:43', '2024-12-13 00:30:43');
INSERT INTO `admin_sys_dict_data` (`id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `default`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (97, 0, '自动同步中', '2', 'admin_sys_api_sync_status', '', '', '', '1', '', '', 1, 1, '2024-12-13 00:30:57', '2024-12-13 00:30:57');
INSERT INTO `admin_sys_dict_data` (`id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `default`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (98, 0, '上次同步异常', '3', 'admin_sys_api_sync_status', '', '', '', '1', '', '', 1, 1, '2024-12-13 00:31:08', '2024-12-13 00:31:08');
COMMIT;

-- ----------------------------
-- Table structure for admin_sys_dict_type
-- ----------------------------
DROP TABLE IF EXISTS `admin_sys_dict_type`;
CREATE TABLE `admin_sys_dict_type` (
  `id` int NOT NULL AUTO_INCREMENT,
  `dict_name` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL,
  `dict_type` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL,
  `status` char(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL,
  `remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL,
  `create_by` int DEFAULT NULL COMMENT '创建者',
  `update_by` int DEFAULT NULL COMMENT '更新者',
  `created_at` datetime DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime DEFAULT NULL COMMENT '最后更新时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=34 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;

-- ----------------------------
-- Records of admin_sys_dict_type
-- ----------------------------
BEGIN;
INSERT INTO `admin_sys_dict_type` (`id`, `dict_name`, `dict_type`, `status`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (1, '管理-开关', 'admin_sys_normal_disable', '0', '系统开关列表', 1, 1, '2021-05-13 19:56:38', '2023-03-11 23:20:35');
INSERT INTO `admin_sys_dict_type` (`id`, `dict_name`, `dict_type`, `status`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (2, '管理-用户性别', 'admin_sys_user_sex', '0', '用户性别列表', 1, 1, '2021-05-13 19:56:38', '2023-03-11 23:21:06');
INSERT INTO `admin_sys_dict_type` (`id`, `dict_name`, `dict_type`, `status`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (3, '管理-菜单状态', 'admin_sys_menu_show_hide', '0', '菜单状态列表', 1, 1, '2021-05-13 19:56:38', '2023-03-11 23:21:02');
INSERT INTO `admin_sys_dict_type` (`id`, `dict_name`, `dict_type`, `status`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (4, '管理-是否', 'admin_sys_yes_no', '0', '系统是否列表', 1, 1, '2021-05-13 19:56:38', '2023-03-11 23:20:58');
INSERT INTO `admin_sys_dict_type` (`id`, `dict_name`, `dict_type`, `status`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (7, '管理-通知类型', 'admin_sys_notice_type', '0', '通知类型列表', 1, 1, '2021-05-13 19:56:38', '2023-03-11 23:20:53');
INSERT INTO `admin_sys_dict_type` (`id`, `dict_name`, `dict_type`, `status`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (8, '管理-状态', 'admin_sys_common_status', '0', '登录状态列表', 1, 1, '2021-05-13 19:56:38', '2023-03-11 23:20:49');
INSERT INTO `admin_sys_dict_type` (`id`, `dict_name`, `dict_type`, `status`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (9, '管理-操作类型', 'admin_sys_oper_type', '0', '操作类型列表', 1, 1, '2021-05-13 19:56:38', '2023-03-11 23:20:42');
INSERT INTO `admin_sys_dict_type` (`id`, `dict_name`, `dict_type`, `status`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (10, '管理-通知状态', 'admin_sys_notice_status', '0', '通知状态列表', 1, 1, '2021-05-13 19:56:38', '2023-03-11 23:20:39');
INSERT INTO `admin_sys_dict_type` (`id`, `dict_name`, `dict_type`, `status`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (11, '管理-基本状态', 'admin_sys_status', '0', '基本通用状态', 1, 1, '2021-07-09 11:39:21', '2023-03-11 23:21:23');
INSERT INTO `admin_sys_dict_type` (`id`, `dict_name`, `dict_type`, `status`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (12, '插件-文件管理-App发布状态', 'plugin_filemgr_publish_status', '2', '', 1, 1, '2021-12-09 12:42:31', '2023-03-11 23:20:01');
INSERT INTO `admin_sys_dict_type` (`id`, `dict_name`, `dict_type`, `status`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (13, '插件-文件管理-App系统平台', 'plugin_filemgr_app_platform', '0', 'App系统平台', 1, 1, '2021-08-13 13:36:40', '2023-03-11 23:20:17');
INSERT INTO `admin_sys_dict_type` (`id`, `dict_name`, `dict_type`, `status`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (14, '插件-文件管理-App类型', 'plugin_filemgr_app_type', '0', 'app属性', 1, 1, '2021-08-13 13:36:40', '2023-03-11 23:20:13');
INSERT INTO `admin_sys_dict_type` (`id`, `dict_name`, `dict_type`, `status`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (15, '插件-文件管理-App下载类型', 'plugin_filemgr_app_download_type', '0', '', 1, 1, '2021-08-13 14:02:03', '2023-03-11 23:20:06');
INSERT INTO `admin_sys_dict_type` (`id`, `dict_name`, `dict_type`, `status`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (16, '管理-接口-类型', 'admin_sys_api_type', '0', '系统', 1, 1, '2022-04-25 23:57:17', '2023-03-01 21:56:34');
INSERT INTO `admin_sys_dict_type` (`id`, `dict_name`, `dict_type`, `status`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (17, '管理-接口-请求方法', 'admin_sys_api_method', '0', '', 1, 1, '2022-04-26 00:03:11', '2023-03-01 21:56:41');
INSERT INTO `admin_sys_dict_type` (`id`, `dict_name`, `dict_type`, `status`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (18, '管理-配置-类型', 'admin_sys_config_type', '1', '1-内置 2-插件 3-应用', 1, 1, '2023-03-01 11:04:56', '2023-03-01 11:08:27');
INSERT INTO `admin_sys_dict_type` (`id`, `dict_name`, `dict_type`, `status`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (19, '管理-配置-是否前台展示', 'admin_sys_config_is_frontend', '1', '1-展示 2-隐藏', 1, 1, '2023-03-01 11:06:28', '2023-03-01 11:08:07');
INSERT INTO `admin_sys_dict_type` (`id`, `dict_name`, `dict_type`, `status`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (20, '管理-登录日志-日志状态', 'admin_sys_loginlog_status', '1', '1-登录 2-退出', 1, 1, '2023-03-01 14:42:56', '2023-03-01 14:42:56');
INSERT INTO `admin_sys_dict_type` (`id`, `dict_name`, `dict_type`, `status`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (21, '管理-角色-数据范围', 'admin_sys_role_data_scope', '1', '1-全部数据权限 2- 自定义数据权限 3-本部门数据权限 4-本部门及以下数据权限 5-仅本人数据权限', 1, 1, '2023-03-04 13:29:21', '2023-03-04 13:29:21');
INSERT INTO `admin_sys_dict_type` (`id`, `dict_name`, `dict_type`, `status`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (22, '管理-模板-go类型', 'admin_sys_gen_go_type', '1', '', 1, 1, '2023-03-07 10:08:07', '2023-03-07 10:08:07');
INSERT INTO `admin_sys_dict_type` (`id`, `dict_name`, `dict_type`, `status`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (23, '管理-模板-查询类型', 'admin_sys_gen_query_type', '1', '', 1, 1, '2023-03-07 10:20:19', '2023-03-07 10:20:19');
INSERT INTO `admin_sys_dict_type` (`id`, `dict_name`, `dict_type`, `status`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (24, '管理-模板-显示类型', 'admin_sys_gen_html_type', '1', '', 1, 1, '2023-03-07 10:23:23', '2023-03-07 10:23:23');
INSERT INTO `admin_sys_dict_type` (`id`, `dict_name`, `dict_type`, `status`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (25, '管理-菜单-类型', 'admin_sys_menu_type', '1', '', 1, 1, '2023-03-08 10:33:32', '2023-03-08 10:33:32');
INSERT INTO `admin_sys_dict_type` (`id`, `dict_name`, `dict_type`, `status`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (26, '应用-用户-等级', 'app_user_level_type', '1', '', 1, 1, '2023-03-08 11:44:48', '2023-03-08 11:44:48');
INSERT INTO `admin_sys_dict_type` (`id`, `dict_name`, `dict_type`, `status`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (27, '应用-用户-资产-资金类型', 'app_money_type', '1', '1-CNY', 1, 1, '2023-03-09 20:24:17', '2023-03-11 14:06:46');
INSERT INTO `admin_sys_dict_type` (`id`, `dict_name`, `dict_type`, `status`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (28, '应用-用户-资产-账变类型', 'app_account_change_type', '1', '1-类型1', 1, 1, '2023-03-09 20:27:33', '2023-03-11 14:06:38');
INSERT INTO `admin_sys_dict_type` (`id`, `dict_name`, `dict_type`, `status`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (29, '应用-用户-行为类型', 'app_user_action_type', '1', '', 1, 1, '2023-03-11 14:06:29', '2023-03-11 14:06:29');
INSERT INTO `admin_sys_dict_type` (`id`, `dict_name`, `dict_type`, `status`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (30, '应用-用户-用户更新类型', 'app_user_by_type', '1', '', 1, 1, '2023-03-11 14:14:06', '2023-03-11 14:14:27');
INSERT INTO `admin_sys_dict_type` (`id`, `dict_name`, `dict_type`, `status`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (31, '插件-消息-验证码类型', 'plugin_msg_code_type', '1', '1-邮箱 2-短信', 1, 1, '2023-03-12 12:12:30', '2023-03-12 12:15:20');
INSERT INTO `admin_sys_dict_type` (`id`, `dict_name`, `dict_type`, `status`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (32, '插件-消息-验证码发送状态', 'plugin_msg_sendstatus', '1', '', 1, 1, '2023-03-12 12:14:56', '2023-03-12 13:23:37');
INSERT INTO `admin_sys_dict_type` (`id`, `dict_name`, `dict_type`, `status`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (33, '管理-接口-同步状态', 'admin_sys_api_sync_status', '1', '', 1, 1, '2024-12-13 00:29:34', '2024-12-13 00:29:34');
COMMIT;

-- ----------------------------
-- Table structure for admin_sys_gen_column
-- ----------------------------
DROP TABLE IF EXISTS `admin_sys_gen_column`;
CREATE TABLE `admin_sys_gen_column` (
  `id` int NOT NULL AUTO_INCREMENT,
  `table_id` int DEFAULT NULL,
  `column_name` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL,
  `column_comment` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL,
  `column_type` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL,
  `go_type` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL,
  `go_field` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL,
  `json_field` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL,
  `is_pk` char(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL,
  `is_required` char(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '该值是否参与新增或者编辑',
  `is_list` char(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '列表',
  `is_query` char(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL,
  `query_type` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL,
  `html_type` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL,
  `dict_type` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL,
  `sort` bigint DEFAULT NULL,
  `remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL,
  `created_at` datetime DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime DEFAULT NULL COMMENT '最后更新时间',
  `create_by` int DEFAULT NULL COMMENT '创建者',
  `update_by` int DEFAULT NULL COMMENT '更新者',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;

-- ----------------------------
-- Records of admin_sys_gen_column
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for admin_sys_gen_table
-- ----------------------------
DROP TABLE IF EXISTS `admin_sys_gen_table`;
CREATE TABLE `admin_sys_gen_table` (
  `id` int NOT NULL AUTO_INCREMENT,
  `table_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '表名',
  `table_comment` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '表描述',
  `class_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '类名',
  `package_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '应用名',
  `module_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '接口名',
  `function_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '功能描述',
  `function_author` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '作者',
  `business_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '业务名',
  `is_plugin` char(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT '1' COMMENT '是否插件 1-是 2-否',
  `remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL,
  `created_at` datetime DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime DEFAULT NULL COMMENT '最后更新时间',
  `create_by` bigint DEFAULT NULL COMMENT '创建者',
  `update_by` bigint DEFAULT NULL COMMENT '更新者',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;

-- ----------------------------
-- Records of admin_sys_gen_table
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for admin_sys_login_log
-- ----------------------------
DROP TABLE IF EXISTS `admin_sys_login_log`;
CREATE TABLE `admin_sys_login_log` (
  `id` int NOT NULL AUTO_INCREMENT COMMENT '主键编码',
  `user_id` int DEFAULT NULL COMMENT '用户编号',
  `ipaddr` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT 'ip地址',
  `login_location` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '归属地',
  `browser` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '浏览器',
  `os` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '系统',
  `agent` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '代理',
  `platform` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '固件',
  `login_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '登录时间',
  `status` char(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '状态 1-登录 2-退出',
  `remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '备注',
  `created_at` datetime DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime DEFAULT NULL COMMENT '最后更新时间',
  `create_by` int DEFAULT NULL COMMENT '创建者',
  `update_by` int DEFAULT NULL COMMENT '更新者',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;

-- ----------------------------
-- Records of admin_sys_login_log
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for admin_sys_menu
-- ----------------------------
DROP TABLE IF EXISTS `admin_sys_menu`;
CREATE TABLE `admin_sys_menu` (
  `id` int NOT NULL AUTO_INCREMENT,
  `title` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL,
  `icon` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL,
  `path` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL,
  `element` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL,
  `redirect` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL,
  `permission` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL,
  `sort` int DEFAULT NULL,
  `parent_id` int DEFAULT NULL COMMENT '上级菜单id',
  `parent_ids` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '上级菜单id集合',
  `menu_type` char(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL,
  `is_keep_alive` char(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '是否缓存 1-是 2-否',
  `is_affix` char(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '是否固定 1-是 2-否',
  `is_hidden` char(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '是否隐藏 1-是 2-否',
  `is_frame` char(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '是否内嵌 1-是 2-否',
  `create_by` int DEFAULT NULL COMMENT '创建者',
  `update_by` int DEFAULT NULL COMMENT '更新者',
  `created_at` datetime DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime DEFAULT NULL COMMENT '最后更新时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=125 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;

-- ----------------------------
-- Records of admin_sys_menu
-- ----------------------------
BEGIN;
INSERT INTO `admin_sys_menu` (`id`, `title`, `icon`, `path`, `element`, `redirect`, `permission`, `sort`, `parent_id`, `parent_ids`, `menu_type`, `is_keep_alive`, `is_affix`, `is_hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (1, '系统管理', 'AlignLeftOutlined', '/admin/sys', 'Layout', '/admin/sys/sys-api', '', 300, 0, '0,', '1', '', '', '2', NULL, 1, 1, '2021-05-20 21:58:46', '2024-11-21 15:57:20');
INSERT INTO `admin_sys_menu` (`id`, `title`, `icon`, `path`, `element`, `redirect`, `permission`, `sort`, `parent_id`, `parent_ids`, `menu_type`, `is_keep_alive`, `is_affix`, `is_hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (2, '用户管理', 'CheckSquareOutlined', '/admin/sys/sys-user', '/admin/sys/user/index', NULL, '', 10, 1, '0,2', '2', '1', '2', '2', '1', 1, 1, '2021-05-20 22:08:45', '2023-04-27 16:32:41');
INSERT INTO `admin_sys_menu` (`id`, `title`, `icon`, `path`, `element`, `redirect`, `permission`, `sort`, `parent_id`, `parent_ids`, `menu_type`, `is_keep_alive`, `is_affix`, `is_hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (3, '新增管理员', 'AlipayCircleFilled', '', '', '', 'admin:sys-user:add', 10, 2, '0,22,', '3', '', '', '', '', 1, 1, '2021-05-20 22:08:45', '2024-12-09 16:23:55');
INSERT INTO `admin_sys_menu` (`id`, `title`, `icon`, `path`, `element`, `redirect`, `permission`, `sort`, `parent_id`, `parent_ids`, `menu_type`, `is_keep_alive`, `is_affix`, `is_hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (4, '查询管理员', 'AlipayCircleFilled', '', '', '', 'admin:sys-user:query', 40, 2, '0,22,', '3', '', '', '', '', 1, 1, '2021-05-20 22:08:45', '2024-12-09 15:16:36');
INSERT INTO `admin_sys_menu` (`id`, `title`, `icon`, `path`, `element`, `redirect`, `permission`, `sort`, `parent_id`, `parent_ids`, `menu_type`, `is_keep_alive`, `is_affix`, `is_hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (5, '修改管理员', 'AlipayCircleFilled', '', '', '', 'admin:sys-user:edit', 30, 2, '0,22,', '3', '', '', '', '', 1, 1, '2021-05-20 22:08:45', '2024-12-09 16:24:08');
INSERT INTO `admin_sys_menu` (`id`, `title`, `icon`, `path`, `element`, `redirect`, `permission`, `sort`, `parent_id`, `parent_ids`, `menu_type`, `is_keep_alive`, `is_affix`, `is_hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (6, '删除管理员', 'AlipayCircleFilled', '', '', '', 'admin:sys-user:remove', 20, 2, '0,22,', '3', '', '', '', '', 1, 1, '2021-05-20 22:08:45', '2024-12-09 15:15:56');
INSERT INTO `admin_sys_menu` (`id`, `title`, `icon`, `path`, `element`, `redirect`, `permission`, `sort`, `parent_id`, `parent_ids`, `menu_type`, `is_keep_alive`, `is_affix`, `is_hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (7, '菜单管理', 'CheckOutlined', '/admin/sys/sys-menu', '/admin/sys/menu/index', NULL, '', 30, 1, '0,1,', '2', '1', '2', '2', '1', 1, 1, '2021-05-20 22:08:45', '2023-04-27 16:32:41');
INSERT INTO `admin_sys_menu` (`id`, `title`, `icon`, `path`, `element`, `redirect`, `permission`, `sort`, `parent_id`, `parent_ids`, `menu_type`, `is_keep_alive`, `is_affix`, `is_hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (8, '角色管理', 'BoxPlotFilled', '/admin/sys/sys-role', '/admin/sys/role/index', NULL, '', 20, 1, '0,1,', '2', '1', '2', '2', '1', 1, 1, '2021-05-20 22:08:45', '2023-04-27 16:32:41');
INSERT INTO `admin_sys_menu` (`id`, `title`, `icon`, `path`, `element`, `redirect`, `permission`, `sort`, `parent_id`, `parent_ids`, `menu_type`, `is_keep_alive`, `is_affix`, `is_hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (9, '部门管理', 'CheckCircleTwoTone', '/admin/sys/sys-dept', '/admin/sys/dept/index', NULL, '', 40, 1, '0,1,', '2', '1', '2', '2', '1', 1, 1, '2021-05-20 22:08:45', '2023-04-27 16:32:41');
INSERT INTO `admin_sys_menu` (`id`, `title`, `icon`, `path`, `element`, `redirect`, `permission`, `sort`, `parent_id`, `parent_ids`, `menu_type`, `is_keep_alive`, `is_affix`, `is_hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (10, '岗位管理', 'BorderRightOutlined', '/admin/sys/sys-post', '/admin/sys/post/index', NULL, '', 50, 1, '0,1,', '2', '1', '2', '2', '1', 1, 1, '2021-05-20 22:08:45', '2023-04-27 16:32:41');
INSERT INTO `admin_sys_menu` (`id`, `title`, `icon`, `path`, `element`, `redirect`, `permission`, `sort`, `parent_id`, `parent_ids`, `menu_type`, `is_keep_alive`, `is_affix`, `is_hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (11, '字典管理', 'ArrowUpOutlined', '/admin/sys/sys-dicttype', '/admin/sys/dicttype/index', NULL, '', 60, 1, '0,1,', '2', '1', '2', '2', '1', 1, 1, '2021-05-20 22:08:45', '2023-04-27 16:32:41');
INSERT INTO `admin_sys_menu` (`id`, `title`, `icon`, `path`, `element`, `redirect`, `permission`, `sort`, `parent_id`, `parent_ids`, `menu_type`, `is_keep_alive`, `is_affix`, `is_hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (12, '字典数据', 'ArrowUpOutlined', '/admin/sys/sys-dictdata', '/admin/sys/dictdata/index', NULL, '', 100, 1, '0,1,', '2', '1', '2', '1', '1', 1, 1, '2021-05-20 22:08:45', '2023-04-27 16:32:41');
INSERT INTO `admin_sys_menu` (`id`, `title`, `icon`, `path`, `element`, `redirect`, `permission`, `sort`, `parent_id`, `parent_ids`, `menu_type`, `is_keep_alive`, `is_affix`, `is_hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (13, '参数管理', 'CaretDownOutlined', '/admin/sys/sys-config', '/admin/sys/config/index', NULL, '', 70, 1, '0,1,', '2', '1', '2', '2', '1', 1, 1, '2021-05-20 22:08:45', '2023-04-27 16:32:41');
INSERT INTO `admin_sys_menu` (`id`, `title`, `icon`, `path`, `element`, `redirect`, `permission`, `sort`, `parent_id`, `parent_ids`, `menu_type`, `is_keep_alive`, `is_affix`, `is_hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (14, '登录日志', 'BlockOutlined', '/admin/sys/sys-loginlog', '/admin/sys/loginlog/index', NULL, '', 90, 1, '0,1,', '2', '1', '2', '2', '1', 1, 1, '2021-05-20 22:08:45', '2023-04-27 16:32:41');
INSERT INTO `admin_sys_menu` (`id`, `title`, `icon`, `path`, `element`, `redirect`, `permission`, `sort`, `parent_id`, `parent_ids`, `menu_type`, `is_keep_alive`, `is_affix`, `is_hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (15, '操作日志', 'CarFilled', '/admin/sys/sys-operalog', '/admin/sys/operlog/index', NULL, '', 120, 1, '0,1,', '2', '1', '2', '2', '1', 1, 1, '2021-05-20 22:08:45', '2023-04-27 16:32:41');
INSERT INTO `admin_sys_menu` (`id`, `title`, `icon`, `path`, `element`, `redirect`, `permission`, `sort`, `parent_id`, `parent_ids`, `menu_type`, `is_keep_alive`, `is_affix`, `is_hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (16, '新增菜单', 'AlipayCircleFilled', '', '', '', 'admin:sys-menu:add', 1, 7, '0,1,7,', '3', '', '', '', '', 1, 1, '2020-04-11 15:52:48', '2024-12-09 15:25:11');
INSERT INTO `admin_sys_menu` (`id`, `title`, `icon`, `path`, `element`, `redirect`, `permission`, `sort`, `parent_id`, `parent_ids`, `menu_type`, `is_keep_alive`, `is_affix`, `is_hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (17, '修改菜单', 'AlipayCircleFilled', '', '', '', 'admin:sys-menu:edit', 1, 7, '0,1,7,', '3', '', '', '', '', 1, 1, '2020-04-11 15:52:48', '2024-12-09 15:25:27');
INSERT INTO `admin_sys_menu` (`id`, `title`, `icon`, `path`, `element`, `redirect`, `permission`, `sort`, `parent_id`, `parent_ids`, `menu_type`, `is_keep_alive`, `is_affix`, `is_hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (18, '查询菜单', 'AlipayCircleFilled', '', '', '', 'admin:sys-menu:query', 1, 7, '0,1,7,', '3', '', '', '', '', 1, 1, '2020-04-11 15:52:48', '2024-12-09 15:25:40');
INSERT INTO `admin_sys_menu` (`id`, `title`, `icon`, `path`, `element`, `redirect`, `permission`, `sort`, `parent_id`, `parent_ids`, `menu_type`, `is_keep_alive`, `is_affix`, `is_hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (19, '删除菜单', 'AlipayCircleFilled', '', '', '', 'admin:sys-menu:remove', 1, 7, '0,1,7,', '3', '', '', '', '', 1, 1, '2020-04-11 15:52:48', '2024-12-09 15:25:51');
INSERT INTO `admin_sys_menu` (`id`, `title`, `icon`, `path`, `element`, `redirect`, `permission`, `sort`, `parent_id`, `parent_ids`, `menu_type`, `is_keep_alive`, `is_affix`, `is_hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (20, '新增角色', 'AlipayCircleFilled', '', '', '', 'admin:sys-role:add', 1, 8, '0,1,8,', '3', '', '', '', '', 1, 1, '2020-04-11 15:52:48', '2024-12-09 15:21:11');
INSERT INTO `admin_sys_menu` (`id`, `title`, `icon`, `path`, `element`, `redirect`, `permission`, `sort`, `parent_id`, `parent_ids`, `menu_type`, `is_keep_alive`, `is_affix`, `is_hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (21, '查询角色', 'AlipayCircleFilled', '', '', '', 'admin:sys-role:query', 1, 8, '0,1,8,', '3', '', '', '', '', 1, 1, '2020-04-11 15:52:48', '2024-12-09 15:20:42');
INSERT INTO `admin_sys_menu` (`id`, `title`, `icon`, `path`, `element`, `redirect`, `permission`, `sort`, `parent_id`, `parent_ids`, `menu_type`, `is_keep_alive`, `is_affix`, `is_hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (22, '修改角色', 'AlipayCircleFilled', '', '', '', 'admin:sys-role:update', 1, 8, '0,1,8,', '3', '', '', '', '', 1, 1, '2020-04-11 15:52:48', '2024-12-09 15:34:37');
INSERT INTO `admin_sys_menu` (`id`, `title`, `icon`, `path`, `element`, `redirect`, `permission`, `sort`, `parent_id`, `parent_ids`, `menu_type`, `is_keep_alive`, `is_affix`, `is_hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (23, '删除角色', 'AlipayCircleFilled', '', '', '', 'admin:sys-role:remove', 1, 8, '0,1,8,', '3', '', '', '', '', 1, 1, '2020-04-11 15:52:48', '2024-12-09 15:24:50');
INSERT INTO `admin_sys_menu` (`id`, `title`, `icon`, `path`, `element`, `redirect`, `permission`, `sort`, `parent_id`, `parent_ids`, `menu_type`, `is_keep_alive`, `is_affix`, `is_hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (24, '查询部门', 'AlipayCircleFilled', '', '', '', 'admin:sys-dept:query', 40, 9, '0,1,9,', '3', '', '', '', '', 1, 1, '2021-05-20 22:08:45', '2024-12-09 15:27:01');
INSERT INTO `admin_sys_menu` (`id`, `title`, `icon`, `path`, `element`, `redirect`, `permission`, `sort`, `parent_id`, `parent_ids`, `menu_type`, `is_keep_alive`, `is_affix`, `is_hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (25, '新增部门', 'AlipayCircleFilled', '', '', '', 'admin:sys-dept:add', 10, 9, '0,1,9,', '3', '', '', '', '', 1, 1, '2021-05-20 22:08:45', '2024-12-09 15:26:07');
INSERT INTO `admin_sys_menu` (`id`, `title`, `icon`, `path`, `element`, `redirect`, `permission`, `sort`, `parent_id`, `parent_ids`, `menu_type`, `is_keep_alive`, `is_affix`, `is_hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (26, '修改部门', 'AlipayCircleFilled', '', '', '', 'admin:sys-dept:edit', 30, 9, '0,1,9,', '3', '', '', '', '', 1, 1, '2021-05-20 22:08:45', '2024-12-09 15:27:15');
INSERT INTO `admin_sys_menu` (`id`, `title`, `icon`, `path`, `element`, `redirect`, `permission`, `sort`, `parent_id`, `parent_ids`, `menu_type`, `is_keep_alive`, `is_affix`, `is_hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (27, '删除部门', 'AlipayCircleFilled', '', '', '', 'admin:sys-dept:remove', 20, 9, '0,1,9,', '3', '', '', '', '', 1, 1, '2021-05-20 22:08:45', '2024-12-09 15:26:19');
INSERT INTO `admin_sys_menu` (`id`, `title`, `icon`, `path`, `element`, `redirect`, `permission`, `sort`, `parent_id`, `parent_ids`, `menu_type`, `is_keep_alive`, `is_affix`, `is_hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (28, '查询岗位', 'AlipayCircleFilled', '', '', '', 'admin:sys-post:query', 0, 10, '0,1,10,', '3', '', '', '', '', 1, 1, '2020-04-11 15:52:48', '2024-12-09 15:29:02');
INSERT INTO `admin_sys_menu` (`id`, `title`, `icon`, `path`, `element`, `redirect`, `permission`, `sort`, `parent_id`, `parent_ids`, `menu_type`, `is_keep_alive`, `is_affix`, `is_hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (29, '新增岗位', 'AlipayCircleFilled', '', '', '', 'admin:sys-post:add', 0, 10, '0,1,10,', '3', '', '', '', '', 1, 1, '2020-04-11 15:52:48', '2024-12-09 15:29:15');
INSERT INTO `admin_sys_menu` (`id`, `title`, `icon`, `path`, `element`, `redirect`, `permission`, `sort`, `parent_id`, `parent_ids`, `menu_type`, `is_keep_alive`, `is_affix`, `is_hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (30, '修改岗位', 'AlipayCircleFilled', '', '', '', 'admin:sys-post:edit', 0, 10, '0,1,10,', '3', '', '', '', '', 1, 1, '2020-04-11 15:52:48', '2024-12-09 15:29:27');
INSERT INTO `admin_sys_menu` (`id`, `title`, `icon`, `path`, `element`, `redirect`, `permission`, `sort`, `parent_id`, `parent_ids`, `menu_type`, `is_keep_alive`, `is_affix`, `is_hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (31, '删除岗位', 'AlipayCircleFilled', '', '', '', 'admin:sys-post:remove', 0, 10, '0,1,10,', '3', '', '', '', '', 1, 1, '2020-04-11 15:52:48', '2024-12-09 15:29:41');
INSERT INTO `admin_sys_menu` (`id`, `title`, `icon`, `path`, `element`, `redirect`, `permission`, `sort`, `parent_id`, `parent_ids`, `menu_type`, `is_keep_alive`, `is_affix`, `is_hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (32, '查询字典', 'AlipayCircleFilled', '', '', NULL, 'admin:sys-dict-type:query', 0, 11, '0,1,11,', '3', NULL, NULL, NULL, NULL, 1, 1, '2020-04-11 15:52:48', '2023-05-09 11:13:23');
INSERT INTO `admin_sys_menu` (`id`, `title`, `icon`, `path`, `element`, `redirect`, `permission`, `sort`, `parent_id`, `parent_ids`, `menu_type`, `is_keep_alive`, `is_affix`, `is_hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (33, '新增类型', 'AlipayCircleFilled', '', '', NULL, 'admin:sys-dict-type:add', 0, 11, '0,1,11,', '3', NULL, NULL, NULL, NULL, 1, 1, '2020-04-11 15:52:48', '2023-05-09 11:13:37');
INSERT INTO `admin_sys_menu` (`id`, `title`, `icon`, `path`, `element`, `redirect`, `permission`, `sort`, `parent_id`, `parent_ids`, `menu_type`, `is_keep_alive`, `is_affix`, `is_hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (34, '修改类型', 'AlipayCircleFilled', '', '', NULL, 'admin:sys-dict-type:edit', 0, 11, '0,1,11,', '3', NULL, NULL, NULL, NULL, 1, 1, '2020-04-11 15:52:48', '2023-05-09 11:13:58');
INSERT INTO `admin_sys_menu` (`id`, `title`, `icon`, `path`, `element`, `redirect`, `permission`, `sort`, `parent_id`, `parent_ids`, `menu_type`, `is_keep_alive`, `is_affix`, `is_hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (35, '删除类型', 'AlipayCircleFilled', '', '', NULL, 'admin:sys-dict-type:remove', 0, 11, '0,1,11,', '3', NULL, NULL, NULL, NULL, 1, 1, '2020-04-11 15:52:48', '2023-05-09 11:15:32');
INSERT INTO `admin_sys_menu` (`id`, `title`, `icon`, `path`, `element`, `redirect`, `permission`, `sort`, `parent_id`, `parent_ids`, `menu_type`, `is_keep_alive`, `is_affix`, `is_hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (36, '查询数据', 'AlipayCircleFilled', '', '', NULL, 'admin:sys-dict-data:query', 0, 12, '0,1,12,', '3', NULL, NULL, NULL, NULL, 1, 1, '2020-04-11 15:52:48', '2023-05-09 11:09:55');
INSERT INTO `admin_sys_menu` (`id`, `title`, `icon`, `path`, `element`, `redirect`, `permission`, `sort`, `parent_id`, `parent_ids`, `menu_type`, `is_keep_alive`, `is_affix`, `is_hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (37, '新增数据', 'AlipayCircleFilled', '', '', NULL, 'admin:sys-dict-data:add', 0, 12, '0,1,12,', '3', NULL, NULL, NULL, NULL, 1, 1, '2020-04-11 15:52:48', '2023-05-09 11:10:05');
INSERT INTO `admin_sys_menu` (`id`, `title`, `icon`, `path`, `element`, `redirect`, `permission`, `sort`, `parent_id`, `parent_ids`, `menu_type`, `is_keep_alive`, `is_affix`, `is_hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (38, '修改数据', 'AlipayCircleFilled', '', '', NULL, 'admin:sys-dict-data:edit', 0, 12, '0,1,12,', '3', NULL, NULL, NULL, NULL, 1, 1, '2020-04-11 15:52:48', '2023-05-09 11:11:11');
INSERT INTO `admin_sys_menu` (`id`, `title`, `icon`, `path`, `element`, `redirect`, `permission`, `sort`, `parent_id`, `parent_ids`, `menu_type`, `is_keep_alive`, `is_affix`, `is_hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (39, '删除数据', 'AlipayCircleFilled', '', '', NULL, 'admin:sys-dict-data:remove', 0, 12, '0,1,12,', '3', NULL, NULL, NULL, NULL, 1, 1, '2020-04-11 15:52:48', '2023-05-09 11:11:43');
INSERT INTO `admin_sys_menu` (`id`, `title`, `icon`, `path`, `element`, `redirect`, `permission`, `sort`, `parent_id`, `parent_ids`, `menu_type`, `is_keep_alive`, `is_affix`, `is_hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (40, '查询参数', 'AlipayCircleFilled', '', '', NULL, 'admin:sys-config:query', 0, 13, '0,1,13,', '3', NULL, NULL, NULL, NULL, 1, 1, '2020-04-11 15:52:48', '2023-05-09 11:25:51');
INSERT INTO `admin_sys_menu` (`id`, `title`, `icon`, `path`, `element`, `redirect`, `permission`, `sort`, `parent_id`, `parent_ids`, `menu_type`, `is_keep_alive`, `is_affix`, `is_hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (41, '新增参数', 'AlipayCircleFilled', '', '', NULL, 'admin:sys-config:add', 0, 13, '0,1,13,', '3', NULL, NULL, NULL, NULL, 1, 1, '2020-04-11 15:52:48', '2023-05-09 11:26:03');
INSERT INTO `admin_sys_menu` (`id`, `title`, `icon`, `path`, `element`, `redirect`, `permission`, `sort`, `parent_id`, `parent_ids`, `menu_type`, `is_keep_alive`, `is_affix`, `is_hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (42, '修改参数', 'AlipayCircleFilled', '', '', NULL, 'admin:sys-config:edit', 0, 13, '0,1,13,', '3', NULL, NULL, NULL, NULL, 1, 1, '2020-04-11 15:52:48', '2023-05-09 11:26:27');
INSERT INTO `admin_sys_menu` (`id`, `title`, `icon`, `path`, `element`, `redirect`, `permission`, `sort`, `parent_id`, `parent_ids`, `menu_type`, `is_keep_alive`, `is_affix`, `is_hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (43, '删除参数', 'AlipayCircleFilled', '', '', NULL, 'admin:sys-config:remove', 0, 13, '0,1,13,', '3', NULL, NULL, NULL, NULL, 1, 1, '2020-04-11 15:52:48', '2023-05-09 11:26:39');
INSERT INTO `admin_sys_menu` (`id`, `title`, `icon`, `path`, `element`, `redirect`, `permission`, `sort`, `parent_id`, `parent_ids`, `menu_type`, `is_keep_alive`, `is_affix`, `is_hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (44, '查询登录日志', 'AlipayCircleFilled', '', '', NULL, 'admin:sys-login-log:query', 0, 14, '0,1,14,', '3', NULL, NULL, NULL, NULL, 1, 1, '2020-04-11 15:52:48', '2023-05-09 11:03:23');
INSERT INTO `admin_sys_menu` (`id`, `title`, `icon`, `path`, `element`, `redirect`, `permission`, `sort`, `parent_id`, `parent_ids`, `menu_type`, `is_keep_alive`, `is_affix`, `is_hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (45, '删除登录日志', 'AlipayCircleFilled', '', '', NULL, 'admin:sys-login-log:remove', 0, 14, '0,1,14,', '3', NULL, NULL, NULL, NULL, 1, 1, '2020-04-11 15:52:48', '2023-05-09 11:03:46');
INSERT INTO `admin_sys_menu` (`id`, `title`, `icon`, `path`, `element`, `redirect`, `permission`, `sort`, `parent_id`, `parent_ids`, `menu_type`, `is_keep_alive`, `is_affix`, `is_hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (46, '查询操作日志', 'AlipayCircleFilled', '', '', NULL, 'admin:sys-oper-log:query', 0, 15, '0,1,15,', '3', NULL, NULL, NULL, NULL, 1, 1, '2020-04-11 15:52:48', '2023-05-09 11:00:44');
INSERT INTO `admin_sys_menu` (`id`, `title`, `icon`, `path`, `element`, `redirect`, `permission`, `sort`, `parent_id`, `parent_ids`, `menu_type`, `is_keep_alive`, `is_affix`, `is_hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (47, '删除操作日志', 'AlipayCircleFilled', '', '', NULL, 'admin:sys-oper-log:remove', 0, 15, '0,1,15,', '3', NULL, NULL, NULL, NULL, 1, 1, '2020-04-11 15:52:48', '2023-05-09 11:00:58');
INSERT INTO `admin_sys_menu` (`id`, `title`, `icon`, `path`, `element`, `redirect`, `permission`, `sort`, `parent_id`, `parent_ids`, `menu_type`, `is_keep_alive`, `is_affix`, `is_hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (48, '代码生成', 'AndroidOutlined', '/admin/sys/tools/sys-gen', '/admin/sys/tools/gen/index', NULL, '', 20, 54, '0,54,', '2', '1', '2', '2', '1', 1, 1, '2020-04-11 15:52:48', '2023-05-09 10:55:36');
INSERT INTO `admin_sys_menu` (`id`, `title`, `icon`, `path`, `element`, `redirect`, `permission`, `sort`, `parent_id`, `parent_ids`, `menu_type`, `is_keep_alive`, `is_affix`, `is_hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (49, '代码生成修改', 'AliwangwangOutlined', '/admin/sys/tools/sys-edit-table', '/admin/sys/tools/gen/edit/index', NULL, '', 100, 54, '0,54,', '2', '1', '2', '1', '1', 1, 1, '2020-04-11 15:52:48', '2023-05-09 11:14:37');
INSERT INTO `admin_sys_menu` (`id`, `title`, `icon`, `path`, `element`, `redirect`, `permission`, `sort`, `parent_id`, `parent_ids`, `menu_type`, `is_keep_alive`, `is_affix`, `is_hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (50, '服务监控', 'ArrowLeftOutlined', '/admin/sys/tools/monitor', '/admin/sys/tools/monitor/index', '', '', 0, 54, '0,54,', '2', '2', '2', '2', '1', 1, 1, '2020-04-14 00:28:19', '2024-12-11 11:03:23');
INSERT INTO `admin_sys_menu` (`id`, `title`, `icon`, `path`, `element`, `redirect`, `permission`, `sort`, `parent_id`, `parent_ids`, `menu_type`, `is_keep_alive`, `is_affix`, `is_hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (51, '接口管理', 'AlertTwoTone', '/admin/sys/sys-api', '/admin/sys/api/index', NULL, '', 0, 1, '0,1,', '2', '1', '2', '2', '1', 1, 1, '2021-05-20 22:08:45', '2023-04-27 16:32:41');
INSERT INTO `admin_sys_menu` (`id`, `title`, `icon`, `path`, `element`, `redirect`, `permission`, `sort`, `parent_id`, `parent_ids`, `menu_type`, `is_keep_alive`, `is_affix`, `is_hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (52, '查询接口', 'AlipayCircleFilled', '', '', '', 'admin:sys-api:query', 40, 51, '0,1,51,', '3', '', '', '', '', 1, 1, '2021-05-20 22:08:45', '2024-12-09 15:06:19');
INSERT INTO `admin_sys_menu` (`id`, `title`, `icon`, `path`, `element`, `redirect`, `permission`, `sort`, `parent_id`, `parent_ids`, `menu_type`, `is_keep_alive`, `is_affix`, `is_hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (53, '修改接口', 'AlipayCircleFilled', '', '', '', 'admin:sys-api:edit', 30, 51, '0,1,51,', '3', '', '', '', '', 1, 1, '2021-05-20 22:08:45', '2024-12-09 15:05:59');
INSERT INTO `admin_sys_menu` (`id`, `title`, `icon`, `path`, `element`, `redirect`, `permission`, `sort`, `parent_id`, `parent_ids`, `menu_type`, `is_keep_alive`, `is_affix`, `is_hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (54, '系统工具', 'CaretRightFilled', '/admin/sys/sys-tools', 'Layout', '/sys-tools/sys-monitor', '', 330, 0, '0,', '1', '', '', '2', NULL, 1, 1, '2021-05-21 11:13:32', '2024-11-21 15:57:40');
INSERT INTO `admin_sys_menu` (`id`, `title`, `icon`, `path`, `element`, `redirect`, `permission`, `sort`, `parent_id`, `parent_ids`, `menu_type`, `is_keep_alive`, `is_affix`, `is_hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (55, '文件管理', 'AlipayOutlined', '/plugins/filemgr', '/index', '/plugins/filemgr/filemgr-app', '', 90, 57, '0,57,', '1', '', '', '2', NULL, 1, 1, '2021-08-13 14:19:11', '2024-11-21 15:56:20');
INSERT INTO `admin_sys_menu` (`id`, `title`, `icon`, `path`, `element`, `redirect`, `permission`, `sort`, `parent_id`, `parent_ids`, `menu_type`, `is_keep_alive`, `is_affix`, `is_hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (56, '内容管理', 'AndroidFilled', '/plugins/content', '/index', '/plugins/content/content-category', '', 60, 57, '0,57,', '1', '', '', '2', NULL, 1, 1, '2021-08-16 18:01:20', '2024-11-21 15:55:48');
INSERT INTO `admin_sys_menu` (`id`, `title`, `icon`, `path`, `element`, `redirect`, `permission`, `sort`, `parent_id`, `parent_ids`, `menu_type`, `is_keep_alive`, `is_affix`, `is_hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (57, '插件管理', 'AmazonCircleFilled', '/plugins', 'Layout', '/plugins/content/content-category', '', 270, 0, '0,', '1', '', '', '2', NULL, 1, 1, '2023-03-07 10:37:37', '2024-11-21 15:55:39');
INSERT INTO `admin_sys_menu` (`id`, `title`, `icon`, `path`, `element`, `redirect`, `permission`, `sort`, `parent_id`, `parent_ids`, `menu_type`, `is_keep_alive`, `is_affix`, `is_hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (58, 'App应用', 'AlipayCircleFilled', '/app', 'Layout', '/app/user/user', '', 30, 0, '0,', '1', '', '', '2', NULL, 1, 1, '2023-03-08 09:27:36', '2024-11-21 16:05:44');
INSERT INTO `admin_sys_menu` (`id`, `title`, `icon`, `path`, `element`, `redirect`, `permission`, `sort`, `parent_id`, `parent_ids`, `menu_type`, `is_keep_alive`, `is_affix`, `is_hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (59, '用户列表', 'AmazonOutlined', '/app/user', '/index', '/app/user/user', '', 30, 58, '0,58,', '1', '', '', '2', NULL, 1, 1, '2023-03-09 14:24:25', '2024-11-21 15:53:06');
INSERT INTO `admin_sys_menu` (`id`, `title`, `icon`, `path`, `element`, `redirect`, `permission`, `sort`, `parent_id`, `parent_ids`, `menu_type`, `is_keep_alive`, `is_affix`, `is_hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (60, '财务管理', 'BackwardFilled', '/app/account', '/index', '/app/user/user-account-log', '', 60, 58, '0,58,', '1', '', '', '2', NULL, 1, 1, '2023-03-09 21:13:23', '2024-11-21 15:53:56');
INSERT INTO `admin_sys_menu` (`id`, `title`, `icon`, `path`, `element`, `redirect`, `permission`, `sort`, `parent_id`, `parent_ids`, `menu_type`, `is_keep_alive`, `is_affix`, `is_hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (61, '用户等级', 'BorderRightOutlined', '/app/user/user-level', '/app/user/user-level/index', NULL, '', 60, 59, '0,58,59,', '2', '1', '2', '2', '1', 1, 1, '2023-03-09 21:33:49', '2023-03-09 23:05:34');
INSERT INTO `admin_sys_menu` (`id`, `title`, `icon`, `path`, `element`, `redirect`, `permission`, `sort`, `parent_id`, `parent_ids`, `menu_type`, `is_keep_alive`, `is_affix`, `is_hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (62, '分页获取用户等级', '', '', '', NULL, 'app:user:user-level:query', 0, 61, '0,58,59,61,,', '3', NULL, NULL, NULL, NULL, 1, 1, '2023-03-09 21:33:49', '2023-05-09 10:22:48');
INSERT INTO `admin_sys_menu` (`id`, `title`, `icon`, `path`, `element`, `redirect`, `permission`, `sort`, `parent_id`, `parent_ids`, `menu_type`, `is_keep_alive`, `is_affix`, `is_hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (63, '创建用户等级', '', '', '', NULL, 'app:user:user-level:add', 0, 61, '0,58,59,61,,', '3', NULL, NULL, NULL, NULL, 1, 1, '2023-03-09 21:33:49', '2023-05-09 10:23:06');
INSERT INTO `admin_sys_menu` (`id`, `title`, `icon`, `path`, `element`, `redirect`, `permission`, `sort`, `parent_id`, `parent_ids`, `menu_type`, `is_keep_alive`, `is_affix`, `is_hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (64, '修改用户等级', '', '', '', NULL, 'app:user:user-level:edit', 0, 61, '0,58,59,61,,', '3', NULL, NULL, NULL, NULL, 1, 1, '2023-03-09 21:33:49', '2023-05-09 10:23:27');
INSERT INTO `admin_sys_menu` (`id`, `title`, `icon`, `path`, `element`, `redirect`, `permission`, `sort`, `parent_id`, `parent_ids`, `menu_type`, `is_keep_alive`, `is_affix`, `is_hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (65, '删除用户等级', '', '', '', NULL, 'app:user:user-level:del', 0, 61, '0,58,59,61,,', '3', NULL, NULL, NULL, NULL, 1, 1, '2023-03-09 21:33:49', '2023-05-09 10:23:44');
INSERT INTO `admin_sys_menu` (`id`, `title`, `icon`, `path`, `element`, `redirect`, `permission`, `sort`, `parent_id`, `parent_ids`, `menu_type`, `is_keep_alive`, `is_affix`, `is_hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (66, '导出用户等级', '', '', '', NULL, 'app:user:user-level:export', 0, 61, '0,58,59,61,,', '3', NULL, NULL, NULL, NULL, 1, 1, '2023-03-09 21:33:49', '2023-05-09 10:23:54');
INSERT INTO `admin_sys_menu` (`id`, `title`, `icon`, `path`, `element`, `redirect`, `permission`, `sort`, `parent_id`, `parent_ids`, `menu_type`, `is_keep_alive`, `is_affix`, `is_hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (67, '账变记录', 'BorderRightOutlined', '/app/user/user-account-log', '/app/user/user-account-log/index', NULL, '', 0, 60, '0,58,60,', '2', '1', '2', '2', '1', 1, 1, '2023-03-09 21:33:51', '2023-03-09 21:35:31');
INSERT INTO `admin_sys_menu` (`id`, `title`, `icon`, `path`, `element`, `redirect`, `permission`, `sort`, `parent_id`, `parent_ids`, `menu_type`, `is_keep_alive`, `is_affix`, `is_hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (68, '分页获取账变记录', '', '', '', NULL, 'app:user:user-account-log:query', 0, 67, '0,58,60,67,', '3', NULL, NULL, NULL, NULL, 1, 1, '2023-03-09 21:33:51', '2023-05-09 10:32:28');
INSERT INTO `admin_sys_menu` (`id`, `title`, `icon`, `path`, `element`, `redirect`, `permission`, `sort`, `parent_id`, `parent_ids`, `menu_type`, `is_keep_alive`, `is_affix`, `is_hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (69, '导出账变记录', '', '', '', NULL, 'app:user:user-account-log:export', 0, 67, '0,58,60,67,', '3', NULL, NULL, NULL, NULL, 1, 1, '2023-03-09 21:33:51', '2023-05-09 10:33:04');
INSERT INTO `admin_sys_menu` (`id`, `title`, `icon`, `path`, `element`, `redirect`, `permission`, `sort`, `parent_id`, `parent_ids`, `menu_type`, `is_keep_alive`, `is_affix`, `is_hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (70, '用户配置', 'BorderRightOutlined', '/app/user/user-conf', '/app/user/user-conf/index', NULL, '', 90, 59, '0,58,59,', '2', '1', '2', '2', '1', 1, 1, '2023-03-09 23:04:40', '2023-03-11 15:02:32');
INSERT INTO `admin_sys_menu` (`id`, `title`, `icon`, `path`, `element`, `redirect`, `permission`, `sort`, `parent_id`, `parent_ids`, `menu_type`, `is_keep_alive`, `is_affix`, `is_hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (71, '分页获取用户配置', '', '', '', NULL, 'app:user:user-conf:query', 0, 70, '0,58,59,70,', '3', NULL, NULL, NULL, NULL, 1, 1, '2023-03-09 23:04:40', '2023-05-09 10:24:15');
INSERT INTO `admin_sys_menu` (`id`, `title`, `icon`, `path`, `element`, `redirect`, `permission`, `sort`, `parent_id`, `parent_ids`, `menu_type`, `is_keep_alive`, `is_affix`, `is_hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (72, '修改用户配置', '', '', '', NULL, 'app:user:user-conf:edit', 0, 70, '0,58,59,70,', '3', NULL, NULL, NULL, NULL, 1, 1, '2023-03-09 23:04:40', '2023-05-09 10:25:16');
INSERT INTO `admin_sys_menu` (`id`, `title`, `icon`, `path`, `element`, `redirect`, `permission`, `sort`, `parent_id`, `parent_ids`, `menu_type`, `is_keep_alive`, `is_affix`, `is_hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (73, '用户管理', 'BorderRightOutlined', '/app/user/user', '/app/user/user/index', '', '', 30, 59, '0,58,59,', '2', '1', '2', '2', '1', 1, 1, '2023-03-09 23:18:49', '2024-12-09 16:46:50');
INSERT INTO `admin_sys_menu` (`id`, `title`, `icon`, `path`, `element`, `redirect`, `permission`, `sort`, `parent_id`, `parent_ids`, `menu_type`, `is_keep_alive`, `is_affix`, `is_hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (74, '分页获取用户管理', '', '', '', '', 'app:user:user:query', 0, 73, '0,58,59,73,', '3', '', '', '', '', 1, 1, '2023-03-09 23:18:49', '2024-12-09 16:46:50');
INSERT INTO `admin_sys_menu` (`id`, `title`, `icon`, `path`, `element`, `redirect`, `permission`, `sort`, `parent_id`, `parent_ids`, `menu_type`, `is_keep_alive`, `is_affix`, `is_hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (75, '创建用户管理', '', '', '', '', 'app:user:user:add', 0, 73, '0,58,59,73,', '3', '', '', '', '', 1, 1, '2023-03-09 23:18:49', '2024-12-09 16:46:50');
INSERT INTO `admin_sys_menu` (`id`, `title`, `icon`, `path`, `element`, `redirect`, `permission`, `sort`, `parent_id`, `parent_ids`, `menu_type`, `is_keep_alive`, `is_affix`, `is_hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (76, '修改用户管理', '', '', '', '', 'app:user:user:edit', 0, 73, '0,58,59,73,', '3', '', '', '', '', 1, 1, '2023-03-09 23:18:49', '2024-12-09 16:46:50');
INSERT INTO `admin_sys_menu` (`id`, `title`, `icon`, `path`, `element`, `redirect`, `permission`, `sort`, `parent_id`, `parent_ids`, `menu_type`, `is_keep_alive`, `is_affix`, `is_hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (77, '导出用户管理', '', '', '', '', 'app:user:user:export', 0, 73, '0,58,59,73,', '3', '', '', '', '', 1, 1, '2023-03-09 23:18:49', '2024-12-09 16:46:50');
INSERT INTO `admin_sys_menu` (`id`, `title`, `icon`, `path`, `element`, `redirect`, `permission`, `sort`, `parent_id`, `parent_ids`, `menu_type`, `is_keep_alive`, `is_affix`, `is_hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (78, '用户行为记录', 'BorderRightOutlined', '/app/user/user-oper-log', '/app/user/user-oper-log/index', NULL, '', 120, 59, '0,58,59,', '2', '1', '2', '2', '1', 1, 1, '2023-03-11 15:00:06', '2023-03-11 15:02:42');
INSERT INTO `admin_sys_menu` (`id`, `title`, `icon`, `path`, `element`, `redirect`, `permission`, `sort`, `parent_id`, `parent_ids`, `menu_type`, `is_keep_alive`, `is_affix`, `is_hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (79, '分页获取用户关键行为日志表', '', '', '', NULL, 'app:user:user-oper-log:query', 0, 78, '0,58,59,78,', '3', NULL, NULL, NULL, NULL, 1, 1, '2023-03-11 15:00:06', '2023-05-09 10:26:19');
INSERT INTO `admin_sys_menu` (`id`, `title`, `icon`, `path`, `element`, `redirect`, `permission`, `sort`, `parent_id`, `parent_ids`, `menu_type`, `is_keep_alive`, `is_affix`, `is_hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (80, '导出用户关键行为日志表', '', '', '', NULL, 'app:user:user-oper-log:export', 0, 78, '0,58,59,78,', '3', NULL, NULL, NULL, NULL, 1, 1, '2023-03-11 15:00:06', '2023-05-09 10:28:23');
INSERT INTO `admin_sys_menu` (`id`, `title`, `icon`, `path`, `element`, `redirect`, `permission`, `sort`, `parent_id`, `parent_ids`, `menu_type`, `is_keep_alive`, `is_affix`, `is_hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (81, '消息管理', 'AlipaySquareFilled', '/plugins/msg', '/index', '/plugins/msg/msg-code', '', 120, 57, '0,57,', '1', '', '', '2', NULL, 1, 1, '2023-03-12 13:27:59', '2024-11-21 15:56:39');
INSERT INTO `admin_sys_menu` (`id`, `title`, `icon`, `path`, `element`, `redirect`, `permission`, `sort`, `parent_id`, `parent_ids`, `menu_type`, `is_keep_alive`, `is_affix`, `is_hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (82, '验证码记录', 'BorderRightOutlined', '/plugins/msg/msg-code', '/plugins/msg/msg-code/index', NULL, '', 0, 81, '0,57,81,', '2', '1', '2', '2', '1', 1, 1, '2023-03-12 21:54:02', '2023-03-12 21:54:32');
INSERT INTO `admin_sys_menu` (`id`, `title`, `icon`, `path`, `element`, `redirect`, `permission`, `sort`, `parent_id`, `parent_ids`, `menu_type`, `is_keep_alive`, `is_affix`, `is_hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (83, '分页获取验证码记录', '', '', '', NULL, 'plugins:msg:msg-code:query', 0, 82, '0,57,81,82,', '3', NULL, NULL, NULL, NULL, 1, 1, '2023-03-12 21:54:02', '2023-05-09 10:34:35');
INSERT INTO `admin_sys_menu` (`id`, `title`, `icon`, `path`, `element`, `redirect`, `permission`, `sort`, `parent_id`, `parent_ids`, `menu_type`, `is_keep_alive`, `is_affix`, `is_hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (84, '公告管理', 'BorderRightOutlined', '/plugins/content/content-announcement', '/plugins/content/content-announcement/index', NULL, '', 90, 56, '0,57,56,', '2', '1', '2', '2', '1', 1, 1, '2023-03-12 22:47:11', '2023-03-12 22:48:08');
INSERT INTO `admin_sys_menu` (`id`, `title`, `icon`, `path`, `element`, `redirect`, `permission`, `sort`, `parent_id`, `parent_ids`, `menu_type`, `is_keep_alive`, `is_affix`, `is_hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (85, '分页获取公告管理', '', '', '', NULL, 'plugins:content:content-announcement:query', 0, 84, '0,57,56,84,', '3', NULL, NULL, NULL, NULL, 1, 1, '2023-03-12 22:47:11', '2023-05-09 10:49:57');
INSERT INTO `admin_sys_menu` (`id`, `title`, `icon`, `path`, `element`, `redirect`, `permission`, `sort`, `parent_id`, `parent_ids`, `menu_type`, `is_keep_alive`, `is_affix`, `is_hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (86, '创建公告管理', '', '', '', NULL, 'plugins:content:content-announcement:add', 0, 84, '0,57,56,84,', '3', NULL, NULL, NULL, NULL, 1, 1, '2023-03-12 22:47:11', '2023-05-09 10:50:13');
INSERT INTO `admin_sys_menu` (`id`, `title`, `icon`, `path`, `element`, `redirect`, `permission`, `sort`, `parent_id`, `parent_ids`, `menu_type`, `is_keep_alive`, `is_affix`, `is_hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (87, '修改公告管理', '', '', '', NULL, 'plugins:content:content-announcement:edit', 0, 84, '0,57,56,84,', '3', NULL, NULL, NULL, NULL, 1, 1, '2023-03-12 22:47:11', '2023-05-09 10:50:39');
INSERT INTO `admin_sys_menu` (`id`, `title`, `icon`, `path`, `element`, `redirect`, `permission`, `sort`, `parent_id`, `parent_ids`, `menu_type`, `is_keep_alive`, `is_affix`, `is_hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (88, '删除公告管理', '', '', '', NULL, 'plugins:content:content-announcement:del', 0, 84, '0,57,56,84,', '3', NULL, NULL, NULL, NULL, 1, 1, '2023-03-12 22:47:11', '2023-05-09 10:50:58');
INSERT INTO `admin_sys_menu` (`id`, `title`, `icon`, `path`, `element`, `redirect`, `permission`, `sort`, `parent_id`, `parent_ids`, `menu_type`, `is_keep_alive`, `is_affix`, `is_hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (89, '导出公告管理', '', '', '', NULL, 'plugins:content:content-announcement:export', 0, 84, '0,57,56,84,', '3', NULL, NULL, NULL, NULL, 1, 1, '2023-03-12 22:47:11', '2023-05-09 10:51:08');
INSERT INTO `admin_sys_menu` (`id`, `title`, `icon`, `path`, `element`, `redirect`, `permission`, `sort`, `parent_id`, `parent_ids`, `menu_type`, `is_keep_alive`, `is_affix`, `is_hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (90, '内容分类', 'BorderRightOutlined', '/plugins/content/content-category', '/plugins/content/content-category/index', NULL, '', 0, 56, '0,57,56,', '2', '1', '2', '2', '1', 1, 1, '2023-03-12 23:17:44', '2023-03-12 23:20:35');
INSERT INTO `admin_sys_menu` (`id`, `title`, `icon`, `path`, `element`, `redirect`, `permission`, `sort`, `parent_id`, `parent_ids`, `menu_type`, `is_keep_alive`, `is_affix`, `is_hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (91, '分页获取内容分类', '', '', '', NULL, 'plugins:content:content-category:query', 0, 90, '0,57,56,90,', '3', NULL, NULL, NULL, NULL, 1, 1, '2023-03-12 23:17:44', '2023-05-09 10:43:56');
INSERT INTO `admin_sys_menu` (`id`, `title`, `icon`, `path`, `element`, `redirect`, `permission`, `sort`, `parent_id`, `parent_ids`, `menu_type`, `is_keep_alive`, `is_affix`, `is_hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (92, '创建内容分类', '', '', '', NULL, 'plugins:content:content-category:add', 0, 90, '0,57,56,90,', '3', NULL, NULL, NULL, NULL, 1, 1, '2023-03-12 23:17:44', '2023-05-09 10:44:14');
INSERT INTO `admin_sys_menu` (`id`, `title`, `icon`, `path`, `element`, `redirect`, `permission`, `sort`, `parent_id`, `parent_ids`, `menu_type`, `is_keep_alive`, `is_affix`, `is_hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (93, '修改内容分类', '', '', '', NULL, 'plugins:content:content-category:edit', 0, 90, '0,57,56,90,', '3', NULL, NULL, NULL, NULL, 1, 1, '2023-03-12 23:17:44', '2023-05-09 10:44:33');
INSERT INTO `admin_sys_menu` (`id`, `title`, `icon`, `path`, `element`, `redirect`, `permission`, `sort`, `parent_id`, `parent_ids`, `menu_type`, `is_keep_alive`, `is_affix`, `is_hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (94, '删除内容分类', '', '', '', NULL, 'plugins:content:content-category:del', 0, 90, '0,57,56,90,', '3', NULL, NULL, NULL, NULL, 1, 1, '2023-03-12 23:17:45', '2023-05-09 10:47:01');
INSERT INTO `admin_sys_menu` (`id`, `title`, `icon`, `path`, `element`, `redirect`, `permission`, `sort`, `parent_id`, `parent_ids`, `menu_type`, `is_keep_alive`, `is_affix`, `is_hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (95, '导出内容分类', '', '', '', NULL, 'plugins:content:content-category:export', 0, 90, '0,57,56,90,', '3', NULL, NULL, NULL, NULL, 1, 1, '2023-03-12 23:17:45', '2023-05-09 10:47:20');
INSERT INTO `admin_sys_menu` (`id`, `title`, `icon`, `path`, `element`, `redirect`, `permission`, `sort`, `parent_id`, `parent_ids`, `menu_type`, `is_keep_alive`, `is_affix`, `is_hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (96, '文章管理', 'BorderRightOutlined', '/plugins/content/content-article', '/plugins/content/content-article/index', NULL, '', 60, 56, '0,57,56,', '2', '1', '2', '2', '1', 1, 1, '2023-03-12 23:52:45', '2023-03-12 23:53:12');
INSERT INTO `admin_sys_menu` (`id`, `title`, `icon`, `path`, `element`, `redirect`, `permission`, `sort`, `parent_id`, `parent_ids`, `menu_type`, `is_keep_alive`, `is_affix`, `is_hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (97, '分页获取文章管理', '', '', '', NULL, 'plugins:content:content-article:query', 0, 96, '0,57,56,96,', '3', NULL, NULL, NULL, NULL, 1, 1, '2023-03-12 23:52:45', '2023-05-09 10:47:48');
INSERT INTO `admin_sys_menu` (`id`, `title`, `icon`, `path`, `element`, `redirect`, `permission`, `sort`, `parent_id`, `parent_ids`, `menu_type`, `is_keep_alive`, `is_affix`, `is_hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (98, '创建文章管理', '', '', '', NULL, 'plugins:content:content-article:add', 0, 96, '0,57,56,96,', '3', NULL, NULL, NULL, NULL, 1, 1, '2023-03-12 23:52:45', '2023-05-09 10:48:03');
INSERT INTO `admin_sys_menu` (`id`, `title`, `icon`, `path`, `element`, `redirect`, `permission`, `sort`, `parent_id`, `parent_ids`, `menu_type`, `is_keep_alive`, `is_affix`, `is_hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (99, '修改文章管理', '', '', '', NULL, 'plugins:content:content-article:edit', 0, 96, '0,57,56,96,', '3', NULL, NULL, NULL, NULL, 1, 1, '2023-03-12 23:52:45', '2023-05-09 10:48:27');
INSERT INTO `admin_sys_menu` (`id`, `title`, `icon`, `path`, `element`, `redirect`, `permission`, `sort`, `parent_id`, `parent_ids`, `menu_type`, `is_keep_alive`, `is_affix`, `is_hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (100, '删除文章管理', '', '', '', NULL, 'plugins:content:content-article:del', 0, 96, '0,57,56,96,', '3', NULL, NULL, NULL, NULL, 1, 1, '2023-03-12 23:52:45', '2023-05-09 10:48:37');
INSERT INTO `admin_sys_menu` (`id`, `title`, `icon`, `path`, `element`, `redirect`, `permission`, `sort`, `parent_id`, `parent_ids`, `menu_type`, `is_keep_alive`, `is_affix`, `is_hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (101, '导出文章管理', '', '', '', NULL, 'plugins:content:content-article:export', 0, 96, '0,57,56,96,', '3', NULL, NULL, NULL, NULL, 1, 1, '2023-03-12 23:52:46', '2023-05-09 10:48:53');
INSERT INTO `admin_sys_menu` (`id`, `title`, `icon`, `path`, `element`, `redirect`, `permission`, `sort`, `parent_id`, `parent_ids`, `menu_type`, `is_keep_alive`, `is_affix`, `is_hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (102, 'App管理', 'BorderRightOutlined', '/plugins/filemgr/filemgr-app', '/plugins/filemgr/filemgr-app/index', NULL, '', 0, 55, '0,57,55,', '2', '1', '2', '2', '1', 1, 1, '2023-03-13 00:55:02', '2023-03-13 00:55:52');
INSERT INTO `admin_sys_menu` (`id`, `title`, `icon`, `path`, `element`, `redirect`, `permission`, `sort`, `parent_id`, `parent_ids`, `menu_type`, `is_keep_alive`, `is_affix`, `is_hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (103, '分页获取App管理', '', '', '', NULL, 'plugins:filemgr:filemgr-app:query', 0, 102, '0,57,55,102,', '3', NULL, NULL, NULL, NULL, 1, 1, '2023-03-13 00:55:02', '2023-05-09 10:51:29');
INSERT INTO `admin_sys_menu` (`id`, `title`, `icon`, `path`, `element`, `redirect`, `permission`, `sort`, `parent_id`, `parent_ids`, `menu_type`, `is_keep_alive`, `is_affix`, `is_hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (104, '创建App管理', '', '', '', NULL, 'plugins:filemgr:filemgr-app:add', 0, 102, '0,57,55,102,', '3', NULL, NULL, NULL, NULL, 1, 1, '2023-03-13 00:55:02', '2023-05-09 10:51:41');
INSERT INTO `admin_sys_menu` (`id`, `title`, `icon`, `path`, `element`, `redirect`, `permission`, `sort`, `parent_id`, `parent_ids`, `menu_type`, `is_keep_alive`, `is_affix`, `is_hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (105, '修改App管理', '', '', '', NULL, 'plugins:filemgr:filemgr-app:edit', 0, 102, '0,57,55,102,', '3', NULL, NULL, NULL, NULL, 1, 1, '2023-03-13 00:55:02', '2023-05-09 10:51:56');
INSERT INTO `admin_sys_menu` (`id`, `title`, `icon`, `path`, `element`, `redirect`, `permission`, `sort`, `parent_id`, `parent_ids`, `menu_type`, `is_keep_alive`, `is_affix`, `is_hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (106, '删除App管理', '', '', '', NULL, 'plugins:filemgr:filemgr-app:del', 0, 102, '0,57,55,102,', '3', NULL, NULL, NULL, NULL, 1, 1, '2023-03-13 00:55:02', '2023-05-09 10:52:06');
INSERT INTO `admin_sys_menu` (`id`, `title`, `icon`, `path`, `element`, `redirect`, `permission`, `sort`, `parent_id`, `parent_ids`, `menu_type`, `is_keep_alive`, `is_affix`, `is_hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (107, '导出App管理', '', '', '', NULL, 'plugins:filemgr:filemgr-app:export', 0, 102, '0,57,55,102,', '3', NULL, NULL, NULL, NULL, 1, 1, '2023-03-13 00:55:02', '2023-05-09 10:52:17');
INSERT INTO `admin_sys_menu` (`id`, `title`, `icon`, `path`, `element`, `redirect`, `permission`, `sort`, `parent_id`, `parent_ids`, `menu_type`, `is_keep_alive`, `is_affix`, `is_hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (108, '国家区号', 'BorderRightOutlined', '/app/user/user-country-code', '/app/user/user-country-code/index', NULL, '', 150, 59, '0,58,59,', '2', '1', '2', '2', '1', 1, 1, '2023-03-14 17:47:44', '2023-03-14 18:06:00');
INSERT INTO `admin_sys_menu` (`id`, `title`, `icon`, `path`, `element`, `redirect`, `permission`, `sort`, `parent_id`, `parent_ids`, `menu_type`, `is_keep_alive`, `is_affix`, `is_hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (109, '分页获取国家电话区号', '', '', '', NULL, 'app:user:user-country-code:query', 0, 108, '0,58,59,108,', '3', NULL, NULL, NULL, NULL, 1, 1, '2023-03-14 17:47:44', '2023-05-09 10:31:00');
INSERT INTO `admin_sys_menu` (`id`, `title`, `icon`, `path`, `element`, `redirect`, `permission`, `sort`, `parent_id`, `parent_ids`, `menu_type`, `is_keep_alive`, `is_affix`, `is_hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (110, '创建国家电话区号', '', '', '', NULL, 'app:user:user-country-code:add', 0, 108, '0,58,59,108,', '3', NULL, NULL, NULL, NULL, 1, 1, '2023-03-14 17:47:44', '2023-05-09 10:31:10');
INSERT INTO `admin_sys_menu` (`id`, `title`, `icon`, `path`, `element`, `redirect`, `permission`, `sort`, `parent_id`, `parent_ids`, `menu_type`, `is_keep_alive`, `is_affix`, `is_hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (111, '修改国家电话区号', '', '', '', NULL, 'app:user:user-country-code:edit', 0, 108, '0,58,59,108,', '3', NULL, NULL, NULL, NULL, 1, 1, '2023-03-14 17:47:44', '2023-05-09 10:31:23');
INSERT INTO `admin_sys_menu` (`id`, `title`, `icon`, `path`, `element`, `redirect`, `permission`, `sort`, `parent_id`, `parent_ids`, `menu_type`, `is_keep_alive`, `is_affix`, `is_hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (112, '删除国家电话区号', '', '', '', NULL, 'app:user:user-country-code:del', 0, 108, '0,58,59,108,', '3', NULL, NULL, NULL, NULL, 1, 1, '2023-03-14 17:47:45', '2023-05-09 10:31:33');
INSERT INTO `admin_sys_menu` (`id`, `title`, `icon`, `path`, `element`, `redirect`, `permission`, `sort`, `parent_id`, `parent_ids`, `menu_type`, `is_keep_alive`, `is_affix`, `is_hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (113, '导出国家电话区号', '', '', '', NULL, 'app:user:user-country-code:export', 0, 108, '0,58,59,108,', '3', NULL, NULL, NULL, NULL, 1, 1, '2023-03-14 17:47:45', '2023-05-09 10:31:45');
INSERT INTO `admin_sys_menu` (`id`, `title`, `icon`, `path`, `element`, `redirect`, `permission`, `sort`, `parent_id`, `parent_ids`, `menu_type`, `is_keep_alive`, `is_affix`, `is_hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (114, '导出操作日志', 'ApartmentOutlined', '', '', '', 'admin:sys-oper-log:export', 0, 15, '0,1,15,', '3', '', '', '', '', 1, 1, '2023-05-09 11:02:50', '2024-12-09 15:32:19');
INSERT INTO `admin_sys_menu` (`id`, `title`, `icon`, `path`, `element`, `redirect`, `permission`, `sort`, `parent_id`, `parent_ids`, `menu_type`, `is_keep_alive`, `is_affix`, `is_hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (115, '登录日志导出', '', '', '', NULL, 'admin:sys-login-log:export', 0, 14, '0,1,14,', '3', NULL, NULL, NULL, NULL, 1, 1, '2023-05-09 11:04:20', '2023-05-09 11:04:20');
INSERT INTO `admin_sys_menu` (`id`, `title`, `icon`, `path`, `element`, `redirect`, `permission`, `sort`, `parent_id`, `parent_ids`, `menu_type`, `is_keep_alive`, `is_affix`, `is_hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (116, '导出数据', '', '', '', NULL, 'admin:sys-dict-type:export', 0, 12, '0,1,12,', '3', NULL, NULL, NULL, NULL, 1, 1, '2023-05-09 11:12:30', '2023-05-09 11:15:14');
INSERT INTO `admin_sys_menu` (`id`, `title`, `icon`, `path`, `element`, `redirect`, `permission`, `sort`, `parent_id`, `parent_ids`, `menu_type`, `is_keep_alive`, `is_affix`, `is_hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (117, '导出类型', 'AliyunOutlined', '', '', '', 'admin:sys-dict-type:export', 0, 11, '0,1,11,', '3', '', '', '', '', 1, 1, '2023-05-09 11:16:13', '2024-12-09 15:29:57');
INSERT INTO `admin_sys_menu` (`id`, `title`, `icon`, `path`, `element`, `redirect`, `permission`, `sort`, `parent_id`, `parent_ids`, `menu_type`, `is_keep_alive`, `is_affix`, `is_hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (118, '导出参数', 'BorderBottomOutlined', '', '', '', 'content:sys-config:export', 0, 13, '0,1,13,', '3', '', '', '', '', 1, 1, '2023-05-09 11:34:20', '2024-12-09 15:31:24');
INSERT INTO `admin_sys_menu` (`id`, `title`, `icon`, `path`, `element`, `redirect`, `permission`, `sort`, `parent_id`, `parent_ids`, `menu_type`, `is_keep_alive`, `is_affix`, `is_hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (119, '首页', 'CarFilled', '/home', '/admin/sys/home/index', '', '', 0, 0, '0,', '2', '2', '1', '2', '1', 1, 1, '2024-11-22 11:34:20', '2024-12-11 11:02:33');
INSERT INTO `admin_sys_menu` (`id`, `title`, `icon`, `path`, `element`, `redirect`, `permission`, `sort`, `parent_id`, `parent_ids`, `menu_type`, `is_keep_alive`, `is_affix`, `is_hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (120, '个人中心', 'CarFilled', '/profile', '/admin/sys/profile/index', '', '', 0, 0, '0,', '2', '2', '2', '1', '1', 1, 1, '2024-11-23 08:20:24', '2024-12-11 11:02:53');
INSERT INTO `admin_sys_menu` (`id`, `title`, `icon`, `path`, `element`, `redirect`, `permission`, `sort`, `parent_id`, `parent_ids`, `menu_type`, `is_keep_alive`, `is_affix`, `is_hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (121, '导出接口', 'AliwangwangOutlined', '', '', '', 'admin:sys-api:query', 0, 51, '0,1,51,', '3', '', '', '', '', 1, 1, '2024-12-09 14:59:51', '2024-12-09 14:59:51');
INSERT INTO `admin_sys_menu` (`id`, `title`, `icon`, `path`, `element`, `redirect`, `permission`, `sort`, `parent_id`, `parent_ids`, `menu_type`, `is_keep_alive`, `is_affix`, `is_hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (122, '修改管理员密码', 'AndroidOutlined', '', '', '', 'admin:sys-user:editPwd', 0, 2, '0,22,', '3', '', '', '', '', 1, 1, '2024-12-09 15:15:05', '2024-12-09 15:15:05');
INSERT INTO `admin_sys_menu` (`id`, `title`, `icon`, `path`, `element`, `redirect`, `permission`, `sort`, `parent_id`, `parent_ids`, `menu_type`, `is_keep_alive`, `is_affix`, `is_hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (123, '修改角色数据范围', 'AppstoreOutlined', '', '', '', 'admin:sys-role:datascope', 0, 8, '0,1,8,', '3', '', '', '', '', 1, 1, '2024-12-09 15:19:53', '2024-12-09 15:34:20');
INSERT INTO `admin_sys_menu` (`id`, `title`, `icon`, `path`, `element`, `redirect`, `permission`, `sort`, `parent_id`, `parent_ids`, `menu_type`, `is_keep_alive`, `is_affix`, `is_hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (124, '系统-删除接口', 'AmazonOutlined', '', '', '', 'admin:sys-api:del', 0, 51, '0,1,51,', '3', '', '', '', '', 1, 1, '2024-12-09 15:37:48', '2024-12-09 15:37:48');
COMMIT;

-- ----------------------------
-- Table structure for admin_sys_menu_api_rule
-- ----------------------------
DROP TABLE IF EXISTS `admin_sys_menu_api_rule`;
CREATE TABLE `admin_sys_menu_api_rule` (
  `admin_sys_menu_menu_id` int NOT NULL,
  `admin_sys_api_id` int NOT NULL COMMENT '主键编码',
  PRIMARY KEY (`admin_sys_menu_menu_id`,`admin_sys_api_id`),
  KEY `fk_admin_sys_menu_api_rule_admin_sys_api` (`admin_sys_api_id`),
  CONSTRAINT `fk_admin_sys_menu_api_rule_admin_sys_api` FOREIGN KEY (`admin_sys_api_id`) REFERENCES `admin_sys_api` (`id`),
  CONSTRAINT `fk_admin_sys_menu_api_rule_admin_sys_menu` FOREIGN KEY (`admin_sys_menu_menu_id`) REFERENCES `admin_sys_menu` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;

-- ----------------------------
-- Records of admin_sys_menu_api_rule
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for admin_sys_oper_log
-- ----------------------------
DROP TABLE IF EXISTS `admin_sys_oper_log`;
CREATE TABLE `admin_sys_oper_log` (
  `id` int NOT NULL AUTO_INCREMENT COMMENT '主键编码',
  `request_method` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '请求方式',
  `user_id` int DEFAULT NULL COMMENT '操作者',
  `oper_url` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '访问地址',
  `oper_ip` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '客户端ip',
  `oper_location` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '访问位置',
  `status` varchar(4) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '操作状态',
  `oper_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '操作时间',
  `json_result` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '返回数据',
  `remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '备注',
  `latency_time` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '耗时',
  `user_agent` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT 'ua',
  `created_at` datetime DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime DEFAULT NULL COMMENT '最后更新时间',
  `create_by` int DEFAULT NULL COMMENT '创建者',
  `update_by` int DEFAULT NULL COMMENT '更新者',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='操作日志';

-- ----------------------------
-- Records of admin_sys_oper_log
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for admin_sys_post
-- ----------------------------
DROP TABLE IF EXISTS `admin_sys_post`;
CREATE TABLE `admin_sys_post` (
  `id` int NOT NULL AUTO_INCREMENT,
  `post_name` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL,
  `post_code` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL,
  `sort` tinyint DEFAULT NULL,
  `status` char(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL,
  `remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL,
  `create_by` int DEFAULT NULL COMMENT '创建者',
  `update_by` int DEFAULT NULL COMMENT '更新者',
  `created_at` datetime DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime DEFAULT NULL COMMENT '最后更新时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;

-- ----------------------------
-- Records of admin_sys_post
-- ----------------------------
BEGIN;
INSERT INTO `admin_sys_post` (`id`, `post_name`, `post_code`, `sort`, `status`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (1, '首席执行官', 'CEO', 0, '1', '首席执行官', 1, 1, '2021-05-13 19:56:38', '2021-05-13 19:56:38');
INSERT INTO `admin_sys_post` (`id`, `post_name`, `post_code`, `sort`, `status`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (2, '首席技术执行官', 'CTO', 2, '1', '首席技术执行官', 1, 1, '2021-05-13 19:56:38', '2021-05-13 19:56:38');
INSERT INTO `admin_sys_post` (`id`, `post_name`, `post_code`, `sort`, `status`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (3, '首席运营官', 'COO', 3, '1', '测试工程师', 1, 1, '2021-05-13 19:56:38', '2021-05-13 19:56:38');
COMMIT;

-- ----------------------------
-- Table structure for admin_sys_role
-- ----------------------------
DROP TABLE IF EXISTS `admin_sys_role`;
CREATE TABLE `admin_sys_role` (
  `id` int NOT NULL AUTO_INCREMENT,
  `role_name` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL,
  `role_key` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL,
  `role_sort` bigint DEFAULT NULL,
  `remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL,
  `data_scope` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL,
  `status` char(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '状态 1-正常 2-停用',
  `create_by` int DEFAULT NULL COMMENT '创建者',
  `update_by` int DEFAULT NULL COMMENT '更新者',
  `created_at` datetime(3) DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime(3) DEFAULT NULL COMMENT '最后更新时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;

-- ----------------------------
-- Records of admin_sys_role
-- ----------------------------
BEGIN;
INSERT INTO `admin_sys_role` (`id`, `role_name`, `role_key`, `role_sort`, `remark`, `data_scope`, `status`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (1, '系统管理员', 'admin', 1, '', '', '1', 1, 1, '2021-05-13 19:56:37.913', '2023-03-03 01:04:03.641');
INSERT INTO `admin_sys_role` (`id`, `role_name`, `role_key`, `role_sort`, `remark`, `data_scope`, `status`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (2, 'test', 'test', 0, '', '', '1', 1, 1, '2023-04-27 14:33:47.437', '2024-12-09 16:46:50.001');
COMMIT;

-- ----------------------------
-- Table structure for admin_sys_role_dept
-- ----------------------------
DROP TABLE IF EXISTS `admin_sys_role_dept`;
CREATE TABLE `admin_sys_role_dept` (
  `role_id` smallint NOT NULL,
  `dept_id` smallint NOT NULL,
  PRIMARY KEY (`role_id`,`dept_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;

-- ----------------------------
-- Records of admin_sys_role_dept
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for admin_sys_role_menu
-- ----------------------------
DROP TABLE IF EXISTS `admin_sys_role_menu`;
CREATE TABLE `admin_sys_role_menu` (
  `role_id` int NOT NULL,
  `menu_id` int NOT NULL,
  PRIMARY KEY (`role_id`,`menu_id`),
  KEY `fk_admin_sys_role_menu_admin_sys_menu` (`menu_id`),
  CONSTRAINT `fk_admin_sys_role_menu_admin_sys_menu` FOREIGN KEY (`menu_id`) REFERENCES `admin_sys_menu` (`id`),
  CONSTRAINT `fk_admin_sys_role_menu_admin_sys_role` FOREIGN KEY (`role_id`) REFERENCES `admin_sys_role` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;

-- ----------------------------
-- Records of admin_sys_role_menu
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for admin_sys_user
-- ----------------------------
DROP TABLE IF EXISTS `admin_sys_user`;
CREATE TABLE `admin_sys_user` (
  `id` int NOT NULL AUTO_INCREMENT COMMENT '编码',
  `username` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '用户名',
  `password` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '密码',
  `nick_name` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '昵称',
  `phone` varchar(11) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '手机号',
  `role_id` int DEFAULT NULL COMMENT '角色ID',
  `salt` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '加盐',
  `avatar` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '头像',
  `sex` char(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '性别',
  `email` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '邮箱',
  `dept_id` int DEFAULT NULL COMMENT '部门',
  `post_id` int DEFAULT NULL COMMENT '岗位',
  `remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '备注',
  `status` varchar(4) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '状态',
  `create_by` int DEFAULT NULL COMMENT '创建者',
  `update_by` int DEFAULT NULL COMMENT '更新者',
  `created_at` datetime DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime DEFAULT NULL COMMENT '最后更新时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;

-- ----------------------------
-- Records of admin_sys_user
-- ----------------------------
BEGIN;
INSERT INTO `admin_sys_user` (`id`, `username`, `password`, `nick_name`, `phone`, `role_id`, `salt`, `avatar`, `sex`, `email`, `dept_id`, `post_id`, `remark`, `status`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (1, 'admin', '$2a$10$40Xa1HapSFE0kJdHV46LPebz6itTy60qfnXc3kFwTPV.qELEJ9k5q', 'admin', '13700000000', 1, '', 'http://www.bitxx.top/images/my_head-touch-icon-next.png', '1', 'admin@admin.com', 1, 1, '', '1', 1, 1, '2021-05-13 19:56:38', '2023-03-14 09:27:36');
INSERT INTO `admin_sys_user` (`id`, `username`, `password`, `nick_name`, `phone`, `role_id`, `salt`, `avatar`, `sex`, `email`, `dept_id`, `post_id`, `remark`, `status`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (2, 'test', '$2a$10$7RrDlHPBnnIpmjFEk9l4BusOVxqPrzk3mcxOX2h9EzI.YAmkzHTB6', 'test', '13711111111', 2, '', 'http://www.bitxx.top/images/my_head-touch-icon-next.png', '1', '13711111111@qq.com', 6, 1, '', '1', 1, 1, '2023-04-27 14:34:57', '2023-04-27 14:34:57');
COMMIT;

-- ----------------------------
-- Table structure for app_user
-- ----------------------------
DROP TABLE IF EXISTS `app_user`;
CREATE TABLE `app_user` (
  `id` int NOT NULL AUTO_INCREMENT COMMENT '用户编码',
  `level_id` int NOT NULL DEFAULT '1' COMMENT '用户等级编号',
  `user_name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '用户昵称',
  `true_name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '真实姓名',
  `money` decimal(30,18) NOT NULL DEFAULT '0.000000000000000000' COMMENT '余额',
  `email` varchar(300) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '电子邮箱',
  `mobile_title` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT '+86' COMMENT '用户手机号国家前缀',
  `mobile` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '手机号码',
  `avatar` varchar(1000) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '头像路径',
  `pay_pwd` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '提现密码',
  `pwd` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '登录密码',
  `ref_code` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '推荐码',
  `parent_id` int NOT NULL DEFAULT '0' COMMENT '父级编号',
  `parent_ids` varchar(1000) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '所有父级编号',
  `tree_sort` int NOT NULL DEFAULT '0' COMMENT '本级排序号（升序）',
  `tree_sorts` varchar(1000) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '0' COMMENT '所有级别排序号',
  `tree_leaf` char(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '0' COMMENT '是否最末级',
  `tree_level` int NOT NULL DEFAULT '0' COMMENT '层次级别',
  `status` char(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '2' COMMENT '状态(1-正常 2-异常)',
  `remark` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '备注信息',
  `create_by` int NOT NULL DEFAULT '0' COMMENT '创建者',
  `update_by` int NOT NULL DEFAULT '0' COMMENT '更新者',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='用户管理';

-- ----------------------------
-- Records of app_user
-- ----------------------------
BEGIN;
INSERT INTO `app_user` (`id`, `level_id`, `user_name`, `true_name`, `money`, `email`, `mobile_title`, `mobile`, `avatar`, `pay_pwd`, `pwd`, `ref_code`, `parent_id`, `parent_ids`, `tree_sort`, `tree_sorts`, `tree_leaf`, `tree_level`, `status`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (1, 1, '- -', '- -', 1.000000000000000000, 'fb0cc809bbed1743bd7d2d8f444e2bae099e69819f4e072f7057bb1e4249bf3d', '86', '6d84b6afd68a5c7188779114f16c46e9', 'http://www.bitxx.top/images/my_head-touch-icon-next.png', '', '', 'akIiWm', 0, '0,', 1, '1,', '2', 1, '1', '', 0, 1, '2023-04-03 21:09:13', '2023-10-19 14:03:37');
INSERT INTO `app_user` (`id`, `level_id`, `user_name`, `true_name`, `money`, `email`, `mobile_title`, `mobile`, `avatar`, `pay_pwd`, `pwd`, `ref_code`, `parent_id`, `parent_ids`, `tree_sort`, `tree_sorts`, `tree_leaf`, `tree_level`, `status`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (2, 2, '- -', '- -', 0.000000000000000000, 'dca887a13d1225ccd447dc52a712861c099e69819f4e072f7057bb1e4249bf3d', '86', '84ace68f39f53a315d8114c61413505d', 'http://www.bitxx.top/images/my_head-touch-icon-next.png', '', '', 'GQFz6v', 1, '0,1,', 1, '1,1,', '1', 2, '1', '', 0, 1, '2023-04-03 21:29:34', '2023-10-19 14:06:49');
INSERT INTO `app_user` (`id`, `level_id`, `user_name`, `true_name`, `money`, `email`, `mobile_title`, `mobile`, `avatar`, `pay_pwd`, `pwd`, `ref_code`, `parent_id`, `parent_ids`, `tree_sort`, `tree_sorts`, `tree_leaf`, `tree_level`, `status`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (3, 1, '- -', '- -', 0.000000000000000000, '4884f3537b62e668d33c6af76ddf6670099e69819f4e072f7057bb1e4249bf3d', '86', 'ff4273c3b1372055923122f9881b651b', 'http://www.bitxx.top/images/my_head-touch-icon-next.png', '', '', 'tT1Fbk', 1, '0,1,', 2, '1,2,', '1', 2, '1', '', 0, 1, '2023-04-03 21:29:35', '2023-10-19 14:06:37');
COMMIT;

-- ----------------------------
-- Table structure for app_user_account_log
-- ----------------------------
DROP TABLE IF EXISTS `app_user_account_log`;
CREATE TABLE `app_user_account_log` (
  `id` int NOT NULL AUTO_INCREMENT COMMENT '账变编号',
  `user_id` int NOT NULL COMMENT '用户编号',
  `change_money` decimal(10,2) NOT NULL DEFAULT '0.00' COMMENT '账变金额',
  `before_money` decimal(30,18) NOT NULL DEFAULT '0.000000000000000000' COMMENT '账变前金额',
  `after_money` decimal(30,18) NOT NULL DEFAULT '0.000000000000000000' COMMENT '账变后金额',
  `money_type` char(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '1' COMMENT '金额类型 1:余额 ',
  `change_type` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '1' COMMENT '帐变类型(1-类型1)',
  `status` char(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT '状态（1正常 2-异常）',
  `create_by` int NOT NULL COMMENT '创建者',
  `created_at` datetime NOT NULL COMMENT '创建时间',
  `update_by` int NOT NULL COMMENT '更新者',
  `updated_date` datetime NOT NULL COMMENT '更新时间',
  `remarks` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '备注信息',
  PRIMARY KEY (`id`),
  KEY `idx_qyc_user_status` (`status`)
) ENGINE=InnoDB AUTO_INCREMENT=10 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='账变记录';

-- ----------------------------
-- Records of app_user_account_log
-- ----------------------------
BEGIN;
INSERT INTO `app_user_account_log` (`id`, `user_id`, `change_money`, `before_money`, `after_money`, `money_type`, `change_type`, `status`, `create_by`, `created_at`, `update_by`, `updated_date`, `remarks`) VALUES (1, 1, 10.00, 0.000000000000010000, 20.000000000000000000, '1', '1', '1', 1, '2023-03-09 22:55:48', 1, '2023-03-09 22:55:51', NULL);
INSERT INTO `app_user_account_log` (`id`, `user_id`, `change_money`, `before_money`, `after_money`, `money_type`, `change_type`, `status`, `create_by`, `created_at`, `update_by`, `updated_date`, `remarks`) VALUES (2, 2, 10.00, 0.000000000000010000, 20.000000000000000000, '1', '1', '1', 1, '2023-03-09 22:55:48', 1, '2023-03-09 22:55:51', NULL);
INSERT INTO `app_user_account_log` (`id`, `user_id`, `change_money`, `before_money`, `after_money`, `money_type`, `change_type`, `status`, `create_by`, `created_at`, `update_by`, `updated_date`, `remarks`) VALUES (3, 1, 10.00, 0.000000000000010000, 20.000000000000000000, '1', '1', '1', 1, '2023-03-09 22:55:48', 1, '2023-03-09 22:55:51', NULL);
INSERT INTO `app_user_account_log` (`id`, `user_id`, `change_money`, `before_money`, `after_money`, `money_type`, `change_type`, `status`, `create_by`, `created_at`, `update_by`, `updated_date`, `remarks`) VALUES (4, 3, 10.00, 0.000000000000010000, 20.000000000000000000, '1', '1', '1', 1, '2023-03-09 22:55:48', 1, '2023-03-09 22:55:51', NULL);
INSERT INTO `app_user_account_log` (`id`, `user_id`, `change_money`, `before_money`, `after_money`, `money_type`, `change_type`, `status`, `create_by`, `created_at`, `update_by`, `updated_date`, `remarks`) VALUES (5, 1, 10.00, 0.000000000000010000, 20.000000000000000000, '1', '1', '1', 1, '2023-03-09 22:55:48', 1, '2023-03-09 22:55:51', NULL);
INSERT INTO `app_user_account_log` (`id`, `user_id`, `change_money`, `before_money`, `after_money`, `money_type`, `change_type`, `status`, `create_by`, `created_at`, `update_by`, `updated_date`, `remarks`) VALUES (6, 2, 10.00, 0.000000000000010000, 20.000000000000000000, '1', '1', '1', 1, '2023-03-09 22:55:48', 1, '2023-03-09 22:55:51', NULL);
INSERT INTO `app_user_account_log` (`id`, `user_id`, `change_money`, `before_money`, `after_money`, `money_type`, `change_type`, `status`, `create_by`, `created_at`, `update_by`, `updated_date`, `remarks`) VALUES (7, 1, 10.00, 0.000000000000010000, 20.000000000000000000, '1', '1', '1', 1, '2023-03-09 22:55:48', 1, '2023-03-09 22:55:51', NULL);
INSERT INTO `app_user_account_log` (`id`, `user_id`, `change_money`, `before_money`, `after_money`, `money_type`, `change_type`, `status`, `create_by`, `created_at`, `update_by`, `updated_date`, `remarks`) VALUES (8, 3, 10.00, 0.000000000000010000, 20.000000000000000000, '1', '1', '1', 1, '2023-03-09 22:55:48', 1, '2023-03-09 22:55:51', NULL);
INSERT INTO `app_user_account_log` (`id`, `user_id`, `change_money`, `before_money`, `after_money`, `money_type`, `change_type`, `status`, `create_by`, `created_at`, `update_by`, `updated_date`, `remarks`) VALUES (9, 1, 10.00, 0.000000000000010000, 20.000000000000000000, '1', '1', '1', 1, '2023-03-09 22:55:48', 1, '2023-03-09 22:55:51', NULL);
COMMIT;

-- ----------------------------
-- Table structure for app_user_conf
-- ----------------------------
DROP TABLE IF EXISTS `app_user_conf`;
CREATE TABLE `app_user_conf` (
  `id` int NOT NULL AUTO_INCREMENT,
  `user_id` int NOT NULL COMMENT '用户id',
  `can_login` char(1) NOT NULL DEFAULT '0' COMMENT '1-允许登陆；2-不允许登陆',
  `remark` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '备注信息',
  `status` char(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '2' COMMENT '状态（1-正常 2-异常）\n',
  `create_by` int NOT NULL DEFAULT '0' COMMENT '创建者',
  `update_by` int NOT NULL DEFAULT '0' COMMENT '更新者',
  `created_at` datetime NOT NULL COMMENT '创建时间',
  `updated_at` datetime NOT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb3 ROW_FORMAT=COMPACT COMMENT='用户配置';

-- ----------------------------
-- Records of app_user_conf
-- ----------------------------
BEGIN;
INSERT INTO `app_user_conf` (`id`, `user_id`, `can_login`, `remark`, `status`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (1, 1, '1', '', '1', 198, 198, '2023-04-03 21:09:13', '2023-04-03 21:09:13');
INSERT INTO `app_user_conf` (`id`, `user_id`, `can_login`, `remark`, `status`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (2, 2, '1', '', '1', 200, 200, '2023-04-03 21:29:34', '2023-04-03 21:29:34');
INSERT INTO `app_user_conf` (`id`, `user_id`, `can_login`, `remark`, `status`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (3, 3, '1', '', '1', 201, 201, '2023-04-03 21:29:35', '2023-04-03 21:29:35');
COMMIT;

-- ----------------------------
-- Table structure for app_user_country_code
-- ----------------------------
DROP TABLE IF EXISTS `app_user_country_code`;
CREATE TABLE `app_user_country_code` (
  `id` int NOT NULL AUTO_INCREMENT,
  `country` varchar(64) NOT NULL DEFAULT '' COMMENT '国家或地区',
  `code` varchar(12) NOT NULL DEFAULT '' COMMENT '区号',
  `status` char(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '2' COMMENT '状态(1-可用 2-停用)',
  `remark` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '备注信息',
  `create_by` int NOT NULL DEFAULT '0' COMMENT '创建者',
  `update_by` int NOT NULL DEFAULT '0' COMMENT '更新者',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=14 DEFAULT CHARSET=utf8mb3 ROW_FORMAT=COMPACT COMMENT='国家区号';

-- ----------------------------
-- Records of app_user_country_code
-- ----------------------------
BEGIN;
INSERT INTO `app_user_country_code` (`id`, `country`, `code`, `status`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (1, '新加坡', '65', '2', '', 1, 1, '2021-06-29 14:10:00', '2021-06-29 14:10:00');
INSERT INTO `app_user_country_code` (`id`, `country`, `code`, `status`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (2, '加拿大', '1', '2', '', 1, 1, '2021-06-29 14:10:21', '2021-06-29 14:10:21');
INSERT INTO `app_user_country_code` (`id`, `country`, `code`, `status`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (3, '韩国', '82', '2', '', 1, 1, '2021-06-29 14:10:36', '2021-06-29 14:10:36');
INSERT INTO `app_user_country_code` (`id`, `country`, `code`, `status`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (4, '日本', '81', '2', '', 1, 1, '2021-06-29 14:10:49', '2021-06-29 14:10:49');
INSERT INTO `app_user_country_code` (`id`, `country`, `code`, `status`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (5, '中国香港', '852', '2', '', 1, 1, '2021-06-29 14:11:02', '2021-06-29 14:11:02');
INSERT INTO `app_user_country_code` (`id`, `country`, `code`, `status`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (6, '中国澳门', '853', '2', '', 1, 1, '2021-06-29 14:11:15', '2021-06-29 14:11:15');
INSERT INTO `app_user_country_code` (`id`, `country`, `code`, `status`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (7, '中国台湾', '886', '2', '', 1, 1, '2021-06-29 14:11:25', '2021-06-29 14:11:25');
INSERT INTO `app_user_country_code` (`id`, `country`, `code`, `status`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (8, '泰国', '66', '2', '', 1, 1, '2021-06-29 14:11:36', '2021-06-29 14:11:36');
INSERT INTO `app_user_country_code` (`id`, `country`, `code`, `status`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (9, '缅甸', '95', '2', '', 1, 1, '2021-06-29 14:11:45', '2021-06-29 14:11:45');
INSERT INTO `app_user_country_code` (`id`, `country`, `code`, `status`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (10, '老挝', '856', '1', '', 1, 1, '2021-06-29 14:11:59', '2023-03-14 21:11:18');
INSERT INTO `app_user_country_code` (`id`, `country`, `code`, `status`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (11, '澳大利亚', '61', '2', '', 1, 1, '2021-06-29 14:12:14', '2021-06-29 14:12:14');
INSERT INTO `app_user_country_code` (`id`, `country`, `code`, `status`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (12, '俄罗斯', '7', '1', '', 1, 1, '2021-06-29 14:12:32', '2023-03-14 21:11:08');
INSERT INTO `app_user_country_code` (`id`, `country`, `code`, `status`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (13, '中国大陆', '86', '1', '', 1, 1, '2021-06-29 14:16:22', '2023-03-14 21:11:03');
COMMIT;

-- ----------------------------
-- Table structure for app_user_level
-- ----------------------------
DROP TABLE IF EXISTS `app_user_level`;
CREATE TABLE `app_user_level` (
  `id` int NOT NULL AUTO_INCREMENT COMMENT '主键',
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT '等级名称',
  `level_type` varchar(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '等级类型',
  `level` int NOT NULL COMMENT '等级',
  `status` char(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '2' COMMENT '状态(1-正常 2-异常)',
  `remark` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '备注信息',
  `create_by` int NOT NULL DEFAULT '0' COMMENT '创建者',
  `update_by` int NOT NULL DEFAULT '0' COMMENT '更新者',
  `created_at` datetime NOT NULL COMMENT '创建时间',
  `updated_at` datetime NOT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='用户等级';

-- ----------------------------
-- Records of app_user_level
-- ----------------------------
BEGIN;
INSERT INTO `app_user_level` (`id`, `name`, `level_type`, `level`, `status`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (1, 'test3', '2', 2, '1', '', 1, 1, '2023-03-09 17:05:24', '2023-03-09 17:05:24');
INSERT INTO `app_user_level` (`id`, `name`, `level_type`, `level`, `status`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (2, 'test34', '1', 1, '1', '', 1, 1, '2023-03-09 17:05:37', '2023-03-09 20:19:19');
COMMIT;

-- ----------------------------
-- Table structure for app_user_oper_log
-- ----------------------------
DROP TABLE IF EXISTS `app_user_oper_log`;
CREATE TABLE `app_user_oper_log` (
  `id` int NOT NULL AUTO_INCREMENT COMMENT '日志编码',
  `user_id` int NOT NULL DEFAULT '1' COMMENT '用户编号',
  `action_type` char(2) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '用户行为类型',
  `by_type` char(2) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '更新用户类型 1-app用户 2-后台用户',
  `status` char(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '2' COMMENT '状态(1-正常 2-异常)',
  `create_by` int NOT NULL DEFAULT '0' COMMENT '创建者',
  `update_by` int NOT NULL DEFAULT '0' COMMENT '更新者',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间',
  `remark` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '备注信息',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin ROW_FORMAT=DYNAMIC COMMENT='用户关键行为日志表';

-- ----------------------------
-- Records of app_user_oper_log
-- ----------------------------
BEGIN;
INSERT INTO `app_user_oper_log` (`id`, `user_id`, `action_type`, `by_type`, `status`, `create_by`, `update_by`, `created_at`, `updated_at`, `remark`) VALUES (1, 1, '', '2', '1', 1, 1, '2023-03-11 15:39:31', '2023-03-11 15:39:31', '');
INSERT INTO `app_user_oper_log` (`id`, `user_id`, `action_type`, `by_type`, `status`, `create_by`, `update_by`, `created_at`, `updated_at`, `remark`) VALUES (2, 2, '', '2', '1', 1, 1, '2023-03-11 15:41:16', '2023-03-11 15:41:16', '');
INSERT INTO `app_user_oper_log` (`id`, `user_id`, `action_type`, `by_type`, `status`, `create_by`, `update_by`, `created_at`, `updated_at`, `remark`) VALUES (3, 3, '', '1', '1', 1, 1, '2023-03-11 15:45:44', '2023-03-11 15:45:44', '');
INSERT INTO `app_user_oper_log` (`id`, `user_id`, `action_type`, `by_type`, `status`, `create_by`, `update_by`, `created_at`, `updated_at`, `remark`) VALUES (4, 1, '', '1', '1', 1, 1, '2023-03-11 15:46:13', '2023-03-11 15:46:13', '');
INSERT INTO `app_user_oper_log` (`id`, `user_id`, `action_type`, `by_type`, `status`, `create_by`, `update_by`, `created_at`, `updated_at`, `remark`) VALUES (5, 3, '2', '1', '1', 1, 1, '2023-03-11 15:54:05', '2023-03-11 15:54:05', '');
INSERT INTO `app_user_oper_log` (`id`, `user_id`, `action_type`, `by_type`, `status`, `create_by`, `update_by`, `created_at`, `updated_at`, `remark`) VALUES (6, 2, '1', '1', '1', 1, 1, '2023-03-11 15:56:36', '2023-03-11 15:56:36', '');
INSERT INTO `app_user_oper_log` (`id`, `user_id`, `action_type`, `by_type`, `status`, `create_by`, `update_by`, `created_at`, `updated_at`, `remark`) VALUES (7, 1, '2', '1', '1', 1, 1, '2023-03-11 16:03:35', '2023-03-11 16:03:35', '');
COMMIT;

-- ----------------------------
-- Table structure for plugins_content_announcement
-- ----------------------------
DROP TABLE IF EXISTS `plugins_content_announcement`;
CREATE TABLE `plugins_content_announcement` (
  `id` int NOT NULL AUTO_INCREMENT COMMENT '主键编码',
  `title` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '标题',
  `content` text CHARACTER SET utf8mb4 COLLATE utf8mb4_bin COMMENT '内容',
  `num` int DEFAULT NULL COMMENT '阅读次数',
  `remark` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '备注信息',
  `status` char(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '0' COMMENT '状态（0正常 1删除 2停用 3冻结）',
  `create_by` int NOT NULL COMMENT '创建者',
  `update_by` int NOT NULL COMMENT '更新者',
  `updated_at` datetime NOT NULL COMMENT '更新时间',
  `created_at` datetime NOT NULL COMMENT '创建时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='公告管理';

-- ----------------------------
-- Records of plugins_content_announcement
-- ----------------------------
BEGIN;
INSERT INTO `plugins_content_announcement` (`id`, `title`, `content`, `num`, `remark`, `status`, `create_by`, `update_by`, `updated_at`, `created_at`) VALUES (1, 'test', '<p>tes</p>', 4, 'test', '1', 1, 1, '2023-02-27 12:36:52', '2023-02-27 11:50:56');
INSERT INTO `plugins_content_announcement` (`id`, `title`, `content`, `num`, `remark`, `status`, `create_by`, `update_by`, `updated_at`, `created_at`) VALUES (2, 'test2', '<p>test</p>', 1, 'test', '1', 1, 1, '2023-02-27 23:49:05', '2023-02-27 23:49:05');
COMMIT;

-- ----------------------------
-- Table structure for plugins_content_article
-- ----------------------------
DROP TABLE IF EXISTS `plugins_content_article`;
CREATE TABLE `plugins_content_article` (
  `id` int NOT NULL AUTO_INCREMENT COMMENT '主键编码',
  `cate_id` int DEFAULT NULL COMMENT '分类编号',
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '名称',
  `content` text CHARACTER SET utf8mb4 COLLATE utf8mb4_bin COMMENT '内容',
  `remark` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '备注信息',
  `status` char(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '0' COMMENT '状态（1-正常 2-异常）',
  `create_by` int NOT NULL COMMENT '创建者',
  `update_by` int NOT NULL COMMENT '更新者',
  `updated_at` datetime NOT NULL COMMENT '更新时间',
  `created_at` datetime NOT NULL COMMENT '创建时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='文章管理';

-- ----------------------------
-- Records of plugins_content_article
-- ----------------------------
BEGIN;
INSERT INTO `plugins_content_article` (`id`, `cate_id`, `name`, `content`, `remark`, `status`, `create_by`, `update_by`, `updated_at`, `created_at`) VALUES (1, 1, 'test', '<p>test</p>', '111', '1', 1, 1, '2023-03-13 00:04:40', '2023-03-13 00:04:40');
COMMIT;

-- ----------------------------
-- Table structure for plugins_content_category
-- ----------------------------
DROP TABLE IF EXISTS `plugins_content_category`;
CREATE TABLE `plugins_content_category` (
  `id` int NOT NULL AUTO_INCREMENT COMMENT '主键编码',
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '名称',
  `status` char(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '0' COMMENT '状态（1-正常 2-异常）',
  `remark` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '备注信息',
  `create_by` int NOT NULL COMMENT '创建者',
  `update_by` int NOT NULL COMMENT '更新者',
  `updated_at` datetime NOT NULL COMMENT '更新时间',
  `created_at` datetime NOT NULL COMMENT '创建时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='文章分类管理';

-- ----------------------------
-- Records of plugins_content_category
-- ----------------------------
BEGIN;
INSERT INTO `plugins_content_category` (`id`, `name`, `status`, `remark`, `create_by`, `update_by`, `updated_at`, `created_at`) VALUES (1, 'test', '1', '', 1, 1, '2023-02-27 23:21:29', '2023-02-27 23:21:29');
INSERT INTO `plugins_content_category` (`id`, `name`, `status`, `remark`, `create_by`, `update_by`, `updated_at`, `created_at`) VALUES (2, 'test2', '1', '', 1, 1, '2023-02-27 23:22:00', '2023-02-27 23:22:00');
INSERT INTO `plugins_content_category` (`id`, `name`, `status`, `remark`, `create_by`, `update_by`, `updated_at`, `created_at`) VALUES (3, 'test23', '1', '', 1, 1, '2023-02-27 23:42:01', '2023-02-27 23:42:01');
COMMIT;

-- ----------------------------
-- Table structure for plugins_filemgr_app
-- ----------------------------
DROP TABLE IF EXISTS `plugins_filemgr_app`;
CREATE TABLE `plugins_filemgr_app` (
  `id` int NOT NULL AUTO_INCREMENT COMMENT '主键',
  `version` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '版本号',
  `platform` char(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '平台 (1-安卓 2-苹果)',
  `app_type` char(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '版本(1-默认)',
  `local_address` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '本地地址',
  `download_type` char(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '下载类型(1-本地 2-外链 3-oss )',
  `download_url` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '下载地址(download_type=1使用)',
  `remark` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '备注信息',
  `status` char(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT '状态（1-已发布 2-待发布）\n',
  `create_by` int NOT NULL COMMENT '创建者',
  `created_at` datetime NOT NULL COMMENT '创建时间',
  `update_by` int NOT NULL COMMENT '更新者',
  `updated_at` datetime NOT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin ROW_FORMAT=DYNAMIC COMMENT='app升级管理';

-- ----------------------------
-- Records of plugins_filemgr_app
-- ----------------------------
BEGIN;
INSERT INTO `plugins_filemgr_app` (`id`, `version`, `platform`, `app_type`, `local_address`, `download_type`, `download_url`, `remark`, `status`, `create_by`, `created_at`, `update_by`, `updated_at`) VALUES (1, '1.0.1', '1', '1', 'files/app/4b6ea3c0-d7fa-49f1-9d50-f9d73caad45f.apk', '3', '', 'test', '1', 1, '2023-03-12 11:34:54', 1, '2023-03-13 01:00:30');
INSERT INTO `plugins_filemgr_app` (`id`, `version`, `platform`, `app_type`, `local_address`, `download_type`, `download_url`, `remark`, `status`, `create_by`, `created_at`, `update_by`, `updated_at`) VALUES (2, '1.0.0', '1', '1', 'files/app/ba7b81c0-e6d2-42ee-82e4-2dcbec720c23.apk', '1', 'http://localhost:9999/files/app/ba7b81c0-e6d2-42ee-82e4-2dcbec720c23.apk', 'test', '1', 1, '2023-03-13 01:06:21', 1, '2023-03-13 01:06:21');
INSERT INTO `plugins_filemgr_app` (`id`, `version`, `platform`, `app_type`, `local_address`, `download_type`, `download_url`, `remark`, `status`, `create_by`, `created_at`, `update_by`, `updated_at`) VALUES (3, '1.0.2', '1', '1', '', '2', 'http://localhost:9999/test.apk', 'test2', '1', 1, '2023-03-13 01:07:00', 1, '2023-03-13 01:07:00');
INSERT INTO `plugins_filemgr_app` (`id`, `version`, `platform`, `app_type`, `local_address`, `download_type`, `download_url`, `remark`, `status`, `create_by`, `created_at`, `update_by`, `updated_at`) VALUES (4, '1.0.3', '1', '1', 'files/app/962bebc9-fdb6-41b5-b62b-b184ee2fd1c0.apk', '3', '', 'test2', '1', 1, '2023-03-13 01:07:24', 1, '2023-03-13 01:07:24');
COMMIT;

-- ----------------------------
-- Table structure for plugins_msg_code
-- ----------------------------
DROP TABLE IF EXISTS `plugins_msg_code`;
CREATE TABLE `plugins_msg_code` (
  `id` int NOT NULL AUTO_INCREMENT COMMENT '验证码编号',
  `user_id` int NOT NULL COMMENT '用户编号',
  `code` varchar(12) NOT NULL DEFAULT '0' COMMENT '验证码',
  `code_type` char(1) NOT NULL DEFAULT '0' COMMENT '验证码类型 1-邮箱；2-短信',
  `remark` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '备注异常',
  `status` char(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '2' COMMENT '验证码状态 1-发送成功 2-发送失败',
  `create_by` int NOT NULL DEFAULT '0' COMMENT '创建者',
  `update_by` int NOT NULL DEFAULT '0' COMMENT '更新者',
  `created_at` datetime NOT NULL COMMENT '创建时间',
  `updated_at` datetime NOT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3 ROW_FORMAT=COMPACT COMMENT='验证码记录';

-- ----------------------------
-- Records of plugins_msg_code
-- ----------------------------
BEGIN;
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
