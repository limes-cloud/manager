package resource

type UpdateResourceScopesRequest struct {
	Keyword    string   `json:"keyword"`
	Scopes     []uint32 `json:"scopes"`
	ResourceId uint32   `json:"resourceId"`
}
