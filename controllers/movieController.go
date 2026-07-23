package controllers

import (
	"net/http"
	"time"

	"github.com/HizkiaHalim/Cinema-Ticketing/repository"
	"github.com/HizkiaHalim/Cinema-Ticketing/services"
	"github.com/gin-gonic/gin"
)

func GetMovieList(c *gin.Context) {
	var input services.MovieListRequest

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Parse the search date
	searchDate, err := time.Parse("2006-01-02", input.SearchDate)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid date format. Use YYYY-MM-DD."})
		return
	}

	movies, err := repository.GetMoviesByDate(searchDate)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch movies"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"movies": movies})
}

func GetMovieDetailOnDate(c *gin.Context) {
	var input services.MovieDetailRequest

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Parse the search date
	searchDate, err := time.Parse("2006-01-02", input.Date)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid date format. Use YYYY-MM-DD."})
		return
	}

	// movies, err := repository.GetMovieForDate(input.Id, searchDate)
	// if err != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch movies"})
	// 	return
	// }

	c.JSON(http.StatusOK, gin.H{"movies": searchDate})
}

func RegisterMovie(c *gin.Context) {
	var input services.MovieRegisterRequest

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Check if movie already exists
	if services.CheckMovieExists(input.Title) {
		c.JSON(http.StatusConflict, gin.H{"error": "Movie with this title already exists"})
		return
	}

	// Register the movie
	if err := services.RegisterMovie(input); err.StatusCode != 200 {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.ErrString})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Movie registered successfully",
		"movie": gin.H{
			"title":       input.Title,
			"description": input.Description,
			"startDate":   input.StartDate,
			"endDate":     input.EndDate,
		},
	})
}

func UpdateMovie(c *gin.Context) {
	var input services.MovieUpdateRequest

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	// Check if movie updated to already existing one
	if services.CheckMovieExistsOnUpdate(input.Title, input.Id) {
		c.JSON(http.StatusConflict, gin.H{"error": "Movie with this title already exists"})
		return
	}

	// Register the movie
	if err := services.UpdateMovie(input); err.StatusCode != 200 {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.ErrString})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Movie updated successfully",
		"movie": gin.H{
			"title":       input.Title,
			"description": input.Description,
			"startDate":   input.StartDate,
			"endDate":     input.EndDate,
		},
	})
}

func DeleteMovie(c *gin.Context) {
	var input services.MovieDeleteRequest

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input: " + err.Error()})
		return
	}

	// Check if movie exists using service (following same pattern as UpdateMovie)
	if services.CheckMovieExistOnDelete(input.Id) {
		c.JSON(http.StatusNotFound, gin.H{"error": "Movie not found"})
		return
	}

	// Delete the movie using repository
	if err := repository.DeleteMovie(input.Id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete movie"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Movie deleted successfully",
		"movie": gin.H{
			"id": input.Id,
		},
	})
}
