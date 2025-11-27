package entity

import (
	"encoding/json"
	"unsafe"

	"github.com/limes-cloud/kratosx/model"
)

// Tenant 租户实体
type Tenant struct {
	model.DeleteModel
	Keyword     string  `gorm:"column:keyword" json:"keyword"`         // 租户标识
	Logo        string  `gorm:"column:logo"  json:"logo"`              // 租户头像
	Name        string  `gorm:"column:name" json:"name"`               // 租户名称
	Status      *bool   `gorm:"column:status" json:"status"`           // 租户状态
	Description string  `gorm:"column:description" json:"description"` // 租户简介
	Weight      *uint32 `gorm:"column:weight" json:"weight"`           // 权重
	Setting     string  `gorm:"column:setting" json:"setting"`         // 租户设置
}

func (t *Tenant) TableName() string {
	return "tenant"
}

func (t *Tenant) GetSetting() *TenantSetting {
	var ts TenantSetting
	_ = json.Unmarshal([]byte(t.Setting), &ts)
	return &ts
}

type TenantSetting struct {
	// QuickBind           bool   `json:"quickBind"`           // 快速绑定开关
	// QuickLogin          bool   `json:"quickLogin"`          // 快速登录开关
	DefaultUserAvatar   string `json:"defaultUserAvatar"`   // 默认用户头像
	DefaultUserPassword string `json:"defaultUserPassword"` // 默认用户密码
	DefaultUserNickname string `json:"defaultUserNickname"` // 默认用户昵称
}

func (ts *TenantSetting) String() string {
	b, _ := json.Marshal(ts)
	return *(*string)(unsafe.Pointer(&b))
}
