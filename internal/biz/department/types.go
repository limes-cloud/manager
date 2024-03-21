package department

type AllDepartmentRequest struct {
	Ids []uint32
}

type AllDepartmentObjectValueRequest struct {
	ObjectId     uint32 `json:"object_id"`
	DepartmentId uint32 `json:"department_id"`
}

type DeleteDepartmentObjectRequest struct {
	ObjectId uint32 `json:"object_id"`
	Value    string `json:"value"`
}
