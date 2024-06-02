/*
 Navicat Premium Data Transfer

 Source Server         : dev
 Source Server Type    : MySQL
 Source Server Version : 80200
 Source Host           : localhost:3306
 Source Schema         : manager_auto

 Target Server Type    : MySQL
 Target Server Version : 80200
 File Encoding         : 65001

 Date: 02/06/2024 16:41:28
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for casbin_rule
-- ----------------------------
DROP TABLE IF EXISTS `casbin_rule`;
CREATE TABLE `casbin_rule` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `ptype` varchar(100) DEFAULT NULL,
  `v0` varchar(100) DEFAULT NULL,
  `v1` varchar(100) DEFAULT NULL,
  `v2` varchar(100) DEFAULT NULL,
  `v3` varchar(100) DEFAULT NULL,
  `v4` varchar(100) DEFAULT NULL,
  `v5` varchar(100) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_casbin_rule` (`ptype`,`v0`,`v1`,`v2`,`v3`,`v4`,`v5`)
) ENGINE=InnoDB AUTO_INCREMENT=591 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Records of casbin_rule
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for department
-- ----------------------------
DROP TABLE IF EXISTS `department`;
CREATE TABLE `department` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `parent_id` bigint unsigned NOT NULL COMMENT '父id',
  `keyword` char(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT '标识',
  `name` varchar(64) NOT NULL COMMENT '名称',
  `description` varchar(256) NOT NULL COMMENT '描述',
  `created_at` bigint unsigned NOT NULL DEFAULT '0' COMMENT '创建时间',
  `updated_at` bigint unsigned NOT NULL DEFAULT '0' COMMENT '修改时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `keyword` (`keyword`),
  KEY `idx_department_created_at` (`created_at`),
  KEY `idx_department_updated_at` (`updated_at`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COMMENT='部门信息';

-- ----------------------------
-- Records of department
-- ----------------------------
BEGIN;
INSERT INTO `department` VALUES (1, 0, 'company', '贵州青橙科技有限公司', '开放合作，拥抱未来', 1713706137, 1713706137);
COMMIT;

-- ----------------------------
-- Table structure for department_closure
-- ----------------------------
DROP TABLE IF EXISTS `department_closure`;
CREATE TABLE `department_closure` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `parent` bigint unsigned NOT NULL COMMENT '部门id',
  `children` bigint unsigned NOT NULL COMMENT '部门id',
  PRIMARY KEY (`id`),
  KEY `parent` (`parent`),
  KEY `children` (`children`),
  CONSTRAINT `department_closure_ibfk_1` FOREIGN KEY (`children`) REFERENCES `department` (`id`) ON DELETE CASCADE,
  CONSTRAINT `department_closure_ibfk_2` FOREIGN KEY (`parent`) REFERENCES `department` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=10 DEFAULT CHARSET=utf8mb4 COMMENT='部门层级信息';

-- ----------------------------
-- Records of department_closure
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for dictionary
-- ----------------------------
DROP TABLE IF EXISTS `dictionary`;
CREATE TABLE `dictionary` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `keyword` char(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT '关键字',
  `name` varchar(64) NOT NULL COMMENT '名称',
  `description` varchar(256) DEFAULT NULL COMMENT '描述',
  `created_at` bigint unsigned NOT NULL DEFAULT '0' COMMENT '创建时间',
  `updated_at` bigint unsigned NOT NULL DEFAULT '0' COMMENT '修改时间',
  `deleted_at` bigint unsigned NOT NULL DEFAULT '0' COMMENT '删除时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `keyword` (`keyword`,`deleted_at`),
  KEY `idx_dictionary_created_at` (`created_at`),
  KEY `idx_dictionary_updated_at` (`updated_at`),
  KEY `idx_dictionary_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COMMENT='字典信息';

-- ----------------------------
-- Records of dictionary
-- ----------------------------
BEGIN;
INSERT INTO `dictionary` VALUES (1, '2', '2', '2', 1717085576, 1717085576, 0);
COMMIT;

-- ----------------------------
-- Table structure for dictionary_value
-- ----------------------------
DROP TABLE IF EXISTS `dictionary_value`;
CREATE TABLE `dictionary_value` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `dictionary_id` bigint unsigned NOT NULL COMMENT '字典id',
  `label` varchar(128) NOT NULL COMMENT '标签',
  `value` varchar(128) NOT NULL COMMENT '标识',
  `status` tinyint(1) DEFAULT '1' COMMENT '状态',
  `weight` int unsigned DEFAULT '0' COMMENT '权重',
  `type` char(32) DEFAULT NULL COMMENT '类型',
  `extra` varchar(512) DEFAULT NULL COMMENT '扩展信息',
  `description` varchar(256) DEFAULT NULL COMMENT '描述',
  `created_at` bigint unsigned NOT NULL DEFAULT '0' COMMENT '创建时间',
  `updated_at` bigint unsigned NOT NULL DEFAULT '0' COMMENT '修改时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `value` (`dictionary_id`,`value`),
  KEY `idx_dictionary_value_created_at` (`created_at`),
  KEY `idx_dictionary_value_updated_at` (`updated_at`),
  KEY `idx_dictionary_value_weight` (`weight`),
  CONSTRAINT `fk_dictionary_value_dict` FOREIGN KEY (`dictionary_id`) REFERENCES `dictionary` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COMMENT='字典值信息';

-- ----------------------------
-- Records of dictionary_value
-- ----------------------------
BEGIN;
INSERT INTO `dictionary_value` VALUES (1, 1, '2', '3', 1, 0, NULL, '1', NULL, 1717086000, 1717308808);
COMMIT;

-- ----------------------------
-- Table structure for job
-- ----------------------------
DROP TABLE IF EXISTS `job`;
CREATE TABLE `job` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `keyword` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT '关键字',
  `name` varchar(64) NOT NULL COMMENT '名称',
  `weight` int unsigned DEFAULT NULL COMMENT '权重',
  `description` varchar(256) NOT NULL COMMENT '描述',
  `created_at` bigint unsigned NOT NULL DEFAULT '0' COMMENT '创建时间',
  `updated_at` bigint unsigned NOT NULL DEFAULT '0' COMMENT '修改时间',
  `deleted_at` bigint unsigned NOT NULL DEFAULT '0' COMMENT '删除时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `keyword` (`keyword`,`deleted_at`),
  KEY `idx_job_weight` (`weight`),
  KEY `idx_job_updated_at` (`updated_at`),
  KEY `idx_job_created_at` (`created_at`),
  KEY `idx_job_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=13 DEFAULT CHARSET=utf8mb4 COMMENT='职位信息';

-- ----------------------------
-- Records of job
-- ----------------------------
BEGIN;
INSERT INTO `job` VALUES (1, 'chairman', '董事长', 0, '董事长', 1713706137, 1717211765, 0);
COMMIT;

-- ----------------------------
-- Table structure for menu
-- ----------------------------
DROP TABLE IF EXISTS `menu`;
CREATE TABLE `menu` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `parent_id` bigint unsigned NOT NULL COMMENT '父id',
  `title` varchar(128) NOT NULL COMMENT '标题',
  `type` char(32) NOT NULL COMMENT '类型',
  `keyword` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '关键词',
  `icon` char(32) DEFAULT NULL COMMENT '图标',
  `api` varchar(128) DEFAULT NULL COMMENT '接口',
  `method` varchar(12) DEFAULT NULL COMMENT '接口方法',
  `path` varchar(128) DEFAULT NULL COMMENT '路径',
  `permission` varchar(128) DEFAULT NULL COMMENT '指令',
  `component` varchar(128) DEFAULT NULL COMMENT '组件',
  `redirect` varchar(128) DEFAULT NULL COMMENT '重定向地址',
  `weight` int unsigned DEFAULT '0' COMMENT '权重',
  `is_hidden` tinyint(1) DEFAULT NULL COMMENT '是否隐藏',
  `is_cache` tinyint(1) DEFAULT NULL COMMENT '是否缓存',
  `is_home` tinyint(1) DEFAULT NULL COMMENT '是否为首页',
  `is_affix` tinyint(1) DEFAULT NULL COMMENT '是否为标签',
  `created_at` bigint unsigned NOT NULL DEFAULT '0' COMMENT '创建时间',
  `updated_at` bigint unsigned NOT NULL DEFAULT '0' COMMENT '修改时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `keyword` (`keyword`),
  UNIQUE KEY `path` (`path`),
  UNIQUE KEY `api_method` (`api`,`method`),
  KEY `idx_menu_created_at` (`created_at`),
  KEY `idx_menu_updated_at` (`updated_at`),
  KEY `idx_menu_weight` (`weight`)
) ENGINE=InnoDB AUTO_INCREMENT=681 DEFAULT CHARSET=utf8mb4 COMMENT='菜单信息';

-- ----------------------------
-- Records of menu
-- ----------------------------
BEGIN;
INSERT INTO `menu` VALUES (624, 0, '管理平台', 'R', 'SystemPlatform', 'apps', NULL, NULL, '/', NULL, 'Layout', NULL, 2, 0, NULL, NULL, NULL, 1717305025, 1717305025);
INSERT INTO `menu` VALUES (625, 624, '首页面板', 'M', 'Dashboard', 'dashboard', NULL, NULL, '/dashboard', NULL, '/dashboard/index', NULL, 0, 0, 1, 1, 1, 1717305025, 1717305025);
INSERT INTO `menu` VALUES (626, 624, '管理中心', 'M', 'SystemPlatformManager', 'apps', NULL, NULL, '/manager', NULL, NULL, NULL, 0, 0, NULL, NULL, NULL, 1717305025, 1717305025);
INSERT INTO `menu` VALUES (627, 626, '基础接口', 'G', 'ManagerBaseApi', 'apps', NULL, NULL, NULL, NULL, NULL, NULL, 0, 1, NULL, NULL, NULL, 1717305025, 1717305025);
INSERT INTO `menu` VALUES (628, 627, '获取用户可见部门树', 'BA', NULL, NULL, '/manager/api/v1/departments', 'GET', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL, 1717305025, 1717305025);
INSERT INTO `menu` VALUES (629, 627, '获取用户可见角色树', 'BA', NULL, NULL, '/manager/api/v1/roles', 'GET', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL, 1717305025, 1717305025);
INSERT INTO `menu` VALUES (630, 627, '获取个人用户信息', 'BA', NULL, NULL, '/manager/api/v1/user/current', 'GET', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL, 1717305025, 1717305025);
INSERT INTO `menu` VALUES (631, 627, '更新个人用户信息', 'BA', NULL, NULL, '/manager/api/v1/user/current/info', 'PUT', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL, 1717305025, 1717305025);
INSERT INTO `menu` VALUES (632, 627, '更新个人用户密码', 'BA', NULL, NULL, '/manager/api/v1/user/current/password', 'PUT', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL, 1717305025, 1717305025);
INSERT INTO `menu` VALUES (633, 627, '保存个人用户设置', 'BA', NULL, NULL, '/manager/api/v1/user/current/setting', 'PUT', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL, 1717305025, 1717305025);
INSERT INTO `menu` VALUES (634, 627, '发送用户验证吗', 'BA', NULL, NULL, '/manager/api/v1/user/current/captcha', 'POST', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL, 1717305025, 1717305025);
INSERT INTO `menu` VALUES (635, 627, '获取用户当前角色菜单', 'BA', NULL, NULL, '/manager/api/v1/menus/by/cur_role', 'GET', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL, 1717305025, 1717305025);
INSERT INTO `menu` VALUES (636, 627, '退出登录', 'BA', NULL, NULL, '/manager/api/v1/user/logout', 'POST', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL, 1717305025, 1717305025);
INSERT INTO `menu` VALUES (637, 627, '刷新token', 'BA', NULL, NULL, '/manager/api/v1/user/token/refresh', 'POST', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL, 1717305025, 1717305025);
INSERT INTO `menu` VALUES (638, 627, '用户登录', 'BA', NULL, NULL, '/manager/api/v1/user/login', 'POST', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL, 1717305025, 1717305025);
INSERT INTO `menu` VALUES (639, 627, '获取登录验证码', 'BA', NULL, NULL, '/manager/api/v1/user/login/captcha', 'POST', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL, 1717305025, 1717305025);
INSERT INTO `menu` VALUES (640, 627, '获取系统基础设置', 'BA', NULL, NULL, '/manager/api/v1/system/setting', 'GET', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL, 1717305025, 1717305025);
INSERT INTO `menu` VALUES (641, 627, '接口鉴权', 'BA', NULL, NULL, '/manager/api/v1/auth', 'POST', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL, 1717305025, 1717305025);
INSERT INTO `menu` VALUES (642, 627, '切换用户角色', 'BA', NULL, NULL, '/manager/api/v1/user/current/role', 'POST', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL, 1717305025, 1717305025);
INSERT INTO `menu` VALUES (643, 626, '字典管理', 'M', 'ManagerDictionary', 'storage', NULL, NULL, '/manager/dictionary', NULL, '/manager/dictionary/index', NULL, 0, NULL, 1, NULL, 1, 1717305025, 1717305025);
INSERT INTO `menu` VALUES (644, 643, '查询字典', 'A', NULL, NULL, '/manager/api/v1/dictionaries', 'GET', NULL, 'manager:dictionary:query', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1717305025, 1717305025);
INSERT INTO `menu` VALUES (645, 643, '新增字典', 'A', NULL, NULL, '/manager/api/v1/dictionary', 'POST', NULL, 'manager:dictionary:add', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1717305025, 1717305025);
INSERT INTO `menu` VALUES (646, 643, '修改字典', 'A', NULL, NULL, '/manager/api/v1/dictionary', 'PUT', NULL, 'manager:dictionary:update', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1717305025, 1717305025);
INSERT INTO `menu` VALUES (647, 643, '删除字典', 'A', NULL, NULL, '/manager/api/v1/dictionary', 'DELETE', NULL, 'manager:dictionary:delete', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1717305025, 1717305025);
INSERT INTO `menu` VALUES (648, 643, '获取字典值', 'A', NULL, NULL, '/manager/api/v1/dictionary_values', 'GET', NULL, 'manager:dictionary:value:query', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1717305025, 1717305025);
INSERT INTO `menu` VALUES (649, 643, '新增字典值', 'A', NULL, NULL, '/manager/api/v1/dictionary_value', 'POST', NULL, 'manager:dictionary:value:add', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1717305025, 1717305025);
INSERT INTO `menu` VALUES (650, 643, '修改字典值', 'A', NULL, NULL, '/manager/api/v1/dictionary_value', 'PUT', NULL, 'manager:dictionary:value:update', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1717305025, 1717305025);
INSERT INTO `menu` VALUES (651, 643, '更新字典值目录状态', 'A', NULL, NULL, '/manager/api/v1/dictionary_value/status', 'PUT', NULL, 'manager:dictionary:value:status', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1717305025, 1717305025);
INSERT INTO `menu` VALUES (652, 643, '删除字典值', 'A', NULL, NULL, '/manager/api/v1/dictionary_value', 'DELETE', NULL, 'manager:dictionary:value:delete', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1717305025, 1717305025);
INSERT INTO `menu` VALUES (653, 626, '菜单管理', 'M', 'ManagerMenu', 'menu', NULL, NULL, '/manager/menu', NULL, '/manager/menu/index', NULL, 0, 0, 1, 1, 1, 1717305025, 1717305025);
INSERT INTO `menu` VALUES (654, 653, '查询菜单', 'A', NULL, NULL, '/manager/api/v1/menus', 'GET', NULL, 'manager:menu:query', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1717305025, 1717305025);
INSERT INTO `menu` VALUES (655, 653, '新增菜单', 'A', NULL, NULL, '/manager/api/v1/menu', 'POST', NULL, 'manager:menu:add', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1717305025, 1717305025);
INSERT INTO `menu` VALUES (656, 653, '修改菜单', 'A', NULL, NULL, '/manager/api/v1/menu', 'PUT', NULL, 'manager:menu:update', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1717305025, 1717305025);
INSERT INTO `menu` VALUES (657, 653, '删除菜单', 'A', NULL, NULL, '/manager/api/v1/menu', 'DELETE', NULL, 'manager:menu:delete', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1717305025, 1717305025);
INSERT INTO `menu` VALUES (658, 626, '职位管理', 'M', 'ManagerJob', 'tag', NULL, NULL, '/manager/job', NULL, '/manager/job/index', NULL, 0, NULL, 1, NULL, 1, 1717305025, 1717305025);
INSERT INTO `menu` VALUES (659, 658, '查询职位', 'A', NULL, NULL, '/manager/api/v1/jobs', 'GET', NULL, 'manager:job:query', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1717305025, 1717305025);
INSERT INTO `menu` VALUES (660, 658, '新增职位', 'A', NULL, NULL, '/manager/api/v1/job', 'POST', NULL, 'manager:job:add', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1717305025, 1717305025);
INSERT INTO `menu` VALUES (661, 658, '修改职位', 'A', NULL, NULL, '/manager/api/v1/job', 'PUT', NULL, 'manager:job:update', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1717305025, 1717305025);
INSERT INTO `menu` VALUES (662, 658, '删除职位', 'A', NULL, NULL, '/manager/api/v1/job', 'DELETE', NULL, 'manager:job:delete', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1717305025, 1717305025);
INSERT INTO `menu` VALUES (663, 626, '部门管理', 'M', 'ManagerDepartment', 'user-group', NULL, NULL, '/manager/department', NULL, '/manager/department/index', NULL, 0, 0, 1, NULL, 1, 1717305025, 1717305025);
INSERT INTO `menu` VALUES (664, 663, '新增部门', 'A', NULL, NULL, '/manager/api/v1/department', 'POST', NULL, 'manager:department:add', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1717305025, 1717305025);
INSERT INTO `menu` VALUES (665, 663, '修改部门', 'A', NULL, NULL, '/manager/api/v1/department', 'PUT', NULL, 'manager:department:update', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1717305025, 1717305025);
INSERT INTO `menu` VALUES (666, 663, '删除部门', 'A', NULL, NULL, '/manager/api/v1/department', 'DELETE', NULL, 'manager:department:delete', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1717305025, 1717305025);
INSERT INTO `menu` VALUES (667, 626, '角色管理', 'M', 'ManagerRole', 'safe', NULL, NULL, '/manager/role', NULL, '/manager/role/index', NULL, 0, NULL, 0, NULL, 1, 1717305025, 1717305025);
INSERT INTO `menu` VALUES (668, 667, '新增角色', 'A', NULL, NULL, '/manager/api/v1/role', 'POST', NULL, 'manager:role:add', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1717305025, 1717305025);
INSERT INTO `menu` VALUES (669, 667, '修改角色', 'A', NULL, NULL, '/manager/api/v1/role', 'PUT', NULL, 'manager:role:update', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1717305025, 1717305025);
INSERT INTO `menu` VALUES (670, 667, '删除角色', 'A', NULL, NULL, '/manager/api/v1/role', 'DELETE', NULL, 'manager:role:delete', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1717305025, 1717305025);
INSERT INTO `menu` VALUES (671, 667, '角色菜单管理', 'G', NULL, NULL, NULL, NULL, NULL, 'manager:role:menu', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1717305025, 1717305025);
INSERT INTO `menu` VALUES (672, 671, '查询角色菜单', 'A', NULL, NULL, '/manager/api/v1/role/menu_ids', 'GET', NULL, 'manager:role:menu:query', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1717305025, 1717305025);
INSERT INTO `menu` VALUES (673, 671, '修改角色菜单', 'A', NULL, NULL, '/manager/api/v1/role/menu', 'POST', NULL, 'manager:role:menu:update', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1717305025, 1717305025);
INSERT INTO `menu` VALUES (674, 626, '用户管理', 'M', 'ManagerUser', 'user', NULL, NULL, '/manager/user', NULL, '/manager/user/index', NULL, 0, NULL, 1, NULL, 1, 1717305025, 1717305025);
INSERT INTO `menu` VALUES (675, 674, '查询用户列表', 'A', NULL, NULL, '/manager/api/v1/users', 'GET', NULL, 'manager:user:query', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1717305025, 1717305025);
INSERT INTO `menu` VALUES (676, 674, '新增用户', 'A', NULL, NULL, '/manager/api/v1/user', 'POST', NULL, 'manager:user:add', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1717305025, 1717305025);
INSERT INTO `menu` VALUES (677, 674, '修改用户', 'A', NULL, NULL, '/manager/api/v1/user', 'PUT', NULL, 'manager:user:update', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1717305025, 1717305025);
INSERT INTO `menu` VALUES (678, 674, '删除用户', 'A', NULL, NULL, '/manager/api/v1/user', 'DELETE', NULL, 'manager:user:delete', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1717305025, 1717305025);
INSERT INTO `menu` VALUES (679, 674, '修改用户状态', 'A', NULL, NULL, '/manager/api/v1/user/status', 'PUT', NULL, 'manager:user:status', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1717305025, 1717305025);
INSERT INTO `menu` VALUES (680, 674, '重置账号密码', 'A', NULL, NULL, '/manager/api/v1/user/password/reset', 'POST', NULL, 'manager:user:reset:password', '', NULL, 0, NULL, NULL, NULL, NULL, 1717305025, 1717305025);
COMMIT;

-- ----------------------------
-- Table structure for menu_closure
-- ----------------------------
DROP TABLE IF EXISTS `menu_closure`;
CREATE TABLE `menu_closure` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `parent` bigint unsigned NOT NULL COMMENT '菜单id',
  `children` bigint unsigned NOT NULL COMMENT '菜单id',
  PRIMARY KEY (`id`),
  KEY `parent` (`parent`),
  KEY `children` (`children`),
  CONSTRAINT `menu_closure_ibfk_1` FOREIGN KEY (`children`) REFERENCES `menu` (`id`) ON DELETE CASCADE,
  CONSTRAINT `menu_closure_ibfk_2` FOREIGN KEY (`parent`) REFERENCES `menu` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=420 DEFAULT CHARSET=utf8mb4 COMMENT='菜单层级信息';

-- ----------------------------
-- Records of menu_closure
-- ----------------------------
BEGIN;
INSERT INTO `menu_closure` VALUES (261, 624, 628);
INSERT INTO `menu_closure` VALUES (262, 626, 628);
INSERT INTO `menu_closure` VALUES (263, 627, 628);
INSERT INTO `menu_closure` VALUES (264, 624, 629);
INSERT INTO `menu_closure` VALUES (265, 626, 629);
INSERT INTO `menu_closure` VALUES (266, 627, 629);
INSERT INTO `menu_closure` VALUES (267, 624, 630);
INSERT INTO `menu_closure` VALUES (268, 626, 630);
INSERT INTO `menu_closure` VALUES (269, 627, 630);
INSERT INTO `menu_closure` VALUES (270, 624, 631);
INSERT INTO `menu_closure` VALUES (271, 626, 631);
INSERT INTO `menu_closure` VALUES (272, 627, 631);
INSERT INTO `menu_closure` VALUES (273, 624, 632);
INSERT INTO `menu_closure` VALUES (274, 626, 632);
INSERT INTO `menu_closure` VALUES (275, 627, 632);
INSERT INTO `menu_closure` VALUES (276, 624, 633);
INSERT INTO `menu_closure` VALUES (277, 626, 633);
INSERT INTO `menu_closure` VALUES (278, 627, 633);
INSERT INTO `menu_closure` VALUES (279, 624, 634);
INSERT INTO `menu_closure` VALUES (280, 626, 634);
INSERT INTO `menu_closure` VALUES (281, 627, 634);
INSERT INTO `menu_closure` VALUES (282, 624, 635);
INSERT INTO `menu_closure` VALUES (283, 626, 635);
INSERT INTO `menu_closure` VALUES (284, 627, 635);
INSERT INTO `menu_closure` VALUES (285, 624, 636);
INSERT INTO `menu_closure` VALUES (286, 626, 636);
INSERT INTO `menu_closure` VALUES (287, 627, 636);
INSERT INTO `menu_closure` VALUES (288, 624, 637);
INSERT INTO `menu_closure` VALUES (289, 626, 637);
INSERT INTO `menu_closure` VALUES (290, 627, 637);
INSERT INTO `menu_closure` VALUES (291, 624, 638);
INSERT INTO `menu_closure` VALUES (292, 626, 638);
INSERT INTO `menu_closure` VALUES (293, 627, 638);
INSERT INTO `menu_closure` VALUES (294, 624, 639);
INSERT INTO `menu_closure` VALUES (295, 626, 639);
INSERT INTO `menu_closure` VALUES (296, 627, 639);
INSERT INTO `menu_closure` VALUES (297, 624, 640);
INSERT INTO `menu_closure` VALUES (298, 626, 640);
INSERT INTO `menu_closure` VALUES (299, 627, 640);
INSERT INTO `menu_closure` VALUES (300, 624, 641);
INSERT INTO `menu_closure` VALUES (301, 626, 641);
INSERT INTO `menu_closure` VALUES (302, 627, 641);
INSERT INTO `menu_closure` VALUES (303, 624, 642);
INSERT INTO `menu_closure` VALUES (304, 626, 642);
INSERT INTO `menu_closure` VALUES (305, 627, 642);
INSERT INTO `menu_closure` VALUES (306, 624, 644);
INSERT INTO `menu_closure` VALUES (307, 626, 644);
INSERT INTO `menu_closure` VALUES (308, 643, 644);
INSERT INTO `menu_closure` VALUES (309, 624, 645);
INSERT INTO `menu_closure` VALUES (310, 626, 645);
INSERT INTO `menu_closure` VALUES (311, 643, 645);
INSERT INTO `menu_closure` VALUES (312, 624, 646);
INSERT INTO `menu_closure` VALUES (313, 626, 646);
INSERT INTO `menu_closure` VALUES (314, 643, 646);
INSERT INTO `menu_closure` VALUES (315, 624, 647);
INSERT INTO `menu_closure` VALUES (316, 626, 647);
INSERT INTO `menu_closure` VALUES (317, 643, 647);
INSERT INTO `menu_closure` VALUES (318, 624, 648);
INSERT INTO `menu_closure` VALUES (319, 626, 648);
INSERT INTO `menu_closure` VALUES (320, 643, 648);
INSERT INTO `menu_closure` VALUES (321, 624, 649);
INSERT INTO `menu_closure` VALUES (322, 626, 649);
INSERT INTO `menu_closure` VALUES (323, 643, 649);
INSERT INTO `menu_closure` VALUES (324, 624, 650);
INSERT INTO `menu_closure` VALUES (325, 626, 650);
INSERT INTO `menu_closure` VALUES (326, 643, 650);
INSERT INTO `menu_closure` VALUES (327, 624, 651);
INSERT INTO `menu_closure` VALUES (328, 626, 651);
INSERT INTO `menu_closure` VALUES (329, 643, 651);
INSERT INTO `menu_closure` VALUES (330, 624, 652);
INSERT INTO `menu_closure` VALUES (331, 626, 652);
INSERT INTO `menu_closure` VALUES (332, 643, 652);
INSERT INTO `menu_closure` VALUES (333, 624, 654);
INSERT INTO `menu_closure` VALUES (334, 626, 654);
INSERT INTO `menu_closure` VALUES (335, 653, 654);
INSERT INTO `menu_closure` VALUES (336, 624, 655);
INSERT INTO `menu_closure` VALUES (337, 626, 655);
INSERT INTO `menu_closure` VALUES (338, 653, 655);
INSERT INTO `menu_closure` VALUES (339, 624, 656);
INSERT INTO `menu_closure` VALUES (340, 626, 656);
INSERT INTO `menu_closure` VALUES (341, 653, 656);
INSERT INTO `menu_closure` VALUES (342, 624, 657);
INSERT INTO `menu_closure` VALUES (343, 626, 657);
INSERT INTO `menu_closure` VALUES (344, 653, 657);
INSERT INTO `menu_closure` VALUES (345, 624, 659);
INSERT INTO `menu_closure` VALUES (346, 626, 659);
INSERT INTO `menu_closure` VALUES (347, 658, 659);
INSERT INTO `menu_closure` VALUES (348, 624, 660);
INSERT INTO `menu_closure` VALUES (349, 626, 660);
INSERT INTO `menu_closure` VALUES (350, 658, 660);
INSERT INTO `menu_closure` VALUES (351, 624, 661);
INSERT INTO `menu_closure` VALUES (352, 626, 661);
INSERT INTO `menu_closure` VALUES (353, 658, 661);
INSERT INTO `menu_closure` VALUES (354, 624, 662);
INSERT INTO `menu_closure` VALUES (355, 626, 662);
INSERT INTO `menu_closure` VALUES (356, 658, 662);
INSERT INTO `menu_closure` VALUES (357, 624, 664);
INSERT INTO `menu_closure` VALUES (358, 626, 664);
INSERT INTO `menu_closure` VALUES (359, 663, 664);
INSERT INTO `menu_closure` VALUES (360, 624, 665);
INSERT INTO `menu_closure` VALUES (361, 626, 665);
INSERT INTO `menu_closure` VALUES (362, 663, 665);
INSERT INTO `menu_closure` VALUES (363, 624, 666);
INSERT INTO `menu_closure` VALUES (364, 626, 666);
INSERT INTO `menu_closure` VALUES (365, 663, 666);
INSERT INTO `menu_closure` VALUES (366, 624, 672);
INSERT INTO `menu_closure` VALUES (367, 626, 672);
INSERT INTO `menu_closure` VALUES (368, 667, 672);
INSERT INTO `menu_closure` VALUES (369, 671, 672);
INSERT INTO `menu_closure` VALUES (370, 624, 673);
INSERT INTO `menu_closure` VALUES (371, 626, 673);
INSERT INTO `menu_closure` VALUES (372, 667, 673);
INSERT INTO `menu_closure` VALUES (373, 671, 673);
INSERT INTO `menu_closure` VALUES (374, 624, 668);
INSERT INTO `menu_closure` VALUES (375, 626, 668);
INSERT INTO `menu_closure` VALUES (376, 667, 668);
INSERT INTO `menu_closure` VALUES (377, 624, 669);
INSERT INTO `menu_closure` VALUES (378, 626, 669);
INSERT INTO `menu_closure` VALUES (379, 667, 669);
INSERT INTO `menu_closure` VALUES (380, 624, 670);
INSERT INTO `menu_closure` VALUES (381, 626, 670);
INSERT INTO `menu_closure` VALUES (382, 667, 670);
INSERT INTO `menu_closure` VALUES (383, 624, 671);
INSERT INTO `menu_closure` VALUES (384, 626, 671);
INSERT INTO `menu_closure` VALUES (385, 667, 671);
INSERT INTO `menu_closure` VALUES (386, 624, 675);
INSERT INTO `menu_closure` VALUES (387, 626, 675);
INSERT INTO `menu_closure` VALUES (388, 674, 675);
INSERT INTO `menu_closure` VALUES (389, 624, 676);
INSERT INTO `menu_closure` VALUES (390, 626, 676);
INSERT INTO `menu_closure` VALUES (391, 674, 676);
INSERT INTO `menu_closure` VALUES (392, 624, 677);
INSERT INTO `menu_closure` VALUES (393, 626, 677);
INSERT INTO `menu_closure` VALUES (394, 674, 677);
INSERT INTO `menu_closure` VALUES (395, 624, 678);
INSERT INTO `menu_closure` VALUES (396, 626, 678);
INSERT INTO `menu_closure` VALUES (397, 674, 678);
INSERT INTO `menu_closure` VALUES (398, 624, 679);
INSERT INTO `menu_closure` VALUES (399, 626, 679);
INSERT INTO `menu_closure` VALUES (400, 674, 679);
INSERT INTO `menu_closure` VALUES (401, 624, 680);
INSERT INTO `menu_closure` VALUES (402, 626, 680);
INSERT INTO `menu_closure` VALUES (403, 674, 680);
INSERT INTO `menu_closure` VALUES (404, 624, 627);
INSERT INTO `menu_closure` VALUES (405, 626, 627);
INSERT INTO `menu_closure` VALUES (406, 624, 643);
INSERT INTO `menu_closure` VALUES (407, 626, 643);
INSERT INTO `menu_closure` VALUES (408, 624, 653);
INSERT INTO `menu_closure` VALUES (409, 626, 653);
INSERT INTO `menu_closure` VALUES (410, 624, 658);
INSERT INTO `menu_closure` VALUES (411, 626, 658);
INSERT INTO `menu_closure` VALUES (412, 624, 663);
INSERT INTO `menu_closure` VALUES (413, 626, 663);
INSERT INTO `menu_closure` VALUES (414, 624, 667);
INSERT INTO `menu_closure` VALUES (415, 626, 667);
INSERT INTO `menu_closure` VALUES (416, 624, 674);
INSERT INTO `menu_closure` VALUES (417, 626, 674);
INSERT INTO `menu_closure` VALUES (418, 624, 625);
INSERT INTO `menu_closure` VALUES (419, 624, 626);
COMMIT;

-- ----------------------------
-- Table structure for role
-- ----------------------------
DROP TABLE IF EXISTS `role`;
CREATE TABLE `role` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `parent_id` bigint unsigned NOT NULL COMMENT '父id',
  `name` varchar(64) NOT NULL COMMENT '名称',
  `keyword` char(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT '标识',
  `status` tinyint(1) NOT NULL DEFAULT '0' COMMENT '状态',
  `description` varchar(128) NOT NULL COMMENT '描述',
  `department_ids` tinytext COMMENT '自定义部门',
  `data_scope` char(32) NOT NULL COMMENT '权限类型',
  `created_at` bigint unsigned NOT NULL DEFAULT '0' COMMENT '创建时间',
  `updated_at` bigint unsigned NOT NULL DEFAULT '0' COMMENT '修改时间',
  `deleted_at` bigint unsigned NOT NULL DEFAULT '0' COMMENT '删除时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `keyword` (`keyword`,`deleted_at`),
  KEY `idx_role_created_at` (`created_at`),
  KEY `idx_role_updated_at` (`updated_at`),
  KEY `idx_role_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COMMENT='角色信息';

-- ----------------------------
-- Records of role
-- ----------------------------
BEGIN;
INSERT INTO `role` VALUES (1, 0, '超级管理员', 'superAdmin', 1, '超级管理员  ', NULL, 'ALL', 1713706137, 1713706137, 0);
COMMIT;

-- ----------------------------
-- Table structure for role_closure
-- ----------------------------
DROP TABLE IF EXISTS `role_closure`;
CREATE TABLE `role_closure` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `parent` bigint unsigned NOT NULL COMMENT '角色id',
  `children` bigint unsigned NOT NULL COMMENT '角色id',
  PRIMARY KEY (`id`),
  KEY `parent` (`parent`),
  KEY `children` (`children`),
  CONSTRAINT `role_closure_ibfk_1` FOREIGN KEY (`children`) REFERENCES `role` (`id`) ON DELETE CASCADE,
  CONSTRAINT `role_closure_ibfk_2` FOREIGN KEY (`parent`) REFERENCES `role` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COMMENT='角色层级信息';

-- ----------------------------
-- Records of role_closure
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for role_menu
-- ----------------------------
DROP TABLE IF EXISTS `role_menu`;
CREATE TABLE `role_menu` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `role_id` bigint unsigned NOT NULL COMMENT '角色id',
  `menu_id` bigint unsigned NOT NULL COMMENT '菜单id',
  PRIMARY KEY (`id`),
  UNIQUE KEY `role_id_2` (`role_id`,`menu_id`),
  KEY `role_id` (`role_id`),
  KEY `menu_id` (`menu_id`),
  CONSTRAINT `role_menu_ibfk_1` FOREIGN KEY (`menu_id`) REFERENCES `menu` (`id`) ON DELETE CASCADE,
  CONSTRAINT `role_menu_ibfk_2` FOREIGN KEY (`role_id`) REFERENCES `role` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=219 DEFAULT CHARSET=utf8mb4 COMMENT='角色菜单信息';

-- ----------------------------
-- Records of role_menu
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `department_id` bigint unsigned NOT NULL COMMENT '部门id',
  `role_id` bigint unsigned NOT NULL COMMENT '角色id',
  `name` char(32) NOT NULL COMMENT '名称',
  `nickname` varchar(64) NOT NULL COMMENT '昵称',
  `gender` char(32) NOT NULL COMMENT '性别',
  `avatar` varchar(256) DEFAULT NULL COMMENT '头像',
  `email` varchar(64) NOT NULL COMMENT '邮箱',
  `phone` char(32) NOT NULL COMMENT '电话',
  `password` varchar(256) NOT NULL COMMENT '密码',
  `status` tinyint(1) DEFAULT '0' COMMENT '状态',
  `setting` tinytext COMMENT '用户设置',
  `token` varchar(512) DEFAULT NULL COMMENT '用户token',
  `logged_at` bigint NOT NULL DEFAULT '0' COMMENT '登陆时间',
  `created_at` bigint NOT NULL DEFAULT '0' COMMENT '创建时间',
  `updated_at` bigint NOT NULL DEFAULT '0' COMMENT '修改时间',
  PRIMARY KEY (`id`),
  KEY `idx_user_updated_at` (`updated_at`),
  KEY `idx_user_created_at` (`created_at`),
  KEY `fk_user_role` (`role_id`),
  KEY `fk_user_department` (`department_id`),
  CONSTRAINT `fk_user_department` FOREIGN KEY (`department_id`) REFERENCES `department` (`id`),
  CONSTRAINT `fk_user_role` FOREIGN KEY (`role_id`) REFERENCES `role` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COMMENT='用户信息';

-- ----------------------------
-- Records of user
-- ----------------------------
BEGIN;
INSERT INTO `user` VALUES (1, 1, 1, '超级管理员', '超级管理员1', 'M', '', '1280291001@qq.com', '18888888888', '$2a$10$9qRJe9KQo6sEcU8ipKg.e.dkl2E7Wy64SigYlgraTAn.1paHFq6W.', 1, '{\"theme\":\"light\",\"themeColor\":\"#165DFF\",\"skin\":\"default\",\"tabBar\":true,\"menuWidth\":300,\"layout\":\"twoColumns\",\"language\":\"zh_CN\",\"animation\":\"gp-fade\"}', 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJkZXBhcnRtZW50SWQiOjEsImRlcGFydG1lbnRLZXl3b3JkIjoiY29tcGFueSIsImV4cCI6MTcxNzMyMzA5NywiaWF0IjoxNzE3MzE1ODk2LCJuYmYiOjE3MTczMTU4OTYsInJvbGVJZCI6MSwicm9sZUtleXdvcmQiOiJzdXBlckFkbWluIiwidXNlcklkIjoxfQ.EauGv5JjbIPde9lf-wURRmhUPQ5FK483xF5WV9WS-9U', 1717315896, 1713706137, 1717315896);
COMMIT;

-- ----------------------------
-- Table structure for user_job
-- ----------------------------
DROP TABLE IF EXISTS `user_job`;
CREATE TABLE `user_job` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `job_id` bigint unsigned NOT NULL COMMENT '职位id',
  `user_id` bigint unsigned NOT NULL COMMENT '用户id',
  PRIMARY KEY (`id`),
  UNIQUE KEY `job_id` (`job_id`,`user_id`),
  KEY `user_id` (`user_id`),
  CONSTRAINT `user_job_ibfk_1` FOREIGN KEY (`job_id`) REFERENCES `job` (`id`) ON DELETE CASCADE,
  CONSTRAINT `user_job_ibfk_2` FOREIGN KEY (`user_id`) REFERENCES `user` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8mb4 COMMENT='用户职位信息';

-- ----------------------------
-- Records of user_job
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for user_role
-- ----------------------------
DROP TABLE IF EXISTS `user_role`;
CREATE TABLE `user_role` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `role_id` bigint unsigned NOT NULL COMMENT '角色id',
  `user_id` bigint unsigned NOT NULL COMMENT '用户id',
  PRIMARY KEY (`id`),
  UNIQUE KEY `role_id` (`role_id`,`user_id`),
  KEY `user_id` (`user_id`),
  CONSTRAINT `user_role_ibfk_1` FOREIGN KEY (`role_id`) REFERENCES `role` (`id`) ON DELETE CASCADE,
  CONSTRAINT `user_role_ibfk_2` FOREIGN KEY (`user_id`) REFERENCES `user` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COMMENT='用户角色信息';

-- ----------------------------
-- Records of user_role
-- ----------------------------
BEGIN;
INSERT INTO `user_role` VALUES (2, 1, 1);
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
