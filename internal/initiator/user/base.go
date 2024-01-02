package user

import (
	"github.com/limes-cloud/kratosx/types"
	"google.golang.org/protobuf/proto"

	"github.com/limes-cloud/manager/consts"
	"github.com/limes-cloud/manager/internal/model"
	"github.com/limes-cloud/manager/pkg/util"
)

var user = &model.User{
	BaseModel: types.BaseModel{
		ID: consts.SuperAdmin,
	},
	RoleID:       consts.SuperAdmin,
	DepartmentID: consts.SuperAdmin,
	Name:         "超级管理员",
	Nickname:     "超级管理员",
	Gender:       "M",
	Phone:        "18888888888",
	Email:        "1280291001@qq.com",
	Password:     util.ParsePwd("123456"),
	Status:       proto.Bool(true),
}
