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
