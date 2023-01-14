package highestscoringword

import "strings"

func solution(s string) string {
	var result string
	var highest int
	stringList := strings.Split(s, " ")
	for _, word := range stringList {
		score := 0
		for _, char := range word {
			score += int(char) - 96
		}
		if score > highest {
			highest = score
			result = word
		}
	}
	return result
}
