package repository

import (
	"github.com/HizkiaHalim/Cinema-Ticketing/initializers"
	"github.com/HizkiaHalim/Cinema-Ticketing/models"
)

func GetExistingShowtime(movie_id uint, studio_id uint, time string) (*models.Showtime, error) {
	var showtime models.Showtime

	result := initializers.DB.Where("UPPER(movieId) = UPPER(?)", movie_id).Where("UPPER(studioId) = UPPER(?)", studio_id).Where("time = time", time).First(showtime)

	if result.Error != nil {
		return nil, result.Error
	}

	return &showtime, nil
}
