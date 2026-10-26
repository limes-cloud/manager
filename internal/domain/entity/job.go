package entity

import "github.com/limes-cloud/kratosx/model"

type Job struct {
	App         string  `json:"app" gorm:"column:app"`
	Keyword     string  `json:"keyword" gorm:"column:keyword"`
	Name        string  `json:"name" gorm:"column:name"`
	Description *string `json:"description" gorm:"column:description"`
	Status      *bool   `json:"status" gorm:"column:status"`
	Weight      uint32  `json:"weight" gorm:"column:weight"`
	model.BaseModel
}
