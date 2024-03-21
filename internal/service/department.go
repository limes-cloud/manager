package service

import (
	"context"

	"github.com/limes-cloud/kratosx"
	"github.com/limes-cloud/kratosx/pkg/util"
	"google.golang.org/protobuf/types/known/emptypb"

	v1 "github.com/limes-cloud/manager/api/department/v1"
	"github.com/limes-cloud/manager/api/errors"
	biz "github.com/limes-cloud/manager/internal/biz/department"
	"github.com/limes-cloud/manager/internal/config"
	data "github.com/limes-cloud/manager/internal/data/department"
	objectData "github.com/limes-cloud/manager/internal/data/object"
)

type DepartmentService struct {
	v1.UnimplementedServiceServer
	uc *biz.UseCase

	conf *config.Config
}

func NewDepartmentService(conf *config.Config) *DepartmentService {
	return &DepartmentService{
		conf: conf,
		uc:   biz.NewUseCase(conf, data.NewRepo(), objectData.NewRepo()),
	}
}

func (d DepartmentService) GetDepartmentTree(ctx context.Context, empty *emptypb.Empty) (*v1.GetDepartmentTreeReply, error) {
	tree, err := d.uc.DepartmentTree(kratosx.MustContext(ctx))
	if err != nil {
		return nil, err
	}

	reply := v1.GetDepartmentTreeReply{}
	if err := util.Transform(tree, &reply.List); err != nil {
		return nil, errors.Transform()
	}

	return &reply, nil
}

func (d DepartmentService) AddDepartment(ctx context.Context, request *v1.AddDepartmentRequest) (*v1.AddDepartmentReply, error) {
	var dep biz.Department
	if err := util.Transform(request, &dep); err != nil {
		return nil, errors.Transform()
	}

	id, err := d.uc.AddDepartment(kratosx.MustContext(ctx), &dep)
	if err != nil {
		return nil, err
	}

	return &v1.AddDepartmentReply{Id: id}, nil
}

func (d DepartmentService) UpdateDepartment(ctx context.Context, request *v1.UpdateDepartmentRequest) (*emptypb.Empty, error) {
	var dep biz.Department
	if err := util.Transform(request, &dep); err != nil {
		return nil, errors.Transform()
	}

	return nil, d.uc.UpdateDepartment(kratosx.MustContext(ctx), &dep)
}

func (d DepartmentService) DeleteDepartment(ctx context.Context, request *v1.DeleteDepartmentRequest) (*emptypb.Empty, error) {
	return nil, d.uc.DeleteDepartment(kratosx.MustContext(ctx), request.Id)
}

func (d DepartmentService) AllDepartmentObjectValue(ctx context.Context, request *v1.AllDepartmentObjectValueRequest) (*v1.AllDepartmentObjectValueReply, error) {
	var dep biz.AllDepartmentObjectValueRequest
	if err := util.Transform(request, &dep); err != nil {
		return nil, errors.Transform()
	}
	list, err := d.uc.AllDepartmentObjectValue(kratosx.MustContext(ctx), &dep)
	return &v1.AllDepartmentObjectValueReply{
		Values: list,
	}, err
}

func (d DepartmentService) ImportDepartmentObject(ctx context.Context, request *v1.ImportDepartmentObjectRequest) (*emptypb.Empty, error) {
	var bucket = make(map[string]bool)
	var list []*biz.DepartmentObject
	for _, value := range request.Values {
		if bucket[value] {
			continue
		}
		bucket[value] = true
		list = append(list, &biz.DepartmentObject{
			DepartmentId: request.DepartmentId,
			ObjectId:     request.ObjectId,
			Value:        value,
		})
	}
	return nil, d.uc.ImportDepartmentObject(kratosx.MustContext(ctx), list)
}

func (d DepartmentService) AddDepartmentObject(ctx context.Context, request *v1.AddDepartmentObjectRequest) (*v1.AddDepartmentObjectReply, error) {
	id, err := d.uc.AddDepartmentObject(kratosx.MustContext(ctx), request.ObjectKeyword, request.Value)
	return &v1.AddDepartmentObjectReply{Id: id}, err
}

func (d DepartmentService) DeleteDepartmentObject(ctx context.Context, request *v1.DeleteDepartmentObjectRequest) (*emptypb.Empty, error) {
	return nil, d.uc.DeleteDepartmentObject(kratosx.MustContext(ctx), request.ObjectKeyword, request.Value)
}
