package service

import (
	"context"

	"github.com/limes-cloud/manager/internal/types"

	mhook "github.com/limes-cloud/kratosx/model/hook"
	"github.com/limes-cloud/manager/internal/core"
	"github.com/limes-cloud/manager/internal/domain/repository"
)

type Scope struct {
	repo     repository.Scope
	userdept repository.UserDept
	tad      repository.TenantAdmin
}

type resp struct {
	ctx      core.Context
	scope    *types.GetScopeResponse
	userdept repository.UserDept
	tad      repository.TenantAdmin
}

func (r *resp) Condition() *mhook.ConditionGroup {
	return nil
	// return r.scope.Rule
}

func (r *resp) Fields() []string {
	return []string{}
	// return r.scope.Fields
}

func NewScope(
	repo repository.Scope,
	userdept repository.UserDept,
	tad repository.TenantAdmin,
) *Scope {
	return &Scope{
		repo:     repo,
		userdept: userdept,
		tad:      tad,
	}
}

func (h *Scope) Hook(ctx context.Context, database string, model string, method string) (mhook.ScopeResponse, error) {
	return &resp{
		ctx:      core.MustContext(ctx),
		scope:    h.repo.GetScope(core.MustContext(ctx), database, model, method),
		userdept: h.userdept,
		tad:      h.tad,
	}, nil
}

func (r *resp) DeptScopes() (bool, []uint32) {
	if r.tad.IsAdmin(r.TenantId(), r.UserId()) {
		return true, nil
	}
	// 获取实体信息
	return r.scope.AllDept, r.scope.DeptScopes
}

func (r *resp) UserDeptId(uid uint32) uint32 {
	return r.userdept.GetUserMainDeptId(uid)
}

func (r *resp) DeptId() uint32 {
	return r.ctx.Auth().DeptId
}

func (r *resp) UserId() uint32 {
	return r.ctx.Auth().UserId
}

func (r *resp) TenantId() uint32 {
	return r.ctx.Auth().TenantId
}
