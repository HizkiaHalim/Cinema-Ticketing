package services

import (
	"fmt"
	"net/http"
	"time"

	"github.com/HizkiaHalim/Cinema-Ticketing/dto"
	"github.com/HizkiaHalim/Cinema-Ticketing/models"
	"github.com/HizkiaHalim/Cinema-Ticketing/repository"
)

type MovieListRequest struct {
	SearchDate string `json:"searchDate" binding:"required"`
}

type MovieDetailRequest struct {
	Id   uint   `json:"movie_id" binding:"required"`
	Date string `json:"date" binding:"required"`
}

type MovieRegisterRequest struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description"`
	StartDate   string `json:"startDate" binding:"required"`
	EndDate     string `json:"endDate" binding:"required"`
}

type MovieUpdateRequest struct {
	Id          uint   `json:"id" binding:"required"`
	Title       string `json:"title" binding:"required"`
	Description string `json:"description"`
	StartDate   string `json:"startDate" binding:"required"`
	EndDate     string `json:"endDate" binding:"required"`
}

type MovieDeleteRequest struct {
	Id uint `json:"id" binding:"required"`
}

func CheckMovieExists(title string) bool {

	movieExists, err := repository.GetMovieByTitle(title)

	if movieExists != nil && err == nil {
		return true
	}

	return false
}

func CheckMovieExistsOnUpdate(title string, id uint) bool {
	movieExists, err := repository.GetMovieByTitleId(title, id)

	if movieExists != nil && err == nil {
		return true
	}

	return false
}

func CheckMovieExistOnDelete(id uint) bool {
	movieExists, err := repository.GetMovieById(id)

	if !movieExists && err == nil {
		return false
	}

	return true
}

func RegisterMovie(input MovieRegisterRequest) ErrorResponse {
	// Parse the start and end dates
	startDate, err := time.Parse("2006-01-02", input.StartDate)
	if err != nil {
		return ErrorResponse{ErrString: "Failed to parse start date! Please use format 'YYYY-MM-DD'.", StatusCode: http.StatusInternalServerError}
	}

	endDate, err := time.Parse("2006-01-02", input.EndDate)
	if err != nil {
		return ErrorResponse{ErrString: "Failed to parse end date! Please use format 'YYYY-MM-DD'.", StatusCode: http.StatusInternalServerError}
	}

	// Create the movie
	movie := &models.Movie{
		Title:       input.Title,
		Description: input.Description,
		StartDate:   startDate,
		EndDate:     endDate,
	}

	if err := repository.CreateMovie(movie).Error; err != nil {
		return ErrorResponse{ErrString: "Failed to register new movie.", StatusCode: http.StatusInternalServerError}
	}

	return ErrorResponse{ErrString: "", StatusCode: http.StatusCreated}
}

func UpdateMovie(input MovieUpdateRequest) ErrorResponse {
	// Parse the start and end dates
	startDate, err := time.Parse("2006-01-02", input.StartDate)
	if err != nil {
		return ErrorResponse{ErrString: "Failed to parse start date! Please use format 'YYYY-MM-DD'.", StatusCode: http.StatusInternalServerError}
	}

	endDate, err := time.Parse("2006-01-02", input.EndDate)
	if err != nil {
		return ErrorResponse{ErrString: "Failed to parse end date! Please use format 'YYYY-MM-DD'.", StatusCode: http.StatusInternalServerError}
	}

	// Update the movie
	movie := &models.Movie{
		ID:          input.Id,
		Title:       input.Title,
		Description: input.Description,
		StartDate:   startDate,
		EndDate:     endDate,
	}

	if err := repository.UpdateMovie(movie).Error; err != nil {
		return ErrorResponse{ErrString: "Failed to update movie.", StatusCode: http.StatusInternalServerError}
	}

	return ErrorResponse{ErrString: "", StatusCode: http.StatusCreated}
}

func GetMovieForDate(id uint, date string) (*dto.MovieDetailResponse, error) {
	// Parse the search date
	parsedDate, err := time.Parse("2006-01-02", date)
	if err != nil {
		return nil, fmt.Errorf("invalid date format")
	}

	// Get movie with preloaded showtimes for specific date
	movie, err := repository.GetMovieForDate(id, parsedDate)

	if err != nil {
		return nil, err
	}

	// Prepare response with booking status
	response := &dto.MovieDetailResponse{
		ID:          movie.ID,
		Title:       movie.Title,
		Description: movie.Description,
		Showtimes:   make([]dto.ShowtimeResponse, len(movie.Showtimes)),
	}

	// Process showtimes with booking status
	for i, showtime := range movie.Showtimes {
		full := showtime.Studio.Capacity == 0

		response.Showtimes[i] = dto.ShowtimeResponse{
			ID:        showtime.Studio.ID,
			StudioNum: showtime.Studio.Studio_num,
			IsFull:    full,
		}
	}

	return response, nil
}
