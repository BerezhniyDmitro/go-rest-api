package entity

import (
	ValueObject "employee/internal/value-object"
	"time"
)

type Employee struct {
	Name      ValueObject.Name  `json:"name"`
	Email     ValueObject.Email `json:"email"`
	CreatedAt time.Time         `json:"created_at"`
}
