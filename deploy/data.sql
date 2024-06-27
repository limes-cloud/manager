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

 Date: 27/06/2024 12:47:31
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
) ENGINE=InnoDB AUTO_INCREMENT=591 DEFAULT CHARSET=utf8mb4 ;

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
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4  COMMENT='部门信息';

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
) ENGINE=InnoDB AUTO_INCREMENT=10 DEFAULT CHARSET=utf8mb4  COMMENT='部门层级信息';

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
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4  COMMENT='字典信息';

-- ----------------------------
-- Records of dictionary
-- ----------------------------
BEGIN;
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
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4  COMMENT='字典值信息';

-- ----------------------------
-- Records of dictionary_value
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for gorm_init
-- ----------------------------
DROP TABLE IF EXISTS `gorm_init`;
CREATE TABLE `gorm_init` (
                             `id` int unsigned NOT NULL AUTO_INCREMENT,
                             `init` tinyint(1) DEFAULT NULL,
                             PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 ;

-- ----------------------------
-- Records of gorm_init
-- ----------------------------
BEGIN;
INSERT INTO `gorm_init` VALUES (1, 1);
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
) ENGINE=InnoDB AUTO_INCREMENT=13 DEFAULT CHARSET=utf8mb4  COMMENT='职位信息';

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
) ENGINE=InnoDB AUTO_INCREMENT=2427 DEFAULT CHARSET=utf8mb4  COMMENT='菜单信息';

-- ----------------------------
-- Records of menu
-- ----------------------------
BEGIN;
INSERT INTO `menu` VALUES (2273, 0, '管理平台', 'R', 'SystemPlatform', 'apps', NULL, NULL, '/', NULL, 'Layout', NULL, 2, 0, NULL, NULL, NULL, 1719290128, 1719290128);
INSERT INTO `menu` VALUES (2274, 2273, '首页面板', 'M', 'Dashboard', 'dashboard', NULL, NULL, '/dashboard', NULL, '/dashboard/index', NULL, 0, 0, 1, 1, 1, 1719290128, 1719290128);
INSERT INTO `menu` VALUES (2275, 2273, '管理中心', 'M', 'SystemPlatformManager', 'apps', NULL, NULL, '/manager', NULL, NULL, NULL, 0, 0, NULL, NULL, NULL, 1719290128, 1719290128);
INSERT INTO `menu` VALUES (2276, 2275, '基础接口', 'G', 'ManagerBaseApi', 'apps', NULL, NULL, NULL, NULL, NULL, NULL, 0, 1, NULL, NULL, NULL, 1719290128, 1719290128);
INSERT INTO `menu` VALUES (2277, 2276, '获取用户可见部门树', 'BA', NULL, NULL, '/manager/api/v1/departments', 'GET', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL, 1719290128, 1719290128);
INSERT INTO `menu` VALUES (2278, 2276, '获取用户可见角色树', 'BA', NULL, NULL, '/manager/api/v1/roles', 'GET', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL, 1719290128, 1719290128);
INSERT INTO `menu` VALUES (2279, 2276, '获取个人用户信息', 'BA', NULL, NULL, '/manager/api/v1/user/current', 'GET', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL, 1719290128, 1719290128);
INSERT INTO `menu` VALUES (2280, 2276, '更新个人用户信息', 'BA', NULL, NULL, '/manager/api/v1/user/current/info', 'PUT', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL, 1719290128, 1719290128);
INSERT INTO `menu` VALUES (2281, 2276, '更新个人用户密码', 'BA', NULL, NULL, '/manager/api/v1/user/current/password', 'PUT', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL, 1719290128, 1719290128);
INSERT INTO `menu` VALUES (2282, 2276, '保存个人用户设置', 'BA', NULL, NULL, '/manager/api/v1/user/current/setting', 'PUT', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL, 1719290128, 1719290128);
INSERT INTO `menu` VALUES (2283, 2276, '发送用户验证吗', 'BA', NULL, NULL, '/manager/api/v1/user/current/captcha', 'POST', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL, 1719290128, 1719290128);
INSERT INTO `menu` VALUES (2284, 2276, '获取用户当前角色菜单', 'BA', NULL, NULL, '/manager/api/v1/menus/by/cur_role', 'GET', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL, 1719290128, 1719290128);
INSERT INTO `menu` VALUES (2285, 2276, '退出登录', 'BA', NULL, NULL, '/manager/api/v1/user/logout', 'POST', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL, 1719290128, 1719290128);
INSERT INTO `menu` VALUES (2286, 2276, '刷新token', 'BA', NULL, NULL, '/manager/api/v1/user/token/refresh', 'POST', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL, 1719290128, 1719290128);
INSERT INTO `menu` VALUES (2287, 2276, '用户登录', 'BA', NULL, NULL, '/manager/api/v1/user/login', 'POST', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL, 1719290128, 1719290128);
INSERT INTO `menu` VALUES (2288, 2276, '获取登录验证码', 'BA', NULL, NULL, '/manager/api/v1/user/login/captcha', 'POST', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL, 1719290128, 1719290128);
INSERT INTO `menu` VALUES (2289, 2276, '获取系统基础设置', 'BA', NULL, NULL, '/manager/api/v1/system/setting', 'GET', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL, 1719290128, 1719290128);
INSERT INTO `menu` VALUES (2290, 2276, '接口鉴权', 'BA', NULL, NULL, '/manager/api/v1/auth', 'POST', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL, 1719290128, 1719290128);
INSERT INTO `menu` VALUES (2291, 2276, '切换用户角色', 'BA', NULL, NULL, '/manager/api/v1/user/current/role', 'POST', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL, 1719290128, 1719290128);
INSERT INTO `menu` VALUES (2292, 2276, '分块上传文件', 'BA', NULL, NULL, '/resource/api/v1/upload', 'POST', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL, 1719290128, 1719290128);
INSERT INTO `menu` VALUES (2293, 2276, '预上传文件', 'BA', NULL, NULL, '/resource/api/v1/prepare_upload', 'POST', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL, 1719290128, 1719290128);
INSERT INTO `menu` VALUES (2294, 2276, '获取字段类型', 'BA', NULL, NULL, '/usercenter/api/v1/field/types', 'GET', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL, 1719290128, 1719290128);
INSERT INTO `menu` VALUES (2295, 2275, '字典管理', 'M', 'ManagerDictionary', 'storage', NULL, NULL, '/manager/dictionary', NULL, '/manager/dictionary/index', NULL, 0, NULL, 1, NULL, 1, 1719290128, 1719290128);
INSERT INTO `menu` VALUES (2296, 2295, '查询字典', 'A', NULL, NULL, '/manager/api/v1/dictionaries', 'GET', NULL, 'manager:dictionary:query', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1719290128, 1719290128);
INSERT INTO `menu` VALUES (2297, 2295, '新增字典', 'A', NULL, NULL, '/manager/api/v1/dictionary', 'POST', NULL, 'manager:dictionary:add', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1719290128, 1719290128);
INSERT INTO `menu` VALUES (2298, 2295, '修改字典', 'A', NULL, NULL, '/manager/api/v1/dictionary', 'PUT', NULL, 'manager:dictionary:update', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1719290128, 1719290128);
INSERT INTO `menu` VALUES (2299, 2295, '删除字典', 'A', NULL, NULL, '/manager/api/v1/dictionary', 'DELETE', NULL, 'manager:dictionary:delete', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1719290128, 1719290128);
INSERT INTO `menu` VALUES (2300, 2295, '获取字典值', 'A', NULL, NULL, '/manager/api/v1/dictionary_values', 'GET', NULL, 'manager:dictionary:value:query', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1719290128, 1719290128);
INSERT INTO `menu` VALUES (2301, 2295, '新增字典值', 'A', NULL, NULL, '/manager/api/v1/dictionary_value', 'POST', NULL, 'manager:dictionary:value:add', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1719290128, 1719290128);
INSERT INTO `menu` VALUES (2302, 2295, '修改字典值', 'A', NULL, NULL, '/manager/api/v1/dictionary_value', 'PUT', NULL, 'manager:dictionary:value:update', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1719290128, 1719290128);
INSERT INTO `menu` VALUES (2303, 2295, '更新字典值目录状态', 'A', NULL, NULL, '/manager/api/v1/dictionary_value/status', 'PUT', NULL, 'manager:dictionary:value:status', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1719290128, 1719290128);
INSERT INTO `menu` VALUES (2304, 2295, '删除字典值', 'A', NULL, NULL, '/manager/api/v1/dictionary_value', 'DELETE', NULL, 'manager:dictionary:value:delete', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1719290128, 1719290128);
INSERT INTO `menu` VALUES (2305, 2275, '菜单管理', 'M', 'ManagerMenu', 'menu', NULL, NULL, '/manager/menu', NULL, '/manager/menu/index', NULL, 0, 0, 1, 1, 1, 1719290128, 1719290128);
INSERT INTO `menu` VALUES (2306, 2305, '查询菜单', 'A', NULL, NULL, '/manager/api/v1/menus', 'GET', NULL, 'manager:menu:query', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1719290128, 1719290128);
INSERT INTO `menu` VALUES (2307, 2305, '新增菜单', 'A', NULL, NULL, '/manager/api/v1/menu', 'POST', NULL, 'manager:menu:add', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1719290128, 1719290128);
INSERT INTO `menu` VALUES (2308, 2305, '修改菜单', 'A', NULL, NULL, '/manager/api/v1/menu', 'PUT', NULL, 'manager:menu:update', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1719290128, 1719290128);
INSERT INTO `menu` VALUES (2309, 2305, '删除菜单', 'A', NULL, NULL, '/manager/api/v1/menu', 'DELETE', NULL, 'manager:menu:delete', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1719290128, 1719290128);
INSERT INTO `menu` VALUES (2310, 2275, '职位管理', 'M', 'ManagerJob', 'tag', NULL, NULL, '/manager/job', NULL, '/manager/job/index', NULL, 0, NULL, 1, NULL, 1, 1719290128, 1719290128);
INSERT INTO `menu` VALUES (2311, 2310, '查询职位', 'A', NULL, NULL, '/manager/api/v1/jobs', 'GET', NULL, 'manager:job:query', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1719290128, 1719290128);
INSERT INTO `menu` VALUES (2312, 2310, '新增职位', 'A', NULL, NULL, '/manager/api/v1/job', 'POST', NULL, 'manager:job:add', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1719290128, 1719290128);
INSERT INTO `menu` VALUES (2313, 2310, '修改职位', 'A', NULL, NULL, '/manager/api/v1/job', 'PUT', NULL, 'manager:job:update', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1719290128, 1719290128);
INSERT INTO `menu` VALUES (2314, 2310, '删除职位', 'A', NULL, NULL, '/manager/api/v1/job', 'DELETE', NULL, 'manager:job:delete', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1719290128, 1719290128);
INSERT INTO `menu` VALUES (2315, 2275, '部门管理', 'M', 'ManagerDepartment', 'user-group', NULL, NULL, '/manager/department', NULL, '/manager/department/index', NULL, 0, 0, 1, NULL, 1, 1719290128, 1719290128);
INSERT INTO `menu` VALUES (2316, 2315, '新增部门', 'A', NULL, NULL, '/manager/api/v1/department', 'POST', NULL, 'manager:department:add', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1719290128, 1719290128);
INSERT INTO `menu` VALUES (2317, 2315, '修改部门', 'A', NULL, NULL, '/manager/api/v1/department', 'PUT', NULL, 'manager:department:update', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1719290128, 1719290128);
INSERT INTO `menu` VALUES (2318, 2315, '删除部门', 'A', NULL, NULL, '/manager/api/v1/department', 'DELETE', NULL, 'manager:department:delete', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1719290128, 1719290128);
INSERT INTO `menu` VALUES (2319, 2275, '角色管理', 'M', 'ManagerRole', 'safe', NULL, NULL, '/manager/role', NULL, '/manager/role/index', NULL, 0, NULL, 0, NULL, 1, 1719290128, 1719290128);
INSERT INTO `menu` VALUES (2320, 2319, '新增角色', 'A', NULL, NULL, '/manager/api/v1/role', 'POST', NULL, 'manager:role:add', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1719290128, 1719290128);
INSERT INTO `menu` VALUES (2321, 2319, '修改角色', 'A', NULL, NULL, '/manager/api/v1/role', 'PUT', NULL, 'manager:role:update', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1719290128, 1719290128);
INSERT INTO `menu` VALUES (2322, 2319, '修改角色状态', 'A', NULL, NULL, '/manager/api/v1/role/status', 'PUT', NULL, 'manager:role:update:status', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1719290128, 1719290128);
INSERT INTO `menu` VALUES (2323, 2319, '删除角色', 'A', NULL, NULL, '/manager/api/v1/role', 'DELETE', NULL, 'manager:role:delete', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1719290128, 1719290128);
INSERT INTO `menu` VALUES (2324, 2319, '角色菜单管理', 'G', NULL, NULL, NULL, NULL, NULL, 'manager:role:menu', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1719290128, 1719290128);
INSERT INTO `menu` VALUES (2325, 2324, '查询角色菜单', 'A', NULL, NULL, '/manager/api/v1/role/menu_ids', 'GET', NULL, 'manager:role:menu:query', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1719290128, 1719290128);
INSERT INTO `menu` VALUES (2326, 2324, '修改角色菜单', 'A', NULL, NULL, '/manager/api/v1/role/menu', 'POST', NULL, 'manager:role:menu:update', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1719290128, 1719290128);
INSERT INTO `menu` VALUES (2327, 2275, '用户管理', 'M', 'ManagerUser', 'user', NULL, NULL, '/manager/user', NULL, '/manager/user/index', NULL, 0, NULL, 1, NULL, 1, 1719290128, 1719290128);
INSERT INTO `menu` VALUES (2328, 2327, '查询用户列表', 'A', NULL, NULL, '/manager/api/v1/users', 'GET', NULL, 'manager:user:query', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1719290128, 1719290128);
INSERT INTO `menu` VALUES (2329, 2327, '新增用户', 'A', NULL, NULL, '/manager/api/v1/user', 'POST', NULL, 'manager:user:add', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1719290128, 1719290128);
INSERT INTO `menu` VALUES (2330, 2327, '修改用户', 'A', NULL, NULL, '/manager/api/v1/user', 'PUT', NULL, 'manager:user:update', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1719290128, 1719290128);
INSERT INTO `menu` VALUES (2331, 2327, '删除用户', 'A', NULL, NULL, '/manager/api/v1/user', 'DELETE', NULL, 'manager:user:delete', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1719290128, 1719290128);
INSERT INTO `menu` VALUES (2332, 2327, '修改用户状态', 'A', NULL, NULL, '/manager/api/v1/user/status', 'PUT', NULL, 'manager:user:status', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1719290128, 1719290128);
INSERT INTO `menu` VALUES (2333, 2327, '重置账号密码', 'A', NULL, NULL, '/manager/api/v1/user/password/reset', 'POST', NULL, 'manager:user:reset:password', '', NULL, 0, NULL, NULL, NULL, NULL, 1719290128, 1719290128);
INSERT INTO `menu` VALUES (2334, 2273, '资源中心', 'M', 'SystemPlatformResource', 'file', NULL, NULL, '/resource', NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL, 1719290128, 1719290128);
INSERT INTO `menu` VALUES (2335, 2334, '文件管理', 'M', 'ResourceFile', 'file', NULL, NULL, '/resource/file', NULL, '/resource/file/index', NULL, 0, 0, 1, NULL, 1, 1719290128, 1719290128);
INSERT INTO `menu` VALUES (2336, 2335, '目录管理', 'G', NULL, NULL, NULL, NULL, NULL, 'resource:directory:group', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1719290128, 1719290128);
INSERT INTO `menu` VALUES (2337, 2336, '查看目录', 'A', NULL, NULL, '/resource/api/v1/directories', 'GET', NULL, 'resource:directory:query', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1719290128, 1719290128);
INSERT INTO `menu` VALUES (2338, 2336, '新增目录', 'A', NULL, NULL, '/resource/api/v1/directory', 'POST', NULL, 'resource:directory:add', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1719290128, 1719290128);
INSERT INTO `menu` VALUES (2339, 2336, '修改目录', 'A', NULL, NULL, '/resource/api/v1/directory', 'PUT', NULL, 'resource:directory:update', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1719290128, 1719290128);
INSERT INTO `menu` VALUES (2340, 2336, '删除目录', 'A', NULL, NULL, '/resource/api/v1/directory', 'DELETE', NULL, 'resource:directory:delete', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1719290128, 1719290128);
INSERT INTO `menu` VALUES (2341, 2335, '文件管理', 'G', NULL, NULL, NULL, NULL, NULL, 'resource:file:group', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1719290128, 1719290128);
INSERT INTO `menu` VALUES (2342, 2341, '查看文件', 'A', NULL, NULL, '/resource/api/v1/files', 'GET', NULL, 'resource:file:query', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1719290128, 1719290128);
INSERT INTO `menu` VALUES (2343, 2341, '修改文件', 'A', NULL, NULL, '/resource/api/v1/file', 'PUT', NULL, 'resource:file:update', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1719290128, 1719290128);
INSERT INTO `menu` VALUES (2344, 2341, '删除文件', 'A', NULL, NULL, '/resource/api/v1/file', 'DELETE', NULL, 'resource:file:delete', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1719290128, 1719290128);
INSERT INTO `menu` VALUES (2345, 2334, '导出管理', 'M', 'ResourceExport', 'export', NULL, NULL, '/resource/export', NULL, '/resource/export/index', NULL, 0, 0, 1, NULL, 1, 1719290128, 1719290128);
INSERT INTO `menu` VALUES (2346, 2345, '查看导出', 'A', NULL, NULL, '/resource/api/v1/exports', 'GET', NULL, 'resource:export:query', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1719290128, 1719290128);
INSERT INTO `menu` VALUES (2347, 2345, '新增导出', 'A', NULL, NULL, '/resource/api/v1/export', 'POST', NULL, 'resource:export:file', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1719290128, 1719290128);
INSERT INTO `menu` VALUES (2348, 2345, '删除导出', 'A', NULL, NULL, '/resource/api/v1/export', 'DELETE', NULL, 'resource:export:delete', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1719290128, 1719290128);
INSERT INTO `menu` VALUES (2349, 2273, '用户中心', 'M', 'SystemPlatformUserCenter', 'user', NULL, NULL, '/usercenter', NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL, 1719290128, 1719290128);
INSERT INTO `menu` VALUES (2350, 2349, '授权渠道', 'M', 'UserCenterChannel', 'mind-mapping', NULL, NULL, '/usercenter/channel', NULL, '/usercenter/channel/index', NULL, 0, 0, 1, NULL, 1, 1719290128, 1719290128);
INSERT INTO `menu` VALUES (2351, 2350, '查看渠道', 'A', NULL, NULL, '/usercenter/api/v1/channels', 'GET', NULL, 'uc:channel:query', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1719290128, 1719290128);
INSERT INTO `menu` VALUES (2352, 2350, '新增渠道', 'A', NULL, NULL, '/usercenter/api/v1/channel', 'POST', NULL, 'uc:channel:add', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1719290128, 1719290128);
INSERT INTO `menu` VALUES (2353, 2350, '修改渠道', 'A', NULL, NULL, '/usercenter/api/v1/channel', 'PUT', NULL, 'uc:channel:update', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1719290128, 1719290128);
INSERT INTO `menu` VALUES (2354, 2350, '修改渠道状态', 'A', NULL, NULL, '/usercenter/api/v1/channel/status', 'PUT', NULL, 'uc:channel:update:status', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1719290128, 1719290128);
INSERT INTO `menu` VALUES (2355, 2350, '删除渠道', 'A', NULL, NULL, '/usercenter/api/v1/channel', 'DELETE', NULL, 'uc:channel:delete', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1719290128, 1719290128);
INSERT INTO `menu` VALUES (2356, 2349, '信息字段', 'M', 'UserCenterField', 'list', NULL, NULL, '/usercenter/field', NULL, '/usercenter/field/index', NULL, 0, 0, 1, 0, 1, 1719290128, 1719290128);
INSERT INTO `menu` VALUES (2357, 2356, '查看字段列表', 'A', NULL, NULL, '/usercenter/api/v1/fields', 'GET', NULL, 'uc:field:query', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1719290128, 1719290128);
INSERT INTO `menu` VALUES (2358, 2356, '新增字段', 'A', NULL, NULL, '/usercenter/api/v1/field', 'POST', NULL, 'uc:field:add', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1719290128, 1719290128);
INSERT INTO `menu` VALUES (2359, 2356, '修改字段', 'A', NULL, NULL, '/usercenter/api/v1/field', 'PUT', NULL, 'uc:field:update', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1719290128, 1719290128);
INSERT INTO `menu` VALUES (2360, 2356, '修改字段状态', 'A', NULL, NULL, '/usercenter/api/v1/field/status', 'PUT', NULL, 'uc:field:update:status', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1719290128, 1719290128);
INSERT INTO `menu` VALUES (2361, 2356, '删除字段', 'A', NULL, NULL, '/usercenter/api/v1/field', 'DELETE', NULL, 'uc:field:delete', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1719290128, 1719290128);
INSERT INTO `menu` VALUES (2362, 2349, '应用管理', 'M', 'UserCenterApp', 'apps', NULL, NULL, '/usercenter/app', NULL, '/usercenter/app/index', NULL, 0, 0, 1, NULL, 1, 1719290128, 1719290128);
INSERT INTO `menu` VALUES (2363, 2362, '查看应用', 'A', NULL, NULL, '/usercenter/api/v1/apps', 'GET', NULL, 'uc:app:query', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1719290128, 1719290128);
INSERT INTO `menu` VALUES (2364, 2362, '新增应用', 'A', NULL, NULL, '/usercenter/api/v1/app', 'POST', NULL, 'uc:app:add', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1719290128, 1719290128);
INSERT INTO `menu` VALUES (2365, 2362, '修改应用', 'A', NULL, NULL, '/usercenter/api/v1/app', 'PUT', NULL, 'uc:app:update', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1719290128, 1719290128);
INSERT INTO `menu` VALUES (2366, 2362, '修改应用状态', 'A', NULL, NULL, '/usercenter/api/v1/app/status', 'PUT', NULL, 'uc:app:update:status', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1719290128, 1719290128);
INSERT INTO `menu` VALUES (2367, 2362, '删除应用', 'A', NULL, NULL, '/usercenter/api/v1/app', 'DELETE', NULL, 'uc:app:delete', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1719290128, 1719290128);
INSERT INTO `menu` VALUES (2368, 2349, '用户管理', 'M', 'UserCenterUser', 'user', NULL, NULL, '/usercenter/user', NULL, '/usercenter/user/index', NULL, 0, 0, 1, NULL, 1, 1719290128, 1719290128);
INSERT INTO `menu` VALUES (2369, 2368, '查看用户', 'A', NULL, NULL, '/usercenter/api/v1/users', 'GET', NULL, 'uc:user:query', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1719290128, 1719290128);
INSERT INTO `menu` VALUES (2370, 2368, '新增用户', 'A', NULL, NULL, '/usercenter/api/v1/user', 'POST', NULL, 'uc:user:add', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1719290128, 1719290128);
INSERT INTO `menu` VALUES (2371, 2368, '导入用户', 'A', NULL, NULL, '/usercenter/api/v1/users', 'POST', NULL, 'uc:user:import', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1719290128, 1719290128);
INSERT INTO `menu` VALUES (2372, 2368, '修改用户', 'A', NULL, NULL, '/usercenter/api/v1/user', 'PUT', NULL, 'uc:user:update', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1719290128, 1719290128);
INSERT INTO `menu` VALUES (2373, 2368, '修改用户状态', 'A', NULL, NULL, '/usercenter/api/v1/user/status', 'PUT', NULL, 'uc:user:update:status', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1719290128, 1719290128);
INSERT INTO `menu` VALUES (2374, 2368, '删除用户', 'A', NULL, NULL, '/usercenter/api/v1/user', 'DELETE', NULL, 'uc:user:delete', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1719290128, 1719290128);
INSERT INTO `menu` VALUES (2375, 2368, '查看用户详细信息', 'G', NULL, NULL, NULL, NULL, NULL, 'uc:user:more', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1719290128, 1719290128);
INSERT INTO `menu` VALUES (2376, 2375, '获取用户应用信息', 'A', NULL, NULL, '/usercenter/api/v1/auths', 'GET', NULL, 'uc:auth:query', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1719290128, 1719290128);
INSERT INTO `menu` VALUES (2377, 2375, '创建用户应用信息', 'A', NULL, NULL, '/usercenter/api/v1/auth', 'POST', NULL, 'uc:auth:add', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1719290128, 1719290128);
INSERT INTO `menu` VALUES (2378, 2375, '修改用户应用状态信息', 'A', NULL, NULL, '/usercenter/api/v1/auth/status', 'PUT', NULL, 'uc:auth:update:status', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1719290128, 1719290128);
INSERT INTO `menu` VALUES (2379, 2375, '删除用户应用信息', 'A', NULL, NULL, '/usercenter/api/v1/auth', 'DELETE', NULL, 'uc:auth:delete', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1719290128, 1719290128);
INSERT INTO `menu` VALUES (2380, 2375, '获取用户渠道信息', 'A', NULL, NULL, '/usercenter/api/v1/oauths', 'GET', NULL, 'uc:oauth:query', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1719290128, 1719290128);
INSERT INTO `menu` VALUES (2381, 2375, '删除用户渠道信息', 'A', NULL, NULL, '/usercenter/api/v1/oauth', 'DELETE', NULL, 'uc:oauth:delete', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1719290128, 1719290128);
INSERT INTO `menu` VALUES (2382, 2375, '获取用户扩展信息', 'A', NULL, NULL, '/usercenter/api/v1/userinfos', 'GET', NULL, 'uc:userinfo:query', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1719290128, 1719290128);
INSERT INTO `menu` VALUES (2383, 2375, '创建用户扩展信息', 'A', NULL, NULL, '/usercenter/api/v1/userinfo', 'POST', NULL, 'uc:userinfo:add', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1719290128, 1719290128);
INSERT INTO `menu` VALUES (2384, 2375, '修改用户扩展信息', 'A', NULL, NULL, '/usercenter/api/v1/userinfo', 'PUT', NULL, 'uc:userinfo:update', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1719290128, 1719290128);
INSERT INTO `menu` VALUES (2385, 2375, '删除用户扩展信息', 'A', NULL, NULL, '/usercenter/api/v1/userinfo', 'DELETE', NULL, 'uc:userinfo:delete', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1719290128, 1719290128);
INSERT INTO `menu` VALUES (2386, 2273, '配置中心', 'M', 'SystemPlatformConfigure', 'code-block', NULL, NULL, '/configure', NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL, 1719290128, 1719290128);
INSERT INTO `menu` VALUES (2387, 2386, '环境管理', 'M', 'ConfigureEnv', 'common', NULL, NULL, '/configure/env', NULL, '/configure/env/index', NULL, 0, NULL, 1, NULL, 1, 1719290128, 1719290128);
INSERT INTO `menu` VALUES (2388, 2387, '查看环境', 'A', NULL, NULL, '/configure/api/v1/envs', 'GET', NULL, 'configure:env:query', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1719290128, 1719290128);
INSERT INTO `menu` VALUES (2389, 2387, '新增环境', 'A', NULL, NULL, '/configure/api/v1/env', 'POST', NULL, 'configure:env:add', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1719290128, 1719290128);
INSERT INTO `menu` VALUES (2390, 2387, '修改环境', 'A', NULL, NULL, '/configure/api/v1/env', 'PUT', NULL, 'configure:env:update', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1719290128, 1719290128);
INSERT INTO `menu` VALUES (2391, 2387, '修改环境状态', 'A', NULL, NULL, '/configure/api/v1/env/status', 'PUT', NULL, 'configure:env:update:status', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1719290128, 1719290128);
INSERT INTO `menu` VALUES (2392, 2387, '删除环境', 'A', NULL, NULL, '/configure/api/v1/env', 'DELETE', NULL, 'configure:env:delete', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1719290128, 1719290128);
INSERT INTO `menu` VALUES (2393, 2387, '查看环境Token', 'A', NULL, NULL, '/configure/api/v1/env/token', 'GET', NULL, 'configure:env:token:query', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1719290128, 1719290128);
INSERT INTO `menu` VALUES (2394, 2387, '重置环境token', 'A', NULL, NULL, '/configure/api/v1/env/token', 'PUT', NULL, 'configure:env:token:reset', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1719290128, 1719290128);
INSERT INTO `menu` VALUES (2395, 2386, '服务管理', 'M', 'ConfigureServer', 'apps', NULL, NULL, '/configure/server', NULL, '/configure/server/index', NULL, 0, NULL, 1, NULL, 1, 1719290128, 1719290128);
INSERT INTO `menu` VALUES (2396, 2395, '查询服务', 'A', NULL, NULL, '/configure/api/v1/servers', 'GET', NULL, 'configure:server:query', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1719290128, 1719290128);
INSERT INTO `menu` VALUES (2397, 2395, '新增服务', 'A', NULL, NULL, '/configure/api/v1/server', 'POST', NULL, 'configure:server:add', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1719290128, 1719290128);
INSERT INTO `menu` VALUES (2398, 2395, '修改服务', 'A', NULL, NULL, '/configure/api/v1/server', 'PUT', NULL, 'configure:server:update', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1719290128, 1719290128);
INSERT INTO `menu` VALUES (2399, 2395, '修改服务状态', 'A', NULL, NULL, '/configure/api/v1/server/status', 'PUT', NULL, 'configure:server:update:status', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1719290128, 1719290128);
INSERT INTO `menu` VALUES (2400, 2395, '删除服务', 'A', NULL, NULL, '/configure/api/v1/server', 'DELETE', NULL, 'configure:server:delete', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1719290128, 1719290128);
INSERT INTO `menu` VALUES (2401, 2386, '业务变量', 'M', 'ConfigureBusiness', 'code', NULL, NULL, '/configure/business', NULL, '/configure/business/index', NULL, 0, NULL, 1, NULL, 1, 1719290128, 1719290128);
INSERT INTO `menu` VALUES (2402, 2401, '查看业务变量', 'A', NULL, NULL, '/configure/api/v1/business', 'GET', NULL, 'configure:business:query', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1719290128, 1719290128);
INSERT INTO `menu` VALUES (2403, 2401, '新增业务变量', 'A', NULL, NULL, '/configure/api/v1/business', 'POST', NULL, 'configure:business:add', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1719290128, 1719290128);
INSERT INTO `menu` VALUES (2404, 2401, '修改业务变量', 'A', NULL, NULL, '/configure/api/v1/business', 'PUT', NULL, 'configure:business:update', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1719290128, 1719290128);
INSERT INTO `menu` VALUES (2405, 2401, '删除业务变量', 'A', NULL, NULL, '/configure/api/v1/business', 'DELETE', NULL, 'configure:business:delete', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1719290128, 1719290128);
INSERT INTO `menu` VALUES (2406, 2401, '查看业务变量值', 'A', NULL, NULL, '/configure/business/values', 'get', NULL, 'configure:business:value:query', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1719290128, 1719290128);
INSERT INTO `menu` VALUES (2407, 2401, '设置业务变量值', 'A', NULL, NULL, '/configure/business/values', 'PUT', NULL, 'configure:business:value:update', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1719290128, 1719290128);
INSERT INTO `menu` VALUES (2408, 2386, '资源变量', 'M', 'ConfigureResource', 'unordered-list', NULL, NULL, '/configure/resource', NULL, '/configure/resource/index', NULL, 0, NULL, 1, NULL, 1, 1719290128, 1719290128);
INSERT INTO `menu` VALUES (2409, 2408, '查看资源', 'A', NULL, NULL, '/configure/api/v1/resources', 'GET', NULL, 'configure:resource:query', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1719290128, 1719290128);
INSERT INTO `menu` VALUES (2410, 2408, '新增资源', 'A', NULL, NULL, '/configure/api/v1/resource', 'POST', NULL, 'configure:resource:add', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1719290128, 1719290128);
INSERT INTO `menu` VALUES (2411, 2408, '修改资源', 'A', NULL, NULL, '/configure/api/v1/resource', 'PUT', NULL, 'configure:resource:update', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1719290128, 1719290128);
INSERT INTO `menu` VALUES (2412, 2408, '删除资源', 'A', NULL, NULL, '/configure/api/v1/resource', 'DELETE', NULL, 'configure:resource:delete', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1719290128, 1719290128);
INSERT INTO `menu` VALUES (2413, 2408, '查看业务变量值', 'A', NULL, NULL, '/configure/resource/values', 'get', NULL, 'configure:resource:value:query', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1719290128, 1719290128);
INSERT INTO `menu` VALUES (2414, 2408, '设置业务变量值', 'A', NULL, NULL, '/configure/resource/values', 'PUT', NULL, 'configure:resource:value:update', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1719290128, 1719290128);
INSERT INTO `menu` VALUES (2415, 2386, '配置模板', 'M', 'ConfgureTemplate', 'file', NULL, NULL, '/configure/template', NULL, '/configure/template/index', NULL, 0, NULL, 1, NULL, 1, 1719290128, 1719290128);
INSERT INTO `menu` VALUES (2416, 2415, '模板管理', 'G', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL, 1719290128, 1719290128);
INSERT INTO `menu` VALUES (2417, 2416, '查看模板历史版本', 'A', NULL, NULL, '/configure/api/v1/templates', 'GET', NULL, 'configure:template:history', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1719290128, 1719290128);
INSERT INTO `menu` VALUES (2418, 2416, '查看指定模板详细数据', 'A', NULL, NULL, '/configure/api/v1/template', 'GET', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL, 1719290128, 1719290128);
INSERT INTO `menu` VALUES (2419, 2416, '查看当前正在使用的模板', 'A', NULL, NULL, '/configure/api/v1/template/current', 'GET', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL, 1719290128, 1719290128);
INSERT INTO `menu` VALUES (2420, 2416, '提交模板', 'A', NULL, NULL, '/configure/api/v1/template', 'POST', NULL, 'configure:template:add', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1719290128, 1719290128);
INSERT INTO `menu` VALUES (2421, 2416, '模板对比', 'A', NULL, NULL, '/configure/api/v1/template/compare', 'POST', NULL, 'configure:template:compare', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1719290128, 1719290128);
INSERT INTO `menu` VALUES (2422, 2416, '切换模板', 'A', NULL, NULL, '/configure/api/v1/template/switch', 'POST', NULL, 'configure:template:switch', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1719290128, 1719290128);
INSERT INTO `menu` VALUES (2423, 2416, '模板预览', 'A', NULL, NULL, '/configure/api/v1/template/preview', 'POST', NULL, 'configure:template:preview', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1719290128, 1719290128);
INSERT INTO `menu` VALUES (2424, 2415, '配置管理', 'G', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL, 1719290128, 1719290128);
INSERT INTO `menu` VALUES (2425, 2424, '配置对比', 'A', NULL, NULL, '/configure/api/v1/configure/compare', 'POST', NULL, 'configure:configure:compare', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1719290128, 1719290128);
INSERT INTO `menu` VALUES (2426, 2424, '同步配置', 'A', NULL, NULL, '/configure/api/v1/configure', 'PUT', NULL, 'configure:configure:sync', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1719290128, 1719290128);
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
) ENGINE=InnoDB AUTO_INCREMENT=5483 DEFAULT CHARSET=utf8mb4  COMMENT='菜单层级信息';

-- ----------------------------
-- Records of menu_closure
-- ----------------------------
BEGIN;
INSERT INTO `menu_closure` VALUES (5024, 2273, 2277);
INSERT INTO `menu_closure` VALUES (5025, 2275, 2277);
INSERT INTO `menu_closure` VALUES (5026, 2276, 2277);
INSERT INTO `menu_closure` VALUES (5027, 2273, 2278);
INSERT INTO `menu_closure` VALUES (5028, 2275, 2278);
INSERT INTO `menu_closure` VALUES (5029, 2276, 2278);
INSERT INTO `menu_closure` VALUES (5030, 2273, 2279);
INSERT INTO `menu_closure` VALUES (5031, 2275, 2279);
INSERT INTO `menu_closure` VALUES (5032, 2276, 2279);
INSERT INTO `menu_closure` VALUES (5033, 2273, 2280);
INSERT INTO `menu_closure` VALUES (5034, 2275, 2280);
INSERT INTO `menu_closure` VALUES (5035, 2276, 2280);
INSERT INTO `menu_closure` VALUES (5036, 2273, 2281);
INSERT INTO `menu_closure` VALUES (5037, 2275, 2281);
INSERT INTO `menu_closure` VALUES (5038, 2276, 2281);
INSERT INTO `menu_closure` VALUES (5039, 2273, 2282);
INSERT INTO `menu_closure` VALUES (5040, 2275, 2282);
INSERT INTO `menu_closure` VALUES (5041, 2276, 2282);
INSERT INTO `menu_closure` VALUES (5042, 2273, 2283);
INSERT INTO `menu_closure` VALUES (5043, 2275, 2283);
INSERT INTO `menu_closure` VALUES (5044, 2276, 2283);
INSERT INTO `menu_closure` VALUES (5045, 2273, 2284);
INSERT INTO `menu_closure` VALUES (5046, 2275, 2284);
INSERT INTO `menu_closure` VALUES (5047, 2276, 2284);
INSERT INTO `menu_closure` VALUES (5048, 2273, 2285);
INSERT INTO `menu_closure` VALUES (5049, 2275, 2285);
INSERT INTO `menu_closure` VALUES (5050, 2276, 2285);
INSERT INTO `menu_closure` VALUES (5051, 2273, 2286);
INSERT INTO `menu_closure` VALUES (5052, 2275, 2286);
INSERT INTO `menu_closure` VALUES (5053, 2276, 2286);
INSERT INTO `menu_closure` VALUES (5054, 2273, 2287);
INSERT INTO `menu_closure` VALUES (5055, 2275, 2287);
INSERT INTO `menu_closure` VALUES (5056, 2276, 2287);
INSERT INTO `menu_closure` VALUES (5057, 2273, 2288);
INSERT INTO `menu_closure` VALUES (5058, 2275, 2288);
INSERT INTO `menu_closure` VALUES (5059, 2276, 2288);
INSERT INTO `menu_closure` VALUES (5060, 2273, 2289);
INSERT INTO `menu_closure` VALUES (5061, 2275, 2289);
INSERT INTO `menu_closure` VALUES (5062, 2276, 2289);
INSERT INTO `menu_closure` VALUES (5063, 2273, 2290);
INSERT INTO `menu_closure` VALUES (5064, 2275, 2290);
INSERT INTO `menu_closure` VALUES (5065, 2276, 2290);
INSERT INTO `menu_closure` VALUES (5066, 2273, 2291);
INSERT INTO `menu_closure` VALUES (5067, 2275, 2291);
INSERT INTO `menu_closure` VALUES (5068, 2276, 2291);
INSERT INTO `menu_closure` VALUES (5069, 2273, 2292);
INSERT INTO `menu_closure` VALUES (5070, 2275, 2292);
INSERT INTO `menu_closure` VALUES (5071, 2276, 2292);
INSERT INTO `menu_closure` VALUES (5072, 2273, 2293);
INSERT INTO `menu_closure` VALUES (5073, 2275, 2293);
INSERT INTO `menu_closure` VALUES (5074, 2276, 2293);
INSERT INTO `menu_closure` VALUES (5075, 2273, 2294);
INSERT INTO `menu_closure` VALUES (5076, 2275, 2294);
INSERT INTO `menu_closure` VALUES (5077, 2276, 2294);
INSERT INTO `menu_closure` VALUES (5078, 2273, 2296);
INSERT INTO `menu_closure` VALUES (5079, 2275, 2296);
INSERT INTO `menu_closure` VALUES (5080, 2295, 2296);
INSERT INTO `menu_closure` VALUES (5081, 2273, 2297);
INSERT INTO `menu_closure` VALUES (5082, 2275, 2297);
INSERT INTO `menu_closure` VALUES (5083, 2295, 2297);
INSERT INTO `menu_closure` VALUES (5084, 2273, 2298);
INSERT INTO `menu_closure` VALUES (5085, 2275, 2298);
INSERT INTO `menu_closure` VALUES (5086, 2295, 2298);
INSERT INTO `menu_closure` VALUES (5087, 2273, 2299);
INSERT INTO `menu_closure` VALUES (5088, 2275, 2299);
INSERT INTO `menu_closure` VALUES (5089, 2295, 2299);
INSERT INTO `menu_closure` VALUES (5090, 2273, 2300);
INSERT INTO `menu_closure` VALUES (5091, 2275, 2300);
INSERT INTO `menu_closure` VALUES (5092, 2295, 2300);
INSERT INTO `menu_closure` VALUES (5093, 2273, 2301);
INSERT INTO `menu_closure` VALUES (5094, 2275, 2301);
INSERT INTO `menu_closure` VALUES (5095, 2295, 2301);
INSERT INTO `menu_closure` VALUES (5096, 2273, 2302);
INSERT INTO `menu_closure` VALUES (5097, 2275, 2302);
INSERT INTO `menu_closure` VALUES (5098, 2295, 2302);
INSERT INTO `menu_closure` VALUES (5099, 2273, 2303);
INSERT INTO `menu_closure` VALUES (5100, 2275, 2303);
INSERT INTO `menu_closure` VALUES (5101, 2295, 2303);
INSERT INTO `menu_closure` VALUES (5102, 2273, 2304);
INSERT INTO `menu_closure` VALUES (5103, 2275, 2304);
INSERT INTO `menu_closure` VALUES (5104, 2295, 2304);
INSERT INTO `menu_closure` VALUES (5105, 2273, 2306);
INSERT INTO `menu_closure` VALUES (5106, 2275, 2306);
INSERT INTO `menu_closure` VALUES (5107, 2305, 2306);
INSERT INTO `menu_closure` VALUES (5108, 2273, 2307);
INSERT INTO `menu_closure` VALUES (5109, 2275, 2307);
INSERT INTO `menu_closure` VALUES (5110, 2305, 2307);
INSERT INTO `menu_closure` VALUES (5111, 2273, 2308);
INSERT INTO `menu_closure` VALUES (5112, 2275, 2308);
INSERT INTO `menu_closure` VALUES (5113, 2305, 2308);
INSERT INTO `menu_closure` VALUES (5114, 2273, 2309);
INSERT INTO `menu_closure` VALUES (5115, 2275, 2309);
INSERT INTO `menu_closure` VALUES (5116, 2305, 2309);
INSERT INTO `menu_closure` VALUES (5117, 2273, 2311);
INSERT INTO `menu_closure` VALUES (5118, 2275, 2311);
INSERT INTO `menu_closure` VALUES (5119, 2310, 2311);
INSERT INTO `menu_closure` VALUES (5120, 2273, 2312);
INSERT INTO `menu_closure` VALUES (5121, 2275, 2312);
INSERT INTO `menu_closure` VALUES (5122, 2310, 2312);
INSERT INTO `menu_closure` VALUES (5123, 2273, 2313);
INSERT INTO `menu_closure` VALUES (5124, 2275, 2313);
INSERT INTO `menu_closure` VALUES (5125, 2310, 2313);
INSERT INTO `menu_closure` VALUES (5126, 2273, 2314);
INSERT INTO `menu_closure` VALUES (5127, 2275, 2314);
INSERT INTO `menu_closure` VALUES (5128, 2310, 2314);
INSERT INTO `menu_closure` VALUES (5129, 2273, 2316);
INSERT INTO `menu_closure` VALUES (5130, 2275, 2316);
INSERT INTO `menu_closure` VALUES (5131, 2315, 2316);
INSERT INTO `menu_closure` VALUES (5132, 2273, 2317);
INSERT INTO `menu_closure` VALUES (5133, 2275, 2317);
INSERT INTO `menu_closure` VALUES (5134, 2315, 2317);
INSERT INTO `menu_closure` VALUES (5135, 2273, 2318);
INSERT INTO `menu_closure` VALUES (5136, 2275, 2318);
INSERT INTO `menu_closure` VALUES (5137, 2315, 2318);
INSERT INTO `menu_closure` VALUES (5138, 2273, 2325);
INSERT INTO `menu_closure` VALUES (5139, 2275, 2325);
INSERT INTO `menu_closure` VALUES (5140, 2319, 2325);
INSERT INTO `menu_closure` VALUES (5141, 2324, 2325);
INSERT INTO `menu_closure` VALUES (5142, 2273, 2326);
INSERT INTO `menu_closure` VALUES (5143, 2275, 2326);
INSERT INTO `menu_closure` VALUES (5144, 2319, 2326);
INSERT INTO `menu_closure` VALUES (5145, 2324, 2326);
INSERT INTO `menu_closure` VALUES (5146, 2273, 2320);
INSERT INTO `menu_closure` VALUES (5147, 2275, 2320);
INSERT INTO `menu_closure` VALUES (5148, 2319, 2320);
INSERT INTO `menu_closure` VALUES (5149, 2273, 2321);
INSERT INTO `menu_closure` VALUES (5150, 2275, 2321);
INSERT INTO `menu_closure` VALUES (5151, 2319, 2321);
INSERT INTO `menu_closure` VALUES (5152, 2273, 2322);
INSERT INTO `menu_closure` VALUES (5153, 2275, 2322);
INSERT INTO `menu_closure` VALUES (5154, 2319, 2322);
INSERT INTO `menu_closure` VALUES (5155, 2273, 2323);
INSERT INTO `menu_closure` VALUES (5156, 2275, 2323);
INSERT INTO `menu_closure` VALUES (5157, 2319, 2323);
INSERT INTO `menu_closure` VALUES (5158, 2273, 2324);
INSERT INTO `menu_closure` VALUES (5159, 2275, 2324);
INSERT INTO `menu_closure` VALUES (5160, 2319, 2324);
INSERT INTO `menu_closure` VALUES (5161, 2273, 2328);
INSERT INTO `menu_closure` VALUES (5162, 2275, 2328);
INSERT INTO `menu_closure` VALUES (5163, 2327, 2328);
INSERT INTO `menu_closure` VALUES (5164, 2273, 2329);
INSERT INTO `menu_closure` VALUES (5165, 2275, 2329);
INSERT INTO `menu_closure` VALUES (5166, 2327, 2329);
INSERT INTO `menu_closure` VALUES (5167, 2273, 2330);
INSERT INTO `menu_closure` VALUES (5168, 2275, 2330);
INSERT INTO `menu_closure` VALUES (5169, 2327, 2330);
INSERT INTO `menu_closure` VALUES (5170, 2273, 2331);
INSERT INTO `menu_closure` VALUES (5171, 2275, 2331);
INSERT INTO `menu_closure` VALUES (5172, 2327, 2331);
INSERT INTO `menu_closure` VALUES (5173, 2273, 2332);
INSERT INTO `menu_closure` VALUES (5174, 2275, 2332);
INSERT INTO `menu_closure` VALUES (5175, 2327, 2332);
INSERT INTO `menu_closure` VALUES (5176, 2273, 2333);
INSERT INTO `menu_closure` VALUES (5177, 2275, 2333);
INSERT INTO `menu_closure` VALUES (5178, 2327, 2333);
INSERT INTO `menu_closure` VALUES (5179, 2273, 2276);
INSERT INTO `menu_closure` VALUES (5180, 2275, 2276);
INSERT INTO `menu_closure` VALUES (5181, 2273, 2295);
INSERT INTO `menu_closure` VALUES (5182, 2275, 2295);
INSERT INTO `menu_closure` VALUES (5183, 2273, 2305);
INSERT INTO `menu_closure` VALUES (5184, 2275, 2305);
INSERT INTO `menu_closure` VALUES (5185, 2273, 2310);
INSERT INTO `menu_closure` VALUES (5186, 2275, 2310);
INSERT INTO `menu_closure` VALUES (5187, 2273, 2315);
INSERT INTO `menu_closure` VALUES (5188, 2275, 2315);
INSERT INTO `menu_closure` VALUES (5189, 2273, 2319);
INSERT INTO `menu_closure` VALUES (5190, 2275, 2319);
INSERT INTO `menu_closure` VALUES (5191, 2273, 2327);
INSERT INTO `menu_closure` VALUES (5192, 2275, 2327);
INSERT INTO `menu_closure` VALUES (5193, 2273, 2337);
INSERT INTO `menu_closure` VALUES (5194, 2334, 2337);
INSERT INTO `menu_closure` VALUES (5195, 2335, 2337);
INSERT INTO `menu_closure` VALUES (5196, 2336, 2337);
INSERT INTO `menu_closure` VALUES (5197, 2273, 2338);
INSERT INTO `menu_closure` VALUES (5198, 2334, 2338);
INSERT INTO `menu_closure` VALUES (5199, 2335, 2338);
INSERT INTO `menu_closure` VALUES (5200, 2336, 2338);
INSERT INTO `menu_closure` VALUES (5201, 2273, 2339);
INSERT INTO `menu_closure` VALUES (5202, 2334, 2339);
INSERT INTO `menu_closure` VALUES (5203, 2335, 2339);
INSERT INTO `menu_closure` VALUES (5204, 2336, 2339);
INSERT INTO `menu_closure` VALUES (5205, 2273, 2340);
INSERT INTO `menu_closure` VALUES (5206, 2334, 2340);
INSERT INTO `menu_closure` VALUES (5207, 2335, 2340);
INSERT INTO `menu_closure` VALUES (5208, 2336, 2340);
INSERT INTO `menu_closure` VALUES (5209, 2273, 2342);
INSERT INTO `menu_closure` VALUES (5210, 2334, 2342);
INSERT INTO `menu_closure` VALUES (5211, 2335, 2342);
INSERT INTO `menu_closure` VALUES (5212, 2341, 2342);
INSERT INTO `menu_closure` VALUES (5213, 2273, 2343);
INSERT INTO `menu_closure` VALUES (5214, 2334, 2343);
INSERT INTO `menu_closure` VALUES (5215, 2335, 2343);
INSERT INTO `menu_closure` VALUES (5216, 2341, 2343);
INSERT INTO `menu_closure` VALUES (5217, 2273, 2344);
INSERT INTO `menu_closure` VALUES (5218, 2334, 2344);
INSERT INTO `menu_closure` VALUES (5219, 2335, 2344);
INSERT INTO `menu_closure` VALUES (5220, 2341, 2344);
INSERT INTO `menu_closure` VALUES (5221, 2273, 2336);
INSERT INTO `menu_closure` VALUES (5222, 2334, 2336);
INSERT INTO `menu_closure` VALUES (5223, 2335, 2336);
INSERT INTO `menu_closure` VALUES (5224, 2273, 2341);
INSERT INTO `menu_closure` VALUES (5225, 2334, 2341);
INSERT INTO `menu_closure` VALUES (5226, 2335, 2341);
INSERT INTO `menu_closure` VALUES (5227, 2273, 2346);
INSERT INTO `menu_closure` VALUES (5228, 2334, 2346);
INSERT INTO `menu_closure` VALUES (5229, 2345, 2346);
INSERT INTO `menu_closure` VALUES (5230, 2273, 2347);
INSERT INTO `menu_closure` VALUES (5231, 2334, 2347);
INSERT INTO `menu_closure` VALUES (5232, 2345, 2347);
INSERT INTO `menu_closure` VALUES (5233, 2273, 2348);
INSERT INTO `menu_closure` VALUES (5234, 2334, 2348);
INSERT INTO `menu_closure` VALUES (5235, 2345, 2348);
INSERT INTO `menu_closure` VALUES (5236, 2273, 2335);
INSERT INTO `menu_closure` VALUES (5237, 2334, 2335);
INSERT INTO `menu_closure` VALUES (5238, 2273, 2345);
INSERT INTO `menu_closure` VALUES (5239, 2334, 2345);
INSERT INTO `menu_closure` VALUES (5240, 2273, 2351);
INSERT INTO `menu_closure` VALUES (5241, 2349, 2351);
INSERT INTO `menu_closure` VALUES (5242, 2350, 2351);
INSERT INTO `menu_closure` VALUES (5243, 2273, 2352);
INSERT INTO `menu_closure` VALUES (5244, 2349, 2352);
INSERT INTO `menu_closure` VALUES (5245, 2350, 2352);
INSERT INTO `menu_closure` VALUES (5246, 2273, 2353);
INSERT INTO `menu_closure` VALUES (5247, 2349, 2353);
INSERT INTO `menu_closure` VALUES (5248, 2350, 2353);
INSERT INTO `menu_closure` VALUES (5249, 2273, 2354);
INSERT INTO `menu_closure` VALUES (5250, 2349, 2354);
INSERT INTO `menu_closure` VALUES (5251, 2350, 2354);
INSERT INTO `menu_closure` VALUES (5252, 2273, 2355);
INSERT INTO `menu_closure` VALUES (5253, 2349, 2355);
INSERT INTO `menu_closure` VALUES (5254, 2350, 2355);
INSERT INTO `menu_closure` VALUES (5255, 2273, 2357);
INSERT INTO `menu_closure` VALUES (5256, 2349, 2357);
INSERT INTO `menu_closure` VALUES (5257, 2356, 2357);
INSERT INTO `menu_closure` VALUES (5258, 2273, 2358);
INSERT INTO `menu_closure` VALUES (5259, 2349, 2358);
INSERT INTO `menu_closure` VALUES (5260, 2356, 2358);
INSERT INTO `menu_closure` VALUES (5261, 2273, 2359);
INSERT INTO `menu_closure` VALUES (5262, 2349, 2359);
INSERT INTO `menu_closure` VALUES (5263, 2356, 2359);
INSERT INTO `menu_closure` VALUES (5264, 2273, 2360);
INSERT INTO `menu_closure` VALUES (5265, 2349, 2360);
INSERT INTO `menu_closure` VALUES (5266, 2356, 2360);
INSERT INTO `menu_closure` VALUES (5267, 2273, 2361);
INSERT INTO `menu_closure` VALUES (5268, 2349, 2361);
INSERT INTO `menu_closure` VALUES (5269, 2356, 2361);
INSERT INTO `menu_closure` VALUES (5270, 2273, 2363);
INSERT INTO `menu_closure` VALUES (5271, 2349, 2363);
INSERT INTO `menu_closure` VALUES (5272, 2362, 2363);
INSERT INTO `menu_closure` VALUES (5273, 2273, 2364);
INSERT INTO `menu_closure` VALUES (5274, 2349, 2364);
INSERT INTO `menu_closure` VALUES (5275, 2362, 2364);
INSERT INTO `menu_closure` VALUES (5276, 2273, 2365);
INSERT INTO `menu_closure` VALUES (5277, 2349, 2365);
INSERT INTO `menu_closure` VALUES (5278, 2362, 2365);
INSERT INTO `menu_closure` VALUES (5279, 2273, 2366);
INSERT INTO `menu_closure` VALUES (5280, 2349, 2366);
INSERT INTO `menu_closure` VALUES (5281, 2362, 2366);
INSERT INTO `menu_closure` VALUES (5282, 2273, 2367);
INSERT INTO `menu_closure` VALUES (5283, 2349, 2367);
INSERT INTO `menu_closure` VALUES (5284, 2362, 2367);
INSERT INTO `menu_closure` VALUES (5285, 2273, 2376);
INSERT INTO `menu_closure` VALUES (5286, 2349, 2376);
INSERT INTO `menu_closure` VALUES (5287, 2368, 2376);
INSERT INTO `menu_closure` VALUES (5288, 2375, 2376);
INSERT INTO `menu_closure` VALUES (5289, 2273, 2377);
INSERT INTO `menu_closure` VALUES (5290, 2349, 2377);
INSERT INTO `menu_closure` VALUES (5291, 2368, 2377);
INSERT INTO `menu_closure` VALUES (5292, 2375, 2377);
INSERT INTO `menu_closure` VALUES (5293, 2273, 2378);
INSERT INTO `menu_closure` VALUES (5294, 2349, 2378);
INSERT INTO `menu_closure` VALUES (5295, 2368, 2378);
INSERT INTO `menu_closure` VALUES (5296, 2375, 2378);
INSERT INTO `menu_closure` VALUES (5297, 2273, 2379);
INSERT INTO `menu_closure` VALUES (5298, 2349, 2379);
INSERT INTO `menu_closure` VALUES (5299, 2368, 2379);
INSERT INTO `menu_closure` VALUES (5300, 2375, 2379);
INSERT INTO `menu_closure` VALUES (5301, 2273, 2380);
INSERT INTO `menu_closure` VALUES (5302, 2349, 2380);
INSERT INTO `menu_closure` VALUES (5303, 2368, 2380);
INSERT INTO `menu_closure` VALUES (5304, 2375, 2380);
INSERT INTO `menu_closure` VALUES (5305, 2273, 2381);
INSERT INTO `menu_closure` VALUES (5306, 2349, 2381);
INSERT INTO `menu_closure` VALUES (5307, 2368, 2381);
INSERT INTO `menu_closure` VALUES (5308, 2375, 2381);
INSERT INTO `menu_closure` VALUES (5309, 2273, 2382);
INSERT INTO `menu_closure` VALUES (5310, 2349, 2382);
INSERT INTO `menu_closure` VALUES (5311, 2368, 2382);
INSERT INTO `menu_closure` VALUES (5312, 2375, 2382);
INSERT INTO `menu_closure` VALUES (5313, 2273, 2383);
INSERT INTO `menu_closure` VALUES (5314, 2349, 2383);
INSERT INTO `menu_closure` VALUES (5315, 2368, 2383);
INSERT INTO `menu_closure` VALUES (5316, 2375, 2383);
INSERT INTO `menu_closure` VALUES (5317, 2273, 2384);
INSERT INTO `menu_closure` VALUES (5318, 2349, 2384);
INSERT INTO `menu_closure` VALUES (5319, 2368, 2384);
INSERT INTO `menu_closure` VALUES (5320, 2375, 2384);
INSERT INTO `menu_closure` VALUES (5321, 2273, 2385);
INSERT INTO `menu_closure` VALUES (5322, 2349, 2385);
INSERT INTO `menu_closure` VALUES (5323, 2368, 2385);
INSERT INTO `menu_closure` VALUES (5324, 2375, 2385);
INSERT INTO `menu_closure` VALUES (5325, 2273, 2369);
INSERT INTO `menu_closure` VALUES (5326, 2349, 2369);
INSERT INTO `menu_closure` VALUES (5327, 2368, 2369);
INSERT INTO `menu_closure` VALUES (5328, 2273, 2370);
INSERT INTO `menu_closure` VALUES (5329, 2349, 2370);
INSERT INTO `menu_closure` VALUES (5330, 2368, 2370);
INSERT INTO `menu_closure` VALUES (5331, 2273, 2371);
INSERT INTO `menu_closure` VALUES (5332, 2349, 2371);
INSERT INTO `menu_closure` VALUES (5333, 2368, 2371);
INSERT INTO `menu_closure` VALUES (5334, 2273, 2372);
INSERT INTO `menu_closure` VALUES (5335, 2349, 2372);
INSERT INTO `menu_closure` VALUES (5336, 2368, 2372);
INSERT INTO `menu_closure` VALUES (5337, 2273, 2373);
INSERT INTO `menu_closure` VALUES (5338, 2349, 2373);
INSERT INTO `menu_closure` VALUES (5339, 2368, 2373);
INSERT INTO `menu_closure` VALUES (5340, 2273, 2374);
INSERT INTO `menu_closure` VALUES (5341, 2349, 2374);
INSERT INTO `menu_closure` VALUES (5342, 2368, 2374);
INSERT INTO `menu_closure` VALUES (5343, 2273, 2375);
INSERT INTO `menu_closure` VALUES (5344, 2349, 2375);
INSERT INTO `menu_closure` VALUES (5345, 2368, 2375);
INSERT INTO `menu_closure` VALUES (5346, 2273, 2350);
INSERT INTO `menu_closure` VALUES (5347, 2349, 2350);
INSERT INTO `menu_closure` VALUES (5348, 2273, 2356);
INSERT INTO `menu_closure` VALUES (5349, 2349, 2356);
INSERT INTO `menu_closure` VALUES (5350, 2273, 2362);
INSERT INTO `menu_closure` VALUES (5351, 2349, 2362);
INSERT INTO `menu_closure` VALUES (5352, 2273, 2368);
INSERT INTO `menu_closure` VALUES (5353, 2349, 2368);
INSERT INTO `menu_closure` VALUES (5354, 2273, 2388);
INSERT INTO `menu_closure` VALUES (5355, 2386, 2388);
INSERT INTO `menu_closure` VALUES (5356, 2387, 2388);
INSERT INTO `menu_closure` VALUES (5357, 2273, 2389);
INSERT INTO `menu_closure` VALUES (5358, 2386, 2389);
INSERT INTO `menu_closure` VALUES (5359, 2387, 2389);
INSERT INTO `menu_closure` VALUES (5360, 2273, 2390);
INSERT INTO `menu_closure` VALUES (5361, 2386, 2390);
INSERT INTO `menu_closure` VALUES (5362, 2387, 2390);
INSERT INTO `menu_closure` VALUES (5363, 2273, 2391);
INSERT INTO `menu_closure` VALUES (5364, 2386, 2391);
INSERT INTO `menu_closure` VALUES (5365, 2387, 2391);
INSERT INTO `menu_closure` VALUES (5366, 2273, 2392);
INSERT INTO `menu_closure` VALUES (5367, 2386, 2392);
INSERT INTO `menu_closure` VALUES (5368, 2387, 2392);
INSERT INTO `menu_closure` VALUES (5369, 2273, 2393);
INSERT INTO `menu_closure` VALUES (5370, 2386, 2393);
INSERT INTO `menu_closure` VALUES (5371, 2387, 2393);
INSERT INTO `menu_closure` VALUES (5372, 2273, 2394);
INSERT INTO `menu_closure` VALUES (5373, 2386, 2394);
INSERT INTO `menu_closure` VALUES (5374, 2387, 2394);
INSERT INTO `menu_closure` VALUES (5375, 2273, 2396);
INSERT INTO `menu_closure` VALUES (5376, 2386, 2396);
INSERT INTO `menu_closure` VALUES (5377, 2395, 2396);
INSERT INTO `menu_closure` VALUES (5378, 2273, 2397);
INSERT INTO `menu_closure` VALUES (5379, 2386, 2397);
INSERT INTO `menu_closure` VALUES (5380, 2395, 2397);
INSERT INTO `menu_closure` VALUES (5381, 2273, 2398);
INSERT INTO `menu_closure` VALUES (5382, 2386, 2398);
INSERT INTO `menu_closure` VALUES (5383, 2395, 2398);
INSERT INTO `menu_closure` VALUES (5384, 2273, 2399);
INSERT INTO `menu_closure` VALUES (5385, 2386, 2399);
INSERT INTO `menu_closure` VALUES (5386, 2395, 2399);
INSERT INTO `menu_closure` VALUES (5387, 2273, 2400);
INSERT INTO `menu_closure` VALUES (5388, 2386, 2400);
INSERT INTO `menu_closure` VALUES (5389, 2395, 2400);
INSERT INTO `menu_closure` VALUES (5390, 2273, 2402);
INSERT INTO `menu_closure` VALUES (5391, 2386, 2402);
INSERT INTO `menu_closure` VALUES (5392, 2401, 2402);
INSERT INTO `menu_closure` VALUES (5393, 2273, 2403);
INSERT INTO `menu_closure` VALUES (5394, 2386, 2403);
INSERT INTO `menu_closure` VALUES (5395, 2401, 2403);
INSERT INTO `menu_closure` VALUES (5396, 2273, 2404);
INSERT INTO `menu_closure` VALUES (5397, 2386, 2404);
INSERT INTO `menu_closure` VALUES (5398, 2401, 2404);
INSERT INTO `menu_closure` VALUES (5399, 2273, 2405);
INSERT INTO `menu_closure` VALUES (5400, 2386, 2405);
INSERT INTO `menu_closure` VALUES (5401, 2401, 2405);
INSERT INTO `menu_closure` VALUES (5402, 2273, 2406);
INSERT INTO `menu_closure` VALUES (5403, 2386, 2406);
INSERT INTO `menu_closure` VALUES (5404, 2401, 2406);
INSERT INTO `menu_closure` VALUES (5405, 2273, 2407);
INSERT INTO `menu_closure` VALUES (5406, 2386, 2407);
INSERT INTO `menu_closure` VALUES (5407, 2401, 2407);
INSERT INTO `menu_closure` VALUES (5408, 2273, 2409);
INSERT INTO `menu_closure` VALUES (5409, 2386, 2409);
INSERT INTO `menu_closure` VALUES (5410, 2408, 2409);
INSERT INTO `menu_closure` VALUES (5411, 2273, 2410);
INSERT INTO `menu_closure` VALUES (5412, 2386, 2410);
INSERT INTO `menu_closure` VALUES (5413, 2408, 2410);
INSERT INTO `menu_closure` VALUES (5414, 2273, 2411);
INSERT INTO `menu_closure` VALUES (5415, 2386, 2411);
INSERT INTO `menu_closure` VALUES (5416, 2408, 2411);
INSERT INTO `menu_closure` VALUES (5417, 2273, 2412);
INSERT INTO `menu_closure` VALUES (5418, 2386, 2412);
INSERT INTO `menu_closure` VALUES (5419, 2408, 2412);
INSERT INTO `menu_closure` VALUES (5420, 2273, 2413);
INSERT INTO `menu_closure` VALUES (5421, 2386, 2413);
INSERT INTO `menu_closure` VALUES (5422, 2408, 2413);
INSERT INTO `menu_closure` VALUES (5423, 2273, 2414);
INSERT INTO `menu_closure` VALUES (5424, 2386, 2414);
INSERT INTO `menu_closure` VALUES (5425, 2408, 2414);
INSERT INTO `menu_closure` VALUES (5426, 2273, 2417);
INSERT INTO `menu_closure` VALUES (5427, 2386, 2417);
INSERT INTO `menu_closure` VALUES (5428, 2415, 2417);
INSERT INTO `menu_closure` VALUES (5429, 2416, 2417);
INSERT INTO `menu_closure` VALUES (5430, 2273, 2418);
INSERT INTO `menu_closure` VALUES (5431, 2386, 2418);
INSERT INTO `menu_closure` VALUES (5432, 2415, 2418);
INSERT INTO `menu_closure` VALUES (5433, 2416, 2418);
INSERT INTO `menu_closure` VALUES (5434, 2273, 2419);
INSERT INTO `menu_closure` VALUES (5435, 2386, 2419);
INSERT INTO `menu_closure` VALUES (5436, 2415, 2419);
INSERT INTO `menu_closure` VALUES (5437, 2416, 2419);
INSERT INTO `menu_closure` VALUES (5438, 2273, 2420);
INSERT INTO `menu_closure` VALUES (5439, 2386, 2420);
INSERT INTO `menu_closure` VALUES (5440, 2415, 2420);
INSERT INTO `menu_closure` VALUES (5441, 2416, 2420);
INSERT INTO `menu_closure` VALUES (5442, 2273, 2421);
INSERT INTO `menu_closure` VALUES (5443, 2386, 2421);
INSERT INTO `menu_closure` VALUES (5444, 2415, 2421);
INSERT INTO `menu_closure` VALUES (5445, 2416, 2421);
INSERT INTO `menu_closure` VALUES (5446, 2273, 2422);
INSERT INTO `menu_closure` VALUES (5447, 2386, 2422);
INSERT INTO `menu_closure` VALUES (5448, 2415, 2422);
INSERT INTO `menu_closure` VALUES (5449, 2416, 2422);
INSERT INTO `menu_closure` VALUES (5450, 2273, 2423);
INSERT INTO `menu_closure` VALUES (5451, 2386, 2423);
INSERT INTO `menu_closure` VALUES (5452, 2415, 2423);
INSERT INTO `menu_closure` VALUES (5453, 2416, 2423);
INSERT INTO `menu_closure` VALUES (5454, 2273, 2425);
INSERT INTO `menu_closure` VALUES (5455, 2386, 2425);
INSERT INTO `menu_closure` VALUES (5456, 2415, 2425);
INSERT INTO `menu_closure` VALUES (5457, 2424, 2425);
INSERT INTO `menu_closure` VALUES (5458, 2273, 2426);
INSERT INTO `menu_closure` VALUES (5459, 2386, 2426);
INSERT INTO `menu_closure` VALUES (5460, 2415, 2426);
INSERT INTO `menu_closure` VALUES (5461, 2424, 2426);
INSERT INTO `menu_closure` VALUES (5462, 2273, 2416);
INSERT INTO `menu_closure` VALUES (5463, 2386, 2416);
INSERT INTO `menu_closure` VALUES (5464, 2415, 2416);
INSERT INTO `menu_closure` VALUES (5465, 2273, 2424);
INSERT INTO `menu_closure` VALUES (5466, 2386, 2424);
INSERT INTO `menu_closure` VALUES (5467, 2415, 2424);
INSERT INTO `menu_closure` VALUES (5468, 2273, 2387);
INSERT INTO `menu_closure` VALUES (5469, 2386, 2387);
INSERT INTO `menu_closure` VALUES (5470, 2273, 2395);
INSERT INTO `menu_closure` VALUES (5471, 2386, 2395);
INSERT INTO `menu_closure` VALUES (5472, 2273, 2401);
INSERT INTO `menu_closure` VALUES (5473, 2386, 2401);
INSERT INTO `menu_closure` VALUES (5474, 2273, 2408);
INSERT INTO `menu_closure` VALUES (5475, 2386, 2408);
INSERT INTO `menu_closure` VALUES (5476, 2273, 2415);
INSERT INTO `menu_closure` VALUES (5477, 2386, 2415);
INSERT INTO `menu_closure` VALUES (5478, 2273, 2274);
INSERT INTO `menu_closure` VALUES (5479, 2273, 2275);
INSERT INTO `menu_closure` VALUES (5480, 2273, 2334);
INSERT INTO `menu_closure` VALUES (5481, 2273, 2349);
INSERT INTO `menu_closure` VALUES (5482, 2273, 2386);
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
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4  COMMENT='角色信息';

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
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4  COMMENT='角色层级信息';

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
) ENGINE=InnoDB AUTO_INCREMENT=219 DEFAULT CHARSET=utf8mb4  COMMENT='角色菜单信息';

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
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4  COMMENT='用户信息';

-- ----------------------------
-- Records of user
-- ----------------------------
BEGIN;
INSERT INTO `user` VALUES (1, 1, 1, '超级管理员', '超级管理员', 'F', '2a0786fe9127b8116bc30ed2ce9581e2', '1280291001@qq.com', '18888888888', '$2a$10$9qRJe9KQo6sEcU8ipKg.e.dkl2E7Wy64SigYlgraTAn.1paHFq6W.', 1, '{\"theme\":\"light\",\"themeColor\":\"#165DFF\",\"skin\":\"default\",\"tabBar\":true,\"menuWidth\":200,\"layout\":\"default\",\"language\":\"zh_CN\",\"animation\":\"gp-fade\"}', 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJkZXBhcnRtZW50SWQiOjEsImRlcGFydG1lbnRLZXl3b3JkIjoiY29tcGFueSIsImV4cCI6MTcxOTM4NTM1NSwiaWF0IjoxNzE5Mzc4MTU0LCJuYmYiOjE3MTkzNzgxNTQsInJvbGVJZCI6MSwicm9sZUtleXdvcmQiOiJzdXBlckFkbWluIiwidXNlcklkIjoxfQ.6uNK-x39GFfK_1FCUpbcfPz83W9gdn_Xtq1zhjGM0xc', 1719378154, 1713706137, 1719378154);
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
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8mb4  COMMENT='用户职位信息';

-- ----------------------------
-- Records of user_job
-- ----------------------------
BEGIN;
INSERT INTO `user_job` VALUES (1, 1, 1);
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
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4  COMMENT='用户角色信息';

-- ----------------------------
-- Records of user_role
-- ----------------------------
BEGIN;
INSERT INTO `user_role` VALUES (1, 1, 1);
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
