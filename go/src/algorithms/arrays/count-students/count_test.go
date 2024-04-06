package countstudents

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCountStudents(t *testing.T) {
	students := []int{1, 1, 0, 0}
	sandwiches := []int{0, 1, 0, 1}

	got := CountStudents(students, sandwiches)
	want := 0

	assert.Equal(t, got, want)

	students = []int{1, 1, 0, 0}
	sandwiches = []int{1, 1, 1, 1}
	got = CountStudents(students, sandwiches)
	want = 2
	assert.Equal(t, got, want)

}
