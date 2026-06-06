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
}

type resp struct {
	ctx      core.Context
	scope    *types.GetScopeResponse
	userdept repository.UserDept
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
) *Scope {
	return &Scope{
		repo:     repo,
		userdept: userdept,
	}
}

func (h *Scope) Hook(ctx context.Context, database string, model string, method string) (mhook.ScopeResponse, error) {
	return &resp{
		ctx:      core.MustContext(ctx),
		scope:    h.repo.GetScope(core.MustContext(ctx), database, model, method),
		userdept: h.userdept,
	}, nil
}

func (r *resp) DeptScopes() (bool, []uint32) {
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

