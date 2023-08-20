package day_one

import (
	"bufio"
	"data_structures_algos_go/problem_solving/advent_of_code_22"
	"fmt"
	"log"
	"sort"
	"strconv"
)

func PartTwo() {
	file, err := advent_of_code_22.ReadInputFile("day_one")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	caloriesTotal := make([]int, 0)
	sum := 0
	for scanner.Scan() {
		s := scanner.Text()
		if s == "" {
			caloriesTotal = append(caloriesTotal, sum)
			sum = 0
		} else {
			n, err := strconv.Atoi(s)
			if err != nil {
				log.Fatal(err)
			}
			sum += n

		}
	}
	sort.Ints(caloriesTotal)
	fmt.Println(sumNums(caloriesTotal[len(caloriesTotal)-3:]))
}

func sumNums(nums []int) int {
	var result int
	for _, num := range nums {
		result += num
	}
	return result
}
