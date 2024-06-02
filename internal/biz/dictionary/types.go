package dictionary

type ListDictionaryRequest struct {
	Page     uint32  `json:"page"`
	PageSize uint32  `json:"pageSize"`
	Order    *string `json:"order"`
	OrderBy  *string `json:"orderBy"`
	Keyword  *string `json:"keyword"`
	Name     *string `json:"name"`
}

type ListDictionaryValueRequest struct {
	Page         uint32  `json:"page"`
	PageSize     uint32  `json:"pageSize"`
	Order        *string `json:"order"`
	OrderBy      *string `json:"orderBy"`
	DictionaryId *uint32 `json:"dictionaryId"`
	Label        *string `json:"label"`
	Value        *string `json:"value"`
	Status       *bool   `json:"status"`
}

type GetDictionaryRequest struct {
	Id      *uint32 `json:"id"`
	Keyword *string `json:"keyword"`
}
