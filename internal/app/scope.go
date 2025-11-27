package app

import (
	"context"
	"encoding/json"

	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"

	"github.com/limes-cloud/manager/internal/infra/dbs"

	"github.com/limes-cloud/manager/api/scope"
	"github.com/limes-cloud/manager/internal/core"
	"github.com/limes-cloud/manager/internal/domain/service"
)

type Scope struct {
	scope.UnimplementedScopeServer
	srv *service.Scope
}

// NewScope 初始化角色应用
func NewScope() *Scope {
	return &Scope{
		srv: service.NewScope(
			dbs.NewScope(),
			dbs.NewUserDept(),
			dbs.NewTenantAdmin(),
		),
	}
}

// init 应用注册
func init() {
	register(func(hs *http.Server, gs *grpc.Server) {
		srv := NewScope()
		scope.RegisterScopeHTTPServer(hs, srv)
		scope.RegisterScopeServer(gs, srv)
	})
}

// GetScope 获取指定角色信息
func (app *Scope) GetScope(c context.Context, req *scope.GetScopeRequest) (*scope.GetScopeReply, error) {
	ctx := core.MustContext(c)

	// 调用服务
	ent, err := app.srv.Hook(ctx, req.Database, req.Model, req.Method)
	if err != nil {
		return nil, err
	}

	// 处理返回数据
	cond, _ := json.Marshal(ent.Condition())
	allDept, depts := ent.DeptScopes()

	reply := &scope.GetScopeReply{
		Rule:       string(cond),
		Fields:     ent.Fields(),
		AllDept:    allDept,
		DeptScopes: depts,
		DeptId:     ent.DeptId(),
		UserId:     ent.UserId(),
		TenantId:   ent.TenantId(),
	}
	return reply, nil
}
