package handler

import (
	"context"

	"github.com/limes-cloud/kratosx"
	"google.golang.org/protobuf/types/known/emptypb"

	v1 "github.com/limes-cloud/manager/api/v1"
)

func (s *Service) GetDict(ctx context.Context, in *v1.GetDictRequest) (*v1.Dict, error) {
	return s.dict.Get(kratosx.MustContext(ctx), in)
}

func (s *Service) PageDict(ctx context.Context, in *v1.PageDictRequest) (*v1.PageDictReply, error) {
	return s.dict.Page(kratosx.MustContext(ctx), in)
}

func (s *Service) AddDict(ctx context.Context, in *v1.AddDictRequest) (*emptypb.Empty, error) {
	return s.dict.Add(kratosx.MustContext(ctx), in)
}

func (s *Service) UpdateDict(ctx context.Context, in *v1.UpdateDictRequest) (*emptypb.Empty, error) {
	return s.dict.Update(kratosx.MustContext(ctx), in)
}

func (s *Service) DeleteDict(ctx context.Context, in *v1.DeleteDictRequest) (*emptypb.Empty, error) {
	return s.dict.Delete(kratosx.MustContext(ctx), in)
}
