package day2

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRuleOne(t *testing.T) {
	data := ReadData("inputs.test")
	var reports []Report = make([]Report, 0, 10)
	for _, val := range data {
		reports = append(reports, NewReport(val))
	}

	expected := []bool{true, true, true, false, true, true}
	for ind, val := range expected {
		reports[ind].CheckRuleOne(-1)
		assert.Equal(t, reports[ind].Rule_one, val)
	}
}

func TestRuleTwo(t *testing.T) {
	data := ReadData("inputs.test")
	var reports []Report = make([]Report, 0, 10)
	for _, val := range data {
		reports = append(reports, NewReport(val))
	}

	expected := []bool{true, false, false, true, false, true}
	for ind, val := range expected {
		reports[ind].CheckRuleTwo(-1)
		assert.Equal(t, reports[ind].Rule_two, val)
	}
}
