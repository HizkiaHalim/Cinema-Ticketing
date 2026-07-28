package models

import (
	"time"

	"gorm.io/gorm"
)

type Movie struct {
	gorm.Model
	ID          uint       `gorm:"primaryKey" json:"id"`
	Title       string     `gorm:"not null" json:"title"`
	Description string     `json:"description"`
	StartDate   time.Time  `json:"start_date"`
	EndDate     time.Time  `json:"end_date"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	Showtimes   []Showtime `gorm:"foreignKey:MovieID" json:"showtimes,omitempty"`
}

type Showtime struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	MovieID   uint      `gorm:"not null" json:"movie_id"`
	StudioID  uint      `gorm:"not null" json:"studio_id"`
	ShowDate  time.Time `gorm:"not null" json:"show_date"`
	ShowTime  string    `gorm:"not null" json:"show_time"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Studio    Studio    `gorm:"foreignKey:StudioID" json:"studio,omitempty"` // Add this line
}
