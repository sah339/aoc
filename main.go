package main

import (
	"fmt"
	"sah339/aoc/day2"
	"sah339/aoc/day3"
)

func main() {
	day_3()
}

func day_3() {
	input := day3.ReadData("./day3/inputs.data")
	input_cleaned := day3.CleanText(input)
	total := day3.Calculate(input_cleaned)

	fmt.Printf("Result of multiplications: %d", total)

}

func day_2() {
	data := day2.ReadData("day2/inputs.data")[:1000]
	var reports []day2.Report = make([]day2.Report, 0, 10)
	for _, val := range data {
		reports = append(reports, day2.NewReport(val))
	}

	fmt.Printf("#'s Reports loaded: %d\n", len(reports))

	totalSafe := 0
	for _, val := range reports {
		for ind := range val.Data {
			val.CheckRuleOne(ind)
			val.CheckRuleTwo(ind)
			if val.Rule_one && val.Rule_two {
				totalSafe += 1
				break
			}
		}
	}
	fmt.Printf("Total safe: %d\n", totalSafe)
}
