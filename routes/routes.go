package routes

import (
	"github.com/gin-gonic/gin"
	"project2/controllers"
	"project2/middlewares"
)

func SetupRoutes(r *gin.Engine) {
	authRoutes := r.Group("/api/auth")
	{
		authRoutes.POST("/signup", controllers.Signup)
		authRoutes.POST("/login", controllers.Login)
	}

	todoRoutes := r.Group("/api/todos")
	{
		todoRoutes.GET("/", controllers.GetTodos)
		todoRoutes.POST("/", middlewares.AuthMiddleware(), controllers.CreateTodo)
		todoRoutes.GET("/:id", controllers.GetSingleTodo)
		todoRoutes.PUT("/:id", middlewares.AuthMiddleware(), controllers.UpdateTodo)
		todoRoutes.DELETE("/:id", middlewares.AuthMiddleware(), controllers.DeleteTodo)
	}

	projectRoutes := r.Group("/api/projects")
	{
		projectRoutes.GET("/", controllers.GetAllProjects)
		projectRoutes.POST("/", middlewares.AuthMiddleware(), controllers.CreateProject)
		projectRoutes.GET("/:id", controllers.GetSingleProject)
		projectRoutes.PUT("/:id", middlewares.AuthMiddleware(), controllers.UpdateProject)
		projectRoutes.DELETE("/:id", middlewares.AuthMiddleware(), controllers.DeleteProject)
	}
	publicationRoutes := r.Group("/api/publications")
	{
		publicationRoutes.GET("/", controllers.GetAllPublications)
		publicationRoutes.POST("/", middlewares.AuthMiddleware(), controllers.CreatePublication)
		publicationRoutes.GET("/:id", controllers.GetSinglePublication)
		publicationRoutes.PUT("/:id", middlewares.AuthMiddleware(), controllers.UpdatePublication)
		publicationRoutes.DELETE("/:id", middlewares.AuthMiddleware(), controllers.DeletePublication)
	}

}