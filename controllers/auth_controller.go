package controllers

import (
	"net/http"

	"project2/models"
	"project2/services"

	"github.com/gin-gonic/gin"
)

type SignUpRequest struct {
	Firstname string `json:"firstname" binding:"required"`
	Lastname  string `json:"lastname" binding:"required"`
	Email     string `json:"email" binding:"required,email"`
	Password  string `json:"password" binding:"required,min=8"`
}

type LoginRequest struct {
	Email string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func Signup(c *gin.Context){
	var req SignUpRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := models.User {
		Firstname: req.Firstname,
		Lastname: req.Lastname,
		Email: req.Email,
		Password: req.Password,
	}

	createdUser, err := services.RegisterUser(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "user created successfully",
		"user": createdUser,
	})
}

func Login(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
		return
	}
	
	user, err := services.AuthenticateUser(req.Email, req.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		return
	}

	token, err := services.GenerateJWT(user.ID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not generate token"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"token": token})
}