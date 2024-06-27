
/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- 数据库： `manager`
--

-- --------------------------------------------------------

--
-- 表的结构 `casbin_rule`
--
DROP TABLE IF EXISTS `casbin_rule`;
CREATE TABLE `casbin_rule` (
                               `id` bigint(20) UNSIGNED NOT NULL,
                               `ptype` varchar(100) DEFAULT NULL,
                               `v0` varchar(100) DEFAULT NULL,
                               `v1` varchar(100) DEFAULT NULL,
                               `v2` varchar(100) DEFAULT NULL,
                               `v3` varchar(100) DEFAULT NULL,
                               `v4` varchar(100) DEFAULT NULL,
                               `v5` varchar(100) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- --------------------------------------------------------

--
-- 表的结构 `department`
--
DROP TABLE IF EXISTS `department`;
CREATE TABLE `department` (
                              `id` bigint(20) UNSIGNED NOT NULL COMMENT '主键ID',
                              `parent_id` bigint(20) UNSIGNED NOT NULL COMMENT '父id',
                              `keyword` char(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT '标识',
                              `name` varchar(64) NOT NULL COMMENT '名称',
                              `description` varchar(256) NOT NULL COMMENT '描述',
                              `created_at` bigint(20) UNSIGNED NOT NULL DEFAULT '0' COMMENT '创建时间',
                              `updated_at` bigint(20) UNSIGNED NOT NULL DEFAULT '0' COMMENT '修改时间'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='部门信息';

--
-- 转存表中的数据 `department`
--

INSERT INTO `department` (`id`, `parent_id`, `keyword`, `name`, `description`, `created_at`, `updated_at`) VALUES
    (1, 0, 'company', '贵州青橙科技有限公司', '开放合作，拥抱未来', 1713706137, 1713706137);

-- --------------------------------------------------------

--
-- 表的结构 `department_closure`
--
DROP TABLE IF EXISTS `department_closure`;
CREATE TABLE `department_closure` (
                                      `id` bigint(20) UNSIGNED NOT NULL COMMENT '主键ID',
                                      `parent` bigint(20) UNSIGNED NOT NULL COMMENT '部门id',
                                      `children` bigint(20) UNSIGNED NOT NULL COMMENT '部门id'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='部门层级信息';

-- --------------------------------------------------------

--
-- 表的结构 `dictionary`
--
DROP TABLE IF EXISTS `dictionary`;
CREATE TABLE `dictionary` (
                              `id` bigint(20) UNSIGNED NOT NULL COMMENT '主键ID',
                              `keyword` char(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT '关键字',
                              `name` varchar(64) NOT NULL COMMENT '名称',
                              `description` varchar(256) DEFAULT NULL COMMENT '描述',
                              `created_at` bigint(20) UNSIGNED NOT NULL DEFAULT '0' COMMENT '创建时间',
                              `updated_at` bigint(20) UNSIGNED NOT NULL DEFAULT '0' COMMENT '修改时间',
                              `deleted_at` bigint(20) UNSIGNED NOT NULL DEFAULT '0' COMMENT '删除时间'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='字典信息';

-- --------------------------------------------------------

--
-- 表的结构 `dictionary_value`
--
DROP TABLE IF EXISTS `dictionary_value`;
CREATE TABLE `dictionary_value` (
                                    `id` bigint(20) UNSIGNED NOT NULL COMMENT '主键ID',
                                    `dictionary_id` bigint(20) UNSIGNED NOT NULL COMMENT '字典id',
                                    `label` varchar(128) NOT NULL COMMENT '标签',
                                    `value` varchar(128) NOT NULL COMMENT '标识',
                                    `status` tinyint(1) DEFAULT '1' COMMENT '状态',
                                    `weight` int(10) UNSIGNED DEFAULT '0' COMMENT '权重',
                                    `type` char(32) DEFAULT NULL COMMENT '类型',
                                    `extra` varchar(512) DEFAULT NULL COMMENT '扩展信息',
                                    `description` varchar(256) DEFAULT NULL COMMENT '描述',
                                    `created_at` bigint(20) UNSIGNED NOT NULL DEFAULT '0' COMMENT '创建时间',
                                    `updated_at` bigint(20) UNSIGNED NOT NULL DEFAULT '0' COMMENT '修改时间'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='字典值信息';

-- --------------------------------------------------------

--
-- 表的结构 `gorm_init`
--
DROP TABLE IF EXISTS `gorm_init`;
CREATE TABLE `gorm_init` (
                             `id` int(10) UNSIGNED NOT NULL,
                             `init` tinyint(1) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- 转存表中的数据 `gorm_init`
--

INSERT INTO `gorm_init` (`id`, `init`) VALUES
    (1, 1);

-- --------------------------------------------------------

--
-- 表的结构 `job`
--
DROP TABLE IF EXISTS `job`;
CREATE TABLE `job` (
                       `id` bigint(20) UNSIGNED NOT NULL COMMENT '主键ID',
                       `keyword` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT '关键字',
                       `name` varchar(64) NOT NULL COMMENT '名称',
                       `weight` int(10) UNSIGNED DEFAULT NULL COMMENT '权重',
                       `description` varchar(256) NOT NULL COMMENT '描述',
                       `created_at` bigint(20) UNSIGNED NOT NULL DEFAULT '0' COMMENT '创建时间',
                       `updated_at` bigint(20) UNSIGNED NOT NULL DEFAULT '0' COMMENT '修改时间',
                       `deleted_at` bigint(20) UNSIGNED NOT NULL DEFAULT '0' COMMENT '删除时间'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='职位信息';

--
-- 转存表中的数据 `job`
--

INSERT INTO `job` (`id`, `keyword`, `name`, `weight`, `description`, `created_at`, `updated_at`, `deleted_at`) VALUES
    (1, 'chairman', '董事长', 0, '董事长', 1713706137, 1717211765, 0);

-- --------------------------------------------------------

--
-- 表的结构 `menu`
--
DROP TABLE IF EXISTS `menu`;
CREATE TABLE `menu` (
                        `id` bigint(20) UNSIGNED NOT NULL COMMENT '主键ID',
                        `parent_id` bigint(20) UNSIGNED NOT NULL COMMENT '父id',
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
                        `weight` int(10) UNSIGNED DEFAULT '0' COMMENT '权重',
                        `is_hidden` tinyint(1) DEFAULT NULL COMMENT '是否隐藏',
                        `is_cache` tinyint(1) DEFAULT NULL COMMENT '是否缓存',
                        `is_home` tinyint(1) DEFAULT NULL COMMENT '是否为首页',
                        `is_affix` tinyint(1) DEFAULT NULL COMMENT '是否为标签',
                        `created_at` bigint(20) UNSIGNED NOT NULL DEFAULT '0' COMMENT '创建时间',
                        `updated_at` bigint(20) UNSIGNED NOT NULL DEFAULT '0' COMMENT '修改时间'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='菜单信息';

--
-- 转存表中的数据 `menu`
--

INSERT INTO `menu` (`id`, `parent_id`, `title`, `type`, `keyword`, `icon`, `api`, `method`, `path`, `permission`, `component`, `redirect`, `weight`, `is_hidden`, `is_cache`, `is_home`, `is_affix`, `created_at`, `updated_at`) VALUES
                                                                                                                                                                                                                                     (2273, 0, '管理平台', 'R', 'SystemPlatform', 'apps', NULL, NULL, '/', NULL, 'Layout', NULL, 2, 0, NULL, 0, NULL, 1719290128, 1719464475),
                                                                                                                                                                                                                                     (2274, 2273, '首页面板', 'M', 'Dashboard', 'dashboard', NULL, NULL, '/dashboard', NULL, '/dashboard/index', NULL, 0, 0, 1, 1, 1, 1719290128, 1719464475),
                                                                                                                                                                                                                                     (2275, 2273, '管理中心', 'M', 'SystemPlatformManager', 'apps', NULL, NULL, '/manager', NULL, NULL, NULL, 0, 0, NULL, 0, NULL, 1719290128, 1719464475),
                                                                                                                                                                                                                                     (2276, 2275, '基础接口', 'G', 'ManagerBaseApi', 'apps', NULL, NULL, NULL, NULL, NULL, NULL, 0, 1, NULL, 0, NULL, 1719290128, 1719464475),
                                                                                                                                                                                                                                     (2277, 2276, '获取用户可见部门树', 'BA', NULL, NULL, '/manager/api/v1/departments', 'GET', NULL, NULL, NULL, NULL, 0, NULL, NULL, 0, NULL, 1719290128, 1719464475),
                                                                                                                                                                                                                                     (2278, 2276, '获取用户可见角色树', 'BA', NULL, NULL, '/manager/api/v1/roles', 'GET', NULL, NULL, NULL, NULL, 0, NULL, NULL, 0, NULL, 1719290128, 1719464475),
                                                                                                                                                                                                                                     (2279, 2276, '获取个人用户信息', 'BA', NULL, NULL, '/manager/api/v1/user/current', 'GET', NULL, NULL, NULL, NULL, 0, NULL, NULL, 0, NULL, 1719290128, 1719464475),
                                                                                                                                                                                                                                     (2280, 2276, '更新个人用户信息', 'BA', NULL, NULL, '/manager/api/v1/user/current/info', 'PUT', NULL, NULL, NULL, NULL, 0, NULL, NULL, 0, NULL, 1719290128, 1719464475),
                                                                                                                                                                                                                                     (2281, 2276, '更新个人用户密码', 'BA', NULL, NULL, '/manager/api/v1/user/current/password', 'PUT', NULL, NULL, NULL, NULL, 0, NULL, NULL, 0, NULL, 1719290128, 1719464475),
                                                                                                                                                                                                                                     (2282, 2276, '保存个人用户设置', 'BA', NULL, NULL, '/manager/api/v1/user/current/setting', 'PUT', NULL, NULL, NULL, NULL, 0, NULL, NULL, 0, NULL, 1719290128, 1719464475),
                                                                                                                                                                                                                                     (2283, 2276, '发送用户验证吗', 'BA', NULL, NULL, '/manager/api/v1/user/current/captcha', 'POST', NULL, NULL, NULL, NULL, 0, NULL, NULL, 0, NULL, 1719290128, 1719464475),
                                                                                                                                                                                                                                     (2284, 2276, '获取用户当前角色菜单', 'BA', NULL, NULL, '/manager/api/v1/menus/by/cur_role', 'GET', NULL, NULL, NULL, NULL, 0, NULL, NULL, 0, NULL, 1719290128, 1719464475),
                                                                                                                                                                                                                                     (2285, 2276, '退出登录', 'BA', NULL, NULL, '/manager/api/v1/user/logout', 'POST', NULL, NULL, NULL, NULL, 0, NULL, NULL, 0, NULL, 1719290128, 1719464475),
                                                                                                                                                                                                                                     (2286, 2276, '刷新token', 'BA', NULL, NULL, '/manager/api/v1/user/token/refresh', 'POST', NULL, NULL, NULL, NULL, 0, NULL, NULL, 0, NULL, 1719290128, 1719464475),
                                                                                                                                                                                                                                     (2287, 2276, '用户登录', 'BA', NULL, NULL, '/manager/api/v1/user/login', 'POST', NULL, NULL, NULL, NULL, 0, NULL, NULL, 0, NULL, 1719290128, 1719464475),
                                                                                                                                                                                                                                     (2288, 2276, '获取登录验证码', 'BA', NULL, NULL, '/manager/api/v1/user/login/captcha', 'POST', NULL, NULL, NULL, NULL, 0, NULL, NULL, 0, NULL, 1719290128, 1719464475),
                                                                                                                                                                                                                                     (2289, 2276, '获取系统基础设置', 'BA', NULL, NULL, '/manager/api/v1/system/setting', 'GET', NULL, NULL, NULL, NULL, 0, NULL, NULL, 0, NULL, 1719290128, 1719464475),
                                                                                                                                                                                                                                     (2290, 2276, '接口鉴权', 'BA', NULL, NULL, '/manager/api/v1/auth', 'POST', NULL, NULL, NULL, NULL, 0, NULL, NULL, 0, NULL, 1719290128, 1719464475),
                                                                                                                                                                                                                                     (2291, 2276, '切换用户角色', 'BA', NULL, NULL, '/manager/api/v1/user/current/role', 'POST', NULL, NULL, NULL, NULL, 0, NULL, NULL, 0, NULL, 1719290128, 1719464475),
                                                                                                                                                                                                                                     (2292, 2276, '分块上传文件', 'BA', NULL, NULL, '/resource/api/v1/upload', 'POST', NULL, NULL, NULL, NULL, 0, NULL, NULL, 0, NULL, 1719290128, 1719464475),
                                                                                                                                                                                                                                     (2293, 2276, '预上传文件', 'BA', NULL, NULL, '/resource/api/v1/prepare_upload', 'POST', NULL, NULL, NULL, NULL, 0, NULL, NULL, 0, NULL, 1719290128, 1719464475),
                                                                                                                                                                                                                                     (2294, 2276, '获取字段类型', 'BA', NULL, NULL, '/usercenter/api/v1/field/types', 'GET', NULL, NULL, NULL, NULL, 0, NULL, NULL, 0, NULL, 1719290128, 1719464475),
                                                                                                                                                                                                                                     (2295, 2275, '字典管理', 'M', 'ManagerDictionary', 'storage', NULL, NULL, '/manager/dictionary', NULL, '/manager/dictionary/index', NULL, 0, NULL, 1, 0, 1, 1719290128, 1719464475),
                                                                                                                                                                                                                                     (2296, 2295, '查询字典', 'A', NULL, NULL, '/manager/api/v1/dictionaries', 'GET', NULL, 'manager:dictionary:query', NULL, NULL, 0, NULL, NULL, 0, NULL, 1719290128, 1719464475),
                                                                                                                                                                                                                                     (2297, 2295, '新增字典', 'A', NULL, NULL, '/manager/api/v1/dictionary', 'POST', NULL, 'manager:dictionary:add', NULL, NULL, 0, NULL, NULL, 0, NULL, 1719290128, 1719464475),
                                                                                                                                                                                                                                     (2298, 2295, '修改字典', 'A', NULL, NULL, '/manager/api/v1/dictionary', 'PUT', NULL, 'manager:dictionary:update', NULL, NULL, 0, NULL, NULL, 0, NULL, 1719290128, 1719464475),
                                                                                                                                                                                                                                     (2299, 2295, '删除字典', 'A', NULL, NULL, '/manager/api/v1/dictionary', 'DELETE', NULL, 'manager:dictionary:delete', NULL, NULL, 0, NULL, NULL, 0, NULL, 1719290128, 1719464475),
                                                                                                                                                                                                                                     (2300, 2295, '获取字典值', 'A', NULL, NULL, '/manager/api/v1/dictionary_values', 'GET', NULL, 'manager:dictionary:value:query', NULL, NULL, 0, NULL, NULL, 0, NULL, 1719290128, 1719464475),
                                                                                                                                                                                                                                     (2301, 2295, '新增字典值', 'A', NULL, NULL, '/manager/api/v1/dictionary_value', 'POST', NULL, 'manager:dictionary:value:add', NULL, NULL, 0, NULL, NULL, 0, NULL, 1719290128, 1719464475),
                                                                                                                                                                                                                                     (2302, 2295, '修改字典值', 'A', NULL, NULL, '/manager/api/v1/dictionary_value', 'PUT', NULL, 'manager:dictionary:value:update', NULL, NULL, 0, NULL, NULL, 0, NULL, 1719290128, 1719464475),
                                                                                                                                                                                                                                     (2303, 2295, '更新字典值目录状态', 'A', NULL, NULL, '/manager/api/v1/dictionary_value/status', 'PUT', NULL, 'manager:dictionary:value:status', NULL, NULL, 0, NULL, NULL, 0, NULL, 1719290128, 1719464475),
                                                                                                                                                                                                                                     (2304, 2295, '删除字典值', 'A', NULL, NULL, '/manager/api/v1/dictionary_value', 'DELETE', NULL, 'manager:dictionary:value:delete', NULL, NULL, 0, NULL, NULL, 0, NULL, 1719290128, 1719464475),
                                                                                                                                                                                                                                     (2305, 2275, '菜单管理', 'M', 'ManagerMenu', 'menu', NULL, NULL, '/manager/menu', NULL, '/manager/menu/index', NULL, 0, 0, 1, 0, 1, 1719290128, 1719464475),
                                                                                                                                                                                                                                     (2306, 2305, '查询菜单', 'A', NULL, NULL, '/manager/api/v1/menus', 'GET', NULL, 'manager:menu:query', NULL, NULL, 0, NULL, NULL, 0, NULL, 1719290128, 1719464475),
                                                                                                                                                                                                                                     (2307, 2305, '新增菜单', 'A', NULL, NULL, '/manager/api/v1/menu', 'POST', NULL, 'manager:menu:add', NULL, NULL, 0, NULL, NULL, 0, NULL, 1719290128, 1719464475),
                                                                                                                                                                                                                                     (2308, 2305, '修改菜单', 'A', NULL, NULL, '/manager/api/v1/menu', 'PUT', NULL, 'manager:menu:update', NULL, NULL, 0, NULL, NULL, 0, NULL, 1719290128, 1719464475),
                                                                                                                                                                                                                                     (2309, 2305, '删除菜单', 'A', NULL, NULL, '/manager/api/v1/menu', 'DELETE', NULL, 'manager:menu:delete', NULL, NULL, 0, NULL, NULL, 0, NULL, 1719290128, 1719464475),
                                                                                                                                                                                                                                     (2310, 2275, '职位管理', 'M', 'ManagerJob', 'tag', NULL, NULL, '/manager/job', NULL, '/manager/job/index', NULL, 0, NULL, 1, 0, 1, 1719290128, 1719464475),
                                                                                                                                                                                                                                     (2311, 2310, '查询职位', 'A', NULL, NULL, '/manager/api/v1/jobs', 'GET', NULL, 'manager:job:query', NULL, NULL, 0, NULL, NULL, 0, NULL, 1719290128, 1719464475),
                                                                                                                                                                                                                                     (2312, 2310, '新增职位', 'A', NULL, NULL, '/manager/api/v1/job', 'POST', NULL, 'manager:job:add', NULL, NULL, 0, NULL, NULL, 0, NULL, 1719290128, 1719464475),
                                                                                                                                                                                                                                     (2313, 2310, '修改职位', 'A', NULL, NULL, '/manager/api/v1/job', 'PUT', NULL, 'manager:job:update', NULL, NULL, 0, NULL, NULL, 0, NULL, 1719290128, 1719464475),
                                                                                                                                                                                                                                     (2314, 2310, '删除职位', 'A', NULL, NULL, '/manager/api/v1/job', 'DELETE', NULL, 'manager:job:delete', NULL, NULL, 0, NULL, NULL, 0, NULL, 1719290128, 1719464475),
                                                                                                                                                                                                                                     (2315, 2275, '部门管理', 'M', 'ManagerDepartment', 'user-group', NULL, NULL, '/manager/department', NULL, '/manager/department/index', NULL, 0, 0, 1, 0, 1, 1719290128, 1719464475),
                                                                                                                                                                                                                                     (2316, 2315, '新增部门', 'A', NULL, NULL, '/manager/api/v1/department', 'POST', NULL, 'manager:department:add', NULL, NULL, 0, NULL, NULL, 0, NULL, 1719290128, 1719464475),
                                                                                                                                                                                                                                     (2317, 2315, '修改部门', 'A', NULL, NULL, '/manager/api/v1/department', 'PUT', NULL, 'manager:department:update', NULL, NULL, 0, NULL, NULL, 0, NULL, 1719290128, 1719464475),
                                                                                                                                                                                                                                     (2318, 2315, '删除部门', 'A', NULL, NULL, '/manager/api/v1/department', 'DELETE', NULL, 'manager:department:delete', NULL, NULL, 0, NULL, NULL, 0, NULL, 1719290128, 1719464475),
                                                                                                                                                                                                                                     (2319, 2275, '角色管理', 'M', 'ManagerRole', 'safe', NULL, NULL, '/manager/role', NULL, '/manager/role/index', NULL, 0, NULL, 0, 0, 1, 1719290128, 1719464475),
                                                                                                                                                                                                                                     (2320, 2319, '新增角色', 'A', NULL, NULL, '/manager/api/v1/role', 'POST', NULL, 'manager:role:add', NULL, NULL, 0, NULL, NULL, 0, NULL, 1719290128, 1719464475),
                                                                                                                                                                                                                                     (2321, 2319, '修改角色', 'A', NULL, NULL, '/manager/api/v1/role', 'PUT', NULL, 'manager:role:update', NULL, NULL, 0, NULL, NULL, 0, NULL, 1719290128, 1719464475),
                                                                                                                                                                                                                                     (2322, 2319, '修改角色状态', 'A', NULL, NULL, '/manager/api/v1/role/status', 'PUT', NULL, 'manager:role:update:status', NULL, NULL, 0, NULL, NULL, 0, NULL, 1719290128, 1719464475),
                                                                                                                                                                                                                                     (2323, 2319, '删除角色', 'A', NULL, NULL, '/manager/api/v1/role', 'DELETE', NULL, 'manager:role:delete', NULL, NULL, 0, NULL, NULL, 0, NULL, 1719290128, 1719464475),
                                                                                                                                                                                                                                     (2324, 2319, '角色菜单管理', 'G', NULL, NULL, NULL, NULL, NULL, 'manager:role:menu', NULL, NULL, 0, NULL, NULL, 0, NULL, 1719290128, 1719464475),
                                                                                                                                                                                                                                     (2325, 2324, '查询角色菜单', 'A', NULL, NULL, '/manager/api/v1/role/menu_ids', 'GET', NULL, 'manager:role:menu:query', NULL, NULL, 0, NULL, NULL, 0, NULL, 1719290128, 1719464475),
                                                                                                                                                                                                                                     (2326, 2324, '修改角色菜单', 'A', NULL, NULL, '/manager/api/v1/role/menu', 'POST', NULL, 'manager:role:menu:update', NULL, NULL, 0, NULL, NULL, 0, NULL, 1719290128, 1719464475),
                                                                                                                                                                                                                                     (2327, 2275, '用户管理', 'M', 'ManagerUser', 'user', NULL, NULL, '/manager/user', NULL, '/manager/user/index', NULL, 0, NULL, 1, 0, 1, 1719290128, 1719464475),
                                                                                                                                                                                                                                     (2328, 2327, '查询用户列表', 'A', NULL, NULL, '/manager/api/v1/users', 'GET', NULL, 'manager:user:query', NULL, NULL, 0, NULL, NULL, 0, NULL, 1719290128, 1719464475),
                                                                                                                                                                                                                                     (2329, 2327, '新增用户', 'A', NULL, NULL, '/manager/api/v1/user', 'POST', NULL, 'manager:user:add', NULL, NULL, 0, NULL, NULL, 0, NULL, 1719290128, 1719464475),
                                                                                                                                                                                                                                     (2330, 2327, '修改用户', 'A', NULL, NULL, '/manager/api/v1/user', 'PUT', NULL, 'manager:user:update', NULL, NULL, 0, NULL, NULL, 0, NULL, 1719290128, 1719464475),
                                                                                                                                                                                                                                     (2331, 2327, '删除用户', 'A', NULL, NULL, '/manager/api/v1/user', 'DELETE', NULL, 'manager:user:delete', NULL, NULL, 0, NULL, NULL, 0, NULL, 1719290128, 1719464475),
                                                                                                                                                                                                                                     (2332, 2327, '修改用户状态', 'A', NULL, NULL, '/manager/api/v1/user/status', 'PUT', NULL, 'manager:user:status', NULL, NULL, 0, NULL, NULL, 0, NULL, 1719290128, 1719464475),
                                                                                                                                                                                                                                     (2333, 2327, '重置账号密码', 'A', NULL, NULL, '/manager/api/v1/user/password/reset', 'POST', NULL, 'manager:user:reset:password', '', NULL, 0, NULL, NULL, 0, NULL, 1719290128, 1719464475),
                                                                                                                                                                                                                                     (2334, 2273, '资源中心', 'M', 'SystemPlatformResource', 'file', NULL, NULL, '/resource', NULL, NULL, NULL, 0, NULL, NULL, 0, NULL, 1719290128, 1719464475),
                                                                                                                                                                                                                                     (2335, 2334, '文件管理', 'M', 'ResourceFile', 'file', NULL, NULL, '/resource/file', NULL, '/resource/file/index', NULL, 0, 0, 1, 0, 1, 1719290128, 1719464475),
                                                                                                                                                                                                                                     (2336, 2335, '目录管理', 'G', NULL, NULL, NULL, NULL, NULL, 'resource:directory:group', NULL, NULL, 0, NULL, NULL, 0, NULL, 1719290128, 1719464475),
                                                                                                                                                                                                                                     (2337, 2336, '查看目录', 'A', NULL, NULL, '/resource/api/v1/directories', 'GET', NULL, 'resource:directory:query', NULL, NULL, 0, NULL, NULL, 0, NULL, 1719290128, 1719464475),
                                                                                                                                                                                                                                     (2338, 2336, '新增目录', 'A', NULL, NULL, '/resource/api/v1/directory', 'POST', NULL, 'resource:directory:add', NULL, NULL, 0, NULL, NULL, 0, NULL, 1719290128, 1719464475),
                                                                                                                                                                                                                                     (2339, 2336, '修改目录', 'A', NULL, NULL, '/resource/api/v1/directory', 'PUT', NULL, 'resource:directory:update', NULL, NULL, 0, NULL, NULL, 0, NULL, 1719290128, 1719464475),
                                                                                                                                                                                                                                     (2340, 2336, '删除目录', 'A', NULL, NULL, '/resource/api/v1/directory', 'DELETE', NULL, 'resource:directory:delete', NULL, NULL, 0, NULL, NULL, 0, NULL, 1719290128, 1719464475),
                                                                                                                                                                                                                                     (2341, 2335, '文件管理', 'G', NULL, NULL, NULL, NULL, NULL, 'resource:file:group', NULL, NULL, 0, NULL, NULL, 0, NULL, 1719290128, 1719464475),
                                                                                                                                                                                                                                     (2342, 2341, '查看文件', 'A', NULL, NULL, '/resource/api/v1/files', 'GET', NULL, 'resource:file:query', NULL, NULL, 0, NULL, NULL, 0, NULL, 1719290128, 1719464475),
                                                                                                                                                                                                                                     (2343, 2341, '修改文件', 'A', NULL, NULL, '/resource/api/v1/file', 'PUT', NULL, 'resource:file:update', NULL, NULL, 0, NULL, NULL, 0, NULL, 1719290128, 1719464475),
                                                                                                                                                                                                                                     (2344, 2341, '删除文件', 'A', NULL, NULL, '/resource/api/v1/file', 'DELETE', NULL, 'resource:file:delete', NULL, NULL, 0, NULL, NULL, 0, NULL, 1719290128, 1719464475),
                                                                                                                                                                                                                                     (2345, 2334, '导出管理', 'M', 'ResourceExport', 'export', NULL, NULL, '/resource/export', NULL, '/resource/export/index', NULL, 0, 0, 1, 0, 1, 1719290128, 1719464475),
                                                                                                                                                                                                                                     (2346, 2345, '查看导出', 'A', NULL, NULL, '/resource/api/v1/exports', 'GET', NULL, 'resource:export:query', NULL, NULL, 0, NULL, NULL, 0, NULL, 1719290128, 1719464475),
                                                                                                                                                                                                                                     (2347, 2345, '新增导出', 'A', NULL, NULL, '/resource/api/v1/export', 'POST', NULL, 'resource:export:file', NULL, NULL, 0, NULL, NULL, 0, NULL, 1719290128, 1719464475),
                                                                                                                                                                                                                                     (2348, 2345, '删除导出', 'A', NULL, NULL, '/resource/api/v1/export', 'DELETE', NULL, 'resource:export:delete', NULL, NULL, 0, NULL, NULL, 0, NULL, 1719290128, 1719464475),
                                                                                                                                                                                                                                     (2349, 2273, '用户中心', 'M', 'SystemPlatformUserCenter', 'user', NULL, NULL, '/usercenter', NULL, NULL, NULL, 0, NULL, NULL, 0, NULL, 1719290128, 1719464475),
                                                                                                                                                                                                                                     (2350, 2349, '授权渠道', 'M', 'UserCenterChannel', 'mind-mapping', NULL, NULL, '/usercenter/channel', NULL, '/usercenter/channel/index', NULL, 0, 0, 1, 0, 1, 1719290128, 1719464475),
                                                                                                                                                                                                                                     (2351, 2350, '查看渠道', 'A', NULL, NULL, '/usercenter/api/v1/channels', 'GET', NULL, 'uc:channel:query', NULL, NULL, 0, NULL, NULL, 0, NULL, 1719290128, 1719464475),
                                                                                                                                                                                                                                     (2352, 2350, '新增渠道', 'A', NULL, NULL, '/usercenter/api/v1/channel', 'POST', NULL, 'uc:channel:add', NULL, NULL, 0, NULL, NULL, 0, NULL, 1719290128, 1719464475),
                                                                                                                                                                                                                                     (2353, 2350, '修改渠道', 'A', NULL, NULL, '/usercenter/api/v1/channel', 'PUT', NULL, 'uc:channel:update', NULL, NULL, 0, NULL, NULL, 0, NULL, 1719290128, 1719464475),
                                                                                                                                                                                                                                     (2354, 2350, '修改渠道状态', 'A', NULL, NULL, '/usercenter/api/v1/channel/status', 'PUT', NULL, 'uc:channel:update:status', NULL, NULL, 0, NULL, NULL, 0, NULL, 1719290128, 1719464475),
                                                                                                                                                                                                                                     (2355, 2350, '删除渠道', 'A', NULL, NULL, '/usercenter/api/v1/channel', 'DELETE', NULL, 'uc:channel:delete', NULL, NULL, 0, NULL, NULL, 0, NULL, 1719290128, 1719464475),
                                                                                                                                                                                                                                     (2356, 2349, '信息字段', 'M', 'UserCenterField', 'list', NULL, NULL, '/usercenter/field', NULL, '/usercenter/field/index', NULL, 0, 0, 1, 0, 1, 1719290128, 1719464475),
                                                                                                                                                                                                                                     (2357, 2356, '查看字段列表', 'A', NULL, NULL, '/usercenter/api/v1/fields', 'GET', NULL, 'uc:field:query', NULL, NULL, 0, NULL, NULL, 0, NULL, 1719290128, 1719464475),
                                                                                                                                                                                                                                     (2358, 2356, '新增字段', 'A', NULL, NULL, '/usercenter/api/v1/field', 'POST', NULL, 'uc:field:add', NULL, NULL, 0, NULL, NULL, 0, NULL, 1719290128, 1719464475),
                                                                                                                                                                                                                                     (2359, 2356, '修改字段', 'A', NULL, NULL, '/usercenter/api/v1/field', 'PUT', NULL, 'uc:field:update', NULL, NULL, 0, NULL, NULL, 0, NULL, 1719290128, 1719464475),
                                                                                                                                                                                                                                     (2360, 2356, '修改字段状态', 'A', NULL, NULL, '/usercenter/api/v1/field/status', 'PUT', NULL, 'uc:field:update:status', NULL, NULL, 0, NULL, NULL, 0, NULL, 1719290128, 1719464475),
                                                                                                                                                                                                                                     (2361, 2356, '删除字段', 'A', NULL, NULL, '/usercenter/api/v1/field', 'DELETE', NULL, 'uc:field:delete', NULL, NULL, 0, NULL, NULL, 0, NULL, 1719290128, 1719464475),
                                                                                                                                                                                                                                     (2362, 2349, '应用管理', 'M', 'UserCenterApp', 'apps', NULL, NULL, '/usercenter/app', NULL, '/usercenter/app/index', NULL, 0, 0, 1, 0, 1, 1719290128, 1719464475),
                                                                                                                                                                                                                                     (2363, 2362, '查看应用', 'A', NULL, NULL, '/usercenter/api/v1/apps', 'GET', NULL, 'uc:app:query', NULL, NULL, 0, NULL, NULL, 0, NULL, 1719290128, 1719464475),
                                                                                                                                                                                                                                     (2364, 2362, '新增应用', 'A', NULL, NULL, '/usercenter/api/v1/app', 'POST', NULL, 'uc:app:add', NULL, NULL, 0, NULL, NULL, 0, NULL, 1719290128, 1719464475),
                                                                                                                                                                                                                                     (2365, 2362, '修改应用', 'A', NULL, NULL, '/usercenter/api/v1/app', 'PUT', NULL, 'uc:app:update', NULL, NULL, 0, NULL, NULL, 0, NULL, 1719290128, 1719464475),
                                                                                                                                                                                                                                     (2366, 2362, '修改应用状态', 'A', NULL, NULL, '/usercenter/api/v1/app/status', 'PUT', NULL, 'uc:app:update:status', NULL, NULL, 0, NULL, NULL, 0, NULL, 1719290128, 1719464475),
                                                                                                                                                                                                                                     (2367, 2362, '删除应用', 'A', NULL, NULL, '/usercenter/api/v1/app', 'DELETE', NULL, 'uc:app:delete', NULL, NULL, 0, NULL, NULL, 0, NULL, 1719290128, 1719464475),
                                                                                                                                                                                                                                     (2368, 2349, '用户管理', 'M', 'UserCenterUser', 'user', NULL, NULL, '/usercenter/user', NULL, '/usercenter/user/index', NULL, 0, 0, 1, 0, 1, 1719290128, 1719464475),
                                                                                                                                                                                                                                     (2369, 2368, '查看用户', 'A', NULL, NULL, '/usercenter/api/v1/users', 'GET', NULL, 'uc:user:query', NULL, NULL, 0, NULL, NULL, 0, NULL, 1719290128, 1719464475),
                                                                                                                                                                                                                                     (2370, 2368, '新增用户', 'A', NULL, NULL, '/usercenter/api/v1/user', 'POST', NULL, 'uc:user:add', NULL, NULL, 0, NULL, NULL, 0, NULL, 1719290128, 1719464475),
                                                                                                                                                                                                                                     (2371, 2368, '导入用户', 'A', NULL, NULL, '/usercenter/api/v1/users', 'POST', NULL, 'uc:user:import', NULL, NULL, 0, NULL, NULL, 0, NULL, 1719290128, 1719464475),
                                                                                                                                                                                                                                     (2372, 2368, '修改用户', 'A', NULL, NULL, '/usercenter/api/v1/user', 'PUT', NULL, 'uc:user:update', NULL, NULL, 0, NULL, NULL, 0, NULL, 1719290128, 1719464475),
                                                                                                                                                                                                                                     (2373, 2368, '修改用户状态', 'A', NULL, NULL, '/usercenter/api/v1/user/status', 'PUT', NULL, 'uc:user:update:status', NULL, NULL, 0, NULL, NULL, 0, NULL, 1719290128, 1719464475),
                                                                                                                                                                                                                                     (2374, 2368, '删除用户', 'A', NULL, NULL, '/usercenter/api/v1/user', 'DELETE', NULL, 'uc:user:delete', NULL, NULL, 0, NULL, NULL, 0, NULL, 1719290128, 1719464475),
                                                                                                                                                                                                                                     (2375, 2368, '查看用户详细信息', 'G', NULL, NULL, NULL, NULL, NULL, 'uc:user:more', NULL, NULL, 0, NULL, NULL, 0, NULL, 1719290128, 1719464475),
                                                                                                                                                                                                                                     (2376, 2375, '获取用户应用信息', 'A', NULL, NULL, '/usercenter/api/v1/auths', 'GET', NULL, 'uc:auth:query', NULL, NULL, 0, NULL, NULL, 0, NULL, 1719290128, 1719464475),
                                                                                                                                                                                                                                     (2377, 2375, '创建用户应用信息', 'A', NULL, NULL, '/usercenter/api/v1/auth', 'POST', NULL, 'uc:auth:add', NULL, NULL, 0, NULL, NULL, 0, NULL, 1719290128, 1719464475),
                                                                                                                                                                                                                                     (2378, 2375, '修改用户应用状态信息', 'A', NULL, NULL, '/usercenter/api/v1/auth/status', 'PUT', NULL, 'uc:auth:update:status', NULL, NULL, 0, NULL, NULL, 0, NULL, 1719290128, 1719464475),
                                                                                                                                                                                                                                     (2379, 2375, '删除用户应用信息', 'A', NULL, NULL, '/usercenter/api/v1/auth', 'DELETE', NULL, 'uc:auth:delete', NULL, NULL, 0, NULL, NULL, 0, NULL, 1719290128, 1719464475),
                                                                                                                                                                                                                                     (2380, 2375, '获取用户渠道信息', 'A', NULL, NULL, '/usercenter/api/v1/oauths', 'GET', NULL, 'uc:oauth:query', NULL, NULL, 0, NULL, NULL, 0, NULL, 1719290128, 1719464475),
                                                                                                                                                                                                                                     (2381, 2375, '删除用户渠道信息', 'A', NULL, NULL, '/usercenter/api/v1/oauth', 'DELETE', NULL, 'uc:oauth:delete', NULL, NULL, 0, NULL, NULL, 0, NULL, 1719290128, 1719464475),
                                                                                                                                                                                                                                     (2382, 2375, '获取用户扩展信息', 'A', NULL, NULL, '/usercenter/api/v1/userinfos', 'GET', NULL, 'uc:userinfo:query', NULL, NULL, 0, NULL, NULL, 0, NULL, 1719290128, 1719464475),
                                                                                                                                                                                                                                     (2383, 2375, '创建用户扩展信息', 'A', NULL, NULL, '/usercenter/api/v1/userinfo', 'POST', NULL, 'uc:userinfo:add', NULL, NULL, 0, NULL, NULL, 0, NULL, 1719290128, 1719464475),
                                                                                                                                                                                                                                     (2384, 2375, '修改用户扩展信息', 'A', NULL, NULL, '/usercenter/api/v1/userinfo', 'PUT', NULL, 'uc:userinfo:update', NULL, NULL, 0, NULL, NULL, 0, NULL, 1719290128, 1719464475),
                                                                                                                                                                                                                                     (2385, 2375, '删除用户扩展信息', 'A', NULL, NULL, '/usercenter/api/v1/userinfo', 'DELETE', NULL, 'uc:userinfo:delete', NULL, NULL, 0, NULL, NULL, 0, NULL, 1719290128, 1719464475),
                                                                                                                                                                                                                                     (2386, 2273, '配置中心', 'M', 'SystemPlatformConfigure', 'code-block', NULL, NULL, '/configure', NULL, NULL, NULL, 0, NULL, NULL, 0, NULL, 1719290128, 1719464475),
                                                                                                                                                                                                                                     (2387, 2386, '环境管理', 'M', 'ConfigureEnv', 'common', NULL, NULL, '/configure/env', NULL, '/configure/env/index', NULL, 0, NULL, 1, 0, 1, 1719290128, 1719464475),
                                                                                                                                                                                                                                     (2388, 2387, '查看环境', 'A', NULL, NULL, '/configure/api/v1/envs', 'GET', NULL, 'configure:env:query', NULL, NULL, 0, NULL, NULL, 0, NULL, 1719290128, 1719464475),
                                                                                                                                                                                                                                     (2389, 2387, '新增环境', 'A', NULL, NULL, '/configure/api/v1/env', 'POST', NULL, 'configure:env:add', NULL, NULL, 0, NULL, NULL, 0, NULL, 1719290128, 1719464475),
                                                                                                                                                                                                                                     (2390, 2387, '修改环境', 'A', NULL, NULL, '/configure/api/v1/env', 'PUT', NULL, 'configure:env:update', NULL, NULL, 0, NULL, NULL, 0, NULL, 1719290128, 1719464475),
                                                                                                                                                                                                                                     (2391, 2387, '修改环境状态', 'A', NULL, NULL, '/configure/api/v1/env/status', 'PUT', NULL, 'configure:env:update:status', NULL, NULL, 0, NULL, NULL, 0, NULL, 1719290128, 1719464475),
                                                                                                                                                                                                                                     (2392, 2387, '删除环境', 'A', NULL, NULL, '/configure/api/v1/env', 'DELETE', NULL, 'configure:env:delete', NULL, NULL, 0, NULL, NULL, 0, NULL, 1719290128, 1719464475),
                                                                                                                                                                                                                                     (2393, 2387, '查看环境Token', 'A', NULL, NULL, '/configure/api/v1/env/token', 'GET', NULL, 'configure:env:token:query', NULL, NULL, 0, NULL, NULL, 0, NULL, 1719290128, 1719464475),
                                                                                                                                                                                                                                     (2394, 2387, '重置环境token', 'A', NULL, NULL, '/configure/api/v1/env/token', 'PUT', NULL, 'configure:env:token:reset', NULL, NULL, 0, NULL, NULL, 0, NULL, 1719290128, 1719464475),
                                                                                                                                                                                                                                     (2395, 2386, '服务管理', 'M', 'ConfigureServer', 'apps', NULL, NULL, '/configure/server', NULL, '/configure/server/index', NULL, 0, NULL, 1, 0, 1, 1719290128, 1719464475),
                                                                                                                                                                                                                                     (2396, 2395, '查询服务', 'A', NULL, NULL, '/configure/api/v1/servers', 'GET', NULL, 'configure:server:query', NULL, NULL, 0, NULL, NULL, 0, NULL, 1719290128, 1719464475),
                                                                                                                                                                                                                                     (2397, 2395, '新增服务', 'A', NULL, NULL, '/configure/api/v1/server', 'POST', NULL, 'configure:server:add', NULL, NULL, 0, NULL, NULL, 0, NULL, 1719290128, 1719464475),
                                                                                                                                                                                                                                     (2398, 2395, '修改服务', 'A', NULL, NULL, '/configure/api/v1/server', 'PUT', NULL, 'configure:server:update', NULL, NULL, 0, NULL, NULL, 0, NULL, 1719290128, 1719464475),
                                                                                                                                                                                                                                     (2399, 2395, '修改服务状态', 'A', NULL, NULL, '/configure/api/v1/server/status', 'PUT', NULL, 'configure:server:update:status', NULL, NULL, 0, NULL, NULL, 0, NULL, 1719290128, 1719464475),
                                                                                                                                                                                                                                     (2400, 2395, '删除服务', 'A', NULL, NULL, '/configure/api/v1/server', 'DELETE', NULL, 'configure:server:delete', NULL, NULL, 0, NULL, NULL, 0, NULL, 1719290128, 1719464475),
                                                                                                                                                                                                                                     (2401, 2386, '业务变量', 'M', 'ConfigureBusiness', 'code', NULL, NULL, '/configure/business', NULL, '/configure/business/index', NULL, 0, NULL, 1, 0, 1, 1719290128, 1719464475),
                                                                                                                                                                                                                                     (2402, 2401, '查看业务变量', 'A', NULL, NULL, '/configure/api/v1/business', 'GET', NULL, 'configure:business:query', NULL, NULL, 0, NULL, NULL, 0, NULL, 1719290128, 1719464475),
                                                                                                                                                                                                                                     (2403, 2401, '新增业务变量', 'A', NULL, NULL, '/configure/api/v1/business', 'POST', NULL, 'configure:business:add', NULL, NULL, 0, NULL, NULL, 0, NULL, 1719290128, 1719464475),
                                                                                                                                                                                                                                     (2404, 2401, '修改业务变量', 'A', NULL, NULL, '/configure/api/v1/business', 'PUT', NULL, 'configure:business:update', NULL, NULL, 0, NULL, NULL, 0, NULL, 1719290128, 1719464475),
                                                                                                                                                                                                                                     (2405, 2401, '删除业务变量', 'A', NULL, NULL, '/configure/api/v1/business', 'DELETE', NULL, 'configure:business:delete', NULL, NULL, 0, NULL, NULL, 0, NULL, 1719290128, 1719464475),
                                                                                                                                                                                                                                     (2406, 2401, '查看业务变量值', 'A', NULL, NULL, '/configure/business/values', 'get', NULL, 'configure:business:value:query', NULL, NULL, 0, NULL, NULL, 0, NULL, 1719290128, 1719464475),
                                                                                                                                                                                                                                     (2407, 2401, '设置业务变量值', 'A', NULL, NULL, '/configure/business/values', 'PUT', NULL, 'configure:business:value:update', NULL, NULL, 0, NULL, NULL, 0, NULL, 1719290128, 1719464475),
                                                                                                                                                                                                                                     (2408, 2386, '资源变量', 'M', 'ConfigureResource', 'unordered-list', NULL, NULL, '/configure/resource', NULL, '/configure/resource/index', NULL, 0, NULL, 1, 0, 1, 1719290128, 1719464475),
                                                                                                                                                                                                                                     (2409, 2408, '查看资源', 'A', NULL, NULL, '/configure/api/v1/resources', 'GET', NULL, 'configure:resource:query', NULL, NULL, 0, NULL, NULL, 0, NULL, 1719290128, 1719464475),
                                                                                                                                                                                                                                     (2410, 2408, '新增资源', 'A', NULL, NULL, '/configure/api/v1/resource', 'POST', NULL, 'configure:resource:add', NULL, NULL, 0, NULL, NULL, 0, NULL, 1719290128, 1719464475),
                                                                                                                                                                                                                                     (2411, 2408, '修改资源', 'A', NULL, NULL, '/configure/api/v1/resource', 'PUT', NULL, 'configure:resource:update', NULL, NULL, 0, NULL, NULL, 0, NULL, 1719290128, 1719464475),
                                                                                                                                                                                                                                     (2412, 2408, '删除资源', 'A', NULL, NULL, '/configure/api/v1/resource', 'DELETE', NULL, 'configure:resource:delete', NULL, NULL, 0, NULL, NULL, 0, NULL, 1719290128, 1719464475),
                                                                                                                                                                                                                                     (2413, 2408, '查看业务变量值', 'A', NULL, NULL, '/configure/resource/values', 'get', NULL, 'configure:resource:value:query', NULL, NULL, 0, NULL, NULL, 0, NULL, 1719290128, 1719464475),
                                                                                                                                                                                                                                     (2414, 2408, '设置业务变量值', 'A', NULL, NULL, '/configure/resource/values', 'PUT', NULL, 'configure:resource:value:update', NULL, NULL, 0, NULL, NULL, 0, NULL, 1719290128, 1719464475),
                                                                                                                                                                                                                                     (2415, 2386, '配置模板', 'M', 'ConfgureTemplate', 'file', NULL, NULL, '/configure/template', NULL, '/configure/template/index', NULL, 0, NULL, 1, 0, 1, 1719290128, 1719464475),
                                                                                                                                                                                                                                     (2416, 2415, '模板管理', 'G', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, 0, NULL, NULL, 0, NULL, 1719290128, 1719464475),
                                                                                                                                                                                                                                     (2417, 2416, '查看模板历史版本', 'A', NULL, NULL, '/configure/api/v1/templates', 'GET', NULL, 'configure:template:history', NULL, NULL, 0, NULL, NULL, 0, NULL, 1719290128, 1719464475),
                                                                                                                                                                                                                                     (2418, 2416, '查看指定模板详细数据', 'A', NULL, NULL, '/configure/api/v1/template', 'GET', NULL, NULL, NULL, NULL, 0, NULL, NULL, 0, NULL, 1719290128, 1719464475),
                                                                                                                                                                                                                                     (2419, 2416, '查看当前正在使用的模板', 'A', NULL, NULL, '/configure/api/v1/template/current', 'GET', NULL, NULL, NULL, NULL, 0, NULL, NULL, 0, NULL, 1719290128, 1719464475),
                                                                                                                                                                                                                                     (2420, 2416, '提交模板', 'A', NULL, NULL, '/configure/api/v1/template', 'POST', NULL, 'configure:template:add', NULL, NULL, 0, NULL, NULL, 0, NULL, 1719290128, 1719464475),
                                                                                                                                                                                                                                     (2421, 2416, '模板对比', 'A', NULL, NULL, '/configure/api/v1/template/compare', 'POST', NULL, 'configure:template:compare', NULL, NULL, 0, NULL, NULL, 0, NULL, 1719290128, 1719464475),
                                                                                                                                                                                                                                     (2422, 2416, '切换模板', 'A', NULL, NULL, '/configure/api/v1/template/switch', 'POST', NULL, 'configure:template:switch', NULL, NULL, 0, NULL, NULL, 0, NULL, 1719290128, 1719464475),
                                                                                                                                                                                                                                     (2423, 2416, '模板预览', 'A', NULL, NULL, '/configure/api/v1/template/preview', 'POST', NULL, 'configure:template:preview', NULL, NULL, 0, NULL, NULL, 0, NULL, 1719290128, 1719464475),
                                                                                                                                                                                                                                     (2424, 2415, '配置管理', 'G', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, 0, NULL, NULL, 0, NULL, 1719290128, 1719464475),
                                                                                                                                                                                                                                     (2425, 2424, '配置对比', 'A', NULL, NULL, '/configure/api/v1/configure/compare', 'POST', NULL, 'configure:configure:compare', NULL, NULL, 0, NULL, NULL, 0, NULL, 1719290128, 1719464475),
                                                                                                                                                                                                                                     (2426, 2424, '同步配置', 'A', NULL, NULL, '/configure/api/v1/configure', 'PUT', NULL, 'configure:configure:sync', NULL, NULL, 0, NULL, NULL, 0, NULL, 1719290128, 1719464475);

-- --------------------------------------------------------

--
-- 表的结构 `menu_closure`
--

CREATE TABLE `menu_closure` (
                                `id` bigint(20) UNSIGNED NOT NULL COMMENT '主键ID',
                                `parent` bigint(20) UNSIGNED NOT NULL COMMENT '菜单id',
                                `children` bigint(20) UNSIGNED NOT NULL COMMENT '菜单id'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='菜单层级信息';

--
-- 转存表中的数据 `menu_closure`
--

INSERT INTO `menu_closure` (`id`, `parent`, `children`) VALUES
                                                            (5024, 2273, 2277),
                                                            (5025, 2275, 2277),
                                                            (5026, 2276, 2277),
                                                            (5027, 2273, 2278),
                                                            (5028, 2275, 2278),
                                                            (5029, 2276, 2278),
                                                            (5030, 2273, 2279),
                                                            (5031, 2275, 2279),
                                                            (5032, 2276, 2279),
                                                            (5033, 2273, 2280),
                                                            (5034, 2275, 2280),
                                                            (5035, 2276, 2280),
                                                            (5036, 2273, 2281),
                                                            (5037, 2275, 2281),
                                                            (5038, 2276, 2281),
                                                            (5039, 2273, 2282),
                                                            (5040, 2275, 2282),
                                                            (5041, 2276, 2282),
                                                            (5042, 2273, 2283),
                                                            (5043, 2275, 2283),
                                                            (5044, 2276, 2283),
                                                            (5045, 2273, 2284),
                                                            (5046, 2275, 2284),
                                                            (5047, 2276, 2284),
                                                            (5048, 2273, 2285),
                                                            (5049, 2275, 2285),
                                                            (5050, 2276, 2285),
                                                            (5051, 2273, 2286),
                                                            (5052, 2275, 2286),
                                                            (5053, 2276, 2286),
                                                            (5054, 2273, 2287),
                                                            (5055, 2275, 2287),
                                                            (5056, 2276, 2287),
                                                            (5057, 2273, 2288),
                                                            (5058, 2275, 2288),
                                                            (5059, 2276, 2288),
                                                            (5060, 2273, 2289),
                                                            (5061, 2275, 2289),
                                                            (5062, 2276, 2289),
                                                            (5063, 2273, 2290),
                                                            (5064, 2275, 2290),
                                                            (5065, 2276, 2290),
                                                            (5066, 2273, 2291),
                                                            (5067, 2275, 2291),
                                                            (5068, 2276, 2291),
                                                            (5069, 2273, 2292),
                                                            (5070, 2275, 2292),
                                                            (5071, 2276, 2292),
                                                            (5072, 2273, 2293),
                                                            (5073, 2275, 2293),
                                                            (5074, 2276, 2293),
                                                            (5075, 2273, 2294),
                                                            (5076, 2275, 2294),
                                                            (5077, 2276, 2294),
                                                            (5078, 2273, 2296),
                                                            (5079, 2275, 2296),
                                                            (5080, 2295, 2296),
                                                            (5081, 2273, 2297),
                                                            (5082, 2275, 2297),
                                                            (5083, 2295, 2297),
                                                            (5084, 2273, 2298),
                                                            (5085, 2275, 2298),
                                                            (5086, 2295, 2298),
                                                            (5087, 2273, 2299),
                                                            (5088, 2275, 2299),
                                                            (5089, 2295, 2299),
                                                            (5090, 2273, 2300),
                                                            (5091, 2275, 2300),
                                                            (5092, 2295, 2300),
                                                            (5093, 2273, 2301),
                                                            (5094, 2275, 2301),
                                                            (5095, 2295, 2301),
                                                            (5096, 2273, 2302),
                                                            (5097, 2275, 2302),
                                                            (5098, 2295, 2302),
                                                            (5099, 2273, 2303),
                                                            (5100, 2275, 2303),
                                                            (5101, 2295, 2303),
                                                            (5102, 2273, 2304),
                                                            (5103, 2275, 2304),
                                                            (5104, 2295, 2304),
                                                            (5105, 2273, 2306),
                                                            (5106, 2275, 2306),
                                                            (5107, 2305, 2306),
                                                            (5108, 2273, 2307),
                                                            (5109, 2275, 2307),
                                                            (5110, 2305, 2307),
                                                            (5111, 2273, 2308),
                                                            (5112, 2275, 2308),
                                                            (5113, 2305, 2308),
                                                            (5114, 2273, 2309),
                                                            (5115, 2275, 2309),
                                                            (5116, 2305, 2309),
                                                            (5117, 2273, 2311),
                                                            (5118, 2275, 2311),
                                                            (5119, 2310, 2311),
                                                            (5120, 2273, 2312),
                                                            (5121, 2275, 2312),
                                                            (5122, 2310, 2312),
                                                            (5123, 2273, 2313),
                                                            (5124, 2275, 2313),
                                                            (5125, 2310, 2313),
                                                            (5126, 2273, 2314),
                                                            (5127, 2275, 2314),
                                                            (5128, 2310, 2314),
                                                            (5129, 2273, 2316),
                                                            (5130, 2275, 2316),
                                                            (5131, 2315, 2316),
                                                            (5132, 2273, 2317),
                                                            (5133, 2275, 2317),
                                                            (5134, 2315, 2317),
                                                            (5135, 2273, 2318),
                                                            (5136, 2275, 2318),
                                                            (5137, 2315, 2318),
                                                            (5138, 2273, 2325),
                                                            (5139, 2275, 2325),
                                                            (5140, 2319, 2325),
                                                            (5141, 2324, 2325),
                                                            (5142, 2273, 2326),
                                                            (5143, 2275, 2326),
                                                            (5144, 2319, 2326),
                                                            (5145, 2324, 2326),
                                                            (5146, 2273, 2320),
                                                            (5147, 2275, 2320),
                                                            (5148, 2319, 2320),
                                                            (5149, 2273, 2321),
                                                            (5150, 2275, 2321),
                                                            (5151, 2319, 2321),
                                                            (5152, 2273, 2322),
                                                            (5153, 2275, 2322),
                                                            (5154, 2319, 2322),
                                                            (5155, 2273, 2323),
                                                            (5156, 2275, 2323),
                                                            (5157, 2319, 2323),
                                                            (5158, 2273, 2324),
                                                            (5159, 2275, 2324),
                                                            (5160, 2319, 2324),
                                                            (5161, 2273, 2328),
                                                            (5162, 2275, 2328),
                                                            (5163, 2327, 2328),
                                                            (5164, 2273, 2329),
                                                            (5165, 2275, 2329),
                                                            (5166, 2327, 2329),
                                                            (5167, 2273, 2330),
                                                            (5168, 2275, 2330),
                                                            (5169, 2327, 2330),
                                                            (5170, 2273, 2331),
                                                            (5171, 2275, 2331),
                                                            (5172, 2327, 2331),
                                                            (5173, 2273, 2332),
                                                            (5174, 2275, 2332),
                                                            (5175, 2327, 2332),
                                                            (5176, 2273, 2333),
                                                            (5177, 2275, 2333),
                                                            (5178, 2327, 2333),
                                                            (5179, 2273, 2276),
                                                            (5180, 2275, 2276),
                                                            (5181, 2273, 2295),
                                                            (5182, 2275, 2295),
                                                            (5183, 2273, 2305),
                                                            (5184, 2275, 2305),
                                                            (5185, 2273, 2310),
                                                            (5186, 2275, 2310),
                                                            (5187, 2273, 2315),
                                                            (5188, 2275, 2315),
                                                            (5189, 2273, 2319),
                                                            (5190, 2275, 2319),
                                                            (5191, 2273, 2327),
                                                            (5192, 2275, 2327),
                                                            (5193, 2273, 2337),
                                                            (5194, 2334, 2337),
                                                            (5195, 2335, 2337),
                                                            (5196, 2336, 2337),
                                                            (5197, 2273, 2338),
                                                            (5198, 2334, 2338),
                                                            (5199, 2335, 2338),
                                                            (5200, 2336, 2338),
                                                            (5201, 2273, 2339),
                                                            (5202, 2334, 2339),
                                                            (5203, 2335, 2339),
                                                            (5204, 2336, 2339),
                                                            (5205, 2273, 2340),
                                                            (5206, 2334, 2340),
                                                            (5207, 2335, 2340),
                                                            (5208, 2336, 2340),
                                                            (5209, 2273, 2342),
                                                            (5210, 2334, 2342),
                                                            (5211, 2335, 2342),
                                                            (5212, 2341, 2342),
                                                            (5213, 2273, 2343),
                                                            (5214, 2334, 2343),
                                                            (5215, 2335, 2343),
                                                            (5216, 2341, 2343),
                                                            (5217, 2273, 2344),
                                                            (5218, 2334, 2344),
                                                            (5219, 2335, 2344),
                                                            (5220, 2341, 2344),
                                                            (5221, 2273, 2336),
                                                            (5222, 2334, 2336),
                                                            (5223, 2335, 2336),
                                                            (5224, 2273, 2341),
                                                            (5225, 2334, 2341),
                                                            (5226, 2335, 2341),
                                                            (5227, 2273, 2346),
                                                            (5228, 2334, 2346),
                                                            (5229, 2345, 2346),
                                                            (5230, 2273, 2347),
                                                            (5231, 2334, 2347),
                                                            (5232, 2345, 2347),
                                                            (5233, 2273, 2348),
                                                            (5234, 2334, 2348),
                                                            (5235, 2345, 2348),
                                                            (5236, 2273, 2335),
                                                            (5237, 2334, 2335),
                                                            (5238, 2273, 2345),
                                                            (5239, 2334, 2345),
                                                            (5240, 2273, 2351),
                                                            (5241, 2349, 2351),
                                                            (5242, 2350, 2351),
                                                            (5243, 2273, 2352),
                                                            (5244, 2349, 2352),
                                                            (5245, 2350, 2352),
                                                            (5246, 2273, 2353),
                                                            (5247, 2349, 2353),
                                                            (5248, 2350, 2353),
                                                            (5249, 2273, 2354),
                                                            (5250, 2349, 2354),
                                                            (5251, 2350, 2354),
                                                            (5252, 2273, 2355),
                                                            (5253, 2349, 2355),
                                                            (5254, 2350, 2355),
                                                            (5255, 2273, 2357),
                                                            (5256, 2349, 2357),
                                                            (5257, 2356, 2357),
                                                            (5258, 2273, 2358),
                                                            (5259, 2349, 2358),
                                                            (5260, 2356, 2358),
                                                            (5261, 2273, 2359),
                                                            (5262, 2349, 2359),
                                                            (5263, 2356, 2359),
                                                            (5264, 2273, 2360),
                                                            (5265, 2349, 2360),
                                                            (5266, 2356, 2360),
                                                            (5267, 2273, 2361),
                                                            (5268, 2349, 2361),
                                                            (5269, 2356, 2361),
                                                            (5270, 2273, 2363),
                                                            (5271, 2349, 2363),
                                                            (5272, 2362, 2363),
                                                            (5273, 2273, 2364),
                                                            (5274, 2349, 2364),
                                                            (5275, 2362, 2364),
                                                            (5276, 2273, 2365),
                                                            (5277, 2349, 2365),
                                                            (5278, 2362, 2365),
                                                            (5279, 2273, 2366),
                                                            (5280, 2349, 2366),
                                                            (5281, 2362, 2366),
                                                            (5282, 2273, 2367),
                                                            (5283, 2349, 2367),
                                                            (5284, 2362, 2367),
                                                            (5285, 2273, 2376),
                                                            (5286, 2349, 2376),
                                                            (5287, 2368, 2376),
                                                            (5288, 2375, 2376),
                                                            (5289, 2273, 2377),
                                                            (5290, 2349, 2377),
                                                            (5291, 2368, 2377),
                                                            (5292, 2375, 2377),
                                                            (5293, 2273, 2378),
                                                            (5294, 2349, 2378),
                                                            (5295, 2368, 2378),
                                                            (5296, 2375, 2378),
                                                            (5297, 2273, 2379),
                                                            (5298, 2349, 2379),
                                                            (5299, 2368, 2379),
                                                            (5300, 2375, 2379),
                                                            (5301, 2273, 2380),
                                                            (5302, 2349, 2380),
                                                            (5303, 2368, 2380),
                                                            (5304, 2375, 2380),
                                                            (5305, 2273, 2381),
                                                            (5306, 2349, 2381),
                                                            (5307, 2368, 2381),
                                                            (5308, 2375, 2381),
                                                            (5309, 2273, 2382),
                                                            (5310, 2349, 2382),
                                                            (5311, 2368, 2382),
                                                            (5312, 2375, 2382),
                                                            (5313, 2273, 2383),
                                                            (5314, 2349, 2383),
                                                            (5315, 2368, 2383),
                                                            (5316, 2375, 2383),
                                                            (5317, 2273, 2384),
                                                            (5318, 2349, 2384),
                                                            (5319, 2368, 2384),
                                                            (5320, 2375, 2384),
                                                            (5321, 2273, 2385),
                                                            (5322, 2349, 2385),
                                                            (5323, 2368, 2385),
                                                            (5324, 2375, 2385),
                                                            (5325, 2273, 2369),
                                                            (5326, 2349, 2369),
                                                            (5327, 2368, 2369),
                                                            (5328, 2273, 2370),
                                                            (5329, 2349, 2370),
                                                            (5330, 2368, 2370),
                                                            (5331, 2273, 2371),
                                                            (5332, 2349, 2371),
                                                            (5333, 2368, 2371),
                                                            (5334, 2273, 2372),
                                                            (5335, 2349, 2372),
                                                            (5336, 2368, 2372),
                                                            (5337, 2273, 2373),
                                                            (5338, 2349, 2373),
                                                            (5339, 2368, 2373),
                                                            (5340, 2273, 2374),
                                                            (5341, 2349, 2374),
                                                            (5342, 2368, 2374),
                                                            (5343, 2273, 2375),
                                                            (5344, 2349, 2375),
                                                            (5345, 2368, 2375),
                                                            (5346, 2273, 2350),
                                                            (5347, 2349, 2350),
                                                            (5348, 2273, 2356),
                                                            (5349, 2349, 2356),
                                                            (5350, 2273, 2362),
                                                            (5351, 2349, 2362),
                                                            (5352, 2273, 2368),
                                                            (5353, 2349, 2368),
                                                            (5354, 2273, 2388),
                                                            (5355, 2386, 2388),
                                                            (5356, 2387, 2388),
                                                            (5357, 2273, 2389),
                                                            (5358, 2386, 2389),
                                                            (5359, 2387, 2389),
                                                            (5360, 2273, 2390),
                                                            (5361, 2386, 2390),
                                                            (5362, 2387, 2390),
                                                            (5363, 2273, 2391),
                                                            (5364, 2386, 2391),
                                                            (5365, 2387, 2391),
                                                            (5366, 2273, 2392),
                                                            (5367, 2386, 2392),
                                                            (5368, 2387, 2392),
                                                            (5369, 2273, 2393),
                                                            (5370, 2386, 2393),
                                                            (5371, 2387, 2393),
                                                            (5372, 2273, 2394),
                                                            (5373, 2386, 2394),
                                                            (5374, 2387, 2394),
                                                            (5375, 2273, 2396),
                                                            (5376, 2386, 2396),
                                                            (5377, 2395, 2396),
                                                            (5378, 2273, 2397),
                                                            (5379, 2386, 2397),
                                                            (5380, 2395, 2397),
                                                            (5381, 2273, 2398),
                                                            (5382, 2386, 2398),
                                                            (5383, 2395, 2398),
                                                            (5384, 2273, 2399),
                                                            (5385, 2386, 2399),
                                                            (5386, 2395, 2399),
                                                            (5387, 2273, 2400),
                                                            (5388, 2386, 2400),
                                                            (5389, 2395, 2400),
                                                            (5390, 2273, 2402),
                                                            (5391, 2386, 2402),
                                                            (5392, 2401, 2402),
                                                            (5393, 2273, 2403),
                                                            (5394, 2386, 2403),
                                                            (5395, 2401, 2403),
                                                            (5396, 2273, 2404),
                                                            (5397, 2386, 2404),
                                                            (5398, 2401, 2404),
                                                            (5399, 2273, 2405),
                                                            (5400, 2386, 2405),
                                                            (5401, 2401, 2405),
                                                            (5402, 2273, 2406),
                                                            (5403, 2386, 2406),
                                                            (5404, 2401, 2406),
                                                            (5405, 2273, 2407),
                                                            (5406, 2386, 2407),
                                                            (5407, 2401, 2407),
                                                            (5408, 2273, 2409),
                                                            (5409, 2386, 2409),
                                                            (5410, 2408, 2409),
                                                            (5411, 2273, 2410),
                                                            (5412, 2386, 2410),
                                                            (5413, 2408, 2410),
                                                            (5414, 2273, 2411),
                                                            (5415, 2386, 2411),
                                                            (5416, 2408, 2411),
                                                            (5417, 2273, 2412),
                                                            (5418, 2386, 2412),
                                                            (5419, 2408, 2412),
                                                            (5420, 2273, 2413),
                                                            (5421, 2386, 2413),
                                                            (5422, 2408, 2413),
                                                            (5423, 2273, 2414),
                                                            (5424, 2386, 2414),
                                                            (5425, 2408, 2414),
                                                            (5426, 2273, 2417),
                                                            (5427, 2386, 2417),
                                                            (5428, 2415, 2417),
                                                            (5429, 2416, 2417),
                                                            (5430, 2273, 2418),
                                                            (5431, 2386, 2418),
                                                            (5432, 2415, 2418),
                                                            (5433, 2416, 2418),
                                                            (5434, 2273, 2419),
                                                            (5435, 2386, 2419),
                                                            (5436, 2415, 2419),
                                                            (5437, 2416, 2419),
                                                            (5438, 2273, 2420),
                                                            (5439, 2386, 2420),
                                                            (5440, 2415, 2420),
                                                            (5441, 2416, 2420),
                                                            (5442, 2273, 2421),
                                                            (5443, 2386, 2421),
                                                            (5444, 2415, 2421),
                                                            (5445, 2416, 2421),
                                                            (5446, 2273, 2422),
                                                            (5447, 2386, 2422),
                                                            (5448, 2415, 2422),
                                                            (5449, 2416, 2422),
                                                            (5450, 2273, 2423),
                                                            (5451, 2386, 2423),
                                                            (5452, 2415, 2423),
                                                            (5453, 2416, 2423),
                                                            (5454, 2273, 2425),
                                                            (5455, 2386, 2425),
                                                            (5456, 2415, 2425),
                                                            (5457, 2424, 2425),
                                                            (5458, 2273, 2426),
                                                            (5459, 2386, 2426),
                                                            (5460, 2415, 2426),
                                                            (5461, 2424, 2426),
                                                            (5462, 2273, 2416),
                                                            (5463, 2386, 2416),
                                                            (5464, 2415, 2416),
                                                            (5465, 2273, 2424),
                                                            (5466, 2386, 2424),
                                                            (5467, 2415, 2424),
                                                            (5468, 2273, 2387),
                                                            (5469, 2386, 2387),
                                                            (5470, 2273, 2395),
                                                            (5471, 2386, 2395),
                                                            (5472, 2273, 2401),
                                                            (5473, 2386, 2401),
                                                            (5474, 2273, 2408),
                                                            (5475, 2386, 2408),
                                                            (5476, 2273, 2415),
                                                            (5477, 2386, 2415),
                                                            (5478, 2273, 2274),
                                                            (5479, 2273, 2275),
                                                            (5480, 2273, 2334),
                                                            (5481, 2273, 2349),
                                                            (5482, 2273, 2386);

-- --------------------------------------------------------

--
-- 表的结构 `role`
--

CREATE TABLE `role` (
                        `id` bigint(20) UNSIGNED NOT NULL COMMENT '主键ID',
                        `parent_id` bigint(20) UNSIGNED NOT NULL COMMENT '父id',
                        `name` varchar(64) NOT NULL COMMENT '名称',
                        `keyword` char(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT '标识',
                        `status` tinyint(1) NOT NULL DEFAULT '0' COMMENT '状态',
                        `description` varchar(128) NOT NULL COMMENT '描述',
                        `department_ids` tinytext COMMENT '自定义部门',
                        `data_scope` char(32) NOT NULL COMMENT '权限类型',
                        `created_at` bigint(20) UNSIGNED NOT NULL DEFAULT '0' COMMENT '创建时间',
                        `updated_at` bigint(20) UNSIGNED NOT NULL DEFAULT '0' COMMENT '修改时间',
                        `deleted_at` bigint(20) UNSIGNED NOT NULL DEFAULT '0' COMMENT '删除时间'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='角色信息';

--
-- 转存表中的数据 `role`
--

INSERT INTO `role` (`id`, `parent_id`, `name`, `keyword`, `status`, `description`, `department_ids`, `data_scope`, `created_at`, `updated_at`, `deleted_at`) VALUES
                                                                                                                                                                 (1, 0, '超级管理员', 'superAdmin', 1, '超级管理员  ', NULL, 'ALL', 1713706137, 1713706137, 0),
                                                                                                                                                                 (5, 1, '2', '3', 1, '4', NULL, 'ALL', 1719464519, 1719464562, 0);

-- --------------------------------------------------------

--
-- 表的结构 `role_closure`
--

CREATE TABLE `role_closure` (
                                `id` bigint(20) UNSIGNED NOT NULL COMMENT '主键ID',
                                `parent` bigint(20) UNSIGNED NOT NULL COMMENT '角色id',
                                `children` bigint(20) UNSIGNED NOT NULL COMMENT '角色id'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='角色层级信息';

--
-- 转存表中的数据 `role_closure`
--

INSERT INTO `role_closure` (`id`, `parent`, `children`) VALUES
    (5, 1, 5);

-- --------------------------------------------------------

--
-- 表的结构 `role_menu`
--

CREATE TABLE `role_menu` (
                             `id` bigint(20) UNSIGNED NOT NULL COMMENT '主键ID',
                             `role_id` bigint(20) UNSIGNED NOT NULL COMMENT '角色id',
                             `menu_id` bigint(20) UNSIGNED NOT NULL COMMENT '菜单id'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='角色菜单信息';

-- --------------------------------------------------------

--
-- 表的结构 `user`
--

CREATE TABLE `user` (
                        `id` bigint(20) UNSIGNED NOT NULL COMMENT '主键ID',
                        `department_id` bigint(20) UNSIGNED NOT NULL COMMENT '部门id',
                        `role_id` bigint(20) UNSIGNED NOT NULL COMMENT '角色id',
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
                        `logged_at` bigint(20) NOT NULL DEFAULT '0' COMMENT '登陆时间',
                        `created_at` bigint(20) NOT NULL DEFAULT '0' COMMENT '创建时间',
                        `updated_at` bigint(20) NOT NULL DEFAULT '0' COMMENT '修改时间'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='用户信息';

--
-- 转存表中的数据 `user`
--

INSERT INTO `user` (`id`, `department_id`, `role_id`, `name`, `nickname`, `gender`, `avatar`, `email`, `phone`, `password`, `status`, `setting`, `token`, `logged_at`, `created_at`, `updated_at`) VALUES
    (1, 1, 1, '超级管理员', '超级管理员', 'F', '2a0786fe9127b8116bc30ed2ce9581e2', '1280291001@qq.com', '18888888888', '$2a$10$9qRJe9KQo6sEcU8ipKg.e.dkl2E7Wy64SigYlgraTAn.1paHFq6W.', 1, '{\"theme\":\"light\",\"themeColor\":\"#165DFF\",\"skin\":\"default\",\"tabBar\":true,\"menuWidth\":200,\"layout\":\"default\",\"language\":\"zh_CN\",\"animation\":\"gp-fade\"}', 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJkZXBhcnRtZW50SWQiOjEsImRlcGFydG1lbnRLZXl3b3JkIjoiY29tcGFueSIsImV4cCI6MTcxOTQ4MDA2NiwiaWF0IjoxNzE5NDcyODY1LCJuYmYiOjE3MTk0NzI4NjUsInJvbGVJZCI6MSwicm9sZUtleXdvcmQiOiJzdXBlckFkbWluIiwidXNlcklkIjoxfQ.BTX_QTdvi6Dc_2ydgUjjYjVXicMvc3flm7O-AkB5-1M', 1719472865, 1713706137, 1719472865);

-- --------------------------------------------------------

--
-- 表的结构 `user_job`
--

CREATE TABLE `user_job` (
                            `id` bigint(20) UNSIGNED NOT NULL COMMENT '主键ID',
                            `job_id` bigint(20) UNSIGNED NOT NULL COMMENT '职位id',
                            `user_id` bigint(20) UNSIGNED NOT NULL COMMENT '用户id'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='用户职位信息';

--
-- 转存表中的数据 `user_job`
--

INSERT INTO `user_job` (`id`, `job_id`, `user_id`) VALUES
    (1, 1, 1);

-- --------------------------------------------------------

--
-- 表的结构 `user_role`
--

CREATE TABLE `user_role` (
                             `id` bigint(20) UNSIGNED NOT NULL COMMENT '主键ID',
                             `role_id` bigint(20) UNSIGNED NOT NULL COMMENT '角色id',
                             `user_id` bigint(20) UNSIGNED NOT NULL COMMENT '用户id'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='用户角色信息';

--
-- 转存表中的数据 `user_role`
--

INSERT INTO `user_role` (`id`, `role_id`, `user_id`) VALUES
    (1, 1, 1);

--
-- 转储表的索引
--

--
-- 表的索引 `casbin_rule`
--
ALTER TABLE `casbin_rule`
    ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `idx_casbin_rule` (`ptype`,`v0`,`v1`,`v2`,`v3`,`v4`,`v5`);

--
-- 表的索引 `department`
--
ALTER TABLE `department`
    ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `keyword` (`keyword`),
  ADD KEY `idx_department_created_at` (`created_at`),
  ADD KEY `idx_department_updated_at` (`updated_at`);

--
-- 表的索引 `department_closure`
--
ALTER TABLE `department_closure`
    ADD PRIMARY KEY (`id`),
  ADD KEY `parent` (`parent`),
  ADD KEY `children` (`children`);

--
-- 表的索引 `dictionary`
--
ALTER TABLE `dictionary`
    ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `keyword` (`keyword`,`deleted_at`),
  ADD KEY `idx_dictionary_created_at` (`created_at`),
  ADD KEY `idx_dictionary_updated_at` (`updated_at`),
  ADD KEY `idx_dictionary_deleted_at` (`deleted_at`);

--
-- 表的索引 `dictionary_value`
--
ALTER TABLE `dictionary_value`
    ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `value` (`dictionary_id`,`value`),
  ADD KEY `idx_dictionary_value_created_at` (`created_at`),
  ADD KEY `idx_dictionary_value_updated_at` (`updated_at`),
  ADD KEY `idx_dictionary_value_weight` (`weight`);

--
-- 表的索引 `gorm_init`
--
ALTER TABLE `gorm_init`
    ADD PRIMARY KEY (`id`);

--
-- 表的索引 `job`
--
ALTER TABLE `job`
    ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `keyword` (`keyword`,`deleted_at`),
  ADD KEY `idx_job_weight` (`weight`),
  ADD KEY `idx_job_updated_at` (`updated_at`),
  ADD KEY `idx_job_created_at` (`created_at`),
  ADD KEY `idx_job_deleted_at` (`deleted_at`);

--
-- 表的索引 `menu`
--
ALTER TABLE `menu`
    ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `keyword` (`keyword`),
  ADD UNIQUE KEY `path` (`path`),
  ADD UNIQUE KEY `api_method` (`api`,`method`),
  ADD KEY `idx_menu_created_at` (`created_at`),
  ADD KEY `idx_menu_updated_at` (`updated_at`),
  ADD KEY `idx_menu_weight` (`weight`);

--
-- 表的索引 `menu_closure`
--
ALTER TABLE `menu_closure`
    ADD PRIMARY KEY (`id`),
  ADD KEY `parent` (`parent`),
  ADD KEY `children` (`children`);

--
-- 表的索引 `role`
--
ALTER TABLE `role`
    ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `keyword` (`keyword`,`deleted_at`),
  ADD KEY `idx_role_created_at` (`created_at`),
  ADD KEY `idx_role_updated_at` (`updated_at`),
  ADD KEY `idx_role_deleted_at` (`deleted_at`);

--
-- 表的索引 `role_closure`
--
ALTER TABLE `role_closure`
    ADD PRIMARY KEY (`id`),
  ADD KEY `parent` (`parent`),
  ADD KEY `children` (`children`);

--
-- 表的索引 `role_menu`
--
ALTER TABLE `role_menu`
    ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `role_id_2` (`role_id`,`menu_id`),
  ADD KEY `role_id` (`role_id`),
  ADD KEY `menu_id` (`menu_id`);

--
-- 表的索引 `user`
--
ALTER TABLE `user`
    ADD PRIMARY KEY (`id`),
  ADD KEY `idx_user_updated_at` (`updated_at`),
  ADD KEY `idx_user_created_at` (`created_at`),
  ADD KEY `fk_user_role` (`role_id`),
  ADD KEY `fk_user_department` (`department_id`);

--
-- 表的索引 `user_job`
--
ALTER TABLE `user_job`
    ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `job_id` (`job_id`,`user_id`),
  ADD KEY `user_id` (`user_id`);

--
-- 表的索引 `user_role`
--
ALTER TABLE `user_role`
    ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `role_id` (`role_id`,`user_id`),
  ADD KEY `user_id` (`user_id`);

--
-- 在导出的表使用AUTO_INCREMENT
--

--
-- 使用表AUTO_INCREMENT `casbin_rule`
--
ALTER TABLE `casbin_rule`
    MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=591;

--
-- 使用表AUTO_INCREMENT `department`
--
ALTER TABLE `department`
    MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键ID', AUTO_INCREMENT=5;

--
-- 使用表AUTO_INCREMENT `department_closure`
--
ALTER TABLE `department_closure`
    MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键ID', AUTO_INCREMENT=10;

--
-- 使用表AUTO_INCREMENT `dictionary`
--
ALTER TABLE `dictionary`
    MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键ID', AUTO_INCREMENT=3;

--
-- 使用表AUTO_INCREMENT `dictionary_value`
--
ALTER TABLE `dictionary_value`
    MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键ID', AUTO_INCREMENT=2;

--
-- 使用表AUTO_INCREMENT `gorm_init`
--
ALTER TABLE `gorm_init`
    MODIFY `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=2;

--
-- 使用表AUTO_INCREMENT `job`
--
ALTER TABLE `job`
    MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键ID', AUTO_INCREMENT=13;

--
-- 使用表AUTO_INCREMENT `menu`
--
ALTER TABLE `menu`
    MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键ID', AUTO_INCREMENT=2427;

--
-- 使用表AUTO_INCREMENT `menu_closure`
--
ALTER TABLE `menu_closure`
    MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键ID', AUTO_INCREMENT=5483;

--
-- 使用表AUTO_INCREMENT `role`
--
ALTER TABLE `role`
    MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键ID', AUTO_INCREMENT=6;

--
-- 使用表AUTO_INCREMENT `role_closure`
--
ALTER TABLE `role_closure`
    MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键ID', AUTO_INCREMENT=6;

--
-- 使用表AUTO_INCREMENT `role_menu`
--
ALTER TABLE `role_menu`
    MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键ID', AUTO_INCREMENT=219;

--
-- 使用表AUTO_INCREMENT `user`
--
ALTER TABLE `user`
    MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键ID', AUTO_INCREMENT=3;

--
-- 使用表AUTO_INCREMENT `user_job`
--
ALTER TABLE `user_job`
    MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键ID', AUTO_INCREMENT=7;

--
-- 使用表AUTO_INCREMENT `user_role`
--
ALTER TABLE `user_role`
    MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键ID', AUTO_INCREMENT=4;

--
-- 限制导出的表
--

--
-- 限制表 `department_closure`
--
ALTER TABLE `department_closure`
    ADD CONSTRAINT `department_closure_ibfk_1` FOREIGN KEY (`children`) REFERENCES `department` (`id`) ON DELETE CASCADE,
  ADD CONSTRAINT `department_closure_ibfk_2` FOREIGN KEY (`parent`) REFERENCES `department` (`id`) ON DELETE CASCADE;

--
-- 限制表 `dictionary_value`
--
ALTER TABLE `dictionary_value`
    ADD CONSTRAINT `fk_dictionary_value_dict` FOREIGN KEY (`dictionary_id`) REFERENCES `dictionary` (`id`) ON DELETE CASCADE;

--
-- 限制表 `menu_closure`
--
ALTER TABLE `menu_closure`
    ADD CONSTRAINT `menu_closure_ibfk_1` FOREIGN KEY (`children`) REFERENCES `menu` (`id`) ON DELETE CASCADE,
  ADD CONSTRAINT `menu_closure_ibfk_2` FOREIGN KEY (`parent`) REFERENCES `menu` (`id`) ON DELETE CASCADE;

--
-- 限制表 `role_closure`
--
ALTER TABLE `role_closure`
    ADD CONSTRAINT `role_closure_ibfk_1` FOREIGN KEY (`children`) REFERENCES `role` (`id`) ON DELETE CASCADE,
  ADD CONSTRAINT `role_closure_ibfk_2` FOREIGN KEY (`parent`) REFERENCES `role` (`id`) ON DELETE CASCADE;

--
-- 限制表 `role_menu`
--
ALTER TABLE `role_menu`
    ADD CONSTRAINT `role_menu_ibfk_1` FOREIGN KEY (`menu_id`) REFERENCES `menu` (`id`) ON DELETE CASCADE,
  ADD CONSTRAINT `role_menu_ibfk_2` FOREIGN KEY (`role_id`) REFERENCES `role` (`id`) ON DELETE CASCADE;

--
-- 限制表 `user`
--
ALTER TABLE `user`
    ADD CONSTRAINT `fk_user_department` FOREIGN KEY (`department_id`) REFERENCES `department` (`id`),
  ADD CONSTRAINT `fk_user_role` FOREIGN KEY (`role_id`) REFERENCES `role` (`id`);

--
-- 限制表 `user_job`
--
ALTER TABLE `user_job`
    ADD CONSTRAINT `user_job_ibfk_1` FOREIGN KEY (`job_id`) REFERENCES `job` (`id`) ON DELETE CASCADE,
  ADD CONSTRAINT `user_job_ibfk_2` FOREIGN KEY (`user_id`) REFERENCES `user` (`id`) ON DELETE CASCADE;

--
-- 限制表 `user_role`
--
ALTER TABLE `user_role`
    ADD CONSTRAINT `user_role_ibfk_1` FOREIGN KEY (`role_id`) REFERENCES `role` (`id`) ON DELETE CASCADE,
  ADD CONSTRAINT `user_role_ibfk_2` FOREIGN KEY (`user_id`) REFERENCES `user` (`id`) ON DELETE CASCADE;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
