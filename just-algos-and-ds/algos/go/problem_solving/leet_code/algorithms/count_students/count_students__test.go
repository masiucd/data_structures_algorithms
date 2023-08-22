package count_students

import "testing"

func TestCountStudents(t *testing.T) {
	students := []int{1, 1, 0, 0}
	sandwiches := []int{0, 1, 0, 1}

	got := countStudents(students, sandwiches)
	want := 0
	if got != want {
		t.Errorf("expected %d got %d", want, got)
	}

	students = []int{1, 1, 1, 0, 0, 1}
	sandwiches = []int{1, 0, 0, 0, 1, 1}
	got = countStudents(students, sandwiches)
	want = 3
	if got != want {
		t.Errorf("expected %d got %d", want, got)
	}

}
