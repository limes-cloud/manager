package dbs

import (
	"context"
	"encoding/json"
	"sync"

	"github.com/limes-cloud/kratosx/model/hook"

	"github.com/limes-cloud/manager/internal/types"

	"github.com/samber/lo"

	"github.com/limes-cloud/manager/internal/domain/entity"

	"github.com/limes-cloud/manager/internal/core"
)

type Scope struct {
	database string
	ta       *TenantApp
	jr       *JobRole
	dr       *DeptRole
	ud       *UserDept
	re       *RoleEntity
	rm       *RoleMenu
	ent      *Entity
	dept     *Dept
	role     *Role
}

var (
	scopeIns  *Scope
	scopeOnce sync.Once
)

func NewScope() *Scope {
	scopeOnce.Do(func() {
		scopeIns = &Scope{
			ta:   NewTenantApp(),
			jr:   NewJobRole(),
			dr:   NewDeptRole(),
			ud:   NewUserDept(),
			re:   NewRoleEntity(),
			rm:   NewRoleMenu(),
			ent:  NewEntity(),
			dept: NewDept(),
			role: NewRole(),
		}

		ctx := core.MustContext(context.Background())
		ctx.BeforeStart("init database", func() {
			scopeIns.database = ctx.DB().Migrator().CurrentDatabase()
		})
	})
	return scopeIns
}

// HasAppScope 获取用户是否具有指定角色的权限
func (s Scope) HasAppScope(ctx core.Context, roleId uint32) bool {
	return s.ta.HasAppScope(ctx.Auth().TenantId, roleId)
}

// HasRoleScope 获取用户是否具有指定角色的权限
func (s Scope) HasRoleScope(ctx core.Context, roleId uint32) bool {
	ids := s.RoleScopes(ctx)
	return lo.Contains(ids, roleId)
}

func (s Scope) RoleScopes(ctx core.Context) []uint32 {
	// 获取当前的所有角色
	rids := s.RoleIds(ctx)
	ids, _ := s.role.GetRoleChildrenIds(ctx, rids)
	return append(rids, ids...)
}

// GetScope 获取的规则
func (s Scope) GetScope(ctx core.Context, database string, entity string, action string) *types.GetScopeResponse {
	reply := &types.GetScopeResponse{}
	eid, has := s.ent.GetEntityIdByName(database, entity)
	if !has {
		return reply
	}

	// 获取当前的的角色id
	rids, drm := s.getRoleDept(ctx)

	var (
		scopes []uint32
		fields []string
		rule   = &hook.ConditionGroup{
			Logic: "OR",
		}
	)
	for _, rid := range rids {
		roleEntity, has := s.re.GetRoleEntityByRDA(rid, eid, action)
		if !has {
			continue
		}

		// 获取部门权限列表
		if !reply.AllDept {
			all, deptIds := s.deptScope(drm[rid], roleEntity)
			scopes = append(scopes, deptIds...)
			reply.AllDept = all
		}

		// 获取字段
		var tfs []uint32
		_ = json.Unmarshal([]byte(roleEntity.Fields), &tfs)
		for _, tf := range tfs {
			fn, ok := s.ent.GetEntityFieldName(tf)
			if !ok {
				continue
			}
			fields = append(fields, fn)
		}

		// 获取规则
		var rfs []uint32
		_ = json.Unmarshal([]byte(roleEntity.Rules), &rfs)
		for _, rf := range rfs {
			rn, ok := s.ent.GetEntityRuleById(rf)
			if !ok {
				continue
			}
			tmp := hook.ConditionGroup{}
			_ = json.Unmarshal([]byte(rn), &tmp)
			rule.Groups = append(rule.Groups, &tmp)
		}
	}

	reply.DeptScopes = lo.Uniq(scopes)
	reply.Fields = lo.Uniq(fields)
	reply.Rule = rule
	if reply.AllDept {
		reply.DeptScopes = nil
	}

	return reply
}

// SystemDeptScopes 获取当前操
func (s Scope) SystemDeptScopes(ctx core.Context, entity string, action string) (bool, []uint32) {
	// 获取当前系统的数据库名称
	eid, has := s.ent.GetEntityIdByName(s.database, entity)
	if !has {
		return false, nil
	}

	// 获取当前的角色id
	rids, drm := s.getRoleDept(ctx)

	// 获取这些角色-实体关系
	var scopes []uint32
	for _, rid := range rids {
		roleEntity, has := s.re.GetRoleEntityByRDA(rid, eid, action)
		if !has {
			continue
		}

		all, deptIds := s.deptScope(drm[rid], roleEntity)
		if all {
			return true, nil
		}
		scopes = append(scopes, deptIds...)
	}

	return false, scopes
}

// DeptScopes 获取当前操
func (s Scope) DeptScopes(ctx core.Context, database string, entity string, action string) (bool, []uint32) {
	eid, has := s.ent.GetEntityIdByName(database, entity)
	if !has {
		return false, nil
	}

	// 获取当前的角色id
	rids, drm := s.getRoleDept(ctx)

	// 获取这些角色-实体关系
	var scopes []uint32
	for _, rid := range rids {
		roleEntity, has := s.re.GetRoleEntityByRDA(rid, eid, action)
		if !has {
			continue
		}

		all, deptIds := s.deptScope(drm[rid], roleEntity)
		if all {
			return true, nil
		}
		scopes = append(scopes, deptIds...)
	}

	return false, scopes
}

func (s Scope) deptScope(deptIds []uint32, re *entity.RoleEntity) (bool, []uint32) {
	switch re.Scope {
	case "ALL": // 所有部门
		return true, nil
	case "CUR": // 当前部门
		return false, deptIds
	case "CUR_DOWN": // 当前及以下部门
		ids := s.dept.GetDeptChildrenIds(deptIds)
		return false, append(deptIds, ids...)
	case "DOWN": // 以下部门
		ids := s.dept.GetDeptChildrenIds(deptIds)
		return false, ids
	case "CUSTOM": // 自定义部门
		var ids []uint32
		_ = json.Unmarshal([]byte(re.Depts), &ids)
		return false, ids
	case "SELF", "NONE": // 自己部门
		return false, nil
	default:
		return false, nil
	}
}

func (s Scope) getRoleDept(ctx core.Context) ([]uint32, map[uint32][]uint32) {
	var (
		jdm  = map[uint32][]uint32{}
		dids = []uint32{ctx.Auth().DeptId}
		jids = []uint32{ctx.Auth().JobId}
	)
	uds := s.ud.ListUserDeptByUserId(ctx.Auth().UserId)
	for _, item := range uds {
		jdm[item.JobId] = append(jdm[item.JobId], item.DeptId)
		dids = append(dids, item.DeptId)
		jids = append(jids, item.JobId)
	}

	var (
		rids []uint32
		drm  = map[uint32][]uint32{}
	)

	// 获取部门-角色
	drs := s.dr.GetDeptRolesByDeptIds(dids)
	for _, item := range drs {
		drm[item.RoleId] = append(drm[item.RoleId], item.DeptId)
		rids = append(rids, item.RoleId)
	}

	// 获取职位-角色
	jrs := s.jr.GetJobRolesByJobIds(jids)
	for _, item := range jrs {
		drm[item.RoleId] = append(drm[item.RoleId], jdm[item.JobId]...)
		rids = append(rids, item.RoleId)
	}

	return rids, drm
}

// RoleIds 获取当前用户的角色id列表
func (s Scope) RoleIds(ctx core.Context) []uint32 {
	// 获取当前的部门id
	var dids, jids []uint32
	list := s.ud.ListUserDeptByUserId(ctx.Auth().UserId)
	for _, item := range list {
		dids = append(dids, item.DeptId)
		jids = append(jids, item.JobId)
	}

	dids = append(dids, ctx.Auth().DeptId)
	jids = append(jids, ctx.Auth().JobId)

	// 获取部门绑定的角色
	rids := s.dr.GetRoleIdsByDeptIds(dids)

	// 获取当前的职位
	rids = append(rids, s.jr.GetRoleIdsByJobIds(jids)...)

	return rids
}

// DeptIds 获取当前用户的部门id列表
func (s Scope) DeptIds(ctx core.Context) []uint32 {
	ids := s.ud.GetDeptIdsByUserId(ctx.Auth().UserId)
	return append(ids, ctx.Auth().DeptId)
}

func (s Scope) HasMenuScope(ctx core.Context, menuId uint32) bool {
	rids := s.RoleIds(ctx)
	for _, rid := range rids {
		if s.rm.hasRoleMenu(rid, menuId) {
			return true
		}
	}
	return false
}
