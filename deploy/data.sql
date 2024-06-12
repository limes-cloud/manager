
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
) ENGINE=InnoDB AUTO_INCREMENT=591 DEFAULT CHARSET=utf8mb4;

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
) ENGINE=InnoDB AUTO_INCREMENT=759 DEFAULT CHARSET=utf8mb4 COMMENT='菜单信息';

-- ----------------------------
-- Records of menu
-- ----------------------------
BEGIN;
INSERT INTO `menu` VALUES (684, 0, '管理平台', 'R', 'SystemPlatform', 'apps', NULL, NULL, '/', NULL, 'Layout', NULL, 2, 0, NULL, NULL, NULL, 1718083051, 1718083051);
INSERT INTO `menu` VALUES (685, 684, '首页面板', 'M', 'Dashboard', 'dashboard', NULL, NULL, '/dashboard', NULL, '/dashboard/index', NULL, 0, 0, 1, 1, 1, 1718083051, 1718083051);
INSERT INTO `menu` VALUES (686, 684, '管理中心', 'M', 'SystemPlatformManager', 'apps', NULL, NULL, '/manager', NULL, NULL, NULL, 0, 0, NULL, NULL, NULL, 1718083051, 1718083051);
INSERT INTO `menu` VALUES (687, 686, '基础接口', 'G', 'ManagerBaseApi', 'apps', NULL, NULL, NULL, NULL, NULL, NULL, 0, 1, NULL, NULL, NULL, 1718083051, 1718083051);
INSERT INTO `menu` VALUES (688, 687, '获取用户可见部门树', 'BA', NULL, NULL, '/manager/api/v1/departments', 'GET', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL, 1718083051, 1718083051);
INSERT INTO `menu` VALUES (689, 687, '获取用户可见角色树', 'BA', NULL, NULL, '/manager/api/v1/roles', 'GET', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL, 1718083051, 1718083051);
INSERT INTO `menu` VALUES (690, 687, '获取个人用户信息', 'BA', NULL, NULL, '/manager/api/v1/user/current', 'GET', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL, 1718083051, 1718083051);
INSERT INTO `menu` VALUES (691, 687, '更新个人用户信息', 'BA', NULL, NULL, '/manager/api/v1/user/current/info', 'PUT', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL, 1718083051, 1718083051);
INSERT INTO `menu` VALUES (692, 687, '更新个人用户密码', 'BA', NULL, NULL, '/manager/api/v1/user/current/password', 'PUT', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL, 1718083051, 1718083051);
INSERT INTO `menu` VALUES (693, 687, '保存个人用户设置', 'BA', NULL, NULL, '/manager/api/v1/user/current/setting', 'PUT', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL, 1718083051, 1718083051);
INSERT INTO `menu` VALUES (694, 687, '发送用户验证吗', 'BA', NULL, NULL, '/manager/api/v1/user/current/captcha', 'POST', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL, 1718083051, 1718083051);
INSERT INTO `menu` VALUES (695, 687, '获取用户当前角色菜单', 'BA', NULL, NULL, '/manager/api/v1/menus/by/cur_role', 'GET', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL, 1718083051, 1718083051);
INSERT INTO `menu` VALUES (696, 687, '退出登录', 'BA', NULL, NULL, '/manager/api/v1/user/logout', 'POST', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL, 1718083051, 1718083051);
INSERT INTO `menu` VALUES (697, 687, '刷新token', 'BA', NULL, NULL, '/manager/api/v1/user/token/refresh', 'POST', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL, 1718083051, 1718083051);
INSERT INTO `menu` VALUES (698, 687, '用户登录', 'BA', NULL, NULL, '/manager/api/v1/user/login', 'POST', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL, 1718083051, 1718083051);
INSERT INTO `menu` VALUES (699, 687, '获取登录验证码', 'BA', NULL, NULL, '/manager/api/v1/user/login/captcha', 'POST', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL, 1718083051, 1718083051);
INSERT INTO `menu` VALUES (700, 687, '获取系统基础设置', 'BA', NULL, NULL, '/manager/api/v1/system/setting', 'GET', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL, 1718083051, 1718083051);
INSERT INTO `menu` VALUES (701, 687, '接口鉴权', 'BA', NULL, NULL, '/manager/api/v1/auth', 'POST', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL, 1718083051, 1718083051);
INSERT INTO `menu` VALUES (702, 687, '切换用户角色', 'BA', NULL, NULL, '/manager/api/v1/user/current/role', 'POST', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL, 1718083051, 1718083051);
INSERT INTO `menu` VALUES (703, 687, '分块上传文件', 'BA', NULL, NULL, '/resource/api/v1/upload', 'POST', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL, 1718083051, 1718083051);
INSERT INTO `menu` VALUES (704, 687, '预上传文件', 'BA', NULL, NULL, '/resource/api/v1/prepare_upload', 'POST', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL, 1718083051, 1718083051);
INSERT INTO `menu` VALUES (705, 686, '字典管理', 'M', 'ManagerDictionary', 'storage', NULL, NULL, '/manager/dictionary', NULL, '/manager/dictionary/index', NULL, 0, NULL, 1, NULL, 1, 1718083051, 1718083051);
INSERT INTO `menu` VALUES (706, 705, '查询字典', 'A', NULL, NULL, '/manager/api/v1/dictionaries', 'GET', NULL, 'manager:dictionary:query', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1718083051, 1718083051);
INSERT INTO `menu` VALUES (707, 705, '新增字典', 'A', NULL, NULL, '/manager/api/v1/dictionary', 'POST', NULL, 'manager:dictionary:add', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1718083051, 1718083051);
INSERT INTO `menu` VALUES (708, 705, '修改字典', 'A', NULL, NULL, '/manager/api/v1/dictionary', 'PUT', NULL, 'manager:dictionary:update', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1718083051, 1718083051);
INSERT INTO `menu` VALUES (709, 705, '删除字典', 'A', NULL, NULL, '/manager/api/v1/dictionary', 'DELETE', NULL, 'manager:dictionary:delete', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1718083051, 1718083051);
INSERT INTO `menu` VALUES (710, 705, '获取字典值', 'A', NULL, NULL, '/manager/api/v1/dictionary_values', 'GET', NULL, 'manager:dictionary:value:query', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1718083051, 1718083051);
INSERT INTO `menu` VALUES (711, 705, '新增字典值', 'A', NULL, NULL, '/manager/api/v1/dictionary_value', 'POST', NULL, 'manager:dictionary:value:add', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1718083051, 1718083051);
INSERT INTO `menu` VALUES (712, 705, '修改字典值', 'A', NULL, NULL, '/manager/api/v1/dictionary_value', 'PUT', NULL, 'manager:dictionary:value:update', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1718083051, 1718083051);
INSERT INTO `menu` VALUES (713, 705, '更新字典值目录状态', 'A', NULL, NULL, '/manager/api/v1/dictionary_value/status', 'PUT', NULL, 'manager:dictionary:value:status', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1718083051, 1718083051);
INSERT INTO `menu` VALUES (714, 705, '删除字典值', 'A', NULL, NULL, '/manager/api/v1/dictionary_value', 'DELETE', NULL, 'manager:dictionary:value:delete', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1718083051, 1718083051);
INSERT INTO `menu` VALUES (715, 686, '菜单管理', 'M', 'ManagerMenu', 'menu', NULL, NULL, '/manager/menu', NULL, '/manager/menu/index', NULL, 0, 0, 1, 1, 1, 1718083051, 1718083051);
INSERT INTO `menu` VALUES (716, 715, '查询菜单', 'A', NULL, NULL, '/manager/api/v1/menus', 'GET', NULL, 'manager:menu:query', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1718083051, 1718083051);
INSERT INTO `menu` VALUES (717, 715, '新增菜单', 'A', NULL, NULL, '/manager/api/v1/menu', 'POST', NULL, 'manager:menu:add', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1718083051, 1718083051);
INSERT INTO `menu` VALUES (718, 715, '修改菜单', 'A', NULL, NULL, '/manager/api/v1/menu', 'PUT', NULL, 'manager:menu:update', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1718083051, 1718083051);
INSERT INTO `menu` VALUES (719, 715, '删除菜单', 'A', NULL, NULL, '/manager/api/v1/menu', 'DELETE', NULL, 'manager:menu:delete', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1718083051, 1718083051);
INSERT INTO `menu` VALUES (720, 686, '职位管理', 'M', 'ManagerJob', 'tag', NULL, NULL, '/manager/job', NULL, '/manager/job/index', NULL, 0, NULL, 1, NULL, 1, 1718083051, 1718083051);
INSERT INTO `menu` VALUES (721, 720, '查询职位', 'A', NULL, NULL, '/manager/api/v1/jobs', 'GET', NULL, 'manager:job:query', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1718083051, 1718083051);
INSERT INTO `menu` VALUES (722, 720, '新增职位', 'A', NULL, NULL, '/manager/api/v1/job', 'POST', NULL, 'manager:job:add', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1718083051, 1718083051);
INSERT INTO `menu` VALUES (723, 720, '修改职位', 'A', NULL, NULL, '/manager/api/v1/job', 'PUT', NULL, 'manager:job:update', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1718083051, 1718083051);
INSERT INTO `menu` VALUES (724, 720, '删除职位', 'A', NULL, NULL, '/manager/api/v1/job', 'DELETE', NULL, 'manager:job:delete', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1718083051, 1718083051);
INSERT INTO `menu` VALUES (725, 686, '部门管理', 'M', 'ManagerDepartment', 'user-group', NULL, NULL, '/manager/department', NULL, '/manager/department/index', NULL, 0, 0, 1, NULL, 1, 1718083051, 1718083051);
INSERT INTO `menu` VALUES (726, 725, '新增部门', 'A', NULL, NULL, '/manager/api/v1/department', 'POST', NULL, 'manager:department:add', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1718083051, 1718083051);
INSERT INTO `menu` VALUES (727, 725, '修改部门', 'A', NULL, NULL, '/manager/api/v1/department', 'PUT', NULL, 'manager:department:update', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1718083051, 1718083051);
INSERT INTO `menu` VALUES (728, 725, '删除部门', 'A', NULL, NULL, '/manager/api/v1/department', 'DELETE', NULL, 'manager:department:delete', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1718083051, 1718083051);
INSERT INTO `menu` VALUES (729, 686, '角色管理', 'M', 'ManagerRole', 'safe', NULL, NULL, '/manager/role', NULL, '/manager/role/index', NULL, 0, NULL, 0, NULL, 1, 1718083051, 1718083051);
INSERT INTO `menu` VALUES (730, 729, '新增角色', 'A', NULL, NULL, '/manager/api/v1/role', 'POST', NULL, 'manager:role:add', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1718083051, 1718083051);
INSERT INTO `menu` VALUES (731, 729, '修改角色', 'A', NULL, NULL, '/manager/api/v1/role', 'PUT', NULL, 'manager:role:update', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1718083051, 1718083051);
INSERT INTO `menu` VALUES (732, 729, '删除角色', 'A', NULL, NULL, '/manager/api/v1/role', 'DELETE', NULL, 'manager:role:delete', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1718083051, 1718083051);
INSERT INTO `menu` VALUES (733, 729, '角色菜单管理', 'G', NULL, NULL, NULL, NULL, NULL, 'manager:role:menu', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1718083051, 1718083051);
INSERT INTO `menu` VALUES (734, 733, '查询角色菜单', 'A', NULL, NULL, '/manager/api/v1/role/menu_ids', 'GET', NULL, 'manager:role:menu:query', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1718083051, 1718083051);
INSERT INTO `menu` VALUES (735, 733, '修改角色菜单', 'A', NULL, NULL, '/manager/api/v1/role/menu', 'POST', NULL, 'manager:role:menu:update', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1718083051, 1718083051);
INSERT INTO `menu` VALUES (736, 686, '用户管理', 'M', 'ManagerUser', 'user', NULL, NULL, '/manager/user', NULL, '/manager/user/index', NULL, 0, NULL, 1, NULL, 1, 1718083051, 1718083051);
INSERT INTO `menu` VALUES (737, 736, '查询用户列表', 'A', NULL, NULL, '/manager/api/v1/users', 'GET', NULL, 'manager:user:query', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1718083051, 1718083051);
INSERT INTO `menu` VALUES (738, 736, '新增用户', 'A', NULL, NULL, '/manager/api/v1/user', 'POST', NULL, 'manager:user:add', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1718083051, 1718083051);
INSERT INTO `menu` VALUES (739, 736, '修改用户', 'A', NULL, NULL, '/manager/api/v1/user', 'PUT', NULL, 'manager:user:update', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1718083051, 1718083051);
INSERT INTO `menu` VALUES (740, 736, '删除用户', 'A', NULL, NULL, '/manager/api/v1/user', 'DELETE', NULL, 'manager:user:delete', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1718083051, 1718083051);
INSERT INTO `menu` VALUES (741, 736, '修改用户状态', 'A', NULL, NULL, '/manager/api/v1/user/status', 'PUT', NULL, 'manager:user:status', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1718083051, 1718083051);
INSERT INTO `menu` VALUES (742, 736, '重置账号密码', 'A', NULL, NULL, '/manager/api/v1/user/password/reset', 'POST', NULL, 'manager:user:reset:password', '', NULL, 0, NULL, NULL, NULL, NULL, 1718083051, 1718083051);
INSERT INTO `menu` VALUES (743, 684, '资源中心', 'M', 'SystemPlatformResource', 'file', NULL, NULL, '/resource', NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL, 1718083051, 1718083051);
INSERT INTO `menu` VALUES (744, 743, '文件管理', 'M', 'ResourceFile', 'file', NULL, NULL, '/resource/file', NULL, '/resource/file/index', NULL, 0, 0, 1, NULL, 1, 1718083051, 1718083051);
INSERT INTO `menu` VALUES (745, 744, '目录管理', 'G', NULL, NULL, NULL, NULL, NULL, 'resource:directory:group', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1718083051, 1718083051);
INSERT INTO `menu` VALUES (746, 745, '查看目录', 'A', NULL, NULL, '/resource/api/v1/directories', 'GET', NULL, 'resource:directory:query', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1718083051, 1718083051);
INSERT INTO `menu` VALUES (747, 745, '新增目录', 'A', NULL, NULL, '/resource/api/v1/directory', 'POST', NULL, 'resource:directory:add', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1718083051, 1718083051);
INSERT INTO `menu` VALUES (748, 745, '修改目录', 'A', NULL, NULL, '/resource/api/v1/directory', 'PUT', NULL, 'resource:directory:update', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1718083051, 1718083051);
INSERT INTO `menu` VALUES (749, 745, '删除目录', 'A', NULL, NULL, '/resource/api/v1/directory', 'DELETE', NULL, 'resource:directory:delete', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1718083051, 1718083051);
INSERT INTO `menu` VALUES (750, 744, '文件管理', 'G', NULL, NULL, NULL, NULL, NULL, 'resource:file:group', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1718083051, 1718083051);
INSERT INTO `menu` VALUES (751, 750, '查看文件', 'A', NULL, NULL, '/resource/api/v1/files', 'GET', NULL, 'resource:file:query', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1718083051, 1718083051);
INSERT INTO `menu` VALUES (752, 750, '修改文件', 'A', NULL, NULL, '/resource/api/v1/file', 'PUT', NULL, 'resource:file:update', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1718083051, 1718083051);
INSERT INTO `menu` VALUES (753, 750, '删除文件', 'A', NULL, NULL, '/resource/api/v1/file', 'DELETE', NULL, 'resource:file:delete', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1718083051, 1718083051);
INSERT INTO `menu` VALUES (754, 743, '导出管理', 'M', 'ResourceExport', 'export', NULL, NULL, '/resource/export', NULL, '/resource/export/index', NULL, 0, 0, 1, NULL, 1, 1718083051, 1718083051);
INSERT INTO `menu` VALUES (755, 754, '查看导出', 'A', NULL, NULL, '/resource/api/v1/exports', 'GET', NULL, 'resource:export:query', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1718083051, 1718083051);
INSERT INTO `menu` VALUES (756, 754, '新增导出', 'A', NULL, NULL, '/resource/api/v1/export', 'POST', NULL, 'resource:export:file', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1718083051, 1718083051);
INSERT INTO `menu` VALUES (757, 754, '修改导出', 'A', NULL, NULL, '/resource/api/v1/export', 'PUT', NULL, 'resource:export:update', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1718083051, 1718083051);
INSERT INTO `menu` VALUES (758, 754, '删除导出', 'A', NULL, NULL, '/resource/api/v1/export', 'DELETE', NULL, 'resource:export:delete', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1718083051, 1718083051);
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
) ENGINE=InnoDB AUTO_INCREMENT=642 DEFAULT CHARSET=utf8mb4 COMMENT='菜单层级信息';

-- ----------------------------
-- Records of menu_closure
-- ----------------------------
BEGIN;
INSERT INTO `menu_closure` VALUES (426, 684, 688);
INSERT INTO `menu_closure` VALUES (427, 686, 688);
INSERT INTO `menu_closure` VALUES (428, 687, 688);
INSERT INTO `menu_closure` VALUES (429, 684, 689);
INSERT INTO `menu_closure` VALUES (430, 686, 689);
INSERT INTO `menu_closure` VALUES (431, 687, 689);
INSERT INTO `menu_closure` VALUES (432, 684, 690);
INSERT INTO `menu_closure` VALUES (433, 686, 690);
INSERT INTO `menu_closure` VALUES (434, 687, 690);
INSERT INTO `menu_closure` VALUES (435, 684, 691);
INSERT INTO `menu_closure` VALUES (436, 686, 691);
INSERT INTO `menu_closure` VALUES (437, 687, 691);
INSERT INTO `menu_closure` VALUES (438, 684, 692);
INSERT INTO `menu_closure` VALUES (439, 686, 692);
INSERT INTO `menu_closure` VALUES (440, 687, 692);
INSERT INTO `menu_closure` VALUES (441, 684, 693);
INSERT INTO `menu_closure` VALUES (442, 686, 693);
INSERT INTO `menu_closure` VALUES (443, 687, 693);
INSERT INTO `menu_closure` VALUES (444, 684, 694);
INSERT INTO `menu_closure` VALUES (445, 686, 694);
INSERT INTO `menu_closure` VALUES (446, 687, 694);
INSERT INTO `menu_closure` VALUES (447, 684, 695);
INSERT INTO `menu_closure` VALUES (448, 686, 695);
INSERT INTO `menu_closure` VALUES (449, 687, 695);
INSERT INTO `menu_closure` VALUES (450, 684, 696);
INSERT INTO `menu_closure` VALUES (451, 686, 696);
INSERT INTO `menu_closure` VALUES (452, 687, 696);
INSERT INTO `menu_closure` VALUES (453, 684, 697);
INSERT INTO `menu_closure` VALUES (454, 686, 697);
INSERT INTO `menu_closure` VALUES (455, 687, 697);
INSERT INTO `menu_closure` VALUES (456, 684, 698);
INSERT INTO `menu_closure` VALUES (457, 686, 698);
INSERT INTO `menu_closure` VALUES (458, 687, 698);
INSERT INTO `menu_closure` VALUES (459, 684, 699);
INSERT INTO `menu_closure` VALUES (460, 686, 699);
INSERT INTO `menu_closure` VALUES (461, 687, 699);
INSERT INTO `menu_closure` VALUES (462, 684, 700);
INSERT INTO `menu_closure` VALUES (463, 686, 700);
INSERT INTO `menu_closure` VALUES (464, 687, 700);
INSERT INTO `menu_closure` VALUES (465, 684, 701);
INSERT INTO `menu_closure` VALUES (466, 686, 701);
INSERT INTO `menu_closure` VALUES (467, 687, 701);
INSERT INTO `menu_closure` VALUES (468, 684, 702);
INSERT INTO `menu_closure` VALUES (469, 686, 702);
INSERT INTO `menu_closure` VALUES (470, 687, 702);
INSERT INTO `menu_closure` VALUES (471, 684, 703);
INSERT INTO `menu_closure` VALUES (472, 686, 703);
INSERT INTO `menu_closure` VALUES (473, 687, 703);
INSERT INTO `menu_closure` VALUES (474, 684, 704);
INSERT INTO `menu_closure` VALUES (475, 686, 704);
INSERT INTO `menu_closure` VALUES (476, 687, 704);
INSERT INTO `menu_closure` VALUES (477, 684, 706);
INSERT INTO `menu_closure` VALUES (478, 686, 706);
INSERT INTO `menu_closure` VALUES (479, 705, 706);
INSERT INTO `menu_closure` VALUES (480, 684, 707);
INSERT INTO `menu_closure` VALUES (481, 686, 707);
INSERT INTO `menu_closure` VALUES (482, 705, 707);
INSERT INTO `menu_closure` VALUES (483, 684, 708);
INSERT INTO `menu_closure` VALUES (484, 686, 708);
INSERT INTO `menu_closure` VALUES (485, 705, 708);
INSERT INTO `menu_closure` VALUES (486, 684, 709);
INSERT INTO `menu_closure` VALUES (487, 686, 709);
INSERT INTO `menu_closure` VALUES (488, 705, 709);
INSERT INTO `menu_closure` VALUES (489, 684, 710);
INSERT INTO `menu_closure` VALUES (490, 686, 710);
INSERT INTO `menu_closure` VALUES (491, 705, 710);
INSERT INTO `menu_closure` VALUES (492, 684, 711);
INSERT INTO `menu_closure` VALUES (493, 686, 711);
INSERT INTO `menu_closure` VALUES (494, 705, 711);
INSERT INTO `menu_closure` VALUES (495, 684, 712);
INSERT INTO `menu_closure` VALUES (496, 686, 712);
INSERT INTO `menu_closure` VALUES (497, 705, 712);
INSERT INTO `menu_closure` VALUES (498, 684, 713);
INSERT INTO `menu_closure` VALUES (499, 686, 713);
INSERT INTO `menu_closure` VALUES (500, 705, 713);
INSERT INTO `menu_closure` VALUES (501, 684, 714);
INSERT INTO `menu_closure` VALUES (502, 686, 714);
INSERT INTO `menu_closure` VALUES (503, 705, 714);
INSERT INTO `menu_closure` VALUES (504, 684, 716);
INSERT INTO `menu_closure` VALUES (505, 686, 716);
INSERT INTO `menu_closure` VALUES (506, 715, 716);
INSERT INTO `menu_closure` VALUES (507, 684, 717);
INSERT INTO `menu_closure` VALUES (508, 686, 717);
INSERT INTO `menu_closure` VALUES (509, 715, 717);
INSERT INTO `menu_closure` VALUES (510, 684, 718);
INSERT INTO `menu_closure` VALUES (511, 686, 718);
INSERT INTO `menu_closure` VALUES (512, 715, 718);
INSERT INTO `menu_closure` VALUES (513, 684, 719);
INSERT INTO `menu_closure` VALUES (514, 686, 719);
INSERT INTO `menu_closure` VALUES (515, 715, 719);
INSERT INTO `menu_closure` VALUES (516, 684, 721);
INSERT INTO `menu_closure` VALUES (517, 686, 721);
INSERT INTO `menu_closure` VALUES (518, 720, 721);
INSERT INTO `menu_closure` VALUES (519, 684, 722);
INSERT INTO `menu_closure` VALUES (520, 686, 722);
INSERT INTO `menu_closure` VALUES (521, 720, 722);
INSERT INTO `menu_closure` VALUES (522, 684, 723);
INSERT INTO `menu_closure` VALUES (523, 686, 723);
INSERT INTO `menu_closure` VALUES (524, 720, 723);
INSERT INTO `menu_closure` VALUES (525, 684, 724);
INSERT INTO `menu_closure` VALUES (526, 686, 724);
INSERT INTO `menu_closure` VALUES (527, 720, 724);
INSERT INTO `menu_closure` VALUES (528, 684, 726);
INSERT INTO `menu_closure` VALUES (529, 686, 726);
INSERT INTO `menu_closure` VALUES (530, 725, 726);
INSERT INTO `menu_closure` VALUES (531, 684, 727);
INSERT INTO `menu_closure` VALUES (532, 686, 727);
INSERT INTO `menu_closure` VALUES (533, 725, 727);
INSERT INTO `menu_closure` VALUES (534, 684, 728);
INSERT INTO `menu_closure` VALUES (535, 686, 728);
INSERT INTO `menu_closure` VALUES (536, 725, 728);
INSERT INTO `menu_closure` VALUES (537, 684, 734);
INSERT INTO `menu_closure` VALUES (538, 686, 734);
INSERT INTO `menu_closure` VALUES (539, 729, 734);
INSERT INTO `menu_closure` VALUES (540, 733, 734);
INSERT INTO `menu_closure` VALUES (541, 684, 735);
INSERT INTO `menu_closure` VALUES (542, 686, 735);
INSERT INTO `menu_closure` VALUES (543, 729, 735);
INSERT INTO `menu_closure` VALUES (544, 733, 735);
INSERT INTO `menu_closure` VALUES (545, 684, 730);
INSERT INTO `menu_closure` VALUES (546, 686, 730);
INSERT INTO `menu_closure` VALUES (547, 729, 730);
INSERT INTO `menu_closure` VALUES (548, 684, 731);
INSERT INTO `menu_closure` VALUES (549, 686, 731);
INSERT INTO `menu_closure` VALUES (550, 729, 731);
INSERT INTO `menu_closure` VALUES (551, 684, 732);
INSERT INTO `menu_closure` VALUES (552, 686, 732);
INSERT INTO `menu_closure` VALUES (553, 729, 732);
INSERT INTO `menu_closure` VALUES (554, 684, 733);
INSERT INTO `menu_closure` VALUES (555, 686, 733);
INSERT INTO `menu_closure` VALUES (556, 729, 733);
INSERT INTO `menu_closure` VALUES (557, 684, 737);
INSERT INTO `menu_closure` VALUES (558, 686, 737);
INSERT INTO `menu_closure` VALUES (559, 736, 737);
INSERT INTO `menu_closure` VALUES (560, 684, 738);
INSERT INTO `menu_closure` VALUES (561, 686, 738);
INSERT INTO `menu_closure` VALUES (562, 736, 738);
INSERT INTO `menu_closure` VALUES (563, 684, 739);
INSERT INTO `menu_closure` VALUES (564, 686, 739);
INSERT INTO `menu_closure` VALUES (565, 736, 739);
INSERT INTO `menu_closure` VALUES (566, 684, 740);
INSERT INTO `menu_closure` VALUES (567, 686, 740);
INSERT INTO `menu_closure` VALUES (568, 736, 740);
INSERT INTO `menu_closure` VALUES (569, 684, 741);
INSERT INTO `menu_closure` VALUES (570, 686, 741);
INSERT INTO `menu_closure` VALUES (571, 736, 741);
INSERT INTO `menu_closure` VALUES (572, 684, 742);
INSERT INTO `menu_closure` VALUES (573, 686, 742);
INSERT INTO `menu_closure` VALUES (574, 736, 742);
INSERT INTO `menu_closure` VALUES (575, 684, 687);
INSERT INTO `menu_closure` VALUES (576, 686, 687);
INSERT INTO `menu_closure` VALUES (577, 684, 705);
INSERT INTO `menu_closure` VALUES (578, 686, 705);
INSERT INTO `menu_closure` VALUES (579, 684, 715);
INSERT INTO `menu_closure` VALUES (580, 686, 715);
INSERT INTO `menu_closure` VALUES (581, 684, 720);
INSERT INTO `menu_closure` VALUES (582, 686, 720);
INSERT INTO `menu_closure` VALUES (583, 684, 725);
INSERT INTO `menu_closure` VALUES (584, 686, 725);
INSERT INTO `menu_closure` VALUES (585, 684, 729);
INSERT INTO `menu_closure` VALUES (586, 686, 729);
INSERT INTO `menu_closure` VALUES (587, 684, 736);
INSERT INTO `menu_closure` VALUES (588, 686, 736);
INSERT INTO `menu_closure` VALUES (589, 684, 746);
INSERT INTO `menu_closure` VALUES (590, 743, 746);
INSERT INTO `menu_closure` VALUES (591, 744, 746);
INSERT INTO `menu_closure` VALUES (592, 745, 746);
INSERT INTO `menu_closure` VALUES (593, 684, 747);
INSERT INTO `menu_closure` VALUES (594, 743, 747);
INSERT INTO `menu_closure` VALUES (595, 744, 747);
INSERT INTO `menu_closure` VALUES (596, 745, 747);
INSERT INTO `menu_closure` VALUES (597, 684, 748);
INSERT INTO `menu_closure` VALUES (598, 743, 748);
INSERT INTO `menu_closure` VALUES (599, 744, 748);
INSERT INTO `menu_closure` VALUES (600, 745, 748);
INSERT INTO `menu_closure` VALUES (601, 684, 749);
INSERT INTO `menu_closure` VALUES (602, 743, 749);
INSERT INTO `menu_closure` VALUES (603, 744, 749);
INSERT INTO `menu_closure` VALUES (604, 745, 749);
INSERT INTO `menu_closure` VALUES (605, 684, 751);
INSERT INTO `menu_closure` VALUES (606, 743, 751);
INSERT INTO `menu_closure` VALUES (607, 744, 751);
INSERT INTO `menu_closure` VALUES (608, 750, 751);
INSERT INTO `menu_closure` VALUES (609, 684, 752);
INSERT INTO `menu_closure` VALUES (610, 743, 752);
INSERT INTO `menu_closure` VALUES (611, 744, 752);
INSERT INTO `menu_closure` VALUES (612, 750, 752);
INSERT INTO `menu_closure` VALUES (613, 684, 753);
INSERT INTO `menu_closure` VALUES (614, 743, 753);
INSERT INTO `menu_closure` VALUES (615, 744, 753);
INSERT INTO `menu_closure` VALUES (616, 750, 753);
INSERT INTO `menu_closure` VALUES (617, 684, 745);
INSERT INTO `menu_closure` VALUES (618, 743, 745);
INSERT INTO `menu_closure` VALUES (619, 744, 745);
INSERT INTO `menu_closure` VALUES (620, 684, 750);
INSERT INTO `menu_closure` VALUES (621, 743, 750);
INSERT INTO `menu_closure` VALUES (622, 744, 750);
INSERT INTO `menu_closure` VALUES (623, 684, 755);
INSERT INTO `menu_closure` VALUES (624, 743, 755);
INSERT INTO `menu_closure` VALUES (625, 754, 755);
INSERT INTO `menu_closure` VALUES (626, 684, 756);
INSERT INTO `menu_closure` VALUES (627, 743, 756);
INSERT INTO `menu_closure` VALUES (628, 754, 756);
INSERT INTO `menu_closure` VALUES (629, 684, 757);
INSERT INTO `menu_closure` VALUES (630, 743, 757);
INSERT INTO `menu_closure` VALUES (631, 754, 757);
INSERT INTO `menu_closure` VALUES (632, 684, 758);
INSERT INTO `menu_closure` VALUES (633, 743, 758);
INSERT INTO `menu_closure` VALUES (634, 754, 758);
INSERT INTO `menu_closure` VALUES (635, 684, 744);
INSERT INTO `menu_closure` VALUES (636, 743, 744);
INSERT INTO `menu_closure` VALUES (637, 684, 754);
INSERT INTO `menu_closure` VALUES (638, 743, 754);
INSERT INTO `menu_closure` VALUES (639, 684, 685);
INSERT INTO `menu_closure` VALUES (640, 684, 686);
INSERT INTO `menu_closure` VALUES (641, 684, 743);
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
INSERT INTO `user` VALUES (1, 1, 1, '超级管理员', '超级管理员', 'F', '2a0786fe9127b8116bc30ed2ce9581e2', '1280291001@qq.com', '18888888888', '$2a$10$9qRJe9KQo6sEcU8ipKg.e.dkl2E7Wy64SigYlgraTAn.1paHFq6W.', 1, '{\"theme\":\"light\",\"themeColor\":\"#165DFF\",\"skin\":\"default\",\"tabBar\":true,\"menuWidth\":300,\"layout\":\"twoColumns\",\"language\":\"zh_CN\",\"animation\":\"gp-fade\"}', 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJkZXBhcnRtZW50SWQiOjEsImRlcGFydG1lbnRLZXl3b3JkIjoiY29tcGFueSIsImV4cCI6MTcxODEyNTIyNSwiaWF0IjoxNzE4MTE4MDI0LCJuYmYiOjE3MTgxMTgwMjQsInJvbGVJZCI6MSwicm9sZUtleXdvcmQiOiJzdXBlckFkbWluIiwidXNlcklkIjoxfQ.FRC2UtxiKo29q0LsbDzVUHpQu3_J2wA01UduTVSXXbw', 1718118024, 1713706137, 1718118024);
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
