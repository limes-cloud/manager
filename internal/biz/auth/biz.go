package auth

import (
	"encoding/json"
	"fmt"

	"github.com/limes-cloud/kratosx"
	"github.com/limes-cloud/kratosx/pkg/util"

	v1 "github.com/limes-cloud/manager/api/auth/v1"
	"github.com/limes-cloud/manager/api/errors"
	depbiz "github.com/limes-cloud/manager/internal/biz/department"
	menubiz "github.com/limes-cloud/manager/internal/biz/menu"
	objectbiz "github.com/limes-cloud/manager/internal/biz/object"
	"github.com/limes-cloud/manager/internal/config"
	"github.com/limes-cloud/manager/internal/pkg/md"
)

type UseCase struct {
	menuRepo menubiz.Repo
	depRepo  depbiz.Repo
	conf     *config.Config
}

const (
	Add = "add"
	In  = "in"
)

func NewUseCase(conf *config.Config, depRepo depbiz.Repo, menuRepo menubiz.Repo) *UseCase {
	return &UseCase{
		depRepo:  depRepo,
		menuRepo: menuRepo,
		conf:     conf,
	}
}

// Auth 接口鉴权
// 后期计划，在这里增加鉴权资源的缓存，避免频繁查询。
func (u *UseCase) Auth(ctx kratosx.Context, in *v1.AuthRequest) (*v1.AuthReply, error) {
	info, err := md.Get(ctx)
	if err != nil {
		return nil, errors.Forbidden()
	}
	reply := &v1.AuthReply{
		Scope:             make(map[string]*v1.AuthReply_Scope),
		UserId:            info.UserId,
		RoleId:            info.RoleId,
		RoleKeyword:       info.RoleKeyword,
		DepartmentId:      info.DepartmentId,
		DepartmentKeyword: info.DepartmentKeyword,
	}

	author := ctx.Authentication().Enforce()

	var skipRoles []string
	_ = ctx.Config().Value("authentication.skipRole").Scan(&skipRoles)

	if util.InList(skipRoles, info.RoleKeyword) {
		return reply, nil
	}

	isAuth, _ := author.Enforce(info.RoleKeyword, in.Path, in.Method)
	if !isAuth {
		return nil, errors.Forbidden()
	}

	// 获取资源对象权限
	menu, _ := u.menuRepo.GetMenuByApi(ctx, in.Path, in.Method)
	if menu.CheckObject == nil || !*menu.CheckObject || menu.CheckObjectRule == nil {
		return reply, nil
	}

	var rule []objectbiz.ObjectRule
	if err := json.Unmarshal([]byte(*menu.CheckObjectRule), &rule); err != nil {
		ctx.Logger().Errorf("parse object rule error %v", err.Error())
		return reply, nil
	}

	for _, item := range rule {
		values, _ := u.depRepo.GetScope(ctx, info.UserId, item.Object)
		switch item.Operate {
		case Add:
			reply.Scope[item.Field] = &v1.AuthReply_Scope{
				List: values,
			}
		case In:
			sv := in.Data.GetStructValue()
			if sv == nil || sv.Fields[item.Field] == nil {
				continue
			}
			if !util.InList(values, fmt.Sprint(sv.Fields[item.Field].AsInterface())) {
				return nil, errors.Forbidden()
			}
		}
	}
	return reply, nil
}
