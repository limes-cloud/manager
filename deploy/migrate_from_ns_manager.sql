-- ============================================================
-- 从 ns_manager 迁移数据到 ns_manager_v2
-- 执行方式: mysql -u root -p --default-character-set=utf8mb4 < migrate_from_ns_manager.sql
-- 说明：
--   1. 补全 ns_manager_v2 中缺失的表注释
--   2. 迁移所有业务数据，去掉 tenant_id 相关列
--   3. 处理 keyword 重复冲突（field/oauther 各有一对重复）
--   4. 排除多租户相关菜单（14条 tenant API 菜单及其 closure）
--   5. user.dept_id/job_id 迁移至 user_dept 表
-- ============================================================

USE ns_manager_v2;

-- ------------------------------------------------------------
-- 补全缺失的表注释
-- ------------------------------------------------------------
ALTER TABLE `casbin_rule` COMMENT='Casbin权限规则表';
ALTER TABLE `user_dept` COMMENT='用户部门表';

-- ============================================================
-- 参照数据（无依赖，先迁）
-- ============================================================

-- app（无 tenant_id，直接复制）
INSERT INTO ns_manager_v2.`app`
SELECT id, type, logo, favicon, keyword, secret, name, show_name, private, status, reason, description, comment, setting, created_at, updated_at
FROM ns_manager.`app`;

-- field（keyword 重复处理：phone 有两条，保留 id=3 的 tenant_id=1 版本，跳过 id=1 的全局版本）
INSERT INTO ns_manager_v2.`field` (id, keyword, type, name, status, required, `unique`, description, created_at, updated_at)
SELECT id, keyword, type, name, status, required, `unique`, description, created_at, updated_at
FROM ns_manager.`field`
WHERE id NOT IN (1)  -- 跳过 id=1（phone keyword 重复，保留 id=3）
ORDER BY id;

-- oauther（keyword 重复处理：yiban 有两条，保留 id=4 的 tenant_id=1 版本，跳过 id=1 的全局版本）
INSERT INTO ns_manager_v2.`oauther` (id, logo, keyword, name, type, ak, sk, description, status, setting, created_at, updated_at)
SELECT id, logo, keyword, name, type, ak, sk, description, status, setting, created_at, updated_at
FROM ns_manager.`oauther`
WHERE id NOT IN (1)  -- 跳过 id=1（yiban keyword 重复，保留 id=4）
ORDER BY id;

-- dept_classify（去掉 tenant_id）
INSERT INTO ns_manager_v2.`dept_classify` (id, keyword, name, description, status, weight, created_at, updated_at)
SELECT id, keyword, name, description, status, weight, created_at, updated_at
FROM ns_manager.`dept_classify`;

-- dept（去掉 tenant_id）
INSERT INTO ns_manager_v2.`dept` (id, classify_id, parent_id, keyword, name, logo, status, description, created_at, updated_at)
SELECT id, classify_id, parent_id, keyword, name, logo, status, description, created_at, updated_at
FROM ns_manager.`dept`;

-- dept_closure（无变化）
INSERT INTO ns_manager_v2.`dept_closure` (id, parent, children)
SELECT id, parent, children
FROM ns_manager.`dept_closure`;

-- job（去掉 tenant_id）
INSERT INTO ns_manager_v2.`job` (id, keyword, name, status, weight, description, created_at, updated_at)
SELECT id, keyword, name, status, weight, description, created_at, updated_at
FROM ns_manager.`job`;

-- role（去掉 tenant_id）
INSERT INTO ns_manager_v2.`role` (id, parent_id, keyword, name, status, description, created_at, updated_at)
SELECT id, parent_id, keyword, name, status, description, created_at, updated_at
FROM ns_manager.`role`;

-- role_closure（无变化）
INSERT INTO ns_manager_v2.`role_closure` (id, parent, children)
SELECT id, parent, children
FROM ns_manager.`role_closure`;

-- dictionary（去掉 tenant_id）
INSERT INTO ns_manager_v2.`dictionary` (id, keyword, type, name, description, created_at, updated_at)
SELECT id, keyword, type, name, description, created_at, updated_at
FROM ns_manager.`dictionary`;

-- dictionary_value（无 tenant_id）
INSERT INTO ns_manager_v2.`dictionary_value` (id, dictionary_id, parent_id, label, value, status, weight, type, extra, description, created_at, updated_at)
SELECT id, dictionary_id, parent_id, label, value, status, weight, type, extra, description, created_at, updated_at
FROM ns_manager.`dictionary_value`;

-- feedback_classify（去掉 tenant_id）
INSERT INTO ns_manager_v2.`feedback_classify` (id, name, created_at, updated_at)
SELECT id, name, created_at, updated_at
FROM ns_manager.`feedback_classify`;

-- notice_classify（去掉 tenant_id）
INSERT INTO ns_manager_v2.`notice_classify` (id, name, logo, weight, created_at, updated_at)
SELECT id, name, logo, weight, created_at, updated_at
FROM ns_manager.`notice_classify`;

-- entity（无 tenant_id，直接复制）
INSERT INTO ns_manager_v2.`entity` (id, app_id, `database`, name, comment, created_at, updated_at)
SELECT id, app_id, `database`, name, comment, created_at, updated_at
FROM ns_manager.`entity`;

-- entity_field（无 tenant_id，直接复制）
INSERT INTO ns_manager_v2.`entity_field` (id, entity_id, name, comment, `index`, created_at, updated_at)
SELECT id, entity_id, name, comment, `index`, created_at, updated_at
FROM ns_manager.`entity_field`;

-- entity_rule（去掉 tenant_id）
INSERT INTO ns_manager_v2.`entity_rule` (id, entity_id, name, description, expression, status, created_at, updated_at)
SELECT id, entity_id, name, description, expression, status, created_at, updated_at
FROM ns_manager.`entity_rule`;

-- ============================================================
-- 菜单数据（排除 tenant 相关菜单，共 14 条）
-- ============================================================

-- menu（排除租户相关菜单：ids 2,32,33,34,35,37,38,39,40,41,43,44,45,46）
INSERT INTO ns_manager_v2.`menu`
  (id, app_id, parent_id, title, type, keyword, icon, api, method, path, permission, component, url, redirect, weight, is_iframe, is_hidden, is_cache, is_home, is_affix, created_at, updated_at)
SELECT
  id, app_id, parent_id, title, type, keyword, icon, api, method, path, permission, component, url, redirect, weight, is_iframe, is_hidden, is_cache, is_home, is_affix, created_at, updated_at
FROM ns_manager.`menu`
WHERE id NOT IN (2, 32, 33, 34, 35, 37, 38, 39, 40, 41, 43, 44, 45, 46);

-- menu_closure（排除引用了被过滤菜单的 closure 条目）
INSERT INTO ns_manager_v2.`menu_closure` (id, parent, children)
SELECT id, parent, children
FROM ns_manager.`menu_closure`
WHERE parent NOT IN (2, 32, 33, 34, 35, 37, 38, 39, 40, 41, 43, 44, 45, 46)
  AND children NOT IN (2, 32, 33, 34, 35, 37, 38, 39, 40, 41, 43, 44, 45, 46);

-- ============================================================
-- 关联/权限数据
-- ============================================================

-- app_field（去掉 tenant_id；只保留 field_id 存在于新库的记录）
INSERT INTO ns_manager_v2.`app_field` (id, app_id, field_id, required, created_at)
SELECT id, app_id, field_id, required, created_at
FROM ns_manager.`app_field`
WHERE field_id IN (SELECT id FROM ns_manager_v2.`field`);

-- app_oauther（去掉 tenant_id；只保留 oauther_id 存在于新库的记录）
INSERT INTO ns_manager_v2.`app_oauther` (id, app_id, oauther_id, type, created_at)
SELECT id, app_id, oauther_id, type, created_at
FROM ns_manager.`app_oauther`
WHERE oauther_id IN (SELECT id FROM ns_manager_v2.`oauther`);

-- role_menu（排除引用了被过滤菜单的条目）
INSERT INTO ns_manager_v2.`role_menu` (id, role_id, menu_id)
SELECT id, role_id, menu_id
FROM ns_manager.`role_menu`
WHERE menu_id IN (SELECT id FROM ns_manager_v2.`menu`);

-- dept_role（无变化）
INSERT INTO ns_manager_v2.`dept_role` (id, role_id, dept_id)
SELECT id, role_id, dept_id
FROM ns_manager.`dept_role`;

-- job_role（无变化）
INSERT INTO ns_manager_v2.`job_role` (id, role_id, job_id)
SELECT id, role_id, job_id
FROM ns_manager.`job_role`;

-- role_entity（无 tenant_id）
INSERT INTO ns_manager_v2.`role_entity` (id, role_id, entity_id, action, scope, depts, fields, rules, created_at, updated_at)
SELECT id, role_id, entity_id, action, scope, depts, fields, rules, created_at, updated_at
FROM ns_manager.`role_entity`;

-- casbin_rule（无变化）
INSERT INTO ns_manager_v2.`casbin_rule` (id, ptype, v0, v1, v2, v3, v4, v5)
SELECT id, ptype, v0, v1, v2, v3, v4, v5
FROM ns_manager.`casbin_rule`;

-- ============================================================
-- 用户数据
-- ============================================================

-- user（去掉 tenant_id / dept_id / job_id，这两列迁移至 user_dept）
INSERT INTO ns_manager_v2.`user` (id, avatar, nickname, signature, username, password, status, reason, logged_at, created_at, updated_at, deleted_at)
SELECT id, avatar, nickname, signature, username, password, status, reason, logged_at, created_at, updated_at, deleted_at
FROM ns_manager.`user`;

-- user_dept（从旧 user 表的 dept_id/job_id 列迁移，跳过 dept_id 或 job_id 为 NULL 的用户）
INSERT INTO ns_manager_v2.`user_dept` (user_id, dept_id, main, job_id, created_at, updated_at)
SELECT id, dept_id, 1, job_id, created_at, updated_at
FROM ns_manager.`user`
WHERE dept_id IS NOT NULL AND job_id IS NOT NULL;

-- user_oauther（去掉 tenant_id）
INSERT INTO ns_manager_v2.`user_oauther` (id, user_id, oauther_id, oid, token, logged_at, expired_at, extra, created_at, updated_at)
SELECT id, user_id, oauther_id, oid, token, logged_at, expired_at, extra, created_at, updated_at
FROM ns_manager.`user_oauther`
WHERE oauther_id IN (SELECT id FROM ns_manager_v2.`oauther`);

-- authorize（去掉 tenant_id）
INSERT INTO ns_manager_v2.`authorize` (id, app_id, user_id, tokens, logged_at, expired_at, created_at)
SELECT id, app_id, user_id, tokens, logged_at, expired_at, created_at
FROM ns_manager.`authorize`;

-- ============================================================
-- userinfo 分片表（去掉 tenant_id；同一 user_id+field 在旧库可能有 tenant_id=0 和 tenant_id=1 各一条，用 INSERT IGNORE 保留先插入的）
-- ============================================================
INSERT IGNORE INTO ns_manager_v2.`userinfo_0` (id, user_id, field, value, value_md5, created_at, updated_at)
SELECT id, user_id, field, value, value_md5, created_at, updated_at FROM ns_manager.`userinfo_0`;

INSERT IGNORE INTO ns_manager_v2.`userinfo_1` (id, user_id, field, value, value_md5, created_at, updated_at)
SELECT id, user_id, field, value, value_md5, created_at, updated_at FROM ns_manager.`userinfo_1` ORDER BY tenant_id DESC;

INSERT IGNORE INTO ns_manager_v2.`userinfo_2` (id, user_id, field, value, value_md5, created_at, updated_at)
SELECT id, user_id, field, value, value_md5, created_at, updated_at FROM ns_manager.`userinfo_2`;

INSERT IGNORE INTO ns_manager_v2.`userinfo_3` (id, user_id, field, value, value_md5, created_at, updated_at)
SELECT id, user_id, field, value, value_md5, created_at, updated_at FROM ns_manager.`userinfo_3`;

INSERT IGNORE INTO ns_manager_v2.`userinfo_4` (id, user_id, field, value, value_md5, created_at, updated_at)
SELECT id, user_id, field, value, value_md5, created_at, updated_at FROM ns_manager.`userinfo_4`;

INSERT IGNORE INTO ns_manager_v2.`userinfo_5` (id, user_id, field, value, value_md5, created_at, updated_at)
SELECT id, user_id, field, value, value_md5, created_at, updated_at FROM ns_manager.`userinfo_5`;

INSERT IGNORE INTO ns_manager_v2.`userinfo_6` (id, user_id, field, value, value_md5, created_at, updated_at)
SELECT id, user_id, field, value, value_md5, created_at, updated_at FROM ns_manager.`userinfo_6`;

INSERT IGNORE INTO ns_manager_v2.`userinfo_7` (id, user_id, field, value, value_md5, created_at, updated_at)
SELECT id, user_id, field, value, value_md5, created_at, updated_at FROM ns_manager.`userinfo_7`;

INSERT IGNORE INTO ns_manager_v2.`userinfo_8` (id, user_id, field, value, value_md5, created_at, updated_at)
SELECT id, user_id, field, value, value_md5, created_at, updated_at FROM ns_manager.`userinfo_8`;

INSERT IGNORE INTO ns_manager_v2.`userinfo_9` (id, user_id, field, value, value_md5, created_at, updated_at)
SELECT id, user_id, field, value, value_md5, created_at, updated_at FROM ns_manager.`userinfo_9`;

-- ============================================================
-- 日志数据（去掉 tenant_id）
-- ============================================================

-- login_log（去掉 tenant_id）
INSERT INTO ns_manager_v2.`login_log` (id, app_id, type, ip, address, device, browser, code, description, dept_id, user_id, created_at)
SELECT id, app_id, type, ip, address, device, browser, code, description, dept_id, user_id, created_at
FROM ns_manager.`login_log`;

-- auth_log（去掉 tenant_id；排除引用被过滤菜单的记录）
INSERT INTO ns_manager_v2.`auth_log` (id, app_id, menu_id, dept_id, user_id, created_at)
SELECT id, app_id, menu_id, dept_id, user_id, created_at
FROM ns_manager.`auth_log`
WHERE menu_id IN (SELECT id FROM ns_manager_v2.`menu`);

-- notice（去掉 tenant_id）
INSERT INTO ns_manager_v2.`notice` (id, app_id, classify_id, title, description, unit, content, is_top, status, created_at, updated_at)
SELECT id, app_id, classify_id, title, description, unit, content, is_top, status, created_at, updated_at
FROM ns_manager.`notice`;

-- notice_user（无 tenant_id）
INSERT INTO ns_manager_v2.`notice_user` (id, notice_id, user_id, created_at)
SELECT id, notice_id, user_id, created_at
FROM ns_manager.`notice_user`;

-- feedback（去掉 tenant_id）
INSERT INTO ns_manager_v2.`feedback` (id, app_id, user_id, classify_id, title, content, status, images, contact, device, platform, version, md5, processed_by, processed_result, created_at, updated_at)
SELECT id, app_id, user_id, classify_id, title, content, status, images, contact, device, platform, version, md5, processed_by, processed_result, created_at, updated_at
FROM ns_manager.`feedback`;

-- ============================================================
-- 验证行数
-- ============================================================
SELECT 'app'             AS tbl, COUNT(*) AS old_rows FROM ns_manager.app
UNION ALL SELECT 'app',            COUNT(*) FROM ns_manager_v2.app
UNION ALL SELECT 'field_old',      COUNT(*) FROM ns_manager.field
UNION ALL SELECT 'field_new',      COUNT(*) FROM ns_manager_v2.field
UNION ALL SELECT 'oauther_old',    COUNT(*) FROM ns_manager.oauther
UNION ALL SELECT 'oauther_new',    COUNT(*) FROM ns_manager_v2.oauther
UNION ALL SELECT 'user_old',       COUNT(*) FROM ns_manager.user
UNION ALL SELECT 'user_new',       COUNT(*) FROM ns_manager_v2.user
UNION ALL SELECT 'menu_old',       COUNT(*) FROM ns_manager.menu
UNION ALL SELECT 'menu_new',       COUNT(*) FROM ns_manager_v2.menu
UNION ALL SELECT 'role_old',       COUNT(*) FROM ns_manager.role
UNION ALL SELECT 'role_new',       COUNT(*) FROM ns_manager_v2.role;
