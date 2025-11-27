package app

import (
	"context"

	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"

	"github.com/limes-cloud/kratosx/pkg/value"

	"github.com/limes-cloud/kratosx/model"
	"github.com/limes-cloud/manager/api/dictionary"
	"github.com/limes-cloud/manager/api/errors"

	"github.com/limes-cloud/kratosx"

	"github.com/limes-cloud/manager/internal/domain/entity"
	"github.com/limes-cloud/manager/internal/domain/service"
	"github.com/limes-cloud/manager/internal/infra/dbs"
	"github.com/limes-cloud/manager/internal/types"
)

type Dictionary struct {
	dictionary.UnimplementedDictionaryServer
	srv *service.Dictionary
}

func NewDictionary() *Dictionary {
	return &Dictionary{
		srv: service.NewDictionary(dbs.NewDictionary()),
	}
}

func init() {
	register(func(hs *http.Server, gs *grpc.Server) {
		srv := NewDictionary()
		dictionary.RegisterDictionaryHTTPServer(hs, srv)
		dictionary.RegisterDictionaryServer(gs, srv)
	})
}

// GetDictionary 获取指定的字典目录
func (s *Dictionary) GetDictionary(c context.Context, req *dictionary.GetDictionaryRequest) (*dictionary.GetDictionaryReply, error) {
	result, err := s.srv.GetDictionary(kratosx.MustContext(c), &types.GetDictionaryRequest{
		Id:      req.Id,
		Keyword: req.Keyword,
	})
	if err != nil {
		return nil, err
	}
	return &dictionary.GetDictionaryReply{
		Id:          result.Id,
		Keyword:     result.Keyword,
		Name:        result.Name,
		Type:        result.Type,
		Description: result.Description,
		CreatedAt:   uint32(result.CreatedAt),
		UpdatedAt:   uint32(result.UpdatedAt),
	}, nil
}

// ListDictionary 获取字典目录列表
func (s *Dictionary) ListDictionary(c context.Context, req *dictionary.ListDictionaryRequest) (*dictionary.ListDictionaryReply, error) {
	ctx := kratosx.MustContext(c)
	list, total, err := s.srv.ListDictionary(ctx, &types.ListDictionaryRequest{
		Page:     req.Page,
		PageSize: req.PageSize,
		Keyword:  req.Keyword,
		Name:     req.Name,
	})
	if err != nil {
		return nil, err
	}

	reply := dictionary.ListDictionaryReply{Total: total}
	for _, item := range list {
		reply.List = append(reply.List, &dictionary.ListDictionaryReply_Dictionary{
			Id:          item.Id,
			Keyword:     item.Keyword,
			Type:        item.Type,
			Name:        item.Name,
			Description: item.Description,
			CreatedAt:   uint32(item.CreatedAt),
			UpdatedAt:   uint32(item.UpdatedAt),
		})
	}
	return &reply, nil
}

// CreateDictionary 创建字典目录
func (s *Dictionary) CreateDictionary(c context.Context, req *dictionary.CreateDictionaryRequest) (*dictionary.CreateDictionaryReply, error) {
	id, err := s.srv.CreateDictionary(kratosx.MustContext(c), &entity.Dictionary{
		Keyword:     req.Keyword,
		Name:        req.Name,
		Type:        req.Type,
		Description: req.Description,
	})
	if err != nil {
		return nil, err
	}
	return &dictionary.CreateDictionaryReply{Id: id}, nil
}

// UpdateDictionary 更新字典目录
func (s *Dictionary) UpdateDictionary(c context.Context, req *dictionary.UpdateDictionaryRequest) (*dictionary.UpdateDictionaryReply, error) {
	if err := s.srv.UpdateDictionary(kratosx.MustContext(c), &entity.Dictionary{
		BaseTenantModel: model.BaseTenantModel{Id: req.Id},
		Keyword:         req.Keyword,
		Name:            req.Name,
		Type:            req.Type,
		Description:     req.Description,
	}); err != nil {
		return nil, err
	}

	return &dictionary.UpdateDictionaryReply{}, nil
}

// DeleteDictionary 删除字典目录
func (s *Dictionary) DeleteDictionary(c context.Context, req *dictionary.DeleteDictionaryRequest) (*dictionary.DeleteDictionaryReply, error) {
	return &dictionary.DeleteDictionaryReply{}, s.srv.DeleteDictionary(kratosx.MustContext(c), req.Id)
}

// ListDictionaryValue 获取字典值目录列表
func (s *Dictionary) ListDictionaryValue(c context.Context, req *dictionary.ListDictionaryValueRequest) (*dictionary.ListDictionaryValueReply, error) {
	ctx := kratosx.MustContext(c)
	result, total, err := s.srv.ListDictionaryValue(ctx, &types.ListDictionaryValueRequest{
		Page:         req.Page,
		PageSize:     req.PageSize,
		DictionaryId: req.DictionaryId,
		Label:        req.Label,
		Value:        req.Value,
		Status:       req.Status,
	})
	if err != nil {
		return nil, err
	}

	reply := dictionary.ListDictionaryValueReply{Total: total}
	if err := value.Transform(result, &reply.List); err != nil {
		return nil, errors.TransformError()
	}
	return &reply, nil
}

// CreateDictionaryValue 创建字典值目录
func (s *Dictionary) CreateDictionaryValue(c context.Context, req *dictionary.CreateDictionaryValueRequest) (*dictionary.CreateDictionaryValueReply, error) {
	id, err := s.srv.CreateDictionaryValue(kratosx.MustContext(c), &entity.DictionaryValue{
		DictionaryId: req.DictionaryId,
		ParentId:     req.ParentId,
		Label:        req.Label,
		Value:        req.Value,
		Status:       req.Status,
		Weight:       req.Weight,
		Type:         req.Type,
		Extra:        req.Extra,
		Description:  req.Description,
	})
	if err != nil {
		return nil, err
	}
	return &dictionary.CreateDictionaryValueReply{Id: id}, nil
}

// UpdateDictionaryValue 更新字典值目录
func (s *Dictionary) UpdateDictionaryValue(c context.Context, req *dictionary.UpdateDictionaryValueRequest) (*dictionary.UpdateDictionaryValueReply, error) {
	if err := s.srv.UpdateDictionaryValue(kratosx.MustContext(c), &entity.DictionaryValue{
		BaseModel:    model.BaseModel{Id: req.Id},
		DictionaryId: req.DictionaryId,
		ParentId:     req.ParentId,
		Label:        req.Label,
		Value:        req.Value,
		Weight:       req.Weight,
		Type:         req.Type,
		Extra:        req.Extra,
		Description:  req.Description,
		Status:       req.Status,
	}); err != nil {
		return nil, err
	}

	return &dictionary.UpdateDictionaryValueReply{}, nil
}

// UpdateDictionaryValueStatus 更新字典值目录状态
func (s *Dictionary) UpdateDictionaryValueStatus(c context.Context, req *dictionary.UpdateDictionaryValueStatusRequest) (*dictionary.UpdateDictionaryValueStatusReply, error) {
	return &dictionary.UpdateDictionaryValueStatusReply{}, s.srv.UpdateDictionaryValueStatus(kratosx.MustContext(c), req.Id, req.Status)
}

// DeleteDictionaryValue 删除字典值目录
func (s *Dictionary) DeleteDictionaryValue(c context.Context, req *dictionary.DeleteDictionaryValueRequest) (*dictionary.DeleteDictionaryValueReply, error) {
	return &dictionary.DeleteDictionaryValueReply{}, s.srv.DeleteDictionaryValue(kratosx.MustContext(c), req.Id)
}

func (s *Dictionary) GetDictionaryValues(c context.Context, req *dictionary.GetDictionaryValuesRequest) (*dictionary.GetDictionaryValuesReply, error) {
	res, err := s.srv.GetDictionaryValues(kratosx.MustContext(c), req.Keywords)
	if err != nil {
		return nil, err
	}

	reply := dictionary.GetDictionaryValuesReply{Dict: make(map[string]*dictionary.GetDictionaryValuesReply_Value)}
	for key, values := range res {
		if reply.Dict[key] == nil {
			reply.Dict[key] = &dictionary.GetDictionaryValuesReply_Value{
				List: make([]*dictionary.GetDictionaryValuesReply_Value_Item, 0),
			}
		}
		if err := value.Transform(values, &reply.Dict[key].List); err != nil {
			return nil, errors.TransformError(err.Error())
		}
	}
	return &reply, nil
}
