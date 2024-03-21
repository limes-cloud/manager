package service

import (
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"

	authV1 "github.com/limes-cloud/manager/api/auth/v1"
	departmentV1 "github.com/limes-cloud/manager/api/department/v1"
	dictionaryV1 "github.com/limes-cloud/manager/api/dictionary/v1"
	jobV1 "github.com/limes-cloud/manager/api/job/v1"
	menuV1 "github.com/limes-cloud/manager/api/menu/v1"
	objectV1 "github.com/limes-cloud/manager/api/object/v1"
	roleV1 "github.com/limes-cloud/manager/api/role/v1"
	settingV1 "github.com/limes-cloud/manager/api/setting/v1"
	userV1 "github.com/limes-cloud/manager/api/user/v1"
	"github.com/limes-cloud/manager/internal/config"
)

func New(c *config.Config, hs *http.Server, gs *grpc.Server) {
	settingService := NewSettingService(c)
	settingV1.RegisterServiceHTTPServer(hs, settingService)
	settingV1.RegisterServiceServer(gs, settingService)

	objectService := NewObjectService(c)
	objectV1.RegisterServiceHTTPServer(hs, objectService)
	objectV1.RegisterServiceServer(gs, objectService)

	dictionaryService := NewDictionaryService(c)
	dictionaryV1.RegisterServiceHTTPServer(hs, dictionaryService)
	dictionaryV1.RegisterServiceServer(gs, dictionaryService)

	roleService := NewRoleService(c)
	roleV1.RegisterServiceHTTPServer(hs, roleService)
	roleV1.RegisterServiceServer(gs, roleService)

	menuService := NewMenuService(c)
	menuV1.RegisterServiceHTTPServer(hs, menuService)
	menuV1.RegisterServiceServer(gs, menuService)

	departmentService := NewDepartmentService(c)
	departmentV1.RegisterServiceHTTPServer(hs, departmentService)
	departmentV1.RegisterServiceServer(gs, departmentService)

	jobService := NewJobService(c)
	jobV1.RegisterServiceHTTPServer(hs, jobService)
	jobV1.RegisterServiceServer(gs, jobService)

	userService := NewUserService(c)
	userV1.RegisterServiceHTTPServer(hs, userService)
	userV1.RegisterServiceServer(gs, userService)

	authService := NewAuthService(c)
	authV1.RegisterServiceHTTPServer(hs, authService)
	authV1.RegisterServiceServer(gs, authService)
}
