package resource

type Resource struct {
	Keyword      string `json:"keyword"`
	DepartmentId uint32 `json:"departmentId"`
	ResourceId   uint32 `json:"resourceId"`
}
