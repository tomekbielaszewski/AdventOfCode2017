package Day_7

import (
	"fmt"
	"testing"
)

func TestPart1_Example(t *testing.T) {
	input := ReadFromFile("example.txt")

	result, _ := GetBottomNodeName(input)

	expected := "tknk"
	if result != expected {
		t.Error("Expected ", expected, " but got ", result)
	}
}

func TestPart1_Solution(t *testing.T) {
	input := ReadFromFile("input.txt")

	result, _ := GetBottomNodeName(input)

	fmt.Printf("Part 1 solution is: %s \n\n", result)
}
