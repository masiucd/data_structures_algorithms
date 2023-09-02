package tower_builder

import "strings"

func towerBuilder(nFloors int) []string {
	var result []string
	for i := 1; i <= nFloors; i++ {
		output := strings.Repeat(" ", nFloors-1) + strings.Repeat("*", i*2-1) + strings.Repeat(" ", nFloors-1)
		result = append(result, "\n", output)
	}
	return result
}
