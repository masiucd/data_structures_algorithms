package number_pf_people_on_the_buss

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNumberOfPeopleOnTheBuss(t *testing.T) {
	actual := solution([][2]int{{10, 0}, {3, 5}, {5, 8}})
	expected := 5
	assert.Equal(t, expected, actual, fmt.Sprintf("Expected %d  actual %d", expected, actual))

	actual = solution([][2]int{{3, 0}, {9, 1}, {4, 10}, {12, 2}, {6, 1}, {7, 10}})
	expected = 17
	assert.Equal(t, expected, actual, fmt.Sprintf("Expected %d  actual %d", expected, actual))
}
