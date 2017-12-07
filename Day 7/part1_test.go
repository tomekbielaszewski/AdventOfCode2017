package Day_7

import (
	"fmt"
	"testing"
)

func TestDay7_Example1(t *testing.T) {
	input := ReadFromFile("example.txt")

	result, _ := GetBottomNodeName(input)

	expected := "tknk"
	fmt.Println("got ", result)
	if result != expected {
		t.Error("Expected ", expected, " but got ", result)
	}
}

func TestDay7_Solution(t *testing.T) {
	input := ReadFromFile("input.txt")

	result, _ := GetBottomNodeName(input)

	fmt.Printf("Part 1 solution is: %s \n\n", result)
}
