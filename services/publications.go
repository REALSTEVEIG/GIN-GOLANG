package services

import (
	"fmt"
	"project2/database"
	"project2/models"
)

func GetAllPublications () ([]models.Publications, error) {
	var publications []models.Publications

	if err := database.DB.Find(&publications).Error; err != nil {
		return nil, err
	}

	return publications, nil
}

func GetSinglePublication (publicationId uint) (models.Publications, error) {
	var publication models.Publications

	if err := database.DB.First(&publication, publicationId).Error ; err != nil {
		return models.Publications{}, err
	}
	return publication, nil
}

func CreatePublication (publicationBody *models.Publications) (models.Publications, error) {
	if err := database.DB.Create(publicationBody).Error; err != nil {
		return models.Publications{}, err
	}
	return *publicationBody, nil
}

func UpdatePublication (publicationId uint, publicationBody *models.Publications) (models.Publications, error) {
	var publication models.Publications

	if err := database.DB.First(&publication, publicationId).Error; err != nil {
		return models.Publications{}, err
	}

	if err := database.DB.Model(&publication).Updates(publicationBody).Error; err != nil {
		return models.Publications{}, err
	}

	return publication, nil
}

func DeletePublication (publicationId uint) (string, error) {
	var publication models.Publications

	if err := database.DB.First(&publication, publicationId).Error; err != nil {
		return "", fmt.Errorf("publication not found: %v", err)
	}

	if err := database.DB.Delete(&publication).Error; err != nil {
		return "", fmt.Errorf("failed to delete publication: %v", err)
	}

	return "Publication deleted successfully", nil
}