package service

import (
	"fmt"
	"strings"

	"github.com/limes-cloud/kratosx/library/db/gormtranserror"
	"github.com/limes-cloud/kratosx/pkg/value"
	"gorm.io/gorm"

	"github.com/limes-cloud/manager/api/errors"
	"github.com/limes-cloud/manager/internal/core"
	"github.com/limes-cloud/manager/internal/domain/entity"
	"github.com/limes-cloud/manager/internal/domain/repository"
	"github.com/limes-cloud/manager/internal/types"
)

type Entity struct {
	repo repository.Entity
	app  repository.App
}

func NewEntity(repo repository.Entity, app ...repository.App) *Entity {
	srv := &Entity{repo: repo}
	if len(app) != 0 {
		srv.app = app[0]
	}
	return srv
}

const (
	RegisterEntityModeAppend = "append"
	RegisterEntityModeSync   = "sync"
)

type registerEntityState struct {
	App       string
	SeenEnt   map[string]struct{}
	SeenField map[uint32]map[string]struct{}
	SeenRule  map[uint32]map[string]struct{}
	Reply     *types.RegisterEntityReply
}

// ListEntityRule 获取租户列表
func (u *Entity) ListEntityRule(ctx core.Context, req *types.ListEntityRuleRequest) ([]*entity.EntityRule, uint32, error) {
	list, total, err := u.repo.ListEntityRule(ctx, req)
	if err != nil {
		ctx.Logger().Warnw("msg", "list entity rule error", "err", err.Error())
		return nil, 0, errors.ListError(err.Error())
	}
	return list, total, nil
}

// CreateEntityRule 创建租户
func (u *Entity) CreateEntityRule(ctx core.Context, req *entity.EntityRule) (uint32, error) {
	id, err := u.repo.CreateEntityRule(ctx, req)
	if err != nil {
		ctx.Logger().Warnw("msg", "create entity rule error", "err", err.Error())
		return 0, errors.CreateError(err.Error())
	}
	return id, nil
}

// UpdateEntityRule 更新租户
func (u *Entity) UpdateEntityRule(ctx core.Context, req *entity.EntityRule) error {
	if err := u.repo.UpdateEntityRule(ctx, req); err != nil {
		ctx.Logger().Warnw("msg", "update entity rule error", "err", err.Error())
		return errors.UpdateError(err.Error())
	}
	return nil
}

// DeleteEntityRule 删除租户
func (u *Entity) DeleteEntityRule(ctx core.Context, id uint32) error {
	if err := u.repo.DeleteEntityRule(ctx, id); err != nil {
		ctx.Logger().Warnw("msg", "delete entity rule error", "err", err.Error())
		return errors.DeleteError(err.Error())
	}
	return nil
}

// ListEntityField 获取租户列表
func (u *Entity) ListEntityField(ctx core.Context, req *types.ListEntityFieldRequest) ([]*entity.EntityField, uint32, error) {
	list, total, err := u.repo.ListEntityField(ctx, req)
	if err != nil {
		ctx.Logger().Warnw("msg", "list entity field error", "err", err.Error())
		return nil, 0, errors.ListError(err.Error())
	}
	return list, total, nil
}

// CreateEntityField 创建租户
func (u *Entity) CreateEntityField(ctx core.Context, req *entity.EntityField) (uint32, error) {
	id, err := u.repo.CreateEntityField(ctx, req)
	if err != nil {
		ctx.Logger().Warnw("msg", "create entity field error", "err", err.Error())
		return 0, errors.CreateError(err.Error())
	}
	return id, nil
}

// UpdateEntityField 更新租户
func (u *Entity) UpdateEntityField(ctx core.Context, req *entity.EntityField) error {
	if err := u.repo.UpdateEntityField(ctx, req); err != nil {
		ctx.Logger().Warnw("msg", "update entity field error", "err", err.Error())
		return errors.UpdateError(err.Error())
	}
	return nil
}

// DeleteEntityField 删除租户
func (u *Entity) DeleteEntityField(ctx core.Context, id uint32) error {
	if err := u.repo.DeleteEntityField(ctx, id); err != nil {
		ctx.Logger().Warnw("msg", "delete entity field error", "err", err.Error())
		return errors.DeleteError(err.Error())
	}
	return nil
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
			for _, field := range item.Fields {
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
		return nil, errors.ListError(err.Error())
	}

	return list, nil
}

// CreateEntity 创建租户
func (u *Entity) CreateEntity(ctx core.Context, req *entity.Entity) (uint32, error) {
	id, err := u.repo.CreateEntity(ctx, req)
	if err != nil {
		ctx.Logger().Warnw("msg", "create entity error", "err", err.Error())
		return 0, errors.CreateError(err.Error())
	}
	return id, nil
}

// UpdateEntity 更新租户
func (u *Entity) UpdateEntity(ctx core.Context, req *entity.Entity) error {
	if err := u.repo.UpdateEntity(ctx, req); err != nil {
		ctx.Logger().Warnw("msg", "update entity error", "err", err.Error())
		return errors.UpdateError(err.Error())
	}
	return nil
}

// DeleteEntity 删除租户
func (u *Entity) DeleteEntity(ctx core.Context, id uint32) error {
	if err := u.repo.DeleteEntity(ctx, id); err != nil {
		ctx.Logger().Warnw("msg", "delete entity error", "err", err.Error())
		return errors.DeleteError(err.Error())
	}
	return nil
}

// RegisterEntity 注册应用实体
func (u *Entity) RegisterEntity(ctx core.Context, req *types.RegisterEntityRequest) (*types.RegisterEntityReply, error) {
	if u.app == nil {
		return nil, errors.SystemError("应用仓储未初始化")
	}

	app, err := u.app.GetAppByKeyword(req.AppKeyword)
	if err != nil {
		ctx.Logger().Warnw("msg", "get app error", "err", err.Error())
		return nil, errors.GetError("应用不存在")
	}

	mode := strings.ToLower(strings.TrimSpace(req.Mode))
	if mode == "" {
		mode = RegisterEntityModeAppend
	}
	if mode != RegisterEntityModeAppend && mode != RegisterEntityModeSync {
		return nil, errors.ParamsError("mode仅支持append或sync")
	}

	state := &registerEntityState{
		App:       app.Keyword,
		SeenEnt:   map[string]struct{}{},
		SeenField: map[uint32]map[string]struct{}{},
		SeenRule:  map[uint32]map[string]struct{}{},
		Reply:     &types.RegisterEntityReply{},
	}
	if err = ctx.Transaction(func(ctx core.Context) error {
		for _, item := range req.Entities {
			if err := u.registerEntityNode(ctx, state, item); err != nil {
				return err
			}
		}
		if mode != RegisterEntityModeSync {
			return nil
		}
		return u.deleteUnregisteredEntities(ctx, state)
	}); err != nil {
		ctx.Logger().Warnw("msg", "register entity error", "err", err.Error())
		return nil, errors.UpdateError(err.Error())
	}

	return state.Reply, nil
}

func (u *Entity) registerEntityNode(
	ctx core.Context,
	state *registerEntityState,
	node *types.RegisterEntityNode,
) error {
	if node == nil {
		return nil
	}

	ent, err := u.toRegisterEntity(state.App, node)
	if err != nil {
		return err
	}

	entKey := u.registerEntityKey(ent.Database, ent.Name)
	if _, ok := state.SeenEnt[entKey]; ok {
		return fmt.Errorf("实体重复: %s", entKey)
	}
	state.SeenEnt[entKey] = struct{}{}

	old, err := u.repo.GetEntityByName(ctx, state.App, ent.Database, ent.Name)
	if gormtranserror.Is(err, gorm.ErrRecordNotFound) {
		old = nil
	} else if err != nil {
		return err
	}

	if old == nil {
		id, err := u.repo.CreateEntity(ctx, ent)
		if err != nil {
			return err
		}
		ent.Id = id
		state.Reply.Created++
	} else {
		ent.Id = old.Id
		if err := u.repo.UpdateRegisteredEntity(ctx, ent); err != nil {
			return err
		}
		state.Reply.Updated++
	}

	if err := u.registerEntityFields(ctx, state, ent.Id, node.Fields); err != nil {
		return err
	}
	return u.registerEntityRules(ctx, state, ent.Id, node.Rules)
}

func (u *Entity) registerEntityFields(
	ctx core.Context,
	state *registerEntityState,
	entityId uint32,
	list []*types.RegisterEntityFieldNode,
) error {
	for ind, item := range list {
		if item == nil {
			continue
		}
		name := strings.TrimSpace(item.Name)
		if name == "" {
			return fmt.Errorf("实体字段名称不能为空")
		}
		if _, ok := state.SeenField[entityId]; !ok {
			state.SeenField[entityId] = map[string]struct{}{}
		}
		if _, ok := state.SeenField[entityId][name]; ok {
			return fmt.Errorf("实体字段重复: %s", name)
		}
		state.SeenField[entityId][name] = struct{}{}

		field := &entity.EntityField{
			EntityId: entityId,
			Name:     name,
			Comment:  strings.TrimSpace(item.Comment),
			Index:    item.Index,
		}
		if field.Index == 0 {
			field.Index = ind
		}

		old, err := u.repo.GetEntityField(ctx, entityId, name)
		if gormtranserror.Is(err, gorm.ErrRecordNotFound) {
			_, err = u.repo.CreateEntityField(ctx, field)
			if err != nil {
				return err
			}
			state.Reply.Created++
			continue
		}
		if err != nil {
			return err
		}
		field.Id = old.Id
		if err := u.repo.UpdateRegisteredEntityField(ctx, field); err != nil {
			return err
		}
		state.Reply.Updated++
	}
	return nil
}

func (u *Entity) registerEntityRules(
	ctx core.Context,
	state *registerEntityState,
	entityId uint32,
	list []*types.RegisterEntityRuleNode,
) error {
	for _, item := range list {
		if item == nil {
			continue
		}
		name := strings.TrimSpace(item.Name)
		if name == "" {
			return fmt.Errorf("实体规则名称不能为空")
		}
		if _, ok := state.SeenRule[entityId]; !ok {
			state.SeenRule[entityId] = map[string]struct{}{}
		}
		if _, ok := state.SeenRule[entityId][name]; ok {
			return fmt.Errorf("实体规则重复: %s", name)
		}
		state.SeenRule[entityId][name] = struct{}{}

		rule := &entity.EntityRule{
			EntityId:    entityId,
			Name:        name,
			Status:      item.Status,
			Expression:  strings.TrimSpace(item.Expression),
			Description: strings.TrimSpace(item.Description),
		}
		old, err := u.repo.GetEntityRule(ctx, entityId, name)
		if gormtranserror.Is(err, gorm.ErrRecordNotFound) {
			_, err = u.repo.CreateEntityRule(ctx, rule)
			if err != nil {
				return err
			}
			state.Reply.Created++
			continue
		}
		if err != nil {
			return err
		}
		rule.Id = old.Id
		if err := u.repo.UpdateRegisteredEntityRule(ctx, rule); err != nil {
			return err
		}
		state.Reply.Updated++
	}
	return nil
}

func (u *Entity) toRegisterEntity(app string, node *types.RegisterEntityNode) (*entity.Entity, error) {
	ent := &entity.Entity{
		App:      app,
		Database: strings.TrimSpace(node.Database),
		Name:     strings.TrimSpace(node.Name),
		Comment:  strings.TrimSpace(node.Comment),
	}
	if ent.Database == "" || ent.Name == "" {
		return nil, fmt.Errorf("实体缺少database或name")
	}
	return ent, nil
}

func (u *Entity) deleteUnregisteredEntities(ctx core.Context, state *registerEntityState) error {
	list, err := u.repo.ListEntity(ctx, &types.ListEntityRequest{App: state.App})
	if err != nil {
		return err
	}

	var entityIds []uint32
	for _, item := range list {
		if _, ok := state.SeenEnt[u.registerEntityKey(item.Database, item.Name)]; ok {
			if err := u.deleteUnregisteredEntityChildren(ctx, state, item); err != nil {
				return err
			}
			continue
		}
		entityIds = append(entityIds, item.Id)
	}
	if len(entityIds) == 0 {
		return nil
	}
	if err := u.repo.DeleteEntities(ctx, entityIds); err != nil {
		return err
	}
	state.Reply.Deleted += uint32(len(entityIds))
	return nil
}

func (u *Entity) deleteUnregisteredEntityChildren(
	ctx core.Context,
	state *registerEntityState,
	ent *entity.Entity,
) error {
	fields, _, err := u.repo.ListEntityField(ctx, &types.ListEntityFieldRequest{EntityId: ent.Id})
	if err != nil {
		return err
	}
	var fieldIds []uint32
	for _, item := range fields {
		if _, ok := state.SeenField[ent.Id][item.Name]; !ok {
			fieldIds = append(fieldIds, item.Id)
		}
	}
	if err := u.repo.DeleteEntityFields(ctx, fieldIds); err != nil {
		return err
	}
	state.Reply.Deleted += uint32(len(fieldIds))

	rules, _, err := u.repo.ListEntityRule(ctx, &types.ListEntityRuleRequest{EntityId: value.Pointer(ent.Id)})
	if err != nil {
		return err
	}
	var ruleIds []uint32
	for _, item := range rules {
		if _, ok := state.SeenRule[ent.Id][item.Name]; !ok {
			ruleIds = append(ruleIds, item.Id)
		}
	}
	if err := u.repo.DeleteEntityRules(ctx, ruleIds); err != nil {
		return err
	}
	state.Reply.Deleted += uint32(len(ruleIds))
	return nil
}

func (u *Entity) registerEntityKey(db string, name string) string {
	return strings.ToLower(strings.TrimSpace(db)) + ":" + strings.ToLower(strings.TrimSpace(name))
}
