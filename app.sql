/*
 Navicat Premium Data Transfer

 Source Server         : admin-prod
 Source Server Type    : MySQL
 Source Server Version : 80024 (8.0.24)
 Source Host           : 47.108.211.104:5997
 Source Schema         : app

 Target Server Type    : MySQL
 Target Server Version : 80024 (8.0.24)
 File Encoding         : 65001

 Date: 14/03/2023 23:01:44
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for app_user
-- ----------------------------
DROP TABLE IF EXISTS `app_user`;
CREATE TABLE `app_user`  (
  `id` int NOT NULL AUTO_INCREMENT COMMENT '用户编码',
  `level_id` int NOT NULL DEFAULT 1 COMMENT '用户等级编号',
  `user_name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '用户昵称',
  `true_name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '真实姓名',
  `money` decimal(30, 18) NOT NULL DEFAULT 0.000000000000000000 COMMENT '余额',
  `email` varchar(300) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL COMMENT '电子邮箱',
  `mobile_title` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT '+86' COMMENT '用户手机号国家前缀',
  `mobile` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL COMMENT '手机号码',
  `avatar` varchar(1000) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL COMMENT '头像路径',
  `pay_pwd` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '提现密码',
  `pwd` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '登录密码',
  `ref_code` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL COMMENT '推荐码',
  `parent_id` int NOT NULL DEFAULT 0 COMMENT '父级编号',
  `parent_ids` varchar(1000) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '所有父级编号',
  `tree_sort` int NOT NULL DEFAULT 0 COMMENT '本级排序号（升序）',
  `tree_sorts` varchar(1000) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '0' COMMENT '所有级别排序号',
  `tree_leaf` char(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '0' COMMENT '是否最末级',
  `tree_level` int NOT NULL DEFAULT 0 COMMENT '层次级别',
  `status` char(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '2' COMMENT '状态(1-正常 2-异常)',
  `remark` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL COMMENT '备注信息',
  `create_by` int NOT NULL DEFAULT 0 COMMENT '创建者',
  `update_by` int NOT NULL DEFAULT 0 COMMENT '更新者',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 197 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_bin COMMENT = '用户管理' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of app_user
-- ----------------------------

-- ----------------------------
-- Table structure for app_user_account_log
-- ----------------------------
DROP TABLE IF EXISTS `app_user_account_log`;
CREATE TABLE `app_user_account_log`  (
  `id` int NOT NULL AUTO_INCREMENT COMMENT '账变编号',
  `user_id` int NOT NULL COMMENT '用户编号',
  `change_money` decimal(10, 2) NOT NULL DEFAULT 0.00 COMMENT '账变金额',
  `before_money` decimal(30, 18) NOT NULL DEFAULT 0.000000000000000000 COMMENT '账变前金额',
  `after_money` decimal(30, 18) NOT NULL DEFAULT 0.000000000000000000 COMMENT '账变后金额',
  `money_type` char(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '1' COMMENT '金额类型 1:余额 ',
  `change_type` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '1' COMMENT '帐变类型(1-类型1)',
  `status` char(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT '状态（1正常 2-异常）',
  `create_by` int NOT NULL COMMENT '创建者',
  `created_at` datetime NOT NULL COMMENT '创建时间',
  `update_by` int NOT NULL COMMENT '更新者',
  `updated_date` datetime NOT NULL COMMENT '更新时间',
  `remarks` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL COMMENT '备注信息',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_qyc_user_status`(`status` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 11 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_bin COMMENT = '账变记录' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of app_user_account_log
-- ----------------------------

-- ----------------------------
-- Table structure for app_user_conf
-- ----------------------------
DROP TABLE IF EXISTS `app_user_conf`;
CREATE TABLE `app_user_conf`  (
  `id` int NOT NULL AUTO_INCREMENT,
  `user_id` int NOT NULL COMMENT '用户id',
  `can_login` char(1) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '0' COMMENT '1-允许登陆；2-不允许登陆',
  `remark` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL COMMENT '备注信息',
  `status` char(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '2' COMMENT '状态（1-正常 2-异常）\n',
  `create_by` int NOT NULL DEFAULT 0 COMMENT '创建者',
  `update_by` int NOT NULL DEFAULT 0 COMMENT '更新者',
  `created_at` datetime NOT NULL COMMENT '创建时间',
  `updated_at` datetime NOT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 176 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '用户配置' ROW_FORMAT = COMPACT;

-- ----------------------------
-- Records of app_user_conf
-- ----------------------------

-- ----------------------------
-- Table structure for app_user_country_code
-- ----------------------------
DROP TABLE IF EXISTS `app_user_country_code`;
CREATE TABLE `app_user_country_code`  (
  `id` int NOT NULL AUTO_INCREMENT,
  `country` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '国家或地区',
  `code` varchar(12) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '区号',
  `status` char(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '2' COMMENT '状态(1-可用 2-停用)',
  `remark` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL COMMENT '备注信息',
  `create_by` int NOT NULL DEFAULT 0 COMMENT '创建者',
  `update_by` int NOT NULL DEFAULT 0 COMMENT '更新者',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 14 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '国家区号' ROW_FORMAT = COMPACT;

-- ----------------------------
-- Records of app_user_country_code
-- ----------------------------
INSERT INTO `app_user_country_code` VALUES (1, '新加坡', '65', '2', '', 1, 1, '2021-06-29 14:10:00', '2021-06-29 14:10:00');
INSERT INTO `app_user_country_code` VALUES (2, '加拿大', '1', '2', '', 1, 1, '2021-06-29 14:10:21', '2021-06-29 14:10:21');
INSERT INTO `app_user_country_code` VALUES (3, '韩国', '82', '2', '', 1, 1, '2021-06-29 14:10:36', '2021-06-29 14:10:36');
INSERT INTO `app_user_country_code` VALUES (4, '日本', '81', '2', '', 1, 1, '2021-06-29 14:10:49', '2021-06-29 14:10:49');
INSERT INTO `app_user_country_code` VALUES (5, '中国香港', '852', '2', '', 1, 1, '2021-06-29 14:11:02', '2021-06-29 14:11:02');
INSERT INTO `app_user_country_code` VALUES (6, '中国澳门', '853', '2', '', 1, 1, '2021-06-29 14:11:15', '2021-06-29 14:11:15');
INSERT INTO `app_user_country_code` VALUES (7, '中国台湾', '886', '2', '', 1, 1, '2021-06-29 14:11:25', '2021-06-29 14:11:25');
INSERT INTO `app_user_country_code` VALUES (8, '泰国', '66', '2', '', 1, 1, '2021-06-29 14:11:36', '2021-06-29 14:11:36');
INSERT INTO `app_user_country_code` VALUES (9, '缅甸', '95', '2', '', 1, 1, '2021-06-29 14:11:45', '2021-06-29 14:11:45');
INSERT INTO `app_user_country_code` VALUES (10, '老挝', '856', '1', '', 1, 1, '2021-06-29 14:11:59', '2023-03-14 21:11:18');
INSERT INTO `app_user_country_code` VALUES (11, '澳大利亚', '61', '2', '', 1, 1, '2021-06-29 14:12:14', '2021-06-29 14:12:14');
INSERT INTO `app_user_country_code` VALUES (12, '俄罗斯', '7', '1', '', 1, 1, '2021-06-29 14:12:32', '2023-03-14 21:11:08');
INSERT INTO `app_user_country_code` VALUES (13, '中国大陆', '86', '1', '', 1, 1, '2021-06-29 14:16:22', '2023-03-14 21:11:03');

-- ----------------------------
-- Table structure for app_user_level
-- ----------------------------
DROP TABLE IF EXISTS `app_user_level`;
CREATE TABLE `app_user_level`  (
  `id` int NOT NULL AUTO_INCREMENT COMMENT '主键',
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT '等级名称',
  `level_type` varchar(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL COMMENT '等级类型',
  `level` int NOT NULL COMMENT '等级',
  `status` char(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '2' COMMENT '状态(1-正常 2-异常)',
  `remark` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL COMMENT '备注信息',
  `create_by` int NOT NULL DEFAULT 0 COMMENT '创建者',
  `update_by` int NOT NULL DEFAULT 0 COMMENT '更新者',
  `created_at` datetime NOT NULL COMMENT '创建时间',
  `updated_at` datetime NOT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 23 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_bin COMMENT = '用户等级' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of app_user_level
-- ----------------------------

-- ----------------------------
-- Table structure for app_user_oper_log
-- ----------------------------
DROP TABLE IF EXISTS `app_user_oper_log`;
CREATE TABLE `app_user_oper_log`  (
  `id` int NOT NULL AUTO_INCREMENT COMMENT '日志编码',
  `user_id` int NOT NULL DEFAULT 1 COMMENT '用户编号',
  `action_type` char(2) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '用户行为类型',
  `by_type` char(2) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '更新用户类型 1-app用户 2-后台用户',
  `status` char(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '2' COMMENT '状态(1-正常 2-异常)',
  `create_by` int NOT NULL DEFAULT 0 COMMENT '创建者',
  `update_by` int NOT NULL DEFAULT 0 COMMENT '更新者',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间',
  `remark` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL COMMENT '备注信息',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 94 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_bin COMMENT = '用户关键行为日志表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of app_user_oper_log
-- ----------------------------

-- ----------------------------
-- Table structure for plugins_content_announcement
-- ----------------------------
DROP TABLE IF EXISTS `plugins_content_announcement`;
CREATE TABLE `plugins_content_announcement`  (
  `id` int NOT NULL AUTO_INCREMENT COMMENT '主键编码',
  `title` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL COMMENT '标题',
  `content` text CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL COMMENT '内容',
  `num` int NULL DEFAULT NULL COMMENT '阅读次数',
  `remark` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL COMMENT '备注信息',
  `status` char(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '0' COMMENT '状态（0正常 1删除 2停用 3冻结）',
  `create_by` int NOT NULL COMMENT '创建者',
  `update_by` int NOT NULL COMMENT '更新者',
  `updated_at` datetime NOT NULL COMMENT '更新时间',
  `created_at` datetime NOT NULL COMMENT '创建时间',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_sys_content_create_by`(`create_by` ASC) USING BTREE,
  INDEX `idx_sys_content_update_by`(`update_by` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 7 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_bin COMMENT = '公告管理' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of plugins_content_announcement
-- ----------------------------

-- ----------------------------
-- Table structure for plugins_content_article
-- ----------------------------
DROP TABLE IF EXISTS `plugins_content_article`;
CREATE TABLE `plugins_content_article`  (
  `id` int NOT NULL AUTO_INCREMENT COMMENT '主键编码',
  `cate_id` int NULL DEFAULT NULL COMMENT '分类编号',
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL COMMENT '名称',
  `content` text CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL COMMENT '内容',
  `remark` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL COMMENT '备注信息',
  `status` char(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '0' COMMENT '状态（1-正常 2-异常）',
  `create_by` int NOT NULL COMMENT '创建者',
  `update_by` int NOT NULL COMMENT '更新者',
  `updated_at` datetime NOT NULL COMMENT '更新时间',
  `created_at` datetime NOT NULL COMMENT '创建时间',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_sys_content_create_by`(`create_by` ASC) USING BTREE,
  INDEX `idx_sys_content_update_by`(`update_by` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 22 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_bin COMMENT = '文章管理' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of plugins_content_article
-- ----------------------------

-- ----------------------------
-- Table structure for plugins_content_category
-- ----------------------------
DROP TABLE IF EXISTS `plugins_content_category`;
CREATE TABLE `plugins_content_category`  (
  `id` int NOT NULL AUTO_INCREMENT COMMENT '主键编码',
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL COMMENT '名称',
  `status` char(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '0' COMMENT '状态（1-正常 2-异常）',
  `remark` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL COMMENT '备注信息',
  `create_by` int NOT NULL COMMENT '创建者',
  `update_by` int NOT NULL COMMENT '更新者',
  `updated_at` datetime NOT NULL COMMENT '更新时间',
  `created_at` datetime NOT NULL COMMENT '创建时间',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_sys_category_create_by`(`create_by` ASC) USING BTREE,
  INDEX `idx_sys_category_update_by`(`update_by` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 14 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_bin COMMENT = '文章分类管理' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of plugins_content_category
-- ----------------------------

-- ----------------------------
-- Table structure for plugins_filemgr_app
-- ----------------------------
DROP TABLE IF EXISTS `plugins_filemgr_app`;
CREATE TABLE `plugins_filemgr_app`  (
  `id` int NOT NULL AUTO_INCREMENT COMMENT '主键',
  `version` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL COMMENT '版本号',
  `platform` char(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL COMMENT '平台 (1-安卓 2-苹果)',
  `app_type` char(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL COMMENT '版本(1-默认)',
  `local_address` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL COMMENT '本地地址',
  `download_num` int NULL DEFAULT 0 COMMENT '下载数量',
  `download_type` char(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL COMMENT '下载类型(1-本地 2-外链 3-oss )',
  `download_url` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL COMMENT '下载地址(download_type=1使用)',
  `remark` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL COMMENT '备注信息',
  `status` char(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT '状态（1-已发布 2-待发布）\n',
  `create_by` int NOT NULL COMMENT '创建者',
  `created_at` datetime NOT NULL COMMENT '创建时间',
  `update_by` int NOT NULL COMMENT '更新者',
  `updated_at` datetime NOT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 7 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_bin COMMENT = 'app升级管理' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of plugins_filemgr_app
-- ----------------------------

-- ----------------------------
-- Table structure for plugins_msg_code
-- ----------------------------
DROP TABLE IF EXISTS `plugins_msg_code`;
CREATE TABLE `plugins_msg_code`  (
  `id` int NOT NULL AUTO_INCREMENT COMMENT '验证码编号',
  `user_id` int NOT NULL COMMENT '用户编号',
  `code` varchar(12) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '0' COMMENT '验证码',
  `code_type` char(1) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '0' COMMENT '验证码类型 1-邮箱；2-短信',
  `remark` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL COMMENT '备注异常',
  `status` char(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '2' COMMENT '验证码状态 1-发送成功 2-发送失败',
  `create_by` int NOT NULL DEFAULT 0 COMMENT '创建者',
  `update_by` int NOT NULL DEFAULT 0 COMMENT '更新者',
  `created_at` datetime NOT NULL COMMENT '创建时间',
  `updated_at` datetime NOT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 174 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '验证码记录' ROW_FORMAT = COMPACT;

-- ----------------------------
-- Records of plugins_msg_code
-- ----------------------------

-- ----------------------------
-- Table structure for sys_api
-- ----------------------------
DROP TABLE IF EXISTS `sys_api`;
CREATE TABLE `sys_api`  (
  `id` int NOT NULL AUTO_INCREMENT COMMENT '主键编码',
  `handle` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL COMMENT 'handle',
  `title` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL COMMENT '标题',
  `path` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL COMMENT '地址',
  `api_type` varchar(16) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL COMMENT '接口类型',
  `action` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL COMMENT '请求类型',
  `created_at` datetime NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime NULL DEFAULT NULL COMMENT '最后更新时间',
  `create_by` int NULL DEFAULT NULL COMMENT '创建者',
  `update_by` int NULL DEFAULT NULL COMMENT '更新者',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_sys_api_create_by`(`create_by` ASC) USING BTREE,
  INDEX `idx_sys_api_update_by`(`update_by` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 260 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_bin ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of sys_api
-- ----------------------------

-- ----------------------------
-- Table structure for sys_casbin_rule
-- ----------------------------
DROP TABLE IF EXISTS `sys_casbin_rule`;
CREATE TABLE `sys_casbin_rule`  (
  `p_type` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `v0` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `v1` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `v2` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `v3` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `v4` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `v5` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  UNIQUE INDEX `idx_sys_casbin_rule`(`p_type` ASC, `v0` ASC, `v1` ASC, `v2` ASC, `v3` ASC, `v4` ASC, `v5` ASC) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of sys_casbin_rule
-- ----------------------------

-- ----------------------------
-- Table structure for sys_config
-- ----------------------------
DROP TABLE IF EXISTS `sys_config`;
CREATE TABLE `sys_config`  (
  `id` int NOT NULL AUTO_INCREMENT COMMENT '主键编码',
  `config_name` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL COMMENT 'ConfigName',
  `config_key` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL COMMENT 'ConfigKey',
  `config_value` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL COMMENT 'ConfigValue',
  `config_type` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL COMMENT 'ConfigType',
  `is_frontend` char(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL COMMENT '是否前台',
  `remark` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL COMMENT 'Remark',
  `create_by` int NULL DEFAULT NULL COMMENT '创建者',
  `update_by` int NULL DEFAULT NULL COMMENT '更新者',
  `created_at` datetime NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime NULL DEFAULT NULL COMMENT '最后更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_sys_config_create_by`(`create_by` ASC) USING BTREE,
  INDEX `idx_sys_config_update_by`(`update_by` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 94 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_bin ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of sys_config
-- ----------------------------
INSERT INTO `sys_config` VALUES (1, '系统内置-皮肤样式', 'sys_index_skinName', 'skin-green', '1', '1', '主框架页-默认皮肤样式名称:蓝色 skin-blue、绿色 skin-green、紫色 skin-purple、红色 skin-red、黄色 skin-yellow', 1, 1, '2021-05-13 19:56:38', '2023-03-11 23:16:02');
INSERT INTO `sys_config` VALUES (2, '系统内置-初始密码', 'sys_user_initPassword', '123456', '1', '1', '用户管理-账号初始密码:123456', 1, 1, '2021-05-13 19:56:38', '2023-03-11 23:16:10');
INSERT INTO `sys_config` VALUES (3, '系统内置-侧栏主题', 'sys_index_sideTheme', 'theme-dark', '1', '1', '主框架页-侧边栏主题:深色主题theme-dark，浅色主题theme-light', 1, 1, '2021-05-13 19:56:38', '2023-03-11 23:16:06');
INSERT INTO `sys_config` VALUES (4, '系统内置-系统名称', 'sys_app_name', 'go-admin后台管理系统', '1', '1', '', 1, 1, '2021-03-17 08:52:06', '2023-03-11 23:16:19');
INSERT INTO `sys_config` VALUES (5, '系统内置-系统logo', 'sys_app_logo', 'http://www.wjblog.top/images/my_head-touch-icon-next.png', '1', '1', '', 1, 1, '2021-03-17 08:53:19', '2023-03-11 23:16:15');
INSERT INTO `sys_config` VALUES (6, '系统内置-单次excel导出数据量', 'sys_max_export_size', '10000', '1', '1', '', 0, 1, '2021-07-28 16:53:48', '2023-03-11 23:15:56');
INSERT INTO `sys_config` VALUES (17, '插件-文件管理-App OSS Bucket', 'plugin_filemgr_app_oss_bucket', '自行配置，阿里OSS', '2', '2', '', 0, 1, '2021-08-13 14:36:23', '2023-03-14 21:35:06');
INSERT INTO `sys_config` VALUES (18, '插件-文件管理-App OSS AccessKeyId', 'plugin_filemgr_app_oss_access_key_id', '自行配置，阿里OSS', '2', '2', '', 0, 1, '2021-08-13 14:37:15', '2023-03-14 21:35:00');
INSERT INTO `sys_config` VALUES (19, '插件-文件管理-App OSS AccessKeySecret', 'plugin_filemgr_app_oss_access_key_secret', '自行配置，阿里OSS', '2', '2', '', 0, 1, '2021-08-13 14:38:00', '2023-03-14 21:34:53');
INSERT INTO `sys_config` VALUES (20, '插件-文件管理-App OSS Endpoint', 'plugin_filemgr_app_oss_endpoint', '自行配置，阿里OSS', '2', '2', '', 0, 1, '2021-08-13 14:38:50', '2023-03-14 21:34:46');
INSERT INTO `sys_config` VALUES (21, '插件-文件管理-App OSS 根目录', 'plugin_filemgr_app_oss_root_path', 'app/', '2', '2', '', 0, 1, '2021-08-13 14:39:31', '2023-03-14 21:34:22');
INSERT INTO `sys_config` VALUES (93, 'App-用户-默认头像', 'app_user_default_avatar', 'http://www.wjblog.top/images/my_head-touch-icon-next.png', '3', '2', '', 1, 1, '2023-03-10 18:07:03', '2023-03-10 18:07:03');

-- ----------------------------
-- Table structure for sys_dept
-- ----------------------------
DROP TABLE IF EXISTS `sys_dept`;
CREATE TABLE `sys_dept`  (
  `id` int NOT NULL AUTO_INCREMENT,
  `parent_id` int NULL DEFAULT NULL,
  `dept_path` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL,
  `dept_name` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL,
  `sort` int NULL DEFAULT NULL,
  `leader` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL,
  `phone` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL,
  `email` varchar(80) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL,
  `status` tinyint NULL DEFAULT NULL,
  `create_by` int NULL DEFAULT NULL COMMENT '创建者',
  `update_by` int NULL DEFAULT NULL COMMENT '更新者',
  `created_at` datetime NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime NULL DEFAULT NULL COMMENT '最后更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_sys_dept_create_by`(`create_by` ASC) USING BTREE,
  INDEX `idx_sys_dept_update_by`(`update_by` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 7 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_bin ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of sys_dept
-- ----------------------------
INSERT INTO `sys_dept` VALUES (1, 0, '0,', 'Admin', 0, 'admin', '', '', 1, 1, 1, '2021-05-13 19:56:38', '2022-05-14 11:20:25');
INSERT INTO `sys_dept` VALUES (2, 1, '0,1,', '研发部', 1, 'admin', '', '', 1, 1, 1, '2021-05-13 19:56:38', '2023-03-04 13:17:45');
INSERT INTO `sys_dept` VALUES (3, 1, '0,1,', '运维部', 1, 'admin', '', '', 1, 1, 1, '2021-05-13 19:56:38', '2023-03-04 13:17:45');
INSERT INTO `sys_dept` VALUES (4, 1, '0,1,', '客服部', 0, 'admin', '', '', 1, 1, 1, '2021-05-13 19:56:38', '2022-05-14 11:20:50');
INSERT INTO `sys_dept` VALUES (5, 1, '0,1,', '人力资源', 3, 'admin', '', '', 1, 1, 1, '2021-05-13 19:56:38', '2022-05-14 11:20:53');
INSERT INTO `sys_dept` VALUES (6, 1, '0,1,', '市场', 10, 'admin', '', '', 1, 1, 1, '2021-12-02 10:13:38', '2021-12-02 10:13:38');

-- ----------------------------
-- Table structure for sys_dict_data
-- ----------------------------
DROP TABLE IF EXISTS `sys_dict_data`;
CREATE TABLE `sys_dict_data`  (
  `id` int NOT NULL AUTO_INCREMENT,
  `dict_sort` int NULL DEFAULT NULL,
  `dict_label` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL,
  `dict_value` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL,
  `dict_type` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL,
  `css_class` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL,
  `list_class` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL,
  `is_default` varchar(8) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL,
  `status` char(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL,
  `default` varchar(8) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL,
  `remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL,
  `create_by` int NULL DEFAULT NULL COMMENT '创建者',
  `update_by` int NULL DEFAULT NULL COMMENT '更新者',
  `created_at` datetime NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime NULL DEFAULT NULL COMMENT '最后更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_sys_dict_data_create_by`(`create_by` ASC) USING BTREE,
  INDEX `idx_sys_dict_data_update_by`(`update_by` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 233 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_bin ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of sys_dict_data
-- ----------------------------
INSERT INTO `sys_dict_data` VALUES (1, 0, '正常', '2', 'sys_normal_disable', '', '', '', '0', '', '系统正常', 1, 1, '2021-05-13 19:56:38', '2022-04-25 00:42:38');
INSERT INTO `sys_dict_data` VALUES (2, 0, '停用', '1', 'sys_normal_disable', '', '', '', '0', '', '系统停用', 1, 1, '2021-05-13 19:56:38', '2021-05-13 19:56:38');
INSERT INTO `sys_dict_data` VALUES (3, 0, '男', '1', 'sys_user_sex', '', '', '', '0', '', '性别男', 1, 1, '2021-05-13 19:56:38', '2021-05-13 19:56:38');
INSERT INTO `sys_dict_data` VALUES (4, 0, '女', '2', 'sys_user_sex', '', '', '', '0', '', '性别女', 1, 1, '2021-05-13 19:56:38', '2021-05-13 19:56:38');
INSERT INTO `sys_dict_data` VALUES (5, 0, '未知', '3', 'sys_user_sex', '', '', '', '0', '', '性别未知', 1, 1, '2021-05-13 19:56:38', '2023-03-05 12:03:33');
INSERT INTO `sys_dict_data` VALUES (6, 0, '显示', '2', 'sys_menu_show_hide', '', '', '', '0', '', '显示菜单', 1, 1, '2021-05-13 19:56:38', '2021-05-13 19:56:38');
INSERT INTO `sys_dict_data` VALUES (7, 0, '隐藏', '1', 'sys_menu_show_hide', '', '', '', '0', '', '隐藏菜单', 1, 1, '2021-05-13 19:56:38', '2021-05-13 19:56:38');
INSERT INTO `sys_dict_data` VALUES (8, 0, '是', '1', 'sys_yes_no', '', '', '', '0', '', '系统默认是', 1, 1, '2021-05-13 19:56:38', '2021-05-13 19:56:38');
INSERT INTO `sys_dict_data` VALUES (9, 0, '否', '2', 'sys_yes_no', '', '', '', '0', '', '系统默认否', 1, 1, '2021-05-13 19:56:38', '2021-05-13 19:56:38');
INSERT INTO `sys_dict_data` VALUES (14, 0, '通知', '1', 'sys_notice_type', '', '', '', '0', '', '通知', 1, 1, '2021-05-13 19:56:38', '2021-05-13 19:56:38');
INSERT INTO `sys_dict_data` VALUES (15, 0, '公告', '2', 'sys_notice_type', '', '', '', '0', '', '公告', 1, 1, '2021-05-13 19:56:38', '2021-05-13 19:56:38');
INSERT INTO `sys_dict_data` VALUES (16, 0, '正常', '2', 'sys_common_status', '', '', '', '0', '', '正常状态', 1, 1, '2021-05-13 19:56:38', '2021-05-13 19:56:38');
INSERT INTO `sys_dict_data` VALUES (17, 0, '关闭', '1', 'sys_common_status', '', '', '', '0', '', '关闭状态', 1, 1, '2021-05-13 19:56:38', '2021-05-13 19:56:38');
INSERT INTO `sys_dict_data` VALUES (18, 0, '新增', '1', 'sys_oper_type', '', '', '', '0', '', '新增操作', 1, 1, '2021-05-13 19:56:38', '2021-05-13 19:56:38');
INSERT INTO `sys_dict_data` VALUES (19, 0, '修改', '2', 'sys_oper_type', '', '', '', '0', '', '修改操作', 1, 1, '2021-05-13 19:56:38', '2021-05-13 19:56:38');
INSERT INTO `sys_dict_data` VALUES (20, 0, '删除', '3', 'sys_oper_type', '', '', '', '0', '', '删除操作', 1, 1, '2021-05-13 19:56:38', '2021-05-13 19:56:38');
INSERT INTO `sys_dict_data` VALUES (21, 0, '授权', '4', 'sys_oper_type', '', '', '', '0', '', '授权操作', 1, 1, '2021-05-13 19:56:38', '2021-05-13 19:56:38');
INSERT INTO `sys_dict_data` VALUES (22, 0, '导出', '5', 'sys_oper_type', '', '', '', '0', '', '导出操作', 1, 1, '2021-05-13 19:56:38', '2021-05-13 19:56:38');
INSERT INTO `sys_dict_data` VALUES (23, 0, '导入', '6', 'sys_oper_type', '', '', '', '0', '', '导入操作', 1, 1, '2021-05-13 19:56:38', '2021-05-13 19:56:38');
INSERT INTO `sys_dict_data` VALUES (24, 0, '强退', '7', 'sys_oper_type', '', '', '', '0', '', '强退操作', 1, 1, '2021-05-13 19:56:38', '2021-05-13 19:56:38');
INSERT INTO `sys_dict_data` VALUES (25, 0, '生成代码', '8', 'sys_oper_type', '', '', '', '0', '', '生成操作', 1, 1, '2021-05-13 19:56:38', '2021-05-13 19:56:38');
INSERT INTO `sys_dict_data` VALUES (26, 0, '清空数据', '9', 'sys_oper_type', '', '', '', '0', '', '清空操作', 1, 1, '2021-05-13 19:56:38', '2021-05-13 19:56:38');
INSERT INTO `sys_dict_data` VALUES (27, 0, '成功', '1', 'sys_notice_status', '', '', '', '0', '', '成功状态', 1, 1, '2021-05-13 19:56:38', '2021-05-13 19:56:38');
INSERT INTO `sys_dict_data` VALUES (28, 0, '失败', '2', 'sys_notice_status', '', '', '', '0', '', '失败状态', 1, 1, '2021-05-13 19:56:38', '2021-05-13 19:56:38');
INSERT INTO `sys_dict_data` VALUES (29, 0, '登录', '10', 'sys_oper_type', '', '', '', '0', '', '登录操作', 1, 1, '2021-05-13 19:56:38', '2021-05-13 19:56:38');
INSERT INTO `sys_dict_data` VALUES (30, 0, '退出', '11', 'sys_oper_type', '', '', '', '0', '', '', 1, 1, '2021-05-13 19:56:38', '2021-05-13 19:56:38');
INSERT INTO `sys_dict_data` VALUES (31, 0, '获取验证码', '12', 'sys_oper_type', '', '', '', '0', '', '获取验证码', 1, 1, '2021-05-13 19:56:38', '2021-05-13 19:56:38');
INSERT INTO `sys_dict_data` VALUES (32, 0, '正常', '1', 'sys_status', '', '', '', '0', '', '', 0, 0, '2021-07-09 11:40:01', '2021-07-09 11:40:01');
INSERT INTO `sys_dict_data` VALUES (33, 0, '停用', '2', 'sys_status', '', '', '', '0', '', '', 0, 0, '2021-07-09 11:40:14', '2021-07-09 11:40:14');
INSERT INTO `sys_dict_data` VALUES (136, 0, '安卓', '1', 'plugin_filemgr_app_platform', '', '', '', '2', '', '', 0, 0, '2021-08-13 13:35:39', '2021-08-13 13:35:39');
INSERT INTO `sys_dict_data` VALUES (137, 0, 'IOS', '2', 'plugin_filemgr_app_platform', '', '', '', '2', '', '', 0, 0, '2021-08-13 13:35:51', '2021-08-13 13:35:51');
INSERT INTO `sys_dict_data` VALUES (138, 0, '类型1', '1', 'plugin_filemgr_app_type', '', '', '', '2', '', '', 0, 0, '2021-08-13 13:37:07', '2021-08-13 13:37:07');
INSERT INTO `sys_dict_data` VALUES (139, 0, '类型2', '2', 'plugin_filemgr_app_type', '', '', '', '2', '', '', 0, 0, '2021-08-13 13:37:19', '2021-08-13 13:37:19');
INSERT INTO `sys_dict_data` VALUES (140, 0, '类型3', '3', 'plugin_filemgr_app_type', '', '', '', '2', '', '', 0, 0, '2021-08-13 13:37:39', '2021-08-13 13:37:39');
INSERT INTO `sys_dict_data` VALUES (141, 0, '本地', '1', 'plugin_filemgr_app_download_type', '', '', '', '2', '', '', 0, 0, '2021-08-13 14:02:44', '2021-08-13 14:02:44');
INSERT INTO `sys_dict_data` VALUES (142, 0, '外链', '2', 'plugin_filemgr_app_download_type', '', '', '', '2', '', '', 0, 0, '2021-08-13 14:02:44', '2021-08-13 14:02:44');
INSERT INTO `sys_dict_data` VALUES (143, 0, 'OSS', '3', 'plugin_filemgr_app_download_type', '', '', '', '2', '', '', 0, 0, '2021-08-13 14:02:33', '2021-08-13 14:02:33');
INSERT INTO `sys_dict_data` VALUES (145, 0, '已发布', '2', 'plugin_filemgr_app_publish_status', '', '', '', '2', '', '', 0, 0, '2021-12-09 12:42:47', '2021-12-09 12:42:47');
INSERT INTO `sys_dict_data` VALUES (146, 0, '待发布', '1', 'plugin_filemgr_app_publish_status', '', '', '', '2', '', '', 0, 0, '2021-12-09 12:42:54', '2021-12-09 12:42:54');
INSERT INTO `sys_dict_data` VALUES (178, 0, '插件', '2', 'sys_api_type', '', '', '', '0', '', '', 1, 1, '2022-04-25 23:58:24', '2023-03-01 21:45:53');
INSERT INTO `sys_dict_data` VALUES (179, 0, '系统', '1', 'sys_api_type', '', '', '', '0', '', '', 1, 1, '2022-04-25 23:58:41', '2023-03-01 21:45:41');
INSERT INTO `sys_dict_data` VALUES (180, 0, 'GET', 'GET', 'sys_api_action', '', '', '', '0', '', '', 1, 1, '2022-04-26 00:03:26', '2022-04-26 00:03:26');
INSERT INTO `sys_dict_data` VALUES (181, 0, 'POST', 'POST', 'sys_api_action', '', '', '', '0', '', '', 1, 1, '2022-04-26 00:03:40', '2022-04-26 00:03:40');
INSERT INTO `sys_dict_data` VALUES (182, 0, 'DELETE', 'DELETE', 'sys_api_action', '', '', '', '0', '', '', 1, 1, '2022-04-26 00:03:49', '2022-04-26 00:03:49');
INSERT INTO `sys_dict_data` VALUES (183, 0, 'PUT', 'PUT', 'sys_api_action', '', '', '', '0', '', '', 1, 1, '2022-04-26 00:04:06', '2022-04-26 00:04:06');
INSERT INTO `sys_dict_data` VALUES (184, 0, 'HEAD', 'HEAD', 'sys_api_action', '', '', '', '0', '', '', 1, 1, '2022-04-26 00:07:02', '2022-04-26 00:07:02');
INSERT INTO `sys_dict_data` VALUES (188, 0, '系统内置', '1', 'sys_config_type', '', '', '', '1', '', '', 1, 1, '2023-03-01 11:05:23', '2023-03-01 11:05:23');
INSERT INTO `sys_dict_data` VALUES (189, 0, '插件', '2', 'sys_config_type', '', '', '', '1', '', '', 1, 1, '2023-03-01 11:05:32', '2023-03-01 11:05:32');
INSERT INTO `sys_dict_data` VALUES (190, 0, '应用', '3', 'sys_config_type', '', '', '', '1', '', '', 1, 1, '2023-03-01 11:05:42', '2023-03-01 11:05:42');
INSERT INTO `sys_dict_data` VALUES (191, 0, '展示', '1', 'sys_config_is_frontend', '', '', '', '1', '', '', 1, 1, '2023-03-01 11:07:49', '2023-03-01 11:07:49');
INSERT INTO `sys_dict_data` VALUES (192, 0, '隐藏', '2', 'sys_config_is_frontend', '', '', '', '1', '', '', 1, 1, '2023-03-01 11:07:56', '2023-03-01 11:07:56');
INSERT INTO `sys_dict_data` VALUES (195, 0, '登录', '1', 'sys_loginlog_status', '', '', '', '1', '', '', 1, 1, '2023-03-01 14:43:04', '2023-03-01 14:43:04');
INSERT INTO `sys_dict_data` VALUES (196, 0, '退出', '2', 'sys_loginlog_status', '', '', '', '1', '', '', 1, 1, '2023-03-01 14:43:10', '2023-03-01 14:43:10');
INSERT INTO `sys_dict_data` VALUES (197, 0, '应用', '3', 'sys_api_type', '', '', '', '1', '', '', 1, 1, '2023-03-01 21:46:01', '2023-03-01 21:46:01');
INSERT INTO `sys_dict_data` VALUES (200, 0, '全部数据权限', '1', 'sys_role_data_scope', '', '', '', '1', '', '', 1, 1, '2023-03-04 13:29:36', '2023-03-04 13:29:36');
INSERT INTO `sys_dict_data` VALUES (201, 0, '自定数据权限', '2', 'sys_role_data_scope', '', '', '', '1', '', '', 1, 1, '2023-03-04 13:29:43', '2023-03-04 13:29:43');
INSERT INTO `sys_dict_data` VALUES (202, 0, '本部门数据权限', '3', 'sys_role_data_scope', '', '', '', '1', '', '', 1, 1, '2023-03-04 13:29:49', '2023-03-04 13:29:49');
INSERT INTO `sys_dict_data` VALUES (203, 0, '本部门及以下数据权限', '4', 'sys_role_data_scope', '', '', '', '1', '', '', 1, 1, '2023-03-04 13:29:56', '2023-03-04 13:29:56');
INSERT INTO `sys_dict_data` VALUES (204, 0, '仅本人数据权限', '5', 'sys_role_data_scope', '', '', '', '1', '', '', 1, 1, '2023-03-04 13:30:04', '2023-03-04 13:30:04');
INSERT INTO `sys_dict_data` VALUES (205, 0, 'int64', 'int64', 'sys_gen_go_type', '', '', '', '1', '', '', 1, 1, '2023-03-07 10:08:26', '2023-03-07 10:08:26');
INSERT INTO `sys_dict_data` VALUES (206, 0, 'int', 'int', 'sys_gen_go_type', '', '', '', '1', '', '', 1, 1, '2023-03-07 10:12:42', '2023-03-07 10:12:42');
INSERT INTO `sys_dict_data` VALUES (207, 0, 'string', 'string', 'sys_gen_go_type', '', '', '', '1', '', '', 1, 1, '2023-03-07 10:13:05', '2023-03-07 10:13:05');
INSERT INTO `sys_dict_data` VALUES (208, 0, 'decimal', 'decimal.Decimal', 'sys_gen_go_type', '', '', '', '1', '', '', 1, 1, '2023-03-07 10:13:16', '2023-03-07 10:13:29');
INSERT INTO `sys_dict_data` VALUES (209, 0, 'time', '*time.Time', 'sys_gen_go_type', '', '', '', '1', '', '', 1, 1, '2023-03-07 10:13:43', '2023-03-07 10:13:43');
INSERT INTO `sys_dict_data` VALUES (210, 0, '=', 'EQ', 'sys_gen_query_type', '', '', '', '1', '', '', 1, 1, '2023-03-07 10:20:53', '2023-03-07 10:20:53');
INSERT INTO `sys_dict_data` VALUES (211, 0, '!=', 'NE', 'sys_gen_query_type', '', '', '', '1', '', '', 1, 1, '2023-03-07 10:21:06', '2023-03-07 10:21:06');
INSERT INTO `sys_dict_data` VALUES (212, 0, '>', 'GT', 'sys_gen_query_type', '', '', '', '1', '', '', 1, 1, '2023-03-07 10:21:20', '2023-03-07 10:21:20');
INSERT INTO `sys_dict_data` VALUES (213, 0, '>=', 'GTE', 'sys_gen_query_type', '', '', '', '1', '', '', 1, 1, '2023-03-07 10:21:33', '2023-03-07 10:21:33');
INSERT INTO `sys_dict_data` VALUES (214, 0, '<', 'LT', 'sys_gen_query_type', '', '', '', '1', '', '', 1, 1, '2023-03-07 10:21:45', '2023-03-07 10:21:45');
INSERT INTO `sys_dict_data` VALUES (215, 0, '<=', 'LTE', 'sys_gen_query_type', '', '', '', '1', '', '', 1, 1, '2023-03-07 10:21:57', '2023-03-07 10:21:57');
INSERT INTO `sys_dict_data` VALUES (216, 0, 'LIKE', 'LIKE', 'sys_gen_query_type', '', '', '', '1', '', '', 1, 1, '2023-03-07 10:22:08', '2023-03-07 10:22:08');
INSERT INTO `sys_dict_data` VALUES (217, 0, '文本框', 'input', 'sys_gen_html_type', '', '', '', '1', '', '', 1, 1, '2023-03-07 10:23:39', '2023-03-07 10:23:39');
INSERT INTO `sys_dict_data` VALUES (218, 0, '下拉框', 'select', 'sys_gen_html_type', '', '', '', '1', '', '', 1, 1, '2023-03-07 10:23:49', '2023-03-07 10:23:49');
INSERT INTO `sys_dict_data` VALUES (219, 0, '单选框', 'radio', 'sys_gen_html_type', '', '', '', '1', '', '', 1, 1, '2023-03-07 10:23:59', '2023-03-07 10:23:59');
INSERT INTO `sys_dict_data` VALUES (220, 0, '文本域', 'textarea', 'sys_gen_html_type', '', '', '', '1', '', '', 1, 1, '2023-03-07 10:24:08', '2023-03-07 10:24:08');
INSERT INTO `sys_dict_data` VALUES (221, 0, '目录', '1', 'sys_menu_type', '', '', '', '1', '', '', 1, 1, '2023-03-08 10:42:00', '2023-03-08 10:42:14');
INSERT INTO `sys_dict_data` VALUES (222, 0, '菜单', '2', 'sys_menu_type', '', '', '', '1', '', '', 1, 1, '2023-03-08 10:42:10', '2023-03-08 10:42:10');
INSERT INTO `sys_dict_data` VALUES (223, 0, '按钮', '3', 'sys_menu_type', '', '', '', '1', '', '', 1, 1, '2023-03-08 10:42:22', '2023-03-08 10:42:22');
INSERT INTO `sys_dict_data` VALUES (224, 0, '类型1', '1', 'app_user_level_type', '', '', '', '1', '', '', 1, 1, '2023-03-08 11:55:57', '2023-03-08 11:55:57');
INSERT INTO `sys_dict_data` VALUES (225, 0, '类型2', '2', 'app_user_level_type', '', '', '', '1', '', '', 1, 1, '2023-03-08 11:56:02', '2023-03-08 11:56:02');
INSERT INTO `sys_dict_data` VALUES (226, 0, '数字文本框', 'numInput', 'sys_gen_html_type', '', '', '', '1', '', '', 1, 1, '2023-03-09 20:12:33', '2023-03-09 20:12:33');
INSERT INTO `sys_dict_data` VALUES (227, 0, 'CNY', '1', 'app_money_type', '', '', '', '1', '', '', 1, 1, '2023-03-09 20:24:26', '2023-03-09 20:24:26');
INSERT INTO `sys_dict_data` VALUES (228, 0, '类型1', '1', 'app_account_change_type', '', '', '', '1', '', '', 1, 1, '2023-03-09 20:27:45', '2023-03-09 20:27:45');
INSERT INTO `sys_dict_data` VALUES (229, 0, '允许用户登录', '1', 'app_user_action_type', '', '', '', '1', '', '', 1, 1, '2023-03-11 14:08:01', '2023-03-11 14:08:01');
INSERT INTO `sys_dict_data` VALUES (230, 0, '禁止用户登录', '2', 'app_user_action_type', '', '', '', '1', '', '', 1, 1, '2023-03-11 14:08:10', '2023-03-11 14:08:10');
INSERT INTO `sys_dict_data` VALUES (231, 0, '后台用户', '1', 'app_user_by_type', '', '', '', '1', '', '', 1, 1, '2023-03-11 14:14:41', '2023-03-11 14:14:41');
INSERT INTO `sys_dict_data` VALUES (232, 0, '前台用户', '2', 'app_user_by_type', '', '', '', '1', '', '', 1, 1, '2023-03-11 14:14:59', '2023-03-11 14:14:59');

-- ----------------------------
-- Table structure for sys_dict_type
-- ----------------------------
DROP TABLE IF EXISTS `sys_dict_type`;
CREATE TABLE `sys_dict_type`  (
  `id` int NOT NULL AUTO_INCREMENT,
  `dict_name` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL,
  `dict_type` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL,
  `status` char(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL,
  `remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL,
  `create_by` int NULL DEFAULT NULL COMMENT '创建者',
  `update_by` int NULL DEFAULT NULL COMMENT '更新者',
  `created_at` datetime NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime NULL DEFAULT NULL COMMENT '最后更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_sys_dict_type_create_by`(`create_by` ASC) USING BTREE,
  INDEX `idx_sys_dict_type_update_by`(`update_by` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 93 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_bin ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of sys_dict_type
-- ----------------------------
INSERT INTO `sys_dict_type` VALUES (1, '系统-开关', 'sys_normal_disable', '0', '系统开关列表', 1, 1, '2021-05-13 19:56:38', '2023-03-11 23:20:35');
INSERT INTO `sys_dict_type` VALUES (2, '系统-用户性别', 'sys_user_sex', '0', '用户性别列表', 1, 1, '2021-05-13 19:56:38', '2023-03-11 23:21:06');
INSERT INTO `sys_dict_type` VALUES (3, '系统-菜单状态', 'sys_menu_show_hide', '0', '菜单状态列表', 1, 1, '2021-05-13 19:56:38', '2023-03-11 23:21:02');
INSERT INTO `sys_dict_type` VALUES (4, '系统-是否', 'sys_yes_no', '0', '系统是否列表', 1, 1, '2021-05-13 19:56:38', '2023-03-11 23:20:58');
INSERT INTO `sys_dict_type` VALUES (7, '系统-通知类型', 'sys_notice_type', '0', '通知类型列表', 1, 1, '2021-05-13 19:56:38', '2023-03-11 23:20:53');
INSERT INTO `sys_dict_type` VALUES (8, '系统-状态', 'sys_common_status', '0', '登录状态列表', 1, 1, '2021-05-13 19:56:38', '2023-03-11 23:20:49');
INSERT INTO `sys_dict_type` VALUES (9, '系统-操作类型', 'sys_oper_type', '0', '操作类型列表', 1, 1, '2021-05-13 19:56:38', '2023-03-11 23:20:42');
INSERT INTO `sys_dict_type` VALUES (10, '系统-通知状态', 'sys_notice_status', '0', '通知状态列表', 1, 1, '2021-05-13 19:56:38', '2023-03-11 23:20:39');
INSERT INTO `sys_dict_type` VALUES (11, '系统-基本状态', 'sys_status', '0', '基本通用状态', 1, 1, '2021-07-09 11:39:21', '2023-03-11 23:21:23');
INSERT INTO `sys_dict_type` VALUES (51, '插件-文件管理-App发布状态', 'plugin_filemgr_publish_status', '2', '', 1, 1, '2021-12-09 12:42:31', '2023-03-11 23:20:01');
INSERT INTO `sys_dict_type` VALUES (58, '插件-文件管理-App系统平台', 'plugin_filemgr_app_platform', '0', 'App系统平台', 1, 1, '2021-08-13 13:36:40', '2023-03-11 23:20:17');
INSERT INTO `sys_dict_type` VALUES (59, '插件-文件管理-App类型', 'plugin_filemgr_app_type', '0', 'app属性', 1, 1, '2021-08-13 13:36:40', '2023-03-11 23:20:13');
INSERT INTO `sys_dict_type` VALUES (60, '插件-文件管理-App下载类型', 'plugin_filemgr_app_download_type', '0', '', 1, 1, '2021-08-13 14:02:03', '2023-03-11 23:20:06');
INSERT INTO `sys_dict_type` VALUES (65, '系统-接口-类型', 'sys_api_type', '0', '系统', 1, 1, '2022-04-25 23:57:17', '2023-03-01 21:56:34');
INSERT INTO `sys_dict_type` VALUES (66, '系统-接口-请求方法', 'sys_api_action', '0', '', 1, 1, '2022-04-26 00:03:11', '2023-03-01 21:56:41');
INSERT INTO `sys_dict_type` VALUES (75, '系统-配置-类型', 'sys_config_type', '1', '1-内置 2-插件 3-应用', 1, 1, '2023-03-01 11:04:56', '2023-03-01 11:08:27');
INSERT INTO `sys_dict_type` VALUES (76, '系统-配置-是否前台展示', 'sys_config_is_frontend', '1', '1-展示 2-隐藏', 1, 1, '2023-03-01 11:06:28', '2023-03-01 11:08:07');
INSERT INTO `sys_dict_type` VALUES (78, '系统-登录日志-日志状态', 'sys_loginlog_status', '1', '1-登录 2-退出', 1, 1, '2023-03-01 14:42:56', '2023-03-01 14:42:56');
INSERT INTO `sys_dict_type` VALUES (81, '系统-角色-数据范围', 'sys_role_data_scope', '1', '1-全部数据权限 2- 自定义数据权限 3-本部门数据权限 4-本部门及以下数据权限 5-仅本人数据权限', 1, 1, '2023-03-04 13:29:21', '2023-03-04 13:29:21');
INSERT INTO `sys_dict_type` VALUES (82, '系统-模板-go类型', 'sys_gen_go_type', '1', '', 1, 1, '2023-03-07 10:08:07', '2023-03-07 10:08:07');
INSERT INTO `sys_dict_type` VALUES (83, '系统-模板-查询类型', 'sys_gen_query_type', '1', '', 1, 1, '2023-03-07 10:20:19', '2023-03-07 10:20:19');
INSERT INTO `sys_dict_type` VALUES (84, '系统-模板-显示类型', 'sys_gen_html_type', '1', '', 1, 1, '2023-03-07 10:23:23', '2023-03-07 10:23:23');
INSERT INTO `sys_dict_type` VALUES (85, '系统-菜单-类型', 'sys_menu_type', '1', '', 1, 1, '2023-03-08 10:33:32', '2023-03-08 10:33:32');
INSERT INTO `sys_dict_type` VALUES (86, 'App-用户-等级', 'app_user_level_type', '1', '', 1, 1, '2023-03-08 11:44:48', '2023-03-08 11:44:48');
INSERT INTO `sys_dict_type` VALUES (87, 'App-用户-资产-资金类型', 'app_money_type', '1', '1-CNY', 1, 1, '2023-03-09 20:24:17', '2023-03-11 14:06:46');
INSERT INTO `sys_dict_type` VALUES (88, 'App-用户-资产-账变类型', 'app_account_change_type', '1', '1-类型1', 1, 1, '2023-03-09 20:27:33', '2023-03-11 14:06:38');
INSERT INTO `sys_dict_type` VALUES (89, 'App-用户-行为类型', 'app_user_action_type', '1', '', 1, 1, '2023-03-11 14:06:29', '2023-03-11 14:06:29');
INSERT INTO `sys_dict_type` VALUES (90, 'App-用户-用户更新类型', 'app_user_by_type', '1', '', 1, 1, '2023-03-11 14:14:06', '2023-03-11 14:14:27');
INSERT INTO `sys_dict_type` VALUES (91, '插件-消息-验证码类型', 'plugin_msg_code_type', '1', '1-邮箱 2-短信', 1, 1, '2023-03-12 12:12:30', '2023-03-12 12:15:20');
INSERT INTO `sys_dict_type` VALUES (92, '插件-消息-验证码发送状态', 'plugin_msg_sendstatus', '1', '', 1, 1, '2023-03-12 12:14:56', '2023-03-12 13:23:37');

-- ----------------------------
-- Table structure for sys_gen_column
-- ----------------------------
DROP TABLE IF EXISTS `sys_gen_column`;
CREATE TABLE `sys_gen_column`  (
  `id` int NOT NULL AUTO_INCREMENT,
  `table_id` int NULL DEFAULT NULL,
  `column_name` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL,
  `column_comment` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL,
  `column_type` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL,
  `go_type` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL,
  `go_field` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL,
  `json_field` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL,
  `is_pk` char(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL,
  `is_required` char(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL COMMENT '该值是否参与新增或者编辑',
  `is_edit` char(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL COMMENT '该值可否二次编辑',
  `is_must` char(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL COMMENT '是否必须填写值 1-是 2-否',
  `is_list` char(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL COMMENT '列表',
  `is_query` char(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL,
  `query_type` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL,
  `html_type` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL,
  `dict_type` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL,
  `sort` bigint NULL DEFAULT NULL,
  `remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL,
  `created_at` datetime NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime NULL DEFAULT NULL COMMENT '最后更新时间',
  `create_by` int NULL DEFAULT NULL COMMENT '创建者',
  `update_by` int NULL DEFAULT NULL COMMENT '更新者',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_sys_columns_create_by`(`create_by` ASC) USING BTREE,
  INDEX `idx_sys_columns_update_by`(`update_by` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 353 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_bin ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of sys_gen_column
-- ----------------------------
INSERT INTO `sys_gen_column` VALUES (201, 8, 'id', '账变编号', 'int', 'int64', 'Id', 'id', '1', '2', '2', NULL, '1', '2', 'EQ', 'input', '', 1, '', '2023-03-09 17:59:56', '2023-03-09 21:40:08', 0, 1);
INSERT INTO `sys_gen_column` VALUES (202, 8, 'user_id', '用户编号', 'int', 'int64', 'UserId', 'userId', '2', '2', '2', NULL, '1', '1', 'EQ', 'input', '', 2, '', '2023-03-09 17:59:56', '2023-03-09 21:38:22', 0, 1);
INSERT INTO `sys_gen_column` VALUES (203, 8, 'change_money', '账变金额', 'decimal(10,2)', 'decimal.Decimal', 'ChangeMoney', 'changeMoney', '2', '2', '2', NULL, '1', '2', 'EQ', 'input', '', 3, '', '2023-03-09 17:59:56', '2023-03-09 21:38:23', 0, 1);
INSERT INTO `sys_gen_column` VALUES (204, 8, 'before_money', '账变前金额', 'decimal(30,18)', 'decimal.Decimal', 'BeforeMoney', 'beforeMoney', '2', '2', '2', NULL, '1', '2', 'EQ', 'input', '', 4, '', '2023-03-09 17:59:56', '2023-03-09 21:38:23', 0, 1);
INSERT INTO `sys_gen_column` VALUES (205, 8, 'after_money', '账变后金额', 'decimal(30,18)', 'decimal.Decimal', 'AfterMoney', 'afterMoney', '2', '2', '2', NULL, '1', '2', 'EQ', 'input', '', 5, '', '2023-03-09 17:59:56', '2023-03-09 21:38:23', 0, 1);
INSERT INTO `sys_gen_column` VALUES (206, 8, 'money_type', '金额类型 1:余额 ', 'char(10)', 'string', 'MoneyType', 'moneyType', '2', '2', '2', NULL, '1', '1', 'EQ', 'select', 'app_money_type', 6, '', '2023-03-09 17:59:56', '2023-03-09 21:38:23', 0, 1);
INSERT INTO `sys_gen_column` VALUES (207, 8, 'change_type', '帐变类型(1-类型1)', 'varchar(30)', 'string', 'ChangeType', 'changeType', '2', '2', '2', NULL, '1', '1', 'EQ', 'select', 'app_account_change_type', 7, '', '2023-03-09 17:59:56', '2023-03-09 21:38:23', 0, 1);
INSERT INTO `sys_gen_column` VALUES (208, 8, 'status', '状态（1正常 2-异常）', 'char(1)', 'string', 'Status', 'status', '2', '2', '2', NULL, NULL, '2', 'EQ', 'select', 'sys_status', 8, '', '2023-03-09 17:59:56', '2023-03-09 21:38:23', 0, 1);
INSERT INTO `sys_gen_column` VALUES (209, 8, 'create_by', '创建者', 'int', 'int64', 'CreateBy', 'createBy', '2', '2', '2', NULL, NULL, '2', 'EQ', 'input', '', 9, '', '2023-03-09 17:59:56', '2023-03-09 21:38:23', 0, 1);
INSERT INTO `sys_gen_column` VALUES (210, 8, 'created_at', '创建时间', 'datetime', '*time.Time', 'CreatedAt', 'createdAt', '2', '2', '2', NULL, '1', '2', 'EQ', 'datetime', '', 10, '', '2023-03-09 17:59:56', '2023-03-09 21:38:23', 0, 1);
INSERT INTO `sys_gen_column` VALUES (211, 8, 'update_by', '更新者', 'int', 'int64', 'UpdateBy', 'updateBy', '2', '2', '2', NULL, NULL, '2', 'EQ', 'input', '', 11, '', '2023-03-09 17:59:56', '2023-03-09 21:38:23', 0, 1);
INSERT INTO `sys_gen_column` VALUES (212, 8, 'updated_date', '更新时间', 'datetime', '*time.Time', 'UpdatedDate', 'updatedDate', '2', '2', '2', NULL, NULL, '2', 'EQ', 'datetime', '', 12, '', '2023-03-09 17:59:56', '2023-03-09 21:38:23', 0, 1);
INSERT INTO `sys_gen_column` VALUES (213, 8, 'remarks', '备注信息', 'varchar(500)', 'string', 'Remarks', 'remarks', '2', '2', '2', NULL, NULL, '2', 'EQ', 'input', '', 13, '', '2023-03-09 17:59:56', '2023-03-09 17:59:56', 0, 0);
INSERT INTO `sys_gen_column` VALUES (214, 9, 'id', '等级编号', 'int', 'int64', 'Id', 'id', '1', '2', '2', NULL, '1', '2', 'EQ', 'input', '', 1, '', '2023-03-09 20:05:43', '2023-03-09 20:17:04', 0, 1);
INSERT INTO `sys_gen_column` VALUES (215, 9, 'name', '等级名称', 'varchar(255)', 'string', 'Name', 'name', '2', '1', '1', '1', '1', '1', 'LIKE', 'input', '', 2, '', '2023-03-09 20:05:43', '2023-03-09 22:47:41', 0, 1);
INSERT INTO `sys_gen_column` VALUES (216, 9, 'level_type', '等级类型', 'varchar(10)', 'string', 'LevelType', 'levelType', '2', '1', '1', '1', '1', '1', 'EQ', 'select', 'app_user_level_type', 3, '', '2023-03-09 20:05:43', '2023-03-09 22:47:41', 0, 1);
INSERT INTO `sys_gen_column` VALUES (217, 9, 'level', '等级', 'int', 'int64', 'Level', 'level', '2', '1', '1', '1', '1', '1', 'EQ', 'numInput', '', 4, '', '2023-03-09 20:05:43', '2023-03-09 22:47:41', 0, 1);
INSERT INTO `sys_gen_column` VALUES (218, 9, 'status', '状态(1-正常 2-异常)', 'char(1)', 'string', 'Status', 'status', '2', '2', '2', NULL, '2', '2', 'EQ', 'select', 'sys_status', 5, '', '2023-03-09 20:05:43', '2023-03-09 20:17:05', 0, 1);
INSERT INTO `sys_gen_column` VALUES (219, 9, 'remark', '备注信息', 'varchar(500)', 'string', 'Remark', 'remark', '2', '2', '2', NULL, '2', '2', 'LIKE', 'input', '', 6, '', '2023-03-09 20:05:43', '2023-03-09 20:08:51', 0, 1);
INSERT INTO `sys_gen_column` VALUES (220, 9, 'create_by', '创建者', 'int', 'int64', 'CreateBy', 'createBy', '2', '2', '2', NULL, '2', '2', 'EQ', 'input', '', 7, '', '2023-03-09 20:05:43', '2023-03-09 20:17:05', 0, 1);
INSERT INTO `sys_gen_column` VALUES (221, 9, 'update_by', '更新者', 'int', 'int64', 'UpdateBy', 'updateBy', '2', '2', '2', NULL, '2', '2', 'EQ', 'input', '', 8, '', '2023-03-09 20:05:43', '2023-03-09 20:17:05', 0, 1);
INSERT INTO `sys_gen_column` VALUES (222, 9, 'created_at', '创建时间', 'datetime', '*time.Time', 'CreatedAt', 'createdAt', '2', '2', '2', NULL, '1', '2', 'EQ', 'datetime', '', 9, '', '2023-03-09 20:05:43', '2023-03-09 20:17:05', 0, 1);
INSERT INTO `sys_gen_column` VALUES (223, 9, 'updated_at', '更新时间', 'datetime', '*time.Time', 'UpdatedAt', 'updatedAt', '2', '2', '2', NULL, '2', '2', 'EQ', 'datetime', '', 10, '', '2023-03-09 20:05:43', '2023-03-09 20:17:05', 0, 1);
INSERT INTO `sys_gen_column` VALUES (224, 10, 'id', '配置编号', 'int', 'int64', 'Id', 'id', '1', '2', '2', '2', '1', '2', 'EQ', 'input', '', 1, '', '2023-03-09 22:59:52', '2023-03-09 23:02:29', 0, 1);
INSERT INTO `sys_gen_column` VALUES (225, 10, 'user_id', '用户编号', 'int', 'int64', 'UserId', 'userId', '2', '1', '2', '1', '1', '1', 'EQ', 'numInput', '', 2, '', '2023-03-09 22:59:52', '2023-03-09 23:09:54', 0, 1);
INSERT INTO `sys_gen_column` VALUES (226, 10, 'can_login', '1-允许登陆；2-不允许登陆', 'char(1)', 'string', 'CanLogin', 'canLogin', '2', '1', '1', '1', '1', '1', 'EQ', 'select', 'sys_yes_no', 3, '', '2023-03-09 22:59:52', '2023-03-09 23:02:29', 0, 1);
INSERT INTO `sys_gen_column` VALUES (227, 10, 'remark', '备注信息', 'varchar(500)', 'string', 'Remark', 'remark', '2', '2', '2', '2', '2', '2', 'EQ', 'input', '', 4, '', '2023-03-09 22:59:52', '2023-03-09 22:59:52', 0, 0);
INSERT INTO `sys_gen_column` VALUES (228, 10, 'status', '状态（1-正常 2-异常）\n', 'char(1)', 'string', 'Status', 'status', '2', '2', '2', '2', '2', '2', 'EQ', 'select', 'sys_status', 5, '', '2023-03-09 22:59:52', '2023-03-09 23:02:29', 0, 1);
INSERT INTO `sys_gen_column` VALUES (229, 10, 'create_by', '创建者', 'int', 'int64', 'CreateBy', 'createBy', '2', '2', '2', '2', '2', '2', 'EQ', 'input', '', 6, '', '2023-03-09 22:59:52', '2023-03-09 23:02:29', 0, 1);
INSERT INTO `sys_gen_column` VALUES (230, 10, 'update_by', '更新者', 'int', 'int64', 'UpdateBy', 'updateBy', '2', '2', '2', '2', '2', '2', 'EQ', 'input', '', 7, '', '2023-03-09 22:59:52', '2023-03-09 23:02:29', 0, 1);
INSERT INTO `sys_gen_column` VALUES (231, 10, 'created_at', '创建时间', 'datetime', '*time.Time', 'CreatedAt', 'createdAt', '2', '2', '2', '2', '1', '2', 'EQ', 'datetime', '', 8, '', '2023-03-09 22:59:52', '2023-03-09 23:02:29', 0, 1);
INSERT INTO `sys_gen_column` VALUES (232, 10, 'updated_at', '更新时间', 'datetime', '*time.Time', 'UpdatedAt', 'updatedAt', '2', '2', '2', '2', '2', '2', 'EQ', 'datetime', '', 9, '', '2023-03-09 22:59:52', '2023-03-09 23:02:29', 0, 1);
INSERT INTO `sys_gen_column` VALUES (233, 11, 'id', '用户编号', 'int', 'int64', 'Id', 'id', '1', '2', '2', '2', '1', '1', 'EQ', 'input', '', 1, '', '2023-03-09 23:12:17', '2023-03-09 23:18:35', 0, 1);
INSERT INTO `sys_gen_column` VALUES (234, 11, 'level_id', '用户等级编号', 'int', 'int64', 'LevelId', 'levelId', '2', '1', '1', '1', '1', '1', 'EQ', 'numInput', '', 2, '', '2023-03-09 23:12:17', '2023-03-09 23:25:14', 0, 1);
INSERT INTO `sys_gen_column` VALUES (235, 11, 'user_name', '用户昵称', 'varchar(100)', 'string', 'UserName', 'userName', '2', '1', '1', '1', '1', '1', 'EQ', 'input', '', 3, '', '2023-03-09 23:12:17', '2023-03-09 23:18:35', 0, 1);
INSERT INTO `sys_gen_column` VALUES (236, 11, 'true_name', '真实姓名', 'varchar(100)', 'string', 'TrueName', 'trueName', '2', '1', '1', '1', '1', '1', 'EQ', 'input', '', 4, '', '2023-03-09 23:12:17', '2023-03-09 23:18:35', 0, 1);
INSERT INTO `sys_gen_column` VALUES (237, 11, 'money', '余额', 'decimal(30,18)', 'decimal.Decimal', 'Money', 'money', '2', '1', '1', '1', '1', '2', 'EQ', 'input', '', 5, '', '2023-03-09 23:12:17', '2023-03-09 23:18:35', 0, 1);
INSERT INTO `sys_gen_column` VALUES (238, 11, 'email', '电子邮箱', 'varchar(300)', 'string', 'Email', 'email', '2', '1', '1', '1', '1', '1', 'EQ', 'input', '', 6, '', '2023-03-09 23:12:17', '2023-03-09 23:18:35', 0, 1);
INSERT INTO `sys_gen_column` VALUES (239, 11, 'mobile_title', '用户手机号国家前缀', 'varchar(255)', 'string', 'MobileTitle', 'mobileTitle', '2', '1', '1', '1', '1', '2', 'EQ', 'input', '', 7, '', '2023-03-09 23:12:17', '2023-03-09 23:18:35', 0, 1);
INSERT INTO `sys_gen_column` VALUES (240, 11, 'mobile', '手机号码', 'varchar(100)', 'string', 'Mobile', 'mobile', '2', '1', '1', '1', '1', '1', 'EQ', 'input', '', 8, '', '2023-03-09 23:12:17', '2023-03-09 23:18:35', 0, 1);
INSERT INTO `sys_gen_column` VALUES (241, 11, 'avatar', '头像路径', 'varchar(1000)', 'string', 'Avatar', 'avatar', '2', '2', '2', '2', '2', '2', 'EQ', 'input', '', 9, '', '2023-03-09 23:12:17', '2023-03-09 23:22:58', 0, 1);
INSERT INTO `sys_gen_column` VALUES (242, 11, 'pay_pwd', '提现密码', 'varchar(100)', 'string', 'PayPwd', 'payPwd', '2', '2', '2', '2', '2', '2', 'EQ', 'input', '', 10, '', '2023-03-09 23:12:17', '2023-03-09 23:22:58', 0, 1);
INSERT INTO `sys_gen_column` VALUES (243, 11, 'pwd', '登录密码', 'varchar(100)', 'string', 'Pwd', 'pwd', '2', '2', '2', '2', '2', '2', 'EQ', 'input', '', 11, '', '2023-03-09 23:12:17', '2023-03-09 23:22:59', 0, 1);
INSERT INTO `sys_gen_column` VALUES (244, 11, 'ref_code', '推荐码', 'varchar(255)', 'string', 'RefCode', 'refCode', '2', '2', '2', '2', '1', '1', 'EQ', 'input', '', 12, '', '2023-03-09 23:12:17', '2023-03-09 23:22:59', 0, 1);
INSERT INTO `sys_gen_column` VALUES (245, 11, 'parent_id', '父级编号', 'int', 'int64', 'ParentId', 'parentId', '2', '2', '2', '2', '1', '1', 'EQ', 'input', '', 13, '', '2023-03-09 23:12:17', '2023-03-09 23:22:59', 0, 1);
INSERT INTO `sys_gen_column` VALUES (246, 11, 'parent_ids', '所有父级编号', 'varchar(1000)', 'string', 'ParentIds', 'parentIds', '2', '2', '2', '2', '2', '2', 'EQ', 'input', '', 14, '', '2023-03-09 23:12:17', '2023-03-09 23:18:36', 0, 1);
INSERT INTO `sys_gen_column` VALUES (247, 11, 'tree_sort', '本级排序号（升序）', 'decimal(10,0)', 'decimal.Decimal', 'TreeSort', 'treeSort', '2', '2', '2', '2', '2', '2', 'EQ', 'input', '', 15, '', '2023-03-09 23:12:17', '2023-03-09 23:18:36', 0, 1);
INSERT INTO `sys_gen_column` VALUES (248, 11, 'tree_sorts', '所有级别排序号', 'varchar(1000)', 'string', 'TreeSorts', 'treeSorts', '2', '2', '2', '2', '2', '2', 'EQ', 'input', '', 16, '', '2023-03-09 23:12:17', '2023-03-09 23:18:36', 0, 1);
INSERT INTO `sys_gen_column` VALUES (249, 11, 'tree_leaf', '是否最末级', 'char(1)', 'string', 'TreeLeaf', 'treeLeaf', '2', '2', '2', '2', '2', '2', 'EQ', 'input', '', 17, '', '2023-03-09 23:12:17', '2023-03-09 23:18:36', 0, 1);
INSERT INTO `sys_gen_column` VALUES (250, 11, 'tree_level', '层次级别', 'int', 'int64', 'TreeLevel', 'treeLevel', '2', '2', '2', '2', '2', '2', 'EQ', 'input', '', 18, '', '2023-03-09 23:12:17', '2023-03-09 23:18:36', 0, 1);
INSERT INTO `sys_gen_column` VALUES (251, 11, 'status', '状态(1-正常 2-异常)', 'char(1)', 'string', 'Status', 'status', '2', '2', '2', '2', '2', '2', 'EQ', 'select', 'sys_status', 19, '', '2023-03-09 23:12:17', '2023-03-09 23:18:36', 0, 1);
INSERT INTO `sys_gen_column` VALUES (252, 11, 'remark', '备注信息', 'varchar(500)', 'string', 'Remark', 'remark', '2', '2', '2', '2', '2', '2', 'EQ', 'input', '', 20, '', '2023-03-09 23:12:17', '2023-03-09 23:12:17', 0, 0);
INSERT INTO `sys_gen_column` VALUES (253, 11, 'create_by', '创建者', 'int', 'int64', 'CreateBy', 'createBy', '2', '2', '2', '2', '0', '2', 'EQ', 'input', '', 21, '', '2023-03-09 23:12:17', '2023-03-09 23:18:36', 0, 1);
INSERT INTO `sys_gen_column` VALUES (254, 11, 'update_by', '更新者', 'int', 'int64', 'UpdateBy', 'updateBy', '2', '2', '2', '2', '0', '2', 'EQ', 'input', '', 22, '', '2023-03-09 23:12:17', '2023-03-09 23:18:36', 0, 1);
INSERT INTO `sys_gen_column` VALUES (255, 11, 'created_at', '创建时间', 'datetime', '*time.Time', 'CreatedAt', 'createdAt', '2', '2', '2', '2', '1', '2', 'EQ', 'datetime', '', 23, '', '2023-03-09 23:12:17', '2023-03-09 23:18:36', 0, 1);
INSERT INTO `sys_gen_column` VALUES (256, 11, 'updated_at', '更新时间', 'datetime', '*time.Time', 'UpdatedAt', 'updatedAt', '2', '2', '2', '2', '0', '2', 'EQ', 'datetime', '', 24, '', '2023-03-09 23:12:17', '2023-03-09 23:18:36', 0, 1);
INSERT INTO `sys_gen_column` VALUES (257, 12, 'action_type', '用户行为类型', 'char(2)', 'string', 'ActionType', 'actionType', '2', '2', '2', '2', '1', '1', 'EQ', 'select', 'app_user_action_type', 1, '', '2023-03-11 14:00:15', '2023-03-11 14:08:37', 0, 1);
INSERT INTO `sys_gen_column` VALUES (258, 12, 'by_type', '更新用户类型 1-app用户 2-后台用户', 'char(2)', 'string', 'ByType', 'byType', '2', '2', '2', '2', '1', '1', 'EQ', 'select', 'app_user_by_type', 2, '', '2023-03-11 14:00:15', '2023-03-11 14:15:30', 0, 1);
INSERT INTO `sys_gen_column` VALUES (259, 12, 'create_by', '创建者', 'int', 'int64', 'CreateBy', 'createBy', '2', '2', '2', '2', '1', '2', 'EQ', 'input', '', 3, '', '2023-03-11 14:00:15', '2023-03-11 14:05:04', 0, 1);
INSERT INTO `sys_gen_column` VALUES (260, 12, 'created_at', '创建时间', 'datetime', '*time.Time', 'CreatedAt', 'createdAt', '2', '2', '2', '2', '1', '2', 'EQ', 'datetime', '', 4, '', '2023-03-11 14:00:15', '2023-03-11 14:05:05', 0, 1);
INSERT INTO `sys_gen_column` VALUES (261, 12, 'id', '日志编码', 'int', 'int64', 'Id', 'id', '1', '2', '2', '2', '1', '2', 'EQ', 'input', '', 5, '', '2023-03-11 14:00:15', '2023-03-11 14:05:05', 0, 1);
INSERT INTO `sys_gen_column` VALUES (262, 12, 'remark', '备注信息', 'varchar(500)', 'string', 'Remark', 'remark', '2', '2', '2', '2', '2', '2', 'EQ', 'input', '', 6, '', '2023-03-11 14:00:15', '2023-03-11 14:00:15', 0, 0);
INSERT INTO `sys_gen_column` VALUES (263, 12, 'status', '状态(1-正常 2-异常)', 'char(1)', 'string', 'Status', 'status', '2', '2', '2', '2', '2', '2', 'EQ', 'select', 'sys_status', 7, '', '2023-03-11 14:00:15', '2023-03-11 14:18:50', 0, 1);
INSERT INTO `sys_gen_column` VALUES (264, 12, 'update_by', '更新者', 'int', 'int64', 'UpdateBy', 'updateBy', '2', '2', '2', '2', '1', '2', 'EQ', 'input', '', 8, '', '2023-03-11 14:00:15', '2023-03-11 14:05:05', 0, 1);
INSERT INTO `sys_gen_column` VALUES (265, 12, 'updated_at', '更新时间', 'datetime', '*time.Time', 'UpdatedAt', 'updatedAt', '2', '2', '2', '2', '1', '2', 'EQ', 'datetime', '', 9, '', '2023-03-11 14:00:15', '2023-03-11 14:05:05', 0, 1);
INSERT INTO `sys_gen_column` VALUES (266, 12, 'user_id', '用户编号', 'int', 'int64', 'UserId', 'userId', '2', '2', '2', '2', '1', '1', 'EQ', 'numInput', '', 10, '', '2023-03-11 14:00:15', '2023-03-11 14:05:05', 0, 1);
INSERT INTO `sys_gen_column` VALUES (292, 15, 'id', '验证码编号', 'int', 'int64', 'Id', 'id', '1', '2', '2', '2', '1', '1', 'EQ', 'input', '', 1, '', '2023-03-12 12:11:09', '2023-03-12 12:14:11', 0, 1);
INSERT INTO `sys_gen_column` VALUES (293, 15, 'user_id', '用户编号', 'int', 'int64', 'UserId', 'userId', '2', '2', '2', '2', '1', '1', 'EQ', 'input', '', 2, '', '2023-03-12 12:11:09', '2023-03-12 12:14:11', 0, 1);
INSERT INTO `sys_gen_column` VALUES (294, 15, 'code', '验证码', 'varchar(12)', 'string', 'Code', 'code', '2', '2', '2', '2', '1', '2', 'EQ', 'input', '', 3, '', '2023-03-12 12:11:09', '2023-03-12 12:14:11', 0, 1);
INSERT INTO `sys_gen_column` VALUES (295, 15, 'code_type', '验证码类型 1-邮箱；2-短信', 'char(1)', 'string', 'CodeType', 'codeType', '2', '2', '2', '2', '1', '1', 'EQ', 'select', 'plugin_msg_code_type', 4, '', '2023-03-12 12:11:09', '2023-03-12 12:16:18', 0, 1);
INSERT INTO `sys_gen_column` VALUES (296, 15, 'remark', '备注异常', 'varchar(500)', 'string', 'Remark', 'remark', '2', '2', '2', '2', '1', '2', 'EQ', 'input', '', 5, '', '2023-03-12 12:11:09', '2023-03-12 12:14:12', 0, 1);
INSERT INTO `sys_gen_column` VALUES (297, 15, 'status', '验证码状态 1-发送成功 2-发送失败', 'char(1)', 'string', 'Status', 'status', '2', '2', '2', '2', '1', '1', 'EQ', 'select', 'plugin_msg_sendstatus', 6, '', '2023-03-12 12:11:09', '2023-03-12 13:44:44', 0, 1);
INSERT INTO `sys_gen_column` VALUES (298, 15, 'create_by', '创建者', 'int', 'int64', 'CreateBy', 'createBy', '2', '2', '2', '2', '2', '2', 'EQ', 'input', '', 7, '', '2023-03-12 12:11:09', '2023-03-12 12:14:12', 0, 1);
INSERT INTO `sys_gen_column` VALUES (299, 15, 'update_by', '更新者', 'int', 'int64', 'UpdateBy', 'updateBy', '2', '2', '2', '2', '2', '2', 'EQ', 'input', '', 8, '', '2023-03-12 12:11:09', '2023-03-12 12:14:12', 0, 1);
INSERT INTO `sys_gen_column` VALUES (300, 15, 'created_at', '创建时间', 'datetime', '*time.Time', 'CreatedAt', 'createdAt', '2', '2', '2', '2', '2', '2', 'EQ', 'datetime', '', 9, '', '2023-03-12 12:11:09', '2023-03-12 12:14:12', 0, 1);
INSERT INTO `sys_gen_column` VALUES (301, 15, 'updated_at', '更新时间', 'datetime', '*time.Time', 'UpdatedAt', 'updatedAt', '2', '2', '2', '2', '2', '2', 'EQ', 'datetime', '', 10, '', '2023-03-12 12:11:09', '2023-03-12 12:14:12', 0, 1);
INSERT INTO `sys_gen_column` VALUES (302, 16, 'id', '公告编号', 'int', 'int64', 'Id', 'id', '1', '2', '2', '2', '1', '2', 'EQ', 'input', '', 1, '', '2023-03-12 22:01:07', '2023-03-12 22:15:05', 0, 1);
INSERT INTO `sys_gen_column` VALUES (303, 16, 'title', '标题', 'varchar(255)', 'string', 'Title', 'title', '2', '1', '1', '1', '1', '1', 'EQ', 'input', '', 2, '', '2023-03-12 22:01:07', '2023-03-12 22:15:05', 0, 1);
INSERT INTO `sys_gen_column` VALUES (304, 16, 'content', '内容', 'text', 'string', 'Content', 'content', '2', '1', '1', '1', '1', '2', 'EQ', 'input', '', 3, '', '2023-03-12 22:01:07', '2023-03-12 22:15:05', 0, 1);
INSERT INTO `sys_gen_column` VALUES (305, 16, 'num', '阅读次数', 'int', 'int64', 'Num', 'num', '2', '1', '1', '1', '1', '2', 'EQ', 'numInput', '', 4, '', '2023-03-12 22:01:07', '2023-03-12 22:51:59', 0, 1);
INSERT INTO `sys_gen_column` VALUES (306, 16, 'remark', '备注信息', 'varchar(500)', 'string', 'Remark', 'remark', '2', '1', '1', '1', '1', '2', 'EQ', 'input', '', 5, '', '2023-03-12 22:01:07', '2023-03-12 22:15:05', 0, 1);
INSERT INTO `sys_gen_column` VALUES (307, 16, 'status', '状态（0正常 1删除 2停用 3冻结）', 'char(1)', 'string', 'Status', 'status', '2', '1', '1', '1', '1', '1', 'EQ', 'select', 'sys_status', 6, '', '2023-03-12 22:01:07', '2023-03-12 22:15:05', 0, 1);
INSERT INTO `sys_gen_column` VALUES (308, 16, 'create_by', '创建者', 'int', 'int64', 'CreateBy', 'createBy', '2', '2', '2', '2', '2', '2', 'EQ', 'input', '', 7, '', '2023-03-12 22:01:07', '2023-03-12 22:15:05', 0, 1);
INSERT INTO `sys_gen_column` VALUES (309, 16, 'update_by', '更新者', 'int', 'int64', 'UpdateBy', 'updateBy', '2', '2', '2', '2', '2', '2', 'EQ', 'input', '', 8, '', '2023-03-12 22:01:07', '2023-03-12 22:15:05', 0, 1);
INSERT INTO `sys_gen_column` VALUES (310, 16, 'updated_at', '更新时间', 'datetime', '*time.Time', 'UpdatedAt', 'updatedAt', '2', '2', '2', '2', '1', '2', 'EQ', 'datetime', '', 9, '', '2023-03-12 22:01:07', '2023-03-12 22:15:06', 0, 1);
INSERT INTO `sys_gen_column` VALUES (311, 16, 'created_at', '创建时间', 'datetime', '*time.Time', 'CreatedAt', 'createdAt', '2', '2', '2', '2', '1', '2', 'EQ', 'datetime', '', 10, '', '2023-03-12 22:01:07', '2023-03-12 22:15:06', 0, 1);
INSERT INTO `sys_gen_column` VALUES (312, 17, 'id', '分类编号', 'int', 'int64', 'Id', 'id', '1', '2', '2', '2', '1', '2', 'EQ', 'input', '', 1, '', '2023-03-12 22:54:51', '2023-03-12 22:57:39', 0, 1);
INSERT INTO `sys_gen_column` VALUES (313, 17, 'name', '分类名称', 'varchar(255)', 'string', 'Name', 'name', '2', '1', '2', '1', '1', '1', 'LIKE', 'input', '', 2, '', '2023-03-12 22:54:51', '2023-03-12 22:57:39', 0, 1);
INSERT INTO `sys_gen_column` VALUES (314, 17, 'status', '状态（1-正常 2-异常）', 'char(1)', 'string', 'Status', 'status', '2', '2', '2', '2', '2', '2', 'EQ', 'select', 'sys_status', 3, '', '2023-03-12 22:54:51', '2023-03-12 22:57:39', 0, 1);
INSERT INTO `sys_gen_column` VALUES (315, 17, 'remark', '备注信息', 'varchar(500)', 'string', 'Remark', 'remark', '2', '2', '2', '2', '2', '2', 'EQ', 'input', '', 4, '', '2023-03-12 22:54:51', '2023-03-12 22:54:51', 0, 0);
INSERT INTO `sys_gen_column` VALUES (316, 17, 'create_by', '更新人编号', 'int', 'int64', 'CreateBy', 'createBy', '2', '2', '2', '2', '1', '2', 'EQ', 'input', '', 5, '', '2023-03-12 22:54:51', '2023-03-12 22:57:39', 0, 1);
INSERT INTO `sys_gen_column` VALUES (317, 17, 'update_by', '更新者', 'int', 'int64', 'UpdateBy', 'updateBy', '2', '2', '2', '2', '2', '2', 'EQ', 'input', '', 6, '', '2023-03-12 22:54:51', '2023-03-12 22:57:39', 0, 1);
INSERT INTO `sys_gen_column` VALUES (318, 17, 'updated_at', '更新时间', 'datetime', '*time.Time', 'UpdatedAt', 'updatedAt', '2', '2', '2', '2', '1', '2', 'EQ', 'datetime', '', 7, '', '2023-03-12 22:54:51', '2023-03-12 22:57:39', 0, 1);
INSERT INTO `sys_gen_column` VALUES (319, 17, 'created_at', '创建时间', 'datetime', '*time.Time', 'CreatedAt', 'createdAt', '2', '2', '2', '2', '1', '2', 'EQ', 'datetime', '', 8, '', '2023-03-12 22:54:51', '2023-03-12 22:57:39', 0, 1);
INSERT INTO `sys_gen_column` VALUES (320, 18, 'id', '文章编号', 'int', 'int64', 'Id', 'id', '1', '2', '2', '2', '1', '1', 'EQ', 'input', '', 1, '', '2023-03-12 23:22:39', '2023-03-12 23:27:47', 0, 1);
INSERT INTO `sys_gen_column` VALUES (321, 18, 'cate_id', '分类编号', 'int', 'int64', 'CateId', 'cateId', '2', '1', '1', '1', '1', '1', 'EQ', 'numInput', '', 2, '', '2023-03-12 23:22:39', '2023-03-12 23:27:47', 0, 1);
INSERT INTO `sys_gen_column` VALUES (322, 18, 'name', '标题', 'varchar(255)', 'string', 'Name', 'name', '2', '1', '1', '1', '1', '1', 'EQ', 'input', '', 3, '', '2023-03-12 23:22:39', '2023-03-12 23:27:47', 0, 1);
INSERT INTO `sys_gen_column` VALUES (323, 18, 'content', '内容', 'text', 'string', 'Content', 'content', '2', '1', '1', '1', '1', '2', 'EQ', 'input', '', 4, '', '2023-03-12 23:22:39', '2023-03-12 23:27:48', 0, 1);
INSERT INTO `sys_gen_column` VALUES (324, 18, 'remark', '备注信息', 'varchar(500)', 'string', 'Remark', 'remark', '2', '1', '1', '1', '1', '2', 'EQ', 'input', '', 5, '', '2023-03-12 23:22:39', '2023-03-12 23:27:48', 0, 1);
INSERT INTO `sys_gen_column` VALUES (325, 18, 'status', '状态（1-正常 2-异常）', 'char(1)', 'string', 'Status', 'status', '2', '1', '1', '1', '1', '2', 'EQ', 'select', 'sys_status', 6, '', '2023-03-12 23:22:39', '2023-03-12 23:27:48', 0, 1);
INSERT INTO `sys_gen_column` VALUES (326, 18, 'create_by', '更新人编号', 'int', 'int64', 'CreateBy', 'createBy', '2', '2', '2', '2', '1', '2', 'EQ', 'input', '', 7, '', '2023-03-12 23:22:39', '2023-03-12 23:27:48', 0, 1);
INSERT INTO `sys_gen_column` VALUES (327, 18, 'update_by', '更新者', 'int', 'int64', 'UpdateBy', 'updateBy', '2', '2', '2', '2', '2', '2', 'EQ', 'input', '', 8, '', '2023-03-12 23:22:39', '2023-03-12 23:27:48', 0, 1);
INSERT INTO `sys_gen_column` VALUES (328, 18, 'updated_at', '更新时间', 'datetime', '*time.Time', 'UpdatedAt', 'updatedAt', '2', '2', '2', '2', '1', '2', 'EQ', 'datetime', '', 9, '', '2023-03-12 23:22:39', '2023-03-12 23:27:48', 0, 1);
INSERT INTO `sys_gen_column` VALUES (329, 18, 'created_at', '创建时间', 'datetime', '*time.Time', 'CreatedAt', 'createdAt', '2', '2', '2', '2', '1', '2', 'EQ', 'datetime', '', 10, '', '2023-03-12 23:22:39', '2023-03-12 23:27:48', 0, 1);
INSERT INTO `sys_gen_column` VALUES (330, 19, 'id', 'App编号', 'int', 'int64', 'Id', 'id', '1', '2', '2', '2', '1', '1', 'EQ', 'input', '', 1, '', '2023-03-13 00:07:25', '2023-03-13 00:12:43', 0, 1);
INSERT INTO `sys_gen_column` VALUES (331, 19, 'version', '版本号', 'varchar(100)', 'string', 'Version', 'version', '2', '1', '1', '1', '1', '1', 'EQ', 'input', '', 2, '', '2023-03-13 00:07:25', '2023-03-13 00:12:43', 0, 1);
INSERT INTO `sys_gen_column` VALUES (332, 19, 'platform', '平台 (1-安卓 2-苹果)', 'char(1)', 'string', 'Platform', 'platform', '2', '1', '1', '1', '1', '1', 'EQ', 'select', 'plugin_filemgr_app_platform', 3, '', '2023-03-13 00:07:25', '2023-03-13 00:12:43', 0, 1);
INSERT INTO `sys_gen_column` VALUES (333, 19, 'app_type', '版本(1-默认)', 'char(1)', 'string', 'AppType', 'appType', '2', '1', '1', '1', '1', '1', 'EQ', 'select', 'plugin_filemgr_app_type', 4, '', '2023-03-13 00:07:25', '2023-03-13 00:12:43', 0, 1);
INSERT INTO `sys_gen_column` VALUES (334, 19, 'local_address', '本地地址', 'varchar(255)', 'string', 'LocalAddress', 'localAddress', '2', '1', '1', '1', '1', '2', 'EQ', 'input', '', 5, '', '2023-03-13 00:07:25', '2023-03-13 00:12:43', 0, 1);
INSERT INTO `sys_gen_column` VALUES (335, 19, 'download_num', '下载数量', 'int', 'int64', 'DownloadNum', 'downloadNum', '2', '1', '1', '1', '1', '2', 'EQ', 'input', '', 6, '', '2023-03-13 00:07:25', '2023-03-13 00:12:43', 0, 1);
INSERT INTO `sys_gen_column` VALUES (336, 19, 'download_type', '下载类型(1-本地 2-外链 3-oss )', 'char(1)', 'string', 'DownloadType', 'downloadType', '2', '1', '1', '1', '1', '1', 'EQ', 'select', 'plugin_filemgr_app_download_type', 7, '', '2023-03-13 00:07:25', '2023-03-13 00:12:44', 0, 1);
INSERT INTO `sys_gen_column` VALUES (337, 19, 'download_url', '下载地址(download_type=1使用)', 'varchar(255)', 'string', 'DownloadUrl', 'downloadUrl', '2', '1', '1', '1', '1', '2', 'EQ', 'input', '', 8, '', '2023-03-13 00:07:25', '2023-03-13 00:12:44', 0, 1);
INSERT INTO `sys_gen_column` VALUES (338, 19, 'remark', '备注信息', 'varchar(500)', 'string', 'Remark', 'remark', '2', '1', '1', '1', '1', '2', 'EQ', 'input', '', 9, '', '2023-03-13 00:07:25', '2023-03-13 00:12:44', 0, 1);
INSERT INTO `sys_gen_column` VALUES (339, 19, 'status', '状态（1-已发布 2-待发布）\n', 'char(1)', 'string', 'Status', 'status', '2', '1', '1', '1', '1', '1', 'EQ', 'select', 'plugin_filemgr_publish_status', 10, '', '2023-03-13 00:07:25', '2023-03-13 00:12:44', 0, 1);
INSERT INTO `sys_gen_column` VALUES (340, 19, 'create_by', '创建者', 'int', 'int64', 'CreateBy', 'createBy', '2', '2', '2', '2', '2', '2', 'EQ', 'input', '', 11, '', '2023-03-13 00:07:25', '2023-03-13 00:12:44', 0, 1);
INSERT INTO `sys_gen_column` VALUES (341, 19, 'created_at', '创建时间', 'datetime', '*time.Time', 'CreatedAt', 'createdAt', '2', '2', '2', '2', '1', '2', 'EQ', 'datetime', '', 12, '', '2023-03-13 00:07:25', '2023-03-13 00:12:44', 0, 1);
INSERT INTO `sys_gen_column` VALUES (342, 19, 'update_by', '更新者', 'int', 'int64', 'UpdateBy', 'updateBy', '2', '2', '2', '2', '2', '2', 'EQ', 'input', '', 13, '', '2023-03-13 00:07:25', '2023-03-13 00:12:44', 0, 1);
INSERT INTO `sys_gen_column` VALUES (343, 19, 'updated_at', '更新时间', 'datetime', '*time.Time', 'UpdatedAt', 'updatedAt', '2', '2', '2', '2', '2', '2', 'EQ', 'datetime', '', 14, '', '2023-03-13 00:07:25', '2023-03-13 00:12:44', 0, 1);
INSERT INTO `sys_gen_column` VALUES (344, 20, 'id', '编号', 'int', 'int64', 'Id', 'id', '1', '2', '2', '2', '1', '2', 'EQ', 'input', '', 1, '', '2023-03-14 17:40:50', '2023-03-14 17:42:59', 0, 1);
INSERT INTO `sys_gen_column` VALUES (345, 20, 'country', '国家地区', 'varchar(64)', 'string', 'Country', 'country', '2', '1', '1', '1', '1', '1', 'LIKE', 'input', '', 2, '', '2023-03-14 17:40:50', '2023-03-14 17:43:00', 0, 1);
INSERT INTO `sys_gen_column` VALUES (346, 20, 'code', '区号', 'varchar(12)', 'string', 'Code', 'code', '2', '1', '1', '1', '1', '1', 'EQ', 'input', '', 3, '', '2023-03-14 17:40:50', '2023-03-14 17:43:00', 0, 1);
INSERT INTO `sys_gen_column` VALUES (347, 20, 'status', '状态(1-可用 2-停用)', 'char(1)', 'string', 'Status', 'status', '2', '1', '1', '1', '1', '1', 'EQ', 'select', 'sys_status', 4, '', '2023-03-14 17:40:50', '2023-03-14 17:43:00', 0, 1);
INSERT INTO `sys_gen_column` VALUES (348, 20, 'remark', '备注信息', 'varchar(500)', 'string', 'Remark', 'remark', '2', '2', '2', '2', '2', '2', 'EQ', 'input', '', 5, '', '2023-03-14 17:40:50', '2023-03-14 17:40:50', 0, 0);
INSERT INTO `sys_gen_column` VALUES (349, 20, 'create_by', '创建者', 'int', 'int64', 'CreateBy', 'createBy', '2', '2', '2', '2', '2', '2', 'EQ', 'input', '', 6, '', '2023-03-14 17:40:50', '2023-03-14 17:43:00', 0, 1);
INSERT INTO `sys_gen_column` VALUES (350, 20, 'update_by', '更新者', 'int', 'int64', 'UpdateBy', 'updateBy', '2', '2', '2', '2', '2', '2', 'EQ', 'input', '', 7, '', '2023-03-14 17:40:50', '2023-03-14 17:43:00', 0, 1);
INSERT INTO `sys_gen_column` VALUES (351, 20, 'created_at', '创建时间', 'datetime', '*time.Time', 'CreatedAt', 'createdAt', '2', '2', '2', '2', '1', '2', 'EQ', 'datetime', '', 8, '', '2023-03-14 17:40:50', '2023-03-14 17:43:00', 0, 1);
INSERT INTO `sys_gen_column` VALUES (352, 20, 'updated_at', '更新时间', 'datetime', '*time.Time', 'UpdatedAt', 'updatedAt', '2', '2', '2', '2', '2', '2', 'EQ', 'datetime', '', 9, '', '2023-03-14 17:40:50', '2023-03-14 17:43:00', 0, 1);

-- ----------------------------
-- Table structure for sys_gen_table
-- ----------------------------
DROP TABLE IF EXISTS `sys_gen_table`;
CREATE TABLE `sys_gen_table`  (
  `id` int NOT NULL AUTO_INCREMENT,
  `table_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL COMMENT '表名',
  `table_comment` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL COMMENT '表描述',
  `class_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL COMMENT '类名',
  `package_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL COMMENT '应用名',
  `module_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL COMMENT '接口名',
  `function_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL COMMENT '功能描述',
  `function_author` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL COMMENT '作者',
  `business_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL COMMENT '业务名',
  `is_plugin` char(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT '1' COMMENT '是否插件 1-是 2-否',
  `remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL,
  `created_at` datetime NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime NULL DEFAULT NULL COMMENT '最后更新时间',
  `create_by` bigint NULL DEFAULT NULL COMMENT '创建者',
  `update_by` bigint NULL DEFAULT NULL COMMENT '更新者',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_sys_tables_create_by`(`create_by` ASC) USING BTREE,
  INDEX `idx_sys_tables_update_by`(`update_by` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 21 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_bin ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of sys_gen_table
-- ----------------------------
INSERT INTO `sys_gen_table` VALUES (8, 'app_user_account_log', '账变记录', 'UserAccountLog', 'app', 'user-account-log', '账变记录', 'Jason', 'user', '2', '', '2023-03-09 17:59:56', '2023-03-09 17:59:56', 0, 0);
INSERT INTO `sys_gen_table` VALUES (9, 'app_user_level', '用户等级', 'UserLevel', 'app', 'user-level', '用户等级', 'Jason', 'user', '2', '', '2023-03-09 20:05:43', '2023-03-09 20:05:43', 0, 0);
INSERT INTO `sys_gen_table` VALUES (10, 'app_user_conf', '用户配置', 'UserConf', 'app', 'user-conf', '用户配置', 'Jason', 'user', '2', '', '2023-03-09 22:59:52', '2023-03-09 22:59:52', 0, 0);
INSERT INTO `sys_gen_table` VALUES (11, 'app_user', '用户管理', 'User', 'app', 'user', '用户管理', 'Jason', 'user', '2', '', '2023-03-09 23:12:17', '2023-03-09 23:12:17', 0, 0);
INSERT INTO `sys_gen_table` VALUES (12, 'app_user_oper_log', '用户关键行为日志表', 'UserOperLog', 'app', 'user-oper-log', '用户关键行为日志', 'Jason', 'user', '2', '', '2023-03-11 14:00:15', '2023-03-11 14:05:04', 0, 1);
INSERT INTO `sys_gen_table` VALUES (15, 'plugins_msg_code', '验证码记录', 'MsgCode', 'plugins', 'msg-code', '验证码记录', 'Jason', 'msg', '1', '', '2023-03-12 12:11:08', '2023-03-12 14:26:24', 0, 1);
INSERT INTO `sys_gen_table` VALUES (16, 'plugins_content_announcement', '公告管理', 'ContentAnnouncement', 'plugins', 'content-announcement', '公告管理', 'Jason', 'content', '2', '', '2023-03-12 22:01:07', '2023-03-12 22:01:07', 0, 0);
INSERT INTO `sys_gen_table` VALUES (17, 'plugins_content_category', '内容分类', 'ContentCategory', 'plugins', 'content-category', '文章分类管理', 'Jason', 'content', '1', '', '2023-03-12 22:54:51', '2023-03-12 22:58:31', 0, 1);
INSERT INTO `sys_gen_table` VALUES (18, 'plugins_content_article', '文章管理', 'ContentArticle', 'plugins', 'content-article', '文章管理', 'Jason', 'content', '1', '', '2023-03-12 23:22:39', '2023-03-12 23:22:39', 0, 0);
INSERT INTO `sys_gen_table` VALUES (19, 'plugins_filemgr_app', 'App管理', 'FilemgrApp', 'plugins', 'filemgr-app', 'App管理', 'Jason', 'filemgr', '1', '', '2023-03-13 00:07:25', '2023-03-13 00:12:43', 0, 1);
INSERT INTO `sys_gen_table` VALUES (20, 'app_user_country_code', '国家电话区号', 'UserCountryCode', 'app', 'user-country-code', '国家区号', 'Jason', 'user', '1', '', '2023-03-14 17:40:50', '2023-03-14 17:43:22', 0, 1);

-- ----------------------------
-- Table structure for sys_login_log
-- ----------------------------
DROP TABLE IF EXISTS `sys_login_log`;
CREATE TABLE `sys_login_log`  (
  `id` int NOT NULL AUTO_INCREMENT COMMENT '主键编码',
  `user_id` int NULL DEFAULT NULL COMMENT '用户编号',
  `ipaddr` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL COMMENT 'ip地址',
  `login_location` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL COMMENT '归属地',
  `browser` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL COMMENT '浏览器',
  `os` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL COMMENT '系统',
  `agent` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL COMMENT '代理',
  `platform` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL COMMENT '固件',
  `login_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '登录时间',
  `status` char(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL COMMENT '状态 1-登录 2-退出',
  `remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL COMMENT '备注',
  `created_at` datetime NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime NULL DEFAULT NULL COMMENT '最后更新时间',
  `create_by` int NULL DEFAULT NULL COMMENT '创建者',
  `update_by` int NULL DEFAULT NULL COMMENT '更新者',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_sys_login_log_create_by`(`create_by` ASC) USING BTREE,
  INDEX `idx_sys_login_log_update_by`(`update_by` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 26 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_bin ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of sys_login_log
-- ----------------------------

-- ----------------------------
-- Table structure for sys_menu
-- ----------------------------
DROP TABLE IF EXISTS `sys_menu`;
CREATE TABLE `sys_menu`  (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL,
  `title` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL,
  `icon` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL,
  `path` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL,
  `menu_type` char(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL,
  `permission` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL,
  `parent_id` int NULL DEFAULT NULL,
  `keep_alive` char(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL COMMENT '是否缓存',
  `breadcrumb` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL,
  `component` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL,
  `sort` int NULL DEFAULT NULL,
  `hidden` char(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL COMMENT '是否隐藏 1-是 2-否',
  `is_frame` char(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT '0' COMMENT '外链 1-是 2-否',
  `create_by` int NULL DEFAULT NULL COMMENT '创建者',
  `update_by` int NULL DEFAULT NULL COMMENT '更新者',
  `created_at` datetime NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime NULL DEFAULT NULL COMMENT '最后更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_sys_menu_create_by`(`create_by` ASC) USING BTREE,
  INDEX `idx_sys_menu_update_by`(`update_by` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 997 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_bin ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of sys_menu
-- ----------------------------
INSERT INTO `sys_menu` VALUES (2, '', '系统管理', 'api-server', '/sys', '1', '', 0, '2', '', 'Layout', 300, '2', '2', 1, 1, '2021-05-20 21:58:46', '2023-03-07 11:12:00');
INSERT INTO `sys_menu` VALUES (3, 'SysUser', '用户管理', 'user', '/sys/sys-user', '2', '', 2, '2', '', '/sys/user/index.vue', 10, '2', '2', 1, 1, '2021-05-20 22:08:45', '2023-03-04 13:17:59');
INSERT INTO `sys_menu` VALUES (43, '', '新增管理员', 'app-group-fill', '', '3', 'admin:sysUser:add', 3, '2', '', '', 10, '2', '2', 1, 1, '2021-05-20 22:08:45', '2023-03-04 13:17:59');
INSERT INTO `sys_menu` VALUES (44, '', '查询管理员', 'app-group-fill', '', '3', 'admin:sysUser:query', 3, '2', '', '', 40, '2', '2', 1, 1, '2021-05-20 22:08:45', '2023-03-04 13:17:59');
INSERT INTO `sys_menu` VALUES (45, '', '修改管理员', 'app-group-fill', '', '3', 'admin:sysUser:edit', 3, '2', '', '', 30, '2', '2', 1, 1, '2021-05-20 22:08:45', '2023-03-04 13:17:59');
INSERT INTO `sys_menu` VALUES (46, '', '删除管理员', 'app-group-fill', '', '3', 'admin:sysUser:remove', 3, '2', '', '', 20, '2', '2', 1, 1, '2021-05-20 22:08:45', '2023-03-04 13:17:59');
INSERT INTO `sys_menu` VALUES (51, 'SysMenu', '菜单管理', 'tree-table', '/sys/sys-menu', '2', 'admin:sysMenu:list', 2, '2', '', '/sys/menu/index.vue', 30, '2', '2', 1, 1, '2021-05-20 22:08:45', '2021-11-29 15:39:04');
INSERT INTO `sys_menu` VALUES (52, 'SysRole', '角色管理', 'peoples', '/sys/sys-role', '2', 'admin:sysRole:list', 2, '2', '', '/sys/role/index.vue', 20, '2', '2', 1, 1, '2021-05-20 22:08:45', '2021-11-29 15:39:04');
INSERT INTO `sys_menu` VALUES (56, 'SysDept', '部门管理', 'tree', '/sys/sys-dept', '2', '', 2, '2', '', '/sys/dept/index.vue', 40, '2', '2', 1, 1, '2021-05-20 22:08:45', '2021-11-29 15:39:04');
INSERT INTO `sys_menu` VALUES (57, 'SysPost', '岗位管理', 'pass', '/sys/sys-post', '2', 'admin:sysPost:list', 2, '2', '', '/sys/post/index.vue', 50, '2', '2', 1, 1, '2021-05-20 22:08:45', '2021-11-29 15:39:04');
INSERT INTO `sys_menu` VALUES (58, 'SysDicttype', '字典管理', 'education', '/sys/sys-dicttype', '2', 'admin:sysDictType:list', 2, '2', '', '/sys/dicttype/index.vue', 60, '2', '2', 1, 1, '2021-05-20 22:08:45', '2021-11-29 15:39:04');
INSERT INTO `sys_menu` VALUES (59, 'SysDictdata', '字典数据', 'education', '/sys/sys-dictdata', '2', 'admin:sysDictData:list', 2, '2', '', '/sys/dictdata/index.vue', 100, '1', '2', 1, 1, '2021-05-20 22:08:45', '2021-11-29 15:39:04');
INSERT INTO `sys_menu` VALUES (62, 'SysConfig', '参数管理', 'swagger', '/sys/sys-config', '2', 'admin:sysConfig:list', 2, '2', '', '/sys/config/index.vue', 70, '2', '2', 1, 1, '2021-05-20 22:08:45', '2021-11-29 15:39:04');
INSERT INTO `sys_menu` VALUES (212, 'SysLoginlog', '登录日志', 'logininfor', '/sys/sys-loginlog', '2', 'admin:sysLoginLog:list', 2, '2', '', '/sys/loginlog/index.vue', 90, '2', '2', 1, 1, '2021-05-20 22:08:45', '2021-11-29 15:39:04');
INSERT INTO `sys_menu` VALUES (216, 'SysOperalog', '操作日志', 'skill', '/sys/sys-operalog', '2', 'admin:sysOperLog:list', 2, '2', '', '/sys/operlog/index.vue', 120, '2', '2', 1, 1, '2021-05-20 22:08:45', '2021-11-29 15:39:04');
INSERT INTO `sys_menu` VALUES (220, '', '新增菜单', 'app-group-fill', '', '3', 'admin:sysMenu:add', 51, '2', '', '', 1, '2', '2', 1, 1, '2020-04-11 15:52:48', '2023-03-04 13:17:59');
INSERT INTO `sys_menu` VALUES (221, '', '修改菜单', 'app-group-fill', '', '3', 'admin:sysMenu:edit', 51, '2', '', '', 1, '2', '2', 1, 1, '2020-04-11 15:52:48', '2021-11-29 15:39:04');
INSERT INTO `sys_menu` VALUES (222, '', '查询菜单', 'app-group-fill', '', '3', 'admin:sysMenu:query', 51, '2', '', '', 1, '2', '2', 1, 1, '2020-04-11 15:52:48', '2021-11-29 15:39:04');
INSERT INTO `sys_menu` VALUES (223, '', '删除菜单', 'app-group-fill', '', '3', 'admin:sysMenu:remove', 51, '2', '', '', 1, '2', '2', 1, 1, '2020-04-11 15:52:48', '2021-11-29 15:39:04');
INSERT INTO `sys_menu` VALUES (224, '', '新增角色', 'app-group-fill', '', '3', 'admin:sysRole:add', 52, '2', '', '', 1, '2', '2', 1, 1, '2020-04-11 15:52:48', '2021-11-29 15:39:04');
INSERT INTO `sys_menu` VALUES (225, '', '查询角色', 'app-group-fill', '', '3', 'admin:sysRole:query', 52, '2', '', '', 1, '2', '2', 1, 1, '2020-04-11 15:52:48', '2023-03-03 13:41:42');
INSERT INTO `sys_menu` VALUES (226, '', '修改角色', 'app-group-fill', '', '3', 'admin:sysRole:update', 52, '2', '', '', 1, '2', '2', 1, 1, '2020-04-11 15:52:48', '2021-11-29 15:39:04');
INSERT INTO `sys_menu` VALUES (227, '', '删除角色', 'app-group-fill', '', '3', 'admin:sysRole:remove', 52, '2', '', '', 1, '2', '2', 1, 1, '2020-04-11 15:52:48', '2021-11-29 15:39:04');
INSERT INTO `sys_menu` VALUES (228, '', '查询部门', 'app-group-fill', '', '3', 'admin:sysDept:query', 56, '2', '', '', 40, '2', '2', 1, 1, '2021-05-20 22:08:45', '2021-11-29 15:39:04');
INSERT INTO `sys_menu` VALUES (229, '', '新增部门', 'app-group-fill', '', '3', 'admin:sysDept:add', 56, '2', '', '', 10, '2', '2', 1, 1, '2021-05-20 22:08:45', '2021-11-29 15:39:04');
INSERT INTO `sys_menu` VALUES (230, '', '修改部门', 'app-group-fill', '', '3', 'admin:sysDept:edit', 56, '2', '', '', 30, '2', '2', 1, 1, '2021-05-20 22:08:45', '2021-11-29 15:39:04');
INSERT INTO `sys_menu` VALUES (231, '', '删除部门', 'app-group-fill', '', '3', 'admin:sysDept:remove', 56, '2', '', '', 20, '2', '2', 1, 1, '2021-05-20 22:08:45', '2021-11-29 15:39:04');
INSERT INTO `sys_menu` VALUES (232, '', '查询岗位', 'app-group-fill', '', '3', 'admin:sysPost:query', 57, '2', '', '', 0, '2', '2', 1, 1, '2020-04-11 15:52:48', '2021-11-29 15:39:04');
INSERT INTO `sys_menu` VALUES (233, '', '新增岗位', 'app-group-fill', '', '3', 'admin:sysPost:add', 57, '2', '', '', 0, '2', '2', 1, 1, '2020-04-11 15:52:48', '2021-11-29 15:39:04');
INSERT INTO `sys_menu` VALUES (234, '', '修改岗位', 'app-group-fill', '', '3', 'admin:sysPost:edit', 57, '2', '', '', 0, '2', '2', 1, 1, '2020-04-11 15:52:48', '2021-11-29 15:39:04');
INSERT INTO `sys_menu` VALUES (235, '', '删除岗位', 'app-group-fill', '', '3', 'admin:sysPost:remove', 57, '2', '', '', 0, '2', '2', 1, 1, '2020-04-11 15:52:48', '2021-11-29 15:39:04');
INSERT INTO `sys_menu` VALUES (236, '', '查询字典', 'app-group-fill', '', '3', 'admin:sysDictType:query', 58, '2', '', '', 0, '2', '2', 1, 1, '2020-04-11 15:52:48', '2021-11-29 15:39:04');
INSERT INTO `sys_menu` VALUES (237, '', '新增类型', 'app-group-fill', '', '3', 'admin:sysDictType:add', 58, '2', '', '', 0, '2', '2', 1, 1, '2020-04-11 15:52:48', '2021-11-29 15:39:04');
INSERT INTO `sys_menu` VALUES (238, '', '修改类型', 'app-group-fill', '', '3', 'admin:sysDictType:edit', 58, '2', '', '', 0, '2', '2', 1, 1, '2020-04-11 15:52:48', '2021-11-29 15:39:04');
INSERT INTO `sys_menu` VALUES (239, '', '删除类型', 'app-group-fill', '', '3', 'admin:sysdicttype:remove', 58, '2', '', '', 0, '2', '2', 1, 1, '2020-04-11 15:52:48', '2021-11-29 15:39:04');
INSERT INTO `sys_menu` VALUES (240, '', '查询数据', 'app-group-fill', '', '3', 'admin:sysDictData:query', 59, '2', '', '', 0, '2', '2', 1, 1, '2020-04-11 15:52:48', '2021-11-29 15:39:04');
INSERT INTO `sys_menu` VALUES (241, '', '新增数据', 'app-group-fill', '', '3', 'admin:sysDictData:add', 59, '2', '', '', 0, '2', '2', 1, 1, '2020-04-11 15:52:48', '2021-11-29 15:39:04');
INSERT INTO `sys_menu` VALUES (242, '', '修改数据', 'app-group-fill', '', '3', 'admin:sysDictData:edit', 59, '2', '', '', 0, '2', '2', 1, 1, '2020-04-11 15:52:48', '2021-11-29 15:39:04');
INSERT INTO `sys_menu` VALUES (243, '', '删除数据', 'app-group-fill', '', '3', 'admin:sysDictData:remove', 59, '2', '', '', 0, '2', '2', 1, 1, '2020-04-11 15:52:48', '2021-11-29 15:39:04');
INSERT INTO `sys_menu` VALUES (244, '', '查询参数', 'app-group-fill', '', '3', 'admin:sysConfig:query', 62, '2', '', '', 0, '2', '2', 1, 1, '2020-04-11 15:52:48', '2021-11-29 15:39:04');
INSERT INTO `sys_menu` VALUES (245, '', '新增参数', 'app-group-fill', '', '3', 'admin:sysConfig:add', 62, '2', '', '', 0, '2', '2', 1, 1, '2020-04-11 15:52:48', '2021-11-29 15:39:04');
INSERT INTO `sys_menu` VALUES (246, '', '修改参数', 'app-group-fill', '', '3', 'admin:sysConfig:edit', 62, '2', '', '', 0, '2', '2', 1, 1, '2020-04-11 15:52:48', '2021-11-29 15:39:04');
INSERT INTO `sys_menu` VALUES (247, '', '删除参数', 'app-group-fill', '', '3', 'admin:sysConfig:remove', 62, '2', '', '', 0, '2', '2', 1, 1, '2020-04-11 15:52:48', '2021-11-29 15:39:04');
INSERT INTO `sys_menu` VALUES (248, '', '查询登录日志', 'app-group-fill', '', '3', 'admin:sysLoginLog:query', 212, '2', '', '', 0, '2', '2', 1, 1, '2020-04-11 15:52:48', '2021-11-29 15:39:04');
INSERT INTO `sys_menu` VALUES (249, '', '删除登录日志', 'app-group-fill', '', '3', 'admin:sysLoginLog:remove', 212, '2', '', '', 0, '2', '2', 1, 1, '2020-04-11 15:52:48', '2021-11-29 15:39:04');
INSERT INTO `sys_menu` VALUES (250, '', '查询操作日志', 'app-group-fill', '', '3', 'admin:sysOperLog:query', 216, '2', '', '', 0, '2', '2', 1, 1, '2020-04-11 15:52:48', '2021-11-29 15:39:04');
INSERT INTO `sys_menu` VALUES (251, '', '删除操作日志', 'app-group-fill', '', '3', 'admin:sysOperLog:remove', 216, '2', '', '', 0, '2', '2', 1, 1, '2020-04-11 15:52:48', '2021-11-29 15:39:04');
INSERT INTO `sys_menu` VALUES (261, 'SysGen', '代码生成', 'code', '/sys-tools/sys-gen', '2', '', 537, '2', '', '/sys/tools/gen/index.vue', 20, '2', '2', 1, 1, '2020-04-11 15:52:48', '2021-09-01 09:37:52');
INSERT INTO `sys_menu` VALUES (262, 'SysEditTable', '代码生成修改', 'build', '/sys-tools/sys-editTable', '2', '', 537, '2', '', '/sys/tools/gen/editTable.vue', 100, '1', '2', 1, 1, '2020-04-11 15:52:48', '2021-09-01 09:38:05');
INSERT INTO `sys_menu` VALUES (269, 'SysMonitor', '服务监控', 'druid', '/sys-tools/sys-monitor', '2', 'admin:monitor:list', 537, '2', '', '/sys/tools/monitor/monitor.vue', 0, '2', '2', 1, 1, '2020-04-14 00:28:19', '2021-06-16 21:26:12');
INSERT INTO `sys_menu` VALUES (528, 'SysApi', '接口管理', 'api-doc', '/sys/sys-api', '2', 'admin:sysApi:list', 2, '2', '', '/sys/api/index.vue', 0, '2', '2', 1, 1, '2021-05-20 22:08:45', '2023-03-08 10:02:50');
INSERT INTO `sys_menu` VALUES (529, '', '查询接口', 'app-group-fill', '', '3', 'admin:sysApi:query', 528, '2', '', '', 40, '2', '2', 1, 1, '2021-05-20 22:08:45', '2021-11-29 15:39:04');
INSERT INTO `sys_menu` VALUES (531, '', '修改接口', 'app-group-fill', '', '3', 'admin:sysApi:edit', 528, '2', '', '', 30, '2', '2', 1, 1, '2021-05-20 22:08:45', '2021-11-29 15:39:04');
INSERT INTO `sys_menu` VALUES (537, '', '系统工具', 'system-tools', '/sys-tools', '1', '', 0, '2', '', 'Layout', 330, '2', '2', 1, 1, '2021-05-21 11:13:32', '2021-07-22 16:04:17');
INSERT INTO `sys_menu` VALUES (772, '', '文件管理', 'base-info', '/plugins/filemgr', '1', '', 843, '2', '', '/index', 90, '2', '2', 1, 1, '2021-08-13 14:19:11', '2023-03-11 23:01:14');
INSERT INTO `sys_menu` VALUES (778, '', '内容管理', 'clipboard', '/plugins/content', '1', '', 843, '2', '', '/index', 60, '2', '2', 1, 1, '2021-08-16 18:01:20', '2023-03-11 23:01:07');
INSERT INTO `sys_menu` VALUES (843, '', '插件管理', 'cascader', '/plugins', '1', '', 0, '', '', 'Layout', 270, '2', '2', 1, 1, '2023-03-07 10:37:37', '2023-03-08 09:27:48');
INSERT INTO `sys_menu` VALUES (844, '', 'App应用', 'app-group-fill', '/app', '1', '', 0, '', '', 'Layout', 0, '2', '2', 1, 1, '2023-03-08 09:27:36', '2023-03-08 09:27:36');
INSERT INTO `sys_menu` VALUES (875, '', '用户列表', 'chart', '/app/user', '1', '', 844, '', '', '/index', 30, '2', '2', 1, 1, '2023-03-09 14:24:25', '2023-03-09 23:20:06');
INSERT INTO `sys_menu` VALUES (886, '', '财务管理', 'eye-open', '/app/account', '1', '', 844, '', '', '/index', 60, '2', '2', 1, 1, '2023-03-09 21:13:23', '2023-03-09 23:20:15');
INSERT INTO `sys_menu` VALUES (887, 'UserLevel', '用户等级', 'pass', '/app/user/user-level', '2', 'app:user:user-level:list', 875, '2', '', '/app/user/user-level/index.vue', 60, '2', '2', 1, 1, '2023-03-09 21:33:49', '2023-03-09 23:05:34');
INSERT INTO `sys_menu` VALUES (888, '', '分页获取用户等级', '', '', '3', 'app:user:user-level:query', 887, '2', '', '', 0, '2', '2', 1, 1, '2023-03-09 21:33:49', '2023-03-09 21:33:49');
INSERT INTO `sys_menu` VALUES (889, '', '创建用户等级', '', '', '3', 'app:user:user-level:add', 887, '2', '', '', 0, '2', '2', 1, 1, '2023-03-09 21:33:49', '2023-03-09 21:33:49');
INSERT INTO `sys_menu` VALUES (890, '', '修改用户等级', '', '', '3', 'app:user:user-level:edit', 887, '2', '', '', 0, '2', '2', 1, 1, '2023-03-09 21:33:49', '2023-03-09 21:33:49');
INSERT INTO `sys_menu` VALUES (891, '', '删除用户等级', '', '', '3', 'app:user:user-level:del', 887, '2', '', '', 0, '2', '2', 1, 1, '2023-03-09 21:33:49', '2023-03-09 21:33:49');
INSERT INTO `sys_menu` VALUES (892, '', '导出用户等级', '', '', '3', 'app:user:user-level:export', 887, '2', '', '', 0, '2', '2', 1, 1, '2023-03-09 21:33:49', '2023-03-09 21:33:49');
INSERT INTO `sys_menu` VALUES (893, 'UserAccountLog', '账变记录', 'pass', '/app/user/user-account-log', '2', 'app:user:user-account-log:list', 886, '2', '', '/app/user/user-account-log/index.vue', 0, '2', '2', 1, 1, '2023-03-09 21:33:51', '2023-03-09 21:35:31');
INSERT INTO `sys_menu` VALUES (894, '', '分页获取账变记录', '', '', '3', 'app:user:user-account-log:query', 893, '2', '', '', 0, '2', '2', 1, 1, '2023-03-09 21:33:51', '2023-03-09 21:33:51');
INSERT INTO `sys_menu` VALUES (895, '', '创建账变记录', '', '', '3', 'app:user:user-account-log:add', 893, '2', '', '', 0, '2', '2', 1, 1, '2023-03-09 21:33:51', '2023-03-09 21:33:51');
INSERT INTO `sys_menu` VALUES (896, '', '修改账变记录', '', '', '3', 'app:user:user-account-log:edit', 893, '2', '', '', 0, '2', '2', 1, 1, '2023-03-09 21:33:51', '2023-03-09 21:33:51');
INSERT INTO `sys_menu` VALUES (897, '', '删除账变记录', '', '', '3', 'app:user:user-account-log:del', 893, '2', '', '', 0, '2', '2', 1, 1, '2023-03-09 21:33:51', '2023-03-09 21:33:51');
INSERT INTO `sys_menu` VALUES (898, '', '导出账变记录', '', '', '3', 'app:user:user-account-log:export', 893, '2', '', '', 0, '2', '2', 1, 1, '2023-03-09 21:33:51', '2023-03-09 21:33:51');
INSERT INTO `sys_menu` VALUES (899, 'UserConf', '用户配置', 'pass', '/app/user/user-conf', '2', 'app:user:user-conf:list', 875, '2', '', '/app/user/user-conf/index.vue', 90, '2', '2', 1, 1, '2023-03-09 23:04:40', '2023-03-11 15:02:32');
INSERT INTO `sys_menu` VALUES (900, '', '分页获取用户配置', '', '', '3', 'app:user:user-conf:query', 899, '2', '', '', 0, '2', '2', 1, 1, '2023-03-09 23:04:40', '2023-03-09 23:04:40');
INSERT INTO `sys_menu` VALUES (901, '', '创建用户配置', '', '', '3', 'app:user:user-conf:add', 899, '2', '', '', 0, '2', '2', 1, 1, '2023-03-09 23:04:40', '2023-03-09 23:04:40');
INSERT INTO `sys_menu` VALUES (902, '', '修改用户配置', '', '', '3', 'app:user:user-conf:edit', 899, '2', '', '', 0, '2', '2', 1, 1, '2023-03-09 23:04:40', '2023-03-09 23:04:40');
INSERT INTO `sys_menu` VALUES (903, '', '删除用户配置', '', '', '3', 'app:user:user-conf:del', 899, '2', '', '', 0, '2', '2', 1, 1, '2023-03-09 23:04:40', '2023-03-09 23:04:40');
INSERT INTO `sys_menu` VALUES (904, '', '导出用户配置', '', '', '3', 'app:user:user-conf:export', 899, '2', '', '', 0, '2', '2', 1, 1, '2023-03-09 23:04:40', '2023-03-09 23:04:40');
INSERT INTO `sys_menu` VALUES (905, 'User', '用户管理', 'pass', '/app/user/user', '2', 'app:user:user:list', 875, '2', '', '/app/user/user/index.vue', 30, '2', '2', 1, 1, '2023-03-09 23:18:49', '2023-03-11 15:01:57');
INSERT INTO `sys_menu` VALUES (906, '', '分页获取用户管理', '', '', '3', 'app:user:user:query', 905, '2', '', '', 0, '2', '2', 1, 1, '2023-03-09 23:18:49', '2023-03-09 23:18:49');
INSERT INTO `sys_menu` VALUES (907, '', '创建用户管理', '', '', '3', 'app:user:user:add', 905, '2', '', '', 0, '2', '2', 1, 1, '2023-03-09 23:18:49', '2023-03-09 23:18:49');
INSERT INTO `sys_menu` VALUES (908, '', '修改用户管理', '', '', '3', 'app:user:user:edit', 905, '2', '', '', 0, '2', '2', 1, 1, '2023-03-09 23:18:49', '2023-03-09 23:18:49');
INSERT INTO `sys_menu` VALUES (909, '', '删除用户管理', '', '', '3', 'app:user:user:del', 905, '2', '', '', 0, '2', '2', 1, 1, '2023-03-09 23:18:49', '2023-03-09 23:18:49');
INSERT INTO `sys_menu` VALUES (910, '', '导出用户管理', '', '', '3', 'app:user:user:export', 905, '2', '', '', 0, '2', '2', 1, 1, '2023-03-09 23:18:49', '2023-03-09 23:18:49');
INSERT INTO `sys_menu` VALUES (911, 'UserOperLog', '用户行为记录', 'pass', '/app/user/user-oper-log', '2', 'app:user:user-oper-log:list', 875, '2', '', '/app/user/user-oper-log/index.vue', 120, '2', '2', 1, 1, '2023-03-11 15:00:06', '2023-03-11 15:02:42');
INSERT INTO `sys_menu` VALUES (912, '', '分页获取用户关键行为日志表', '', '', '3', 'app:user:user-oper-log:query', 911, '2', '', '', 0, '2', '2', 1, 1, '2023-03-11 15:00:06', '2023-03-11 15:00:06');
INSERT INTO `sys_menu` VALUES (913, '', '创建用户关键行为日志表', '', '', '3', 'app:user:user-oper-log:add', 911, '2', '', '', 0, '2', '2', 1, 1, '2023-03-11 15:00:06', '2023-03-11 15:00:06');
INSERT INTO `sys_menu` VALUES (914, '', '修改用户关键行为日志表', '', '', '3', 'app:user:user-oper-log:edit', 911, '2', '', '', 0, '2', '2', 1, 1, '2023-03-11 15:00:06', '2023-03-11 15:00:06');
INSERT INTO `sys_menu` VALUES (915, '', '删除用户关键行为日志表', '', '', '3', 'app:user:user-oper-log:del', 911, '2', '', '', 0, '2', '2', 1, 1, '2023-03-11 15:00:06', '2023-03-11 15:00:06');
INSERT INTO `sys_menu` VALUES (916, '', '导出用户关键行为日志表', '', '', '3', 'app:user:user-oper-log:export', 911, '2', '', '', 0, '2', '2', 1, 1, '2023-03-11 15:00:06', '2023-03-11 15:00:06');
INSERT INTO `sys_menu` VALUES (918, '', '消息管理', 'batch-update', '/plugins/msg', '1', '', 843, '', '', '/index', 0, '2', '2', 1, 1, '2023-03-12 13:27:59', '2023-03-12 19:52:02');
INSERT INTO `sys_menu` VALUES (961, 'MsgCode', '验证码记录', 'pass', '/plugins/msg/msg-code', '2', 'plugins:msg:msg-code:list', 918, '2', '', '/plugins/msg/msg-code/index.vue', 0, '2', '2', 1, 1, '2023-03-12 21:54:02', '2023-03-12 21:54:32');
INSERT INTO `sys_menu` VALUES (962, '', '分页获取验证码记录', '', '', '3', 'plugins:msg:msg-code:query', 961, '2', '', '', 0, '2', '2', 1, 1, '2023-03-12 21:54:02', '2023-03-12 21:54:02');
INSERT INTO `sys_menu` VALUES (963, '', '创建验证码记录', '', '', '3', 'plugins:msg:msg-code:add', 961, '2', '', '', 0, '2', '2', 1, 1, '2023-03-12 21:54:02', '2023-03-12 21:54:02');
INSERT INTO `sys_menu` VALUES (964, '', '修改验证码记录', '', '', '3', 'plugins:msg:msg-code:edit', 961, '2', '', '', 0, '2', '2', 1, 1, '2023-03-12 21:54:02', '2023-03-12 21:54:02');
INSERT INTO `sys_menu` VALUES (965, '', '删除验证码记录', '', '', '3', 'plugins:msg:msg-code:del', 961, '2', '', '', 0, '2', '2', 1, 1, '2023-03-12 21:54:02', '2023-03-12 21:54:02');
INSERT INTO `sys_menu` VALUES (966, '', '导出验证码记录', '', '', '3', 'plugins:msg:msg-code:export', 961, '2', '', '', 0, '2', '2', 1, 1, '2023-03-12 21:54:02', '2023-03-12 21:54:02');
INSERT INTO `sys_menu` VALUES (967, 'ContentAnnouncement', '公告管理', 'pass', '/plugins/content/content-announcement', '2', 'plugins:content:content-announcement:list', 778, '2', '', '/plugins/content/content-announcement/index.vue', 90, '2', '2', 1, 1, '2023-03-12 22:47:11', '2023-03-12 22:48:08');
INSERT INTO `sys_menu` VALUES (968, '', '分页获取公告管理', '', '', '3', 'plugins:content:content-announcement:query', 967, '2', '', '', 0, '2', '2', 1, 1, '2023-03-12 22:47:11', '2023-03-12 22:47:11');
INSERT INTO `sys_menu` VALUES (969, '', '创建公告管理', '', '', '3', 'plugins:content:content-announcement:add', 967, '2', '', '', 0, '2', '2', 1, 1, '2023-03-12 22:47:11', '2023-03-12 22:47:11');
INSERT INTO `sys_menu` VALUES (970, '', '修改公告管理', '', '', '3', 'plugins:content:content-announcement:edit', 967, '2', '', '', 0, '2', '2', 1, 1, '2023-03-12 22:47:11', '2023-03-12 22:47:11');
INSERT INTO `sys_menu` VALUES (971, '', '删除公告管理', '', '', '3', 'plugins:content:content-announcement:del', 967, '2', '', '', 0, '2', '2', 1, 1, '2023-03-12 22:47:11', '2023-03-12 22:47:11');
INSERT INTO `sys_menu` VALUES (972, '', '导出公告管理', '', '', '3', 'plugins:content:content-announcement:export', 967, '2', '', '', 0, '2', '2', 1, 1, '2023-03-12 22:47:11', '2023-03-12 22:47:11');
INSERT INTO `sys_menu` VALUES (973, 'ContentCategory', '内容分类', 'pass', '/plugins/content/content-category', '2', 'plugins:content:content-category:list', 778, '2', '', '/plugins/content/content-category/index.vue', 0, '2', '2', 1, 1, '2023-03-12 23:17:44', '2023-03-12 23:20:35');
INSERT INTO `sys_menu` VALUES (974, '', '分页获取内容分类', '', '', '3', 'plugins:content:content-category:query', 973, '2', '', '', 0, '2', '2', 1, 1, '2023-03-12 23:17:44', '2023-03-12 23:17:44');
INSERT INTO `sys_menu` VALUES (975, '', '创建内容分类', '', '', '3', 'plugins:content:content-category:add', 973, '2', '', '', 0, '2', '2', 1, 1, '2023-03-12 23:17:44', '2023-03-12 23:17:44');
INSERT INTO `sys_menu` VALUES (976, '', '修改内容分类', '', '', '3', 'plugins:content:content-category:edit', 973, '2', '', '', 0, '2', '2', 1, 1, '2023-03-12 23:17:44', '2023-03-12 23:17:44');
INSERT INTO `sys_menu` VALUES (977, '', '删除内容分类', '', '', '3', 'plugins:content:content-category:del', 973, '2', '', '', 0, '2', '2', 1, 1, '2023-03-12 23:17:45', '2023-03-12 23:17:45');
INSERT INTO `sys_menu` VALUES (978, '', '导出内容分类', '', '', '3', 'plugins:content:content-category:export', 973, '2', '', '', 0, '2', '2', 1, 1, '2023-03-12 23:17:45', '2023-03-12 23:17:45');
INSERT INTO `sys_menu` VALUES (979, 'ContentArticle', '文章管理', 'pass', '/plugins/content/content-article', '2', 'plugins:content:content-article:list', 778, '2', '', '/plugins/content/content-article/index.vue', 60, '2', '2', 1, 1, '2023-03-12 23:52:45', '2023-03-12 23:53:12');
INSERT INTO `sys_menu` VALUES (980, '', '分页获取文章管理', '', '', '3', 'plugins:content:content-article:query', 979, '2', '', '', 0, '2', '2', 1, 1, '2023-03-12 23:52:45', '2023-03-12 23:52:45');
INSERT INTO `sys_menu` VALUES (981, '', '创建文章管理', '', '', '3', 'plugins:content:content-article:add', 979, '2', '', '', 0, '2', '2', 1, 1, '2023-03-12 23:52:45', '2023-03-12 23:52:45');
INSERT INTO `sys_menu` VALUES (982, '', '修改文章管理', '', '', '3', 'plugins:content:content-article:edit', 979, '2', '', '', 0, '2', '2', 1, 1, '2023-03-12 23:52:45', '2023-03-12 23:52:45');
INSERT INTO `sys_menu` VALUES (983, '', '删除文章管理', '', '', '3', 'plugins:content:content-article:del', 979, '2', '', '', 0, '2', '2', 1, 1, '2023-03-12 23:52:45', '2023-03-12 23:52:45');
INSERT INTO `sys_menu` VALUES (984, '', '导出文章管理', '', '', '3', 'plugins:content:content-article:export', 979, '2', '', '', 0, '2', '2', 1, 1, '2023-03-12 23:52:46', '2023-03-12 23:52:46');
INSERT INTO `sys_menu` VALUES (985, 'FilemgrApp', 'App管理', 'pass', '/plugins/filemgr/filemgr-app', '2', 'plugins:filemgr:filemgr-app:list', 772, '2', '', '/plugins/filemgr/filemgr-app/index.vue', 0, '2', '2', 1, 1, '2023-03-13 00:55:02', '2023-03-13 00:55:52');
INSERT INTO `sys_menu` VALUES (986, '', '分页获取App管理', '', '', '3', 'plugins:filemgr:filemgr-app:query', 985, '2', '', '', 0, '2', '2', 1, 1, '2023-03-13 00:55:02', '2023-03-13 00:55:02');
INSERT INTO `sys_menu` VALUES (987, '', '创建App管理', '', '', '3', 'plugins:filemgr:filemgr-app:add', 985, '2', '', '', 0, '2', '2', 1, 1, '2023-03-13 00:55:02', '2023-03-13 00:55:02');
INSERT INTO `sys_menu` VALUES (988, '', '修改App管理', '', '', '3', 'plugins:filemgr:filemgr-app:edit', 985, '2', '', '', 0, '2', '2', 1, 1, '2023-03-13 00:55:02', '2023-03-13 00:55:02');
INSERT INTO `sys_menu` VALUES (989, '', '删除App管理', '', '', '3', 'plugins:filemgr:filemgr-app:del', 985, '2', '', '', 0, '2', '2', 1, 1, '2023-03-13 00:55:02', '2023-03-13 00:55:02');
INSERT INTO `sys_menu` VALUES (990, '', '导出App管理', '', '', '3', 'plugins:filemgr:filemgr-app:export', 985, '2', '', '', 0, '2', '2', 1, 1, '2023-03-13 00:55:02', '2023-03-13 00:55:02');
INSERT INTO `sys_menu` VALUES (991, 'UserCountryCode', '国家区号', 'pass', '/app/user/user-country-code', '2', 'app:user:user-country-code:list', 875, '2', '', '/app/user/user-country-code/index.vue', 150, '2', '2', 1, 1, '2023-03-14 17:47:44', '2023-03-14 18:06:00');
INSERT INTO `sys_menu` VALUES (992, '', '分页获取国家电话区号', '', '', '3', 'app:user:user-country-code:query', 991, '2', '', '', 0, '2', '2', 1, 1, '2023-03-14 17:47:44', '2023-03-14 17:47:44');
INSERT INTO `sys_menu` VALUES (993, '', '创建国家电话区号', '', '', '3', 'app:user:user-country-code:add', 991, '2', '', '', 0, '2', '2', 1, 1, '2023-03-14 17:47:44', '2023-03-14 17:47:44');
INSERT INTO `sys_menu` VALUES (994, '', '修改国家电话区号', '', '', '3', 'app:user:user-country-code:edit', 991, '2', '', '', 0, '2', '2', 1, 1, '2023-03-14 17:47:44', '2023-03-14 17:47:44');
INSERT INTO `sys_menu` VALUES (995, '', '删除国家电话区号', '', '', '3', 'app:user:user-country-code:del', 991, '2', '', '', 0, '2', '2', 1, 1, '2023-03-14 17:47:45', '2023-03-14 17:47:45');
INSERT INTO `sys_menu` VALUES (996, '', '导出国家电话区号', '', '', '3', 'app:user:user-country-code:export', 991, '2', '', '', 0, '2', '2', 1, 1, '2023-03-14 17:47:45', '2023-03-14 17:47:45');

-- ----------------------------
-- Table structure for sys_menu_api_rule
-- ----------------------------
DROP TABLE IF EXISTS `sys_menu_api_rule`;
CREATE TABLE `sys_menu_api_rule`  (
  `sys_menu_menu_id` int NOT NULL,
  `sys_api_id` int NOT NULL COMMENT '主键编码',
  PRIMARY KEY (`sys_menu_menu_id`, `sys_api_id`) USING BTREE,
  INDEX `fk_sys_menu_api_rule_sys_api`(`sys_api_id` ASC) USING BTREE,
  CONSTRAINT `fk_sys_menu_api_rule_sys_api` FOREIGN KEY (`sys_api_id`) REFERENCES `sys_api` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
  CONSTRAINT `fk_sys_menu_api_rule_sys_menu` FOREIGN KEY (`sys_menu_menu_id`) REFERENCES `sys_menu` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_bin ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of sys_menu_api_rule
-- ----------------------------

-- ----------------------------
-- Table structure for sys_oper_log
-- ----------------------------
DROP TABLE IF EXISTS `sys_oper_log`;
CREATE TABLE `sys_oper_log`  (
  `id` int NOT NULL AUTO_INCREMENT COMMENT '主键编码',
  `request_method` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL COMMENT '请求方式',
  `user_id` int NULL DEFAULT NULL COMMENT '操作者',
  `oper_url` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL COMMENT '访问地址',
  `oper_ip` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL COMMENT '客户端ip',
  `oper_location` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL COMMENT '访问位置',
  `status` varchar(4) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL COMMENT '操作状态',
  `oper_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '操作时间',
  `json_result` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL COMMENT '返回数据',
  `remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL COMMENT '备注',
  `latency_time` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL COMMENT '耗时',
  `user_agent` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL COMMENT 'ua',
  `created_at` datetime NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime NULL DEFAULT NULL COMMENT '最后更新时间',
  `create_by` int NULL DEFAULT NULL COMMENT '创建者',
  `update_by` int NULL DEFAULT NULL COMMENT '更新者',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_sys_opera_log_create_by`(`create_by` ASC) USING BTREE,
  INDEX `idx_sys_opera_log_update_by`(`update_by` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 11074 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_bin COMMENT = '操作日志' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of sys_oper_log
-- ----------------------------

-- ----------------------------
-- Table structure for sys_post
-- ----------------------------
DROP TABLE IF EXISTS `sys_post`;
CREATE TABLE `sys_post`  (
  `id` int NOT NULL AUTO_INCREMENT,
  `post_name` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL,
  `post_code` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL,
  `sort` tinyint NULL DEFAULT NULL,
  `status` char(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL,
  `remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL,
  `create_by` int NULL DEFAULT NULL COMMENT '创建者',
  `update_by` int NULL DEFAULT NULL COMMENT '更新者',
  `created_at` datetime NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime NULL DEFAULT NULL COMMENT '最后更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_sys_post_create_by`(`create_by` ASC) USING BTREE,
  INDEX `idx_sys_post_update_by`(`update_by` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 4 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_bin ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of sys_post
-- ----------------------------
INSERT INTO `sys_post` VALUES (1, '首席执行官', 'CEO', 0, '1', '首席执行官', 1, 1, '2021-05-13 19:56:38', '2021-05-13 19:56:38');
INSERT INTO `sys_post` VALUES (2, '首席技术执行官', 'CTO', 2, '1', '首席技术执行官', 1, 1, '2021-05-13 19:56:38', '2021-05-13 19:56:38');
INSERT INTO `sys_post` VALUES (3, '首席运营官', 'COO', 3, '1', '测试工程师', 1, 1, '2021-05-13 19:56:38', '2021-05-13 19:56:38');

-- ----------------------------
-- Table structure for sys_role
-- ----------------------------
DROP TABLE IF EXISTS `sys_role`;
CREATE TABLE `sys_role`  (
  `id` int NOT NULL AUTO_INCREMENT,
  `role_name` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL,
  `role_key` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL,
  `role_sort` bigint NULL DEFAULT NULL,
  `remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL,
  `data_scope` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL,
  `status` char(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL COMMENT '状态 1-正常 2-停用',
  `create_by` int NULL DEFAULT NULL COMMENT '创建者',
  `update_by` int NULL DEFAULT NULL COMMENT '更新者',
  `created_at` datetime(3) NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime(3) NULL DEFAULT NULL COMMENT '最后更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_sys_role_create_by`(`create_by` ASC) USING BTREE,
  INDEX `idx_sys_role_update_by`(`update_by` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 18 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_bin ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of sys_role
-- ----------------------------
INSERT INTO `sys_role` VALUES (1, '系统管理员', 'admin', 1, '', '', '1', 1, 1, '2021-05-13 19:56:37.913', '2023-03-03 01:04:03.641');

-- ----------------------------
-- Table structure for sys_role_dept
-- ----------------------------
DROP TABLE IF EXISTS `sys_role_dept`;
CREATE TABLE `sys_role_dept`  (
  `role_id` smallint NOT NULL,
  `dept_id` smallint NOT NULL,
  PRIMARY KEY (`role_id`, `dept_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_bin ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of sys_role_dept
-- ----------------------------

-- ----------------------------
-- Table structure for sys_role_menu
-- ----------------------------
DROP TABLE IF EXISTS `sys_role_menu`;
CREATE TABLE `sys_role_menu`  (
  `role_id` int NOT NULL,
  `menu_id` int NOT NULL,
  PRIMARY KEY (`role_id`, `menu_id`) USING BTREE,
  INDEX `fk_sys_role_menu_sys_menu`(`menu_id` ASC) USING BTREE,
  CONSTRAINT `fk_sys_role_menu_sys_menu` FOREIGN KEY (`menu_id`) REFERENCES `sys_menu` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
  CONSTRAINT `fk_sys_role_menu_sys_role` FOREIGN KEY (`role_id`) REFERENCES `sys_role` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_bin ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of sys_role_menu
-- ----------------------------

-- ----------------------------
-- Table structure for sys_user
-- ----------------------------
DROP TABLE IF EXISTS `sys_user`;
CREATE TABLE `sys_user`  (
  `id` int NOT NULL AUTO_INCREMENT COMMENT '编码',
  `username` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL COMMENT '用户名',
  `password` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL COMMENT '密码',
  `nick_name` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL COMMENT '昵称',
  `phone` varchar(11) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL COMMENT '手机号',
  `role_id` int NULL DEFAULT NULL COMMENT '角色ID',
  `salt` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL COMMENT '加盐',
  `avatar` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL COMMENT '头像',
  `sex` char(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL COMMENT '性别',
  `email` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL COMMENT '邮箱',
  `dept_id` int NULL DEFAULT NULL COMMENT '部门',
  `post_id` int NULL DEFAULT NULL COMMENT '岗位',
  `remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL COMMENT '备注',
  `status` varchar(4) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL COMMENT '状态',
  `create_by` int NULL DEFAULT NULL COMMENT '创建者',
  `update_by` int NULL DEFAULT NULL COMMENT '更新者',
  `created_at` datetime NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime NULL DEFAULT NULL COMMENT '最后更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_sys_user_create_by`(`create_by` ASC) USING BTREE,
  INDEX `idx_sys_user_update_by`(`update_by` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 10 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_bin ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of sys_user
-- ----------------------------
INSERT INTO `sys_user` VALUES (1, 'admin', '$2a$10$ZFMtvc.ROtYwk2UNOaBLCOrpr.Mq/i1ae4PVZfoWgHTb4ffORW/lm', 'admin', '13700000000', 1, '', '/files/admin/avatar/5226ae82-349c-48cc-b312-b19f78233086.jpg', '1', 'admin@admin.com', 1, 1, '', '1', 1, 1, '2021-05-13 19:56:38', '2023-03-14 09:27:36');

SET FOREIGN_KEY_CHECKS = 1;
