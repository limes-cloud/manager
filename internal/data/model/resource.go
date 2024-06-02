package model

type Resource struct {
	Keyword      string `json:"keyword" gorm:"column:keyword"`
	DepartmentId uint32 `json:"departmentIds" gorm:"column:department_id"`
	ResourceId   uint32 `json:"resourceId" gorm:"column:resource_id"`
}
