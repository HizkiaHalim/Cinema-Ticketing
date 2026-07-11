package models

import (
	"time"

	"gorm.io/gorm"
)

type Movie struct {
	gorm.Model
	ID          uint      `gorm:"primaryKey" json:"id"`
	Title       string    `gorm:"not null" json:"title"`
	Description string    `json:"description"`
	StartDate   time.Time `json:"start_date"`
	EndDate     time.Time `json:"end_date"`
}
