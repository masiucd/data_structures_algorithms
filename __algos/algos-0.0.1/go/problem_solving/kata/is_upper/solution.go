package is_upper

import "strings"

type MyString string

func (s MyString) IsUpperCase() bool {
	return string(s) == strings.ToUpper(string(s))
}

func (s MyString) IsUpperCase2() bool {
	for _, chr := range string(s) {
		if chr >= 'a' && chr <= 'z' {
			return false
		}
	}

	return true
}
