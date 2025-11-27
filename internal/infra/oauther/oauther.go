package oauther

import (
	"errors"
	"sort"

	"github.com/limes-cloud/manager/internal/domain/entity"
	"github.com/limes-cloud/manager/internal/domain/repository"
	"github.com/limes-cloud/manager/internal/infra/oauther/email"
	"github.com/limes-cloud/manager/internal/infra/oauther/yiban"
	"github.com/limes-cloud/manager/internal/types"
)

var ins = make(map[string]*oauther)

type (
	InitFunc func(req *entity.OAuther) (repository.OAutherFunc, error)

	Name interface {
		Keyword() string
		Name() string
	}
)

type oauther struct {
	keyword string
	name    string
	initFn  InitFunc
}

func register(keyword string, name string, init InitFunc) {
	ins[keyword] = &oauther{
		keyword: keyword,
		name:    name,
		initFn:  init,
	}
}

func init() {
	register("yiban", "易班", yiban.NewYiBan)
	register("email", "邮箱", email.NewEmail)
}

type OAuther struct{}

func New() *OAuther {
	return &OAuther{}
}

// GetOAuther 获取指定的授权器
func (oa OAuther) GetOAuther(req *entity.OAuther) (repository.OAutherFunc, error) {
	if req == nil || ins[req.Type] == nil {
		return nil, errors.New("not fount oauther " + req.Type)
	}
	return ins[req.Type].initFn(req)
}

// ListOAutherType 获取实现的授权器名称
func (oa OAuther) ListOAutherType() []*types.OAutherType {
	var list []*types.OAutherType
	for _, item := range ins {
		list = append(list, &types.OAutherType{
			Keyword: item.keyword,
			Name:    item.name,
		})
	}
	sort.Slice(list, func(i, j int) bool {
		return list[i].Keyword > list[j].Keyword
	})
	return list
}
