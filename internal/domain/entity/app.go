package entity

import (
	"encoding/json"
	"fmt"
	"unsafe"

	"github.com/limes-cloud/kratosx/model"
	"github.com/limes-cloud/kratosx/pkg/value"
)

const (
	AppTypeBase = "base" // 基础设施：提供菜单与数据实体，不承载组织数据
	AppTypeApp  = "app"  // 标准应用：承载角色/部门/职位等组织数据
)

// AppTypes 全部合法的应用类型
var AppTypes = []string{AppTypeBase, AppTypeApp}

type App struct {
	Type        string  `json:"type" gorm:"column:type"`
	Logo        string  `json:"logo" gorm:"column:logo"`
	Keyword     string  `json:"keyword" gorm:"column:keyword"`
	Favicon     string  `json:"favicon" gorm:"column:favicon"`
	Name        string  `json:"name" gorm:"column:name"`
	ShowName    string  `json:"showName" gorm:"column:show_name"`
	Status      *bool   `json:"status" gorm:"column:status"`
	Reason      *string `json:"reason" gorm:"column:reason"`
	Private     *bool   `json:"private" gorm:"column:private"`
	Loginable   *bool   `json:"loginable" gorm:"column:loginable"` // 是否可作为登录入口；NULL/false = 不可登录
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

type AppSetting struct {
	JWT     AppSettingJWT `json:"jwt"`
	Modules []string      `json:"modules"` // 订阅的基础设施应用 keyword 列表，取值必须是 type=base 的应用
}

// TableName 获取表名称
func (app *App) TableName() string {
	return "app"
}

func (app *App) String() string {
	b, _ := json.Marshal(app)
	return *(*string)(unsafe.Pointer(&b))
}

// GetSetting 获取应用配置，解析失败时返回零值配置
// 注意：配置参与权限判断时应改用 ParseSetting，以便记录解析错误
func (app *App) GetSetting() *AppSetting {
	setting, _ := app.ParseSetting()
	return setting
}

// ParseSetting 获取应用配置并返回解析错误，供调用方记录日志
// 解析失败会导致 Modules 为空（表现为未订阅任何模块），静默忽略会造成难以定位的鉴权拒绝
func (app *App) ParseSetting() (*AppSetting, error) {
	var setting AppSetting
	if app.Setting == "" {
		return &setting, nil
	}
	if err := json.Unmarshal([]byte(app.Setting), &setting); err != nil {
		return &setting, fmt.Errorf("parse app %s setting: %w", app.Keyword, err)
	}
	return &setting, nil
}

// IsLoginable 判断应用是否可作为登录入口
func (app *App) IsLoginable() bool {
	return value.Value(app.Loginable)
}

// IsBase 判断应用是否为基础设施
func (app *App) IsBase() bool {
	return app.Type == AppTypeBase
}

// ScopeModules 获取应用的模块可见范围：自身 + 已订阅的基础设施应用
// 用于菜单下发与接口鉴权时判断某个菜单所属模块是否在当前应用的可触达范围内
func (app *App) ScopeModules() []string {
	setting, _ := app.ParseSetting()
	scopes := make([]string, 0, len(setting.Modules)+1)
	scopes = append(scopes, app.Keyword)
	for _, item := range setting.Modules {
		if item != "" && item != app.Keyword {
			scopes = append(scopes, item)
		}
	}
	return scopes
}

// HasModule 判断指定模块是否在当前应用的可触达范围内
func (app *App) HasModule(module string) bool {
	return value.InList(app.ScopeModules(), module)
}
