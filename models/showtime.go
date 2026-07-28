package models

import (
	"time"
)

type Showtime struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	MovieID   uint      `gorm:"not null" json:"movie_id"`
	StudioID  uint      `gorm:"not null" json:"studio_id"`
	ShowDate  time.Time `gorm:"not null" json:"show_date"`
	ShowTime  string    `gorm:"not null" json:"show_time"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Movie     Movie     `gorm:"foreignKey:MovieID" json:"movie,omitempty"`
	Studio    Studio    `gorm:"foreignKey:StudioID" json:"studio,omitempty"`
}
