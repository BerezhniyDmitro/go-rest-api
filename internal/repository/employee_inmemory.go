package repository

import (
	"employee/internal/entity"
	ValueObject "employee/internal/value-object"
	"sync"
)

type EmployeeInMemoryRepository struct {
	counter   int
	employees map[ValueObject.Email]entity.Employee
	sync.Mutex
}

func NewEmployeeInMemoryRepository() *EmployeeInMemoryRepository {
	counter := 0
	data := make(map[ValueObject.Email]entity.Employee)

	return &EmployeeInMemoryRepository{counter: counter, employees: data}
}

func (repository *EmployeeInMemoryRepository) Save(employee *entity.Employee) {
	repository.Lock()

	repository.employees[employee.Email] = *employee
	repository.counter++

	repository.Unlock()
}

func (repository *EmployeeInMemoryRepository) Get(email ValueObject.Email) *entity.Employee {
	employee := repository.employees[email]

	return &employee
}

func (repository *EmployeeInMemoryRepository) GetAll() map[ValueObject.Email]entity.Employee {
	return repository.employees
}
