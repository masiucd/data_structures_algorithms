package accum

import "strings"

func solution(s string) string {
	s = strings.ToLower(s)
	var result string
	for i, v := range s {
		result += strings.Repeat(string(v), i+1)
		result += "-"
	}

	ssS := strings.Split(result[:len(result)-1], "-")
	var xs []string
	for _, v := range ssS {
		res := strings.ToUpper(string(v[0])) + v
		res = res[:len(res)-1]
		xs = append(xs, res)

	}
	return strings.Join(xs, "-")
}
