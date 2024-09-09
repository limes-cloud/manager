package app

import (
	"context"

	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/limes-cloud/kratosx"
	"github.com/limes-cloud/kratosx/pkg/valx"
	ktypes "github.com/limes-cloud/kratosx/types"

	pb "github.com/limes-cloud/manager/api/manager/department/v1"
	"github.com/limes-cloud/manager/api/manager/errors"
	"github.com/limes-cloud/manager/internal/conf"
	"github.com/limes-cloud/manager/internal/domain/entity"
	"github.com/limes-cloud/manager/internal/domain/service"
	"github.com/limes-cloud/manager/internal/infra/dbs"
	"github.com/limes-cloud/manager/internal/types"
)

type Department struct {
	pb.UnimplementedDepartmentServer
	srv *service.Department
}

func NewDepartment(conf *conf.Config) *Department {
	return &Department{
		srv: service.NewDepartment(conf, dbs.NewDepartment()),
	}
}

func init() {
	register(func(c *conf.Config, hs *http.Server, gs *grpc.Server) {
		srv := NewDepartment(c)
		pb.RegisterDepartmentHTTPServer(hs, srv)
		pb.RegisterDepartmentServer(gs, srv)
	})
}

// GetDepartment 获取指定的部门信息
func (s *Department) GetDepartment(c context.Context, req *pb.GetDepartmentRequest) (*pb.GetDepartmentReply, error) {
	var (
		ent *entity.Department
		err error
	)
	switch req.Params.(type) {
	case *pb.GetDepartmentRequest_Id:
		ent, err = s.srv.GetDepartment(kratosx.MustContext(c), req.GetId())
	case *pb.GetDepartmentRequest_Keyword:
		ent, err = s.srv.GetDepartmentByKeyword(kratosx.MustContext(c), req.GetKeyword())
	default:
		return nil, errors.ParamsError()
	}
	if err != nil {
		return nil, err
	}

	return &pb.GetDepartmentReply{
		Id:          ent.Id,
		ParentId:    ent.ParentId,
		Name:        ent.Name,
		Keyword:     ent.Keyword,
		Description: ent.Description,
		CreatedAt:   uint32(ent.CreatedAt),
		UpdatedAt:   uint32(ent.UpdatedAt),
	}, nil
}

// ListDepartment 获取部门信息列表
func (s *Department) ListDepartment(c context.Context, req *pb.ListDepartmentRequest) (*pb.ListDepartmentReply, error) {
	ctx := kratosx.MustContext(c)
	result, err := s.srv.ListCurrentDepartment(ctx, &types.ListDepartmentRequest{
		Name:    req.Name,
		Keyword: req.Keyword,
	})
	if err != nil {
		return nil, err
	}

	reply := pb.ListDepartmentReply{}
	if err := valx.Transform(result, &reply.List); err != nil {
		ctx.Logger().Warnw("msg", "reply transform err", "err", err.Error())
		return nil, errors.TransformError()
	}

	return &reply, nil
}

// CreateDepartment 创建部门信息
func (s *Department) CreateDepartment(c context.Context, req *pb.CreateDepartmentRequest) (*pb.CreateDepartmentReply, error) {
	id, err := s.srv.CreateDepartment(kratosx.MustContext(c), &entity.Department{
		ParentId:    req.ParentId,
		Name:        req.Name,
		Keyword:     req.Keyword,
		Description: req.Description,
	})
	if err != nil {
		return nil, err
	}

	return &pb.CreateDepartmentReply{Id: id}, nil
}

// UpdateDepartment 更新部门信息
func (s *Department) UpdateDepartment(c context.Context, req *pb.UpdateDepartmentRequest) (*pb.UpdateDepartmentReply, error) {
	if err := s.srv.UpdateDepartment(kratosx.MustContext(c), &entity.Department{
		BaseModel:   ktypes.BaseModel{Id: req.Id},
		ParentId:    req.ParentId,
		Name:        req.Name,
		Description: req.Description,
	}); err != nil {
		return nil, err
	}

	return &pb.UpdateDepartmentReply{}, nil
}

// DeleteDepartment 删除部门信息
func (s *Department) DeleteDepartment(c context.Context, req *pb.DeleteDepartmentRequest) (*pb.DeleteDepartmentReply, error) {
	if err := s.srv.DeleteDepartment(kratosx.MustContext(c), req.Id); err != nil {
		return nil, err
	}
	return &pb.DeleteDepartmentReply{}, nil
}
