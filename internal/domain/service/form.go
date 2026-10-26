package service

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/limes-cloud/manager/api/errors"
	"github.com/limes-cloud/manager/internal/core"
	"github.com/limes-cloud/manager/internal/domain/entity"
	"github.com/limes-cloud/manager/internal/domain/repository"
	"github.com/limes-cloud/manager/internal/types"
)

type Form struct {
	repo repository.UserForm
}

func NewForm(repo repository.UserForm) *Form {
	return &Form{repo: repo}
}

// GetUserForm 获取表单
func (s *Form) GetUserForm(ctx core.Context, req *types.GetUserFormRequest) (*entity.UserForm, error) {
	res, err := s.repo.GetUserForm(ctx, req)
	if err != nil {
		ctx.Logger().Warnw("msg", "get user form error", "err", err.Error())
		return nil, errors.GetError()
	}
	return res, nil
}

// ListUserForm 分页查询表单列表
func (s *Form) ListUserForm(ctx core.Context, req *types.ListUserFormRequest) ([]*entity.UserForm, uint32, error) {
	list, total, err := s.repo.ListUserForm(ctx, req)
	if err != nil {
		ctx.Logger().Warnw("msg", "list user form error", "err", err.Error())
		return nil, 0, errors.ListError()
	}
	return list, total, nil
}

// CreateUserForm 创建表单
func (s *Form) CreateUserForm(ctx core.Context, ent *entity.UserForm) (uint32, error) {
	status := false
	ent.Status = &status
	ent.Version = 1
	id, err := s.repo.CreateUserForm(ctx, ent)
	if err != nil {
		ctx.Logger().Warnw("msg", "create user form error", "err", err.Error())
		return 0, errors.CreateError()
	}
	return id, nil
}

// UpdateUserForm 更新表单（version 自增，status 回到草稿）
func (s *Form) UpdateUserForm(ctx core.Context, ent *entity.UserForm) error {
	// 先获取现有记录取 version
	existing, err := s.repo.GetUserForm(ctx, &types.GetUserFormRequest{Id: &ent.Id})
	if err != nil {
		ctx.Logger().Warnw("msg", "get user form for update error", "err", err.Error())
		return errors.GetError()
	}
	ent.Version = existing.Version + 1
	status := false
	ent.Status = &status
	if err := s.repo.UpdateUserForm(ctx, ent); err != nil {
		ctx.Logger().Warnw("msg", "update user form error", "err", err.Error())
		return errors.UpdateError()
	}
	return nil
}

// PublishUserForm 发布表单
func (s *Form) PublishUserForm(ctx core.Context, id uint32) error {
	ent := &entity.UserForm{}
	ent.Id = id
	status := true
	ent.Status = &status
	if err := s.repo.UpdateUserForm(ctx, ent); err != nil {
		ctx.Logger().Warnw("msg", "publish user form error", "err", err.Error())
		return errors.UpdateError()
	}
	return nil
}

// DeleteUserForm 删除表单
func (s *Form) DeleteUserForm(ctx core.Context, id uint32) error {
	if err := s.repo.DeleteUserForm(ctx, id); err != nil {
		ctx.Logger().Warnw("msg", "delete user form error", "err", err.Error())
		return errors.DeleteError()
	}
	return nil
}

// RemoteOptionItem 远程选项结果
type RemoteOptionItem struct {
	Label string `json:"label"`
	Value string `json:"value"`
}

// FetchRemoteOptionsRequest 远程选项请求参数
type FetchRemoteOptionsRequest struct {
	URL        string            `json:"url"`
	Method     string            `json:"method"`
	DataPath   string            `json:"dataPath"`
	LabelField string            `json:"labelField"`
	ValueField string            `json:"valueField"`
	Headers    map[string]string `json:"headers"`
	Params     string            `json:"params"`
	Body       string            `json:"body"`
}

// FetchRemoteOptions 远程数据代理
func (s *Form) FetchRemoteOptions(ctx core.Context, req *FetchRemoteOptionsRequest) ([]*RemoteOptionItem, error) {
	// 构造请求 URL
	url := req.URL
	if req.Params != "" {
		if strings.Contains(url, "?") {
			url += "&" + req.Params
		} else {
			url += "?" + req.Params
		}
	}

	// 设置默认值
	method := req.Method
	if method == "" {
		method = "GET"
	}
	labelField := req.LabelField
	if labelField == "" {
		labelField = "label"
	}
	valueField := req.ValueField
	if valueField == "" {
		valueField = "value"
	}

	// 构造 HTTP 请求
	var bodyReader io.Reader
	if method == "POST" && req.Body != "" {
		bodyReader = strings.NewReader(req.Body)
	}

	httpReq, err := http.NewRequestWithContext(ctx.Ctx(), method, url, bodyReader)
	if err != nil {
		ctx.Logger().Warnw("msg", "create remote request error", "err", err.Error())
		return nil, errors.SystemError("构造远程请求失败")
	}

	// 设置请求头
	for k, v := range req.Headers {
		httpReq.Header.Set(k, v)
	}
	if method == "POST" && httpReq.Header.Get("Content-Type") == "" {
		httpReq.Header.Set("Content-Type", "application/json")
	}

	// 发起请求（10s 超时）
	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(httpReq)
	if err != nil {
		ctx.Logger().Warnw("msg", "fetch remote options error", "url", url, "err", err.Error())
		return nil, errors.SystemError("远程请求失败")
	}
	defer resp.Body.Close()

	// 读取响应
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		ctx.Logger().Warnw("msg", "read remote response error", "err", err.Error())
		return nil, errors.SystemError("读取远程响应失败")
	}

	// 解析 JSON
	var jsonData any
	if err := json.Unmarshal(respBody, &jsonData); err != nil {
		ctx.Logger().Warnw("msg", "parse remote response error", "err", err.Error())
		return nil, errors.SystemError("解析远程响应失败")
	}

	// 按 dataPath 提取数据
	data := jsonData
	if req.DataPath != "" {
		parts := strings.Split(req.DataPath, ".")
		for _, part := range parts {
			if m, ok := data.(map[string]any); ok {
				data = m[part]
			} else {
				ctx.Logger().Warnw("msg", "remote data path invalid", "dataPath", req.DataPath)
				return nil, errors.SystemError("数据路径无效")
			}
		}
	}

	// 转换为选项列表
	arr, ok := data.([]any)
	if !ok {
		ctx.Logger().Warnw("msg", "remote data not array", "dataPath", req.DataPath)
		return nil, errors.SystemError("远程数据格式错误，期望数组")
	}

	items := make([]*RemoteOptionItem, 0, len(arr))
	for _, raw := range arr {
		m, ok := raw.(map[string]any)
		if !ok {
			continue
		}
		items = append(items, &RemoteOptionItem{
			Label: fmt.Sprintf("%v", m[labelField]),
			Value: fmt.Sprintf("%v", m[valueField]),
		})
	}

	return items, nil
}
