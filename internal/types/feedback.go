package types

type ListFeedbackCategoryRequest struct {
	Page     uint32  `json:"page"`
	PageSize uint32  `json:"pageSize"`
	Order    *string `json:"order"`
	OrderBy  *string `json:"orderBy"`
	Name     *string `json:"name"`
}

type ListFeedbackRequest struct {
	Page       uint32   `json:"page"`
	PageSize   uint32   `json:"pageSize"`
	Order      *string  `json:"order"`
	OrderBy    *string  `json:"orderBy"`
	AppId      *uint32  `json:"appId"`
	CategoryId *uint32  `json:"categoryId"`
	Status     *string  `json:"status"`
	Platform   *string  `json:"platform"`
	AppIds     []uint32 `json:"appIds"`
}
