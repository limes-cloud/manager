package role

import (
	"github.com/limes-cloud/kratosx/types"
	"google.golang.org/protobuf/proto"

	"github.com/limes-cloud/manager/consts"
	"github.com/limes-cloud/manager/internal/model"
)

var role = &model.Role{
	BaseModel: types.BaseModel{
		ID: consts.SuperAdmin,
	},
	Name:        "超级管理员",
	Keyword:     "superAdmin",
	Status:      proto.Bool(true),
	Description: proto.String("超级管理员  "),
	DataScope:   consts.DataScopeAll,
}
