package initiator

import (
	"context"

	"github.com/limes-cloud/kratosx"

	"github.com/limes-cloud/manager/config"
	"github.com/limes-cloud/manager/internal/initiator/auth"
	"github.com/limes-cloud/manager/internal/initiator/department"
	"github.com/limes-cloud/manager/internal/initiator/dict"
	"github.com/limes-cloud/manager/internal/initiator/job"
	"github.com/limes-cloud/manager/internal/initiator/menu"
	"github.com/limes-cloud/manager/internal/initiator/migrate"
	"github.com/limes-cloud/manager/internal/initiator/role"
	"github.com/limes-cloud/manager/internal/initiator/user"
)

type Initiator struct {
	conf *config.Config
}

func New(conf *config.Config) *Initiator {
	return &Initiator{
		conf: conf,
	}
}

// Run 执行系统初始化
func (a *Initiator) Run() error {
	ctx := kratosx.MustContext(context.Background())

	// 自动迁移
	migrate.Init(ctx, a.conf)

	// 菜单初始化
	menu.Init(ctx, a.conf)

	// 角色初始化
	role.Init(ctx, a.conf)

	// 初始化部门数据
	department.Init(ctx, a.conf)

	// 职位初始化
	job.Init(ctx, a.conf)

	// 用户初始化
	user.Init(ctx, a.conf)

	// 字典初始化
	dict.Init(ctx, a.conf)

	// 鉴权器初始化
	auth.Init(ctx, a.conf)

	return nil
}
