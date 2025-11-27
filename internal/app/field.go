package app

import (
	"context"

	"github.com/limes-cloud/kratosx/pkg/value"

	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"

	"github.com/limes-cloud/manager/internal/core"

	"github.com/limes-cloud/kratosx/model"

	"github.com/limes-cloud/kratosx/model/page"

	"github.com/limes-cloud/manager/api/field"

	"github.com/limes-cloud/manager/internal/domain/entity"
	"github.com/limes-cloud/manager/internal/domain/service"
	"github.com/limes-cloud/manager/internal/infra/dbs"
	"github.com/limes-cloud/manager/internal/types"
)

type Field struct {
	field.UnimplementedFieldServer
	srv *service.Field
}

func NewField() *Field {
	return &Field{
		srv: service.NewField(
			dbs.NewField(),
		),
	}
}

func init() {
	register(func(hs *http.Server, gs *grpc.Server) {
		srv := NewField()
		field.RegisterFieldHTTPServer(hs, srv)
		field.RegisterFieldServer(gs, srv)
	})
}

// ListFieldType 获取字段类型列表
func (fd *Field) ListFieldType(c context.Context, req *field.ListFieldTypeRequest) (*field.ListFieldTypeReply, error) {
	list := fd.srv.ListFieldType()
	reply := field.ListFieldTypeReply{}
	for _, item := range list {
		reply.List = append(reply.List, &field.ListFieldTypeReply_Type{
			Name: item.Name,
			Type: item.Type,
		})
	}

	return &reply, nil
}

// ListRequiredField 获取必填用户字段列表
func (fd *Field) ListRequiredField(c context.Context, _ *field.ListRequiredFieldRequest) (*field.ListRequiredFieldReply, error) {
	list, _, err := fd.srv.ListField(core.MustContext(c), &types.ListFieldRequest{
		Status:   value.Pointer(true),
		Required: value.Pointer(true),
	})
	if err != nil {
		return nil, err
	}
	reply := field.ListRequiredFieldReply{}
	for _, item := range list {
		reply.List = append(reply.List, &field.ListRequiredFieldReply_Field{
			Keyword: item.Keyword,
			Type:    item.Type,
			Name:    item.Name,
		})
	}
	return &reply, nil
}

// ListField 获取用户字段列表
func (fd *Field) ListField(c context.Context, req *field.ListFieldRequest) (*field.ListFieldReply, error) {
	list, total, err := fd.srv.ListField(core.MustContext(c), &types.ListFieldRequest{
		Search: &page.Search{
			Page:     req.Page,
			PageSize: req.PageSize,
			Order:    req.Order,
			OrderBy:  req.OrderBy,
		},
		Keyword: req.Keyword,
		Name:    req.Name,
		Status:  req.Status,
	})
	if err != nil {
		return nil, err
	}
	reply := field.ListFieldReply{Total: total}
	for _, item := range list {
		reply.List = append(reply.List, &field.ListFieldReply_Field{
			Id:          item.Id,
			Keyword:     item.Keyword,
			Type:        item.Type,
			Name:        item.Name,
			Status:      item.Status,
			Required:    item.Required,
			Unique:      item.Unique,
			Description: item.Description,
			CreatedAt:   uint32(item.CreatedAt),
			UpdatedAt:   uint32(item.UpdatedAt),
		})
	}
	return &reply, nil
}

// CreateField 创建用户字段
func (fd *Field) CreateField(c context.Context, req *field.CreateFieldRequest) (*field.CreateFieldReply, error) {
	id, err := fd.srv.CreateField(core.MustContext(c), &entity.Field{
		Keyword:     req.Keyword,
		Type:        req.Type,
		Name:        req.Name,
		Required:    &req.Required,
		Unique:      &req.Unique,
		Description: req.Description,
	})
	if err != nil {
		return nil, err
	}
	return &field.CreateFieldReply{Id: id}, nil
}

// UpdateField 更新用户字段
func (fd *Field) UpdateField(c context.Context, req *field.UpdateFieldRequest) (*field.UpdateFieldReply, error) {
	if err := fd.srv.UpdateField(core.MustContext(c), &entity.Field{
		BaseTenantModel: model.BaseTenantModel{Id: req.Id},
		Keyword:         req.Keyword,
		Type:            req.Type,
		Name:            req.Name,
		Status:          req.Status,
		Required:        req.Required,
		Unique:          req.Unique,
		Description:     req.Description,
	}); err != nil {
		return nil, err
	}
	return &field.UpdateFieldReply{}, nil
}

// DeleteField 删除用户字段
func (fd *Field) DeleteField(c context.Context, req *field.DeleteFieldRequest) (*field.DeleteFieldReply, error) {
	if err := fd.srv.DeleteField(core.MustContext(c), req.Id); err != nil {
		return nil, err
	}
	return &field.DeleteFieldReply{}, nil
}
