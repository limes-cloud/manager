package dictionary

import (
	"github.com/limes-cloud/kratosx"

	"github.com/limes-cloud/manager/api/manager/errors"
	"github.com/limes-cloud/manager/internal/conf"
)

type UseCase struct {
	conf *conf.Config
	repo Repo
}

func NewUseCase(config *conf.Config, repo Repo) *UseCase {
	return &UseCase{conf: config, repo: repo}
}

// ListDictionary 获取字典目录列表
func (u *UseCase) ListDictionary(ctx kratosx.Context, req *ListDictionaryRequest) ([]*Dictionary, uint32, error) {
	list, total, err := u.repo.ListDictionary(ctx, req)
	if err != nil {
		return nil, 0, errors.ListError(err.Error())
	}
	return list, total, nil
}

// CreateDictionary 创建字典目录
func (u *UseCase) CreateDictionary(ctx kratosx.Context, req *Dictionary) (uint32, error) {
	id, err := u.repo.CreateDictionary(ctx, req)
	if err != nil {
		return 0, errors.CreateError(err.Error())
	}
	return id, nil
}

// UpdateDictionary 更新字典目录
func (u *UseCase) UpdateDictionary(ctx kratosx.Context, req *Dictionary) error {
	if err := u.repo.UpdateDictionary(ctx, req); err != nil {
		return errors.UpdateError(err.Error())
	}
	return nil
}

// DeleteDictionary 删除字典目录
func (u *UseCase) DeleteDictionary(ctx kratosx.Context, ids []uint32) (uint32, error) {
	total, err := u.repo.DeleteDictionary(ctx, ids)
	if err != nil {
		return 0, errors.DeleteError(err.Error())
	}
	return total, nil
}

// ListDictionaryValue 获取字典值目录列表
func (u *UseCase) ListDictionaryValue(ctx kratosx.Context, req *ListDictionaryValueRequest) ([]*DictionaryValue, uint32, error) {
	list, total, err := u.repo.ListDictionaryValue(ctx, req)
	if err != nil {
		return nil, 0, errors.ListError(err.Error())
	}
	return list, total, nil
}

// CreateDictionaryValue 创建字典值目录
func (u *UseCase) CreateDictionaryValue(ctx kratosx.Context, req *DictionaryValue) (uint32, error) {
	id, err := u.repo.CreateDictionaryValue(ctx, req)
	if err != nil {
		return 0, errors.CreateError(err.Error())
	}
	return id, nil
}

// UpdateDictionaryValue 更新字典值目录
func (u *UseCase) UpdateDictionaryValue(ctx kratosx.Context, req *DictionaryValue) error {
	if err := u.repo.UpdateDictionaryValue(ctx, req); err != nil {
		return errors.UpdateError(err.Error())
	}
	return nil
}

// UpdateDictionaryValueStatus 更新字典值目录状态
func (u *UseCase) UpdateDictionaryValueStatus(ctx kratosx.Context, id uint32, status bool) error {
	if err := u.repo.UpdateDictionaryValueStatus(ctx, id, status); err != nil {
		return errors.UpdateError(err.Error())
	}
	return nil
}

// DeleteDictionaryValue 删除字典值目录
func (u *UseCase) DeleteDictionaryValue(ctx kratosx.Context, ids []uint32) (uint32, error) {
	total, err := u.repo.DeleteDictionaryValue(ctx, ids)
	if err != nil {
		return 0, errors.DeleteError(err.Error())
	}
	return total, nil
}

// GetDictionary 获取指定的字典目录
func (u *UseCase) GetDictionary(ctx kratosx.Context, req *GetDictionaryRequest) (*Dictionary, error) {
	var (
		res *Dictionary
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
func (u *UseCase) GetDictionaryValues(ctx kratosx.Context, keywords []string) (map[string][]*DictionaryValue, error) {
	var reply = make(map[string][]*DictionaryValue)
	for _, key := range keywords {
		values, err := u.repo.AllDictionaryValue(ctx, key)
		if err != nil {
			return nil, errors.GetError(err.Error())
		}
		reply[key] = values
	}
	return reply, nil
}
