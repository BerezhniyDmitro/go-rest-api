package value_object

import (
	"fmt"
	"unicode"
)

type Name struct {
	name string
}

func NewName(name string) (Name, error) {
	if len(name) > 20 || name == "" {
		return Name{}, fmt.Errorf("name is invalid, must be withn 20 charcters and non-empty")
	}

	if !isLetter(name) {
		return Name{}, fmt.Errorf("name is invalid, should only contain letters")
	}

	return Name{name: name}, nil
}

func (n Name) ToString() string {
	return n.name
}

func isLetter(s string) bool {
	for _, r := range s {
		if !unicode.IsLetter(r) {
			return false
		}
	}
	return true
}
