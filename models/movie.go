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

// make a seat map struct
type Showtime struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	MovieID     uint      `gorm:"not null" json:"movie_id"`
	ShowDate    time.Time `gorm:"not null" json:"show_date"`
	ShowTime    string    `gorm:"not null" json:"show_time"`
	TotalSeats  int       `gorm:"not null" json:"total_seats"`
	BookedSeats int       `gorm:"not null" json:"booked_seats"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
