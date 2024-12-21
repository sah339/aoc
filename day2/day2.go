package day2

import (
	"bufio"
	"math"
	"os"
	"strconv"
	"strings"
)

type Report struct {
	Data     []int
	Rule_one bool
	Rule_two bool
	Safe     bool
}

func NewReport(data []int) Report {
	r := Report{
		Data:     data,
		Rule_one: false,
		Rule_two: false,
		Safe:     false,
	}
	return r
}

func (r *Report) CheckRuleOne(skip int) {
	r.Rule_one = false
	var overallIncreasing bool
	trendFound := false
	var prev, cur int
	prev = -1
	for ind, val := range r.Data {
		if ind == skip {
			continue
		}

		if prev == -1 {
			// need to check a change so need a previous element
			prev = val
			continue
		}
		cur = val
		if trendFound {
			// check if the current trend matches the overall
			positiveChange := cur > prev
			if positiveChange != overallIncreasing {
				// if not then we can stop early
				r.Rule_two = false
				return
			}
		} else {
			// no overall trend set so set it
			overallIncreasing = cur > prev
			trendFound = true
		}
		prev = cur
	}
	// if we make it through all trend checks then the rule is passed
	r.Rule_one = true
}

func (r *Report) CheckRuleTwo(skip int) {
	var prev, cur int
	prev = -1
	for ind, val := range r.Data {
		if ind == skip {
			continue
		}

		if prev == -1 {
			prev = val
			continue
		}
		cur = val

		diff := int(math.Abs(float64(prev - cur)))
		acceptable := 1 <= diff && diff <= 3
		if !acceptable {
			r.Rule_two = false
			return
		}
		prev = cur
	}
	r.Rule_two = true
}

func ReadData(path string) [][]int {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var data [][]int = make([][]int, 1)
	var i int = 0
	for scanner.Scan() {
		data = append(data, make([]int, 0, 10))
		s_data := strings.Split(scanner.Text(), " ")
		for _, val := range s_data {
			num, err := strconv.Atoi(val)
			if err != nil {
				panic(err)
			}
			data[i] = append(data[i], num)
		}
		i++
	}
	return data
}
