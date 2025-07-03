package dbs

import (
	"errors"
	"sync"
	"time"

	json "github.com/json-iterator/go"
	"github.com/limes-cloud/kratosx"
	"github.com/limes-cloud/kratosx/pkg/lock"
	"github.com/limes-cloud/kratosx/pkg/valx"

	"github.com/limes-cloud/manager/internal/domain/entity"
	"github.com/limes-cloud/manager/internal/types"
)

type Role struct {
}

type roleDataScope struct {
	scope string
	depId uint32
	data  []uint32
}

const (
	roleLockKV = "lock:role:kv"
	roleKV     = "get:role:kv"
)

var (
	roleIns  *Role
	roleOnce sync.Once
)

func NewRole() *Role {
	roleOnce.Do(func() {
		roleIns = &Role{}
	})
	return roleIns
}

func (infra *Role) delCache(ctx kratosx.Context) {
	// 删除缓存
	time.Sleep(1 * time.Second)
	ctx.Redis().Del(ctx, roleKV)
}

// ListCacheRole 获取指定的数据
func (infra *Role) ListCacheRole(ctx kratosx.Context) ([]*entity.Role, error) {
	var reply []*entity.Role
	locker := lock.New(ctx.Redis(), roleLockKV)
	err := locker.AcquireFunc(
		ctx,
		// 查询缓存
		func() error {
			res, err := ctx.Redis().Get(ctx, roleKV).Result()
			if err != nil {
				return err
			}
			if err := json.Unmarshal([]byte(res), &reply); err != nil {
				return err
			}
			return nil
		},
		// 从数据库读取，并设置到缓存
		func() error {
			if err := ctx.DB().
				Select([]string{
					"id", "parent_id", "name", "keyword", "data_scope", "department_ids", "job_scope", "job_ids",
				}).
				Where("status=true").
				Find(&reply).Error; err != nil {
				return err
			}

			// 序列化存储
			b, _ := json.Marshal(reply)

			return ctx.Redis().Set(ctx, roleKV, string(b), -1).Err()
		},
	)
	if err != nil {
		return nil, err
	}
	return reply, nil
}

// GetCacheRoleKeyMap 获取缓存的key映射
func (infra *Role) GetCacheRoleKeyMap(ctx kratosx.Context) (map[string]*entity.Role, error) {
	list, err := infra.ListCacheRole(ctx)
	if err != nil {
		return nil, err
	}
	var m = map[string]*entity.Role{}
	for _, item := range list {
		m[item.Keyword] = item
	}
	return m, nil
}

// GetCacheRoleIDMap 获取缓存的id映射
func (infra *Role) GetCacheRoleIDMap(ctx kratosx.Context) (map[uint32]*entity.Role, error) {
	list, err := infra.ListCacheRole(ctx)
	if err != nil {
		return nil, err
	}
	var m = map[uint32]*entity.Role{}
	for _, item := range list {
		m[item.Id] = item
	}
	return m, nil
}

// GetRole 获取指定的数据
func (infra *Role) GetRole(ctx kratosx.Context, id uint32) (*entity.Role, error) {
	var role = entity.Role{}
	return &role, ctx.DB().First(&role, id).Error
}

// GetRoleByKeyword 获取指定数据
func (infra *Role) GetRoleByKeyword(ctx kratosx.Context, keyword string) (*entity.Role, error) {
	var role = entity.Role{}
	return &role, ctx.DB().Where("keyword = ?", keyword).First(&role).Error
}

// ListRole 获取列表
func (infra *Role) ListRole(ctx kratosx.Context, req *types.ListRoleRequest) ([]*entity.Role, error) {
	var (
		es []*entity.Role
		fs = []string{"*"}
	)

	db := ctx.DB().Model(entity.Role{}).Select(fs)
	if req.Ids != nil {
		db = db.Where("id in ?", req.Ids)
	}
	if req.Name != nil {
		db = db.Where("name LIKE ?", "%"+*req.Name+"%")
	}
	if req.Keyword != nil {
		db = db.Where("keyword = ?", *req.Keyword)
	}

	return es, db.Find(&es).Error
}

// CreateRole 创建数据
func (infra *Role) CreateRole(ctx kratosx.Context, role *entity.Role) (uint32, error) {
	// 删除缓存
	defer infra.delCache(ctx)

	return role.Id, ctx.Transaction(func(ctx kratosx.Context) error {
		if err := ctx.DB().Create(role).Error; err != nil {
			return err
		}
		return infra.appendRoleChildren(ctx, role.ParentId, role.Id)
	})
}

// UpdateRole 更新数据
func (infra *Role) UpdateRole(ctx kratosx.Context, req *entity.Role) error {
	if req.Id == req.ParentId {
		return errors.New("父级不能为自己")
	}
	old, err := infra.GetRole(ctx, req.Id)
	if err != nil {
		return err
	}

	// 删除缓存
	defer infra.delCache(ctx)

	return ctx.Transaction(func(ctx kratosx.Context) error {
		if old.ParentId != req.ParentId {
			if err := infra.removeRoleParent(ctx, req.Id); err != nil {
				return err
			}
			if err := infra.appendRoleChildren(ctx, req.ParentId, req.Id); err != nil {
				return err
			}
		}
		return ctx.DB().Updates(req).Error
	})
}

// UpdateRoleStatus 更新数据状态
func (infra *Role) UpdateRoleStatus(ctx kratosx.Context, id uint32, status bool) error {
	// 删除缓存
	defer infra.delCache(ctx)

	return ctx.DB().Model(entity.Role{}).Where("id=?", id).Update("status", status).Error
}

// DeleteRole 删除数据
func (infra *Role) DeleteRole(ctx kratosx.Context, id uint32) error {
	// 删除缓存
	defer infra.delCache(ctx)

	ids, err := infra.GetRoleChildrenIds(ctx, []uint32{id})
	if err != nil {
		return err
	}
	ids = append(ids, id)
	return ctx.DB().Where("id in ?", ids).Delete(&entity.Role{}).Error
}

// GetRoleChildrenIds 获取指定id的所有子id
func (infra *Role) GetRoleChildrenIds(ctx kratosx.Context, pids []uint32) ([]uint32, error) {
	var ids []uint32
	return ids, ctx.DB().Model(entity.RoleClosure{}).
		Select("children").
		Where("parent in ?", pids).
		Scan(&ids).Error
}

// GetRoleParentIds 获取指定id的所有父id
func (infra *Role) GetRoleParentIds(ctx kratosx.Context, id uint32) ([]uint32, error) {
	var ids []uint32
	return ids, ctx.DB().Model(entity.RoleClosure{}).
		Select("parent").
		Where("children=?", id).
		Scan(&ids).Error
}

// appendRoleChildren 添加id到指定的父id下
func (infra *Role) appendRoleChildren(ctx kratosx.Context, pid uint32, id uint32) error {
	list := []*entity.RoleClosure{
		{
			Parent:   pid,
			Children: id,
		},
	}
	ids, _ := infra.GetRoleParentIds(ctx, pid)
	for _, item := range ids {
		list = append(list, &entity.RoleClosure{
			Parent:   item,
			Children: id,
		})
	}
	return ctx.DB().Create(&list).Error
}

// removeRoleParent 删除指定id的所有父层级
func (infra *Role) removeRoleParent(ctx kratosx.Context, id uint32) error {
	return ctx.DB().Delete(&entity.RoleClosure{}, "children=?", id).Error
}

// GetRoleMenuIds 获取指定角色的所有id
func (infra *Role) GetRoleMenuIds(ctx kratosx.Context, rids []uint32) ([]uint32, error) {
	var ids []uint32
	return ids, ctx.DB().Model(entity.RoleMenu{}).
		Select("menu_id").
		Where("role_id in ?", rids).
		Scan(&ids).Error
}

// UpdateRoleMenu 更新所有角色的id
func (infra *Role) UpdateRoleMenu(ctx kratosx.Context, roleId uint32, menuIds []uint32) error {
	var list []*entity.RoleMenu
	for _, mid := range menuIds {
		list = append(list, &entity.RoleMenu{
			RoleId: roleId,
			MenuId: mid,
		})
	}

	// 删除缓存
	defer infra.delCache(ctx)

	return ctx.Transaction(func(ctx kratosx.Context) error {
		if err := ctx.DB().Delete(entity.RoleMenu{}, "role_id=?", roleId).Error; err != nil {
			return err
		}
		if err := ctx.DB().Create(&list).Error; err != nil {
			return err
		}
		return nil
	})
}

func (infra *Role) GetRoleChildrenKeywords(ctx kratosx.Context, id uint32) ([]string, error) {
	ids, err := infra.GetRoleChildrenIds(ctx, []uint32{id})
	if err != nil {
		return nil, err
	}
	ids = append(ids, id)

	// 获取全部keyword
	var keywords []string
	return keywords, ctx.DB().Model(entity.Role{}).
		Select("keyword").
		Where("id in ?", ids).
		Scan(&keywords).Error
}

// GetRoleScopeByUID 获取角色组的的下级角色权限
func (infra *Role) GetRoleScopeByUID(ctx kratosx.Context, uid uint32) (bool, []uint32, error) {
	if uid == 1 {
		return true, nil, nil
	}

	rids, err := infra.getRoleIdsByUID(ctx, uid)
	if err != nil {
		return false, nil, err
	}

	ids, err := infra.GetRoleChildrenIds(ctx, rids)
	if err != nil {
		return false, nil, err
	}
	ids = append(ids, rids...)
	return false, valx.Unique(ids), nil
}

// GetRoleScope 获取角色组的的下级角色权限
func (infra *Role) GetRoleScope(ctx kratosx.Context, rids []uint32) (bool, []uint32, error) {
	if valx.InList(rids, 1) {
		return true, nil, nil
	}
	ids, err := infra.GetRoleChildrenIds(ctx, rids)
	if err != nil {
		return false, nil, err
	}
	ids = append(ids, rids...)
	return false, valx.Unique(ids), nil
}

// HasRolePurview 判断能否查看当前角色
func (infra *Role) HasRolePurview(ctx kratosx.Context, uid uint32, rids []uint32) (bool, error) {
	// 获取当前用户的所有角色
	var roleIds []uint32
	if err := ctx.DB().
		Model(entity.UserRole{}).
		Select([]string{"role_id"}).
		Where("user_id=?", uid).
		Scan(&roleIds).Error; err != nil {
		return false, err
	}

	all, scopes, err := infra.GetRoleScope(ctx, roleIds)
	if err != nil {
		return false, err
	}
	if all {
		return true, nil
	}

	uni := valx.New(scopes)
	for _, id := range rids {
		if !uni.Has(id) {
			return false, nil
		}
	}
	return true, nil
}

// GetDataScope 获取角色组的指定菜单的数据权限
func (infra *Role) GetDataScope(ctx kratosx.Context, uid uint32) (bool, []uint32, error) {
	if uid == 1 {
		return true, nil, nil
	}
	// 获取key映射的角色信息
	roleMap, err := infra.GetCacheRoleKeyMap(ctx)
	if err != nil {
		return false, nil, err
	}

	// 获取用户对应的部门角色
	var depRoles []*entity.DepartmentRole
	err = ctx.DB().Model(entity.DepartmentRole{}).
		Where("department_id in (?)",
			ctx.DB().Model(entity.UserDepartment{}).
				Select("department_id").
				Where("user_id = ?", uid),
		).
		Find(&depRoles).Error
	if err != nil {
		return false, nil, err
	}
	for _, item := range depRoles {
		if item.RoleId == 1 {
			return true, nil, nil
		}
	}

	// 获取用户单独设置的角色权限
	var userRoleIds []uint32
	err = ctx.DB().Model(entity.UserRole{}).
		Select("role_id").
		Where("user_id = ?", uid).
		Scan(&userRoleIds).Error
	if err != nil {
		return false, nil, err
	}
	if valx.InList(userRoleIds, 1) {
		return true, nil, nil
	}

	// 没有任何部门的权限
	if len(depRoles) == 0 && len(userRoleIds) == 0 {
		return false, []uint32{}, nil
	}

	// 转换角色为id作为映射
	var riMap = map[uint32]*entity.Role{}
	for _, item := range roleMap {
		riMap[item.Id] = item
	}

	// 合并角色权限
	dataScopes, isAll := infra.mergeRoleDataScope(riMap, depRoles, userRoleIds)
	if isAll {
		return true, nil, err
	}

	// 获取需要获取下级部门的都有部门ID
	var (
		depIds          []uint32
		needChildDepIds []uint32
	)
	for _, item := range dataScopes {
		if valx.InList([]string{
			entity.ScopeAssignAll,
			entity.ScopeCurrentDown,
			entity.ScopeDown,
		}, item.scope) {
			needChildDepIds = append(needChildDepIds, item.depId)
		}
		switch item.scope {
		// 包含当前部门的，直接追加
		case entity.ScopeCurrentDown, entity.ScopeCurrent:
			depIds = append(depIds, item.depId)
		// 指定部门的直接追加
		case entity.ScopeCustom:
			depIds = append(depIds, item.data...)
		}
	}

	// 查询这些部门的所有下级部门
	var depChildIds []uint32
	if err := ctx.DB().Model(entity.DepartmentClosure{}).
		Select("children").
		Where("parent in ?", needChildDepIds).
		Scan(&depChildIds).Error; err != nil {
		return false, nil, err
	}

	// 获取全部有权限的部门
	depIds = append(depIds, depChildIds...)

	return false, valx.Unique(depIds), nil
}

func (infra *Role) HasDataPurview(ctx kratosx.Context, uid uint32, rids []uint32) (bool, error) {
	all, scopes, err := infra.GetDataScope(ctx, uid)
	if err != nil {
		return false, err
	}
	if all {
		return true, nil
	}

	uni := valx.New(scopes)
	for _, id := range rids {
		if !uni.Has(id) {
			return false, nil
		}
	}
	return true, nil
}

// // GetRoleMenuDataScope 获取角色组的指定菜单的数据权限
//
//	func (infra *Role) GetRoleMenuDataScope(ctx kratosx.Context, uid uint32, path, method string) (bool, []uint32, error) {
//		// 查询这个菜单哪些角色拥有权限
//		var roleKeys []string
//		policies := ctx.Authentication().Enforce().GetFilteredPolicy(1, path, method)
//		for _, policy := range policies {
//			if len(policy) == 0 {
//				continue
//			}
//			roleKeys = append(roleKeys, policy[0])
//		}
//
//		// 获取key映射的角色信息
//		roleMap, err := infra.GetCacheRoleKeyMap(ctx)
//		if err != nil {
//			return false, nil, err
//		}
//
//		// 获取角色具有菜单权限的角色id
//		var allowRoleIds []uint32
//		for _, key := range roleKeys {
//			if roleMap[key] == nil {
//				continue
//			}
//			allowRoleIds = append(allowRoleIds, roleMap[key].Id)
//		}
//
//		// 获取用户对应的部门角色
//		var depRoles []*entity.DepartmentRole
//		err = ctx.DB().Model(entity.DepartmentRole{}).
//			Where("department_id in (?)",
//				ctx.DB().Model(entity.UserDepartment{}).
//					Select("department_id").
//					Where("user_id = ?", uid),
//			).
//			Where("role_id in ?", allowRoleIds).
//			Find(&depRoles).Error
//		if err != nil {
//			return false, nil, err
//		}
//
//		// 获取用户单独设置的角色权限
//		var userRoleIds []uint32
//		err = ctx.DB().Model(entity.UserRole{}).
//			Select("role_id").
//			Where("user_id = ?", uid).
//			Where("role_id in ?", allowRoleIds).
//			Scan(&userRoleIds).Error
//		if err != nil {
//			return false, nil, err
//		}
//
//		// 没有任何部门的权限
//		if len(depRoles) == 0 && len(userRoleIds) == 0 {
//			return false, []uint32{}, nil
//		}
//
//		// 转换角色为id作为映射
//		var riMap map[uint32]*entity.Role
//		for _, item := range roleMap {
//			riMap[item.Id] = item
//		}
//
//		// 合并角色权限
//		dataScopes, isAll := infra.mergeRoleDataScope(riMap, depRoles, userRoleIds)
//		if isAll {
//			return true, nil, err
//		}
//
//		// 获取需要获取下级部门的都有部门ID
//		var (
//			depIds          []uint32
//			needChildDepIds []uint32
//		)
//		for _, item := range dataScopes {
//			if valx.InList([]string{
//				entity.ScopeAssignAll,
//				entity.ScopeCurrentDown,
//				entity.ScopeDown,
//			}, item.scope) {
//				needChildDepIds = append(needChildDepIds, item.depId)
//			}
//			switch item.scope {
//			// 包含当前部门的，直接追加
//			case entity.ScopeCurrentDown, entity.ScopeCurrent:
//				depIds = append(depIds, item.depId)
//			// 指定部门的直接追加
//			case entity.ScopeCustom:
//				depIds = append(depIds, item.data...)
//			}
//		}
//
//		// 查询这些部门的所有下级部门
//		var depChildIds []uint32
//		if err := ctx.DB().Model(entity.DepartmentClosure{}).
//			Select("children").
//			Where("parent in ?", needChildDepIds).
//			Scan(&depChildIds).Error; err != nil {
//			return false, nil, err
//		}
//
//		// 获取全部有权限的部门
//		depIds = append(depIds, depChildIds...)
//
//		return false, valx.Unique(depIds), nil
//	}

func (infra *Role) GetEnableRoleIdsByUID(ctx kratosx.Context, uid uint32) ([]uint32, error) {
	// 获取key映射的角色信息
	roleMap, err := infra.GetCacheRoleIDMap(ctx)
	if err != nil {
		return nil, err
	}

	rids, err := infra.getRoleIdsByUID(ctx, uid)
	if err != nil {
		return nil, err
	}

	var res []uint32
	for _, rid := range rids {
		if roleMap[rid] != nil {
			res = append(res, rid)
		}
	}
	return res, nil
}

func (infra *Role) getRoleIdsByUID(ctx kratosx.Context, uid uint32) ([]uint32, error) {
	// 获取用户对应的部门角色
	var depRoleIds []uint32
	err := ctx.DB().Model(entity.DepartmentRole{}).
		Select("role_id").
		Where("department_id in (?)",
			ctx.DB().Model(entity.UserDepartment{}).
				Select("department_id").
				Where("user_id = ?", uid),
		).
		Find(&depRoleIds).Error
	if err != nil {
		return nil, err
	}

	// 获取用户单独设置的角色权限
	var userRoleIds []uint32
	err = ctx.DB().Model(entity.UserRole{}).
		Select("role_id").
		Where("user_id = ?", uid).
		Scan(&userRoleIds).Error
	if err != nil {
		return nil, err
	}
	return append(depRoleIds, userRoleIds...), nil
}

func (infra *Role) AllRoleKeywordByMenuId(ctx kratosx.Context, id uint32) ([]string, error) {
	var (
		keys []string
		ids  []uint32
	)

	if err := ctx.DB().Model(entity.RoleMenu{}).
		Scan("menu_id").
		Where("role_id=?", id).
		Scan(&ids).Error; err != nil {
		return nil, err
	}

	return keys, ctx.DB().Model(entity.Role{}).Select("keyword").
		Where("id in ?", ids).
		Scan(&keys).
		Error
}

func (infra *Role) mergeUserRoleScope(roles []*entity.Role) ([]*roleDataScope, string) {
	var (
		scopes   []*roleDataScope
		leave    = -1
		maxScope string
	)

	for _, item := range roles {
		tl, ok := entity.ScopeLeaves[item.DataScope]
		if !ok {
			continue
		}

		// 指定获取部门下的所有
		if item.DataScope == entity.ScopeAssignAll {
			depids := item.GetCustomDataIds()
			if len(depids) == 0 {
				continue
			}
			scopes = append(scopes, &roleDataScope{
				scope: item.DataScope,
				depId: depids[0],
			})
			continue
		}

		// 指定部门
		if item.DataScope == entity.ScopeCustom {
			scopes = append(scopes, &roleDataScope{
				scope: item.DataScope,
				data:  item.GetCustomDataIds(),
			})
			continue
		}

		// 其他权限，则计算一个最终最大的权限值
		if tl < leave {
			continue
		}
		maxScope = item.DataScope
	}

	return scopes, maxScope
}

func (infra *Role) mergeDepUserRoleScope(
	roleMap map[uint32]*entity.Role,
	depRoles []*entity.DepartmentRole,
	maxScope string,
) ([]*roleDataScope, bool) {
	var scopes []*roleDataScope

	for _, item := range depRoles {
		// 获取角色
		role := roleMap[item.RoleId]
		if role == nil {
			continue
		}

		// 指定获取部门下的所有
		if role.DataScope == entity.ScopeAssignAll {
			depids := role.GetCustomDataIds()
			if len(depids) == 0 {
				continue
			}
			scopes = append(scopes, &roleDataScope{
				scope: role.DataScope,
				depId: depids[0],
			})
			continue
		}

		// 指定部门
		if role.DataScope == entity.ScopeCustom {
			scopes = append(scopes, &roleDataScope{
				scope: role.DataScope,
				data:  role.GetCustomDataIds(),
			})
			continue
		}

		// 其他权限，则选择大的一个权限
		if maxScope != "" && entity.ScopeLeaves[role.DataScope] < entity.ScopeLeaves[maxScope] {
			role.DataScope = maxScope
		}

		// 是顶级部门，且包含当前下级，直接是全部权限
		if item.DepartmentId == 1 && role.DataScope == entity.ScopeCurrentDown {
			return nil, true
		}

		scopes = append(scopes, &roleDataScope{
			scope: role.DataScope,
			depId: item.DepartmentId,
		})
	}

	return scopes, false
}

// 计算角色权限
func (infra *Role) mergeRoleDataScope(
	roleMap map[uint32]*entity.Role,
	depRoles []*entity.DepartmentRole,
	userRoleIds []uint32,
) ([]*roleDataScope, bool) {
	// 获取用户角色
	var userRoles []*entity.Role
	for _, id := range userRoleIds {
		if roleMap[id] == nil {
			continue
		}
		userRoles = append(userRoles, roleMap[id])
	}

	// 合并用户数据权限,返回参数：自定义权限，影响全局的权限
	dataScopes, maxScope := infra.mergeUserRoleScope(userRoles)

	// 合并用户角色和部门角色
	depDataScopes, isAll := infra.mergeDepUserRoleScope(roleMap, depRoles, maxScope)

	// 合并权限
	return append(dataScopes, depDataScopes...), isAll
}
