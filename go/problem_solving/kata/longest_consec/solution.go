package longest_consec

import "strings"

func solution(input []string, k int) string {
	longest := ""
	if k > 0 && len(input) >= k {
		for i := 0; i < len(input)-k+1; i++ {
			s := strings.Join(input[i:i+k], "")
			if len(s) > len(longest) {
				longest = s
			}
		}
	}
	return longest
}
