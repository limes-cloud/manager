package job

type ListJobRequest struct {
	Page     uint32  `json:"page"`
	PageSize uint32  `json:"pageSize"`
	Order    *string `json:"order"`
	OrderBy  *string `json:"orderBy"`
	Keyword  *string `json:"keyword"`
	Name     *string `json:"name"`
}

type GetJobRequest struct {
	Id      *uint32 `json:"id"`
	Keyword *string `json:"keyword"`
}
