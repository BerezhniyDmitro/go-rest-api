package vo

import (
	ValueObject "employee/internal/value-object"
	"testing"
)

func TestNewEmail(t *testing.T) {
	testEmail := "berezhniy.d@gmail.com"
	email, _ := ValueObject.NewEmail(testEmail)
	if email.ToString() != testEmail {
		t.Errorf("Value object email work is incorrect got %s excepted %s", testEmail, email)
	}
}

func TestNotValidNewEmail(t *testing.T) {
	email, _ := ValueObject.NewEmail("not-correct-email")
	if email.ToString() != "" {
		t.Errorf("Value object email failed validation")
	}
}
