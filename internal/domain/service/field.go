package service

import (
	"github.com/limes-cloud/manager/api/errors"
	"github.com/limes-cloud/manager/internal/core"
	"google.golang.org/protobuf/proto"

	"github.com/limes-cloud/manager/internal/domain/entity"
	"github.com/limes-cloud/manager/internal/domain/repository"
	"github.com/limes-cloud/manager/internal/pkg/field"
	"github.com/limes-cloud/manager/internal/types"
)

type Field struct {
	repo repository.Field
}

func NewField(
	repo repository.Field,
) *Field {
	return &Field{
		repo: repo,
	}
}

// ListFieldType 获取支持的全部字段类型
func (srv *Field) ListFieldType() []*types.FieldType {
	var list []*types.FieldType
	tps := field.New().GetTypes()

	for key, tp := range tps {
		list = append(list, &types.FieldType{Type: key, Name: tp.Name()})
	}
	return list
}

// ListField 获取用户字段列表
func (srv *Field) ListField(ctx core.Context, req *types.ListFieldRequest) ([]*entity.Field, uint32, error) {
	list, total, err := srv.repo.ListField(ctx, req)
	if err != nil {
		ctx.Logger().Warnw("msg", "list field error", "err", err.Error())
		return nil, 0, errors.ListError(err.Error())
	}
	return list, total, nil
}

// CreateField 创建用户字段
func (srv *Field) CreateField(ctx core.Context, field *entity.Field) (uint32, error) {
	field.Status = proto.Bool(false)
	id, err := srv.repo.CreateField(ctx, field)
	if err != nil {
		ctx.Logger().Warnw("msg", "create field error", "err", err.Error())
		return 0, errors.CreateError(err.Error())
	}
	return id, nil
}

// UpdateField 更新用户字段
func (srv *Field) UpdateField(ctx core.Context, field *entity.Field) error {
	if err := srv.repo.UpdateField(ctx, field); err != nil {
		ctx.Logger().Warnw("msg", "update field error", "err", err.Error())
		return errors.UpdateError(err.Error())
	}
	return nil
}

// DeleteField 删除用户字段
func (srv *Field) DeleteField(ctx core.Context, id uint32) error {
	if err := srv.repo.DeleteField(ctx, id); err != nil {
		ctx.Logger().Warnw("msg", "delete field error", "err", err.Error())
		return errors.DeleteError()
	}
	return nil
}
