package service

import (
	"github.com/limes-cloud/kratosx/library/db/gormtranserror"
	"github.com/limes-cloud/kratosx/pkg/value"
	"github.com/limes-cloud/manager/api/errors"
	"github.com/limes-cloud/manager/internal/core"
	"github.com/limes-cloud/manager/internal/domain/entity"
	"github.com/limes-cloud/manager/internal/domain/repository"
	"github.com/limes-cloud/manager/internal/types"
	"gorm.io/gorm"
)

type Entity struct {
	repo repository.Entity
}

func NewEntity(repo repository.Entity) *Entity {
	return &Entity{repo: repo}
}

// ListEntityRule 获取租户列表
func (u *Entity) ListEntityRule(ctx core.Context, req *types.ListEntityRuleRequest) ([]*entity.EntityRule, uint32, error) {
	list, total, err := u.repo.ListEntityRule(ctx, req)
	if err != nil {
		ctx.Logger().Warnw("msg", "list entity rule error", "err", err.Error())
		return nil, 0, errors.ListError()
	}
	return list, total, nil
}

// CreateEntityRule 创建租户
func (u *Entity) CreateEntityRule(ctx core.Context, req *entity.EntityRule) (uint32, error) {
	id, err := u.repo.CreateEntityRule(ctx, req)
	if err != nil {
		ctx.Logger().Warnw("msg", "create entity rule error", "err", err.Error())
		return 0, errors.CreateError()
	}
	return id, nil
}

// UpdateEntityRule 更新租户
func (u *Entity) UpdateEntityRule(ctx core.Context, req *entity.EntityRule) error {
	if err := u.repo.UpdateEntityRule(ctx, req); err != nil {
		ctx.Logger().Warnw("msg", "update entity rule error", "err", err.Error())
		return errors.UpdateError()
	}
	return nil
}

// DeleteEntityRule 删除租户
func (u *Entity) DeleteEntityRule(ctx core.Context, id uint32) error {
	if err := u.repo.DeleteEntityRule(ctx, id); err != nil {
		ctx.Logger().Warnw("msg", "delete entity rule error", "err", err.Error())
		return errors.DeleteError()
	}
	return nil
}

// ListEntityField 获取租户列表
func (u *Entity) ListEntityField(ctx core.Context, req *types.ListEntityFieldRequest) ([]*entity.EntityField, uint32, error) {
	list, total, err := u.repo.ListEntityField(ctx, req)
	if err != nil {
		ctx.Logger().Warnw("msg", "list entity field error", "err", err.Error())
		return nil, 0, errors.ListError()
	}
	return list, total, nil
}

// CreateEntityField 创建租户
func (u *Entity) CreateEntityField(ctx core.Context, req *entity.EntityField) (uint32, error) {
	id, err := u.repo.CreateEntityField(ctx, req)
	if err != nil {
		ctx.Logger().Warnw("msg", "create entity field error", "err", err.Error())
		return 0, errors.CreateError()
	}
	return id, nil
}

// UpdateEntityField 更新租户
func (u *Entity) UpdateEntityField(ctx core.Context, req *entity.EntityField) error {
	if err := u.repo.UpdateEntityField(ctx, req); err != nil {
		ctx.Logger().Warnw("msg", "update entity field error", "err", err.Error())
		return errors.UpdateError()
	}
	return nil
}

// DeleteEntityField 删除租户
func (u *Entity) DeleteEntityField(ctx core.Context, id uint32) error {
	if err := u.repo.DeleteEntityField(ctx, id); err != nil {
		ctx.Logger().Warnw("msg", "delete entity field error", "err", err.Error())
		return errors.DeleteError()
	}
	return nil
}

// ImportEntity 导入实体列表
func (u *Entity) ImportEntity(ctx core.Context, ents []*entity.Entity) error {
	return ctx.Transaction(func(ctx core.Context) error {
		for _, ent := range ents {
			fields := ent.Fields
			ent.Fields = nil

			insertFields := func(entId uint32, fs []*entity.EntityField) error {
				for ind, field := range fs {
					// 查询是否存在字段
					_, err := u.repo.GetEntityField(ctx, entId, field.Name)
					if err != nil && !gormtranserror.Is(err, gorm.ErrRecordNotFound) {
						return err
					}
					if err == nil {
						continue
					}
					field.Index = ind
					field.EntityId = entId
					if _, err := u.repo.CreateEntityField(ctx, field); err != nil {
						return err
					}
				}
				return nil
			}

			// 查询是否存在实体
			eid, has := u.repo.GetEntityIdByName(ent.Database, ent.Name)

			// 存在则直接插入实体字段
			if has {
				if err := insertFields(eid, fields); err != nil {
					return err
				}
				continue
			}

			// 不存在则创建
			id, err := u.repo.CreateEntity(ctx, ent)
			if err != nil {
				return err
			}
			if err := insertFields(id, fields); err != nil {
				return err
			}
		}
		return nil
	})
}

// LoadEntity 获取租户列表
func (u *Entity) LoadEntity(ctx core.Context) ([]*entity.Entity, error) {
	selectTable := []string{
		"app",
		"dept",
		"dictionary",
		"feedback",
		"job",
		"user",
	}
	filterColumn := []string{
		"tenant_id",
		"deleted_at",
	}

	var list []*entity.Entity
	// 加载数据库的全部信息
	res := ctx.Database().Entities()
	for _, item := range res {
		if value.InList(selectTable, item.Name) {
			ent := &entity.Entity{
				Database: item.Database,
				Name:     item.Name,
				Comment:  item.Comment,
			}

			var fields []*entity.EntityField
			// 过滤字段
			for _, field := range item.Columns {
				if value.InList(filterColumn, field.Name) {
					continue
				}
				fields = append(fields, &entity.EntityField{
					Name:    field.Name,
					Comment: field.Comment,
				})
			}

			ent.Fields = fields
			list = append(list, ent)
		}
	}

	return list, nil
}

// ListEntity 获取租户列表
func (u *Entity) ListEntity(ctx core.Context, req *types.ListEntityRequest) ([]*entity.Entity, error) {
	list, err := u.repo.ListEntity(ctx, req)
	if err != nil {
		ctx.Logger().Warnw("msg", "list entity error", "err", err.Error())
		return nil, errors.ListError()
	}

	return list, nil
}

// CreateEntity 创建租户
func (u *Entity) CreateEntity(ctx core.Context, req *entity.Entity) (uint32, error) {
	id, err := u.repo.CreateEntity(ctx, req)
	if err != nil {
		ctx.Logger().Warnw("msg", "create entity error", "err", err.Error())
		return 0, errors.CreateError()
	}
	return id, nil
}

// UpdateEntity 更新租户
func (u *Entity) UpdateEntity(ctx core.Context, req *entity.Entity) error {
	if err := u.repo.UpdateEntity(ctx, req); err != nil {
		ctx.Logger().Warnw("msg", "update entity error", "err", err.Error())
		return errors.UpdateError()
	}
	return nil
}

// DeleteEntity 删除租户
func (u *Entity) DeleteEntity(ctx core.Context, id uint32) error {
	if err := u.repo.DeleteEntity(ctx, id); err != nil {
		ctx.Logger().Warnw("msg", "delete entity error", "err", err.Error())
		return errors.DeleteError()
	}
	return nil
}
