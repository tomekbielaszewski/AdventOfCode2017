package Day_8

import (
	"fmt"
	"testing"
)

func TestPart2_Example1(t *testing.T) {
	input := ReadFromFile("example.txt")

	result := ComputeLargestRegisterValueEver(input)

	expected := 10
	if result != expected {
		t.Error("Expected ", expected, " but got ", result)
	}
}

func TestPart2_Solution(t *testing.T) {
	input := ReadFromFile("input.txt")

	result := ComputeLargestRegisterValueEver(input)

	fmt.Printf("Part 1 solution is: %d \n\n", result)
}
