package types

type ListMenuRequest struct {
	AppId      *uint32  `json:"appId"`
	InIds      []uint32 `json:"inIds"`
	InTypes    []string `json:"inTypes"`
	NotInTypes []string `json:"notInTypes"`

	// 中间转换参数
	OnlyMenu bool `json:"onlyMenu"`
}

type CacheMenu struct {
	Id    uint32 `json:"id"`
	Type  string `json:"type"`
	AppId uint32 `json:"appId"`
	Name  string `json:"name"`
}
