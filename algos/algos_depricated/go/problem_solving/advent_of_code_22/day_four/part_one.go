package day_four

import (
	"bufio"
	"data_structures_algos_go/problem_solving/advent_of_code_22"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

//For example, 2-8 fully contains 3-7
// and 6-6 is fully contained by 4-6

func PartOne() {
	file, err := advent_of_code_22.ReadInputFile("day_four")
	if err != nil {
		panic(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(file)

	var count int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		pairs := strings.Split(line, ",")
		first, second := pairs[0], pairs[1]
		x := strings.Split(first, "-")
		x2 := strings.Split(second, "-")
		n1, _ := strconv.Atoi(x[0])
		n2, _ := strconv.Atoi(x[1])
		n3, _ := strconv.Atoi(x2[0])
		n4, _ := strconv.Atoi(x2[1])
		list1 := getRangeList(n1, n2)
		list2 := getRangeList(n3, n4)
		fmt.Println(list1, list2)
		if every(list1, list2) {
			count++
		}
		if every(list2, list1) {
			count++
		}

	}
	//472 to low

	fmt.Println(count)

}

func getRangeList(from, to int) []int {
	var list []int
	for i := from; i <= to; i++ {
		list = append(list, i)
	}
	return list
}

func every(list1, list2 []int) bool {
	//get the longest ist before
	var toIterate []int
	var toBeCompared []int
	if len(list1) > len(list2) {
		toIterate = list1
		toBeCompared = list2
	} else {
		toIterate = list2
		toBeCompared = list1
	}

	for i := 0; i < len(toBeCompared); i++ {
		if toIterate[i] != toBeCompared[i] {
			return false
		}
	}
	return true
}
