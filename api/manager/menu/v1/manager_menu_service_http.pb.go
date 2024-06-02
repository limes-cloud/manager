// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.7.3
// - protoc             v4.24.4
// source: api/manager/menu/manager_menu_service.proto

package v1

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

const OperationMenuCreateMenu = "/manager.api.manager.menu.v1.Menu/CreateMenu"
const OperationMenuDeleteMenu = "/manager.api.manager.menu.v1.Menu/DeleteMenu"
const OperationMenuListMenu = "/manager.api.manager.menu.v1.Menu/ListMenu"
const OperationMenuListMenuByCurRole = "/manager.api.manager.menu.v1.Menu/ListMenuByCurRole"
const OperationMenuUpdateMenu = "/manager.api.manager.menu.v1.Menu/UpdateMenu"

type MenuHTTPServer interface {
	// CreateMenu CreateMenu 创建菜单信息
	CreateMenu(context.Context, *CreateMenuRequest) (*CreateMenuReply, error)
	// DeleteMenu DeleteMenu 删除菜单信息
	DeleteMenu(context.Context, *DeleteMenuRequest) (*DeleteMenuReply, error)
	// ListMenu ListMenu 获取菜单信息列表
	ListMenu(context.Context, *ListMenuRequest) (*ListMenuReply, error)
	// ListMenuByCurRole ListMenuByCurRole 获取菜单信息列表
	ListMenuByCurRole(context.Context, *ListMenuByCurRoleRequest) (*ListMenuByCurRoleReply, error)
	// UpdateMenu UpdateMenu 更新菜单信息
	UpdateMenu(context.Context, *UpdateMenuRequest) (*UpdateMenuReply, error)
}

func RegisterMenuHTTPServer(s *http.Server, srv MenuHTTPServer) {
	r := s.Route("/")
	r.GET("/manager/api/v1/menus", _Menu_ListMenu0_HTTP_Handler(srv))
	r.GET("/manager/api/v1/menus/by/cur_role", _Menu_ListMenuByCurRole0_HTTP_Handler(srv))
	r.POST("/manager/api/v1/menu", _Menu_CreateMenu0_HTTP_Handler(srv))
	r.PUT("/manager/api/v1/menu", _Menu_UpdateMenu0_HTTP_Handler(srv))
	r.DELETE("/manager/api/v1/menu", _Menu_DeleteMenu0_HTTP_Handler(srv))
}

func _Menu_ListMenu0_HTTP_Handler(srv MenuHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ListMenuRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationMenuListMenu)
		h := ctx.Middleware(func(ctx context.Context, req any) (any, error) {
			return srv.ListMenu(ctx, req.(*ListMenuRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ListMenuReply)
		return ctx.Result(200, reply)
	}
}

func _Menu_ListMenuByCurRole0_HTTP_Handler(srv MenuHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ListMenuByCurRoleRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationMenuListMenuByCurRole)
		h := ctx.Middleware(func(ctx context.Context, req any) (any, error) {
			return srv.ListMenuByCurRole(ctx, req.(*ListMenuByCurRoleRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ListMenuByCurRoleReply)
		return ctx.Result(200, reply)
	}
}

func _Menu_CreateMenu0_HTTP_Handler(srv MenuHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CreateMenuRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationMenuCreateMenu)
		h := ctx.Middleware(func(ctx context.Context, req any) (any, error) {
			return srv.CreateMenu(ctx, req.(*CreateMenuRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*CreateMenuReply)
		return ctx.Result(200, reply)
	}
}

func _Menu_UpdateMenu0_HTTP_Handler(srv MenuHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UpdateMenuRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationMenuUpdateMenu)
		h := ctx.Middleware(func(ctx context.Context, req any) (any, error) {
			return srv.UpdateMenu(ctx, req.(*UpdateMenuRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UpdateMenuReply)
		return ctx.Result(200, reply)
	}
}

func _Menu_DeleteMenu0_HTTP_Handler(srv MenuHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in DeleteMenuRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationMenuDeleteMenu)
		h := ctx.Middleware(func(ctx context.Context, req any) (any, error) {
			return srv.DeleteMenu(ctx, req.(*DeleteMenuRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*DeleteMenuReply)
		return ctx.Result(200, reply)
	}
}

type MenuHTTPClient interface {
	CreateMenu(ctx context.Context, req *CreateMenuRequest, opts ...http.CallOption) (rsp *CreateMenuReply, err error)
	DeleteMenu(ctx context.Context, req *DeleteMenuRequest, opts ...http.CallOption) (rsp *DeleteMenuReply, err error)
	ListMenu(ctx context.Context, req *ListMenuRequest, opts ...http.CallOption) (rsp *ListMenuReply, err error)
	ListMenuByCurRole(ctx context.Context, req *ListMenuByCurRoleRequest, opts ...http.CallOption) (rsp *ListMenuByCurRoleReply, err error)
	UpdateMenu(ctx context.Context, req *UpdateMenuRequest, opts ...http.CallOption) (rsp *UpdateMenuReply, err error)
}

type MenuHTTPClientImpl struct {
	cc *http.Client
}

func NewMenuHTTPClient(client *http.Client) MenuHTTPClient {
	return &MenuHTTPClientImpl{client}
}

func (c *MenuHTTPClientImpl) CreateMenu(ctx context.Context, in *CreateMenuRequest, opts ...http.CallOption) (*CreateMenuReply, error) {
	var out CreateMenuReply
	pattern := "/manager/api/v1/menu"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationMenuCreateMenu))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *MenuHTTPClientImpl) DeleteMenu(ctx context.Context, in *DeleteMenuRequest, opts ...http.CallOption) (*DeleteMenuReply, error) {
	var out DeleteMenuReply
	pattern := "/manager/api/v1/menu"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationMenuDeleteMenu))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "DELETE", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *MenuHTTPClientImpl) ListMenu(ctx context.Context, in *ListMenuRequest, opts ...http.CallOption) (*ListMenuReply, error) {
	var out ListMenuReply
	pattern := "/manager/api/v1/menus"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationMenuListMenu))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *MenuHTTPClientImpl) ListMenuByCurRole(ctx context.Context, in *ListMenuByCurRoleRequest, opts ...http.CallOption) (*ListMenuByCurRoleReply, error) {
	var out ListMenuByCurRoleReply
	pattern := "/manager/api/v1/menus/by/cur_role"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationMenuListMenuByCurRole))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *MenuHTTPClientImpl) UpdateMenu(ctx context.Context, in *UpdateMenuRequest, opts ...http.CallOption) (*UpdateMenuReply, error) {
	var out UpdateMenuReply
	pattern := "/manager/api/v1/menu"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationMenuUpdateMenu))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PUT", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
