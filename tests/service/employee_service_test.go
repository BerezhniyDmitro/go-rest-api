package service

import (
	"employee/internal/dto"
	"employee/internal/entity"
	"employee/internal/service"
	ValueObject "employee/internal/value-object"
	"testing"
)

type MockEmployeeRepository struct {
	employees map[ValueObject.Email]entity.Employee
}

func NewMockEmployeeRepository() *MockEmployeeRepository {
	//реализация совпадает так как основной реп должен был бы использовать базу данных, а не память
	data := make(map[ValueObject.Email]entity.Employee)
	return &MockEmployeeRepository{data}
}

func (m MockEmployeeRepository) Save(e *entity.Employee) {
	m.employees[e.Email] = *e
}

func (m MockEmployeeRepository) Get(e ValueObject.Email) *entity.Employee {
	employee := m.employees[e]

	return &employee
}

func (m *MockEmployeeRepository) GetAll() map[ValueObject.Email]entity.Employee {
	return m.employees
}

func TestCreateEmployee(t *testing.T) {
	employeeService := service.NewEmployeeService(NewMockEmployeeRepository())

	name := "test"
	email := "test@gmail.com"
	createEmployeeDto := dto.CreateEmployeeDTO{Name: name, Email: email}
	employee, err := employeeService.CreateEmployee(createEmployeeDto)

	if err != nil {
		t.Errorf("EmployeeService failed create employee %s", err)
	}

	employeeNameToString := employee.Name.ToString()
	if employeeNameToString != name {
		t.Errorf(
			"EmployeeService failed create employee, error name excepted %s, actual %s",
			name,
			employeeNameToString,
		)
	}

	employeeEmailToString := employee.Email.ToString()
	if employeeEmailToString != email {
		t.Errorf(
			"EmployeeService failed create employee, error email excepted %s, actual %s",
			name,
			employeeNameToString,
		)
	}
}

func TestErrorToDuplicateCreateEmployee(t *testing.T) {
	employeeService := service.NewEmployeeService(NewMockEmployeeRepository())

	name := "test"
	email := "test@gmail.com"
	createEmployeeDto := dto.CreateEmployeeDTO{Name: name, Email: email}

	_, err := employeeService.CreateEmployee(createEmployeeDto)
	employee2, err2 := employeeService.CreateEmployee(createEmployeeDto)

	if err != nil {
		t.Errorf("EmployeeService failed create employee %s", err)
	}

	if err2 == nil {
		t.Errorf("EmployeeService create duplicate employee %s", employee2)
	}
}
