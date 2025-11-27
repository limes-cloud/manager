package entity

import "github.com/limes-cloud/kratosx/model"

type UserDept struct {
	UserId uint32 `json:"userId" gorm:"column:user_id"`
	DeptId uint32 `json:"deptId" gorm:"column:dept_id;hook:dept[dept]"`
	JobId  uint32 `json:"jobId" gorm:"column:job_id"`
	Main   *bool  `json:"main"  gorm:"column:main"`
	Dept   *Dept  `json:"dept" gorm:"foreignKey:dept_id;references:id"`
	Job    *Job   `json:"job" gorm:"foreignKey:job_id;references:id"`
	model.BaseModel
}
