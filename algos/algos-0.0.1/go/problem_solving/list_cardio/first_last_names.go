package list_cardiopackage

import "fmt"

//Give us an array of the inventors first and last names

func firstLastNames(input []Inventor) []string {
	var result []string

	for _, v := range input {
		fullName := fmt.Sprintf("%s %s | ", v.first, v.last)
		result = append(result, fullName)
	}

	return result
}
