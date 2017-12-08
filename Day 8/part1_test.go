package Day_8

import (
	"fmt"
	"testing"
)

func TestPart1_Example1(t *testing.T) {
	input := ReadFromFile("example.txt")

	result := ComputeLargestRegisterValue(input)

	expected := 1
	if result != expected {
		t.Error("Expected ", expected, " but got ", result)
	}
}

func TestPart1_Solution(t *testing.T) {
	input := ReadFromFile("input.txt")

	result := ComputeLargestRegisterValue(input)

	fmt.Printf("Part 1 solution is: %d \n\n", result)
}
