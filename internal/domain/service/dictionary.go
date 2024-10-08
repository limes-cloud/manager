package service

import (
	"github.com/limes-cloud/kratosx"
	"github.com/limes-cloud/kratosx/pkg/tree"
	"google.golang.org/protobuf/proto"

	"github.com/limes-cloud/manager/api/manager/errors"
	"github.com/limes-cloud/manager/internal/conf"
	"github.com/limes-cloud/manager/internal/domain/entity"
	"github.com/limes-cloud/manager/internal/domain/repository"
	"github.com/limes-cloud/manager/internal/types"
)

const (
	dataTypeList = "list"
	dataTypeTree = "tree"
)

type Dictionary struct {
	conf *conf.Config
	repo repository.Dictionary
}

func NewDictionary(config *conf.Config, repo repository.Dictionary) *Dictionary {
	return &Dictionary{conf: config, repo: repo}
}

// ListDictionary 获取字典目录列表
func (u *Dictionary) ListDictionary(ctx kratosx.Context, req *types.ListDictionaryRequest) ([]*entity.Dictionary, uint32, error) {
	list, total, err := u.repo.ListDictionary(ctx, req)
	if err != nil {
		ctx.Logger().Warnw("msg", "list dictionary error", "err", err.Error())
		return nil, 0, errors.ListError()
	}
	return list, total, nil
}

// CreateDictionary 创建字典目录
func (u *Dictionary) CreateDictionary(ctx kratosx.Context, req *entity.Dictionary) (uint32, error) {
	id, err := u.repo.CreateDictionary(ctx, req)
	if err != nil {
		ctx.Logger().Warnw("msg", "create dictionary error", "err", err.Error())
		return 0, errors.CreateError(err.Error())
	}
	return id, nil
}

// UpdateDictionary 更新字典目录
func (u *Dictionary) UpdateDictionary(ctx kratosx.Context, req *entity.Dictionary) error {
	if err := u.repo.UpdateDictionary(ctx, req); err != nil {
		ctx.Logger().Warnw("msg", "update dictionary error", "err", err.Error())
		return errors.UpdateError(err.Error())
	}
	return nil
}

// DeleteDictionary 删除字典目录
func (u *Dictionary) DeleteDictionary(ctx kratosx.Context, id uint32) error {
	if err := u.repo.DeleteDictionary(ctx, id); err != nil {
		ctx.Logger().Warnw("msg", "delete dictionary error", "err", err.Error())
		return errors.DeleteError(err.Error())
	}
	return nil
}

// ListDictionaryValue 获取字典值目录列表
func (u *Dictionary) ListDictionaryValue(ctx kratosx.Context, req *types.ListDictionaryValueRequest) ([]*entity.DictionaryValue, uint32, error) {
	// 获取任务数据类型
	task, err := u.repo.GetDictionary(ctx, req.DictionaryId)
	if err != nil {
		return nil, 0, errors.ListError(err.Error())
	}
	if task.Type == dataTypeList {
		list, total, err := u.repo.ListDictionaryValue(ctx, req)
		if err != nil {
			ctx.Logger().Warnw("msg", "list dictionary error", "err", err.Error())
			return nil, 0, errors.ListError(err.Error())
		}
		return list, total, nil
	}
	list, err := u.repo.AllDictionaryValue(ctx, &types.AllDictionaryValueRequest{
		DictionaryId: req.DictionaryId,
	})
	if err != nil {
		ctx.Logger().Warnw("msg", "list dictionary error", "err", err.Error())
		return nil, 0, errors.ListError(err.Error())
	}
	total := uint32(len(list))
	return tree.BuildArrayTree(list), total, nil
}

// CreateDictionaryValue 创建字典值目录
func (u *Dictionary) CreateDictionaryValue(ctx kratosx.Context, req *entity.DictionaryValue) (uint32, error) {
	id, err := u.repo.CreateDictionaryValue(ctx, req)
	if err != nil {
		ctx.Logger().Warnw("msg", "create dictionary error", "err", err.Error())
		return 0, errors.CreateError(err.Error())
	}
	return id, nil
}

// UpdateDictionaryValue 更新字典值目录
func (u *Dictionary) UpdateDictionaryValue(ctx kratosx.Context, dictValue *entity.DictionaryValue) error {
	// 获取字典是否为树
	dictionary, err := u.repo.GetDictionary(ctx, dictValue.DictionaryId)
	if err != nil {
		return err
	}

	// 是数的情况下，不能选择自己作为父节点
	if dictionary.Type == "tree" && dictValue.Id == dictValue.ParentId {
		return errors.UpdateError("不能选择自己作为父节点")
	}

	if err := u.repo.UpdateDictionaryValue(ctx, dictValue); err != nil {
		ctx.Logger().Warnw("msg", "update dictionary error", "err", err.Error())
		return errors.UpdateError(err.Error())
	}
	return nil
}

// UpdateDictionaryValueStatus 更新字典值目录状态
func (u *Dictionary) UpdateDictionaryValueStatus(ctx kratosx.Context, id uint32, status bool) error {
	if status {
		if err := u.repo.UpdateDictionaryValueStatus(ctx, id, status); err != nil {
			ctx.Logger().Warnw("msg", "update dictionary value error", "err", err.Error())
			return errors.UpdateError(err.Error())
		}
		return nil
	}

	// 获取值所属字典id
	oldValue, err := u.repo.GetDictionaryValue(ctx, id)
	if err != nil {
		return errors.UpdateError(err.Error())
	}

	// 获取字典信息
	dictionary, err := u.repo.GetDictionary(ctx, oldValue.DictionaryId)
	if err != nil {
		return errors.UpdateError(err.Error())
	}

	// 如果是列表直接更新
	if dictionary.Type == dataTypeList {
		if err := u.repo.UpdateDictionaryValueStatus(ctx, id, status); err != nil {
			ctx.Logger().Warnw("msg", "update dictionary value error", "err", err.Error())
			return errors.UpdateError(err.Error())
		}
		return nil
	}

	list, err := u.repo.AllDictionaryValue(ctx, &types.AllDictionaryValueRequest{DictionaryId: oldValue.DictionaryId})
	if err != nil {
		return errors.UpdateError(err.Error())
	}
	// 如果是树，则更新下属所有节点状态
	td := tree.BuildTreeByID(list, oldValue.Id)
	ids := tree.GetTreeID(td)
	for _, iid := range ids {
		if err := u.repo.UpdateDictionaryValueStatus(ctx, iid, status); err != nil {
			ctx.Logger().Warnw("msg", "update dictionary value error", "err", err.Error())
			return errors.UpdateError(err.Error())
		}
	}
	return nil
}

// DeleteDictionaryValue 删除字典值目录
func (u *Dictionary) DeleteDictionaryValue(ctx kratosx.Context, id uint32) error {
	if err := u.repo.DeleteDictionaryValue(ctx, id); err != nil {
		ctx.Logger().Warnw("msg", "delete dictionary value error", "err", err.Error())
		return errors.DeleteError(err.Error())
	}
	return nil
}

// GetDictionary 获取指定的字典目录
func (u *Dictionary) GetDictionary(ctx kratosx.Context, req *types.GetDictionaryRequest) (*entity.Dictionary, error) {
	var (
		res *entity.Dictionary
		err error
	)

	if req.Id != nil {
		res, err = u.repo.GetDictionary(ctx, *req.Id)
	} else if req.Keyword != nil {
		res, err = u.repo.GetDictionaryByKeyword(ctx, *req.Keyword)
	} else {
		return nil, errors.ParamsError()
	}

	if err != nil {
		return nil, errors.GetError(err.Error())
	}
	return res, nil
}

// GetDictionaryValues 获取字典值目录列表
func (u *Dictionary) GetDictionaryValues(ctx kratosx.Context, keywords []string) (map[string][]*entity.DictionaryValue, error) {
	var reply = make(map[string][]*entity.DictionaryValue)
	for _, key := range keywords {
		// 获取keyword对应的id
		dictionary, err := u.repo.GetDictionaryByKeyword(ctx, key)
		if err != nil {
			ctx.Logger().Warnw("msg", "not found key", "key", key, "err", err.Error())
			continue
		}

		values, err := u.repo.AllDictionaryValue(ctx, &types.AllDictionaryValueRequest{
			DictionaryId: dictionary.Id,
			Status:       proto.Bool(true),
		})
		if err != nil {
			ctx.Logger().Warnw("msg", "not found key value", "key", key, "err", err.Error())
			continue
		}

		if dictionary.Type == dataTypeTree {
			reply[key] = tree.BuildArrayTree(values)
		} else {
			reply[key] = values
		}
	}
	return reply, nil
}
