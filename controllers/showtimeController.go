package controllers

import (
	"net/http"

	"github.com/HizkiaHalim/Cinema-Ticketing/services"
	"github.com/gin-gonic/gin"
)

func AddShowtime(c *gin.Context) {
	var input services.AddShowtimeRequest

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	// Check if showtime already exists
	if services.CheckShowtimeExists(input.MovieId, input.StudioId, input.Time) {
		c.JSON(http.StatusConflict, gin.H{"error": "There are overlapping showtime"})
		return
	}

	// Register the showtime
	result := services.RegisterShowtime(input)

	if result.StatusCode != 200 {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.ErrString})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Showtime registered successfully",
		"Showtime": gin.H{
			"movie":  result.Title,
			"studio": result.Description,
			"time":   input.Time,
		},
	})
}
