package day6

import (
	"testing"
	"fmt"
)

func TestPart1_Example1(t *testing.T) {
	input := [4]int{0, 2, 7, 0}

	result := CalculateCycles(input[:])

	expected := 5
	if result != expected {
		t.Error("Expected ", expected, " but got ", result)
	}
}

func TestPart1_Solution(t *testing.T) {
	input := ReadFromFile("input.txt")

	result := CalculateCycles(input)

	fmt.Printf("Part 1 solution is: %d \n\n", result)
}
