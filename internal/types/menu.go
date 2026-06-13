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

type RegisterMenuRequest struct {
	AppKeyword string              `json:"appKeyword"`
	Mode       string              `json:"mode"`
	Menus      []*RegisterMenuNode `json:"menus"`
}

type RegisterMenuNode struct {
	Key        *string             `json:"key"`
	Title      string              `json:"title"`
	Type       string              `json:"type"`
	Icon       *string             `json:"icon"`
	Api        *string             `json:"api"`
	Method     *string             `json:"method"`
	Path       *string             `json:"path"`
	Permission *string             `json:"permission"`
	Component  *string             `json:"component"`
	Url        *string             `json:"url"`
	Redirect   *string             `json:"redirect"`
	Weight     *int32              `json:"weight"`
	IsIframe   *bool               `json:"isIframe"`
	IsHidden   *bool               `json:"isHidden"`
	IsCache    *bool               `json:"isCache"`
	IsHome     *bool               `json:"isHome"`
	IsAffix    *bool               `json:"isAffix"`
	Children   []*RegisterMenuNode `json:"children"`
}

type RegisterMenuReply struct {
	Created uint32 `json:"created"`
	Updated uint32 `json:"updated"`
	Deleted uint32 `json:"deleted"`
}
