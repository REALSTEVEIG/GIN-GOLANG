package services

import (
	"fmt"
	"project2/database"
	"project2/models"
)

func GetAllTodos () ([]models.Todo, error) {
	var todos []models.Todo

	if err := database.DB.Find(&todos).Error; err != nil {
		return nil, err
	}
	return todos, nil
}

func CreateTodo(todo *models.Todo) error {
	if err := database.DB.Create(todo).Error; err != nil {
		return err
	}
	return nil
}

func GetSingleTodo(todoId uint) (models.Todo, error) {
	var todo models.Todo

	if err := database.DB.First(&todo, todoId).Error; err != nil {
		return models.Todo{}, err
	}

	return todo, nil
}

func UpdateSingleTodo(todoId uint, todoData *models.Todo) (models.Todo, error) {
	var todo models.Todo

	if err := database.DB.First(&todo, todoId).Error; err != nil {
		return models.Todo{}, err
	}

	if err := database.DB.Model(&todo).Updates(todoData).Error; err != nil {
		return models.Todo{}, err
	}

	return todo, nil
}

func DeleteTodo(todoId uint) (string, error) {
	var todo models.Todo

	if err := database.DB.First(&todo, todoId).Error; err != nil {
		return "", fmt.Errorf("todo not found: %v", err)
	}
	if err := database.DB.Delete(&todo).Error; err != nil {
		return "", fmt.Errorf("failed to delete todo: %v", err)
	}
	return "Todo deleted successfully", nil
}