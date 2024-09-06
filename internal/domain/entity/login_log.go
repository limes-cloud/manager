package entity

import "github.com/limes-cloud/kratosx/types"

type LoginLog struct {
	types.CreateModel
	Username    string `json:"username" gorm:"column:username"`
	Type        string `json:"type" gorm:"column:type"`
	IP          string `json:"ip" gorm:"column:ip"`
	Address     string `json:"address" gorm:"column:address"`
	Browser     string `json:"browser" gorm:"column:browser"`
	Device      string `json:"device" gorm:"column:device"`
	Code        int    `json:"code" gorm:"column:code"`
	Description string `json:"description" gorm:"column:description"`
}
