package services

import (
	"net/http"
	"time"

	"github.com/HizkiaHalim/Cinema-Ticketing/models"
	"github.com/HizkiaHalim/Cinema-Ticketing/repository"
)

type MovieListRequest struct {
	SearchDate string `json:"searchDate" binding:"required"`
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
