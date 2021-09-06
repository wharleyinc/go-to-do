package router

import (
	"os"

	"github.com/gin-gonic/gin"

	"wharleyinc.com/to-do/services"
)

func Router() {

	router := gin.Default()
	router.GET("/todos", services.GetAllTodos)
	router.POST("/todos", services.CreateTodo)
	router.GET("/todos/:id", services.GetTodo)
	router.PUT("/todos/:id", services.UpdateTodo)
	router.DELETE("/todos/:id", services.DeleteTodo)
	/* router.GET("/todos", services.FindByIDWale) */
	/* router.GET("/todos/:id", services.GetTodoByID) */
	/* router.PUT("/todos/:id", services.editTodoByID) */
	/* router.POST("/todos", services.CreateTodos) */
	/* router.POST("/todos/:id", services.deleteTodoByID) */
	/* router.DELETE("/todos", services.deleteAllTodos) */

	// By default, gin serves on :8080 unless a PORT environment variable was defined/set
	// Logic to use 8080 if on local, and PORT environment variable if on Heroku
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Default port if not specified
	}

	router.Run(":" + port)
	// router.Run("localhost:8080") for a hard coded port

}
