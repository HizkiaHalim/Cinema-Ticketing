package controllers

import (
	"net/http"

	"github.com/HizkiaHalim/Cinema-Ticketing/repository"
	"github.com/HizkiaHalim/Cinema-Ticketing/services"
	"github.com/HizkiaHalim/Cinema-Ticketing/utils"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// server HealthCheck
func HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  "healthy",
		"message": "Cinema Ticketing API is running smoothly",
	})
}

// handle user registration
func SignUp(c *gin.Context) {
	var input services.SignUpRequest

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Check if user already exists
	if services.CheckUserExists(input.Email) {
		c.JSON(http.StatusConflict, gin.H{"error": "User with this email already exists"})
		return
	}

	// Register the user
	if err := services.RegisterUser(input); err.StatusCode != 200 {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register user, " + err.ErrString})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "User created successfully",
		"user": gin.H{
			"name":  input.Name,
			"email": input.Email,
		},
	})
}

// handle user authentication
func Login(c *gin.Context) {
	var input services.LoginRequest

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Find user by email
	user, err := repository.GetUserByEmail(input.Email)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	// Check password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password))

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	// Generate JWT token
	token, err := utils.GenerateJWT(user.ID, user.Email)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Login successful",
		"user": gin.H{
			"id":    user.ID,
			"name":  user.Name,
			"email": user.Email,
		},
		"token": token,
	})
}

func RegisterAdmin(c *gin.Context) {
	var input services.SignUpRequest

	// Check if current user is admin

	// Input validation
	if err := c.ShouldBind(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	// Check if user already exists
	if services.CheckUserExists(input.Email) {
		c.JSON(http.StatusConflict, gin.H{"error": "Admin with this email already exists"})
		return
	}

	// Register admin
	if err := services.RegisterAdmin(input); err.StatusCode != 200 {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register admin, " + err.ErrString})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Admin created successfully",
		"user": gin.H{
			"name":  input.Name,
			"email": input.Email,
		},
	})
}
