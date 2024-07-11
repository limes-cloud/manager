package resource

type UpdateResourceRequest struct {
	UserId        uint32   `json:"userId"`
	Keyword       string   `json:"keyword"`
	ResourceId    uint32   `json:"resourceId"`
	DepartmentIds []uint32 `json:"departmentIds"`
}

type GetResourceRequest struct {
	UserId     uint32 `json:"userId"`
	Keyword    string `json:"keyword"`
	ResourceId uint32 `json:"resourceId"`
}
