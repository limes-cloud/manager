-- phpMyAdmin SQL Dump
-- version 5.2.0
-- https://www.phpmyadmin.net/
--
-- 主机： localhost
-- 生成日期： 2024-08-22 16:27:08
-- 服务器版本： 5.7.40-log
-- PHP 版本： 7.4.33

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


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

--
-- 转存表中的数据 `casbin_rule`
--

INSERT INTO `casbin_rule` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES
                                                                                  (738, 'p', 'test', '/manager/api/v1/department', 'DELETE', '', '', ''),
                                                                                  (736, 'p', 'test', '/manager/api/v1/department', 'POST', '', '', ''),
                                                                                  (737, 'p', 'test', '/manager/api/v1/department', 'PUT', '', '', ''),
                                                                                  (719, 'p', 'test', '/manager/api/v1/dictionaries', 'GET', '', '', ''),
                                                                                  (722, 'p', 'test', '/manager/api/v1/dictionary', 'DELETE', '', '', ''),
                                                                                  (720, 'p', 'test', '/manager/api/v1/dictionary', 'POST', '', '', ''),
                                                                                  (721, 'p', 'test', '/manager/api/v1/dictionary', 'PUT', '', '', ''),
                                                                                  (727, 'p', 'test', '/manager/api/v1/dictionary_value', 'DELETE', '', '', ''),
                                                                                  (724, 'p', 'test', '/manager/api/v1/dictionary_value', 'POST', '', '', ''),
                                                                                  (725, 'p', 'test', '/manager/api/v1/dictionary_value', 'PUT', '', '', ''),
                                                                                  (726, 'p', 'test', '/manager/api/v1/dictionary_value/status', 'PUT', '', '', ''),
                                                                                  (723, 'p', 'test', '/manager/api/v1/dictionary_values', 'GET', '', '', ''),
                                                                                  (735, 'p', 'test', '/manager/api/v1/job', 'DELETE', '', '', ''),
                                                                                  (733, 'p', 'test', '/manager/api/v1/job', 'POST', '', '', ''),
                                                                                  (734, 'p', 'test', '/manager/api/v1/job', 'PUT', '', '', ''),
                                                                                  (732, 'p', 'test', '/manager/api/v1/jobs', 'GET', '', '', ''),
                                                                                  (731, 'p', 'test', '/manager/api/v1/menu', 'DELETE', '', '', ''),
                                                                                  (729, 'p', 'test', '/manager/api/v1/menu', 'POST', '', '', ''),
                                                                                  (730, 'p', 'test', '/manager/api/v1/menu', 'PUT', '', '', ''),
                                                                                  (728, 'p', 'test', '/manager/api/v1/menus', 'GET', '', '', ''),
                                                                                  (742, 'p', 'test', '/manager/api/v1/role', 'DELETE', '', '', ''),
                                                                                  (739, 'p', 'test', '/manager/api/v1/role', 'POST', '', '', ''),
                                                                                  (740, 'p', 'test', '/manager/api/v1/role', 'PUT', '', '', ''),
                                                                                  (744, 'p', 'test', '/manager/api/v1/role/menu', 'POST', '', '', ''),
                                                                                  (743, 'p', 'test', '/manager/api/v1/role/menu_ids', 'GET', '', '', ''),
                                                                                  (741, 'p', 'test', '/manager/api/v1/role/status', 'PUT', '', '', ''),
                                                                                  (748, 'p', 'test', '/manager/api/v1/user', 'DELETE', '', '', ''),
                                                                                  (746, 'p', 'test', '/manager/api/v1/user', 'POST', '', '', ''),
                                                                                  (747, 'p', 'test', '/manager/api/v1/user', 'PUT', '', '', ''),
                                                                                  (750, 'p', 'test', '/manager/api/v1/user/password/reset', 'POST', '', '', ''),
                                                                                  (749, 'p', 'test', '/manager/api/v1/user/status', 'PUT', '', '', ''),
                                                                                  (745, 'p', 'test', '/manager/api/v1/users', 'GET', '', '', '');

-- --------------------------------------------------------

--
-- 表的结构 `department`
--

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
                                                                                                               (1, 0, 'company', '贵州青橙科技有限公司', '开放合作，拥抱未来', 1713706137, 1713706137),
                                                                                                               (5, 1, '123', '12', '12', 1722220242, 1722220242);

-- --------------------------------------------------------

--
-- 表的结构 `department_closure`
--

CREATE TABLE `department_closure` (
                                      `id` bigint(20) UNSIGNED NOT NULL COMMENT '主键ID',
                                      `parent` bigint(20) UNSIGNED NOT NULL COMMENT '部门id',
                                      `children` bigint(20) UNSIGNED NOT NULL COMMENT '部门id'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='部门层级信息';

--
-- 转存表中的数据 `department_closure`
--

INSERT INTO `department_closure` (`id`, `parent`, `children`) VALUES
    (10, 1, 5);

-- --------------------------------------------------------

--
-- 表的结构 `dictionary`
--

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
                                                                                                                                                                                                                                     (4010, 4009, '首页面板', 'M', 'Dashboard', 'dashboard', NULL, NULL, '/dashboard', NULL, '/dashboard/index', NULL, 0, 0, 1, 1, 1, 1720778192, 1720778192),
                                                                                                                                                                                                                                     (4011, 4009, '管理中心', 'M', 'SystemPlatformManager', 'apps', NULL, NULL, '/manager', NULL, NULL, NULL, 0, 0, NULL, NULL, NULL, 1720778192, 1720778192),
                                                                                                                                                                                                                                     (4012, 4011, '基础接口', 'G', 'ManagerBaseApi', 'apps', NULL, NULL, NULL, NULL, NULL, NULL, 0, 1, NULL, NULL, NULL, 1720778192, 1720778192),
                                                                                                                                                                                                                                     (4013, 4012, '获取用户可见部门树', 'BA', NULL, NULL, '/manager/api/v1/departments', 'GET', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL, 1720778192, 1720778192),
                                                                                                                                                                                                                                     (4014, 4012, '获取用户可见角色树', 'BA', NULL, NULL, '/manager/api/v1/roles', 'GET', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL, 1720778192, 1720778192),
                                                                                                                                                                                                                                     (4015, 4012, '获取个人用户信息', 'BA', NULL, NULL, '/manager/api/v1/user/current', 'GET', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL, 1720778192, 1720778192),
                                                                                                                                                                                                                                     (4016, 4012, '更新个人用户信息', 'BA', NULL, NULL, '/manager/api/v1/user/current/info', 'PUT', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL, 1720778192, 1720778192),
                                                                                                                                                                                                                                     (4017, 4012, '更新个人用户密码', 'BA', NULL, NULL, '/manager/api/v1/user/current/password', 'PUT', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL, 1720778192, 1720778192),
                                                                                                                                                                                                                                     (4018, 4012, '保存个人用户设置', 'BA', NULL, NULL, '/manager/api/v1/user/current/setting', 'PUT', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL, 1720778192, 1720778192),
                                                                                                                                                                                                                                     (4019, 4012, '发送用户验证吗', 'BA', NULL, NULL, '/manager/api/v1/user/current/captcha', 'POST', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL, 1720778192, 1720778192),
                                                                                                                                                                                                                                     (4020, 4012, '获取用户当前角色菜单', 'BA', NULL, NULL, '/manager/api/v1/menus/by/cur_role', 'GET', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL, 1720778192, 1720778192),
                                                                                                                                                                                                                                     (4021, 4012, '退出登录', 'BA', NULL, NULL, '/manager/api/v1/user/logout', 'POST', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL, 1720778192, 1720778192),
                                                                                                                                                                                                                                     (4022, 4012, '刷新token', 'BA', NULL, NULL, '/manager/api/v1/user/token/refresh', 'POST', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL, 1720778192, 1720778192),
                                                                                                                                                                                                                                     (4023, 4012, '用户登录', 'BA', NULL, NULL, '/manager/api/v1/user/login', 'POST', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL, 1720778192, 1720778192),
                                                                                                                                                                                                                                     (4024, 4012, '获取登录验证码', 'BA', NULL, NULL, '/manager/api/v1/user/login/captcha', 'POST', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL, 1720778192, 1720778192),
                                                                                                                                                                                                                                     (4025, 4012, '获取系统基础设置', 'BA', NULL, NULL, '/manager/api/v1/system/setting', 'GET', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL, 1720778192, 1720778192),
                                                                                                                                                                                                                                     (4026, 4012, '接口鉴权', 'BA', NULL, NULL, '/manager/api/v1/auth', 'POST', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL, 1720778192, 1720778192),
                                                                                                                                                                                                                                     (4027, 4012, '切换用户角色', 'BA', NULL, NULL, '/manager/api/v1/user/current/role', 'POST', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL, 1720778192, 1720778192),
                                                                                                                                                                                                                                     (4028, 4012, '分块上传文件', 'BA', NULL, NULL, '/resource/api/v1/upload', 'POST', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL, 1720778192, 1720778192),
                                                                                                                                                                                                                                     (4029, 4012, '预上传文件', 'BA', NULL, NULL, '/resource/api/v1/prepare_upload', 'POST', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL, 1720778192, 1720778192),
                                                                                                                                                                                                                                     (4030, 4012, '获取字段类型', 'BA', NULL, NULL, '/usercenter/api/v1/field/types', 'GET', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL, 1720778192, 1720778192),
                                                                                                                                                                                                                                     (4031, 4012, '查询资源权限', 'BA', NULL, NULL, '/manager/api/v1/resource', 'GET', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL, 1720778192, 1720778192),
                                                                                                                                                                                                                                     (4032, 4012, '获取渠道类型', 'BA', NULL, NULL, '/usercenter/api/v1/channel/types', 'GET', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL, 1720778192, 1720778192),
                                                                                                                                                                                                                                     (4033, 4011, '字典管理', 'M', 'ManagerDictionary', 'storage', NULL, NULL, '/manager/dictionary', NULL, '/manager/dictionary/index', NULL, 0, NULL, 1, NULL, 1, 1720778192, 1720778192),
                                                                                                                                                                                                                                     (4034, 4033, '查询字典', 'A', NULL, NULL, '/manager/api/v1/dictionaries', 'GET', NULL, 'manager:dictionary:query', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1720778192, 1720778192),
                                                                                                                                                                                                                                     (4035, 4033, '新增字典', 'A', NULL, NULL, '/manager/api/v1/dictionary', 'POST', NULL, 'manager:dictionary:add', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1720778192, 1720778192),
                                                                                                                                                                                                                                     (4036, 4033, '修改字典', 'A', NULL, NULL, '/manager/api/v1/dictionary', 'PUT', NULL, 'manager:dictionary:update', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1720778192, 1720778192),
                                                                                                                                                                                                                                     (4037, 4033, '删除字典', 'A', NULL, NULL, '/manager/api/v1/dictionary', 'DELETE', NULL, 'manager:dictionary:delete', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1720778192, 1720778192),
                                                                                                                                                                                                                                     (4038, 4033, '获取字典值', 'A', NULL, NULL, '/manager/api/v1/dictionary_values', 'GET', NULL, 'manager:dictionary:value:query', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1720778192, 1720778192),
                                                                                                                                                                                                                                     (4039, 4033, '新增字典值', 'A', NULL, NULL, '/manager/api/v1/dictionary_value', 'POST', NULL, 'manager:dictionary:value:add', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1720778192, 1720778192),
                                                                                                                                                                                                                                     (4040, 4033, '修改字典值', 'A', NULL, NULL, '/manager/api/v1/dictionary_value', 'PUT', NULL, 'manager:dictionary:value:update', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1720778192, 1720778192),
                                                                                                                                                                                                                                     (4041, 4033, '更新字典值目录状态', 'A', NULL, NULL, '/manager/api/v1/dictionary_value/status', 'PUT', NULL, 'manager:dictionary:value:status', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1720778192, 1720778192),
                                                                                                                                                                                                                                     (4042, 4033, '删除字典值', 'A', NULL, NULL, '/manager/api/v1/dictionary_value', 'DELETE', NULL, 'manager:dictionary:value:delete', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1720778192, 1720778192),
                                                                                                                                                                                                                                     (4043, 4011, '菜单管理', 'M', 'ManagerMenu', 'menu', NULL, NULL, '/manager/menu', NULL, '/manager/menu/index', NULL, 0, 0, 1, 1, 1, 1720778192, 1720778192),
                                                                                                                                                                                                                                     (4044, 4043, '查询菜单', 'A', NULL, NULL, '/manager/api/v1/menus', 'GET', NULL, 'manager:menu:query', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1720778192, 1720778192),
                                                                                                                                                                                                                                     (4045, 4043, '新增菜单', 'A', NULL, NULL, '/manager/api/v1/menu', 'POST', NULL, 'manager:menu:add', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1720778192, 1720778192),
                                                                                                                                                                                                                                     (4046, 4043, '修改菜单', 'A', NULL, NULL, '/manager/api/v1/menu', 'PUT', NULL, 'manager:menu:update', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1720778192, 1720778192),
                                                                                                                                                                                                                                     (4047, 4043, '删除菜单', 'A', NULL, NULL, '/manager/api/v1/menu', 'DELETE', NULL, 'manager:menu:delete', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1720778192, 1720778192),
                                                                                                                                                                                                                                     (4048, 4011, '职位管理', 'M', 'ManagerJob', 'tag', NULL, NULL, '/manager/job', NULL, '/manager/job/index', NULL, 0, NULL, 1, NULL, 1, 1720778192, 1720778192),
                                                                                                                                                                                                                                     (4049, 4048, '查询职位', 'A', NULL, NULL, '/manager/api/v1/jobs', 'GET', NULL, 'manager:job:query', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1720778192, 1720778192),
                                                                                                                                                                                                                                     (4050, 4048, '新增职位', 'A', NULL, NULL, '/manager/api/v1/job', 'POST', NULL, 'manager:job:add', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1720778192, 1720778192),
                                                                                                                                                                                                                                     (4051, 4048, '修改职位', 'A', NULL, NULL, '/manager/api/v1/job', 'PUT', NULL, 'manager:job:update', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1720778192, 1720778192),
                                                                                                                                                                                                                                     (4052, 4048, '删除职位', 'A', NULL, NULL, '/manager/api/v1/job', 'DELETE', NULL, 'manager:job:delete', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1720778192, 1720778192),
                                                                                                                                                                                                                                     (4053, 4011, '部门管理', 'M', 'ManagerDepartment', 'user-group', NULL, NULL, '/manager/department', NULL, '/manager/department/index', NULL, 0, 0, 1, NULL, 1, 1720778192, 1720778192),
                                                                                                                                                                                                                                     (4054, 4053, '新增部门', 'A', NULL, NULL, '/manager/api/v1/department', 'POST', NULL, 'manager:department:add', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1720778192, 1720778192),
                                                                                                                                                                                                                                     (4055, 4053, '修改部门', 'A', NULL, NULL, '/manager/api/v1/department', 'PUT', NULL, 'manager:department:update', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1720778192, 1720778192),
                                                                                                                                                                                                                                     (4056, 4053, '删除部门', 'A', NULL, NULL, '/manager/api/v1/department', 'DELETE', NULL, 'manager:department:delete', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1720778192, 1720778192),
                                                                                                                                                                                                                                     (4057, 4011, '角色管理', 'M', 'ManagerRole', 'safe', NULL, NULL, '/manager/role', NULL, '/manager/role/index', NULL, 0, NULL, 0, NULL, 1, 1720778192, 1720778192),
                                                                                                                                                                                                                                     (4058, 4057, '新增角色', 'A', NULL, NULL, '/manager/api/v1/role', 'POST', NULL, 'manager:role:add', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1720778192, 1720778192),
                                                                                                                                                                                                                                     (4059, 4057, '修改角色', 'A', NULL, NULL, '/manager/api/v1/role', 'PUT', NULL, 'manager:role:update', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1720778192, 1720778192),
                                                                                                                                                                                                                                     (4060, 4057, '修改角色状态', 'A', NULL, NULL, '/manager/api/v1/role/status', 'PUT', NULL, 'manager:role:update:status', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1720778192, 1720778192),
                                                                                                                                                                                                                                     (4061, 4057, '删除角色', 'A', NULL, NULL, '/manager/api/v1/role', 'DELETE', NULL, 'manager:role:delete', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1720778192, 1720778192),
                                                                                                                                                                                                                                     (4062, 4057, '角色菜单管理', 'G', NULL, NULL, NULL, NULL, NULL, 'manager:role:menu', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1720778192, 1720778192),
                                                                                                                                                                                                                                     (4063, 4062, '查询角色菜单', 'A', NULL, NULL, '/manager/api/v1/role/menu_ids', 'GET', NULL, 'manager:role:menu:query', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1720778192, 1720778192),
                                                                                                                                                                                                                                     (4064, 4062, '修改角色菜单', 'A', NULL, NULL, '/manager/api/v1/role/menu', 'POST', NULL, 'manager:role:menu:update', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1720778192, 1720778192),
                                                                                                                                                                                                                                     (4065, 4011, '用户管理', 'M', 'ManagerUser', 'user', NULL, NULL, '/manager/user', NULL, '/manager/user/index', NULL, 0, NULL, 1, NULL, 1, 1720778192, 1720778192),
                                                                                                                                                                                                                                     (4066, 4065, '查询用户列表', 'A', NULL, NULL, '/manager/api/v1/users', 'GET', NULL, 'manager:user:query', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1720778192, 1720778192),
                                                                                                                                                                                                                                     (4067, 4065, '新增用户', 'A', NULL, NULL, '/manager/api/v1/user', 'POST', NULL, 'manager:user:add', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1720778192, 1720778192),
                                                                                                                                                                                                                                     (4068, 4065, '修改用户', 'A', NULL, NULL, '/manager/api/v1/user', 'PUT', NULL, 'manager:user:update', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1720778192, 1720778192),
                                                                                                                                                                                                                                     (4069, 4065, '删除用户', 'A', NULL, NULL, '/manager/api/v1/user', 'DELETE', NULL, 'manager:user:delete', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1720778192, 1720778192),
                                                                                                                                                                                                                                     (4070, 4065, '修改用户状态', 'A', NULL, NULL, '/manager/api/v1/user/status', 'PUT', NULL, 'manager:user:status', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1720778192, 1720778192),
                                                                                                                                                                                                                                     (4071, 4065, '重置账号密码', 'A', NULL, NULL, '/manager/api/v1/user/password/reset', 'POST', NULL, 'manager:user:reset:password', '', NULL, 0, NULL, NULL, NULL, NULL, 1720778192, 1720778192),
                                                                                                                                                                                                                                     (4072, 4009, '资源中心', 'M', 'SystemPlatformResource', 'file', NULL, NULL, '/resource', NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL, 1720778192, 1720778192),
                                                                                                                                                                                                                                     (4073, 4072, '文件管理', 'M', 'ResourceFile', 'file', NULL, NULL, '/resource/file', NULL, '/resource/file/index', NULL, 0, 0, 1, NULL, 1, 1720778192, 1720778192),
                                                                                                                                                                                                                                     (4074, 4073, '目录管理', 'G', NULL, NULL, NULL, NULL, NULL, 'resource:directory:group', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1720778192, 1720778192),
                                                                                                                                                                                                                                     (4075, 4074, '查看目录', 'A', NULL, NULL, '/resource/api/v1/directories', 'GET', NULL, 'resource:directory:query', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1720778192, 1720778192),
                                                                                                                                                                                                                                     (4076, 4074, '新增目录', 'A', NULL, NULL, '/resource/api/v1/directory', 'POST', NULL, 'resource:directory:add', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1720778192, 1720778192),
                                                                                                                                                                                                                                     (4077, 4074, '修改目录', 'A', NULL, NULL, '/resource/api/v1/directory', 'PUT', NULL, 'resource:directory:update', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1720778192, 1720778192),
                                                                                                                                                                                                                                     (4078, 4074, '删除目录', 'A', NULL, NULL, '/resource/api/v1/directory', 'DELETE', NULL, 'resource:directory:delete', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1720778192, 1720778192),
                                                                                                                                                                                                                                     (4079, 4073, '文件管理', 'G', NULL, NULL, NULL, NULL, NULL, 'resource:file:group', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1720778192, 1720778192),
                                                                                                                                                                                                                                     (4080, 4079, '查看文件', 'A', NULL, NULL, '/resource/api/v1/files', 'GET', NULL, 'resource:file:query', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1720778192, 1720778192),
                                                                                                                                                                                                                                     (4081, 4079, '修改文件', 'A', NULL, NULL, '/resource/api/v1/file', 'PUT', NULL, 'resource:file:update', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1720778192, 1720778192),
                                                                                                                                                                                                                                     (4082, 4079, '删除文件', 'A', NULL, NULL, '/resource/api/v1/file', 'DELETE', NULL, 'resource:file:delete', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1720778192, 1720778192),
                                                                                                                                                                                                                                     (4083, 4072, '导出管理', 'M', 'ResourceExport', 'export', NULL, NULL, '/resource/export', NULL, '/resource/export/index', NULL, 0, 0, 1, NULL, 1, 1720778192, 1720778192),
                                                                                                                                                                                                                                     (4084, 4083, '查看导出', 'A', NULL, NULL, '/resource/api/v1/exports', 'GET', NULL, 'resource:export:query', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1720778192, 1720778192),
                                                                                                                                                                                                                                     (4085, 4083, '新增导出', 'A', NULL, NULL, '/resource/api/v1/export', 'POST', NULL, 'resource:export:file', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1720778192, 1720778192),
                                                                                                                                                                                                                                     (4086, 4083, '删除导出', 'A', NULL, NULL, '/resource/api/v1/export', 'DELETE', NULL, 'resource:export:delete', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1720778192, 1720778192),
                                                                                                                                                                                                                                     (4087, 4009, '用户中心', 'M', 'SystemPlatformUserCenter', 'user', NULL, NULL, '/usercenter', NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL, 1720778192, 1720778192),
                                                                                                                                                                                                                                     (4088, 4087, '授权渠道', 'M', 'UserCenterChannel', 'mind-mapping', NULL, NULL, '/usercenter/channel', NULL, '/usercenter/channel/index', NULL, 0, 0, 1, NULL, 1, 1720778192, 1720778192),
                                                                                                                                                                                                                                     (4089, 4088, '查看渠道', 'A', NULL, NULL, '/usercenter/api/v1/channels', 'GET', NULL, 'uc:channel:query', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1720778192, 1720778192),
                                                                                                                                                                                                                                     (4090, 4088, '新增渠道', 'A', NULL, NULL, '/usercenter/api/v1/channel', 'POST', NULL, 'uc:channel:add', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1720778192, 1720778192),
                                                                                                                                                                                                                                     (4091, 4088, '修改渠道', 'A', NULL, NULL, '/usercenter/api/v1/channel', 'PUT', NULL, 'uc:channel:update', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1720778192, 1720778192),
                                                                                                                                                                                                                                     (4092, 4088, '修改渠道状态', 'A', NULL, NULL, '/usercenter/api/v1/channel/status', 'PUT', NULL, 'uc:channel:update:status', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1720778192, 1720778192),
                                                                                                                                                                                                                                     (4093, 4088, '删除渠道', 'A', NULL, NULL, '/usercenter/api/v1/channel', 'DELETE', NULL, 'uc:channel:delete', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1720778192, 1720778192),
                                                                                                                                                                                                                                     (4094, 4087, '信息字段', 'M', 'UserCenterField', 'list', NULL, NULL, '/usercenter/field', NULL, '/usercenter/field/index', NULL, 0, 0, 1, 0, 1, 1720778192, 1720778192),
                                                                                                                                                                                                                                     (4095, 4094, '查看字段列表', 'A', NULL, NULL, '/usercenter/api/v1/fields', 'GET', NULL, 'uc:field:query', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1720778192, 1720778192),
                                                                                                                                                                                                                                     (4096, 4094, '新增字段', 'A', NULL, NULL, '/usercenter/api/v1/field', 'POST', NULL, 'uc:field:add', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1720778192, 1720778192),
                                                                                                                                                                                                                                     (4097, 4094, '修改字段', 'A', NULL, NULL, '/usercenter/api/v1/field', 'PUT', NULL, 'uc:field:update', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1720778192, 1720778192),
                                                                                                                                                                                                                                     (4098, 4094, '修改字段状态', 'A', NULL, NULL, '/usercenter/api/v1/field/status', 'PUT', NULL, 'uc:field:update:status', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1720778192, 1720778192),
                                                                                                                                                                                                                                     (4099, 4094, '删除字段', 'A', NULL, NULL, '/usercenter/api/v1/field', 'DELETE', NULL, 'uc:field:delete', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1720778192, 1720778192),
                                                                                                                                                                                                                                     (4100, 4087, '应用管理', 'M', 'UserCenterApp', 'apps', NULL, NULL, '/usercenter/app', NULL, '/usercenter/app/index', NULL, 0, 0, 1, NULL, 1, 1720778192, 1720778192),
                                                                                                                                                                                                                                     (4101, 4100, '查看应用', 'A', NULL, NULL, '/usercenter/api/v1/apps', 'GET', NULL, 'uc:app:query', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1720778192, 1720778192),
                                                                                                                                                                                                                                     (4102, 4100, '设置应用资源权限', 'A', NULL, NULL, '/manager/api/v1/resource/uc_app', 'PUT', NULL, 'uc:app:resource:permission', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1720778192, 1720778192),
                                                                                                                                                                                                                                     (4103, 4100, '新增应用', 'A', NULL, NULL, '/usercenter/api/v1/app', 'POST', NULL, 'uc:app:add', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1720778192, 1720778192),
                                                                                                                                                                                                                                     (4104, 4100, '修改应用', 'A', NULL, NULL, '/usercenter/api/v1/app', 'PUT', NULL, 'uc:app:update', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1720778192, 1720778192),
                                                                                                                                                                                                                                     (4105, 4100, '修改应用状态', 'A', NULL, NULL, '/usercenter/api/v1/app/status', 'PUT', NULL, 'uc:app:update:status', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1720778192, 1720778192),
                                                                                                                                                                                                                                     (4106, 4100, '删除应用', 'A', NULL, NULL, '/usercenter/api/v1/app', 'DELETE', NULL, 'uc:app:delete', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1720778192, 1720778192),
                                                                                                                                                                                                                                     (4107, 4087, '反馈管理', 'M', 'UserCenterFeedback', 'question-circle', NULL, NULL, '/usercenter/feedback', NULL, '/usercenter/feedback/index', NULL, 0, 0, 1, NULL, 1, 1720778192, 1720778192),
                                                                                                                                                                                                                                     (4108, 4107, '查看反馈分类', 'A', NULL, NULL, '/usercenter/api/v1/feedback_categories', 'GET', NULL, 'uc:feedback:category:query', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1720778192, 1720778192),
                                                                                                                                                                                                                                     (4109, 4107, '新增反馈渠道', 'A', NULL, NULL, '/usercenter/api/v1/feedback_category', 'POST', NULL, 'uc:feedback:category:add', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1720778192, 1720778192),
                                                                                                                                                                                                                                     (4110, 4107, '修改反馈渠道', 'A', NULL, NULL, '/usercenter/api/v1/feedback_category', 'PUT', NULL, 'uc:feedback:category:update', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1720778192, 1720778192),
                                                                                                                                                                                                                                     (4111, 4107, '删除反馈渠道', 'A', NULL, NULL, '/usercenter/api/v1/feedback_category', 'DELETE', NULL, 'uc:feedback:category:delete', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1720778192, 1720778192),
                                                                                                                                                                                                                                     (4112, 4107, '查看反馈', 'A', NULL, NULL, '/usercenter/api/v1/feedbacks', 'GET', NULL, 'uc:feedback:query', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1720778192, 1720778192),
                                                                                                                                                                                                                                     (4113, 4107, '新增反馈', 'A', NULL, NULL, '/usercenter/api/v1/feedback', 'POST', NULL, 'uc:feedback:add', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1720778192, 1720778192),
                                                                                                                                                                                                                                     (4114, 4107, '修改反馈', 'A', NULL, NULL, '/usercenter/api/v1/feedback', 'PUT', NULL, 'uc:feedback:update', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1720778192, 1720778192),
                                                                                                                                                                                                                                     (4115, 4107, '删除反馈', 'A', NULL, NULL, '/usercenter/api/v1/feedback', 'DELETE', NULL, 'uc:feedback:delete', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1720778192, 1720778192),
                                                                                                                                                                                                                                     (4116, 4087, '用户管理', 'M', 'UserCenterUser', 'user', NULL, NULL, '/usercenter/user', NULL, '/usercenter/user/index', NULL, 0, 0, 1, NULL, 1, 1720778192, 1720778192),
                                                                                                                                                                                                                                     (4117, 4116, '查看用户', 'A', NULL, NULL, '/usercenter/api/v1/users', 'GET', NULL, 'uc:user:query', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1720778192, 1720778192),
                                                                                                                                                                                                                                     (4118, 4116, '新增用户', 'A', NULL, NULL, '/usercenter/api/v1/user', 'POST', NULL, 'uc:user:add', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1720778192, 1720778192),
                                                                                                                                                                                                                                     (4119, 4116, '导入用户', 'A', NULL, NULL, '/usercenter/api/v1/users', 'POST', NULL, 'uc:user:import', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1720778192, 1720778192),
                                                                                                                                                                                                                                     (4120, 4116, '修改用户', 'A', NULL, NULL, '/usercenter/api/v1/user', 'PUT', NULL, 'uc:user:update', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1720778192, 1720778192),
                                                                                                                                                                                                                                     (4121, 4116, '修改用户状态', 'A', NULL, NULL, '/usercenter/api/v1/user/status', 'PUT', NULL, 'uc:user:update:status', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1720778192, 1720778192),
                                                                                                                                                                                                                                     (4122, 4116, '删除用户', 'A', NULL, NULL, '/usercenter/api/v1/user', 'DELETE', NULL, 'uc:user:delete', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1720778192, 1720778192),
                                                                                                                                                                                                                                     (4123, 4116, '查看用户详细信息', 'G', NULL, NULL, NULL, NULL, NULL, 'uc:user:more', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1720778192, 1720778192),
                                                                                                                                                                                                                                     (4124, 4123, '获取用户应用信息', 'A', NULL, NULL, '/usercenter/api/v1/auths', 'GET', NULL, 'uc:auth:query', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1720778192, 1720778192),
                                                                                                                                                                                                                                     (4125, 4123, '创建用户应用信息', 'A', NULL, NULL, '/usercenter/api/v1/auth', 'POST', NULL, 'uc:auth:add', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1720778192, 1720778192),
                                                                                                                                                                                                                                     (4126, 4123, '修改用户应用状态信息', 'A', NULL, NULL, '/usercenter/api/v1/auth/status', 'PUT', NULL, 'uc:auth:update:status', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1720778192, 1720778192),
                                                                                                                                                                                                                                     (4127, 4123, '删除用户应用信息', 'A', NULL, NULL, '/usercenter/api/v1/auth', 'DELETE', NULL, 'uc:auth:delete', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1720778192, 1720778192),
                                                                                                                                                                                                                                     (4128, 4123, '获取用户渠道信息', 'A', NULL, NULL, '/usercenter/api/v1/oauths', 'GET', NULL, 'uc:oauth:query', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1720778192, 1720778192),
                                                                                                                                                                                                                                     (4129, 4123, '删除用户渠道信息', 'A', NULL, NULL, '/usercenter/api/v1/oauth', 'DELETE', NULL, 'uc:oauth:delete', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1720778192, 1720778192),
                                                                                                                                                                                                                                     (4130, 4123, '获取用户扩展信息', 'A', NULL, NULL, '/usercenter/api/v1/userinfos', 'GET', NULL, 'uc:userinfo:query', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1720778192, 1720778192),
                                                                                                                                                                                                                                     (4131, 4123, '创建用户扩展信息', 'A', NULL, NULL, '/usercenter/api/v1/userinfo', 'POST', NULL, 'uc:userinfo:add', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1720778192, 1720778192),
                                                                                                                                                                                                                                     (4132, 4123, '修改用户扩展信息', 'A', NULL, NULL, '/usercenter/api/v1/userinfo', 'PUT', NULL, 'uc:userinfo:update', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1720778192, 1720778192),
                                                                                                                                                                                                                                     (4133, 4123, '删除用户扩展信息', 'A', NULL, NULL, '/usercenter/api/v1/userinfo', 'DELETE', NULL, 'uc:userinfo:delete', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1720778192, 1720778192),
                                                                                                                                                                                                                                     (4134, 4009, '配置中心', 'M', 'SystemPlatformConfigure', 'code-block', NULL, NULL, '/configure', NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL, 1720778192, 1720778192),
                                                                                                                                                                                                                                     (4135, 4134, '环境管理', 'M', 'ConfigureEnv', 'common', NULL, NULL, '/configure/env', NULL, '/configure/env/index', NULL, 0, NULL, 1, NULL, 1, 1720778192, 1720778192),
                                                                                                                                                                                                                                     (4136, 4135, '查看环境', 'A', NULL, NULL, '/configure/api/v1/envs', 'GET', NULL, 'configure:env:query', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1720778192, 1720778192),
                                                                                                                                                                                                                                     (4137, 4135, '设置环境资源权限', 'A', NULL, NULL, '/manager/api/v1/resource/cfg_env', 'PUT', NULL, 'configure:env:resource:permission', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1720778192, 1720778192),
                                                                                                                                                                                                                                     (4138, 4135, '新增环境', 'A', NULL, NULL, '/configure/api/v1/env', 'POST', NULL, 'configure:env:add', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1720778192, 1720778192),
                                                                                                                                                                                                                                     (4139, 4135, '修改环境', 'A', NULL, NULL, '/configure/api/v1/env', 'PUT', NULL, 'configure:env:update', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1720778192, 1720778192),
                                                                                                                                                                                                                                     (4140, 4135, '修改环境状态', 'A', NULL, NULL, '/configure/api/v1/env/status', 'PUT', NULL, 'configure:env:update:status', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1720778192, 1720778192),
                                                                                                                                                                                                                                     (4141, 4135, '删除环境', 'A', NULL, NULL, '/configure/api/v1/env', 'DELETE', NULL, 'configure:env:delete', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1720778192, 1720778192),
                                                                                                                                                                                                                                     (4142, 4135, '查看环境Token', 'A', NULL, NULL, '/configure/api/v1/env/token', 'GET', NULL, 'configure:env:token:query', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1720778192, 1720778192),
                                                                                                                                                                                                                                     (4143, 4135, '重置环境token', 'A', NULL, NULL, '/configure/api/v1/env/token', 'PUT', NULL, 'configure:env:token:reset', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1720778192, 1720778192),
                                                                                                                                                                                                                                     (4144, 4134, '服务管理', 'M', 'ConfigureServer', 'apps', NULL, NULL, '/configure/server', NULL, '/configure/server/index', NULL, 0, NULL, 1, NULL, 1, 1720778192, 1720778192),
                                                                                                                                                                                                                                     (4145, 4144, '查询服务', 'A', NULL, NULL, '/configure/api/v1/servers', 'GET', NULL, 'configure:server:query', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1720778192, 1720778192),
                                                                                                                                                                                                                                     (4146, 4144, '设置服务资源权限', 'A', NULL, NULL, '/manager/api/v1/resource/cfg_server', 'PUT', NULL, 'configure:server:resource:permission', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1720778192, 1720778192),
                                                                                                                                                                                                                                     (4147, 4144, '新增服务', 'A', NULL, NULL, '/configure/api/v1/server', 'POST', NULL, 'configure:server:add', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1720778192, 1720778192),
                                                                                                                                                                                                                                     (4148, 4144, '修改服务', 'A', NULL, NULL, '/configure/api/v1/server', 'PUT', NULL, 'configure:server:update', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1720778192, 1720778192),
                                                                                                                                                                                                                                     (4149, 4144, '修改服务状态', 'A', NULL, NULL, '/configure/api/v1/server/status', 'PUT', NULL, 'configure:server:update:status', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1720778192, 1720778192),
                                                                                                                                                                                                                                     (4150, 4144, '删除服务', 'A', NULL, NULL, '/configure/api/v1/server', 'DELETE', NULL, 'configure:server:delete', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1720778192, 1720778192),
                                                                                                                                                                                                                                     (4151, 4134, '业务变量', 'M', 'ConfigureBusiness', 'code', NULL, NULL, '/configure/business', NULL, '/configure/business/index', NULL, 0, NULL, 1, NULL, 1, 1720778192, 1720778192),
                                                                                                                                                                                                                                     (4152, 4151, '查看业务变量', 'A', NULL, NULL, '/configure/api/v1/businesses', 'GET', NULL, 'configure:business:query', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1720778192, 1720778192),
                                                                                                                                                                                                                                     (4153, 4151, '新增业务变量', 'A', NULL, NULL, '/configure/api/v1/business', 'POST', NULL, 'configure:business:add', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1720778192, 1720778192),
                                                                                                                                                                                                                                     (4154, 4151, '修改业务变量', 'A', NULL, NULL, '/configure/api/v1/business', 'PUT', NULL, 'configure:business:update', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1720778192, 1720778192),
                                                                                                                                                                                                                                     (4155, 4151, '删除业务变量', 'A', NULL, NULL, '/configure/api/v1/business', 'DELETE', NULL, 'configure:business:delete', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1720778192, 1720778192),
                                                                                                                                                                                                                                     (4156, 4151, '查看业务变量值', 'A', NULL, NULL, '/configure/business/values', 'get', NULL, 'configure:business:value:query', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1720778192, 1720778192),
                                                                                                                                                                                                                                     (4157, 4151, '设置业务变量值', 'A', NULL, NULL, '/configure/business/values', 'PUT', NULL, 'configure:business:value:update', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1720778192, 1720778192),
                                                                                                                                                                                                                                     (4158, 4134, '资源变量', 'M', 'ConfigureResource', 'unordered-list', NULL, NULL, '/configure/resource', NULL, '/configure/resource/index', NULL, 0, NULL, 1, NULL, 1, 1720778192, 1720778192),
                                                                                                                                                                                                                                     (4159, 4158, '查看资源', 'A', NULL, NULL, '/configure/api/v1/resources', 'GET', NULL, 'configure:resource:query', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1720778192, 1720778192),
                                                                                                                                                                                                                                     (4160, 4158, '新增资源', 'A', NULL, NULL, '/configure/api/v1/resource', 'POST', NULL, 'configure:resource:add', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1720778192, 1720778192),
                                                                                                                                                                                                                                     (4161, 4158, '修改资源', 'A', NULL, NULL, '/configure/api/v1/resource', 'PUT', NULL, 'configure:resource:update', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1720778192, 1720778192),
                                                                                                                                                                                                                                     (4162, 4158, '删除资源', 'A', NULL, NULL, '/configure/api/v1/resource', 'DELETE', NULL, 'configure:resource:delete', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1720778192, 1720778192),
                                                                                                                                                                                                                                     (4163, 4158, '查看业务变量值', 'A', NULL, NULL, '/configure/resource/values', 'get', NULL, 'configure:resource:value:query', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1720778192, 1720778192),
                                                                                                                                                                                                                                     (4164, 4158, '设置业务变量值', 'A', NULL, NULL, '/configure/resource/values', 'PUT', NULL, 'configure:resource:value:update', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1720778192, 1720778192),
                                                                                                                                                                                                                                     (4165, 4134, '配置模板', 'M', 'ConfgureTemplate', 'file', NULL, NULL, '/configure/template', NULL, '/configure/template/index', NULL, 0, NULL, 1, NULL, 1, 1720778192, 1720778192),
                                                                                                                                                                                                                                     (4166, 4165, '模板管理', 'G', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL, 1720778192, 1720778192),
                                                                                                                                                                                                                                     (4167, 4166, '查看模板历史版本', 'A', NULL, NULL, '/configure/api/v1/templates', 'GET', NULL, 'configure:template:history', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1720778192, 1720778192),
                                                                                                                                                                                                                                     (4168, 4166, '查看指定模板详细数据', 'A', NULL, NULL, '/configure/api/v1/template', 'GET', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL, 1720778192, 1720778192),
                                                                                                                                                                                                                                     (4169, 4166, '查看当前正在使用的模板', 'A', NULL, NULL, '/configure/api/v1/template/current', 'GET', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL, 1720778192, 1720778192),
                                                                                                                                                                                                                                     (4170, 4166, '提交模板', 'A', NULL, NULL, '/configure/api/v1/template', 'POST', NULL, 'configure:template:add', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1720778192, 1720778192),
                                                                                                                                                                                                                                     (4171, 4166, '模板对比', 'A', NULL, NULL, '/configure/api/v1/template/compare', 'POST', NULL, 'configure:template:compare', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1720778192, 1720778192),
                                                                                                                                                                                                                                     (4172, 4166, '切换模板', 'A', NULL, NULL, '/configure/api/v1/template/switch', 'POST', NULL, 'configure:template:switch', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1720778192, 1720778192),
                                                                                                                                                                                                                                     (4173, 4166, '模板预览', 'A', NULL, NULL, '/configure/api/v1/template/preview', 'POST', NULL, 'configure:template:preview', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1720778192, 1720778192),
                                                                                                                                                                                                                                     (4174, 4165, '配置管理', 'G', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL, 1720778192, 1720778192),
                                                                                                                                                                                                                                     (4175, 4174, '配置对比', 'A', NULL, NULL, '/configure/api/v1/configure/compare', 'POST', NULL, 'configure:configure:compare', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1720778192, 1720778192),
                                                                                                                                                                                                                                     (4176, 4174, '同步配置', 'A', NULL, NULL, '/configure/api/v1/configure', 'PUT', NULL, 'configure:configure:sync', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1720778192, 1720778192),
                                                                                                                                                                                                                                     (4177, 4009, '定时任务', 'M', 'SystemPlatformCron', 'schedule', NULL, NULL, '/cron', NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL, 1720778192, 1720778192),
                                                                                                                                                                                                                                     (4178, 4177, '节点管理', 'M', 'Worker', 'common', NULL, NULL, '/cron/worker', NULL, '/cron/worker/index', NULL, 0, 0, 1, NULL, 1, 1720778192, 1720778192),
                                                                                                                                                                                                                                     (4179, 4178, '查看节点分组', 'A', NULL, NULL, '/cron/api/v1/worker_groups', 'GET', NULL, 'cron:worker:group:query', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1720778192, 1720778192),
                                                                                                                                                                                                                                     (4180, 4178, '新增节点分组', 'A', NULL, NULL, '/cron/api/v1/worker_group', 'POST', NULL, 'cron:worker:group:add', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1720778192, 1720778192),
                                                                                                                                                                                                                                     (4181, 4178, '修改节点分组', 'A', NULL, NULL, '/cron/api/v1/worker_group', 'PUT', NULL, 'cron:worker:group:update', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1720778192, 1720778192),
                                                                                                                                                                                                                                     (4182, 4178, '删除节点分组', 'A', NULL, NULL, '/cron/api/v1/worker_group', 'DELETE', NULL, 'cron:worker:group:delete', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1720778192, 1720778192),
                                                                                                                                                                                                                                     (4183, 4178, '查看节点', 'A', NULL, NULL, '/cron/api/v1/workers', 'GET', NULL, 'cron:worker:query', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1720778192, 1720778192),
                                                                                                                                                                                                                                     (4184, 4178, '新增节点', 'A', NULL, NULL, '/cron/api/v1/worker', 'POST', NULL, 'cron:worker:add', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1720778192, 1720778192),
                                                                                                                                                                                                                                     (4185, 4178, '修改节点', 'A', NULL, NULL, '/cron/api/v1/worker', 'PUT', NULL, 'cron:worker:update', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1720778192, 1720778192),
                                                                                                                                                                                                                                     (4186, 4178, '删除节点', 'A', NULL, NULL, '/cron/api/v1/worker', 'DELETE', NULL, 'cron:worker:delete', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1720778192, 1720778192),
                                                                                                                                                                                                                                     (4187, 4178, '更新节点状态', 'A', NULL, NULL, '/cron/api/v1/worker/status', 'PUT', NULL, 'cron:worker:update:status', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1720778192, 1720778192),
                                                                                                                                                                                                                                     (4188, 4177, '任务管理', 'M', 'Task', 'computer', NULL, NULL, '/cron/task', NULL, '/cron/task/index', NULL, 0, 0, 1, NULL, 1, 1720778192, 1720778192),
                                                                                                                                                                                                                                     (4189, 4188, '查看任务分组', 'A', NULL, NULL, '/cron/api/v1/task_groups', 'GET', NULL, 'cron:task:group:query', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1720778192, 1720778192),
                                                                                                                                                                                                                                     (4190, 4188, '新增任务分组', 'A', NULL, NULL, '/cron/api/v1/task_group', 'POST', NULL, 'cron:task:group:add', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1720778192, 1720778192),
                                                                                                                                                                                                                                     (4191, 4188, '修改任务分组', 'A', NULL, NULL, '/cron/api/v1/task_group', 'PUT', NULL, 'cron:task:group:update', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1720778192, 1720778192),
                                                                                                                                                                                                                                     (4192, 4188, '删除任务分组', 'A', NULL, NULL, '/cron/api/v1/task_group', 'DELETE', NULL, 'cron:task:group:delete', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1720778192, 1720778192),
                                                                                                                                                                                                                                     (4193, 4188, '查看任务', 'A', NULL, NULL, '/cron/api/v1/tasks', 'GET', NULL, 'cron:task:query', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1720778192, 1720778192),
                                                                                                                                                                                                                                     (4194, 4188, '新增任务', 'A', NULL, NULL, '/cron/api/v1/task', 'POST', NULL, 'cron:task:add', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1720778192, 1720778192),
                                                                                                                                                                                                                                     (4195, 4188, '立即执行任务', 'A', NULL, NULL, '/cron/api/v1/task/exec', 'POST', NULL, 'cron:task:exec', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1720778192, 1720778192),
                                                                                                                                                                                                                                     (4196, 4188, '取消执行任务', 'A', NULL, NULL, '/cron/api/v1/task/cancel', 'POST', NULL, 'cron:task:cancel', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1720778192, 1720778192),
                                                                                                                                                                                                                                     (4197, 4188, '修改任务', 'A', NULL, NULL, '/cron/api/v1/task', 'PUT', NULL, 'cron:task:update', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1720778192, 1720778192),
                                                                                                                                                                                                                                     (4198, 4188, '删除任务', 'A', NULL, NULL, '/cron/api/v1/task', 'DELETE', NULL, 'cron:task:delete', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1720778192, 1720778192),
                                                                                                                                                                                                                                     (4199, 4188, '任务状态管理', 'A', NULL, NULL, '/cron/api/v1/task/status', 'PUT', NULL, 'cron:task:update:status', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1720778192, 1720778192),
                                                                                                                                                                                                                                     (4200, 4188, '任务日志', 'G', NULL, NULL, NULL, NULL, NULL, 'cron:task:log', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1720778192, 1720778192),
                                                                                                                                                                                                                                     (4201, 4200, '获取任务日志分页', 'A', NULL, NULL, '/cron/api/v1/logs', 'GET', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL, 1720778192, 1720778192),
                                                                                                                                                                                                                                     (4202, 4200, '获取任务日志详情', 'A', NULL, NULL, '/cron/api/v1/log', 'GET', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL, 1720778192, 1720778192);

-- --------------------------------------------------------

--
-- 表的结构 `menu_closure`
--

CREATE TABLE `menu_closure` (
                                `id` bigint(20) UNSIGNED NOT NULL COMMENT '主键ID',
                                `parent` bigint(20) UNSIGNED NOT NULL COMMENT '菜单id',
                                `children` bigint(20) UNSIGNED NOT NULL COMMENT '菜单id'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='菜单层级信息';

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
                                                                                                                                                                 (6, 1, 'test', 'test', 1, 'test', NULL, 'ALL', 1720638166, 1720638166, 0);

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
    (6, 1, 6);

-- --------------------------------------------------------

--
-- 表的结构 `role_menu`
--

CREATE TABLE `role_menu` (
                             `id` bigint(20) UNSIGNED NOT NULL COMMENT '主键ID',
                             `role_id` bigint(20) UNSIGNED NOT NULL COMMENT '角色id',
                             `menu_id` bigint(20) UNSIGNED NOT NULL COMMENT '菜单id'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='角色菜单信息';

--
-- 转存表中的数据 `role_menu`
--

INSERT INTO `role_menu` (`id`, `role_id`, `menu_id`) VALUES
                                                         (382, 6, 4011),
                                                         (383, 6, 4033),
                                                         (384, 6, 4034),
                                                         (385, 6, 4035),
                                                         (386, 6, 4036),
                                                         (387, 6, 4037),
                                                         (388, 6, 4038),
                                                         (389, 6, 4039),
                                                         (390, 6, 4040),
                                                         (391, 6, 4041),
                                                         (392, 6, 4042),
                                                         (393, 6, 4043),
                                                         (394, 6, 4044),
                                                         (395, 6, 4045),
                                                         (396, 6, 4046),
                                                         (397, 6, 4047),
                                                         (398, 6, 4048),
                                                         (399, 6, 4049),
                                                         (400, 6, 4050),
                                                         (401, 6, 4051),
                                                         (402, 6, 4052),
                                                         (403, 6, 4053),
                                                         (404, 6, 4054),
                                                         (405, 6, 4055),
                                                         (406, 6, 4056),
                                                         (407, 6, 4057),
                                                         (408, 6, 4058),
                                                         (409, 6, 4059),
                                                         (410, 6, 4060),
                                                         (411, 6, 4061),
                                                         (412, 6, 4062),
                                                         (413, 6, 4063),
                                                         (414, 6, 4064),
                                                         (415, 6, 4065),
                                                         (416, 6, 4066),
                                                         (417, 6, 4067),
                                                         (418, 6, 4068),
                                                         (419, 6, 4069),
                                                         (420, 6, 4070),
                                                         (421, 6, 4071);

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
                                                                                                                                                                                                       (1, 1, 1, '超级管理员', '超级管理员', 'F', '25c560e9cae8b86eb58828f457760bc6', '1280291001@qq.com', '18888888888', '$2a$10$9qRJe9KQo6sEcU8ipKg.e.dkl2E7Wy64SigYlgraTAn.1paHFq6W.', 1, '{\"theme\":\"light\",\"themeColor\":\"#5F6A82\",\"skin\":\"default\",\"tabBar\":true,\"menuWidth\":200,\"layout\":\"default\",\"language\":\"zh_CN\",\"animation\":\"gp-fade\"}', 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJkZXBhcnRtZW50SWQiOjEsImRlcGFydG1lbnRLZXl3b3JkIjoiY29tcGFueSIsImV4cCI6MTcyNDMxNTkwMSwiaWF0IjoxNzI0MzA4NzAwLCJuYmYiOjE3MjQzMDg3MDAsInJvbGVJZCI6MSwicm9sZUtleXdvcmQiOiJzdXBlckFkbWluIiwidXNlcklkIjoxfQ.j5cVgS2JBUutIaZGszZOvSNf7tUBogOIS64d3EDqeIs', 1724308700, 1713706137, 1724308700),
                                                                                                                                                                                                       (4, 1, 6, 'test2', 'test2', 'M', '', '16666666666@gmail.com', '16666666666', '$2a$10$UsKv0fmTZ6q0NiK0yzkE9OQ1XYG8gPccYigjuZGEZmBNpixDrfJGO', 1, NULL, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJkZXBhcnRtZW50SWQiOjEsImRlcGFydG1lbnRLZXl3b3JkIjoiY29tcGFueSIsImV4cCI6MTcyMDY0NTU2NSwiaWF0IjoxNzIwNjM4MzY0LCJuYmYiOjE3MjA2MzgzNjQsInJvbGVJZCI6Niwicm9sZUtleXdvcmQiOiJ0ZXN0IiwidXNlcklkIjo0fQ.JaZqdifR8t2wvU3FcBzawex2vkESgVPfJPjayMvgURs', 1720638364, 1720638351, 1722413160);

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
                                                       (1, 1, 1),
                                                       (9, 1, 4);

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
                                                         (1, 1, 1),
                                                         (6, 6, 4);

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
    MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=751;

--
-- 使用表AUTO_INCREMENT `department`
--
ALTER TABLE `department`
    MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键ID', AUTO_INCREMENT=6;

--
-- 使用表AUTO_INCREMENT `department_closure`
--
ALTER TABLE `department_closure`
    MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键ID', AUTO_INCREMENT=11;

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
    MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键ID', AUTO_INCREMENT=4203;

--
-- 使用表AUTO_INCREMENT `menu_closure`
--
ALTER TABLE `menu_closure`
    MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键ID', AUTO_INCREMENT=6018;

--
-- 使用表AUTO_INCREMENT `role`
--
ALTER TABLE `role`
    MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键ID', AUTO_INCREMENT=7;

--
-- 使用表AUTO_INCREMENT `role_closure`
--
ALTER TABLE `role_closure`
    MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键ID', AUTO_INCREMENT=7;

--
-- 使用表AUTO_INCREMENT `role_menu`
--
ALTER TABLE `role_menu`
    MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键ID', AUTO_INCREMENT=423;

--
-- 使用表AUTO_INCREMENT `user`
--
ALTER TABLE `user`
    MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键ID', AUTO_INCREMENT=5;

--
-- 使用表AUTO_INCREMENT `user_job`
--
ALTER TABLE `user_job`
    MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键ID', AUTO_INCREMENT=10;

--
-- 使用表AUTO_INCREMENT `user_role`
--
ALTER TABLE `user_role`
    MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键ID', AUTO_INCREMENT=7;

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
