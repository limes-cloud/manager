package entity

import (
	"encoding/json"
	"unsafe"

	"github.com/limes-cloud/kratosx/model"
)

type App struct {
	Type        string  `json:"type" gorm:"column:type"`
	Logo        string  `json:"logo" gorm:"column:logo"`
	Keyword     string  `json:"keyword" gorm:"column:keyword"`
	Favicon     string  `json:"favicon" gorm:"column:favicon"`
	Secret      string  `json:"secret" gorm:"column:secret"`
	Name        string  `json:"name" gorm:"column:name"`
	ShowName    string  `json:"showName" gorm:"column:show_name"`
	Status      *bool   `json:"status" gorm:"column:status"`
	Reason      *string `json:"reason" gorm:"column:reason"`
	Private     *bool   `json:"private" gorm:"column:private"`
	Description *string `json:"description" gorm:"column:description"`
	Comment     *string `json:"comment" gorm:"column:comment"`
	Setting     string  `json:"setting" gorm:"column:setting"`
	model.BaseModel
}

type AppSettingJWT struct {
	Secret         string `json:"secret"`         // 加密密钥
	Expire         int64  `json:"expire"`         // 过期时间
	Renewal        int64  `json:"renewal"`        // 续期时间
	UniqueDevice   bool   `json:"uniqueDevice"`   // 唯一设备
	UniquePlatform bool   `json:"uniquePlatform"` // 唯一平台
}

type AppSettingTenant struct {
	Mode string `json:"mode"`
}

type AppSetting struct {
	JWT    AppSettingJWT    `json:"jwt"`
	Tenant AppSettingTenant `json:"tenant"`
}

// TableName 获取表名称
func (app *App) TableName() string {
	return "app"
}

func (app *App) String() string {
	b, _ := json.Marshal(app)
	return *(*string)(unsafe.Pointer(&b))
}

func (app *App) GetSetting() *AppSetting {
	var setting AppSetting
	_ = json.Unmarshal([]byte(app.Setting), &setting)
	return &setting
}
