package auth

import (
	"github.com/limes-cloud/kratosx"
	"github.com/limes-cloud/kratosx/pkg/valx"
	"github.com/limes-cloud/manager/api/manager/errors"

	"github.com/limes-cloud/manager/internal/conf"
	"github.com/limes-cloud/manager/internal/pkg/md"
)

type UseCase struct {
	conf *conf.Config
}

func NewUseCase(conf *conf.Config) *UseCase {
	return &UseCase{
		conf: conf,
	}
}

// Auth 外部接口鉴权
func (u *UseCase) Auth(ctx kratosx.Context, in *AuthRequest) (*md.Auth, error) {
	info, err := md.Get(ctx)
	if err != nil {
		return nil, errors.ForbiddenError()
	}

	if valx.InList(ctx.Config().App().Authentication.SkipRole, info.RoleKeyword) {
		return info, nil
	}

	author := ctx.Authentication().Enforce()
	isAuth, _ := author.Enforce(info.RoleKeyword, in.Path, in.Method)
	if !isAuth {
		return nil, errors.ForbiddenError()
	}

	return info, nil
}
