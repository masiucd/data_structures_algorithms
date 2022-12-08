package advent_of_code_22

import (
	"os"
	"path/filepath"
)

// ReadInputFile for example ->  day = day_one
func ReadInputFile(day string) (*os.File, error) {
	base := "./problem_solving"
	file, err := os.Open(filepath.Join(base, "advent_of_code_22", day, "input.txt"))
	return file, err
}
