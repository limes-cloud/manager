package scope

import (
	"context"
	"encoding/json"

	"github.com/limes-cloud/kratosx"

	"github.com/limes-cloud/kratosx/model/hook"
)

type h struct {
	ctx    kratosx.Context
	client ScopeClient
	db     string
	model  string
	method string
	data   *scopeData
}

type scopeData struct {
	AllDept    bool
	DeptScopes []uint32
	Rule       *hook.ConditionGroup
	Fields     []string
	UserId     uint32
	DeptId     uint32
	TenantId   uint32
}

func Hook(c context.Context, db, model, method string) (hook.ScopeResponse, error) {
	ctx := kratosx.MustContext(c)
	conn, err := ctx.GrpcConn("Manager")
	if err != nil {
		return nil, err
	}
	ins := &h{
		client: NewScopeClient(conn),
		db:     db,
		model:  model,
		method: method,
	}

	reply, err := ins.client.GetScope(ctx, &GetScopeRequest{
		Database: db,
		Model:    model,
		Method:   method,
	})
	if err != nil {
		return nil, err
	}

	rules := hook.ConditionGroup{}
	_ = json.Unmarshal([]byte(reply.Rule), &rules)

	ins.data = &scopeData{
		AllDept:    reply.AllDept,
		DeptScopes: reply.DeptScopes,
		Rule:       &rules,
		Fields:     reply.Fields,
		UserId:     reply.UserId,
		DeptId:     reply.DeptId,
		TenantId:   reply.TenantId,
	}

	return ins, nil
}

func (r *h) Condition() *hook.ConditionGroup {
	return r.data.Rule
}

func (r *h) Fields() []string {
	return r.data.Fields
}

func (r *h) DeptScopes() (bool, []uint32) {
	// 获取实体信息
	return r.data.AllDept, r.data.DeptScopes
}

func (r *h) UserDeptId(u uint32) uint32 {
	reply, err := r.client.GetUserDeptId(r.ctx, &GetUserDeptIdRequest{
		UserId: u,
	})
	if err != nil {
		return 0
	}

	return reply.DeptId
}

func (r *h) DeptId() uint32 {
	return r.data.DeptId
}

func (r *h) UserId() uint32 {
	return r.data.UserId
}

func (r *h) TenantId() uint32 {
	return r.data.TenantId
}
