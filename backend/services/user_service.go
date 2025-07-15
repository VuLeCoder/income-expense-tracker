package services

import (
	"errors"
	"server/databases"
	"server/models"
)

func GetUserByUsernameAndPassword(username string, password string) (*models.User, error) {
	user := &models.User{}

	err := databases.DB.Where("username = ? AND password = ?", username, password).First(user).Error
	if err != nil {
		return nil, errors.New("Không tìm thấy người dùng")
	}

	return user, nil
}