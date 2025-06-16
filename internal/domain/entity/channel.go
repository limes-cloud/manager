package entity

import (
	"encoding/json"

	"github.com/limes-cloud/kratosx/types"
)

type Channel struct {
	Logo        string `json:"logo" gorm:"column:logo"`
	Keyword     string `json:"keyword" gorm:"column:keyword"`
	Name        string `json:"name" gorm:"column:name"`
	Status      *bool  `json:"status" gorm:"column:status"`
	Admin       *bool  `json:"admin" gorm:"column:admin"`
	Ak          string `json:"ak" gorm:"column:ak"`
	Sk          string `json:"sk" gorm:"column:sk"`
	Extra       string `json:"extra" gorm:"column:extra"`
	Description string `json:"description" gorm:"column:description"`
	LogoUrl     string `json:"logoUrl" gorm:"-"`
	types.BaseModel
}

type ChannelExtra struct {
	CallBack string
}

func (c Channel) GetExtra() *ChannelExtra {
	var extra ChannelExtra
	_ = json.Unmarshal([]byte(c.Extra), &extra)
	return &extra
}

func (extra ChannelExtra) ToString() string {
	b, _ := json.Marshal(extra)
	return string(b)
}
