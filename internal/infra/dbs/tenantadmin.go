package dbs

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"sync"

	"github.com/limes-cloud/kratosx/pkg/cache"

	"github.com/limes-cloud/kratosx"
	"github.com/limes-cloud/kratosx/model/page"
	"github.com/limes-cloud/manager/internal/core"
	"github.com/limes-cloud/manager/internal/domain/entity"
	"github.com/limes-cloud/manager/internal/types"
)

type TenantAdmin struct {
	cache     *cache.Cache[string, struct{}]
	cacheMenu *cache.Cache[string, struct{}]
}

var (
	tadIns      *TenantAdmin
	tadOnce     sync.Once
	tadCacheKey = "tenantadmin"
)

func NewTenantAdmin() *TenantAdmin {
	tadOnce.Do(func() {
		tadIns = &TenantAdmin{}
		ctx := core.MustContext(
			context.Background(),
			kratosx.WithTrace("cache", "tenant admin"),
			kratosx.WithSkipDBHook(),
		)

		ctx.BeforeStart("load cache tenant admin", func() {
			tadIns.cache = cache.NewCacheAndInit[string, struct{}](
				ctx,
				ctx.Redis(),
				tadCacheKey,
				cache.HookLoad(func() (map[string]struct{}, error) {
					return tadIns.load(ctx)
				}),
			)
		})
	})
	return tadIns
}

func (ta *TenantAdmin) IsAdmin(tid uint32, uid uint32) bool {
	_, is := ta.cache.Get(ta.getCacheKey(tid, uid))
	return is
}

// ListTenantAdmin 获取admin列表
func (ta *TenantAdmin) ListTenantAdmin(ctx core.Context, req *types.ListTenantAdminRequest) ([]*entity.TenantAdmin, uint32, error) {
	var (
		list  []*entity.TenantAdmin
		total int64
	)

	db := ctx.DB().
		Model(entity.TenantAdmin{}).
		Joins("User").
		Where("tenant_admin.tenant_id=?", req.TenantId)

	db = db.Where("User.id is not null")

	if req.Search != nil {
		// 查询条件下数据总数
		if err := db.Count(&total).Error; err != nil {
			return nil, 0, err
		}

		// 分页查询
		db = page.SearchScopes(db, req.Search)
	}
	return list, uint32(total), db.Find(&list).Error
}

// CreateTenantAdmin 创建租户的应用信息
func (ta *TenantAdmin) CreateTenantAdmin(ctx core.Context, ent *entity.TenantAdmin) (uint32, error) {
	op := ta.cache.OP(ctx)
	err := ctx.Transaction(func(ctx core.Context) error {
		if err := ctx.DB().Create(ent).Error; err != nil {
			return err
		}
		return op.Put(ta.getCacheKey(ent.TenantId, ent.UserId), struct{}{}).Do()
	})
	if err != nil {
		return 0, err
	}
	return ent.Id, nil
}

// DeleteTenantAdmin 删除租户的应用
func (ta *TenantAdmin) DeleteTenantAdmin(ctx core.Context, req *types.DeleteTenantAdminRequest) error {
	op := ta.cache.OP(ctx)
	return ctx.Transaction(func(ctx core.Context) error {
		if err := ctx.DB().
			Where("tenant_id=? AND user_id=?", req.TenantId, req.UserId).
			Delete(&entity.TenantAdmin{}).Error; err != nil {
			return err
		}

		return op.Delete(ta.getCacheKey(req.TenantId, req.UserId)).Do()
	})
}

// load 加载全量的数据
func (ta *TenantAdmin) load(ctx core.Context) (map[string]struct{}, error) {
	var list []*entity.TenantAdmin
	if err := ctx.DB().Model(&entity.TenantAdmin{}).Find(&list).Error; err != nil {
		return nil, err
	}

	bucket := map[string]struct{}{}
	for _, item := range list {
		bucket[ta.getCacheKey(item.TenantId, item.UserId)] = struct{}{}
	}

	return bucket, nil
}

// getCacheKey 获取缓存的key
func (ta *TenantAdmin) getCacheKey(tid uint32, aid uint32) string {
	return fmt.Sprintf("%d:%d", tid, aid)
}

// HasAdminScope 判断是否具有指定的应用权限
func (ta *TenantAdmin) HasAdminScope(tid uint32, aid uint32) bool {
	key := ta.getCacheKey(tid, aid)
	_, ok := ta.cache.Get(key)
	return ok
}

// splitCacheKey 通过key获取id
func (ta *TenantAdmin) splitCacheKey(key string) (uint32, uint32) {
	arr := strings.Split(key, ":")
	if len(arr) != 2 {
		return 0, 0
	}
	rid, _ := strconv.Atoi(arr[0])
	mid, _ := strconv.Atoi(arr[1])
	return uint32(rid), uint32(mid)
}

func (ta *TenantAdmin) GetAdminIds(tid uint32) []uint32 {
	var (
		ids  = make([]uint32, 0)
		keys = ta.cache.Keys()
	)

	for _, key := range keys {
		trid, aid := ta.splitCacheKey(key)
		if trid == tid {
			ids = append(ids, aid)
		}
	}
	return ids
}
