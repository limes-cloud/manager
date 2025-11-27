package core

import (
	"context"

	types2 "github.com/limes-cloud/manager/internal/types"

	"github.com/limes-cloud/kratosx"
	"github.com/limes-cloud/manager/api/errors"
)

type Context struct {
	kratosx.Context
}

func (c Context) Transaction(fn func(ctx Context) error, name ...string) error {
	return c.Context.Transaction(func(ctx kratosx.Context) error {
		return fn(Context{Context: ctx})
	}, name...)
}

func (Context) Config() *Conf {
	return conf
}

func (c Context) Auth() *types2.AuthorizeInfo {
	val, _ := c.Context.Value(types2.InfoKey).(*types2.AuthorizeInfo)
	if val == nil {
		c.Exit(errors.NotLoginError())
	}
	return val
}

func (c Context) Clone() Context {
	return MustContext(context.WithoutCancel(c.Context))
}

func MustContext(ctx context.Context, opts ...kratosx.ContextOptionFunc) Context {
	return Context{
		Context: kratosx.MustContext(ctx, opts...),
	}
}
