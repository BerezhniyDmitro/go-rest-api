package _interface

import (
	"employee/internal/dto"
	"employee/internal/entity"
	ValueObject "employee/internal/value-object"
)

type EmployeeRepositoryInterface interface {
	Save(e *entity.Employee)
	Get(e ValueObject.Email) *entity.Employee
	GetAll() map[ValueObject.Email]entity.Employee
}

type CommandEmployeeInterface interface {
	CreateEmployee(createEmployeeDTO dto.CreateEmployeeDTO) (*entity.Employee, error)
}

type QueryEmployeeInterface interface {
	GetAll() map[ValueObject.Email]entity.Employee
}
