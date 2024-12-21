package day3

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func CleanText(input string) string {
	pattern := regexp.MustCompile(`(?P<sequence>mul\([0-9]+,[0-9]+\))`)

	template := []byte("$sequence")

	out := []byte{}

	all_submatches := pattern.FindAllStringSubmatchIndex(input, -1)

	for _, submatch := range all_submatches {
		out = pattern.Expand(out, template, []byte(input), submatch)
	}

	return string(out)
}

func Calculate(input string) int {
	pattern := regexp.MustCompile(`(?P<num1>\d+),(?P<num2>\d+)`)

	template := []byte("$num1 $num2")

	all_submatches := pattern.FindAllStringSubmatchIndex(input, -1)

	out := 0

	for _, submatch := range all_submatches {
		temp := []byte{}

		temp = pattern.Expand(temp, template, []byte(input), submatch)

		strs := strings.Split(string(temp), " ")
		nums := []int{}
		for _, val := range strs {
			n, err := strconv.Atoi(val)
			if err != nil {
				panic(err)
			}
			nums = append(nums, n)
		}
		if len(nums) != 2 {
			log.Panicf("Error reading strings to ints, expected 2 ints, got: %#v", nums)
		}

		out += nums[0] * nums[1]
	}

	return out
}

func ReadData(path string) string {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var data strings.Builder
	for scanner.Scan() {
		_, err := data.WriteString(scanner.Text())
		if err != nil {
			log.Panicf("Error building data string: %s\n", err)
		}
	}
	return data.String()
}
