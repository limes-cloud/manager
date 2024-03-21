package dictionary

import (
	"github.com/limes-cloud/kratosx"

	"github.com/limes-cloud/manager/api/errors"
	"github.com/limes-cloud/manager/internal/config"
)

type UseCase struct {
	repo Repo
	conf *config.Config
}

func NewUseCase(conf *config.Config, repo Repo) *UseCase {
	return &UseCase{
		repo: repo,
		conf: conf,
	}
}

// PageDictionary 获取分页字典信息
func (u *UseCase) PageDictionary(ctx kratosx.Context, in *PageDictionaryRequest) ([]*Dictionary, uint32, error) {
	list, total, err := u.repo.PageDictionary(ctx, in)
	if err != nil {
		return nil, 0, errors.DatabaseFormat(err.Error())
	}
	return list, total, nil
}

// AddDictionary 添加字典信息
func (u *UseCase) AddDictionary(ctx kratosx.Context, in *Dictionary) (uint32, error) {
	id, err := u.repo.AddDictionary(ctx, in)
	if err != nil {
		return 0, errors.DatabaseFormat(err.Error())
	}
	return id, nil
}

// UpdateDictionary 更新字典信息
func (u *UseCase) UpdateDictionary(ctx kratosx.Context, in *Dictionary) error {
	if err := u.repo.UpdateDictionary(ctx, in); err != nil {
		return errors.DatabaseFormat(err.Error())
	}
	return nil
}

// DeleteDictionary 删除字典信息
func (u *UseCase) DeleteDictionary(ctx kratosx.Context, id uint32) error {
	if err := u.repo.DeleteDictionary(ctx, id); err != nil {
		return errors.Database()
	}
	return nil
}

// GetDictionaryValue 获取指定字典值
func (u *UseCase) GetDictionaryValue(ctx kratosx.Context, keyword string) ([]*DictionaryValue, error) {
	list, err := u.repo.GetDictionaryValue(ctx, keyword)
	if err != nil {
		return nil, errors.DatabaseFormat(err.Error())
	}
	return list, nil
}

// PageDictionaryValue 获取分页字典值
func (u *UseCase) PageDictionaryValue(ctx kratosx.Context, in *PageDictionaryValueRequest) ([]*DictionaryValue, uint32, error) {
	list, total, err := u.repo.PageDictionaryValue(ctx, in)
	if err != nil {
		return nil, 0, errors.DatabaseFormat(err.Error())
	}
	return list, total, nil
}

// AddDictionaryValue 删除字典值
func (u *UseCase) AddDictionaryValue(ctx kratosx.Context, in *DictionaryValue) (uint32, error) {
	id, err := u.repo.AddDictionaryValue(ctx, in)
	if err != nil {
		return 0, errors.DatabaseFormat(err.Error())
	}
	return id, nil
}

// UpdateDictionaryValue 更新字典值
func (u *UseCase) UpdateDictionaryValue(ctx kratosx.Context, in *DictionaryValue) error {
	if err := u.repo.UpdateDictionaryValue(ctx, in); err != nil {
		return errors.DatabaseFormat(err.Error())
	}
	return nil
}

// DeleteDictionaryValue 删除字典值
func (u *UseCase) DeleteDictionaryValue(ctx kratosx.Context, id uint32) error {
	if err := u.repo.DeleteDictionaryValue(ctx, id); err != nil {
		return errors.Database()
	}
	return nil
}
