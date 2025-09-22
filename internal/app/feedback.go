package app

import (
	"context"

	"github.com/limes-cloud/kratosx/pkg/value"

	"github.com/limes-cloud/kratosx/model"

	"github.com/limes-cloud/kratosx/model/page"

	"github.com/limes-cloud/manager/internal/core"

	"github.com/limes-cloud/manager/api/errors"
	"github.com/limes-cloud/manager/api/feedback"

	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/limes-cloud/manager/internal/domain/entity"
	"github.com/limes-cloud/manager/internal/domain/service"
	"github.com/limes-cloud/manager/internal/infra/dbs"
	"github.com/limes-cloud/manager/internal/types"
)

type Feedback struct {
	feedback.UnimplementedFeedbackServer
	srv *service.Feedback
}

func NewFeedback() *Feedback {
	return &Feedback{
		srv: service.NewFeedback(
			dbs.NewFeedback(),
		),
	}
}

func init() {
	register(func(hs *http.Server, gs *grpc.Server) {
		srv := NewFeedback()
		feedback.RegisterFeedbackHTTPServer(hs, srv)
		feedback.RegisterFeedbackServer(gs, srv)
	})
}

// ListFeedbackCategory 获取反馈建议分类列表
func (fb *Feedback) ListFeedbackCategory(c context.Context, req *feedback.ListFeedbackCategoryRequest) (*feedback.ListFeedbackCategoryReply, error) {
	list, total, err := fb.srv.ListFeedbackCategory(core.MustContext(c), &types.ListFeedbackCategoryRequest{
		Search: page.Search{
			Page:     req.Page,
			PageSize: req.PageSize,
			Order:    req.Order,
			OrderBy:  req.OrderBy,
		},
		Name: req.Name,
	})
	if err != nil {
		return nil, err
	}
	reply := feedback.ListFeedbackCategoryReply{Total: total}
	for _, item := range list {
		reply.List = append(reply.List, &feedback.ListFeedbackCategoryReply_FeedbackCategory{
			Id:        item.Id,
			Name:      item.Name,
			CreatedAt: uint32(item.CreatedAt),
		})
	}
	return &reply, nil
}

// CreateFeedbackCategory 创建反馈建议分类
func (fb *Feedback) CreateFeedbackCategory(c context.Context, req *feedback.CreateFeedbackCategoryRequest) (*feedback.CreateFeedbackCategoryReply, error) {
	id, err := fb.srv.CreateFeedbackCategory(core.MustContext(c), &entity.FeedbackCategory{
		Name: req.Name,
	})
	if err != nil {
		return nil, err
	}
	return &feedback.CreateFeedbackCategoryReply{Id: id}, nil
}

// UpdateFeedbackCategory 更新反馈建议分类
func (fb *Feedback) UpdateFeedbackCategory(c context.Context, req *feedback.UpdateFeedbackCategoryRequest) (*feedback.UpdateFeedbackCategoryReply, error) {
	if err := fb.srv.UpdateFeedbackCategory(core.MustContext(c), &entity.FeedbackCategory{
		BaseTenantModel: model.BaseTenantModel{Id: req.Id},
		Name:            req.Name,
	}); err != nil {
		return nil, err
	}
	return &feedback.UpdateFeedbackCategoryReply{}, nil
}

// DeleteFeedbackCategory 删除反馈建议分类
func (fb *Feedback) DeleteFeedbackCategory(c context.Context, req *feedback.DeleteFeedbackCategoryRequest) (*feedback.DeleteFeedbackCategoryReply, error) {
	if err := fb.srv.DeleteFeedbackCategory(core.MustContext(c), req.Id); err != nil {
		return nil, err
	}
	return &feedback.DeleteFeedbackCategoryReply{}, nil
}

// ListFeedback 获取反馈建议列表
func (fb *Feedback) ListFeedback(c context.Context, req *feedback.ListFeedbackRequest) (*feedback.ListFeedbackReply, error) {
	list, total, err := fb.srv.ListFeedback(core.MustContext(c), &types.ListFeedbackRequest{
		Search: page.Search{
			Page:     req.Page,
			PageSize: req.PageSize,
			Order:    req.Order,
			OrderBy:  req.OrderBy,
		},
		AppId:      req.AppId,
		CategoryId: req.CategoryId,
		Status:     req.Status,
		Platform:   req.Platform,
	})
	if err != nil {
		return nil, err
	}
	reply := feedback.ListFeedbackReply{Total: total}
	if err := value.Transform(list, &reply.List); err != nil {
		return nil, errors.TransformError()
	}
	return &reply, nil
}

// CreateFeedback 创建反馈建议
func (fb *Feedback) CreateFeedback(c context.Context, req *feedback.CreateFeedbackRequest) (*feedback.CreateFeedbackReply, error) {
	id, err := fb.srv.CreateFeedback(core.MustContext(c), &entity.Feedback{
		AppId:      req.AppId,
		CategoryId: req.CategoryId,
		Title:      req.Title,
		Content:    req.Content,
		Images:     req.Images,
		Contact:    req.Contact,
		Device:     req.Device,
		Platform:   req.Platform,
		Version:    req.Version,
	})
	if err != nil {
		return nil, err
	}
	return &feedback.CreateFeedbackReply{Id: id}, nil
}

// DeleteFeedback 删除反馈建议
func (fb *Feedback) DeleteFeedback(c context.Context, req *feedback.DeleteFeedbackRequest) (*feedback.DeleteFeedbackReply, error) {
	if err := fb.srv.DeleteFeedback(core.MustContext(c), req.Id); err != nil {
		return nil, err
	}
	return &feedback.DeleteFeedbackReply{}, nil
}

// UpdateFeedback 更新反馈建议
func (fb *Feedback) UpdateFeedback(c context.Context, req *feedback.UpdateFeedbackRequest) (*feedback.UpdateFeedbackReply, error) {
	if err := fb.srv.UpdateFeedback(core.MustContext(c), &entity.Feedback{
		BaseTenantUserModel: model.BaseTenantUserModel{Id: req.Id},
		Status:              req.Status,
		ProcessedResult:     req.ProcessedResult,
	}); err != nil {
		return nil, err
	}
	return &feedback.UpdateFeedbackReply{}, nil
}
