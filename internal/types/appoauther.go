package types

import "github.com/limes-cloud/kratosx/model/page"

type ListAppOAutherRequest struct {
	*page.Search
	AppId    uint32  `json:"appId"`
	Keyword  *string `json:"keyword"`
	Name     *string `json:"name"`
	TenantId *uint32 `json:"tenantId"`
}
