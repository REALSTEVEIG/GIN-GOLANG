package controllers

import (
	"net/http"
	"project2/models"
	"project2/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetTodos(c *gin.Context) {
	todos, err := services.GetAllTodos()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, todos)
}

func CreateTodo(c *gin.Context) {
	var todo models.Todo
	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := services.CreateTodo(&todo); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, todo)
}

func GetSingleTodo (c *gin.Context) {
	todoId, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid todo ID"})
		return
	}

	todo, err := services.GetSingleTodo(uint(todoId))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, todo)
}

func UpdateTodo (c *gin.Context) {
	todoId, err := strconv.Atoi(c.Param("id"));

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Invalid todo ID"})
		return
	}

	var todoData models.Todo

	if err := c.ShouldBindJSON(&todoData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Invalid todo request"})
		return
	}

	updatedTodo, err := services.UpdateSingleTodo(uint(todoId), &todoData)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, updatedTodo)
}

func DeleteTodo(c *gin.Context) {
	todoId, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Todo ID"})
		return
	}

	message, err := services.DeleteTodo(uint(todoId))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": message})
}