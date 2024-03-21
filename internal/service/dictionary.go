package service

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/limes-cloud/kratosx"
	"github.com/limes-cloud/kratosx/pkg/util"

	v1 "github.com/limes-cloud/manager/api/dictionary/v1"
	"github.com/limes-cloud/manager/api/errors"
	biz "github.com/limes-cloud/manager/internal/biz/dictionary"
	"github.com/limes-cloud/manager/internal/config"
	data "github.com/limes-cloud/manager/internal/data/dictionary"
)

type DictionaryService struct {
	v1.UnimplementedServiceServer
	uc   *biz.UseCase
	conf *config.Config
}

func NewDictionaryService(conf *config.Config) *DictionaryService {
	return &DictionaryService{
		conf: conf,
		uc:   biz.NewUseCase(conf, data.NewRepo()),
	}
}

func (s *DictionaryService) PageDictionary(ctx context.Context, request *v1.PageDictionaryRequest) (*v1.PageDictionaryReply, error) {
	var req biz.PageDictionaryRequest
	if err := util.Transform(request, &req); err != nil {
		return nil, errors.Transform()
	}

	list, total, err := s.uc.PageDictionary(kratosx.MustContext(ctx), &req)
	if err != nil {
		return nil, err
	}

	reply := v1.PageDictionaryReply{Total: total}
	if err := util.Transform(list, &reply.List); err != nil {
		return nil, errors.Transform()
	}

	return &reply, nil
}

func (s *DictionaryService) AddDictionary(ctx context.Context, request *v1.AddDictionaryRequest) (*v1.AddDictionaryReply, error) {
	var req biz.Dictionary
	if err := util.Transform(request, &req); err != nil {
		return nil, errors.Transform()
	}

	id, err := s.uc.AddDictionary(kratosx.MustContext(ctx), &req)
	if err != nil {
		return nil, err
	}
	return &v1.AddDictionaryReply{Id: id}, nil
}

func (s *DictionaryService) UpdateDictionary(ctx context.Context, request *v1.UpdateDictionaryRequest) (*empty.Empty, error) {
	var req biz.Dictionary
	if err := util.Transform(request, &req); err != nil {
		return nil, errors.Transform()
	}

	return nil, s.uc.UpdateDictionary(kratosx.MustContext(ctx), &req)
}

func (s *DictionaryService) DeleteDictionary(ctx context.Context, request *v1.DeleteDictionaryRequest) (*empty.Empty, error) {
	return nil, s.uc.DeleteDictionary(kratosx.MustContext(ctx), request.Id)
}

func (s *DictionaryService) GetDictionaryValue(ctx context.Context, request *v1.GetDictionaryValueRequest) (*v1.GetDictionaryValueReply, error) {
	list, err := s.uc.GetDictionaryValue(kratosx.MustContext(ctx), request.Keyword)
	if err != nil {
		return nil, err
	}

	reply := &v1.GetDictionaryValueReply{Dict: make(map[string]string)}
	for _, value := range list {
		reply.Dict[value.Value] = value.Label
	}

	if err = util.Transform(list, &reply.List); err != nil {
		return nil, errors.Transform()
	}
	return reply, nil
}

func (s *DictionaryService) PageDictionaryValue(ctx context.Context, request *v1.PageDictionaryValueRequest) (*v1.PageDictionaryValueReply, error) {
	var req biz.PageDictionaryValueRequest
	if err := util.Transform(request, &req); err != nil {
		return nil, errors.Transform()
	}

	list, total, err := s.uc.PageDictionaryValue(kratosx.MustContext(ctx), &req)
	if err != nil {
		return nil, err
	}

	reply := v1.PageDictionaryValueReply{Total: total}
	if err := util.Transform(list, &reply.List); err != nil {
		return nil, errors.Transform()
	}

	return &reply, nil
}

func (s *DictionaryService) AddDictionaryValue(ctx context.Context, request *v1.AddDictionaryValueRequest) (*v1.AddDictionaryValueReply, error) {
	var req biz.DictionaryValue
	if err := util.Transform(request, &req); err != nil {
		return nil, errors.Transform()
	}

	id, err := s.uc.AddDictionaryValue(kratosx.MustContext(ctx), &req)
	if err != nil {
		return nil, err
	}
	return &v1.AddDictionaryValueReply{Id: id}, nil
}

func (s *DictionaryService) UpdateDictionaryValue(ctx context.Context, request *v1.UpdateDictionaryValueRequest) (*empty.Empty, error) {
	var req biz.DictionaryValue
	if err := util.Transform(request, &req); err != nil {
		return nil, errors.Transform()
	}

	return nil, s.uc.UpdateDictionaryValue(kratosx.MustContext(ctx), &req)
}

func (s *DictionaryService) DeleteDictionaryValue(ctx context.Context, request *v1.DeleteDictionaryValueRequest) (*empty.Empty, error) {
	return nil, s.uc.DeleteDictionaryValue(kratosx.MustContext(ctx), request.Id)
}
