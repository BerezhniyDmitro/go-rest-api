package response

type JsonEmployeeDto struct {
	Name      string `json:"name"`
	Email     string `json:"email"`
	CreatedAt string `json:"created_at"`
}

func NewJsonEmployeeDto(name string, email string, createdAt string) *JsonEmployeeDto {
	return &JsonEmployeeDto{Name: name, Email: email, CreatedAt: createdAt}
}
