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

// ListDepartmentClassify 获取任务分组列表
func (s *Department) ListDepartmentClassify(c context.Context, req *pb.ListDepartmentClassifyRequest) (*pb.ListDepartmentClassifyReply, error) {
	list, total, err := s.srv.ListDepartmentClassify(kratosx.MustContext(c), &types.ListDepartmentClassifyRequest{
		Page:     req.Page,
		PageSize: req.PageSize,
		Order:    req.Order,
		OrderBy:  req.OrderBy,
		Name:     req.Name,
	})
	if err != nil {
		return nil, err
	}
	reply := pb.ListDepartmentClassifyReply{Total: total}
	for _, item := range list {
		reply.List = append(reply.List, &pb.ListDepartmentClassifyReply_DepartmentClassify{
			Id:          item.Id,
			Name:        item.Name,
			Description: item.Description,
			CreatedAt:   uint32(item.CreatedAt),
			UpdatedAt:   uint32(item.UpdatedAt),
		})
	}
	return &reply, nil
}

// CreateDepartmentClassify 创建任务分组
func (s *Department) CreateDepartmentClassify(c context.Context, req *pb.CreateDepartmentClassifyRequest) (*pb.CreateDepartmentClassifyReply, error) {
	id, err := s.srv.CreateDepartmentClassify(kratosx.MustContext(c), &entity.DepartmentClassify{
		Name:        req.Name,
		Description: req.Description,
	})
	if err != nil {
		return nil, err
	}
	return &pb.CreateDepartmentClassifyReply{Id: id}, nil
}

// UpdateDepartmentClassify 更新任务分组
func (s *Department) UpdateDepartmentClassify(c context.Context, req *pb.UpdateDepartmentClassifyRequest) (*pb.UpdateDepartmentClassifyReply, error) {
	if err := s.srv.UpdateDepartmentClassify(kratosx.MustContext(c), &entity.DepartmentClassify{
		BaseModel:   ktypes.BaseModel{Id: req.Id},
		Name:        req.Name,
		Description: req.Description,
	}); err != nil {
		return nil, err
	}
	return &pb.UpdateDepartmentClassifyReply{}, nil
}

// DeleteDepartmentClassify 删除任务分组
func (s *Department) DeleteDepartmentClassify(c context.Context, req *pb.DeleteDepartmentClassifyRequest) (*pb.DeleteDepartmentClassifyReply, error) {
	err := s.srv.DeleteDepartmentClassify(kratosx.MustContext(c), req.Id)
	if err != nil {
		return nil, err
	}
	return &pb.DeleteDepartmentClassifyReply{}, nil
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
		ClassifyId:  ent.ClassifyId,
		Name:        ent.Name,
		Keyword:     ent.Keyword,
		Description: ent.Description,
		Classify: &pb.DepartmentClassify{
			Id:   ent.Classify.Id,
			Name: ent.Classify.Name,
		},
		CreatedAt: uint32(ent.CreatedAt),
		UpdatedAt: uint32(ent.UpdatedAt),
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
		ClassifyId:  req.ClassifyId,
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
		ClassifyId:  req.ClassifyId,
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
