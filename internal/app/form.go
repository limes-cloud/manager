package app

import (
	"context"

	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/limes-cloud/kratosx/model/page"

	"github.com/limes-cloud/manager/api/form"
	"github.com/limes-cloud/manager/internal/core"
	"github.com/limes-cloud/manager/internal/domain/entity"
	"github.com/limes-cloud/manager/internal/domain/service"
	"github.com/limes-cloud/manager/internal/infra/dbs"
	"github.com/limes-cloud/manager/internal/types"
)

type Form struct {
	form.UnimplementedFormServer
	srv        *service.Form
	sceneSrv   *service.FormScene
	appFormSrv *service.AppForm
}

func NewForm() *Form {
	userFormRepo := dbs.NewUserForm()
	return &Form{
		srv:        service.NewForm(userFormRepo),
		sceneSrv:   service.NewFormScene(dbs.NewFormScene()),
		appFormSrv: service.NewAppForm(dbs.NewAppForm(), userFormRepo),
	}
}

func init() {
	register(func(hs *http.Server, gs *grpc.Server) {
		srv := NewForm()
		form.RegisterFormHTTPServer(hs, srv)
		form.RegisterFormServer(gs, srv)
	})
}

// ListFormScene 查询场景列表
func (h *Form) ListFormScene(c context.Context, req *form.ListFormSceneRequest) (*form.ListFormSceneResponse, error) {
	ctx := core.MustContext(c)

	in := &types.ListFormSceneRequest{
		Search: &page.Search{
			Page:     req.Page,
			PageSize: req.PageSize,
		},
	}
	if req.Keyword != nil {
		in.Keyword = req.Keyword
	}
	if req.Name != nil {
		in.Name = req.Name
	}

	list, total, err := h.sceneSrv.ListFormScene(ctx, in)
	if err != nil {
		return nil, err
	}

	reply := &form.ListFormSceneResponse{Total: total}
	for _, item := range list {
		reply.List = append(reply.List, toProtoFormScene(item))
	}
	return reply, nil
}

// CreateFormScene 创建场景
func (h *Form) CreateFormScene(c context.Context, req *form.CreateFormSceneRequest) (*form.CreateFormSceneResponse, error) {
	ctx := core.MustContext(c)

	ent := &entity.FormScene{
		Keyword: req.Keyword,
		Name:    req.Name,
	}
	if req.Description != nil {
		ent.Description = req.Description
	}

	id, err := h.sceneSrv.CreateFormScene(ctx, ent)
	if err != nil {
		return nil, err
	}
	return &form.CreateFormSceneResponse{Id: id}, nil
}

// UpdateFormScene 更新场景
func (h *Form) UpdateFormScene(c context.Context, req *form.UpdateFormSceneRequest) (*form.UpdateFormSceneResponse, error) {
	ctx := core.MustContext(c)

	ent := &entity.FormScene{}
	ent.Id = req.Id
	if req.Keyword != nil {
		ent.Keyword = *req.Keyword
	}
	if req.Name != nil {
		ent.Name = *req.Name
	}
	if req.Description != nil {
		ent.Description = req.Description
	}

	if err := h.sceneSrv.UpdateFormScene(ctx, ent); err != nil {
		return nil, err
	}
	return &form.UpdateFormSceneResponse{}, nil
}

// DeleteFormScene 删除场景
func (h *Form) DeleteFormScene(c context.Context, req *form.DeleteFormSceneRequest) (*form.DeleteFormSceneResponse, error) {
	ctx := core.MustContext(c)
	if err := h.sceneSrv.DeleteFormScene(ctx, req.Id); err != nil {
		return nil, err
	}
	return &form.DeleteFormSceneResponse{}, nil
}

// ListUserForm 分页查询表单列表
func (h *Form) ListUserForm(c context.Context, req *form.ListUserFormRequest) (*form.ListUserFormResponse, error) {
	ctx := core.MustContext(c)

	in := &types.ListUserFormRequest{
		Search: &page.Search{
			Page:     req.Page,
			PageSize: req.PageSize,
		},
	}
	if req.Scene != nil {
		in.Scene = req.Scene
	}
	if req.Name != nil {
		in.Name = req.Name
	}
	if req.Status != nil {
		in.Status = req.Status
	}

	list, total, err := h.srv.ListUserForm(ctx, in)
	if err != nil {
		return nil, err
	}

	reply := &form.ListUserFormResponse{Total: total}
	for _, item := range list {
		reply.List = append(reply.List, toProtoUserForm(item))
	}
	return reply, nil
}

// GetUserForm 获取单个表单
func (h *Form) GetUserForm(c context.Context, req *form.GetUserFormRequest) (*form.GetUserFormResponse, error) {
	ctx := core.MustContext(c)

	in := &types.GetUserFormRequest{
		Id:    req.Id,
		Scene: req.Scene,
	}

	res, err := h.srv.GetUserForm(ctx, in)
	if err != nil {
		return nil, err
	}

	return &form.GetUserFormResponse{Data: toProtoUserForm(res)}, nil
}

// CreateUserForm 创建表单
func (h *Form) CreateUserForm(c context.Context, req *form.CreateUserFormRequest) (*form.CreateUserFormResponse, error) {
	ctx := core.MustContext(c)

	ent := &entity.UserForm{
		Scene:  req.Scene,
		Name:   req.Name,
		Schema: req.Schema,
	}
	if req.Description != nil {
		ent.Description = req.Description
	}

	id, err := h.srv.CreateUserForm(ctx, ent)
	if err != nil {
		return nil, err
	}
	return &form.CreateUserFormResponse{Id: id}, nil
}

// UpdateUserForm 更新表单
func (h *Form) UpdateUserForm(c context.Context, req *form.UpdateUserFormRequest) (*form.UpdateUserFormResponse, error) {
	ctx := core.MustContext(c)

	ent := &entity.UserForm{}
	ent.Id = req.Id
	if req.Scene != nil {
		ent.Scene = *req.Scene
	}
	if req.Name != nil {
		ent.Name = *req.Name
	}
	if req.Description != nil {
		ent.Description = req.Description
	}
	if req.Schema != nil {
		ent.Schema = *req.Schema
	}

	if err := h.srv.UpdateUserForm(ctx, ent); err != nil {
		return nil, err
	}
	return &form.UpdateUserFormResponse{}, nil
}

// PublishUserForm 发布表单
func (h *Form) PublishUserForm(c context.Context, req *form.PublishUserFormRequest) (*form.PublishUserFormResponse, error) {
	ctx := core.MustContext(c)
	if err := h.srv.PublishUserForm(ctx, req.Id); err != nil {
		return nil, err
	}
	return &form.PublishUserFormResponse{}, nil
}

// DeleteUserForm 删除表单
func (h *Form) DeleteUserForm(c context.Context, req *form.DeleteUserFormRequest) (*form.DeleteUserFormResponse, error) {
	ctx := core.MustContext(c)
	if err := h.srv.DeleteUserForm(ctx, req.Id); err != nil {
		return nil, err
	}
	return &form.DeleteUserFormResponse{}, nil
}

// ListAppForm 查询某应用下的表单关联列表
func (h *Form) ListAppForm(c context.Context, req *form.ListAppFormRequest) (*form.ListAppFormResponse, error) {
	ctx := core.MustContext(c)

	in := &types.ListAppFormRequest{
		AppKeyword: req.AppKeyword,
	}

	list, err := h.appFormSrv.ListAppForm(ctx, in)
	if err != nil {
		return nil, err
	}

	reply := &form.ListAppFormResponse{}
	for _, item := range list {
		reply.List = append(reply.List, &form.AppFormItem{
			Id:         item.Id,
			AppKeyword: item.AppKeyword,
			FormId:     item.FormId,
			Scene:      item.Scene,
			CreatedAt:  uint32(item.CreatedAt),
		})
	}
	return reply, nil
}

// CreateAppForm 分配表单到应用
func (h *Form) CreateAppForm(c context.Context, req *form.CreateAppFormRequest) (*form.CreateAppFormResponse, error) {
	ctx := core.MustContext(c)

	ent := &entity.AppForm{
		AppKeyword: req.AppKeyword,
		FormId:     req.FormId,
	}

	id, err := h.appFormSrv.CreateAppForm(ctx, ent)
	if err != nil {
		return nil, err
	}
	return &form.CreateAppFormResponse{Id: id}, nil
}

// DeleteAppForm 取消分配
func (h *Form) DeleteAppForm(c context.Context, req *form.DeleteAppFormRequest) (*form.DeleteAppFormResponse, error) {
	ctx := core.MustContext(c)
	if err := h.appFormSrv.DeleteAppForm(ctx, req.Id); err != nil {
		return nil, err
	}
	return &form.DeleteAppFormResponse{}, nil
}

// GetUserFormByApp 通过应用和场景获取表单（运行时用）
func (h *Form) GetUserFormByApp(c context.Context, req *form.GetUserFormByAppRequest) (*form.GetUserFormByAppResponse, error) {
	ctx := core.MustContext(c)

	uf, err := h.appFormSrv.GetUserFormByApp(ctx, req.AppKeyword, req.Scene)
	if err != nil {
		return nil, err
	}

	return &form.GetUserFormByAppResponse{Data: toProtoUserForm(uf)}, nil
}

// FetchRemoteOptions 远程数据代理（表单编辑器远程选项加载）
func (h *Form) FetchRemoteOptions(c context.Context, req *form.FetchRemoteOptionsRequest) (*form.FetchRemoteOptionsResponse, error) {
	ctx := core.MustContext(c)

	srvReq := &service.FetchRemoteOptionsRequest{
		URL:        req.Url,
		Method:     req.Method,
		DataPath:   req.DataPath,
		LabelField: req.LabelField,
		ValueField: req.ValueField,
		Headers:    req.Headers,
		Params:     req.Params,
		Body:       req.Body,
	}

	items, err := h.srv.FetchRemoteOptions(ctx, srvReq)
	if err != nil {
		return nil, err
	}

	reply := &form.FetchRemoteOptionsResponse{
		Items: make([]*form.RemoteOptionItem, 0, len(items)),
	}
	for _, item := range items {
		reply.Items = append(reply.Items, &form.RemoteOptionItem{
			Label: item.Label,
			Value: item.Value,
		})
	}
	return reply, nil
}

// toProtoUserForm entity → proto 转换
func toProtoUserForm(ent *entity.UserForm) *form.UserFormItem {
	item := &form.UserFormItem{
		Id:        ent.Id,
		Scene:     ent.Scene,
		Name:      ent.Name,
		Schema:    ent.Schema,
		Version:   ent.Version,
		CreatedAt: uint32(ent.CreatedAt),
		UpdatedAt: uint32(ent.UpdatedAt),
	}
	if ent.Status != nil {
		item.Status = *ent.Status
	}
	if ent.Description != nil {
		item.Description = ent.Description
	}
	return item
}

// toProtoFormScene entity → proto 转换
func toProtoFormScene(ent *entity.FormScene) *form.FormSceneItem {
	item := &form.FormSceneItem{
		Id:        ent.Id,
		Keyword:   ent.Keyword,
		Name:      ent.Name,
		CreatedAt: uint32(ent.CreatedAt),
		UpdatedAt: uint32(ent.UpdatedAt),
	}
	if ent.Description != nil {
		item.Description = ent.Description
	}
	return item
}
