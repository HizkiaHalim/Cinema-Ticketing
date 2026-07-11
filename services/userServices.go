package services

import (
	"net/http"
	"time"

	"github.com/HizkiaHalim/Cinema-Ticketing/initializers"
	"github.com/HizkiaHalim/Cinema-Ticketing/models"
	"github.com/HizkiaHalim/Cinema-Ticketing/repository"
	"golang.org/x/crypto/bcrypt"
)

type SignUpRequest struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type ErrorResponse struct {
	StatusCode int    `json:"status_code"`
	ErrString  string `json:"error"`
}

func CheckUserExists(email string) bool {
	var existingUser models.User

	userExists := initializers.DB.Where("email = ?", email).First(&existingUser).Error

	if userExists != nil {
		return true
	}
	return false
}

func RegisterUser(input SignUpRequest) ErrorResponse {
	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return ErrorResponse{ErrString: "Failed to hash password", StatusCode: http.StatusInternalServerError}
	}

	// Create new user
	user := models.User{
		Name:      input.Name,
		Email:     input.Email,
		Password:  string(hashedPassword),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if err := repository.CreateUser(&user); err != nil {
		return ErrorResponse{ErrString: "Failed to create user", StatusCode: http.StatusInternalServerError}
	}

	return ErrorResponse{ErrString: "", StatusCode: http.StatusCreated}
}
