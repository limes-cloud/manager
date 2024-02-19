package handler

import (
	"context"

	"github.com/limes-cloud/kratosx"
	"google.golang.org/protobuf/types/known/emptypb"

	v1 "github.com/limes-cloud/manager/api/v1"
)

func (s *Service) GetDictValue(ctx context.Context, in *v1.GetDictValueRequest) (*v1.GetDictValueReply, error) {
	return s.dictValue.GetValue(kratosx.MustContext(ctx), in)
}

func (s *Service) PageDictValue(ctx context.Context, in *v1.PageDictValueRequest) (*v1.PageDictValueReply, error) {
	return s.dictValue.Page(kratosx.MustContext(ctx), in)
}

func (s *Service) AddDictValue(ctx context.Context, in *v1.AddDictValueRequest) (*emptypb.Empty, error) {
	return s.dictValue.Add(kratosx.MustContext(ctx), in)
}

func (s *Service) ImportDictValue(ctx context.Context, in *v1.ImportDictValueRequest) (*v1.ImportDictValueReply, error) {
	return s.dictValue.Import(kratosx.MustContext(ctx), in)
}

func (s *Service) UpdateDictValue(ctx context.Context, in *v1.UpdateDictValueRequest) (*emptypb.Empty, error) {
	return s.dictValue.Update(kratosx.MustContext(ctx), in)
}

func (s *Service) DeleteDictValue(ctx context.Context, in *v1.DeleteDictValueRequest) (*emptypb.Empty, error) {
	return s.dictValue.Delete(kratosx.MustContext(ctx), in)
}
