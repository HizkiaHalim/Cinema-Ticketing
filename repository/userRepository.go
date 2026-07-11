package repository

import (
	"github.com/HizkiaHalim/Cinema-Ticketing/initializers"
	"github.com/HizkiaHalim/Cinema-Ticketing/models"
)

func GetUserByEmail(email string) (*models.User, error) {
	var user models.User

	result := initializers.DB.Where("email = ?", email).First(&user)

	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}

func CreateUser(user *models.User) error {
	result := initializers.DB.Create(&user)

	if result.Error != nil {
		return result.Error
	}

	return nil
}
