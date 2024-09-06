/*
 Navicat Premium Data Transfer

 Source Server         : dev
 Source Server Type    : MySQL
 Source Server Version : 80300
 Source Host           : localhost:3306
 Source Schema         : manager

 Target Server Type    : MySQL
 Target Server Version : 80300
 File Encoding         : 65001

 Date: 30/08/2024 10:42:56
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
) ENGINE=InnoDB AUTO_INCREMENT=1973 DEFAULT CHARSET=utf8mb4 ;

-- ----------------------------
-- Records of casbin_rule
-- ----------------------------
BEGIN;
INSERT INTO `casbin_rule` VALUES (1794, 'p', '3', '/configure/api/v1/business', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES (1792, 'p', '3', '/configure/api/v1/business', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (1793, 'p', '3', '/configure/api/v1/business', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES (1791, 'p', '3', '/configure/api/v1/businesses', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (1811, 'p', '3', '/configure/api/v1/configure', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES (1810, 'p', '3', '/configure/api/v1/configure/compare', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (1782, 'p', '3', '/configure/api/v1/env', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES (1779, 'p', '3', '/configure/api/v1/env', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (1780, 'p', '3', '/configure/api/v1/env', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES (1781, 'p', '3', '/configure/api/v1/env/status', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES (1783, 'p', '3', '/configure/api/v1/env/token', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (1784, 'p', '3', '/configure/api/v1/env/token', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES (1777, 'p', '3', '/configure/api/v1/envs', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (1800, 'p', '3', '/configure/api/v1/resource', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES (1798, 'p', '3', '/configure/api/v1/resource', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (1799, 'p', '3', '/configure/api/v1/resource', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES (1797, 'p', '3', '/configure/api/v1/resources', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (1790, 'p', '3', '/configure/api/v1/server', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES (1787, 'p', '3', '/configure/api/v1/server', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (1788, 'p', '3', '/configure/api/v1/server', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES (1789, 'p', '3', '/configure/api/v1/server/status', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES (1785, 'p', '3', '/configure/api/v1/servers', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (1804, 'p', '3', '/configure/api/v1/template', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (1806, 'p', '3', '/configure/api/v1/template', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (1807, 'p', '3', '/configure/api/v1/template/compare', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (1805, 'p', '3', '/configure/api/v1/template/current', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (1809, 'p', '3', '/configure/api/v1/template/preview', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (1808, 'p', '3', '/configure/api/v1/template/switch', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (1803, 'p', '3', '/configure/api/v1/templates', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (1795, 'p', '3', '/configure/business/values', 'get', '', '', '');
INSERT INTO `casbin_rule` VALUES (1796, 'p', '3', '/configure/business/values', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES (1801, 'p', '3', '/configure/resource/values', 'get', '', '', '');
INSERT INTO `casbin_rule` VALUES (1802, 'p', '3', '/configure/resource/values', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES (1833, 'p', '3', '/cron/api/v1/log', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (1832, 'p', '3', '/cron/api/v1/logs', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (1830, 'p', '3', '/cron/api/v1/task', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES (1826, 'p', '3', '/cron/api/v1/task', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (1829, 'p', '3', '/cron/api/v1/task', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES (1824, 'p', '3', '/cron/api/v1/task_group', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES (1822, 'p', '3', '/cron/api/v1/task_group', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (1823, 'p', '3', '/cron/api/v1/task_group', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES (1821, 'p', '3', '/cron/api/v1/task_groups', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (1828, 'p', '3', '/cron/api/v1/task/cancel', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (1827, 'p', '3', '/cron/api/v1/task/exec', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (1831, 'p', '3', '/cron/api/v1/task/status', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES (1825, 'p', '3', '/cron/api/v1/tasks', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (1819, 'p', '3', '/cron/api/v1/worker', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES (1817, 'p', '3', '/cron/api/v1/worker', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (1818, 'p', '3', '/cron/api/v1/worker', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES (1815, 'p', '3', '/cron/api/v1/worker_group', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES (1813, 'p', '3', '/cron/api/v1/worker_group', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (1814, 'p', '3', '/cron/api/v1/worker_group', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES (1812, 'p', '3', '/cron/api/v1/worker_groups', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (1820, 'p', '3', '/cron/api/v1/worker/status', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES (1816, 'p', '3', '/cron/api/v1/workers', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (1714, 'p', '3', '/manager/api/v1/department', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES (1712, 'p', '3', '/manager/api/v1/department', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (1713, 'p', '3', '/manager/api/v1/department', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES (1695, 'p', '3', '/manager/api/v1/dictionaries', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (1698, 'p', '3', '/manager/api/v1/dictionary', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES (1696, 'p', '3', '/manager/api/v1/dictionary', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (1697, 'p', '3', '/manager/api/v1/dictionary', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES (1703, 'p', '3', '/manager/api/v1/dictionary_value', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES (1700, 'p', '3', '/manager/api/v1/dictionary_value', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (1701, 'p', '3', '/manager/api/v1/dictionary_value', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES (1702, 'p', '3', '/manager/api/v1/dictionary_value/status', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES (1699, 'p', '3', '/manager/api/v1/dictionary_values', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (1711, 'p', '3', '/manager/api/v1/job', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES (1709, 'p', '3', '/manager/api/v1/job', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (1710, 'p', '3', '/manager/api/v1/job', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES (1708, 'p', '3', '/manager/api/v1/jobs', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (1707, 'p', '3', '/manager/api/v1/menu', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES (1705, 'p', '3', '/manager/api/v1/menu', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (1706, 'p', '3', '/manager/api/v1/menu', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES (1704, 'p', '3', '/manager/api/v1/menus', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (1778, 'p', '3', '/manager/api/v1/resource/cfg_env', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES (1786, 'p', '3', '/manager/api/v1/resource/cfg_server', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES (1748, 'p', '3', '/manager/api/v1/resource/uc_app', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES (1718, 'p', '3', '/manager/api/v1/role', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES (1715, 'p', '3', '/manager/api/v1/role', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (1716, 'p', '3', '/manager/api/v1/role', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES (1720, 'p', '3', '/manager/api/v1/role/menu', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (1719, 'p', '3', '/manager/api/v1/role/menu_ids', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (1717, 'p', '3', '/manager/api/v1/role/status', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES (1724, 'p', '3', '/manager/api/v1/user', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES (1722, 'p', '3', '/manager/api/v1/user', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (1723, 'p', '3', '/manager/api/v1/user', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES (1726, 'p', '3', '/manager/api/v1/user/password/reset', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (1725, 'p', '3', '/manager/api/v1/user/status', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES (1721, 'p', '3', '/manager/api/v1/users', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (1727, 'p', '3', '/resource/api/v1/directories', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (1730, 'p', '3', '/resource/api/v1/directory', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES (1728, 'p', '3', '/resource/api/v1/directory', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (1729, 'p', '3', '/resource/api/v1/directory', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES (1736, 'p', '3', '/resource/api/v1/export', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES (1735, 'p', '3', '/resource/api/v1/export', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (1734, 'p', '3', '/resource/api/v1/exports', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (1733, 'p', '3', '/resource/api/v1/file', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES (1732, 'p', '3', '/resource/api/v1/file', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES (1731, 'p', '3', '/resource/api/v1/files', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (1752, 'p', '3', '/usercenter/api/v1/app', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES (1749, 'p', '3', '/usercenter/api/v1/app', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (1750, 'p', '3', '/usercenter/api/v1/app', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES (1751, 'p', '3', '/usercenter/api/v1/app/status', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES (1747, 'p', '3', '/usercenter/api/v1/apps', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (1770, 'p', '3', '/usercenter/api/v1/auth', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES (1768, 'p', '3', '/usercenter/api/v1/auth', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (1769, 'p', '3', '/usercenter/api/v1/auth/status', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES (1767, 'p', '3', '/usercenter/api/v1/auths', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (1741, 'p', '3', '/usercenter/api/v1/channel', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES (1738, 'p', '3', '/usercenter/api/v1/channel', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (1739, 'p', '3', '/usercenter/api/v1/channel', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES (1740, 'p', '3', '/usercenter/api/v1/channel/status', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES (1737, 'p', '3', '/usercenter/api/v1/channels', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (1760, 'p', '3', '/usercenter/api/v1/feedback', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES (1758, 'p', '3', '/usercenter/api/v1/feedback', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (1759, 'p', '3', '/usercenter/api/v1/feedback', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES (1753, 'p', '3', '/usercenter/api/v1/feedback_categories', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (1756, 'p', '3', '/usercenter/api/v1/feedback_category', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES (1754, 'p', '3', '/usercenter/api/v1/feedback_category', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (1755, 'p', '3', '/usercenter/api/v1/feedback_category', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES (1757, 'p', '3', '/usercenter/api/v1/feedbacks', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (1746, 'p', '3', '/usercenter/api/v1/field', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES (1743, 'p', '3', '/usercenter/api/v1/field', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (1744, 'p', '3', '/usercenter/api/v1/field', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES (1745, 'p', '3', '/usercenter/api/v1/field/status', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES (1742, 'p', '3', '/usercenter/api/v1/fields', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (1772, 'p', '3', '/usercenter/api/v1/oauth', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES (1771, 'p', '3', '/usercenter/api/v1/oauths', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (1766, 'p', '3', '/usercenter/api/v1/user', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES (1762, 'p', '3', '/usercenter/api/v1/user', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (1764, 'p', '3', '/usercenter/api/v1/user', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES (1765, 'p', '3', '/usercenter/api/v1/user/status', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES (1776, 'p', '3', '/usercenter/api/v1/userinfo', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES (1774, 'p', '3', '/usercenter/api/v1/userinfo', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (1775, 'p', '3', '/usercenter/api/v1/userinfo', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES (1773, 'p', '3', '/usercenter/api/v1/userinfos', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (1761, 'p', '3', '/usercenter/api/v1/users', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (1763, 'p', '3', '/usercenter/api/v1/users', 'POST', '', '', '');
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
) ENGINE=InnoDB AUTO_INCREMENT=11 DEFAULT CHARSET=utf8mb4  COMMENT='部门信息';

-- ----------------------------
-- Records of department
-- ----------------------------
BEGIN;
INSERT INTO `department` VALUES (1, 0, 'company', '贵州青橙科技有限公司', '开放合作，拥抱未来', 1713706137, 1713706137);
INSERT INTO `department` VALUES (5, 1, 'dep_1', '下级测试部门', '下级测试部门', 1720685640, 1721838411);
INSERT INTO `department` VALUES (6, 5, 'dep_2', '下级测试部门2', '下级测试部门', 1720685653, 1720685653);
INSERT INTO `department` VALUES (7, 1, 'dep3', '下级测试部门3', '下级测试部门2', 1720685663, 1720685663);
INSERT INTO `department` VALUES (8, 7, 'dep4', '下级测试部门4', '下级测试部门2', 1720685670, 1720685670);
INSERT INTO `department` VALUES (9, 6, 'dep5', '下级测试部门5', '下级测试部门2', 1720685679, 1720685679);
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
) ENGINE=InnoDB AUTO_INCREMENT=22 DEFAULT CHARSET=utf8mb4  COMMENT='部门层级信息';

-- ----------------------------
-- Records of department_closure
-- ----------------------------
BEGIN;
INSERT INTO `department_closure` VALUES (10, 1, 5);
INSERT INTO `department_closure` VALUES (11, 5, 6);
INSERT INTO `department_closure` VALUES (12, 1, 6);
INSERT INTO `department_closure` VALUES (13, 1, 7);
INSERT INTO `department_closure` VALUES (14, 7, 8);
INSERT INTO `department_closure` VALUES (15, 1, 8);
INSERT INTO `department_closure` VALUES (16, 6, 9);
INSERT INTO `department_closure` VALUES (17, 5, 9);
INSERT INTO `department_closure` VALUES (18, 1, 9);
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
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8mb4  COMMENT='字典信息';

-- ----------------------------
-- Records of dictionary
-- ----------------------------
BEGIN;
INSERT INTO `dictionary` VALUES (4, 't2', 't2', 't', 1721835689, 1721837018, 0);
INSERT INTO `dictionary` VALUES (6, '1', '1', '1', 1721964102, 1721964102, 0);
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
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4  COMMENT='字典值信息';

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
) ENGINE=InnoDB AUTO_INCREMENT=14 DEFAULT CHARSET=utf8mb4  COMMENT='职位信息';

-- ----------------------------
-- Records of job
-- ----------------------------
BEGIN;
INSERT INTO `job` VALUES (1, 'chairman', '董事长', 2, '董事长', 1713706137, 1721838228, 0);
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
) ENGINE=InnoDB AUTO_INCREMENT=7660 DEFAULT CHARSET=utf8mb4  COMMENT='菜单信息';

-- ----------------------------
-- Records of menu
-- ----------------------------
BEGIN;
INSERT INTO `menu` VALUES (7368, 0, '管理平台', 'R', 'SystemPlatform', 'apps', NULL, NULL, '/', NULL, 'Layout', NULL, 2, 0, NULL, NULL, NULL, 1724036084, 1724036084);
INSERT INTO `menu` VALUES (7369, 7368, '首页面板', 'M', 'Dashboard', 'dashboard', NULL, NULL, '/dashboard', NULL, '/dashboard/index', NULL, 0, 0, 1, 1, 1, 1724036084, 1724036084);
INSERT INTO `menu` VALUES (7370, 7368, '管理中心', 'M', 'SystemPlatformManager', 'apps', NULL, NULL, '/manager', NULL, NULL, NULL, 0, 0, NULL, NULL, NULL, 1724036084, 1724036084);
INSERT INTO `menu` VALUES (7371, 7370, '基础接口', 'G', 'ManagerBaseApi', 'apps', NULL, NULL, NULL, NULL, NULL, NULL, 0, 1, NULL, NULL, NULL, 1724036084, 1724036084);
INSERT INTO `menu` VALUES (7372, 7371, '获取用户可见部门树', 'BA', NULL, NULL, '/manager/api/v1/departments', 'GET', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL, 1724036084, 1724036084);
INSERT INTO `menu` VALUES (7373, 7371, '获取用户可见角色树', 'BA', NULL, NULL, '/manager/api/v1/roles', 'GET', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL, 1724036084, 1724036084);
INSERT INTO `menu` VALUES (7374, 7371, '获取个人用户信息', 'BA', NULL, NULL, '/manager/api/v1/user/current', 'GET', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL, 1724036084, 1724036084);
INSERT INTO `menu` VALUES (7375, 7371, '更新个人用户信息', 'BA', NULL, NULL, '/manager/api/v1/user/current/info', 'PUT', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL, 1724036084, 1724036084);
INSERT INTO `menu` VALUES (7376, 7371, '更新个人用户密码', 'BA', NULL, NULL, '/manager/api/v1/user/current/password', 'PUT', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL, 1724036084, 1724036084);
INSERT INTO `menu` VALUES (7377, 7371, '保存个人用户设置', 'BA', NULL, NULL, '/manager/api/v1/user/current/setting', 'PUT', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL, 1724036084, 1724036084);
INSERT INTO `menu` VALUES (7378, 7371, '发送用户验证吗', 'BA', NULL, NULL, '/manager/api/v1/user/current/captcha', 'POST', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL, 1724036084, 1724036084);
INSERT INTO `menu` VALUES (7379, 7371, '获取用户当前角色菜单', 'BA', NULL, NULL, '/manager/api/v1/menus/by/cur_role', 'GET', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL, 1724036084, 1724036084);
INSERT INTO `menu` VALUES (7380, 7371, '退出登录', 'BA', NULL, NULL, '/manager/api/v1/user/logout', 'POST', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL, 1724036084, 1724036084);
INSERT INTO `menu` VALUES (7381, 7371, '刷新token', 'BA', NULL, NULL, '/manager/api/v1/user/token/refresh', 'POST', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL, 1724036084, 1724036084);
INSERT INTO `menu` VALUES (7382, 7371, '用户登录', 'BA', NULL, NULL, '/manager/api/v1/user/login', 'POST', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL, 1724036084, 1724036084);
INSERT INTO `menu` VALUES (7383, 7371, '获取登录验证码', 'BA', NULL, NULL, '/manager/api/v1/user/login/captcha', 'POST', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL, 1724036084, 1724036084);
INSERT INTO `menu` VALUES (7384, 7371, '获取系统基础设置', 'BA', NULL, NULL, '/manager/api/v1/system/setting', 'GET', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL, 1724036084, 1724036084);
INSERT INTO `menu` VALUES (7385, 7371, '接口鉴权', 'BA', NULL, NULL, '/manager/api/v1/auth', 'POST', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL, 1724036084, 1724036084);
INSERT INTO `menu` VALUES (7386, 7371, '切换用户角色', 'BA', NULL, NULL, '/manager/api/v1/user/current/role', 'POST', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL, 1724036084, 1724036084);
INSERT INTO `menu` VALUES (7387, 7371, '分块上传文件', 'BA', NULL, NULL, '/resource/api/v1/upload', 'POST', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL, 1724036084, 1724036084);
INSERT INTO `menu` VALUES (7388, 7371, '预上传文件', 'BA', NULL, NULL, '/resource/api/v1/prepare_upload', 'POST', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL, 1724036084, 1724036084);
INSERT INTO `menu` VALUES (7389, 7371, '获取字段类型', 'BA', NULL, NULL, '/usercenter/api/v1/field/types', 'GET', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL, 1724036084, 1724036084);
INSERT INTO `menu` VALUES (7390, 7371, '查询资源权限', 'BA', NULL, NULL, '/manager/api/v1/resource', 'GET', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL, 1724036084, 1724036084);
INSERT INTO `menu` VALUES (7391, 7371, '获取渠道类型', 'BA', NULL, NULL, '/usercenter/api/v1/channel/types', 'GET', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL, 1724036084, 1724036084);
INSERT INTO `menu` VALUES (7392, 7370, '字典管理', 'M', 'ManagerDictionary', 'storage', NULL, NULL, '/manager/dictionary', NULL, '/manager/dictionary/index', NULL, 0, NULL, 1, NULL, 1, 1724036084, 1724036084);
INSERT INTO `menu` VALUES (7393, 7392, '查询字典', 'A', NULL, NULL, '/manager/api/v1/dictionaries', 'GET', NULL, 'manager:dictionary:query', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1724036084, 1724036084);
INSERT INTO `menu` VALUES (7394, 7392, '新增字典', 'A', NULL, NULL, '/manager/api/v1/dictionary', 'POST', NULL, 'manager:dictionary:add', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1724036084, 1724036084);
INSERT INTO `menu` VALUES (7395, 7392, '修改字典', 'A', NULL, NULL, '/manager/api/v1/dictionary', 'PUT', NULL, 'manager:dictionary:update', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1724036084, 1724036084);
INSERT INTO `menu` VALUES (7396, 7392, '删除字典', 'A', NULL, NULL, '/manager/api/v1/dictionary', 'DELETE', NULL, 'manager:dictionary:delete', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1724036084, 1724036084);
INSERT INTO `menu` VALUES (7397, 7392, '获取字典值', 'A', NULL, NULL, '/manager/api/v1/dictionary_values', 'GET', NULL, 'manager:dictionary:value:query', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1724036084, 1724036084);
INSERT INTO `menu` VALUES (7398, 7392, '新增字典值', 'A', NULL, NULL, '/manager/api/v1/dictionary_value', 'POST', NULL, 'manager:dictionary:value:add', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1724036084, 1724036084);
INSERT INTO `menu` VALUES (7399, 7392, '修改字典值', 'A', NULL, NULL, '/manager/api/v1/dictionary_value', 'PUT', NULL, 'manager:dictionary:value:update', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1724036084, 1724036084);
INSERT INTO `menu` VALUES (7400, 7392, '更新字典值目录状态', 'A', NULL, NULL, '/manager/api/v1/dictionary_value/status', 'PUT', NULL, 'manager:dictionary:value:status', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1724036084, 1724036084);
INSERT INTO `menu` VALUES (7401, 7392, '删除字典值', 'A', NULL, NULL, '/manager/api/v1/dictionary_value', 'DELETE', NULL, 'manager:dictionary:value:delete', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1724036084, 1724036084);
INSERT INTO `menu` VALUES (7402, 7370, '菜单管理', 'M', 'ManagerMenu', 'menu', NULL, NULL, '/manager/menu', NULL, '/manager/menu/index', NULL, 0, 0, 1, 1, 1, 1724036084, 1724036084);
INSERT INTO `menu` VALUES (7403, 7402, '查询菜单', 'A', NULL, NULL, '/manager/api/v1/menus', 'GET', NULL, 'manager:menu:query', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1724036084, 1724036084);
INSERT INTO `menu` VALUES (7404, 7402, '新增菜单', 'A', NULL, NULL, '/manager/api/v1/menu', 'POST', NULL, 'manager:menu:add', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1724036084, 1724036084);
INSERT INTO `menu` VALUES (7405, 7402, '修改菜单', 'A', NULL, NULL, '/manager/api/v1/menu', 'PUT', NULL, 'manager:menu:update', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1724036084, 1724036084);
INSERT INTO `menu` VALUES (7406, 7402, '删除菜单', 'A', NULL, NULL, '/manager/api/v1/menu', 'DELETE', NULL, 'manager:menu:delete', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1724036084, 1724036084);
INSERT INTO `menu` VALUES (7407, 7370, '职位管理', 'M', 'ManagerJob', 'tag', NULL, NULL, '/manager/job', NULL, '/manager/job/index', NULL, 0, NULL, 1, NULL, 1, 1724036084, 1724036084);
INSERT INTO `menu` VALUES (7408, 7407, '查询职位', 'A', NULL, NULL, '/manager/api/v1/jobs', 'GET', NULL, 'manager:job:query', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1724036084, 1724036084);
INSERT INTO `menu` VALUES (7409, 7407, '新增职位', 'A', NULL, NULL, '/manager/api/v1/job', 'POST', NULL, 'manager:job:add', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1724036084, 1724036084);
INSERT INTO `menu` VALUES (7410, 7407, '修改职位', 'A', NULL, NULL, '/manager/api/v1/job', 'PUT', NULL, 'manager:job:update', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1724036084, 1724036084);
INSERT INTO `menu` VALUES (7411, 7407, '删除职位', 'A', NULL, NULL, '/manager/api/v1/job', 'DELETE', NULL, 'manager:job:delete', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1724036084, 1724036084);
INSERT INTO `menu` VALUES (7412, 7370, '部门管理', 'M', 'ManagerDepartment', 'user-group', NULL, NULL, '/manager/department', NULL, '/manager/department/index', NULL, 0, 0, 1, NULL, 1, 1724036084, 1724036084);
INSERT INTO `menu` VALUES (7413, 7412, '新增部门', 'A', NULL, NULL, '/manager/api/v1/department', 'POST', NULL, 'manager:department:add', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1724036084, 1724036084);
INSERT INTO `menu` VALUES (7414, 7412, '修改部门', 'A', NULL, NULL, '/manager/api/v1/department', 'PUT', NULL, 'manager:department:update', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1724036084, 1724036084);
INSERT INTO `menu` VALUES (7415, 7412, '删除部门', 'A', NULL, NULL, '/manager/api/v1/department', 'DELETE', NULL, 'manager:department:delete', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1724036084, 1724036084);
INSERT INTO `menu` VALUES (7416, 7370, '角色管理', 'M', 'ManagerRole', 'safe', NULL, NULL, '/manager/role', NULL, '/manager/role/index', NULL, 0, NULL, 0, NULL, 1, 1724036084, 1724036084);
INSERT INTO `menu` VALUES (7417, 7416, '新增角色', 'A', NULL, NULL, '/manager/api/v1/role', 'POST', NULL, 'manager:role:add', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1724036084, 1724036084);
INSERT INTO `menu` VALUES (7418, 7416, '修改角色', 'A', NULL, NULL, '/manager/api/v1/role', 'PUT', NULL, 'manager:role:update', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1724036084, 1724036084);
INSERT INTO `menu` VALUES (7419, 7416, '修改角色状态', 'A', NULL, NULL, '/manager/api/v1/role/status', 'PUT', NULL, 'manager:role:update:status', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1724036084, 1724036084);
INSERT INTO `menu` VALUES (7420, 7416, '删除角色', 'A', NULL, NULL, '/manager/api/v1/role', 'DELETE', NULL, 'manager:role:delete', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1724036084, 1724036084);
INSERT INTO `menu` VALUES (7421, 7416, '角色菜单管理', 'G', NULL, NULL, NULL, NULL, NULL, 'manager:role:menu', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1724036084, 1724036084);
INSERT INTO `menu` VALUES (7422, 7421, '查询角色菜单', 'A', NULL, NULL, '/manager/api/v1/role/menu_ids', 'GET', NULL, 'manager:role:menu:query', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1724036084, 1724036084);
INSERT INTO `menu` VALUES (7423, 7421, '修改角色菜单', 'A', NULL, NULL, '/manager/api/v1/role/menu', 'POST', NULL, 'manager:role:menu:update', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1724036084, 1724036084);
INSERT INTO `menu` VALUES (7424, 7370, '用户管理', 'M', 'ManagerUser', 'user', NULL, NULL, '/manager/user', NULL, '/manager/user/index', NULL, 0, NULL, 1, NULL, 1, 1724036084, 1724036084);
INSERT INTO `menu` VALUES (7425, 7424, '查询用户列表', 'A', NULL, NULL, '/manager/api/v1/users', 'GET', NULL, 'manager:user:query', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1724036084, 1724036084);
INSERT INTO `menu` VALUES (7426, 7424, '新增用户', 'A', NULL, NULL, '/manager/api/v1/user', 'POST', NULL, 'manager:user:add', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1724036084, 1724036084);
INSERT INTO `menu` VALUES (7427, 7424, '修改用户', 'A', NULL, NULL, '/manager/api/v1/user', 'PUT', NULL, 'manager:user:update', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1724036084, 1724036084);
INSERT INTO `menu` VALUES (7428, 7424, '删除用户', 'A', NULL, NULL, '/manager/api/v1/user', 'DELETE', NULL, 'manager:user:delete', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1724036084, 1724036084);
INSERT INTO `menu` VALUES (7429, 7424, '修改用户状态', 'A', NULL, NULL, '/manager/api/v1/user/status', 'PUT', NULL, 'manager:user:status', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1724036084, 1724036084);
INSERT INTO `menu` VALUES (7430, 7424, '重置账号密码', 'A', NULL, NULL, '/manager/api/v1/user/password/reset', 'POST', NULL, 'manager:user:reset:password', '', NULL, 0, NULL, NULL, NULL, NULL, 1724036084, 1724036084);
INSERT INTO `menu` VALUES (7431, 7368, '资源中心', 'M', 'SystemPlatformResource', 'file', NULL, NULL, '/resource', NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL, 1724036084, 1724036084);
INSERT INTO `menu` VALUES (7432, 7431, '文件管理', 'M', 'ResourceFile', 'file', NULL, NULL, '/resource/file', NULL, '/resource/file/index', NULL, 0, 0, 1, NULL, 1, 1724036084, 1724036084);
INSERT INTO `menu` VALUES (7433, 7432, '目录管理', 'G', NULL, NULL, NULL, NULL, NULL, 'resource:directory:group', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1724036084, 1724036084);
INSERT INTO `menu` VALUES (7434, 7433, '查看目录', 'A', NULL, NULL, '/resource/api/v1/directories', 'GET', NULL, 'resource:directory:query', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1724036084, 1724036084);
INSERT INTO `menu` VALUES (7435, 7433, '新增目录', 'A', NULL, NULL, '/resource/api/v1/directory', 'POST', NULL, 'resource:directory:add', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1724036084, 1724036084);
INSERT INTO `menu` VALUES (7436, 7433, '修改目录', 'A', NULL, NULL, '/resource/api/v1/directory', 'PUT', NULL, 'resource:directory:update', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1724036084, 1724036084);
INSERT INTO `menu` VALUES (7437, 7433, '删除目录', 'A', NULL, NULL, '/resource/api/v1/directory', 'DELETE', NULL, 'resource:directory:delete', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1724036084, 1724036084);
INSERT INTO `menu` VALUES (7438, 7432, '文件管理', 'G', NULL, NULL, NULL, NULL, NULL, 'resource:file:group', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1724036084, 1724036084);
INSERT INTO `menu` VALUES (7439, 7438, '查看文件', 'A', NULL, NULL, '/resource/api/v1/files', 'GET', NULL, 'resource:file:query', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1724036084, 1724036084);
INSERT INTO `menu` VALUES (7440, 7438, '修改文件', 'A', NULL, NULL, '/resource/api/v1/file', 'PUT', NULL, 'resource:file:update', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1724036084, 1724036084);
INSERT INTO `menu` VALUES (7441, 7438, '删除文件', 'A', NULL, NULL, '/resource/api/v1/file', 'DELETE', NULL, 'resource:file:delete', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1724036084, 1724036084);
INSERT INTO `menu` VALUES (7442, 7431, '导出管理', 'M', 'ResourceExport', 'export', NULL, NULL, '/resource/export', NULL, '/resource/export/index', NULL, 0, 0, 1, NULL, 1, 1724036084, 1724036084);
INSERT INTO `menu` VALUES (7443, 7442, '查看导出', 'A', NULL, NULL, '/resource/api/v1/exports', 'GET', NULL, 'resource:export:query', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1724036084, 1724036084);
INSERT INTO `menu` VALUES (7444, 7442, '新增导出', 'A', NULL, NULL, '/resource/api/v1/export', 'POST', NULL, 'resource:export:file', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1724036084, 1724036084);
INSERT INTO `menu` VALUES (7445, 7442, '删除导出', 'A', NULL, NULL, '/resource/api/v1/export', 'DELETE', NULL, 'resource:export:delete', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1724036084, 1724036084);
INSERT INTO `menu` VALUES (7446, 7368, '用户中心', 'M', 'SystemPlatformUserCenter', 'user', NULL, NULL, '/usercenter', NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL, 1724036084, 1724036084);
INSERT INTO `menu` VALUES (7447, 7446, '授权渠道', 'M', 'UserCenterChannel', 'mind-mapping', NULL, NULL, '/usercenter/channel', NULL, '/usercenter/channel/index', NULL, 0, 0, 1, NULL, 1, 1724036084, 1724036084);
INSERT INTO `menu` VALUES (7448, 7447, '查看渠道', 'A', NULL, NULL, '/usercenter/api/v1/channels', 'GET', NULL, 'uc:channel:query', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1724036084, 1724036084);
INSERT INTO `menu` VALUES (7449, 7447, '新增渠道', 'A', NULL, NULL, '/usercenter/api/v1/channel', 'POST', NULL, 'uc:channel:add', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1724036084, 1724036084);
INSERT INTO `menu` VALUES (7450, 7447, '修改渠道', 'A', NULL, NULL, '/usercenter/api/v1/channel', 'PUT', NULL, 'uc:channel:update', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1724036084, 1724036084);
INSERT INTO `menu` VALUES (7451, 7447, '修改渠道状态', 'A', NULL, NULL, '/usercenter/api/v1/channel/status', 'PUT', NULL, 'uc:channel:update:status', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1724036084, 1724036084);
INSERT INTO `menu` VALUES (7452, 7447, '删除渠道', 'A', NULL, NULL, '/usercenter/api/v1/channel', 'DELETE', NULL, 'uc:channel:delete', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1724036084, 1724036084);
INSERT INTO `menu` VALUES (7453, 7446, '信息字段', 'M', 'UserCenterField', 'list', NULL, NULL, '/usercenter/field', NULL, '/usercenter/field/index', NULL, 0, 0, 1, 0, 1, 1724036084, 1724036084);
INSERT INTO `menu` VALUES (7454, 7453, '查看字段列表', 'A', NULL, NULL, '/usercenter/api/v1/fields', 'GET', NULL, 'uc:field:query', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1724036084, 1724036084);
INSERT INTO `menu` VALUES (7455, 7453, '新增字段', 'A', NULL, NULL, '/usercenter/api/v1/field', 'POST', NULL, 'uc:field:add', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1724036084, 1724036084);
INSERT INTO `menu` VALUES (7456, 7453, '修改字段', 'A', NULL, NULL, '/usercenter/api/v1/field', 'PUT', NULL, 'uc:field:update', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1724036084, 1724036084);
INSERT INTO `menu` VALUES (7457, 7453, '修改字段状态', 'A', NULL, NULL, '/usercenter/api/v1/field/status', 'PUT', NULL, 'uc:field:update:status', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1724036084, 1724036084);
INSERT INTO `menu` VALUES (7458, 7453, '删除字段', 'A', NULL, NULL, '/usercenter/api/v1/field', 'DELETE', NULL, 'uc:field:delete', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1724036084, 1724036084);
INSERT INTO `menu` VALUES (7459, 7446, '应用管理', 'M', 'UserCenterApp', 'apps', NULL, NULL, '/usercenter/app', NULL, '/usercenter/app/index', NULL, 0, 0, 1, NULL, 1, 1724036084, 1724036084);
INSERT INTO `menu` VALUES (7460, 7459, '查看应用', 'A', NULL, NULL, '/usercenter/api/v1/apps', 'GET', NULL, 'uc:app:query', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1724036084, 1724036084);
INSERT INTO `menu` VALUES (7461, 7459, '设置应用资源权限', 'A', NULL, NULL, '/manager/api/v1/resource/uc_app', 'PUT', NULL, 'uc:app:resource:permission', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1724036084, 1724036084);
INSERT INTO `menu` VALUES (7462, 7459, '新增应用', 'A', NULL, NULL, '/usercenter/api/v1/app', 'POST', NULL, 'uc:app:add', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1724036084, 1724036084);
INSERT INTO `menu` VALUES (7463, 7459, '修改应用', 'A', NULL, NULL, '/usercenter/api/v1/app', 'PUT', NULL, 'uc:app:update', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1724036084, 1724036084);
INSERT INTO `menu` VALUES (7464, 7459, '修改应用状态', 'A', NULL, NULL, '/usercenter/api/v1/app/status', 'PUT', NULL, 'uc:app:update:status', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1724036084, 1724036084);
INSERT INTO `menu` VALUES (7465, 7459, '删除应用', 'A', NULL, NULL, '/usercenter/api/v1/app', 'DELETE', NULL, 'uc:app:delete', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1724036084, 1724036084);
INSERT INTO `menu` VALUES (7466, 7446, '反馈管理', 'M', 'UserCenterFeedback', 'question-circle', NULL, NULL, '/usercenter/feedback', NULL, '/usercenter/feedback/index', NULL, 0, 0, 1, NULL, 1, 1724036084, 1724036084);
INSERT INTO `menu` VALUES (7467, 7466, '查看反馈分类', 'A', NULL, NULL, '/usercenter/api/v1/feedback_categories', 'GET', NULL, 'uc:feedback:category:query', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1724036084, 1724036084);
INSERT INTO `menu` VALUES (7468, 7466, '新增反馈渠道', 'A', NULL, NULL, '/usercenter/api/v1/feedback_category', 'POST', NULL, 'uc:feedback:category:add', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1724036084, 1724036084);
INSERT INTO `menu` VALUES (7469, 7466, '修改反馈渠道', 'A', NULL, NULL, '/usercenter/api/v1/feedback_category', 'PUT', NULL, 'uc:feedback:category:update', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1724036084, 1724036084);
INSERT INTO `menu` VALUES (7470, 7466, '删除反馈渠道', 'A', NULL, NULL, '/usercenter/api/v1/feedback_category', 'DELETE', NULL, 'uc:feedback:category:delete', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1724036084, 1724036084);
INSERT INTO `menu` VALUES (7471, 7466, '查看反馈', 'A', NULL, NULL, '/usercenter/api/v1/feedbacks', 'GET', NULL, 'uc:feedback:query', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1724036084, 1724036084);
INSERT INTO `menu` VALUES (7472, 7466, '新增反馈', 'A', NULL, NULL, '/usercenter/api/v1/feedback', 'POST', NULL, 'uc:feedback:add', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1724036084, 1724036084);
INSERT INTO `menu` VALUES (7473, 7466, '修改反馈', 'A', NULL, NULL, '/usercenter/api/v1/feedback', 'PUT', NULL, 'uc:feedback:update', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1724036084, 1724036084);
INSERT INTO `menu` VALUES (7474, 7466, '删除反馈', 'A', NULL, NULL, '/usercenter/api/v1/feedback', 'DELETE', NULL, 'uc:feedback:delete', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1724036084, 1724036084);
INSERT INTO `menu` VALUES (7475, 7446, '用户管理', 'M', 'UserCenterUser', 'user', NULL, NULL, '/usercenter/user', NULL, '/usercenter/user/index', NULL, 0, 0, 1, NULL, 1, 1724036084, 1724036084);
INSERT INTO `menu` VALUES (7476, 7475, '查看用户', 'A', NULL, NULL, '/usercenter/api/v1/users', 'GET', NULL, 'uc:user:query', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1724036084, 1724036084);
INSERT INTO `menu` VALUES (7477, 7475, '新增用户', 'A', NULL, NULL, '/usercenter/api/v1/user', 'POST', NULL, 'uc:user:add', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1724036084, 1724036084);
INSERT INTO `menu` VALUES (7478, 7475, '导入用户', 'A', NULL, NULL, '/usercenter/api/v1/users', 'POST', NULL, 'uc:user:import', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1724036084, 1724036084);
INSERT INTO `menu` VALUES (7479, 7475, '修改用户', 'A', NULL, NULL, '/usercenter/api/v1/user', 'PUT', NULL, 'uc:user:update', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1724036084, 1724036084);
INSERT INTO `menu` VALUES (7480, 7475, '修改用户状态', 'A', NULL, NULL, '/usercenter/api/v1/user/status', 'PUT', NULL, 'uc:user:update:status', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1724036084, 1724036084);
INSERT INTO `menu` VALUES (7481, 7475, '删除用户', 'A', NULL, NULL, '/usercenter/api/v1/user', 'DELETE', NULL, 'uc:user:delete', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1724036084, 1724036084);
INSERT INTO `menu` VALUES (7482, 7475, '查看用户详细信息', 'G', NULL, NULL, NULL, NULL, NULL, 'uc:user:more', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1724036084, 1724036084);
INSERT INTO `menu` VALUES (7483, 7482, '获取用户应用信息', 'A', NULL, NULL, '/usercenter/api/v1/auths', 'GET', NULL, 'uc:auth:query', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1724036084, 1724036084);
INSERT INTO `menu` VALUES (7484, 7482, '创建用户应用信息', 'A', NULL, NULL, '/usercenter/api/v1/auth', 'POST', NULL, 'uc:auth:add', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1724036084, 1724036084);
INSERT INTO `menu` VALUES (7485, 7482, '修改用户应用状态信息', 'A', NULL, NULL, '/usercenter/api/v1/auth/status', 'PUT', NULL, 'uc:auth:update:status', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1724036084, 1724036084);
INSERT INTO `menu` VALUES (7486, 7482, '删除用户应用信息', 'A', NULL, NULL, '/usercenter/api/v1/auth', 'DELETE', NULL, 'uc:auth:delete', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1724036084, 1724036084);
INSERT INTO `menu` VALUES (7487, 7482, '获取用户渠道信息', 'A', NULL, NULL, '/usercenter/api/v1/oauths', 'GET', NULL, 'uc:oauth:query', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1724036084, 1724036084);
INSERT INTO `menu` VALUES (7488, 7482, '删除用户渠道信息', 'A', NULL, NULL, '/usercenter/api/v1/oauth', 'DELETE', NULL, 'uc:oauth:delete', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1724036084, 1724036084);
INSERT INTO `menu` VALUES (7489, 7482, '获取用户扩展信息', 'A', NULL, NULL, '/usercenter/api/v1/userinfos', 'GET', NULL, 'uc:userinfo:query', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1724036084, 1724036084);
INSERT INTO `menu` VALUES (7490, 7482, '创建用户扩展信息', 'A', NULL, NULL, '/usercenter/api/v1/userinfo', 'POST', NULL, 'uc:userinfo:add', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1724036084, 1724036084);
INSERT INTO `menu` VALUES (7491, 7482, '修改用户扩展信息', 'A', NULL, NULL, '/usercenter/api/v1/userinfo', 'PUT', NULL, 'uc:userinfo:update', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1724036084, 1724036084);
INSERT INTO `menu` VALUES (7492, 7482, '删除用户扩展信息', 'A', NULL, NULL, '/usercenter/api/v1/userinfo', 'DELETE', NULL, 'uc:userinfo:delete', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1724036084, 1724036084);
INSERT INTO `menu` VALUES (7493, 7368, '配置中心', 'M', 'SystemPlatformConfigure', 'code-block', NULL, NULL, '/configure', NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL, 1724036084, 1724036084);
INSERT INTO `menu` VALUES (7494, 7493, '环境管理', 'M', 'ConfigureEnv', 'common', NULL, NULL, '/configure/env', NULL, '/configure/env/index', NULL, 0, NULL, 1, NULL, 1, 1724036084, 1724036084);
INSERT INTO `menu` VALUES (7495, 7494, '查看环境', 'A', NULL, NULL, '/configure/api/v1/envs', 'GET', NULL, 'configure:env:query', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1724036084, 1724036084);
INSERT INTO `menu` VALUES (7496, 7494, '设置环境资源权限', 'A', NULL, NULL, '/manager/api/v1/resource/cfg_env', 'PUT', NULL, 'configure:env:resource:permission', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1724036084, 1724036084);
INSERT INTO `menu` VALUES (7497, 7494, '新增环境', 'A', NULL, NULL, '/configure/api/v1/env', 'POST', NULL, 'configure:env:add', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1724036084, 1724036084);
INSERT INTO `menu` VALUES (7498, 7494, '修改环境', 'A', NULL, NULL, '/configure/api/v1/env', 'PUT', NULL, 'configure:env:update', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1724036084, 1724036084);
INSERT INTO `menu` VALUES (7499, 7494, '修改环境状态', 'A', NULL, NULL, '/configure/api/v1/env/status', 'PUT', NULL, 'configure:env:update:status', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1724036084, 1724036084);
INSERT INTO `menu` VALUES (7500, 7494, '删除环境', 'A', NULL, NULL, '/configure/api/v1/env', 'DELETE', NULL, 'configure:env:delete', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1724036084, 1724036084);
INSERT INTO `menu` VALUES (7501, 7494, '查看环境Token', 'A', NULL, NULL, '/configure/api/v1/env/token', 'GET', NULL, 'configure:env:token:query', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1724036084, 1724036084);
INSERT INTO `menu` VALUES (7502, 7494, '重置环境token', 'A', NULL, NULL, '/configure/api/v1/env/token', 'PUT', NULL, 'configure:env:token:reset', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1724036084, 1724036084);
INSERT INTO `menu` VALUES (7503, 7493, '服务管理', 'M', 'ConfigureServer', 'apps', NULL, NULL, '/configure/server', NULL, '/configure/server/index', NULL, 0, NULL, 1, NULL, 1, 1724036084, 1724036084);
INSERT INTO `menu` VALUES (7504, 7503, '查询服务', 'A', NULL, NULL, '/configure/api/v1/servers', 'GET', NULL, 'configure:server:query', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1724036084, 1724036084);
INSERT INTO `menu` VALUES (7505, 7503, '设置服务资源权限', 'A', NULL, NULL, '/manager/api/v1/resource/cfg_server', 'PUT', NULL, 'configure:server:resource:permission', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1724036084, 1724036084);
INSERT INTO `menu` VALUES (7506, 7503, '新增服务', 'A', NULL, NULL, '/configure/api/v1/server', 'POST', NULL, 'configure:server:add', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1724036084, 1724036084);
INSERT INTO `menu` VALUES (7507, 7503, '修改服务', 'A', NULL, NULL, '/configure/api/v1/server', 'PUT', NULL, 'configure:server:update', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1724036084, 1724036084);
INSERT INTO `menu` VALUES (7508, 7503, '修改服务状态', 'A', NULL, NULL, '/configure/api/v1/server/status', 'PUT', NULL, 'configure:server:update:status', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1724036084, 1724036084);
INSERT INTO `menu` VALUES (7509, 7503, '删除服务', 'A', NULL, NULL, '/configure/api/v1/server', 'DELETE', NULL, 'configure:server:delete', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1724036084, 1724036084);
INSERT INTO `menu` VALUES (7510, 7493, '业务变量', 'M', 'ConfigureBusiness', 'code', NULL, NULL, '/configure/business', NULL, '/configure/business/index', NULL, 0, NULL, 1, NULL, 1, 1724036084, 1724036084);
INSERT INTO `menu` VALUES (7511, 7510, '查看业务变量', 'A', NULL, NULL, '/configure/api/v1/businesses', 'GET', NULL, 'configure:business:query', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1724036084, 1724036084);
INSERT INTO `menu` VALUES (7512, 7510, '新增业务变量', 'A', NULL, NULL, '/configure/api/v1/business', 'POST', NULL, 'configure:business:add', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1724036084, 1724036084);
INSERT INTO `menu` VALUES (7513, 7510, '修改业务变量', 'A', NULL, NULL, '/configure/api/v1/business', 'PUT', NULL, 'configure:business:update', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1724036084, 1724036084);
INSERT INTO `menu` VALUES (7514, 7510, '删除业务变量', 'A', NULL, NULL, '/configure/api/v1/business', 'DELETE', NULL, 'configure:business:delete', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1724036084, 1724036084);
INSERT INTO `menu` VALUES (7515, 7510, '查看业务变量值', 'A', NULL, NULL, '/configure/business/values', 'get', NULL, 'configure:business:value:query', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1724036084, 1724036084);
INSERT INTO `menu` VALUES (7516, 7510, '设置业务变量值', 'A', NULL, NULL, '/configure/business/values', 'PUT', NULL, 'configure:business:value:update', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1724036084, 1724036084);
INSERT INTO `menu` VALUES (7517, 7493, '资源变量', 'M', 'ConfigureResource', 'unordered-list', NULL, NULL, '/configure/resource', NULL, '/configure/resource/index', NULL, 0, NULL, 1, NULL, 1, 1724036084, 1724036084);
INSERT INTO `menu` VALUES (7518, 7517, '查看资源', 'A', NULL, NULL, '/configure/api/v1/resources', 'GET', NULL, 'configure:resource:query', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1724036084, 1724036084);
INSERT INTO `menu` VALUES (7519, 7517, '新增资源', 'A', NULL, NULL, '/configure/api/v1/resource', 'POST', NULL, 'configure:resource:add', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1724036084, 1724036084);
INSERT INTO `menu` VALUES (7520, 7517, '修改资源', 'A', NULL, NULL, '/configure/api/v1/resource', 'PUT', NULL, 'configure:resource:update', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1724036084, 1724036084);
INSERT INTO `menu` VALUES (7521, 7517, '删除资源', 'A', NULL, NULL, '/configure/api/v1/resource', 'DELETE', NULL, 'configure:resource:delete', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1724036084, 1724036084);
INSERT INTO `menu` VALUES (7522, 7517, '查看业务变量值', 'A', NULL, NULL, '/configure/resource/values', 'get', NULL, 'configure:resource:value:query', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1724036084, 1724036084);
INSERT INTO `menu` VALUES (7523, 7517, '设置业务变量值', 'A', NULL, NULL, '/configure/resource/values', 'PUT', NULL, 'configure:resource:value:update', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1724036084, 1724036084);
INSERT INTO `menu` VALUES (7524, 7493, '配置模板', 'M', 'ConfgureTemplate', 'file', NULL, NULL, '/configure/template', NULL, '/configure/template/index', NULL, 0, NULL, 1, NULL, 1, 1724036084, 1724036084);
INSERT INTO `menu` VALUES (7525, 7524, '模板管理', 'G', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL, 1724036084, 1724036084);
INSERT INTO `menu` VALUES (7526, 7525, '查看模板历史版本', 'A', NULL, NULL, '/configure/api/v1/templates', 'GET', NULL, 'configure:template:history', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1724036084, 1724036084);
INSERT INTO `menu` VALUES (7527, 7525, '查看指定模板详细数据', 'A', NULL, NULL, '/configure/api/v1/template', 'GET', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL, 1724036084, 1724036084);
INSERT INTO `menu` VALUES (7528, 7525, '查看当前正在使用的模板', 'A', NULL, NULL, '/configure/api/v1/template/current', 'GET', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL, 1724036084, 1724036084);
INSERT INTO `menu` VALUES (7529, 7525, '提交模板', 'A', NULL, NULL, '/configure/api/v1/template', 'POST', NULL, 'configure:template:add', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1724036084, 1724036084);
INSERT INTO `menu` VALUES (7530, 7525, '模板对比', 'A', NULL, NULL, '/configure/api/v1/template/compare', 'POST', NULL, 'configure:template:compare', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1724036084, 1724036084);
INSERT INTO `menu` VALUES (7531, 7525, '切换模板', 'A', NULL, NULL, '/configure/api/v1/template/switch', 'POST', NULL, 'configure:template:switch', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1724036084, 1724036084);
INSERT INTO `menu` VALUES (7532, 7525, '模板预览', 'A', NULL, NULL, '/configure/api/v1/template/preview', 'POST', NULL, 'configure:template:preview', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1724036084, 1724036084);
INSERT INTO `menu` VALUES (7533, 7524, '配置管理', 'G', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL, 1724036084, 1724036084);
INSERT INTO `menu` VALUES (7534, 7533, '配置对比', 'A', NULL, NULL, '/configure/api/v1/configure/compare', 'POST', NULL, 'configure:configure:compare', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1724036084, 1724036084);
INSERT INTO `menu` VALUES (7535, 7533, '同步配置', 'A', NULL, NULL, '/configure/api/v1/configure', 'PUT', NULL, 'configure:configure:sync', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1724036084, 1724036084);
INSERT INTO `menu` VALUES (7536, 7368, '定时任务', 'M', 'SystemPlatformCron', 'schedule', NULL, NULL, '/cron', NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL, 1724036084, 1724036084);
INSERT INTO `menu` VALUES (7537, 7536, '节点管理', 'M', 'Worker', 'common', NULL, NULL, '/cron/worker', NULL, '/cron/worker/index', NULL, 0, 0, 1, NULL, 1, 1724036084, 1724036084);
INSERT INTO `menu` VALUES (7538, 7537, '查看节点分组', 'A', NULL, NULL, '/cron/api/v1/worker_groups', 'GET', NULL, 'cron:worker:group:query', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1724036084, 1724036084);
INSERT INTO `menu` VALUES (7539, 7537, '新增节点分组', 'A', NULL, NULL, '/cron/api/v1/worker_group', 'POST', NULL, 'cron:worker:group:add', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1724036084, 1724036084);
INSERT INTO `menu` VALUES (7540, 7537, '修改节点分组', 'A', NULL, NULL, '/cron/api/v1/worker_group', 'PUT', NULL, 'cron:worker:group:update', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1724036084, 1724036084);
INSERT INTO `menu` VALUES (7541, 7537, '删除节点分组', 'A', NULL, NULL, '/cron/api/v1/worker_group', 'DELETE', NULL, 'cron:worker:group:delete', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1724036084, 1724036084);
INSERT INTO `menu` VALUES (7542, 7537, '查看节点', 'A', NULL, NULL, '/cron/api/v1/workers', 'GET', NULL, 'cron:worker:query', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1724036084, 1724036084);
INSERT INTO `menu` VALUES (7543, 7537, '新增节点', 'A', NULL, NULL, '/cron/api/v1/worker', 'POST', NULL, 'cron:worker:add', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1724036084, 1724036084);
INSERT INTO `menu` VALUES (7544, 7537, '修改节点', 'A', NULL, NULL, '/cron/api/v1/worker', 'PUT', NULL, 'cron:worker:update', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1724036084, 1724036084);
INSERT INTO `menu` VALUES (7545, 7537, '删除节点', 'A', NULL, NULL, '/cron/api/v1/worker', 'DELETE', NULL, 'cron:worker:delete', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1724036084, 1724036084);
INSERT INTO `menu` VALUES (7546, 7537, '更新节点状态', 'A', NULL, NULL, '/cron/api/v1/worker/status', 'PUT', NULL, 'cron:worker:update:status', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1724036084, 1724036084);
INSERT INTO `menu` VALUES (7547, 7536, '任务管理', 'M', 'Task', 'computer', NULL, NULL, '/cron/task', NULL, '/cron/task/index', NULL, 0, 0, 1, NULL, 1, 1724036084, 1724036084);
INSERT INTO `menu` VALUES (7548, 7547, '查看任务分组', 'A', NULL, NULL, '/cron/api/v1/task_groups', 'GET', NULL, 'cron:task:group:query', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1724036084, 1724036084);
INSERT INTO `menu` VALUES (7549, 7547, '新增任务分组', 'A', NULL, NULL, '/cron/api/v1/task_group', 'POST', NULL, 'cron:task:group:add', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1724036084, 1724036084);
INSERT INTO `menu` VALUES (7550, 7547, '修改任务分组', 'A', NULL, NULL, '/cron/api/v1/task_group', 'PUT', NULL, 'cron:task:group:update', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1724036084, 1724036084);
INSERT INTO `menu` VALUES (7551, 7547, '删除任务分组', 'A', NULL, NULL, '/cron/api/v1/task_group', 'DELETE', NULL, 'cron:task:group:delete', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1724036084, 1724036084);
INSERT INTO `menu` VALUES (7552, 7547, '查看任务', 'A', NULL, NULL, '/cron/api/v1/tasks', 'GET', NULL, 'cron:task:query', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1724036084, 1724036084);
INSERT INTO `menu` VALUES (7553, 7547, '新增任务', 'A', NULL, NULL, '/cron/api/v1/task', 'POST', NULL, 'cron:task:add', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1724036084, 1724036084);
INSERT INTO `menu` VALUES (7554, 7547, '立即执行任务', 'A', NULL, NULL, '/cron/api/v1/task/exec', 'POST', NULL, 'cron:task:exec', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1724036084, 1724036084);
INSERT INTO `menu` VALUES (7555, 7547, '取消执行任务', 'A', NULL, NULL, '/cron/api/v1/task/cancel', 'POST', NULL, 'cron:task:cancel', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1724036084, 1724036084);
INSERT INTO `menu` VALUES (7556, 7547, '修改任务', 'A', NULL, NULL, '/cron/api/v1/task', 'PUT', NULL, 'cron:task:update', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1724036084, 1724036084);
INSERT INTO `menu` VALUES (7557, 7547, '删除任务', 'A', NULL, NULL, '/cron/api/v1/task', 'DELETE', NULL, 'cron:task:delete', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1724036084, 1724036084);
INSERT INTO `menu` VALUES (7558, 7547, '任务状态管理', 'A', NULL, NULL, '/cron/api/v1/task/status', 'PUT', NULL, 'cron:task:update:status', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1724036084, 1724036084);
INSERT INTO `menu` VALUES (7559, 7547, '任务日志', 'G', NULL, NULL, NULL, NULL, NULL, 'cron:task:log', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1724036084, 1724036084);
INSERT INTO `menu` VALUES (7560, 7559, '获取任务日志分页', 'A', NULL, NULL, '/cron/api/v1/logs', 'GET', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL, 1724036084, 1724036084);
INSERT INTO `menu` VALUES (7561, 7559, '获取任务日志详情', 'A', NULL, NULL, '/cron/api/v1/log', 'GET', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL, 1724036084, 1724036084);
INSERT INTO `menu` VALUES (7562, 0, '应用平台', 'R', 'AppPlatform', 'apps', NULL, NULL, '/app', NULL, 'Layout', NULL, 2, 0, NULL, NULL, NULL, 1724036084, 1724036084);
INSERT INTO `menu` VALUES (7563, 7562, '首页面板', 'M', 'AppDashboard', 'dashboard', NULL, NULL, '/app/dashboard', NULL, '/dashboard/index', NULL, 0, 0, 1, 1, 1, 1724036084, 1724036084);
INSERT INTO `menu` VALUES (7564, 7562, '信号灯系统', 'M', 'PartyAffairs', 'apps', NULL, NULL, '/party-affairs', NULL, NULL, NULL, 0, 0, NULL, NULL, NULL, 1724036084, 1724036084);
INSERT INTO `menu` VALUES (7565, 7564, '轮播管理', 'M', 'PartyAffairsBanner', 'list', NULL, NULL, '/partyaffairs/banner', NULL, '/partyaffairs/banner/index', NULL, 0, NULL, 1, NULL, 1, 1724036084, 1724036084);
INSERT INTO `menu` VALUES (7566, 7565, '查看轮播', 'A', NULL, NULL, '/partyaffairs/api/v1/banners', 'GET', NULL, 'partyaffairs:banner:query', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1724036084, 1724036084);
INSERT INTO `menu` VALUES (7567, 7565, '新增轮播', 'A', NULL, NULL, '/partyaffairs/api/v1/banner', 'POST', NULL, 'partyaffairs:banner:add', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1724036084, 1724036084);
INSERT INTO `menu` VALUES (7568, 7565, '修改轮播', 'A', NULL, NULL, '/partyaffairs/api/v1/banner', 'PUT', NULL, 'partyaffairs:banner:update', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1724036084, 1724036084);
INSERT INTO `menu` VALUES (7569, 7565, '删除轮播', 'A', NULL, NULL, '/partyaffairs/api/v1/banner', 'DELETE', NULL, 'partyaffairs:banner:delete', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1724036084, 1724036084);
INSERT INTO `menu` VALUES (7570, 7564, '通知管理', 'M', 'PartyAffairsNotice', 'notification', NULL, NULL, '/partyaffairs/notice', NULL, '/partyaffairs/notice/index', NULL, 0, NULL, 1, NULL, 1, 1724036084, 1724036084);
INSERT INTO `menu` VALUES (7571, 7570, '查看通知', 'A', NULL, NULL, '/partyaffairs/api/v1/notices', 'GET', NULL, 'partyaffairs:notice:query', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1724036084, 1724036084);
INSERT INTO `menu` VALUES (7572, 7570, '查看通知阅读情况', 'A', NULL, NULL, '/partyaffairs/api/v1/notice/users', 'GET', NULL, 'partyaffairs:notice:user:query', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1724036084, 1724036084);
INSERT INTO `menu` VALUES (7573, 7570, '新增通知', 'A', NULL, NULL, '/partyaffairs/api/v1/notice', 'POST', NULL, 'partyaffairs:notice:add', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1724036084, 1724036084);
INSERT INTO `menu` VALUES (7574, 7570, '修改通知', 'A', NULL, NULL, '/partyaffairs/api/v1/notice', 'PUT', NULL, 'partyaffairs:notice:update', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1724036084, 1724036084);
INSERT INTO `menu` VALUES (7575, 7570, '删除通知', 'A', NULL, NULL, '/partyaffairs/api/v1/notice', 'DELETE', NULL, 'partyaffairs:notice:delete', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1724036084, 1724036084);
INSERT INTO `menu` VALUES (7576, 7564, '新闻管理', 'M', 'PartyAffairsNews', 'book', NULL, NULL, '/partyaffairs/news', NULL, NULL, NULL, 0, 0, NULL, NULL, NULL, 1724036084, 1724036084);
INSERT INTO `menu` VALUES (7577, 7576, '板块管理', 'M', 'PartyAffairsNewsClassify', 'menu', NULL, NULL, '/partyaffairs/news/classify', NULL, '/partyaffairs/news/classify/index', NULL, 0, NULL, 1, NULL, 1, 1724036084, 1724036084);
INSERT INTO `menu` VALUES (7578, 7577, '查看板块', 'A', NULL, NULL, '/partyaffairs/api/v1/news/classify', 'GET', NULL, 'partyaffairs:news:classify:query', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1724036084, 1724036084);
INSERT INTO `menu` VALUES (7579, 7577, '新增板块', 'A', NULL, NULL, '/partyaffairs/api/v1/news/classify', 'POST', NULL, 'partyaffairs:news:classify:add', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1724036084, 1724036084);
INSERT INTO `menu` VALUES (7580, 7577, '修改板块', 'A', NULL, NULL, '/partyaffairs/api/v1/news/classify', 'PUT', NULL, 'partyaffairs:news:classify:update', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1724036084, 1724036084);
INSERT INTO `menu` VALUES (7581, 7577, '删除板块', 'A', NULL, NULL, '/partyaffairs/api/v1/news/classify', 'DELETE', NULL, 'partyaffairs:news:classify:delete', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1724036084, 1724036084);
INSERT INTO `menu` VALUES (7582, 7576, '内容管理', 'M', 'PartyAffairsNewsContent', 'book', NULL, NULL, '/partyaffairs/news/content', NULL, '/partyaffairs/news/content/index', NULL, 0, NULL, 1, NULL, 1, 1724036084, 1724036084);
INSERT INTO `menu` VALUES (7583, 7582, '查看内容列表', 'A', NULL, NULL, '/partyaffairs/api/v1/news/contents', 'GET', NULL, 'partyaffairs:news:content:query', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1724036084, 1724036084);
INSERT INTO `menu` VALUES (7584, 7582, '查看指定内容', 'A', NULL, NULL, '/partyaffairs/api/v1/news/content', 'GET', NULL, 'partyaffairs:news:content:info', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1724036084, 1724036084);
INSERT INTO `menu` VALUES (7585, 7582, '新增内容', 'A', NULL, NULL, '/partyaffairs/api/v1/news/content', 'POST', NULL, 'partyaffairs:news:content:add', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1724036084, 1724036084);
INSERT INTO `menu` VALUES (7586, 7582, '修改内容', 'A', NULL, NULL, '/partyaffairs/api/v1/news/content', 'PUT', NULL, 'partyaffairs:news:content:update', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1724036084, 1724036084);
INSERT INTO `menu` VALUES (7587, 7582, '删除内容', 'A', NULL, NULL, '/partyaffairs/api/v1/news/content', 'DELETE', NULL, 'partyaffairs:news:content:delete', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1724036084, 1724036084);
INSERT INTO `menu` VALUES (7588, 7576, '评论管理', 'M', 'PartyAffairsNewsComment', 'book', NULL, NULL, '/partyaffairs/news/comment', NULL, '/partyaffairs/news/comment/index', NULL, 0, 1, NULL, NULL, NULL, 1724036084, 1724036084);
INSERT INTO `menu` VALUES (7589, 7588, '查看评论列表', 'A', NULL, NULL, '/partyaffairs/api/v1/news/comments', 'GET', NULL, 'partyaffairs:news:comment:query', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1724036084, 1724036084);
INSERT INTO `menu` VALUES (7590, 7588, '删除评论', 'A', NULL, NULL, '/partyaffairs/api/v1/news/comment', 'DELETE', NULL, 'partyaffairs:news:comment:delete', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1724036084, 1724036084);
INSERT INTO `menu` VALUES (7591, 7564, '资料管理', 'M', 'PartyAffairsResource', 'public', NULL, NULL, '/partyaffairs/resource', NULL, NULL, NULL, 0, 0, NULL, NULL, NULL, 1724036084, 1724036084);
INSERT INTO `menu` VALUES (7592, 7591, '板块管理', 'M', 'PartyAffairsResourceClassify', 'menu', NULL, NULL, '/partyaffairs/resource/classify', NULL, '/partyaffairs/resource/classify/index', NULL, 0, NULL, 1, NULL, 1, 1724036084, 1724036084);
INSERT INTO `menu` VALUES (7593, 7592, '查看板块', 'A', NULL, NULL, '/partyaffairs/api/v1/resource/classify', 'GET', NULL, 'partyaffairs:resource:classify:query', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1724036084, 1724036084);
INSERT INTO `menu` VALUES (7594, 7592, '新增板块', 'A', NULL, NULL, '/partyaffairs/api/v1/resource/classify', 'POST', NULL, 'partyaffairs:resource:classify:add', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1724036084, 1724036084);
INSERT INTO `menu` VALUES (7595, 7592, '修改板块', 'A', NULL, NULL, '/partyaffairs/api/v1/resource/classify', 'PUT', NULL, 'partyaffairs:resource:classify:update', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1724036084, 1724036084);
INSERT INTO `menu` VALUES (7596, 7592, '删除板块', 'A', NULL, NULL, '/partyaffairs/api/v1/resource/classify', 'DELETE', NULL, 'partyaffairs:resource:classify:delete', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1724036084, 1724036084);
INSERT INTO `menu` VALUES (7597, 7591, '内容管理', 'M', 'PartyAffairsResourceContent', 'public', NULL, NULL, '/partyaffairs/resource/content', NULL, '/partyaffairs/resource/content/index', NULL, 0, NULL, 1, NULL, 1, 1724036084, 1724036084);
INSERT INTO `menu` VALUES (7598, 7597, '查看内容列表', 'A', NULL, NULL, '/partyaffairs/api/v1/resource/contents', 'GET', NULL, 'partyaffairs:resource:content:query', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1724036084, 1724036084);
INSERT INTO `menu` VALUES (7599, 7597, '新增内容', 'A', NULL, NULL, '/partyaffairs/api/v1/resource/content', 'POST', NULL, 'partyaffairs:resource:content:add', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1724036084, 1724036084);
INSERT INTO `menu` VALUES (7600, 7597, '修改内容', 'A', NULL, NULL, '/partyaffairs/api/v1/resource/content', 'PUT', NULL, 'partyaffairs:resource:content:update', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1724036084, 1724036084);
INSERT INTO `menu` VALUES (7601, 7597, '删除内容', 'A', NULL, NULL, '/partyaffairs/api/v1/resource/content', 'DELETE', NULL, 'partyaffairs:resource:content:delete', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1724036084, 1724036084);
INSERT INTO `menu` VALUES (7602, 7564, '任务管理', 'M', 'PartyAffairsTask', 'layers', NULL, NULL, '/partyaffairs/task', NULL, NULL, NULL, 0, 0, NULL, NULL, NULL, 1724036084, 1724036084);
INSERT INTO `menu` VALUES (7603, 7602, '任务配置', 'M', 'PartyAffairsTaskConfigure', 'unordered-list', NULL, NULL, '/partyaffairs/task/configure', NULL, '/partyaffairs/task/configure/index', NULL, 0, NULL, 1, NULL, 1, 1724036084, 1724036084);
INSERT INTO `menu` VALUES (7604, 7603, '查看配置', 'A', NULL, NULL, '/partyaffairs/api/v1/task/configure', 'GET', NULL, 'partyaffairs:task:configure:query', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1724036084, 1724036084);
INSERT INTO `menu` VALUES (7605, 7603, '新增配置', 'A', NULL, NULL, '/partyaffairs/api/v1/task/configure', 'POST', NULL, 'partyaffairs:task:configure:add', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1724036084, 1724036084);
INSERT INTO `menu` VALUES (7606, 7603, '修改配置', 'A', NULL, NULL, '/partyaffairs/api/v1/task/configure', 'PUT', NULL, 'partyaffairs:task:configure:update', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1724036084, 1724036084);
INSERT INTO `menu` VALUES (7607, 7603, '删除配置', 'A', NULL, NULL, '/partyaffairs/api/v1/task/configure', 'DELETE', NULL, 'partyaffairs:task:configure:delete', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1724036084, 1724036084);
INSERT INTO `menu` VALUES (7608, 7602, '任务统计', 'M', 'PartyAffairsTaskContent', 'computer', NULL, NULL, '/partyaffairs/task/value', NULL, '/partyaffairs/task/value/index', NULL, 0, NULL, 1, NULL, 1, 1724036084, 1724036084);
INSERT INTO `menu` VALUES (7609, 7608, '查看任务填写列表', 'A', NULL, NULL, '/partyaffairs/api/v1/task/values', 'GET', NULL, 'partyaffairs:task:value:query', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1724036084, 1724036084);
INSERT INTO `menu` VALUES (7610, 7608, '查看指定任务填写详情', 'A', NULL, NULL, '/partyaffairs/api/v1/task/value', 'GET', NULL, 'partyaffairs:task:value:query', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1724036084, 1724036084);
INSERT INTO `menu` VALUES (7611, 7608, '导出任务填写详情', 'A', NULL, NULL, '/partyaffairs/api/v1/task/values', 'POST', NULL, 'partyaffairs:task:value:export', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1724036084, 1724036084);
INSERT INTO `menu` VALUES (7612, 7608, '修改任务填写详情', 'A', NULL, NULL, '/partyaffairs/api/v1/task/value', 'PUT', NULL, 'partyaffairs:task:value:update', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1724036084, 1724036084);
INSERT INTO `menu` VALUES (7613, 7608, '删除任务填写详情', 'A', NULL, NULL, '/partyaffairs/api/v1/task/value', 'DELETE', NULL, 'partyaffairs:task:value:delete', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1724036084, 1724036084);
INSERT INTO `menu` VALUES (7614, 7564, '视频管理', 'M', 'PartyAffairsVideo', 'video-camera', NULL, NULL, '/partyaffairs/video', NULL, NULL, NULL, 0, 0, NULL, NULL, NULL, 1724036084, 1724036084);
INSERT INTO `menu` VALUES (7615, 7614, '板块管理', 'M', 'PartyAffairsVideoClassify', 'menu', NULL, NULL, '/partyaffairs/video/classify', NULL, '/partyaffairs/video/classify/index', NULL, 0, NULL, 1, NULL, 1, 1724036084, 1724036084);
INSERT INTO `menu` VALUES (7616, 7615, '查看板块', 'A', NULL, NULL, '/partyaffairs/api/v1/video/classify', 'GET', NULL, 'partyaffairs:video:classify:query', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1724036084, 1724036084);
INSERT INTO `menu` VALUES (7617, 7615, '新增板块', 'A', NULL, NULL, '/partyaffairs/api/v1/video/classify', 'POST', NULL, 'partyaffairs:video:classify:add', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1724036084, 1724036084);
INSERT INTO `menu` VALUES (7618, 7615, '修改板块', 'A', NULL, NULL, '/partyaffairs/api/v1/video/classify', 'PUT', NULL, 'partyaffairs:video:classify:update', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1724036084, 1724036084);
INSERT INTO `menu` VALUES (7619, 7615, '删除板块', 'A', NULL, NULL, '/partyaffairs/api/v1/video/classify', 'DELETE', NULL, 'partyaffairs:video:classify:delete', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1724036084, 1724036084);
INSERT INTO `menu` VALUES (7620, 7614, '内容管理', 'M', 'PartyAffairsVideoContent', 'video-camera', NULL, NULL, '/partyaffairs/video/content', NULL, '/partyaffairs/video/content/index', NULL, 0, NULL, 1, NULL, 1, 1724036084, 1724036084);
INSERT INTO `menu` VALUES (7621, 7620, '查看内容列表', 'A', NULL, NULL, '/partyaffairs/api/v1/video/contents', 'GET', NULL, 'partyaffairs:video:content:query', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1724036084, 1724036084);
INSERT INTO `menu` VALUES (7622, 7620, '新增内容', 'A', NULL, NULL, '/partyaffairs/api/v1/video/content', 'POST', NULL, 'partyaffairs:video:content:add', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1724036084, 1724036084);
INSERT INTO `menu` VALUES (7623, 7620, '修改内容', 'A', NULL, NULL, '/partyaffairs/api/v1/video/content', 'PUT', NULL, 'partyaffairs:video:content:update', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1724036084, 1724036084);
INSERT INTO `menu` VALUES (7624, 7620, '删除内容', 'A', NULL, NULL, '/partyaffairs/api/v1/video/content', 'DELETE', NULL, 'partyaffairs:video:content:delete', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1724036084, 1724036084);
INSERT INTO `menu` VALUES (7625, 7620, '进度统计', 'A', NULL, NULL, '/partyaffairs/api/v1/video/process', 'GET', NULL, 'partyaffairs:video:content:process', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1724036084, 1724036084);
INSERT INTO `menu` VALUES (7626, 7562, '贫困资助系统', 'M', 'poverty', 'apps', NULL, NULL, '/poverty', NULL, NULL, NULL, 0, 0, NULL, NULL, NULL, 1724036084, 1724036084);
INSERT INTO `menu` VALUES (7627, 7626, '轮播管理', 'M', 'PovertyBanner', 'list', NULL, NULL, '/poverty/banner', NULL, '/poverty/banner/index', NULL, 0, NULL, 1, NULL, 1, 1724036084, 1724036084);
INSERT INTO `menu` VALUES (7628, 7627, '查看轮播', 'A', NULL, NULL, '/poverty/api/v1/banners', 'GET', NULL, 'poverty:banner:query', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1724036084, 1724036084);
INSERT INTO `menu` VALUES (7629, 7627, '新增轮播', 'A', NULL, NULL, '/poverty/api/v1/banner', 'POST', NULL, 'poverty:banner:add', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1724036084, 1724036084);
INSERT INTO `menu` VALUES (7630, 7627, '修改轮播', 'A', NULL, NULL, '/poverty/api/v1/banner', 'PUT', NULL, 'poverty:banner:update', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1724036084, 1724036084);
INSERT INTO `menu` VALUES (7631, 7627, '删除轮播', 'A', NULL, NULL, '/poverty/api/v1/banner', 'DELETE', NULL, 'poverty:banner:delete', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1724036084, 1724036084);
INSERT INTO `menu` VALUES (7632, 7626, '通知管理', 'M', 'PovertyNotice', 'notification', NULL, NULL, '/poverty/notice', NULL, '/poverty/notice/index', NULL, 0, NULL, 1, NULL, 1, 1724036084, 1724036084);
INSERT INTO `menu` VALUES (7633, 7632, '查看指定通知', 'A', NULL, NULL, '/poverty/api/v1/notice', 'GET', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL, 1724036084, 1724036084);
INSERT INTO `menu` VALUES (7634, 7632, '查看通知列表', 'A', NULL, NULL, '/poverty/api/v1/notices', 'GET', NULL, 'poverty:notice:query', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1724036084, 1724036084);
INSERT INTO `menu` VALUES (7635, 7632, '新增通知', 'A', NULL, NULL, '/poverty/api/v1/notice', 'POST', NULL, 'poverty:notice:add', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1724036084, 1724036084);
INSERT INTO `menu` VALUES (7636, 7632, '修改通知', 'A', NULL, NULL, '/poverty/api/v1/notice', 'PUT', NULL, 'poverty:notice:update', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1724036084, 1724036084);
INSERT INTO `menu` VALUES (7637, 7632, '删除通知', 'A', NULL, NULL, '/poverty/api/v1/notice', 'DELETE', NULL, 'poverty:notice:delete', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1724036084, 1724036084);
INSERT INTO `menu` VALUES (7638, 7626, '资讯管理', 'M', 'PovertyInformation', 'book', NULL, NULL, '/poverty/information', NULL, '/poverty/information/index', NULL, 0, 0, NULL, NULL, NULL, 1724036084, 1724036084);
INSERT INTO `menu` VALUES (7639, 7638, '查看资讯列表', 'A', NULL, NULL, '/poverty/api/v1/informations', 'GET', NULL, 'poverty:information:query', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1724036084, 1724036084);
INSERT INTO `menu` VALUES (7640, 7638, '查看指定资讯', 'A', NULL, NULL, '/poverty/api/v1/information', 'GET', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL, 1724036084, 1724036084);
INSERT INTO `menu` VALUES (7641, 7638, '新增资讯', 'A', NULL, NULL, '/poverty/api/v1/information', 'POST', NULL, 'poverty:information:add', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1724036084, 1724036084);
INSERT INTO `menu` VALUES (7642, 7638, '修改资讯', 'A', NULL, NULL, '/poverty/api/v1/information', 'PUT', NULL, 'poverty:information:update', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1724036084, 1724036084);
INSERT INTO `menu` VALUES (7643, 7638, '删除资讯', 'A', NULL, NULL, '/poverty/api/v1/information', 'DELETE', NULL, 'poverty:information:delete', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1724036084, 1724036084);
INSERT INTO `menu` VALUES (7644, 7626, '活动管理', 'M', 'PovertyActivity', 'unordered-list', NULL, NULL, '/poverty/activity', NULL, '/poverty/activity/index', NULL, 0, 0, NULL, NULL, NULL, 1724036084, 1724036084);
INSERT INTO `menu` VALUES (7645, 7644, '查看活动列表', 'A', NULL, NULL, '/poverty/api/v1/activitys', 'GET', NULL, 'poverty:activity:query', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1724036084, 1724036084);
INSERT INTO `menu` VALUES (7646, 7644, '查看指定活动', 'A', NULL, NULL, '/poverty/api/v1/activity', 'GET', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL, 1724036084, 1724036084);
INSERT INTO `menu` VALUES (7647, 7644, '新增活动', 'A', NULL, NULL, '/poverty/api/v1/activity', 'POST', NULL, 'poverty:activity:add', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1724036084, 1724036084);
INSERT INTO `menu` VALUES (7648, 7644, '修改活动', 'A', NULL, NULL, '/poverty/api/v1/activity', 'PUT', NULL, 'poverty:activity:update', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1724036084, 1724036084);
INSERT INTO `menu` VALUES (7649, 7644, '删除活动', 'A', NULL, NULL, '/poverty/api/v1/activity', 'DELETE', NULL, 'poverty:activity:delete', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1724036084, 1724036084);
INSERT INTO `menu` VALUES (7650, 7626, '咨询记录', 'M', 'PovertyChat', 'robot', NULL, NULL, '/poverty/chat', NULL, '/poverty/chat/index', NULL, 0, 0, NULL, NULL, NULL, 1724036084, 1724036084);
INSERT INTO `menu` VALUES (7651, 7650, '查看咨询记录', 'A', NULL, NULL, '/poverty/api/v1/chat/records', 'GET', NULL, 'poverty:chat:query', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1724036084, 1724036084);
INSERT INTO `menu` VALUES (7652, 7626, '文件管理', 'M', 'PovertyResource', 'public', NULL, NULL, '/poverty/resource', NULL, '/poverty/resource/index', NULL, 0, 0, NULL, NULL, NULL, 1724036084, 1724036084);
INSERT INTO `menu` VALUES (7653, 7652, '查看文件列表', 'A', NULL, NULL, '/poverty/api/v1/resources', 'GET', NULL, 'poverty:resource:query', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1724036084, 1724036084);
INSERT INTO `menu` VALUES (7654, 7652, '新增文件', 'A', NULL, NULL, '/poverty/api/v1/resource', 'POST', NULL, 'poverty:resource:add', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1724036084, 1724036084);
INSERT INTO `menu` VALUES (7655, 7652, '修改文件', 'A', NULL, NULL, '/poverty/api/v1/resource', 'PUT', NULL, 'poverty:resource:update', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1724036084, 1724036084);
INSERT INTO `menu` VALUES (7656, 7652, '删除文件', 'A', NULL, NULL, '/poverty/api/v1/resource', 'DELETE', NULL, 'poverty:resource:delete', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1724036084, 1724036084);
INSERT INTO `menu` VALUES (7657, 7626, '留言记录', 'M', 'PovertyComment', 'message', NULL, NULL, '/poverty/comment', NULL, '/poverty/comment/index', NULL, 0, 0, NULL, NULL, NULL, 1724036084, 1724036084);
INSERT INTO `menu` VALUES (7658, 7657, '查看留言', 'A', NULL, NULL, '/poverty/api/v1/comments', 'GET', NULL, 'poverty:comment:query', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1724036084, 1724036084);
INSERT INTO `menu` VALUES (7659, 7657, '删除留言', 'A', NULL, NULL, '/poverty/api/v1/comment', 'DELETE', NULL, 'poverty:comment:delete', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1724036084, 1724036084);
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
) ENGINE=InnoDB AUTO_INCREMENT=21295 DEFAULT CHARSET=utf8mb4  COMMENT='菜单层级信息';

-- ----------------------------
-- Records of menu_closure
-- ----------------------------
BEGIN;
INSERT INTO `menu_closure` VALUES (20410, 7368, 7372);
INSERT INTO `menu_closure` VALUES (20411, 7370, 7372);
INSERT INTO `menu_closure` VALUES (20412, 7371, 7372);
INSERT INTO `menu_closure` VALUES (20413, 7368, 7373);
INSERT INTO `menu_closure` VALUES (20414, 7370, 7373);
INSERT INTO `menu_closure` VALUES (20415, 7371, 7373);
INSERT INTO `menu_closure` VALUES (20416, 7368, 7374);
INSERT INTO `menu_closure` VALUES (20417, 7370, 7374);
INSERT INTO `menu_closure` VALUES (20418, 7371, 7374);
INSERT INTO `menu_closure` VALUES (20419, 7368, 7375);
INSERT INTO `menu_closure` VALUES (20420, 7370, 7375);
INSERT INTO `menu_closure` VALUES (20421, 7371, 7375);
INSERT INTO `menu_closure` VALUES (20422, 7368, 7376);
INSERT INTO `menu_closure` VALUES (20423, 7370, 7376);
INSERT INTO `menu_closure` VALUES (20424, 7371, 7376);
INSERT INTO `menu_closure` VALUES (20425, 7368, 7377);
INSERT INTO `menu_closure` VALUES (20426, 7370, 7377);
INSERT INTO `menu_closure` VALUES (20427, 7371, 7377);
INSERT INTO `menu_closure` VALUES (20428, 7368, 7378);
INSERT INTO `menu_closure` VALUES (20429, 7370, 7378);
INSERT INTO `menu_closure` VALUES (20430, 7371, 7378);
INSERT INTO `menu_closure` VALUES (20431, 7368, 7379);
INSERT INTO `menu_closure` VALUES (20432, 7370, 7379);
INSERT INTO `menu_closure` VALUES (20433, 7371, 7379);
INSERT INTO `menu_closure` VALUES (20434, 7368, 7380);
INSERT INTO `menu_closure` VALUES (20435, 7370, 7380);
INSERT INTO `menu_closure` VALUES (20436, 7371, 7380);
INSERT INTO `menu_closure` VALUES (20437, 7368, 7381);
INSERT INTO `menu_closure` VALUES (20438, 7370, 7381);
INSERT INTO `menu_closure` VALUES (20439, 7371, 7381);
INSERT INTO `menu_closure` VALUES (20440, 7368, 7382);
INSERT INTO `menu_closure` VALUES (20441, 7370, 7382);
INSERT INTO `menu_closure` VALUES (20442, 7371, 7382);
INSERT INTO `menu_closure` VALUES (20443, 7368, 7383);
INSERT INTO `menu_closure` VALUES (20444, 7370, 7383);
INSERT INTO `menu_closure` VALUES (20445, 7371, 7383);
INSERT INTO `menu_closure` VALUES (20446, 7368, 7384);
INSERT INTO `menu_closure` VALUES (20447, 7370, 7384);
INSERT INTO `menu_closure` VALUES (20448, 7371, 7384);
INSERT INTO `menu_closure` VALUES (20449, 7368, 7385);
INSERT INTO `menu_closure` VALUES (20450, 7370, 7385);
INSERT INTO `menu_closure` VALUES (20451, 7371, 7385);
INSERT INTO `menu_closure` VALUES (20452, 7368, 7386);
INSERT INTO `menu_closure` VALUES (20453, 7370, 7386);
INSERT INTO `menu_closure` VALUES (20454, 7371, 7386);
INSERT INTO `menu_closure` VALUES (20455, 7368, 7387);
INSERT INTO `menu_closure` VALUES (20456, 7370, 7387);
INSERT INTO `menu_closure` VALUES (20457, 7371, 7387);
INSERT INTO `menu_closure` VALUES (20458, 7368, 7388);
INSERT INTO `menu_closure` VALUES (20459, 7370, 7388);
INSERT INTO `menu_closure` VALUES (20460, 7371, 7388);
INSERT INTO `menu_closure` VALUES (20461, 7368, 7389);
INSERT INTO `menu_closure` VALUES (20462, 7370, 7389);
INSERT INTO `menu_closure` VALUES (20463, 7371, 7389);
INSERT INTO `menu_closure` VALUES (20464, 7368, 7390);
INSERT INTO `menu_closure` VALUES (20465, 7370, 7390);
INSERT INTO `menu_closure` VALUES (20466, 7371, 7390);
INSERT INTO `menu_closure` VALUES (20467, 7368, 7391);
INSERT INTO `menu_closure` VALUES (20468, 7370, 7391);
INSERT INTO `menu_closure` VALUES (20469, 7371, 7391);
INSERT INTO `menu_closure` VALUES (20470, 7368, 7393);
INSERT INTO `menu_closure` VALUES (20471, 7370, 7393);
INSERT INTO `menu_closure` VALUES (20472, 7392, 7393);
INSERT INTO `menu_closure` VALUES (20473, 7368, 7394);
INSERT INTO `menu_closure` VALUES (20474, 7370, 7394);
INSERT INTO `menu_closure` VALUES (20475, 7392, 7394);
INSERT INTO `menu_closure` VALUES (20476, 7368, 7395);
INSERT INTO `menu_closure` VALUES (20477, 7370, 7395);
INSERT INTO `menu_closure` VALUES (20478, 7392, 7395);
INSERT INTO `menu_closure` VALUES (20479, 7368, 7396);
INSERT INTO `menu_closure` VALUES (20480, 7370, 7396);
INSERT INTO `menu_closure` VALUES (20481, 7392, 7396);
INSERT INTO `menu_closure` VALUES (20482, 7368, 7397);
INSERT INTO `menu_closure` VALUES (20483, 7370, 7397);
INSERT INTO `menu_closure` VALUES (20484, 7392, 7397);
INSERT INTO `menu_closure` VALUES (20485, 7368, 7398);
INSERT INTO `menu_closure` VALUES (20486, 7370, 7398);
INSERT INTO `menu_closure` VALUES (20487, 7392, 7398);
INSERT INTO `menu_closure` VALUES (20488, 7368, 7399);
INSERT INTO `menu_closure` VALUES (20489, 7370, 7399);
INSERT INTO `menu_closure` VALUES (20490, 7392, 7399);
INSERT INTO `menu_closure` VALUES (20491, 7368, 7400);
INSERT INTO `menu_closure` VALUES (20492, 7370, 7400);
INSERT INTO `menu_closure` VALUES (20493, 7392, 7400);
INSERT INTO `menu_closure` VALUES (20494, 7368, 7401);
INSERT INTO `menu_closure` VALUES (20495, 7370, 7401);
INSERT INTO `menu_closure` VALUES (20496, 7392, 7401);
INSERT INTO `menu_closure` VALUES (20497, 7368, 7403);
INSERT INTO `menu_closure` VALUES (20498, 7370, 7403);
INSERT INTO `menu_closure` VALUES (20499, 7402, 7403);
INSERT INTO `menu_closure` VALUES (20500, 7368, 7404);
INSERT INTO `menu_closure` VALUES (20501, 7370, 7404);
INSERT INTO `menu_closure` VALUES (20502, 7402, 7404);
INSERT INTO `menu_closure` VALUES (20503, 7368, 7405);
INSERT INTO `menu_closure` VALUES (20504, 7370, 7405);
INSERT INTO `menu_closure` VALUES (20505, 7402, 7405);
INSERT INTO `menu_closure` VALUES (20506, 7368, 7406);
INSERT INTO `menu_closure` VALUES (20507, 7370, 7406);
INSERT INTO `menu_closure` VALUES (20508, 7402, 7406);
INSERT INTO `menu_closure` VALUES (20509, 7368, 7408);
INSERT INTO `menu_closure` VALUES (20510, 7370, 7408);
INSERT INTO `menu_closure` VALUES (20511, 7407, 7408);
INSERT INTO `menu_closure` VALUES (20512, 7368, 7409);
INSERT INTO `menu_closure` VALUES (20513, 7370, 7409);
INSERT INTO `menu_closure` VALUES (20514, 7407, 7409);
INSERT INTO `menu_closure` VALUES (20515, 7368, 7410);
INSERT INTO `menu_closure` VALUES (20516, 7370, 7410);
INSERT INTO `menu_closure` VALUES (20517, 7407, 7410);
INSERT INTO `menu_closure` VALUES (20518, 7368, 7411);
INSERT INTO `menu_closure` VALUES (20519, 7370, 7411);
INSERT INTO `menu_closure` VALUES (20520, 7407, 7411);
INSERT INTO `menu_closure` VALUES (20521, 7368, 7413);
INSERT INTO `menu_closure` VALUES (20522, 7370, 7413);
INSERT INTO `menu_closure` VALUES (20523, 7412, 7413);
INSERT INTO `menu_closure` VALUES (20524, 7368, 7414);
INSERT INTO `menu_closure` VALUES (20525, 7370, 7414);
INSERT INTO `menu_closure` VALUES (20526, 7412, 7414);
INSERT INTO `menu_closure` VALUES (20527, 7368, 7415);
INSERT INTO `menu_closure` VALUES (20528, 7370, 7415);
INSERT INTO `menu_closure` VALUES (20529, 7412, 7415);
INSERT INTO `menu_closure` VALUES (20530, 7368, 7422);
INSERT INTO `menu_closure` VALUES (20531, 7370, 7422);
INSERT INTO `menu_closure` VALUES (20532, 7416, 7422);
INSERT INTO `menu_closure` VALUES (20533, 7421, 7422);
INSERT INTO `menu_closure` VALUES (20534, 7368, 7423);
INSERT INTO `menu_closure` VALUES (20535, 7370, 7423);
INSERT INTO `menu_closure` VALUES (20536, 7416, 7423);
INSERT INTO `menu_closure` VALUES (20537, 7421, 7423);
INSERT INTO `menu_closure` VALUES (20538, 7368, 7417);
INSERT INTO `menu_closure` VALUES (20539, 7370, 7417);
INSERT INTO `menu_closure` VALUES (20540, 7416, 7417);
INSERT INTO `menu_closure` VALUES (20541, 7368, 7418);
INSERT INTO `menu_closure` VALUES (20542, 7370, 7418);
INSERT INTO `menu_closure` VALUES (20543, 7416, 7418);
INSERT INTO `menu_closure` VALUES (20544, 7368, 7419);
INSERT INTO `menu_closure` VALUES (20545, 7370, 7419);
INSERT INTO `menu_closure` VALUES (20546, 7416, 7419);
INSERT INTO `menu_closure` VALUES (20547, 7368, 7420);
INSERT INTO `menu_closure` VALUES (20548, 7370, 7420);
INSERT INTO `menu_closure` VALUES (20549, 7416, 7420);
INSERT INTO `menu_closure` VALUES (20550, 7368, 7421);
INSERT INTO `menu_closure` VALUES (20551, 7370, 7421);
INSERT INTO `menu_closure` VALUES (20552, 7416, 7421);
INSERT INTO `menu_closure` VALUES (20553, 7368, 7425);
INSERT INTO `menu_closure` VALUES (20554, 7370, 7425);
INSERT INTO `menu_closure` VALUES (20555, 7424, 7425);
INSERT INTO `menu_closure` VALUES (20556, 7368, 7426);
INSERT INTO `menu_closure` VALUES (20557, 7370, 7426);
INSERT INTO `menu_closure` VALUES (20558, 7424, 7426);
INSERT INTO `menu_closure` VALUES (20559, 7368, 7427);
INSERT INTO `menu_closure` VALUES (20560, 7370, 7427);
INSERT INTO `menu_closure` VALUES (20561, 7424, 7427);
INSERT INTO `menu_closure` VALUES (20562, 7368, 7428);
INSERT INTO `menu_closure` VALUES (20563, 7370, 7428);
INSERT INTO `menu_closure` VALUES (20564, 7424, 7428);
INSERT INTO `menu_closure` VALUES (20565, 7368, 7429);
INSERT INTO `menu_closure` VALUES (20566, 7370, 7429);
INSERT INTO `menu_closure` VALUES (20567, 7424, 7429);
INSERT INTO `menu_closure` VALUES (20568, 7368, 7430);
INSERT INTO `menu_closure` VALUES (20569, 7370, 7430);
INSERT INTO `menu_closure` VALUES (20570, 7424, 7430);
INSERT INTO `menu_closure` VALUES (20571, 7368, 7371);
INSERT INTO `menu_closure` VALUES (20572, 7370, 7371);
INSERT INTO `menu_closure` VALUES (20573, 7368, 7392);
INSERT INTO `menu_closure` VALUES (20574, 7370, 7392);
INSERT INTO `menu_closure` VALUES (20575, 7368, 7402);
INSERT INTO `menu_closure` VALUES (20576, 7370, 7402);
INSERT INTO `menu_closure` VALUES (20577, 7368, 7407);
INSERT INTO `menu_closure` VALUES (20578, 7370, 7407);
INSERT INTO `menu_closure` VALUES (20579, 7368, 7412);
INSERT INTO `menu_closure` VALUES (20580, 7370, 7412);
INSERT INTO `menu_closure` VALUES (20581, 7368, 7416);
INSERT INTO `menu_closure` VALUES (20582, 7370, 7416);
INSERT INTO `menu_closure` VALUES (20583, 7368, 7424);
INSERT INTO `menu_closure` VALUES (20584, 7370, 7424);
INSERT INTO `menu_closure` VALUES (20585, 7368, 7434);
INSERT INTO `menu_closure` VALUES (20586, 7431, 7434);
INSERT INTO `menu_closure` VALUES (20587, 7432, 7434);
INSERT INTO `menu_closure` VALUES (20588, 7433, 7434);
INSERT INTO `menu_closure` VALUES (20589, 7368, 7435);
INSERT INTO `menu_closure` VALUES (20590, 7431, 7435);
INSERT INTO `menu_closure` VALUES (20591, 7432, 7435);
INSERT INTO `menu_closure` VALUES (20592, 7433, 7435);
INSERT INTO `menu_closure` VALUES (20593, 7368, 7436);
INSERT INTO `menu_closure` VALUES (20594, 7431, 7436);
INSERT INTO `menu_closure` VALUES (20595, 7432, 7436);
INSERT INTO `menu_closure` VALUES (20596, 7433, 7436);
INSERT INTO `menu_closure` VALUES (20597, 7368, 7437);
INSERT INTO `menu_closure` VALUES (20598, 7431, 7437);
INSERT INTO `menu_closure` VALUES (20599, 7432, 7437);
INSERT INTO `menu_closure` VALUES (20600, 7433, 7437);
INSERT INTO `menu_closure` VALUES (20601, 7368, 7439);
INSERT INTO `menu_closure` VALUES (20602, 7431, 7439);
INSERT INTO `menu_closure` VALUES (20603, 7432, 7439);
INSERT INTO `menu_closure` VALUES (20604, 7438, 7439);
INSERT INTO `menu_closure` VALUES (20605, 7368, 7440);
INSERT INTO `menu_closure` VALUES (20606, 7431, 7440);
INSERT INTO `menu_closure` VALUES (20607, 7432, 7440);
INSERT INTO `menu_closure` VALUES (20608, 7438, 7440);
INSERT INTO `menu_closure` VALUES (20609, 7368, 7441);
INSERT INTO `menu_closure` VALUES (20610, 7431, 7441);
INSERT INTO `menu_closure` VALUES (20611, 7432, 7441);
INSERT INTO `menu_closure` VALUES (20612, 7438, 7441);
INSERT INTO `menu_closure` VALUES (20613, 7368, 7433);
INSERT INTO `menu_closure` VALUES (20614, 7431, 7433);
INSERT INTO `menu_closure` VALUES (20615, 7432, 7433);
INSERT INTO `menu_closure` VALUES (20616, 7368, 7438);
INSERT INTO `menu_closure` VALUES (20617, 7431, 7438);
INSERT INTO `menu_closure` VALUES (20618, 7432, 7438);
INSERT INTO `menu_closure` VALUES (20619, 7368, 7443);
INSERT INTO `menu_closure` VALUES (20620, 7431, 7443);
INSERT INTO `menu_closure` VALUES (20621, 7442, 7443);
INSERT INTO `menu_closure` VALUES (20622, 7368, 7444);
INSERT INTO `menu_closure` VALUES (20623, 7431, 7444);
INSERT INTO `menu_closure` VALUES (20624, 7442, 7444);
INSERT INTO `menu_closure` VALUES (20625, 7368, 7445);
INSERT INTO `menu_closure` VALUES (20626, 7431, 7445);
INSERT INTO `menu_closure` VALUES (20627, 7442, 7445);
INSERT INTO `menu_closure` VALUES (20628, 7368, 7432);
INSERT INTO `menu_closure` VALUES (20629, 7431, 7432);
INSERT INTO `menu_closure` VALUES (20630, 7368, 7442);
INSERT INTO `menu_closure` VALUES (20631, 7431, 7442);
INSERT INTO `menu_closure` VALUES (20632, 7368, 7448);
INSERT INTO `menu_closure` VALUES (20633, 7446, 7448);
INSERT INTO `menu_closure` VALUES (20634, 7447, 7448);
INSERT INTO `menu_closure` VALUES (20635, 7368, 7449);
INSERT INTO `menu_closure` VALUES (20636, 7446, 7449);
INSERT INTO `menu_closure` VALUES (20637, 7447, 7449);
INSERT INTO `menu_closure` VALUES (20638, 7368, 7450);
INSERT INTO `menu_closure` VALUES (20639, 7446, 7450);
INSERT INTO `menu_closure` VALUES (20640, 7447, 7450);
INSERT INTO `menu_closure` VALUES (20641, 7368, 7451);
INSERT INTO `menu_closure` VALUES (20642, 7446, 7451);
INSERT INTO `menu_closure` VALUES (20643, 7447, 7451);
INSERT INTO `menu_closure` VALUES (20644, 7368, 7452);
INSERT INTO `menu_closure` VALUES (20645, 7446, 7452);
INSERT INTO `menu_closure` VALUES (20646, 7447, 7452);
INSERT INTO `menu_closure` VALUES (20647, 7368, 7454);
INSERT INTO `menu_closure` VALUES (20648, 7446, 7454);
INSERT INTO `menu_closure` VALUES (20649, 7453, 7454);
INSERT INTO `menu_closure` VALUES (20650, 7368, 7455);
INSERT INTO `menu_closure` VALUES (20651, 7446, 7455);
INSERT INTO `menu_closure` VALUES (20652, 7453, 7455);
INSERT INTO `menu_closure` VALUES (20653, 7368, 7456);
INSERT INTO `menu_closure` VALUES (20654, 7446, 7456);
INSERT INTO `menu_closure` VALUES (20655, 7453, 7456);
INSERT INTO `menu_closure` VALUES (20656, 7368, 7457);
INSERT INTO `menu_closure` VALUES (20657, 7446, 7457);
INSERT INTO `menu_closure` VALUES (20658, 7453, 7457);
INSERT INTO `menu_closure` VALUES (20659, 7368, 7458);
INSERT INTO `menu_closure` VALUES (20660, 7446, 7458);
INSERT INTO `menu_closure` VALUES (20661, 7453, 7458);
INSERT INTO `menu_closure` VALUES (20662, 7368, 7460);
INSERT INTO `menu_closure` VALUES (20663, 7446, 7460);
INSERT INTO `menu_closure` VALUES (20664, 7459, 7460);
INSERT INTO `menu_closure` VALUES (20665, 7368, 7461);
INSERT INTO `menu_closure` VALUES (20666, 7446, 7461);
INSERT INTO `menu_closure` VALUES (20667, 7459, 7461);
INSERT INTO `menu_closure` VALUES (20668, 7368, 7462);
INSERT INTO `menu_closure` VALUES (20669, 7446, 7462);
INSERT INTO `menu_closure` VALUES (20670, 7459, 7462);
INSERT INTO `menu_closure` VALUES (20671, 7368, 7463);
INSERT INTO `menu_closure` VALUES (20672, 7446, 7463);
INSERT INTO `menu_closure` VALUES (20673, 7459, 7463);
INSERT INTO `menu_closure` VALUES (20674, 7368, 7464);
INSERT INTO `menu_closure` VALUES (20675, 7446, 7464);
INSERT INTO `menu_closure` VALUES (20676, 7459, 7464);
INSERT INTO `menu_closure` VALUES (20677, 7368, 7465);
INSERT INTO `menu_closure` VALUES (20678, 7446, 7465);
INSERT INTO `menu_closure` VALUES (20679, 7459, 7465);
INSERT INTO `menu_closure` VALUES (20680, 7368, 7467);
INSERT INTO `menu_closure` VALUES (20681, 7446, 7467);
INSERT INTO `menu_closure` VALUES (20682, 7466, 7467);
INSERT INTO `menu_closure` VALUES (20683, 7368, 7468);
INSERT INTO `menu_closure` VALUES (20684, 7446, 7468);
INSERT INTO `menu_closure` VALUES (20685, 7466, 7468);
INSERT INTO `menu_closure` VALUES (20686, 7368, 7469);
INSERT INTO `menu_closure` VALUES (20687, 7446, 7469);
INSERT INTO `menu_closure` VALUES (20688, 7466, 7469);
INSERT INTO `menu_closure` VALUES (20689, 7368, 7470);
INSERT INTO `menu_closure` VALUES (20690, 7446, 7470);
INSERT INTO `menu_closure` VALUES (20691, 7466, 7470);
INSERT INTO `menu_closure` VALUES (20692, 7368, 7471);
INSERT INTO `menu_closure` VALUES (20693, 7446, 7471);
INSERT INTO `menu_closure` VALUES (20694, 7466, 7471);
INSERT INTO `menu_closure` VALUES (20695, 7368, 7472);
INSERT INTO `menu_closure` VALUES (20696, 7446, 7472);
INSERT INTO `menu_closure` VALUES (20697, 7466, 7472);
INSERT INTO `menu_closure` VALUES (20698, 7368, 7473);
INSERT INTO `menu_closure` VALUES (20699, 7446, 7473);
INSERT INTO `menu_closure` VALUES (20700, 7466, 7473);
INSERT INTO `menu_closure` VALUES (20701, 7368, 7474);
INSERT INTO `menu_closure` VALUES (20702, 7446, 7474);
INSERT INTO `menu_closure` VALUES (20703, 7466, 7474);
INSERT INTO `menu_closure` VALUES (20704, 7368, 7483);
INSERT INTO `menu_closure` VALUES (20705, 7446, 7483);
INSERT INTO `menu_closure` VALUES (20706, 7475, 7483);
INSERT INTO `menu_closure` VALUES (20707, 7482, 7483);
INSERT INTO `menu_closure` VALUES (20708, 7368, 7484);
INSERT INTO `menu_closure` VALUES (20709, 7446, 7484);
INSERT INTO `menu_closure` VALUES (20710, 7475, 7484);
INSERT INTO `menu_closure` VALUES (20711, 7482, 7484);
INSERT INTO `menu_closure` VALUES (20712, 7368, 7485);
INSERT INTO `menu_closure` VALUES (20713, 7446, 7485);
INSERT INTO `menu_closure` VALUES (20714, 7475, 7485);
INSERT INTO `menu_closure` VALUES (20715, 7482, 7485);
INSERT INTO `menu_closure` VALUES (20716, 7368, 7486);
INSERT INTO `menu_closure` VALUES (20717, 7446, 7486);
INSERT INTO `menu_closure` VALUES (20718, 7475, 7486);
INSERT INTO `menu_closure` VALUES (20719, 7482, 7486);
INSERT INTO `menu_closure` VALUES (20720, 7368, 7487);
INSERT INTO `menu_closure` VALUES (20721, 7446, 7487);
INSERT INTO `menu_closure` VALUES (20722, 7475, 7487);
INSERT INTO `menu_closure` VALUES (20723, 7482, 7487);
INSERT INTO `menu_closure` VALUES (20724, 7368, 7488);
INSERT INTO `menu_closure` VALUES (20725, 7446, 7488);
INSERT INTO `menu_closure` VALUES (20726, 7475, 7488);
INSERT INTO `menu_closure` VALUES (20727, 7482, 7488);
INSERT INTO `menu_closure` VALUES (20728, 7368, 7489);
INSERT INTO `menu_closure` VALUES (20729, 7446, 7489);
INSERT INTO `menu_closure` VALUES (20730, 7475, 7489);
INSERT INTO `menu_closure` VALUES (20731, 7482, 7489);
INSERT INTO `menu_closure` VALUES (20732, 7368, 7490);
INSERT INTO `menu_closure` VALUES (20733, 7446, 7490);
INSERT INTO `menu_closure` VALUES (20734, 7475, 7490);
INSERT INTO `menu_closure` VALUES (20735, 7482, 7490);
INSERT INTO `menu_closure` VALUES (20736, 7368, 7491);
INSERT INTO `menu_closure` VALUES (20737, 7446, 7491);
INSERT INTO `menu_closure` VALUES (20738, 7475, 7491);
INSERT INTO `menu_closure` VALUES (20739, 7482, 7491);
INSERT INTO `menu_closure` VALUES (20740, 7368, 7492);
INSERT INTO `menu_closure` VALUES (20741, 7446, 7492);
INSERT INTO `menu_closure` VALUES (20742, 7475, 7492);
INSERT INTO `menu_closure` VALUES (20743, 7482, 7492);
INSERT INTO `menu_closure` VALUES (20744, 7368, 7476);
INSERT INTO `menu_closure` VALUES (20745, 7446, 7476);
INSERT INTO `menu_closure` VALUES (20746, 7475, 7476);
INSERT INTO `menu_closure` VALUES (20747, 7368, 7477);
INSERT INTO `menu_closure` VALUES (20748, 7446, 7477);
INSERT INTO `menu_closure` VALUES (20749, 7475, 7477);
INSERT INTO `menu_closure` VALUES (20750, 7368, 7478);
INSERT INTO `menu_closure` VALUES (20751, 7446, 7478);
INSERT INTO `menu_closure` VALUES (20752, 7475, 7478);
INSERT INTO `menu_closure` VALUES (20753, 7368, 7479);
INSERT INTO `menu_closure` VALUES (20754, 7446, 7479);
INSERT INTO `menu_closure` VALUES (20755, 7475, 7479);
INSERT INTO `menu_closure` VALUES (20756, 7368, 7480);
INSERT INTO `menu_closure` VALUES (20757, 7446, 7480);
INSERT INTO `menu_closure` VALUES (20758, 7475, 7480);
INSERT INTO `menu_closure` VALUES (20759, 7368, 7481);
INSERT INTO `menu_closure` VALUES (20760, 7446, 7481);
INSERT INTO `menu_closure` VALUES (20761, 7475, 7481);
INSERT INTO `menu_closure` VALUES (20762, 7368, 7482);
INSERT INTO `menu_closure` VALUES (20763, 7446, 7482);
INSERT INTO `menu_closure` VALUES (20764, 7475, 7482);
INSERT INTO `menu_closure` VALUES (20765, 7368, 7447);
INSERT INTO `menu_closure` VALUES (20766, 7446, 7447);
INSERT INTO `menu_closure` VALUES (20767, 7368, 7453);
INSERT INTO `menu_closure` VALUES (20768, 7446, 7453);
INSERT INTO `menu_closure` VALUES (20769, 7368, 7459);
INSERT INTO `menu_closure` VALUES (20770, 7446, 7459);
INSERT INTO `menu_closure` VALUES (20771, 7368, 7466);
INSERT INTO `menu_closure` VALUES (20772, 7446, 7466);
INSERT INTO `menu_closure` VALUES (20773, 7368, 7475);
INSERT INTO `menu_closure` VALUES (20774, 7446, 7475);
INSERT INTO `menu_closure` VALUES (20775, 7368, 7495);
INSERT INTO `menu_closure` VALUES (20776, 7493, 7495);
INSERT INTO `menu_closure` VALUES (20777, 7494, 7495);
INSERT INTO `menu_closure` VALUES (20778, 7368, 7496);
INSERT INTO `menu_closure` VALUES (20779, 7493, 7496);
INSERT INTO `menu_closure` VALUES (20780, 7494, 7496);
INSERT INTO `menu_closure` VALUES (20781, 7368, 7497);
INSERT INTO `menu_closure` VALUES (20782, 7493, 7497);
INSERT INTO `menu_closure` VALUES (20783, 7494, 7497);
INSERT INTO `menu_closure` VALUES (20784, 7368, 7498);
INSERT INTO `menu_closure` VALUES (20785, 7493, 7498);
INSERT INTO `menu_closure` VALUES (20786, 7494, 7498);
INSERT INTO `menu_closure` VALUES (20787, 7368, 7499);
INSERT INTO `menu_closure` VALUES (20788, 7493, 7499);
INSERT INTO `menu_closure` VALUES (20789, 7494, 7499);
INSERT INTO `menu_closure` VALUES (20790, 7368, 7500);
INSERT INTO `menu_closure` VALUES (20791, 7493, 7500);
INSERT INTO `menu_closure` VALUES (20792, 7494, 7500);
INSERT INTO `menu_closure` VALUES (20793, 7368, 7501);
INSERT INTO `menu_closure` VALUES (20794, 7493, 7501);
INSERT INTO `menu_closure` VALUES (20795, 7494, 7501);
INSERT INTO `menu_closure` VALUES (20796, 7368, 7502);
INSERT INTO `menu_closure` VALUES (20797, 7493, 7502);
INSERT INTO `menu_closure` VALUES (20798, 7494, 7502);
INSERT INTO `menu_closure` VALUES (20799, 7368, 7504);
INSERT INTO `menu_closure` VALUES (20800, 7493, 7504);
INSERT INTO `menu_closure` VALUES (20801, 7503, 7504);
INSERT INTO `menu_closure` VALUES (20802, 7368, 7505);
INSERT INTO `menu_closure` VALUES (20803, 7493, 7505);
INSERT INTO `menu_closure` VALUES (20804, 7503, 7505);
INSERT INTO `menu_closure` VALUES (20805, 7368, 7506);
INSERT INTO `menu_closure` VALUES (20806, 7493, 7506);
INSERT INTO `menu_closure` VALUES (20807, 7503, 7506);
INSERT INTO `menu_closure` VALUES (20808, 7368, 7507);
INSERT INTO `menu_closure` VALUES (20809, 7493, 7507);
INSERT INTO `menu_closure` VALUES (20810, 7503, 7507);
INSERT INTO `menu_closure` VALUES (20811, 7368, 7508);
INSERT INTO `menu_closure` VALUES (20812, 7493, 7508);
INSERT INTO `menu_closure` VALUES (20813, 7503, 7508);
INSERT INTO `menu_closure` VALUES (20814, 7368, 7509);
INSERT INTO `menu_closure` VALUES (20815, 7493, 7509);
INSERT INTO `menu_closure` VALUES (20816, 7503, 7509);
INSERT INTO `menu_closure` VALUES (20817, 7368, 7511);
INSERT INTO `menu_closure` VALUES (20818, 7493, 7511);
INSERT INTO `menu_closure` VALUES (20819, 7510, 7511);
INSERT INTO `menu_closure` VALUES (20820, 7368, 7512);
INSERT INTO `menu_closure` VALUES (20821, 7493, 7512);
INSERT INTO `menu_closure` VALUES (20822, 7510, 7512);
INSERT INTO `menu_closure` VALUES (20823, 7368, 7513);
INSERT INTO `menu_closure` VALUES (20824, 7493, 7513);
INSERT INTO `menu_closure` VALUES (20825, 7510, 7513);
INSERT INTO `menu_closure` VALUES (20826, 7368, 7514);
INSERT INTO `menu_closure` VALUES (20827, 7493, 7514);
INSERT INTO `menu_closure` VALUES (20828, 7510, 7514);
INSERT INTO `menu_closure` VALUES (20829, 7368, 7515);
INSERT INTO `menu_closure` VALUES (20830, 7493, 7515);
INSERT INTO `menu_closure` VALUES (20831, 7510, 7515);
INSERT INTO `menu_closure` VALUES (20832, 7368, 7516);
INSERT INTO `menu_closure` VALUES (20833, 7493, 7516);
INSERT INTO `menu_closure` VALUES (20834, 7510, 7516);
INSERT INTO `menu_closure` VALUES (20835, 7368, 7518);
INSERT INTO `menu_closure` VALUES (20836, 7493, 7518);
INSERT INTO `menu_closure` VALUES (20837, 7517, 7518);
INSERT INTO `menu_closure` VALUES (20838, 7368, 7519);
INSERT INTO `menu_closure` VALUES (20839, 7493, 7519);
INSERT INTO `menu_closure` VALUES (20840, 7517, 7519);
INSERT INTO `menu_closure` VALUES (20841, 7368, 7520);
INSERT INTO `menu_closure` VALUES (20842, 7493, 7520);
INSERT INTO `menu_closure` VALUES (20843, 7517, 7520);
INSERT INTO `menu_closure` VALUES (20844, 7368, 7521);
INSERT INTO `menu_closure` VALUES (20845, 7493, 7521);
INSERT INTO `menu_closure` VALUES (20846, 7517, 7521);
INSERT INTO `menu_closure` VALUES (20847, 7368, 7522);
INSERT INTO `menu_closure` VALUES (20848, 7493, 7522);
INSERT INTO `menu_closure` VALUES (20849, 7517, 7522);
INSERT INTO `menu_closure` VALUES (20850, 7368, 7523);
INSERT INTO `menu_closure` VALUES (20851, 7493, 7523);
INSERT INTO `menu_closure` VALUES (20852, 7517, 7523);
INSERT INTO `menu_closure` VALUES (20853, 7368, 7526);
INSERT INTO `menu_closure` VALUES (20854, 7493, 7526);
INSERT INTO `menu_closure` VALUES (20855, 7524, 7526);
INSERT INTO `menu_closure` VALUES (20856, 7525, 7526);
INSERT INTO `menu_closure` VALUES (20857, 7368, 7527);
INSERT INTO `menu_closure` VALUES (20858, 7493, 7527);
INSERT INTO `menu_closure` VALUES (20859, 7524, 7527);
INSERT INTO `menu_closure` VALUES (20860, 7525, 7527);
INSERT INTO `menu_closure` VALUES (20861, 7368, 7528);
INSERT INTO `menu_closure` VALUES (20862, 7493, 7528);
INSERT INTO `menu_closure` VALUES (20863, 7524, 7528);
INSERT INTO `menu_closure` VALUES (20864, 7525, 7528);
INSERT INTO `menu_closure` VALUES (20865, 7368, 7529);
INSERT INTO `menu_closure` VALUES (20866, 7493, 7529);
INSERT INTO `menu_closure` VALUES (20867, 7524, 7529);
INSERT INTO `menu_closure` VALUES (20868, 7525, 7529);
INSERT INTO `menu_closure` VALUES (20869, 7368, 7530);
INSERT INTO `menu_closure` VALUES (20870, 7493, 7530);
INSERT INTO `menu_closure` VALUES (20871, 7524, 7530);
INSERT INTO `menu_closure` VALUES (20872, 7525, 7530);
INSERT INTO `menu_closure` VALUES (20873, 7368, 7531);
INSERT INTO `menu_closure` VALUES (20874, 7493, 7531);
INSERT INTO `menu_closure` VALUES (20875, 7524, 7531);
INSERT INTO `menu_closure` VALUES (20876, 7525, 7531);
INSERT INTO `menu_closure` VALUES (20877, 7368, 7532);
INSERT INTO `menu_closure` VALUES (20878, 7493, 7532);
INSERT INTO `menu_closure` VALUES (20879, 7524, 7532);
INSERT INTO `menu_closure` VALUES (20880, 7525, 7532);
INSERT INTO `menu_closure` VALUES (20881, 7368, 7534);
INSERT INTO `menu_closure` VALUES (20882, 7493, 7534);
INSERT INTO `menu_closure` VALUES (20883, 7524, 7534);
INSERT INTO `menu_closure` VALUES (20884, 7533, 7534);
INSERT INTO `menu_closure` VALUES (20885, 7368, 7535);
INSERT INTO `menu_closure` VALUES (20886, 7493, 7535);
INSERT INTO `menu_closure` VALUES (20887, 7524, 7535);
INSERT INTO `menu_closure` VALUES (20888, 7533, 7535);
INSERT INTO `menu_closure` VALUES (20889, 7368, 7525);
INSERT INTO `menu_closure` VALUES (20890, 7493, 7525);
INSERT INTO `menu_closure` VALUES (20891, 7524, 7525);
INSERT INTO `menu_closure` VALUES (20892, 7368, 7533);
INSERT INTO `menu_closure` VALUES (20893, 7493, 7533);
INSERT INTO `menu_closure` VALUES (20894, 7524, 7533);
INSERT INTO `menu_closure` VALUES (20895, 7368, 7494);
INSERT INTO `menu_closure` VALUES (20896, 7493, 7494);
INSERT INTO `menu_closure` VALUES (20897, 7368, 7503);
INSERT INTO `menu_closure` VALUES (20898, 7493, 7503);
INSERT INTO `menu_closure` VALUES (20899, 7368, 7510);
INSERT INTO `menu_closure` VALUES (20900, 7493, 7510);
INSERT INTO `menu_closure` VALUES (20901, 7368, 7517);
INSERT INTO `menu_closure` VALUES (20902, 7493, 7517);
INSERT INTO `menu_closure` VALUES (20903, 7368, 7524);
INSERT INTO `menu_closure` VALUES (20904, 7493, 7524);
INSERT INTO `menu_closure` VALUES (20905, 7368, 7538);
INSERT INTO `menu_closure` VALUES (20906, 7536, 7538);
INSERT INTO `menu_closure` VALUES (20907, 7537, 7538);
INSERT INTO `menu_closure` VALUES (20908, 7368, 7539);
INSERT INTO `menu_closure` VALUES (20909, 7536, 7539);
INSERT INTO `menu_closure` VALUES (20910, 7537, 7539);
INSERT INTO `menu_closure` VALUES (20911, 7368, 7540);
INSERT INTO `menu_closure` VALUES (20912, 7536, 7540);
INSERT INTO `menu_closure` VALUES (20913, 7537, 7540);
INSERT INTO `menu_closure` VALUES (20914, 7368, 7541);
INSERT INTO `menu_closure` VALUES (20915, 7536, 7541);
INSERT INTO `menu_closure` VALUES (20916, 7537, 7541);
INSERT INTO `menu_closure` VALUES (20917, 7368, 7542);
INSERT INTO `menu_closure` VALUES (20918, 7536, 7542);
INSERT INTO `menu_closure` VALUES (20919, 7537, 7542);
INSERT INTO `menu_closure` VALUES (20920, 7368, 7543);
INSERT INTO `menu_closure` VALUES (20921, 7536, 7543);
INSERT INTO `menu_closure` VALUES (20922, 7537, 7543);
INSERT INTO `menu_closure` VALUES (20923, 7368, 7544);
INSERT INTO `menu_closure` VALUES (20924, 7536, 7544);
INSERT INTO `menu_closure` VALUES (20925, 7537, 7544);
INSERT INTO `menu_closure` VALUES (20926, 7368, 7545);
INSERT INTO `menu_closure` VALUES (20927, 7536, 7545);
INSERT INTO `menu_closure` VALUES (20928, 7537, 7545);
INSERT INTO `menu_closure` VALUES (20929, 7368, 7546);
INSERT INTO `menu_closure` VALUES (20930, 7536, 7546);
INSERT INTO `menu_closure` VALUES (20931, 7537, 7546);
INSERT INTO `menu_closure` VALUES (20932, 7368, 7560);
INSERT INTO `menu_closure` VALUES (20933, 7536, 7560);
INSERT INTO `menu_closure` VALUES (20934, 7547, 7560);
INSERT INTO `menu_closure` VALUES (20935, 7559, 7560);
INSERT INTO `menu_closure` VALUES (20936, 7368, 7561);
INSERT INTO `menu_closure` VALUES (20937, 7536, 7561);
INSERT INTO `menu_closure` VALUES (20938, 7547, 7561);
INSERT INTO `menu_closure` VALUES (20939, 7559, 7561);
INSERT INTO `menu_closure` VALUES (20940, 7368, 7548);
INSERT INTO `menu_closure` VALUES (20941, 7536, 7548);
INSERT INTO `menu_closure` VALUES (20942, 7547, 7548);
INSERT INTO `menu_closure` VALUES (20943, 7368, 7549);
INSERT INTO `menu_closure` VALUES (20944, 7536, 7549);
INSERT INTO `menu_closure` VALUES (20945, 7547, 7549);
INSERT INTO `menu_closure` VALUES (20946, 7368, 7550);
INSERT INTO `menu_closure` VALUES (20947, 7536, 7550);
INSERT INTO `menu_closure` VALUES (20948, 7547, 7550);
INSERT INTO `menu_closure` VALUES (20949, 7368, 7551);
INSERT INTO `menu_closure` VALUES (20950, 7536, 7551);
INSERT INTO `menu_closure` VALUES (20951, 7547, 7551);
INSERT INTO `menu_closure` VALUES (20952, 7368, 7552);
INSERT INTO `menu_closure` VALUES (20953, 7536, 7552);
INSERT INTO `menu_closure` VALUES (20954, 7547, 7552);
INSERT INTO `menu_closure` VALUES (20955, 7368, 7553);
INSERT INTO `menu_closure` VALUES (20956, 7536, 7553);
INSERT INTO `menu_closure` VALUES (20957, 7547, 7553);
INSERT INTO `menu_closure` VALUES (20958, 7368, 7554);
INSERT INTO `menu_closure` VALUES (20959, 7536, 7554);
INSERT INTO `menu_closure` VALUES (20960, 7547, 7554);
INSERT INTO `menu_closure` VALUES (20961, 7368, 7555);
INSERT INTO `menu_closure` VALUES (20962, 7536, 7555);
INSERT INTO `menu_closure` VALUES (20963, 7547, 7555);
INSERT INTO `menu_closure` VALUES (20964, 7368, 7556);
INSERT INTO `menu_closure` VALUES (20965, 7536, 7556);
INSERT INTO `menu_closure` VALUES (20966, 7547, 7556);
INSERT INTO `menu_closure` VALUES (20967, 7368, 7557);
INSERT INTO `menu_closure` VALUES (20968, 7536, 7557);
INSERT INTO `menu_closure` VALUES (20969, 7547, 7557);
INSERT INTO `menu_closure` VALUES (20970, 7368, 7558);
INSERT INTO `menu_closure` VALUES (20971, 7536, 7558);
INSERT INTO `menu_closure` VALUES (20972, 7547, 7558);
INSERT INTO `menu_closure` VALUES (20973, 7368, 7559);
INSERT INTO `menu_closure` VALUES (20974, 7536, 7559);
INSERT INTO `menu_closure` VALUES (20975, 7547, 7559);
INSERT INTO `menu_closure` VALUES (20976, 7368, 7537);
INSERT INTO `menu_closure` VALUES (20977, 7536, 7537);
INSERT INTO `menu_closure` VALUES (20978, 7368, 7547);
INSERT INTO `menu_closure` VALUES (20979, 7536, 7547);
INSERT INTO `menu_closure` VALUES (20980, 7368, 7369);
INSERT INTO `menu_closure` VALUES (20981, 7368, 7370);
INSERT INTO `menu_closure` VALUES (20982, 7368, 7431);
INSERT INTO `menu_closure` VALUES (20983, 7368, 7446);
INSERT INTO `menu_closure` VALUES (20984, 7368, 7493);
INSERT INTO `menu_closure` VALUES (20985, 7368, 7536);
INSERT INTO `menu_closure` VALUES (20986, 7562, 7566);
INSERT INTO `menu_closure` VALUES (20987, 7564, 7566);
INSERT INTO `menu_closure` VALUES (20988, 7565, 7566);
INSERT INTO `menu_closure` VALUES (20989, 7562, 7567);
INSERT INTO `menu_closure` VALUES (20990, 7564, 7567);
INSERT INTO `menu_closure` VALUES (20991, 7565, 7567);
INSERT INTO `menu_closure` VALUES (20992, 7562, 7568);
INSERT INTO `menu_closure` VALUES (20993, 7564, 7568);
INSERT INTO `menu_closure` VALUES (20994, 7565, 7568);
INSERT INTO `menu_closure` VALUES (20995, 7562, 7569);
INSERT INTO `menu_closure` VALUES (20996, 7564, 7569);
INSERT INTO `menu_closure` VALUES (20997, 7565, 7569);
INSERT INTO `menu_closure` VALUES (20998, 7562, 7571);
INSERT INTO `menu_closure` VALUES (20999, 7564, 7571);
INSERT INTO `menu_closure` VALUES (21000, 7570, 7571);
INSERT INTO `menu_closure` VALUES (21001, 7562, 7572);
INSERT INTO `menu_closure` VALUES (21002, 7564, 7572);
INSERT INTO `menu_closure` VALUES (21003, 7570, 7572);
INSERT INTO `menu_closure` VALUES (21004, 7562, 7573);
INSERT INTO `menu_closure` VALUES (21005, 7564, 7573);
INSERT INTO `menu_closure` VALUES (21006, 7570, 7573);
INSERT INTO `menu_closure` VALUES (21007, 7562, 7574);
INSERT INTO `menu_closure` VALUES (21008, 7564, 7574);
INSERT INTO `menu_closure` VALUES (21009, 7570, 7574);
INSERT INTO `menu_closure` VALUES (21010, 7562, 7575);
INSERT INTO `menu_closure` VALUES (21011, 7564, 7575);
INSERT INTO `menu_closure` VALUES (21012, 7570, 7575);
INSERT INTO `menu_closure` VALUES (21013, 7562, 7578);
INSERT INTO `menu_closure` VALUES (21014, 7564, 7578);
INSERT INTO `menu_closure` VALUES (21015, 7576, 7578);
INSERT INTO `menu_closure` VALUES (21016, 7577, 7578);
INSERT INTO `menu_closure` VALUES (21017, 7562, 7579);
INSERT INTO `menu_closure` VALUES (21018, 7564, 7579);
INSERT INTO `menu_closure` VALUES (21019, 7576, 7579);
INSERT INTO `menu_closure` VALUES (21020, 7577, 7579);
INSERT INTO `menu_closure` VALUES (21021, 7562, 7580);
INSERT INTO `menu_closure` VALUES (21022, 7564, 7580);
INSERT INTO `menu_closure` VALUES (21023, 7576, 7580);
INSERT INTO `menu_closure` VALUES (21024, 7577, 7580);
INSERT INTO `menu_closure` VALUES (21025, 7562, 7581);
INSERT INTO `menu_closure` VALUES (21026, 7564, 7581);
INSERT INTO `menu_closure` VALUES (21027, 7576, 7581);
INSERT INTO `menu_closure` VALUES (21028, 7577, 7581);
INSERT INTO `menu_closure` VALUES (21029, 7562, 7583);
INSERT INTO `menu_closure` VALUES (21030, 7564, 7583);
INSERT INTO `menu_closure` VALUES (21031, 7576, 7583);
INSERT INTO `menu_closure` VALUES (21032, 7582, 7583);
INSERT INTO `menu_closure` VALUES (21033, 7562, 7584);
INSERT INTO `menu_closure` VALUES (21034, 7564, 7584);
INSERT INTO `menu_closure` VALUES (21035, 7576, 7584);
INSERT INTO `menu_closure` VALUES (21036, 7582, 7584);
INSERT INTO `menu_closure` VALUES (21037, 7562, 7585);
INSERT INTO `menu_closure` VALUES (21038, 7564, 7585);
INSERT INTO `menu_closure` VALUES (21039, 7576, 7585);
INSERT INTO `menu_closure` VALUES (21040, 7582, 7585);
INSERT INTO `menu_closure` VALUES (21041, 7562, 7586);
INSERT INTO `menu_closure` VALUES (21042, 7564, 7586);
INSERT INTO `menu_closure` VALUES (21043, 7576, 7586);
INSERT INTO `menu_closure` VALUES (21044, 7582, 7586);
INSERT INTO `menu_closure` VALUES (21045, 7562, 7587);
INSERT INTO `menu_closure` VALUES (21046, 7564, 7587);
INSERT INTO `menu_closure` VALUES (21047, 7576, 7587);
INSERT INTO `menu_closure` VALUES (21048, 7582, 7587);
INSERT INTO `menu_closure` VALUES (21049, 7562, 7589);
INSERT INTO `menu_closure` VALUES (21050, 7564, 7589);
INSERT INTO `menu_closure` VALUES (21051, 7576, 7589);
INSERT INTO `menu_closure` VALUES (21052, 7588, 7589);
INSERT INTO `menu_closure` VALUES (21053, 7562, 7590);
INSERT INTO `menu_closure` VALUES (21054, 7564, 7590);
INSERT INTO `menu_closure` VALUES (21055, 7576, 7590);
INSERT INTO `menu_closure` VALUES (21056, 7588, 7590);
INSERT INTO `menu_closure` VALUES (21057, 7562, 7577);
INSERT INTO `menu_closure` VALUES (21058, 7564, 7577);
INSERT INTO `menu_closure` VALUES (21059, 7576, 7577);
INSERT INTO `menu_closure` VALUES (21060, 7562, 7582);
INSERT INTO `menu_closure` VALUES (21061, 7564, 7582);
INSERT INTO `menu_closure` VALUES (21062, 7576, 7582);
INSERT INTO `menu_closure` VALUES (21063, 7562, 7588);
INSERT INTO `menu_closure` VALUES (21064, 7564, 7588);
INSERT INTO `menu_closure` VALUES (21065, 7576, 7588);
INSERT INTO `menu_closure` VALUES (21066, 7562, 7593);
INSERT INTO `menu_closure` VALUES (21067, 7564, 7593);
INSERT INTO `menu_closure` VALUES (21068, 7591, 7593);
INSERT INTO `menu_closure` VALUES (21069, 7592, 7593);
INSERT INTO `menu_closure` VALUES (21070, 7562, 7594);
INSERT INTO `menu_closure` VALUES (21071, 7564, 7594);
INSERT INTO `menu_closure` VALUES (21072, 7591, 7594);
INSERT INTO `menu_closure` VALUES (21073, 7592, 7594);
INSERT INTO `menu_closure` VALUES (21074, 7562, 7595);
INSERT INTO `menu_closure` VALUES (21075, 7564, 7595);
INSERT INTO `menu_closure` VALUES (21076, 7591, 7595);
INSERT INTO `menu_closure` VALUES (21077, 7592, 7595);
INSERT INTO `menu_closure` VALUES (21078, 7562, 7596);
INSERT INTO `menu_closure` VALUES (21079, 7564, 7596);
INSERT INTO `menu_closure` VALUES (21080, 7591, 7596);
INSERT INTO `menu_closure` VALUES (21081, 7592, 7596);
INSERT INTO `menu_closure` VALUES (21082, 7562, 7598);
INSERT INTO `menu_closure` VALUES (21083, 7564, 7598);
INSERT INTO `menu_closure` VALUES (21084, 7591, 7598);
INSERT INTO `menu_closure` VALUES (21085, 7597, 7598);
INSERT INTO `menu_closure` VALUES (21086, 7562, 7599);
INSERT INTO `menu_closure` VALUES (21087, 7564, 7599);
INSERT INTO `menu_closure` VALUES (21088, 7591, 7599);
INSERT INTO `menu_closure` VALUES (21089, 7597, 7599);
INSERT INTO `menu_closure` VALUES (21090, 7562, 7600);
INSERT INTO `menu_closure` VALUES (21091, 7564, 7600);
INSERT INTO `menu_closure` VALUES (21092, 7591, 7600);
INSERT INTO `menu_closure` VALUES (21093, 7597, 7600);
INSERT INTO `menu_closure` VALUES (21094, 7562, 7601);
INSERT INTO `menu_closure` VALUES (21095, 7564, 7601);
INSERT INTO `menu_closure` VALUES (21096, 7591, 7601);
INSERT INTO `menu_closure` VALUES (21097, 7597, 7601);
INSERT INTO `menu_closure` VALUES (21098, 7562, 7592);
INSERT INTO `menu_closure` VALUES (21099, 7564, 7592);
INSERT INTO `menu_closure` VALUES (21100, 7591, 7592);
INSERT INTO `menu_closure` VALUES (21101, 7562, 7597);
INSERT INTO `menu_closure` VALUES (21102, 7564, 7597);
INSERT INTO `menu_closure` VALUES (21103, 7591, 7597);
INSERT INTO `menu_closure` VALUES (21104, 7562, 7604);
INSERT INTO `menu_closure` VALUES (21105, 7564, 7604);
INSERT INTO `menu_closure` VALUES (21106, 7602, 7604);
INSERT INTO `menu_closure` VALUES (21107, 7603, 7604);
INSERT INTO `menu_closure` VALUES (21108, 7562, 7605);
INSERT INTO `menu_closure` VALUES (21109, 7564, 7605);
INSERT INTO `menu_closure` VALUES (21110, 7602, 7605);
INSERT INTO `menu_closure` VALUES (21111, 7603, 7605);
INSERT INTO `menu_closure` VALUES (21112, 7562, 7606);
INSERT INTO `menu_closure` VALUES (21113, 7564, 7606);
INSERT INTO `menu_closure` VALUES (21114, 7602, 7606);
INSERT INTO `menu_closure` VALUES (21115, 7603, 7606);
INSERT INTO `menu_closure` VALUES (21116, 7562, 7607);
INSERT INTO `menu_closure` VALUES (21117, 7564, 7607);
INSERT INTO `menu_closure` VALUES (21118, 7602, 7607);
INSERT INTO `menu_closure` VALUES (21119, 7603, 7607);
INSERT INTO `menu_closure` VALUES (21120, 7562, 7609);
INSERT INTO `menu_closure` VALUES (21121, 7564, 7609);
INSERT INTO `menu_closure` VALUES (21122, 7602, 7609);
INSERT INTO `menu_closure` VALUES (21123, 7608, 7609);
INSERT INTO `menu_closure` VALUES (21124, 7562, 7610);
INSERT INTO `menu_closure` VALUES (21125, 7564, 7610);
INSERT INTO `menu_closure` VALUES (21126, 7602, 7610);
INSERT INTO `menu_closure` VALUES (21127, 7608, 7610);
INSERT INTO `menu_closure` VALUES (21128, 7562, 7611);
INSERT INTO `menu_closure` VALUES (21129, 7564, 7611);
INSERT INTO `menu_closure` VALUES (21130, 7602, 7611);
INSERT INTO `menu_closure` VALUES (21131, 7608, 7611);
INSERT INTO `menu_closure` VALUES (21132, 7562, 7612);
INSERT INTO `menu_closure` VALUES (21133, 7564, 7612);
INSERT INTO `menu_closure` VALUES (21134, 7602, 7612);
INSERT INTO `menu_closure` VALUES (21135, 7608, 7612);
INSERT INTO `menu_closure` VALUES (21136, 7562, 7613);
INSERT INTO `menu_closure` VALUES (21137, 7564, 7613);
INSERT INTO `menu_closure` VALUES (21138, 7602, 7613);
INSERT INTO `menu_closure` VALUES (21139, 7608, 7613);
INSERT INTO `menu_closure` VALUES (21140, 7562, 7603);
INSERT INTO `menu_closure` VALUES (21141, 7564, 7603);
INSERT INTO `menu_closure` VALUES (21142, 7602, 7603);
INSERT INTO `menu_closure` VALUES (21143, 7562, 7608);
INSERT INTO `menu_closure` VALUES (21144, 7564, 7608);
INSERT INTO `menu_closure` VALUES (21145, 7602, 7608);
INSERT INTO `menu_closure` VALUES (21146, 7562, 7616);
INSERT INTO `menu_closure` VALUES (21147, 7564, 7616);
INSERT INTO `menu_closure` VALUES (21148, 7614, 7616);
INSERT INTO `menu_closure` VALUES (21149, 7615, 7616);
INSERT INTO `menu_closure` VALUES (21150, 7562, 7617);
INSERT INTO `menu_closure` VALUES (21151, 7564, 7617);
INSERT INTO `menu_closure` VALUES (21152, 7614, 7617);
INSERT INTO `menu_closure` VALUES (21153, 7615, 7617);
INSERT INTO `menu_closure` VALUES (21154, 7562, 7618);
INSERT INTO `menu_closure` VALUES (21155, 7564, 7618);
INSERT INTO `menu_closure` VALUES (21156, 7614, 7618);
INSERT INTO `menu_closure` VALUES (21157, 7615, 7618);
INSERT INTO `menu_closure` VALUES (21158, 7562, 7619);
INSERT INTO `menu_closure` VALUES (21159, 7564, 7619);
INSERT INTO `menu_closure` VALUES (21160, 7614, 7619);
INSERT INTO `menu_closure` VALUES (21161, 7615, 7619);
INSERT INTO `menu_closure` VALUES (21162, 7562, 7621);
INSERT INTO `menu_closure` VALUES (21163, 7564, 7621);
INSERT INTO `menu_closure` VALUES (21164, 7614, 7621);
INSERT INTO `menu_closure` VALUES (21165, 7620, 7621);
INSERT INTO `menu_closure` VALUES (21166, 7562, 7622);
INSERT INTO `menu_closure` VALUES (21167, 7564, 7622);
INSERT INTO `menu_closure` VALUES (21168, 7614, 7622);
INSERT INTO `menu_closure` VALUES (21169, 7620, 7622);
INSERT INTO `menu_closure` VALUES (21170, 7562, 7623);
INSERT INTO `menu_closure` VALUES (21171, 7564, 7623);
INSERT INTO `menu_closure` VALUES (21172, 7614, 7623);
INSERT INTO `menu_closure` VALUES (21173, 7620, 7623);
INSERT INTO `menu_closure` VALUES (21174, 7562, 7624);
INSERT INTO `menu_closure` VALUES (21175, 7564, 7624);
INSERT INTO `menu_closure` VALUES (21176, 7614, 7624);
INSERT INTO `menu_closure` VALUES (21177, 7620, 7624);
INSERT INTO `menu_closure` VALUES (21178, 7562, 7625);
INSERT INTO `menu_closure` VALUES (21179, 7564, 7625);
INSERT INTO `menu_closure` VALUES (21180, 7614, 7625);
INSERT INTO `menu_closure` VALUES (21181, 7620, 7625);
INSERT INTO `menu_closure` VALUES (21182, 7562, 7615);
INSERT INTO `menu_closure` VALUES (21183, 7564, 7615);
INSERT INTO `menu_closure` VALUES (21184, 7614, 7615);
INSERT INTO `menu_closure` VALUES (21185, 7562, 7620);
INSERT INTO `menu_closure` VALUES (21186, 7564, 7620);
INSERT INTO `menu_closure` VALUES (21187, 7614, 7620);
INSERT INTO `menu_closure` VALUES (21188, 7562, 7565);
INSERT INTO `menu_closure` VALUES (21189, 7564, 7565);
INSERT INTO `menu_closure` VALUES (21190, 7562, 7570);
INSERT INTO `menu_closure` VALUES (21191, 7564, 7570);
INSERT INTO `menu_closure` VALUES (21192, 7562, 7576);
INSERT INTO `menu_closure` VALUES (21193, 7564, 7576);
INSERT INTO `menu_closure` VALUES (21194, 7562, 7591);
INSERT INTO `menu_closure` VALUES (21195, 7564, 7591);
INSERT INTO `menu_closure` VALUES (21196, 7562, 7602);
INSERT INTO `menu_closure` VALUES (21197, 7564, 7602);
INSERT INTO `menu_closure` VALUES (21198, 7562, 7614);
INSERT INTO `menu_closure` VALUES (21199, 7564, 7614);
INSERT INTO `menu_closure` VALUES (21200, 7562, 7628);
INSERT INTO `menu_closure` VALUES (21201, 7626, 7628);
INSERT INTO `menu_closure` VALUES (21202, 7627, 7628);
INSERT INTO `menu_closure` VALUES (21203, 7562, 7629);
INSERT INTO `menu_closure` VALUES (21204, 7626, 7629);
INSERT INTO `menu_closure` VALUES (21205, 7627, 7629);
INSERT INTO `menu_closure` VALUES (21206, 7562, 7630);
INSERT INTO `menu_closure` VALUES (21207, 7626, 7630);
INSERT INTO `menu_closure` VALUES (21208, 7627, 7630);
INSERT INTO `menu_closure` VALUES (21209, 7562, 7631);
INSERT INTO `menu_closure` VALUES (21210, 7626, 7631);
INSERT INTO `menu_closure` VALUES (21211, 7627, 7631);
INSERT INTO `menu_closure` VALUES (21212, 7562, 7633);
INSERT INTO `menu_closure` VALUES (21213, 7626, 7633);
INSERT INTO `menu_closure` VALUES (21214, 7632, 7633);
INSERT INTO `menu_closure` VALUES (21215, 7562, 7634);
INSERT INTO `menu_closure` VALUES (21216, 7626, 7634);
INSERT INTO `menu_closure` VALUES (21217, 7632, 7634);
INSERT INTO `menu_closure` VALUES (21218, 7562, 7635);
INSERT INTO `menu_closure` VALUES (21219, 7626, 7635);
INSERT INTO `menu_closure` VALUES (21220, 7632, 7635);
INSERT INTO `menu_closure` VALUES (21221, 7562, 7636);
INSERT INTO `menu_closure` VALUES (21222, 7626, 7636);
INSERT INTO `menu_closure` VALUES (21223, 7632, 7636);
INSERT INTO `menu_closure` VALUES (21224, 7562, 7637);
INSERT INTO `menu_closure` VALUES (21225, 7626, 7637);
INSERT INTO `menu_closure` VALUES (21226, 7632, 7637);
INSERT INTO `menu_closure` VALUES (21227, 7562, 7639);
INSERT INTO `menu_closure` VALUES (21228, 7626, 7639);
INSERT INTO `menu_closure` VALUES (21229, 7638, 7639);
INSERT INTO `menu_closure` VALUES (21230, 7562, 7640);
INSERT INTO `menu_closure` VALUES (21231, 7626, 7640);
INSERT INTO `menu_closure` VALUES (21232, 7638, 7640);
INSERT INTO `menu_closure` VALUES (21233, 7562, 7641);
INSERT INTO `menu_closure` VALUES (21234, 7626, 7641);
INSERT INTO `menu_closure` VALUES (21235, 7638, 7641);
INSERT INTO `menu_closure` VALUES (21236, 7562, 7642);
INSERT INTO `menu_closure` VALUES (21237, 7626, 7642);
INSERT INTO `menu_closure` VALUES (21238, 7638, 7642);
INSERT INTO `menu_closure` VALUES (21239, 7562, 7643);
INSERT INTO `menu_closure` VALUES (21240, 7626, 7643);
INSERT INTO `menu_closure` VALUES (21241, 7638, 7643);
INSERT INTO `menu_closure` VALUES (21242, 7562, 7645);
INSERT INTO `menu_closure` VALUES (21243, 7626, 7645);
INSERT INTO `menu_closure` VALUES (21244, 7644, 7645);
INSERT INTO `menu_closure` VALUES (21245, 7562, 7646);
INSERT INTO `menu_closure` VALUES (21246, 7626, 7646);
INSERT INTO `menu_closure` VALUES (21247, 7644, 7646);
INSERT INTO `menu_closure` VALUES (21248, 7562, 7647);
INSERT INTO `menu_closure` VALUES (21249, 7626, 7647);
INSERT INTO `menu_closure` VALUES (21250, 7644, 7647);
INSERT INTO `menu_closure` VALUES (21251, 7562, 7648);
INSERT INTO `menu_closure` VALUES (21252, 7626, 7648);
INSERT INTO `menu_closure` VALUES (21253, 7644, 7648);
INSERT INTO `menu_closure` VALUES (21254, 7562, 7649);
INSERT INTO `menu_closure` VALUES (21255, 7626, 7649);
INSERT INTO `menu_closure` VALUES (21256, 7644, 7649);
INSERT INTO `menu_closure` VALUES (21257, 7562, 7651);
INSERT INTO `menu_closure` VALUES (21258, 7626, 7651);
INSERT INTO `menu_closure` VALUES (21259, 7650, 7651);
INSERT INTO `menu_closure` VALUES (21260, 7562, 7653);
INSERT INTO `menu_closure` VALUES (21261, 7626, 7653);
INSERT INTO `menu_closure` VALUES (21262, 7652, 7653);
INSERT INTO `menu_closure` VALUES (21263, 7562, 7654);
INSERT INTO `menu_closure` VALUES (21264, 7626, 7654);
INSERT INTO `menu_closure` VALUES (21265, 7652, 7654);
INSERT INTO `menu_closure` VALUES (21266, 7562, 7655);
INSERT INTO `menu_closure` VALUES (21267, 7626, 7655);
INSERT INTO `menu_closure` VALUES (21268, 7652, 7655);
INSERT INTO `menu_closure` VALUES (21269, 7562, 7656);
INSERT INTO `menu_closure` VALUES (21270, 7626, 7656);
INSERT INTO `menu_closure` VALUES (21271, 7652, 7656);
INSERT INTO `menu_closure` VALUES (21272, 7562, 7658);
INSERT INTO `menu_closure` VALUES (21273, 7626, 7658);
INSERT INTO `menu_closure` VALUES (21274, 7657, 7658);
INSERT INTO `menu_closure` VALUES (21275, 7562, 7659);
INSERT INTO `menu_closure` VALUES (21276, 7626, 7659);
INSERT INTO `menu_closure` VALUES (21277, 7657, 7659);
INSERT INTO `menu_closure` VALUES (21278, 7562, 7627);
INSERT INTO `menu_closure` VALUES (21279, 7626, 7627);
INSERT INTO `menu_closure` VALUES (21280, 7562, 7632);
INSERT INTO `menu_closure` VALUES (21281, 7626, 7632);
INSERT INTO `menu_closure` VALUES (21282, 7562, 7638);
INSERT INTO `menu_closure` VALUES (21283, 7626, 7638);
INSERT INTO `menu_closure` VALUES (21284, 7562, 7644);
INSERT INTO `menu_closure` VALUES (21285, 7626, 7644);
INSERT INTO `menu_closure` VALUES (21286, 7562, 7650);
INSERT INTO `menu_closure` VALUES (21287, 7626, 7650);
INSERT INTO `menu_closure` VALUES (21288, 7562, 7652);
INSERT INTO `menu_closure` VALUES (21289, 7626, 7652);
INSERT INTO `menu_closure` VALUES (21290, 7562, 7657);
INSERT INTO `menu_closure` VALUES (21291, 7626, 7657);
INSERT INTO `menu_closure` VALUES (21292, 7562, 7563);
INSERT INTO `menu_closure` VALUES (21293, 7562, 7564);
INSERT INTO `menu_closure` VALUES (21294, 7562, 7626);
COMMIT;

-- ----------------------------
-- Table structure for resource
-- ----------------------------
DROP TABLE IF EXISTS `resource`;
CREATE TABLE `resource` (
                            `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '主键ID',
                            `keyword` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT '关键字',
                            `department_id` bigint unsigned NOT NULL COMMENT '部门id',
                            `resource_id` bigint unsigned NOT NULL COMMENT '资源id',
                            PRIMARY KEY (`id`),
                            UNIQUE KEY `department_id` (`keyword`,`department_id`,`resource_id`),
                            KEY `department_id_2` (`department_id`),
                            CONSTRAINT `resource_ibfk_1` FOREIGN KEY (`department_id`) REFERENCES `department` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ;

-- ----------------------------
-- Records of resource
-- ----------------------------
BEGIN;
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
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=utf8mb4  COMMENT='角色信息';

-- ----------------------------
-- Records of role
-- ----------------------------
BEGIN;
INSERT INTO `role` VALUES (1, 0, '超级管理员', 'superAdmin', 1, '超级管理员  ', NULL, 'ALL', 1713706137, 1713706137, 0);
INSERT INTO `role` VALUES (5, 1, '21', '3', 1, '412', NULL, 'CUR_DOWN', 1719464519, 1721837751, 0);
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
) ENGINE=InnoDB AUTO_INCREMENT=12 DEFAULT CHARSET=utf8mb4  COMMENT='角色层级信息';

-- ----------------------------
-- Records of role_closure
-- ----------------------------
BEGIN;
INSERT INTO `role_closure` VALUES (5, 1, 5);
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
) ENGINE=InnoDB AUTO_INCREMENT=2289 DEFAULT CHARSET=utf8mb4  COMMENT='角色菜单信息';

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
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4  COMMENT='用户信息';

-- ----------------------------
-- Records of user
-- ----------------------------
BEGIN;
INSERT INTO `user` VALUES (1, 1, 1, '超级管理员', '超级管理员', 'F', 'a9f224627346905e258d771e4043f921', '1280291001@qq.com', '18888888888', '$2a$10$9qRJe9KQo6sEcU8ipKg.e.dkl2E7Wy64SigYlgraTAn.1paHFq6W.', 1, '{\"theme\":\"light\",\"themeColor\":\"#165DFF\",\"skin\":\"default\",\"tabBar\":true,\"menuWidth\":200,\"layout\":\"default\",\"language\":\"zh_CN\",\"animation\":\"gp-fade\"}', 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJkZXBhcnRtZW50SWQiOjEsImRlcGFydG1lbnRLZXl3b3JkIjoiY29tcGFueSIsImV4cCI6MTcyNDc1MTczNywiaWF0IjoxNzI0NzQ0NTM2LCJuYmYiOjE3MjQ3NDQ1MzYsInJvbGVJZCI6MSwicm9sZUtleXdvcmQiOiJzdXBlckFkbWluIiwidXNlcklkIjoxfQ.kZ2y12xxhJXRdfIv_43vkrE0E23XI62TufeYtSBQWZo', 1724744536, 1713706137, 1724744536);
INSERT INTO `user` VALUES (4, 5, 5, '1', '1', 'F', '', '31@q.com', '18286219255', '$2a$10$lpaNcnRMxC3oqVx4sYr16OvLy5BY1j8iFyUPyqAnprnC1wK/QsySq', 1, NULL, NULL, 0, 1721840505, 1721840505);
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
) ENGINE=InnoDB AUTO_INCREMENT=13 DEFAULT CHARSET=utf8mb4  COMMENT='用户职位信息';

-- ----------------------------
-- Records of user_job
-- ----------------------------
BEGIN;
INSERT INTO `user_job` VALUES (1, 1, 1);
INSERT INTO `user_job` VALUES (12, 1, 4);
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
) ENGINE=InnoDB AUTO_INCREMENT=11 DEFAULT CHARSET=utf8mb4  COMMENT='用户角色信息';

-- ----------------------------
-- Records of user_role
-- ----------------------------
BEGIN;
INSERT INTO `user_role` VALUES (1, 1, 1);
INSERT INTO `user_role` VALUES (10, 5, 4);
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
