package router

import (
	"github.com/gin-gonic/gin"

	"wharleyinc.com/to-do/services"
)

func Router() {

	router := gin.Default()
	router.GET("/todos", services.GetTodos)
	router.GET("/todos", services.FindByIDWale)
	/* router.GET("/todos/:id", services.GetTodoByID) */
	/* router.PUT("/todos/:id", services.editTodoByID) */
	router.POST("/todos", services.CreateTodos)
	/* 	router.POST("/todos/:id", services.deleteTodoByID)
	   	router.DELETE("/todos", services.deleteAllTodos) */

	router.Run("localhost:8080")
}
