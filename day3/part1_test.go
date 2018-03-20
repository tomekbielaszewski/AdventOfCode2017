package Day_3

import (
	"fmt"
	"testing"
)

func TestPart1_Example1(t *testing.T) {
	input := 1

	result := CalculateDistance1(input)

	expected := 0
	if result != expected {
		t.Error("Expected ", expected, " but got ", result)
	}
}

func TestPart1_Example2(t *testing.T) {
	input := 12

	result := CalculateDistance1(input)

	expected := 3
	if result != expected {
		t.Error("Expected ", expected, " but got ", result)
	}
}

func TestPart1_Example3(t *testing.T) {
	input := 23

	result := CalculateDistance1(input)

	expected := 2
	if result != expected {
		t.Error("Expected ", expected, " but got ", result)
	}
}

func TestPart1_Example4(t *testing.T) {
	input := 1024

	result := CalculateDistance1(input)

	expected := 31
	if result != expected {
		t.Error("Expected ", expected, " but got ", result)
	}
}

func TestPart1_Solution(t *testing.T) {
	input := 347991

	result := CalculateDistance1(input)

	fmt.Printf("Part 1 solution is: %d \n\n", result)
}
