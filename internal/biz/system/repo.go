package system

import "github.com/limes-cloud/kratosx"

type Repo interface {
	// GetDictionaryValues 获取指定标识集合的值
	GetDictionaryValues(ctx kratosx.Context, keywords []string) ([]*DictionaryValue, error)
}
