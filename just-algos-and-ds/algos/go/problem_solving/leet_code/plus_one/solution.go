package plusone

import reverselist "data_structures_algos_go/problem_solving/reverse_list"


func plusOne(digits []int) []int {
	var result = reverselist.Reverse  (digits)
	carry, i := 1, 0
	for carry > 0 {
		// make sure is in bounce
		if i < len(result) {
			if result[i] == 9 {
				result[i] = 0
			} else {
				result[i] ++
				carry = 0
			}
		}else{
			result = append(result, 1)
			carry = 0
		}
		i++
	}
	result = reverselist.Reverse(result)
	return result
}
