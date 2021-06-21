package value_object

import (
	"fmt"
	"net/mail"
)

type Email struct {
	email string
}

func NewEmail(email string) (Email, error) {
	_, err := mail.ParseAddress(email)
	if err != nil {
		return Email{}, fmt.Errorf("email is not valid, please input correct email")
	}

	return Email{email: email}, nil
}

func (e Email) ToString() string {
	return e.email
}
