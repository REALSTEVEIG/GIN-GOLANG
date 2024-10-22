package services

import (
	"fmt"
	"project2/database"
	"project2/models"
)

func GetAllProjects() ([]models.Project, error) {
	var projects []models.Project

	if err := database.DB.Find(&projects).Error; err != nil {
		return nil, err
	}

	return projects, nil
}

func GetSingleProject (projectId uint) (models.Project, error) {
	var project models.Project

	if err := database.DB.First(&project, projectId).Error; err != nil {
		return models.Project{}, err
	}

	return project, nil
} 

func CreateProject (project *models.Project) error {
	if err := database.DB.Create(project).Error; err != nil {
		return err
	}
	return nil
}

func UpdateProject (projectId uint, projectData *models.Project) (models.Project, error) {
	var project models.Project

	if err:= database.DB.First(&project, projectId).Error; err != nil {
		return models.Project{}, err
	}

	if err := database.DB.Model(&project).Updates(projectData).Error; err != nil {
		return models.Project{}, err
	}

	return project, nil
}

func DeleteProject (projectId uint) (string, error) {
	var project models.Project

	if err := database.DB.First(&project, projectId).Error; err != nil {
		return "", fmt.Errorf("project not found: %v", err)
	}

	if err := database.DB.Delete(&project).Error; err != nil {
		return "", fmt.Errorf("failed to delete project: %v", err)
	}

	return "Project deleted successfully", nil
}