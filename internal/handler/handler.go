package handler

import (
	v1 "github.com/limes-cloud/manager/api/v1"
	"github.com/limes-cloud/manager/config"
	"github.com/limes-cloud/manager/internal/logic"
)

type Service struct {
	v1.UnimplementedServiceServer
	role       *logic.Role
	menu       *logic.Menu
	department *logic.Department
	user       *logic.User
	userRole   *logic.UserRole
	roleMenu   *logic.RoleMenu
	auth       *logic.Auth
	setting    *logic.Setting
}

func New(conf *config.Config) *Service {
	if err := logic.NewInit(conf).Init(); err != nil {
		panic(err)
	}

	return &Service{
		role:       logic.NewRole(conf),
		menu:       logic.NewMenu(conf),
		department: logic.NewDepartment(conf),
		user:       logic.NewUser(conf),
		userRole:   logic.NewUserRole(conf),
		roleMenu:   logic.NewRoleMenu(conf),
		auth:       logic.NewAuth(conf),
		setting:    logic.NewSetting(conf),
	}
}
