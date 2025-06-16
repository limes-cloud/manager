package types

type GetAppRequest struct {
	Id      *uint32 `json:"id"`
	Keyword *string `json:"keyword"`
}

type ListAppRequest struct {
	Page     uint32   `json:"page"`
	PageSize uint32   `json:"pageSize"`
	Order    *string  `json:"order"`
	OrderBy  *string  `json:"orderBy"`
	Keyword  *string  `json:"keyword"`
	Name     *string  `json:"name"`
	Status   *bool    `json:"status"`
	Ids      []uint32 `json:"ids"`
}

type UpdateAppStatusRequest struct {
	Id          uint32  `json:"id"`
	Status      bool    `json:"status"`
	DisableDesc *string `json:"disableDesc"`
}
