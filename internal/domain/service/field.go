package service

import (
	"google.golang.org/protobuf/proto"

	"github.com/limes-cloud/manager/api/errors"
	"github.com/limes-cloud/manager/internal/core"

	"github.com/limes-cloud/manager/internal/domain/entity"
	"github.com/limes-cloud/manager/internal/domain/repository"
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

// ListFieldOptions 获取组件类型和校验规则枚举
func (srv *Field) ListFieldOptions() *types.FieldOptionsReply {
	return &types.FieldOptionsReply{
		Types: []types.FieldOptionItem{
			{Value: "input", Name: "单行文本"},
			{Value: "textarea", Name: "多行文本"},
			{Value: "number", Name: "数字输入"},
			{Value: "password", Name: "密码输入"},
			{Value: "radio", Name: "单项选择"},
			{Value: "checkbox", Name: "多项选择"},
			{Value: "select", Name: "下拉选择"},
			{Value: "time", Name: "时间选择"},
			{Value: "date", Name: "日期选择"},
			{Value: "datetime", Name: "日期时间"},
			{Value: "upload", Name: "文件上传"},
		},
		Rules: []types.FieldOptionItem{
			{Value: "phone", Name: "手机号"},
			{Value: "email", Name: "邮箱"},
			{Value: "idcard", Name: "身份证"},
			{Value: "regex", Name: "正则表达式"},
			{Value: "enum", Name: "枚举值"},
			{Value: "range", Name: "数值范围"},
			{Value: "length", Name: "字符长度"},
		},
	}
}

// ListField 获取字段列表
func (srv *Field) ListField(ctx core.Context, req *types.ListFieldRequest) ([]*entity.Field, uint32, error) {
	list, total, err := srv.repo.ListField(ctx, req)
	if err != nil {
		ctx.Logger().Warnw("msg", "list field error", "err", err.Error())
		return nil, 0, errors.ListError(err.Error())
	}
	return list, total, nil
}

// CreateField 创建字段
func (srv *Field) CreateField(ctx core.Context, f *entity.Field) (uint32, error) {
	f.Status = proto.Bool(false)
	id, err := srv.repo.CreateField(ctx, f)
	if err != nil {
		ctx.Logger().Warnw("msg", "create field error", "err", err.Error())
		return 0, errors.CreateError(err.Error())
	}
	return id, nil
}

// UpdateField 更新字段
func (srv *Field) UpdateField(ctx core.Context, f *entity.Field) error {
	if err := srv.repo.UpdateField(ctx, f); err != nil {
		ctx.Logger().Warnw("msg", "update field error", "err", err.Error())
		return errors.UpdateError(err.Error())
	}
	return nil
}

// DeleteField 删除字段
func (srv *Field) DeleteField(ctx core.Context, id uint32) error {
	if err := srv.repo.DeleteField(ctx, id); err != nil {
		ctx.Logger().Warnw("msg", "delete field error", "err", err.Error())
		return errors.DeleteError()
	}
	return nil
}
