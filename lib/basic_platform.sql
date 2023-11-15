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
) ENGINE=InnoDB AUTO_INCREMENT=48 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Records of casbin_rule
-- ----------------------------
BEGIN;
INSERT INTO `casbin_rule` VALUES (29, 'p', 'admin', '/basic/v1/menu', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES (27, 'p', 'admin', '/basic/v1/menu', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (28, 'p', 'admin', '/basic/v1/menu', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES (33, 'p', 'admin', '/v1/department', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES (31, 'p', 'admin', '/v1/department', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (32, 'p', 'admin', '/v1/department', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES (30, 'p', 'admin', '/v1/department/tree', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (26, 'p', 'admin', '/v1/menu/tree', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (37, 'p', 'admin', '/v1/role', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES (35, 'p', 'admin', '/v1/role', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (36, 'p', 'admin', '/v1/role', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES (39, 'p', 'admin', '/v1/role/menu', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (38, 'p', 'admin', '/v1/role/menu/ids', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (34, 'p', 'admin', '/v1/role/tree', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (43, 'p', 'admin', '/v1/user', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES (41, 'p', 'admin', '/v1/user', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (42, 'p', 'admin', '/v1/user', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES (45, 'p', 'admin', '/v1/user/disable', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (44, 'p', 'admin', '/v1/user/enable', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (46, 'p', 'admin', '/v1/user/password/reset', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (47, 'p', 'admin', '/v1/user/roles', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (40, 'p', 'admin', '/v1/users', 'GET', '', '', '');
COMMIT;

-- ----------------------------
-- Table structure for department
-- ----------------------------
DROP TABLE IF EXISTS `department`;
CREATE TABLE `department` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '自增id',
  `parent_id` bigint unsigned NOT NULL COMMENT '父部门',
  `name` char(64) NOT NULL COMMENT '部门名称',
  `keyword` char(32) CHARACTER SET utf8mb3 COLLATE utf8mb3_bin NOT NULL COMMENT '部门关键字',
  `description` varchar(256) NOT NULL COMMENT '部门备注',
  `created_at` bigint unsigned DEFAULT NULL COMMENT '创建时间',
  `updated_at` bigint unsigned DEFAULT NULL COMMENT '修改时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `keyword` (`keyword`)
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8mb3;

-- ----------------------------
-- Records of department
-- ----------------------------
BEGIN;
INSERT INTO `department` VALUES (1, 0, '贵州青橙科技', 'company', '贵州青橙科技公司简介', 20231022184455, NULL);
INSERT INTO `department` VALUES (2, 1, '研发部', 'development', '这是一个研发部111', 1698061206, 1698061705);
COMMIT;

-- ----------------------------
-- Table structure for menu
-- ----------------------------
DROP TABLE IF EXISTS `menu`;
CREATE TABLE `menu` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '自增id',
  `app` char(32) CHARACTER SET utf8mb3 COLLATE utf8mb3_bin NOT NULL COMMENT '所属服务',
  `parent_id` bigint unsigned NOT NULL COMMENT '父菜单',
  `title` varchar(128) NOT NULL COMMENT '菜单标题',
  `type` char(32) NOT NULL COMMENT '菜单类型',
  `keyword` char(64) DEFAULT NULL COMMENT '菜单关键词',
  `icon` char(32) DEFAULT NULL COMMENT '菜单图标',
  `path` varchar(128) DEFAULT NULL COMMENT '菜单路径',
  `permission` varchar(128) DEFAULT NULL COMMENT '菜单指令',
  `component` varchar(256) DEFAULT NULL COMMENT '菜单组件地址',
  `redirect` varchar(128) DEFAULT NULL COMMENT '重定向地址',
  `api` varchar(128) DEFAULT NULL COMMENT '接口api',
  `method` char(32) DEFAULT NULL COMMENT '接口方法',
  `weight` int DEFAULT '0' COMMENT '菜单权重',
  `is_hidden` tinyint(1) DEFAULT '0' COMMENT '是否隐藏',
  `is_cache` tinyint(1) DEFAULT '0' COMMENT '是否缓存',
  `is_home` tinyint(1) DEFAULT '0' COMMENT '是否为首页',
  `is_affix` tinyint(1) DEFAULT '1' COMMENT '是否加入tag',
  `created_at` bigint unsigned DEFAULT NULL COMMENT '创建时间',
  `updated_at` bigint unsigned DEFAULT NULL COMMENT '修改时间',
  PRIMARY KEY (`id`),
) ENGINE=InnoDB AUTO_INCREMENT=44 DEFAULT CHARSET=utf8mb3;

-- ----------------------------
-- Records of menu
-- ----------------------------
BEGIN;
INSERT INTO `menu` VALUES (1, 'manager', 0, '基础平台', 'R', 'Basic', 'apps', '/basic', NULL, 'Layout', NULL, NULL, NULL, 0, 0, NULL, NULL, NULL, 1698809470, 1698988713);
INSERT INTO `menu` VALUES (2, 'manager', 1, '基础接口', 'G', 'BasicBaseApi', 'apps', NULL, 'BA', NULL, NULL, NULL, NULL, 0, 1, NULL, NULL, NULL, 1699066881, 1699098618);
INSERT INTO `menu` VALUES (3, 'manager', 1, '菜单管理', 'M', 'BasicMenu', 'menu', '/basic/menu', NULL, 'basic/menu/index', NULL, NULL, NULL, 0, 0, 1, 1, 1, 1698809470, 1698809470);
INSERT INTO `menu` VALUES (4, 'manager', 1, '部门管理', 'M', 'BasicDepartment', 'user-group', '/basic/department', NULL, 'basic/department/index', NULL, NULL, NULL, 0, 0, 1, NULL, 1, 1698809470, 1698988624);
INSERT INTO `menu` VALUES (5, 'manager', 1, '角色管理', 'M', 'BasicRole', 'safe', '/basic/role', NULL, 'basic/role/index', NULL, NULL, NULL, 0, NULL, 0, NULL, 1, 1698816203, 1698988607);
INSERT INTO `menu` VALUES (6, 'manager', 1, '用户管理', 'M', 'BasicUser', 'user', '/basic/user', NULL, 'basic/user/index', NULL, NULL, NULL, 0, NULL, 1, NULL, 1, 1698893392, 1698893392);
INSERT INTO `menu` VALUES (7, 'Configure', 0, '配置中心', 'R', 'Configure', 'code-block', '/configure', NULL, 'Layout', NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL, 1698979506, 1698979506);
INSERT INTO `menu` VALUES (8, 'Configure', 7, '环境管理', 'M', 'ConfigureEnv', 'common', '/configure/env1', NULL, 'configure/env/index', NULL, NULL, NULL, 0, NULL, 1, NULL, 1, 1698986828, 1699102077);
INSERT INTO `menu` VALUES (9, 'manager', 3, '查询菜单', 'A', NULL, NULL, '', 'basic:menu:query', NULL, NULL, '/v1/menu/tree', 'GET', 0, NULL, NULL, NULL, NULL, 1699066727, 1699088143);
INSERT INTO `menu` VALUES (12, 'manager', 3, '新增菜单', 'A', NULL, NULL, NULL, 'basic:menu:add', NULL, NULL, '/basic/v1/menu', 'POST', 0, NULL, NULL, NULL, NULL, 1699067220, 1699067220);
INSERT INTO `menu` VALUES (13, 'manager', 3, '修改菜单', 'A', NULL, NULL, NULL, 'basic:menu:update', NULL, NULL, '/basic/v1/menu', 'PUT', 0, NULL, NULL, NULL, NULL, 1699067245, 1699067245);
INSERT INTO `menu` VALUES (14, 'manager', 3, '删除菜单', 'A', NULL, NULL, NULL, 'basic:menu:delete', NULL, NULL, '/basic/v1/menu', 'DELETE', 0, NULL, NULL, NULL, NULL, 1699067278, 1699067278);
INSERT INTO `menu` VALUES (15, 'manager', 4, '查看部门', 'A', NULL, NULL, NULL, 'basic:department:query', NULL, NULL, '/v1/department/tree', 'GET', 0, NULL, NULL, NULL, NULL, 1699090258, 1699090289);
INSERT INTO `menu` VALUES (16, 'manager', 4, '新增部门', 'A', NULL, NULL, NULL, 'basic:department:add', NULL, NULL, '/v1/department', 'POST', 0, NULL, NULL, NULL, NULL, 1699090283, 1699090283);
INSERT INTO `menu` VALUES (17, 'manager', 4, '修改部门', 'A', NULL, NULL, NULL, 'basic:department:update', NULL, NULL, '/v1/department', 'PUT', 0, NULL, NULL, NULL, NULL, 1699090316, 1699090316);
INSERT INTO `menu` VALUES (18, 'manager', 4, '删除部门', 'A', NULL, NULL, NULL, 'basic:department:delete', NULL, NULL, '/v1/department', 'DELETE', 0, NULL, NULL, NULL, NULL, 1699090340, 1699090340);
INSERT INTO `menu` VALUES (19, 'manager', 5, '查询角色', 'A', NULL, NULL, NULL, 'basic:role:query', NULL, NULL, '/v1/role/tree', 'GET', 0, NULL, NULL, NULL, NULL, 1699090425, 1699090425);
INSERT INTO `menu` VALUES (20, 'manager', 5, '新增角色', 'A', NULL, NULL, NULL, 'basic:role:add', NULL, NULL, '/v1/role', 'POST', 0, NULL, NULL, NULL, NULL, 1699090629, 1699090629);
INSERT INTO `menu` VALUES (21, 'manager', 5, '修改角色', 'A', NULL, NULL, NULL, 'basic:role:update', NULL, NULL, '/v1/role', 'PUT', 0, NULL, NULL, NULL, NULL, 1699090656, 1699090656);
INSERT INTO `menu` VALUES (22, 'manager', 5, '删除角色', 'A', NULL, NULL, NULL, 'basic:role:delete', NULL, NULL, '/v1/role', 'DELETE', 0, NULL, NULL, NULL, NULL, 1699090676, 1699090676);
INSERT INTO `menu` VALUES (23, 'manager', 5, '角色菜单管理', 'G', NULL, NULL, NULL, 'basic:role:menu', NULL, NULL, '/v1/role/menu/tree', 'GET', 0, NULL, NULL, NULL, NULL, 1699090828, 1699090896);
INSERT INTO `menu` VALUES (24, 'manager', 23, '查询角色菜单', 'A', NULL, NULL, NULL, 'basic:role:menu:query', NULL, NULL, '/v1/role/menu/ids', 'GET', 0, NULL, NULL, NULL, NULL, 1699090923, 1699108809);
INSERT INTO `menu` VALUES (25, 'manager', 23, '修改角色菜单', 'A', NULL, NULL, NULL, 'basic:role:menu:update', NULL, NULL, '/v1/role/menu', 'POST', 0, NULL, NULL, NULL, NULL, 1699091405, 1699091405);
INSERT INTO `menu` VALUES (26, 'manager', 6, '查询用户', 'A', NULL, NULL, NULL, 'basic:user:query', NULL, NULL, '/v1/users', 'GET', 0, NULL, NULL, NULL, NULL, 1699091944, 1699091944);
INSERT INTO `menu` VALUES (27, 'manager', 6, '新增用户', 'A', NULL, NULL, NULL, 'basic:user:add', NULL, NULL, '/v1/user', 'POST', 0, NULL, NULL, NULL, NULL, 1699092466, 1699092476);
INSERT INTO `menu` VALUES (28, 'manager', 6, '修改用户', 'A', NULL, NULL, NULL, 'basic:user:update', NULL, NULL, '/v1/user', 'PUT', 0, NULL, NULL, NULL, NULL, 1699092495, 1699092495);
INSERT INTO `menu` VALUES (29, 'manager', 6, '删除用户', 'A', NULL, NULL, NULL, 'basic:user:delete', NULL, NULL, '/v1/user', 'DELETE', 0, NULL, NULL, NULL, NULL, 1699092515, 1699092515);
INSERT INTO `menu` VALUES (30, 'manager', 6, '启用/禁用用户', 'G', NULL, NULL, NULL, 'basic:user:status', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL, 1699092540, 1699092540);
INSERT INTO `menu` VALUES (31, 'manager', 30, '启用用户', 'A', NULL, NULL, NULL, 'basic:user:status:enable', NULL, NULL, '/v1/user/enable', 'POST', 0, NULL, NULL, NULL, NULL, 1699092863, 1699092863);
INSERT INTO `menu` VALUES (32, 'manager', 30, '禁用用户', 'A', NULL, NULL, NULL, 'basic:user:status:disable', NULL, NULL, '/v1/user/disable', 'POST', 0, NULL, NULL, NULL, NULL, 1699092902, 1699092902);
INSERT INTO `menu` VALUES (33, 'manager', 6, '重置账号密码', 'A', NULL, NULL, NULL, 'basic:user:reset:password', '', NULL, '/v1/user/password/reset', 'POST', 0, NULL, NULL, NULL, NULL, 1699093021, 1699107168);
INSERT INTO `menu` VALUES (34, 'manager', 2, '获取用户部门', 'BA', NULL, NULL, NULL, NULL, NULL, NULL, '/v1/department/tree', 'GET', 0, NULL, NULL, NULL, NULL, 1699095870, 1699095870);
INSERT INTO `menu` VALUES (35, 'manager', 2, '获取个人用户信息', 'BA', NULL, NULL, NULL, NULL, NULL, NULL, '/v1/user/current', 'GET', 0, NULL, NULL, NULL, NULL, 1699095899, 1699095899);
INSERT INTO `menu` VALUES (36, 'manager', 2, '获取用户可见角色树', 'BA', NULL, NULL, NULL, NULL, NULL, NULL, '/v1/role/tree', 'GET', 0, NULL, NULL, NULL, NULL, 1699095965, 1699109755);
INSERT INTO `menu` VALUES (37, 'manager', 2, '获取用户菜单', 'BA', NULL, NULL, NULL, NULL, NULL, NULL, '/v1/role/menu/tree', 'GET', 0, NULL, NULL, NULL, NULL, 1699096007, 1699096007);
INSERT INTO `menu` VALUES (38, 'manager', 2, '退出登录', 'BA', NULL, NULL, NULL, NULL, NULL, NULL, '/v1/logout', 'POST', 0, NULL, NULL, NULL, NULL, 1699096034, 1699096034);
INSERT INTO `menu` VALUES (39, 'manager', 2, '刷新token', 'BA', NULL, NULL, NULL, NULL, NULL, NULL, '/v1/token/refresh', 'POST', 0, NULL, NULL, NULL, NULL, 1699096056, 1699096056);
INSERT INTO `menu` VALUES (40, 'manager', 2, '用户登录', 'BA', NULL, NULL, NULL, NULL, NULL, NULL, '/v1/login', 'POST', 0, NULL, NULL, NULL, NULL, 1699108346, 1699108346);
INSERT INTO `menu` VALUES (41, 'manager', 2, '获取登录验证码', 'BA', NULL, NULL, NULL, NULL, NULL, NULL, '/v1/login/captcha', 'POST', 0, NULL, NULL, NULL, NULL, 1699108367, 1699108367);
INSERT INTO `menu` VALUES (43, 'manager', 28, '获取用户角色列表', 'A', NULL, NULL, NULL, 'basic:user:update', NULL, NULL, '/v1/user/roles', 'GET', 0, NULL, NULL, NULL, NULL, 1699110046, 1699110046);
COMMIT;

-- ----------------------------
-- Table structure for role
-- ----------------------------
DROP TABLE IF EXISTS `role`;
CREATE TABLE `role` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '自增id',
  `parent_id` bigint unsigned NOT NULL COMMENT '父角色',
  `name` char(64) NOT NULL COMMENT '角色名称',
  `keyword` char(32) CHARACTER SET utf8mb3 COLLATE utf8mb3_bin NOT NULL COMMENT '角色关键字',
  `status` tinyint(1) NOT NULL COMMENT '角色状态',
  `description` varchar(256) NOT NULL COMMENT '角色备注',
  `department_ids` text COMMENT '自定义权限部门id',
  `data_scope` char(32) NOT NULL COMMENT '数据权限',
  `created_at` bigint unsigned DEFAULT NULL COMMENT '创建时间',
  `updated_at` bigint unsigned DEFAULT NULL COMMENT '修改时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `keyword` (`keyword`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8mb3;

-- ----------------------------
-- Records of role
-- ----------------------------
BEGIN;
INSERT INTO `role` VALUES (1, 0, '超级管理员', 'super_admin', 1, '超级管理员', NULL, 'ALL', 20231022184455, NULL);
INSERT INTO `role` VALUES (6, 1, '管理员', 'admin', 1, '管理员', '1,2', 'CUR_DOWN', 1698203397, 1699097350);
COMMIT;

-- ----------------------------
-- Table structure for role_menu
-- ----------------------------
DROP TABLE IF EXISTS `role_menu`;
CREATE TABLE `role_menu` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '自增id',
  `role_id` bigint unsigned NOT NULL COMMENT '角色id',
  `menu_id` bigint unsigned NOT NULL COMMENT '菜单id',
  `created_at` bigint unsigned NOT NULL COMMENT '新增时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `role_id` (`role_id`,`menu_id`)
) ENGINE=InnoDB AUTO_INCREMENT=79 DEFAULT CHARSET=utf8mb3;

-- ----------------------------
-- Records of role_menu
-- ----------------------------
BEGIN;
INSERT INTO `role_menu` VALUES (50, 6, 9, 1699110056);
INSERT INTO `role_menu` VALUES (51, 6, 12, 1699110056);
INSERT INTO `role_menu` VALUES (52, 6, 13, 1699110056);
INSERT INTO `role_menu` VALUES (53, 6, 14, 1699110056);
INSERT INTO `role_menu` VALUES (54, 6, 3, 1699110056);
INSERT INTO `role_menu` VALUES (55, 6, 15, 1699110056);
INSERT INTO `role_menu` VALUES (56, 6, 16, 1699110056);
INSERT INTO `role_menu` VALUES (57, 6, 17, 1699110056);
INSERT INTO `role_menu` VALUES (58, 6, 18, 1699110056);
INSERT INTO `role_menu` VALUES (59, 6, 4, 1699110056);
INSERT INTO `role_menu` VALUES (60, 6, 19, 1699110056);
INSERT INTO `role_menu` VALUES (61, 6, 20, 1699110056);
INSERT INTO `role_menu` VALUES (62, 6, 21, 1699110056);
INSERT INTO `role_menu` VALUES (63, 6, 22, 1699110056);
INSERT INTO `role_menu` VALUES (64, 6, 24, 1699110056);
INSERT INTO `role_menu` VALUES (65, 6, 25, 1699110056);
INSERT INTO `role_menu` VALUES (66, 6, 23, 1699110056);
INSERT INTO `role_menu` VALUES (67, 6, 5, 1699110056);
INSERT INTO `role_menu` VALUES (68, 6, 26, 1699110056);
INSERT INTO `role_menu` VALUES (69, 6, 27, 1699110056);
INSERT INTO `role_menu` VALUES (70, 6, 29, 1699110056);
INSERT INTO `role_menu` VALUES (71, 6, 31, 1699110056);
INSERT INTO `role_menu` VALUES (72, 6, 32, 1699110056);
INSERT INTO `role_menu` VALUES (73, 6, 30, 1699110056);
INSERT INTO `role_menu` VALUES (74, 6, 33, 1699110056);
INSERT INTO `role_menu` VALUES (75, 6, 1, 1699110056);
INSERT INTO `role_menu` VALUES (76, 6, 6, 1699110056);
INSERT INTO `role_menu` VALUES (77, 6, 28, 1699110056);
INSERT INTO `role_menu` VALUES (78, 6, 43, 1699110056);
COMMIT;

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '自增id',
  `department_id` bigint unsigned NOT NULL COMMENT '部门id',
  `role_id` bigint unsigned NOT NULL COMMENT '最近登录的角色id',
  `name` char(64) NOT NULL COMMENT '用户姓名',
  `avatar` varchar(128) NOT NULL COMMENT '用户头像',
  `nickname` char(64) NOT NULL COMMENT '用户昵称',
  `gender` enum('M','F','U') NOT NULL COMMENT '用户性别',
  `phone` char(11) NOT NULL COMMENT '用户电话',
  `email` char(64) NOT NULL COMMENT '用户邮箱',
  `password` varchar(256) NOT NULL COMMENT '用户密码',
  `status` tinyint(1) DEFAULT '1' COMMENT '用户状态',
  `token` text COMMENT '用户当前token',
  `disabled` varchar(128) DEFAULT NULL COMMENT '禁用原因',
  `last_login` bigint DEFAULT NULL COMMENT '最后登录时间',
  `created_at` bigint unsigned DEFAULT NULL COMMENT '创建时间',
  `updated_at` bigint unsigned DEFAULT NULL COMMENT '修改时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `phone` (`phone`),
  UNIQUE KEY `email` (`email`),
  KEY `role_id` (`role_id`),
  KEY `department_id` (`department_id`),
  CONSTRAINT `user_ibfk_1` FOREIGN KEY (`role_id`) REFERENCES `role` (`id`),
  CONSTRAINT `user_ibfk_2` FOREIGN KEY (`department_id`) REFERENCES `department` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8mb3;

-- ----------------------------
-- Records of user
-- ----------------------------
BEGIN;
INSERT INTO `user` VALUES (1, 1, 1, '超级管理员', '', '超级管理员', 'M', '18286219254', '1280291001@qq.com', '$2a$10$CSknnZYIYbXmjtJ/.jO9B.EUGn8Qn.YOiam0A5Gt6dKqOSQEA8aIO', 1, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJkZXBhcnRtZW50X2lkIjoxLCJkZXBhcnRtZW50X2tleXdvcmQiOiJjb21wYW55IiwiZXhwIjoxNjk5MTE2OTI5LCJpYXQiOjE2OTkxMDk3MjgsIm5iZiI6MTY5OTEwOTcyOCwicm9sZV9pZCI6MSwicm9sZV9rZXl3b3JkIjoic3VwZXJfYWRtaW4iLCJ1c2VyX2lkIjoxfQ.HRjzAPJjHsssiKcpEd5NjZ89XLbwxpH66bZoXgKlGic', '', 1699109728, 1698924586, 1699109728);
INSERT INTO `user` VALUES (4, 1, 6, '测试用户', '', '测试用户', 'M', '18286219255', '128029101@qq.com', '$2a$10$rO2zRtrWIDLZtv44EwnpwOOZHVTHmQPFnjw/zxsTM.tKs7irzxJ8S', 1, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJkZXBhcnRtZW50X2lkIjoxLCJkZXBhcnRtZW50X2tleXdvcmQiOiJjb21wYW55IiwiZXhwIjoxNjk5MTE1OTY2LCJpYXQiOjE2OTkxMDg3NjUsIm5iZiI6MTY5OTEwODc2NSwicm9sZV9pZCI6Niwicm9sZV9rZXl3b3JkIjoiYWRtaW4iLCJ1c2VyX2lkIjo0fQ.nlPk494KwMKXkMD9SiNRgdvUz8DJNMBsjPTRJRWM2nc', '', 1699108765, 1698933364, 1699108765);
COMMIT;

-- ----------------------------
-- Table structure for user_role
-- ----------------------------
DROP TABLE IF EXISTS `user_role`;
CREATE TABLE `user_role` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '自增id',
  `user_id` bigint unsigned NOT NULL COMMENT '用户id',
  `role_id` bigint unsigned NOT NULL COMMENT '角色id',
  `created_at` bigint unsigned DEFAULT NULL COMMENT '创建时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `user_id` (`user_id`,`role_id`),
  KEY `role_id` (`role_id`),
  CONSTRAINT `user_role_ibfk_1` FOREIGN KEY (`role_id`) REFERENCES `role` (`id`) ON DELETE CASCADE,
  CONSTRAINT `user_role_ibfk_2` FOREIGN KEY (`user_id`) REFERENCES `user` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8mb3;

-- ----------------------------
-- Records of user_role
-- ----------------------------
BEGIN;
INSERT INTO `user_role` VALUES (7, 4, 6, 1699083860);
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
