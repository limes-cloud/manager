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
                                                                                  (996, 'p', '3', '/application/api/v1/auth', 'DELETE', '', '', ''),
                                                                                  (994, 'p', '3', '/application/api/v1/auth', 'POST', '', '', ''),
                                                                                  (995, 'p', '3', '/application/api/v1/auth/status', 'PUT', '', '', ''),
                                                                                  (993, 'p', '3', '/application/api/v1/auths', 'GET', '', '', ''),
                                                                                  (986, 'p', '3', '/application/api/v1/feedback', 'DELETE', '', '', ''),
                                                                                  (984, 'p', '3', '/application/api/v1/feedback', 'POST', '', '', ''),
                                                                                  (985, 'p', '3', '/application/api/v1/feedback', 'PUT', '', '', ''),
                                                                                  (983, 'p', '3', '/application/api/v1/feedbacks', 'GET', '', '', ''),
                                                                                  (979, 'p', '3', '/application/api/v1/feedback_categories', 'GET', '', '', ''),
                                                                                  (982, 'p', '3', '/application/api/v1/feedback_category', 'DELETE', '', '', ''),
                                                                                  (980, 'p', '3', '/application/api/v1/feedback_category', 'POST', '', '', ''),
                                                                                  (981, 'p', '3', '/application/api/v1/feedback_category', 'PUT', '', '', ''),
                                                                                  (998, 'p', '3', '/application/api/v1/oauth', 'DELETE', '', '', ''),
                                                                                  (997, 'p', '3', '/application/api/v1/oauths', 'GET', '', '', ''),
                                                                                  (992, 'p', '3', '/application/api/v1/user', 'DELETE', '', '', ''),
                                                                                  (988, 'p', '3', '/application/api/v1/user', 'POST', '', '', ''),
                                                                                  (990, 'p', '3', '/application/api/v1/user', 'PUT', '', '', ''),
                                                                                  (991, 'p', '3', '/application/api/v1/user/status', 'PUT', '', '', ''),
                                                                                  (1002, 'p', '3', '/application/api/v1/userinfo', 'DELETE', '', '', ''),
                                                                                  (1000, 'p', '3', '/application/api/v1/userinfo', 'POST', '', '', ''),
                                                                                  (1001, 'p', '3', '/application/api/v1/userinfo', 'PUT', '', '', ''),
                                                                                  (999, 'p', '3', '/application/api/v1/userinfos', 'GET', '', '', ''),
                                                                                  (987, 'p', '3', '/application/api/v1/users', 'GET', '', '', ''),
                                                                                  (989, 'p', '3', '/application/api/v1/users', 'POST', '', '', ''),
                                                                                  (972, 'p', '3', '/manager/api/v1/department', 'DELETE', '', '', ''),
                                                                                  (970, 'p', '3', '/manager/api/v1/department', 'POST', '', '', ''),
                                                                                  (971, 'p', '3', '/manager/api/v1/department', 'PUT', '', '', ''),
                                                                                  (976, 'p', '3', '/manager/api/v1/role', 'DELETE', '', '', ''),
                                                                                  (973, 'p', '3', '/manager/api/v1/role', 'POST', '', '', ''),
                                                                                  (974, 'p', '3', '/manager/api/v1/role', 'PUT', '', '', ''),
                                                                                  (978, 'p', '3', '/manager/api/v1/role/menu', 'POST', '', '', ''),
                                                                                  (977, 'p', '3', '/manager/api/v1/role/menu_ids', 'GET', '', '', ''),
                                                                                  (975, 'p', '3', '/manager/api/v1/role/status', 'PUT', '', '', ''),
                                                                                  (1006, 'p', '3', '/partyaffairs/api/v1/banner', 'DELETE', '', '', ''),
                                                                                  (1004, 'p', '3', '/partyaffairs/api/v1/banner', 'POST', '', '', ''),
                                                                                  (1005, 'p', '3', '/partyaffairs/api/v1/banner', 'PUT', '', '', ''),
                                                                                  (1003, 'p', '3', '/partyaffairs/api/v1/banners', 'GET', '', '', ''),
                                                                                  (1021, 'p', '3', '/partyaffairs/api/v1/information', 'DELETE', '', '', ''),
                                                                                  (1018, 'p', '3', '/partyaffairs/api/v1/information', 'GET', '', '', ''),
                                                                                  (1019, 'p', '3', '/partyaffairs/api/v1/information', 'POST', '', '', ''),
                                                                                  (1020, 'p', '3', '/partyaffairs/api/v1/information', 'PUT', '', '', ''),
                                                                                  (1013, 'p', '3', '/partyaffairs/api/v1/information/classifies', 'GET', '', '', ''),
                                                                                  (1016, 'p', '3', '/partyaffairs/api/v1/information/classify', 'DELETE', '', '', ''),
                                                                                  (1014, 'p', '3', '/partyaffairs/api/v1/information/classify', 'POST', '', '', ''),
                                                                                  (1015, 'p', '3', '/partyaffairs/api/v1/information/classify', 'PUT', '', '', ''),
                                                                                  (1017, 'p', '3', '/partyaffairs/api/v1/informations', 'GET', '', '', ''),
                                                                                  (1012, 'p', '3', '/partyaffairs/api/v1/notice', 'DELETE', '', '', ''),
                                                                                  (1009, 'p', '3', '/partyaffairs/api/v1/notice', 'POST', '', '', ''),
                                                                                  (1011, 'p', '3', '/partyaffairs/api/v1/notice', 'PUT', '', '', ''),
                                                                                  (1010, 'p', '3', '/partyaffairs/api/v1/notice/push', 'POST', '', '', ''),
                                                                                  (1008, 'p', '3', '/partyaffairs/api/v1/notice/users', 'GET', '', '', ''),
                                                                                  (1007, 'p', '3', '/partyaffairs/api/v1/notices', 'GET', '', '', ''),
                                                                                  (1030, 'p', '3', '/partyaffairs/api/v1/resource', 'DELETE', '', '', ''),
                                                                                  (1027, 'p', '3', '/partyaffairs/api/v1/resource', 'GET', '', '', ''),
                                                                                  (1028, 'p', '3', '/partyaffairs/api/v1/resource', 'POST', '', '', ''),
                                                                                  (1029, 'p', '3', '/partyaffairs/api/v1/resource', 'PUT', '', '', ''),
                                                                                  (1022, 'p', '3', '/partyaffairs/api/v1/resource/classifies', 'GET', '', '', ''),
                                                                                  (1025, 'p', '3', '/partyaffairs/api/v1/resource/classify', 'DELETE', '', '', ''),
                                                                                  (1023, 'p', '3', '/partyaffairs/api/v1/resource/classify', 'POST', '', '', ''),
                                                                                  (1024, 'p', '3', '/partyaffairs/api/v1/resource/classify', 'PUT', '', '', ''),
                                                                                  (1026, 'p', '3', '/partyaffairs/api/v1/resources', 'GET', '', '', ''),
                                                                                  (1035, 'p', '3', '/partyaffairs/api/v1/task', 'DELETE', '', '', ''),
                                                                                  (1032, 'p', '3', '/partyaffairs/api/v1/task', 'GET', '', '', ''),
                                                                                  (1033, 'p', '3', '/partyaffairs/api/v1/task', 'POST', '', '', ''),
                                                                                  (1034, 'p', '3', '/partyaffairs/api/v1/task', 'PUT', '', '', ''),
                                                                                  (1040, 'p', '3', '/partyaffairs/api/v1/task/value', 'DELETE', '', '', ''),
                                                                                  (1037, 'p', '3', '/partyaffairs/api/v1/task/value', 'GET', '', '', ''),
                                                                                  (1039, 'p', '3', '/partyaffairs/api/v1/task/value', 'PUT', '', '', ''),
                                                                                  (1036, 'p', '3', '/partyaffairs/api/v1/task/values', 'GET', '', '', ''),
                                                                                  (1038, 'p', '3', '/partyaffairs/api/v1/task/values', 'POST', '', '', ''),
                                                                                  (1031, 'p', '3', '/partyaffairs/api/v1/tasks', 'GET', '', '', ''),
                                                                                  (906, 'p', 'povertyAdmin', '/configure/api/v1/business', 'DELETE', '', '', ''),
                                                                                  (904, 'p', 'povertyAdmin', '/configure/api/v1/business', 'POST', '', '', ''),
                                                                                  (905, 'p', 'povertyAdmin', '/configure/api/v1/business', 'PUT', '', '', ''),
                                                                                  (903, 'p', 'povertyAdmin', '/configure/api/v1/businesses', 'GET', '', '', ''),
                                                                                  (923, 'p', 'povertyAdmin', '/configure/api/v1/configure', 'PUT', '', '', ''),
                                                                                  (922, 'p', 'povertyAdmin', '/configure/api/v1/configure/compare', 'POST', '', '', ''),
                                                                                  (894, 'p', 'povertyAdmin', '/configure/api/v1/env', 'DELETE', '', '', ''),
                                                                                  (891, 'p', 'povertyAdmin', '/configure/api/v1/env', 'POST', '', '', ''),
                                                                                  (892, 'p', 'povertyAdmin', '/configure/api/v1/env', 'PUT', '', '', ''),
                                                                                  (893, 'p', 'povertyAdmin', '/configure/api/v1/env/status', 'PUT', '', '', ''),
                                                                                  (895, 'p', 'povertyAdmin', '/configure/api/v1/env/token', 'GET', '', '', ''),
                                                                                  (896, 'p', 'povertyAdmin', '/configure/api/v1/env/token', 'PUT', '', '', ''),
                                                                                  (889, 'p', 'povertyAdmin', '/configure/api/v1/envs', 'GET', '', '', ''),
                                                                                  (912, 'p', 'povertyAdmin', '/configure/api/v1/resource', 'DELETE', '', '', ''),
                                                                                  (910, 'p', 'povertyAdmin', '/configure/api/v1/resource', 'POST', '', '', ''),
                                                                                  (911, 'p', 'povertyAdmin', '/configure/api/v1/resource', 'PUT', '', '', ''),
                                                                                  (909, 'p', 'povertyAdmin', '/configure/api/v1/resources', 'GET', '', '', ''),
                                                                                  (902, 'p', 'povertyAdmin', '/configure/api/v1/server', 'DELETE', '', '', ''),
                                                                                  (899, 'p', 'povertyAdmin', '/configure/api/v1/server', 'POST', '', '', ''),
                                                                                  (900, 'p', 'povertyAdmin', '/configure/api/v1/server', 'PUT', '', '', ''),
                                                                                  (901, 'p', 'povertyAdmin', '/configure/api/v1/server/status', 'PUT', '', '', ''),
                                                                                  (897, 'p', 'povertyAdmin', '/configure/api/v1/servers', 'GET', '', '', ''),
                                                                                  (916, 'p', 'povertyAdmin', '/configure/api/v1/template', 'GET', '', '', ''),
                                                                                  (918, 'p', 'povertyAdmin', '/configure/api/v1/template', 'POST', '', '', ''),
                                                                                  (919, 'p', 'povertyAdmin', '/configure/api/v1/template/compare', 'POST', '', '', ''),
                                                                                  (917, 'p', 'povertyAdmin', '/configure/api/v1/template/current', 'GET', '', '', ''),
                                                                                  (921, 'p', 'povertyAdmin', '/configure/api/v1/template/preview', 'POST', '', '', ''),
                                                                                  (920, 'p', 'povertyAdmin', '/configure/api/v1/template/switch', 'POST', '', '', ''),
                                                                                  (915, 'p', 'povertyAdmin', '/configure/api/v1/templates', 'GET', '', '', ''),
                                                                                  (907, 'p', 'povertyAdmin', '/configure/business/values', 'get', '', '', ''),
                                                                                  (908, 'p', 'povertyAdmin', '/configure/business/values', 'PUT', '', '', ''),
                                                                                  (913, 'p', 'povertyAdmin', '/configure/resource/values', 'get', '', '', ''),
                                                                                  (914, 'p', 'povertyAdmin', '/configure/resource/values', 'PUT', '', '', ''),
                                                                                  (945, 'p', 'povertyAdmin', '/cron/api/v1/log', 'GET', '', '', ''),
                                                                                  (944, 'p', 'povertyAdmin', '/cron/api/v1/logs', 'GET', '', '', ''),
                                                                                  (942, 'p', 'povertyAdmin', '/cron/api/v1/task', 'DELETE', '', '', ''),
                                                                                  (938, 'p', 'povertyAdmin', '/cron/api/v1/task', 'POST', '', '', ''),
                                                                                  (941, 'p', 'povertyAdmin', '/cron/api/v1/task', 'PUT', '', '', ''),
                                                                                  (940, 'p', 'povertyAdmin', '/cron/api/v1/task/cancel', 'POST', '', '', ''),
                                                                                  (939, 'p', 'povertyAdmin', '/cron/api/v1/task/exec', 'POST', '', '', ''),
                                                                                  (943, 'p', 'povertyAdmin', '/cron/api/v1/task/status', 'PUT', '', '', ''),
                                                                                  (937, 'p', 'povertyAdmin', '/cron/api/v1/tasks', 'GET', '', '', ''),
                                                                                  (936, 'p', 'povertyAdmin', '/cron/api/v1/task_group', 'DELETE', '', '', ''),
                                                                                  (934, 'p', 'povertyAdmin', '/cron/api/v1/task_group', 'POST', '', '', ''),
                                                                                  (935, 'p', 'povertyAdmin', '/cron/api/v1/task_group', 'PUT', '', '', ''),
                                                                                  (933, 'p', 'povertyAdmin', '/cron/api/v1/task_groups', 'GET', '', '', ''),
                                                                                  (931, 'p', 'povertyAdmin', '/cron/api/v1/worker', 'DELETE', '', '', ''),
                                                                                  (929, 'p', 'povertyAdmin', '/cron/api/v1/worker', 'POST', '', '', ''),
                                                                                  (930, 'p', 'povertyAdmin', '/cron/api/v1/worker', 'PUT', '', '', ''),
                                                                                  (932, 'p', 'povertyAdmin', '/cron/api/v1/worker/status', 'PUT', '', '', ''),
                                                                                  (928, 'p', 'povertyAdmin', '/cron/api/v1/workers', 'GET', '', '', ''),
                                                                                  (927, 'p', 'povertyAdmin', '/cron/api/v1/worker_group', 'DELETE', '', '', ''),
                                                                                  (925, 'p', 'povertyAdmin', '/cron/api/v1/worker_group', 'POST', '', '', ''),
                                                                                  (926, 'p', 'povertyAdmin', '/cron/api/v1/worker_group', 'PUT', '', '', ''),
                                                                                  (924, 'p', 'povertyAdmin', '/cron/api/v1/worker_groups', 'GET', '', '', ''),
                                                                                  (840, 'p', 'povertyAdmin', '/manager/api/v1/department', 'DELETE', '', '', ''),
                                                                                  (838, 'p', 'povertyAdmin', '/manager/api/v1/department', 'POST', '', '', ''),
                                                                                  (839, 'p', 'povertyAdmin', '/manager/api/v1/department', 'PUT', '', '', ''),
                                                                                  (825, 'p', 'povertyAdmin', '/manager/api/v1/dictionaries', 'GET', '', '', ''),
                                                                                  (828, 'p', 'povertyAdmin', '/manager/api/v1/dictionary', 'DELETE', '', '', ''),
                                                                                  (826, 'p', 'povertyAdmin', '/manager/api/v1/dictionary', 'POST', '', '', ''),
                                                                                  (827, 'p', 'povertyAdmin', '/manager/api/v1/dictionary', 'PUT', '', '', ''),
                                                                                  (833, 'p', 'povertyAdmin', '/manager/api/v1/dictionary_value', 'DELETE', '', '', ''),
                                                                                  (830, 'p', 'povertyAdmin', '/manager/api/v1/dictionary_value', 'POST', '', '', ''),
                                                                                  (831, 'p', 'povertyAdmin', '/manager/api/v1/dictionary_value', 'PUT', '', '', ''),
                                                                                  (832, 'p', 'povertyAdmin', '/manager/api/v1/dictionary_value/status', 'PUT', '', '', ''),
                                                                                  (829, 'p', 'povertyAdmin', '/manager/api/v1/dictionary_values', 'GET', '', '', ''),
                                                                                  (837, 'p', 'povertyAdmin', '/manager/api/v1/job', 'DELETE', '', '', ''),
                                                                                  (835, 'p', 'povertyAdmin', '/manager/api/v1/job', 'POST', '', '', ''),
                                                                                  (836, 'p', 'povertyAdmin', '/manager/api/v1/job', 'PUT', '', '', ''),
                                                                                  (834, 'p', 'povertyAdmin', '/manager/api/v1/jobs', 'GET', '', '', ''),
                                                                                  (890, 'p', 'povertyAdmin', '/manager/api/v1/resource/cfg_env', 'PUT', '', '', ''),
                                                                                  (898, 'p', 'povertyAdmin', '/manager/api/v1/resource/cfg_server', 'PUT', '', '', ''),
                                                                                  (844, 'p', 'povertyAdmin', '/manager/api/v1/role', 'DELETE', '', '', ''),
                                                                                  (841, 'p', 'povertyAdmin', '/manager/api/v1/role', 'POST', '', '', ''),
                                                                                  (842, 'p', 'povertyAdmin', '/manager/api/v1/role', 'PUT', '', '', ''),
                                                                                  (846, 'p', 'povertyAdmin', '/manager/api/v1/role/menu', 'POST', '', '', ''),
                                                                                  (845, 'p', 'povertyAdmin', '/manager/api/v1/role/menu_ids', 'GET', '', '', ''),
                                                                                  (843, 'p', 'povertyAdmin', '/manager/api/v1/role/status', 'PUT', '', '', ''),
                                                                                  (850, 'p', 'povertyAdmin', '/manager/api/v1/user', 'DELETE', '', '', ''),
                                                                                  (848, 'p', 'povertyAdmin', '/manager/api/v1/user', 'POST', '', '', ''),
                                                                                  (849, 'p', 'povertyAdmin', '/manager/api/v1/user', 'PUT', '', '', ''),
                                                                                  (852, 'p', 'povertyAdmin', '/manager/api/v1/user/password/reset', 'POST', '', '', ''),
                                                                                  (851, 'p', 'povertyAdmin', '/manager/api/v1/user/status', 'PUT', '', '', ''),
                                                                                  (847, 'p', 'povertyAdmin', '/manager/api/v1/users', 'GET', '', '', ''),
                                                                                  (964, 'p', 'povertyAdmin', '/poverty/api/v1/activity', 'DELETE', '', '', ''),
                                                                                  (961, 'p', 'povertyAdmin', '/poverty/api/v1/activity', 'GET', '', '', ''),
                                                                                  (962, 'p', 'povertyAdmin', '/poverty/api/v1/activity', 'POST', '', '', ''),
                                                                                  (963, 'p', 'povertyAdmin', '/poverty/api/v1/activity', 'PUT', '', '', ''),
                                                                                  (960, 'p', 'povertyAdmin', '/poverty/api/v1/activitys', 'GET', '', '', ''),
                                                                                  (949, 'p', 'povertyAdmin', '/poverty/api/v1/banner', 'DELETE', '', '', ''),
                                                                                  (947, 'p', 'povertyAdmin', '/poverty/api/v1/banner', 'POST', '', '', ''),
                                                                                  (948, 'p', 'povertyAdmin', '/poverty/api/v1/banner', 'PUT', '', '', ''),
                                                                                  (946, 'p', 'povertyAdmin', '/poverty/api/v1/banners', 'GET', '', '', ''),
                                                                                  (965, 'p', 'povertyAdmin', '/poverty/api/v1/chat/records', 'GET', '', '', ''),
                                                                                  (959, 'p', 'povertyAdmin', '/poverty/api/v1/information', 'DELETE', '', '', ''),
                                                                                  (956, 'p', 'povertyAdmin', '/poverty/api/v1/information', 'GET', '', '', ''),
                                                                                  (957, 'p', 'povertyAdmin', '/poverty/api/v1/information', 'POST', '', '', ''),
                                                                                  (958, 'p', 'povertyAdmin', '/poverty/api/v1/information', 'PUT', '', '', ''),
                                                                                  (955, 'p', 'povertyAdmin', '/poverty/api/v1/informations', 'GET', '', '', ''),
                                                                                  (954, 'p', 'povertyAdmin', '/poverty/api/v1/notice', 'DELETE', '', '', ''),
                                                                                  (950, 'p', 'povertyAdmin', '/poverty/api/v1/notice', 'GET', '', '', ''),
                                                                                  (952, 'p', 'povertyAdmin', '/poverty/api/v1/notice', 'POST', '', '', ''),
                                                                                  (953, 'p', 'povertyAdmin', '/poverty/api/v1/notice', 'PUT', '', '', ''),
                                                                                  (951, 'p', 'povertyAdmin', '/poverty/api/v1/notices', 'GET', '', '', ''),
                                                                                  (969, 'p', 'povertyAdmin', '/poverty/api/v1/resource', 'DELETE', '', '', ''),
                                                                                  (967, 'p', 'povertyAdmin', '/poverty/api/v1/resource', 'POST', '', '', ''),
                                                                                  (968, 'p', 'povertyAdmin', '/poverty/api/v1/resource', 'PUT', '', '', ''),
                                                                                  (966, 'p', 'povertyAdmin', '/poverty/api/v1/resources', 'GET', '', '', ''),
                                                                                  (853, 'p', 'povertyAdmin', '/resource/api/v1/directories', 'GET', '', '', ''),
                                                                                  (856, 'p', 'povertyAdmin', '/resource/api/v1/directory', 'DELETE', '', '', ''),
                                                                                  (854, 'p', 'povertyAdmin', '/resource/api/v1/directory', 'POST', '', '', ''),
                                                                                  (855, 'p', 'povertyAdmin', '/resource/api/v1/directory', 'PUT', '', '', ''),
                                                                                  (862, 'p', 'povertyAdmin', '/resource/api/v1/export', 'DELETE', '', '', ''),
                                                                                  (861, 'p', 'povertyAdmin', '/resource/api/v1/export', 'POST', '', '', ''),
                                                                                  (860, 'p', 'povertyAdmin', '/resource/api/v1/exports', 'GET', '', '', ''),
                                                                                  (859, 'p', 'povertyAdmin', '/resource/api/v1/file', 'DELETE', '', '', ''),
                                                                                  (858, 'p', 'povertyAdmin', '/resource/api/v1/file', 'PUT', '', '', ''),
                                                                                  (857, 'p', 'povertyAdmin', '/resource/api/v1/files', 'GET', '', '', ''),
                                                                                  (864, 'p', 'povertyAdmin', '/usercenter/api/v1/app', 'PUT', '', '', ''),
                                                                                  (863, 'p', 'povertyAdmin', '/usercenter/api/v1/apps', 'GET', '', '', ''),
                                                                                  (882, 'p', 'povertyAdmin', '/usercenter/api/v1/auth', 'DELETE', '', '', ''),
                                                                                  (880, 'p', 'povertyAdmin', '/usercenter/api/v1/auth', 'POST', '', '', ''),
                                                                                  (881, 'p', 'povertyAdmin', '/usercenter/api/v1/auth/status', 'PUT', '', '', ''),
                                                                                  (879, 'p', 'povertyAdmin', '/usercenter/api/v1/auths', 'GET', '', '', ''),
                                                                                  (872, 'p', 'povertyAdmin', '/usercenter/api/v1/feedback', 'DELETE', '', '', ''),
                                                                                  (870, 'p', 'povertyAdmin', '/usercenter/api/v1/feedback', 'POST', '', '', ''),
                                                                                  (871, 'p', 'povertyAdmin', '/usercenter/api/v1/feedback', 'PUT', '', '', ''),
                                                                                  (869, 'p', 'povertyAdmin', '/usercenter/api/v1/feedbacks', 'GET', '', '', ''),
                                                                                  (865, 'p', 'povertyAdmin', '/usercenter/api/v1/feedback_categories', 'GET', '', '', ''),
                                                                                  (868, 'p', 'povertyAdmin', '/usercenter/api/v1/feedback_category', 'DELETE', '', '', ''),
                                                                                  (866, 'p', 'povertyAdmin', '/usercenter/api/v1/feedback_category', 'POST', '', '', ''),
                                                                                  (867, 'p', 'povertyAdmin', '/usercenter/api/v1/feedback_category', 'PUT', '', '', ''),
                                                                                  (884, 'p', 'povertyAdmin', '/usercenter/api/v1/oauth', 'DELETE', '', '', ''),
                                                                                  (883, 'p', 'povertyAdmin', '/usercenter/api/v1/oauths', 'GET', '', '', ''),
                                                                                  (878, 'p', 'povertyAdmin', '/usercenter/api/v1/user', 'DELETE', '', '', ''),
                                                                                  (874, 'p', 'povertyAdmin', '/usercenter/api/v1/user', 'POST', '', '', ''),
                                                                                  (876, 'p', 'povertyAdmin', '/usercenter/api/v1/user', 'PUT', '', '', ''),
                                                                                  (877, 'p', 'povertyAdmin', '/usercenter/api/v1/user/status', 'PUT', '', '', ''),
                                                                                  (888, 'p', 'povertyAdmin', '/usercenter/api/v1/userinfo', 'DELETE', '', '', ''),
                                                                                  (886, 'p', 'povertyAdmin', '/usercenter/api/v1/userinfo', 'POST', '', '', ''),
                                                                                  (887, 'p', 'povertyAdmin', '/usercenter/api/v1/userinfo', 'PUT', '', '', ''),
                                                                                  (885, 'p', 'povertyAdmin', '/usercenter/api/v1/userinfos', 'GET', '', '', ''),
                                                                                  (873, 'p', 'povertyAdmin', '/usercenter/api/v1/users', 'GET', '', '', ''),
                                                                                  (875, 'p', 'povertyAdmin', '/usercenter/api/v1/users', 'POST', '', '', '');

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
                                                                                                               (5, 1, 'gzy_school', '子部门', '子部门', 1720877602, 1720877602);

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

--
-- 转存表中的数据 `dictionary`
--

INSERT INTO `dictionary` (`id`, `keyword`, `name`, `description`, `created_at`, `updated_at`, `deleted_at`) VALUES
    (3, 'keyword', '关键词', '关键词', 1728103179, 1728103179, 0);

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
    (1, 'admin', '管理员', 0, '管理员', 1713706137, 1720877675, 0);

-- --------------------------------------------------------

--
-- 表的结构 `login_log`
--

CREATE TABLE `login_log` (
                             `id` bigint(20) UNSIGNED NOT NULL COMMENT '主键ID',
                             `username` varchar(128) NOT NULL COMMENT '账号',
                             `type` char(32) NOT NULL COMMENT '类型',
                             `ip` char(64) NOT NULL COMMENT 'ip',
                             `address` varchar(128) NOT NULL COMMENT '地址',
                             `device` varchar(128) NOT NULL COMMENT '设备',
                             `browser` varchar(128) NOT NULL COMMENT '浏览器',
                             `code` int(11) NOT NULL COMMENT '状态码',
                             `description` varchar(128) NOT NULL COMMENT '结果',
                             `created_at` bigint(20) UNSIGNED NOT NULL DEFAULT '0' COMMENT '创建时间'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='登录日志';

--
-- 转存表中的数据 `login_log`
--

INSERT INTO `login_log` (`id`, `username`, `type`, `ip`, `address`, `device`, `browser`, `code`, `description`, `created_at`) VALUES
                                                                                                                                  (1, '1280291001@qq.com', 'email', '1.204.170.140', '贵州省遵义市 电信', 'Windows 10.0', 'Edge 129.0.0.0', 200, '登陆成功', 1728106620),
                                                                                                                                  (2, '1280291001@qq.com', 'email', '117.189.213.192', '贵州省 移动', 'macOS 10.15.7', 'Chrome 112.0.0.0', 200, '登陆成功', 1728110502),
                                                                                                                                  (3, '1280291001@qq.com', 'email', '1.204.170.140', '贵州省遵义市 电信', 'Windows 10.0', 'Edge 129.0.0.0', 200, '登陆成功', 1728188469),
                                                                                                                                  (4, '1280291001@qq.com', 'email', '124.203.131.221', '贵州省遵义市 长城宽带', 'macOS 10.15.7', 'Chrome 112.0.0.0', 200, '登陆成功', 1728221083),
                                                                                                                                  (5, '18302586522', 'phone', '39.144.230.37', '贵州省贵阳市 移动', 'Android 14', 'Chrome 126.0.6478.188', 500, 'error: code = 500 reason = PasswordError message = 密码错误 metadata = map[] cause = <nil>', 1728479005),
                                                                                                                                  (6, '18302586522', 'phone', '39.144.230.37', '贵州省贵阳市 移动', 'Android 14', 'Chrome 126.0.6478.188', 200, '登陆成功', 1728479026),
                                                                                                                                  (7, '18302586522', 'phone', '117.189.26.117', '贵州省贵阳市 移动', 'Android 14', 'Chrome 126.0.6478.188', 200, '登陆成功', 1728479109),
                                                                                                                                  (8, '18302586522', 'phone', '117.189.26.117', '贵州省贵阳市 移动', 'Android 14', 'Chrome 126.0.6478.188', 200, '登陆成功', 1728479138),
                                                                                                                                  (9, '1280291001@qq.com', 'email', '220.181.3.151', '北京市 电信', 'macOS 10.15.7', 'Chrome 112.0.0.0', 200, '登陆成功', 1728479279),
                                                                                                                                  (10, '1280291001@qq.com', 'email', '117.188.5.57', '贵州省贵阳市 移通', 'Windows 10.0', 'Edge 129.0.0.0', 200, '登陆成功', 1728486586),
                                                                                                                                  (11, '1280291001@qq.com', 'email', '111.201.4.162', '北京市大兴区 联通', 'macOS 10.15.7', 'Chrome 112.0.0.0', 200, '登陆成功', 1728487276),
                                                                                                                                  (12, '1280291001@qq.com', 'email', '111.201.4.162', '北京市大兴区 联通', 'macOS 10.15.7', 'Chrome 112.0.0.0', 200, '登陆成功', 1728496602),
                                                                                                                                  (13, '1280291001@qq.com', 'email', '39.144.231.82', ' 中国移动', 'Windows 10.0', 'Edge 129.0.0.0', 200, '登陆成功', 1728526525),
                                                                                                                                  (14, '2371703004@qq.com', 'email', '39.144.231.82', ' 中国移动', 'Linux x86_64', 'XiaoMi MiuiBrowser/18.6.60929', 200, '登陆成功', 1728530309),
                                                                                                                                  (15, '1280291001@qq.com', 'email', '117.189.21.181', '贵州省贵阳市 移动', 'Windows 10.0', 'Edge 129.0.0.0', 200, '登陆成功', 1728553625),
                                                                                                                                  (16, '1280291001@qq.com', 'email', '220.181.3.151', '北京市 电信', 'macOS 10.15.7', 'Chrome 112.0.0.0', 200, '登陆成功', 1728560319),
                                                                                                                                  (17, '1280291001@qq.com', 'email', '117.189.21.181', '贵州省贵阳市 移动', 'Windows 10.0', 'Edge 129.0.0.0', 200, '登陆成功', 1728569885),
                                                                                                                                  (18, '1280291001@qq.com', 'email', '220.181.3.151', '北京市 电信', 'macOS 10.15.7', 'Chrome 112.0.0.0', 200, '登陆成功', 1728615573),
                                                                                                                                  (19, '1280291001@qq.com', 'email', '111.201.4.162', '北京市大兴区 联通', 'macOS 10.15.7', 'Chrome 112.0.0.0', 200, '登陆成功', 1728653350),
                                                                                                                                  (20, '1280291001@qq.com', 'email', '117.189.21.124', '贵州省贵阳市 移动', 'Windows 10.0', 'Edge 129.0.0.0', 200, '登陆成功', 1728721021),
                                                                                                                                  (21, '1280291001@qq.com', 'email', '117.188.5.179', '贵州省贵阳市 移通', 'Windows 10.0', 'Edge 129.0.0.0', 200, '登陆成功', 1728825231),
                                                                                                                                  (22, '1280291001@qq.com', 'email', '117.189.21.248', '贵州省贵阳市 移动', 'Windows 10.0', 'Edge 129.0.0.0', 200, '登陆成功', 1728918165),
                                                                                                                                  (23, '1280291001@qq.com', 'email', '111.201.4.162', '北京市大兴区 联通', 'macOS 10.15.7', 'Chrome 112.0.0.0', 200, '登陆成功', 1728997242),
                                                                                                                                  (24, '1280291001@qq.com', 'email', '117.189.26.64', '贵州省贵阳市 移动', 'Windows 10.0', 'Edge 129.0.0.0', 200, '登陆成功', 1729086052),
                                                                                                                                  (25, '1280291001@qq.com', 'email', '111.201.4.162', '北京市大兴区 联通', 'macOS 10.15.7', 'Chrome 112.0.0.0', 200, '登陆成功', 1729086604),
                                                                                                                                  (26, '1280291001@qq.com', 'email', '39.144.230.24', '贵州省贵阳市 移动', 'Windows 10.0', 'Edge 129.0.0.0', 200, '登陆成功', 1729131596),
                                                                                                                                  (27, '1280291001@qq.com', 'email', '220.181.3.151', '北京市 电信', 'macOS 10.15.7', 'Chrome 112.0.0.0', 200, '登陆成功', 1729133483),
                                                                                                                                  (28, '1280291001@qq.com', 'email', '111.201.4.162', '北京市大兴区 联通', 'macOS 10.15.7', 'Chrome 112.0.0.0', 200, '登陆成功', 1729183805),
                                                                                                                                  (29, '1280291001@qq.com', 'email', '220.181.3.151', '北京市 电信', 'macOS 10.15.7', 'Chrome 112.0.0.0', 200, '登陆成功', 1729250550),
                                                                                                                                  (30, '1280291001@qq.com', 'email', '111.201.4.162', '北京市大兴区 联通', 'macOS 10.15.7', 'Chrome 112.0.0.0', 200, '登陆成功', 1729417912),
                                                                                                                                  (31, '1280291001@qq.com', 'email', '117.188.5.26', '贵州省贵阳市 移通', 'Windows 10.0', 'Edge 130.0.0.0', 200, '登陆成功', 1729665117),
                                                                                                                                  (32, '1280291001@qq.com', 'email', '220.181.3.151', '北京市 电信', 'macOS 10.15.7', 'Chrome 130.0.0.0', 200, '登陆成功', 1729669031);

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
                                                                                                                                                                                                                                     (3446, 0, '管理平台', 'R', 'SystemPlatform', 'desktop', NULL, NULL, '/', NULL, 'Layout', NULL, 2, 0, NULL, NULL, NULL, 1728056544, 1728056544),
                                                                                                                                                                                                                                     (3447, 3446, '首页面板', 'M', 'Dashboard', 'dashboard', NULL, NULL, '/dashboard', NULL, '/dashboard/index', NULL, 0, 0, 1, 1, 1, 1728056544, 1728056544),
                                                                                                                                                                                                                                     (3448, 3446, '管理中心', 'M', 'SystemPlatformManager', 'desktop', NULL, NULL, '/manager', NULL, NULL, NULL, 0, 0, NULL, NULL, NULL, 1728056544, 1728056544),
                                                                                                                                                                                                                                     (3449, 3448, '基础接口', 'G', 'ManagerBaseApi', 'apps', NULL, NULL, NULL, NULL, NULL, NULL, 0, 1, NULL, NULL, NULL, 1728056544, 1728056544),
                                                                                                                                                                                                                                     (3450, 3449, '获取用户可见部门树', 'BA', NULL, NULL, '/manager/api/v1/departments', 'GET', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL, 1728056544, 1728056544),
                                                                                                                                                                                                                                     (3451, 3449, '获取用户可见角色树', 'BA', NULL, NULL, '/manager/api/v1/roles', 'GET', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL, 1728056544, 1728056544),
                                                                                                                                                                                                                                     (3452, 3449, '获取个人用户信息', 'BA', NULL, NULL, '/manager/api/v1/user/current', 'GET', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL, 1728056544, 1728056544),
                                                                                                                                                                                                                                     (3453, 3449, '更新个人用户信息', 'BA', NULL, NULL, '/manager/api/v1/user/current/info', 'PUT', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL, 1728056544, 1728056544),
                                                                                                                                                                                                                                     (3454, 3449, '更新个人用户密码', 'BA', NULL, NULL, '/manager/api/v1/user/current/password', 'PUT', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL, 1728056544, 1728056544),
                                                                                                                                                                                                                                     (3455, 3449, '保存个人用户设置', 'BA', NULL, NULL, '/manager/api/v1/user/current/setting', 'PUT', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL, 1728056544, 1728056544),
                                                                                                                                                                                                                                     (3456, 3449, '发送用户验证吗', 'BA', NULL, NULL, '/manager/api/v1/user/current/captcha', 'POST', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL, 1728056544, 1728056544),
                                                                                                                                                                                                                                     (3457, 3449, '获取用户当前角色菜单', 'BA', NULL, NULL, '/manager/api/v1/menus/by/cur_role', 'GET', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL, 1728056544, 1728056544),
                                                                                                                                                                                                                                     (3458, 3449, '退出登录', 'BA', NULL, NULL, '/manager/api/v1/user/logout', 'POST', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL, 1728056544, 1728056544),
                                                                                                                                                                                                                                     (3459, 3449, '刷新token', 'BA', NULL, NULL, '/manager/api/v1/user/token/refresh', 'POST', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL, 1728056544, 1728056544),
                                                                                                                                                                                                                                     (3460, 3449, '用户登录', 'BA', NULL, NULL, '/manager/api/v1/user/login', 'POST', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL, 1728056544, 1728056544),
                                                                                                                                                                                                                                     (3461, 3449, '获取登录验证码', 'BA', NULL, NULL, '/manager/api/v1/user/login/captcha', 'POST', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL, 1728056544, 1728056544),
                                                                                                                                                                                                                                     (3462, 3449, '获取系统基础设置', 'BA', NULL, NULL, '/manager/api/v1/system/setting', 'GET', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL, 1728056544, 1728056544),
                                                                                                                                                                                                                                     (3463, 3449, '接口鉴权', 'BA', NULL, NULL, '/manager/api/v1/auth', 'POST', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL, 1728056544, 1728056544),
                                                                                                                                                                                                                                     (3464, 3449, '切换用户角色', 'BA', NULL, NULL, '/manager/api/v1/user/current/role', 'POST', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL, 1728056544, 1728056544),
                                                                                                                                                                                                                                     (3465, 3449, '分块上传文件', 'BA', NULL, NULL, '/resource/api/v1/upload', 'POST', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL, 1728056544, 1728056544),
                                                                                                                                                                                                                                     (3466, 3449, '预上传文件', 'BA', NULL, NULL, '/resource/api/v1/prepare_upload', 'POST', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL, 1728056544, 1728056544),
                                                                                                                                                                                                                                     (3467, 3449, '获取字段类型', 'BA', NULL, NULL, '/application/api/v1/field/types', 'GET', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL, 1728056544, 1728056544),
                                                                                                                                                                                                                                     (3468, 3449, '查询资源权限', 'BA', NULL, NULL, '/manager/api/v1/resource', 'GET', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL, 1728056544, 1728056544),
                                                                                                                                                                                                                                     (3469, 3449, '获取渠道类型', 'BA', NULL, NULL, '/application/api/v1/channel/types', 'GET', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL, 1728056544, 1728056544),
                                                                                                                                                                                                                                     (3470, 3448, '字典管理', 'M', 'ManagerDictionary', 'storage', NULL, NULL, '/manager/dictionary', NULL, '/manager/dictionary/index', NULL, 0, NULL, 1, NULL, 1, 1728056544, 1728056544),
                                                                                                                                                                                                                                     (3471, 3470, '查询字典', 'A', NULL, NULL, '/manager/api/v1/dictionaries', 'GET', NULL, 'manager:dictionary:query', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1728056544, 1728056544),
                                                                                                                                                                                                                                     (3472, 3470, '新增字典', 'A', NULL, NULL, '/manager/api/v1/dictionary', 'POST', NULL, 'manager:dictionary:add', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1728056544, 1728056544),
                                                                                                                                                                                                                                     (3473, 3470, '修改字典', 'A', NULL, NULL, '/manager/api/v1/dictionary', 'PUT', NULL, 'manager:dictionary:update', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1728056544, 1728056544),
                                                                                                                                                                                                                                     (3474, 3470, '删除字典', 'A', NULL, NULL, '/manager/api/v1/dictionary', 'DELETE', NULL, 'manager:dictionary:delete', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1728056544, 1728056544),
                                                                                                                                                                                                                                     (3475, 3470, '获取字典值', 'A', NULL, NULL, '/manager/api/v1/dictionary_values', 'GET', NULL, 'manager:dictionary:value:query', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1728056544, 1728056544),
                                                                                                                                                                                                                                     (3476, 3470, '新增字典值', 'A', NULL, NULL, '/manager/api/v1/dictionary_value', 'POST', NULL, 'manager:dictionary:value:add', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1728056544, 1728056544),
                                                                                                                                                                                                                                     (3477, 3470, '修改字典值', 'A', NULL, NULL, '/manager/api/v1/dictionary_value', 'PUT', NULL, 'manager:dictionary:value:update', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1728056544, 1728056544),
                                                                                                                                                                                                                                     (3478, 3470, '更新字典值目录状态', 'A', NULL, NULL, '/manager/api/v1/dictionary_value/status', 'PUT', NULL, 'manager:dictionary:value:status', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1728056544, 1728056544),
                                                                                                                                                                                                                                     (3479, 3470, '删除字典值', 'A', NULL, NULL, '/manager/api/v1/dictionary_value', 'DELETE', NULL, 'manager:dictionary:value:delete', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1728056544, 1728056544),
                                                                                                                                                                                                                                     (3480, 3448, '菜单管理', 'M', 'ManagerMenu', 'menu', NULL, NULL, '/manager/menu', NULL, '/manager/menu/index', NULL, 0, 0, 1, 1, 1, 1728056544, 1728056544),
                                                                                                                                                                                                                                     (3481, 3480, '查询菜单', 'A', NULL, NULL, '/manager/api/v1/menus', 'GET', NULL, 'manager:menu:query', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1728056544, 1728056544),
                                                                                                                                                                                                                                     (3482, 3480, '新增菜单', 'A', NULL, NULL, '/manager/api/v1/menu', 'POST', NULL, 'manager:menu:add', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1728056544, 1728056544),
                                                                                                                                                                                                                                     (3483, 3480, '修改菜单', 'A', NULL, NULL, '/manager/api/v1/menu', 'PUT', NULL, 'manager:menu:update', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1728056544, 1728056544),
                                                                                                                                                                                                                                     (3484, 3480, '删除菜单', 'A', NULL, NULL, '/manager/api/v1/menu', 'DELETE', NULL, 'manager:menu:delete', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1728056544, 1728056544),
                                                                                                                                                                                                                                     (3485, 3448, '职位管理', 'M', 'ManagerJob', 'tag', NULL, NULL, '/manager/job', NULL, '/manager/job/index', NULL, 0, NULL, 1, NULL, 1, 1728056544, 1728056544),
                                                                                                                                                                                                                                     (3486, 3485, '查询职位', 'A', NULL, NULL, '/manager/api/v1/jobs', 'GET', NULL, 'manager:job:query', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1728056544, 1728056544),
                                                                                                                                                                                                                                     (3487, 3485, '新增职位', 'A', NULL, NULL, '/manager/api/v1/job', 'POST', NULL, 'manager:job:add', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1728056544, 1728056544),
                                                                                                                                                                                                                                     (3488, 3485, '修改职位', 'A', NULL, NULL, '/manager/api/v1/job', 'PUT', NULL, 'manager:job:update', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1728056544, 1728056544),
                                                                                                                                                                                                                                     (3489, 3485, '删除职位', 'A', NULL, NULL, '/manager/api/v1/job', 'DELETE', NULL, 'manager:job:delete', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1728056544, 1728056544),
                                                                                                                                                                                                                                     (3490, 3448, '部门管理', 'M', 'ManagerDepartment', 'user-group', NULL, NULL, '/manager/department', NULL, '/manager/department/index', NULL, 0, 0, 1, NULL, 1, 1728056544, 1728056544),
                                                                                                                                                                                                                                     (3491, 3490, '新增部门', 'A', NULL, NULL, '/manager/api/v1/department', 'POST', NULL, 'manager:department:add', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1728056544, 1728056544),
                                                                                                                                                                                                                                     (3492, 3490, '修改部门', 'A', NULL, NULL, '/manager/api/v1/department', 'PUT', NULL, 'manager:department:update', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1728056544, 1728056544),
                                                                                                                                                                                                                                     (3493, 3490, '删除部门', 'A', NULL, NULL, '/manager/api/v1/department', 'DELETE', NULL, 'manager:department:delete', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1728056544, 1728056544),
                                                                                                                                                                                                                                     (3494, 3448, '角色管理', 'M', 'ManagerRole', 'safe', NULL, NULL, '/manager/role', NULL, '/manager/role/index', NULL, 0, NULL, 0, NULL, 1, 1728056544, 1728056544),
                                                                                                                                                                                                                                     (3495, 3494, '新增角色', 'A', NULL, NULL, '/manager/api/v1/role', 'POST', NULL, 'manager:role:add', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1728056544, 1728056544),
                                                                                                                                                                                                                                     (3496, 3494, '修改角色', 'A', NULL, NULL, '/manager/api/v1/role', 'PUT', NULL, 'manager:role:update', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1728056544, 1728056544),
                                                                                                                                                                                                                                     (3497, 3494, '修改角色状态', 'A', NULL, NULL, '/manager/api/v1/role/status', 'PUT', NULL, 'manager:role:update:status', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1728056544, 1728056544),
                                                                                                                                                                                                                                     (3498, 3494, '删除角色', 'A', NULL, NULL, '/manager/api/v1/role', 'DELETE', NULL, 'manager:role:delete', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1728056544, 1728056544),
                                                                                                                                                                                                                                     (3499, 3494, '角色菜单管理', 'G', NULL, NULL, NULL, NULL, NULL, 'manager:role:menu', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1728056544, 1728056544),
                                                                                                                                                                                                                                     (3500, 3499, '查询角色菜单', 'A', NULL, NULL, '/manager/api/v1/role/menu_ids', 'GET', NULL, 'manager:role:menu:query', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1728056544, 1728056544),
                                                                                                                                                                                                                                     (3501, 3499, '修改角色菜单', 'A', NULL, NULL, '/manager/api/v1/role/menu', 'POST', NULL, 'manager:role:menu:update', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1728056544, 1728056544),
                                                                                                                                                                                                                                     (3502, 3448, '用户管理', 'M', 'ManagerUser', 'user', NULL, NULL, '/manager/user', NULL, '/manager/user/index', NULL, 0, NULL, 1, NULL, 1, 1728056544, 1728056544),
                                                                                                                                                                                                                                     (3503, 3502, '查询用户列表', 'A', NULL, NULL, '/manager/api/v1/users', 'GET', NULL, 'manager:user:query', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1728056544, 1728056544),
                                                                                                                                                                                                                                     (3504, 3502, '新增用户', 'A', NULL, NULL, '/manager/api/v1/user', 'POST', NULL, 'manager:user:add', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1728056544, 1728056544),
                                                                                                                                                                                                                                     (3505, 3502, '修改用户', 'A', NULL, NULL, '/manager/api/v1/user', 'PUT', NULL, 'manager:user:update', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1728056544, 1728056544),
                                                                                                                                                                                                                                     (3506, 3502, '删除用户', 'A', NULL, NULL, '/manager/api/v1/user', 'DELETE', NULL, 'manager:user:delete', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1728056544, 1728056544),
                                                                                                                                                                                                                                     (3507, 3502, '修改用户状态', 'A', NULL, NULL, '/manager/api/v1/user/status', 'PUT', NULL, 'manager:user:status', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1728056544, 1728056544),
                                                                                                                                                                                                                                     (3508, 3502, '重置账号密码', 'A', NULL, NULL, '/manager/api/v1/user/password/reset', 'POST', NULL, 'manager:user:reset:password', '', NULL, 0, NULL, NULL, NULL, NULL, 1728056544, 1728056544),
                                                                                                                                                                                                                                     (3509, 3448, '登陆日志', 'M', 'ManagerLoginLog', 'history', NULL, NULL, '/manager/log', NULL, '/manager/log/index', NULL, 0, NULL, 1, NULL, 1, 1728056544, 1728056544),
                                                                                                                                                                                                                                     (3510, 3509, '查询登陆列表', 'A', NULL, NULL, '/manager/api/v1/user/login/logs', 'GET', NULL, 'manager:login:log:query', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1728056544, 1728056544),
                                                                                                                                                                                                                                     (3511, 3446, '资源中心', 'M', 'SystemPlatformResource', 'file', NULL, NULL, '/resource', NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL, 1728056544, 1728056544),
                                                                                                                                                                                                                                     (3512, 3511, '文件管理', 'M', 'ResourceFile', 'file', NULL, NULL, '/resource/file', NULL, '/resource/file/index', NULL, 0, 0, 1, NULL, 1, 1728056544, 1728056544),
                                                                                                                                                                                                                                     (3513, 3512, '目录管理', 'G', NULL, NULL, NULL, NULL, NULL, 'resource:directory:group', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1728056544, 1728056544),
                                                                                                                                                                                                                                     (3514, 3513, '查看目录', 'A', NULL, NULL, '/resource/api/v1/directories', 'GET', NULL, 'resource:directory:query', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1728056544, 1728056544),
                                                                                                                                                                                                                                     (3515, 3513, '新增目录', 'A', NULL, NULL, '/resource/api/v1/directory', 'POST', NULL, 'resource:directory:add', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1728056544, 1728056544),
                                                                                                                                                                                                                                     (3516, 3513, '修改目录', 'A', NULL, NULL, '/resource/api/v1/directory', 'PUT', NULL, 'resource:directory:update', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1728056544, 1728056544),
                                                                                                                                                                                                                                     (3517, 3513, '删除目录', 'A', NULL, NULL, '/resource/api/v1/directory', 'DELETE', NULL, 'resource:directory:delete', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1728056544, 1728056544),
                                                                                                                                                                                                                                     (3518, 3512, '文件管理', 'G', NULL, NULL, NULL, NULL, NULL, 'resource:file:group', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1728056544, 1728056544),
                                                                                                                                                                                                                                     (3519, 3518, '查看文件', 'A', NULL, NULL, '/resource/api/v1/files', 'GET', NULL, 'resource:file:query', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1728056544, 1728056544),
                                                                                                                                                                                                                                     (3520, 3518, '修改文件', 'A', NULL, NULL, '/resource/api/v1/file', 'PUT', NULL, 'resource:file:update', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1728056544, 1728056544),
                                                                                                                                                                                                                                     (3521, 3518, '删除文件', 'A', NULL, NULL, '/resource/api/v1/file', 'DELETE', NULL, 'resource:file:delete', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1728056544, 1728056544),
                                                                                                                                                                                                                                     (3522, 3511, '导出管理', 'M', 'ResourceExport', 'export', NULL, NULL, '/resource/export', NULL, '/resource/export/index', NULL, 0, 0, 1, NULL, 1, 1728056544, 1728056544),
                                                                                                                                                                                                                                     (3523, 3522, '查看导出', 'A', NULL, NULL, '/resource/api/v1/exports', 'GET', NULL, 'resource:export:query', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1728056544, 1728056544),
                                                                                                                                                                                                                                     (3524, 3522, '新增导出', 'A', NULL, NULL, '/resource/api/v1/export', 'POST', NULL, 'resource:export:file', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1728056544, 1728056544),
                                                                                                                                                                                                                                     (3525, 3522, '删除导出', 'A', NULL, NULL, '/resource/api/v1/export', 'DELETE', NULL, 'resource:export:delete', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1728056544, 1728056544),
                                                                                                                                                                                                                                     (3526, 3446, '应用中心', 'M', 'SystemPlatformApplication', 'apps', NULL, NULL, '/application', NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL, 1728056544, 1728056544),
                                                                                                                                                                                                                                     (3527, 3526, '授权渠道', 'M', 'ApplicationChannel', 'mind-mapping', NULL, NULL, '/application/channel', NULL, '/application/channel/index', NULL, 0, 0, 1, NULL, 1, 1728056544, 1728056544),
                                                                                                                                                                                                                                     (3528, 3527, '查看渠道', 'A', NULL, NULL, '/application/api/v1/channels', 'GET', NULL, 'application:channel:query', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1728056544, 1728056544),
                                                                                                                                                                                                                                     (3529, 3527, '新增渠道', 'A', NULL, NULL, '/application/api/v1/channel', 'POST', NULL, 'application:channel:add', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1728056544, 1728056544),
                                                                                                                                                                                                                                     (3530, 3527, '修改渠道', 'A', NULL, NULL, '/application/api/v1/channel', 'PUT', NULL, 'application:channel:update', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1728056544, 1728056544),
                                                                                                                                                                                                                                     (3531, 3527, '删除渠道', 'A', NULL, NULL, '/application/api/v1/channel', 'DELETE', NULL, 'application:channel:delete', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1728056544, 1728056544),
                                                                                                                                                                                                                                     (3532, 3526, '信息字段', 'M', 'ApplicationField', 'list', NULL, NULL, '/application/field', NULL, '/application/field/index', NULL, 0, 0, 1, 0, 1, 1728056544, 1728056544),
                                                                                                                                                                                                                                     (3533, 3532, '查看字段列表', 'A', NULL, NULL, '/application/api/v1/fields', 'GET', NULL, 'application:field:query', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1728056544, 1728056544),
                                                                                                                                                                                                                                     (3534, 3532, '新增字段', 'A', NULL, NULL, '/application/api/v1/field', 'POST', NULL, 'application:field:add', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1728056544, 1728056544),
                                                                                                                                                                                                                                     (3535, 3532, '修改字段', 'A', NULL, NULL, '/application/api/v1/field', 'PUT', NULL, 'application:field:update', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1728056544, 1728056544),
                                                                                                                                                                                                                                     (3536, 3532, '修改字段状态', 'A', NULL, NULL, '/application/api/v1/field/status', 'PUT', NULL, 'application:field:update:status', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1728056544, 1728056544),
                                                                                                                                                                                                                                     (3537, 3532, '删除字段', 'A', NULL, NULL, '/application/api/v1/field', 'DELETE', NULL, 'application:field:delete', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1728056544, 1728056544),
                                                                                                                                                                                                                                     (3538, 3526, '应用管理', 'M', 'ApplicationApp', 'apps', NULL, NULL, '/application/app', NULL, '/application/app/index', NULL, 0, 0, 1, NULL, 1, 1728056544, 1728056544),
                                                                                                                                                                                                                                     (3539, 3538, '查看应用', 'A', NULL, NULL, '/application/api/v1/apps', 'GET', NULL, 'application:app:query', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1728056544, 1728056544),
                                                                                                                                                                                                                                     (3540, 3538, '设置应用资源权限', 'A', NULL, NULL, '/manager/api/v1/resource/uc_app', 'PUT', NULL, 'application:app:resource:permission', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1728056544, 1728056544),
                                                                                                                                                                                                                                     (3541, 3538, '新增应用', 'A', NULL, NULL, '/application/api/v1/app', 'POST', NULL, 'application:app:add', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1728056544, 1728056544),
                                                                                                                                                                                                                                     (3542, 3538, '修改应用', 'A', NULL, NULL, '/application/api/v1/app', 'PUT', NULL, 'application:app:update', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1728056544, 1728056544),
                                                                                                                                                                                                                                     (3543, 3538, '修改应用状态', 'A', NULL, NULL, '/application/api/v1/app/status', 'PUT', NULL, 'application:app:update:status', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1728056544, 1728056544),
                                                                                                                                                                                                                                     (3544, 3538, '删除应用', 'A', NULL, NULL, '/application/api/v1/app', 'DELETE', NULL, 'application:app:delete', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1728056544, 1728056544),
                                                                                                                                                                                                                                     (3545, 3526, '反馈管理', 'M', 'ApplicationFeedback', 'question-circle', NULL, NULL, '/application/feedback', NULL, '/application/feedback/index', NULL, 0, 0, 1, NULL, 1, 1728056544, 1728056544),
                                                                                                                                                                                                                                     (3546, 3545, '查看反馈分类', 'A', NULL, NULL, '/application/api/v1/feedback_categories', 'GET', NULL, 'application:feedback:category:query', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1728056544, 1728056544),
                                                                                                                                                                                                                                     (3547, 3545, '新增反馈渠道', 'A', NULL, NULL, '/application/api/v1/feedback_category', 'POST', NULL, 'application:feedback:category:add', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1728056544, 1728056544),
                                                                                                                                                                                                                                     (3548, 3545, '修改反馈渠道', 'A', NULL, NULL, '/application/api/v1/feedback_category', 'PUT', NULL, 'application:feedback:category:update', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1728056544, 1728056544),
                                                                                                                                                                                                                                     (3549, 3545, '删除反馈渠道', 'A', NULL, NULL, '/application/api/v1/feedback_category', 'DELETE', NULL, 'application:feedback:category:delete', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1728056544, 1728056544),
                                                                                                                                                                                                                                     (3550, 3545, '查看反馈', 'A', NULL, NULL, '/application/api/v1/feedbacks', 'GET', NULL, 'application:feedback:query', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1728056544, 1728056544),
                                                                                                                                                                                                                                     (3551, 3545, '新增反馈', 'A', NULL, NULL, '/application/api/v1/feedback', 'POST', NULL, 'application:feedback:add', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1728056544, 1728056544),
                                                                                                                                                                                                                                     (3552, 3545, '修改反馈', 'A', NULL, NULL, '/application/api/v1/feedback', 'PUT', NULL, 'application:feedback:update', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1728056544, 1728056544),
                                                                                                                                                                                                                                     (3553, 3545, '删除反馈', 'A', NULL, NULL, '/application/api/v1/feedback', 'DELETE', NULL, 'application:feedback:delete', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1728056544, 1728056544),
                                                                                                                                                                                                                                     (3554, 3526, '用户管理', 'M', 'ApplicationUser', 'user', NULL, NULL, '/application/user', NULL, '/application/user/index', NULL, 0, 0, 1, NULL, 1, 1728056544, 1728056544),
                                                                                                                                                                                                                                     (3555, 3554, '查看用户', 'A', NULL, NULL, '/application/api/v1/users', 'GET', NULL, 'application:user:query', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1728056544, 1728056544),
                                                                                                                                                                                                                                     (3556, 3554, '新增用户', 'A', NULL, NULL, '/application/api/v1/user', 'POST', NULL, 'application:user:add', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1728056544, 1728056544),
                                                                                                                                                                                                                                     (3557, 3554, '导入用户', 'A', NULL, NULL, '/application/api/v1/users', 'POST', NULL, 'application:user:import', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1728056544, 1728056544),
                                                                                                                                                                                                                                     (3558, 3554, '修改用户', 'A', NULL, NULL, '/application/api/v1/user', 'PUT', NULL, 'application:user:update', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1728056544, 1728056544),
                                                                                                                                                                                                                                     (3559, 3554, '修改用户状态', 'A', NULL, NULL, '/application/api/v1/user/status', 'PUT', NULL, 'application:user:update:status', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1728056544, 1728056544),
                                                                                                                                                                                                                                     (3560, 3554, '删除用户', 'A', NULL, NULL, '/application/api/v1/user', 'DELETE', NULL, 'application:user:delete', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1728056544, 1728056544),
                                                                                                                                                                                                                                     (3561, 3554, '查看用户详细信息', 'G', NULL, NULL, NULL, NULL, NULL, 'application:user:more', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1728056544, 1728056544),
                                                                                                                                                                                                                                     (3562, 3561, '获取用户应用信息', 'A', NULL, NULL, '/application/api/v1/auths', 'GET', NULL, 'application:auth:query', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1728056544, 1728056544),
                                                                                                                                                                                                                                     (3563, 3561, '创建用户应用信息', 'A', NULL, NULL, '/application/api/v1/auth', 'POST', NULL, 'application:auth:add', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1728056544, 1728056544),
                                                                                                                                                                                                                                     (3564, 3561, '修改用户应用状态信息', 'A', NULL, NULL, '/application/api/v1/auth/status', 'PUT', NULL, 'application:auth:update:status', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1728056544, 1728056544),
                                                                                                                                                                                                                                     (3565, 3561, '删除用户应用信息', 'A', NULL, NULL, '/application/api/v1/auth', 'DELETE', NULL, 'application:auth:delete', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1728056544, 1728056544),
                                                                                                                                                                                                                                     (3566, 3561, '获取用户渠道信息', 'A', NULL, NULL, '/application/api/v1/oauths', 'GET', NULL, 'application:oauth:query', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1728056544, 1728056544),
                                                                                                                                                                                                                                     (3567, 3561, '删除用户渠道信息', 'A', NULL, NULL, '/application/api/v1/oauth', 'DELETE', NULL, 'application:oauth:delete', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1728056544, 1728056544),
                                                                                                                                                                                                                                     (3568, 3561, '获取用户扩展信息', 'A', NULL, NULL, '/application/api/v1/userinfos', 'GET', NULL, 'application:userinfo:query', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1728056544, 1728056544),
                                                                                                                                                                                                                                     (3569, 3561, '创建用户扩展信息', 'A', NULL, NULL, '/application/api/v1/userinfo', 'POST', NULL, 'application:userinfo:add', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1728056544, 1728056544),
                                                                                                                                                                                                                                     (3570, 3561, '修改用户扩展信息', 'A', NULL, NULL, '/application/api/v1/userinfo', 'PUT', NULL, 'application:userinfo:update', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1728056544, 1728056544),
                                                                                                                                                                                                                                     (3571, 3561, '删除用户扩展信息', 'A', NULL, NULL, '/application/api/v1/userinfo', 'DELETE', NULL, 'application:userinfo:delete', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1728056544, 1728056544),
                                                                                                                                                                                                                                     (3572, 3446, '配置中心', 'M', 'SystemPlatformConfigure', 'code-block', NULL, NULL, '/configure', NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL, 1728056544, 1728056544),
                                                                                                                                                                                                                                     (3573, 3572, '环境管理', 'M', 'ConfigureEnv', 'common', NULL, NULL, '/configure/env', NULL, '/configure/env/index', NULL, 0, NULL, 1, NULL, 1, 1728056544, 1728056544),
                                                                                                                                                                                                                                     (3574, 3573, '查看环境', 'A', NULL, NULL, '/configure/api/v1/envs', 'GET', NULL, 'configure:env:query', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1728056544, 1728056544),
                                                                                                                                                                                                                                     (3575, 3573, '设置环境资源权限', 'A', NULL, NULL, '/manager/api/v1/resource/cfg_env', 'PUT', NULL, 'configure:env:resource:permission', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1728056544, 1728056544),
                                                                                                                                                                                                                                     (3576, 3573, '新增环境', 'A', NULL, NULL, '/configure/api/v1/env', 'POST', NULL, 'configure:env:add', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1728056544, 1728056544),
                                                                                                                                                                                                                                     (3577, 3573, '修改环境', 'A', NULL, NULL, '/configure/api/v1/env', 'PUT', NULL, 'configure:env:update', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1728056544, 1728056544),
                                                                                                                                                                                                                                     (3578, 3573, '修改环境状态', 'A', NULL, NULL, '/configure/api/v1/env/status', 'PUT', NULL, 'configure:env:update:status', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1728056544, 1728056544),
                                                                                                                                                                                                                                     (3579, 3573, '删除环境', 'A', NULL, NULL, '/configure/api/v1/env', 'DELETE', NULL, 'configure:env:delete', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1728056544, 1728056544),
                                                                                                                                                                                                                                     (3580, 3573, '查看环境Token', 'A', NULL, NULL, '/configure/api/v1/env/token', 'GET', NULL, 'configure:env:token:query', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1728056544, 1728056544),
                                                                                                                                                                                                                                     (3581, 3573, '重置环境token', 'A', NULL, NULL, '/configure/api/v1/env/token', 'PUT', NULL, 'configure:env:token:reset', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1728056544, 1728056544),
                                                                                                                                                                                                                                     (3582, 3572, '服务管理', 'M', 'ConfigureServer', 'apps', NULL, NULL, '/configure/server', NULL, '/configure/server/index', NULL, 0, NULL, 1, NULL, 1, 1728056544, 1728056544),
                                                                                                                                                                                                                                     (3583, 3582, '查询服务', 'A', NULL, NULL, '/configure/api/v1/servers', 'GET', NULL, 'configure:server:query', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1728056544, 1728056544),
                                                                                                                                                                                                                                     (3584, 3582, '设置服务资源权限', 'A', NULL, NULL, '/manager/api/v1/resource/cfg_server', 'PUT', NULL, 'configure:server:resource:permission', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1728056544, 1728056544),
                                                                                                                                                                                                                                     (3585, 3582, '新增服务', 'A', NULL, NULL, '/configure/api/v1/server', 'POST', NULL, 'configure:server:add', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1728056544, 1728056544),
                                                                                                                                                                                                                                     (3586, 3582, '修改服务', 'A', NULL, NULL, '/configure/api/v1/server', 'PUT', NULL, 'configure:server:update', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1728056544, 1728056544),
                                                                                                                                                                                                                                     (3587, 3582, '修改服务状态', 'A', NULL, NULL, '/configure/api/v1/server/status', 'PUT', NULL, 'configure:server:update:status', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1728056544, 1728056544),
                                                                                                                                                                                                                                     (3588, 3582, '删除服务', 'A', NULL, NULL, '/configure/api/v1/server', 'DELETE', NULL, 'configure:server:delete', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1728056544, 1728056544),
                                                                                                                                                                                                                                     (3589, 3572, '业务变量', 'M', 'ConfigureBusiness', 'code', NULL, NULL, '/configure/business', NULL, '/configure/business/index', NULL, 0, NULL, 1, NULL, 1, 1728056544, 1728056544),
                                                                                                                                                                                                                                     (3590, 3589, '查看业务变量', 'A', NULL, NULL, '/configure/api/v1/businesses', 'GET', NULL, 'configure:business:query', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1728056544, 1728056544),
                                                                                                                                                                                                                                     (3591, 3589, '新增业务变量', 'A', NULL, NULL, '/configure/api/v1/business', 'POST', NULL, 'configure:business:add', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1728056544, 1728056544),
                                                                                                                                                                                                                                     (3592, 3589, '修改业务变量', 'A', NULL, NULL, '/configure/api/v1/business', 'PUT', NULL, 'configure:business:update', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1728056544, 1728056544),
                                                                                                                                                                                                                                     (3593, 3589, '删除业务变量', 'A', NULL, NULL, '/configure/api/v1/business', 'DELETE', NULL, 'configure:business:delete', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1728056544, 1728056544),
                                                                                                                                                                                                                                     (3594, 3589, '查看业务变量值', 'A', NULL, NULL, '/configure/business/values', 'get', NULL, 'configure:business:value:query', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1728056544, 1728056544),
                                                                                                                                                                                                                                     (3595, 3589, '设置业务变量值', 'A', NULL, NULL, '/configure/business/values', 'PUT', NULL, 'configure:business:value:update', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1728056544, 1728056544),
                                                                                                                                                                                                                                     (3596, 3572, '资源变量', 'M', 'ConfigureResource', 'unordered-list', NULL, NULL, '/configure/resource', NULL, '/configure/resource/index', NULL, 0, NULL, 1, NULL, 1, 1728056544, 1728056544),
                                                                                                                                                                                                                                     (3597, 3596, '查看资源', 'A', NULL, NULL, '/configure/api/v1/resources', 'GET', NULL, 'configure:resource:query', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1728056544, 1728056544),
                                                                                                                                                                                                                                     (3598, 3596, '新增资源', 'A', NULL, NULL, '/configure/api/v1/resource', 'POST', NULL, 'configure:resource:add', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1728056544, 1728056544),
                                                                                                                                                                                                                                     (3599, 3596, '修改资源', 'A', NULL, NULL, '/configure/api/v1/resource', 'PUT', NULL, 'configure:resource:update', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1728056544, 1728056544),
                                                                                                                                                                                                                                     (3600, 3596, '删除资源', 'A', NULL, NULL, '/configure/api/v1/resource', 'DELETE', NULL, 'configure:resource:delete', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1728056544, 1728056544),
                                                                                                                                                                                                                                     (3601, 3596, '查看业务变量值', 'A', NULL, NULL, '/configure/resource/values', 'get', NULL, 'configure:resource:value:query', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1728056544, 1728056544),
                                                                                                                                                                                                                                     (3602, 3596, '设置业务变量值', 'A', NULL, NULL, '/configure/resource/values', 'PUT', NULL, 'configure:resource:value:update', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1728056544, 1728056544),
                                                                                                                                                                                                                                     (3603, 3572, '配置模板', 'M', 'ConfgureTemplate', 'file', NULL, NULL, '/configure/template', NULL, '/configure/template/index', NULL, 0, NULL, 1, NULL, 1, 1728056544, 1728056544),
                                                                                                                                                                                                                                     (3604, 3603, '模板管理', 'G', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL, 1728056544, 1728056544),
                                                                                                                                                                                                                                     (3605, 3604, '查看模板历史版本', 'A', NULL, NULL, '/configure/api/v1/templates', 'GET', NULL, 'configure:template:history', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1728056544, 1728056544),
                                                                                                                                                                                                                                     (3606, 3604, '查看指定模板详细数据', 'A', NULL, NULL, '/configure/api/v1/template', 'GET', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL, 1728056544, 1728056544),
                                                                                                                                                                                                                                     (3607, 3604, '查看当前正在使用的模板', 'A', NULL, NULL, '/configure/api/v1/template/current', 'GET', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL, 1728056544, 1728056544),
                                                                                                                                                                                                                                     (3608, 3604, '提交模板', 'A', NULL, NULL, '/configure/api/v1/template', 'POST', NULL, 'configure:template:add', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1728056544, 1728056544),
                                                                                                                                                                                                                                     (3609, 3604, '模板对比', 'A', NULL, NULL, '/configure/api/v1/template/compare', 'POST', NULL, 'configure:template:compare', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1728056544, 1728056544),
                                                                                                                                                                                                                                     (3610, 3604, '切换模板', 'A', NULL, NULL, '/configure/api/v1/template/switch', 'POST', NULL, 'configure:template:switch', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1728056544, 1728056544),
                                                                                                                                                                                                                                     (3611, 3604, '模板预览', 'A', NULL, NULL, '/configure/api/v1/template/preview', 'POST', NULL, 'configure:template:preview', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1728056544, 1728056544),
                                                                                                                                                                                                                                     (3612, 3603, '配置管理', 'G', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL, 1728056544, 1728056544),
                                                                                                                                                                                                                                     (3613, 3612, '配置对比', 'A', NULL, NULL, '/configure/api/v1/configure/compare', 'POST', NULL, 'configure:configure:compare', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1728056544, 1728056544),
                                                                                                                                                                                                                                     (3614, 3612, '同步配置', 'A', NULL, NULL, '/configure/api/v1/configure', 'PUT', NULL, 'configure:configure:sync', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1728056544, 1728056544),
                                                                                                                                                                                                                                     (3615, 3446, '定时任务', 'M', 'SystemPlatformCron', 'schedule', NULL, NULL, '/cron', NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL, 1728056544, 1728056544),
                                                                                                                                                                                                                                     (3616, 3615, '节点管理', 'M', 'Worker', 'common', NULL, NULL, '/cron/worker', NULL, '/cron/worker/index', NULL, 0, 0, 1, NULL, 1, 1728056544, 1728056544),
                                                                                                                                                                                                                                     (3617, 3616, '查看节点分组', 'A', NULL, NULL, '/cron/api/v1/worker_groups', 'GET', NULL, 'cron:worker:group:query', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1728056544, 1728056544),
                                                                                                                                                                                                                                     (3618, 3616, '新增节点分组', 'A', NULL, NULL, '/cron/api/v1/worker_group', 'POST', NULL, 'cron:worker:group:add', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1728056544, 1728056544),
                                                                                                                                                                                                                                     (3619, 3616, '修改节点分组', 'A', NULL, NULL, '/cron/api/v1/worker_group', 'PUT', NULL, 'cron:worker:group:update', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1728056544, 1728056544),
                                                                                                                                                                                                                                     (3620, 3616, '删除节点分组', 'A', NULL, NULL, '/cron/api/v1/worker_group', 'DELETE', NULL, 'cron:worker:group:delete', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1728056544, 1728056544),
                                                                                                                                                                                                                                     (3621, 3616, '查看节点', 'A', NULL, NULL, '/cron/api/v1/workers', 'GET', NULL, 'cron:worker:query', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1728056544, 1728056544),
                                                                                                                                                                                                                                     (3622, 3616, '新增节点', 'A', NULL, NULL, '/cron/api/v1/worker', 'POST', NULL, 'cron:worker:add', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1728056544, 1728056544),
                                                                                                                                                                                                                                     (3623, 3616, '修改节点', 'A', NULL, NULL, '/cron/api/v1/worker', 'PUT', NULL, 'cron:worker:update', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1728056544, 1728056544),
                                                                                                                                                                                                                                     (3624, 3616, '删除节点', 'A', NULL, NULL, '/cron/api/v1/worker', 'DELETE', NULL, 'cron:worker:delete', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1728056544, 1728056544),
                                                                                                                                                                                                                                     (3625, 3616, '更新节点状态', 'A', NULL, NULL, '/cron/api/v1/worker/status', 'PUT', NULL, 'cron:worker:update:status', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1728056544, 1728056544),
                                                                                                                                                                                                                                     (3626, 3615, '任务管理', 'M', 'Task', 'computer', NULL, NULL, '/cron/task', NULL, '/cron/task/index', NULL, 0, 0, 1, NULL, 1, 1728056544, 1728056544),
                                                                                                                                                                                                                                     (3627, 3626, '查看任务分组', 'A', NULL, NULL, '/cron/api/v1/task_groups', 'GET', NULL, 'cron:task:group:query', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1728056544, 1728056544),
                                                                                                                                                                                                                                     (3628, 3626, '新增任务分组', 'A', NULL, NULL, '/cron/api/v1/task_group', 'POST', NULL, 'cron:task:group:add', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1728056544, 1728056544),
                                                                                                                                                                                                                                     (3629, 3626, '修改任务分组', 'A', NULL, NULL, '/cron/api/v1/task_group', 'PUT', NULL, 'cron:task:group:update', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1728056544, 1728056544),
                                                                                                                                                                                                                                     (3630, 3626, '删除任务分组', 'A', NULL, NULL, '/cron/api/v1/task_group', 'DELETE', NULL, 'cron:task:group:delete', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1728056544, 1728056544),
                                                                                                                                                                                                                                     (3631, 3626, '查看任务', 'A', NULL, NULL, '/cron/api/v1/tasks', 'GET', NULL, 'cron:task:query', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1728056544, 1728056544),
                                                                                                                                                                                                                                     (3632, 3626, '新增任务', 'A', NULL, NULL, '/cron/api/v1/task', 'POST', NULL, 'cron:task:add', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1728056544, 1728056544),
                                                                                                                                                                                                                                     (3633, 3626, '立即执行任务', 'A', NULL, NULL, '/cron/api/v1/task/exec', 'POST', NULL, 'cron:task:exec', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1728056544, 1728056544),
                                                                                                                                                                                                                                     (3634, 3626, '取消执行任务', 'A', NULL, NULL, '/cron/api/v1/task/cancel', 'POST', NULL, 'cron:task:cancel', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1728056544, 1728056544),
                                                                                                                                                                                                                                     (3635, 3626, '修改任务', 'A', NULL, NULL, '/cron/api/v1/task', 'PUT', NULL, 'cron:task:update', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1728056544, 1728056544),
                                                                                                                                                                                                                                     (3636, 3626, '删除任务', 'A', NULL, NULL, '/cron/api/v1/task', 'DELETE', NULL, 'cron:task:delete', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1728056544, 1728056544),
                                                                                                                                                                                                                                     (3637, 3626, '任务状态管理', 'A', NULL, NULL, '/cron/api/v1/task/status', 'PUT', NULL, 'cron:task:update:status', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1728056544, 1728056544),
                                                                                                                                                                                                                                     (3638, 3626, '任务日志', 'G', NULL, NULL, NULL, NULL, NULL, 'cron:task:log', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1728056544, 1728056544),
                                                                                                                                                                                                                                     (3639, 3638, '获取任务日志分页', 'A', NULL, NULL, '/cron/api/v1/logs', 'GET', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL, 1728056544, 1728056544),
                                                                                                                                                                                                                                     (3640, 3638, '获取任务日志详情', 'A', NULL, NULL, '/cron/api/v1/log', 'GET', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL, 1728056544, 1728056544),
                                                                                                                                                                                                                                     (3641, 0, '应用平台', 'R', 'AppPlatform', 'apps', NULL, NULL, '/app', NULL, 'Layout', NULL, 2, 0, NULL, NULL, NULL, 1728056544, 1728056544),
                                                                                                                                                                                                                                     (3642, 3641, '首页面板', 'M', 'AppDashboard', 'dashboard', NULL, NULL, '/app/dashboard', NULL, '/dashboard/index', NULL, 0, 0, 1, 1, 1, 1728056544, 1728056544),
                                                                                                                                                                                                                                     (3643, 3641, '信号灯系统', 'M', 'PartyAffairs', 'apps', NULL, NULL, '/party-affairs', NULL, NULL, NULL, 0, 0, NULL, NULL, NULL, 1728056544, 1728056544),
                                                                                                                                                                                                                                     (3644, 3643, '轮播管理', 'M', 'PartyAffairsBanner', 'list', NULL, NULL, '/partyaffairs/banner', NULL, '/partyaffairs/banner/index', NULL, 0, NULL, 1, NULL, 1, 1728056544, 1728056544),
                                                                                                                                                                                                                                     (3645, 3644, '查看轮播', 'A', NULL, NULL, '/partyaffairs/api/v1/banners', 'GET', NULL, 'partyaffairs:banner:query', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1728056544, 1728056544),
                                                                                                                                                                                                                                     (3646, 3644, '新增轮播', 'A', NULL, NULL, '/partyaffairs/api/v1/banner', 'POST', NULL, 'partyaffairs:banner:add', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1728056544, 1728056544),
                                                                                                                                                                                                                                     (3647, 3644, '修改轮播', 'A', NULL, NULL, '/partyaffairs/api/v1/banner', 'PUT', NULL, 'partyaffairs:banner:update', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1728056544, 1728056544),
                                                                                                                                                                                                                                     (3648, 3644, '删除轮播', 'A', NULL, NULL, '/partyaffairs/api/v1/banner', 'DELETE', NULL, 'partyaffairs:banner:delete', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1728056544, 1728056544),
                                                                                                                                                                                                                                     (3649, 3643, '通知管理', 'M', 'PartyAffairsNotice', 'notification', NULL, NULL, '/partyaffairs/notice', NULL, '/partyaffairs/notice/index', NULL, 0, NULL, 1, NULL, 1, 1728056544, 1728056544),
                                                                                                                                                                                                                                     (3650, 3649, '查看通知', 'A', NULL, NULL, '/partyaffairs/api/v1/notices', 'GET', NULL, 'partyaffairs:notice:query', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1728056544, 1728056544),
                                                                                                                                                                                                                                     (3651, 3649, '查看通知阅读情况', 'A', NULL, NULL, '/partyaffairs/api/v1/notice/users', 'GET', NULL, 'partyaffairs:notice:user:query', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1728056544, 1728056544),
                                                                                                                                                                                                                                     (3652, 3649, '新增通知', 'A', NULL, NULL, '/partyaffairs/api/v1/notice', 'POST', NULL, 'partyaffairs:notice:add', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1728056544, 1728056544),
                                                                                                                                                                                                                                     (3653, 3649, '推送通知', 'A', NULL, NULL, '/partyaffairs/api/v1/notice/push', 'POST', NULL, 'partyaffairs:notice:push', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1728056544, 1728056544),
                                                                                                                                                                                                                                     (3654, 3649, '修改通知', 'A', NULL, NULL, '/partyaffairs/api/v1/notice', 'PUT', NULL, 'partyaffairs:notice:update', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1728056544, 1728056544),
                                                                                                                                                                                                                                     (3655, 3649, '删除通知', 'A', NULL, NULL, '/partyaffairs/api/v1/notice', 'DELETE', NULL, 'partyaffairs:notice:delete', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1728056544, 1728056544),
                                                                                                                                                                                                                                     (3656, 3643, '资讯管理', 'M', 'PartyAffairsInformation', 'book', NULL, NULL, '/partyaffairs/information', NULL, '/partyaffairs/information/index', NULL, 0, NULL, 1, NULL, 1, 1728056544, 1728056544),
                                                                                                                                                                                                                                     (3657, 3656, '查看资讯分类', 'A', NULL, NULL, '/partyaffairs/api/v1/information/classifies', 'GET', NULL, 'partyaffairs:information:classify:query', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1728056544, 1728056544),
                                                                                                                                                                                                                                     (3658, 3656, '新增资讯分类', 'A', NULL, NULL, '/partyaffairs/api/v1/information/classify', 'POST', NULL, 'partyaffairs:information:classify:add', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1728056544, 1728056544),
                                                                                                                                                                                                                                     (3659, 3656, '修改资讯分类', 'A', NULL, NULL, '/partyaffairs/api/v1/information/classify', 'PUT', NULL, 'partyaffairs:information:classify:update', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1728056544, 1728056544),
                                                                                                                                                                                                                                     (3660, 3656, '删除资讯分类', 'A', NULL, NULL, '/partyaffairs/api/v1/information/classify', 'DELETE', NULL, 'partyaffairs:information:classify:delete', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1728056544, 1728056544),
                                                                                                                                                                                                                                     (3661, 3656, '查看资讯内容列表', 'A', NULL, NULL, '/partyaffairs/api/v1/informations', 'GET', NULL, 'partyaffairs:information:query', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1728056544, 1728056544),
                                                                                                                                                                                                                                     (3662, 3656, '查看资讯内容', 'A', NULL, NULL, '/partyaffairs/api/v1/information', 'GET', NULL, 'partyaffairs:information:info', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1728056544, 1728056544),
                                                                                                                                                                                                                                     (3663, 3656, '新增资讯内容', 'A', NULL, NULL, '/partyaffairs/api/v1/information', 'POST', NULL, 'partyaffairs:information:add', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1728056544, 1728056544),
                                                                                                                                                                                                                                     (3664, 3656, '修改资讯内容', 'A', NULL, NULL, '/partyaffairs/api/v1/information', 'PUT', NULL, 'partyaffairs:information:update', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1728056544, 1728056544),
                                                                                                                                                                                                                                     (3665, 3656, '删除资讯内容', 'A', NULL, NULL, '/partyaffairs/api/v1/information', 'DELETE', NULL, 'partyaffairs:information:delete', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1728056544, 1728056544),
                                                                                                                                                                                                                                     (3666, 3643, '资料管理', 'M', 'PartyAffairsResource', 'public', NULL, NULL, '/partyaffairs/resource', NULL, '/partyaffairs/resource/index', NULL, 0, NULL, 1, NULL, 1, 1728056544, 1728056544),
                                                                                                                                                                                                                                     (3667, 3666, '查看资料分类', 'A', NULL, NULL, '/partyaffairs/api/v1/resource/classifies', 'GET', NULL, 'partyaffairs:resource:classify:query', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1728056544, 1728056544),
                                                                                                                                                                                                                                     (3668, 3666, '新增资料分类', 'A', NULL, NULL, '/partyaffairs/api/v1/resource/classify', 'POST', NULL, 'partyaffairs:resource:classify:add', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1728056544, 1728056544),
                                                                                                                                                                                                                                     (3669, 3666, '修改资料分类', 'A', NULL, NULL, '/partyaffairs/api/v1/resource/classify', 'PUT', NULL, 'partyaffairs:resource:classify:update', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1728056544, 1728056544),
                                                                                                                                                                                                                                     (3670, 3666, '删除资料分类', 'A', NULL, NULL, '/partyaffairs/api/v1/resource/classify', 'DELETE', NULL, 'partyaffairs:resource:classify:delete', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1728056544, 1728056544),
                                                                                                                                                                                                                                     (3671, 3666, '查看资料内容列表', 'A', NULL, NULL, '/partyaffairs/api/v1/resources', 'GET', NULL, 'partyaffairs:resource:query', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1728056544, 1728056544),
                                                                                                                                                                                                                                     (3672, 3666, '查看资料内容', 'A', NULL, NULL, '/partyaffairs/api/v1/resource', 'GET', NULL, 'partyaffairs:resource:info', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1728056544, 1728056544),
                                                                                                                                                                                                                                     (3673, 3666, '新增资料内容', 'A', NULL, NULL, '/partyaffairs/api/v1/resource', 'POST', NULL, 'partyaffairs:resource:add', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1728056544, 1728056544),
                                                                                                                                                                                                                                     (3674, 3666, '修改资料内容', 'A', NULL, NULL, '/partyaffairs/api/v1/resource', 'PUT', NULL, 'partyaffairs:resource:update', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1728056544, 1728056544),
                                                                                                                                                                                                                                     (3675, 3666, '删除资料内容', 'A', NULL, NULL, '/partyaffairs/api/v1/resource', 'DELETE', NULL, 'partyaffairs:resource:delete', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1728056544, 1728056544),
                                                                                                                                                                                                                                     (3676, 3643, '任务管理', 'M', 'PartyAffairsTask', 'layers', NULL, NULL, '/partyaffairs/task', NULL, '/partyaffairs/task/index', NULL, 0, NULL, 1, NULL, 1, 1728056544, 1728056544),
                                                                                                                                                                                                                                     (3677, 3676, '查看配置列表', 'A', NULL, NULL, '/partyaffairs/api/v1/tasks', 'GET', NULL, 'partyaffairs:task:query', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1728056544, 1728056544),
                                                                                                                                                                                                                                     (3678, 3676, '查看配置', 'A', NULL, NULL, '/partyaffairs/api/v1/task', 'GET', NULL, 'partyaffairs:task:query', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1728056544, 1728056544),
                                                                                                                                                                                                                                     (3679, 3676, '新增配置', 'A', NULL, NULL, '/partyaffairs/api/v1/task', 'POST', NULL, 'partyaffairs:task:add', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1728056544, 1728056544),
                                                                                                                                                                                                                                     (3680, 3676, '修改配置', 'A', NULL, NULL, '/partyaffairs/api/v1/task', 'PUT', NULL, 'partyaffairs:task:update', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1728056544, 1728056544),
                                                                                                                                                                                                                                     (3681, 3676, '删除配置', 'A', NULL, NULL, '/partyaffairs/api/v1/task', 'DELETE', NULL, 'partyaffairs:task:delete', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1728056544, 1728056544),
                                                                                                                                                                                                                                     (3682, 3676, '查看任务填写列表', 'A', NULL, NULL, '/partyaffairs/api/v1/task/values', 'GET', NULL, 'partyaffairs:task:value:query', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1728056544, 1728056544),
                                                                                                                                                                                                                                     (3683, 3676, '查看指定任务填写详情', 'A', NULL, NULL, '/partyaffairs/api/v1/task/value', 'GET', NULL, 'partyaffairs:task:value:query', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1728056544, 1728056544),
                                                                                                                                                                                                                                     (3684, 3676, '导出任务填写详情', 'A', NULL, NULL, '/partyaffairs/api/v1/task/values', 'POST', NULL, 'partyaffairs:task:value:export', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1728056544, 1728056544),
                                                                                                                                                                                                                                     (3685, 3676, '修改任务填写详情', 'A', NULL, NULL, '/partyaffairs/api/v1/task/value', 'PUT', NULL, 'partyaffairs:task:value:update', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1728056544, 1728056544),
                                                                                                                                                                                                                                     (3686, 3676, '删除任务填写详情', 'A', NULL, NULL, '/partyaffairs/api/v1/task/value', 'DELETE', NULL, 'partyaffairs:task:value:delete', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1728056544, 1728056544),
                                                                                                                                                                                                                                     (3687, 3641, '贫困资助系统', 'M', 'poverty', 'apps', NULL, NULL, '/poverty', NULL, NULL, NULL, 0, 0, NULL, NULL, NULL, 1728056544, 1728056544),
                                                                                                                                                                                                                                     (3688, 3687, '轮播管理', 'M', 'PovertyBanner', 'list', NULL, NULL, '/poverty/banner', NULL, '/poverty/banner/index', NULL, 0, NULL, 1, NULL, 1, 1728056544, 1728056544),
                                                                                                                                                                                                                                     (3689, 3688, '查看轮播', 'A', NULL, NULL, '/poverty/api/v1/banners', 'GET', NULL, 'poverty:banner:query', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1728056544, 1728056544),
                                                                                                                                                                                                                                     (3690, 3688, '新增轮播', 'A', NULL, NULL, '/poverty/api/v1/banner', 'POST', NULL, 'poverty:banner:add', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1728056544, 1728056544),
                                                                                                                                                                                                                                     (3691, 3688, '修改轮播', 'A', NULL, NULL, '/poverty/api/v1/banner', 'PUT', NULL, 'poverty:banner:update', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1728056544, 1728056544),
                                                                                                                                                                                                                                     (3692, 3688, '删除轮播', 'A', NULL, NULL, '/poverty/api/v1/banner', 'DELETE', NULL, 'poverty:banner:delete', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1728056544, 1728056544),
                                                                                                                                                                                                                                     (3693, 3687, '通知管理', 'M', 'PovertyNotice', 'notification', NULL, NULL, '/poverty/notice', NULL, '/poverty/notice/index', NULL, 0, NULL, 1, NULL, 1, 1728056544, 1728056544),
                                                                                                                                                                                                                                     (3694, 3693, '查看指定通知', 'A', NULL, NULL, '/poverty/api/v1/notice', 'GET', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL, 1728056544, 1728056544),
                                                                                                                                                                                                                                     (3695, 3693, '查看通知列表', 'A', NULL, NULL, '/poverty/api/v1/notices', 'GET', NULL, 'poverty:notice:query', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1728056544, 1728056544),
                                                                                                                                                                                                                                     (3696, 3693, '新增通知', 'A', NULL, NULL, '/poverty/api/v1/notice', 'POST', NULL, 'poverty:notice:add', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1728056544, 1728056544),
                                                                                                                                                                                                                                     (3697, 3693, '修改通知', 'A', NULL, NULL, '/poverty/api/v1/notice', 'PUT', NULL, 'poverty:notice:update', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1728056544, 1728056544),
                                                                                                                                                                                                                                     (3698, 3693, '删除通知', 'A', NULL, NULL, '/poverty/api/v1/notice', 'DELETE', NULL, 'poverty:notice:delete', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1728056544, 1728056544),
                                                                                                                                                                                                                                     (3699, 3687, '资讯管理', 'M', 'PovertyInformation', 'book', NULL, NULL, '/poverty/information', NULL, '/poverty/information/index', NULL, 0, 0, NULL, NULL, NULL, 1728056544, 1728056544),
                                                                                                                                                                                                                                     (3700, 3699, '查看资讯列表', 'A', NULL, NULL, '/poverty/api/v1/informations', 'GET', NULL, 'poverty:information:query', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1728056544, 1728056544),
                                                                                                                                                                                                                                     (3701, 3699, '查看指定资讯', 'A', NULL, NULL, '/poverty/api/v1/information', 'GET', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL, 1728056544, 1728056544),
                                                                                                                                                                                                                                     (3702, 3699, '新增资讯', 'A', NULL, NULL, '/poverty/api/v1/information', 'POST', NULL, 'poverty:information:add', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1728056544, 1728056544),
                                                                                                                                                                                                                                     (3703, 3699, '修改资讯', 'A', NULL, NULL, '/poverty/api/v1/information', 'PUT', NULL, 'poverty:information:update', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1728056544, 1728056544),
                                                                                                                                                                                                                                     (3704, 3699, '删除资讯', 'A', NULL, NULL, '/poverty/api/v1/information', 'DELETE', NULL, 'poverty:information:delete', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1728056544, 1728056544),
                                                                                                                                                                                                                                     (3705, 3687, '活动管理', 'M', 'PovertyActivity', 'unordered-list', NULL, NULL, '/poverty/activity', NULL, '/poverty/activity/index', NULL, 0, 0, NULL, NULL, NULL, 1728056544, 1728056544),
                                                                                                                                                                                                                                     (3706, 3705, '查看活动列表', 'A', NULL, NULL, '/poverty/api/v1/activitys', 'GET', NULL, 'poverty:activity:query', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1728056544, 1728056544),
                                                                                                                                                                                                                                     (3707, 3705, '查看指定活动', 'A', NULL, NULL, '/poverty/api/v1/activity', 'GET', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL, 1728056544, 1728056544),
                                                                                                                                                                                                                                     (3708, 3705, '新增活动', 'A', NULL, NULL, '/poverty/api/v1/activity', 'POST', NULL, 'poverty:activity:add', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1728056544, 1728056544),
                                                                                                                                                                                                                                     (3709, 3705, '修改活动', 'A', NULL, NULL, '/poverty/api/v1/activity', 'PUT', NULL, 'poverty:activity:update', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1728056544, 1728056544),
                                                                                                                                                                                                                                     (3710, 3705, '删除活动', 'A', NULL, NULL, '/poverty/api/v1/activity', 'DELETE', NULL, 'poverty:activity:delete', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1728056544, 1728056544),
                                                                                                                                                                                                                                     (3711, 3687, '咨询记录', 'M', 'PovertyChat', 'robot', NULL, NULL, '/poverty/chat', NULL, '/poverty/chat/index', NULL, 0, 0, NULL, NULL, NULL, 1728056544, 1728056544),
                                                                                                                                                                                                                                     (3712, 3711, '查看咨询记录', 'A', NULL, NULL, '/poverty/api/v1/chat/records', 'GET', NULL, 'poverty:chat:query', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1728056544, 1728056544),
                                                                                                                                                                                                                                     (3713, 3687, '文件管理', 'M', 'PovertyResource', 'public', NULL, NULL, '/poverty/resource', NULL, '/poverty/resource/index', NULL, 0, 0, NULL, NULL, NULL, 1728056544, 1728056544),
                                                                                                                                                                                                                                     (3714, 3713, '查看文件列表', 'A', NULL, NULL, '/poverty/api/v1/resources', 'GET', NULL, 'poverty:resource:query', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1728056544, 1728056544),
                                                                                                                                                                                                                                     (3715, 3713, '新增文件', 'A', NULL, NULL, '/poverty/api/v1/resource', 'POST', NULL, 'poverty:resource:add', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1728056544, 1728056544),
                                                                                                                                                                                                                                     (3716, 3713, '修改文件', 'A', NULL, NULL, '/poverty/api/v1/resource', 'PUT', NULL, 'poverty:resource:update', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1728056544, 1728056544),
                                                                                                                                                                                                                                     (3717, 3713, '删除文件', 'A', NULL, NULL, '/poverty/api/v1/resource', 'DELETE', NULL, 'poverty:resource:delete', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1728056544, 1728056544),
                                                                                                                                                                                                                                     (3718, 3687, '留言记录', 'M', 'PovertyComment', 'message', NULL, NULL, '/poverty/comment', NULL, '/poverty/comment/index', NULL, 0, 0, NULL, NULL, NULL, 1728056544, 1728056544),
                                                                                                                                                                                                                                     (3719, 3718, '查看留言', 'A', NULL, NULL, '/poverty/api/v1/comments', 'GET', NULL, 'poverty:comment:query', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1728056544, 1728056544),
                                                                                                                                                                                                                                     (3720, 3718, '删除留言', 'A', NULL, NULL, '/poverty/api/v1/comment', 'DELETE', NULL, 'poverty:comment:delete', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1728056544, 1728056544);

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
                                                            (33395, 3446, 3450),
                                                            (33396, 3448, 3450),
                                                            (33397, 3449, 3450),
                                                            (33398, 3446, 3451),
                                                            (33399, 3448, 3451),
                                                            (33400, 3449, 3451),
                                                            (33401, 3446, 3452),
                                                            (33402, 3448, 3452),
                                                            (33403, 3449, 3452),
                                                            (33404, 3446, 3453),
                                                            (33405, 3448, 3453),
                                                            (33406, 3449, 3453),
                                                            (33407, 3446, 3454),
                                                            (33408, 3448, 3454),
                                                            (33409, 3449, 3454),
                                                            (33410, 3446, 3455),
                                                            (33411, 3448, 3455),
                                                            (33412, 3449, 3455),
                                                            (33413, 3446, 3456),
                                                            (33414, 3448, 3456),
                                                            (33415, 3449, 3456),
                                                            (33416, 3446, 3457),
                                                            (33417, 3448, 3457),
                                                            (33418, 3449, 3457),
                                                            (33419, 3446, 3458),
                                                            (33420, 3448, 3458),
                                                            (33421, 3449, 3458),
                                                            (33422, 3446, 3459),
                                                            (33423, 3448, 3459),
                                                            (33424, 3449, 3459),
                                                            (33425, 3446, 3460),
                                                            (33426, 3448, 3460),
                                                            (33427, 3449, 3460),
                                                            (33428, 3446, 3461),
                                                            (33429, 3448, 3461),
                                                            (33430, 3449, 3461),
                                                            (33431, 3446, 3462),
                                                            (33432, 3448, 3462),
                                                            (33433, 3449, 3462),
                                                            (33434, 3446, 3463),
                                                            (33435, 3448, 3463),
                                                            (33436, 3449, 3463),
                                                            (33437, 3446, 3464),
                                                            (33438, 3448, 3464),
                                                            (33439, 3449, 3464),
                                                            (33440, 3446, 3465),
                                                            (33441, 3448, 3465),
                                                            (33442, 3449, 3465),
                                                            (33443, 3446, 3466),
                                                            (33444, 3448, 3466),
                                                            (33445, 3449, 3466),
                                                            (33446, 3446, 3467),
                                                            (33447, 3448, 3467),
                                                            (33448, 3449, 3467),
                                                            (33449, 3446, 3468),
                                                            (33450, 3448, 3468),
                                                            (33451, 3449, 3468),
                                                            (33452, 3446, 3469),
                                                            (33453, 3448, 3469),
                                                            (33454, 3449, 3469),
                                                            (33455, 3446, 3471),
                                                            (33456, 3448, 3471),
                                                            (33457, 3470, 3471),
                                                            (33458, 3446, 3472),
                                                            (33459, 3448, 3472),
                                                            (33460, 3470, 3472),
                                                            (33461, 3446, 3473),
                                                            (33462, 3448, 3473),
                                                            (33463, 3470, 3473),
                                                            (33464, 3446, 3474),
                                                            (33465, 3448, 3474),
                                                            (33466, 3470, 3474),
                                                            (33467, 3446, 3475),
                                                            (33468, 3448, 3475),
                                                            (33469, 3470, 3475),
                                                            (33470, 3446, 3476),
                                                            (33471, 3448, 3476),
                                                            (33472, 3470, 3476),
                                                            (33473, 3446, 3477),
                                                            (33474, 3448, 3477),
                                                            (33475, 3470, 3477),
                                                            (33476, 3446, 3478),
                                                            (33477, 3448, 3478),
                                                            (33478, 3470, 3478),
                                                            (33479, 3446, 3479),
                                                            (33480, 3448, 3479),
                                                            (33481, 3470, 3479),
                                                            (33482, 3446, 3481),
                                                            (33483, 3448, 3481),
                                                            (33484, 3480, 3481),
                                                            (33485, 3446, 3482),
                                                            (33486, 3448, 3482),
                                                            (33487, 3480, 3482),
                                                            (33488, 3446, 3483),
                                                            (33489, 3448, 3483),
                                                            (33490, 3480, 3483),
                                                            (33491, 3446, 3484),
                                                            (33492, 3448, 3484),
                                                            (33493, 3480, 3484),
                                                            (33494, 3446, 3486),
                                                            (33495, 3448, 3486),
                                                            (33496, 3485, 3486),
                                                            (33497, 3446, 3487),
                                                            (33498, 3448, 3487),
                                                            (33499, 3485, 3487),
                                                            (33500, 3446, 3488),
                                                            (33501, 3448, 3488),
                                                            (33502, 3485, 3488),
                                                            (33503, 3446, 3489),
                                                            (33504, 3448, 3489),
                                                            (33505, 3485, 3489),
                                                            (33506, 3446, 3491),
                                                            (33507, 3448, 3491),
                                                            (33508, 3490, 3491),
                                                            (33509, 3446, 3492),
                                                            (33510, 3448, 3492),
                                                            (33511, 3490, 3492),
                                                            (33512, 3446, 3493),
                                                            (33513, 3448, 3493),
                                                            (33514, 3490, 3493),
                                                            (33515, 3446, 3500),
                                                            (33516, 3448, 3500),
                                                            (33517, 3494, 3500),
                                                            (33518, 3499, 3500),
                                                            (33519, 3446, 3501),
                                                            (33520, 3448, 3501),
                                                            (33521, 3494, 3501),
                                                            (33522, 3499, 3501),
                                                            (33523, 3446, 3495),
                                                            (33524, 3448, 3495),
                                                            (33525, 3494, 3495),
                                                            (33526, 3446, 3496),
                                                            (33527, 3448, 3496),
                                                            (33528, 3494, 3496),
                                                            (33529, 3446, 3497),
                                                            (33530, 3448, 3497),
                                                            (33531, 3494, 3497),
                                                            (33532, 3446, 3498),
                                                            (33533, 3448, 3498),
                                                            (33534, 3494, 3498),
                                                            (33535, 3446, 3499),
                                                            (33536, 3448, 3499),
                                                            (33537, 3494, 3499),
                                                            (33538, 3446, 3503),
                                                            (33539, 3448, 3503),
                                                            (33540, 3502, 3503),
                                                            (33541, 3446, 3504),
                                                            (33542, 3448, 3504),
                                                            (33543, 3502, 3504),
                                                            (33544, 3446, 3505),
                                                            (33545, 3448, 3505),
                                                            (33546, 3502, 3505),
                                                            (33547, 3446, 3506),
                                                            (33548, 3448, 3506),
                                                            (33549, 3502, 3506),
                                                            (33550, 3446, 3507),
                                                            (33551, 3448, 3507),
                                                            (33552, 3502, 3507),
                                                            (33553, 3446, 3508),
                                                            (33554, 3448, 3508),
                                                            (33555, 3502, 3508),
                                                            (33556, 3446, 3510),
                                                            (33557, 3448, 3510),
                                                            (33558, 3509, 3510),
                                                            (33559, 3446, 3449),
                                                            (33560, 3448, 3449),
                                                            (33561, 3446, 3470),
                                                            (33562, 3448, 3470),
                                                            (33563, 3446, 3480),
                                                            (33564, 3448, 3480),
                                                            (33565, 3446, 3485),
                                                            (33566, 3448, 3485),
                                                            (33567, 3446, 3490),
                                                            (33568, 3448, 3490),
                                                            (33569, 3446, 3494),
                                                            (33570, 3448, 3494),
                                                            (33571, 3446, 3502),
                                                            (33572, 3448, 3502),
                                                            (33573, 3446, 3509),
                                                            (33574, 3448, 3509),
                                                            (33575, 3446, 3514),
                                                            (33576, 3511, 3514),
                                                            (33577, 3512, 3514),
                                                            (33578, 3513, 3514),
                                                            (33579, 3446, 3515),
                                                            (33580, 3511, 3515),
                                                            (33581, 3512, 3515),
                                                            (33582, 3513, 3515),
                                                            (33583, 3446, 3516),
                                                            (33584, 3511, 3516),
                                                            (33585, 3512, 3516),
                                                            (33586, 3513, 3516),
                                                            (33587, 3446, 3517),
                                                            (33588, 3511, 3517),
                                                            (33589, 3512, 3517),
                                                            (33590, 3513, 3517),
                                                            (33591, 3446, 3519),
                                                            (33592, 3511, 3519),
                                                            (33593, 3512, 3519),
                                                            (33594, 3518, 3519),
                                                            (33595, 3446, 3520),
                                                            (33596, 3511, 3520),
                                                            (33597, 3512, 3520),
                                                            (33598, 3518, 3520),
                                                            (33599, 3446, 3521),
                                                            (33600, 3511, 3521),
                                                            (33601, 3512, 3521),
                                                            (33602, 3518, 3521),
                                                            (33603, 3446, 3513),
                                                            (33604, 3511, 3513),
                                                            (33605, 3512, 3513),
                                                            (33606, 3446, 3518),
                                                            (33607, 3511, 3518),
                                                            (33608, 3512, 3518),
                                                            (33609, 3446, 3523),
                                                            (33610, 3511, 3523),
                                                            (33611, 3522, 3523),
                                                            (33612, 3446, 3524),
                                                            (33613, 3511, 3524),
                                                            (33614, 3522, 3524),
                                                            (33615, 3446, 3525),
                                                            (33616, 3511, 3525),
                                                            (33617, 3522, 3525),
                                                            (33618, 3446, 3512),
                                                            (33619, 3511, 3512),
                                                            (33620, 3446, 3522),
                                                            (33621, 3511, 3522),
                                                            (33622, 3446, 3528),
                                                            (33623, 3526, 3528),
                                                            (33624, 3527, 3528),
                                                            (33625, 3446, 3529),
                                                            (33626, 3526, 3529),
                                                            (33627, 3527, 3529),
                                                            (33628, 3446, 3530),
                                                            (33629, 3526, 3530),
                                                            (33630, 3527, 3530),
                                                            (33631, 3446, 3531),
                                                            (33632, 3526, 3531),
                                                            (33633, 3527, 3531),
                                                            (33634, 3446, 3533),
                                                            (33635, 3526, 3533),
                                                            (33636, 3532, 3533),
                                                            (33637, 3446, 3534),
                                                            (33638, 3526, 3534),
                                                            (33639, 3532, 3534),
                                                            (33640, 3446, 3535),
                                                            (33641, 3526, 3535),
                                                            (33642, 3532, 3535),
                                                            (33643, 3446, 3536),
                                                            (33644, 3526, 3536),
                                                            (33645, 3532, 3536),
                                                            (33646, 3446, 3537),
                                                            (33647, 3526, 3537),
                                                            (33648, 3532, 3537),
                                                            (33649, 3446, 3539),
                                                            (33650, 3526, 3539),
                                                            (33651, 3538, 3539),
                                                            (33652, 3446, 3540),
                                                            (33653, 3526, 3540),
                                                            (33654, 3538, 3540),
                                                            (33655, 3446, 3541),
                                                            (33656, 3526, 3541),
                                                            (33657, 3538, 3541),
                                                            (33658, 3446, 3542),
                                                            (33659, 3526, 3542),
                                                            (33660, 3538, 3542),
                                                            (33661, 3446, 3543),
                                                            (33662, 3526, 3543),
                                                            (33663, 3538, 3543),
                                                            (33664, 3446, 3544),
                                                            (33665, 3526, 3544),
                                                            (33666, 3538, 3544),
                                                            (33667, 3446, 3546),
                                                            (33668, 3526, 3546),
                                                            (33669, 3545, 3546),
                                                            (33670, 3446, 3547),
                                                            (33671, 3526, 3547),
                                                            (33672, 3545, 3547),
                                                            (33673, 3446, 3548),
                                                            (33674, 3526, 3548),
                                                            (33675, 3545, 3548),
                                                            (33676, 3446, 3549),
                                                            (33677, 3526, 3549),
                                                            (33678, 3545, 3549),
                                                            (33679, 3446, 3550),
                                                            (33680, 3526, 3550),
                                                            (33681, 3545, 3550),
                                                            (33682, 3446, 3551),
                                                            (33683, 3526, 3551),
                                                            (33684, 3545, 3551),
                                                            (33685, 3446, 3552),
                                                            (33686, 3526, 3552),
                                                            (33687, 3545, 3552),
                                                            (33688, 3446, 3553),
                                                            (33689, 3526, 3553),
                                                            (33690, 3545, 3553),
                                                            (33691, 3446, 3562),
                                                            (33692, 3526, 3562),
                                                            (33693, 3554, 3562),
                                                            (33694, 3561, 3562),
                                                            (33695, 3446, 3563),
                                                            (33696, 3526, 3563),
                                                            (33697, 3554, 3563),
                                                            (33698, 3561, 3563),
                                                            (33699, 3446, 3564),
                                                            (33700, 3526, 3564),
                                                            (33701, 3554, 3564),
                                                            (33702, 3561, 3564),
                                                            (33703, 3446, 3565),
                                                            (33704, 3526, 3565),
                                                            (33705, 3554, 3565),
                                                            (33706, 3561, 3565),
                                                            (33707, 3446, 3566),
                                                            (33708, 3526, 3566),
                                                            (33709, 3554, 3566),
                                                            (33710, 3561, 3566),
                                                            (33711, 3446, 3567),
                                                            (33712, 3526, 3567),
                                                            (33713, 3554, 3567),
                                                            (33714, 3561, 3567),
                                                            (33715, 3446, 3568),
                                                            (33716, 3526, 3568),
                                                            (33717, 3554, 3568),
                                                            (33718, 3561, 3568),
                                                            (33719, 3446, 3569),
                                                            (33720, 3526, 3569),
                                                            (33721, 3554, 3569),
                                                            (33722, 3561, 3569),
                                                            (33723, 3446, 3570),
                                                            (33724, 3526, 3570),
                                                            (33725, 3554, 3570),
                                                            (33726, 3561, 3570),
                                                            (33727, 3446, 3571),
                                                            (33728, 3526, 3571),
                                                            (33729, 3554, 3571),
                                                            (33730, 3561, 3571),
                                                            (33731, 3446, 3555),
                                                            (33732, 3526, 3555),
                                                            (33733, 3554, 3555),
                                                            (33734, 3446, 3556),
                                                            (33735, 3526, 3556),
                                                            (33736, 3554, 3556),
                                                            (33737, 3446, 3557),
                                                            (33738, 3526, 3557),
                                                            (33739, 3554, 3557),
                                                            (33740, 3446, 3558),
                                                            (33741, 3526, 3558),
                                                            (33742, 3554, 3558),
                                                            (33743, 3446, 3559),
                                                            (33744, 3526, 3559),
                                                            (33745, 3554, 3559),
                                                            (33746, 3446, 3560),
                                                            (33747, 3526, 3560),
                                                            (33748, 3554, 3560),
                                                            (33749, 3446, 3561),
                                                            (33750, 3526, 3561),
                                                            (33751, 3554, 3561),
                                                            (33752, 3446, 3527),
                                                            (33753, 3526, 3527),
                                                            (33754, 3446, 3532),
                                                            (33755, 3526, 3532),
                                                            (33756, 3446, 3538),
                                                            (33757, 3526, 3538),
                                                            (33758, 3446, 3545),
                                                            (33759, 3526, 3545),
                                                            (33760, 3446, 3554),
                                                            (33761, 3526, 3554),
                                                            (33762, 3446, 3574),
                                                            (33763, 3572, 3574),
                                                            (33764, 3573, 3574),
                                                            (33765, 3446, 3575),
                                                            (33766, 3572, 3575),
                                                            (33767, 3573, 3575),
                                                            (33768, 3446, 3576),
                                                            (33769, 3572, 3576),
                                                            (33770, 3573, 3576),
                                                            (33771, 3446, 3577),
                                                            (33772, 3572, 3577),
                                                            (33773, 3573, 3577),
                                                            (33774, 3446, 3578),
                                                            (33775, 3572, 3578),
                                                            (33776, 3573, 3578),
                                                            (33777, 3446, 3579),
                                                            (33778, 3572, 3579),
                                                            (33779, 3573, 3579),
                                                            (33780, 3446, 3580),
                                                            (33781, 3572, 3580),
                                                            (33782, 3573, 3580),
                                                            (33783, 3446, 3581),
                                                            (33784, 3572, 3581),
                                                            (33785, 3573, 3581),
                                                            (33786, 3446, 3583),
                                                            (33787, 3572, 3583),
                                                            (33788, 3582, 3583),
                                                            (33789, 3446, 3584),
                                                            (33790, 3572, 3584),
                                                            (33791, 3582, 3584),
                                                            (33792, 3446, 3585),
                                                            (33793, 3572, 3585),
                                                            (33794, 3582, 3585),
                                                            (33795, 3446, 3586),
                                                            (33796, 3572, 3586),
                                                            (33797, 3582, 3586),
                                                            (33798, 3446, 3587),
                                                            (33799, 3572, 3587),
                                                            (33800, 3582, 3587),
                                                            (33801, 3446, 3588),
                                                            (33802, 3572, 3588),
                                                            (33803, 3582, 3588),
                                                            (33804, 3446, 3590),
                                                            (33805, 3572, 3590),
                                                            (33806, 3589, 3590),
                                                            (33807, 3446, 3591),
                                                            (33808, 3572, 3591),
                                                            (33809, 3589, 3591),
                                                            (33810, 3446, 3592),
                                                            (33811, 3572, 3592),
                                                            (33812, 3589, 3592),
                                                            (33813, 3446, 3593),
                                                            (33814, 3572, 3593),
                                                            (33815, 3589, 3593),
                                                            (33816, 3446, 3594),
                                                            (33817, 3572, 3594),
                                                            (33818, 3589, 3594),
                                                            (33819, 3446, 3595),
                                                            (33820, 3572, 3595),
                                                            (33821, 3589, 3595),
                                                            (33822, 3446, 3597),
                                                            (33823, 3572, 3597),
                                                            (33824, 3596, 3597),
                                                            (33825, 3446, 3598),
                                                            (33826, 3572, 3598),
                                                            (33827, 3596, 3598),
                                                            (33828, 3446, 3599),
                                                            (33829, 3572, 3599),
                                                            (33830, 3596, 3599),
                                                            (33831, 3446, 3600),
                                                            (33832, 3572, 3600),
                                                            (33833, 3596, 3600),
                                                            (33834, 3446, 3601),
                                                            (33835, 3572, 3601),
                                                            (33836, 3596, 3601),
                                                            (33837, 3446, 3602),
                                                            (33838, 3572, 3602),
                                                            (33839, 3596, 3602),
                                                            (33840, 3446, 3605),
                                                            (33841, 3572, 3605),
                                                            (33842, 3603, 3605),
                                                            (33843, 3604, 3605),
                                                            (33844, 3446, 3606),
                                                            (33845, 3572, 3606),
                                                            (33846, 3603, 3606),
                                                            (33847, 3604, 3606),
                                                            (33848, 3446, 3607),
                                                            (33849, 3572, 3607),
                                                            (33850, 3603, 3607),
                                                            (33851, 3604, 3607),
                                                            (33852, 3446, 3608),
                                                            (33853, 3572, 3608),
                                                            (33854, 3603, 3608),
                                                            (33855, 3604, 3608),
                                                            (33856, 3446, 3609),
                                                            (33857, 3572, 3609),
                                                            (33858, 3603, 3609),
                                                            (33859, 3604, 3609),
                                                            (33860, 3446, 3610),
                                                            (33861, 3572, 3610),
                                                            (33862, 3603, 3610),
                                                            (33863, 3604, 3610),
                                                            (33864, 3446, 3611),
                                                            (33865, 3572, 3611),
                                                            (33866, 3603, 3611),
                                                            (33867, 3604, 3611),
                                                            (33868, 3446, 3613),
                                                            (33869, 3572, 3613),
                                                            (33870, 3603, 3613),
                                                            (33871, 3612, 3613),
                                                            (33872, 3446, 3614),
                                                            (33873, 3572, 3614),
                                                            (33874, 3603, 3614),
                                                            (33875, 3612, 3614),
                                                            (33876, 3446, 3604),
                                                            (33877, 3572, 3604),
                                                            (33878, 3603, 3604),
                                                            (33879, 3446, 3612),
                                                            (33880, 3572, 3612),
                                                            (33881, 3603, 3612),
                                                            (33882, 3446, 3573),
                                                            (33883, 3572, 3573),
                                                            (33884, 3446, 3582),
                                                            (33885, 3572, 3582),
                                                            (33886, 3446, 3589),
                                                            (33887, 3572, 3589),
                                                            (33888, 3446, 3596),
                                                            (33889, 3572, 3596),
                                                            (33890, 3446, 3603),
                                                            (33891, 3572, 3603),
                                                            (33892, 3446, 3617),
                                                            (33893, 3615, 3617),
                                                            (33894, 3616, 3617),
                                                            (33895, 3446, 3618),
                                                            (33896, 3615, 3618),
                                                            (33897, 3616, 3618),
                                                            (33898, 3446, 3619),
                                                            (33899, 3615, 3619),
                                                            (33900, 3616, 3619),
                                                            (33901, 3446, 3620),
                                                            (33902, 3615, 3620),
                                                            (33903, 3616, 3620),
                                                            (33904, 3446, 3621),
                                                            (33905, 3615, 3621),
                                                            (33906, 3616, 3621),
                                                            (33907, 3446, 3622),
                                                            (33908, 3615, 3622),
                                                            (33909, 3616, 3622),
                                                            (33910, 3446, 3623),
                                                            (33911, 3615, 3623),
                                                            (33912, 3616, 3623),
                                                            (33913, 3446, 3624),
                                                            (33914, 3615, 3624),
                                                            (33915, 3616, 3624),
                                                            (33916, 3446, 3625),
                                                            (33917, 3615, 3625),
                                                            (33918, 3616, 3625),
                                                            (33919, 3446, 3639),
                                                            (33920, 3615, 3639),
                                                            (33921, 3626, 3639),
                                                            (33922, 3638, 3639),
                                                            (33923, 3446, 3640),
                                                            (33924, 3615, 3640),
                                                            (33925, 3626, 3640),
                                                            (33926, 3638, 3640),
                                                            (33927, 3446, 3627),
                                                            (33928, 3615, 3627),
                                                            (33929, 3626, 3627),
                                                            (33930, 3446, 3628),
                                                            (33931, 3615, 3628),
                                                            (33932, 3626, 3628),
                                                            (33933, 3446, 3629),
                                                            (33934, 3615, 3629),
                                                            (33935, 3626, 3629),
                                                            (33936, 3446, 3630),
                                                            (33937, 3615, 3630),
                                                            (33938, 3626, 3630),
                                                            (33939, 3446, 3631),
                                                            (33940, 3615, 3631),
                                                            (33941, 3626, 3631),
                                                            (33942, 3446, 3632),
                                                            (33943, 3615, 3632),
                                                            (33944, 3626, 3632),
                                                            (33945, 3446, 3633),
                                                            (33946, 3615, 3633),
                                                            (33947, 3626, 3633),
                                                            (33948, 3446, 3634),
                                                            (33949, 3615, 3634),
                                                            (33950, 3626, 3634),
                                                            (33951, 3446, 3635),
                                                            (33952, 3615, 3635),
                                                            (33953, 3626, 3635),
                                                            (33954, 3446, 3636),
                                                            (33955, 3615, 3636),
                                                            (33956, 3626, 3636),
                                                            (33957, 3446, 3637),
                                                            (33958, 3615, 3637),
                                                            (33959, 3626, 3637),
                                                            (33960, 3446, 3638),
                                                            (33961, 3615, 3638),
                                                            (33962, 3626, 3638),
                                                            (33963, 3446, 3616),
                                                            (33964, 3615, 3616),
                                                            (33965, 3446, 3626),
                                                            (33966, 3615, 3626),
                                                            (33967, 3446, 3447),
                                                            (33968, 3446, 3448),
                                                            (33969, 3446, 3511),
                                                            (33970, 3446, 3526),
                                                            (33971, 3446, 3572),
                                                            (33972, 3446, 3615),
                                                            (33973, 3641, 3645),
                                                            (33974, 3643, 3645),
                                                            (33975, 3644, 3645),
                                                            (33976, 3641, 3646),
                                                            (33977, 3643, 3646),
                                                            (33978, 3644, 3646),
                                                            (33979, 3641, 3647),
                                                            (33980, 3643, 3647),
                                                            (33981, 3644, 3647),
                                                            (33982, 3641, 3648),
                                                            (33983, 3643, 3648),
                                                            (33984, 3644, 3648),
                                                            (33985, 3641, 3650),
                                                            (33986, 3643, 3650),
                                                            (33987, 3649, 3650),
                                                            (33988, 3641, 3651),
                                                            (33989, 3643, 3651),
                                                            (33990, 3649, 3651),
                                                            (33991, 3641, 3652),
                                                            (33992, 3643, 3652),
                                                            (33993, 3649, 3652),
                                                            (33994, 3641, 3653),
                                                            (33995, 3643, 3653),
                                                            (33996, 3649, 3653),
                                                            (33997, 3641, 3654),
                                                            (33998, 3643, 3654),
                                                            (33999, 3649, 3654),
                                                            (34000, 3641, 3655),
                                                            (34001, 3643, 3655),
                                                            (34002, 3649, 3655),
                                                            (34003, 3641, 3657),
                                                            (34004, 3643, 3657),
                                                            (34005, 3656, 3657),
                                                            (34006, 3641, 3658),
                                                            (34007, 3643, 3658),
                                                            (34008, 3656, 3658),
                                                            (34009, 3641, 3659),
                                                            (34010, 3643, 3659),
                                                            (34011, 3656, 3659),
                                                            (34012, 3641, 3660),
                                                            (34013, 3643, 3660),
                                                            (34014, 3656, 3660),
                                                            (34015, 3641, 3661),
                                                            (34016, 3643, 3661),
                                                            (34017, 3656, 3661),
                                                            (34018, 3641, 3662),
                                                            (34019, 3643, 3662),
                                                            (34020, 3656, 3662),
                                                            (34021, 3641, 3663),
                                                            (34022, 3643, 3663),
                                                            (34023, 3656, 3663),
                                                            (34024, 3641, 3664),
                                                            (34025, 3643, 3664),
                                                            (34026, 3656, 3664),
                                                            (34027, 3641, 3665),
                                                            (34028, 3643, 3665),
                                                            (34029, 3656, 3665),
                                                            (34030, 3641, 3667),
                                                            (34031, 3643, 3667),
                                                            (34032, 3666, 3667),
                                                            (34033, 3641, 3668),
                                                            (34034, 3643, 3668),
                                                            (34035, 3666, 3668),
                                                            (34036, 3641, 3669),
                                                            (34037, 3643, 3669),
                                                            (34038, 3666, 3669),
                                                            (34039, 3641, 3670),
                                                            (34040, 3643, 3670),
                                                            (34041, 3666, 3670),
                                                            (34042, 3641, 3671),
                                                            (34043, 3643, 3671),
                                                            (34044, 3666, 3671),
                                                            (34045, 3641, 3672),
                                                            (34046, 3643, 3672),
                                                            (34047, 3666, 3672),
                                                            (34048, 3641, 3673),
                                                            (34049, 3643, 3673),
                                                            (34050, 3666, 3673),
                                                            (34051, 3641, 3674),
                                                            (34052, 3643, 3674),
                                                            (34053, 3666, 3674),
                                                            (34054, 3641, 3675),
                                                            (34055, 3643, 3675),
                                                            (34056, 3666, 3675),
                                                            (34057, 3641, 3677),
                                                            (34058, 3643, 3677),
                                                            (34059, 3676, 3677),
                                                            (34060, 3641, 3678),
                                                            (34061, 3643, 3678),
                                                            (34062, 3676, 3678),
                                                            (34063, 3641, 3679),
                                                            (34064, 3643, 3679),
                                                            (34065, 3676, 3679),
                                                            (34066, 3641, 3680),
                                                            (34067, 3643, 3680),
                                                            (34068, 3676, 3680),
                                                            (34069, 3641, 3681),
                                                            (34070, 3643, 3681),
                                                            (34071, 3676, 3681),
                                                            (34072, 3641, 3682),
                                                            (34073, 3643, 3682),
                                                            (34074, 3676, 3682),
                                                            (34075, 3641, 3683),
                                                            (34076, 3643, 3683),
                                                            (34077, 3676, 3683),
                                                            (34078, 3641, 3684),
                                                            (34079, 3643, 3684),
                                                            (34080, 3676, 3684),
                                                            (34081, 3641, 3685),
                                                            (34082, 3643, 3685),
                                                            (34083, 3676, 3685),
                                                            (34084, 3641, 3686),
                                                            (34085, 3643, 3686),
                                                            (34086, 3676, 3686),
                                                            (34087, 3641, 3644),
                                                            (34088, 3643, 3644),
                                                            (34089, 3641, 3649),
                                                            (34090, 3643, 3649),
                                                            (34091, 3641, 3656),
                                                            (34092, 3643, 3656),
                                                            (34093, 3641, 3666),
                                                            (34094, 3643, 3666),
                                                            (34095, 3641, 3676),
                                                            (34096, 3643, 3676),
                                                            (34097, 3641, 3689),
                                                            (34098, 3687, 3689),
                                                            (34099, 3688, 3689),
                                                            (34100, 3641, 3690),
                                                            (34101, 3687, 3690),
                                                            (34102, 3688, 3690),
                                                            (34103, 3641, 3691),
                                                            (34104, 3687, 3691),
                                                            (34105, 3688, 3691),
                                                            (34106, 3641, 3692),
                                                            (34107, 3687, 3692),
                                                            (34108, 3688, 3692),
                                                            (34109, 3641, 3694),
                                                            (34110, 3687, 3694),
                                                            (34111, 3693, 3694),
                                                            (34112, 3641, 3695),
                                                            (34113, 3687, 3695),
                                                            (34114, 3693, 3695),
                                                            (34115, 3641, 3696),
                                                            (34116, 3687, 3696),
                                                            (34117, 3693, 3696),
                                                            (34118, 3641, 3697),
                                                            (34119, 3687, 3697),
                                                            (34120, 3693, 3697),
                                                            (34121, 3641, 3698),
                                                            (34122, 3687, 3698),
                                                            (34123, 3693, 3698),
                                                            (34124, 3641, 3700),
                                                            (34125, 3687, 3700),
                                                            (34126, 3699, 3700),
                                                            (34127, 3641, 3701),
                                                            (34128, 3687, 3701),
                                                            (34129, 3699, 3701),
                                                            (34130, 3641, 3702),
                                                            (34131, 3687, 3702),
                                                            (34132, 3699, 3702),
                                                            (34133, 3641, 3703),
                                                            (34134, 3687, 3703),
                                                            (34135, 3699, 3703),
                                                            (34136, 3641, 3704),
                                                            (34137, 3687, 3704),
                                                            (34138, 3699, 3704),
                                                            (34139, 3641, 3706),
                                                            (34140, 3687, 3706),
                                                            (34141, 3705, 3706),
                                                            (34142, 3641, 3707),
                                                            (34143, 3687, 3707),
                                                            (34144, 3705, 3707),
                                                            (34145, 3641, 3708),
                                                            (34146, 3687, 3708),
                                                            (34147, 3705, 3708),
                                                            (34148, 3641, 3709),
                                                            (34149, 3687, 3709),
                                                            (34150, 3705, 3709),
                                                            (34151, 3641, 3710),
                                                            (34152, 3687, 3710),
                                                            (34153, 3705, 3710),
                                                            (34154, 3641, 3712),
                                                            (34155, 3687, 3712),
                                                            (34156, 3711, 3712),
                                                            (34157, 3641, 3714),
                                                            (34158, 3687, 3714),
                                                            (34159, 3713, 3714),
                                                            (34160, 3641, 3715),
                                                            (34161, 3687, 3715),
                                                            (34162, 3713, 3715),
                                                            (34163, 3641, 3716),
                                                            (34164, 3687, 3716),
                                                            (34165, 3713, 3716),
                                                            (34166, 3641, 3717),
                                                            (34167, 3687, 3717),
                                                            (34168, 3713, 3717),
                                                            (34169, 3641, 3719),
                                                            (34170, 3687, 3719),
                                                            (34171, 3718, 3719),
                                                            (34172, 3641, 3720),
                                                            (34173, 3687, 3720),
                                                            (34174, 3718, 3720),
                                                            (34175, 3641, 3688),
                                                            (34176, 3687, 3688),
                                                            (34177, 3641, 3693),
                                                            (34178, 3687, 3693),
                                                            (34179, 3641, 3699),
                                                            (34180, 3687, 3699),
                                                            (34181, 3641, 3705),
                                                            (34182, 3687, 3705),
                                                            (34183, 3641, 3711),
                                                            (34184, 3687, 3711),
                                                            (34185, 3641, 3713),
                                                            (34186, 3687, 3713),
                                                            (34187, 3641, 3718),
                                                            (34188, 3687, 3718),
                                                            (34189, 3641, 3642),
                                                            (34190, 3641, 3643),
                                                            (34191, 3641, 3687);

-- --------------------------------------------------------

--
-- 表的结构 `resource`
--

CREATE TABLE `resource` (
                            `id` bigint(20) UNSIGNED NOT NULL COMMENT '主键ID',
                            `keyword` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT '关键字',
                            `department_id` bigint(20) UNSIGNED NOT NULL COMMENT '部门id',
                            `resource_id` bigint(20) UNSIGNED NOT NULL COMMENT '资源id'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='资源权限信息';

--
-- 转存表中的数据 `resource`
--

INSERT INTO `resource` (`id`, `keyword`, `department_id`, `resource_id`) VALUES
                                                                             (2, 'uc_app', 5, 1),
                                                                             (1, 'uc_app', 5, 2);

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
                                                                                                                                                                 (5, 1, '信号灯管理员', '3', 1, '管理员', NULL, 'CUR_DOWN', 1719464519, 1723030092, 0),
                                                                                                                                                                 (6, 1, '资助系统管理员', 'povertyAdmin', 1, '资助系统管理员', NULL, 'CUR_DOWN', 1723030136, 1723030136, 0);

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
                                                            (5, 1, 5),
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
                                                         (797, 5, 3446),
                                                         (798, 5, 3448),
                                                         (786, 5, 3490),
                                                         (783, 5, 3491),
                                                         (784, 5, 3492),
                                                         (785, 5, 3493),
                                                         (787, 5, 3494),
                                                         (788, 5, 3495),
                                                         (789, 5, 3496),
                                                         (790, 5, 3497),
                                                         (791, 5, 3498),
                                                         (792, 5, 3499),
                                                         (793, 5, 3500),
                                                         (794, 5, 3501),
                                                         (796, 5, 3526),
                                                         (782, 5, 3545),
                                                         (774, 5, 3546),
                                                         (775, 5, 3547),
                                                         (776, 5, 3548),
                                                         (777, 5, 3549),
                                                         (778, 5, 3550),
                                                         (779, 5, 3551),
                                                         (780, 5, 3552),
                                                         (781, 5, 3553),
                                                         (773, 5, 3554),
                                                         (756, 5, 3555),
                                                         (757, 5, 3556),
                                                         (758, 5, 3557),
                                                         (759, 5, 3558),
                                                         (760, 5, 3559),
                                                         (761, 5, 3560),
                                                         (772, 5, 3561),
                                                         (762, 5, 3562),
                                                         (763, 5, 3563),
                                                         (764, 5, 3564),
                                                         (765, 5, 3565),
                                                         (766, 5, 3566),
                                                         (767, 5, 3567),
                                                         (768, 5, 3568),
                                                         (769, 5, 3569),
                                                         (770, 5, 3570),
                                                         (771, 5, 3571),
                                                         (795, 5, 3641),
                                                         (755, 5, 3643),
                                                         (716, 5, 3644),
                                                         (712, 5, 3645),
                                                         (713, 5, 3646),
                                                         (714, 5, 3647),
                                                         (715, 5, 3648),
                                                         (723, 5, 3649),
                                                         (717, 5, 3650),
                                                         (718, 5, 3651),
                                                         (719, 5, 3652),
                                                         (720, 5, 3653),
                                                         (721, 5, 3654),
                                                         (722, 5, 3655),
                                                         (733, 5, 3656),
                                                         (724, 5, 3657),
                                                         (725, 5, 3658),
                                                         (726, 5, 3659),
                                                         (727, 5, 3660),
                                                         (728, 5, 3661),
                                                         (729, 5, 3662),
                                                         (730, 5, 3663),
                                                         (731, 5, 3664),
                                                         (732, 5, 3665),
                                                         (743, 5, 3666),
                                                         (734, 5, 3667),
                                                         (735, 5, 3668),
                                                         (736, 5, 3669),
                                                         (737, 5, 3670),
                                                         (738, 5, 3671),
                                                         (739, 5, 3672),
                                                         (740, 5, 3673),
                                                         (741, 5, 3674),
                                                         (742, 5, 3675),
                                                         (754, 5, 3676),
                                                         (744, 5, 3677),
                                                         (745, 5, 3678),
                                                         (746, 5, 3679),
                                                         (747, 5, 3680),
                                                         (748, 5, 3681),
                                                         (749, 5, 3682),
                                                         (750, 5, 3683),
                                                         (751, 5, 3684),
                                                         (752, 5, 3685),
                                                         (753, 5, 3686);

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

INSERT INTO `user` VALUES (1, 1, 1, '超级管理员', '超级管理员', 'F', 'a9f224627346905e258d771e4043f921', '1280291001@qq.com', '18888888888', '$2a$10$9qRJe9KQo6sEcU8ipKg.e.dkl2E7Wy64SigYlgraTAn.1paHFq6W.', 1, '{\"theme\":\"light\",\"themeColor\":\"#165DFF\",\"skin\":\"default\",\"tabBar\":true,\"menuWidth\":200,\"layout\":\"default\",\"language\":\"zh_CN\",\"animation\":\"gp-fade\"}', 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJkZXBhcnRtZW50SWQiOjEsImRlcGFydG1lbnRLZXl3b3JkIjoiY29tcGFueSIsImV4cCI6MTcyNDc1MTczNywiaWF0IjoxNzI0NzQ0NTM2LCJuYmYiOjE3MjQ3NDQ1MzYsInJvbGVJZCI6MSwicm9sZUtleXdvcmQiOiJzdXBlckFkbWluIiwidXNlcklkIjoxfQ.kZ2y12xxhJXRdfIv_43vkrE0E23XI62TufeYtSBQWZo', 1724744536, 1713706137, 1724744536);

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

INSERT INTO `user_job` (`id`, `job_id`, `user_id`) VALUES (1, 1, 1);

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

INSERT INTO `user_role` (`id`, `role_id`, `user_id`) VALUES (1, 1, 1);

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
-- 表的索引 `login_log`
--
ALTER TABLE `login_log`
    ADD PRIMARY KEY (`id`),
  ADD KEY `created_at` (`created_at`);

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
-- 表的索引 `resource`
--
ALTER TABLE `resource`
    ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `department_id` (`keyword`,`department_id`,`resource_id`),
  ADD KEY `department_id_2` (`department_id`);

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
    MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=1041;

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
    MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键ID', AUTO_INCREMENT=4;

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
-- 使用表AUTO_INCREMENT `login_log`
--
ALTER TABLE `login_log`
    MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键ID', AUTO_INCREMENT=33;

--
-- 使用表AUTO_INCREMENT `menu`
--
ALTER TABLE `menu`
    MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键ID', AUTO_INCREMENT=7368;

--
-- 使用表AUTO_INCREMENT `menu_closure`
--
ALTER TABLE `menu_closure`
    MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键ID', AUTO_INCREMENT=34192;

--
-- 使用表AUTO_INCREMENT `resource`
--
ALTER TABLE `resource`
    MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键ID', AUTO_INCREMENT=3;

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
    MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键ID', AUTO_INCREMENT=799;

--
-- 使用表AUTO_INCREMENT `user`
--
ALTER TABLE `user`
    MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键ID', AUTO_INCREMENT=6;

--
-- 使用表AUTO_INCREMENT `user_job`
--
ALTER TABLE `user_job`
    MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键ID', AUTO_INCREMENT=11;

--
-- 使用表AUTO_INCREMENT `user_role`
--
ALTER TABLE `user_role`
    MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键ID', AUTO_INCREMENT=8;

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
-- 限制表 `resource`
--
ALTER TABLE `resource`
    ADD CONSTRAINT `resource_ibfk_1` FOREIGN KEY (`department_id`) REFERENCES `department` (`id`) ON DELETE CASCADE;

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
