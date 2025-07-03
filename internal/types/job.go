package types

type ListJobRequest struct {
	Keyword *string  `json:"keyword"`
	Name    *string  `json:"name"`
	Ids     []uint32 `json:"ids"`
	RootId  *uint32  `json:"rootId"`
}
