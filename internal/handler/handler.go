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
	job        *logic.Job
	dict       *logic.Dict
	dictValue  *logic.DictValue
}

func New(conf *config.Config) *Service {
	return &Service{
		role:       logic.NewRole(conf),
		menu:       logic.NewMenu(conf),
		department: logic.NewDepartment(conf),
		user:       logic.NewUser(conf),
		userRole:   logic.NewUserRole(conf),
		roleMenu:   logic.NewRoleMenu(conf),
		auth:       logic.NewAuth(conf),
		setting:    logic.NewSetting(conf),
		job:        logic.NewJob(conf),
		dict:       logic.NewDict(conf),
		dictValue:  logic.NewDictValue(conf),
	}
}
