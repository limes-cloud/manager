package entity

import "github.com/limes-cloud/kratosx/model"

type LoginLog struct {
	model.CreateModel
	AppId       uint32 `json:"appId" gorm:"column:app_id"`
	DeptId      uint32 `json:"deptId" gorm:"column:dept_id"`
	UserId      uint32 `json:"userId" gorm:"column:user_id"`
	Type        string `json:"type" gorm:"column:type"`
	IP          string `json:"ip" gorm:"column:ip"`
	Address     string `json:"address" gorm:"column:address"`
	Browser     string `json:"browser" gorm:"column:browser"`
	Device      string `json:"device" gorm:"column:device"`
	Code        int    `json:"code" gorm:"column:code"`
	Description string `json:"description" gorm:"column:description"`
	User        *User  `json:"user" gorm:"foreignKey:user_id;references:id"`
	App         *App   `json:"app" gorm:"foreignKey:app_id;references:id"`
}

type AuthLog struct {
	model.CreateModel
	AppId  uint32 `json:"appId" gorm:"column:app_id"`
	DeptId uint32 `json:"deptId" gorm:"column:dept_id"`
	UserId uint32 `json:"userId" gorm:"column:user_id"`
	MenuId uint32 `json:"menuId" gorm:"column:menu_id"`
	App    *App   `json:"app" gorm:"foreignKey:app_id;references:id"`
	User   *User  `json:"user" gorm:"foreignKey:user_id;references:id"`
	Menu   *Menu  `json:"menu" gorm:"foreignKey:menu_id;references:id"`
}
