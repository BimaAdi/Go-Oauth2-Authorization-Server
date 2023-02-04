package repository

import (
	"time"

	"github.com/BimaAdi/Oauth2AuthorizationServer/core"
	"github.com/BimaAdi/Oauth2AuthorizationServer/models"
)

func GetUserById(id string) (models.User, error) {
	user := models.User{}
	if err := models.DBConn.Where("id = ? AND deleted_at IS NULL", id).First(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}

func CreateUser(username string, email string, password string, isActive bool, isSuperuser bool, createdAt time.Time, updatedAt *time.Time) (models.User, error) {
	hashedPassword, err := core.HashPassword(password)
	if err != nil {
		return models.User{}, err
	}

	newUser := models.User{
		Email:       email,
		Username:    username,
		Password:    hashedPassword,
		IsActive:    isActive,
		IsSuperuser: isSuperuser,
		CreatedAt:   createdAt,
		UpdatedAt:   updatedAt,
		DeletedAt:   nil,
	}

	if err := models.DBConn.Create(&newUser).Error; err != nil {
		return newUser, err
	}
	return newUser, nil
}

func UpdateUser(updatedUser models.User, email string, username string, password *string, isActive bool, isSuperUser bool) (models.User, error) {
	// Hashed Password
	if password != nil {
		rawPassword := password
		hashedPassword, err := core.HashPassword(*rawPassword)
		if err != nil {
			return models.User{}, err
		}
		updatedUser.Password = hashedPassword
	}

	// Update data
	updatedUser.Email = email
	updatedUser.Username = username
	updatedUser.IsActive = isActive
	updatedUser.IsSuperuser = isSuperUser
	now := time.Now()
	updatedUser.UpdatedAt = &now
	if err := models.DBConn.Save(&updatedUser).Error; err != nil {
		return updatedUser, err
	}
	return updatedUser, nil
}

func DeleteUser(user models.User) (models.User, error) {
	now := time.Now()
	user.DeletedAt = &now
	if err := models.DBConn.Save(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}
