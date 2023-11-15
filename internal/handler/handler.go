package handler

import (
	v1 "manager/api/v1"
	"manager/config"
	"manager/internal/logic"
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
	}
}
