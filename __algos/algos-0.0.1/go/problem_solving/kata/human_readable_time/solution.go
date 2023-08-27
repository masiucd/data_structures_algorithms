package humanreadabletime

import "fmt"

func solution(seconds int) (result string) {
	hours := seconds / 3600
	minutes := (seconds % 3600) / 60
	seconds = seconds % 60
	result = fmt.Sprintf("%02d:%02d:%02d", hours, minutes, seconds)
	return result
}
