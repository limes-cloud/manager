package entity

import (
	"github.com/limes-cloud/kratosx/model"
)

type Feedback struct {
	AppId           uint32            `json:"appId" gorm:"column:app_id"`
	UserId          uint32            `json:"userId" gorm:"column:user_id"`
	ClassifyId      uint32            `json:"classifyId" gorm:"column:classify_id"`
	Title           string            `json:"title" gorm:"column:title"`
	Content         string            `json:"content" gorm:"column:content"`
	Status          string            `json:"status" gorm:"column:status"`
	Images          *string           `json:"images" gorm:"column:images"`
	Contact         *string           `json:"contact" gorm:"column:contact"`
	Device          string            `json:"device" gorm:"column:device"`
	Platform        string            `json:"platform" gorm:"column:platform"`
	Version         string            `json:"version" gorm:"column:version"`
	Md5             string            `json:"md5" gorm:"column:md5"`
	ProcessedBy     *uint32           `json:"processedBy" gorm:"column:processed_by"`
	ProcessedResult *string           `json:"processedResult" gorm:"column:processed_result"`
	App             *App              `json:"app"`
	User            *User             `json:"user"`
	Classify        *FeedbackClassify `json:"classify" gorm:"foreignKey:classify_id;references:id"`
	ImageUrls       []string          `json:"imageUrls" gorm:"-"`
	model.BaseTenantModel
}

type FeedbackClassify struct {
	Name string `json:"name" gorm:"column:name"`
	model.BaseTenantModel
}
