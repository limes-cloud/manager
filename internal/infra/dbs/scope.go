package dbs

import (
	"sync"

	"github.com/limes-cloud/manager/internal/core"
)

type Scope struct {
	ta  *TenantApp
	jr  *JobRole
	dr  *DeptRole
	ud  *UserDept
	re  *RoleEntity
	ent *Entity
}

func (s Scope) IsSuperAdmin(ctx core.Context) bool {
	return true
}

var (
	scopeIns  *Scope
	scopeOnce sync.Once
)

func NewScope() *Scope {
	scopeOnce.Do(func() {
		scopeIns = &Scope{
			ta: NewTenantApp(),
			jr: NewJobRole(),
			dr: NewDeptRole(),
			ud: NewUserDept(),
		}
	})
	return scopeIns
}

func (s Scope) AppScope(_ core.Context, tid uint32) []uint32 {
	return s.ta.GetAppIds(tid)
}

// HasAppScope 判断当前用户是否具有应用权限
func (s Scope) HasAppScope(ctx core.Context, appId uint32) bool {
	tid := ctx.Auth().TenantId
	return s.ta.HasAppScope(tid, appId)
}

// HasRoleScope 获取用户是否具有指定角色的权限
func (s Scope) HasRoleScope(ctx core.Context, roleId uint32) bool {
	// 用户角色
	// 部门角色
	// 职位角色
	return true
}

func (s Scope) RoleScopes(ctx core.Context) []uint32 {
	return []uint32{1}
}

func (s Scope) HasDeptScope(ctx core.Context, entity string, deptId uint32) bool {
	// db := ctx.DB().Migrator().CurrentDatabase()
	// name := ctx.Context.Config().App().Database[0].Connect.DBName
	// s.re.GetRoleEntity()

	return true
	// TODO implement me
	panic("implement me")
}

// func (s Scope) defaultDeptScopes()
// db+name => entity
// rule => entity
func (s Scope) DeptScopes(ctx core.Context, entity string) (bool, []uint32) {
	return true, nil
	// return s.ud.GetDeptIdsByUserId(ctx.Auth().UserId)

	//rdm := s.getRoleDept(ctx)
	//for rid, deptIds := range rdm {
	//	s.ent.GetEntityByName(ctx, "user_center", entity)
	//}
}

//func (s Scope) DeptScopes(ctx core.Context, database string, entity string, action string) (bool, []uint32) {
//	return true, nil
//	// return s.ud.GetDeptIdsByUserId(ctx.Auth().UserId)
//
//	//rdm := s.getRoleDept(ctx)
//	//for rid, deptIds := range rdm {
//	//	s.ent.GetEntityByName(ctx, "user_center", entity)
//	//}
//}

//func (s Scope) getRoleDept(ctx core.Context) map[uint32][]uint32 {
//	// roleid => deptids
//}

// RoleIds 获取当前的角色id列表
func (s Scope) RoleIds(ctx core.Context) []uint32 {
	// 获取当前的部门id
	dids, jids := s.ud.GetDeptsAndRolesByUserId(ctx.Auth().UserId)
	dids = append(dids, ctx.Auth().DeptId)
	jids = append(jids, ctx.Auth().JobId)

	// 获取部门绑定的角色
	rids := s.dr.GetRoleIdsByDeptIds(dids)

	// 获取当前的职位
	rids = append(rids, s.jr.GetRoleIdsByJobIds(dids)...)

	return rids
}

// DeptIds 获取当前用户的部门id
func (s Scope) DeptIds(ctx core.Context) []uint32 {
	return s.ud.GetDeptIdsByUserId(ctx.Auth().UserId)
}
