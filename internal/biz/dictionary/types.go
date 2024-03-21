package dictionary

type PageDictionaryRequest struct {
	Page     uint32  `json:"page"`
	PageSize uint32  `json:"page_size"`
	Keyword  *string `json:"keyword"`
	Name     *string `json:"name"`
}

type PageDictionaryValueRequest struct {
	Page         uint32  `json:"page"`
	PageSize     uint32  `json:"page_size"`
	DictionaryId uint32  `json:"dictionary_id"`
	Label        *string `json:"label"`
	Value        *string `json:"value"`
}
