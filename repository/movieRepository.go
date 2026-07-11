package repository

import (
	"time"

	"github.com/HizkiaHalim/Cinema-Ticketing/initializers"
	"github.com/HizkiaHalim/Cinema-Ticketing/models"
)

func GetMoviesByDate(searchDate time.Time) ([]models.Movie, error) {
	var movies []models.Movie

	result := initializers.DB.Where("start_date >= ?", searchDate).Where("end_date <= ?", searchDate).Find(&movies)

	if result.Error != nil {
		return nil, result.Error
	}

	return movies, nil
}
