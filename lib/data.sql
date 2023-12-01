
BEGIN;
INSERT INTO `menu` VALUES (1, 'Manager', 0, '管理平台', 'R', 'Manager', 'apps', '/manager', NULL, 'Layout', NULL, NULL, NULL, 0, 0, NULL, NULL, NULL, 1698809470, 1698988713);
INSERT INTO `menu` VALUES (2, 'Manager', 1, '基础接口', 'G', 'BaseApi', 'apps', NULL, 'BA', NULL, NULL, NULL, NULL, 0, 1, NULL, NULL, NULL, 1699066881, 1699098618);
INSERT INTO `menu` VALUES (3, 'Manager', 1, '菜单管理', 'M', 'ManagerMenu', 'menu', '/manager/menu', NULL, 'manager/menu/index', NULL, NULL, NULL, 0, 0, 1, 1, 1, 1698809470, 1698809470);
INSERT INTO `menu` VALUES (4, 'Manager', 1, '部门管理', 'M', 'ManagerDepartment', 'user-group', '/manager/department', NULL, 'manager/department/index', NULL, NULL, NULL, 0, 0, 1, NULL, 1, 1698809470, 1698988624);
INSERT INTO `menu` VALUES (5, 'Manager', 1, '角色管理', 'M', 'ManagerRole', 'safe', '/manager/role', NULL, 'manager/role/index', NULL, NULL, NULL, 0, NULL, 0, NULL, 1, 1698816203, 1698988607);
INSERT INTO `menu` VALUES (6, 'Manager', 1, '用户管理', 'M', 'ManagerUser', 'user', '/manager/user', NULL, 'manager/user/index', NULL, NULL, NULL, 0, NULL, 1, NULL, 1, 1698893392, 1698893392);
INSERT INTO `menu` VALUES (7, 'Configure', 0, '配置中心', 'R', 'Configure', 'code-block', '/configure', NULL, 'Layout', NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL, 1698979506, 1698979506);
INSERT INTO `menu` VALUES (8, 'Configure', 7, '环境管理', 'M', 'ConfigureEnv', 'common', '/configure/env', NULL, 'configure/env/index', NULL, NULL, NULL, 0, NULL, 1, NULL, 1, 1698986828, 1699159819);
INSERT INTO `menu` VALUES (9, 'Manager', 3, '查询菜单', 'A', NULL, NULL, '', 'manager:menu:query', NULL, NULL, '/manager/v1/menu/tree', 'GET', 0, NULL, NULL, NULL, NULL, 1699066727, 1699088143);
INSERT INTO `menu` VALUES (12, 'Manager', 3, '新增菜单', 'A', NULL, NULL, NULL, 'manager:menu:add', NULL, NULL, '/manager/v1/menu', 'POST', 0, NULL, NULL, NULL, NULL, 1699067220, 1699067220);
INSERT INTO `menu` VALUES (13, 'Manager', 3, '修改菜单', 'A', NULL, NULL, NULL, 'manager:menu:update', NULL, NULL, '/manager/v1/menu', 'PUT', 0, NULL, NULL, NULL, NULL, 1699067245, 1699067245);
INSERT INTO `menu` VALUES (14, 'Manager', 3, '删除菜单', 'A', NULL, NULL, NULL, 'manager:menu:delete', NULL, NULL, '/manager/v1/menu', 'DELETE', 0, NULL, NULL, NULL, NULL, 1699067278, 1699067278);
INSERT INTO `menu` VALUES (15, 'Manager', 4, '查看部门', 'A', NULL, NULL, NULL, 'manager:department:query', NULL, NULL, '/manager/v1/department/tree', 'GET', 0, NULL, NULL, NULL, NULL, 1699090258, 1699090289);
INSERT INTO `menu` VALUES (16, 'Manager', 4, '新增部门', 'A', NULL, NULL, NULL, 'manager:department:add', NULL, NULL, '/manager/v1/department', 'POST', 0, NULL, NULL, NULL, NULL, 1699090283, 1699090283);
INSERT INTO `menu` VALUES (17, 'Manager', 4, '修改部门', 'A', NULL, NULL, NULL, 'manager:department:update', NULL, NULL, '/manager/v1/department', 'PUT', 0, NULL, NULL, NULL, NULL, 1699090316, 1699090316);
INSERT INTO `menu` VALUES (18, 'Manager', 4, '删除部门', 'A', NULL, NULL, NULL, 'manager:department:delete', NULL, NULL, '/manager/v1/department', 'DELETE', 0, NULL, NULL, NULL, NULL, 1699090340, 1699090340);
INSERT INTO `menu` VALUES (19, 'Manager', 5, '查询角色', 'A', NULL, NULL, NULL, 'manager:role:query', NULL, NULL, '/manager/v1/role/tree', 'GET', 0, NULL, NULL, NULL, NULL, 1699090425, 1699090425);
INSERT INTO `menu` VALUES (20, 'Manager', 5, '新增角色', 'A', NULL, NULL, NULL, 'manager:role:add', NULL, NULL, '/manager/v1/role', 'POST', 0, NULL, NULL, NULL, NULL, 1699090629, 1699090629);
INSERT INTO `menu` VALUES (21, 'Manager', 5, '修改角色', 'A', NULL, NULL, NULL, 'manager:role:update', NULL, NULL, '/manager/v1/role', 'PUT', 0, NULL, NULL, NULL, NULL, 1699090656, 1699090656);
INSERT INTO `menu` VALUES (22, 'Manager', 5, '删除角色', 'A', NULL, NULL, NULL, 'manager:role:delete', NULL, NULL, '/manager/v1/role', 'DELETE', 0, NULL, NULL, NULL, NULL, 1699090676, 1699090676);
INSERT INTO `menu` VALUES (23, 'Manager', 5, '角色菜单管理', 'G', NULL, NULL, NULL, 'manager:role:menu', NULL, NULL, '/manager/v1/role/menu/tree', 'GET', 0, NULL, NULL, NULL, NULL, 1699090828, 1699090896);
INSERT INTO `menu` VALUES (24, 'Manager', 23, '查询角色菜单', 'A', NULL, NULL, NULL, 'manager:role:menu:query', NULL, NULL, '/manager/v1/role/menu/ids', 'GET', 0, NULL, NULL, NULL, NULL, 1699090923, 1699108809);
INSERT INTO `menu` VALUES (25, 'Manager', 23, '修改角色菜单', 'A', NULL, NULL, NULL, 'manager:role:menu:update', NULL, NULL, '/manager/v1/role/menu', 'POST', 0, NULL, NULL, NULL, NULL, 1699091405, 1699091405);
INSERT INTO `menu` VALUES (26, 'Manager', 6, '查询用户', 'A', NULL, NULL, NULL, 'manager:user:query', NULL, NULL, '/manager/v1/users', 'GET', 0, NULL, NULL, NULL, NULL, 1699091944, 1699091944);
INSERT INTO `menu` VALUES (27, 'Manager', 6, '新增用户', 'A', NULL, NULL, NULL, 'manager:user:add', NULL, NULL, '/manager/v1/user', 'POST', 0, NULL, NULL, NULL, NULL, 1699092466, 1699092476);
INSERT INTO `menu` VALUES (28, 'Manager', 6, '修改用户', 'A', NULL, NULL, NULL, 'manager:user:update', NULL, NULL, '/manager/v1/user', 'PUT', 0, NULL, NULL, NULL, NULL, 1699092495, 1699092495);
INSERT INTO `menu` VALUES (29, 'Manager', 6, '删除用户', 'A', NULL, NULL, NULL, 'manager:user:delete', NULL, NULL, '/manager/v1/user', 'DELETE', 0, NULL, NULL, NULL, NULL, 1699092515, 1699092515);
INSERT INTO `menu` VALUES (30, 'Manager', 6, '启用/禁用用户', 'G', NULL, NULL, NULL, 'manager:user:status', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL, 1699092540, 1699092540);
INSERT INTO `menu` VALUES (31, 'Manager', 30, '启用用户', 'A', NULL, NULL, NULL, 'manager:user:status:enable', NULL, NULL, '/manager/v1/user/enable', 'POST', 0, NULL, NULL, NULL, NULL, 1699092863, 1699092863);
INSERT INTO `menu` VALUES (32, 'Manager', 30, '禁用用户', 'A', NULL, NULL, NULL, 'manager:user:status:disable', NULL, NULL, '/manager/v1/user/disable', 'POST', 0, NULL, NULL, NULL, NULL, 1699092902, 1699092902);
INSERT INTO `menu` VALUES (33, 'Manager', 6, '重置账号密码', 'A', NULL, NULL, NULL, 'manager:user:reset:password', '', NULL, '/manager/v1/user/password/reset', 'POST', 0, NULL, NULL, NULL, NULL, 1699093021, 1699107168);
INSERT INTO `menu` VALUES (34, 'Manager', 2, '获取用户部门', 'BA', NULL, NULL, NULL, NULL, NULL, NULL, '/manager/v1/department/tree', 'GET', 0, NULL, NULL, NULL, NULL, 1699095870, 1699095870);
INSERT INTO `menu` VALUES (35, 'Manager', 2, '获取个人用户信息', 'BA', NULL, NULL, NULL, NULL, NULL, NULL, '/manager/v1/user/current', 'GET', 0, NULL, NULL, NULL, NULL, 1699095899, 1699095899);
INSERT INTO `menu` VALUES (36, 'Manager', 2, '获取用户可见角色树', 'BA', NULL, NULL, NULL, NULL, NULL, NULL, '/manager/v1/role/tree', 'GET', 0, NULL, NULL, NULL, NULL, 1699095965, 1699109755);
INSERT INTO `menu` VALUES (37, 'Manager', 2, '获取用户菜单', 'BA', NULL, NULL, NULL, NULL, NULL, NULL, '/manager/v1/role/menu/tree', 'GET', 0, NULL, NULL, NULL, NULL, 1699096007, 1699096007);
INSERT INTO `menu` VALUES (38, 'Manager', 2, '退出登录', 'BA', NULL, NULL, NULL, NULL, NULL, NULL, '/manager/v1/logout', 'POST', 0, NULL, NULL, NULL, NULL, 1699096034, 1699096034);
INSERT INTO `menu` VALUES (39, 'Manager', 2, '刷新token', 'BA', NULL, NULL, NULL, NULL, NULL, NULL, '/manager/v1/token/refresh', 'POST', 0, NULL, NULL, NULL, NULL, 1699096056, 1699096056);
INSERT INTO `menu` VALUES (40, 'Manager', 2, '用户登录', 'BA', NULL, NULL, NULL, NULL, NULL, NULL, '/manager/v1/login', 'POST', 0, NULL, NULL, NULL, NULL, 1699108346, 1699108346);
INSERT INTO `menu` VALUES (41, 'Manager', 2, '获取登录验证码', 'BA', NULL, NULL, NULL, NULL, NULL, NULL, '/manager/v1/login/captcha', 'POST', 0, NULL, NULL, NULL, NULL, 1699108367, 1699108367);
INSERT INTO `menu` VALUES (43, 'Manager', 28, '获取用户角色列表', 'A', NULL, NULL, NULL, 'manager:user:update', NULL, NULL, '/manager/v1/user/roles', 'GET', 0, NULL, NULL, NULL, NULL, 1699110046, 1699110046);
INSERT INTO `menu` VALUES (44, 'Configure', 8, '查看环境', 'A', NULL, NULL, NULL, 'configure:environment:query', NULL, NULL, '/configure/v1/environments', 'GET', 0, NULL, NULL, NULL, NULL, 1699157386, 1699158860);
INSERT INTO `menu` VALUES (45, 'Configure', 8, '新增环境', 'A', NULL, NULL, NULL, 'configure:environment:add', NULL, NULL, '/configure/v1/environment', 'POST', 0, NULL, NULL, NULL, NULL, 1699157558, 1699157558);
INSERT INTO `menu` VALUES (46, 'Configure', 8, '修改环境', 'A', NULL, NULL, NULL, 'configure:environment:update', NULL, NULL, '/configure/v1/environment', 'PUT', 0, NULL, NULL, NULL, NULL, 1699157828, 1699157828);
INSERT INTO `menu` VALUES (47, 'Configure', 8, '删除环境', 'A', NULL, NULL, NULL, 'configure:environment:delete', NULL, NULL, '/configure/v1/environment', 'DELETE', 0, NULL, NULL, NULL, NULL, 1699157963, 1699157963);
INSERT INTO `menu` VALUES (48, 'Configure', 8, '查看环境Token', 'A', NULL, NULL, NULL, 'configure:environment:token:query', NULL, NULL, '/configure/v1/environment/token', 'GET', 0, NULL, NULL, NULL, NULL, 1699158028, 1699158464);
INSERT INTO `menu` VALUES (49, 'Configure', 8, '重置环境token', 'A', NULL, NULL, NULL, 'configure:environment:token:reset', NULL, NULL, '/configure/v1/environment/token', 'PUT', 0, NULL, NULL, NULL, NULL, 1699158347, 1699158355);
INSERT INTO `menu` VALUES (50, 'Configure', 7, '服务管理', 'M', 'ConfigureServer', 'apps', '/configure/server', NULL, 'configure/server/index', NULL, NULL, NULL, 0, NULL, 1, NULL, 1, 1699159786, 1699159786);
INSERT INTO `menu` VALUES (51, 'Configure', 50, '查询服务', 'A', NULL, NULL, NULL, 'configure:server:query', NULL, NULL, '/configure/v1/servers', 'GET', 0, NULL, NULL, NULL, NULL, 1699185102, 1699185102);
INSERT INTO `menu` VALUES (52, 'Configure', 50, '新增服务', 'A', NULL, NULL, NULL, 'configure:server:add', NULL, NULL, '/configure/v1/server', 'POST', 0, NULL, NULL, NULL, NULL, 1699185138, 1699185138);
INSERT INTO `menu` VALUES (53, 'Configure', 50, '修改服务', 'A', NULL, NULL, NULL, 'configure:server:update', NULL, NULL, '/configure/v1/server', 'PUT', 0, NULL, NULL, NULL, NULL, 1699185164, 1699185164);
INSERT INTO `menu` VALUES (54, 'Configure', 50, '删除服务', 'A', NULL, NULL, NULL, 'configure:server:delete', NULL, NULL, '/configure/v1/server', 'DELETE', 0, NULL, NULL, NULL, NULL, 1699185208, 1699185208);
INSERT INTO `menu` VALUES (55, 'Configure', 7, '资源变量', 'M', 'ConfigureResource', 'unordered-list', '/configure/resource', NULL, 'configure/resource/index', NULL, NULL, NULL, 0, NULL, 1, NULL, 1, 1699186200, 1699186200);
INSERT INTO `menu` VALUES (56, 'Configure', 55, '查看资源', 'A', NULL, NULL, NULL, 'configure:resource:query', NULL, NULL, '/configure/v1/resources', 'GET', 0, NULL, NULL, NULL, NULL, 1699186608, 1699186608);
INSERT INTO `menu` VALUES (57, 'Configure', 55, '新增资源', 'A', NULL, NULL, NULL, 'configure:resource:add', NULL, NULL, '/configure/v1/resource', 'POST', 0, NULL, NULL, NULL, NULL, 1699186633, 1699186633);
INSERT INTO `menu` VALUES (58, 'Configure', 55, '修改资源', 'A', NULL, NULL, NULL, 'configure:resource:update', NULL, NULL, '/configure/v1/resource', 'PUT', 0, NULL, NULL, NULL, NULL, 1699186656, 1699186656);
INSERT INTO `menu` VALUES (59, 'Configure', 55, '删除资源', 'A', NULL, NULL, NULL, 'configure:resource:delete', NULL, NULL, '/configure/v1/resource', 'DELETE', 0, NULL, NULL, NULL, NULL, 1699186676, 1699186676);
INSERT INTO `menu` VALUES (60, 'Configure', 55, '资源变量值配置', 'G', NULL, NULL, NULL, 'configure:resource:value', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL, 1699248053, 1699248053);
INSERT INTO `menu` VALUES (61, 'Configure', 7, '业务变量', 'M', 'ConfigureBusiness', 'code', '/configure/business', NULL, 'configure/business/index', NULL, NULL, NULL, 0, NULL, 1, NULL, 1, 1699373901, 1699413051);
INSERT INTO `menu` VALUES (62, 'Configure', 61, '查看业务变量', 'A', NULL, NULL, NULL, 'configure:business:query', NULL, NULL, '/configure/v1/business', 'GET', 0, NULL, NULL, NULL, NULL, 1699413104, 1699413104);
INSERT INTO `menu` VALUES (63, 'Configure', 61, '新增业务变量', 'A', NULL, NULL, NULL, 'configure:business:add', NULL, NULL, '/configure/v1/business', 'POST', 0, NULL, NULL, NULL, NULL, 1699418799, 1699418799);
INSERT INTO `menu` VALUES (64, 'Configure', 61, '修改业务变量', 'A', NULL, NULL, NULL, 'configure:business:update', NULL, NULL, '/configure/v1/business', 'PUT', 0, NULL, NULL, NULL, NULL, 1699418856, 1699418864);
INSERT INTO `menu` VALUES (65, 'Configure', 61, '删除业务变量', 'A', NULL, NULL, NULL, 'configure:business:delete', NULL, NULL, '/configure/v1/business', 'DELETE', 0, NULL, NULL, NULL, NULL, 1699418892, 1699418892);
INSERT INTO `menu` VALUES (66, 'Configure', 7, '配置模板', 'M', 'ConfgureTemplate', 'file', '/configure/template', NULL, 'configure/template/index', NULL, NULL, NULL, 0, NULL, 1, NULL, 1, 1699458029, 1699458029);
INSERT INTO `menu` VALUES (67, 'Configure', 66, '配置预览', 'G', NULL, NULL, NULL, 'configure:template:preview', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL, 1699608352, 1699608458);
INSERT INTO `menu` VALUES (68, 'Configure', 66, '提交模板', 'A', NULL, NULL, NULL, 'configure:template:add', NULL, NULL, '/configure/v1/template', 'POST', 0, NULL, NULL, NULL, NULL, 1699608517, 1699608517);
INSERT INTO `menu` VALUES (69, 'Configure', 66, '同步配置', 'G', NULL, NULL, NULL, 'configure:configure:sync', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL, 1699608570, 1699608570);
INSERT INTO `menu` VALUES (70, 'Configure', 67, '测试环境', 'A', NULL, NULL, NULL, 'configure:template:preview:TEST', NULL, NULL, '/configure/v1/template/preview/TEST', 'POST', 0, NULL, NULL, NULL, NULL, 1699609506, 1699615583);
INSERT INTO `menu` VALUES (71, 'Configure', 67, '预发布环境', 'A', NULL, NULL, NULL, 'configure:template:preview:PRE', NULL, NULL, '/configure/v1/template/preview/PRE', 'POST', 0, NULL, NULL, NULL, NULL, 1699609570, 1699615595);
INSERT INTO `menu` VALUES (72, 'Configure', 67, '回归测试环境', 'A', NULL, NULL, NULL, 'configure:template:preview:BETA', NULL, NULL, '/configure/v1/template/preview/BETA', 'POST', 0, NULL, NULL, NULL, NULL, 1699609609, 1699615607);
INSERT INTO `menu` VALUES (73, 'Configure', 67, '生产环境', 'A', NULL, NULL, NULL, 'configure:template:preview:PROD', NULL, NULL, '/configure/v1/template/preview/PROD', 'POST', 0, NULL, NULL, NULL, NULL, 1699609649, 1699615618);
INSERT INTO `menu` VALUES (74, 'Configure', 69, '测试环境', 'A', NULL, NULL, NULL, 'configure:configure:sync:TEST', NULL, NULL, '/configure/v1/configure/TEST', 'PUT', 0, NULL, NULL, NULL, NULL, 1699609728, 1699615634);
INSERT INTO `menu` VALUES (75, 'Configure', 69, '预发布环境', 'A', NULL, NULL, NULL, 'configure:configure:sync:PRE', NULL, NULL, '/configure/v1/configure/PRE', 'PUT', 0, NULL, NULL, NULL, NULL, 1699609761, 1699615644);
INSERT INTO `menu` VALUES (76, 'Configure', 69, '回归测试环境', 'A', NULL, NULL, NULL, 'configure:configure:sync:BETA', NULL, NULL, '/configure/v1/configure/BETA', 'PUT', 0, NULL, NULL, NULL, NULL, 1699609789, 1699615655);
INSERT INTO `menu` VALUES (77, 'Configure', 69, '生产环境', 'A', NULL, NULL, NULL, 'configure:configure:sync:PROD', NULL, NULL, '/configure/v1/configure/PROD', 'PUT', 0, NULL, NULL, NULL, NULL, 1699609871, 1699615673);
INSERT INTO `menu` VALUES (78, 'Configure', 61, '配置业务变量值', 'A', NULL, NULL, NULL, 'configure:business:value', NULL, NULL, '/configure/business/value', 'PUT', 0, NULL, NULL, NULL, NULL, 1699617844, 1699617844);
INSERT INTO `menu` VALUES (79, 'Resource', 0, '资源中心', 'R', 'Resource', 'file', '/resource', NULL, 'Layout', NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL, 1700744316, 1700744316);
INSERT INTO `menu` VALUES (80, 'Resource', 79, '文件管理', 'M', 'File', 'file', '/resource/file', NULL, 'resource/file/index', NULL, NULL, NULL, 0, 0, 1, NULL, 1, 1700744655, 1700790343);
INSERT INTO `menu` VALUES (81, 'Resource', 80, '目录管理', 'G', NULL, NULL, NULL, 'resource:directory:group', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL, 1701420127, 1701420545);
INSERT INTO `menu` VALUES (82, 'Resource', 81, '查看目录', 'A', NULL, NULL, NULL, 'resource:directory:query', NULL, NULL, '/resource/v1/directory', 'GET', 0, NULL, NULL, NULL, NULL, 1701420313, 1701420313);
INSERT INTO `menu` VALUES (83, 'Resource', 81, '新增目录', 'A', NULL, NULL, NULL, 'resource:directory:add', NULL, NULL, '/resource/v1/directory', 'POST', 0, NULL, NULL, NULL, NULL, 1701420340, 1701420340);
INSERT INTO `menu` VALUES (84, 'Resource', 81, '修改目录', 'A', NULL, NULL, NULL, 'resource:directory:update', NULL, NULL, '/resource/v1/directory', 'PUT', 0, NULL, NULL, NULL, NULL, 1701420366, 1701420366);
INSERT INTO `menu` VALUES (85, 'Resource', 81, '删除目录', 'A', NULL, NULL, NULL, 'resource:directory:delete', NULL, NULL, '/resource/v1/directory', 'DELETE', 0, NULL, NULL, NULL, NULL, 1701420405, 1701420405);
INSERT INTO `menu` VALUES (86, 'Resource', 80, '文件管理', 'G', NULL, NULL, NULL, 'resource:file:group', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL, 1701420536, 1701420536);
INSERT INTO `menu` VALUES (87, 'Resource', 86, '查看文件', 'A', NULL, NULL, NULL, 'resource:file:query', NULL, NULL, '/resource/v1/files', 'GET', 0, NULL, NULL, NULL, NULL, 1701420717, 1701420717);
INSERT INTO `menu` VALUES (88, 'Resource', 92, '分块上传文件', 'A', NULL, NULL, NULL, 'resource:file:upload', NULL, NULL, '/resource/v1/upload', 'POST', 0, NULL, NULL, NULL, NULL, 1701420751, 1701421037);
INSERT INTO `menu` VALUES (89, 'Resource', 86, '修改文件', 'A', NULL, NULL, NULL, 'resource:file:upload', NULL, NULL, '/resource/v1/file', 'PUT', 0, NULL, NULL, NULL, NULL, 1701420801, 1701420801);
INSERT INTO `menu` VALUES (90, 'Resource', 81, '删除文件', 'A', NULL, NULL, NULL, 'resource:file:delete', NULL, NULL, '/resource/v1/file', 'DELETE', 0, NULL, NULL, NULL, NULL, 1701420836, 1701420836);
INSERT INTO `menu` VALUES (91, 'Resource', 86, '删除文件', 'A', NULL, NULL, NULL, 'resource:file:delete', NULL, NULL, '/resource/v1/file', 'DELETE', 0, NULL, NULL, NULL, NULL, 1701420980, 1701420980);
INSERT INTO `menu` VALUES (92, 'Resource', 86, '上传文件', 'G', NULL, NULL, NULL, 'resource:file:upload:group', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL, 1701421014, 1701421014);
INSERT INTO `menu` VALUES (93, 'Resource', 92, '预上传文件', 'A', NULL, NULL, NULL, 'resource:file:upload', NULL, NULL, '/resource/v1/upload/prepare', 'POST', 0, NULL, NULL, NULL, NULL, 1701421098, 1701421098);
COMMIT;

DOCKERFILE_PATH: Dockerfile
DOCKER_IMAGE_NAME:basic-platform
DOCKER_REPO_NAME: Dockerfile
DOCKER_IMAGE_VERSION:${GIT_LOCAL_BRANCH}-${GIT_COMMIT_SHORT}
DOCKER_BUILD_CONTEXT:.
REMOTE_HOST: 101.34.229.39
REMOTE_CRED:.
REMOTE_SSH_PORT:22
REMOTE_USER_NAME:root
DOCKER_RUN_PORT:7001
DOCKER_LISTEN_PORT:80
APP_VERSION:${GIT_COMMIT_SHORT}
APP_NAME:Resource
CONFIG_TOKEN:5B655B7D4A51BF79C974C9F27C27D992
CONFIG_HOST:localhost:6082