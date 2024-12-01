package controllers

import (
	"net/http"
	"project2/models"
	"project2/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllPublications (c *gin.Context) {
	publications, err := services.GetAllPublications();

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, publications)
} 

func GetSinglePublication (c *gin.Context) {
	publicationId, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid project Id"})
		return
	}

	publication, err := services.GetSinglePublication(uint(publicationId))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, publication)
}

func CreatePublication(c *gin.Context) {
    var publication models.Publications

    if err := c.ShouldBindJSON(&publication); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    createdPublication, err := services.CreatePublication(&publication)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusCreated, createdPublication)
}

func UpdatePublication (c *gin.Context) {
	publicationId, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid publication ID"})
	}

	var publicationData models.Publications

	if err := c.ShouldBindJSON(&publicationData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input data"})
		return
	}

	updatePublication, err := services.UpdatePublication(uint(publicationId), &publicationData)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, updatePublication)
}

func DeletePublication (c *gin.Context) {
	publicationId, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid publication ID"})
		return
	}

	message, err := services.DeletePublication(uint(publicationId))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": message})
}