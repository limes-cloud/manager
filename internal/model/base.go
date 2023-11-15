package model

import "gorm.io/gorm"

type Scopes func(db *gorm.DB) *gorm.DB

type PageOptions struct {
	Page     int32
	PageSize int32
	Scopes   Scopes
}

type CreateModel struct {
	ID        uint32 `json:"id"`
	CreatedAt uint32 `json:"created_at,omitempty"`
}

type BaseModel struct {
	ID        uint32 `json:"id"`
	CreatedAt uint32 `json:"created_at,omitempty"`
	UpdatedAt uint32 `json:"updated_at,omitempty"`
}

type DeleteModel struct {
	ID        uint32         `json:"id"`
	CreatedAt uint32         `json:"created_at,omitempty"`
	UpdatedAt uint32         `json:"updated_at,omitempty"`
	DeletedAt gorm.DeletedAt `json:"-"`
}
