package countstudents

func CountStudents(students []int, sandwiches []int) int {
	preferenceCount := [2]int{0, 0}
	for _, student := range students {
		preferenceCount[student]++
	}
	for _, sandwich := range sandwiches {
		if preferenceCount[sandwich] > 0 {
			preferenceCount[sandwich]--
		} else {
			break
		}
	}
	return preferenceCount[0] + preferenceCount[1]
}

/**
*Input: students = [1,1,0,0], sandwiches = [0,1,0,1]
Output: 0
Explanation:
- Front student leaves the top sandwich and returns to the end of the line making students = [1,0,0,1].
- Front student leaves the top sandwich and returns to the end of the line making students = [0,0,1,1].
- Front student takes the top sandwich and leaves the line making students = [0,1,1] and sandwiches = [1,0,1].
- Front student leaves the top sandwich and returns to the end of the line making students = [1,1,0].
- Front student takes the top sandwich and leaves the line making students = [1,0] and sandwiches = [0,1].
- Front student leaves the top sandwich and returns to the end of the line making students = [0,1].
- Front student takes the top sandwich and leaves the line making students = [1] and sandwiches = [1].
- Front student takes the top sandwich and leaves the line making students = [] and sandwiches = [].
Hence all students are able to eat.
*/
