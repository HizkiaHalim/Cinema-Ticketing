package models

import (
	"time"

	"gorm.io/gorm"
)

type Studio struct {
	gorm.Model
	ID         uint      `gorm:"primaryKey" json:"id"`
	Studio_num string    `gorm:"not null" json:"studio_num"`
	Capacity   int       `gorm:"not null" json:"capacity"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
