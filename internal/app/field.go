package app

import (
	"context"

	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/limes-cloud/kratosx"
	ktypes "github.com/limes-cloud/kratosx/types"

	pb "github.com/limes-cloud/manager/api/manager/field/v1"
	"github.com/limes-cloud/manager/internal/conf"
	"github.com/limes-cloud/manager/internal/domain/entity"
	"github.com/limes-cloud/manager/internal/domain/service"
	"github.com/limes-cloud/manager/internal/infra/dbs"
	"github.com/limes-cloud/manager/internal/types"
)

type Field struct {
	pb.UnimplementedFieldServer
	srv *service.Field
}

func NewField(conf *conf.Config) *Field {
	return &Field{
		srv: service.NewField(
			conf,
			dbs.NewField(),
		),
	}
}

func init() {
	register(func(c *conf.Config, hs *http.Server, gs *grpc.Server) {
		srv := NewField(c)
		pb.RegisterFieldHTTPServer(hs, srv)
		pb.RegisterFieldServer(gs, srv)
	})
}

// ListFieldType 获取字段类型列表
func (fd *Field) ListFieldType(c context.Context, req *pb.ListFieldTypeRequest) (*pb.ListFieldTypeReply, error) {
	list := fd.srv.ListFieldType()
	var reply = pb.ListFieldTypeReply{}
	for _, item := range list {
		reply.List = append(reply.List, &pb.ListFieldTypeReply_Type{
			Name: item.Name,
			Type: item.Type,
		})
	}

	return &reply, nil
}

// ListField 获取用户字段列表
func (fd *Field) ListField(c context.Context, req *pb.ListFieldRequest) (*pb.ListFieldReply, error) {
	list, total, err := fd.srv.ListField(kratosx.MustContext(c), &types.ListFieldRequest{
		Page:     req.Page,
		PageSize: req.PageSize,
		Order:    req.Order,
		OrderBy:  req.OrderBy,
		Keyword:  req.Keyword,
		Name:     req.Name,
		Status:   req.Status,
	})
	if err != nil {
		return nil, err
	}
	reply := pb.ListFieldReply{Total: total}
	for _, item := range list {
		reply.List = append(reply.List, &pb.ListFieldReply_Field{
			Id:          item.Id,
			Keyword:     item.Keyword,
			Type:        item.Type,
			Name:        item.Name,
			Status:      item.Status,
			Description: item.Description,
			CreatedAt:   uint32(item.CreatedAt),
			UpdatedAt:   uint32(item.UpdatedAt),
		})
	}
	return &reply, nil
}

// CreateField 创建用户字段
func (fd *Field) CreateField(c context.Context, req *pb.CreateFieldRequest) (*pb.CreateFieldReply, error) {
	id, err := fd.srv.CreateField(kratosx.MustContext(c), &entity.Field{
		Keyword:     req.Keyword,
		Type:        req.Type,
		Name:        req.Name,
		Description: req.Description,
	})
	if err != nil {
		return nil, err
	}
	return &pb.CreateFieldReply{Id: id}, nil
}

// UpdateField 更新用户字段
func (fd *Field) UpdateField(c context.Context, req *pb.UpdateFieldRequest) (*pb.UpdateFieldReply, error) {
	if err := fd.srv.UpdateField(kratosx.MustContext(c), &entity.Field{
		BaseModel:   ktypes.BaseModel{Id: req.Id},
		Keyword:     req.Keyword,
		Type:        req.Type,
		Name:        req.Name,
		Status:      req.Status,
		Description: req.Description,
	}); err != nil {
		return nil, err
	}
	return &pb.UpdateFieldReply{}, nil
}

// DeleteField 删除用户字段
func (fd *Field) DeleteField(c context.Context, req *pb.DeleteFieldRequest) (*pb.DeleteFieldReply, error) {
	if err := fd.srv.DeleteField(kratosx.MustContext(c), req.Id); err != nil {
		return nil, err
	}
	return &pb.DeleteFieldReply{}, nil
}
