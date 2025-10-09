package dbs

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"sync"

	"github.com/limes-cloud/kratosx/model/page"

	"github.com/limes-cloud/kratosx"
	"github.com/limes-cloud/kratosx/pkg/cache"
	"github.com/limes-cloud/manager/internal/core"
	"github.com/limes-cloud/manager/internal/domain/entity"
	"github.com/limes-cloud/manager/internal/types"
	"gorm.io/gorm/clause"
)

type UserDept struct {
	cache *cache.Cache[string, struct{}]
}

var (
	udIns      *UserDept
	udOnce     sync.Once
	udCacheKey = "userdept"
)

func NewUserDept() *UserDept {
	udOnce.Do(func() {
		// 初始化属性
		udIns = &UserDept{}

		ctx := core.MustContext(
			context.Background(),
			kratosx.WithTrace("cache", "user dept"),
			kratosx.WithSkipDBHook(),
		)
		ctx.BeforeStart("load cache user dept", func() {
			udIns.cache = cache.NewCacheAndInit[string, struct{}](
				ctx,
				ctx.Redis(),
				udCacheKey,
				cache.HookLoad(func() (map[string]struct{}, error) {
					return udIns.load(ctx)
				}),
			)
		})
	})
	return udIns
}

// ListDeptUser 获取指定角色绑定的所有部门
func (ud *UserDept) ListDeptUser(ctx core.Context, req *types.ListDeptUserRequest) ([]*entity.User, uint32, error) {
	var (
		total int64
		list  []*entity.User
	)
	db := ctx.DB().Model(&entity.User{}).
		Joins("left join user_dept on user.id = user_dept.user_id").
		Joins("Classify", ctx.DB().Where("Classify.status=1")).
		Where("user_dept.dept_id is not null").
		Where("Classify.id is not null").
		Where("user.status = 1")

	if req.Name != nil {
		db = db.Where("user.name like ?", *req.Name+"%")
	}

	if err := db.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	db = page.SearchScopes(db, &req.Search)
	return list, uint32(total), db.Find(&list).Error
}

// ListUserDept 获取指定部门的角色列表
func (ud *UserDept) ListUserDept(ctx core.Context, req *types.ListUserDeptRequest) ([]*entity.Dept, uint32, error) {
	var (
		total int64
		list  []*entity.Dept
	)
	db := ctx.DB().
		Model(&entity.Dept{}).
		Joins("left join user_dept on dept.id = user_dept.dept_id").
		Where("user_dept.dept_id is not null").
		Where("dept.status = 1")
	if req.Name != nil {
		db = db.Where("dept.name like ?", *req.Name+"%")
	}
	if req.InDeptIds != nil {
		db = db.Where("dept.id in ?", req.InDeptIds)
	}
	if err := db.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	db = page.SearchScopes(db, &req.Search)
	return list, uint32(total), db.Find(&list).Error
}

// CreateUserDept 批量绑定部门部门角色关系
func (ud *UserDept) CreateUserDept(ctx core.Context, ent *entity.UserDept) error {
	op := ud.cache.OP(ctx)
	return ctx.Transaction(func(ctx core.Context) error {
		if err := ctx.DB().Model(&entity.UserDept{}).
			Clauses(clause.OnConflict{UpdateAll: true}).
			Create(&ent).Error; err != nil {
			return err
		}

		return op.Puts(ud.transvals([]*entity.UserDept{ent})).Do()
	})
}

// DeleteUserDept 获取指定菜单的角色列表
func (ud *UserDept) DeleteUserDept(ctx core.Context, ent *entity.UserDept) error {
	op := ud.cache.OP(ctx)
	return ctx.Transaction(func(ctx core.Context) error {
		if err := ctx.DB().Where("dept_id = ? AND user_id = ?", ent.DeptId, ent.UserId).Delete(&entity.UserDept{}).Error; err != nil {
			return err
		}

		// 从缓存中删除
		return op.Deletes(ud.transkeys([]*entity.UserDept{ent})).Do()
	})
}

// GetDeptIdsByUserId 获取指定部门绑定的所有角色ID
func (ud *UserDept) GetDeptIdsByUserId(userId uint32) []uint32 {
	var (
		ids  = make([]uint32, 0)
		keys = ud.cache.Keys()
	)

	for _, key := range keys {
		uid, did, _ := ud.splitCacheKey(key)
		if did != 0 && uid == userId {
			ids = append(ids, did)
		}

	}
	return ids
}

// ListUserDeptByUserId 获取指定部门绑定的所有角色ID
func (ud *UserDept) ListUserDeptByUserId(id uint32) []*entity.UserDept {
	var (
		list []*entity.UserDept
		keys = ud.cache.Keys()
	)

	for _, key := range keys {
		uid, did, jid := ud.splitCacheKey(key)
		if uid == id {
			list = append(list, &entity.UserDept{
				UserId: uid,
				DeptId: did,
				JobId:  jid,
			})
		}

	}
	return list
}

// GetUserIdsByDeptId 获取指定角色绑定的所有部门ID
func (ud *UserDept) GetUserIdsByDeptId(deptId uint32) []uint32 {
	var (
		ids  = make([]uint32, 0)
		keys = ud.cache.Keys()
	)

	for _, key := range keys {
		uid, did, _ := ud.splitCacheKey(key)
		if uid != 0 && did == deptId {
			ids = append(ids, uid)
		}
	}
	return ids
}

func (ud *UserDept) transvals(uds []*entity.UserDept) map[string]struct{} {
	kvs := map[string]struct{}{}
	for _, item := range uds {
		kvs[ud.getCacheKey(item.UserId, item.DeptId, item.JobId)] = struct{}{}
	}
	return kvs
}

func (ud *UserDept) transkeys(uds []*entity.UserDept) []string {
	var keys []string
	for _, item := range uds {
		keys = append(keys, ud.getCacheKey(item.UserId, item.DeptId, item.JobId))
	}
	return keys
}

// load 加载全量的数据
func (ud *UserDept) load(ctx core.Context) (map[string]struct{}, error) {
	var list []*entity.UserDept
	if err := ctx.DB().Model(&entity.UserDept{}).Find(&list).Error; err != nil {
		return nil, err
	}

	bucket := map[string]struct{}{}
	for _, item := range list {
		bucket[ud.getCacheKey(item.UserId, item.DeptId, item.JobId)] = struct{}{}
	}

	return bucket, nil
}

// getCacheKey 获取缓存的key
func (ud *UserDept) getCacheKey(userId uint32, deptId uint32, jobId uint32) string {
	return fmt.Sprintf("%d:%d:%d", userId, deptId, jobId)
}

// splitCacheKey 通过key获取id
func (ud *UserDept) splitCacheKey(key string) (uint32, uint32, uint32) {
	arr := strings.Split(key, ":")
	if len(arr) != 2 {
		return 0, 0, 0
	}
	userId, _ := strconv.Atoi(arr[0])
	deptId, _ := strconv.Atoi(arr[1])
	jobId, _ := strconv.Atoi(arr[2])
	return uint32(userId), uint32(deptId), uint32(jobId)
}
