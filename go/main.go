package main

import "fmt"

func main() {
	fmt.Println("res  ===> ", solution("AAAABBBCCDAA"))
}

func solution(input string) string {

	var result string
	count := 1
	for i := 1; i <= len(input); i++ {
		currentChar := input[i-1]
		if i <= len(input)-1&& currentChar == input[i] {
			count++
		} else {
			addCharAndAmount(&result, count, currentChar)
			count = 1
		}
	}
	return result
}

func addCharAndAmount(result *string, count int, char uint8) {
	*result += fmt.Sprintf("%d", count)
	*result += string(char)
}
