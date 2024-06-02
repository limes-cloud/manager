package service

import (
	"context"

	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/limes-cloud/kratosx"
	"github.com/limes-cloud/kratosx/pkg/valx"
	pb "github.com/limes-cloud/manager/api/manager/department/v1"
	"github.com/limes-cloud/manager/api/manager/errors"
	"github.com/limes-cloud/manager/internal/biz/department"
	"github.com/limes-cloud/manager/internal/conf"
	"github.com/limes-cloud/manager/internal/data"
)

type DepartmentService struct {
	pb.UnimplementedDepartmentServer
	uc *department.UseCase
}

func NewDepartmentService(conf *conf.Config) *DepartmentService {
	return &DepartmentService{
		uc: department.NewUseCase(conf, data.NewDepartmentRepo()),
	}
}

func init() {
	register(func(c *conf.Config, hs *http.Server, gs *grpc.Server) {
		srv := NewDepartmentService(c)
		pb.RegisterDepartmentHTTPServer(hs, srv)
		pb.RegisterDepartmentServer(gs, srv)
	})
}

// GetDepartment 获取指定的部门信息
func (s *DepartmentService) GetDepartment(c context.Context, req *pb.GetDepartmentRequest) (*pb.GetDepartmentReply, error) {
	var (
		in  = department.GetDepartmentRequest{}
		ctx = kratosx.MustContext(c)
	)

	if err := valx.Transform(req, &in); err != nil {
		ctx.Logger().Warnw("msg", "request transform err", "err", err.Error())
		return nil, errors.TransformError()
	}

	result, err := s.uc.GetDepartment(ctx, &in)
	if err != nil {
		return nil, err
	}

	reply := pb.GetDepartmentReply{}
	if err := valx.Transform(result, &reply); err != nil {
		ctx.Logger().Warnw("msg", "reply transform err", "err", err.Error())
		return nil, errors.TransformError()
	}
	return &reply, nil
}

// ListDepartment 获取部门信息列表
func (s *DepartmentService) ListDepartment(c context.Context, req *pb.ListDepartmentRequest) (*pb.ListDepartmentReply, error) {
	var (
		in  = department.ListDepartmentRequest{}
		ctx = kratosx.MustContext(c)
	)

	if err := valx.Transform(req, &in); err != nil {
		ctx.Logger().Warnw("msg", "req transform err", "err", err.Error())
		return nil, errors.TransformError()
	}

	result, total, err := s.uc.ListDepartment(ctx, &in)
	if err != nil {
		return nil, err
	}

	reply := pb.ListDepartmentReply{Total: total}
	if err := valx.Transform(result, &reply.List); err != nil {
		ctx.Logger().Warnw("msg", "reply transform err", "err", err.Error())
		return nil, errors.TransformError()
	}

	return &reply, nil
}

// CreateDepartment 创建部门信息
func (s *DepartmentService) CreateDepartment(c context.Context, req *pb.CreateDepartmentRequest) (*pb.CreateDepartmentReply, error) {
	var (
		in  = department.Department{}
		ctx = kratosx.MustContext(c)
	)

	if err := valx.Transform(req, &in); err != nil {
		ctx.Logger().Warnw("msg", "req transform err", "err", err.Error())
		return nil, errors.TransformError()
	}

	id, err := s.uc.CreateDepartment(ctx, &in)
	if err != nil {
		return nil, err
	}

	return &pb.CreateDepartmentReply{Id: id}, nil
}

// UpdateDepartment 更新部门信息
func (s *DepartmentService) UpdateDepartment(c context.Context, req *pb.UpdateDepartmentRequest) (*pb.UpdateDepartmentReply, error) {
	var (
		in  = department.Department{}
		ctx = kratosx.MustContext(c)
	)

	if err := valx.Transform(req, &in); err != nil {
		ctx.Logger().Warnw("msg", "req transform err", "err", err.Error())
		return nil, errors.TransformError()
	}

	if err := s.uc.UpdateDepartment(ctx, &in); err != nil {
		return nil, err
	}

	return &pb.UpdateDepartmentReply{}, nil
}

// DeleteDepartment 删除部门信息
func (s *DepartmentService) DeleteDepartment(c context.Context, req *pb.DeleteDepartmentRequest) (*pb.DeleteDepartmentReply, error) {
	total, err := s.uc.DeleteDepartment(kratosx.MustContext(c), req.Ids)
	if err != nil {
		return nil, err
	}
	return &pb.DeleteDepartmentReply{Total: total}, nil
}
