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
	"gorm.io/gorm/clause"
)

type JobRole struct {
	cache *cache.Cache[string, struct{}]
}

var (
	jrIns      *JobRole
	jrOnce     sync.Once
	jrCacheKey = "jobrole"
)

func NewJobRole() *JobRole {
	jrOnce.Do(func() {
		// 监听变更时间
		ctx := core.MustContext(context.Background(), kratosx.WithTrace("cache", "job role"))

		// 初始化属性
		jrIns = &JobRole{}

		// 缓存相关操作
		{
			lc := cache.NewCache[string, struct{}](
				ctx.Redis(),
				jrCacheKey,
				cache.HookLoad(func() (map[string]struct{}, error) {
					return jrIns.load(ctx)
				}),
			)

			// 加载缓存,加载失败则直接报错，避免线上隐式错误。
			if err := lc.Init(ctx); err != nil {
				panic(err)
			}
			// 监听缓存变更
			go lc.Subscribe(ctx)
			// 定期修复缓存
			go lc.Repair(ctx)

			jrIns.cache = lc
		}
	})
	return jrIns
}

// ListRoleJob 获取指定角色绑定的所有部门
func (jr *JobRole) ListRoleJob(ctx core.Context, req *types.ListRoleJobRequest) ([]*entity.Job, uint32, error) {
	var (
		total int64
		list  []*entity.Job
	)
	db := ctx.DB().Model(&entity.Job{}).
		Joins("left join job_role on job.id = job_role.job_id").
		Joins("Classify", ctx.DB().Where("Classify.status=1")).
		Where("job_role.role_id is not null").
		Where("Classify.id is not null").
		Where("job.status = 1")

	if req.Name != nil {
		db = db.Where("job.name like ?", *req.Name+"%")
	}
	if req.InJobIds != nil {
		db = db.Where("job.id in ?", req.InJobIds)
	}

	if err := db.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	db = page.SearchScopes(db, &req.Search)
	return list, uint32(total), db.Find(&list).Error
}

// ListJobRole 获取指定部门的角色列表
func (jr *JobRole) ListJobRole(ctx core.Context, req *types.ListJobRoleRequest) ([]*entity.Role, uint32, error) {
	var (
		total int64
		list  []*entity.Role
	)
	db := ctx.DB().
		Model(&entity.Role{}).
		Joins("left join job_role on role.id = job_role.role_id").
		Where("job_role.role_id is not null").
		Where("role.status = 1")
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

// CreateJobRoles 批量绑定部门部门角色关系
func (jr *JobRole) CreateJobRoles(ctx core.Context, jrs []*entity.JobRole) error {
	op := jr.cache.OP(ctx)
	return ctx.Transaction(func(ctx core.Context) error {
		if err := ctx.DB().Model(&entity.JobRole{}).
			Clauses(clause.OnConflict{UpdateAll: true}).
			Create(&jrs).Error; err != nil {
			return err
		}

		return op.Puts(jr.transvals(jrs)).Do()
	})
}

// DeleteJobRoles 获取指定菜单的角色列表
func (jr *JobRole) DeleteJobRoles(ctx core.Context, jrs []*entity.JobRole) error {
	op := jr.cache.OP(ctx)
	return ctx.Transaction(func(ctx core.Context) error {
		db := ctx.DB()
		for _, item := range jrs {
			db = db.Or(ctx.DB().Where("role_id = ? AND job_id = ?", item.RoleId, item.JobId))
		}
		if err := db.Delete(&entity.JobRole{}).Error; err != nil {
			return err
		}

		// 从缓存中删除
		return op.Deletes(jr.transkeys(jrs)).Do()
	})
}

// GetRoleIdsByJobIds 获取指定部门绑定的所有角色ID
func (jr *JobRole) GetRoleIdsByJobIds(jobIds []uint32) []uint32 {
	var (
		ids  []uint32
		keys = jr.cache.Keys()
	)

	for _, key := range keys {
		did, rid := jr.splitCacheKey(key)
		if rid != 0 && value.InList(jobIds, did) {
			ids = append(ids, rid)
		}
	}
	return ids
}

// GetJobIdsByRoleIds 获取指定角色绑定的所有部门ID
func (jr *JobRole) GetJobIdsByRoleIds(roleIds []uint32) []uint32 {
	var (
		ids  []uint32
		keys = jr.cache.Keys()
	)

	for _, key := range keys {
		did, rid := jr.splitCacheKey(key)
		if did != 0 && value.InList(roleIds, rid) {
			ids = append(ids, did)
		}
	}
	return ids
}

func (jr *JobRole) transvals(jrs []*entity.JobRole) map[string]struct{} {
	kvs := map[string]struct{}{}
	for _, item := range jrs {
		kvs[jr.getCacheKey(item.JobId, item.RoleId)] = struct{}{}
	}
	return kvs
}

func (jr *JobRole) transkeys(jrs []*entity.JobRole) []string {
	var keys []string
	for _, item := range jrs {
		keys = append(keys, jr.getCacheKey(item.JobId, item.RoleId))
	}
	return keys
}

// load 加载全量的数据
func (jr *JobRole) load(ctx core.Context) (map[string]struct{}, error) {
	var list []*entity.JobRole
	if err := ctx.DB().Model(&entity.JobRole{}).Find(&list).Error; err != nil {
		return nil, err
	}

	bucket := map[string]struct{}{}
	for _, item := range list {
		bucket[jr.getCacheKey(item.JobId, item.RoleId)] = struct{}{}
	}

	return bucket, nil
}

// getCacheKey 获取缓存的key
func (jr *JobRole) getCacheKey(jobId uint32, roleId uint32) string {
	return fmt.Sprintf("%d:%d", jobId, roleId)
}

// splitCacheKey 通过key获取id
func (jr *JobRole) splitCacheKey(key string) (uint32, uint32) {
	arr := strings.Split(key, ":")
	if len(arr) != 2 {
		return 0, 0
	}
	jobId, _ := strconv.Atoi(arr[0])
	roleId, _ := strconv.Atoi(arr[1])
	return uint32(jobId), uint32(roleId)
}
