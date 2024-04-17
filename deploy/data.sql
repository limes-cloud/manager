/*
 Navicat Premium Data Transfer

 Source Server         : dev
 Source Server Type    : MySQL
 Source Server Version : 80200
 Source Host           : localhost:3306
 Source Schema         : manager

 Target Server Type    : MySQL
 Target Server Version : 80200
 File Encoding         : 65001

 Date: 17/04/2024 12:47:12
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
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

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
  `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `created_at` bigint DEFAULT NULL COMMENT '创建时间',
  `updated_at` bigint DEFAULT NULL COMMENT '修改时间',
  `parent_id` int unsigned NOT NULL COMMENT '父id',
  `keyword` varchar(32) NOT NULL COMMENT '关键字',
  `name` varchar(32) NOT NULL COMMENT '名称',
  `description` varchar(256) NOT NULL COMMENT '描述',
  PRIMARY KEY (`id`),
  UNIQUE KEY `keyword` (`keyword`),
  UNIQUE KEY `name` (`name`),
  KEY `idx_department_created_at` (`created_at`),
  KEY `idx_department_updated_at` (`updated_at`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='部门信息';

-- ----------------------------
-- Records of department
-- ----------------------------
BEGIN;
INSERT INTO `department` VALUES (1, 1713290381, 1713290381, 0, 'company', '贵州青橙科技有限公司', '开放合作，拥抱未来');
COMMIT;

-- ----------------------------
-- Table structure for department_closure
-- ----------------------------
DROP TABLE IF EXISTS `department_closure`;
CREATE TABLE `department_closure` (
  `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `parent` int unsigned NOT NULL COMMENT '部门id',
  `children` int unsigned NOT NULL COMMENT '部门id',
  PRIMARY KEY (`id`),
  KEY `fk_department_closure_parent_department` (`parent`),
  KEY `fk_department_closure_children_department` (`children`),
  CONSTRAINT `fk_department_closure_children_department` FOREIGN KEY (`children`) REFERENCES `department` (`id`) ON DELETE CASCADE,
  CONSTRAINT `fk_department_closure_parent_department` FOREIGN KEY (`parent`) REFERENCES `department` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='部门层级信息';

-- ----------------------------
-- Records of department_closure
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for department_object
-- ----------------------------
DROP TABLE IF EXISTS `department_object`;
CREATE TABLE `department_object` (
  `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `department_id` int unsigned NOT NULL COMMENT '部门资源id',
  `object_id` int unsigned NOT NULL COMMENT '资源对象id',
  `value` varchar(64) NOT NULL COMMENT '资源对象值',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='部门资源';

-- ----------------------------
-- Records of department_object
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for dictionary
-- ----------------------------
DROP TABLE IF EXISTS `dictionary`;
CREATE TABLE `dictionary` (
  `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `created_at` bigint DEFAULT NULL COMMENT '创建时间',
  `updated_at` bigint DEFAULT NULL COMMENT '修改时间',
  `keyword` varchar(32) NOT NULL COMMENT '关键字',
  `name` varchar(32) NOT NULL COMMENT '名称',
  `description` varchar(256) NOT NULL COMMENT '描述',
  PRIMARY KEY (`id`),
  UNIQUE KEY `keyword` (`keyword`),
  UNIQUE KEY `name` (`name`),
  KEY `idx_dictionary_updated_at` (`updated_at`),
  KEY `idx_dictionary_created_at` (`created_at`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='字典信息';

-- ----------------------------
-- Records of dictionary
-- ----------------------------
BEGIN;
INSERT INTO `dictionary` VALUES (1, 1713290381, 1713290381, 'gender', '性别', '性别选项');
COMMIT;

-- ----------------------------
-- Table structure for dictionary_value
-- ----------------------------
DROP TABLE IF EXISTS `dictionary_value`;
CREATE TABLE `dictionary_value` (
  `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `created_at` bigint DEFAULT NULL COMMENT '创建时间',
  `updated_at` bigint DEFAULT NULL COMMENT '修改时间',
  `dictionary_id` int unsigned NOT NULL COMMENT '字典id',
  `label` varchar(128) NOT NULL COMMENT '标签',
  `value` varchar(128) NOT NULL COMMENT '值',
  `status` tinyint(1) DEFAULT '1' COMMENT '状态',
  `weight` int unsigned DEFAULT NULL COMMENT '权重',
  `type` varchar(32) DEFAULT NULL COMMENT '字典类型',
  `extra` varchar(256) DEFAULT NULL COMMENT '扩展信息',
  `description` varchar(256) DEFAULT NULL COMMENT '描述',
  PRIMARY KEY (`id`),
  UNIQUE KEY `label` (`label`),
  UNIQUE KEY `value` (`value`),
  UNIQUE KEY `dv` (`dictionary_id`,`value`),
  KEY `idx_dictionary_value_created_at` (`created_at`),
  KEY `idx_dictionary_value_updated_at` (`updated_at`),
  CONSTRAINT `fk_dictionary_value_dict` FOREIGN KEY (`dictionary_id`) REFERENCES `dictionary` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='字典值信息';

-- ----------------------------
-- Records of dictionary_value
-- ----------------------------
BEGIN;
INSERT INTO `dictionary_value` VALUES (1, 1713290381, 1713290381, 1, '男', 'M', 1, 0, '', '', '男性');
INSERT INTO `dictionary_value` VALUES (2, 1713290381, 1713290381, 1, '女', 'F', 1, 0, '', '', '女性');
INSERT INTO `dictionary_value` VALUES (3, 1713290381, 1713290381, 1, '未知', 'U', 1, 0, '', '', '未知性别');
COMMIT;

-- ----------------------------
-- Table structure for job
-- ----------------------------
DROP TABLE IF EXISTS `job`;
CREATE TABLE `job` (
  `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `created_at` bigint DEFAULT NULL COMMENT '创建时间',
  `updated_at` bigint DEFAULT NULL COMMENT '修改时间',
  `keyword` varchar(32) NOT NULL COMMENT '关键字',
  `name` varchar(32) NOT NULL COMMENT '名称',
  `weight` int unsigned DEFAULT NULL COMMENT '权重',
  `description` varchar(256) NOT NULL COMMENT '描述',
  PRIMARY KEY (`id`),
  UNIQUE KEY `keyword` (`keyword`),
  UNIQUE KEY `name` (`name`),
  KEY `idx_job_created_at` (`created_at`),
  KEY `idx_job_updated_at` (`updated_at`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='职位信息';

-- ----------------------------
-- Records of job
-- ----------------------------
BEGIN;
INSERT INTO `job` VALUES (1, 1713290381, 1713290381, 'chairman', '董事长', 0, '董事长');
COMMIT;

-- ----------------------------
-- Table structure for menu
-- ----------------------------
DROP TABLE IF EXISTS `menu`;
CREATE TABLE `menu` (
  `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `created_at` bigint DEFAULT NULL COMMENT '创建时间',
  `updated_at` bigint DEFAULT NULL COMMENT '修改时间',
  `parent_id` int unsigned NOT NULL COMMENT '父id',
  `app` char(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT '服务',
  `title` varchar(128) NOT NULL COMMENT '标题',
  `type` char(32) NOT NULL COMMENT '类型',
  `keyword` varchar(64) DEFAULT NULL COMMENT '关键词',
  `icon` char(32) DEFAULT NULL COMMENT '图标',
  `api` varchar(128) DEFAULT NULL COMMENT '接口',
  `method` varchar(12) DEFAULT NULL COMMENT '接口方法',
  `path` varchar(128) DEFAULT NULL COMMENT '路径',
  `permission` varchar(128) DEFAULT NULL COMMENT '指令',
  `check_object` tinyint(1) DEFAULT NULL COMMENT '校验资源对象',
  `check_object_rule` varchar(256) DEFAULT NULL COMMENT '校验规则',
  `component` varchar(128) DEFAULT NULL COMMENT '组件',
  `redirect` varchar(128) DEFAULT NULL COMMENT '重定向地址',
  `weight` int unsigned DEFAULT NULL COMMENT '权重',
  `is_hidden` tinyint(1) DEFAULT NULL COMMENT '是否隐藏',
  `is_cache` tinyint(1) DEFAULT NULL COMMENT '是否缓存',
  `is_home` tinyint(1) DEFAULT NULL COMMENT '是否为首页',
  `is_affix` tinyint(1) DEFAULT NULL COMMENT '是否为标签',
  PRIMARY KEY (`id`),
  UNIQUE KEY `keyword` (`keyword`),
  UNIQUE KEY `path` (`path`),
  KEY `idx_menu_created_at` (`created_at`),
  KEY `idx_menu_updated_at` (`updated_at`)
) ENGINE=InnoDB AUTO_INCREMENT=264 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='菜单信息';

-- ----------------------------
-- Records of menu
-- ----------------------------
BEGIN;
INSERT INTO `menu` VALUES (1, 1713290381, 1713290381, 0, 'SystemPlatform', '管理平台', 'R', 'SystemPlatform', 'apps', NULL, NULL, '/', NULL, NULL, NULL, 'Layout', NULL, 2, 0, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (2, 1713290381, 1713290381, 1, 'SystemPlatform', '首页面板', 'M', 'Dashboard', 'dashboard', NULL, NULL, '/dashboard', NULL, NULL, NULL, '/dashboard/index', NULL, 0, 0, 1, 1, 1);
INSERT INTO `menu` VALUES (3, 1713290381, 1713290381, 1, 'SystemPlatform', '管理中心', 'M', 'SystemPlatformManager', 'apps', NULL, NULL, '/manager', NULL, NULL, NULL, NULL, NULL, 0, 0, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (4, 1713290381, 1713290381, 3, 'SystemPlatform', '基础接口', 'G', 'BaseApi', 'apps', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, 0, 1, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (5, 1713290381, 1713290381, 4, 'SystemPlatform', '获取用户部门', 'BA', NULL, NULL, '/manager/v1/department/tree', 'GET', NULL, NULL, NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (6, 1713290381, 1713290381, 4, 'SystemPlatform', '获取个人用户信息', 'BA', NULL, NULL, '/manager/v1/user/current', 'GET', NULL, NULL, NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (7, 1713290381, 1713290381, 4, 'SystemPlatform', '获取用户可见角色树', 'BA', NULL, NULL, '/manager/v1/role/tree', 'GET', NULL, NULL, NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (8, 1713290381, 1713290381, 4, 'SystemPlatform', '获取用户菜单', 'BA', NULL, NULL, '/manager/v1/menu/tree/from/role', 'GET', NULL, NULL, NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (9, 1713290381, 1713290381, 4, 'SystemPlatform', '退出登录', 'BA', NULL, NULL, '/manager/v1/user/logout', 'POST', NULL, NULL, NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (10, 1713290381, 1713290381, 4, 'SystemPlatform', '刷新token', 'BA', NULL, NULL, '/manager/v1/user/token/refresh', 'POST', NULL, NULL, NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (11, 1713290381, 1713290381, 4, 'SystemPlatform', '用户登录', 'BA', NULL, NULL, '/manager/v1/user/login', 'POST', NULL, NULL, NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (12, 1713290381, 1713290381, 4, 'SystemPlatform', '获取登录验证码', 'BA', NULL, NULL, '/manager/v1/user/login/captcha', 'POST', NULL, NULL, NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (13, 1713290381, 1713290381, 4, 'SystemPlatform', '获取系统基础设置', 'BA', NULL, NULL, '/manager/v1/setting', 'GET', NULL, NULL, NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (14, 1713290381, 1713290381, 4, 'SystemPlatform', '接口鉴权', 'BA', NULL, NULL, '/manager/v1/auth', 'POST', NULL, NULL, NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (15, 1713290381, 1713290381, 4, 'SystemPlatform', '切换用户角色', 'BA', NULL, NULL, '/manager/v1/user/role/switch', 'POST', NULL, NULL, NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (16, 1713290381, 1713290381, 4, 'SystemPlatform', '新增部门资源对象', 'BA', NULL, NULL, '/manager/v1/department/object', 'POST', NULL, NULL, NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (17, 1713290381, 1713290381, 4, 'SystemPlatform', '删除部门资源对象', 'BA', NULL, NULL, '/manager/v1/department/object', 'DELETE', NULL, NULL, NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (18, 1713290381, 1713290381, 3, 'SystemPlatform', '资源对象', 'M', 'Object', 'common', NULL, NULL, '/manager/object', NULL, NULL, NULL, '/manager/object/index', NULL, 0, NULL, 1, NULL, 1);
INSERT INTO `menu` VALUES (19, 1713290381, 1713290381, 18, 'SystemPlatform', '查询对象', 'A', NULL, NULL, '/manager/v1/objects', 'GET', NULL, 'manager:object:query', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (20, 1713290381, 1713290381, 18, 'SystemPlatform', '新增对象', 'A', NULL, NULL, '/manager/v1/object', 'POST', NULL, 'manager:object:add', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (21, 1713290381, 1713290381, 18, 'SystemPlatform', '修改对象', 'A', NULL, NULL, '/manager/v1/object', 'PUT', NULL, 'manager:object:update', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (22, 1713290381, 1713290381, 18, 'SystemPlatform', '删除对象', 'A', NULL, NULL, '/manager/v1/object', 'DELETE', NULL, 'manager:object:delete', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (23, 1713290381, 1713290381, 3, 'SystemPlatform', '字典管理', 'M', 'Dict', 'storage', NULL, NULL, '/manager/dictionary', NULL, NULL, NULL, '/manager/dictionary/index', NULL, 0, NULL, 1, NULL, 1);
INSERT INTO `menu` VALUES (24, 1713290381, 1713290381, 23, 'SystemPlatform', '查询字典', 'A', NULL, NULL, '/manager/v1/dictionaries', 'GET', NULL, 'manager:dictionary:query', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (25, 1713290381, 1713290381, 23, 'SystemPlatform', '新增字典', 'A', NULL, NULL, '/manager/v1/dictionary', 'POST', NULL, 'manager:dictionary:add', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (26, 1713290381, 1713290381, 23, 'SystemPlatform', '修改字典', 'A', NULL, NULL, '/manager/v1/dictionary', 'PUT', NULL, 'manager:dictionary:update', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (27, 1713290381, 1713290381, 23, 'SystemPlatform', '删除字典', 'A', NULL, NULL, '/manager/v1/dictionary', 'DELETE', NULL, 'manager:dictionary:delete', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (28, 1713290381, 1713290381, 23, 'SystemPlatform', '刷新字典值', 'A', NULL, NULL, '/manager/v1/dictionary/values', 'GET', NULL, 'manager:dictionary:value:query', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (29, 1713290381, 1713290381, 23, 'SystemPlatform', '新增字典值', 'A', NULL, NULL, '/manager/v1/dictionary/value', 'POST', NULL, 'manager:dictionary:value:add', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (30, 1713290381, 1713290381, 23, 'SystemPlatform', '修改字典值', 'A', NULL, NULL, '/manager/v1/dictionary/value', 'PUT', NULL, 'manager:dictionary:value:update', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (31, 1713290381, 1713290381, 23, 'SystemPlatform', '删除字典值', 'A', NULL, NULL, '/manager/v1/dictionary/value', 'DELETE', NULL, 'manager:dictionary:value:delete', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (32, 1713290381, 1713290381, 3, 'SystemPlatform', '菜单管理', 'M', 'ManagerMenu', 'menu', NULL, NULL, '/manager/menu', NULL, NULL, NULL, '/manager/menu/index', NULL, 0, 0, 1, 1, 1);
INSERT INTO `menu` VALUES (33, 1713290381, 1713290381, 32, 'SystemPlatform', '查询菜单', 'A', NULL, NULL, '/manager/v1/menu/tree', 'GET', NULL, 'manager:menu:query', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (34, 1713290381, 1713290381, 32, 'SystemPlatform', '新增菜单', 'A', NULL, NULL, '/manager/v1/menu', 'POST', NULL, 'manager:menu:add', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (35, 1713290381, 1713290381, 32, 'SystemPlatform', '修改菜单', 'A', NULL, NULL, '/manager/v1/menu', 'PUT', NULL, 'manager:menu:update', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (36, 1713290381, 1713290381, 32, 'SystemPlatform', '删除菜单', 'A', NULL, NULL, '/manager/v1/menu', 'DELETE', NULL, 'manager:menu:delete', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (37, 1713290381, 1713290381, 3, 'SystemPlatform', '职位管理', 'M', 'JobUser', 'tag', NULL, NULL, '/manager/job', NULL, NULL, NULL, '/manager/job/index', NULL, 0, NULL, 1, NULL, 1);
INSERT INTO `menu` VALUES (38, 1713290381, 1713290381, 37, 'SystemPlatform', '查询职位', 'A', NULL, NULL, '/manager/v1/jobs', 'GET', NULL, 'manager:job:query', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (39, 1713290381, 1713290381, 37, 'SystemPlatform', '新增职位', 'A', NULL, NULL, '/manager/v1/job', 'POST', NULL, 'manager:job:add', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (40, 1713290381, 1713290381, 37, 'SystemPlatform', '修改职位', 'A', NULL, NULL, '/manager/v1/job', 'PUT', NULL, 'manager:job:update', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (41, 1713290381, 1713290381, 37, 'SystemPlatform', '删除职位', 'A', NULL, NULL, '/manager/v1/job', 'DELETE', NULL, 'manager:job:delete', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (42, 1713290381, 1713290381, 3, 'SystemPlatform', '部门管理', 'M', 'ManagerDepartment', 'user-group', NULL, NULL, '/manager/department', NULL, NULL, NULL, '/manager/department/index', NULL, 0, 0, 1, NULL, 1);
INSERT INTO `menu` VALUES (43, 1713290381, 1713290381, 42, 'SystemPlatform', '查看部门', 'A', NULL, NULL, '/manager/v1/department/tree', 'GET', NULL, 'manager:department:query', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (44, 1713290381, 1713290381, 42, 'SystemPlatform', '新增部门', 'A', NULL, NULL, '/manager/v1/department', 'POST', NULL, 'manager:department:add', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (45, 1713290381, 1713290381, 42, 'SystemPlatform', '修改部门', 'A', NULL, NULL, '/manager/v1/department', 'PUT', NULL, 'manager:department:update', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (46, 1713290381, 1713290381, 42, 'SystemPlatform', '删除部门', 'A', NULL, NULL, '/manager/v1/department', 'DELETE', NULL, 'manager:department:delete', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (47, 1713290381, 1713290381, 42, 'SystemPlatform', '查看部门资源对象', 'A', NULL, NULL, '/manager/v1/department/objects', 'GET', NULL, 'manager:department:object:query', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (48, 1713290381, 1713290381, 42, 'SystemPlatform', '设置部门资源对象', 'A', NULL, NULL, '/manager/v1/department/objects', 'POST', NULL, 'manager:department:object:import', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (49, 1713290381, 1713290381, 3, 'SystemPlatform', '角色管理', 'M', 'ManagerRole', 'safe', NULL, NULL, '/manager/role', NULL, NULL, NULL, '/manager/role/index', NULL, 0, NULL, 0, NULL, 1);
INSERT INTO `menu` VALUES (50, 1713290381, 1713290381, 49, 'SystemPlatform', '查询角色', 'A', NULL, NULL, '/manager/v1/role/tree', 'GET', NULL, 'manager:role:query', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (51, 1713290381, 1713290381, 49, 'SystemPlatform', '新增角色', 'A', NULL, NULL, '/manager/v1/role', 'POST', NULL, 'manager:role:add', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (52, 1713290381, 1713290381, 49, 'SystemPlatform', '修改角色', 'A', NULL, NULL, '/manager/v1/role', 'PUT', NULL, 'manager:role:update', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (53, 1713290381, 1713290381, 49, 'SystemPlatform', '删除角色', 'A', NULL, NULL, '/manager/v1/role', 'DELETE', NULL, 'manager:role:delete', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (54, 1713290381, 1713290381, 49, 'SystemPlatform', '角色菜单管理', 'G', NULL, NULL, '/manager/v1/role/menu/tree', 'GET', NULL, 'manager:role:menu', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (55, 1713290381, 1713290381, 54, 'SystemPlatform', '查询角色菜单', 'A', NULL, NULL, '/manager/v1/role/menu/ids', 'GET', NULL, 'manager:role:menu:query', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (56, 1713290381, 1713290381, 54, 'SystemPlatform', '修改角色菜单', 'A', NULL, NULL, '/manager/v1/role/menu', 'POST', NULL, 'manager:role:menu:update', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (57, 1713290381, 1713290381, 3, 'SystemPlatform', '用户管理', 'M', 'ManagerUser', 'user', NULL, NULL, '/manager/user', NULL, NULL, NULL, '/manager/user/index', NULL, 0, NULL, 1, NULL, 1);
INSERT INTO `menu` VALUES (58, 1713290381, 1713290381, 57, 'SystemPlatform', '查询用户', 'A', NULL, NULL, '/manager/v1/users', 'GET', NULL, 'manager:user:query', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (59, 1713290381, 1713290381, 57, 'SystemPlatform', '新增用户', 'A', NULL, NULL, '/manager/v1/user', 'POST', NULL, 'manager:user:add', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (60, 1713290381, 1713290381, 57, 'SystemPlatform', '修改用户', 'A', NULL, NULL, '/manager/v1/user', 'PUT', NULL, 'manager:user:update', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (61, 1713290381, 1713290381, 60, 'SystemPlatform', '获取用户角色列表', 'A', NULL, NULL, '/manager/v1/user/roles', 'GET', NULL, 'manager:user:update', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (62, 1713290381, 1713290381, 60, 'SystemPlatform', '获取用户部门列表', 'A', NULL, NULL, '/manager/v1/user/jobs', 'POST', NULL, 'manager:user:update', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (63, 1713290381, 1713290381, 57, 'SystemPlatform', '删除用户', 'A', NULL, NULL, '/manager/v1/user', 'DELETE', NULL, 'manager:user:delete', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (64, 1713290381, 1713290381, 57, 'SystemPlatform', '启用/禁用用户', 'G', NULL, NULL, NULL, NULL, NULL, 'manager:user:status', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (65, 1713290381, 1713290381, 64, 'SystemPlatform', '启用用户', 'A', NULL, NULL, '/manager/v1/user/enable', 'POST', NULL, 'manager:user:status:enable', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (66, 1713290381, 1713290381, 64, 'SystemPlatform', '禁用用户', 'A', NULL, NULL, '/manager/v1/user/disable', 'POST', NULL, 'manager:user:status:disable', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (67, 1713290381, 1713290381, 57, 'SystemPlatform', '重置账号密码', 'A', NULL, NULL, '/manager/v1/user/password/reset', 'POST', NULL, 'manager:user:reset:password', NULL, NULL, '', NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (68, 1713290381, 1713290381, 1, 'SystemPlatform', '配置中心', 'M', 'SystemPlatformConfigure', 'code-block', NULL, NULL, '/configure', NULL, NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (69, 1713290381, 1713290381, 68, 'SystemPlatform', '环境管理', 'M', 'ConfigureEnv', 'common', NULL, NULL, '/configure/env', NULL, NULL, NULL, '/configure/env/index', NULL, 0, NULL, 1, NULL, 1);
INSERT INTO `menu` VALUES (70, 1713290381, 1713290381, 69, 'SystemPlatform', '查看环境', 'A', NULL, NULL, '/configure/v1/envs', 'GET', NULL, 'configure:env:query', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (71, 1713290381, 1713290381, 69, 'SystemPlatform', '新增环境', 'A', NULL, NULL, '/configure/v1/env', 'POST', NULL, 'configure:env:add', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (72, 1713290381, 1713290381, 69, 'SystemPlatform', '修改环境', 'A', NULL, NULL, '/configure/v1/env', 'PUT', NULL, 'configure:env:update', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (73, 1713290381, 1713290381, 69, 'SystemPlatform', '删除环境', 'A', NULL, NULL, '/configure/v1/env', 'DELETE', NULL, 'configure:env:delete', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (74, 1713290381, 1713290381, 69, 'SystemPlatform', '查看环境Token', 'A', NULL, NULL, '/configure/v1/env/token', 'GET', NULL, 'configure:env:token:query', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (75, 1713290381, 1713290381, 69, 'SystemPlatform', '重置环境token', 'A', NULL, NULL, '/configure/v1/env/token', 'PUT', NULL, 'configure:env:token:reset', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (76, 1713290381, 1713290381, 68, 'SystemPlatform', '服务管理', 'M', 'ConfigureServer', 'apps', NULL, NULL, '/configure/server', NULL, NULL, NULL, '/configure/server/index', NULL, 0, NULL, 1, NULL, 1);
INSERT INTO `menu` VALUES (77, 1713290381, 1713290381, 76, 'SystemPlatform', '查询服务', 'A', NULL, NULL, '/configure/v1/servers', 'GET', NULL, 'configure:server:query', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (78, 1713290381, 1713290381, 76, 'SystemPlatform', '新增服务', 'A', NULL, NULL, '/configure/v1/server', 'POST', NULL, 'configure:server:add', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (79, 1713290381, 1713290381, 76, 'SystemPlatform', '修改服务', 'A', NULL, NULL, '/configure/v1/server', 'PUT', NULL, 'configure:server:update', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (80, 1713290381, 1713290381, 76, 'SystemPlatform', '删除服务', 'A', NULL, NULL, '/configure/v1/server', 'DELETE', NULL, 'configure:server:delete', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (81, 1713290381, 1713290381, 68, 'SystemPlatform', '资源变量', 'M', 'ConfigureResource', 'unordered-list', NULL, NULL, '/configure/resource', NULL, NULL, NULL, '/configure/resource/index', NULL, 0, NULL, 1, NULL, 1);
INSERT INTO `menu` VALUES (82, 1713290381, 1713290381, 81, 'SystemPlatform', '查看资源', 'A', NULL, NULL, '/configure/v1/resources', 'GET', NULL, 'configure:resource:query', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (83, 1713290381, 1713290381, 81, 'SystemPlatform', '新增资源', 'A', NULL, NULL, '/configure/v1/resource', 'POST', NULL, 'configure:resource:add', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (84, 1713290381, 1713290381, 81, 'SystemPlatform', '修改资源', 'A', NULL, NULL, '/configure/v1/resource', 'PUT', NULL, 'configure:resource:update', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (85, 1713290381, 1713290381, 81, 'SystemPlatform', '删除资源', 'A', NULL, NULL, '/configure/v1/resource', 'DELETE', NULL, 'configure:resource:delete', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (86, 1713290381, 1713290381, 81, 'SystemPlatform', '资源变量值配置', 'G', NULL, NULL, NULL, NULL, NULL, 'configure:resource:value', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (87, 1713290381, 1713290381, 68, 'SystemPlatform', '业务变量', 'M', 'ConfigureBusiness', 'code', NULL, NULL, '/configure/business', NULL, NULL, NULL, '/configure/business/index', NULL, 0, NULL, 1, NULL, 1);
INSERT INTO `menu` VALUES (88, 1713290381, 1713290381, 87, 'SystemPlatform', '查看业务变量', 'A', NULL, NULL, '/configure/v1/business', 'GET', NULL, 'configure:business:query', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (89, 1713290381, 1713290381, 87, 'SystemPlatform', '新增业务变量', 'A', NULL, NULL, '/configure/v1/business', 'POST', NULL, 'configure:business:add', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (90, 1713290381, 1713290381, 87, 'SystemPlatform', '修改业务变量', 'A', NULL, NULL, '/configure/v1/business', 'PUT', NULL, 'configure:business:update', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (91, 1713290381, 1713290381, 87, 'SystemPlatform', '删除业务变量', 'A', NULL, NULL, '/configure/v1/business', 'DELETE', NULL, 'configure:business:delete', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (92, 1713290381, 1713290381, 87, 'SystemPlatform', '配置业务变量值', 'A', NULL, NULL, '/configure/business/value', 'PUT', NULL, 'configure:business:value', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (93, 1713290381, 1713290381, 68, 'SystemPlatform', '配置模板', 'M', 'ConfgureTemplate', 'file', NULL, NULL, '/configure/template', NULL, NULL, NULL, '/configure/template/index', NULL, 0, NULL, 1, NULL, 1);
INSERT INTO `menu` VALUES (94, 1713290381, 1713290381, 93, 'SystemPlatform', '模板管理', 'G', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (95, 1713290381, 1713290381, 94, 'SystemPlatform', '提交模板', 'A', NULL, NULL, '/configure/v1/template', 'POST', NULL, 'configure:template:add', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (96, 1713290381, 1713290381, 94, 'SystemPlatform', '查看模板历史版本', 'A', NULL, NULL, '/configure/v1/templates', 'GET', NULL, 'configure:template:history', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (97, 1713290381, 1713290381, 94, 'SystemPlatform', '模板对比', 'A', NULL, NULL, '/configure/v1/template/compare', 'POST', NULL, 'configure:template:compare', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (98, 1713290381, 1713290381, 94, 'SystemPlatform', '切换模板', 'A', NULL, NULL, '/configure/v1/template/switch', 'POST', NULL, 'configure:template:switch', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (99, 1713290381, 1713290381, 93, 'SystemPlatform', '配置管理', 'G', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (100, 1713290381, 1713290381, 99, 'SystemPlatform', '配置对比', 'A', NULL, NULL, '/configure/v1/configure/compare', 'POST', NULL, 'configure:configure:compare', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (101, 1713290381, 1713290381, 99, 'SystemPlatform', '配置预览', 'A', NULL, NULL, '/configure/v1/template/preview', 'POST', NULL, 'configure:template:preview', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (102, 1713290381, 1713290381, 99, 'SystemPlatform', '同步配置', 'A', NULL, NULL, '/configure/v1/configure', 'PUT', NULL, 'configure:configure:sync', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (103, 1713290381, 1713290381, 1, 'SystemPlatform', '资源中心', 'M', 'SystemPlatformResource', 'file', NULL, NULL, '/resource', NULL, NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (104, 1713290381, 1713290381, 103, 'SystemPlatform', '文件管理', 'M', 'File', 'file', NULL, NULL, '/resource/file', NULL, NULL, NULL, '/resource/file/index', NULL, 0, 0, 1, NULL, 1);
INSERT INTO `menu` VALUES (105, 1713290381, 1713290381, 104, 'SystemPlatform', '目录管理', 'G', NULL, NULL, NULL, NULL, NULL, 'resource:directory:group', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (106, 1713290381, 1713290381, 105, 'SystemPlatform', '查看目录', 'A', NULL, NULL, '/resource/v1/directory', 'GET', NULL, 'resource:directory:query', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (107, 1713290381, 1713290381, 105, 'SystemPlatform', '新增目录', 'A', NULL, NULL, '/resource/v1/directory', 'POST', NULL, 'resource:directory:add', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (108, 1713290381, 1713290381, 105, 'SystemPlatform', '修改目录', 'A', NULL, NULL, '/resource/v1/directory', 'PUT', NULL, 'resource:directory:update', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (109, 1713290381, 1713290381, 105, 'SystemPlatform', '删除目录', 'A', NULL, NULL, '/resource/v1/directory', 'DELETE', NULL, 'resource:directory:delete', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (110, 1713290381, 1713290381, 104, 'SystemPlatform', '文件管理', 'G', NULL, NULL, NULL, NULL, NULL, 'resource:file:group', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (111, 1713290381, 1713290381, 110, 'SystemPlatform', '查看文件', 'A', NULL, NULL, '/resource/v1/files', 'GET', NULL, 'resource:file:query', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (112, 1713290381, 1713290381, 110, 'SystemPlatform', '修改文件', 'A', NULL, NULL, '/resource/v1/file', 'PUT', NULL, 'resource:file:upload', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (113, 1713290381, 1713290381, 110, 'SystemPlatform', '删除文件', 'A', NULL, NULL, '/resource/v1/file', 'DELETE', NULL, 'resource:file:delete', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (114, 1713290381, 1713290381, 110, 'SystemPlatform', '上传文件', 'G', NULL, NULL, NULL, NULL, NULL, 'resource:file:upload:group', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (115, 1713290381, 1713290381, 114, 'SystemPlatform', '分块上传文件', 'A', NULL, NULL, '/resource/v1/upload', 'POST', NULL, 'resource:file:upload', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (116, 1713290381, 1713290381, 114, 'SystemPlatform', '预上传文件', 'A', NULL, NULL, '/resource/v1/upload/prepare', 'POST', NULL, 'resource:file:upload', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (117, 1713290381, 1713290381, 1, 'SystemPlatform', '定时任务', 'M', 'SystemPlatformCron', 'schedule', NULL, NULL, '/cron', NULL, NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (118, 1713290381, 1713290381, 117, 'SystemPlatform', '节点管理', 'M', 'Worker', 'common', NULL, NULL, '/cron/worker', NULL, NULL, NULL, '/cron/worker/index', NULL, 0, 0, 1, NULL, 1);
INSERT INTO `menu` VALUES (119, 1713290381, 1713290381, 118, 'SystemPlatform', '查看节点分组', 'A', NULL, NULL, '/cron/v1/worker/groups', 'GET', NULL, 'cron:worker:group:query', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (120, 1713290381, 1713290381, 118, 'SystemPlatform', '新增节点分组', 'A', NULL, NULL, '/cron/v1/worker/group', 'POST', NULL, 'cron:worker:group:add', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (121, 1713290381, 1713290381, 118, 'SystemPlatform', '修改节点分组', 'A', NULL, NULL, '/cron/v1/worker/group', 'PUT', NULL, 'cron:worker:group:update', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (122, 1713290381, 1713290381, 118, 'SystemPlatform', '删除节点分组', 'A', NULL, NULL, '/cron/v1/worker/group', 'DELETE', NULL, 'cron:worker:group:delete', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (123, 1713290381, 1713290381, 118, 'SystemPlatform', '查看节点', 'A', NULL, NULL, '/cron/v1/workers', 'GET', NULL, 'cron:worker:query', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (124, 1713290381, 1713290381, 118, 'SystemPlatform', '新增节点', 'A', NULL, NULL, '/cron/v1/worker', 'POST', NULL, 'cron:worker:add', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (125, 1713290381, 1713290381, 118, 'SystemPlatform', '修改节点', 'A', NULL, NULL, '/cron/v1/worker', 'PUT', NULL, 'cron:worker:update', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (126, 1713290381, 1713290381, 118, 'SystemPlatform', '删除节点', 'A', NULL, NULL, '/cron/v1/worker', 'DELETE', NULL, 'cron:worker:delete', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (127, 1713290381, 1713290381, 118, 'SystemPlatform', '节点状态管理', 'G', NULL, NULL, NULL, NULL, NULL, 'cron:worker:status', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (128, 1713290381, 1713290381, 127, 'SystemPlatform', '启用节点', 'A', NULL, NULL, '/cron/v1/worker/enable', 'POST', NULL, NULL, NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (129, 1713290381, 1713290381, 127, 'SystemPlatform', '启用节点', 'A', NULL, NULL, '/cron/v1/worker/enable', 'POST', NULL, NULL, NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (130, 1713290381, 1713290381, 117, 'SystemPlatform', '任务管理', 'M', 'Task', 'computer', NULL, NULL, '/cron/task', NULL, NULL, NULL, '/cron/task/index', NULL, 0, 0, 1, NULL, 1);
INSERT INTO `menu` VALUES (131, 1713290381, 1713290381, 130, 'SystemPlatform', '查看任务分组', 'A', NULL, NULL, '/cron/v1/task/groups', 'GET', NULL, 'cron:task:group:query', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (132, 1713290381, 1713290381, 130, 'SystemPlatform', '新增任务分组', 'A', NULL, NULL, '/cron/v1/task/group', 'POST', NULL, 'cron:task:group:add', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (133, 1713290381, 1713290381, 130, 'SystemPlatform', '修改任务分组', 'A', NULL, NULL, '/cron/v1/task/group', 'PUT', NULL, 'cron:task:group:update', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (134, 1713290381, 1713290381, 130, 'SystemPlatform', '删除任务分组', 'A', NULL, NULL, '/cron/v1/task/group', 'DELETE', NULL, 'cron:task:group:delete', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (135, 1713290381, 1713290381, 130, 'SystemPlatform', '查看任务', 'A', NULL, NULL, '/cron/v1/tasks', 'GET', NULL, 'cron:task:query', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (136, 1713290381, 1713290381, 130, 'SystemPlatform', '新增任务', 'A', NULL, NULL, '/cron/v1/task', 'POST', NULL, 'cron:task:add', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (137, 1713290381, 1713290381, 130, 'SystemPlatform', '立即执行任务', 'A', NULL, NULL, '/cron/v1/task/exec', 'POST', NULL, 'cron:task:exec', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (138, 1713290381, 1713290381, 130, 'SystemPlatform', '取消执行任务', 'A', NULL, NULL, '/cron/v1/task/cancel', 'POST', NULL, 'cron:task:cancel', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (139, 1713290381, 1713290381, 130, 'SystemPlatform', '修改任务', 'A', NULL, NULL, '/cron/v1/task', 'PUT', NULL, 'cron:task:update', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (140, 1713290381, 1713290381, 130, 'SystemPlatform', '删除任务', 'A', NULL, NULL, '/cron/v1/task', 'DELETE', NULL, 'cron:task:delete', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (141, 1713290381, 1713290381, 130, 'SystemPlatform', '任务状态管理', 'G', NULL, NULL, NULL, NULL, NULL, 'cron:task:status', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (142, 1713290381, 1713290381, 141, 'SystemPlatform', '启用任务', 'A', NULL, NULL, '/cron/v1/task/enable', 'POST', NULL, NULL, NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (143, 1713290381, 1713290381, 141, 'SystemPlatform', '禁用任务', 'A', NULL, NULL, '/cron/v1/task/disable', 'POST', NULL, NULL, NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (144, 1713290381, 1713290381, 130, 'SystemPlatform', '任务日志', 'G', NULL, NULL, NULL, NULL, NULL, 'cron:task:log', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (145, 1713290381, 1713290381, 144, 'SystemPlatform', '获取任务日志分页', 'A', NULL, NULL, '/cron/v1/logs', 'GET', NULL, NULL, NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (146, 1713290381, 1713290381, 144, 'SystemPlatform', '获取任务日志详情', 'A', NULL, NULL, '/cron/v1/log', 'GET', NULL, NULL, NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (147, 1713290381, 1713290381, 1, 'SystemPlatform', '用户中心', 'M', 'SystemPlatformUserCenter', 'user', NULL, NULL, '/user-center', NULL, NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (148, 1713290381, 1713290381, 147, 'SystemPlatform', '协议管理', 'M', 'UserCenterAgreement', 'bookmark', NULL, NULL, '/user-center/agreement', NULL, NULL, NULL, NULL, NULL, 0, 0, 1, NULL, 1);
INSERT INTO `menu` VALUES (149, 1713290381, 1713290381, 148, 'SystemPlatform', '协议内容', 'M', 'UserCenterAgreementContent', 'bookmark', NULL, NULL, '/user-center/agreement/content', NULL, NULL, NULL, '/user-center/agreement/content/index', NULL, 0, 0, 1, NULL, 1);
INSERT INTO `menu` VALUES (150, 1713290381, 1713290381, 149, 'SystemPlatform', '查看协议列表', 'A', NULL, NULL, '/user-center/v1/agreement/contents', 'GET', NULL, 'uc:agreement:content:query', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (151, 1713290381, 1713290381, 149, 'SystemPlatform', '查看协议详细信息', 'A', NULL, NULL, '/user-center/v1/agreement/content', 'GET', NULL, 'uc:agreement:content:query', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (152, 1713290381, 1713290381, 149, 'SystemPlatform', '新增协议', 'A', NULL, NULL, '/user-center/v1/agreement/content', 'POST', NULL, 'uc:agreement:content:add', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (153, 1713290381, 1713290381, 149, 'SystemPlatform', '修改协议', 'A', NULL, NULL, '/user-center/v1/agreement/content', 'PUT', NULL, 'uc:agreement:content:update', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (154, 1713290381, 1713290381, 149, 'SystemPlatform', '删除协议', 'A', NULL, NULL, '/user-center/v1/agreement/content', 'DELETE', NULL, 'uc:agreement:content:delete', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (155, 1713290381, 1713290381, 148, 'SystemPlatform', '协议场景', 'M', 'UserCenterAgreementScene', 'common', NULL, NULL, '/user-center/agreement/scene', NULL, NULL, NULL, '/user-center/agreement/scene/index', NULL, 0, 0, 1, NULL, 1);
INSERT INTO `menu` VALUES (156, 1713290381, 1713290381, 155, 'SystemPlatform', '查看场景列表', 'A', NULL, NULL, '/user-center/v1/agreement/scenes', 'GET', NULL, 'uc:agreement:scene:query', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (157, 1713290381, 1713290381, 155, 'SystemPlatform', '查看场景详细信息', 'A', NULL, NULL, '/user-center/v1/agreement/scene', 'GET', NULL, 'uc:agreement:scene:query', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (158, 1713290381, 1713290381, 155, 'SystemPlatform', '新增场景', 'A', NULL, NULL, '/user-center/v1/agreement/scene', 'POST', NULL, 'uc:agreement:scene:add', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (159, 1713290381, 1713290381, 155, 'SystemPlatform', '修改场景', 'A', NULL, NULL, '/user-center/v1/agreement/scene', 'PUT', NULL, 'uc:agreement:scene:update', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (160, 1713290381, 1713290381, 155, 'SystemPlatform', '删除场景', 'A', NULL, NULL, '/user-center/v1/agreement/scene', 'DELETE', NULL, 'uc:agreement:scene:delete', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (161, 1713290381, 1713290381, 147, 'SystemPlatform', '授权渠道', 'M', 'UserCenterChannel', 'mind-mapping', NULL, NULL, '/user-center/channel', NULL, NULL, NULL, '/user-center/channel/index', NULL, 0, 0, 1, NULL, 1);
INSERT INTO `menu` VALUES (162, 1713290381, 1713290381, 161, 'SystemPlatform', '查看渠道', 'A', NULL, NULL, '/user-center/v1/channels', 'GET', NULL, 'uc:channel:query', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (163, 1713290381, 1713290381, 161, 'SystemPlatform', '新增渠道', 'A', NULL, NULL, '/user-center/v1/channel', 'POST', NULL, 'uc:channel:add', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (164, 1713290381, 1713290381, 161, 'SystemPlatform', '修改渠道', 'A', NULL, NULL, '/user-center/v1/channel', 'PUT', NULL, 'uc:channel:update', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (165, 1713290381, 1713290381, 161, 'SystemPlatform', '删除渠道', 'A', NULL, NULL, '/user-center/v1/channel', 'DELETE', NULL, 'uc:channel:delete', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (166, 1713290381, 1713290381, 147, 'SystemPlatform', '信息字段', 'M', 'UserCenterField', 'list', NULL, NULL, '/user-center/field', NULL, NULL, NULL, '/user-center/field/index', NULL, 0, 0, 1, NULL, 1);
INSERT INTO `menu` VALUES (167, 1713290381, 1713290381, 166, 'SystemPlatform', '查看字段', 'A', NULL, NULL, '/user-center/v1/fields', 'GET', NULL, 'uc:field:query', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (168, 1713290381, 1713290381, 166, 'SystemPlatform', '新增字段', 'A', NULL, NULL, '/user-center/v1/field', 'POST', NULL, 'uc:field:add', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (169, 1713290381, 1713290381, 166, 'SystemPlatform', '修改字段', 'A', NULL, NULL, '/user-center/v1/field', 'PUT', NULL, 'uc:field:update', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (170, 1713290381, 1713290381, 166, 'SystemPlatform', '删除字段', 'A', NULL, NULL, '/user-center/v1/field', 'DELETE', NULL, 'uc:field:delete', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (171, 1713290381, 1713290381, 147, 'SystemPlatform', '应用管理', 'M', 'UserCenterApp', 'apps', NULL, NULL, '/user-center/app', NULL, NULL, NULL, '/user-center/app/index', NULL, 0, 0, 1, NULL, 1);
INSERT INTO `menu` VALUES (172, 1713290381, 1713290381, 171, 'SystemPlatform', '查看应用', 'A', NULL, NULL, '/user-center/v1/apps', 'GET', NULL, 'uc:app:query', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (173, 1713290381, 1713290381, 171, 'SystemPlatform', '新增应用', 'A', NULL, NULL, '/user-center/v1/app', 'POST', NULL, 'uc:app:add', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (174, 1713290381, 1713290381, 171, 'SystemPlatform', '修改应用', 'A', NULL, NULL, '/user-center/v1/app', 'PUT', NULL, 'uc:app:update', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (175, 1713290381, 1713290381, 171, 'SystemPlatform', '删除应用', 'A', NULL, NULL, '/user-center/v1/app', 'DELETE', NULL, 'uc:app:delete', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (176, 1713290381, 1713290381, 147, 'SystemPlatform', '用户管理', 'M', 'user-center', 'user', NULL, NULL, '/user-center/user', NULL, NULL, NULL, '/user-center/user/index', NULL, 0, 0, 1, NULL, 1);
INSERT INTO `menu` VALUES (177, 1713290381, 1713290381, 176, 'SystemPlatform', '查看用户', 'A', NULL, NULL, '/user-center/v1/users', 'GET', NULL, 'uc:user:query', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (178, 1713290381, 1713290381, 176, 'SystemPlatform', '新增用户', 'A', NULL, NULL, '/user-center/v1/user', 'POST', NULL, 'uc:user:add', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (179, 1713290381, 1713290381, 176, 'SystemPlatform', '导入用户', 'A', NULL, NULL, '/user-center/v1/users', 'POST', NULL, 'uc:user:import', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (180, 1713290381, 1713290381, 176, 'SystemPlatform', '修改用户', 'A', NULL, NULL, '/user-center/v1/user', 'PUT', NULL, 'uc:user:update', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (181, 1713290381, 1713290381, 176, 'SystemPlatform', '删除用户', 'A', NULL, NULL, '/user-center/v1/user', 'DELETE', NULL, 'uc:user:delete', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (182, 1713290381, 1713290381, 176, 'SystemPlatform', '下线用户', 'A', NULL, NULL, '/user-center/v1/user/offline', 'POST', NULL, 'uc:user:offline', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (183, 1713290381, 1713290381, 176, 'SystemPlatform', '用户状态管理', 'G', NULL, NULL, NULL, NULL, NULL, 'uc:user:status', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (184, 1713290381, 1713290381, 183, 'SystemPlatform', '禁用用户', 'A', NULL, NULL, '/user-center/v1/user/disable', 'POST', NULL, 'uc:user:disable', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (185, 1713290381, 1713290381, 183, 'SystemPlatform', '启用用户', 'A', NULL, NULL, '/user-center/v1/user/enable', 'POST', NULL, 'uc:user:enable', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (186, 1713290381, 1713290381, 0, 'AppPlatform', '应用平台', 'R', 'AppPlatform', 'apps', NULL, NULL, '/app', NULL, NULL, NULL, 'Layout', NULL, 2, 0, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (187, 1713290381, 1713290381, 186, 'AppPlatform', '首页面板', 'M', 'AppDashboard', 'dashboard', NULL, NULL, '/app/dashboard', NULL, NULL, NULL, '/dashboard/index', NULL, 0, 0, 1, 1, 1);
INSERT INTO `menu` VALUES (188, 1713290381, 1713290381, 186, 'AppPlatform', '信号灯系统', 'M', 'PartyAffairs', 'apps', NULL, NULL, '/party-affairs', NULL, NULL, NULL, NULL, NULL, 0, 0, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (189, 1713290381, 1713290381, 188, 'AppPlatform', '轮播管理', 'M', 'PartyAffairsBanner', 'list', NULL, NULL, '/party-affairs/banner', NULL, NULL, NULL, '/party-affairs/banner/index', NULL, 0, NULL, 1, NULL, 1);
INSERT INTO `menu` VALUES (190, 1713290381, 1713290381, 189, 'AppPlatform', '查看轮播', 'A', NULL, NULL, '/party-affairs/v1/banners', 'GET', NULL, 'party-affairs:banner:query', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (191, 1713290381, 1713290381, 189, 'AppPlatform', '新增轮播', 'A', NULL, NULL, '/party-affairs/v1/banner', 'POST', NULL, 'party-affairs:banner:add', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (192, 1713290381, 1713290381, 189, 'AppPlatform', '修改轮播', 'A', NULL, NULL, '/party-affairs/v1/banner', 'PUT', NULL, 'party-affairs:banner:update', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (193, 1713290381, 1713290381, 189, 'AppPlatform', '删除轮播', 'A', NULL, NULL, '/party-affairs/v1/banner', 'DELETE', NULL, 'party-affairs:banner:delete', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (194, 1713290381, 1713290381, 188, 'AppPlatform', '通知管理', 'M', 'PartyAffairsNotice', 'notification', NULL, NULL, '/party-affairs/notice', NULL, NULL, NULL, '/party-affairs/notice/index', NULL, 0, NULL, 1, NULL, 1);
INSERT INTO `menu` VALUES (195, 1713290381, 1713290381, 194, 'AppPlatform', '查看通知', 'A', NULL, NULL, '/party-affairs/v1/notices', 'GET', NULL, 'party-affairs:notice:query', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (196, 1713290381, 1713290381, 194, 'AppPlatform', '查看通知阅读情况', 'A', NULL, NULL, '/party-affairs/v1/notice/users', 'GET', NULL, 'party-affairs:notice:user:query', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (197, 1713290381, 1713290381, 194, 'AppPlatform', '新增通知', 'A', NULL, NULL, '/party-affairs/v1/notice', 'POST', NULL, 'party-affairs:notice:add', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (198, 1713290381, 1713290381, 194, 'AppPlatform', '修改通知', 'A', NULL, NULL, '/party-affairs/v1/notice', 'PUT', NULL, 'party-affairs:notice:update', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (199, 1713290381, 1713290381, 194, 'AppPlatform', '删除通知', 'A', NULL, NULL, '/party-affairs/v1/notice', 'DELETE', NULL, 'party-affairs:notice:delete', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (200, 1713290381, 1713290381, 188, 'AppPlatform', '新闻管理', 'M', 'PartyAffairsNews', 'book', NULL, NULL, '/party-affairs/news', NULL, NULL, NULL, NULL, NULL, 0, 0, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (201, 1713290381, 1713290381, 200, 'AppPlatform', '板块管理', 'M', 'PartyAffairsNewsClassify', 'menu', NULL, NULL, '/party-affairs/news/classify', NULL, NULL, NULL, '/party-affairs/news/classify/index', NULL, 0, NULL, 1, NULL, 1);
INSERT INTO `menu` VALUES (202, 1713290381, 1713290381, 201, 'AppPlatform', '查看板块', 'A', NULL, NULL, '/party-affairs/v1/news/classify', 'GET', NULL, 'party-affairs:news:classify:query', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (203, 1713290381, 1713290381, 201, 'AppPlatform', '新增板块', 'A', NULL, NULL, '/party-affairs/v1/news/classify', 'POST', NULL, 'party-affairs:news:classify:add', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (204, 1713290381, 1713290381, 201, 'AppPlatform', '修改板块', 'A', NULL, NULL, '/party-affairs/v1/news/classify', 'PUT', NULL, 'party-affairs:news:classify:update', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (205, 1713290381, 1713290381, 201, 'AppPlatform', '删除板块', 'A', NULL, NULL, '/party-affairs/v1/news/classify', 'DELETE', NULL, 'party-affairs:news:classify:delete', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (206, 1713290381, 1713290381, 200, 'AppPlatform', '内容管理', 'M', 'PartyAffairsNewsContent', 'book', NULL, NULL, '/party-affairs/news/content', NULL, NULL, NULL, '/party-affairs/news/content/index', NULL, 0, NULL, 1, NULL, 1);
INSERT INTO `menu` VALUES (207, 1713290381, 1713290381, 206, 'AppPlatform', '查看内容列表', 'A', NULL, NULL, '/party-affairs/v1/news/contents', 'GET', NULL, 'party-affairs:news:content:query', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (208, 1713290381, 1713290381, 206, 'AppPlatform', '查看指定内容', 'A', NULL, NULL, '/party-affairs/v1/news/content', 'GET', NULL, 'party-affairs:news:content:info', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (209, 1713290381, 1713290381, 206, 'AppPlatform', '新增内容', 'A', NULL, NULL, '/party-affairs/v1/news/content', 'POST', NULL, 'party-affairs:news:content:add', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (210, 1713290381, 1713290381, 206, 'AppPlatform', '修改内容', 'A', NULL, NULL, '/party-affairs/v1/news/content', 'PUT', NULL, 'party-affairs:news:content:update', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (211, 1713290381, 1713290381, 206, 'AppPlatform', '删除内容', 'A', NULL, NULL, '/party-affairs/v1/news/content', 'DELETE', NULL, 'party-affairs:news:content:delete', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (212, 1713290381, 1713290381, 200, 'AppPlatform', '评论管理', 'M', 'PartyAffairsNewsComment', 'book', NULL, NULL, '/party-affairs/news/comment', NULL, NULL, NULL, '/party-affairs/news/comment/index', NULL, 0, 1, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (213, 1713290381, 1713290381, 212, 'AppPlatform', '查看评论列表', 'A', NULL, NULL, '/party-affairs/v1/news/comments', 'GET', NULL, 'party-affairs:news:comment:query', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (214, 1713290381, 1713290381, 212, 'AppPlatform', '删除评论', 'A', NULL, NULL, '/party-affairs/v1/news/comment', 'DELETE', NULL, 'party-affairs:news:comment:delete', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (215, 1713290381, 1713290381, 188, 'AppPlatform', '资料管理', 'M', 'PartyAffairsResource', 'public', NULL, NULL, '/party-affairs/resource', NULL, NULL, NULL, NULL, NULL, 0, 0, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (216, 1713290381, 1713290381, 215, 'AppPlatform', '板块管理', 'M', 'PartyAffairsResourceClassify', 'menu', NULL, NULL, '/party-affairs/resource/classify', NULL, NULL, NULL, '/party-affairs/resource/classify/index', NULL, 0, NULL, 1, NULL, 1);
INSERT INTO `menu` VALUES (217, 1713290381, 1713290381, 216, 'AppPlatform', '查看板块', 'A', NULL, NULL, '/party-affairs/v1/resource/classify', 'GET', NULL, 'party-affairs:resource:classify:query', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (218, 1713290381, 1713290381, 216, 'AppPlatform', '新增板块', 'A', NULL, NULL, '/party-affairs/v1/resource/classify', 'POST', NULL, 'party-affairs:resource:classify:add', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (219, 1713290381, 1713290381, 216, 'AppPlatform', '修改板块', 'A', NULL, NULL, '/party-affairs/v1/resource/classify', 'PUT', NULL, 'party-affairs:resource:classify:update', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (220, 1713290381, 1713290381, 216, 'AppPlatform', '删除板块', 'A', NULL, NULL, '/party-affairs/v1/resource/classify', 'DELETE', NULL, 'party-affairs:resource:classify:delete', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (221, 1713290381, 1713290381, 215, 'AppPlatform', '内容管理', 'M', 'PartyAffairsResourceContent', 'public', NULL, NULL, '/party-affairs/resource/content', NULL, NULL, NULL, '/party-affairs/resource/content/index', NULL, 0, NULL, 1, NULL, 1);
INSERT INTO `menu` VALUES (222, 1713290381, 1713290381, 221, 'AppPlatform', '查看内容列表', 'A', NULL, NULL, '/party-affairs/v1/resource/contents', 'GET', NULL, 'party-affairs:resource:content:query', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (223, 1713290381, 1713290381, 221, 'AppPlatform', '新增内容', 'A', NULL, NULL, '/party-affairs/v1/resource/content', 'POST', NULL, 'party-affairs:resource:content:add', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (224, 1713290381, 1713290381, 221, 'AppPlatform', '修改内容', 'A', NULL, NULL, '/party-affairs/v1/resource/content', 'PUT', NULL, 'party-affairs:resource:content:update', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (225, 1713290381, 1713290381, 221, 'AppPlatform', '删除内容', 'A', NULL, NULL, '/party-affairs/v1/resource/content', 'DELETE', NULL, 'party-affairs:resource:content:delete', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (226, 1713290381, 1713290381, 188, 'AppPlatform', '任务管理', 'M', 'PartyAffairsTask', 'layers', NULL, NULL, '/party-affairs/task', NULL, NULL, NULL, NULL, NULL, 0, 0, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (227, 1713290381, 1713290381, 226, 'AppPlatform', '任务配置', 'M', 'PartyAffairsTaskConfigure', 'unordered-list', NULL, NULL, '/party-affairs/task/configure', NULL, NULL, NULL, '/party-affairs/task/configure/index', NULL, 0, NULL, 1, NULL, 1);
INSERT INTO `menu` VALUES (228, 1713290381, 1713290381, 227, 'AppPlatform', '查看配置', 'A', NULL, NULL, '/party-affairs/v1/task/configure', 'GET', NULL, 'party-affairs:task:configure:query', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (229, 1713290381, 1713290381, 227, 'AppPlatform', '新增配置', 'A', NULL, NULL, '/party-affairs/v1/task/configure', 'POST', NULL, 'party-affairs:task:configure:add', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (230, 1713290381, 1713290381, 227, 'AppPlatform', '修改配置', 'A', NULL, NULL, '/party-affairs/v1/task/configure', 'PUT', NULL, 'party-affairs:task:configure:update', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (231, 1713290381, 1713290381, 227, 'AppPlatform', '删除配置', 'A', NULL, NULL, '/party-affairs/v1/task/configure', 'DELETE', NULL, 'party-affairs:task:configure:delete', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (232, 1713290381, 1713290381, 226, 'AppPlatform', '任务统计', 'M', 'PartyAffairsTaskContent', 'computer', NULL, NULL, '/party-affairs/task/value', NULL, NULL, NULL, '/party-affairs/task/value/index', NULL, 0, NULL, 1, NULL, 1);
INSERT INTO `menu` VALUES (233, 1713290381, 1713290381, 232, 'AppPlatform', '查看任务填写列表', 'A', NULL, NULL, '/party-affairs/v1/task/values', 'GET', NULL, 'party-affairs:task:value:query', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (234, 1713290381, 1713290381, 232, 'AppPlatform', '查看指定任务填写详情', 'A', NULL, NULL, '/party-affairs/v1/task/value', 'GET', NULL, 'party-affairs:task:value:query', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (235, 1713290381, 1713290381, 232, 'AppPlatform', '修改任务填写详情', 'A', NULL, NULL, '/party-affairs/v1/task/value', 'PUT', NULL, 'party-affairs:task:value:update', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (236, 1713290381, 1713290381, 232, 'AppPlatform', '删除任务填写详情', 'A', NULL, NULL, '/party-affairs/v1/task/value', 'DELETE', NULL, 'party-affairs:task:value:delete', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (237, 1713290381, 1713290381, 188, 'AppPlatform', '视频管理', 'M', 'PartyAffairsVideo', 'video-camera', NULL, NULL, '/party-affairs/video', NULL, NULL, NULL, NULL, NULL, 0, 0, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (238, 1713290381, 1713290381, 237, 'AppPlatform', '板块管理', 'M', 'PartyAffairsVideoClassify', 'menu', NULL, NULL, '/party-affairs/video/classify', NULL, NULL, NULL, '/party-affairs/video/classify/index', NULL, 0, NULL, 1, NULL, 1);
INSERT INTO `menu` VALUES (239, 1713290381, 1713290381, 238, 'AppPlatform', '查看板块', 'A', NULL, NULL, '/party-affairs/v1/video/classify', 'GET', NULL, 'party-affairs:video:classify:query', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (240, 1713290381, 1713290381, 238, 'AppPlatform', '新增板块', 'A', NULL, NULL, '/party-affairs/v1/video/classify', 'POST', NULL, 'party-affairs:video:classify:add', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (241, 1713290381, 1713290381, 238, 'AppPlatform', '修改板块', 'A', NULL, NULL, '/party-affairs/v1/video/classify', 'PUT', NULL, 'party-affairs:video:classify:update', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (242, 1713290381, 1713290381, 238, 'AppPlatform', '删除板块', 'A', NULL, NULL, '/party-affairs/v1/video/classify', 'DELETE', NULL, 'party-affairs:video:classify:delete', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (243, 1713290381, 1713290381, 237, 'AppPlatform', '内容管理', 'M', 'PartyAffairsVideoContent', 'video-camera', NULL, NULL, '/party-affairs/video/content', NULL, NULL, NULL, '/party-affairs/video/content/index', NULL, 0, NULL, 1, NULL, 1);
INSERT INTO `menu` VALUES (244, 1713290381, 1713290381, 243, 'AppPlatform', '查看内容列表', 'A', NULL, NULL, '/party-affairs/v1/video/contents', 'GET', NULL, 'party-affairs:video:content:query', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (245, 1713290381, 1713290381, 243, 'AppPlatform', '新增内容', 'A', NULL, NULL, '/party-affairs/v1/video/content', 'POST', NULL, 'party-affairs:video:content:add', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (246, 1713290381, 1713290381, 243, 'AppPlatform', '修改内容', 'A', NULL, NULL, '/party-affairs/v1/video/content', 'PUT', NULL, 'party-affairs:video:content:update', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (247, 1713290381, 1713290381, 243, 'AppPlatform', '删除内容', 'A', NULL, NULL, '/party-affairs/v1/video/content', 'DELETE', NULL, 'party-affairs:video:content:delete', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (248, 1713290381, 1713290381, 243, 'AppPlatform', '进度统计', 'A', NULL, NULL, '/party-affairs/v1/video/process', 'GET', NULL, 'party-affairs:video:content:process', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (249, 1713290381, 1713290381, 0, 'HunanPlatform', '算法组件', 'M', 'Hunan', 'share-alt', NULL, NULL, '/hunan', NULL, NULL, NULL, 'Layout', NULL, 0, 0, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (250, 1713290381, 1713290381, 249, 'HunanPlatform', '全域价值链生态资源评估', 'M', 'HunanUniverse', 'common', NULL, NULL, '/hunan/universe', NULL, NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, 1);
INSERT INTO `menu` VALUES (251, 1713290381, 1713290381, 250, 'HunanPlatform', '用户偏好采集', 'M', 'HunanUniversePersonality', 'user', NULL, NULL, '/hunan/universe/personality', NULL, NULL, NULL, '/hunan/universe/personality/index', NULL, 0, NULL, 1, NULL, 1);
INSERT INTO `menu` VALUES (252, 1713290381, 1713290381, 250, 'HunanPlatform', '供应商评估结果', 'M', 'HunanUniverseDesigning', 'computer', NULL, NULL, '/hunan/universe/designing', NULL, NULL, NULL, '/hunan/universe/designing/index', NULL, 0, NULL, 1, NULL, 1);
INSERT INTO `menu` VALUES (253, 1713290381, 1713290381, 249, 'HunanPlatform', '生态资源重构', 'M', 'HunanNatural', 'compass', NULL, NULL, '/hunan/natural', NULL, NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, 1);
INSERT INTO `menu` VALUES (254, 1713290381, 1713290381, 253, 'HunanPlatform', '生态资源重构', 'M', 'HunanNaturalPreferences', 'compass', NULL, NULL, '/hunan/natural/preferences', NULL, NULL, NULL, '/hunan/natural/preferences/index', NULL, 0, NULL, 1, NULL, 1);
INSERT INTO `menu` VALUES (255, 1713290381, 1713290381, 253, 'HunanPlatform', '供应商指标变动', 'M', 'HunanNaturalVariable', 'bar-chart', NULL, NULL, '/hunan/natural/variable', NULL, NULL, NULL, '/hunan/natural/variable/index', NULL, 0, NULL, 1, NULL, 1);
INSERT INTO `menu` VALUES (256, 1713290381, 1713290381, 249, 'HunanPlatform', '资源配置', 'M', 'HunanConfig', 'cloud', NULL, NULL, '/hunan/config', NULL, NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, 1);
INSERT INTO `menu` VALUES (257, 1713290381, 1713290381, 256, 'HunanPlatform', '生态资源配置', 'M', 'HunanConfigEcological', 'cloud', NULL, NULL, '/hunan/config/ecological', NULL, NULL, NULL, '/hunan/config/ecological/index', NULL, 0, NULL, 0, NULL, 1);
INSERT INTO `menu` VALUES (258, 1713290381, 1713290381, 249, 'HunanPlatform', '信息管理', 'M', 'HunanAdmin', 'settings', NULL, NULL, '/hunan/admin', NULL, NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, 1);
INSERT INTO `menu` VALUES (259, 1713290381, 1713290381, 258, 'HunanPlatform', '供应商库', 'M', 'HunanAdminMachine', 'user-group', NULL, NULL, '/hunan/admin/machine', NULL, NULL, NULL, '/hunan/admin/machine/index', NULL, 0, NULL, 0, NULL, 1);
INSERT INTO `menu` VALUES (260, 1713290381, 1713290381, 258, 'HunanPlatform', '供应商库（管理员）', 'M', 'HunanAdminMachineCopy', 'user-group', NULL, NULL, '/hunan/admin/machine_copy', NULL, NULL, NULL, '/hunan/admin/machine/index', NULL, 0, NULL, 0, NULL, 1);
INSERT INTO `menu` VALUES (261, 1713290381, 1713290381, 258, 'HunanPlatform', '配置记录', 'M', 'HunanAdminRecord', 'save', NULL, NULL, '/hunan/admin/record', NULL, NULL, NULL, '/hunan/admin/record/index', NULL, 0, NULL, 0, NULL, 1);
INSERT INTO `menu` VALUES (262, 1713290381, 1713290381, 258, 'HunanPlatform', '配置记录（管理员）', 'M', 'HunanAdminRecordCopy', 'save', NULL, NULL, '/hunan/admin/record_copy', NULL, NULL, NULL, '/hunan/admin/record/index', NULL, 0, NULL, 0, NULL, 1);
INSERT INTO `menu` VALUES (263, 1713290381, 1713290381, 258, 'HunanPlatform', '用户管理', 'M', 'HunanAdminUser', 'user', NULL, NULL, '/hunan/admin/user', NULL, NULL, NULL, '/hunan/admin/user/index', NULL, 0, NULL, 0, NULL, 1);
COMMIT;

-- ----------------------------
-- Table structure for object
-- ----------------------------
DROP TABLE IF EXISTS `object`;
CREATE TABLE `object` (
  `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `created_at` bigint DEFAULT NULL COMMENT '创建时间',
  `updated_at` bigint DEFAULT NULL COMMENT '修改时间',
  `keyword` varchar(32) NOT NULL COMMENT '关键字',
  `name` varchar(32) NOT NULL COMMENT '名称',
  `api` varchar(128) NOT NULL COMMENT '请求路径',
  `method` varchar(6) NOT NULL COMMENT '请求方法',
  `params` varchar(512) DEFAULT NULL COMMENT '请求参数',
  `label` varchar(32) NOT NULL COMMENT '标签字段',
  `value` varchar(32) NOT NULL COMMENT '值字段',
  `description` varchar(256) NOT NULL COMMENT '描述',
  PRIMARY KEY (`id`),
  KEY `idx_object_created_at` (`created_at`),
  KEY `idx_object_updated_at` (`updated_at`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='资源对象';

-- ----------------------------
-- Records of object
-- ----------------------------
BEGIN;
INSERT INTO `object` VALUES (1, 1713290381, 1713290381, 'env', '环境资源对象', '/configure/v1/envs', 'GET', '', 'list.name', 'list.id', '环境资源对象');
INSERT INTO `object` VALUES (2, 1713290381, 1713290381, 'server', '服务资源对象', '/configure/v1/servers', 'GET', '{\"page\":1,\"page_size\":50}', 'list.name', 'list.id', '服务资源对象');
INSERT INTO `object` VALUES (3, 1713290381, 1713290381, 'app', '应用资源对象', '/user-center/v1/apps', 'GET', '{\"page\":1,\"page_size\":50}', 'list.name', 'list.id', '应用资源对象');
COMMIT;

-- ----------------------------
-- Table structure for role
-- ----------------------------
DROP TABLE IF EXISTS `role`;
CREATE TABLE `role` (
  `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `created_at` bigint DEFAULT NULL COMMENT '创建时间',
  `updated_at` bigint DEFAULT NULL COMMENT '修改时间',
  `parent_id` int unsigned NOT NULL COMMENT '父id',
  `name` char(64) NOT NULL COMMENT '名称',
  `keyword` char(32) NOT NULL COMMENT '关键字',
  `status` tinyint(1) NOT NULL DEFAULT '0' COMMENT '状态',
  `description` varchar(128) NOT NULL COMMENT '描述',
  `department_ids` varchar(256) DEFAULT NULL COMMENT '自定义管理部门',
  `data_scope` char(32) NOT NULL COMMENT '权限类型',
  PRIMARY KEY (`id`),
  KEY `idx_role_created_at` (`created_at`),
  KEY `idx_role_updated_at` (`updated_at`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='角色信息';

-- ----------------------------
-- Records of role
-- ----------------------------
BEGIN;
INSERT INTO `role` VALUES (1, 1713290381, 1713290381, 0, '超级管理员', 'superAdmin', 1, '超级管理员  ', NULL, 'ALL');
COMMIT;

-- ----------------------------
-- Table structure for role_closure
-- ----------------------------
DROP TABLE IF EXISTS `role_closure`;
CREATE TABLE `role_closure` (
  `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `parent` int unsigned NOT NULL COMMENT '用户id',
  `children` int unsigned NOT NULL COMMENT '用户id',
  PRIMARY KEY (`id`),
  KEY `fk_role_closure_parent_role` (`parent`),
  KEY `fk_role_closure_children_role` (`children`),
  CONSTRAINT `fk_role_closure_children_role` FOREIGN KEY (`children`) REFERENCES `role` (`id`) ON DELETE CASCADE,
  CONSTRAINT `fk_role_closure_parent_role` FOREIGN KEY (`parent`) REFERENCES `role` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='角色层级信息';

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
  `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `created_at` bigint DEFAULT NULL COMMENT '创建时间',
  `role_id` int unsigned NOT NULL COMMENT '角色id',
  `menu_id` int unsigned NOT NULL COMMENT '菜单id',
  PRIMARY KEY (`id`),
  KEY `idx_role_menu_created_at` (`created_at`),
  KEY `fk_role_menu_role` (`role_id`),
  KEY `fk_role_menu_menu` (`menu_id`),
  CONSTRAINT `fk_role_menu_menu` FOREIGN KEY (`menu_id`) REFERENCES `menu` (`id`) ON DELETE CASCADE,
  CONSTRAINT `fk_role_menu_role` FOREIGN KEY (`role_id`) REFERENCES `role` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='角色菜单信息';

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
  `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `created_at` bigint DEFAULT NULL COMMENT '创建时间',
  `updated_at` bigint DEFAULT NULL COMMENT '修改时间',
  `department_id` int unsigned NOT NULL COMMENT '部门id',
  `role_id` int unsigned NOT NULL COMMENT '角色id',
  `name` varchar(32) NOT NULL COMMENT '名称',
  `nickname` varchar(64) NOT NULL COMMENT '昵称',
  `gender` varchar(12) NOT NULL COMMENT '性别',
  `phone` varchar(11) NOT NULL COMMENT '手机号码',
  `password` varchar(256) NOT NULL COMMENT '密码',
  `avatar` varchar(128) DEFAULT NULL COMMENT '头像',
  `email` varchar(64) NOT NULL COMMENT '邮箱',
  `status` tinyint(1) DEFAULT '0' COMMENT '状态',
  `disabled` varchar(64) DEFAULT NULL COMMENT '禁用原因',
  `last_login` bigint DEFAULT NULL COMMENT '禁用原因',
  `token` varchar(512) DEFAULT NULL COMMENT '禁用原因',
  PRIMARY KEY (`id`),
  KEY `idx_user_created_at` (`created_at`),
  KEY `idx_user_updated_at` (`updated_at`),
  KEY `fk_user_role` (`role_id`),
  KEY `fk_user_department` (`department_id`),
  CONSTRAINT `fk_user_department` FOREIGN KEY (`department_id`) REFERENCES `department` (`id`),
  CONSTRAINT `fk_user_role` FOREIGN KEY (`role_id`) REFERENCES `role` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='用户信息';

-- ----------------------------
-- Records of user
-- ----------------------------
BEGIN;
INSERT INTO `user` VALUES (1, 1713290381, 1713324610, 1, 1, '超级管理员', '超级管理员', 'M', '18888888888', '$2a$10$tF5cAQLfgmlCzA2T3YzEIOl0iBPVs7f8WEwFICQZNU0QnGFd1/4n.', '', '1280291001@qq.com', 1, NULL, 1713324610, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJkZXBhcnRtZW50X2lkIjoxLCJkZXBhcnRtZW50X2tleXdvcmQiOiJjb21wYW55IiwiZXhwIjoxNzEzMzMxODExLCJpYXQiOjE3MTMzMjQ2MTAsIm5iZiI6MTcxMzMyNDYxMCwicm9sZV9pZCI6MSwicm9sZV9rZXl3b3JkIjoic3VwZXJBZG1pbiIsInVzZXJfaWQiOjF9.ZhvwqM5PejNsSz1oT4JlWasAFoPRjAmN7FmpKhsQKoU');
COMMIT;

-- ----------------------------
-- Table structure for user_job
-- ----------------------------
DROP TABLE IF EXISTS `user_job`;
CREATE TABLE `user_job` (
  `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `created_at` bigint DEFAULT NULL COMMENT '创建时间',
  `job_id` int unsigned NOT NULL COMMENT '角色id',
  `user_id` int unsigned NOT NULL COMMENT '菜单id',
  PRIMARY KEY (`id`),
  KEY `idx_user_job_created_at` (`created_at`),
  KEY `fk_user_job_job` (`job_id`),
  CONSTRAINT `fk_user_job_job` FOREIGN KEY (`job_id`) REFERENCES `job` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='用户职位信息';

-- ----------------------------
-- Records of user_job
-- ----------------------------
BEGIN;
INSERT INTO `user_job` VALUES (1, 1713290381, 1, 1);
COMMIT;

-- ----------------------------
-- Table structure for user_role
-- ----------------------------
DROP TABLE IF EXISTS `user_role`;
CREATE TABLE `user_role` (
  `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `created_at` bigint DEFAULT NULL COMMENT '创建时间',
  `role_id` int unsigned NOT NULL COMMENT '角色id',
  `user_id` int unsigned NOT NULL COMMENT '菜单id',
  PRIMARY KEY (`id`),
  KEY `idx_user_role_created_at` (`created_at`),
  KEY `fk_user_role_role` (`role_id`),
  CONSTRAINT `fk_user_role_role` FOREIGN KEY (`role_id`) REFERENCES `role` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='用户角色信息';

-- ----------------------------
-- Records of user_role
-- ----------------------------
BEGIN;
INSERT INTO `user_role` VALUES (1, 1713290381, 1, 1);
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
