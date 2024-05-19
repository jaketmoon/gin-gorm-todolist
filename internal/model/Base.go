package model

import (
	"gorm.io/gorm"
	"time"
)

type Base struct {
	ID        string         `gorm:"primary_key" json:"id"`
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}
