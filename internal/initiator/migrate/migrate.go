package migrate

import (
	"github.com/limes-cloud/kratosx"
	gte "github.com/limes-cloud/kratosx/library/db/gormtranserror"

	"github.com/limes-cloud/manager/config"
	"github.com/limes-cloud/manager/internal/model"
)

func Init(ctx kratosx.Context, config *config.Config) {
	db := ctx.DB().Set("gorm:table_options", "ENGINE=InnoDB CHARSET=utf8mb4")
	_ = db.Set("gorm:table_options", "COMMENT='菜单信息' ENGINE=InnoDB CHARSET=utf8mb4").AutoMigrate(model.Menu{})
	_ = db.Set("gorm:table_options", "COMMENT='角色信息' ENGINE=InnoDB CHARSET=utf8mb4").AutoMigrate(model.Role{})
	_ = db.Set("gorm:table_options", "COMMENT='角色层级信息' ENGINE=InnoDB CHARSET=utf8mb4").AutoMigrate(model.RoleClosure{})
	_ = db.Set("gorm:table_options", "COMMENT='角色菜单信息' ENGINE=InnoDB CHARSET=utf8mb4").AutoMigrate(model.RoleMenu{})
	_ = db.Set("gorm:table_options", "COMMENT='部门信息' ENGINE=InnoDB CHARSET=utf8mb4").AutoMigrate(model.Department{})
	_ = db.Set("gorm:table_options", "COMMENT='部门层级信息' ENGINE=InnoDB CHARSET=utf8mb4").AutoMigrate(model.DepartmentClosure{})
	_ = db.Set("gorm:table_options", "COMMENT='职位信息' ENGINE=InnoDB CHARSET=utf8mb4").AutoMigrate(model.Job{})
	_ = db.Set("gorm:table_options", "COMMENT='用户信息' ENGINE=InnoDB CHARSET=utf8mb4").AutoMigrate(model.User{})
	_ = db.Set("gorm:table_options", "COMMENT='用户角色信息' ENGINE=InnoDB CHARSET=utf8mb4").AutoMigrate(model.UserRole{})
	_ = db.Set("gorm:table_options", "COMMENT='用户职位信息' ENGINE=InnoDB CHARSET=utf8mb4").AutoMigrate(model.UserJob{})
	_ = db.Set("gorm:table_options", "COMMENT='字典信息' ENGINE=InnoDB CHARSET=utf8mb4").AutoMigrate(model.Dict{})
	_ = db.Set("gorm:table_options", "COMMENT='字典值信息' ENGINE=InnoDB CHARSET=utf8mb4").AutoMigrate(model.DictValue{})

	// 重新载入gorm错误插件
	gte.NewGlobalGormErrorPlugin().Initialize(ctx.DB())
}
