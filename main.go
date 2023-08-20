package main


import (
	handlers "maintmp/internal/handlers"
	"github.com/labstack/echo/v4"
	
	"github.com/swaggo/echo-swagger"
	_ "maintmp/docs"  
)

// @title Swagger Example
// @version 1.0
// @description This is a aaa.

// @host localhost:5000
// @BasePath /
// @name todolist
func main() {
	
	
	e := echo.New()
	
	e.GET("/task", handlers.HandleGet)
	e.GET("/task/all", handlers.HandleGetAll)
    e.POST("/task", handlers.HandlePost)
	e.PUT("/task", handlers.HandleUpdate)
	e.DELETE("/task", handlers.HandleDelete)
	e.GET("/swagger/*", echoSwagger.WrapHandler)
	
	e.Start(":5000")
	
}