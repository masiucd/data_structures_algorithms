package count_students

//https://leetcode.com/problems/number-of-students-unable-to-eat-lunch/

func countStudents(students []int, sandwiches []int) int {
	var miss = 0
	for miss != len(students) { // if queue as made a full round then stop
		if students[0] == sandwiches[0] {
			students = students[1:]
			sandwiches = sandwiches[1:]
			miss = 0
		} else {
			var student int
			student, students = students[0], students[1:]
			students = append(students, student)
			miss++
		}
	}
	return len(students)
}
