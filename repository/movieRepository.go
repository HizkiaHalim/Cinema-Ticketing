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

func GetMovieByTitle(title string) (*models.Movie, error) {
	var movie models.Movie

	result := initializers.DB.Where("UPPER(title) = UPPER(?)", title).First(&movie)

	if result.Error != nil {
		return nil, result.Error
	}

	return &movie, nil
}

func GetMovieByTitleId(title string, id uint) (*models.Movie, error) {
	var movie models.Movie

	result := initializers.DB.Where("UPPER(title) = UPPER(?)", title).Where("id != ?", id).First(&movie)

	if result.Error != nil {
		return nil, result.Error
	}

	return &movie, nil
}

func GetMovieById(id uint) (bool, error) {
	var count int64
	result := initializers.DB.Model(&models.Movie{}).Where("id = ?", id).Count(&count)

	if result.Error != nil {
		return false, result.Error
	}

	return count > 0, nil
}

func CreateMovie(movie *models.Movie) error {
	result := initializers.DB.Create(&movie)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func UpdateMovie(movie *models.Movie) error {
	result := initializers.DB.Save(&movie)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func DeleteMovie(id uint) error {
	result := initializers.DB.Delete(&models.Movie{}, id)

	if result.Error != nil {
		return result.Error
	}

	return nil
}
