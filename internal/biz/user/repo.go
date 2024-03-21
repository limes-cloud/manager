package user

import "github.com/limes-cloud/kratosx"

type Repo interface {
	GetUser(ctx kratosx.Context, id uint32) (*User, error)
	GetUserByPhone(ctx kratosx.Context, phone string) (*User, error)
	GetUserByEmail(ctx kratosx.Context, email string) (*User, error)
	PageUser(ctx kratosx.Context, in *PageUserRequest) ([]*User, uint32, error)
	HasRole(ctx kratosx.Context, uid, rid uint32) bool
	AddUser(ctx kratosx.Context, in *User) (uint32, error)
	UpdateUser(ctx kratosx.Context, in *User) error
	DeleteUser(ctx kratosx.Context, id uint32) error
}
