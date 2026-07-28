package services

import "github.com/HizkiaHalim/Cinema-Ticketing/repository"

type AddShowtimeRequest struct {
	MovieId  uint   `json:"movie_id" binding:"required"`
	StudioId uint   `json:"studio_id" binding:"required"`
	Time     string `json:"showtime" binding:"required"`
}

func CheckShowtimeExists(movie_id uint, showtime_id uint, time string) bool {
	showtimeExists, err := repository.GetExistingShowtime(movie_id, showtime_id, time)

	if showtimeExists != nil && err == nil {
		return true
	}

	return false
}
