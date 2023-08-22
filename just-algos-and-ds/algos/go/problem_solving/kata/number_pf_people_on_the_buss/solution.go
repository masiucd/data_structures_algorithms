package number_pf_people_on_the_buss

func solution(busStops [][2]int) int {
	var totalOnTheBuss int
	var totalOffTheBuss int
	for _, b := range busStops {
		totalOnTheBuss += b[0]
		totalOffTheBuss += b[1]
	}
	return totalOnTheBuss - totalOffTheBuss
}
