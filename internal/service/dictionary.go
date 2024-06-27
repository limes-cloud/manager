package service

import (
	"context"

	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/limes-cloud/kratosx"
	"github.com/limes-cloud/kratosx/pkg/valx"

	pb "github.com/limes-cloud/manager/api/manager/dictionary/v1"
	"github.com/limes-cloud/manager/api/manager/errors"
	"github.com/limes-cloud/manager/internal/biz/dictionary"
	"github.com/limes-cloud/manager/internal/conf"
	"github.com/limes-cloud/manager/internal/data"
)

type DictionaryService struct {
	pb.UnimplementedDictionaryServer
	uc *dictionary.UseCase
}

func NewDictionaryService(conf *conf.Config) *DictionaryService {
	return &DictionaryService{
		uc: dictionary.NewUseCase(conf, data.NewDictionaryRepo()),
	}
}

func init() {
	register(func(c *conf.Config, hs *http.Server, gs *grpc.Server) {
		srv := NewDictionaryService(c)
		pb.RegisterDictionaryHTTPServer(hs, srv)
		pb.RegisterDictionaryServer(gs, srv)
	})
}

// GetDictionary 获取指定的字典目录
func (s *DictionaryService) GetDictionary(c context.Context, req *pb.GetDictionaryRequest) (*pb.GetDictionaryReply, error) {
	var (
		in  = dictionary.GetDictionaryRequest{}
		ctx = kratosx.MustContext(c)
	)

	if err := valx.Transform(req, &in); err != nil {
		ctx.Logger().Warnw("msg", "req transform err", "err", err.Error())
		return nil, errors.TransformError()
	}

	result, err := s.uc.GetDictionary(ctx, &in)
	if err != nil {
		return nil, err
	}

	reply := pb.GetDictionaryReply{}
	if err := valx.Transform(result, &reply); err != nil {
		ctx.Logger().Warnw("msg", "reply transform err", "err", err.Error())
		return nil, errors.TransformError()
	}
	return &reply, nil
}

// ListDictionary 获取字典目录列表
func (s *DictionaryService) ListDictionary(c context.Context, req *pb.ListDictionaryRequest) (*pb.ListDictionaryReply, error) {
	var (
		in  = dictionary.ListDictionaryRequest{}
		ctx = kratosx.MustContext(c)
	)

	if err := valx.Transform(req, &in); err != nil {
		ctx.Logger().Warnw("msg", "req transform err", "err", err.Error())
		return nil, errors.TransformError()
	}

	result, total, err := s.uc.ListDictionary(ctx, &in)
	if err != nil {
		return nil, err
	}

	reply := pb.ListDictionaryReply{Total: total}
	if err := valx.Transform(result, &reply.List); err != nil {
		ctx.Logger().Warnw("msg", "reply transform err", "err", err.Error())
		return nil, errors.TransformError()
	}

	return &reply, nil
}

// CreateDictionary 创建字典目录
func (s *DictionaryService) CreateDictionary(c context.Context, req *pb.CreateDictionaryRequest) (*pb.CreateDictionaryReply, error) {
	var (
		in  = dictionary.Dictionary{}
		ctx = kratosx.MustContext(c)
	)

	if err := valx.Transform(req, &in); err != nil {
		ctx.Logger().Warnw("msg", "req transform err", "err", err.Error())
		return nil, errors.TransformError()
	}

	id, err := s.uc.CreateDictionary(ctx, &in)
	if err != nil {
		return nil, err
	}

	return &pb.CreateDictionaryReply{Id: id}, nil
}

// UpdateDictionary 更新字典目录
func (s *DictionaryService) UpdateDictionary(c context.Context, req *pb.UpdateDictionaryRequest) (*pb.UpdateDictionaryReply, error) {
	var (
		in  = dictionary.Dictionary{}
		ctx = kratosx.MustContext(c)
	)

	if err := valx.Transform(req, &in); err != nil {
		ctx.Logger().Warnw("msg", "req transform err", "err", err.Error())
		return nil, errors.TransformError()
	}

	if err := s.uc.UpdateDictionary(ctx, &in); err != nil {
		return nil, err
	}

	return &pb.UpdateDictionaryReply{}, nil
}

// DeleteDictionary 删除字典目录
func (s *DictionaryService) DeleteDictionary(c context.Context, req *pb.DeleteDictionaryRequest) (*pb.DeleteDictionaryReply, error) {
	total, err := s.uc.DeleteDictionary(kratosx.MustContext(c), req.Ids)
	if err != nil {
		return nil, err
	}
	return &pb.DeleteDictionaryReply{Total: total}, nil
}

// ListDictionaryValue 获取字典值目录列表
func (s *DictionaryService) ListDictionaryValue(c context.Context, req *pb.ListDictionaryValueRequest) (*pb.ListDictionaryValueReply, error) {
	var (
		in  = dictionary.ListDictionaryValueRequest{}
		ctx = kratosx.MustContext(c)
	)

	if err := valx.Transform(req, &in); err != nil {
		ctx.Logger().Warnw("msg", "req transform err", "err", err.Error())
		return nil, errors.TransformError()
	}

	result, total, err := s.uc.ListDictionaryValue(ctx, &in)
	if err != nil {
		return nil, err
	}

	reply := pb.ListDictionaryValueReply{Total: total}
	if err := valx.Transform(result, &reply.List); err != nil {
		ctx.Logger().Warnw("msg", "reply transform err", "err", err.Error())
		return nil, errors.TransformError()
	}

	return &reply, nil
}

// CreateDictionaryValue 创建字典值目录
func (s *DictionaryService) CreateDictionaryValue(c context.Context, req *pb.CreateDictionaryValueRequest) (*pb.CreateDictionaryValueReply, error) {
	var (
		in  = dictionary.DictionaryValue{}
		ctx = kratosx.MustContext(c)
	)

	if err := valx.Transform(req, &in); err != nil {
		ctx.Logger().Warnw("msg", "req transform err", "err", err.Error())
		return nil, errors.TransformError()
	}

	id, err := s.uc.CreateDictionaryValue(ctx, &in)
	if err != nil {
		return nil, err
	}

	return &pb.CreateDictionaryValueReply{Id: id}, nil
}

// UpdateDictionaryValue 更新字典值目录
func (s *DictionaryService) UpdateDictionaryValue(c context.Context, req *pb.UpdateDictionaryValueRequest) (*pb.UpdateDictionaryValueReply, error) {
	var (
		in  = dictionary.DictionaryValue{}
		ctx = kratosx.MustContext(c)
	)

	if err := valx.Transform(req, &in); err != nil {
		ctx.Logger().Warnw("msg", "req transform err", "err", err.Error())
		return nil, errors.TransformError()
	}

	if err := s.uc.UpdateDictionaryValue(ctx, &in); err != nil {
		return nil, err
	}

	return &pb.UpdateDictionaryValueReply{}, nil
}

// UpdateDictionaryValueStatus 更新字典值目录状态
func (s *DictionaryService) UpdateDictionaryValueStatus(c context.Context, req *pb.UpdateDictionaryValueStatusRequest) (*pb.UpdateDictionaryValueStatusReply, error) {
	return &pb.UpdateDictionaryValueStatusReply{}, s.uc.UpdateDictionaryValueStatus(kratosx.MustContext(c), req.Id, req.Status)
}

// DeleteDictionaryValue 删除字典值目录
func (s *DictionaryService) DeleteDictionaryValue(c context.Context, req *pb.DeleteDictionaryValueRequest) (*pb.DeleteDictionaryValueReply, error) {
	total, err := s.uc.DeleteDictionaryValue(kratosx.MustContext(c), req.Ids)
	if err != nil {
		return nil, err
	}
	return &pb.DeleteDictionaryValueReply{Total: total}, nil
}

func (s *DictionaryService) GetDictionaryValues(c context.Context, req *pb.GetDictionaryValuesRequest) (*pb.GetDictionaryValuesReply, error) {
	var (
		ctx = kratosx.MustContext(c)
	)
	res, err := s.uc.GetDictionaryValues(ctx, req.Keywords)
	if err != nil {
		return nil, err
	}

	reply := pb.GetDictionaryValuesReply{Dict: make(map[string]*pb.GetDictionaryValuesReply_Value)}
	for key, values := range res {
		for _, val := range values {
			reply.Dict[key].List = append(reply.Dict[key].List, &pb.GetDictionaryValuesReply_Value_Item{
				Label:       val.Label,
				Value:       val.Value,
				Type:        val.Type,
				Extra:       val.Extra,
				Description: val.Description,
			})
		}
	}
	return &reply, nil
}
