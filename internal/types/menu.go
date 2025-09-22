package types

type ListMenuRequest struct {
	AppId      *uint32  `json:"appId"`
	InIds      []uint32 `json:"inIds"`
	InTypes    []string `json:"inTypes"`
	NotInTypes []string `json:"notInTypes"`
	FilterRoot bool     `json:"filterRoot"`
}
