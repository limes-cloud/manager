package entity

import "github.com/limes-cloud/kratosx/model"

type LoginLog struct {
	model.CreateTenantUserModel
	Type        string `json:"type" gorm:"column:type"`
	IP          string `json:"ip" gorm:"column:ip"`
	Address     string `json:"address" gorm:"column:address"`
	Browser     string `json:"browser" gorm:"column:browser"`
	Device      string `json:"device" gorm:"column:device"`
	Code        int    `json:"code" gorm:"column:code"`
	Description string `json:"description" gorm:"column:description"`
}

type AuthLog struct {
	model.CreateTenantUserModel
	Type    string `json:"type" gorm:"column:type"`
	IP      string `json:"ip" gorm:"column:ip"`
	Address string `json:"address" gorm:"column:address"`
	Browser string `json:"browser" gorm:"column:browser"`
	Device  string `json:"device" gorm:"column:device"`
	Path    string `json:"path" gorm:"column:path"`
	Method  string `json:"method" gorm:"column:method"`
	Name    string `json:"name" gorm:"column:method"`
}
