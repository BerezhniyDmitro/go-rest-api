package vo

import (
	ValueObject "employee/internal/value-object"
	"testing"
)

func TestNewName(t *testing.T) {
	testName := "testName"
	name, _ := ValueObject.NewName(testName)
	if name.ToString() != testName {
		t.Errorf("Value object name work is incorrect got %s excepted %s", testName, name)
	}
}

func TestEmptyValueNewName(t *testing.T) {
	name, _ := ValueObject.NewName("")
	if name.ToString() != "" {
		t.Errorf("Value object validation failed got empty value, exepted %s", name.ToString())
	}
}

func TestLongValueNewName(t *testing.T) {
	name, _ := ValueObject.NewName("big_name_20_characters_in_name")
	if name.ToString() != "" {
		t.Errorf("Value object validation failed got length < 20 charactesrs value, exepted %s", name.ToString())
	}
}

func TestNotLeterValueNewName(t *testing.T) {
	name, _ := ValueObject.NewName("_*nameIsNotvalid)(!")
	if name.ToString() != "" {
		t.Errorf("Value object validation failed got only letter value, exepted %s", name.ToString())
	}
}
