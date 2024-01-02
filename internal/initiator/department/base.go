package department

import (
	"github.com/limes-cloud/kratosx/types"

	"github.com/limes-cloud/manager/consts"
	"github.com/limes-cloud/manager/internal/model"
)

var department = &model.Department{
	BaseModel: types.BaseModel{
		ID: consts.SuperAdmin,
	},
	Name:        "贵州青橙科技有限公司",
	Keyword:     "company",
	Description: "开放合作，拥抱未来",
}
