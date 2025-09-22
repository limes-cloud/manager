package hook

import (
	"context"

	"github.com/limes-cloud/manager/internal/domain/repository"

	"github.com/limes-cloud/manager/internal/domain/middleware/auth"

	mhook "github.com/limes-cloud/kratosx/model/hook"
	"github.com/limes-cloud/manager/internal/core"
)

type hook struct {
	ta  repository.TenantApp
	jr  repository.JobRole
	dr  repository.DeptRole
	ud  repository.UserDept
	re  repository.RoleEntity
	ent repository.Entity
}

type resp struct {
	ctx      core.Context
	database string
	model    string
	method   string
	hook     *hook
}

func (r *resp) UserDeptId(u uint32) uint32 {
	// TODO implement me
	panic("implement me")
}

type Hook interface {
	Hook(ctx context.Context, database string, model string, method string) (mhook.ScopeResponse, error)
}

func New() Hook {
	return &hook{}
}

func (h *hook) Hook(ctx context.Context, database string, model string, method string) (mhook.ScopeResponse, error) {
	if !auth.HasLogin(ctx) {
		return nil, nil
	}
	return &resp{
		ctx:      core.MustContext(ctx),
		database: database,
		model:    model,
		method:   method,
		hook:     h,
	}, nil
}

func (r *resp) DeptScopes() (bool, []uint32) {
	return false, []uint32{r.ctx.Auth().DeptId}
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
