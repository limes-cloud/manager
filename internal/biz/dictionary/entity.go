package dictionary

type Dictionary struct {
	Id          uint32  `json:"id"`
	Keyword     string  `json:"keyword"`
	Name        string  `json:"name"`
	Description *string `json:"description"`
	CreatedAt   int64   `json:"createdAt"`
	UpdatedAt   int64   `json:"updatedAt"`
}

type DictionaryValue struct {
	Id           uint32  `json:"id"`
	DictionaryId uint32  `json:"dictionaryId"`
	Label        string  `json:"label"`
	Value        string  `json:"value"`
	Status       *bool   `json:"status"`
	Weight       *int32  `json:"weight"`
	Type         *string `json:"type"`
	Extra        *string `json:"extra"`
	Description  *string `json:"description"`
	CreatedAt    int64   `json:"createdAt"`
	UpdatedAt    int64   `json:"updatedAt"`
}
