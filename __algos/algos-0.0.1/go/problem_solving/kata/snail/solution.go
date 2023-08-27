package snail

func snail(snaipMap [][]int) []int {
	var result []int
	for len(snaipMap) > 0 {
		result = append(result, snaipMap[0]...)
		snaipMap = snaipMap[1:]
		if len(snaipMap) == 0 {
			break
		}
		for i := 0; i < len(snaipMap); i++ {
			result = append(result, snaipMap[i][len(snaipMap[i])-1])
			snaipMap[i] = snaipMap[i][:len(snaipMap[i])-1]
		}
		if len(snaipMap[0]) == 0 {
			break
		}
		for i := len(snaipMap) - 1; i >= 0; i-- {
			result = append(result, snaipMap[len(snaipMap)-1][i])
		}
		snaipMap = snaipMap[:len(snaipMap)-1]
		if len(snaipMap) == 0 {
			break
		}
		for i := len(snaipMap) - 1; i >= 0; i-- {
			result = append(result, snaipMap[i][0])
			snaipMap[i] = snaipMap[i][1:]
		}
	}
	return result

}
