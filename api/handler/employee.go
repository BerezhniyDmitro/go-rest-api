package handler

import (
	"employee/api/response"
	"employee/internal/dto"
	_interface "employee/internal/interface"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CreateEmployeeHandler struct {
	service _interface.CommandEmployeeInterface
}

func NewCreateEmployeeHandler(service _interface.CommandEmployeeInterface) *CreateEmployeeHandler {
	return &CreateEmployeeHandler{service: service}
}

func (employeeHandler CreateEmployeeHandler) Handle(c *gin.Context) {
	var createEmployeeDto dto.CreateEmployeeDTO

	if err := c.BindJSON(&createEmployeeDto); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	employee, err := employeeHandler.service.CreateEmployee(createEmployeeDto)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, response.ErrorResponse{
			Status:  http.StatusUnprocessableEntity,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, response.NewJsonEmployeeDto(
		employee.Name.ToString(),
		employee.Email.ToString(),
		employee.CreatedAt.Format(http.TimeFormat),
	))
}

type GetAllEmployeeHandler struct {
	service _interface.QueryEmployeeInterface
}

func (g GetAllEmployeeHandler) Handle(c *gin.Context) {
	employees := g.service.GetAll()
	employeesToResponse := make(map[int]interface{})

	counter := 0
	for _, employee := range employees {
		employeesToResponse[counter] = response.NewJsonEmployeeDto(
			employee.Name.ToString(),
			employee.Email.ToString(),
			employee.CreatedAt.Format(http.TimeFormat),
		)
		counter++
	}

	c.JSON(http.StatusOK, employeesToResponse)
}

func NewGetAllEmployeeHandler(service _interface.QueryEmployeeInterface) *GetAllEmployeeHandler {
	return &GetAllEmployeeHandler{service: service}
}
