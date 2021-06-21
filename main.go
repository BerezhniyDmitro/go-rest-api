package main

import (
	"employee/api/handler"
	"employee/api/middleware"
	"employee/internal/repository"
	"employee/internal/service"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	employeeRepository := repository.NewEmployeeInMemoryRepository()
	employeeService := service.NewEmployeeService(employeeRepository)
	createEmployeeHandler := handler.NewCreateEmployeeHandler(employeeService)
	queryEmployeeHandler := handler.NewGetAllEmployeeHandler(employeeService)

	router.Use(middleware.HttpLoggerMiddleware)

	router.POST("/employee", createEmployeeHandler.Handle)
	router.GET("/employee", queryEmployeeHandler.Handle)

	err := router.Run()
	if err != nil {
		panic(err)
	}
}
