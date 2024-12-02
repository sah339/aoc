package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {

	doPart1 := flag.Bool("one", false, "Include to output part 1 answer")
	doPart2 := flag.Bool("two", false, "Include to output part 2 answers")
	flag.Parse()

	if *doPart1 {
		_ = part1()
	}
	if *doPart2 {
		_ = part2()
	}
	if !*doPart1 && !*doPart2 {
		fmt.Printf("No flags provided, which parts should be computed")
	}

	return
}

func part1() int {
	left, right := readData("inputs.txt")

	slices.Sort(left)
	slices.Sort(right)

	var totalDistance int

	for i := 0; i < len(left); i++ {
		distance := left[i] - right[i]
		if distance < 0 {
			distance = -distance
		}

		totalDistance += distance
	}

	fmt.Printf("Total distance: %d\n", totalDistance)
	return totalDistance
}

func part2() int {
	left, right := readData("inputs.txt")

	// Create map containing occurence data
	rightOccurences := make(map[int]int)
	for _, val := range right {
		rightOccurences[val] += 1
	}

	// Calculate similarity
	var similarity int
	for _, val := range left {
		similarity += val * rightOccurences[val]
	}

	fmt.Printf("Total similarity: %d\n", similarity)
	return similarity
}

func readData(path string) ([]int, []int) {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var left []int
	var right []int

	for scanner.Scan() {
		st := strings.Split(scanner.Text(), "   ")
		l, err := strconv.Atoi(st[0])
		if err != nil {
			panic(err)
		}
		r, err := strconv.Atoi(st[1])
		if err != nil {
			panic(err)
		}
		left = append(left, l)
		right = append(right, r)
	}
	return left, right
}
