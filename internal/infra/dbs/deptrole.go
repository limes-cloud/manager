package dbs

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"sync"

	"github.com/limes-cloud/kratosx/pkg/value"

	"github.com/limes-cloud/kratosx/model/page"

	"github.com/limes-cloud/kratosx"
	"github.com/limes-cloud/kratosx/pkg/cache"
	"github.com/limes-cloud/manager/internal/core"
	"github.com/limes-cloud/manager/internal/domain/entity"
	"github.com/limes-cloud/manager/internal/types"
)

type DeptRole struct {
	cache *cache.Cache[string, struct{}]
}

var (
	drIns      *DeptRole
	drOnce     sync.Once
	drCacheKey = "deptrole"
)

func NewDeptRole() *DeptRole {
	drOnce.Do(func() {
		// 监听变更时间
		drIns = &DeptRole{}
		ctx := core.MustContext(
			context.Background(),
			kratosx.WithTrace("cache", "dept role"),
			kratosx.WithSkipDBHook(),
		)

		ctx.BeforeStart("load cache dept role", func() {
			drIns.cache = cache.NewCacheAndInit[string, struct{}](
				ctx,
				ctx.Redis(),
				drCacheKey,
				cache.HookLoad(func() (map[string]struct{}, error) {
					return drIns.load(ctx)
				}),
			)
		})
	})
	return drIns
}

// ListDeptRole 获取指定部门的角色列表
func (dr *DeptRole) ListDeptRole(ctx core.Context, req *types.ListDeptRoleRequest) ([]*entity.Role, uint32, error) {
	var (
		total int64
		list  []*entity.Role
	)
	db := ctx.DB().
		Model(&entity.Role{}).
		Joins("left join dept_role on role.id = dept_role.role_id").
		Where("dept_role.role_id is not null").
		Where("role.status = 1").
		Where("dept_role.dept_id = ?", req.DeptId)
	if req.Name != nil {
		db = db.Where("role.name like ?", *req.Name+"%")
	}
	if req.InRoleIds != nil {
		db = db.Where("role.id in ?", req.InRoleIds)
	}

	if err := db.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	db = page.SearchScopes(db, &req.Search)
	return list, uint32(total), db.Find(&list).Error
}

// CreateDeptRole 创建部门角色
func (dr *DeptRole) CreateDeptRole(ctx core.Context, deptId, roleId uint32) error {
	ent := &entity.DeptRole{
		DeptId: deptId,
		RoleId: roleId,
	}

	op := dr.cache.OP(ctx)
	return ctx.Transaction(func(ctx core.Context) error {
		if err := ctx.DB().Model(&entity.DeptRole{}).Create(&ent).Error; err != nil {
			return err
		}
		return op.Puts(dr.transvals([]*entity.DeptRole{ent})).Do()
	})
}

// DeleteDeptRole 获取指定部门的角色列表
func (dr *DeptRole) DeleteDeptRole(ctx core.Context, deptId, roleId uint32) error {
	ent := &entity.DeptRole{
		DeptId: deptId,
		RoleId: roleId,
	}
	op := dr.cache.OP(ctx)
	return ctx.Transaction(func(ctx core.Context) error {
		db := ctx.DB()
		if err := db.
			Where("dept_id = ?", deptId).
			Where("role_id = ?", roleId).
			Delete(ent).Error; err != nil {
			return err
		}

		// 从缓存中删除
		return op.Deletes(dr.transkeys([]*entity.DeptRole{ent})).Do()
	})
}

// GetDeptRolesByDeptIds 获取指定部门绑定的所有角色ID
func (dr *DeptRole) GetDeptRolesByDeptIds(deptIds []uint32) []*entity.DeptRole {
	var (
		list []*entity.DeptRole
		keys = dr.cache.Keys()
	)

	for _, key := range keys {
		did, rid := dr.splitCacheKey(key)
		if rid != 0 && value.InList(deptIds, did) {
			list = append(list, &entity.DeptRole{
				DeptId: did,
				RoleId: rid,
			})
		}
	}
	return list
}

// GetRoleIdsByDeptIds 获取指定部门绑定的所有角色ID
func (dr *DeptRole) GetRoleIdsByDeptIds(deptIds []uint32) []uint32 {
	var (
		ids  = make([]uint32, 0)
		keys = dr.cache.Keys()
	)

	for _, key := range keys {
		did, rid := dr.splitCacheKey(key)
		if rid != 0 && value.InList(deptIds, did) {
			ids = append(ids, rid)
		}
	}
	return ids
}

// GetDeptIdsByRoleIds 获取指定角色绑定的所有部门ID
func (dr *DeptRole) GetDeptIdsByRoleIds(roleIds []uint32) []uint32 {
	var (
		ids  = make([]uint32, 0)
		keys = dr.cache.Keys()
	)

	for _, key := range keys {
		did, rid := dr.splitCacheKey(key)
		if did != 0 && value.InList(roleIds, rid) {
			ids = append(ids, did)
		}
	}
	return ids
}

func (dr *DeptRole) transvals(drs []*entity.DeptRole) map[string]struct{} {
	kvs := map[string]struct{}{}
	for _, item := range drs {
		kvs[dr.getCacheKey(item.DeptId, item.RoleId)] = struct{}{}
	}
	return kvs
}

func (dr *DeptRole) transkeys(drs []*entity.DeptRole) []string {
	var keys []string
	for _, item := range drs {
		keys = append(keys, dr.getCacheKey(item.DeptId, item.RoleId))
	}
	return keys
}

// load 加载全量的数据
func (dr *DeptRole) load(ctx core.Context) (map[string]struct{}, error) {
	var list []*entity.DeptRole
	if err := ctx.DB().Model(&entity.DeptRole{}).Find(&list).Error; err != nil {
		return nil, err
	}

	bucket := map[string]struct{}{}
	for _, item := range list {
		bucket[dr.getCacheKey(item.DeptId, item.RoleId)] = struct{}{}
	}

	return bucket, nil
}

// getCacheKey 获取缓存的key
func (dr *DeptRole) getCacheKey(deptId uint32, roleId uint32) string {
	return fmt.Sprintf("%d:%d", deptId, roleId)
}

// splitCacheKey 通过key获取id
func (dr *DeptRole) splitCacheKey(key string) (uint32, uint32) {
	arr := strings.Split(key, ":")
	if len(arr) != 2 {
		return 0, 0
	}
	deptId, _ := strconv.Atoi(arr[0])
	roleId, _ := strconv.Atoi(arr[1])
	return uint32(deptId), uint32(roleId)
}
