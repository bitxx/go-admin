/*
 Navicat Premium Data Transfer

 Source Server         : Jason
 Source Server Type    : MySQL
 Source Server Version : 80033
 Source Host           : 127.0.0.1:3306
 Source Schema         : app

 Target Server Type    : MySQL
 Target Server Version : 80033
 File Encoding         : 65001

 Date: 19/10/2023 14:15:16
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

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
  `can_login` char(1) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '0' COMMENT '1-允许登陆；2-不允许登陆',
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
  `code` varchar(12) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '区号',
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
INSERT INTO `plugins_content_article` (`id`, `cate_id`, `name`, `content`, `remark`, `status`, `create_by`, `update_by`, `updated_at`, `created_at`) VALUES (1,1, 'test', '<p>test</p>', '111', '1', 1, 1, '2023-03-13 00:04:40', '2023-03-13 00:04:40');
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
  `download_num` int DEFAULT '0' COMMENT '下载数量',
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
INSERT INTO `plugins_filemgr_app` (`id`, `version`, `platform`, `app_type`, `local_address`, `download_num`, `download_type`, `download_url`, `remark`, `status`, `create_by`, `created_at`, `update_by`, `updated_at`) VALUES (1, '1.0.1', '1', '1', 'files/app/4b6ea3c0-d7fa-49f1-9d50-f9d73caad45f.apk', 0, '3', '', 'test', '1', 1, '2023-03-12 11:34:54', 1, '2023-03-13 01:00:30');
INSERT INTO `plugins_filemgr_app` (`id`, `version`, `platform`, `app_type`, `local_address`, `download_num`, `download_type`, `download_url`, `remark`, `status`, `create_by`, `created_at`, `update_by`, `updated_at`) VALUES (2, '1.0.0', '1', '1', 'files/app/ba7b81c0-e6d2-42ee-82e4-2dcbec720c23.apk', 0, '1', 'http://localhost:9999/files/app/ba7b81c0-e6d2-42ee-82e4-2dcbec720c23.apk', 'test', '1', 1, '2023-03-13 01:06:21', 1, '2023-03-13 01:06:21');
INSERT INTO `plugins_filemgr_app` (`id`, `version`, `platform`, `app_type`, `local_address`, `download_num`, `download_type`, `download_url`, `remark`, `status`, `create_by`, `created_at`, `update_by`, `updated_at`) VALUES (3, '1.0.2', '1', '1', '', 0, '2', 'http://localhost:9999/test.apk', 'test2', '1', 1, '2023-03-13 01:07:00', 1, '2023-03-13 01:07:00');
INSERT INTO `plugins_filemgr_app` (`id`, `version`, `platform`, `app_type`, `local_address`, `download_num`, `download_type`, `download_url`, `remark`, `status`, `create_by`, `created_at`, `update_by`, `updated_at`) VALUES (4, '1.0.3', '1', '1', 'files/app/962bebc9-fdb6-41b5-b62b-b184ee2fd1c0.apk', 0, '3', '', 'test2', '1', 1, '2023-03-13 01:07:24', 1, '2023-03-13 01:07:24');
INSERT INTO `plugins_filemgr_app` (`id`, `version`, `platform`, `app_type`, `local_address`, `download_num`, `download_type`, `download_url`, `remark`, `status`, `create_by`, `created_at`, `update_by`, `updated_at`) VALUES (5, '2.0.0', '1', '1', 'files/app/9ab7880d-4f64-40ba-8804-3023ea5e93ff.apk', 0, '1', 'https://localhost:8888/files/app/9ab7880d-4f64-40ba-8804-3023ea5e93ff.apk', 'test', '2', 1, '2023-03-14 17:29:09', 1, '2023-03-14 17:29:18');
COMMIT;

-- ----------------------------
-- Table structure for plugins_msg_code
-- ----------------------------
DROP TABLE IF EXISTS `plugins_msg_code`;
CREATE TABLE `plugins_msg_code` (
  `id` int NOT NULL AUTO_INCREMENT COMMENT '验证码编号',
  `user_id` int NOT NULL COMMENT '用户编号',
  `code` varchar(12) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '0' COMMENT '验证码',
  `code_type` char(1) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '0' COMMENT '验证码类型 1-邮箱；2-短信',
  `remark` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '备注异常',
  `status` char(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '2' COMMENT '验证码状态 1-发送成功 2-发送失败',
  `create_by` int NOT NULL DEFAULT '0' COMMENT '创建者',
  `update_by` int NOT NULL DEFAULT '0' COMMENT '更新者',
  `created_at` datetime NOT NULL COMMENT '创建时间',
  `updated_at` datetime NOT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb3 ROW_FORMAT=COMPACT COMMENT='验证码记录';

-- ----------------------------
-- Records of plugins_msg_code
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for sys_api
-- ----------------------------
DROP TABLE IF EXISTS `sys_api`;
CREATE TABLE `sys_api` (
  `id` int NOT NULL AUTO_INCREMENT COMMENT '主键编码',
  `handle` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT 'handle',
  `title` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '标题',
  `path` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '地址',
  `api_type` varchar(16) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '接口类型',
  `action` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '请求类型',
  `created_at` datetime DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime DEFAULT NULL COMMENT '最后更新时间',
  `create_by` int DEFAULT NULL COMMENT '创建者',
  `update_by` int DEFAULT NULL COMMENT '更新者',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=153 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;

-- ----------------------------
-- Records of sys_api
-- ----------------------------
BEGIN;
INSERT INTO `sys_api` (`id`, `handle`, `title`, `path`, `api_type`, `action`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (1, 'go-admin/app/admin/apis.SysTables.GetPage-fm', '系统-表信息列表获取', '/admin-api/v1/sys/table', '1', 'GET', '2023-05-08 16:51:10', '2023-05-09 09:53:14', 0, 1);
INSERT INTO `sys_api` (`id`, `handle`, `title`, `path`, `api_type`, `action`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (2, 'go-admin/app/admin/apis.SysTables.DownloadCode-fm', '系统-下载代码', '/admin-api/v1/sys/table/gen/download/:id', '1', 'GET', '2023-05-08 16:51:10', '2023-05-09 10:55:36', 0, 1);
INSERT INTO `sys_api` (`id`, `handle`, `title`, `path`, `api_type`, `action`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (3, 'go-admin/app/admin/apis.SysTables.GenDB-fm', '系统-表信息详情', '/admin-api/v1/sys/table/gen/db/:id', '1', 'GET', '2023-05-08 16:51:10', '2023-05-09 09:53:25', 0, 1);
INSERT INTO `sys_api` (`id`, `handle`, `title`, `path`, `api_type`, `action`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (4, 'go-admin/app/admin/apis.SysTables.GenCode-fm', '系统-表生成代码', '/admin-api/v1/sys/table/gen/:id', '1', 'GET', '2023-05-08 16:51:11', '2023-05-09 10:55:36', 0, 1);
INSERT INTO `sys_api` (`id`, `handle`, `title`, `path`, `api_type`, `action`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (5, 'go-admin/app/admin/apis.SysTables.GetDBTablePage-fm', '系统-获取数据库表', '/admin-api/v1/sys/table/dbtables', '1', 'GET', '2023-05-08 16:51:11', '2023-05-09 10:55:36', 0, 1);
INSERT INTO `sys_api` (`id`, `handle`, `title`, `path`, `api_type`, `action`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (6, 'go-admin/app/admin/apis.SysTables.Preview-fm', '系统-代码预览', '/admin-api/v1/sys/table/preview/:id', '1', 'GET', '2023-05-08 16:51:11', '2023-05-09 10:55:36', 0, 1);
INSERT INTO `sys_api` (`id`, `handle`, `title`, `path`, `api_type`, `action`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (7, 'go-admin/app/admin/apis.SysTables.Get-fm', '系统-表信息详情获取', '/admin-api/v1/sys/table/:id', '1', 'GET', '2023-05-08 16:51:12', '2023-05-09 09:52:44', 0, 1);
INSERT INTO `sys_api` (`id`, `handle`, `title`, `path`, `api_type`, `action`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (8, 'go-admin/app/admin/apis.SysDept.GetList-fm', '系统-部门列表获取', '/admin-api/v1/sys/dept', '1', 'GET', '2023-05-08 16:51:12', '2023-05-09 09:52:50', 0, 1);
INSERT INTO `sys_api` (`id`, `handle`, `title`, `path`, `api_type`, `action`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (9, 'go-admin/app/admin/apis.SysDept.Get-fm', '系统-部门详情获取', '/admin-api/v1/sys/dept/:id', '1', 'GET', '2023-05-08 16:51:12', '2023-05-09 09:52:55', 0, 1);
INSERT INTO `sys_api` (`id`, `handle`, `title`, `path`, `api_type`, `action`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (10, 'go-admin/app/admin/apis.SysDept.Get2Tree-fm', '系统-树部门', '/admin-api/v1/sys/deptTree', '1', 'GET', '2023-05-08 16:51:13', '2023-05-09 09:52:27', 0, 1);
INSERT INTO `sys_api` (`id`, `handle`, `title`, `path`, `api_type`, `action`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (11, 'go-admin/app/admin/apis.SysPost.GetPage-fm', '系统-岗位列表获取', '/admin-api/v1/sys/post', '1', 'GET', '2023-05-08 16:51:13', '2023-05-09 09:52:31', 0, 1);
INSERT INTO `sys_api` (`id`, `handle`, `title`, `path`, `api_type`, `action`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (12, 'go-admin/app/admin/apis.SysPost.Export-fm', '系统-部门列表导出', '/admin-api/v1/sys/post/export', '1', 'GET', '2023-05-08 16:51:13', '2023-05-09 09:52:39', 0, 1);
INSERT INTO `sys_api` (`id`, `handle`, `title`, `path`, `api_type`, `action`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (13, 'go-admin/app/admin/apis.SysPost.Get-fm', '系统-岗位详情获取', '/admin-api/v1/sys/post/:id', '1', 'GET', '2023-05-08 16:51:14', '2023-05-09 09:52:07', 0, 1);
INSERT INTO `sys_api` (`id`, `handle`, `title`, `path`, `api_type`, `action`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (14, 'go-admin/app/admin/apis.SysMenu.GetMenuTreeSelect-fm', '系统-树菜单角色', '/admin-api/v1/sys/roleMenuTreeselect/:roleId', '1', 'GET', '2023-05-08 16:51:14', '2023-05-09 09:52:12', 0, 1);
INSERT INTO `sys_api` (`id`, `handle`, `title`, `path`, `api_type`, `action`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (15, 'go-admin/app/admin/apis.SysDept.GetDeptTreeRoleSelect-fm', '系统-树部门角色', '/admin-api/v1/sys/roleDeptTreeselect/:roleId', '1', 'GET', '2023-05-08 16:51:14', '2023-05-09 09:52:17', 0, 1);
INSERT INTO `sys_api` (`id`, `handle`, `title`, `path`, `api_type`, `action`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (16, 'go-admin/app/admin/apis.SysApi.GetPage-fm', '系统-接口列表获取', '/admin-api/v1/sys-api', '1', 'GET', '2023-05-08 16:51:15', '2023-05-09 09:51:46', 0, 1);
INSERT INTO `sys_api` (`id`, `handle`, `title`, `path`, `api_type`, `action`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (17, 'go-admin/app/admin/apis.SysApi.Export-fm', '系统-接口列表导出', '/admin-api/v1/sys-api/export', '1', 'GET', '2023-05-08 16:51:15', '2023-05-09 09:51:56', 0, 1);
INSERT INTO `sys_api` (`id`, `handle`, `title`, `path`, `api_type`, `action`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (18, 'go-admin/app/admin/apis.SysApi.Get-fm', '系统-接口详情获取', '/admin-api/v1/sys-api/:id', '1', 'GET', '2023-05-08 16:51:15', '2023-05-09 09:52:02', 0, 1);
INSERT INTO `sys_api` (`id`, `handle`, `title`, `path`, `api_type`, `action`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (19, 'go-admin/app/admin/apis.SysLoginLog.GetPage-fm', '系统-登录日志列表获取', '/admin-api/v1/sys-login-log', '1', 'GET', '2023-05-08 16:51:16', '2023-05-09 11:03:23', 0, 1);
INSERT INTO `sys_api` (`id`, `handle`, `title`, `path`, `api_type`, `action`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (20, 'go-admin/app/admin/apis.SysLoginLog.Export-fm', '系统-登录日志列表导出', '/admin-api/v1/sys-login-log/export', '1', 'GET', '2023-05-08 16:51:16', '2023-05-09 09:51:32', 0, 1);
INSERT INTO `sys_api` (`id`, `handle`, `title`, `path`, `api_type`, `action`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (21, 'go-admin/app/admin/apis.SysLoginLog.Get-fm', '系统-登录日志详情获取', '/admin-api/v1/sys-login-log/:id', '1', 'GET', '2023-05-08 16:51:16', '2023-05-09 09:51:36', 0, 1);
INSERT INTO `sys_api` (`id`, `handle`, `title`, `path`, `api_type`, `action`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (22, 'go-admin/app/admin/apis.SysOperLog.GetPage-fm', '系统-操作日志列表获取', '/admin-api/v1/sys-oper-log', '1', 'GET', '2023-05-08 16:51:17', '2023-05-09 11:00:44', 0, 1);
INSERT INTO `sys_api` (`id`, `handle`, `title`, `path`, `api_type`, `action`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (23, 'go-admin/app/admin/apis.SysOperLog.Export-fm', '系统-操作日志列表导出', '/admin-api/v1/sys-oper-log/export', '1', 'GET', '2023-05-08 16:51:17', '2023-05-09 09:50:53', 0, 1);
INSERT INTO `sys_api` (`id`, `handle`, `title`, `path`, `api_type`, `action`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (24, 'go-admin/app/admin/apis.SysOperLog.Get-fm', '系统-操作日志详情获取', '/admin-api/v1/sys-oper-log/:id', '1', 'GET', '2023-05-08 16:51:17', '2023-05-09 09:51:23', 0, 1);
INSERT INTO `sys_api` (`id`, `handle`, `title`, `path`, `api_type`, `action`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (25, 'go-admin/app/admin/apis.SysUser.GetPage-fm', '系统-用户列表获取', '/admin-api/v1/sys-user', '1', 'GET', '2023-05-08 16:51:18', '2023-05-09 09:50:22', 0, 1);
INSERT INTO `sys_api` (`id`, `handle`, `title`, `path`, `api_type`, `action`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (26, 'go-admin/app/admin/apis.SysUser.GetProfile-fm', '系统-登录用户信息获取', '/admin-api/v1/sys-user/profile', '1', 'GET', '2023-05-08 16:51:18', '2023-05-09 09:50:29', 0, 1);
INSERT INTO `sys_api` (`id`, `handle`, `title`, `path`, `api_type`, `action`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (27, 'go-admin/app/admin/apis.SysUser.Get-fm', '系统-用户详情获取', '/admin-api/v1/sys-user/:id', '1', 'GET', '2023-05-08 16:51:18', '2023-05-09 09:50:35', 0, 1);
INSERT INTO `sys_api` (`id`, `handle`, `title`, `path`, `api_type`, `action`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (28, 'go-admin/app/admin/apis.SysRuntimeConfig.GetConfig-fm', '系统-获取运行时配置', '/admin-api/v1/sysRuntimeConfig/getConfig', '1', 'GET', '2023-05-08 16:51:19', '2023-05-09 09:49:57', 0, 1);
INSERT INTO `sys_api` (`id`, `handle`, `title`, `path`, `api_type`, `action`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (29, 'go-admin/app/admin/apis.ServerMonitor.ServerInfo-fm', '系统-获取运行环境状态', '/admin-api/v1/server-monitor', '1', 'GET', '2023-05-08 16:51:19', '2023-05-09 09:50:03', 0, 1);
INSERT INTO `sys_api` (`id`, `handle`, `title`, `path`, `api_type`, `action`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (30, 'go-admin/app/app/user/apis.User.GetPage-fm', '应用-用户列表获取', '/admin-api/v1/app/user/user', '3', 'GET', '2023-05-08 16:51:19', '2023-05-09 09:50:11', 0, 1);
INSERT INTO `sys_api` (`id`, `handle`, `title`, `path`, `api_type`, `action`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (31, 'go-admin/app/app/user/apis.UserCountryCode.GetPage-fm', '应用-国家区号列表获取', '/admin-api/v1/app/user/user-country-code', '3', 'GET', '2023-05-08 16:51:20', '2023-05-09 10:31:00', 0, 1);
INSERT INTO `sys_api` (`id`, `handle`, `title`, `path`, `api_type`, `action`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (32, 'go-admin/app/app/user/apis.UserCountryCode.Export-fm', '应用-国家区号列表导出', '/admin-api/v1/app/user/user-country-code/export', '3', 'GET', '2023-05-08 16:51:20', '2023-05-09 10:31:46', 0, 1);
INSERT INTO `sys_api` (`id`, `handle`, `title`, `path`, `api_type`, `action`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (33, 'go-admin/app/app/user/apis.UserCountryCode.Get-fm', '应用-国家区号详情获取', '/admin-api/v1/app/user/user-country-code/:id', '3', 'GET', '2023-05-08 16:51:20', '2023-05-09 10:31:24', 0, 1);
INSERT INTO `sys_api` (`id`, `handle`, `title`, `path`, `api_type`, `action`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (34, 'go-admin/app/app/user/apis.UserConf.GetPage-fm', '应用-用户配置列表获取', '/admin-api/v1/app/user/user-conf', '3', 'GET', '2023-05-08 16:51:20', '2023-05-09 10:24:15', 0, 1);
INSERT INTO `sys_api` (`id`, `handle`, `title`, `path`, `api_type`, `action`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (35, 'go-admin/app/app/user/apis.UserConf.Get-fm', '应用-用户配置详情获取', '/admin-api/v1/app/user/user-conf/:id', '3', 'GET', '2023-05-08 16:51:21', '2023-05-09 10:25:16', 0, 1);
INSERT INTO `sys_api` (`id`, `handle`, `title`, `path`, `api_type`, `action`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (36, 'go-admin/app/app/user/apis.UserAccountLog.GetPage-fm', '应用-用户账变日志列表获取', '/admin-api/v1/app/user/user-account-log', '3', 'GET', '2023-05-08 16:51:21', '2023-05-09 10:32:28', 0, 1);
INSERT INTO `sys_api` (`id`, `handle`, `title`, `path`, `api_type`, `action`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (37, 'go-admin/app/app/user/apis.UserAccountLog.Export-fm', '应用-用户账变日志列表导出', '/admin-api/v1/app/user/user-account-log/export', '3', 'GET', '2023-05-08 16:51:21', '2023-05-09 10:33:04', 0, 1);
INSERT INTO `sys_api` (`id`, `handle`, `title`, `path`, `api_type`, `action`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (38, 'go-admin/app/app/user/apis.UserAccountLog.Get-fm', '应用-用户账变列表详情获取', '/admin-api/v1/app/user/user-account-log/:id', '3', 'GET', '2023-05-08 16:51:22', '2023-05-09 09:49:10', 0, 1);
INSERT INTO `sys_api` (`id`, `handle`, `title`, `path`, `api_type`, `action`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (39, 'go-admin/app/app/user/apis.UserLevel.GetPage-fm', '应用-用户等级列表获取', '/admin-api/v1/app/user/user-level', '3', 'GET', '2023-05-08 16:51:22', '2023-05-09 10:22:48', 0, 1);
INSERT INTO `sys_api` (`id`, `handle`, `title`, `path`, `api_type`, `action`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (40, 'go-admin/app/app/user/apis.UserLevel.Export-fm', '应用-用户等级列表导出', '/admin-api/v1/app/user/user-level/export', '3', 'GET', '2023-05-08 16:51:22', '2023-05-09 10:23:54', 0, 1);
INSERT INTO `sys_api` (`id`, `handle`, `title`, `path`, `api_type`, `action`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (41, 'go-admin/app/app/user/apis.UserLevel.Get-fm', '应用-用户等级详情获取', '/admin-api/v1/app/user/user-level/:id', '3', 'GET', '2023-05-08 16:51:23', '2023-05-09 10:23:27', 0, 1);
INSERT INTO `sys_api` (`id`, `handle`, `title`, `path`, `api_type`, `action`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (42, 'go-admin/app/app/user/apis.UserOperLog.GetPage-fm', '应用-用户操作日志列表获取', '/admin-api/v1/app/user/user-oper-log', '3', 'GET', '2023-05-08 16:51:23', '2023-05-09 10:26:19', 0, 1);
INSERT INTO `sys_api` (`id`, `handle`, `title`, `path`, `api_type`, `action`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (43, 'go-admin/app/app/user/apis.UserOperLog.Export-fm', '应用-用户操作日志列表导出', '/admin-api/v1/app/user/user-oper-log/export', '3', 'GET', '2023-05-08 16:51:23', '2023-05-09 10:28:23', 0, 1);
INSERT INTO `sys_api` (`id`, `handle`, `title`, `path`, `api_type`, `action`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (44, 'go-admin/app/app/user/apis.UserOperLog.Get-fm', '应用-用户操作日志详情获取', '/admin-api/v1/app/user/user-oper-log/:id', '3', 'GET', '2023-05-08 16:51:24', '2023-05-09 09:47:59', 0, 1);
INSERT INTO `sys_api` (`id`, `handle`, `title`, `path`, `api_type`, `action`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (45, 'go-admin/app/app/user/apis.User.Export-fm', '应用-用户列表导出', '/admin-api/v1/app/user/user/export', '3', 'GET', '2023-05-08 16:51:24', '2023-05-09 09:48:04', 0, 1);
INSERT INTO `sys_api` (`id`, `handle`, `title`, `path`, `api_type`, `action`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (46, 'go-admin/app/app/user/apis.User.Get-fm', '应用-用户详情获取', '/admin-api/v1/app/user/user/:id', '3', 'GET', '2023-05-08 16:51:24', '2023-05-09 09:48:08', 0, 1);
INSERT INTO `sys_api` (`id`, `handle`, `title`, `path`, `api_type`, `action`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (47, 'go-admin/app/admin/apis.SysConfig.GetSysConfigBySysApp-fm', '系统-获取所有系统后台业务配置', '/admin-api/v1/app-config', '1', 'GET', '2023-05-08 16:51:25', '2023-05-09 09:47:48', 0, 1);
INSERT INTO `sys_api` (`id`, `handle`, `title`, `path`, `api_type`, `action`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (48, 'go-admin/app/plugins/content/apis.ContentAnnouncement.GetPage-fm', '插件-内容管理-公告列表获取', '/admin-api/v1/plugins/content/content-announcement', '2', 'GET', '2023-05-08 16:51:25', '2023-05-09 10:49:57', 0, 1);
INSERT INTO `sys_api` (`id`, `handle`, `title`, `path`, `api_type`, `action`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (49, 'go-admin/app/plugins/content/apis.ContentAnnouncement.Export-fm', '插件-内容管理-公告列表导出', '/admin-api/v1/plugins/content/content-announcement/export', '2', 'GET', '2023-05-08 16:51:25', '2023-05-09 10:52:17', 0, 1);
INSERT INTO `sys_api` (`id`, `handle`, `title`, `path`, `api_type`, `action`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (50, 'go-admin/app/plugins/content/apis.ContentAnnouncement.Get-fm', '插件-内容管理-公告详情获取', '/admin-api/v1/plugins/content/content-announcement/:id', '2', 'GET', '2023-05-08 16:51:26', '2023-05-09 10:51:56', 0, 1);
INSERT INTO `sys_api` (`id`, `handle`, `title`, `path`, `api_type`, `action`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (51, 'go-admin/app/plugins/content/apis.ContentArticle.GetPage-fm', '插件-内容管理-文章列表获取', '/admin-api/v1/plugins/content/content-article', '2', 'GET', '2023-05-08 16:51:26', '2023-05-09 10:47:48', 0, 1);
INSERT INTO `sys_api` (`id`, `handle`, `title`, `path`, `api_type`, `action`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (52, 'go-admin/app/plugins/content/apis.ContentArticle.Export-fm', '插件-内容管理-文章列表导出', '/admin-api/v1/plugins/content/content-article/export', '2', 'GET', '2023-05-08 16:51:26', '2023-05-09 10:48:53', 0, 1);
INSERT INTO `sys_api` (`id`, `handle`, `title`, `path`, `api_type`, `action`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (53, 'go-admin/app/plugins/content/apis.ContentArticle.Get-fm', '插件-内容管理-文章详情获取', '/admin-api/v1/plugins/content/content-article/:id', '2', 'GET', '2023-05-08 16:51:26', '2023-05-09 10:48:27', 0, 1);
INSERT INTO `sys_api` (`id`, `handle`, `title`, `path`, `api_type`, `action`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (54, 'go-admin/app/plugins/content/apis.ContentCategory.GetPage-fm', '插件-内容管理-文章分类列表获取', '/admin-api/v1/plugins/content/content-category', '2', 'GET', '2023-05-08 16:51:27', '2023-05-09 10:43:56', 0, 1);
INSERT INTO `sys_api` (`id`, `handle`, `title`, `path`, `api_type`, `action`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (55, 'go-admin/app/plugins/content/apis.ContentCategory.Export-fm', '插件-内容管理-文章分类列表导出', '/admin-api/v1/plugins/content/content-category/export', '2', 'GET', '2023-05-08 16:51:27', '2023-05-09 10:47:20', 0, 1);
INSERT INTO `sys_api` (`id`, `handle`, `title`, `path`, `api_type`, `action`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (56, 'go-admin/app/plugins/content/apis.ContentCategory.Get-fm', '插件-内容管理-文章分类详情获取', '/admin-api/v1/plugins/content/content-category/:id', '2', 'GET', '2023-05-08 16:51:27', '2023-05-09 10:44:33', 0, 1);
INSERT INTO `sys_api` (`id`, `handle`, `title`, `path`, `api_type`, `action`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (57, 'go-admin/app/plugins/filemgr/apis.FilemgrApp.GetPage-fm', '插件-文件管理-app列表获取', '/admin-api/v1/plugins/filemgr/filemgr-app', '2', 'GET', '2023-05-08 16:51:28', '2023-05-09 10:51:29', 0, 1);
INSERT INTO `sys_api` (`id`, `handle`, `title`, `path`, `api_type`, `action`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (58, 'go-admin/app/plugins/filemgr/apis.FilemgrApp.Export-fm', '插件-文件管理-app数据导出', '/admin-api/v1/plugins/filemgr/filemgr-app/export', '2', 'GET', '2023-05-08 16:51:28', '2023-05-09 09:46:30', 0, 1);
INSERT INTO `sys_api` (`id`, `handle`, `title`, `path`, `api_type`, `action`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (59, 'go-admin/app/plugins/filemgr/apis.FilemgrApp.Get-fm', '插件-文件管理-app详情获取', '/admin-api/v1/plugins/filemgr/filemgr-app/:id', '2', 'GET', '2023-05-08 16:51:29', '2023-05-09 09:45:44', 0, 1);
INSERT INTO `sys_api` (`id`, `handle`, `title`, `path`, `api_type`, `action`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (60, 'go-admin/app/plugins/msg/apis.MsgCode.GetPage-fm', '插件-消息-验证码列表获取', '/admin-api/v1/plugins/msg/msg-code', '2', 'GET', '2023-05-08 16:51:29', '2023-05-09 10:34:35', 0, 1);
INSERT INTO `sys_api` (`id`, `handle`, `title`, `path`, `api_type`, `action`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (61, 'go-admin/app/plugins/msg/apis.MsgCode.Get-fm', '插件-消息-验证码详情获取', '/admin-api/v1/plugins/msg/msg-code/:id', '2', 'GET', '2023-05-08 16:51:29', '2023-05-09 09:46:08', 0, 1);
INSERT INTO `sys_api` (`id`, `handle`, `title`, `path`, `api_type`, `action`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (62, 'go-admin/app/admin/apis.SysDictType.GetPage-fm', '系统-字典类型列表获取', '/admin-api/v1/dict/type', '1', 'GET', '2023-05-08 16:51:30', '2023-05-09 11:13:23', 0, 1);
INSERT INTO `sys_api` (`id`, `handle`, `title`, `path`, `api_type`, `action`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (63, 'go-admin/app/admin/apis.SysDictType.Export-fm', '系统-字典类型列表导出', '/admin-api/v1/dict/type/export', '1', 'GET', '2023-05-08 16:51:30', '2023-05-09 11:16:42', 0, 1);
INSERT INTO `sys_api` (`id`, `handle`, `title`, `path`, `api_type`, `action`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (64, 'go-admin/app/admin/apis.SysDictType.Get-fm', '系统-字典类型详情获取', '/admin-api/v1/dict/type/:id', '1', 'GET', '2023-05-08 16:51:30', '2023-05-09 09:44:58', 0, 1);
INSERT INTO `sys_api` (`id`, `handle`, `title`, `path`, `api_type`, `action`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (65, 'go-admin/app/admin/apis.SysDictType.GetAll-fm', '系统-字典类型获取', '/admin-api/v1/dict/type-option-select', '1', 'GET', '2023-05-08 16:51:31', '2023-05-09 11:13:58', 0, 1);
INSERT INTO `sys_api` (`id`, `handle`, `title`, `path`, `api_type`, `action`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (66, 'go-admin/app/admin/apis.SysDictData.GetPage-fm', '系统-字典数据列表获取', '/admin-api/v1/dict/data', '1', 'GET', '2023-05-08 16:51:31', '2023-05-09 11:09:55', 0, 1);
INSERT INTO `sys_api` (`id`, `handle`, `title`, `path`, `api_type`, `action`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (67, 'go-admin/app/admin/apis.SysDictData.Export-fm', '系统-字典数据列表导出', '/admin-api/v1/dict/data/export', '1', 'GET', '2023-05-08 16:51:31', '2023-05-09 11:15:14', 0, 1);
INSERT INTO `sys_api` (`id`, `handle`, `title`, `path`, `api_type`, `action`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (68, 'go-admin/app/admin/apis.SysDictData.Get-fm', '系统-获取字典数据', '/admin-api/v1/dict/data/:id', '1', 'GET', '2023-05-08 16:51:32', '2023-05-09 11:11:11', 0, 1);
INSERT INTO `sys_api` (`id`, `handle`, `title`, `path`, `api_type`, `action`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (69, 'go-admin/app/admin/apis.SysDictData.GetSysDictDataAll-fm', '系统-获取所有字典数据', '/admin-api/v1/dict-data/option-select', '1', 'GET', '2023-05-08 16:51:32', '2023-05-09 11:14:37', 0, 1);
INSERT INTO `sys_api` (`id`, `handle`, `title`, `path`, `api_type`, `action`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (70, 'go-admin/app/admin/apis.SysConfig.GetPage-fm', '系统-配置列表获取', '/admin-api/v1/config', '1', 'GET', '2023-05-08 16:51:32', '2023-05-09 11:25:52', 0, 1);
INSERT INTO `sys_api` (`id`, `handle`, `title`, `path`, `api_type`, `action`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (71, 'go-admin/app/admin/apis.SysConfig.Export-fm', '系统-业务配置导出', '/admin-api/v1/config/export', '1', 'GET', '2023-05-08 16:51:33', '2023-05-09 11:34:38', 0, 1);
INSERT INTO `sys_api` (`id`, `handle`, `title`, `path`, `api_type`, `action`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (72, 'go-admin/app/admin/apis.SysConfig.Get-fm', '系统-配置详情', '/admin-api/v1/config/:id', '1', 'GET', '2023-05-08 16:51:33', '2023-05-09 11:26:27', 0, 1);
INSERT INTO `sys_api` (`id`, `handle`, `title`, `path`, `api_type`, `action`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (73, 'go-admin/app/admin/apis.SysConfig.GetSysConfigByKey-fm', '系统-根据key获取系统配置详情', '/admin-api/v1/configKey/:configKey', '1', 'GET', '2023-05-08 16:51:33', '2023-05-09 09:44:11', 0, 1);
INSERT INTO `sys_api` (`id`, `handle`, `title`, `path`, `api_type`, `action`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (74, 'go-admin/app/admin/apis.SysUser.GenCaptcha-fm', '系统-获取验证码', '/admin-api/v1/captcha', '1', 'GET', '2023-05-08 16:51:34', '2023-05-09 09:43:43', 0, 1);
INSERT INTO `sys_api` (`id`, `handle`, `title`, `path`, `api_type`, `action`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (75, 'go-admin/app/admin/apis.SysMenu.GetPage-fm', '系统-菜单列表', '/admin-api/v1/menu', '1', 'GET', '2023-05-08 16:51:34', '2023-05-09 09:43:48', 0, 1);
INSERT INTO `sys_api` (`id`, `handle`, `title`, `path`, `api_type`, `action`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (76, 'go-admin/app/admin/apis.SysMenu.Get-fm', '系统-菜单详情', '/admin-api/v1/menu/:id', '1', 'GET', '2023-05-08 16:51:34', '2023-05-09 09:43:53', 0, 1);
INSERT INTO `sys_api` (`id`, `handle`, `title`, `path`, `api_type`, `action`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (77, 'go-admin/app/admin/apis.SysMenu.GetMenuRole-fm', '系统-获取当前登录账户的菜单', '/admin-api/v1/menurole', '1', 'GET', '2023-05-08 16:51:35', '2023-05-09 09:43:26', 0, 1);
INSERT INTO `sys_api` (`id`, `handle`, `title`, `path`, `api_type`, `action`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (78, 'go-admin/common/core/tools/transfer.Handler.func1', '系统-性能监控', '/admin-api/v1/metrics', '1', 'GET', '2023-05-08 16:51:35', '2023-05-09 10:52:46', 0, 1);
INSERT INTO `sys_api` (`id`, `handle`, `title`, `path`, `api_type`, `action`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (79, 'go-admin/app/admin/apis.SysRole.GetPage-fm', '系统-角色列表', '/admin-api/v1/role', '1', 'GET', '2023-05-08 16:51:35', '2023-05-09 09:43:37', 0, 1);
INSERT INTO `sys_api` (`id`, `handle`, `title`, `path`, `api_type`, `action`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (80, 'go-admin/app/admin/apis.SysRole.Get-fm', '', '/admin-api/v1/role/:id', '', 'GET', '2023-05-08 16:51:36', '2023-05-08 16:51:36', 0, 0);
INSERT INTO `sys_api` (`id`, `handle`, `title`, `path`, `api_type`, `action`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (81, 'go-admin/app/admin/router.registerMonitorRouter.func1', '系统-健康检测', '/admin-api/v1/health', '1', 'GET', '2023-05-08 16:51:36', '2023-05-09 09:42:43', 0, 1);
INSERT INTO `sys_api` (`id`, `handle`, `title`, `path`, `api_type`, `action`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (85, 'go-admin/app/admin/router.Ping', '系统-ping', '/info', '1', 'GET', '2023-05-08 16:51:37', '2023-05-09 09:43:11', 0, 1);
INSERT INTO `sys_api` (`id`, `handle`, `title`, `path`, `api_type`, `action`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (86, 'github.com/gin-gonic/gin.(*RouterGroup).createStaticHandler.func1', '系统-本地文件', '/files/*filepath', '1', 'GET', '2023-05-08 16:51:37', '2023-05-09 09:42:38', 0, 1);
INSERT INTO `sys_api` (`id`, `handle`, `title`, `path`, `api_type`, `action`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (87, 'github.com/gin-gonic/gin.(*RouterGroup).createStaticHandler.func1', '系统-静态文件获取', '/static/*filepath', '1', 'GET', '2023-05-08 16:51:38', '2023-05-09 09:42:16', 0, 1);
INSERT INTO `sys_api` (`id`, `handle`, `title`, `path`, `api_type`, `action`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (88, 'go-admin/app/admin/apis.SysDept.Insert-fm', '系统-部门新增', '/admin-api/v1/sys/dept', '1', 'POST', '2023-05-08 16:51:38', '2023-05-09 09:42:21', 0, 1);
INSERT INTO `sys_api` (`id`, `handle`, `title`, `path`, `api_type`, `action`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (89, 'go-admin/app/admin/apis.SysTables.Insert-fm', '系统-表数据新增', '/admin-api/v1/sys/table', '1', 'POST', '2023-05-08 16:51:38', '2023-05-09 10:55:36', 0, 1);
INSERT INTO `sys_api` (`id`, `handle`, `title`, `path`, `api_type`, `action`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (90, 'go-admin/app/admin/apis.SysPost.Insert-fm', '系统-岗位新增', '/admin-api/v1/sys/post', '1', 'POST', '2023-05-08 16:51:39', '2023-05-09 09:42:05', 0, 1);
INSERT INTO `sys_api` (`id`, `handle`, `title`, `path`, `api_type`, `action`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (91, 'go-admin/app/admin/apis.SysUser.Insert-fm', '系统-用户新增', '/admin-api/v1/sys-user', '1', 'POST', '2023-05-08 16:51:39', '2023-05-09 09:39:49', 0, 1);
INSERT INTO `sys_api` (`id`, `handle`, `title`, `path`, `api_type`, `action`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (92, 'go-admin/app/admin/apis.SysUser.InsetAvatar-fm', '系统-用户头像新增', '/admin-api/v1/sys-user/avatar', '1', 'POST', '2023-05-08 16:51:39', '2023-05-09 09:41:55', 0, 1);
INSERT INTO `sys_api` (`id`, `handle`, `title`, `path`, `api_type`, `action`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (93, 'go-admin/app/admin/apis.(*SysUser).LogOut-fm', '', '/admin-api/v1/sys-user/logout', '', 'POST', '2023-05-08 16:51:39', '2023-05-08 16:51:39', 0, 0);
INSERT INTO `sys_api` (`id`, `handle`, `title`, `path`, `api_type`, `action`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (94, 'go-admin/app/plugins/content/apis.ContentAnnouncement.Insert-fm', '插件-内容管理-公告新增', '/admin-api/v1/plugins/content/content-announcement', '2', 'POST', '2023-05-08 16:51:40', '2023-05-09 10:51:41', 0, 1);
INSERT INTO `sys_api` (`id`, `handle`, `title`, `path`, `api_type`, `action`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (95, 'go-admin/app/plugins/content/apis.ContentArticle.Insert-fm', '插件-内容管理-文章新增', '/admin-api/v1/plugins/content/content-article', '2', 'POST', '2023-05-08 16:51:40', '2023-05-09 10:48:03', 0, 1);
INSERT INTO `sys_api` (`id`, `handle`, `title`, `path`, `api_type`, `action`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (96, 'go-admin/app/plugins/content/apis.ContentCategory.Insert-fm', '插件-内容管理-文章分类新增', '/admin-api/v1/plugins/content/content-category', '2', 'POST', '2023-05-08 16:51:40', '2023-05-09 10:44:14', 0, 1);
INSERT INTO `sys_api` (`id`, `handle`, `title`, `path`, `api_type`, `action`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (97, 'go-admin/app/plugins/filemgr/apis.FilemgrApp.Insert-fm', '插件-文件管理-app数据新增', '/admin-api/v1/plugins/filemgr/filemgr-app', '2', 'POST', '2023-05-08 16:51:41', '2023-05-09 09:39:20', 0, 1);
INSERT INTO `sys_api` (`id`, `handle`, `title`, `path`, `api_type`, `action`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (98, 'go-admin/app/plugins/filemgr/apis.FilemgrApp.Upload-fm', '插件-文件管理-app文件上传', '/admin-api/v1/plugins/filemgr/filemgr-app/upload', '2', 'POST', '2023-05-08 16:51:41', '2023-05-09 09:39:11', 0, 1);
INSERT INTO `sys_api` (`id`, `handle`, `title`, `path`, `api_type`, `action`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (99, 'go-admin/app/app/user/apis.User.Insert-fm', '应用-用户新增', '/admin-api/v1/app/user/user', '3', 'POST', '2023-05-08 16:51:41', '2023-05-09 09:38:49', 0, 1);
INSERT INTO `sys_api` (`id`, `handle`, `title`, `path`, `api_type`, `action`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (100, 'go-admin/app/app/user/apis.UserCountryCode.Insert-fm', '应用-国家区号新增', '/admin-api/v1/app/user/user-country-code', '3', 'POST', '2023-05-08 16:51:42', '2023-05-09 10:31:10', 0, 1);
INSERT INTO `sys_api` (`id`, `handle`, `title`, `path`, `api_type`, `action`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (101, 'go-admin/app/app/user/apis.UserLevel.Insert-fm', '应用-用户等级新增', '/admin-api/v1/app/user/user-level', '3', 'POST', '2023-05-08 16:51:42', '2023-05-09 10:23:06', 0, 1);
INSERT INTO `sys_api` (`id`, `handle`, `title`, `path`, `api_type`, `action`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (102, 'go-admin/app/admin/apis.SysDictData.Insert-fm', '系统-字典新增', '/admin-api/v1/dict/data', '1', 'POST', '2023-05-08 16:51:42', '2023-05-09 11:10:05', 0, 1);
INSERT INTO `sys_api` (`id`, `handle`, `title`, `path`, `api_type`, `action`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (103, 'go-admin/app/admin/apis.SysDictType.Insert-fm', '系统-字典类型新增', '/admin-api/v1/dict/type', '1', 'POST', '2023-05-08 16:51:43', '2023-05-09 11:13:37', 0, 1);
INSERT INTO `sys_api` (`id`, `handle`, `title`, `path`, `api_type`, `action`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (104, 'go-admin/app/admin/apis.SysUser.Login-fm', '系统-用户登陆', '/admin-api/v1/login', '1', 'POST', '2023-05-08 16:51:43', '2023-05-09 09:38:16', 0, 1);
INSERT INTO `sys_api` (`id`, `handle`, `title`, `path`, `api_type`, `action`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (105, 'go-admin/app/admin/apis.SysConfig.Insert-fm', '系统-业务配置新增', '/admin-api/v1/config', '1', 'POST', '2023-05-08 16:51:43', '2023-05-09 11:26:03', 0, 1);
INSERT INTO `sys_api` (`id`, `handle`, `title`, `path`, `api_type`, `action`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (106, 'go-admin/app/admin/apis.SysMenu.Insert-fm', '系统-菜单新增', '/admin-api/v1/menu', '1', 'POST', '2023-05-08 16:51:44', '2023-05-09 09:37:50', 0, 1);
INSERT INTO `sys_api` (`id`, `handle`, `title`, `path`, `api_type`, `action`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (107, 'go-admin/app/admin/apis.SysRole.Insert-fm', '系统-角色新增', '/admin-api/v1/role', '1', 'POST', '2023-05-08 16:51:44', '2023-05-09 09:38:01', 0, 1);
INSERT INTO `sys_api` (`id`, `handle`, `title`, `path`, `api_type`, `action`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (108, 'go-admin/app/admin/apis.SysUser.Update-fm', '系统-用户更新', '/admin-api/v1/sys-user', '1', 'PUT', '2023-05-08 16:51:44', '2023-05-09 09:37:55', 0, 1);
INSERT INTO `sys_api` (`id`, `handle`, `title`, `path`, `api_type`, `action`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (109, 'go-admin/app/admin/apis.SysUser.UpdateSelfEmail-fm', '系统-用户邮箱更新', '/admin-api/v1/sys-user/updateSelfEmail', '1', 'PUT', '2023-05-08 16:51:45', '2023-05-09 09:37:40', 0, 1);
INSERT INTO `sys_api` (`id`, `handle`, `title`, `path`, `api_type`, `action`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (110, 'go-admin/app/admin/apis.SysUser.UpdateSelfPhone-fm', '系统-用户手机号更新', '/admin-api/v1/sys-user/updateSelfPhone', '1', 'PUT', '2023-05-08 16:51:45', '2023-05-09 09:37:35', 0, 1);
INSERT INTO `sys_api` (`id`, `handle`, `title`, `path`, `api_type`, `action`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (111, 'go-admin/app/admin/apis.SysUser.UpdateSelfNickName-fm', '系统-用户昵称更新', '/admin-api/v1/sys-user/updateSelfNickName', '1', 'PUT', '2023-05-08 16:51:45', '2023-05-09 09:37:45', 0, 1);
INSERT INTO `sys_api` (`id`, `handle`, `title`, `path`, `api_type`, `action`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (112, 'go-admin/app/admin/apis.SysUser.UpdatePwd-fm', '系统-更新密码', '/admin-api/v1/sys-user/pwd/set', '1', 'PUT', '2023-05-08 16:51:46', '2023-05-09 09:37:15', 0, 1);
INSERT INTO `sys_api` (`id`, `handle`, `title`, `path`, `api_type`, `action`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (113, 'go-admin/app/admin/apis.SysUser.ResetPwd-fm', '系统-重置密码', '/admin-api/v1/sys-user/pwd/reset', '1', 'PUT', '2023-05-08 16:51:46', '2023-05-09 09:37:10', 0, 1);
INSERT INTO `sys_api` (`id`, `handle`, `title`, `path`, `api_type`, `action`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (114, 'go-admin/app/admin/apis.SysUser.UpdateStatus-fm', '系统-用户状态更新', '/admin-api/v1/sys-user/status', '1', 'PUT', '2023-05-08 16:51:46', '2023-05-09 09:37:21', 0, 1);
INSERT INTO `sys_api` (`id`, `handle`, `title`, `path`, `api_type`, `action`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (115, 'go-admin/app/admin/apis.SysApi.Update-fm', '系统-接口信息更新', '/admin-api/v1/sys-api/:id', '1', 'PUT', '2023-05-08 16:51:46', '2023-05-09 09:37:30', 0, 1);
INSERT INTO `sys_api` (`id`, `handle`, `title`, `path`, `api_type`, `action`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (116, 'go-admin/app/admin/apis.SysDept.Update-fm', '系统-部门更新', '/admin-api/v1/sys/dept/:id', '1', 'PUT', '2023-05-08 16:51:47', '2023-05-09 09:36:55', 0, 1);
INSERT INTO `sys_api` (`id`, `handle`, `title`, `path`, `api_type`, `action`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (117, 'go-admin/app/admin/apis.SysTables.Update-fm', '系统-表信息更新', '/admin-api/v1/sys/table/:id', '1', 'PUT', '2023-05-08 16:51:47', '2023-05-09 11:14:37', 0, 1);
INSERT INTO `sys_api` (`id`, `handle`, `title`, `path`, `api_type`, `action`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (118, 'go-admin/app/admin/apis.SysPost.Update-fm', '系统-岗位更新', '/admin-api/v1/sys/post/:id', '1', 'PUT', '2023-05-08 16:51:47', '2023-05-09 09:37:05', 0, 1);
INSERT INTO `sys_api` (`id`, `handle`, `title`, `path`, `api_type`, `action`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (119, 'go-admin/app/app/user/apis.UserConf.Update-fm', '应用-用户配置更新', '/admin-api/v1/app/user/user-conf/:id', '3', 'PUT', '2023-05-08 16:51:48', '2023-05-09 10:25:16', 0, 1);
INSERT INTO `sys_api` (`id`, `handle`, `title`, `path`, `api_type`, `action`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (120, 'go-admin/app/app/user/apis.UserCountryCode.Update-fm', '应用-国家区号更新', '/admin-api/v1/app/user/user-country-code/:id', '3', 'PUT', '2023-05-08 16:51:48', '2023-05-09 10:31:24', 0, 1);
INSERT INTO `sys_api` (`id`, `handle`, `title`, `path`, `api_type`, `action`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (121, 'go-admin/app/app/user/apis.UserLevel.Update-fm', '应用-用户等级更新', '/admin-api/v1/app/user/user-level/:id', '3', 'PUT', '2023-05-08 16:51:48', '2023-05-09 10:23:27', 0, 1);
INSERT INTO `sys_api` (`id`, `handle`, `title`, `path`, `api_type`, `action`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (122, 'go-admin/app/app/user/apis.User.Update-fm', '用户更新', '/admin-api/v1/app/user/user/:id', '3', 'PUT', '2023-05-08 16:51:49', '2023-05-08 22:36:28', 0, 1);
INSERT INTO `sys_api` (`id`, `handle`, `title`, `path`, `api_type`, `action`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (123, 'go-admin/app/plugins/content/apis.ContentAnnouncement.Update-fm', '插件-内容管理-公告更新', '/admin-api/v1/plugins/content/content-announcement/:id', '2', 'PUT', '2023-05-08 16:51:49', '2023-05-09 10:51:56', 0, 1);
INSERT INTO `sys_api` (`id`, `handle`, `title`, `path`, `api_type`, `action`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (124, 'go-admin/app/plugins/content/apis.ContentArticle.Update-fm', '插件-内容管理-文章更新', '/admin-api/v1/plugins/content/content-article/:id', '2', 'PUT', '2023-05-08 16:51:49', '2023-05-09 10:48:27', 0, 1);
INSERT INTO `sys_api` (`id`, `handle`, `title`, `path`, `api_type`, `action`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (125, 'go-admin/app/plugins/content/apis.ContentCategory.Update-fm', '插件-内容管理-文章分类更新', '/admin-api/v1/plugins/content/content-category/:id', '2', 'PUT', '2023-05-08 16:51:50', '2023-05-09 10:45:07', 0, 1);
INSERT INTO `sys_api` (`id`, `handle`, `title`, `path`, `api_type`, `action`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (126, 'go-admin/app/plugins/filemgr/apis.FilemgrApp.Update-fm', '插件-文件管理-app信息更新', '/admin-api/v1/plugins/filemgr/filemgr-app/:id', '2', 'PUT', '2023-05-08 16:51:50', '2023-05-09 09:40:50', 0, 1);
INSERT INTO `sys_api` (`id`, `handle`, `title`, `path`, `api_type`, `action`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (127, 'go-admin/app/admin/apis.SysRole.Update-fm', '系统-角色更新', '/admin-api/v1/role/:id', '1', 'PUT', '2023-05-08 16:51:50', '2023-05-09 09:36:11', 0, 1);
INSERT INTO `sys_api` (`id`, `handle`, `title`, `path`, `api_type`, `action`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (128, 'go-admin/app/admin/apis.SysRole.Update2Status-fm', '系统-角色状态更新', '/admin-api/v1/role-status', '1', 'PUT', '2023-05-08 16:51:51', '2023-05-09 09:35:38', 0, 1);
INSERT INTO `sys_api` (`id`, `handle`, `title`, `path`, `api_type`, `action`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (129, 'go-admin/app/admin/apis.SysRole.Update2DataScope-fm', '系统-数据范围更新', '/admin-api/v1/roledatascope', '1', 'PUT', '2023-05-08 16:51:51', '2023-05-09 09:35:46', 0, 1);
INSERT INTO `sys_api` (`id`, `handle`, `title`, `path`, `api_type`, `action`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (130, 'go-admin/app/admin/apis.SysDictData.Update-fm', '系统-字典数据更新', '/admin-api/v1/dict/data/:id', '1', 'PUT', '2023-05-08 16:51:51', '2023-05-09 11:11:11', 0, 1);
INSERT INTO `sys_api` (`id`, `handle`, `title`, `path`, `api_type`, `action`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (131, 'go-admin/app/admin/apis.SysDictType.Update-fm', '系统-字典类型更新', '/admin-api/v1/dict/type/:id', '1', 'PUT', '2023-05-08 16:51:51', '2023-05-09 11:13:58', 0, 1);
INSERT INTO `sys_api` (`id`, `handle`, `title`, `path`, `api_type`, `action`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (132, 'go-admin/app/admin/apis.SysConfig.Update-fm', '系统-业务配置更新', '/admin-api/v1/config/:id', '1', 'PUT', '2023-05-08 16:51:52', '2023-05-09 11:26:27', 0, 1);
INSERT INTO `sys_api` (`id`, `handle`, `title`, `path`, `api_type`, `action`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (133, 'go-admin/app/admin/apis.SysMenu.Update-fm', '系统-菜单更新', '/admin-api/v1/menu/:id', '1', 'PUT', '2023-05-08 16:51:52', '2023-05-09 09:35:23', 0, 1);
INSERT INTO `sys_api` (`id`, `handle`, `title`, `path`, `api_type`, `action`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (134, 'go-admin/app/admin/apis.SysApi.Delete-fm', '', '/admin-api/v1/sys-api', '', 'DELETE', '2023-05-08 16:51:52', '2023-05-08 16:51:52', 0, 0);
INSERT INTO `sys_api` (`id`, `handle`, `title`, `path`, `api_type`, `action`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (135, 'go-admin/app/admin/apis.SysLoginLog.Delete-fm', '系统-登录日志删除', '/admin-api/v1/sys-login-log', '1', 'DELETE', '2023-05-08 16:51:53', '2023-05-09 11:03:46', 0, 1);
INSERT INTO `sys_api` (`id`, `handle`, `title`, `path`, `api_type`, `action`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (136, 'go-admin/app/admin/apis.SysOperLog.Delete-fm', '系统-操作日志删除', '/admin-api/v1/sys-oper-log', '1', 'DELETE', '2023-05-08 16:51:53', '2023-05-09 11:00:58', 0, 1);
INSERT INTO `sys_api` (`id`, `handle`, `title`, `path`, `api_type`, `action`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (137, 'go-admin/app/admin/apis.SysUser.Delete-fm', '系统-用户删除', '/admin-api/v1/sys-user', '1', 'DELETE', '2023-05-08 16:51:53', '2023-05-09 09:35:14', 0, 1);
INSERT INTO `sys_api` (`id`, `handle`, `title`, `path`, `api_type`, `action`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (138, 'go-admin/app/admin/apis.SysDept.Delete-fm', '系统-部门删除', '/admin-api/v1/sys/dept', '1', 'DELETE', '2023-05-08 16:51:54', '2023-05-09 09:34:52', 0, 1);
INSERT INTO `sys_api` (`id`, `handle`, `title`, `path`, `api_type`, `action`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (139, 'go-admin/app/admin/apis.SysTables.Delete-fm', '系统-表信息删除', '/admin-api/v1/sys/table', '1', 'DELETE', '2023-05-08 16:51:54', '2023-05-09 10:55:36', 0, 1);
INSERT INTO `sys_api` (`id`, `handle`, `title`, `path`, `api_type`, `action`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (140, 'go-admin/app/admin/apis.SysPost.Delete-fm', '系统-岗位删除', '/admin-api/v1/sys/post', '1', 'DELETE', '2023-05-08 16:51:54', '2023-05-09 09:34:57', 0, 1);
INSERT INTO `sys_api` (`id`, `handle`, `title`, `path`, `api_type`, `action`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (141, 'go-admin/app/plugins/content/apis.ContentAnnouncement.Delete-fm', '插件-内容管理-公告删除', '/admin-api/v1/plugins/content/content-announcement', '2', 'DELETE', '2023-05-08 16:51:55', '2023-05-09 10:52:06', 0, 1);
INSERT INTO `sys_api` (`id`, `handle`, `title`, `path`, `api_type`, `action`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (142, 'go-admin/app/plugins/content/apis.ContentArticle.Delete-fm', '插件-内容管理-文章删除', '/admin-api/v1/plugins/content/content-article', '2', 'DELETE', '2023-05-08 16:51:55', '2023-05-09 10:48:37', 0, 1);
INSERT INTO `sys_api` (`id`, `handle`, `title`, `path`, `api_type`, `action`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (143, 'go-admin/app/plugins/content/apis.ContentCategory.Delete-fm', '插件-内容管理-文章分类删除', '/admin-api/v1/plugins/content/content-category', '2', 'DELETE', '2023-05-08 16:51:55', '2023-05-09 10:47:01', 0, 1);
INSERT INTO `sys_api` (`id`, `handle`, `title`, `path`, `api_type`, `action`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (144, 'go-admin/app/plugins/filemgr/apis.FilemgrApp.Delete-fm', '插件-文件管理-app删除', '/admin-api/v1/plugins/filemgr/filemgr-app', '2', 'DELETE', '2023-05-08 16:51:56', '2023-05-09 09:41:38', 0, 1);
INSERT INTO `sys_api` (`id`, `handle`, `title`, `path`, `api_type`, `action`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (145, 'go-admin/app/app/user/apis.UserCountryCode.Delete-fm', '应用-国家区号删除', '/admin-api/v1/app/user/user-country-code', '3', 'DELETE', '2023-05-08 16:51:56', '2023-05-09 10:31:33', 0, 1);
INSERT INTO `sys_api` (`id`, `handle`, `title`, `path`, `api_type`, `action`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (146, 'go-admin/app/app/user/apis.UserLevel.Delete-fm', '应用-用户等级删除', '/admin-api/v1/app/user/user-level', '3', 'DELETE', '2023-05-08 16:51:56', '2023-05-09 10:23:44', 0, 1);
INSERT INTO `sys_api` (`id`, `handle`, `title`, `path`, `api_type`, `action`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (147, 'go-admin/app/admin/apis.SysDictData.Delete-fm', '系统-删除字典数据', '/admin-api/v1/dict/data', '1', 'DELETE', '2023-05-08 16:51:56', '2023-05-09 11:11:43', 0, 1);
INSERT INTO `sys_api` (`id`, `handle`, `title`, `path`, `api_type`, `action`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (148, 'go-admin/app/admin/apis.SysDictType.Delete-fm', '系统-删除字典类型', '/admin-api/v1/dict/type', '1', 'DELETE', '2023-05-08 16:51:57', '2023-05-09 11:15:32', 0, 1);
INSERT INTO `sys_api` (`id`, `handle`, `title`, `path`, `api_type`, `action`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (149, 'go-admin/app/admin/apis.SysConfig.Delete-fm', '系统-删除业务配置', '/admin-api/v1/config', '1', 'DELETE', '2023-05-08 16:51:57', '2023-05-09 11:26:39', 0, 1);
INSERT INTO `sys_api` (`id`, `handle`, `title`, `path`, `api_type`, `action`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (150, 'go-admin/app/admin/apis.SysMenu.Delete-fm', '系统-删除菜单', '/admin-api/v1/menu', '1', 'DELETE', '2023-05-08 16:51:57', '2023-05-09 09:28:26', 0, 1);
INSERT INTO `sys_api` (`id`, `handle`, `title`, `path`, `api_type`, `action`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (151, 'go-admin/app/admin/apis.SysRole.Delete-fm', '系统-删除角色', '/admin-api/v1/role', '1', 'DELETE', '2023-05-08 16:51:58', '2023-05-09 09:28:16', 0, 1);
INSERT INTO `sys_api` (`id`, `handle`, `title`, `path`, `api_type`, `action`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (152, 'github.com/gin-gonic/gin.(*RouterGroup).createStaticHandler.func1', '系统-静态资源', '/static/*filepath', '1', 'HEAD', '2023-05-08 16:51:58', '2023-05-09 09:28:11', 0, 1);
COMMIT;

-- ----------------------------
-- Table structure for sys_casbin_rule
-- ----------------------------
DROP TABLE IF EXISTS `sys_casbin_rule`;
CREATE TABLE `sys_casbin_rule` (
  `p_type` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `v0` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `v1` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `v2` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `v3` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `v4` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `v5` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  UNIQUE KEY `idx_sys_casbin_rule` (`p_type`,`v0`,`v1`,`v2`,`v3`,`v4`,`v5`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
-- Records of sys_casbin_rule
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for sys_config
-- ----------------------------
DROP TABLE IF EXISTS `sys_config`;
CREATE TABLE `sys_config` (
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
) ENGINE=InnoDB AUTO_INCREMENT=28 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;

-- ----------------------------
-- Records of sys_config
-- ----------------------------
BEGIN;
INSERT INTO `sys_config` (`id`, `config_name`, `config_key`, `config_value`, `config_type`, `is_frontend`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (1, '系统内置-皮肤样式', 'sys_index_skinName', 'skin-green', '1', '1', '主框架页-默认皮肤样式名称:蓝色 skin-blue、绿色 skin-green、紫色 skin-purple、红色 skin-red、黄色 skin-yellow', 1, 1, '2021-05-13 19:56:38', '2023-03-11 23:16:02');
INSERT INTO `sys_config` (`id`, `config_name`, `config_key`, `config_value`, `config_type`, `is_frontend`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (2, '系统内置-初始密码', 'sys_user_initPassword', '123456', '1', '1', '用户管理-账号初始密码:123456', 1, 1, '2021-05-13 19:56:38', '2023-03-11 23:16:10');
INSERT INTO `sys_config` (`id`, `config_name`, `config_key`, `config_value`, `config_type`, `is_frontend`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (3, '系统内置-侧栏主题', 'sys_index_sideTheme', 'theme-dark', '1', '1', '主框架页-侧边栏主题:深色主题theme-dark，浅色主题theme-light', 1, 1, '2021-05-13 19:56:38', '2023-03-11 23:16:06');
INSERT INTO `sys_config` (`id`, `config_name`, `config_key`, `config_value`, `config_type`, `is_frontend`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (4, '系统内置-系统名称', 'sys_app_name', 'go-admin后台管理系统', '1', '1', '', 1, 1, '2021-03-17 08:52:06', '2023-03-11 23:16:19');
INSERT INTO `sys_config` (`id`, `config_name`, `config_key`, `config_value`, `config_type`, `is_frontend`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (5, '系统内置-系统logo', 'sys_app_logo', 'http://www.bitxx.top/images/my_head-touch-icon-next.png', '1', '1', '', 1, 1, '2021-03-17 08:53:19', '2023-03-11 23:16:15');
INSERT INTO `sys_config` (`id`, `config_name`, `config_key`, `config_value`, `config_type`, `is_frontend`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (6, '系统内置-单次excel导出数据量', 'sys_max_export_size', '10000', '1', '1', '', 0, 1, '2021-07-28 16:53:48', '2023-03-11 23:15:56');
INSERT INTO `sys_config` (`id`, `config_name`, `config_key`, `config_value`, `config_type`, `is_frontend`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (7, '插件-文件管理-App OSS Bucket', 'plugin_filemgr_app_oss_bucket', '请自行配置', '2', '2', '', 0, 1, '2021-08-13 14:36:23', '2023-03-11 23:14:45');
INSERT INTO `sys_config` (`id`, `config_name`, `config_key`, `config_value`, `config_type`, `is_frontend`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (8, '插件-文件管理-App OSS AccessKeyId', 'plugin_filemgr_app_oss_access_key_id', '请自行配置', '2', '2', '', 0, 1, '2021-08-13 14:37:15', '2023-03-11 23:14:41');
INSERT INTO `sys_config` (`id`, `config_name`, `config_key`, `config_value`, `config_type`, `is_frontend`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (9, '插件-文件管理-App OSS AccessKeySecret', 'plugin_filemgr_app_oss_access_key_secret', '请自行配置', '2', '2', '', 0, 1, '2021-08-13 14:38:00', '2023-03-11 23:14:33');
INSERT INTO `sys_config` (`id`, `config_name`, `config_key`, `config_value`, `config_type`, `is_frontend`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (10, '插件-文件管理-App OSS Endpoint', 'plugin_filemgr_app_oss_endpoint', '请自行配置', '2', '2', '', 0, 1, '2021-08-13 14:38:50', '2023-03-11 23:14:28');
INSERT INTO `sys_config` (`id`, `config_name`, `config_key`, `config_value`, `config_type`, `is_frontend`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (11, '插件-文件管理-App OSS 根目录', 'plugin_filemgr_app_oss_root_path', 'testfile/', '2', '2', '', 0, 1, '2021-08-13 14:39:31', '2023-03-11 23:14:22');
INSERT INTO `sys_config` (`id`, `config_name`, `config_key`, `config_value`, `config_type`, `is_frontend`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (12, '中文通用短信模板', 'sys_sms_template_cn', '请自行配置', '2', '2', '', 0, 0, '2021-08-24 15:16:02', '2021-08-24 15:16:02');
INSERT INTO `sys_config` (`id`, `config_name`, `config_key`, `config_value`, `config_type`, `is_frontend`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (13, '英文通用短信模板', 'sys_sms_template_en', '请自行配置', '2', '2', '', 0, 0, '2021-08-24 15:16:33', '2022-03-16 14:28:02');
INSERT INTO `sys_config` (`id`, `config_name`, `config_key`, `config_value`, `config_type`, `is_frontend`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (14, '短信是否开启', 'sys_sms_open', '1', '2', '2', '0-开启  1-关闭', 0, 0, '2021-08-24 15:34:23', '2022-03-18 15:27:18');
INSERT INTO `sys_config` (`id`, `config_name`, `config_key`, `config_value`, `config_type`, `is_frontend`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (15, '短信签名', 'sys_sms_sign_name', '请自行配置', '2', '2', '', 0, 0, '2021-08-24 15:53:47', '2021-08-24 15:53:47');
INSERT INTO `sys_config` (`id`, `config_name`, `config_key`, `config_value`, `config_type`, `is_frontend`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (16, '短信secret', 'sys_sms_secret', '请自行配置', '2', '2', '', 0, 1, '2021-08-24 15:54:31', '2022-04-25 11:51:36');
INSERT INTO `sys_config` (`id`, `config_name`, `config_key`, `config_value`, `config_type`, `is_frontend`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (17, '短信Key', 'sys_sms_key', '请自行配置', '2', '2', '', 0, 0, '2021-08-24 15:57:20', '2021-08-24 15:57:20');
INSERT INTO `sys_config` (`id`, `config_name`, `config_key`, `config_value`, `config_type`, `is_frontend`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (18, '邮箱是否开启', 'sys_email_open', '1', '2', '2', '0-开启 1-关闭', 0, 0, '2021-09-01 17:23:28', '2022-03-23 14:54:19');
INSERT INTO `sys_config` (`id`, `config_name`, `config_key`, `config_value`, `config_type`, `is_frontend`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (19, '邮箱配置-服务商域名', 'msg_email_smtp_server', '请自行配置', '2', '2', '', 0, 0, '2021-09-01 17:58:28', '2021-09-01 17:58:28');
INSERT INTO `sys_config` (`id`, `config_name`, `config_key`, `config_value`, `config_type`, `is_frontend`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (20, '邮箱配置-发送人邮箱', 'msg_email_send_address', '请自行配置', '2', '2', '', 0, 0, '2021-09-01 17:59:42', '2021-09-01 17:59:42');
INSERT INTO `sys_config` (`id`, `config_name`, `config_key`, `config_value`, `config_type`, `is_frontend`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (21, '邮箱配置-授权码(密码)', 'msg_email_auth', '请自行配置', '2', '2', '', 0, 0, '2021-09-01 18:00:49', '2021-09-01 18:00:49');
INSERT INTO `sys_config` (`id`, `config_name`, `config_key`, `config_value`, `config_type`, `is_frontend`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (22, '邮箱配置-默认主题', 'msg_email_send_subject', '请自行配置', '2', '2', '', 0, 0, '2021-09-01 18:01:27', '2021-09-01 18:01:27');
INSERT INTO `sys_config` (`id`, `config_name`, `config_key`, `config_value`, `config_type`, `is_frontend`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (23, '邮箱配置-发送人姓名', 'msg_email_send_username', '请自行配置', '2', '2', '', 0, 0, '2021-09-01 18:02:02', '2021-09-01 18:02:02');
INSERT INTO `sys_config` (`id`, `config_name`, `config_key`, `config_value`, `config_type`, `is_frontend`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (24, '邮箱配置-默认内容', 'msg_email_send_content', '验证码5分钟有效,验证码提供他人可能导致账号被盗，请勿转发或泄露,您的验证码是:', '2', '2', '', 0, 0, '2021-09-01 18:02:37', '2021-09-01 18:02:37');
INSERT INTO `sys_config` (`id`, `config_name`, `config_key`, `config_value`, `config_type`, `is_frontend`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (25, '邮箱配置-端口号', 'msg_email_port', '465', '2', '2', '', 0, 0, '2021-09-01 18:03:05', '2021-09-01 18:03:05');
INSERT INTO `sys_config` (`id`, `config_name`, `config_key`, `config_value`, `config_type`, `is_frontend`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (26, '邮箱配置-是否https通信', 'msg_email_https', '0', '2', '2', '0-https 1-http', 0, 0, '2021-09-02 09:13:18', '2021-09-02 09:13:18');
INSERT INTO `sys_config` (`id`, `config_name`, `config_key`, `config_value`, `config_type`, `is_frontend`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (27, 'App-用户-默认头像', 'app_user_default_avatar', 'http://www.bitxx.top/images/my_head-touch-icon-next.png', '3', '2', '', 1, 1, '2023-03-10 18:07:03', '2023-03-10 18:07:03');
COMMIT;

-- ----------------------------
-- Table structure for sys_dept
-- ----------------------------
DROP TABLE IF EXISTS `sys_dept`;
CREATE TABLE `sys_dept` (
  `id` int NOT NULL AUTO_INCREMENT,
  `parent_id` int DEFAULT NULL,
  `dept_path` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL,
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
-- Records of sys_dept
-- ----------------------------
BEGIN;
INSERT INTO `sys_dept` (`id`, `parent_id`, `dept_path`, `dept_name`, `sort`, `leader`, `phone`, `email`, `status`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (1, 0, '0,', 'Admin', 0, 'admin', '', '', 1, 1, 1, '2021-05-13 19:56:38', '2022-05-14 11:20:25');
INSERT INTO `sys_dept` (`id`, `parent_id`, `dept_path`, `dept_name`, `sort`, `leader`, `phone`, `email`, `status`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (2, 1, '0,1,', '研发部', 1, 'admin', '', '', 1, 1, 1, '2021-05-13 19:56:38', '2023-03-04 13:17:45');
INSERT INTO `sys_dept` (`id`, `parent_id`, `dept_path`, `dept_name`, `sort`, `leader`, `phone`, `email`, `status`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (3, 1, '0,1,', '运维部', 1, 'admin', '', '', 1, 1, 1, '2021-05-13 19:56:38', '2023-03-04 13:17:45');
INSERT INTO `sys_dept` (`id`, `parent_id`, `dept_path`, `dept_name`, `sort`, `leader`, `phone`, `email`, `status`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (4, 1, '0,1,', '客服部', 0, 'admin', '', '', 1, 1, 1, '2021-05-13 19:56:38', '2022-05-14 11:20:50');
INSERT INTO `sys_dept` (`id`, `parent_id`, `dept_path`, `dept_name`, `sort`, `leader`, `phone`, `email`, `status`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (5, 1, '0,1,', '人力资源', 3, 'admin', '', '', 1, 1, 1, '2021-05-13 19:56:38', '2022-05-14 11:20:53');
INSERT INTO `sys_dept` (`id`, `parent_id`, `dept_path`, `dept_name`, `sort`, `leader`, `phone`, `email`, `status`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (6, 1, '0,1,', '市场', 10, 'admin', '', '', 1, 1, 1, '2021-12-02 10:13:38', '2021-12-02 10:13:38');
COMMIT;

-- ----------------------------
-- Table structure for sys_dict_data
-- ----------------------------
DROP TABLE IF EXISTS `sys_dict_data`;
CREATE TABLE `sys_dict_data` (
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
) ENGINE=InnoDB AUTO_INCREMENT=92 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;

-- ----------------------------
-- Records of sys_dict_data
-- ----------------------------
BEGIN;
INSERT INTO `sys_dict_data` (`id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `default`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (1, 0, '正常', '2', 'sys_normal_disable', '', '', '', '0', '', '系统正常', 1, 1, '2021-05-13 19:56:38', '2022-04-25 00:42:38');
INSERT INTO `sys_dict_data` (`id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `default`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (2, 0, '停用', '1', 'sys_normal_disable', '', '', '', '0', '', '系统停用', 1, 1, '2021-05-13 19:56:38', '2021-05-13 19:56:38');
INSERT INTO `sys_dict_data` (`id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `default`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (3, 0, '男', '1', 'sys_user_sex', '', '', '', '0', '', '性别男', 1, 1, '2021-05-13 19:56:38', '2021-05-13 19:56:38');
INSERT INTO `sys_dict_data` (`id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `default`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (4, 0, '女', '2', 'sys_user_sex', '', '', '', '0', '', '性别女', 1, 1, '2021-05-13 19:56:38', '2021-05-13 19:56:38');
INSERT INTO `sys_dict_data` (`id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `default`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (5, 0, '未知', '3', 'sys_user_sex', '', '', '', '0', '', '性别未知', 1, 1, '2021-05-13 19:56:38', '2023-03-05 12:03:33');
INSERT INTO `sys_dict_data` (`id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `default`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (6, 0, '显示', '2', 'sys_menu_show_hide', '', '', '', '0', '', '显示菜单', 1, 1, '2021-05-13 19:56:38', '2021-05-13 19:56:38');
INSERT INTO `sys_dict_data` (`id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `default`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (7, 0, '隐藏', '1', 'sys_menu_show_hide', '', '', '', '0', '', '隐藏菜单', 1, 1, '2021-05-13 19:56:38', '2021-05-13 19:56:38');
INSERT INTO `sys_dict_data` (`id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `default`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (8, 0, '是', '1', 'sys_yes_no', '', '', '', '0', '', '系统默认是', 1, 1, '2021-05-13 19:56:38', '2021-05-13 19:56:38');
INSERT INTO `sys_dict_data` (`id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `default`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (9, 0, '否', '2', 'sys_yes_no', '', '', '', '0', '', '系统默认否', 1, 1, '2021-05-13 19:56:38', '2021-05-13 19:56:38');
INSERT INTO `sys_dict_data` (`id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `default`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (10, 0, '通知', '1', 'sys_notice_type', '', '', '', '0', '', '通知', 1, 1, '2021-05-13 19:56:38', '2021-05-13 19:56:38');
INSERT INTO `sys_dict_data` (`id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `default`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (11, 0, '公告', '2', 'sys_notice_type', '', '', '', '0', '', '公告', 1, 1, '2021-05-13 19:56:38', '2021-05-13 19:56:38');
INSERT INTO `sys_dict_data` (`id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `default`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (12, 0, '正常', '2', 'sys_common_status', '', '', '', '0', '', '正常状态', 1, 1, '2021-05-13 19:56:38', '2021-05-13 19:56:38');
INSERT INTO `sys_dict_data` (`id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `default`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (13, 0, '关闭', '1', 'sys_common_status', '', '', '', '0', '', '关闭状态', 1, 1, '2021-05-13 19:56:38', '2021-05-13 19:56:38');
INSERT INTO `sys_dict_data` (`id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `default`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (14, 0, '新增', '1', 'sys_oper_type', '', '', '', '0', '', '新增操作', 1, 1, '2021-05-13 19:56:38', '2021-05-13 19:56:38');
INSERT INTO `sys_dict_data` (`id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `default`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (15, 0, '修改', '2', 'sys_oper_type', '', '', '', '0', '', '修改操作', 1, 1, '2021-05-13 19:56:38', '2021-05-13 19:56:38');
INSERT INTO `sys_dict_data` (`id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `default`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (16, 0, '删除', '3', 'sys_oper_type', '', '', '', '0', '', '删除操作', 1, 1, '2021-05-13 19:56:38', '2021-05-13 19:56:38');
INSERT INTO `sys_dict_data` (`id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `default`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (17, 0, '授权', '4', 'sys_oper_type', '', '', '', '0', '', '授权操作', 1, 1, '2021-05-13 19:56:38', '2021-05-13 19:56:38');
INSERT INTO `sys_dict_data` (`id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `default`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (18, 0, '导出', '5', 'sys_oper_type', '', '', '', '0', '', '导出操作', 1, 1, '2021-05-13 19:56:38', '2021-05-13 19:56:38');
INSERT INTO `sys_dict_data` (`id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `default`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (19, 0, '导入', '6', 'sys_oper_type', '', '', '', '0', '', '导入操作', 1, 1, '2021-05-13 19:56:38', '2021-05-13 19:56:38');
INSERT INTO `sys_dict_data` (`id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `default`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (20, 0, '强退', '7', 'sys_oper_type', '', '', '', '0', '', '强退操作', 1, 1, '2021-05-13 19:56:38', '2021-05-13 19:56:38');
INSERT INTO `sys_dict_data` (`id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `default`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (21, 0, '生成代码', '8', 'sys_oper_type', '', '', '', '0', '', '生成操作', 1, 1, '2021-05-13 19:56:38', '2021-05-13 19:56:38');
INSERT INTO `sys_dict_data` (`id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `default`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (22, 0, '清空数据', '9', 'sys_oper_type', '', '', '', '0', '', '清空操作', 1, 1, '2021-05-13 19:56:38', '2021-05-13 19:56:38');
INSERT INTO `sys_dict_data` (`id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `default`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (23, 0, '成功', '1', 'sys_notice_status', '', '', '', '0', '', '成功状态', 1, 1, '2021-05-13 19:56:38', '2021-05-13 19:56:38');
INSERT INTO `sys_dict_data` (`id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `default`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (24, 0, '失败', '2', 'sys_notice_status', '', '', '', '0', '', '失败状态', 1, 1, '2021-05-13 19:56:38', '2021-05-13 19:56:38');
INSERT INTO `sys_dict_data` (`id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `default`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (25, 0, '登录', '10', 'sys_oper_type', '', '', '', '0', '', '登录操作', 1, 1, '2021-05-13 19:56:38', '2021-05-13 19:56:38');
INSERT INTO `sys_dict_data` (`id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `default`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (26, 0, '退出', '11', 'sys_oper_type', '', '', '', '0', '', '', 1, 1, '2021-05-13 19:56:38', '2021-05-13 19:56:38');
INSERT INTO `sys_dict_data` (`id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `default`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (27, 0, '获取验证码', '12', 'sys_oper_type', '', '', '', '0', '', '获取验证码', 1, 1, '2021-05-13 19:56:38', '2021-05-13 19:56:38');
INSERT INTO `sys_dict_data` (`id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `default`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (28, 0, '正常', '1', 'sys_status', '', '', '', '0', '', '', 0, 0, '2021-07-09 11:40:01', '2021-07-09 11:40:01');
INSERT INTO `sys_dict_data` (`id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `default`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (29, 0, '停用', '2', 'sys_status', '', '', '', '0', '', '', 0, 0, '2021-07-09 11:40:14', '2021-07-09 11:40:14');
INSERT INTO `sys_dict_data` (`id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `default`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (30, 0, '安卓', '1', 'plugin_filemgr_app_platform', '', '', '', '2', '', '', 0, 0, '2021-08-13 13:35:39', '2021-08-13 13:35:39');
INSERT INTO `sys_dict_data` (`id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `default`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (31, 0, 'IOS', '2', 'plugin_filemgr_app_platform', '', '', '', '2', '', '', 0, 0, '2021-08-13 13:35:51', '2021-08-13 13:35:51');
INSERT INTO `sys_dict_data` (`id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `default`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (32, 0, '类型1', '1', 'plugin_filemgr_app_type', '', '', '', '2', '', '', 0, 0, '2021-08-13 13:37:07', '2021-08-13 13:37:07');
INSERT INTO `sys_dict_data` (`id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `default`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (33, 0, '类型2', '2', 'plugin_filemgr_app_type', '', '', '', '2', '', '', 0, 0, '2021-08-13 13:37:19', '2021-08-13 13:37:19');
INSERT INTO `sys_dict_data` (`id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `default`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (34, 0, '类型3', '3', 'plugin_filemgr_app_type', '', '', '', '2', '', '', 0, 0, '2021-08-13 13:37:39', '2021-08-13 13:37:39');
INSERT INTO `sys_dict_data` (`id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `default`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (35, 0, '本地', '1', 'plugin_filemgr_app_download_type', '', '', '', '2', '', '', 0, 0, '2021-08-13 14:02:44', '2021-08-13 14:02:44');
INSERT INTO `sys_dict_data` (`id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `default`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (36, 0, '外链', '2', 'plugin_filemgr_app_download_type', '', '', '', '2', '', '', 0, 0, '2021-08-13 14:02:44', '2021-08-13 14:02:44');
INSERT INTO `sys_dict_data` (`id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `default`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (37, 0, 'OSS', '3', 'plugin_filemgr_app_download_type', '', '', '', '2', '', '', 0, 0, '2021-08-13 14:02:33', '2021-08-13 14:02:33');
INSERT INTO `sys_dict_data` (`id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `default`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (38, 0, '已发布', '2', 'plugin_filemgr_app_publish_status', '', '', '', '2', '', '', 0, 0, '2021-12-09 12:42:47', '2021-12-09 12:42:47');
INSERT INTO `sys_dict_data` (`id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `default`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (39, 0, '待发布', '1', 'plugin_filemgr_app_publish_status', '', '', '', '2', '', '', 0, 0, '2021-12-09 12:42:54', '2021-12-09 12:42:54');
INSERT INTO `sys_dict_data` (`id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `default`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (40, 0, '插件', '2', 'sys_api_type', '', '', '', '0', '', '', 1, 1, '2022-04-25 23:58:24', '2023-03-01 21:45:53');
INSERT INTO `sys_dict_data` (`id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `default`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (41, 0, '系统', '1', 'sys_api_type', '', '', '', '0', '', '', 1, 1, '2022-04-25 23:58:41', '2023-03-01 21:45:41');
INSERT INTO `sys_dict_data` (`id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `default`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (42, 0, 'GET', 'GET', 'sys_api_action', '', '', '', '0', '', '', 1, 1, '2022-04-26 00:03:26', '2022-04-26 00:03:26');
INSERT INTO `sys_dict_data` (`id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `default`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (43, 0, 'POST', 'POST', 'sys_api_action', '', '', '', '0', '', '', 1, 1, '2022-04-26 00:03:40', '2022-04-26 00:03:40');
INSERT INTO `sys_dict_data` (`id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `default`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (44, 0, 'DELETE', 'DELETE', 'sys_api_action', '', '', '', '0', '', '', 1, 1, '2022-04-26 00:03:49', '2022-04-26 00:03:49');
INSERT INTO `sys_dict_data` (`id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `default`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (45, 0, 'PUT', 'PUT', 'sys_api_action', '', '', '', '0', '', '', 1, 1, '2022-04-26 00:04:06', '2022-04-26 00:04:06');
INSERT INTO `sys_dict_data` (`id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `default`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (46, 0, 'HEAD', 'HEAD', 'sys_api_action', '', '', '', '0', '', '', 1, 1, '2022-04-26 00:07:02', '2022-04-26 00:07:02');
INSERT INTO `sys_dict_data` (`id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `default`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (47, 0, '系统内置', '1', 'sys_config_type', '', '', '', '1', '', '', 1, 1, '2023-03-01 11:05:23', '2023-03-01 11:05:23');
INSERT INTO `sys_dict_data` (`id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `default`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (48, 0, '插件', '2', 'sys_config_type', '', '', '', '1', '', '', 1, 1, '2023-03-01 11:05:32', '2023-03-01 11:05:32');
INSERT INTO `sys_dict_data` (`id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `default`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (49, 0, '应用', '3', 'sys_config_type', '', '', '', '1', '', '', 1, 1, '2023-03-01 11:05:42', '2023-03-01 11:05:42');
INSERT INTO `sys_dict_data` (`id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `default`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (50, 0, '展示', '1', 'sys_config_is_frontend', '', '', '', '1', '', '', 1, 1, '2023-03-01 11:07:49', '2023-03-01 11:07:49');
INSERT INTO `sys_dict_data` (`id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `default`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (51, 0, '隐藏', '2', 'sys_config_is_frontend', '', '', '', '1', '', '', 1, 1, '2023-03-01 11:07:56', '2023-03-01 11:07:56');
INSERT INTO `sys_dict_data` (`id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `default`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (52, 0, '登录', '1', 'sys_loginlog_status', '', '', '', '1', '', '', 1, 1, '2023-03-01 14:43:04', '2023-03-01 14:43:04');
INSERT INTO `sys_dict_data` (`id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `default`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (53, 0, '退出', '2', 'sys_loginlog_status', '', '', '', '1', '', '', 1, 1, '2023-03-01 14:43:10', '2023-03-01 14:43:10');
INSERT INTO `sys_dict_data` (`id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `default`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (54, 0, '应用', '3', 'sys_api_type', '', '', '', '1', '', '', 1, 1, '2023-03-01 21:46:01', '2023-03-01 21:46:01');
INSERT INTO `sys_dict_data` (`id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `default`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (55, 0, '全部数据权限', '1', 'sys_role_data_scope', '', '', '', '1', '', '', 1, 1, '2023-03-04 13:29:36', '2023-03-04 13:29:36');
INSERT INTO `sys_dict_data` (`id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `default`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (56, 0, '自定数据权限', '2', 'sys_role_data_scope', '', '', '', '1', '', '', 1, 1, '2023-03-04 13:29:43', '2023-03-04 13:29:43');
INSERT INTO `sys_dict_data` (`id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `default`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (57, 0, '本部门数据权限', '3', 'sys_role_data_scope', '', '', '', '1', '', '', 1, 1, '2023-03-04 13:29:49', '2023-03-04 13:29:49');
INSERT INTO `sys_dict_data` (`id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `default`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (58, 0, '本部门及以下数据权限', '4', 'sys_role_data_scope', '', '', '', '1', '', '', 1, 1, '2023-03-04 13:29:56', '2023-03-04 13:29:56');
INSERT INTO `sys_dict_data` (`id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `default`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (59, 0, '仅本人数据权限', '5', 'sys_role_data_scope', '', '', '', '1', '', '', 1, 1, '2023-03-04 13:30:04', '2023-03-04 13:30:04');
INSERT INTO `sys_dict_data` (`id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `default`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (60, 0, 'int64', 'int64', 'sys_gen_go_type', '', '', '', '1', '', '', 1, 1, '2023-03-07 10:08:26', '2023-03-07 10:08:26');
INSERT INTO `sys_dict_data` (`id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `default`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (61, 0, 'int', 'int', 'sys_gen_go_type', '', '', '', '1', '', '', 1, 1, '2023-03-07 10:12:42', '2023-03-07 10:12:42');
INSERT INTO `sys_dict_data` (`id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `default`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (62, 0, 'string', 'string', 'sys_gen_go_type', '', '', '', '1', '', '', 1, 1, '2023-03-07 10:13:05', '2023-03-07 10:13:05');
INSERT INTO `sys_dict_data` (`id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `default`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (63, 0, 'decimal', 'decimal.Decimal', 'sys_gen_go_type', '', '', '', '1', '', '', 1, 1, '2023-03-07 10:13:16', '2023-03-07 10:13:29');
INSERT INTO `sys_dict_data` (`id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `default`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (64, 0, 'time', '*time.Time', 'sys_gen_go_type', '', '', '', '1', '', '', 1, 1, '2023-03-07 10:13:43', '2023-03-07 10:13:43');
INSERT INTO `sys_dict_data` (`id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `default`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (65, 0, '=', 'EQ', 'sys_gen_query_type', '', '', '', '1', '', '', 1, 1, '2023-03-07 10:20:53', '2023-03-07 10:20:53');
INSERT INTO `sys_dict_data` (`id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `default`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (66, 0, '!=', 'NE', 'sys_gen_query_type', '', '', '', '1', '', '', 1, 1, '2023-03-07 10:21:06', '2023-03-07 10:21:06');
INSERT INTO `sys_dict_data` (`id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `default`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (67, 0, '>', 'GT', 'sys_gen_query_type', '', '', '', '1', '', '', 1, 1, '2023-03-07 10:21:20', '2023-03-07 10:21:20');
INSERT INTO `sys_dict_data` (`id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `default`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (68, 0, '>=', 'GTE', 'sys_gen_query_type', '', '', '', '1', '', '', 1, 1, '2023-03-07 10:21:33', '2023-03-07 10:21:33');
INSERT INTO `sys_dict_data` (`id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `default`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (69, 0, '<', 'LT', 'sys_gen_query_type', '', '', '', '1', '', '', 1, 1, '2023-03-07 10:21:45', '2023-03-07 10:21:45');
INSERT INTO `sys_dict_data` (`id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `default`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (70, 0, '<=', 'LTE', 'sys_gen_query_type', '', '', '', '1', '', '', 1, 1, '2023-03-07 10:21:57', '2023-03-07 10:21:57');
INSERT INTO `sys_dict_data` (`id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `default`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (71, 0, 'LIKE', 'LIKE', 'sys_gen_query_type', '', '', '', '1', '', '', 1, 1, '2023-03-07 10:22:08', '2023-03-07 10:22:08');
INSERT INTO `sys_dict_data` (`id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `default`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (72, 0, '文本框', 'input', 'sys_gen_html_type', '', '', '', '1', '', '', 1, 1, '2023-03-07 10:23:39', '2023-03-07 10:23:39');
INSERT INTO `sys_dict_data` (`id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `default`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (73, 0, '下拉框', 'select', 'sys_gen_html_type', '', '', '', '1', '', '', 1, 1, '2023-03-07 10:23:49', '2023-03-07 10:23:49');
INSERT INTO `sys_dict_data` (`id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `default`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (74, 0, '单选框', 'radio', 'sys_gen_html_type', '', '', '', '1', '', '', 1, 1, '2023-03-07 10:23:59', '2023-03-07 10:23:59');
INSERT INTO `sys_dict_data` (`id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `default`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (75, 0, '文本域', 'textarea', 'sys_gen_html_type', '', '', '', '1', '', '', 1, 1, '2023-03-07 10:24:08', '2023-03-07 10:24:08');
INSERT INTO `sys_dict_data` (`id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `default`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (76, 0, '目录', '1', 'sys_menu_type', '', '', '', '1', '', '', 1, 1, '2023-03-08 10:42:00', '2023-03-08 10:42:14');
INSERT INTO `sys_dict_data` (`id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `default`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (77, 0, '菜单', '2', 'sys_menu_type', '', '', '', '1', '', '', 1, 1, '2023-03-08 10:42:10', '2023-03-08 10:42:10');
INSERT INTO `sys_dict_data` (`id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `default`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (78, 0, '按钮', '3', 'sys_menu_type', '', '', '', '1', '', '', 1, 1, '2023-03-08 10:42:22', '2023-03-08 10:42:22');
INSERT INTO `sys_dict_data` (`id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `default`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (79, 0, '类型1', '1', 'app_user_level_type', '', '', '', '1', '', '', 1, 1, '2023-03-08 11:55:57', '2023-03-08 11:55:57');
INSERT INTO `sys_dict_data` (`id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `default`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (80, 0, '类型2', '2', 'app_user_level_type', '', '', '', '1', '', '', 1, 1, '2023-03-08 11:56:02', '2023-03-08 11:56:02');
INSERT INTO `sys_dict_data` (`id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `default`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (81, 0, '数字文本框', 'numInput', 'sys_gen_html_type', '', '', '', '1', '', '', 1, 1, '2023-03-09 20:12:33', '2023-03-09 20:12:33');
INSERT INTO `sys_dict_data` (`id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `default`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (82, 0, 'CNY', '1', 'app_money_type', '', '', '', '1', '', '', 1, 1, '2023-03-09 20:24:26', '2023-03-09 20:24:26');
INSERT INTO `sys_dict_data` (`id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `default`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (83, 0, '类型1', '1', 'app_account_change_type', '', '', '', '1', '', '', 1, 1, '2023-03-09 20:27:45', '2023-03-09 20:27:45');
INSERT INTO `sys_dict_data` (`id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `default`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (84, 0, '允许用户登录', '1', 'app_user_action_type', '', '', '', '1', '', '', 1, 1, '2023-03-11 14:08:01', '2023-03-11 14:08:01');
INSERT INTO `sys_dict_data` (`id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `default`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (85, 0, '禁止用户登录', '2', 'app_user_action_type', '', '', '', '1', '', '', 1, 1, '2023-03-11 14:08:10', '2023-03-11 14:08:10');
INSERT INTO `sys_dict_data` (`id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `default`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (86, 0, '后台用户', '1', 'app_user_by_type', '', '', '', '1', '', '', 1, 1, '2023-03-11 14:14:41', '2023-03-11 14:14:41');
INSERT INTO `sys_dict_data` (`id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `default`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (87, 0, '前台用户', '2', 'app_user_by_type', '', '', '', '1', '', '', 1, 1, '2023-03-11 14:14:59', '2023-03-11 14:14:59');
INSERT INTO `sys_dict_data` (`id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `default`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (88, 0, '发送成功', '1', 'plugin_msg_sendstatus', '', '', '', '1', '', '', 1, 1, '2023-09-26 10:42:22', '2023-09-26 10:42:22');
INSERT INTO `sys_dict_data` (`id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `default`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (89, 0, '发送失败', '2', 'plugin_msg_sendstatus', '', '', '', '1', '', '', 1, 1, '2023-09-26 10:42:31', '2023-09-26 10:42:31');
INSERT INTO `sys_dict_data` (`id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `default`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (90, 0, '邮箱', '1', 'plugin_msg_code_type', '', '', '', '1', '', '', 1, 1, '2023-09-26 10:42:58', '2023-09-26 10:42:58');
INSERT INTO `sys_dict_data` (`id`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `default`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (91, 0, '短信', '2', 'plugin_msg_code_type', '', '', '', '1', '', '', 1, 1, '2023-09-26 10:43:04', '2023-09-26 10:43:04');
COMMIT;

-- ----------------------------
-- Table structure for sys_dict_type
-- ----------------------------
DROP TABLE IF EXISTS `sys_dict_type`;
CREATE TABLE `sys_dict_type` (
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
) ENGINE=InnoDB AUTO_INCREMENT=33 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;

-- ----------------------------
-- Records of sys_dict_type
-- ----------------------------
BEGIN;
INSERT INTO `sys_dict_type` (`id`, `dict_name`, `dict_type`, `status`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (1, '系统-开关', 'sys_normal_disable', '0', '系统开关列表', 1, 1, '2021-05-13 19:56:38', '2023-03-11 23:20:35');
INSERT INTO `sys_dict_type` (`id`, `dict_name`, `dict_type`, `status`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (2, '系统-用户性别', 'sys_user_sex', '0', '用户性别列表', 1, 1, '2021-05-13 19:56:38', '2023-03-11 23:21:06');
INSERT INTO `sys_dict_type` (`id`, `dict_name`, `dict_type`, `status`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (3, '系统-菜单状态', 'sys_menu_show_hide', '0', '菜单状态列表', 1, 1, '2021-05-13 19:56:38', '2023-03-11 23:21:02');
INSERT INTO `sys_dict_type` (`id`, `dict_name`, `dict_type`, `status`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (4, '系统-是否', 'sys_yes_no', '0', '系统是否列表', 1, 1, '2021-05-13 19:56:38', '2023-03-11 23:20:58');
INSERT INTO `sys_dict_type` (`id`, `dict_name`, `dict_type`, `status`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (7, '系统-通知类型', 'sys_notice_type', '0', '通知类型列表', 1, 1, '2021-05-13 19:56:38', '2023-03-11 23:20:53');
INSERT INTO `sys_dict_type` (`id`, `dict_name`, `dict_type`, `status`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (8, '系统-状态', 'sys_common_status', '0', '登录状态列表', 1, 1, '2021-05-13 19:56:38', '2023-03-11 23:20:49');
INSERT INTO `sys_dict_type` (`id`, `dict_name`, `dict_type`, `status`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (9, '系统-操作类型', 'sys_oper_type', '0', '操作类型列表', 1, 1, '2021-05-13 19:56:38', '2023-03-11 23:20:42');
INSERT INTO `sys_dict_type` (`id`, `dict_name`, `dict_type`, `status`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (10, '系统-通知状态', 'sys_notice_status', '0', '通知状态列表', 1, 1, '2021-05-13 19:56:38', '2023-03-11 23:20:39');
INSERT INTO `sys_dict_type` (`id`, `dict_name`, `dict_type`, `status`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (11, '系统-基本状态', 'sys_status', '0', '基本通用状态', 1, 1, '2021-07-09 11:39:21', '2023-03-11 23:21:23');
INSERT INTO `sys_dict_type` (`id`, `dict_name`, `dict_type`, `status`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (12, '插件-文件管理-App发布状态', 'plugin_filemgr_publish_status', '2', '', 1, 1, '2021-12-09 12:42:31', '2023-03-11 23:20:01');
INSERT INTO `sys_dict_type` (`id`, `dict_name`, `dict_type`, `status`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (13, '插件-文件管理-App系统平台', 'plugin_filemgr_app_platform', '0', 'App系统平台', 1, 1, '2021-08-13 13:36:40', '2023-03-11 23:20:17');
INSERT INTO `sys_dict_type` (`id`, `dict_name`, `dict_type`, `status`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (14, '插件-文件管理-App类型', 'plugin_filemgr_app_type', '0', 'app属性', 1, 1, '2021-08-13 13:36:40', '2023-03-11 23:20:13');
INSERT INTO `sys_dict_type` (`id`, `dict_name`, `dict_type`, `status`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (15, '插件-文件管理-App下载类型', 'plugin_filemgr_app_download_type', '0', '', 1, 1, '2021-08-13 14:02:03', '2023-03-11 23:20:06');
INSERT INTO `sys_dict_type` (`id`, `dict_name`, `dict_type`, `status`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (16, '系统-接口-类型', 'sys_api_type', '0', '系统', 1, 1, '2022-04-25 23:57:17', '2023-03-01 21:56:34');
INSERT INTO `sys_dict_type` (`id`, `dict_name`, `dict_type`, `status`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (17, '系统-接口-请求方法', 'sys_api_action', '0', '', 1, 1, '2022-04-26 00:03:11', '2023-03-01 21:56:41');
INSERT INTO `sys_dict_type` (`id`, `dict_name`, `dict_type`, `status`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (18, '系统-配置-类型', 'sys_config_type', '1', '1-内置 2-插件 3-应用', 1, 1, '2023-03-01 11:04:56', '2023-03-01 11:08:27');
INSERT INTO `sys_dict_type` (`id`, `dict_name`, `dict_type`, `status`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (19, '系统-配置-是否前台展示', 'sys_config_is_frontend', '1', '1-展示 2-隐藏', 1, 1, '2023-03-01 11:06:28', '2023-03-01 11:08:07');
INSERT INTO `sys_dict_type` (`id`, `dict_name`, `dict_type`, `status`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (20, '系统-登录日志-日志状态', 'sys_loginlog_status', '1', '1-登录 2-退出', 1, 1, '2023-03-01 14:42:56', '2023-03-01 14:42:56');
INSERT INTO `sys_dict_type` (`id`, `dict_name`, `dict_type`, `status`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (21, '系统-角色-数据范围', 'sys_role_data_scope', '1', '1-全部数据权限 2- 自定义数据权限 3-本部门数据权限 4-本部门及以下数据权限 5-仅本人数据权限', 1, 1, '2023-03-04 13:29:21', '2023-03-04 13:29:21');
INSERT INTO `sys_dict_type` (`id`, `dict_name`, `dict_type`, `status`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (22, '系统-模板-go类型', 'sys_gen_go_type', '1', '', 1, 1, '2023-03-07 10:08:07', '2023-03-07 10:08:07');
INSERT INTO `sys_dict_type` (`id`, `dict_name`, `dict_type`, `status`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (23, '系统-模板-查询类型', 'sys_gen_query_type', '1', '', 1, 1, '2023-03-07 10:20:19', '2023-03-07 10:20:19');
INSERT INTO `sys_dict_type` (`id`, `dict_name`, `dict_type`, `status`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (24, '系统-模板-显示类型', 'sys_gen_html_type', '1', '', 1, 1, '2023-03-07 10:23:23', '2023-03-07 10:23:23');
INSERT INTO `sys_dict_type` (`id`, `dict_name`, `dict_type`, `status`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (25, '系统-菜单-类型', 'sys_menu_type', '1', '', 1, 1, '2023-03-08 10:33:32', '2023-03-08 10:33:32');
INSERT INTO `sys_dict_type` (`id`, `dict_name`, `dict_type`, `status`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (26, 'App-用户-等级', 'app_user_level_type', '1', '', 1, 1, '2023-03-08 11:44:48', '2023-03-08 11:44:48');
INSERT INTO `sys_dict_type` (`id`, `dict_name`, `dict_type`, `status`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (27, 'App-用户-资产-资金类型', 'app_money_type', '1', '1-CNY', 1, 1, '2023-03-09 20:24:17', '2023-03-11 14:06:46');
INSERT INTO `sys_dict_type` (`id`, `dict_name`, `dict_type`, `status`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (28, 'App-用户-资产-账变类型', 'app_account_change_type', '1', '1-类型1', 1, 1, '2023-03-09 20:27:33', '2023-03-11 14:06:38');
INSERT INTO `sys_dict_type` (`id`, `dict_name`, `dict_type`, `status`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (29, 'App-用户-行为类型', 'app_user_action_type', '1', '', 1, 1, '2023-03-11 14:06:29', '2023-03-11 14:06:29');
INSERT INTO `sys_dict_type` (`id`, `dict_name`, `dict_type`, `status`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (30, 'App-用户-用户更新类型', 'app_user_by_type', '1', '', 1, 1, '2023-03-11 14:14:06', '2023-03-11 14:14:27');
INSERT INTO `sys_dict_type` (`id`, `dict_name`, `dict_type`, `status`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (31, '插件-消息-验证码类型', 'plugin_msg_code_type', '1', '1-邮箱 2-短信', 1, 1, '2023-03-12 12:12:30', '2023-03-12 12:15:20');
INSERT INTO `sys_dict_type` (`id`, `dict_name`, `dict_type`, `status`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (32, '插件-消息-验证码发送状态', 'plugin_msg_sendstatus', '1', '', 1, 1, '2023-03-12 12:14:56', '2023-03-12 13:23:37');
COMMIT;

-- ----------------------------
-- Table structure for sys_gen_column
-- ----------------------------
DROP TABLE IF EXISTS `sys_gen_column`;
CREATE TABLE `sys_gen_column` (
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
  `is_edit` char(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '该值可否二次编辑',
  `is_must` char(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '是否必须填写值 1-是 2-否',
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
) ENGINE=InnoDB AUTO_INCREMENT=128 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;

-- ----------------------------
-- Records of sys_gen_column
-- ----------------------------
BEGIN;
INSERT INTO `sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_edit`, `is_must`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (1, 1, 'id', '账变编号', 'int', 'int64', 'Id', 'id', '1', '2', '2', NULL, '1', '2', 'EQ', 'input', '', 1, '', '2023-03-09 17:59:56', '2023-03-09 21:40:08', 0, 1);
INSERT INTO `sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_edit`, `is_must`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (2, 1, 'user_id', '用户编号', 'int', 'int64', 'UserId', 'userId', '2', '2', '2', NULL, '1', '1', 'EQ', 'input', '', 2, '', '2023-03-09 17:59:56', '2023-03-09 21:38:22', 0, 1);
INSERT INTO `sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_edit`, `is_must`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (3, 1, 'change_money', '账变金额', 'decimal(10,2)', 'decimal.Decimal', 'ChangeMoney', 'changeMoney', '2', '2', '2', NULL, '1', '2', 'EQ', 'input', '', 3, '', '2023-03-09 17:59:56', '2023-03-09 21:38:23', 0, 1);
INSERT INTO `sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_edit`, `is_must`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (4, 1, 'before_money', '账变前金额', 'decimal(30,18)', 'decimal.Decimal', 'BeforeMoney', 'beforeMoney', '2', '2', '2', NULL, '1', '2', 'EQ', 'input', '', 4, '', '2023-03-09 17:59:56', '2023-03-09 21:38:23', 0, 1);
INSERT INTO `sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_edit`, `is_must`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (5, 1, 'after_money', '账变后金额', 'decimal(30,18)', 'decimal.Decimal', 'AfterMoney', 'afterMoney', '2', '2', '2', NULL, '1', '2', 'EQ', 'input', '', 5, '', '2023-03-09 17:59:56', '2023-03-09 21:38:23', 0, 1);
INSERT INTO `sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_edit`, `is_must`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (6, 1, 'money_type', '金额类型 1:余额 ', 'char(10)', 'string', 'MoneyType', 'moneyType', '2', '2', '2', NULL, '1', '1', 'EQ', 'select', 'app_money_type', 6, '', '2023-03-09 17:59:56', '2023-03-09 21:38:23', 0, 1);
INSERT INTO `sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_edit`, `is_must`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (7, 1, 'change_type', '帐变类型(1-类型1)', 'varchar(30)', 'string', 'ChangeType', 'changeType', '2', '2', '2', NULL, '1', '1', 'EQ', 'select', 'app_account_change_type', 7, '', '2023-03-09 17:59:56', '2023-03-09 21:38:23', 0, 1);
INSERT INTO `sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_edit`, `is_must`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (8, 1, 'status', '状态（1正常 2-异常）', 'char(1)', 'string', 'Status', 'status', '2', '2', '2', NULL, NULL, '2', 'EQ', 'select', 'sys_status', 8, '', '2023-03-09 17:59:56', '2023-03-09 21:38:23', 0, 1);
INSERT INTO `sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_edit`, `is_must`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (9, 1, 'create_by', '创建者', 'int', 'int64', 'CreateBy', 'createBy', '2', '2', '2', NULL, NULL, '2', 'EQ', 'input', '', 9, '', '2023-03-09 17:59:56', '2023-03-09 21:38:23', 0, 1);
INSERT INTO `sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_edit`, `is_must`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (10, 1, 'created_at', '创建时间', 'datetime', '*time.Time', 'CreatedAt', 'createdAt', '2', '2', '2', NULL, '1', '2', 'EQ', 'datetime', '', 10, '', '2023-03-09 17:59:56', '2023-03-09 21:38:23', 0, 1);
INSERT INTO `sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_edit`, `is_must`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (11, 1, 'update_by', '更新者', 'int', 'int64', 'UpdateBy', 'updateBy', '2', '2', '2', NULL, NULL, '2', 'EQ', 'input', '', 11, '', '2023-03-09 17:59:56', '2023-03-09 21:38:23', 0, 1);
INSERT INTO `sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_edit`, `is_must`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (12, 1, 'updated_date', '更新时间', 'datetime', '*time.Time', 'UpdatedDate', 'updatedDate', '2', '2', '2', NULL, NULL, '2', 'EQ', 'datetime', '', 12, '', '2023-03-09 17:59:56', '2023-03-09 21:38:23', 0, 1);
INSERT INTO `sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_edit`, `is_must`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (13, 1, 'remarks', '备注信息', 'varchar(500)', 'string', 'Remarks', 'remarks', '2', '2', '2', NULL, NULL, '2', 'EQ', 'input', '', 13, '', '2023-03-09 17:59:56', '2023-03-09 17:59:56', 0, 0);
INSERT INTO `sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_edit`, `is_must`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (14, 2, 'id', '等级编号', 'int', 'int64', 'Id', 'id', '1', '2', '2', NULL, '1', '2', 'EQ', 'input', '', 1, '', '2023-03-09 20:05:43', '2023-03-09 20:17:04', 0, 1);
INSERT INTO `sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_edit`, `is_must`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (15, 2, 'name', '等级名称', 'varchar(255)', 'string', 'Name', 'name', '2', '1', '1', '1', '1', '1', 'LIKE', 'input', '', 2, '', '2023-03-09 20:05:43', '2023-03-09 22:47:41', 0, 1);
INSERT INTO `sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_edit`, `is_must`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (16, 2, 'level_type', '等级类型', 'varchar(10)', 'string', 'LevelType', 'levelType', '2', '1', '1', '1', '1', '1', 'EQ', 'select', 'app_user_level_type', 3, '', '2023-03-09 20:05:43', '2023-03-09 22:47:41', 0, 1);
INSERT INTO `sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_edit`, `is_must`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (17, 2, 'level', '等级', 'int', 'int64', 'Level', 'level', '2', '1', '1', '1', '1', '1', 'EQ', 'numInput', '', 4, '', '2023-03-09 20:05:43', '2023-03-09 22:47:41', 0, 1);
INSERT INTO `sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_edit`, `is_must`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (18, 2, 'status', '状态(1-正常 2-异常)', 'char(1)', 'string', 'Status', 'status', '2', '2', '2', NULL, '2', '2', 'EQ', 'select', 'sys_status', 5, '', '2023-03-09 20:05:43', '2023-03-09 20:17:05', 0, 1);
INSERT INTO `sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_edit`, `is_must`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (19, 2, 'remark', '备注信息', 'varchar(500)', 'string', 'Remark', 'remark', '2', '2', '2', NULL, '2', '2', 'LIKE', 'input', '', 6, '', '2023-03-09 20:05:43', '2023-03-09 20:08:51', 0, 1);
INSERT INTO `sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_edit`, `is_must`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (20, 2, 'create_by', '创建者', 'int', 'int64', 'CreateBy', 'createBy', '2', '2', '2', NULL, '2', '2', 'EQ', 'input', '', 7, '', '2023-03-09 20:05:43', '2023-03-09 20:17:05', 0, 1);
INSERT INTO `sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_edit`, `is_must`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (21, 2, 'update_by', '更新者', 'int', 'int64', 'UpdateBy', 'updateBy', '2', '2', '2', NULL, '2', '2', 'EQ', 'input', '', 8, '', '2023-03-09 20:05:43', '2023-03-09 20:17:05', 0, 1);
INSERT INTO `sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_edit`, `is_must`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (22, 2, 'created_at', '创建时间', 'datetime', '*time.Time', 'CreatedAt', 'createdAt', '2', '2', '2', NULL, '1', '2', 'EQ', 'datetime', '', 9, '', '2023-03-09 20:05:43', '2023-03-09 20:17:05', 0, 1);
INSERT INTO `sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_edit`, `is_must`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (23, 2, 'updated_at', '更新时间', 'datetime', '*time.Time', 'UpdatedAt', 'updatedAt', '2', '2', '2', NULL, '2', '2', 'EQ', 'datetime', '', 10, '', '2023-03-09 20:05:43', '2023-03-09 20:17:05', 0, 1);
INSERT INTO `sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_edit`, `is_must`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (24, 3, 'id', '配置编号', 'int', 'int64', 'Id', 'id', '1', '2', '2', '2', '1', '2', 'EQ', 'input', '', 1, '', '2023-03-09 22:59:52', '2023-03-09 23:02:29', 0, 1);
INSERT INTO `sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_edit`, `is_must`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (25, 3, 'user_id', '用户编号', 'int', 'int64', 'UserId', 'userId', '2', '1', '2', '1', '1', '1', 'EQ', 'numInput', '', 2, '', '2023-03-09 22:59:52', '2023-03-09 23:09:54', 0, 1);
INSERT INTO `sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_edit`, `is_must`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (26, 3, 'can_login', '1-允许登陆；2-不允许登陆', 'char(1)', 'string', 'CanLogin', 'canLogin', '2', '1', '1', '1', '1', '1', 'EQ', 'select', 'sys_yes_no', 3, '', '2023-03-09 22:59:52', '2023-03-09 23:02:29', 0, 1);
INSERT INTO `sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_edit`, `is_must`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (27, 3, 'remark', '备注信息', 'varchar(500)', 'string', 'Remark', 'remark', '2', '2', '2', '2', '2', '2', 'EQ', 'input', '', 4, '', '2023-03-09 22:59:52', '2023-03-09 22:59:52', 0, 0);
INSERT INTO `sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_edit`, `is_must`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (28, 3, 'status', '状态（1-正常 2-异常）\n', 'char(1)', 'string', 'Status', 'status', '2', '2', '2', '2', '2', '2', 'EQ', 'select', 'sys_status', 5, '', '2023-03-09 22:59:52', '2023-03-09 23:02:29', 0, 1);
INSERT INTO `sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_edit`, `is_must`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (29, 3, 'create_by', '创建者', 'int', 'int64', 'CreateBy', 'createBy', '2', '2', '2', '2', '2', '2', 'EQ', 'input', '', 6, '', '2023-03-09 22:59:52', '2023-03-09 23:02:29', 0, 1);
INSERT INTO `sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_edit`, `is_must`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (30, 3, 'update_by', '更新者', 'int', 'int64', 'UpdateBy', 'updateBy', '2', '2', '2', '2', '2', '2', 'EQ', 'input', '', 7, '', '2023-03-09 22:59:52', '2023-03-09 23:02:29', 0, 1);
INSERT INTO `sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_edit`, `is_must`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (31, 3, 'created_at', '创建时间', 'datetime', '*time.Time', 'CreatedAt', 'createdAt', '2', '2', '2', '2', '1', '2', 'EQ', 'datetime', '', 8, '', '2023-03-09 22:59:52', '2023-03-09 23:02:29', 0, 1);
INSERT INTO `sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_edit`, `is_must`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (32, 3, 'updated_at', '更新时间', 'datetime', '*time.Time', 'UpdatedAt', 'updatedAt', '2', '2', '2', '2', '2', '2', 'EQ', 'datetime', '', 9, '', '2023-03-09 22:59:52', '2023-03-09 23:02:29', 0, 1);
INSERT INTO `sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_edit`, `is_must`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (33, 4, 'id', '用户编号', 'int', 'int64', 'Id', 'id', '1', '2', '2', '2', '1', '1', 'EQ', 'input', '', 1, '', '2023-03-09 23:12:17', '2023-03-09 23:18:35', 0, 1);
INSERT INTO `sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_edit`, `is_must`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (34, 4, 'level_id', '用户等级编号', 'int', 'int64', 'LevelId', 'levelId', '2', '1', '1', '1', '1', '1', 'EQ', 'numInput', '', 2, '', '2023-03-09 23:12:17', '2023-03-09 23:25:14', 0, 1);
INSERT INTO `sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_edit`, `is_must`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (35, 4, 'user_name', '用户昵称', 'varchar(100)', 'string', 'UserName', 'userName', '2', '1', '1', '1', '1', '1', 'EQ', 'input', '', 3, '', '2023-03-09 23:12:17', '2023-03-09 23:18:35', 0, 1);
INSERT INTO `sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_edit`, `is_must`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (36, 4, 'true_name', '真实姓名', 'varchar(100)', 'string', 'TrueName', 'trueName', '2', '1', '1', '1', '1', '1', 'EQ', 'input', '', 4, '', '2023-03-09 23:12:17', '2023-03-09 23:18:35', 0, 1);
INSERT INTO `sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_edit`, `is_must`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (37, 4, 'money', '余额', 'decimal(30,18)', 'decimal.Decimal', 'Money', 'money', '2', '1', '1', '1', '1', '2', 'EQ', 'input', '', 5, '', '2023-03-09 23:12:17', '2023-03-09 23:18:35', 0, 1);
INSERT INTO `sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_edit`, `is_must`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (38, 4, 'email', '电子邮箱', 'varchar(300)', 'string', 'Email', 'email', '2', '1', '1', '1', '1', '1', 'EQ', 'input', '', 6, '', '2023-03-09 23:12:17', '2023-03-09 23:18:35', 0, 1);
INSERT INTO `sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_edit`, `is_must`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (39, 4, 'mobile_title', '用户手机号国家前缀', 'varchar(255)', 'string', 'MobileTitle', 'mobileTitle', '2', '1', '1', '1', '1', '2', 'EQ', 'input', '', 7, '', '2023-03-09 23:12:17', '2023-03-09 23:18:35', 0, 1);
INSERT INTO `sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_edit`, `is_must`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (40, 4, 'mobile', '手机号码', 'varchar(100)', 'string', 'Mobile', 'mobile', '2', '1', '1', '1', '1', '1', 'EQ', 'input', '', 8, '', '2023-03-09 23:12:17', '2023-03-09 23:18:35', 0, 1);
INSERT INTO `sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_edit`, `is_must`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (41, 4, 'avatar', '头像路径', 'varchar(1000)', 'string', 'Avatar', 'avatar', '2', '2', '2', '2', '2', '2', 'EQ', 'input', '', 9, '', '2023-03-09 23:12:17', '2023-03-09 23:22:58', 0, 1);
INSERT INTO `sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_edit`, `is_must`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (42, 4, 'pay_pwd', '提现密码', 'varchar(100)', 'string', 'PayPwd', 'payPwd', '2', '2', '2', '2', '2', '2', 'EQ', 'input', '', 10, '', '2023-03-09 23:12:17', '2023-03-09 23:22:58', 0, 1);
INSERT INTO `sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_edit`, `is_must`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (43, 4, 'pwd', '登录密码', 'varchar(100)', 'string', 'Pwd', 'pwd', '2', '2', '2', '2', '2', '2', 'EQ', 'input', '', 11, '', '2023-03-09 23:12:17', '2023-03-09 23:22:59', 0, 1);
INSERT INTO `sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_edit`, `is_must`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (44, 4, 'ref_code', '推荐码', 'varchar(255)', 'string', 'RefCode', 'refCode', '2', '2', '2', '2', '1', '1', 'EQ', 'input', '', 12, '', '2023-03-09 23:12:17', '2023-03-09 23:22:59', 0, 1);
INSERT INTO `sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_edit`, `is_must`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (45, 4, 'parent_id', '父级编号', 'int', 'int64', 'ParentId', 'parentId', '2', '2', '2', '2', '1', '1', 'EQ', 'input', '', 13, '', '2023-03-09 23:12:17', '2023-03-09 23:22:59', 0, 1);
INSERT INTO `sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_edit`, `is_must`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (46, 4, 'parent_ids', '所有父级编号', 'varchar(1000)', 'string', 'ParentIds', 'parentIds', '2', '2', '2', '2', '2', '2', 'EQ', 'input', '', 14, '', '2023-03-09 23:12:17', '2023-03-09 23:18:36', 0, 1);
INSERT INTO `sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_edit`, `is_must`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (47, 4, 'tree_sort', '本级排序号（升序）', 'decimal(10,0)', 'decimal.Decimal', 'TreeSort', 'treeSort', '2', '2', '2', '2', '2', '2', 'EQ', 'input', '', 15, '', '2023-03-09 23:12:17', '2023-03-09 23:18:36', 0, 1);
INSERT INTO `sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_edit`, `is_must`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (48, 4, 'tree_sorts', '所有级别排序号', 'varchar(1000)', 'string', 'TreeSorts', 'treeSorts', '2', '2', '2', '2', '2', '2', 'EQ', 'input', '', 16, '', '2023-03-09 23:12:17', '2023-03-09 23:18:36', 0, 1);
INSERT INTO `sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_edit`, `is_must`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (49, 4, 'tree_leaf', '是否最末级', 'char(1)', 'string', 'TreeLeaf', 'treeLeaf', '2', '2', '2', '2', '2', '2', 'EQ', 'input', '', 17, '', '2023-03-09 23:12:17', '2023-03-09 23:18:36', 0, 1);
INSERT INTO `sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_edit`, `is_must`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (50, 4, 'tree_level', '层次级别', 'int', 'int64', 'TreeLevel', 'treeLevel', '2', '2', '2', '2', '2', '2', 'EQ', 'input', '', 18, '', '2023-03-09 23:12:17', '2023-03-09 23:18:36', 0, 1);
INSERT INTO `sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_edit`, `is_must`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (51, 4, 'status', '状态(1-正常 2-异常)', 'char(1)', 'string', 'Status', 'status', '2', '2', '2', '2', '2', '2', 'EQ', 'select', 'sys_status', 19, '', '2023-03-09 23:12:17', '2023-03-09 23:18:36', 0, 1);
INSERT INTO `sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_edit`, `is_must`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (52, 4, 'remark', '备注信息', 'varchar(500)', 'string', 'Remark', 'remark', '2', '2', '2', '2', '2', '2', 'EQ', 'input', '', 20, '', '2023-03-09 23:12:17', '2023-03-09 23:12:17', 0, 0);
INSERT INTO `sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_edit`, `is_must`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (53, 4, 'create_by', '创建者', 'int', 'int64', 'CreateBy', 'createBy', '2', '2', '2', '2', '0', '2', 'EQ', 'input', '', 21, '', '2023-03-09 23:12:17', '2023-03-09 23:18:36', 0, 1);
INSERT INTO `sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_edit`, `is_must`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (54, 4, 'update_by', '更新者', 'int', 'int64', 'UpdateBy', 'updateBy', '2', '2', '2', '2', '0', '2', 'EQ', 'input', '', 22, '', '2023-03-09 23:12:17', '2023-03-09 23:18:36', 0, 1);
INSERT INTO `sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_edit`, `is_must`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (55, 4, 'created_at', '创建时间', 'datetime', '*time.Time', 'CreatedAt', 'createdAt', '2', '2', '2', '2', '1', '2', 'EQ', 'datetime', '', 23, '', '2023-03-09 23:12:17', '2023-03-09 23:18:36', 0, 1);
INSERT INTO `sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_edit`, `is_must`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (56, 4, 'updated_at', '更新时间', 'datetime', '*time.Time', 'UpdatedAt', 'updatedAt', '2', '2', '2', '2', '0', '2', 'EQ', 'datetime', '', 24, '', '2023-03-09 23:12:17', '2023-03-09 23:18:36', 0, 1);
INSERT INTO `sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_edit`, `is_must`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (57, 5, 'action_type', '用户行为类型', 'char(2)', 'string', 'ActionType', 'actionType', '2', '2', '2', '2', '1', '1', 'EQ', 'select', 'app_user_action_type', 1, '', '2023-03-11 14:00:15', '2023-03-11 14:08:37', 0, 1);
INSERT INTO `sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_edit`, `is_must`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (58, 5, 'by_type', '更新用户类型 1-app用户 2-后台用户', 'char(2)', 'string', 'ByType', 'byType', '2', '2', '2', '2', '1', '1', 'EQ', 'select', 'app_user_by_type', 2, '', '2023-03-11 14:00:15', '2023-03-11 14:15:30', 0, 1);
INSERT INTO `sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_edit`, `is_must`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (59, 5, 'create_by', '创建者', 'int', 'int64', 'CreateBy', 'createBy', '2', '2', '2', '2', '1', '2', 'EQ', 'input', '', 3, '', '2023-03-11 14:00:15', '2023-03-11 14:05:04', 0, 1);
INSERT INTO `sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_edit`, `is_must`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (60, 5, 'created_at', '创建时间', 'datetime', '*time.Time', 'CreatedAt', 'createdAt', '2', '2', '2', '2', '1', '2', 'EQ', 'datetime', '', 4, '', '2023-03-11 14:00:15', '2023-03-11 14:05:05', 0, 1);
INSERT INTO `sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_edit`, `is_must`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (61, 5, 'id', '日志编码', 'int', 'int64', 'Id', 'id', '1', '2', '2', '2', '1', '2', 'EQ', 'input', '', 5, '', '2023-03-11 14:00:15', '2023-03-11 14:05:05', 0, 1);
INSERT INTO `sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_edit`, `is_must`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (62, 5, 'remark', '备注信息', 'varchar(500)', 'string', 'Remark', 'remark', '2', '2', '2', '2', '2', '2', 'EQ', 'input', '', 6, '', '2023-03-11 14:00:15', '2023-03-11 14:00:15', 0, 0);
INSERT INTO `sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_edit`, `is_must`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (63, 5, 'status', '状态(1-正常 2-异常)', 'char(1)', 'string', 'Status', 'status', '2', '2', '2', '2', '2', '2', 'EQ', 'select', 'sys_status', 7, '', '2023-03-11 14:00:15', '2023-03-11 14:18:50', 0, 1);
INSERT INTO `sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_edit`, `is_must`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (64, 5, 'update_by', '更新者', 'int', 'int64', 'UpdateBy', 'updateBy', '2', '2', '2', '2', '1', '2', 'EQ', 'input', '', 8, '', '2023-03-11 14:00:15', '2023-03-11 14:05:05', 0, 1);
INSERT INTO `sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_edit`, `is_must`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (65, 5, 'updated_at', '更新时间', 'datetime', '*time.Time', 'UpdatedAt', 'updatedAt', '2', '2', '2', '2', '1', '2', 'EQ', 'datetime', '', 9, '', '2023-03-11 14:00:15', '2023-03-11 14:05:05', 0, 1);
INSERT INTO `sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_edit`, `is_must`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (66, 5, 'user_id', '用户编号', 'int', 'int64', 'UserId', 'userId', '2', '2', '2', '2', '1', '1', 'EQ', 'numInput', '', 10, '', '2023-03-11 14:00:15', '2023-03-11 14:05:05', 0, 1);
INSERT INTO `sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_edit`, `is_must`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (67, 6, 'id', '验证码编号', 'int', 'int64', 'Id', 'id', '1', '2', '2', '2', '1', '1', 'EQ', 'input', '', 1, '', '2023-03-12 12:11:09', '2023-03-12 12:14:11', 0, 1);
INSERT INTO `sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_edit`, `is_must`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (68, 6, 'user_id', '用户编号', 'int', 'int64', 'UserId', 'userId', '2', '2', '2', '2', '1', '1', 'EQ', 'input', '', 2, '', '2023-03-12 12:11:09', '2023-03-12 12:14:11', 0, 1);
INSERT INTO `sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_edit`, `is_must`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (69, 6, 'code', '验证码', 'varchar(12)', 'string', 'Code', 'code', '2', '2', '2', '2', '1', '2', 'EQ', 'input', '', 3, '', '2023-03-12 12:11:09', '2023-03-12 12:14:11', 0, 1);
INSERT INTO `sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_edit`, `is_must`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (70, 6, 'code_type', '验证码类型 1-邮箱；2-短信', 'char(1)', 'string', 'CodeType', 'codeType', '2', '2', '2', '2', '1', '1', 'EQ', 'select', 'plugin_msg_code_type', 4, '', '2023-03-12 12:11:09', '2023-03-12 12:16:18', 0, 1);
INSERT INTO `sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_edit`, `is_must`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (71, 6, 'remark', '备注异常', 'varchar(500)', 'string', 'Remark', 'remark', '2', '2', '2', '2', '1', '2', 'EQ', 'input', '', 5, '', '2023-03-12 12:11:09', '2023-03-12 12:14:12', 0, 1);
INSERT INTO `sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_edit`, `is_must`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (72, 6, 'status', '验证码状态 1-发送成功 2-发送失败', 'char(1)', 'string', 'Status', 'status', '2', '2', '2', '2', '1', '1', 'EQ', 'select', 'plugin_msg_sendstatus', 6, '', '2023-03-12 12:11:09', '2023-03-12 13:44:44', 0, 1);
INSERT INTO `sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_edit`, `is_must`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (73, 6, 'create_by', '创建者', 'int', 'int64', 'CreateBy', 'createBy', '2', '2', '2', '2', '2', '2', 'EQ', 'input', '', 7, '', '2023-03-12 12:11:09', '2023-03-12 12:14:12', 0, 1);
INSERT INTO `sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_edit`, `is_must`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (74, 6, 'update_by', '更新者', 'int', 'int64', 'UpdateBy', 'updateBy', '2', '2', '2', '2', '2', '2', 'EQ', 'input', '', 8, '', '2023-03-12 12:11:09', '2023-03-12 12:14:12', 0, 1);
INSERT INTO `sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_edit`, `is_must`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (75, 6, 'created_at', '创建时间', 'datetime', '*time.Time', 'CreatedAt', 'createdAt', '2', '2', '2', '2', '2', '2', 'EQ', 'datetime', '', 9, '', '2023-03-12 12:11:09', '2023-03-12 12:14:12', 0, 1);
INSERT INTO `sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_edit`, `is_must`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (76, 6, 'updated_at', '更新时间', 'datetime', '*time.Time', 'UpdatedAt', 'updatedAt', '2', '2', '2', '2', '2', '2', 'EQ', 'datetime', '', 10, '', '2023-03-12 12:11:09', '2023-03-12 12:14:12', 0, 1);
INSERT INTO `sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_edit`, `is_must`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (77, 7, 'id', '公告编号', 'int', 'int64', 'Id', 'id', '1', '2', '2', '2', '1', '2', 'EQ', 'input', '', 1, '', '2023-03-12 22:01:07', '2023-03-12 22:15:05', 0, 1);
INSERT INTO `sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_edit`, `is_must`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (78, 7, 'title', '标题', 'varchar(255)', 'string', 'Title', 'title', '2', '1', '1', '1', '1', '1', 'EQ', 'input', '', 2, '', '2023-03-12 22:01:07', '2023-03-12 22:15:05', 0, 1);
INSERT INTO `sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_edit`, `is_must`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (79, 7, 'content', '内容', 'text', 'string', 'Content', 'content', '2', '1', '1', '1', '1', '2', 'EQ', 'input', '', 3, '', '2023-03-12 22:01:07', '2023-03-12 22:15:05', 0, 1);
INSERT INTO `sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_edit`, `is_must`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (80, 7, 'num', '阅读次数', 'int', 'int64', 'Num', 'num', '2', '1', '1', '1', '1', '2', 'EQ', 'numInput', '', 4, '', '2023-03-12 22:01:07', '2023-03-12 22:51:59', 0, 1);
INSERT INTO `sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_edit`, `is_must`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (81, 7, 'remark', '备注信息', 'varchar(500)', 'string', 'Remark', 'remark', '2', '1', '1', '1', '1', '2', 'EQ', 'input', '', 5, '', '2023-03-12 22:01:07', '2023-03-12 22:15:05', 0, 1);
INSERT INTO `sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_edit`, `is_must`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (82, 7, 'status', '状态（0正常 1删除 2停用 3冻结）', 'char(1)', 'string', 'Status', 'status', '2', '1', '1', '1', '1', '1', 'EQ', 'select', 'sys_status', 6, '', '2023-03-12 22:01:07', '2023-03-12 22:15:05', 0, 1);
INSERT INTO `sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_edit`, `is_must`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (83, 7, 'create_by', '创建者', 'int', 'int64', 'CreateBy', 'createBy', '2', '2', '2', '2', '2', '2', 'EQ', 'input', '', 7, '', '2023-03-12 22:01:07', '2023-03-12 22:15:05', 0, 1);
INSERT INTO `sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_edit`, `is_must`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (84, 7, 'update_by', '更新者', 'int', 'int64', 'UpdateBy', 'updateBy', '2', '2', '2', '2', '2', '2', 'EQ', 'input', '', 8, '', '2023-03-12 22:01:07', '2023-03-12 22:15:05', 0, 1);
INSERT INTO `sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_edit`, `is_must`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (85, 7, 'updated_at', '更新时间', 'datetime', '*time.Time', 'UpdatedAt', 'updatedAt', '2', '2', '2', '2', '1', '2', 'EQ', 'datetime', '', 9, '', '2023-03-12 22:01:07', '2023-03-12 22:15:06', 0, 1);
INSERT INTO `sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_edit`, `is_must`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (86, 7, 'created_at', '创建时间', 'datetime', '*time.Time', 'CreatedAt', 'createdAt', '2', '2', '2', '2', '1', '2', 'EQ', 'datetime', '', 10, '', '2023-03-12 22:01:07', '2023-03-12 22:15:06', 0, 1);
INSERT INTO `sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_edit`, `is_must`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (87, 8, 'id', '分类编号', 'int', 'int64', 'Id', 'id', '1', '2', '2', '2', '1', '2', 'EQ', 'input', '', 1, '', '2023-03-12 22:54:51', '2023-03-12 22:57:39', 0, 1);
INSERT INTO `sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_edit`, `is_must`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (88, 8, 'name', '分类名称', 'varchar(255)', 'string', 'Name', 'name', '2', '1', '2', '1', '1', '1', 'LIKE', 'input', '', 2, '', '2023-03-12 22:54:51', '2023-03-12 22:57:39', 0, 1);
INSERT INTO `sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_edit`, `is_must`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (89, 8, 'status', '状态（1-正常 2-异常）', 'char(1)', 'string', 'Status', 'status', '2', '2', '2', '2', '2', '2', 'EQ', 'select', 'sys_status', 3, '', '2023-03-12 22:54:51', '2023-03-12 22:57:39', 0, 1);
INSERT INTO `sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_edit`, `is_must`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (90, 8, 'remark', '备注信息', 'varchar(500)', 'string', 'Remark', 'remark', '2', '2', '2', '2', '2', '2', 'EQ', 'input', '', 4, '', '2023-03-12 22:54:51', '2023-03-12 22:54:51', 0, 0);
INSERT INTO `sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_edit`, `is_must`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (91, 8, 'create_by', '更新人编号', 'int', 'int64', 'CreateBy', 'createBy', '2', '2', '2', '2', '1', '2', 'EQ', 'input', '', 5, '', '2023-03-12 22:54:51', '2023-03-12 22:57:39', 0, 1);
INSERT INTO `sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_edit`, `is_must`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (92, 8, 'update_by', '更新者', 'int', 'int64', 'UpdateBy', 'updateBy', '2', '2', '2', '2', '2', '2', 'EQ', 'input', '', 6, '', '2023-03-12 22:54:51', '2023-03-12 22:57:39', 0, 1);
INSERT INTO `sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_edit`, `is_must`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (93, 8, 'updated_at', '更新时间', 'datetime', '*time.Time', 'UpdatedAt', 'updatedAt', '2', '2', '2', '2', '1', '2', 'EQ', 'datetime', '', 7, '', '2023-03-12 22:54:51', '2023-03-12 22:57:39', 0, 1);
INSERT INTO `sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_edit`, `is_must`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (94, 8, 'created_at', '创建时间', 'datetime', '*time.Time', 'CreatedAt', 'createdAt', '2', '2', '2', '2', '1', '2', 'EQ', 'datetime', '', 8, '', '2023-03-12 22:54:51', '2023-03-12 22:57:39', 0, 1);
INSERT INTO `sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_edit`, `is_must`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (95, 9, 'id', '文章编号', 'int', 'int64', 'Id', 'id', '1', '2', '2', '2', '1', '1', 'EQ', 'input', '', 1, '', '2023-03-12 23:22:39', '2023-03-12 23:27:47', 0, 1);
INSERT INTO `sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_edit`, `is_must`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (96, 9, 'cate_id', '分类编号', 'int', 'int64', 'CateId', 'cateId', '2', '1', '1', '1', '1', '1', 'EQ', 'numInput', '', 2, '', '2023-03-12 23:22:39', '2023-03-12 23:27:47', 0, 1);
INSERT INTO `sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_edit`, `is_must`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (97, 9, 'name', '标题', 'varchar(255)', 'string', 'Name', 'name', '2', '1', '1', '1', '1', '1', 'EQ', 'input', '', 3, '', '2023-03-12 23:22:39', '2023-03-12 23:27:47', 0, 1);
INSERT INTO `sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_edit`, `is_must`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (98, 9, 'content', '内容', 'text', 'string', 'Content', 'content', '2', '1', '1', '1', '1', '2', 'EQ', 'input', '', 4, '', '2023-03-12 23:22:39', '2023-03-12 23:27:48', 0, 1);
INSERT INTO `sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_edit`, `is_must`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (99, 9, 'remark', '备注信息', 'varchar(500)', 'string', 'Remark', 'remark', '2', '1', '1', '1', '1', '2', 'EQ', 'input', '', 5, '', '2023-03-12 23:22:39', '2023-03-12 23:27:48', 0, 1);
INSERT INTO `sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_edit`, `is_must`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (100, 9, 'status', '状态（1-正常 2-异常）', 'char(1)', 'string', 'Status', 'status', '2', '1', '1', '1', '1', '2', 'EQ', 'select', 'sys_status', 6, '', '2023-03-12 23:22:39', '2023-03-12 23:27:48', 0, 1);
INSERT INTO `sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_edit`, `is_must`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (101, 9, 'create_by', '更新人编号', 'int', 'int64', 'CreateBy', 'createBy', '2', '2', '2', '2', '1', '2', 'EQ', 'input', '', 7, '', '2023-03-12 23:22:39', '2023-03-12 23:27:48', 0, 1);
INSERT INTO `sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_edit`, `is_must`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (102, 9, 'update_by', '更新者', 'int', 'int64', 'UpdateBy', 'updateBy', '2', '2', '2', '2', '2', '2', 'EQ', 'input', '', 8, '', '2023-03-12 23:22:39', '2023-03-12 23:27:48', 0, 1);
INSERT INTO `sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_edit`, `is_must`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (103, 9, 'updated_at', '更新时间', 'datetime', '*time.Time', 'UpdatedAt', 'updatedAt', '2', '2', '2', '2', '1', '2', 'EQ', 'datetime', '', 9, '', '2023-03-12 23:22:39', '2023-03-12 23:27:48', 0, 1);
INSERT INTO `sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_edit`, `is_must`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (104, 9, 'created_at', '创建时间', 'datetime', '*time.Time', 'CreatedAt', 'createdAt', '2', '2', '2', '2', '1', '2', 'EQ', 'datetime', '', 10, '', '2023-03-12 23:22:39', '2023-03-12 23:27:48', 0, 1);
INSERT INTO `sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_edit`, `is_must`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (105, 10, 'id', 'App编号', 'int', 'int64', 'Id', 'id', '1', '2', '2', '2', '1', '1', 'EQ', 'input', '', 1, '', '2023-03-13 00:07:25', '2023-03-13 00:12:43', 0, 1);
INSERT INTO `sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_edit`, `is_must`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (106, 10, 'version', '版本号', 'varchar(100)', 'string', 'Version', 'version', '2', '1', '1', '1', '1', '1', 'EQ', 'input', '', 2, '', '2023-03-13 00:07:25', '2023-03-13 00:12:43', 0, 1);
INSERT INTO `sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_edit`, `is_must`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (107, 10, 'platform', '平台 (1-安卓 2-苹果)', 'char(1)', 'string', 'Platform', 'platform', '2', '1', '1', '1', '1', '1', 'EQ', 'select', 'plugin_filemgr_app_platform', 3, '', '2023-03-13 00:07:25', '2023-03-13 00:12:43', 0, 1);
INSERT INTO `sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_edit`, `is_must`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (108, 10, 'app_type', '版本(1-默认)', 'char(1)', 'string', 'AppType', 'appType', '2', '1', '1', '1', '1', '1', 'EQ', 'select', 'plugin_filemgr_app_type', 4, '', '2023-03-13 00:07:25', '2023-03-13 00:12:43', 0, 1);
INSERT INTO `sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_edit`, `is_must`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (109, 10, 'local_address', '本地地址', 'varchar(255)', 'string', 'LocalAddress', 'localAddress', '2', '1', '1', '1', '1', '2', 'EQ', 'input', '', 5, '', '2023-03-13 00:07:25', '2023-03-13 00:12:43', 0, 1);
INSERT INTO `sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_edit`, `is_must`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (110, 10, 'download_num', '下载数量', 'int', 'int64', 'DownloadNum', 'downloadNum', '2', '1', '1', '1', '1', '2', 'EQ', 'input', '', 6, '', '2023-03-13 00:07:25', '2023-03-13 00:12:43', 0, 1);
INSERT INTO `sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_edit`, `is_must`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (111, 10, 'download_type', '下载类型(1-本地 2-外链 3-oss )', 'char(1)', 'string', 'DownloadType', 'downloadType', '2', '1', '1', '1', '1', '1', 'EQ', 'select', 'plugin_filemgr_app_download_type', 7, '', '2023-03-13 00:07:25', '2023-03-13 00:12:44', 0, 1);
INSERT INTO `sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_edit`, `is_must`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (112, 10, 'download_url', '下载地址(download_type=1使用)', 'varchar(255)', 'string', 'DownloadUrl', 'downloadUrl', '2', '1', '1', '1', '1', '2', 'EQ', 'input', '', 8, '', '2023-03-13 00:07:25', '2023-03-13 00:12:44', 0, 1);
INSERT INTO `sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_edit`, `is_must`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (113, 10, 'remark', '备注信息', 'varchar(500)', 'string', 'Remark', 'remark', '2', '1', '1', '1', '1', '2', 'EQ', 'input', '', 9, '', '2023-03-13 00:07:25', '2023-03-13 00:12:44', 0, 1);
INSERT INTO `sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_edit`, `is_must`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (114, 10, 'status', '状态（1-已发布 2-待发布）\n', 'char(1)', 'string', 'Status', 'status', '2', '1', '1', '1', '1', '1', 'EQ', 'select', 'plugin_filemgr_publish_status', 10, '', '2023-03-13 00:07:25', '2023-03-13 00:12:44', 0, 1);
INSERT INTO `sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_edit`, `is_must`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (115, 10, 'create_by', '创建者', 'int', 'int64', 'CreateBy', 'createBy', '2', '2', '2', '2', '2', '2', 'EQ', 'input', '', 11, '', '2023-03-13 00:07:25', '2023-03-13 00:12:44', 0, 1);
INSERT INTO `sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_edit`, `is_must`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (116, 10, 'created_at', '创建时间', 'datetime', '*time.Time', 'CreatedAt', 'createdAt', '2', '2', '2', '2', '1', '2', 'EQ', 'datetime', '', 12, '', '2023-03-13 00:07:25', '2023-03-13 00:12:44', 0, 1);
INSERT INTO `sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_edit`, `is_must`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (117, 10, 'update_by', '更新者', 'int', 'int64', 'UpdateBy', 'updateBy', '2', '2', '2', '2', '2', '2', 'EQ', 'input', '', 13, '', '2023-03-13 00:07:25', '2023-03-13 00:12:44', 0, 1);
INSERT INTO `sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_edit`, `is_must`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (118, 10, 'updated_at', '更新时间', 'datetime', '*time.Time', 'UpdatedAt', 'updatedAt', '2', '2', '2', '2', '2', '2', 'EQ', 'datetime', '', 14, '', '2023-03-13 00:07:25', '2023-03-13 00:12:44', 0, 1);
INSERT INTO `sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_edit`, `is_must`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (119, 11, 'id', '编号', 'int', 'int64', 'Id', 'id', '1', '2', '2', '2', '1', '2', 'EQ', 'input', '', 1, '', '2023-03-14 17:40:50', '2023-03-14 17:42:59', 0, 1);
INSERT INTO `sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_edit`, `is_must`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (120, 11, 'country', '国家地区', 'varchar(64)', 'string', 'Country', 'country', '2', '1', '1', '1', '1', '1', 'LIKE', 'input', '', 2, '', '2023-03-14 17:40:50', '2023-03-14 17:43:00', 0, 1);
INSERT INTO `sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_edit`, `is_must`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (121, 11, 'code', '区号', 'varchar(12)', 'string', 'Code', 'code', '2', '1', '1', '1', '1', '1', 'EQ', 'input', '', 3, '', '2023-03-14 17:40:50', '2023-03-14 17:43:00', 0, 1);
INSERT INTO `sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_edit`, `is_must`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (122, 11, 'status', '状态(1-可用 2-停用)', 'char(1)', 'string', 'Status', 'status', '2', '1', '1', '1', '1', '1', 'EQ', 'select', 'sys_status', 4, '', '2023-03-14 17:40:50', '2023-03-14 17:43:00', 0, 1);
INSERT INTO `sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_edit`, `is_must`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (123, 11, 'remark', '备注信息', 'varchar(500)', 'string', 'Remark', 'remark', '2', '2', '2', '2', '2', '2', 'EQ', 'input', '', 5, '', '2023-03-14 17:40:50', '2023-03-14 17:40:50', 0, 0);
INSERT INTO `sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_edit`, `is_must`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (124, 11, 'create_by', '创建者', 'int', 'int64', 'CreateBy', 'createBy', '2', '2', '2', '2', '2', '2', 'EQ', 'input', '', 6, '', '2023-03-14 17:40:50', '2023-03-14 17:43:00', 0, 1);
INSERT INTO `sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_edit`, `is_must`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (125, 11, 'update_by', '更新者', 'int', 'int64', 'UpdateBy', 'updateBy', '2', '2', '2', '2', '2', '2', 'EQ', 'input', '', 7, '', '2023-03-14 17:40:50', '2023-03-14 17:43:00', 0, 1);
INSERT INTO `sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_edit`, `is_must`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (126, 11, 'created_at', '创建时间', 'datetime', '*time.Time', 'CreatedAt', 'createdAt', '2', '2', '2', '2', '1', '2', 'EQ', 'datetime', '', 8, '', '2023-03-14 17:40:50', '2023-03-14 17:43:00', 0, 1);
INSERT INTO `sys_gen_column` (`id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_required`, `is_edit`, `is_must`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (127, 11, 'updated_at', '更新时间', 'datetime', '*time.Time', 'UpdatedAt', 'updatedAt', '2', '2', '2', '2', '2', '2', 'EQ', 'datetime', '', 9, '', '2023-03-14 17:40:50', '2023-03-14 17:43:00', 0, 1);
COMMIT;

-- ----------------------------
-- Table structure for sys_gen_table
-- ----------------------------
DROP TABLE IF EXISTS `sys_gen_table`;
CREATE TABLE `sys_gen_table` (
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
) ENGINE=InnoDB AUTO_INCREMENT=12 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;

-- ----------------------------
-- Records of sys_gen_table
-- ----------------------------
BEGIN;
INSERT INTO `sys_gen_table` (`id`, `table_name`, `table_comment`, `class_name`, `package_name`, `module_name`, `function_name`, `function_author`, `business_name`, `is_plugin`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (1, 'app_user_account_log', '账变记录', 'UserAccountLog', 'app', 'user-account-log', '账变记录', 'Jason', 'user', '2', '', '2023-03-09 17:59:56', '2023-03-09 17:59:56', 0, 0);
INSERT INTO `sys_gen_table` (`id`, `table_name`, `table_comment`, `class_name`, `package_name`, `module_name`, `function_name`, `function_author`, `business_name`, `is_plugin`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (2, 'app_user_level', '用户等级', 'UserLevel', 'app', 'user-level', '用户等级', 'Jason', 'user', '2', '', '2023-03-09 20:05:43', '2023-03-09 20:05:43', 0, 0);
INSERT INTO `sys_gen_table` (`id`, `table_name`, `table_comment`, `class_name`, `package_name`, `module_name`, `function_name`, `function_author`, `business_name`, `is_plugin`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (3, 'app_user_conf', '用户配置', 'UserConf', 'app', 'user-conf', '用户配置', 'Jason', 'user', '2', '', '2023-03-09 22:59:52', '2023-03-09 22:59:52', 0, 0);
INSERT INTO `sys_gen_table` (`id`, `table_name`, `table_comment`, `class_name`, `package_name`, `module_name`, `function_name`, `function_author`, `business_name`, `is_plugin`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (4, 'app_user', '用户管理', 'User', 'app', 'user', '用户管理', 'Jason', 'user', '2', '', '2023-03-09 23:12:17', '2023-03-09 23:12:17', 0, 0);
INSERT INTO `sys_gen_table` (`id`, `table_name`, `table_comment`, `class_name`, `package_name`, `module_name`, `function_name`, `function_author`, `business_name`, `is_plugin`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (5, 'app_user_oper_log', '用户关键行为日志表', 'UserOperLog', 'app', 'user-oper-log', '用户关键行为日志', 'Jason', 'user', '2', '', '2023-03-11 14:00:15', '2023-03-11 14:05:04', 0, 1);
INSERT INTO `sys_gen_table` (`id`, `table_name`, `table_comment`, `class_name`, `package_name`, `module_name`, `function_name`, `function_author`, `business_name`, `is_plugin`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (6, 'plugins_msg_code', '验证码记录', 'MsgCode', 'plugins', 'msg-code', '验证码记录', 'Jason', 'msg', '1', '', '2023-03-12 12:11:08', '2023-03-12 14:26:24', 0, 1);
INSERT INTO `sys_gen_table` (`id`, `table_name`, `table_comment`, `class_name`, `package_name`, `module_name`, `function_name`, `function_author`, `business_name`, `is_plugin`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (7, 'plugins_content_announcement', '公告管理', 'ContentAnnouncement', 'plugins', 'content-announcement', '公告管理', 'Jason', 'content', '2', '', '2023-03-12 22:01:07', '2023-03-12 22:01:07', 0, 0);
INSERT INTO `sys_gen_table` (`id`, `table_name`, `table_comment`, `class_name`, `package_name`, `module_name`, `function_name`, `function_author`, `business_name`, `is_plugin`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (8, 'plugins_content_category', '内容分类', 'ContentCategory', 'plugins', 'content-category', '文章分类管理', 'Jason', 'content', '1', '', '2023-03-12 22:54:51', '2023-03-12 22:58:31', 0, 1);
INSERT INTO `sys_gen_table` (`id`, `table_name`, `table_comment`, `class_name`, `package_name`, `module_name`, `function_name`, `function_author`, `business_name`, `is_plugin`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (9, 'plugins_content_article', '文章管理', 'ContentArticle', 'plugins', 'content-article', '文章管理', 'Jason', 'content', '1', '', '2023-03-12 23:22:39', '2023-03-12 23:22:39', 0, 0);
INSERT INTO `sys_gen_table` (`id`, `table_name`, `table_comment`, `class_name`, `package_name`, `module_name`, `function_name`, `function_author`, `business_name`, `is_plugin`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (10, 'plugins_filemgr_app', 'App管理', 'FilemgrApp', 'plugins', 'filemgr-app', 'App管理', 'Jason', 'filemgr', '1', '', '2023-03-13 00:07:25', '2023-03-13 00:12:43', 0, 1);
INSERT INTO `sys_gen_table` (`id`, `table_name`, `table_comment`, `class_name`, `package_name`, `module_name`, `function_name`, `function_author`, `business_name`, `is_plugin`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (11, 'app_user_country_code', '国家电话区号', 'UserCountryCode', 'app', 'user-country-code', '国家区号', 'Jason', 'user', '1', '', '2023-03-14 17:40:50', '2023-03-14 17:43:22', 0, 1);
COMMIT;

-- ----------------------------
-- Table structure for sys_login_log
-- ----------------------------
DROP TABLE IF EXISTS `sys_login_log`;
CREATE TABLE `sys_login_log` (
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
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;

-- ----------------------------
-- Records of sys_login_log
-- ----------------------------
BEGIN;
INSERT INTO `sys_login_log` (`id`, `user_id`, `ipaddr`, `login_location`, `browser`, `os`, `agent`, `platform`, `login_time`, `status`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (1,1, '127.0.0.1', '内部IP', 'Chrome 110.0.0.0', 'Intel Mac OS X 10_15_7', 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.0.0 Safari/537.36 Edg/110.0.1587.50', 'Macintosh', '2023-03-01 14:50:52', '2', '退出成功', '2023-03-01 14:50:52', '2023-03-01 14:50:52', 0, 0);
INSERT INTO `sys_login_log` (`id`, `user_id`, `ipaddr`, `login_location`, `browser`, `os`, `agent`, `platform`, `login_time`, `status`, `remark`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (2, 2, '127.0.0.1', '内部IP', 'Chrome 110.0.0.0', 'Intel Mac OS X 10_15_7', 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.0.0 Safari/537.36 Edg/110.0.1587.57', 'Macintosh', '2023-03-02 12:35:06', '1', '登录操作', '2023-03-02 12:35:06', '2023-03-02 12:35:06', 0, 0);
COMMIT;

-- ----------------------------
-- Table structure for sys_menu
-- ----------------------------
DROP TABLE IF EXISTS `sys_menu`;
CREATE TABLE `sys_menu` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL,
  `title` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL,
  `icon` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL,
  `path` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL,
  `component` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL,
  `redirect` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL,
  `menu_type` char(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL,
  `permission` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL,
  `sort` int DEFAULT NULL,
  `parent_id` int DEFAULT NULL COMMENT '上级菜单id',
  `parent_ids` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '上级菜单id集合',
  `keep_alive` char(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '是否缓存 1-是 2-否',
  `is_affix` char(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '是否固定 1-是 2-否',
  `hidden` char(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '是否隐藏 1-是 2-否',
  `is_frame` char(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '是否外链 1-是 2-否',
  `create_by` int DEFAULT NULL COMMENT '创建者',
  `update_by` int DEFAULT NULL COMMENT '更新者',
  `created_at` datetime DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime DEFAULT NULL COMMENT '最后更新时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=119 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;

-- ----------------------------
-- Records of sys_menu
-- ----------------------------
BEGIN;
INSERT INTO `sys_menu` (`id`, `name`, `title`, `icon`, `path`, `component`, `redirect`, `menu_type`, `permission`, `sort`, `parent_id`, `parent_ids`, `keep_alive`, `is_affix`, `hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (1, '', '系统管理', 'api-server', '/sys', 'Layout', '/sys/sys-api', '1', '', 300, 0, '0,', '', '', '2', '2', 1, 1, '2021-05-20 21:58:46', '2024-11-21 15:57:20');
INSERT INTO `sys_menu` (`id`, `name`, `title`, `icon`, `path`, `component`, `redirect`, `menu_type`, `permission`, `sort`, `parent_id`, `parent_ids`, `keep_alive`, `is_affix`, `hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (2, 'SysUser', '用户管理', 'user', '/sys/sys-user', '/sys/user/index.vue', NULL, '2', '', 10, 1, '0,2', '2', '2', '2', '2', 1, 1, '2021-05-20 22:08:45', '2023-04-27 16:32:41');
INSERT INTO `sys_menu` (`id`, `name`, `title`, `icon`, `path`, `component`, `redirect`, `menu_type`, `permission`, `sort`, `parent_id`, `parent_ids`, `keep_alive`, `is_affix`, `hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (3, '', '新增管理员', 'app-group-fill', '', '', NULL, '3', 'admin:sysUser:add', 10, 2, '0,1,2,', NULL, NULL, NULL, NULL, 1, 1, '2021-05-20 22:08:45', '2023-04-27 16:40:14');
INSERT INTO `sys_menu` (`id`, `name`, `title`, `icon`, `path`, `component`, `redirect`, `menu_type`, `permission`, `sort`, `parent_id`, `parent_ids`, `keep_alive`, `is_affix`, `hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (4, '', '查询管理员', 'app-group-fill', '', '', NULL, '3', 'admin:sysUser:query', 40, 2, '0,1,2,', NULL, NULL, NULL, NULL, 1, 1, '2021-05-20 22:08:45', '2023-04-27 16:32:41');
INSERT INTO `sys_menu` (`id`, `name`, `title`, `icon`, `path`, `component`, `redirect`, `menu_type`, `permission`, `sort`, `parent_id`, `parent_ids`, `keep_alive`, `is_affix`, `hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (5, '', '修改管理员', 'app-group-fill', '', '', NULL, '3', 'admin:sysUser:edit', 30, 2, '0,1,2,', NULL, NULL, NULL, NULL, 1, 1, '2021-05-20 22:08:45', '2023-04-27 16:32:41');
INSERT INTO `sys_menu` (`id`, `name`, `title`, `icon`, `path`, `component`, `redirect`, `menu_type`, `permission`, `sort`, `parent_id`, `parent_ids`, `keep_alive`, `is_affix`, `hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (6, '', '删除管理员', 'app-group-fill', '', '', NULL, '3', 'admin:sysUser:remove', 20, 2, '0,1,2,', NULL, NULL, NULL, NULL, 1, 1, '2021-05-20 22:08:45', '2023-04-27 16:32:41');
INSERT INTO `sys_menu` (`id`, `name`, `title`, `icon`, `path`, `component`, `redirect`, `menu_type`, `permission`, `sort`, `parent_id`, `parent_ids`, `keep_alive`, `is_affix`, `hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (7, 'SysMenu', '菜单管理', 'tree-table', '/sys/sys-menu', '/sys/menu/index.vue', NULL, '2', '', 30, 1, '0,1,', '2', '2', '2', '2', 1, 1, '2021-05-20 22:08:45', '2023-04-27 16:32:41');
INSERT INTO `sys_menu` (`id`, `name`, `title`, `icon`, `path`, `component`, `redirect`, `menu_type`, `permission`, `sort`, `parent_id`, `parent_ids`, `keep_alive`, `is_affix`, `hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (8, 'SysRole', '角色管理', 'peoples', '/sys/sys-role', '/sys/role/index.vue', NULL, '2', '', 20, 1, '0,1,', '2', '2', '2', '2', 1, 1, '2021-05-20 22:08:45', '2023-04-27 16:32:41');
INSERT INTO `sys_menu` (`id`, `name`, `title`, `icon`, `path`, `component`, `redirect`, `menu_type`, `permission`, `sort`, `parent_id`, `parent_ids`, `keep_alive`, `is_affix`, `hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (9, 'SysDept', '部门管理', 'tree', '/sys/sys-dept', '/sys/dept/index.vue', NULL, '2', '', 40, 1, '0,1,', '2', '2', '2', '2', 1, 1, '2021-05-20 22:08:45', '2023-04-27 16:32:41');
INSERT INTO `sys_menu` (`id`, `name`, `title`, `icon`, `path`, `component`, `redirect`, `menu_type`, `permission`, `sort`, `parent_id`, `parent_ids`, `keep_alive`, `is_affix`, `hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (10, 'SysPost', '岗位管理', 'pass', '/sys/sys-post', '/sys/post/index.vue', NULL, '2', '', 50, 1, '0,1,', '2', '2', '2', '2', 1, 1, '2021-05-20 22:08:45', '2023-04-27 16:32:41');
INSERT INTO `sys_menu` (`id`, `name`, `title`, `icon`, `path`, `component`, `redirect`, `menu_type`, `permission`, `sort`, `parent_id`, `parent_ids`, `keep_alive`, `is_affix`, `hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (11, 'SysDicttype', '字典管理', 'education', '/sys/sys-dicttype', '/sys/dicttype/index.vue', NULL, '2', '', 60, 1, '0,1,', '2', '2', '2', '2', 1, 1, '2021-05-20 22:08:45', '2023-04-27 16:32:41');
INSERT INTO `sys_menu` (`id`, `name`, `title`, `icon`, `path`, `component`, `redirect`, `menu_type`, `permission`, `sort`, `parent_id`, `parent_ids`, `keep_alive`, `is_affix`, `hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (12, 'SysDictdata', '字典数据', 'education', '/sys/sys-dictdata', '/sys/dictdata/index.vue', NULL, '2', '', 100, 1, '0,1,', '2', '2', '1', '2', 1, 1, '2021-05-20 22:08:45', '2023-04-27 16:32:41');
INSERT INTO `sys_menu` (`id`, `name`, `title`, `icon`, `path`, `component`, `redirect`, `menu_type`, `permission`, `sort`, `parent_id`, `parent_ids`, `keep_alive`, `is_affix`, `hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (13, 'SysConfig', '参数管理', 'swagger', '/sys/sys-config', '/sys/config/index.vue', NULL, '2', '', 70, 1, '0,1,', '2', '2', '2', '2', 1, 1, '2021-05-20 22:08:45', '2023-04-27 16:32:41');
INSERT INTO `sys_menu` (`id`, `name`, `title`, `icon`, `path`, `component`, `redirect`, `menu_type`, `permission`, `sort`, `parent_id`, `parent_ids`, `keep_alive`, `is_affix`, `hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (14, 'SysLoginlog', '登录日志', 'logininfor', '/sys/sys-loginlog', '/sys/loginlog/index.vue', NULL, '2', '', 90, 1, '0,1,', '2', '2', '2', '2', 1, 1, '2021-05-20 22:08:45', '2023-04-27 16:32:41');
INSERT INTO `sys_menu` (`id`, `name`, `title`, `icon`, `path`, `component`, `redirect`, `menu_type`, `permission`, `sort`, `parent_id`, `parent_ids`, `keep_alive`, `is_affix`, `hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (15, 'SysOperalog', '操作日志', 'skill', '/sys/sys-operalog', '/sys/operlog/index.vue', NULL, '2', '', 120, 1, '0,1,', '2', '2', '2', '2', 1, 1, '2021-05-20 22:08:45', '2023-04-27 16:32:41');
INSERT INTO `sys_menu` (`id`, `name`, `title`, `icon`, `path`, `component`, `redirect`, `menu_type`, `permission`, `sort`, `parent_id`, `parent_ids`, `keep_alive`, `is_affix`, `hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (16, '', '新增菜单', 'app-group-fill', '', '', NULL, '3', 'admin:sysMenu:add', 1, 7, '0,1,7,', NULL, NULL, NULL, NULL, 1, 1, '2020-04-11 15:52:48', '2023-04-27 16:32:41');
INSERT INTO `sys_menu` (`id`, `name`, `title`, `icon`, `path`, `component`, `redirect`, `menu_type`, `permission`, `sort`, `parent_id`, `parent_ids`, `keep_alive`, `is_affix`, `hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (17, '', '修改菜单', 'app-group-fill', '', '', NULL, '3', 'admin:sysMenu:edit', 1, 7, '0,1,7,', NULL, NULL, NULL, NULL, 1, 1, '2020-04-11 15:52:48', '2023-04-27 16:32:41');
INSERT INTO `sys_menu` (`id`, `name`, `title`, `icon`, `path`, `component`, `redirect`, `menu_type`, `permission`, `sort`, `parent_id`, `parent_ids`, `keep_alive`, `is_affix`, `hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (18, '', '查询菜单', 'app-group-fill', '', '', NULL, '3', 'admin:sysMenu:query', 1, 7, '0,1,7,', NULL, NULL, NULL, NULL, 1, 1, '2020-04-11 15:52:48', '2023-04-27 16:32:41');
INSERT INTO `sys_menu` (`id`, `name`, `title`, `icon`, `path`, `component`, `redirect`, `menu_type`, `permission`, `sort`, `parent_id`, `parent_ids`, `keep_alive`, `is_affix`, `hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (19, '', '删除菜单', 'app-group-fill', '', '', NULL, '3', 'admin:sysMenu:remove', 1, 7, '0,1,7,', NULL, NULL, NULL, NULL, 1, 1, '2020-04-11 15:52:48', '2023-04-27 16:32:41');
INSERT INTO `sys_menu` (`id`, `name`, `title`, `icon`, `path`, `component`, `redirect`, `menu_type`, `permission`, `sort`, `parent_id`, `parent_ids`, `keep_alive`, `is_affix`, `hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (20, '', '新增角色', 'app-group-fill', '', '', NULL, '3', 'admin:sysRole:add', 1, 8, '0,1,8,', NULL, NULL, NULL, NULL, 1, 1, '2020-04-11 15:52:48', '2023-04-27 16:32:41');
INSERT INTO `sys_menu` (`id`, `name`, `title`, `icon`, `path`, `component`, `redirect`, `menu_type`, `permission`, `sort`, `parent_id`, `parent_ids`, `keep_alive`, `is_affix`, `hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (21, '', '查询角色', 'app-group-fill', '', '', NULL, '3', 'admin:sysRole:query', 1, 8, '0,1,8,', NULL, NULL, NULL, NULL, 1, 1, '2020-04-11 15:52:48', '2023-04-27 16:32:41');
INSERT INTO `sys_menu` (`id`, `name`, `title`, `icon`, `path`, `component`, `redirect`, `menu_type`, `permission`, `sort`, `parent_id`, `parent_ids`, `keep_alive`, `is_affix`, `hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (22, '', '修改角色', 'app-group-fill', '', '', NULL, '3', 'admin:sysRole:update', 1, 8, '0,1,8,', NULL, NULL, NULL, NULL, 1, 1, '2020-04-11 15:52:48', '2023-04-27 16:32:41');
INSERT INTO `sys_menu` (`id`, `name`, `title`, `icon`, `path`, `component`, `redirect`, `menu_type`, `permission`, `sort`, `parent_id`, `parent_ids`, `keep_alive`, `is_affix`, `hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (23, '', '删除角色', 'app-group-fill', '', '', NULL, '3', 'admin:sysRole:remove', 1, 8, '0,1,8,', NULL, NULL, NULL, NULL, 1, 1, '2020-04-11 15:52:48', '2023-04-27 16:32:41');
INSERT INTO `sys_menu` (`id`, `name`, `title`, `icon`, `path`, `component`, `redirect`, `menu_type`, `permission`, `sort`, `parent_id`, `parent_ids`, `keep_alive`, `is_affix`, `hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (24, '', '查询部门', 'app-group-fill', '', '', NULL, '3', 'admin:sysDept:query', 40, 9, '0,1,9,', NULL, NULL, NULL, NULL, 1, 1, '2021-05-20 22:08:45', '2023-04-27 16:32:41');
INSERT INTO `sys_menu` (`id`, `name`, `title`, `icon`, `path`, `component`, `redirect`, `menu_type`, `permission`, `sort`, `parent_id`, `parent_ids`, `keep_alive`, `is_affix`, `hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (25, '', '新增部门', 'app-group-fill', '', '', NULL, '3', 'admin:sysDept:add', 10, 9, '0,1,9,', NULL, NULL, NULL, NULL, 1, 1, '2021-05-20 22:08:45', '2023-04-27 16:32:41');
INSERT INTO `sys_menu` (`id`, `name`, `title`, `icon`, `path`, `component`, `redirect`, `menu_type`, `permission`, `sort`, `parent_id`, `parent_ids`, `keep_alive`, `is_affix`, `hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (26, '', '修改部门', 'app-group-fill', '', '', NULL, '3', 'admin:sysDept:edit', 30, 9, '0,1,9,', NULL, NULL, NULL, NULL, 1, 1, '2021-05-20 22:08:45', '2023-04-27 16:32:41');
INSERT INTO `sys_menu` (`id`, `name`, `title`, `icon`, `path`, `component`, `redirect`, `menu_type`, `permission`, `sort`, `parent_id`, `parent_ids`, `keep_alive`, `is_affix`, `hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (27, '', '删除部门', 'app-group-fill', '', '', NULL, '3', 'admin:sysDept:remove', 20, 9, '0,1,9,', NULL, NULL, NULL, NULL, 1, 1, '2021-05-20 22:08:45', '2023-04-27 16:32:41');
INSERT INTO `sys_menu` (`id`, `name`, `title`, `icon`, `path`, `component`, `redirect`, `menu_type`, `permission`, `sort`, `parent_id`, `parent_ids`, `keep_alive`, `is_affix`, `hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (28, '', '查询岗位', 'app-group-fill', '', '', NULL, '3', 'admin:sysPost:query', 0, 10, '0,1,10,', NULL, NULL, NULL, NULL, 1, 1, '2020-04-11 15:52:48', '2023-04-27 16:32:41');
INSERT INTO `sys_menu` (`id`, `name`, `title`, `icon`, `path`, `component`, `redirect`, `menu_type`, `permission`, `sort`, `parent_id`, `parent_ids`, `keep_alive`, `is_affix`, `hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (29, '', '新增岗位', 'app-group-fill', '', '', NULL, '3', 'admin:sysPost:add', 0, 10, '0,1,10,', NULL, NULL, NULL, NULL, 1, 1, '2020-04-11 15:52:48', '2023-04-27 16:32:41');
INSERT INTO `sys_menu` (`id`, `name`, `title`, `icon`, `path`, `component`, `redirect`, `menu_type`, `permission`, `sort`, `parent_id`, `parent_ids`, `keep_alive`, `is_affix`, `hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (30, '', '修改岗位', 'app-group-fill', '', '', NULL, '3', 'admin:sysPost:edit', 0, 10, '0,1,10,', NULL, NULL, NULL, NULL, 1, 1, '2020-04-11 15:52:48', '2023-04-27 16:32:41');
INSERT INTO `sys_menu` (`id`, `name`, `title`, `icon`, `path`, `component`, `redirect`, `menu_type`, `permission`, `sort`, `parent_id`, `parent_ids`, `keep_alive`, `is_affix`, `hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (31, '', '删除岗位', 'app-group-fill', '', '', NULL, '3', 'admin:sysPost:remove', 0, 10, '0,1,10,', NULL, NULL, NULL, NULL, 1, 1, '2020-04-11 15:52:48', '2023-04-27 16:32:41');
INSERT INTO `sys_menu` (`id`, `name`, `title`, `icon`, `path`, `component`, `redirect`, `menu_type`, `permission`, `sort`, `parent_id`, `parent_ids`, `keep_alive`, `is_affix`, `hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (32, '', '查询字典', 'app-group-fill', '', '', NULL, '3', 'admin:sysDictType:query', 0, 11, '0,1,11,', NULL, NULL, NULL, NULL, 1, 1, '2020-04-11 15:52:48', '2023-05-09 11:13:23');
INSERT INTO `sys_menu` (`id`, `name`, `title`, `icon`, `path`, `component`, `redirect`, `menu_type`, `permission`, `sort`, `parent_id`, `parent_ids`, `keep_alive`, `is_affix`, `hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (33, '', '新增类型', 'app-group-fill', '', '', NULL, '3', 'admin:sysDictType:add', 0, 11, '0,1,11,', NULL, NULL, NULL, NULL, 1, 1, '2020-04-11 15:52:48', '2023-05-09 11:13:37');
INSERT INTO `sys_menu` (`id`, `name`, `title`, `icon`, `path`, `component`, `redirect`, `menu_type`, `permission`, `sort`, `parent_id`, `parent_ids`, `keep_alive`, `is_affix`, `hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (34, '', '修改类型', 'app-group-fill', '', '', NULL, '3', 'admin:sysDictType:edit', 0, 11, '0,1,11,', NULL, NULL, NULL, NULL, 1, 1, '2020-04-11 15:52:48', '2023-05-09 11:13:58');
INSERT INTO `sys_menu` (`id`, `name`, `title`, `icon`, `path`, `component`, `redirect`, `menu_type`, `permission`, `sort`, `parent_id`, `parent_ids`, `keep_alive`, `is_affix`, `hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (35, '', '删除类型', 'app-group-fill', '', '', NULL, '3', 'admin:sysdicttype:remove', 0, 11, '0,1,11,', NULL, NULL, NULL, NULL, 1, 1, '2020-04-11 15:52:48', '2023-05-09 11:15:32');
INSERT INTO `sys_menu` (`id`, `name`, `title`, `icon`, `path`, `component`, `redirect`, `menu_type`, `permission`, `sort`, `parent_id`, `parent_ids`, `keep_alive`, `is_affix`, `hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (36, '', '查询数据', 'app-group-fill', '', '', NULL, '3', 'admin:sysDictData:query', 0, 12, '0,1,12,', NULL, NULL, NULL, NULL, 1, 1, '2020-04-11 15:52:48', '2023-05-09 11:09:55');
INSERT INTO `sys_menu` (`id`, `name`, `title`, `icon`, `path`, `component`, `redirect`, `menu_type`, `permission`, `sort`, `parent_id`, `parent_ids`, `keep_alive`, `is_affix`, `hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (37, '', '新增数据', 'app-group-fill', '', '', NULL, '3', 'admin:sysDictData:add', 0, 12, '0,1,12,', NULL, NULL, NULL, NULL, 1, 1, '2020-04-11 15:52:48', '2023-05-09 11:10:05');
INSERT INTO `sys_menu` (`id`, `name`, `title`, `icon`, `path`, `component`, `redirect`, `menu_type`, `permission`, `sort`, `parent_id`, `parent_ids`, `keep_alive`, `is_affix`, `hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (38, '', '修改数据', 'app-group-fill', '', '', NULL, '3', 'admin:sysDictData:edit', 0, 12, '0,1,12,', NULL, NULL, NULL, NULL, 1, 1, '2020-04-11 15:52:48', '2023-05-09 11:11:11');
INSERT INTO `sys_menu` (`id`, `name`, `title`, `icon`, `path`, `component`, `redirect`, `menu_type`, `permission`, `sort`, `parent_id`, `parent_ids`, `keep_alive`, `is_affix`, `hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (39, '', '删除数据', 'app-group-fill', '', '', NULL, '3', 'admin:sysDictData:remove', 0, 12, '0,1,12,', NULL, NULL, NULL, NULL, 1, 1, '2020-04-11 15:52:48', '2023-05-09 11:11:43');
INSERT INTO `sys_menu` (`id`, `name`, `title`, `icon`, `path`, `component`, `redirect`, `menu_type`, `permission`, `sort`, `parent_id`, `parent_ids`, `keep_alive`, `is_affix`, `hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (40, '', '查询参数', 'app-group-fill', '', '', NULL, '3', 'admin:sysConfig:query', 0, 13, '0,1,13,', NULL, NULL, NULL, NULL, 1, 1, '2020-04-11 15:52:48', '2023-05-09 11:25:51');
INSERT INTO `sys_menu` (`id`, `name`, `title`, `icon`, `path`, `component`, `redirect`, `menu_type`, `permission`, `sort`, `parent_id`, `parent_ids`, `keep_alive`, `is_affix`, `hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (41, '', '新增参数', 'app-group-fill', '', '', NULL, '3', 'admin:sysConfig:add', 0, 13, '0,1,13,', NULL, NULL, NULL, NULL, 1, 1, '2020-04-11 15:52:48', '2023-05-09 11:26:03');
INSERT INTO `sys_menu` (`id`, `name`, `title`, `icon`, `path`, `component`, `redirect`, `menu_type`, `permission`, `sort`, `parent_id`, `parent_ids`, `keep_alive`, `is_affix`, `hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (42, '', '修改参数', 'app-group-fill', '', '', NULL, '3', 'admin:sysConfig:edit', 0, 13, '0,1,13,', NULL, NULL, NULL, NULL, 1, 1, '2020-04-11 15:52:48', '2023-05-09 11:26:27');
INSERT INTO `sys_menu` (`id`, `name`, `title`, `icon`, `path`, `component`, `redirect`, `menu_type`, `permission`, `sort`, `parent_id`, `parent_ids`, `keep_alive`, `is_affix`, `hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (43, '', '删除参数', 'app-group-fill', '', '', NULL, '3', 'admin:sysConfig:remove', 0, 13, '0,1,13,', NULL, NULL, NULL, NULL, 1, 1, '2020-04-11 15:52:48', '2023-05-09 11:26:39');
INSERT INTO `sys_menu` (`id`, `name`, `title`, `icon`, `path`, `component`, `redirect`, `menu_type`, `permission`, `sort`, `parent_id`, `parent_ids`, `keep_alive`, `is_affix`, `hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (44, '', '查询登录日志', 'app-group-fill', '', '', NULL, '3', 'admin:sysLoginLog:query', 0, 14, '0,1,14,', NULL, NULL, NULL, NULL, 1, 1, '2020-04-11 15:52:48', '2023-05-09 11:03:23');
INSERT INTO `sys_menu` (`id`, `name`, `title`, `icon`, `path`, `component`, `redirect`, `menu_type`, `permission`, `sort`, `parent_id`, `parent_ids`, `keep_alive`, `is_affix`, `hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (45, '', '删除登录日志', 'app-group-fill', '', '', NULL, '3', 'admin:sysLoginLog:remove', 0, 14, '0,1,14,', NULL, NULL, NULL, NULL, 1, 1, '2020-04-11 15:52:48', '2023-05-09 11:03:46');
INSERT INTO `sys_menu` (`id`, `name`, `title`, `icon`, `path`, `component`, `redirect`, `menu_type`, `permission`, `sort`, `parent_id`, `parent_ids`, `keep_alive`, `is_affix`, `hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (46, '', '查询操作日志', 'app-group-fill', '', '', NULL, '3', 'admin:sysOperLog:query', 0, 15, '0,1,15,', NULL, NULL, NULL, NULL, 1, 1, '2020-04-11 15:52:48', '2023-05-09 11:00:44');
INSERT INTO `sys_menu` (`id`, `name`, `title`, `icon`, `path`, `component`, `redirect`, `menu_type`, `permission`, `sort`, `parent_id`, `parent_ids`, `keep_alive`, `is_affix`, `hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (47, '', '删除操作日志', 'app-group-fill', '', '', NULL, '3', 'admin:sysOperLog:remove', 0, 15, '0,1,15,', NULL, NULL, NULL, NULL, 1, 1, '2020-04-11 15:52:48', '2023-05-09 11:00:58');
INSERT INTO `sys_menu` (`id`, `name`, `title`, `icon`, `path`, `component`, `redirect`, `menu_type`, `permission`, `sort`, `parent_id`, `parent_ids`, `keep_alive`, `is_affix`, `hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (48, 'SysGen', '代码生成', 'code', '/sys-tools/sys-gen', '/sys/tools/gen/index.vue', NULL, '2', '', 20, 54, '0,54,', '2', '2', '2', '2', 1, 1, '2020-04-11 15:52:48', '2023-05-09 10:55:36');
INSERT INTO `sys_menu` (`id`, `name`, `title`, `icon`, `path`, `component`, `redirect`, `menu_type`, `permission`, `sort`, `parent_id`, `parent_ids`, `keep_alive`, `is_affix`, `hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (49, 'SysEditTable', '代码生成修改', 'build', '/sys-tools/sys-editTable', '/sys/tools/gen/editTable.vue', NULL, '2', '', 100, 54, '0,54,', '2', '2', '1', '2', 1, 1, '2020-04-11 15:52:48', '2023-05-09 11:14:37');
INSERT INTO `sys_menu` (`id`, `name`, `title`, `icon`, `path`, `component`, `redirect`, `menu_type`, `permission`, `sort`, `parent_id`, `parent_ids`, `keep_alive`, `is_affix`, `hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (50, 'SysMonitor', '服务监控', 'druid', '/sys-tools/sys-monitor', '/sys/tools/monitor/monitor.vue', NULL, '2', '', 0, 54, '0,54,', '2', '2', '2', '2', 1, 1, '2020-04-14 00:28:19', '2023-05-09 10:52:46');
INSERT INTO `sys_menu` (`id`, `name`, `title`, `icon`, `path`, `component`, `redirect`, `menu_type`, `permission`, `sort`, `parent_id`, `parent_ids`, `keep_alive`, `is_affix`, `hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (51, 'SysApi', '接口管理', 'api-doc', '/sys/sys-api', '/sys/api/index.vue', NULL, '2', '', 0, 1, '0,1,', '2', '2', '2', '2', 1, 1, '2021-05-20 22:08:45', '2023-04-27 16:32:41');
INSERT INTO `sys_menu` (`id`, `name`, `title`, `icon`, `path`, `component`, `redirect`, `menu_type`, `permission`, `sort`, `parent_id`, `parent_ids`, `keep_alive`, `is_affix`, `hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (52, '', '查询接口', 'app-group-fill', '', '', NULL, '3', 'admin:sysApi:query', 40, 51, '0,1,51,', NULL, NULL, NULL, NULL, 1, 1, '2021-05-20 22:08:45', '2023-04-27 16:32:41');
INSERT INTO `sys_menu` (`id`, `name`, `title`, `icon`, `path`, `component`, `redirect`, `menu_type`, `permission`, `sort`, `parent_id`, `parent_ids`, `keep_alive`, `is_affix`, `hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (53, '', '修改接口', 'app-group-fill', '', '', NULL, '3', 'admin:sysApi:edit', 30, 51, '0,1,51,', NULL, NULL, NULL, NULL, 1, 1, '2021-05-20 22:08:45', '2023-04-27 16:32:41');
INSERT INTO `sys_menu` (`id`, `name`, `title`, `icon`, `path`, `component`, `redirect`, `menu_type`, `permission`, `sort`, `parent_id`, `parent_ids`, `keep_alive`, `is_affix`, `hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (54, '', '系统工具', 'system-tools', '/sys-tools', 'Layout', '/sys-tools/sys-monitor', '1', '', 330, 0, '0,', '', '', '2', '2', 1, 1, '2021-05-21 11:13:32', '2024-11-21 15:57:40');
INSERT INTO `sys_menu` (`id`, `name`, `title`, `icon`, `path`, `component`, `redirect`, `menu_type`, `permission`, `sort`, `parent_id`, `parent_ids`, `keep_alive`, `is_affix`, `hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (55, '', '文件管理', 'base-info', '/plugins/filemgr', '/index', '/plugins/filemgr/filemgr-app', '1', '', 90, 57, '0,57,', '', '', '2', '2', 1, 1, '2021-08-13 14:19:11', '2024-11-21 15:56:20');
INSERT INTO `sys_menu` (`id`, `name`, `title`, `icon`, `path`, `component`, `redirect`, `menu_type`, `permission`, `sort`, `parent_id`, `parent_ids`, `keep_alive`, `is_affix`, `hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (56, '', '内容管理', 'clipboard', '/plugins/content', '/index', '/plugins/content/content-category', '1', '', 60, 57, '0,57,', '', '', '2', '2', 1, 1, '2021-08-16 18:01:20', '2024-11-21 15:55:48');
INSERT INTO `sys_menu` (`id`, `name`, `title`, `icon`, `path`, `component`, `redirect`, `menu_type`, `permission`, `sort`, `parent_id`, `parent_ids`, `keep_alive`, `is_affix`, `hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (57, '', '插件管理', 'cascader', '/plugins', 'Layout', '/plugins/content/content-category', '1', '', 270, 0, '0,', '', '', '2', '2', 1, 1, '2023-03-07 10:37:37', '2024-11-21 15:55:39');
INSERT INTO `sys_menu` (`id`, `name`, `title`, `icon`, `path`, `component`, `redirect`, `menu_type`, `permission`, `sort`, `parent_id`, `parent_ids`, `keep_alive`, `is_affix`, `hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (58, '', 'App应用', 'app-group-fill', '/app', 'Layout', '/app/user/user', '1', '', 0, 0, '0,', '', '', '2', '2', 1, 1, '2023-03-08 09:27:36', '2024-11-21 16:05:44');
INSERT INTO `sys_menu` (`id`, `name`, `title`, `icon`, `path`, `component`, `redirect`, `menu_type`, `permission`, `sort`, `parent_id`, `parent_ids`, `keep_alive`, `is_affix`, `hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (59, '', '用户列表', 'chart', '/app/user', '/index', '/app/user/user', '1', '', 30, 58, '0,58,', '', '', '2', '2', 1, 1, '2023-03-09 14:24:25', '2024-11-21 15:53:06');
INSERT INTO `sys_menu` (`id`, `name`, `title`, `icon`, `path`, `component`, `redirect`, `menu_type`, `permission`, `sort`, `parent_id`, `parent_ids`, `keep_alive`, `is_affix`, `hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (60, '', '财务管理', 'eye-open', '/app/account', '/index', '/app/user/user-account-log', '1', '', 60, 58, '0,58,', '', '', '2', '2', 1, 1, '2023-03-09 21:13:23', '2024-11-21 15:53:56');
INSERT INTO `sys_menu` (`id`, `name`, `title`, `icon`, `path`, `component`, `redirect`, `menu_type`, `permission`, `sort`, `parent_id`, `parent_ids`, `keep_alive`, `is_affix`, `hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (61, 'UserLevel', '用户等级', 'pass', '/app/user/user-level', '/app/user/user-level/index.vue', NULL, '2', '', 60, 59, '0,58,59,', '2', '2', '2', '2', 1, 1, '2023-03-09 21:33:49', '2023-03-09 23:05:34');
INSERT INTO `sys_menu` (`id`, `name`, `title`, `icon`, `path`, `component`, `redirect`, `menu_type`, `permission`, `sort`, `parent_id`, `parent_ids`, `keep_alive`, `is_affix`, `hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (62, '', '分页获取用户等级', '', '', '', NULL, '3', 'app:user:user-level:query', 0, 61, '0,58,59,61,,', NULL, NULL, NULL, NULL, 1, 1, '2023-03-09 21:33:49', '2023-05-09 10:22:48');
INSERT INTO `sys_menu` (`id`, `name`, `title`, `icon`, `path`, `component`, `redirect`, `menu_type`, `permission`, `sort`, `parent_id`, `parent_ids`, `keep_alive`, `is_affix`, `hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (63, '', '创建用户等级', '', '', '', NULL, '3', 'app:user:user-level:add', 0, 61, '0,58,59,61,,', NULL, NULL, NULL, NULL, 1, 1, '2023-03-09 21:33:49', '2023-05-09 10:23:06');
INSERT INTO `sys_menu` (`id`, `name`, `title`, `icon`, `path`, `component`, `redirect`, `menu_type`, `permission`, `sort`, `parent_id`, `parent_ids`, `keep_alive`, `is_affix`, `hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (64, '', '修改用户等级', '', '', '', NULL, '3', 'app:user:user-level:edit', 0, 61, '0,58,59,61,,', NULL, NULL, NULL, NULL, 1, 1, '2023-03-09 21:33:49', '2023-05-09 10:23:27');
INSERT INTO `sys_menu` (`id`, `name`, `title`, `icon`, `path`, `component`, `redirect`, `menu_type`, `permission`, `sort`, `parent_id`, `parent_ids`, `keep_alive`, `is_affix`, `hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (65, '', '删除用户等级', '', '', '', NULL, '3', 'app:user:user-level:del', 0, 61, '0,58,59,61,,', NULL, NULL, NULL, NULL, 1, 1, '2023-03-09 21:33:49', '2023-05-09 10:23:44');
INSERT INTO `sys_menu` (`id`, `name`, `title`, `icon`, `path`, `component`, `redirect`, `menu_type`, `permission`, `sort`, `parent_id`, `parent_ids`, `keep_alive`, `is_affix`, `hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (66, '', '导出用户等级', '', '', '', NULL, '3', 'app:user:user-level:export', 0, 61, '0,58,59,61,,', NULL, NULL, NULL, NULL, 1, 1, '2023-03-09 21:33:49', '2023-05-09 10:23:54');
INSERT INTO `sys_menu` (`id`, `name`, `title`, `icon`, `path`, `component`, `redirect`, `menu_type`, `permission`, `sort`, `parent_id`, `parent_ids`, `keep_alive`, `is_affix`, `hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (67, 'UserAccountLog', '账变记录', 'pass', '/app/user/user-account-log', '/app/user/user-account-log/index.vue', NULL, '2', '', 0, 60, '0,58,60,', '2', '2', '2', '2', 1, 1, '2023-03-09 21:33:51', '2023-03-09 21:35:31');
INSERT INTO `sys_menu` (`id`, `name`, `title`, `icon`, `path`, `component`, `redirect`, `menu_type`, `permission`, `sort`, `parent_id`, `parent_ids`, `keep_alive`, `is_affix`, `hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (68, '', '分页获取账变记录', '', '', '', NULL, '3', 'app:user:user-account-log:query', 0, 67, '0,58,60,67,', NULL, NULL, NULL, NULL, 1, 1, '2023-03-09 21:33:51', '2023-05-09 10:32:28');
INSERT INTO `sys_menu` (`id`, `name`, `title`, `icon`, `path`, `component`, `redirect`, `menu_type`, `permission`, `sort`, `parent_id`, `parent_ids`, `keep_alive`, `is_affix`, `hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (69, '', '导出账变记录', '', '', '', NULL, '3', 'app:user:user-account-log:export', 0, 67, '0,58,60,67,', NULL, NULL, NULL, NULL, 1, 1, '2023-03-09 21:33:51', '2023-05-09 10:33:04');
INSERT INTO `sys_menu` (`id`, `name`, `title`, `icon`, `path`, `component`, `redirect`, `menu_type`, `permission`, `sort`, `parent_id`, `parent_ids`, `keep_alive`, `is_affix`, `hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (70, 'UserConf', '用户配置', 'pass', '/app/user/user-conf', '/app/user/user-conf/index.vue', NULL, '2', '', 90, 59, '0,58,59,', '2', '2', '2', '2', 1, 1, '2023-03-09 23:04:40', '2023-03-11 15:02:32');
INSERT INTO `sys_menu` (`id`, `name`, `title`, `icon`, `path`, `component`, `redirect`, `menu_type`, `permission`, `sort`, `parent_id`, `parent_ids`, `keep_alive`, `is_affix`, `hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (71, '', '分页获取用户配置', '', '', '', NULL, '3', 'app:user:user-conf:query', 0, 70, '0,58,59,70,', NULL, NULL, NULL, NULL, 1, 1, '2023-03-09 23:04:40', '2023-05-09 10:24:15');
INSERT INTO `sys_menu` (`id`, `name`, `title`, `icon`, `path`, `component`, `redirect`, `menu_type`, `permission`, `sort`, `parent_id`, `parent_ids`, `keep_alive`, `is_affix`, `hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (72, '', '修改用户配置', '', '', '', NULL, '3', 'app:user:user-conf:edit', 0, 70, '0,58,59,70,', NULL, NULL, NULL, NULL, 1, 1, '2023-03-09 23:04:40', '2023-05-09 10:25:16');
INSERT INTO `sys_menu` (`id`, `name`, `title`, `icon`, `path`, `component`, `redirect`, `menu_type`, `permission`, `sort`, `parent_id`, `parent_ids`, `keep_alive`, `is_affix`, `hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (73, 'User', '用户管理', 'pass', '/app/user/user', '/app/user/user/index.vue', NULL, '2', '', 30, 59, '0,58,59,', '2', '2', '2', '2', 1, 1, '2023-03-09 23:18:49', '2023-05-08 16:19:31');
INSERT INTO `sys_menu` (`id`, `name`, `title`, `icon`, `path`, `component`, `redirect`, `menu_type`, `permission`, `sort`, `parent_id`, `parent_ids`, `keep_alive`, `is_affix`, `hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (74, '', '分页获取用户管理', '', '', '', NULL, '3', 'app:user:user:query', 0, 73, '0,58,59,73,', NULL, NULL, NULL, NULL, 1, 1, '2023-03-09 23:18:49', '2023-05-08 22:30:33');
INSERT INTO `sys_menu` (`id`, `name`, `title`, `icon`, `path`, `component`, `redirect`, `menu_type`, `permission`, `sort`, `parent_id`, `parent_ids`, `keep_alive`, `is_affix`, `hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (75, '', '创建用户管理', '', '', '', NULL, '3', 'app:user:user:add', 0, 73, '0,58,59,73,', NULL, NULL, NULL, NULL, 1, 1, '2023-03-09 23:18:49', '2023-05-08 22:30:51');
INSERT INTO `sys_menu` (`id`, `name`, `title`, `icon`, `path`, `component`, `redirect`, `menu_type`, `permission`, `sort`, `parent_id`, `parent_ids`, `keep_alive`, `is_affix`, `hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (76, '', '修改用户管理', '', '', '', NULL, '3', 'app:user:user:edit', 0, 73, '0,58,59,73,', NULL, NULL, NULL, NULL, 1, 1, '2023-03-09 23:18:49', '2023-05-08 22:36:28');
INSERT INTO `sys_menu` (`id`, `name`, `title`, `icon`, `path`, `component`, `redirect`, `menu_type`, `permission`, `sort`, `parent_id`, `parent_ids`, `keep_alive`, `is_affix`, `hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (77, '', '导出用户管理', '', '', '', NULL, '3', 'app:user:user:export', 0, 73, '0,58,59,73,', NULL, NULL, NULL, NULL, 1, 1, '2023-03-09 23:18:49', '2023-05-08 22:37:31');
INSERT INTO `sys_menu` (`id`, `name`, `title`, `icon`, `path`, `component`, `redirect`, `menu_type`, `permission`, `sort`, `parent_id`, `parent_ids`, `keep_alive`, `is_affix`, `hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (78, 'UserOperLog', '用户行为记录', 'pass', '/app/user/user-oper-log', '/app/user/user-oper-log/index.vue', NULL, '2', '', 120, 59, '0,58,59,', '2', '2', '2', '2', 1, 1, '2023-03-11 15:00:06', '2023-03-11 15:02:42');
INSERT INTO `sys_menu` (`id`, `name`, `title`, `icon`, `path`, `component`, `redirect`, `menu_type`, `permission`, `sort`, `parent_id`, `parent_ids`, `keep_alive`, `is_affix`, `hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (79, '', '分页获取用户关键行为日志表', '', '', '', NULL, '3', 'app:user:user-oper-log:query', 0, 78, '0,58,59,78,', NULL, NULL, NULL, NULL, 1, 1, '2023-03-11 15:00:06', '2023-05-09 10:26:19');
INSERT INTO `sys_menu` (`id`, `name`, `title`, `icon`, `path`, `component`, `redirect`, `menu_type`, `permission`, `sort`, `parent_id`, `parent_ids`, `keep_alive`, `is_affix`, `hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (80, '', '导出用户关键行为日志表', '', '', '', NULL, '3', 'app:user:user-oper-log:export', 0, 78, '0,58,59,78,', NULL, NULL, NULL, NULL, 1, 1, '2023-03-11 15:00:06', '2023-05-09 10:28:23');
INSERT INTO `sys_menu` (`id`, `name`, `title`, `icon`, `path`, `component`, `redirect`, `menu_type`, `permission`, `sort`, `parent_id`, `parent_ids`, `keep_alive`, `is_affix`, `hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (81, '', '消息管理', 'batch-update', '/plugins/msg', '/index', '/plugins/msg/msg-code', '1', '', 120, 57, '0,57,', '', '', '2', '2', 1, 1, '2023-03-12 13:27:59', '2024-11-21 15:56:39');
INSERT INTO `sys_menu` (`id`, `name`, `title`, `icon`, `path`, `component`, `redirect`, `menu_type`, `permission`, `sort`, `parent_id`, `parent_ids`, `keep_alive`, `is_affix`, `hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (82, 'MsgCode', '验证码记录', 'pass', '/plugins/msg/msg-code', '/plugins/msg/msg-code/index.vue', NULL, '2', '', 0, 81, '0,57,81,', '2', '2', '2', '2', 1, 1, '2023-03-12 21:54:02', '2023-03-12 21:54:32');
INSERT INTO `sys_menu` (`id`, `name`, `title`, `icon`, `path`, `component`, `redirect`, `menu_type`, `permission`, `sort`, `parent_id`, `parent_ids`, `keep_alive`, `is_affix`, `hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (83, '', '分页获取验证码记录', '', '', '', NULL, '3', 'plugins:msg:msg-code:query', 0, 82, '0,57,81,82,', NULL, NULL, NULL, NULL, 1, 1, '2023-03-12 21:54:02', '2023-05-09 10:34:35');
INSERT INTO `sys_menu` (`id`, `name`, `title`, `icon`, `path`, `component`, `redirect`, `menu_type`, `permission`, `sort`, `parent_id`, `parent_ids`, `keep_alive`, `is_affix`, `hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (84, 'ContentAnnouncement', '公告管理', 'pass', '/plugins/content/content-announcement', '/plugins/content/content-announcement/index.vue', NULL, '2', '', 90, 56, '0,57,56,', '2', '2', '2', '2', 1, 1, '2023-03-12 22:47:11', '2023-03-12 22:48:08');
INSERT INTO `sys_menu` (`id`, `name`, `title`, `icon`, `path`, `component`, `redirect`, `menu_type`, `permission`, `sort`, `parent_id`, `parent_ids`, `keep_alive`, `is_affix`, `hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (85, '', '分页获取公告管理', '', '', '', NULL, '3', 'plugins:content:content-announcement:query', 0, 84, '0,57,56,84,', NULL, NULL, NULL, NULL, 1, 1, '2023-03-12 22:47:11', '2023-05-09 10:49:57');
INSERT INTO `sys_menu` (`id`, `name`, `title`, `icon`, `path`, `component`, `redirect`, `menu_type`, `permission`, `sort`, `parent_id`, `parent_ids`, `keep_alive`, `is_affix`, `hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (86, '', '创建公告管理', '', '', '', NULL, '3', 'plugins:content:content-announcement:add', 0, 84, '0,57,56,84,', NULL, NULL, NULL, NULL, 1, 1, '2023-03-12 22:47:11', '2023-05-09 10:50:13');
INSERT INTO `sys_menu` (`id`, `name`, `title`, `icon`, `path`, `component`, `redirect`, `menu_type`, `permission`, `sort`, `parent_id`, `parent_ids`, `keep_alive`, `is_affix`, `hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (87, '', '修改公告管理', '', '', '', NULL, '3', 'plugins:content:content-announcement:edit', 0, 84, '0,57,56,84,', NULL, NULL, NULL, NULL, 1, 1, '2023-03-12 22:47:11', '2023-05-09 10:50:39');
INSERT INTO `sys_menu` (`id`, `name`, `title`, `icon`, `path`, `component`, `redirect`, `menu_type`, `permission`, `sort`, `parent_id`, `parent_ids`, `keep_alive`, `is_affix`, `hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (88, '', '删除公告管理', '', '', '', NULL, '3', 'plugins:content:content-announcement:del', 0, 84, '0,57,56,84,', NULL, NULL, NULL, NULL, 1, 1, '2023-03-12 22:47:11', '2023-05-09 10:50:58');
INSERT INTO `sys_menu` (`id`, `name`, `title`, `icon`, `path`, `component`, `redirect`, `menu_type`, `permission`, `sort`, `parent_id`, `parent_ids`, `keep_alive`, `is_affix`, `hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (89, '', '导出公告管理', '', '', '', NULL, '3', 'plugins:content:content-announcement:export', 0, 84, '0,57,56,84,', NULL, NULL, NULL, NULL, 1, 1, '2023-03-12 22:47:11', '2023-05-09 10:51:08');
INSERT INTO `sys_menu` (`id`, `name`, `title`, `icon`, `path`, `component`, `redirect`, `menu_type`, `permission`, `sort`, `parent_id`, `parent_ids`, `keep_alive`, `is_affix`, `hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (90, 'ContentCategory', '内容分类', 'pass', '/plugins/content/content-category', '/plugins/content/content-category/index.vue', NULL, '2', '', 0, 56, '0,57,56,', '2', '2', '2', '2', 1, 1, '2023-03-12 23:17:44', '2023-03-12 23:20:35');
INSERT INTO `sys_menu` (`id`, `name`, `title`, `icon`, `path`, `component`, `redirect`, `menu_type`, `permission`, `sort`, `parent_id`, `parent_ids`, `keep_alive`, `is_affix`, `hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (91, '', '分页获取内容分类', '', '', '', NULL, '3', 'plugins:content:content-category:query', 0, 90, '0,57,56,90,', NULL, NULL, NULL, NULL, 1, 1, '2023-03-12 23:17:44', '2023-05-09 10:43:56');
INSERT INTO `sys_menu` (`id`, `name`, `title`, `icon`, `path`, `component`, `redirect`, `menu_type`, `permission`, `sort`, `parent_id`, `parent_ids`, `keep_alive`, `is_affix`, `hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (92, '', '创建内容分类', '', '', '', NULL, '3', 'plugins:content:content-category:add', 0, 90, '0,57,56,90,', NULL, NULL, NULL, NULL, 1, 1, '2023-03-12 23:17:44', '2023-05-09 10:44:14');
INSERT INTO `sys_menu` (`id`, `name`, `title`, `icon`, `path`, `component`, `redirect`, `menu_type`, `permission`, `sort`, `parent_id`, `parent_ids`, `keep_alive`, `is_affix`, `hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (93, '', '修改内容分类', '', '', '', NULL, '3', 'plugins:content:content-category:edit', 0, 90, '0,57,56,90,', NULL, NULL, NULL, NULL, 1, 1, '2023-03-12 23:17:44', '2023-05-09 10:44:33');
INSERT INTO `sys_menu` (`id`, `name`, `title`, `icon`, `path`, `component`, `redirect`, `menu_type`, `permission`, `sort`, `parent_id`, `parent_ids`, `keep_alive`, `is_affix`, `hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (94, '', '删除内容分类', '', '', '', NULL, '3', 'plugins:content:content-category:del', 0, 90, '0,57,56,90,', NULL, NULL, NULL, NULL, 1, 1, '2023-03-12 23:17:45', '2023-05-09 10:47:01');
INSERT INTO `sys_menu` (`id`, `name`, `title`, `icon`, `path`, `component`, `redirect`, `menu_type`, `permission`, `sort`, `parent_id`, `parent_ids`, `keep_alive`, `is_affix`, `hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (95, '', '导出内容分类', '', '', '', NULL, '3', 'plugins:content:content-category:export', 0, 90, '0,57,56,90,', NULL, NULL, NULL, NULL, 1, 1, '2023-03-12 23:17:45', '2023-05-09 10:47:20');
INSERT INTO `sys_menu` (`id`, `name`, `title`, `icon`, `path`, `component`, `redirect`, `menu_type`, `permission`, `sort`, `parent_id`, `parent_ids`, `keep_alive`, `is_affix`, `hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (96, 'ContentArticle', '文章管理', 'pass', '/plugins/content/content-article', '/plugins/content/content-article/index.vue', NULL, '2', '', 60, 56, '0,57,56,', '2', '2', '2', '2', 1, 1, '2023-03-12 23:52:45', '2023-03-12 23:53:12');
INSERT INTO `sys_menu` (`id`, `name`, `title`, `icon`, `path`, `component`, `redirect`, `menu_type`, `permission`, `sort`, `parent_id`, `parent_ids`, `keep_alive`, `is_affix`, `hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (97, '', '分页获取文章管理', '', '', '', NULL, '3', 'plugins:content:content-article:query', 0, 96, '0,57,56,96,', NULL, NULL, NULL, NULL, 1, 1, '2023-03-12 23:52:45', '2023-05-09 10:47:48');
INSERT INTO `sys_menu` (`id`, `name`, `title`, `icon`, `path`, `component`, `redirect`, `menu_type`, `permission`, `sort`, `parent_id`, `parent_ids`, `keep_alive`, `is_affix`, `hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (98, '', '创建文章管理', '', '', '', NULL, '3', 'plugins:content:content-article:add', 0, 96, '0,57,56,96,', NULL, NULL, NULL, NULL, 1, 1, '2023-03-12 23:52:45', '2023-05-09 10:48:03');
INSERT INTO `sys_menu` (`id`, `name`, `title`, `icon`, `path`, `component`, `redirect`, `menu_type`, `permission`, `sort`, `parent_id`, `parent_ids`, `keep_alive`, `is_affix`, `hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (99, '', '修改文章管理', '', '', '', NULL, '3', 'plugins:content:content-article:edit', 0, 96, '0,57,56,96,', NULL, NULL, NULL, NULL, 1, 1, '2023-03-12 23:52:45', '2023-05-09 10:48:27');
INSERT INTO `sys_menu` (`id`, `name`, `title`, `icon`, `path`, `component`, `redirect`, `menu_type`, `permission`, `sort`, `parent_id`, `parent_ids`, `keep_alive`, `is_affix`, `hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (100, '', '删除文章管理', '', '', '', NULL, '3', 'plugins:content:content-article:del', 0, 96, '0,57,56,96,', NULL, NULL, NULL, NULL, 1, 1, '2023-03-12 23:52:45', '2023-05-09 10:48:37');
INSERT INTO `sys_menu` (`id`, `name`, `title`, `icon`, `path`, `component`, `redirect`, `menu_type`, `permission`, `sort`, `parent_id`, `parent_ids`, `keep_alive`, `is_affix`, `hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (101, '', '导出文章管理', '', '', '', NULL, '3', 'plugins:content:content-article:export', 0, 96, '0,57,56,96,', NULL, NULL, NULL, NULL, 1, 1, '2023-03-12 23:52:46', '2023-05-09 10:48:53');
INSERT INTO `sys_menu` (`id`, `name`, `title`, `icon`, `path`, `component`, `redirect`, `menu_type`, `permission`, `sort`, `parent_id`, `parent_ids`, `keep_alive`, `is_affix`, `hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (102, 'FilemgrApp', 'App管理', 'pass', '/plugins/filemgr/filemgr-app', '/plugins/filemgr/filemgr-app/index.vue', NULL, '2', '', 0, 55, '0,57,55,', '2', '2', '2', '2', 1, 1, '2023-03-13 00:55:02', '2023-03-13 00:55:52');
INSERT INTO `sys_menu` (`id`, `name`, `title`, `icon`, `path`, `component`, `redirect`, `menu_type`, `permission`, `sort`, `parent_id`, `parent_ids`, `keep_alive`, `is_affix`, `hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (103, '', '分页获取App管理', '', '', '', NULL, '3', 'plugins:filemgr:filemgr-app:query', 0, 102, '0,57,55,102,', NULL, NULL, NULL, NULL, 1, 1, '2023-03-13 00:55:02', '2023-05-09 10:51:29');
INSERT INTO `sys_menu` (`id`, `name`, `title`, `icon`, `path`, `component`, `redirect`, `menu_type`, `permission`, `sort`, `parent_id`, `parent_ids`, `keep_alive`, `is_affix`, `hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (104, '', '创建App管理', '', '', '', NULL, '3', 'plugins:filemgr:filemgr-app:add', 0, 102, '0,57,55,102,', NULL, NULL, NULL, NULL, 1, 1, '2023-03-13 00:55:02', '2023-05-09 10:51:41');
INSERT INTO `sys_menu` (`id`, `name`, `title`, `icon`, `path`, `component`, `redirect`, `menu_type`, `permission`, `sort`, `parent_id`, `parent_ids`, `keep_alive`, `is_affix`, `hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (105, '', '修改App管理', '', '', '', NULL, '3', 'plugins:filemgr:filemgr-app:edit', 0, 102, '0,57,55,102,', NULL, NULL, NULL, NULL, 1, 1, '2023-03-13 00:55:02', '2023-05-09 10:51:56');
INSERT INTO `sys_menu` (`id`, `name`, `title`, `icon`, `path`, `component`, `redirect`, `menu_type`, `permission`, `sort`, `parent_id`, `parent_ids`, `keep_alive`, `is_affix`, `hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (106, '', '删除App管理', '', '', '', NULL, '3', 'plugins:filemgr:filemgr-app:del', 0, 102, '0,57,55,102,', NULL, NULL, NULL, NULL, 1, 1, '2023-03-13 00:55:02', '2023-05-09 10:52:06');
INSERT INTO `sys_menu` (`id`, `name`, `title`, `icon`, `path`, `component`, `redirect`, `menu_type`, `permission`, `sort`, `parent_id`, `parent_ids`, `keep_alive`, `is_affix`, `hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (107, '', '导出App管理', '', '', '', NULL, '3', 'plugins:filemgr:filemgr-app:export', 0, 102, '0,57,55,102,', NULL, NULL, NULL, NULL, 1, 1, '2023-03-13 00:55:02', '2023-05-09 10:52:17');
INSERT INTO `sys_menu` (`id`, `name`, `title`, `icon`, `path`, `component`, `redirect`, `menu_type`, `permission`, `sort`, `parent_id`, `parent_ids`, `keep_alive`, `is_affix`, `hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (108, 'UserCountryCode', '国家区号', 'pass', '/app/user/user-country-code', '/app/user/user-country-code/index.vue', NULL, '2', '', 150, 59, '0,58,59,', '2', '2', '2', '2', 1, 1, '2023-03-14 17:47:44', '2023-03-14 18:06:00');
INSERT INTO `sys_menu` (`id`, `name`, `title`, `icon`, `path`, `component`, `redirect`, `menu_type`, `permission`, `sort`, `parent_id`, `parent_ids`, `keep_alive`, `is_affix`, `hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (109, '', '分页获取国家电话区号', '', '', '', NULL, '3', 'app:user:user-country-code:query', 0, 108, '0,58,59,108,', NULL, NULL, NULL, NULL, 1, 1, '2023-03-14 17:47:44', '2023-05-09 10:31:00');
INSERT INTO `sys_menu` (`id`, `name`, `title`, `icon`, `path`, `component`, `redirect`, `menu_type`, `permission`, `sort`, `parent_id`, `parent_ids`, `keep_alive`, `is_affix`, `hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (110, '', '创建国家电话区号', '', '', '', NULL, '3', 'app:user:user-country-code:add', 0, 108, '0,58,59,108,', NULL, NULL, NULL, NULL, 1, 1, '2023-03-14 17:47:44', '2023-05-09 10:31:10');
INSERT INTO `sys_menu` (`id`, `name`, `title`, `icon`, `path`, `component`, `redirect`, `menu_type`, `permission`, `sort`, `parent_id`, `parent_ids`, `keep_alive`, `is_affix`, `hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (111, '', '修改国家电话区号', '', '', '', NULL, '3', 'app:user:user-country-code:edit', 0, 108, '0,58,59,108,', NULL, NULL, NULL, NULL, 1, 1, '2023-03-14 17:47:44', '2023-05-09 10:31:23');
INSERT INTO `sys_menu` (`id`, `name`, `title`, `icon`, `path`, `component`, `redirect`, `menu_type`, `permission`, `sort`, `parent_id`, `parent_ids`, `keep_alive`, `is_affix`, `hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (112, '', '删除国家电话区号', '', '', '', NULL, '3', 'app:user:user-country-code:del', 0, 108, '0,58,59,108,', NULL, NULL, NULL, NULL, 1, 1, '2023-03-14 17:47:45', '2023-05-09 10:31:33');
INSERT INTO `sys_menu` (`id`, `name`, `title`, `icon`, `path`, `component`, `redirect`, `menu_type`, `permission`, `sort`, `parent_id`, `parent_ids`, `keep_alive`, `is_affix`, `hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (113, '', '导出国家电话区号', '', '', '', NULL, '3', 'app:user:user-country-code:export', 0, 108, '0,58,59,108,', NULL, NULL, NULL, NULL, 1, 1, '2023-03-14 17:47:45', '2023-05-09 10:31:45');
INSERT INTO `sys_menu` (`id`, `name`, `title`, `icon`, `path`, `component`, `redirect`, `menu_type`, `permission`, `sort`, `parent_id`, `parent_ids`, `keep_alive`, `is_affix`, `hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (114, '', '导出操作日志', '', '', '', NULL, '3', 'admin:sysOperLog:export', 0, 15, '0,1,15,', NULL, NULL, NULL, NULL, 1, 1, '2023-05-09 11:02:50', '2023-05-09 11:07:07');
INSERT INTO `sys_menu` (`id`, `name`, `title`, `icon`, `path`, `component`, `redirect`, `menu_type`, `permission`, `sort`, `parent_id`, `parent_ids`, `keep_alive`, `is_affix`, `hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (115, '', '登录日志导出', '', '', '', NULL, '3', 'admin:sysLoginLog:export', 0, 14, '0,1,14,', NULL, NULL, NULL, NULL, 1, 1, '2023-05-09 11:04:20', '2023-05-09 11:04:20');
INSERT INTO `sys_menu` (`id`, `name`, `title`, `icon`, `path`, `component`, `redirect`, `menu_type`, `permission`, `sort`, `parent_id`, `parent_ids`, `keep_alive`, `is_affix`, `hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (116, '', '导出数据', 'app-group-fill', '', '', NULL, '3', 'admin:sysDictType:export', 0, 12, '0,1,12,', NULL, NULL, NULL, NULL, 1, 1, '2023-05-09 11:12:30', '2023-05-09 11:15:14');
INSERT INTO `sys_menu` (`id`, `name`, `title`, `icon`, `path`, `component`, `redirect`, `menu_type`, `permission`, `sort`, `parent_id`, `parent_ids`, `keep_alive`, `is_affix`, `hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (117, '', '导出类型', 'app-group-fill', '', '', NULL, '3', 'admin:sysDictType:export', 0, 11, '0,1,11,', NULL, NULL, NULL, NULL, 1, 1, '2023-05-09 11:16:13', '2023-05-09 11:16:42');
INSERT INTO `sys_menu` (`id`, `name`, `title`, `icon`, `path`, `component`, `redirect`, `menu_type`, `permission`, `sort`, `parent_id`, `parent_ids`, `keep_alive`, `is_affix`, `hidden`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (118, '', '导出参数', '', '', '', NULL, '3', 'content:sysConfig:export', 0, 13, '0,1,13,', NULL, NULL, NULL, NULL, 1, 1, '2023-05-09 11:34:20', '2023-05-09 11:34:38');
COMMIT;

-- ----------------------------
-- Table structure for sys_menu_api_rule
-- ----------------------------
DROP TABLE IF EXISTS `sys_menu_api_rule`;
CREATE TABLE `sys_menu_api_rule` (
  `sys_menu_menu_id` int NOT NULL,
  `sys_api_id` int NOT NULL COMMENT '主键编码',
  PRIMARY KEY (`sys_menu_menu_id`,`sys_api_id`),
  KEY `fk_sys_menu_api_rule_sys_api` (`sys_api_id`),
  CONSTRAINT `fk_sys_menu_api_rule_sys_api` FOREIGN KEY (`sys_api_id`) REFERENCES `sys_api` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
  CONSTRAINT `fk_sys_menu_api_rule_sys_menu` FOREIGN KEY (`sys_menu_menu_id`) REFERENCES `sys_menu` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;

-- ----------------------------
-- Records of sys_menu_api_rule
-- ----------------------------
BEGIN;
INSERT INTO `sys_menu_api_rule` (`sys_menu_menu_id`, `sys_api_id`) VALUES (48, 2);
INSERT INTO `sys_menu_api_rule` (`sys_menu_menu_id`, `sys_api_id`) VALUES (48, 4);
INSERT INTO `sys_menu_api_rule` (`sys_menu_menu_id`, `sys_api_id`) VALUES (48, 5);
INSERT INTO `sys_menu_api_rule` (`sys_menu_menu_id`, `sys_api_id`) VALUES (48, 6);
INSERT INTO `sys_menu_api_rule` (`sys_menu_menu_id`, `sys_api_id`) VALUES (44, 19);
INSERT INTO `sys_menu_api_rule` (`sys_menu_menu_id`, `sys_api_id`) VALUES (46, 22);
INSERT INTO `sys_menu_api_rule` (`sys_menu_menu_id`, `sys_api_id`) VALUES (74, 30);
INSERT INTO `sys_menu_api_rule` (`sys_menu_menu_id`, `sys_api_id`) VALUES (109, 31);
INSERT INTO `sys_menu_api_rule` (`sys_menu_menu_id`, `sys_api_id`) VALUES (113, 32);
INSERT INTO `sys_menu_api_rule` (`sys_menu_menu_id`, `sys_api_id`) VALUES (111, 33);
INSERT INTO `sys_menu_api_rule` (`sys_menu_menu_id`, `sys_api_id`) VALUES (71, 34);
INSERT INTO `sys_menu_api_rule` (`sys_menu_menu_id`, `sys_api_id`) VALUES (72, 35);
INSERT INTO `sys_menu_api_rule` (`sys_menu_menu_id`, `sys_api_id`) VALUES (68, 36);
INSERT INTO `sys_menu_api_rule` (`sys_menu_menu_id`, `sys_api_id`) VALUES (69, 37);
INSERT INTO `sys_menu_api_rule` (`sys_menu_menu_id`, `sys_api_id`) VALUES (62, 39);
INSERT INTO `sys_menu_api_rule` (`sys_menu_menu_id`, `sys_api_id`) VALUES (66, 40);
INSERT INTO `sys_menu_api_rule` (`sys_menu_menu_id`, `sys_api_id`) VALUES (64, 41);
INSERT INTO `sys_menu_api_rule` (`sys_menu_menu_id`, `sys_api_id`) VALUES (79, 42);
INSERT INTO `sys_menu_api_rule` (`sys_menu_menu_id`, `sys_api_id`) VALUES (80, 43);
INSERT INTO `sys_menu_api_rule` (`sys_menu_menu_id`, `sys_api_id`) VALUES (77, 45);
INSERT INTO `sys_menu_api_rule` (`sys_menu_menu_id`, `sys_api_id`) VALUES (76, 46);
INSERT INTO `sys_menu_api_rule` (`sys_menu_menu_id`, `sys_api_id`) VALUES (85, 48);
INSERT INTO `sys_menu_api_rule` (`sys_menu_menu_id`, `sys_api_id`) VALUES (89, 49);
INSERT INTO `sys_menu_api_rule` (`sys_menu_menu_id`, `sys_api_id`) VALUES (107, 49);
INSERT INTO `sys_menu_api_rule` (`sys_menu_menu_id`, `sys_api_id`) VALUES (87, 50);
INSERT INTO `sys_menu_api_rule` (`sys_menu_menu_id`, `sys_api_id`) VALUES (105, 50);
INSERT INTO `sys_menu_api_rule` (`sys_menu_menu_id`, `sys_api_id`) VALUES (97, 51);
INSERT INTO `sys_menu_api_rule` (`sys_menu_menu_id`, `sys_api_id`) VALUES (101, 52);
INSERT INTO `sys_menu_api_rule` (`sys_menu_menu_id`, `sys_api_id`) VALUES (99, 53);
INSERT INTO `sys_menu_api_rule` (`sys_menu_menu_id`, `sys_api_id`) VALUES (91, 54);
INSERT INTO `sys_menu_api_rule` (`sys_menu_menu_id`, `sys_api_id`) VALUES (95, 55);
INSERT INTO `sys_menu_api_rule` (`sys_menu_menu_id`, `sys_api_id`) VALUES (93, 56);
INSERT INTO `sys_menu_api_rule` (`sys_menu_menu_id`, `sys_api_id`) VALUES (103, 57);
INSERT INTO `sys_menu_api_rule` (`sys_menu_menu_id`, `sys_api_id`) VALUES (83, 60);
INSERT INTO `sys_menu_api_rule` (`sys_menu_menu_id`, `sys_api_id`) VALUES (32, 62);
INSERT INTO `sys_menu_api_rule` (`sys_menu_menu_id`, `sys_api_id`) VALUES (114, 63);
INSERT INTO `sys_menu_api_rule` (`sys_menu_menu_id`, `sys_api_id`) VALUES (117, 63);
INSERT INTO `sys_menu_api_rule` (`sys_menu_menu_id`, `sys_api_id`) VALUES (34, 65);
INSERT INTO `sys_menu_api_rule` (`sys_menu_menu_id`, `sys_api_id`) VALUES (36, 66);
INSERT INTO `sys_menu_api_rule` (`sys_menu_menu_id`, `sys_api_id`) VALUES (114, 67);
INSERT INTO `sys_menu_api_rule` (`sys_menu_menu_id`, `sys_api_id`) VALUES (116, 67);
INSERT INTO `sys_menu_api_rule` (`sys_menu_menu_id`, `sys_api_id`) VALUES (38, 68);
INSERT INTO `sys_menu_api_rule` (`sys_menu_menu_id`, `sys_api_id`) VALUES (49, 69);
INSERT INTO `sys_menu_api_rule` (`sys_menu_menu_id`, `sys_api_id`) VALUES (40, 70);
INSERT INTO `sys_menu_api_rule` (`sys_menu_menu_id`, `sys_api_id`) VALUES (118, 71);
INSERT INTO `sys_menu_api_rule` (`sys_menu_menu_id`, `sys_api_id`) VALUES (42, 72);
INSERT INTO `sys_menu_api_rule` (`sys_menu_menu_id`, `sys_api_id`) VALUES (50, 78);
INSERT INTO `sys_menu_api_rule` (`sys_menu_menu_id`, `sys_api_id`) VALUES (48, 89);
INSERT INTO `sys_menu_api_rule` (`sys_menu_menu_id`, `sys_api_id`) VALUES (86, 94);
INSERT INTO `sys_menu_api_rule` (`sys_menu_menu_id`, `sys_api_id`) VALUES (104, 94);
INSERT INTO `sys_menu_api_rule` (`sys_menu_menu_id`, `sys_api_id`) VALUES (98, 95);
INSERT INTO `sys_menu_api_rule` (`sys_menu_menu_id`, `sys_api_id`) VALUES (92, 96);
INSERT INTO `sys_menu_api_rule` (`sys_menu_menu_id`, `sys_api_id`) VALUES (75, 99);
INSERT INTO `sys_menu_api_rule` (`sys_menu_menu_id`, `sys_api_id`) VALUES (110, 100);
INSERT INTO `sys_menu_api_rule` (`sys_menu_menu_id`, `sys_api_id`) VALUES (63, 101);
INSERT INTO `sys_menu_api_rule` (`sys_menu_menu_id`, `sys_api_id`) VALUES (37, 102);
INSERT INTO `sys_menu_api_rule` (`sys_menu_menu_id`, `sys_api_id`) VALUES (33, 103);
INSERT INTO `sys_menu_api_rule` (`sys_menu_menu_id`, `sys_api_id`) VALUES (41, 105);
INSERT INTO `sys_menu_api_rule` (`sys_menu_menu_id`, `sys_api_id`) VALUES (48, 117);
INSERT INTO `sys_menu_api_rule` (`sys_menu_menu_id`, `sys_api_id`) VALUES (49, 117);
INSERT INTO `sys_menu_api_rule` (`sys_menu_menu_id`, `sys_api_id`) VALUES (72, 119);
INSERT INTO `sys_menu_api_rule` (`sys_menu_menu_id`, `sys_api_id`) VALUES (111, 120);
INSERT INTO `sys_menu_api_rule` (`sys_menu_menu_id`, `sys_api_id`) VALUES (64, 121);
INSERT INTO `sys_menu_api_rule` (`sys_menu_menu_id`, `sys_api_id`) VALUES (76, 122);
INSERT INTO `sys_menu_api_rule` (`sys_menu_menu_id`, `sys_api_id`) VALUES (87, 123);
INSERT INTO `sys_menu_api_rule` (`sys_menu_menu_id`, `sys_api_id`) VALUES (105, 123);
INSERT INTO `sys_menu_api_rule` (`sys_menu_menu_id`, `sys_api_id`) VALUES (99, 124);
INSERT INTO `sys_menu_api_rule` (`sys_menu_menu_id`, `sys_api_id`) VALUES (93, 125);
INSERT INTO `sys_menu_api_rule` (`sys_menu_menu_id`, `sys_api_id`) VALUES (38, 130);
INSERT INTO `sys_menu_api_rule` (`sys_menu_menu_id`, `sys_api_id`) VALUES (34, 131);
INSERT INTO `sys_menu_api_rule` (`sys_menu_menu_id`, `sys_api_id`) VALUES (42, 132);
INSERT INTO `sys_menu_api_rule` (`sys_menu_menu_id`, `sys_api_id`) VALUES (45, 135);
INSERT INTO `sys_menu_api_rule` (`sys_menu_menu_id`, `sys_api_id`) VALUES (47, 136);
INSERT INTO `sys_menu_api_rule` (`sys_menu_menu_id`, `sys_api_id`) VALUES (48, 139);
INSERT INTO `sys_menu_api_rule` (`sys_menu_menu_id`, `sys_api_id`) VALUES (88, 141);
INSERT INTO `sys_menu_api_rule` (`sys_menu_menu_id`, `sys_api_id`) VALUES (106, 141);
INSERT INTO `sys_menu_api_rule` (`sys_menu_menu_id`, `sys_api_id`) VALUES (100, 142);
INSERT INTO `sys_menu_api_rule` (`sys_menu_menu_id`, `sys_api_id`) VALUES (94, 143);
INSERT INTO `sys_menu_api_rule` (`sys_menu_menu_id`, `sys_api_id`) VALUES (112, 145);
INSERT INTO `sys_menu_api_rule` (`sys_menu_menu_id`, `sys_api_id`) VALUES (65, 146);
INSERT INTO `sys_menu_api_rule` (`sys_menu_menu_id`, `sys_api_id`) VALUES (39, 147);
INSERT INTO `sys_menu_api_rule` (`sys_menu_menu_id`, `sys_api_id`) VALUES (35, 148);
INSERT INTO `sys_menu_api_rule` (`sys_menu_menu_id`, `sys_api_id`) VALUES (43, 149);
COMMIT;

-- ----------------------------
-- Table structure for sys_oper_log
-- ----------------------------
DROP TABLE IF EXISTS `sys_oper_log`;
CREATE TABLE `sys_oper_log` (
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
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='操作日志';

-- ----------------------------
-- Records of sys_oper_log
-- ----------------------------
BEGIN;
INSERT INTO `sys_oper_log` (`id`, `request_method`, `user_id`, `oper_url`, `oper_ip`, `oper_location`, `status`, `oper_time`, `json_result`, `remark`, `latency_time`, `user_agent`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (2, 'GET', 1, '/admin-api/v1/menurole', '127.0.0.1', '内部IP', '200', '2023-03-01 16:13:48', '{\"requestId\":\"9795559d-2619-48ac-854b-d3c826e69c5c\",\"code\":200,\"data\":[{\"menuId\":778,\"name\":\"content', '', '347.447748ms', 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.0.0 Safari/537.36 Edg/110.0.1587.50', '2023-03-01 16:13:48', '2023-03-01 16:13:48', 0, 0);
INSERT INTO `sys_oper_log` (`id`, `request_method`, `user_id`, `oper_url`, `oper_ip`, `oper_location`, `status`, `oper_time`, `json_result`, `remark`, `latency_time`, `user_agent`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (3, 'GET', 1, '/admin-api/v1/dict-data/option-select?dictType=sys_config_type', '127.0.0.1', '内部IP', '200', '2023-03-01 16:13:48', '{\"requestId\":\"de938987-da4e-4101-be07-6886e6a1bccd\",\"code\":200,\"msg\":\"操作成功\",\"data\":[{\"id\":18', '', '54.264428ms', 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.0.0 Safari/537.36 Edg/110.0.1587.50', '2023-03-01 16:13:48', '2023-03-01 16:13:48', 0, 0);
INSERT INTO `sys_oper_log` (`id`, `request_method`, `user_id`, `oper_url`, `oper_ip`, `oper_location`, `status`, `oper_time`, `json_result`, `remark`, `latency_time`, `user_agent`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (4, 'GET', 1, '/admin-api/v1/config?pageIndex=1&pageSize=10', '127.0.0.1', '内部IP', '200', '2023-03-01 16:13:48', '{\"requestId\":\"968f9eaf-23c5-48d2-9ece-e568ceece919\",\"code\":200,\"msg\":\"操作成功\",\"data\":{\"count\":', '', '54.989684ms', 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.0.0 Safari/537.36 Edg/110.0.1587.50', '2023-03-01 16:13:48', '2023-03-01 16:13:48', 0, 0);
COMMIT;

-- ----------------------------
-- Table structure for sys_post
-- ----------------------------
DROP TABLE IF EXISTS `sys_post`;
CREATE TABLE `sys_post` (
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
-- Records of sys_post
-- ----------------------------
BEGIN;
INSERT INTO `sys_post` (`id`, `post_name`, `post_code`, `sort`, `status`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (1, '首席执行官', 'CEO', 0, '1', '首席执行官', 1, 1, '2021-05-13 19:56:38', '2021-05-13 19:56:38');
INSERT INTO `sys_post` (`id`, `post_name`, `post_code`, `sort`, `status`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (2, '首席技术执行官', 'CTO', 2, '1', '首席技术执行官', 1, 1, '2021-05-13 19:56:38', '2021-05-13 19:56:38');
INSERT INTO `sys_post` (`id`, `post_name`, `post_code`, `sort`, `status`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (3, '首席运营官', 'COO', 3, '1', '测试工程师', 1, 1, '2021-05-13 19:56:38', '2021-05-13 19:56:38');
COMMIT;

-- ----------------------------
-- Table structure for sys_role
-- ----------------------------
DROP TABLE IF EXISTS `sys_role`;
CREATE TABLE `sys_role` (
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
-- Records of sys_role
-- ----------------------------
BEGIN;
INSERT INTO `sys_role` (`id`, `role_name`, `role_key`, `role_sort`, `remark`, `data_scope`, `status`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (1, '系统管理员', 'admin', 1, '', '', '1', 1, 1, '2021-05-13 19:56:37.913', '2023-03-03 01:04:03.641');
INSERT INTO `sys_role` (`id`, `role_name`, `role_key`, `role_sort`, `remark`, `data_scope`, `status`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (2, 'test', 'test', 0, '', '', '1', 1, 1, '2023-04-27 14:33:47.437', '2023-04-27 16:40:14.207');
COMMIT;

-- ----------------------------
-- Table structure for sys_role_dept
-- ----------------------------
DROP TABLE IF EXISTS `sys_role_dept`;
CREATE TABLE `sys_role_dept` (
  `role_id` smallint NOT NULL,
  `dept_id` smallint NOT NULL,
  PRIMARY KEY (`role_id`,`dept_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;

-- ----------------------------
-- Records of sys_role_dept
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for sys_role_menu
-- ----------------------------
DROP TABLE IF EXISTS `sys_role_menu`;
CREATE TABLE `sys_role_menu` (
  `role_id` int NOT NULL,
  `menu_id` int NOT NULL,
  PRIMARY KEY (`role_id`,`menu_id`),
  KEY `fk_sys_role_menu_sys_menu` (`menu_id`),
  CONSTRAINT `fk_sys_role_menu_sys_menu` FOREIGN KEY (`menu_id`) REFERENCES `sys_menu` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
  CONSTRAINT `fk_sys_role_menu_sys_role` FOREIGN KEY (`role_id`) REFERENCES `sys_role` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;

-- ----------------------------
-- Records of sys_role_menu
-- ----------------------------
BEGIN;
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`) VALUES (2, 43);
COMMIT;

-- ----------------------------
-- Table structure for sys_user
-- ----------------------------
DROP TABLE IF EXISTS `sys_user`;
CREATE TABLE `sys_user` (
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
-- Records of sys_user
-- ----------------------------
BEGIN;
INSERT INTO `sys_user` (`id`, `username`, `password`, `nick_name`, `phone`, `role_id`, `salt`, `avatar`, `sex`, `email`, `dept_id`, `post_id`, `remark`, `status`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (1, 'admin', '$2a$10$ZFMtvc.ROtYwk2UNOaBLCOrpr.Mq/i1ae4PVZfoWgHTb4ffORW/lm', 'admin', '13700000000', 1, '', 'http://www.bitxx.top/images/my_head-touch-icon-next.png', '1', 'admin@admin.com', 1, 1, '', '1', 1, 1, '2021-05-13 19:56:38', '2023-03-14 09:27:36');
INSERT INTO `sys_user` (`id`, `username`, `password`, `nick_name`, `phone`, `role_id`, `salt`, `avatar`, `sex`, `email`, `dept_id`, `post_id`, `remark`, `status`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (2, 'test', '$2a$10$IZufMe1mAFjoghFGrsrz9OEtXEmLmvlaJfcNCWe7UieboiMb9Sl72', 'test', '13711111111', 2, '', 'http://www.bitxx.top/images/my_head-touch-icon-next.png', '1', '13711111111@qq.com', 6, 1, '', '1', 1, 1, '2023-04-27 14:34:57', '2023-04-27 14:34:57');
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
