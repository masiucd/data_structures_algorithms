package strings_ending

import "strings"

func solution(str, ending string) bool {
	return strings.HasSuffix(str, ending)

}

func solutionTwo(str, ending string) bool {
	if len(ending) > len(str) {
		return false
	}
	return str[len(str)-len(ending):] == ending
}
