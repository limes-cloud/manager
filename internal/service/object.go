package service

import (
	"context"

	"github.com/limes-cloud/kratosx"
	"github.com/limes-cloud/kratosx/pkg/util"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/limes-cloud/manager/api/errors"
	v1 "github.com/limes-cloud/manager/api/object/v1"
	biz "github.com/limes-cloud/manager/internal/biz/object"
	"github.com/limes-cloud/manager/internal/config"
	data "github.com/limes-cloud/manager/internal/data/object"
)

type ObjectService struct {
	v1.UnimplementedServiceServer
	uc *biz.UseCase

	conf *config.Config
}

func NewObjectService(conf *config.Config) *ObjectService {
	return &ObjectService{
		conf: conf,
		uc:   biz.NewUseCase(conf, data.NewRepo()),
	}
}

func (o ObjectService) AddObject(ctx context.Context, request *v1.AddObjectRequest) (*v1.AddObjectReply, error) {
	var object biz.Object
	if err := util.Transform(request, &object); err != nil {
		return nil, errors.Transform()
	}

	id, err := o.uc.AddObject(kratosx.MustContext(ctx), &object)
	if err != nil {
		return nil, err
	}

	return &v1.AddObjectReply{Id: id}, nil
}

func (o ObjectService) DeleteObject(ctx context.Context, request *v1.DeleteObjectRequest) (*emptypb.Empty, error) {
	return nil, o.uc.DeleteObject(kratosx.MustContext(ctx), request.Id)
}

func (o ObjectService) GetObject(ctx context.Context, request *v1.GetObjectRequest) (*v1.Object, error) {
	obj, err := o.uc.GetObjectById(kratosx.MustContext(ctx), request.Id)
	if err != nil {
		return nil, err
	}

	reply := v1.Object{}
	if err := util.Transform(obj, &reply); err != nil {
		return nil, errors.Transform()
	}

	return &reply, nil
}

func (o ObjectService) PageObject(ctx context.Context, request *v1.PageObjectRequest) (*v1.PageObjectReply, error) {
	var in biz.PageObjectRequest
	if err := util.Transform(request, &in); err != nil {
		return nil, errors.Transform()
	}

	list, total, err := o.uc.PageObject(kratosx.MustContext(ctx), &in)
	if err != nil {
		return nil, err
	}

	reply := v1.PageObjectReply{Total: total}
	if err := util.Transform(list, &reply.List); err != nil {
		return nil, errors.Transform()
	}

	return &reply, nil
}

func (o ObjectService) UpdateObject(ctx context.Context, request *v1.UpdateObjectRequest) (*emptypb.Empty, error) {
	var object biz.Object
	if err := util.Transform(request, &object); err != nil {
		return nil, errors.Transform()
	}

	return nil, o.uc.UpdateObject(kratosx.MustContext(ctx), &object)
}
