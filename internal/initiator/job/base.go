package job

import (
	"github.com/limes-cloud/kratosx/types"
	"google.golang.org/protobuf/proto"

	"github.com/limes-cloud/manager/consts"
	"github.com/limes-cloud/manager/internal/model"
)

var job = &model.Job{
	BaseModel: types.BaseModel{
		ID: consts.SuperAdmin,
	},
	Name:        "董事长",
	Keyword:     "chairman",
	Description: "董事长",
	Weight:      proto.Uint32(0),
}
