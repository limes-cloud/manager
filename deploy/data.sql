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

 Date: 30/04/2024 15:47:08
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
) ENGINE=InnoDB AUTO_INCREMENT=1699 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of casbin_rule
-- ----------------------------
BEGIN;
INSERT INTO `casbin_rule` VALUES (1626, 'p', 'admin', '/configure/business/value', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES (1625, 'p', 'admin', '/configure/v1/business', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES (1622, 'p', 'admin', '/configure/v1/business', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (1623, 'p', 'admin', '/configure/v1/business', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (1624, 'p', 'admin', '/configure/v1/business', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES (1635, 'p', 'admin', '/configure/v1/configure', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES (1634, 'p', 'admin', '/configure/v1/configure/compare', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (1611, 'p', 'admin', '/configure/v1/env', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES (1609, 'p', 'admin', '/configure/v1/env', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (1610, 'p', 'admin', '/configure/v1/env', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES (1612, 'p', 'admin', '/configure/v1/env/token', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (1613, 'p', 'admin', '/configure/v1/env/token', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES (1608, 'p', 'admin', '/configure/v1/envs', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (1621, 'p', 'admin', '/configure/v1/resource', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES (1619, 'p', 'admin', '/configure/v1/resource', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (1620, 'p', 'admin', '/configure/v1/resource', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES (1618, 'p', 'admin', '/configure/v1/resources', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (1617, 'p', 'admin', '/configure/v1/server', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES (1615, 'p', 'admin', '/configure/v1/server', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (1616, 'p', 'admin', '/configure/v1/server', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES (1614, 'p', 'admin', '/configure/v1/servers', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (1628, 'p', 'admin', '/configure/v1/template', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (1630, 'p', 'admin', '/configure/v1/template', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (1631, 'p', 'admin', '/configure/v1/template/compare', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (1629, 'p', 'admin', '/configure/v1/template/current', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (1633, 'p', 'admin', '/configure/v1/template/preview', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (1632, 'p', 'admin', '/configure/v1/template/switch', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (1627, 'p', 'admin', '/configure/v1/templates', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (1668, 'p', 'admin', '/cron/v1/log', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (1667, 'p', 'admin', '/cron/v1/logs', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (1664, 'p', 'admin', '/cron/v1/task', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES (1660, 'p', 'admin', '/cron/v1/task', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (1663, 'p', 'admin', '/cron/v1/task', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES (1662, 'p', 'admin', '/cron/v1/task/cancel', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (1666, 'p', 'admin', '/cron/v1/task/disable', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (1665, 'p', 'admin', '/cron/v1/task/enable', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (1661, 'p', 'admin', '/cron/v1/task/exec', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (1658, 'p', 'admin', '/cron/v1/task/group', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES (1656, 'p', 'admin', '/cron/v1/task/group', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (1657, 'p', 'admin', '/cron/v1/task/group', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES (1655, 'p', 'admin', '/cron/v1/task/groups', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (1659, 'p', 'admin', '/cron/v1/tasks', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (1652, 'p', 'admin', '/cron/v1/worker', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES (1650, 'p', 'admin', '/cron/v1/worker', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (1651, 'p', 'admin', '/cron/v1/worker', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES (1654, 'p', 'admin', '/cron/v1/worker/disable', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (1653, 'p', 'admin', '/cron/v1/worker/enable', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (1648, 'p', 'admin', '/cron/v1/worker/group', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES (1646, 'p', 'admin', '/cron/v1/worker/group', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (1647, 'p', 'admin', '/cron/v1/worker/group', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES (1645, 'p', 'admin', '/cron/v1/worker/groups', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (1649, 'p', 'admin', '/cron/v1/workers', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (1590, 'p', 'admin', '/manager/v1/department', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES (1588, 'p', 'admin', '/manager/v1/department', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (1589, 'p', 'admin', '/manager/v1/department', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES (1591, 'p', 'admin', '/manager/v1/department/objects', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (1592, 'p', 'admin', '/manager/v1/department/objects', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (1587, 'p', 'admin', '/manager/v1/department/tree', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (1571, 'p', 'admin', '/manager/v1/dictionaries', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (1574, 'p', 'admin', '/manager/v1/dictionary', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES (1572, 'p', 'admin', '/manager/v1/dictionary', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (1573, 'p', 'admin', '/manager/v1/dictionary', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES (1578, 'p', 'admin', '/manager/v1/dictionary/value', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES (1576, 'p', 'admin', '/manager/v1/dictionary/value', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (1577, 'p', 'admin', '/manager/v1/dictionary/value', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES (1575, 'p', 'admin', '/manager/v1/dictionary/values', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (1586, 'p', 'admin', '/manager/v1/job', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES (1584, 'p', 'admin', '/manager/v1/job', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (1585, 'p', 'admin', '/manager/v1/job', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES (1583, 'p', 'admin', '/manager/v1/jobs', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (1582, 'p', 'admin', '/manager/v1/menu', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES (1580, 'p', 'admin', '/manager/v1/menu', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (1581, 'p', 'admin', '/manager/v1/menu', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES (1579, 'p', 'admin', '/manager/v1/menu/tree', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (1570, 'p', 'admin', '/manager/v1/object', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES (1568, 'p', 'admin', '/manager/v1/object', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (1569, 'p', 'admin', '/manager/v1/object', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES (1567, 'p', 'admin', '/manager/v1/objects', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (1596, 'p', 'admin', '/manager/v1/role', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES (1594, 'p', 'admin', '/manager/v1/role', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (1595, 'p', 'admin', '/manager/v1/role', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES (1598, 'p', 'admin', '/manager/v1/role/menu', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (1597, 'p', 'admin', '/manager/v1/role/menu/ids', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (1593, 'p', 'admin', '/manager/v1/role/tree', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (1604, 'p', 'admin', '/manager/v1/user', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES (1600, 'p', 'admin', '/manager/v1/user', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (1601, 'p', 'admin', '/manager/v1/user', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES (1606, 'p', 'admin', '/manager/v1/user/disable', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (1605, 'p', 'admin', '/manager/v1/user/enable', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (1603, 'p', 'admin', '/manager/v1/user/jobs', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (1607, 'p', 'admin', '/manager/v1/user/password/reset', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (1602, 'p', 'admin', '/manager/v1/user/roles', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (1599, 'p', 'admin', '/manager/v1/users', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (1636, 'p', 'admin', '/resource/v1/directories', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (1639, 'p', 'admin', '/resource/v1/directory', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES (1637, 'p', 'admin', '/resource/v1/directory', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (1638, 'p', 'admin', '/resource/v1/directory', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES (1642, 'p', 'admin', '/resource/v1/file', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES (1641, 'p', 'admin', '/resource/v1/file', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES (1640, 'p', 'admin', '/resource/v1/files', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (1643, 'p', 'admin', '/resource/v1/upload', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (1644, 'p', 'admin', '/resource/v1/upload/prepare', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (1673, 'p', 'admin', '/user-center/v1/agreement/content', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES (1670, 'p', 'admin', '/user-center/v1/agreement/content', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (1671, 'p', 'admin', '/user-center/v1/agreement/content', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (1672, 'p', 'admin', '/user-center/v1/agreement/content', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES (1669, 'p', 'admin', '/user-center/v1/agreement/contents', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (1678, 'p', 'admin', '/user-center/v1/agreement/scene', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES (1675, 'p', 'admin', '/user-center/v1/agreement/scene', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (1676, 'p', 'admin', '/user-center/v1/agreement/scene', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (1677, 'p', 'admin', '/user-center/v1/agreement/scene', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES (1674, 'p', 'admin', '/user-center/v1/agreement/scenes', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (1690, 'p', 'admin', '/user-center/v1/app', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES (1688, 'p', 'admin', '/user-center/v1/app', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (1689, 'p', 'admin', '/user-center/v1/app', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES (1687, 'p', 'admin', '/user-center/v1/apps', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (1682, 'p', 'admin', '/user-center/v1/channel', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES (1680, 'p', 'admin', '/user-center/v1/channel', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (1681, 'p', 'admin', '/user-center/v1/channel', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES (1679, 'p', 'admin', '/user-center/v1/channels', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (1686, 'p', 'admin', '/user-center/v1/field', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES (1684, 'p', 'admin', '/user-center/v1/field', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (1685, 'p', 'admin', '/user-center/v1/field', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES (1683, 'p', 'admin', '/user-center/v1/fields', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (1695, 'p', 'admin', '/user-center/v1/user', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES (1692, 'p', 'admin', '/user-center/v1/user', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (1694, 'p', 'admin', '/user-center/v1/user', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES (1697, 'p', 'admin', '/user-center/v1/user/disable', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (1698, 'p', 'admin', '/user-center/v1/user/enable', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (1696, 'p', 'admin', '/user-center/v1/user/offline', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (1691, 'p', 'admin', '/user-center/v1/users', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (1693, 'p', 'admin', '/user-center/v1/users', 'POST', '', '', '');
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
                              KEY `idx_department_updated_at` (`updated_at`),
                              KEY `idx_department_created_at` (`created_at`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COMMENT='部门信息';

-- ----------------------------
-- Records of department
-- ----------------------------
BEGIN;
INSERT INTO `department` VALUES (1, 1713706137, 1713706137, 0, 'company', '贵州青橙科技有限公司', '开放合作，拥抱未来');
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
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='部门层级信息';

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
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='部门资源';

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
                              KEY `idx_dictionary_created_at` (`created_at`),
                              KEY `idx_dictionary_updated_at` (`updated_at`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COMMENT='字典信息';

-- ----------------------------
-- Records of dictionary
-- ----------------------------
BEGIN;
INSERT INTO `dictionary` VALUES (1, 1713706137, 1713706137, 'gender', '性别', '性别选项');
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
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COMMENT='字典值信息';

-- ----------------------------
-- Records of dictionary_value
-- ----------------------------
BEGIN;
INSERT INTO `dictionary_value` VALUES (1, 1713706137, 1713706137, 1, '男', 'M', 1, 0, '', '', '男性');
INSERT INTO `dictionary_value` VALUES (2, 1713706137, 1713706137, 1, '女', 'F', 1, 0, '', '', '女性');
INSERT INTO `dictionary_value` VALUES (3, 1713706137, 1713706137, 1, '未知', 'U', 1, 0, '', '', '未知性别');
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
                       KEY `idx_job_updated_at` (`updated_at`),
                       KEY `idx_job_created_at` (`created_at`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COMMENT='职位信息';

-- ----------------------------
-- Records of job
-- ----------------------------
BEGIN;
INSERT INTO `job` VALUES (1, 1713706137, 1713706137, 'chairman', '董事长', 0, '董事长');
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
) ENGINE=InnoDB AUTO_INCREMENT=772 DEFAULT CHARSET=utf8mb4 COMMENT='菜单信息';

-- ----------------------------
-- Records of menu
-- ----------------------------
BEGIN;
INSERT INTO `menu` VALUES (522, 1714462150, 1714462150, 0, '管理平台', 'R', 'SystemPlatform', 'apps', NULL, NULL, '/', NULL, NULL, NULL, 'Layout', NULL, 2, 0, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (523, 1714462150, 1714462150, 522, '首页面板', 'M', 'Dashboard', 'dashboard', NULL, NULL, '/dashboard', NULL, NULL, NULL, '/dashboard/index', NULL, 0, 0, 1, 1, 1);
INSERT INTO `menu` VALUES (524, 1714462150, 1714462150, 522, '管理中心', 'M', 'SystemPlatformManager', 'apps', NULL, NULL, '/manager', NULL, NULL, NULL, NULL, NULL, 0, 0, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (525, 1714462150, 1714462150, 524, '基础接口', 'G', 'BaseApi', 'apps', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, 0, 1, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (526, 1714462150, 1714462150, 525, '获取用户部门', 'BA', NULL, NULL, '/manager/v1/department/tree', 'GET', NULL, NULL, NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (527, 1714462150, 1714462150, 525, '获取个人用户信息', 'BA', NULL, NULL, '/manager/v1/user/current', 'GET', NULL, NULL, NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (528, 1714462150, 1714462150, 525, '获取用户可见角色树', 'BA', NULL, NULL, '/manager/v1/role/tree', 'GET', NULL, NULL, NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (529, 1714462150, 1714462150, 525, '获取用户菜单', 'BA', NULL, NULL, '/manager/v1/menu/tree/from/role', 'GET', NULL, NULL, NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (530, 1714462150, 1714462150, 525, '退出登录', 'BA', NULL, NULL, '/manager/v1/user/logout', 'POST', NULL, NULL, NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (531, 1714462150, 1714462150, 525, '刷新token', 'BA', NULL, NULL, '/manager/v1/user/token/refresh', 'POST', NULL, NULL, NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (532, 1714462150, 1714462150, 525, '用户登录', 'BA', NULL, NULL, '/manager/v1/user/login', 'POST', NULL, NULL, NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (533, 1714462150, 1714462150, 525, '获取登录验证码', 'BA', NULL, NULL, '/manager/v1/user/login/captcha', 'POST', NULL, NULL, NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (534, 1714462150, 1714462150, 525, '获取系统基础设置', 'BA', NULL, NULL, '/manager/v1/setting', 'GET', NULL, NULL, NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (535, 1714462150, 1714462150, 525, '接口鉴权', 'BA', NULL, NULL, '/manager/v1/auth', 'POST', NULL, NULL, NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (536, 1714462150, 1714462150, 525, '切换用户角色', 'BA', NULL, NULL, '/manager/v1/user/role/switch', 'POST', NULL, NULL, NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (537, 1714462150, 1714462150, 525, '新增部门资源对象', 'BA', NULL, NULL, '/manager/v1/department/object', 'POST', NULL, NULL, NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (538, 1714462150, 1714462150, 525, '删除部门资源对象', 'BA', NULL, NULL, '/manager/v1/department/object', 'DELETE', NULL, NULL, NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (539, 1714462150, 1714462150, 524, '资源对象', 'M', 'Object', 'common', NULL, NULL, '/manager/object', NULL, NULL, NULL, '/manager/object/index', NULL, 0, NULL, 1, NULL, 1);
INSERT INTO `menu` VALUES (540, 1714462150, 1714462150, 539, '查询对象', 'A', NULL, NULL, '/manager/v1/objects', 'GET', NULL, 'manager:object:query', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (541, 1714462150, 1714462150, 539, '新增对象', 'A', NULL, NULL, '/manager/v1/object', 'POST', NULL, 'manager:object:add', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (542, 1714462150, 1714462150, 539, '修改对象', 'A', NULL, NULL, '/manager/v1/object', 'PUT', NULL, 'manager:object:update', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (543, 1714462150, 1714462150, 539, '删除对象', 'A', NULL, NULL, '/manager/v1/object', 'DELETE', NULL, 'manager:object:delete', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (544, 1714462150, 1714462150, 524, '字典管理', 'M', 'Dict', 'storage', NULL, NULL, '/manager/dictionary', NULL, NULL, NULL, '/manager/dictionary/index', NULL, 0, NULL, 1, NULL, 1);
INSERT INTO `menu` VALUES (545, 1714462150, 1714462150, 544, '查询字典', 'A', NULL, NULL, '/manager/v1/dictionaries', 'GET', NULL, 'manager:dictionary:query', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (546, 1714462150, 1714462150, 544, '新增字典', 'A', NULL, NULL, '/manager/v1/dictionary', 'POST', NULL, 'manager:dictionary:add', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (547, 1714462150, 1714462150, 544, '修改字典', 'A', NULL, NULL, '/manager/v1/dictionary', 'PUT', NULL, 'manager:dictionary:update', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (548, 1714462150, 1714462150, 544, '删除字典', 'A', NULL, NULL, '/manager/v1/dictionary', 'DELETE', NULL, 'manager:dictionary:delete', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (549, 1714462150, 1714462150, 544, '刷新字典值', 'A', NULL, NULL, '/manager/v1/dictionary/values', 'GET', NULL, 'manager:dictionary:value:query', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (550, 1714462150, 1714462150, 544, '新增字典值', 'A', NULL, NULL, '/manager/v1/dictionary/value', 'POST', NULL, 'manager:dictionary:value:add', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (551, 1714462150, 1714462150, 544, '修改字典值', 'A', NULL, NULL, '/manager/v1/dictionary/value', 'PUT', NULL, 'manager:dictionary:value:update', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (552, 1714462150, 1714462150, 544, '删除字典值', 'A', NULL, NULL, '/manager/v1/dictionary/value', 'DELETE', NULL, 'manager:dictionary:value:delete', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (553, 1714462150, 1714462150, 524, '菜单管理', 'M', 'ManagerMenu', 'menu', NULL, NULL, '/manager/menu', NULL, NULL, NULL, '/manager/menu/index', NULL, 0, 0, 1, 1, 1);
INSERT INTO `menu` VALUES (554, 1714462150, 1714462150, 553, '查询菜单', 'A', NULL, NULL, '/manager/v1/menu/tree', 'GET', NULL, 'manager:menu:query', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (555, 1714462150, 1714462150, 553, '新增菜单', 'A', NULL, NULL, '/manager/v1/menu', 'POST', NULL, 'manager:menu:add', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (556, 1714462150, 1714462150, 553, '修改菜单', 'A', NULL, NULL, '/manager/v1/menu', 'PUT', NULL, 'manager:menu:update', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (557, 1714462150, 1714462150, 553, '删除菜单', 'A', NULL, NULL, '/manager/v1/menu', 'DELETE', NULL, 'manager:menu:delete', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (558, 1714462150, 1714462150, 524, '职位管理', 'M', 'JobUser', 'tag', NULL, NULL, '/manager/job', NULL, NULL, NULL, '/manager/job/index', NULL, 0, NULL, 1, NULL, 1);
INSERT INTO `menu` VALUES (559, 1714462150, 1714462150, 558, '查询职位', 'A', NULL, NULL, '/manager/v1/jobs', 'GET', NULL, 'manager:job:query', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (560, 1714462150, 1714462150, 558, '新增职位', 'A', NULL, NULL, '/manager/v1/job', 'POST', NULL, 'manager:job:add', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (561, 1714462150, 1714462150, 558, '修改职位', 'A', NULL, NULL, '/manager/v1/job', 'PUT', NULL, 'manager:job:update', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (562, 1714462150, 1714462150, 558, '删除职位', 'A', NULL, NULL, '/manager/v1/job', 'DELETE', NULL, 'manager:job:delete', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (563, 1714462150, 1714462150, 524, '部门管理', 'M', 'ManagerDepartment', 'user-group', NULL, NULL, '/manager/department', NULL, NULL, NULL, '/manager/department/index', NULL, 0, 0, 1, NULL, 1);
INSERT INTO `menu` VALUES (564, 1714462150, 1714462150, 563, '查看部门', 'A', NULL, NULL, '/manager/v1/department/tree', 'GET', NULL, 'manager:department:query', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (565, 1714462150, 1714462150, 563, '新增部门', 'A', NULL, NULL, '/manager/v1/department', 'POST', NULL, 'manager:department:add', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (566, 1714462150, 1714462150, 563, '修改部门', 'A', NULL, NULL, '/manager/v1/department', 'PUT', NULL, 'manager:department:update', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (567, 1714462150, 1714462150, 563, '删除部门', 'A', NULL, NULL, '/manager/v1/department', 'DELETE', NULL, 'manager:department:delete', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (568, 1714462150, 1714462150, 563, '查看部门资源对象', 'A', NULL, NULL, '/manager/v1/department/objects', 'GET', NULL, 'manager:department:object:query', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (569, 1714462150, 1714462150, 563, '设置部门资源对象', 'A', NULL, NULL, '/manager/v1/department/objects', 'POST', NULL, 'manager:department:object:import', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (570, 1714462150, 1714462150, 524, '角色管理', 'M', 'ManagerRole', 'safe', NULL, NULL, '/manager/role', NULL, NULL, NULL, '/manager/role/index', NULL, 0, NULL, 0, NULL, 1);
INSERT INTO `menu` VALUES (571, 1714462150, 1714462150, 570, '查询角色', 'A', NULL, NULL, '/manager/v1/role/tree', 'GET', NULL, 'manager:role:query', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (572, 1714462150, 1714462150, 570, '新增角色', 'A', NULL, NULL, '/manager/v1/role', 'POST', NULL, 'manager:role:add', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (573, 1714462150, 1714462150, 570, '修改角色', 'A', NULL, NULL, '/manager/v1/role', 'PUT', NULL, 'manager:role:update', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (574, 1714462150, 1714462150, 570, '删除角色', 'A', NULL, NULL, '/manager/v1/role', 'DELETE', NULL, 'manager:role:delete', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (575, 1714462150, 1714462150, 570, '角色菜单管理', 'G', NULL, NULL, '/manager/v1/role/menu/tree', 'GET', NULL, 'manager:role:menu', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (576, 1714462150, 1714462150, 575, '查询角色菜单', 'A', NULL, NULL, '/manager/v1/role/menu/ids', 'GET', NULL, 'manager:role:menu:query', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (577, 1714462150, 1714462150, 575, '修改角色菜单', 'A', NULL, NULL, '/manager/v1/role/menu', 'POST', NULL, 'manager:role:menu:update', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (578, 1714462150, 1714462150, 524, '用户管理', 'M', 'ManagerUser', 'user', NULL, NULL, '/manager/user', NULL, NULL, NULL, '/manager/user/index', NULL, 0, NULL, 1, NULL, 1);
INSERT INTO `menu` VALUES (579, 1714462150, 1714462150, 578, '查询用户', 'A', NULL, NULL, '/manager/v1/users', 'GET', NULL, 'manager:user:query', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (580, 1714462150, 1714462150, 578, '新增用户', 'A', NULL, NULL, '/manager/v1/user', 'POST', NULL, 'manager:user:add', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (581, 1714462150, 1714462150, 578, '修改用户', 'A', NULL, NULL, '/manager/v1/user', 'PUT', NULL, 'manager:user:update', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (582, 1714462150, 1714462150, 581, '获取用户角色列表', 'A', NULL, NULL, '/manager/v1/user/roles', 'GET', NULL, 'manager:user:update', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (583, 1714462150, 1714462150, 581, '获取用户部门列表', 'A', NULL, NULL, '/manager/v1/user/jobs', 'POST', NULL, 'manager:user:update', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (584, 1714462150, 1714462150, 578, '删除用户', 'A', NULL, NULL, '/manager/v1/user', 'DELETE', NULL, 'manager:user:delete', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (585, 1714462150, 1714462150, 578, '启用/禁用用户', 'G', NULL, NULL, NULL, NULL, NULL, 'manager:user:status', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (586, 1714462150, 1714462150, 585, '启用用户', 'A', NULL, NULL, '/manager/v1/user/enable', 'POST', NULL, 'manager:user:status:enable', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (587, 1714462150, 1714462150, 585, '禁用用户', 'A', NULL, NULL, '/manager/v1/user/disable', 'POST', NULL, 'manager:user:status:disable', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (588, 1714462150, 1714462150, 578, '重置账号密码', 'A', NULL, NULL, '/manager/v1/user/password/reset', 'POST', NULL, 'manager:user:reset:password', NULL, NULL, '', NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (589, 1714462150, 1714462150, 522, '配置中心', 'M', 'SystemPlatformConfigure', 'code-block', NULL, NULL, '/configure', NULL, NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (590, 1714462150, 1714462150, 589, '环境管理', 'M', 'ConfigureEnv', 'common', NULL, NULL, '/configure/env', NULL, NULL, NULL, '/configure/env/index', NULL, 0, NULL, 1, NULL, 1);
INSERT INTO `menu` VALUES (591, 1714462150, 1714462150, 590, '查看环境', 'A', NULL, NULL, '/configure/v1/envs', 'GET', NULL, 'configure:env:query', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (592, 1714462150, 1714462150, 590, '新增环境', 'A', NULL, NULL, '/configure/v1/env', 'POST', NULL, 'configure:env:add', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (593, 1714462150, 1714462150, 590, '修改环境', 'A', NULL, NULL, '/configure/v1/env', 'PUT', NULL, 'configure:env:update', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (594, 1714462150, 1714462150, 590, '删除环境', 'A', NULL, NULL, '/configure/v1/env', 'DELETE', NULL, 'configure:env:delete', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (595, 1714462150, 1714462150, 590, '查看环境Token', 'A', NULL, NULL, '/configure/v1/env/token', 'GET', NULL, 'configure:env:token:query', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (596, 1714462150, 1714462150, 590, '重置环境token', 'A', NULL, NULL, '/configure/v1/env/token', 'PUT', NULL, 'configure:env:token:reset', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (597, 1714462150, 1714462150, 589, '服务管理', 'M', 'ConfigureServer', 'apps', NULL, NULL, '/configure/server', NULL, NULL, NULL, '/configure/server/index', NULL, 0, NULL, 1, NULL, 1);
INSERT INTO `menu` VALUES (598, 1714462150, 1714462150, 597, '查询服务', 'A', NULL, NULL, '/configure/v1/servers', 'GET', NULL, 'configure:server:query', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (599, 1714462150, 1714462150, 597, '新增服务', 'A', NULL, NULL, '/configure/v1/server', 'POST', NULL, 'configure:server:add', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (600, 1714462150, 1714462150, 597, '修改服务', 'A', NULL, NULL, '/configure/v1/server', 'PUT', NULL, 'configure:server:update', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (601, 1714462150, 1714462150, 597, '删除服务', 'A', NULL, NULL, '/configure/v1/server', 'DELETE', NULL, 'configure:server:delete', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (602, 1714462150, 1714462150, 589, '资源变量', 'M', 'ConfigureResource', 'unordered-list', NULL, NULL, '/configure/resource', NULL, NULL, NULL, '/configure/resource/index', NULL, 0, NULL, 1, NULL, 1);
INSERT INTO `menu` VALUES (603, 1714462150, 1714462150, 602, '查看资源', 'A', NULL, NULL, '/configure/v1/resources', 'GET', NULL, 'configure:resource:query', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (604, 1714462150, 1714462150, 602, '新增资源', 'A', NULL, NULL, '/configure/v1/resource', 'POST', NULL, 'configure:resource:add', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (605, 1714462150, 1714462150, 602, '修改资源', 'A', NULL, NULL, '/configure/v1/resource', 'PUT', NULL, 'configure:resource:update', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (606, 1714462150, 1714462150, 602, '删除资源', 'A', NULL, NULL, '/configure/v1/resource', 'DELETE', NULL, 'configure:resource:delete', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (607, 1714462150, 1714462150, 602, '资源变量值配置', 'G', NULL, NULL, NULL, NULL, NULL, 'configure:resource:value', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (608, 1714462150, 1714462150, 589, '业务变量', 'M', 'ConfigureBusiness', 'code', NULL, NULL, '/configure/business', NULL, NULL, NULL, '/configure/business/index', NULL, 0, NULL, 1, NULL, 1);
INSERT INTO `menu` VALUES (609, 1714462150, 1714462150, 608, '查看业务变量', 'A', NULL, NULL, '/configure/v1/business', 'GET', NULL, 'configure:business:query', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (610, 1714462150, 1714462150, 608, '新增业务变量', 'A', NULL, NULL, '/configure/v1/business', 'POST', NULL, 'configure:business:add', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (611, 1714462150, 1714462150, 608, '修改业务变量', 'A', NULL, NULL, '/configure/v1/business', 'PUT', NULL, 'configure:business:update', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (612, 1714462150, 1714462150, 608, '删除业务变量', 'A', NULL, NULL, '/configure/v1/business', 'DELETE', NULL, 'configure:business:delete', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (613, 1714462150, 1714462150, 608, '配置业务变量值', 'A', NULL, NULL, '/configure/business/value', 'PUT', NULL, 'configure:business:value', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (614, 1714462150, 1714462150, 589, '配置模板', 'M', 'ConfgureTemplate', 'file', NULL, NULL, '/configure/template', NULL, NULL, NULL, '/configure/template/index', NULL, 0, NULL, 1, NULL, 1);
INSERT INTO `menu` VALUES (615, 1714462150, 1714462150, 614, '模板管理', 'G', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (616, 1714462150, 1714462150, 615, '查看模板历史版本', 'A', NULL, NULL, '/configure/v1/templates', 'GET', NULL, 'configure:template:history', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (617, 1714462150, 1714462150, 615, '查看指定模板详细数据', 'A', NULL, NULL, '/configure/v1/template', 'GET', NULL, NULL, NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (618, 1714462150, 1714462150, 615, '查看当前正在使用的模板', 'A', NULL, NULL, '/configure/v1/template/current', 'GET', NULL, NULL, NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (619, 1714462150, 1714462150, 615, '提交模板', 'A', NULL, NULL, '/configure/v1/template', 'POST', NULL, 'configure:template:add', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (620, 1714462150, 1714462150, 615, '模板对比', 'A', NULL, NULL, '/configure/v1/template/compare', 'POST', NULL, 'configure:template:compare', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (621, 1714462150, 1714462150, 615, '切换模板', 'A', NULL, NULL, '/configure/v1/template/switch', 'POST', NULL, 'configure:template:switch', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (622, 1714462150, 1714462150, 615, '模板预览', 'A', NULL, NULL, '/configure/v1/template/preview', 'POST', NULL, 'configure:template:preview', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (623, 1714462150, 1714462150, 614, '配置管理', 'G', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (624, 1714462150, 1714462150, 623, '配置对比', 'A', NULL, NULL, '/configure/v1/configure/compare', 'POST', NULL, 'configure:configure:compare', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (625, 1714462150, 1714462150, 623, '同步配置', 'A', NULL, NULL, '/configure/v1/configure', 'PUT', NULL, 'configure:configure:sync', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (626, 1714462150, 1714462150, 522, '资源中心', 'M', 'SystemPlatformResource', 'file', NULL, NULL, '/resource', NULL, NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (627, 1714462150, 1714462150, 626, '文件管理', 'M', 'File', 'file', NULL, NULL, '/resource/file', NULL, NULL, NULL, '/resource/file/index', NULL, 0, 0, 1, NULL, 1);
INSERT INTO `menu` VALUES (628, 1714462150, 1714462150, 627, '目录管理', 'G', NULL, NULL, NULL, NULL, NULL, 'resource:directory:group', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (629, 1714462150, 1714462150, 628, '查看目录', 'A', NULL, NULL, '/resource/v1/directories', 'GET', NULL, 'resource:directory:query', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (630, 1714462150, 1714462150, 628, '新增目录', 'A', NULL, NULL, '/resource/v1/directory', 'POST', NULL, 'resource:directory:add', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (631, 1714462150, 1714462150, 628, '修改目录', 'A', NULL, NULL, '/resource/v1/directory', 'PUT', NULL, 'resource:directory:update', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (632, 1714462150, 1714462150, 628, '删除目录', 'A', NULL, NULL, '/resource/v1/directory', 'DELETE', NULL, 'resource:directory:delete', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (633, 1714462150, 1714462150, 627, '文件管理', 'G', NULL, NULL, NULL, NULL, NULL, 'resource:file:group', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (634, 1714462150, 1714462150, 633, '查看文件', 'A', NULL, NULL, '/resource/v1/files', 'GET', NULL, 'resource:file:query', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (635, 1714462150, 1714462150, 633, '修改文件', 'A', NULL, NULL, '/resource/v1/file', 'PUT', NULL, 'resource:file:upload', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (636, 1714462150, 1714462150, 633, '删除文件', 'A', NULL, NULL, '/resource/v1/file', 'DELETE', NULL, 'resource:file:delete', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (637, 1714462150, 1714462150, 633, '上传文件', 'G', NULL, NULL, NULL, NULL, NULL, 'resource:file:upload:group', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (638, 1714462150, 1714462150, 637, '分块上传文件', 'A', NULL, NULL, '/resource/v1/upload', 'POST', NULL, 'resource:file:upload', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (639, 1714462150, 1714462150, 637, '预上传文件', 'A', NULL, NULL, '/resource/v1/upload/prepare', 'POST', NULL, 'resource:file:upload', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (640, 1714462150, 1714462150, 522, '定时任务', 'M', 'SystemPlatformCron', 'schedule', NULL, NULL, '/cron', NULL, NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (641, 1714462150, 1714462150, 640, '节点管理', 'M', 'Worker', 'common', NULL, NULL, '/cron/worker', NULL, NULL, NULL, '/cron/worker/index', NULL, 0, 0, 1, NULL, 1);
INSERT INTO `menu` VALUES (642, 1714462150, 1714462150, 641, '查看节点分组', 'A', NULL, NULL, '/cron/v1/worker/groups', 'GET', NULL, 'cron:worker:group:query', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (643, 1714462150, 1714462150, 641, '新增节点分组', 'A', NULL, NULL, '/cron/v1/worker/group', 'POST', NULL, 'cron:worker:group:add', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (644, 1714462150, 1714462150, 641, '修改节点分组', 'A', NULL, NULL, '/cron/v1/worker/group', 'PUT', NULL, 'cron:worker:group:update', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (645, 1714462150, 1714462150, 641, '删除节点分组', 'A', NULL, NULL, '/cron/v1/worker/group', 'DELETE', NULL, 'cron:worker:group:delete', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (646, 1714462150, 1714462150, 641, '查看节点', 'A', NULL, NULL, '/cron/v1/workers', 'GET', NULL, 'cron:worker:query', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (647, 1714462150, 1714462150, 641, '新增节点', 'A', NULL, NULL, '/cron/v1/worker', 'POST', NULL, 'cron:worker:add', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (648, 1714462150, 1714462150, 641, '修改节点', 'A', NULL, NULL, '/cron/v1/worker', 'PUT', NULL, 'cron:worker:update', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (649, 1714462150, 1714462150, 641, '删除节点', 'A', NULL, NULL, '/cron/v1/worker', 'DELETE', NULL, 'cron:worker:delete', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (650, 1714462150, 1714462150, 641, '节点状态管理', 'G', NULL, NULL, NULL, NULL, NULL, 'cron:worker:status', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (651, 1714462150, 1714462150, 650, '启用节点', 'A', NULL, NULL, '/cron/v1/worker/enable', 'POST', NULL, 'cron:worker:status:enable', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (652, 1714462150, 1714462150, 650, '启用节点', 'A', NULL, NULL, '/cron/v1/worker/disable', 'POST', NULL, 'cron:worker:status:disable', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (653, 1714462150, 1714462150, 640, '任务管理', 'M', 'Task', 'computer', NULL, NULL, '/cron/task', NULL, NULL, NULL, '/cron/task/index', NULL, 0, 0, 1, NULL, 1);
INSERT INTO `menu` VALUES (654, 1714462150, 1714462150, 653, '查看任务分组', 'A', NULL, NULL, '/cron/v1/task/groups', 'GET', NULL, 'cron:task:group:query', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (655, 1714462150, 1714462150, 653, '新增任务分组', 'A', NULL, NULL, '/cron/v1/task/group', 'POST', NULL, 'cron:task:group:add', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (656, 1714462150, 1714462150, 653, '修改任务分组', 'A', NULL, NULL, '/cron/v1/task/group', 'PUT', NULL, 'cron:task:group:update', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (657, 1714462150, 1714462150, 653, '删除任务分组', 'A', NULL, NULL, '/cron/v1/task/group', 'DELETE', NULL, 'cron:task:group:delete', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (658, 1714462150, 1714462150, 653, '查看任务', 'A', NULL, NULL, '/cron/v1/tasks', 'GET', NULL, 'cron:task:query', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (659, 1714462150, 1714462150, 653, '新增任务', 'A', NULL, NULL, '/cron/v1/task', 'POST', NULL, 'cron:task:add', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (660, 1714462150, 1714462150, 653, '立即执行任务', 'A', NULL, NULL, '/cron/v1/task/exec', 'POST', NULL, 'cron:task:exec', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (661, 1714462150, 1714462150, 653, '取消执行任务', 'A', NULL, NULL, '/cron/v1/task/cancel', 'POST', NULL, 'cron:task:cancel', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (662, 1714462150, 1714462150, 653, '修改任务', 'A', NULL, NULL, '/cron/v1/task', 'PUT', NULL, 'cron:task:update', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (663, 1714462150, 1714462150, 653, '删除任务', 'A', NULL, NULL, '/cron/v1/task', 'DELETE', NULL, 'cron:task:delete', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (664, 1714462150, 1714462150, 653, '任务状态管理', 'G', NULL, NULL, NULL, NULL, NULL, 'cron:task:status', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (665, 1714462150, 1714462150, 664, '启用任务', 'A', NULL, NULL, '/cron/v1/task/enable', 'POST', NULL, NULL, NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (666, 1714462150, 1714462150, 664, '禁用任务', 'A', NULL, NULL, '/cron/v1/task/disable', 'POST', NULL, NULL, NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (667, 1714462150, 1714462150, 653, '任务日志', 'G', NULL, NULL, NULL, NULL, NULL, 'cron:task:log', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (668, 1714462150, 1714462150, 667, '获取任务日志分页', 'A', NULL, NULL, '/cron/v1/logs', 'GET', NULL, NULL, NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (669, 1714462150, 1714462150, 667, '获取任务日志详情', 'A', NULL, NULL, '/cron/v1/log', 'GET', NULL, NULL, NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (670, 1714462150, 1714462150, 522, '用户中心', 'M', 'SystemPlatformUserCenter', 'user', NULL, NULL, '/user-center', NULL, NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (671, 1714462150, 1714462150, 670, '协议管理', 'M', 'UserCenterAgreement', 'bookmark', NULL, NULL, '/user-center/agreement', NULL, NULL, NULL, NULL, NULL, 0, 0, 1, NULL, 1);
INSERT INTO `menu` VALUES (672, 1714462150, 1714462150, 671, '协议内容', 'M', 'UserCenterAgreementContent', 'bookmark', NULL, NULL, '/user-center/agreement/content', NULL, NULL, NULL, '/user-center/agreement/content/index', NULL, 0, 0, 1, NULL, 1);
INSERT INTO `menu` VALUES (673, 1714462150, 1714462150, 672, '查看协议列表', 'A', NULL, NULL, '/user-center/v1/agreement/contents', 'GET', NULL, 'uc:agreement:content:query', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (674, 1714462150, 1714462150, 672, '查看协议详细信息', 'A', NULL, NULL, '/user-center/v1/agreement/content', 'GET', NULL, 'uc:agreement:content:query', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (675, 1714462150, 1714462150, 672, '新增协议', 'A', NULL, NULL, '/user-center/v1/agreement/content', 'POST', NULL, 'uc:agreement:content:add', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (676, 1714462150, 1714462150, 672, '修改协议', 'A', NULL, NULL, '/user-center/v1/agreement/content', 'PUT', NULL, 'uc:agreement:content:update', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (677, 1714462150, 1714462150, 672, '删除协议', 'A', NULL, NULL, '/user-center/v1/agreement/content', 'DELETE', NULL, 'uc:agreement:content:delete', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (678, 1714462150, 1714462150, 671, '协议场景', 'M', 'UserCenterAgreementScene', 'common', NULL, NULL, '/user-center/agreement/scene', NULL, NULL, NULL, '/user-center/agreement/scene/index', NULL, 0, 0, 1, NULL, 1);
INSERT INTO `menu` VALUES (679, 1714462150, 1714462150, 678, '查看场景列表', 'A', NULL, NULL, '/user-center/v1/agreement/scenes', 'GET', NULL, 'uc:agreement:scene:query', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (680, 1714462150, 1714462150, 678, '查看场景详细信息', 'A', NULL, NULL, '/user-center/v1/agreement/scene', 'GET', NULL, 'uc:agreement:scene:query', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (681, 1714462150, 1714462150, 678, '新增场景', 'A', NULL, NULL, '/user-center/v1/agreement/scene', 'POST', NULL, 'uc:agreement:scene:add', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (682, 1714462150, 1714462150, 678, '修改场景', 'A', NULL, NULL, '/user-center/v1/agreement/scene', 'PUT', NULL, 'uc:agreement:scene:update', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (683, 1714462150, 1714462150, 678, '删除场景', 'A', NULL, NULL, '/user-center/v1/agreement/scene', 'DELETE', NULL, 'uc:agreement:scene:delete', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (684, 1714462150, 1714462150, 670, '授权渠道', 'M', 'UserCenterChannel', 'mind-mapping', NULL, NULL, '/user-center/channel', NULL, NULL, NULL, '/user-center/channel/index', NULL, 0, 0, 1, NULL, 1);
INSERT INTO `menu` VALUES (685, 1714462150, 1714462150, 684, '查看渠道', 'A', NULL, NULL, '/user-center/v1/channels', 'GET', NULL, 'uc:channel:query', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (686, 1714462150, 1714462150, 684, '新增渠道', 'A', NULL, NULL, '/user-center/v1/channel', 'POST', NULL, 'uc:channel:add', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (687, 1714462150, 1714462150, 684, '修改渠道', 'A', NULL, NULL, '/user-center/v1/channel', 'PUT', NULL, 'uc:channel:update', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (688, 1714462150, 1714462150, 684, '删除渠道', 'A', NULL, NULL, '/user-center/v1/channel', 'DELETE', NULL, 'uc:channel:delete', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (689, 1714462150, 1714462150, 670, '信息字段', 'M', 'UserCenterField', 'list', NULL, NULL, '/user-center/field', NULL, NULL, NULL, '/user-center/field/index', NULL, 0, 0, 1, NULL, 1);
INSERT INTO `menu` VALUES (690, 1714462150, 1714462150, 689, '查看字段', 'A', NULL, NULL, '/user-center/v1/fields', 'GET', NULL, 'uc:field:query', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (691, 1714462150, 1714462150, 689, '新增字段', 'A', NULL, NULL, '/user-center/v1/field', 'POST', NULL, 'uc:field:add', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (692, 1714462150, 1714462150, 689, '修改字段', 'A', NULL, NULL, '/user-center/v1/field', 'PUT', NULL, 'uc:field:update', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (693, 1714462150, 1714462150, 689, '删除字段', 'A', NULL, NULL, '/user-center/v1/field', 'DELETE', NULL, 'uc:field:delete', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (694, 1714462150, 1714462150, 670, '应用管理', 'M', 'UserCenterApp', 'apps', NULL, NULL, '/user-center/app', NULL, NULL, NULL, '/user-center/app/index', NULL, 0, 0, 1, NULL, 1);
INSERT INTO `menu` VALUES (695, 1714462150, 1714462150, 694, '查看应用', 'A', NULL, NULL, '/user-center/v1/apps', 'GET', NULL, 'uc:app:query', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (696, 1714462150, 1714462150, 694, '新增应用', 'A', NULL, NULL, '/user-center/v1/app', 'POST', NULL, 'uc:app:add', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (697, 1714462150, 1714462150, 694, '修改应用', 'A', NULL, NULL, '/user-center/v1/app', 'PUT', NULL, 'uc:app:update', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (698, 1714462150, 1714462150, 694, '删除应用', 'A', NULL, NULL, '/user-center/v1/app', 'DELETE', NULL, 'uc:app:delete', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (699, 1714462150, 1714462150, 670, '用户管理', 'M', 'user-center', 'user', NULL, NULL, '/user-center/user', NULL, NULL, NULL, '/user-center/user/index', NULL, 0, 0, 1, NULL, 1);
INSERT INTO `menu` VALUES (700, 1714462150, 1714462150, 699, '查看用户', 'A', NULL, NULL, '/user-center/v1/users', 'GET', NULL, 'uc:user:query', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (701, 1714462150, 1714462150, 699, '新增用户', 'A', NULL, NULL, '/user-center/v1/user', 'POST', NULL, 'uc:user:add', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (702, 1714462150, 1714462150, 699, '导入用户', 'A', NULL, NULL, '/user-center/v1/users', 'POST', NULL, 'uc:user:import', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (703, 1714462150, 1714462150, 699, '修改用户', 'A', NULL, NULL, '/user-center/v1/user', 'PUT', NULL, 'uc:user:update', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (704, 1714462150, 1714462150, 699, '删除用户', 'A', NULL, NULL, '/user-center/v1/user', 'DELETE', NULL, 'uc:user:delete', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (705, 1714462150, 1714462150, 699, '下线用户', 'A', NULL, NULL, '/user-center/v1/user/offline', 'POST', NULL, 'uc:user:offline', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (706, 1714462150, 1714462150, 699, '用户状态管理', 'G', NULL, NULL, NULL, NULL, NULL, 'uc:user:status', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (707, 1714462150, 1714462150, 706, '禁用用户', 'A', NULL, NULL, '/user-center/v1/user/disable', 'POST', NULL, 'uc:user:disable', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (708, 1714462150, 1714462150, 706, '启用用户', 'A', NULL, NULL, '/user-center/v1/user/enable', 'POST', NULL, 'uc:user:enable', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (709, 1714462150, 1714462150, 0, '应用平台', 'R', 'AppPlatform', 'apps', NULL, NULL, '/app', NULL, NULL, NULL, 'Layout', NULL, 2, 0, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (710, 1714462150, 1714462150, 709, '首页面板', 'M', 'AppDashboard', 'dashboard', NULL, NULL, '/app/dashboard', NULL, NULL, NULL, '/dashboard/index', NULL, 0, 0, 1, 1, 1);
INSERT INTO `menu` VALUES (711, 1714462150, 1714462150, 709, '信号灯系统', 'M', 'PartyAffairs', 'apps', NULL, NULL, '/party-affairs', NULL, NULL, NULL, NULL, NULL, 0, 0, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (712, 1714462150, 1714462150, 711, '轮播管理', 'M', 'PartyAffairsBanner', 'list', NULL, NULL, '/party-affairs/banner', NULL, NULL, NULL, '/party-affairs/banner/index', NULL, 0, NULL, 1, NULL, 1);
INSERT INTO `menu` VALUES (713, 1714462150, 1714462150, 712, '查看轮播', 'A', NULL, NULL, '/party-affairs/v1/banners', 'GET', NULL, 'party-affairs:banner:query', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (714, 1714462150, 1714462150, 712, '新增轮播', 'A', NULL, NULL, '/party-affairs/v1/banner', 'POST', NULL, 'party-affairs:banner:add', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (715, 1714462150, 1714462150, 712, '修改轮播', 'A', NULL, NULL, '/party-affairs/v1/banner', 'PUT', NULL, 'party-affairs:banner:update', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (716, 1714462150, 1714462150, 712, '删除轮播', 'A', NULL, NULL, '/party-affairs/v1/banner', 'DELETE', NULL, 'party-affairs:banner:delete', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (717, 1714462150, 1714462150, 711, '通知管理', 'M', 'PartyAffairsNotice', 'notification', NULL, NULL, '/party-affairs/notice', NULL, NULL, NULL, '/party-affairs/notice/index', NULL, 0, NULL, 1, NULL, 1);
INSERT INTO `menu` VALUES (718, 1714462150, 1714462150, 717, '查看通知', 'A', NULL, NULL, '/party-affairs/v1/notices', 'GET', NULL, 'party-affairs:notice:query', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (719, 1714462150, 1714462150, 717, '查看通知阅读情况', 'A', NULL, NULL, '/party-affairs/v1/notice/users', 'GET', NULL, 'party-affairs:notice:user:query', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (720, 1714462150, 1714462150, 717, '新增通知', 'A', NULL, NULL, '/party-affairs/v1/notice', 'POST', NULL, 'party-affairs:notice:add', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (721, 1714462150, 1714462150, 717, '修改通知', 'A', NULL, NULL, '/party-affairs/v1/notice', 'PUT', NULL, 'party-affairs:notice:update', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (722, 1714462150, 1714462150, 717, '删除通知', 'A', NULL, NULL, '/party-affairs/v1/notice', 'DELETE', NULL, 'party-affairs:notice:delete', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (723, 1714462150, 1714462150, 711, '新闻管理', 'M', 'PartyAffairsNews', 'book', NULL, NULL, '/party-affairs/news', NULL, NULL, NULL, NULL, NULL, 0, 0, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (724, 1714462150, 1714462150, 723, '板块管理', 'M', 'PartyAffairsNewsClassify', 'menu', NULL, NULL, '/party-affairs/news/classify', NULL, NULL, NULL, '/party-affairs/news/classify/index', NULL, 0, NULL, 1, NULL, 1);
INSERT INTO `menu` VALUES (725, 1714462150, 1714462150, 724, '查看板块', 'A', NULL, NULL, '/party-affairs/v1/news/classify', 'GET', NULL, 'party-affairs:news:classify:query', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (726, 1714462150, 1714462150, 724, '新增板块', 'A', NULL, NULL, '/party-affairs/v1/news/classify', 'POST', NULL, 'party-affairs:news:classify:add', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (727, 1714462150, 1714462150, 724, '修改板块', 'A', NULL, NULL, '/party-affairs/v1/news/classify', 'PUT', NULL, 'party-affairs:news:classify:update', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (728, 1714462150, 1714462150, 724, '删除板块', 'A', NULL, NULL, '/party-affairs/v1/news/classify', 'DELETE', NULL, 'party-affairs:news:classify:delete', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (729, 1714462150, 1714462150, 723, '内容管理', 'M', 'PartyAffairsNewsContent', 'book', NULL, NULL, '/party-affairs/news/content', NULL, NULL, NULL, '/party-affairs/news/content/index', NULL, 0, NULL, 1, NULL, 1);
INSERT INTO `menu` VALUES (730, 1714462150, 1714462150, 729, '查看内容列表', 'A', NULL, NULL, '/party-affairs/v1/news/contents', 'GET', NULL, 'party-affairs:news:content:query', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (731, 1714462150, 1714462150, 729, '查看指定内容', 'A', NULL, NULL, '/party-affairs/v1/news/content', 'GET', NULL, 'party-affairs:news:content:info', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (732, 1714462150, 1714462150, 729, '新增内容', 'A', NULL, NULL, '/party-affairs/v1/news/content', 'POST', NULL, 'party-affairs:news:content:add', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (733, 1714462150, 1714462150, 729, '修改内容', 'A', NULL, NULL, '/party-affairs/v1/news/content', 'PUT', NULL, 'party-affairs:news:content:update', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (734, 1714462150, 1714462150, 729, '删除内容', 'A', NULL, NULL, '/party-affairs/v1/news/content', 'DELETE', NULL, 'party-affairs:news:content:delete', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (735, 1714462150, 1714462150, 723, '评论管理', 'M', 'PartyAffairsNewsComment', 'book', NULL, NULL, '/party-affairs/news/comment', NULL, NULL, NULL, '/party-affairs/news/comment/index', NULL, 0, 1, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (736, 1714462150, 1714462150, 735, '查看评论列表', 'A', NULL, NULL, '/party-affairs/v1/news/comments', 'GET', NULL, 'party-affairs:news:comment:query', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (737, 1714462150, 1714462150, 735, '删除评论', 'A', NULL, NULL, '/party-affairs/v1/news/comment', 'DELETE', NULL, 'party-affairs:news:comment:delete', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (738, 1714462150, 1714462150, 711, '资料管理', 'M', 'PartyAffairsResource', 'public', NULL, NULL, '/party-affairs/resource', NULL, NULL, NULL, NULL, NULL, 0, 0, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (739, 1714462150, 1714462150, 738, '板块管理', 'M', 'PartyAffairsResourceClassify', 'menu', NULL, NULL, '/party-affairs/resource/classify', NULL, NULL, NULL, '/party-affairs/resource/classify/index', NULL, 0, NULL, 1, NULL, 1);
INSERT INTO `menu` VALUES (740, 1714462150, 1714462150, 739, '查看板块', 'A', NULL, NULL, '/party-affairs/v1/resource/classify', 'GET', NULL, 'party-affairs:resource:classify:query', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (741, 1714462150, 1714462150, 739, '新增板块', 'A', NULL, NULL, '/party-affairs/v1/resource/classify', 'POST', NULL, 'party-affairs:resource:classify:add', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (742, 1714462150, 1714462150, 739, '修改板块', 'A', NULL, NULL, '/party-affairs/v1/resource/classify', 'PUT', NULL, 'party-affairs:resource:classify:update', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (743, 1714462150, 1714462150, 739, '删除板块', 'A', NULL, NULL, '/party-affairs/v1/resource/classify', 'DELETE', NULL, 'party-affairs:resource:classify:delete', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (744, 1714462150, 1714462150, 738, '内容管理', 'M', 'PartyAffairsResourceContent', 'public', NULL, NULL, '/party-affairs/resource/content', NULL, NULL, NULL, '/party-affairs/resource/content/index', NULL, 0, NULL, 1, NULL, 1);
INSERT INTO `menu` VALUES (745, 1714462150, 1714462150, 744, '查看内容列表', 'A', NULL, NULL, '/party-affairs/v1/resource/contents', 'GET', NULL, 'party-affairs:resource:content:query', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (746, 1714462150, 1714462150, 744, '新增内容', 'A', NULL, NULL, '/party-affairs/v1/resource/content', 'POST', NULL, 'party-affairs:resource:content:add', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (747, 1714462150, 1714462150, 744, '修改内容', 'A', NULL, NULL, '/party-affairs/v1/resource/content', 'PUT', NULL, 'party-affairs:resource:content:update', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (748, 1714462150, 1714462150, 744, '删除内容', 'A', NULL, NULL, '/party-affairs/v1/resource/content', 'DELETE', NULL, 'party-affairs:resource:content:delete', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (749, 1714462150, 1714462150, 711, '任务管理', 'M', 'PartyAffairsTask', 'layers', NULL, NULL, '/party-affairs/task', NULL, NULL, NULL, NULL, NULL, 0, 0, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (750, 1714462150, 1714462150, 749, '任务配置', 'M', 'PartyAffairsTaskConfigure', 'unordered-list', NULL, NULL, '/party-affairs/task/configure', NULL, NULL, NULL, '/party-affairs/task/configure/index', NULL, 0, NULL, 1, NULL, 1);
INSERT INTO `menu` VALUES (751, 1714462150, 1714462150, 750, '查看配置', 'A', NULL, NULL, '/party-affairs/v1/task/configure', 'GET', NULL, 'party-affairs:task:configure:query', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (752, 1714462150, 1714462150, 750, '新增配置', 'A', NULL, NULL, '/party-affairs/v1/task/configure', 'POST', NULL, 'party-affairs:task:configure:add', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (753, 1714462150, 1714462150, 750, '修改配置', 'A', NULL, NULL, '/party-affairs/v1/task/configure', 'PUT', NULL, 'party-affairs:task:configure:update', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (754, 1714462150, 1714462150, 750, '删除配置', 'A', NULL, NULL, '/party-affairs/v1/task/configure', 'DELETE', NULL, 'party-affairs:task:configure:delete', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (755, 1714462150, 1714462150, 749, '任务统计', 'M', 'PartyAffairsTaskContent', 'computer', NULL, NULL, '/party-affairs/task/value', NULL, NULL, NULL, '/party-affairs/task/value/index', NULL, 0, NULL, 1, NULL, 1);
INSERT INTO `menu` VALUES (756, 1714462150, 1714462150, 755, '查看任务填写列表', 'A', NULL, NULL, '/party-affairs/v1/task/values', 'GET', NULL, 'party-affairs:task:value:query', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (757, 1714462150, 1714462150, 755, '查看指定任务填写详情', 'A', NULL, NULL, '/party-affairs/v1/task/value', 'GET', NULL, 'party-affairs:task:value:query', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (758, 1714462150, 1714462150, 755, '修改任务填写详情', 'A', NULL, NULL, '/party-affairs/v1/task/value', 'PUT', NULL, 'party-affairs:task:value:update', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (759, 1714462150, 1714462150, 755, '删除任务填写详情', 'A', NULL, NULL, '/party-affairs/v1/task/value', 'DELETE', NULL, 'party-affairs:task:value:delete', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (760, 1714462150, 1714462150, 711, '视频管理', 'M', 'PartyAffairsVideo', 'video-camera', NULL, NULL, '/party-affairs/video', NULL, NULL, NULL, NULL, NULL, 0, 0, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (761, 1714462150, 1714462150, 760, '板块管理', 'M', 'PartyAffairsVideoClassify', 'menu', NULL, NULL, '/party-affairs/video/classify', NULL, NULL, NULL, '/party-affairs/video/classify/index', NULL, 0, NULL, 1, NULL, 1);
INSERT INTO `menu` VALUES (762, 1714462150, 1714462150, 761, '查看板块', 'A', NULL, NULL, '/party-affairs/v1/video/classify', 'GET', NULL, 'party-affairs:video:classify:query', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (763, 1714462150, 1714462150, 761, '新增板块', 'A', NULL, NULL, '/party-affairs/v1/video/classify', 'POST', NULL, 'party-affairs:video:classify:add', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (764, 1714462150, 1714462150, 761, '修改板块', 'A', NULL, NULL, '/party-affairs/v1/video/classify', 'PUT', NULL, 'party-affairs:video:classify:update', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (765, 1714462150, 1714462150, 761, '删除板块', 'A', NULL, NULL, '/party-affairs/v1/video/classify', 'DELETE', NULL, 'party-affairs:video:classify:delete', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (766, 1714462150, 1714462150, 760, '内容管理', 'M', 'PartyAffairsVideoContent', 'video-camera', NULL, NULL, '/party-affairs/video/content', NULL, NULL, NULL, '/party-affairs/video/content/index', NULL, 0, NULL, 1, NULL, 1);
INSERT INTO `menu` VALUES (767, 1714462150, 1714462150, 766, '查看内容列表', 'A', NULL, NULL, '/party-affairs/v1/video/contents', 'GET', NULL, 'party-affairs:video:content:query', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (768, 1714462150, 1714462150, 766, '新增内容', 'A', NULL, NULL, '/party-affairs/v1/video/content', 'POST', NULL, 'party-affairs:video:content:add', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (769, 1714462150, 1714462150, 766, '修改内容', 'A', NULL, NULL, '/party-affairs/v1/video/content', 'PUT', NULL, 'party-affairs:video:content:update', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (770, 1714462150, 1714462150, 766, '删除内容', 'A', NULL, NULL, '/party-affairs/v1/video/content', 'DELETE', NULL, 'party-affairs:video:content:delete', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (771, 1714462150, 1714462150, 766, '进度统计', 'A', NULL, NULL, '/party-affairs/v1/video/process', 'GET', NULL, 'party-affairs:video:content:process', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
COMMIT;

-- ----------------------------
-- Table structure for menu_copy1
-- ----------------------------
DROP TABLE IF EXISTS `menu_copy1`;
CREATE TABLE `menu_copy1` (
                              `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT '主键ID',
                              `created_at` bigint DEFAULT NULL COMMENT '创建时间',
                              `updated_at` bigint DEFAULT NULL COMMENT '修改时间',
                              `parent_id` int unsigned NOT NULL COMMENT '父id',
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
) ENGINE=InnoDB AUTO_INCREMENT=772 DEFAULT CHARSET=utf8mb4 COMMENT='菜单信息';

-- ----------------------------
-- Records of menu_copy1
-- ----------------------------
BEGIN;
INSERT INTO `menu_copy1` VALUES (522, 1714462150, 1714462150, 0, '管理平台', 'R', 'SystemPlatform', 'apps', NULL, NULL, '/', NULL, NULL, NULL, 'Layout', NULL, 2, 0, NULL, NULL, NULL);
INSERT INTO `menu_copy1` VALUES (523, 1714462150, 1714462150, 522, '首页面板', 'M', 'Dashboard', 'dashboard', NULL, NULL, '/dashboard', NULL, NULL, NULL, '/dashboard/index', NULL, 0, 0, 1, 1, 1);
INSERT INTO `menu_copy1` VALUES (524, 1714462150, 1714462150, 522, '管理中心', 'M', 'SystemPlatformManager', 'apps', NULL, NULL, '/manager', NULL, NULL, NULL, NULL, NULL, 0, 0, NULL, NULL, NULL);
INSERT INTO `menu_copy1` VALUES (525, 1714462150, 1714462150, 524, '基础接口', 'G', 'BaseApi', 'apps', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, 0, 1, NULL, NULL, NULL);
INSERT INTO `menu_copy1` VALUES (526, 1714462150, 1714462150, 525, '获取用户部门', 'BA', NULL, NULL, '/manager/v1/department/tree', 'GET', NULL, NULL, NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu_copy1` VALUES (527, 1714462150, 1714462150, 525, '获取个人用户信息', 'BA', NULL, NULL, '/manager/v1/user/current', 'GET', NULL, NULL, NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu_copy1` VALUES (528, 1714462150, 1714462150, 525, '获取用户可见角色树', 'BA', NULL, NULL, '/manager/v1/role/tree', 'GET', NULL, NULL, NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu_copy1` VALUES (529, 1714462150, 1714462150, 525, '获取用户菜单', 'BA', NULL, NULL, '/manager/v1/menu/tree/from/role', 'GET', NULL, NULL, NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu_copy1` VALUES (530, 1714462150, 1714462150, 525, '退出登录', 'BA', NULL, NULL, '/manager/v1/user/logout', 'POST', NULL, NULL, NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu_copy1` VALUES (531, 1714462150, 1714462150, 525, '刷新token', 'BA', NULL, NULL, '/manager/v1/user/token/refresh', 'POST', NULL, NULL, NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu_copy1` VALUES (532, 1714462150, 1714462150, 525, '用户登录', 'BA', NULL, NULL, '/manager/v1/user/login', 'POST', NULL, NULL, NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu_copy1` VALUES (533, 1714462150, 1714462150, 525, '获取登录验证码', 'BA', NULL, NULL, '/manager/v1/user/login/captcha', 'POST', NULL, NULL, NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu_copy1` VALUES (534, 1714462150, 1714462150, 525, '获取系统基础设置', 'BA', NULL, NULL, '/manager/v1/setting', 'GET', NULL, NULL, NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu_copy1` VALUES (535, 1714462150, 1714462150, 525, '接口鉴权', 'BA', NULL, NULL, '/manager/v1/auth', 'POST', NULL, NULL, NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu_copy1` VALUES (536, 1714462150, 1714462150, 525, '切换用户角色', 'BA', NULL, NULL, '/manager/v1/user/role/switch', 'POST', NULL, NULL, NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu_copy1` VALUES (537, 1714462150, 1714462150, 525, '新增部门资源对象', 'BA', NULL, NULL, '/manager/v1/department/object', 'POST', NULL, NULL, NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu_copy1` VALUES (538, 1714462150, 1714462150, 525, '删除部门资源对象', 'BA', NULL, NULL, '/manager/v1/department/object', 'DELETE', NULL, NULL, NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu_copy1` VALUES (539, 1714462150, 1714462150, 524, '资源对象', 'M', 'Object', 'common', NULL, NULL, '/manager/object', NULL, NULL, NULL, '/manager/object/index', NULL, 0, NULL, 1, NULL, 1);
INSERT INTO `menu_copy1` VALUES (540, 1714462150, 1714462150, 539, '查询对象', 'A', NULL, NULL, '/manager/v1/objects', 'GET', NULL, 'manager:object:query', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu_copy1` VALUES (541, 1714462150, 1714462150, 539, '新增对象', 'A', NULL, NULL, '/manager/v1/object', 'POST', NULL, 'manager:object:add', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu_copy1` VALUES (542, 1714462150, 1714462150, 539, '修改对象', 'A', NULL, NULL, '/manager/v1/object', 'PUT', NULL, 'manager:object:update', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu_copy1` VALUES (543, 1714462150, 1714462150, 539, '删除对象', 'A', NULL, NULL, '/manager/v1/object', 'DELETE', NULL, 'manager:object:delete', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu_copy1` VALUES (544, 1714462150, 1714462150, 524, '字典管理', 'M', 'Dict', 'storage', NULL, NULL, '/manager/dictionary', NULL, NULL, NULL, '/manager/dictionary/index', NULL, 0, NULL, 1, NULL, 1);
INSERT INTO `menu_copy1` VALUES (545, 1714462150, 1714462150, 544, '查询字典', 'A', NULL, NULL, '/manager/v1/dictionaries', 'GET', NULL, 'manager:dictionary:query', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu_copy1` VALUES (546, 1714462150, 1714462150, 544, '新增字典', 'A', NULL, NULL, '/manager/v1/dictionary', 'POST', NULL, 'manager:dictionary:add', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu_copy1` VALUES (547, 1714462150, 1714462150, 544, '修改字典', 'A', NULL, NULL, '/manager/v1/dictionary', 'PUT', NULL, 'manager:dictionary:update', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu_copy1` VALUES (548, 1714462150, 1714462150, 544, '删除字典', 'A', NULL, NULL, '/manager/v1/dictionary', 'DELETE', NULL, 'manager:dictionary:delete', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu_copy1` VALUES (549, 1714462150, 1714462150, 544, '刷新字典值', 'A', NULL, NULL, '/manager/v1/dictionary/values', 'GET', NULL, 'manager:dictionary:value:query', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu_copy1` VALUES (550, 1714462150, 1714462150, 544, '新增字典值', 'A', NULL, NULL, '/manager/v1/dictionary/value', 'POST', NULL, 'manager:dictionary:value:add', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu_copy1` VALUES (551, 1714462150, 1714462150, 544, '修改字典值', 'A', NULL, NULL, '/manager/v1/dictionary/value', 'PUT', NULL, 'manager:dictionary:value:update', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu_copy1` VALUES (552, 1714462150, 1714462150, 544, '删除字典值', 'A', NULL, NULL, '/manager/v1/dictionary/value', 'DELETE', NULL, 'manager:dictionary:value:delete', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu_copy1` VALUES (553, 1714462150, 1714462150, 524, '菜单管理', 'M', 'ManagerMenu', 'menu', NULL, NULL, '/manager/menu', NULL, NULL, NULL, '/manager/menu/index', NULL, 0, 0, 1, 1, 1);
INSERT INTO `menu_copy1` VALUES (554, 1714462150, 1714462150, 553, '查询菜单', 'A', NULL, NULL, '/manager/v1/menu/tree', 'GET', NULL, 'manager:menu:query', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu_copy1` VALUES (555, 1714462150, 1714462150, 553, '新增菜单', 'A', NULL, NULL, '/manager/v1/menu', 'POST', NULL, 'manager:menu:add', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu_copy1` VALUES (556, 1714462150, 1714462150, 553, '修改菜单', 'A', NULL, NULL, '/manager/v1/menu', 'PUT', NULL, 'manager:menu:update', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu_copy1` VALUES (557, 1714462150, 1714462150, 553, '删除菜单', 'A', NULL, NULL, '/manager/v1/menu', 'DELETE', NULL, 'manager:menu:delete', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu_copy1` VALUES (558, 1714462150, 1714462150, 524, '职位管理', 'M', 'JobUser', 'tag', NULL, NULL, '/manager/job', NULL, NULL, NULL, '/manager/job/index', NULL, 0, NULL, 1, NULL, 1);
INSERT INTO `menu_copy1` VALUES (559, 1714462150, 1714462150, 558, '查询职位', 'A', NULL, NULL, '/manager/v1/jobs', 'GET', NULL, 'manager:job:query', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu_copy1` VALUES (560, 1714462150, 1714462150, 558, '新增职位', 'A', NULL, NULL, '/manager/v1/job', 'POST', NULL, 'manager:job:add', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu_copy1` VALUES (561, 1714462150, 1714462150, 558, '修改职位', 'A', NULL, NULL, '/manager/v1/job', 'PUT', NULL, 'manager:job:update', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu_copy1` VALUES (562, 1714462150, 1714462150, 558, '删除职位', 'A', NULL, NULL, '/manager/v1/job', 'DELETE', NULL, 'manager:job:delete', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu_copy1` VALUES (563, 1714462150, 1714462150, 524, '部门管理', 'M', 'ManagerDepartment', 'user-group', NULL, NULL, '/manager/department', NULL, NULL, NULL, '/manager/department/index', NULL, 0, 0, 1, NULL, 1);
INSERT INTO `menu_copy1` VALUES (564, 1714462150, 1714462150, 563, '查看部门', 'A', NULL, NULL, '/manager/v1/department/tree', 'GET', NULL, 'manager:department:query', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu_copy1` VALUES (565, 1714462150, 1714462150, 563, '新增部门', 'A', NULL, NULL, '/manager/v1/department', 'POST', NULL, 'manager:department:add', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu_copy1` VALUES (566, 1714462150, 1714462150, 563, '修改部门', 'A', NULL, NULL, '/manager/v1/department', 'PUT', NULL, 'manager:department:update', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu_copy1` VALUES (567, 1714462150, 1714462150, 563, '删除部门', 'A', NULL, NULL, '/manager/v1/department', 'DELETE', NULL, 'manager:department:delete', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu_copy1` VALUES (568, 1714462150, 1714462150, 563, '查看部门资源对象', 'A', NULL, NULL, '/manager/v1/department/objects', 'GET', NULL, 'manager:department:object:query', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu_copy1` VALUES (569, 1714462150, 1714462150, 563, '设置部门资源对象', 'A', NULL, NULL, '/manager/v1/department/objects', 'POST', NULL, 'manager:department:object:import', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu_copy1` VALUES (570, 1714462150, 1714462150, 524, '角色管理', 'M', 'ManagerRole', 'safe', NULL, NULL, '/manager/role', NULL, NULL, NULL, '/manager/role/index', NULL, 0, NULL, 0, NULL, 1);
INSERT INTO `menu_copy1` VALUES (571, 1714462150, 1714462150, 570, '查询角色', 'A', NULL, NULL, '/manager/v1/role/tree', 'GET', NULL, 'manager:role:query', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu_copy1` VALUES (572, 1714462150, 1714462150, 570, '新增角色', 'A', NULL, NULL, '/manager/v1/role', 'POST', NULL, 'manager:role:add', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu_copy1` VALUES (573, 1714462150, 1714462150, 570, '修改角色', 'A', NULL, NULL, '/manager/v1/role', 'PUT', NULL, 'manager:role:update', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu_copy1` VALUES (574, 1714462150, 1714462150, 570, '删除角色', 'A', NULL, NULL, '/manager/v1/role', 'DELETE', NULL, 'manager:role:delete', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu_copy1` VALUES (575, 1714462150, 1714462150, 570, '角色菜单管理', 'G', NULL, NULL, '/manager/v1/role/menu/tree', 'GET', NULL, 'manager:role:menu', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu_copy1` VALUES (576, 1714462150, 1714462150, 575, '查询角色菜单', 'A', NULL, NULL, '/manager/v1/role/menu/ids', 'GET', NULL, 'manager:role:menu:query', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu_copy1` VALUES (577, 1714462150, 1714462150, 575, '修改角色菜单', 'A', NULL, NULL, '/manager/v1/role/menu', 'POST', NULL, 'manager:role:menu:update', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu_copy1` VALUES (578, 1714462150, 1714462150, 524, '用户管理', 'M', 'ManagerUser', 'user', NULL, NULL, '/manager/user', NULL, NULL, NULL, '/manager/user/index', NULL, 0, NULL, 1, NULL, 1);
INSERT INTO `menu_copy1` VALUES (579, 1714462150, 1714462150, 578, '查询用户', 'A', NULL, NULL, '/manager/v1/users', 'GET', NULL, 'manager:user:query', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu_copy1` VALUES (580, 1714462150, 1714462150, 578, '新增用户', 'A', NULL, NULL, '/manager/v1/user', 'POST', NULL, 'manager:user:add', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu_copy1` VALUES (581, 1714462150, 1714462150, 578, '修改用户', 'A', NULL, NULL, '/manager/v1/user', 'PUT', NULL, 'manager:user:update', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu_copy1` VALUES (582, 1714462150, 1714462150, 581, '获取用户角色列表', 'A', NULL, NULL, '/manager/v1/user/roles', 'GET', NULL, 'manager:user:update', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu_copy1` VALUES (583, 1714462150, 1714462150, 581, '获取用户部门列表', 'A', NULL, NULL, '/manager/v1/user/jobs', 'POST', NULL, 'manager:user:update', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu_copy1` VALUES (584, 1714462150, 1714462150, 578, '删除用户', 'A', NULL, NULL, '/manager/v1/user', 'DELETE', NULL, 'manager:user:delete', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu_copy1` VALUES (585, 1714462150, 1714462150, 578, '启用/禁用用户', 'G', NULL, NULL, NULL, NULL, NULL, 'manager:user:status', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu_copy1` VALUES (586, 1714462150, 1714462150, 585, '启用用户', 'A', NULL, NULL, '/manager/v1/user/enable', 'POST', NULL, 'manager:user:status:enable', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu_copy1` VALUES (587, 1714462150, 1714462150, 585, '禁用用户', 'A', NULL, NULL, '/manager/v1/user/disable', 'POST', NULL, 'manager:user:status:disable', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu_copy1` VALUES (588, 1714462150, 1714462150, 578, '重置账号密码', 'A', NULL, NULL, '/manager/v1/user/password/reset', 'POST', NULL, 'manager:user:reset:password', NULL, NULL, '', NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu_copy1` VALUES (589, 1714462150, 1714462150, 522, '配置中心', 'M', 'SystemPlatformConfigure', 'code-block', NULL, NULL, '/configure', NULL, NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu_copy1` VALUES (590, 1714462150, 1714462150, 589, '环境管理', 'M', 'ConfigureEnv', 'common', NULL, NULL, '/configure/env', NULL, NULL, NULL, '/configure/env/index', NULL, 0, NULL, 1, NULL, 1);
INSERT INTO `menu_copy1` VALUES (591, 1714462150, 1714462150, 590, '查看环境', 'A', NULL, NULL, '/configure/v1/envs', 'GET', NULL, 'configure:env:query', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu_copy1` VALUES (592, 1714462150, 1714462150, 590, '新增环境', 'A', NULL, NULL, '/configure/v1/env', 'POST', NULL, 'configure:env:add', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu_copy1` VALUES (593, 1714462150, 1714462150, 590, '修改环境', 'A', NULL, NULL, '/configure/v1/env', 'PUT', NULL, 'configure:env:update', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu_copy1` VALUES (594, 1714462150, 1714462150, 590, '删除环境', 'A', NULL, NULL, '/configure/v1/env', 'DELETE', NULL, 'configure:env:delete', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu_copy1` VALUES (595, 1714462150, 1714462150, 590, '查看环境Token', 'A', NULL, NULL, '/configure/v1/env/token', 'GET', NULL, 'configure:env:token:query', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu_copy1` VALUES (596, 1714462150, 1714462150, 590, '重置环境token', 'A', NULL, NULL, '/configure/v1/env/token', 'PUT', NULL, 'configure:env:token:reset', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu_copy1` VALUES (597, 1714462150, 1714462150, 589, '服务管理', 'M', 'ConfigureServer', 'apps', NULL, NULL, '/configure/server', NULL, NULL, NULL, '/configure/server/index', NULL, 0, NULL, 1, NULL, 1);
INSERT INTO `menu_copy1` VALUES (598, 1714462150, 1714462150, 597, '查询服务', 'A', NULL, NULL, '/configure/v1/servers', 'GET', NULL, 'configure:server:query', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu_copy1` VALUES (599, 1714462150, 1714462150, 597, '新增服务', 'A', NULL, NULL, '/configure/v1/server', 'POST', NULL, 'configure:server:add', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu_copy1` VALUES (600, 1714462150, 1714462150, 597, '修改服务', 'A', NULL, NULL, '/configure/v1/server', 'PUT', NULL, 'configure:server:update', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu_copy1` VALUES (601, 1714462150, 1714462150, 597, '删除服务', 'A', NULL, NULL, '/configure/v1/server', 'DELETE', NULL, 'configure:server:delete', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu_copy1` VALUES (602, 1714462150, 1714462150, 589, '资源变量', 'M', 'ConfigureResource', 'unordered-list', NULL, NULL, '/configure/resource', NULL, NULL, NULL, '/configure/resource/index', NULL, 0, NULL, 1, NULL, 1);
INSERT INTO `menu_copy1` VALUES (603, 1714462150, 1714462150, 602, '查看资源', 'A', NULL, NULL, '/configure/v1/resources', 'GET', NULL, 'configure:resource:query', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu_copy1` VALUES (604, 1714462150, 1714462150, 602, '新增资源', 'A', NULL, NULL, '/configure/v1/resource', 'POST', NULL, 'configure:resource:add', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu_copy1` VALUES (605, 1714462150, 1714462150, 602, '修改资源', 'A', NULL, NULL, '/configure/v1/resource', 'PUT', NULL, 'configure:resource:update', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu_copy1` VALUES (606, 1714462150, 1714462150, 602, '删除资源', 'A', NULL, NULL, '/configure/v1/resource', 'DELETE', NULL, 'configure:resource:delete', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu_copy1` VALUES (607, 1714462150, 1714462150, 602, '资源变量值配置', 'G', NULL, NULL, NULL, NULL, NULL, 'configure:resource:value', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu_copy1` VALUES (608, 1714462150, 1714462150, 589, '业务变量', 'M', 'ConfigureBusiness', 'code', NULL, NULL, '/configure/business', NULL, NULL, NULL, '/configure/business/index', NULL, 0, NULL, 1, NULL, 1);
INSERT INTO `menu_copy1` VALUES (609, 1714462150, 1714462150, 608, '查看业务变量', 'A', NULL, NULL, '/configure/v1/business', 'GET', NULL, 'configure:business:query', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu_copy1` VALUES (610, 1714462150, 1714462150, 608, '新增业务变量', 'A', NULL, NULL, '/configure/v1/business', 'POST', NULL, 'configure:business:add', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu_copy1` VALUES (611, 1714462150, 1714462150, 608, '修改业务变量', 'A', NULL, NULL, '/configure/v1/business', 'PUT', NULL, 'configure:business:update', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu_copy1` VALUES (612, 1714462150, 1714462150, 608, '删除业务变量', 'A', NULL, NULL, '/configure/v1/business', 'DELETE', NULL, 'configure:business:delete', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu_copy1` VALUES (613, 1714462150, 1714462150, 608, '配置业务变量值', 'A', NULL, NULL, '/configure/business/value', 'PUT', NULL, 'configure:business:value', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu_copy1` VALUES (614, 1714462150, 1714462150, 589, '配置模板', 'M', 'ConfgureTemplate', 'file', NULL, NULL, '/configure/template', NULL, NULL, NULL, '/configure/template/index', NULL, 0, NULL, 1, NULL, 1);
INSERT INTO `menu_copy1` VALUES (615, 1714462150, 1714462150, 614, '模板管理', 'G', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu_copy1` VALUES (616, 1714462150, 1714462150, 615, '查看模板历史版本', 'A', NULL, NULL, '/configure/v1/templates', 'GET', NULL, 'configure:template:history', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu_copy1` VALUES (617, 1714462150, 1714462150, 615, '查看指定模板详细数据', 'A', NULL, NULL, '/configure/v1/template', 'GET', NULL, NULL, NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu_copy1` VALUES (618, 1714462150, 1714462150, 615, '查看当前正在使用的模板', 'A', NULL, NULL, '/configure/v1/template/current', 'GET', NULL, NULL, NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu_copy1` VALUES (619, 1714462150, 1714462150, 615, '提交模板', 'A', NULL, NULL, '/configure/v1/template', 'POST', NULL, 'configure:template:add', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu_copy1` VALUES (620, 1714462150, 1714462150, 615, '模板对比', 'A', NULL, NULL, '/configure/v1/template/compare', 'POST', NULL, 'configure:template:compare', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu_copy1` VALUES (621, 1714462150, 1714462150, 615, '切换模板', 'A', NULL, NULL, '/configure/v1/template/switch', 'POST', NULL, 'configure:template:switch', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu_copy1` VALUES (622, 1714462150, 1714462150, 615, '模板预览', 'A', NULL, NULL, '/configure/v1/template/preview', 'POST', NULL, 'configure:template:preview', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu_copy1` VALUES (623, 1714462150, 1714462150, 614, '配置管理', 'G', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu_copy1` VALUES (624, 1714462150, 1714462150, 623, '配置对比', 'A', NULL, NULL, '/configure/v1/configure/compare', 'POST', NULL, 'configure:configure:compare', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu_copy1` VALUES (625, 1714462150, 1714462150, 623, '同步配置', 'A', NULL, NULL, '/configure/v1/configure', 'PUT', NULL, 'configure:configure:sync', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu_copy1` VALUES (626, 1714462150, 1714462150, 522, '资源中心', 'M', 'SystemPlatformResource', 'file', NULL, NULL, '/resource', NULL, NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu_copy1` VALUES (627, 1714462150, 1714462150, 626, '文件管理', 'M', 'File', 'file', NULL, NULL, '/resource/file', NULL, NULL, NULL, '/resource/file/index', NULL, 0, 0, 1, NULL, 1);
INSERT INTO `menu_copy1` VALUES (628, 1714462150, 1714462150, 627, '目录管理', 'G', NULL, NULL, NULL, NULL, NULL, 'resource:directory:group', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu_copy1` VALUES (629, 1714462150, 1714462150, 628, '查看目录', 'A', NULL, NULL, '/resource/v1/directories', 'GET', NULL, 'resource:directory:query', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu_copy1` VALUES (630, 1714462150, 1714462150, 628, '新增目录', 'A', NULL, NULL, '/resource/v1/directory', 'POST', NULL, 'resource:directory:add', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu_copy1` VALUES (631, 1714462150, 1714462150, 628, '修改目录', 'A', NULL, NULL, '/resource/v1/directory', 'PUT', NULL, 'resource:directory:update', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu_copy1` VALUES (632, 1714462150, 1714462150, 628, '删除目录', 'A', NULL, NULL, '/resource/v1/directory', 'DELETE', NULL, 'resource:directory:delete', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu_copy1` VALUES (633, 1714462150, 1714462150, 627, '文件管理', 'G', NULL, NULL, NULL, NULL, NULL, 'resource:file:group', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu_copy1` VALUES (634, 1714462150, 1714462150, 633, '查看文件', 'A', NULL, NULL, '/resource/v1/files', 'GET', NULL, 'resource:file:query', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu_copy1` VALUES (635, 1714462150, 1714462150, 633, '修改文件', 'A', NULL, NULL, '/resource/v1/file', 'PUT', NULL, 'resource:file:upload', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu_copy1` VALUES (636, 1714462150, 1714462150, 633, '删除文件', 'A', NULL, NULL, '/resource/v1/file', 'DELETE', NULL, 'resource:file:delete', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu_copy1` VALUES (637, 1714462150, 1714462150, 633, '上传文件', 'G', NULL, NULL, NULL, NULL, NULL, 'resource:file:upload:group', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu_copy1` VALUES (638, 1714462150, 1714462150, 637, '分块上传文件', 'A', NULL, NULL, '/resource/v1/upload', 'POST', NULL, 'resource:file:upload', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu_copy1` VALUES (639, 1714462150, 1714462150, 637, '预上传文件', 'A', NULL, NULL, '/resource/v1/upload/prepare', 'POST', NULL, 'resource:file:upload', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu_copy1` VALUES (640, 1714462150, 1714462150, 522, '定时任务', 'M', 'SystemPlatformCron', 'schedule', NULL, NULL, '/cron', NULL, NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu_copy1` VALUES (641, 1714462150, 1714462150, 640, '节点管理', 'M', 'Worker', 'common', NULL, NULL, '/cron/worker', NULL, NULL, NULL, '/cron/worker/index', NULL, 0, 0, 1, NULL, 1);
INSERT INTO `menu_copy1` VALUES (642, 1714462150, 1714462150, 641, '查看节点分组', 'A', NULL, NULL, '/cron/v1/worker/groups', 'GET', NULL, 'cron:worker:group:query', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu_copy1` VALUES (643, 1714462150, 1714462150, 641, '新增节点分组', 'A', NULL, NULL, '/cron/v1/worker/group', 'POST', NULL, 'cron:worker:group:add', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu_copy1` VALUES (644, 1714462150, 1714462150, 641, '修改节点分组', 'A', NULL, NULL, '/cron/v1/worker/group', 'PUT', NULL, 'cron:worker:group:update', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu_copy1` VALUES (645, 1714462150, 1714462150, 641, '删除节点分组', 'A', NULL, NULL, '/cron/v1/worker/group', 'DELETE', NULL, 'cron:worker:group:delete', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu_copy1` VALUES (646, 1714462150, 1714462150, 641, '查看节点', 'A', NULL, NULL, '/cron/v1/workers', 'GET', NULL, 'cron:worker:query', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu_copy1` VALUES (647, 1714462150, 1714462150, 641, '新增节点', 'A', NULL, NULL, '/cron/v1/worker', 'POST', NULL, 'cron:worker:add', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu_copy1` VALUES (648, 1714462150, 1714462150, 641, '修改节点', 'A', NULL, NULL, '/cron/v1/worker', 'PUT', NULL, 'cron:worker:update', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu_copy1` VALUES (649, 1714462150, 1714462150, 641, '删除节点', 'A', NULL, NULL, '/cron/v1/worker', 'DELETE', NULL, 'cron:worker:delete', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu_copy1` VALUES (650, 1714462150, 1714462150, 641, '节点状态管理', 'G', NULL, NULL, NULL, NULL, NULL, 'cron:worker:status', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu_copy1` VALUES (651, 1714462150, 1714462150, 650, '启用节点', 'A', NULL, NULL, '/cron/v1/worker/enable', 'POST', NULL, 'cron:worker:status:enable', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu_copy1` VALUES (652, 1714462150, 1714462150, 650, '启用节点', 'A', NULL, NULL, '/cron/v1/worker/disable', 'POST', NULL, 'cron:worker:status:disable', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu_copy1` VALUES (653, 1714462150, 1714462150, 640, '任务管理', 'M', 'Task', 'computer', NULL, NULL, '/cron/task', NULL, NULL, NULL, '/cron/task/index', NULL, 0, 0, 1, NULL, 1);
INSERT INTO `menu_copy1` VALUES (654, 1714462150, 1714462150, 653, '查看任务分组', 'A', NULL, NULL, '/cron/v1/task/groups', 'GET', NULL, 'cron:task:group:query', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu_copy1` VALUES (655, 1714462150, 1714462150, 653, '新增任务分组', 'A', NULL, NULL, '/cron/v1/task/group', 'POST', NULL, 'cron:task:group:add', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu_copy1` VALUES (656, 1714462150, 1714462150, 653, '修改任务分组', 'A', NULL, NULL, '/cron/v1/task/group', 'PUT', NULL, 'cron:task:group:update', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu_copy1` VALUES (657, 1714462150, 1714462150, 653, '删除任务分组', 'A', NULL, NULL, '/cron/v1/task/group', 'DELETE', NULL, 'cron:task:group:delete', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu_copy1` VALUES (658, 1714462150, 1714462150, 653, '查看任务', 'A', NULL, NULL, '/cron/v1/tasks', 'GET', NULL, 'cron:task:query', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu_copy1` VALUES (659, 1714462150, 1714462150, 653, '新增任务', 'A', NULL, NULL, '/cron/v1/task', 'POST', NULL, 'cron:task:add', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu_copy1` VALUES (660, 1714462150, 1714462150, 653, '立即执行任务', 'A', NULL, NULL, '/cron/v1/task/exec', 'POST', NULL, 'cron:task:exec', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu_copy1` VALUES (661, 1714462150, 1714462150, 653, '取消执行任务', 'A', NULL, NULL, '/cron/v1/task/cancel', 'POST', NULL, 'cron:task:cancel', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu_copy1` VALUES (662, 1714462150, 1714462150, 653, '修改任务', 'A', NULL, NULL, '/cron/v1/task', 'PUT', NULL, 'cron:task:update', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu_copy1` VALUES (663, 1714462150, 1714462150, 653, '删除任务', 'A', NULL, NULL, '/cron/v1/task', 'DELETE', NULL, 'cron:task:delete', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu_copy1` VALUES (664, 1714462150, 1714462150, 653, '任务状态管理', 'G', NULL, NULL, NULL, NULL, NULL, 'cron:task:status', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu_copy1` VALUES (665, 1714462150, 1714462150, 664, '启用任务', 'A', NULL, NULL, '/cron/v1/task/enable', 'POST', NULL, NULL, NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu_copy1` VALUES (666, 1714462150, 1714462150, 664, '禁用任务', 'A', NULL, NULL, '/cron/v1/task/disable', 'POST', NULL, NULL, NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu_copy1` VALUES (667, 1714462150, 1714462150, 653, '任务日志', 'G', NULL, NULL, NULL, NULL, NULL, 'cron:task:log', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu_copy1` VALUES (668, 1714462150, 1714462150, 667, '获取任务日志分页', 'A', NULL, NULL, '/cron/v1/logs', 'GET', NULL, NULL, NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu_copy1` VALUES (669, 1714462150, 1714462150, 667, '获取任务日志详情', 'A', NULL, NULL, '/cron/v1/log', 'GET', NULL, NULL, NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu_copy1` VALUES (670, 1714462150, 1714462150, 522, '用户中心', 'M', 'SystemPlatformUserCenter', 'user', NULL, NULL, '/user-center', NULL, NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu_copy1` VALUES (671, 1714462150, 1714462150, 670, '协议管理', 'M', 'UserCenterAgreement', 'bookmark', NULL, NULL, '/user-center/agreement', NULL, NULL, NULL, NULL, NULL, 0, 0, 1, NULL, 1);
INSERT INTO `menu_copy1` VALUES (672, 1714462150, 1714462150, 671, '协议内容', 'M', 'UserCenterAgreementContent', 'bookmark', NULL, NULL, '/user-center/agreement/content', NULL, NULL, NULL, '/user-center/agreement/content/index', NULL, 0, 0, 1, NULL, 1);
INSERT INTO `menu_copy1` VALUES (673, 1714462150, 1714462150, 672, '查看协议列表', 'A', NULL, NULL, '/user-center/v1/agreement/contents', 'GET', NULL, 'uc:agreement:content:query', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu_copy1` VALUES (674, 1714462150, 1714462150, 672, '查看协议详细信息', 'A', NULL, NULL, '/user-center/v1/agreement/content', 'GET', NULL, 'uc:agreement:content:query', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu_copy1` VALUES (675, 1714462150, 1714462150, 672, '新增协议', 'A', NULL, NULL, '/user-center/v1/agreement/content', 'POST', NULL, 'uc:agreement:content:add', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu_copy1` VALUES (676, 1714462150, 1714462150, 672, '修改协议', 'A', NULL, NULL, '/user-center/v1/agreement/content', 'PUT', NULL, 'uc:agreement:content:update', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu_copy1` VALUES (677, 1714462150, 1714462150, 672, '删除协议', 'A', NULL, NULL, '/user-center/v1/agreement/content', 'DELETE', NULL, 'uc:agreement:content:delete', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu_copy1` VALUES (678, 1714462150, 1714462150, 671, '协议场景', 'M', 'UserCenterAgreementScene', 'common', NULL, NULL, '/user-center/agreement/scene', NULL, NULL, NULL, '/user-center/agreement/scene/index', NULL, 0, 0, 1, NULL, 1);
INSERT INTO `menu_copy1` VALUES (679, 1714462150, 1714462150, 678, '查看场景列表', 'A', NULL, NULL, '/user-center/v1/agreement/scenes', 'GET', NULL, 'uc:agreement:scene:query', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu_copy1` VALUES (680, 1714462150, 1714462150, 678, '查看场景详细信息', 'A', NULL, NULL, '/user-center/v1/agreement/scene', 'GET', NULL, 'uc:agreement:scene:query', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu_copy1` VALUES (681, 1714462150, 1714462150, 678, '新增场景', 'A', NULL, NULL, '/user-center/v1/agreement/scene', 'POST', NULL, 'uc:agreement:scene:add', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu_copy1` VALUES (682, 1714462150, 1714462150, 678, '修改场景', 'A', NULL, NULL, '/user-center/v1/agreement/scene', 'PUT', NULL, 'uc:agreement:scene:update', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu_copy1` VALUES (683, 1714462150, 1714462150, 678, '删除场景', 'A', NULL, NULL, '/user-center/v1/agreement/scene', 'DELETE', NULL, 'uc:agreement:scene:delete', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu_copy1` VALUES (684, 1714462150, 1714462150, 670, '授权渠道', 'M', 'UserCenterChannel', 'mind-mapping', NULL, NULL, '/user-center/channel', NULL, NULL, NULL, '/user-center/channel/index', NULL, 0, 0, 1, NULL, 1);
INSERT INTO `menu_copy1` VALUES (685, 1714462150, 1714462150, 684, '查看渠道', 'A', NULL, NULL, '/user-center/v1/channels', 'GET', NULL, 'uc:channel:query', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu_copy1` VALUES (686, 1714462150, 1714462150, 684, '新增渠道', 'A', NULL, NULL, '/user-center/v1/channel', 'POST', NULL, 'uc:channel:add', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu_copy1` VALUES (687, 1714462150, 1714462150, 684, '修改渠道', 'A', NULL, NULL, '/user-center/v1/channel', 'PUT', NULL, 'uc:channel:update', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu_copy1` VALUES (688, 1714462150, 1714462150, 684, '删除渠道', 'A', NULL, NULL, '/user-center/v1/channel', 'DELETE', NULL, 'uc:channel:delete', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu_copy1` VALUES (689, 1714462150, 1714462150, 670, '信息字段', 'M', 'UserCenterField', 'list', NULL, NULL, '/user-center/field', NULL, NULL, NULL, '/user-center/field/index', NULL, 0, 0, 1, NULL, 1);
INSERT INTO `menu_copy1` VALUES (690, 1714462150, 1714462150, 689, '查看字段', 'A', NULL, NULL, '/user-center/v1/fields', 'GET', NULL, 'uc:field:query', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu_copy1` VALUES (691, 1714462150, 1714462150, 689, '新增字段', 'A', NULL, NULL, '/user-center/v1/field', 'POST', NULL, 'uc:field:add', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu_copy1` VALUES (692, 1714462150, 1714462150, 689, '修改字段', 'A', NULL, NULL, '/user-center/v1/field', 'PUT', NULL, 'uc:field:update', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu_copy1` VALUES (693, 1714462150, 1714462150, 689, '删除字段', 'A', NULL, NULL, '/user-center/v1/field', 'DELETE', NULL, 'uc:field:delete', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu_copy1` VALUES (694, 1714462150, 1714462150, 670, '应用管理', 'M', 'UserCenterApp', 'apps', NULL, NULL, '/user-center/app', NULL, NULL, NULL, '/user-center/app/index', NULL, 0, 0, 1, NULL, 1);
INSERT INTO `menu_copy1` VALUES (695, 1714462150, 1714462150, 694, '查看应用', 'A', NULL, NULL, '/user-center/v1/apps', 'GET', NULL, 'uc:app:query', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu_copy1` VALUES (696, 1714462150, 1714462150, 694, '新增应用', 'A', NULL, NULL, '/user-center/v1/app', 'POST', NULL, 'uc:app:add', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu_copy1` VALUES (697, 1714462150, 1714462150, 694, '修改应用', 'A', NULL, NULL, '/user-center/v1/app', 'PUT', NULL, 'uc:app:update', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu_copy1` VALUES (698, 1714462150, 1714462150, 694, '删除应用', 'A', NULL, NULL, '/user-center/v1/app', 'DELETE', NULL, 'uc:app:delete', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu_copy1` VALUES (699, 1714462150, 1714462150, 670, '用户管理', 'M', 'user-center', 'user', NULL, NULL, '/user-center/user', NULL, NULL, NULL, '/user-center/user/index', NULL, 0, 0, 1, NULL, 1);
INSERT INTO `menu_copy1` VALUES (700, 1714462150, 1714462150, 699, '查看用户', 'A', NULL, NULL, '/user-center/v1/users', 'GET', NULL, 'uc:user:query', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu_copy1` VALUES (701, 1714462150, 1714462150, 699, '新增用户', 'A', NULL, NULL, '/user-center/v1/user', 'POST', NULL, 'uc:user:add', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu_copy1` VALUES (702, 1714462150, 1714462150, 699, '导入用户', 'A', NULL, NULL, '/user-center/v1/users', 'POST', NULL, 'uc:user:import', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu_copy1` VALUES (703, 1714462150, 1714462150, 699, '修改用户', 'A', NULL, NULL, '/user-center/v1/user', 'PUT', NULL, 'uc:user:update', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu_copy1` VALUES (704, 1714462150, 1714462150, 699, '删除用户', 'A', NULL, NULL, '/user-center/v1/user', 'DELETE', NULL, 'uc:user:delete', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu_copy1` VALUES (705, 1714462150, 1714462150, 699, '下线用户', 'A', NULL, NULL, '/user-center/v1/user/offline', 'POST', NULL, 'uc:user:offline', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu_copy1` VALUES (706, 1714462150, 1714462150, 699, '用户状态管理', 'G', NULL, NULL, NULL, NULL, NULL, 'uc:user:status', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu_copy1` VALUES (707, 1714462150, 1714462150, 706, '禁用用户', 'A', NULL, NULL, '/user-center/v1/user/disable', 'POST', NULL, 'uc:user:disable', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu_copy1` VALUES (708, 1714462150, 1714462150, 706, '启用用户', 'A', NULL, NULL, '/user-center/v1/user/enable', 'POST', NULL, 'uc:user:enable', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu_copy1` VALUES (709, 1714462150, 1714462150, 0, '应用平台', 'R', 'AppPlatform', 'apps', NULL, NULL, '/app', NULL, NULL, NULL, 'Layout', NULL, 2, 0, NULL, NULL, NULL);
INSERT INTO `menu_copy1` VALUES (710, 1714462150, 1714462150, 709, '首页面板', 'M', 'AppDashboard', 'dashboard', NULL, NULL, '/app/dashboard', NULL, NULL, NULL, '/dashboard/index', NULL, 0, 0, 1, 1, 1);
INSERT INTO `menu_copy1` VALUES (711, 1714462150, 1714462150, 709, '信号灯系统', 'M', 'PartyAffairs', 'apps', NULL, NULL, '/party-affairs', NULL, NULL, NULL, NULL, NULL, 0, 0, NULL, NULL, NULL);
INSERT INTO `menu_copy1` VALUES (712, 1714462150, 1714462150, 711, '轮播管理', 'M', 'PartyAffairsBanner', 'list', NULL, NULL, '/party-affairs/banner', NULL, NULL, NULL, '/party-affairs/banner/index', NULL, 0, NULL, 1, NULL, 1);
INSERT INTO `menu_copy1` VALUES (713, 1714462150, 1714462150, 712, '查看轮播', 'A', NULL, NULL, '/party-affairs/v1/banners', 'GET', NULL, 'party-affairs:banner:query', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu_copy1` VALUES (714, 1714462150, 1714462150, 712, '新增轮播', 'A', NULL, NULL, '/party-affairs/v1/banner', 'POST', NULL, 'party-affairs:banner:add', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu_copy1` VALUES (715, 1714462150, 1714462150, 712, '修改轮播', 'A', NULL, NULL, '/party-affairs/v1/banner', 'PUT', NULL, 'party-affairs:banner:update', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu_copy1` VALUES (716, 1714462150, 1714462150, 712, '删除轮播', 'A', NULL, NULL, '/party-affairs/v1/banner', 'DELETE', NULL, 'party-affairs:banner:delete', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu_copy1` VALUES (717, 1714462150, 1714462150, 711, '通知管理', 'M', 'PartyAffairsNotice', 'notification', NULL, NULL, '/party-affairs/notice', NULL, NULL, NULL, '/party-affairs/notice/index', NULL, 0, NULL, 1, NULL, 1);
INSERT INTO `menu_copy1` VALUES (718, 1714462150, 1714462150, 717, '查看通知', 'A', NULL, NULL, '/party-affairs/v1/notices', 'GET', NULL, 'party-affairs:notice:query', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu_copy1` VALUES (719, 1714462150, 1714462150, 717, '查看通知阅读情况', 'A', NULL, NULL, '/party-affairs/v1/notice/users', 'GET', NULL, 'party-affairs:notice:user:query', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu_copy1` VALUES (720, 1714462150, 1714462150, 717, '新增通知', 'A', NULL, NULL, '/party-affairs/v1/notice', 'POST', NULL, 'party-affairs:notice:add', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu_copy1` VALUES (721, 1714462150, 1714462150, 717, '修改通知', 'A', NULL, NULL, '/party-affairs/v1/notice', 'PUT', NULL, 'party-affairs:notice:update', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu_copy1` VALUES (722, 1714462150, 1714462150, 717, '删除通知', 'A', NULL, NULL, '/party-affairs/v1/notice', 'DELETE', NULL, 'party-affairs:notice:delete', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu_copy1` VALUES (723, 1714462150, 1714462150, 711, '新闻管理', 'M', 'PartyAffairsNews', 'book', NULL, NULL, '/party-affairs/news', NULL, NULL, NULL, NULL, NULL, 0, 0, NULL, NULL, NULL);
INSERT INTO `menu_copy1` VALUES (724, 1714462150, 1714462150, 723, '板块管理', 'M', 'PartyAffairsNewsClassify', 'menu', NULL, NULL, '/party-affairs/news/classify', NULL, NULL, NULL, '/party-affairs/news/classify/index', NULL, 0, NULL, 1, NULL, 1);
INSERT INTO `menu_copy1` VALUES (725, 1714462150, 1714462150, 724, '查看板块', 'A', NULL, NULL, '/party-affairs/v1/news/classify', 'GET', NULL, 'party-affairs:news:classify:query', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu_copy1` VALUES (726, 1714462150, 1714462150, 724, '新增板块', 'A', NULL, NULL, '/party-affairs/v1/news/classify', 'POST', NULL, 'party-affairs:news:classify:add', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu_copy1` VALUES (727, 1714462150, 1714462150, 724, '修改板块', 'A', NULL, NULL, '/party-affairs/v1/news/classify', 'PUT', NULL, 'party-affairs:news:classify:update', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu_copy1` VALUES (728, 1714462150, 1714462150, 724, '删除板块', 'A', NULL, NULL, '/party-affairs/v1/news/classify', 'DELETE', NULL, 'party-affairs:news:classify:delete', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu_copy1` VALUES (729, 1714462150, 1714462150, 723, '内容管理', 'M', 'PartyAffairsNewsContent', 'book', NULL, NULL, '/party-affairs/news/content', NULL, NULL, NULL, '/party-affairs/news/content/index', NULL, 0, NULL, 1, NULL, 1);
INSERT INTO `menu_copy1` VALUES (730, 1714462150, 1714462150, 729, '查看内容列表', 'A', NULL, NULL, '/party-affairs/v1/news/contents', 'GET', NULL, 'party-affairs:news:content:query', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu_copy1` VALUES (731, 1714462150, 1714462150, 729, '查看指定内容', 'A', NULL, NULL, '/party-affairs/v1/news/content', 'GET', NULL, 'party-affairs:news:content:info', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu_copy1` VALUES (732, 1714462150, 1714462150, 729, '新增内容', 'A', NULL, NULL, '/party-affairs/v1/news/content', 'POST', NULL, 'party-affairs:news:content:add', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu_copy1` VALUES (733, 1714462150, 1714462150, 729, '修改内容', 'A', NULL, NULL, '/party-affairs/v1/news/content', 'PUT', NULL, 'party-affairs:news:content:update', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu_copy1` VALUES (734, 1714462150, 1714462150, 729, '删除内容', 'A', NULL, NULL, '/party-affairs/v1/news/content', 'DELETE', NULL, 'party-affairs:news:content:delete', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu_copy1` VALUES (735, 1714462150, 1714462150, 723, '评论管理', 'M', 'PartyAffairsNewsComment', 'book', NULL, NULL, '/party-affairs/news/comment', NULL, NULL, NULL, '/party-affairs/news/comment/index', NULL, 0, 1, NULL, NULL, NULL);
INSERT INTO `menu_copy1` VALUES (736, 1714462150, 1714462150, 735, '查看评论列表', 'A', NULL, NULL, '/party-affairs/v1/news/comments', 'GET', NULL, 'party-affairs:news:comment:query', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu_copy1` VALUES (737, 1714462150, 1714462150, 735, '删除评论', 'A', NULL, NULL, '/party-affairs/v1/news/comment', 'DELETE', NULL, 'party-affairs:news:comment:delete', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu_copy1` VALUES (738, 1714462150, 1714462150, 711, '资料管理', 'M', 'PartyAffairsResource', 'public', NULL, NULL, '/party-affairs/resource', NULL, NULL, NULL, NULL, NULL, 0, 0, NULL, NULL, NULL);
INSERT INTO `menu_copy1` VALUES (739, 1714462150, 1714462150, 738, '板块管理', 'M', 'PartyAffairsResourceClassify', 'menu', NULL, NULL, '/party-affairs/resource/classify', NULL, NULL, NULL, '/party-affairs/resource/classify/index', NULL, 0, NULL, 1, NULL, 1);
INSERT INTO `menu_copy1` VALUES (740, 1714462150, 1714462150, 739, '查看板块', 'A', NULL, NULL, '/party-affairs/v1/resource/classify', 'GET', NULL, 'party-affairs:resource:classify:query', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu_copy1` VALUES (741, 1714462150, 1714462150, 739, '新增板块', 'A', NULL, NULL, '/party-affairs/v1/resource/classify', 'POST', NULL, 'party-affairs:resource:classify:add', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu_copy1` VALUES (742, 1714462150, 1714462150, 739, '修改板块', 'A', NULL, NULL, '/party-affairs/v1/resource/classify', 'PUT', NULL, 'party-affairs:resource:classify:update', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu_copy1` VALUES (743, 1714462150, 1714462150, 739, '删除板块', 'A', NULL, NULL, '/party-affairs/v1/resource/classify', 'DELETE', NULL, 'party-affairs:resource:classify:delete', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu_copy1` VALUES (744, 1714462150, 1714462150, 738, '内容管理', 'M', 'PartyAffairsResourceContent', 'public', NULL, NULL, '/party-affairs/resource/content', NULL, NULL, NULL, '/party-affairs/resource/content/index', NULL, 0, NULL, 1, NULL, 1);
INSERT INTO `menu_copy1` VALUES (745, 1714462150, 1714462150, 744, '查看内容列表', 'A', NULL, NULL, '/party-affairs/v1/resource/contents', 'GET', NULL, 'party-affairs:resource:content:query', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu_copy1` VALUES (746, 1714462150, 1714462150, 744, '新增内容', 'A', NULL, NULL, '/party-affairs/v1/resource/content', 'POST', NULL, 'party-affairs:resource:content:add', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu_copy1` VALUES (747, 1714462150, 1714462150, 744, '修改内容', 'A', NULL, NULL, '/party-affairs/v1/resource/content', 'PUT', NULL, 'party-affairs:resource:content:update', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu_copy1` VALUES (748, 1714462150, 1714462150, 744, '删除内容', 'A', NULL, NULL, '/party-affairs/v1/resource/content', 'DELETE', NULL, 'party-affairs:resource:content:delete', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu_copy1` VALUES (749, 1714462150, 1714462150, 711, '任务管理', 'M', 'PartyAffairsTask', 'layers', NULL, NULL, '/party-affairs/task', NULL, NULL, NULL, NULL, NULL, 0, 0, NULL, NULL, NULL);
INSERT INTO `menu_copy1` VALUES (750, 1714462150, 1714462150, 749, '任务配置', 'M', 'PartyAffairsTaskConfigure', 'unordered-list', NULL, NULL, '/party-affairs/task/configure', NULL, NULL, NULL, '/party-affairs/task/configure/index', NULL, 0, NULL, 1, NULL, 1);
INSERT INTO `menu_copy1` VALUES (751, 1714462150, 1714462150, 750, '查看配置', 'A', NULL, NULL, '/party-affairs/v1/task/configure', 'GET', NULL, 'party-affairs:task:configure:query', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu_copy1` VALUES (752, 1714462150, 1714462150, 750, '新增配置', 'A', NULL, NULL, '/party-affairs/v1/task/configure', 'POST', NULL, 'party-affairs:task:configure:add', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu_copy1` VALUES (753, 1714462150, 1714462150, 750, '修改配置', 'A', NULL, NULL, '/party-affairs/v1/task/configure', 'PUT', NULL, 'party-affairs:task:configure:update', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu_copy1` VALUES (754, 1714462150, 1714462150, 750, '删除配置', 'A', NULL, NULL, '/party-affairs/v1/task/configure', 'DELETE', NULL, 'party-affairs:task:configure:delete', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu_copy1` VALUES (755, 1714462150, 1714462150, 749, '任务统计', 'M', 'PartyAffairsTaskContent', 'computer', NULL, NULL, '/party-affairs/task/value', NULL, NULL, NULL, '/party-affairs/task/value/index', NULL, 0, NULL, 1, NULL, 1);
INSERT INTO `menu_copy1` VALUES (756, 1714462150, 1714462150, 755, '查看任务填写列表', 'A', NULL, NULL, '/party-affairs/v1/task/values', 'GET', NULL, 'party-affairs:task:value:query', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu_copy1` VALUES (757, 1714462150, 1714462150, 755, '查看指定任务填写详情', 'A', NULL, NULL, '/party-affairs/v1/task/value', 'GET', NULL, 'party-affairs:task:value:query', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu_copy1` VALUES (758, 1714462150, 1714462150, 755, '修改任务填写详情', 'A', NULL, NULL, '/party-affairs/v1/task/value', 'PUT', NULL, 'party-affairs:task:value:update', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu_copy1` VALUES (759, 1714462150, 1714462150, 755, '删除任务填写详情', 'A', NULL, NULL, '/party-affairs/v1/task/value', 'DELETE', NULL, 'party-affairs:task:value:delete', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu_copy1` VALUES (760, 1714462150, 1714462150, 711, '视频管理', 'M', 'PartyAffairsVideo', 'video-camera', NULL, NULL, '/party-affairs/video', NULL, NULL, NULL, NULL, NULL, 0, 0, NULL, NULL, NULL);
INSERT INTO `menu_copy1` VALUES (761, 1714462150, 1714462150, 760, '板块管理', 'M', 'PartyAffairsVideoClassify', 'menu', NULL, NULL, '/party-affairs/video/classify', NULL, NULL, NULL, '/party-affairs/video/classify/index', NULL, 0, NULL, 1, NULL, 1);
INSERT INTO `menu_copy1` VALUES (762, 1714462150, 1714462150, 761, '查看板块', 'A', NULL, NULL, '/party-affairs/v1/video/classify', 'GET', NULL, 'party-affairs:video:classify:query', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu_copy1` VALUES (763, 1714462150, 1714462150, 761, '新增板块', 'A', NULL, NULL, '/party-affairs/v1/video/classify', 'POST', NULL, 'party-affairs:video:classify:add', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu_copy1` VALUES (764, 1714462150, 1714462150, 761, '修改板块', 'A', NULL, NULL, '/party-affairs/v1/video/classify', 'PUT', NULL, 'party-affairs:video:classify:update', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu_copy1` VALUES (765, 1714462150, 1714462150, 761, '删除板块', 'A', NULL, NULL, '/party-affairs/v1/video/classify', 'DELETE', NULL, 'party-affairs:video:classify:delete', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu_copy1` VALUES (766, 1714462150, 1714462150, 760, '内容管理', 'M', 'PartyAffairsVideoContent', 'video-camera', NULL, NULL, '/party-affairs/video/content', NULL, NULL, NULL, '/party-affairs/video/content/index', NULL, 0, NULL, 1, NULL, 1);
INSERT INTO `menu_copy1` VALUES (767, 1714462150, 1714462150, 766, '查看内容列表', 'A', NULL, NULL, '/party-affairs/v1/video/contents', 'GET', NULL, 'party-affairs:video:content:query', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu_copy1` VALUES (768, 1714462150, 1714462150, 766, '新增内容', 'A', NULL, NULL, '/party-affairs/v1/video/content', 'POST', NULL, 'party-affairs:video:content:add', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu_copy1` VALUES (769, 1714462150, 1714462150, 766, '修改内容', 'A', NULL, NULL, '/party-affairs/v1/video/content', 'PUT', NULL, 'party-affairs:video:content:update', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu_copy1` VALUES (770, 1714462150, 1714462150, 766, '删除内容', 'A', NULL, NULL, '/party-affairs/v1/video/content', 'DELETE', NULL, 'party-affairs:video:content:delete', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
INSERT INTO `menu_copy1` VALUES (771, 1714462150, 1714462150, 766, '进度统计', 'A', NULL, NULL, '/party-affairs/v1/video/process', 'GET', NULL, 'party-affairs:video:content:process', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);
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
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COMMENT='资源对象';

-- ----------------------------
-- Records of object
-- ----------------------------
BEGIN;
INSERT INTO `object` VALUES (1, 1713706137, 1713706137, 'env', '环境资源对象', '/configure/v1/envs', 'GET', '', 'list.name', 'list.id', '环境资源对象');
INSERT INTO `object` VALUES (2, 1713706137, 1713706137, 'server', '服务资源对象', '/configure/v1/servers', 'GET', '{\"page\":1,\"page_size\":50}', 'list.name', 'list.id', '服务资源对象');
INSERT INTO `object` VALUES (3, 1713706137, 1713706137, 'app', '应用资源对象', '/user-center/v1/apps', 'GET', '{\"page\":1,\"page_size\":50}', 'list.name', 'list.id', '应用资源对象');
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
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COMMENT='角色信息';

-- ----------------------------
-- Records of role
-- ----------------------------
BEGIN;
INSERT INTO `role` VALUES (1, 1713706137, 1713706137, 0, '超级管理员', 'superAdmin', 1, '超级管理员  ', NULL, 'ALL');
INSERT INTO `role` VALUES (2, 1714444901, 1714444901, 1, '管理员', 'admin', 1, '管理员', NULL, 'CUR');
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
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COMMENT='角色层级信息';

-- ----------------------------
-- Records of role_closure
-- ----------------------------
BEGIN;
INSERT INTO `role_closure` VALUES (1, 1, 2);
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
) ENGINE=InnoDB AUTO_INCREMENT=1444 DEFAULT CHARSET=utf8mb4 COMMENT='角色菜单信息';

-- ----------------------------
-- Records of role_menu
-- ----------------------------
BEGIN;
INSERT INTO `role_menu` VALUES (1271, 1714462165, 2, 522);
INSERT INTO `role_menu` VALUES (1272, 1714462165, 2, 523);
INSERT INTO `role_menu` VALUES (1273, 1714462165, 2, 524);
INSERT INTO `role_menu` VALUES (1274, 1714462165, 2, 539);
INSERT INTO `role_menu` VALUES (1275, 1714462165, 2, 540);
INSERT INTO `role_menu` VALUES (1276, 1714462165, 2, 541);
INSERT INTO `role_menu` VALUES (1277, 1714462165, 2, 542);
INSERT INTO `role_menu` VALUES (1278, 1714462165, 2, 543);
INSERT INTO `role_menu` VALUES (1279, 1714462165, 2, 544);
INSERT INTO `role_menu` VALUES (1280, 1714462165, 2, 545);
INSERT INTO `role_menu` VALUES (1281, 1714462165, 2, 546);
INSERT INTO `role_menu` VALUES (1282, 1714462165, 2, 547);
INSERT INTO `role_menu` VALUES (1283, 1714462165, 2, 548);
INSERT INTO `role_menu` VALUES (1284, 1714462165, 2, 549);
INSERT INTO `role_menu` VALUES (1285, 1714462165, 2, 550);
INSERT INTO `role_menu` VALUES (1286, 1714462165, 2, 551);
INSERT INTO `role_menu` VALUES (1287, 1714462165, 2, 552);
INSERT INTO `role_menu` VALUES (1288, 1714462165, 2, 553);
INSERT INTO `role_menu` VALUES (1289, 1714462165, 2, 554);
INSERT INTO `role_menu` VALUES (1290, 1714462165, 2, 555);
INSERT INTO `role_menu` VALUES (1291, 1714462165, 2, 556);
INSERT INTO `role_menu` VALUES (1292, 1714462165, 2, 557);
INSERT INTO `role_menu` VALUES (1293, 1714462165, 2, 558);
INSERT INTO `role_menu` VALUES (1294, 1714462165, 2, 559);
INSERT INTO `role_menu` VALUES (1295, 1714462165, 2, 560);
INSERT INTO `role_menu` VALUES (1296, 1714462165, 2, 561);
INSERT INTO `role_menu` VALUES (1297, 1714462165, 2, 562);
INSERT INTO `role_menu` VALUES (1298, 1714462165, 2, 563);
INSERT INTO `role_menu` VALUES (1299, 1714462165, 2, 564);
INSERT INTO `role_menu` VALUES (1300, 1714462165, 2, 565);
INSERT INTO `role_menu` VALUES (1301, 1714462165, 2, 566);
INSERT INTO `role_menu` VALUES (1302, 1714462165, 2, 567);
INSERT INTO `role_menu` VALUES (1303, 1714462165, 2, 568);
INSERT INTO `role_menu` VALUES (1304, 1714462165, 2, 569);
INSERT INTO `role_menu` VALUES (1305, 1714462165, 2, 570);
INSERT INTO `role_menu` VALUES (1306, 1714462165, 2, 571);
INSERT INTO `role_menu` VALUES (1307, 1714462165, 2, 572);
INSERT INTO `role_menu` VALUES (1308, 1714462165, 2, 573);
INSERT INTO `role_menu` VALUES (1309, 1714462165, 2, 574);
INSERT INTO `role_menu` VALUES (1310, 1714462165, 2, 575);
INSERT INTO `role_menu` VALUES (1311, 1714462165, 2, 576);
INSERT INTO `role_menu` VALUES (1312, 1714462165, 2, 577);
INSERT INTO `role_menu` VALUES (1313, 1714462165, 2, 578);
INSERT INTO `role_menu` VALUES (1314, 1714462165, 2, 579);
INSERT INTO `role_menu` VALUES (1315, 1714462165, 2, 580);
INSERT INTO `role_menu` VALUES (1316, 1714462165, 2, 581);
INSERT INTO `role_menu` VALUES (1317, 1714462165, 2, 582);
INSERT INTO `role_menu` VALUES (1318, 1714462165, 2, 583);
INSERT INTO `role_menu` VALUES (1319, 1714462165, 2, 584);
INSERT INTO `role_menu` VALUES (1320, 1714462165, 2, 585);
INSERT INTO `role_menu` VALUES (1321, 1714462165, 2, 586);
INSERT INTO `role_menu` VALUES (1322, 1714462165, 2, 587);
INSERT INTO `role_menu` VALUES (1323, 1714462165, 2, 588);
INSERT INTO `role_menu` VALUES (1324, 1714462165, 2, 589);
INSERT INTO `role_menu` VALUES (1325, 1714462165, 2, 590);
INSERT INTO `role_menu` VALUES (1326, 1714462165, 2, 591);
INSERT INTO `role_menu` VALUES (1327, 1714462165, 2, 592);
INSERT INTO `role_menu` VALUES (1328, 1714462165, 2, 593);
INSERT INTO `role_menu` VALUES (1329, 1714462165, 2, 594);
INSERT INTO `role_menu` VALUES (1330, 1714462165, 2, 595);
INSERT INTO `role_menu` VALUES (1331, 1714462165, 2, 596);
INSERT INTO `role_menu` VALUES (1332, 1714462165, 2, 597);
INSERT INTO `role_menu` VALUES (1333, 1714462165, 2, 598);
INSERT INTO `role_menu` VALUES (1334, 1714462165, 2, 599);
INSERT INTO `role_menu` VALUES (1335, 1714462165, 2, 600);
INSERT INTO `role_menu` VALUES (1336, 1714462165, 2, 601);
INSERT INTO `role_menu` VALUES (1337, 1714462165, 2, 602);
INSERT INTO `role_menu` VALUES (1338, 1714462165, 2, 603);
INSERT INTO `role_menu` VALUES (1339, 1714462165, 2, 604);
INSERT INTO `role_menu` VALUES (1340, 1714462165, 2, 605);
INSERT INTO `role_menu` VALUES (1341, 1714462165, 2, 606);
INSERT INTO `role_menu` VALUES (1342, 1714462165, 2, 607);
INSERT INTO `role_menu` VALUES (1343, 1714462165, 2, 608);
INSERT INTO `role_menu` VALUES (1344, 1714462165, 2, 609);
INSERT INTO `role_menu` VALUES (1345, 1714462165, 2, 610);
INSERT INTO `role_menu` VALUES (1346, 1714462165, 2, 611);
INSERT INTO `role_menu` VALUES (1347, 1714462165, 2, 612);
INSERT INTO `role_menu` VALUES (1348, 1714462165, 2, 613);
INSERT INTO `role_menu` VALUES (1349, 1714462165, 2, 614);
INSERT INTO `role_menu` VALUES (1350, 1714462165, 2, 615);
INSERT INTO `role_menu` VALUES (1351, 1714462165, 2, 616);
INSERT INTO `role_menu` VALUES (1352, 1714462165, 2, 617);
INSERT INTO `role_menu` VALUES (1353, 1714462165, 2, 618);
INSERT INTO `role_menu` VALUES (1354, 1714462165, 2, 619);
INSERT INTO `role_menu` VALUES (1355, 1714462165, 2, 620);
INSERT INTO `role_menu` VALUES (1356, 1714462165, 2, 621);
INSERT INTO `role_menu` VALUES (1357, 1714462165, 2, 622);
INSERT INTO `role_menu` VALUES (1358, 1714462165, 2, 623);
INSERT INTO `role_menu` VALUES (1359, 1714462165, 2, 624);
INSERT INTO `role_menu` VALUES (1360, 1714462165, 2, 625);
INSERT INTO `role_menu` VALUES (1361, 1714462165, 2, 626);
INSERT INTO `role_menu` VALUES (1362, 1714462165, 2, 627);
INSERT INTO `role_menu` VALUES (1363, 1714462165, 2, 628);
INSERT INTO `role_menu` VALUES (1364, 1714462165, 2, 629);
INSERT INTO `role_menu` VALUES (1365, 1714462165, 2, 630);
INSERT INTO `role_menu` VALUES (1366, 1714462165, 2, 631);
INSERT INTO `role_menu` VALUES (1367, 1714462165, 2, 632);
INSERT INTO `role_menu` VALUES (1368, 1714462165, 2, 633);
INSERT INTO `role_menu` VALUES (1369, 1714462165, 2, 634);
INSERT INTO `role_menu` VALUES (1370, 1714462165, 2, 635);
INSERT INTO `role_menu` VALUES (1371, 1714462165, 2, 636);
INSERT INTO `role_menu` VALUES (1372, 1714462165, 2, 637);
INSERT INTO `role_menu` VALUES (1373, 1714462165, 2, 638);
INSERT INTO `role_menu` VALUES (1374, 1714462165, 2, 639);
INSERT INTO `role_menu` VALUES (1375, 1714462165, 2, 640);
INSERT INTO `role_menu` VALUES (1376, 1714462165, 2, 641);
INSERT INTO `role_menu` VALUES (1377, 1714462165, 2, 642);
INSERT INTO `role_menu` VALUES (1378, 1714462165, 2, 643);
INSERT INTO `role_menu` VALUES (1379, 1714462165, 2, 644);
INSERT INTO `role_menu` VALUES (1380, 1714462165, 2, 645);
INSERT INTO `role_menu` VALUES (1381, 1714462165, 2, 646);
INSERT INTO `role_menu` VALUES (1382, 1714462165, 2, 647);
INSERT INTO `role_menu` VALUES (1383, 1714462165, 2, 648);
INSERT INTO `role_menu` VALUES (1384, 1714462165, 2, 649);
INSERT INTO `role_menu` VALUES (1385, 1714462165, 2, 650);
INSERT INTO `role_menu` VALUES (1386, 1714462165, 2, 651);
INSERT INTO `role_menu` VALUES (1387, 1714462165, 2, 652);
INSERT INTO `role_menu` VALUES (1388, 1714462165, 2, 653);
INSERT INTO `role_menu` VALUES (1389, 1714462165, 2, 654);
INSERT INTO `role_menu` VALUES (1390, 1714462165, 2, 655);
INSERT INTO `role_menu` VALUES (1391, 1714462165, 2, 656);
INSERT INTO `role_menu` VALUES (1392, 1714462165, 2, 657);
INSERT INTO `role_menu` VALUES (1393, 1714462165, 2, 658);
INSERT INTO `role_menu` VALUES (1394, 1714462165, 2, 659);
INSERT INTO `role_menu` VALUES (1395, 1714462165, 2, 660);
INSERT INTO `role_menu` VALUES (1396, 1714462165, 2, 661);
INSERT INTO `role_menu` VALUES (1397, 1714462165, 2, 662);
INSERT INTO `role_menu` VALUES (1398, 1714462165, 2, 663);
INSERT INTO `role_menu` VALUES (1399, 1714462165, 2, 664);
INSERT INTO `role_menu` VALUES (1400, 1714462165, 2, 665);
INSERT INTO `role_menu` VALUES (1401, 1714462165, 2, 666);
INSERT INTO `role_menu` VALUES (1402, 1714462165, 2, 667);
INSERT INTO `role_menu` VALUES (1403, 1714462165, 2, 668);
INSERT INTO `role_menu` VALUES (1404, 1714462165, 2, 669);
INSERT INTO `role_menu` VALUES (1405, 1714462165, 2, 670);
INSERT INTO `role_menu` VALUES (1406, 1714462165, 2, 671);
INSERT INTO `role_menu` VALUES (1407, 1714462165, 2, 672);
INSERT INTO `role_menu` VALUES (1408, 1714462165, 2, 673);
INSERT INTO `role_menu` VALUES (1409, 1714462165, 2, 674);
INSERT INTO `role_menu` VALUES (1410, 1714462165, 2, 675);
INSERT INTO `role_menu` VALUES (1411, 1714462165, 2, 676);
INSERT INTO `role_menu` VALUES (1412, 1714462165, 2, 677);
INSERT INTO `role_menu` VALUES (1413, 1714462165, 2, 678);
INSERT INTO `role_menu` VALUES (1414, 1714462165, 2, 679);
INSERT INTO `role_menu` VALUES (1415, 1714462165, 2, 680);
INSERT INTO `role_menu` VALUES (1416, 1714462165, 2, 681);
INSERT INTO `role_menu` VALUES (1417, 1714462165, 2, 682);
INSERT INTO `role_menu` VALUES (1418, 1714462165, 2, 683);
INSERT INTO `role_menu` VALUES (1419, 1714462165, 2, 684);
INSERT INTO `role_menu` VALUES (1420, 1714462165, 2, 685);
INSERT INTO `role_menu` VALUES (1421, 1714462165, 2, 686);
INSERT INTO `role_menu` VALUES (1422, 1714462165, 2, 687);
INSERT INTO `role_menu` VALUES (1423, 1714462165, 2, 688);
INSERT INTO `role_menu` VALUES (1424, 1714462165, 2, 689);
INSERT INTO `role_menu` VALUES (1425, 1714462165, 2, 690);
INSERT INTO `role_menu` VALUES (1426, 1714462165, 2, 691);
INSERT INTO `role_menu` VALUES (1427, 1714462165, 2, 692);
INSERT INTO `role_menu` VALUES (1428, 1714462165, 2, 693);
INSERT INTO `role_menu` VALUES (1429, 1714462165, 2, 694);
INSERT INTO `role_menu` VALUES (1430, 1714462165, 2, 695);
INSERT INTO `role_menu` VALUES (1431, 1714462165, 2, 696);
INSERT INTO `role_menu` VALUES (1432, 1714462165, 2, 697);
INSERT INTO `role_menu` VALUES (1433, 1714462165, 2, 698);
INSERT INTO `role_menu` VALUES (1434, 1714462165, 2, 699);
INSERT INTO `role_menu` VALUES (1435, 1714462165, 2, 700);
INSERT INTO `role_menu` VALUES (1436, 1714462165, 2, 701);
INSERT INTO `role_menu` VALUES (1437, 1714462165, 2, 702);
INSERT INTO `role_menu` VALUES (1438, 1714462165, 2, 703);
INSERT INTO `role_menu` VALUES (1439, 1714462165, 2, 704);
INSERT INTO `role_menu` VALUES (1440, 1714462165, 2, 705);
INSERT INTO `role_menu` VALUES (1441, 1714462165, 2, 706);
INSERT INTO `role_menu` VALUES (1442, 1714462165, 2, 707);
INSERT INTO `role_menu` VALUES (1443, 1714462165, 2, 708);
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
INSERT INTO `user` VALUES (1, 1713706137, 1714456004, 1, 1, '超级管理员', '超级管理员', 'M', '18888888888', '$2a$10$9qRJe9KQo6sEcU8ipKg.e.dkl2E7Wy64SigYlgraTAn.1paHFq6W.', '', '1280291001@qq.com', 1, NULL, 1714456004, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJkZXBhcnRtZW50X2lkIjoxLCJkZXBhcnRtZW50X2tleXdvcmQiOiJjb21wYW55IiwiZXhwIjoxNzE0NDYzMjA1LCJpYXQiOjE3MTQ0NTYwMDQsIm5iZiI6MTcxNDQ1NjAwNCwicm9sZV9pZCI6MSwicm9sZV9rZXl3b3JkIjoic3VwZXJBZG1pbiIsInVzZXJfaWQiOjF9.q3GLKFp9Ax2MrOsksLhkGxU-WSow312GO0e2rUr6XXQ');
INSERT INTO `user` VALUES (2, 1714444948, 1714461993, 1, 2, 'testing', 'testing', 'M', '18286219255', '$2a$10$D2e/mU8JLpInF1foiuDWaepkECMzG.J/o.T7Je.c8LV4w4ziPVr.a', '', '1280291001@qq.com', 1, NULL, 1714461993, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJkZXBhcnRtZW50X2lkIjoxLCJkZXBhcnRtZW50X2tleXdvcmQiOiJjb21wYW55IiwiZXhwIjoxNzE0NDY5MTk0LCJpYXQiOjE3MTQ0NjE5OTMsIm5iZiI6MTcxNDQ2MTk5Mywicm9sZV9pZCI6Miwicm9sZV9rZXl3b3JkIjoiYWRtaW4iLCJ1c2VyX2lkIjoyfQ.jZzRK4svID1gr7H0R7-FeJpAKEO5tgx9dkOAD1yBHLo');
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
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COMMENT='用户职位信息';

-- ----------------------------
-- Records of user_job
-- ----------------------------
BEGIN;
INSERT INTO `user_job` VALUES (1, 1713706137, 1, 1);
INSERT INTO `user_job` VALUES (2, 1714444948, 1, 2);
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
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COMMENT='用户角色信息';

-- ----------------------------
-- Records of user_role
-- ----------------------------
BEGIN;
INSERT INTO `user_role` VALUES (1, 1713706137, 1, 1);
INSERT INTO `user_role` VALUES (2, 1714444948, 2, 2);
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
