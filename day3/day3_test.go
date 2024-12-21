package day3

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadData(t *testing.T) {
	input := ReadData("inputs.test")
	expected := "xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))"

	assert.Equal(t, input, expected)
}

func TestCleanText(t *testing.T) {
	input := ReadData("inputs.test")
	expected := "mul(2,4)mul(5,5)mul(11,8)mul(8,5)"

	input_cleaned := CleanText(input)

	assert.Equal(t, expected, input_cleaned)
}

func TestCalculate(t *testing.T) {
	input := "mul(2,4)mul(5,5)mul(11,8)mul(8,5)"
	expected := 161

	output := Calculate(input)
	assert.Equal(t, expected, output)
}
