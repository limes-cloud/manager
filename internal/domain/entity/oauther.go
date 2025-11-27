package entity

import (
	"encoding/json"

	"github.com/limes-cloud/kratosx/model"
)

type OAuther struct {
	Logo        string `json:"logo" gorm:"column:logo"`               // 渠道logo
	Keyword     string `json:"keyword" gorm:"column:keyword"`         // 渠道标识
	Type        string `json:"type"  gorm:"column:type"`              // 渠道类型
	Name        string `json:"name" gorm:"column:name"`               // 渠道名称
	Status      *bool  `json:"status" gorm:"column:status"`           // 启用状态
	Ak          string `json:"ak" gorm:"column:ak"`                   // 渠道AK
	Sk          string `json:"sk" gorm:"column:sk"`                   // 渠道SK
	Setting     string `json:"extra" gorm:"column:setting"`           // 渠道设置
	Description string `json:"description" gorm:"column:description"` // 描述信息
	model.BaseTenantModel
}

func (c *OAuther) TableName() string {
	return "oauther"
}

type OAutherSetting struct {
	Callback string `json:"callback"`
	Email    struct {
		Host     string
		Port     int
		Subject  string
		From     string
		Template string
	}
}

func (c *OAuther) GetSetting() *OAutherSetting {
	var st OAutherSetting
	_ = json.Unmarshal([]byte(c.Setting), &st)
	return &st
}

func (extra OAutherSetting) ToString() string {
	b, _ := json.Marshal(extra)
	return string(b)
}
