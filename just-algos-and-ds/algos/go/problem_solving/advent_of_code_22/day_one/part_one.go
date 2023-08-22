package day_one

import (
	"bufio"
	"data_structures_algos_go/problem_solving/advent_of_code_22"
	"fmt"
	"log"
	"os"
	"strconv"
)

func PartOne() {
	file, err := advent_of_code_22.ReadInputFile("day_one")
	if err != nil {
		fmt.Println(err)

	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(file)
	scanner := bufio.NewScanner(file)
	total := 0
	max := 0
	for scanner.Scan() {
		s := scanner.Text()
		if s != "" {
			n, _ := strconv.Atoi(s)
			total += n
			if total > max {
				max = total
			}
		} else {
			total = 0
		}
	}
	fmt.Println(max)
}
