package service

import (
	"employee/internal/dto"
	"employee/internal/entity"
	_interface "employee/internal/interface"
	ValueObject "employee/internal/value-object"
	"fmt"
	"time"
)

type EmployeeService struct {
	repository _interface.EmployeeRepositoryInterface
}

func NewEmployeeService(repository _interface.EmployeeRepositoryInterface) *EmployeeService {
	return &EmployeeService{repository: repository}
}

func (s EmployeeService) CreateEmployee(createEmployeeDTO dto.CreateEmployeeDTO) (*entity.Employee, error) {
	var employee entity.Employee
	emailValueObject, err := ValueObject.NewEmail(createEmployeeDTO.Email)
	if err != nil {
		return &employee, err
	}

	nameValueObject, err := ValueObject.NewName(createEmployeeDTO.Name)
	if err != nil {
		return &employee, err
	}

	employeeFromStorage := s.repository.Get(emailValueObject)
	if employeeFromStorage.Email.ToString() == emailValueObject.ToString() {
		return &employee, fmt.Errorf("employee with email " + createEmployeeDTO.Email + " is exist")
	}

	employee.Email = emailValueObject
	employee.Name = nameValueObject
	employee.CreatedAt = time.Now()

	s.repository.Save(&employee)

	return &employee, nil
}

func (s EmployeeService) GetAll() map[ValueObject.Email]entity.Employee {
	return s.repository.GetAll()
}
