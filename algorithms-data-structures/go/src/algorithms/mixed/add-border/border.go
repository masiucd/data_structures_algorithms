package addborder

import "strings"

func addBorder(picture []string) []string {
	wallLength := len(picture[0]) + 2
	border := strings.Repeat("*", wallLength)
	xs := []string{border}
	for _, v := range picture {
		xs = append(xs, "*"+v+"*")
	}
	xs = append(xs, border)
	return xs
}
