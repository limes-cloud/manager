package entity

import "github.com/limes-cloud/kratosx/model"

type NoticeClassify struct {
	Name   string `json:"name" gorm:"column:name"`
	Logo   string `json:"logo" gorm:"column:logo"`
	Weight uint32 `json:"weight" gorm:"column:weight"`
	model.BaseTenantModel
}

type Notice struct {
	model.BaseTenantModel
	AppId       uint32          `gorm:"column:app_id" json:"appId"`
	ClassifyId  uint32          `gorm:"column:classify_id" json:"classifyId"`
	Title       string          `gorm:"column:title" json:"title"`             // 通知标题
	Description string          `gorm:"column:description" json:"description"` // 通知简介
	Unit        string          `gorm:"column:unit" json:"unit"`               // 发布单位
	Content     string          `gorm:"column:content" json:"content"`         // 通知内容
	IsTop       *bool           `gorm:"column:is_top" json:"isTop"`            // 是否置顶
	Status      *bool           `gorm:"column:status" json:"status"`           // 通知状态
	App         *App            `gorm:"foreignKey:app_id;references:id" json:"app"`
	Classify    *NoticeClassify `gorm:"foreignKey:classify_id;references:id" json:"classify"`
}

type NoticeUser struct {
	NoticeId  uint32 `gorm:"column:notice_id" json:"noticeId"`
	UserId    uint32 `gorm:"column:user_id" json:"userId"`
	CreatedAt int64  `gorm:"column:created_at" json:"createdAt" `
}
