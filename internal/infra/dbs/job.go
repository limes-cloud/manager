package dbs

import (
	"errors"
	"sync"

	"github.com/limes-cloud/kratosx"
	"github.com/limes-cloud/kratosx/pkg/valx"

	"github.com/limes-cloud/manager/internal/domain/entity"
	"github.com/limes-cloud/manager/internal/types"
)

type Job struct {
}

var (
	jobIns  *Job
	jobOnce sync.Once
)

func NewJob() *Job {
	jobOnce.Do(func() {
		jobIns = &Job{}
	})
	return jobIns
}

// GetJob 获取指定的数据
func (infra *Job) GetJob(ctx kratosx.Context, id uint32) (*entity.Job, error) {
	var job = entity.Job{}
	return &job, ctx.DB().First(&job, id).Error
}

// GetJobByKeyword 获取指定数据
func (infra *Job) GetJobByKeyword(ctx kratosx.Context, keyword string) (*entity.Job, error) {
	var job = entity.Job{}
	return &job, ctx.DB().Where("keyword = ?", keyword).First(&job).Error
}

// ListJob 获取列表
func (infra *Job) ListJob(ctx kratosx.Context, req *types.ListJobRequest) ([]*entity.Job, error) {
	var (
		es []*entity.Job
		fs = []string{"*"}
	)

	db := ctx.DB().Model(entity.Job{}).Select(fs)
	if req.Ids != nil {
		db = db.Where("id in ?", req.Ids)
	}
	if req.Name != nil {
		db = db.Where("name LIKE ?", *req.Name+"%")
	}
	if req.RootId != nil {
		ids, _ := infra.GetJobChildrenIds(ctx, []uint32{*req.RootId})
		db = db.Where("id in ?", ids)
	}

	return es, db.Find(&es).Error
}

// CreateJob 创建数据
func (infra *Job) CreateJob(ctx kratosx.Context, job *entity.Job) (uint32, error) {
	return job.Id, ctx.Transaction(func(ctx kratosx.Context) error {
		if err := ctx.DB().Create(job).Error; err != nil {
			return err
		}
		return infra.appendJobChildren(ctx, job.ParentId, job.Id)
	})
}

// UpdateJob 更新数据
func (infra *Job) UpdateJob(ctx kratosx.Context, req *entity.Job) error {
	if req.Id == req.ParentId {
		return errors.New("父级不能为自己")
	}
	old, err := infra.GetJob(ctx, req.Id)
	if err != nil {
		return err
	}

	return ctx.Transaction(func(ctx kratosx.Context) error {
		if old.ParentId != req.ParentId {
			if err := infra.removeJobParent(ctx, req.Id); err != nil {
				return err
			}
			if err := infra.appendJobChildren(ctx, req.ParentId, req.Id); err != nil {
				return err
			}
		}
		return ctx.DB().Updates(req).Error
	})
}

// DeleteJob 删除数据
func (infra *Job) DeleteJob(ctx kratosx.Context, id uint32) error {
	ids, err := infra.GetJobChildrenIds(ctx, []uint32{id})
	if err != nil {
		return err
	}
	ids = append(ids, id)
	return ctx.DB().Where("id in ?", ids).Delete(&entity.Job{}).Error
}

// GetJobChildrenIds 获取指定id的所有子id
func (infra *Job) GetJobChildrenIds(ctx kratosx.Context, pids []uint32) ([]uint32, error) {
	var ids []uint32
	return ids, ctx.DB().Model(entity.JobClosure{}).
		Select("children").
		Where("parent in ?", pids).
		Scan(&ids).Error
}

// GetJobChildrenIdsByIds 获取指定id的所有子id
func (infra *Job) GetJobChildrenIdsByIds(ctx kratosx.Context, pids []uint32) ([]uint32, error) {
	var ids []uint32
	return ids, ctx.DB().Model(entity.JobClosure{}).
		Select("children").
		Where("parent in ?", pids).
		Scan(&ids).Error
}

// GetJobParentIds 获取指定id的所有父id
func (infra *Job) GetJobParentIds(ctx kratosx.Context, id uint32) ([]uint32, error) {
	var ids []uint32
	return ids, ctx.DB().Model(entity.JobClosure{}).
		Select("parent").
		Where("children=?", id).
		Scan(&ids).Error
}

// appendJobChildren 添加id到指定的父id下
func (infra *Job) appendJobChildren(ctx kratosx.Context, pid uint32, id uint32) error {
	list := []*entity.JobClosure{
		{
			Parent:   pid,
			Children: id,
		},
	}
	ids, _ := infra.GetJobParentIds(ctx, pid)
	for _, item := range ids {
		list = append(list, &entity.JobClosure{
			Parent:   item,
			Children: id,
		})
	}
	return ctx.DB().Model(entity.JobClosure{}).Create(&list).Error
}

// removeJobParent 删除指定id的所有父层级
func (infra *Job) removeJobParent(ctx kratosx.Context, id uint32) error {
	return ctx.DB().Delete(&entity.JobClosure{}, "children=?", id).Error
}

func (infra *Job) GetJobChildrenKeywords(ctx kratosx.Context, id uint32) ([]string, error) {
	ids, err := infra.GetJobChildrenIds(ctx, []uint32{id})
	if err != nil {
		return nil, err
	}
	ids = append(ids, id)

	// 获取全部keyword
	var keywords []string
	return keywords, ctx.DB().Model(entity.Job{}).
		Select("keyword").
		Where("id in ?", ids).
		Scan(&keywords).Error
}

func (infra *Job) mergeRoleScope(roles []*entity.Role) (string, []uint32) {
	var (
		scope = ""
		leave = 99
		ids   []uint32
	)

	for _, item := range roles {
		tl, ok := entity.ScopeLeaves[item.DataScope]
		if !ok {
			continue
		}
		if item.DataScope == entity.ScopeCustom {
			ids = append(ids, item.GetCustomJobIds()...)
		}
		if tl > leave {
			continue
		}
		scope = item.DataScope
	}
	return scope, ids
}

func (infra *Job) GetJobDataScope(ctx kratosx.Context, uid uint32) (bool, []uint32, error) {
	if uid == 1 {
		return true, nil, nil
	}

	// 获取用户职位
	var jobIds []uint32
	if err := ctx.DB().
		Model(entity.UserJob{}).
		Select([]string{"job_id"}).
		Where("user_id = ?", uid).
		Scan(&jobIds).Error; err != nil {
		return false, nil, err
	}

	// 获取用户角色
	var roleIds []uint32
	if err := ctx.DB().
		Model(entity.UserRole{}).
		Select([]string{"role_id"}).
		Where("user_id=?", uid).
		Scan(&roleIds).Error; err != nil {
		return false, nil, err
	}

	// 获取全部角色
	var roles []*entity.Role
	if err := ctx.DB().
		Select([]string{"job_scope", "job_ids"}).
		First(&roles, "id in ?", roleIds).Error; err != nil {
		return false, nil, err
	}

	// 合并角色权限
	scope, cids := infra.mergeRoleScope(roles)

	if scope == entity.ScopeAssignAll {
		return true, []uint32{}, nil
	}

	// 如果当前部门是最顶层部门，且权限为ScopeCurrentDown
	if scope == entity.ScopeCurrentDown && valx.InList(jobIds, 1) {
		return true, []uint32{}, nil
	}

	switch scope {
	case entity.ScopeCurrent:
		ids := append(jobIds, cids...)
		return false, ids, nil
	case entity.ScopeCurrentDown:
		ids, err := infra.GetJobChildrenIds(ctx, jobIds)
		if err != nil {
			return false, nil, err
		}
		ids = append(ids, append(jobIds, cids...)...)
		return false, ids, nil
	case entity.ScopeDown:
		ids, err := infra.GetJobChildrenIds(ctx, jobIds)
		if err != nil {
			return false, nil, err
		}
		ids = append(ids, cids...)
		return false, ids, nil
	case entity.ScopeCustom:
		return false, cids, nil
	}
	return false, []uint32{}, nil
}

func (infra *Job) HasJobPurview(ctx kratosx.Context, uid uint32, jids []uint32) (bool, error) {
	all, scopes, err := infra.GetJobDataScope(ctx, uid)
	if err != nil {
		return false, err
	}
	if all {
		return true, nil
	}

	uni := valx.New(scopes)
	for _, id := range jids {
		if !uni.Has(id) {
			return false, nil
		}
	}
	return true, nil
}
