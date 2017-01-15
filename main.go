package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func solve(numbers []int, min, max int) int {
	size := len(numbers)
	hash := make(map[int]bool, size)

	for i := 0; i < size; i++ {
		hash[numbers[i]] = true
	}

	solutions := make(map[int][]int)

	for t := min; t < max+1; t++ {
		if t%100 == 0 {
			fmt.Printf("Current: %vs: @%v\n", t, len(solutions))
		}

		for x := range hash {
			_, exists := solutions[t]
			if exists {
				continue
			}

			y := t - x

			if x == y {
				continue
			}

			if hash[y] {
				solutions[t] = []int{x, y}
			}
		}
	}

	return len(solutions)
}

func main() {
	solution := solve(load("pa61_sort.txt"), -10000, 10000)
	fmt.Printf("Solution: %v\n", solution)
}

func load(s string) []int {
	inputBytes, err := ioutil.ReadFile(s)

	poof(err)

	inputString := string(inputBytes[:])
	rows := strings.Split(inputString, "\n")

	// Lose the last row because I'm too lazy to either delete the trailing \n or to read about an unexpected behavior in golang ioutil.ReadFile()
	size := len(rows) - 1

	fmt.Printf("size: %v\n", size)

	data := make([]int, size)

	fmt.Println("Importing")

	for j := 0; j < size; j++ {
		is := rows[j]
		i, err := strconv.Atoi(strings.TrimSpace(is))

		poof(err)

		data[j] = i
	}

	fmt.Println("Imported")

	return data
}

func poof(err error) {
	if err != nil {
		panic(err)
	}
}
