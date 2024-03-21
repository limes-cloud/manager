package dictionary

import "github.com/limes-cloud/kratosx"

type Repo interface {
	PageDictionary(ctx kratosx.Context, in *PageDictionaryRequest) ([]*Dictionary, uint32, error)
	AddDictionary(ctx kratosx.Context, in *Dictionary) (uint32, error)
	UpdateDictionary(ctx kratosx.Context, in *Dictionary) error
	DeleteDictionary(ctx kratosx.Context, id uint32) error
	GetDictionaryValue(ctx kratosx.Context, keyword string) ([]*DictionaryValue, error)
	PageDictionaryValue(ctx kratosx.Context, in *PageDictionaryValueRequest) ([]*DictionaryValue, uint32, error)
	AddDictionaryValue(ctx kratosx.Context, in *DictionaryValue) (uint32, error)
	UpdateDictionaryValue(ctx kratosx.Context, in *DictionaryValue) error
	DeleteDictionaryValue(ctx kratosx.Context, id uint32) error
}
