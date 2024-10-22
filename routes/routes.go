package routes

import (
	"project2/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	todoRoutes := r.Group("/api/todos")
	{
		todoRoutes.GET("/", controllers.GetTodos)
		todoRoutes.POST("/", controllers.CreateTodo)
		todoRoutes.GET("/:id", controllers.GetSingleTodo)
		todoRoutes.PUT("/:id", controllers.UpdateTodo)
		todoRoutes.DELETE("/:id", controllers.DeleteTodo)
	}

	projectRoutes := r.Group("/api/projects")
	{
		projectRoutes.GET("/", controllers.GetAllProjects)
		projectRoutes.POST("/", controllers.CreateProject)
		projectRoutes.GET("/:id", controllers.GetSingleProject)
		projectRoutes.PUT("/:id", controllers.UpdateProject)
		projectRoutes.DELETE("/:id", controllers.DeleteProject)
	}
}