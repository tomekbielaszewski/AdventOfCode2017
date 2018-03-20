package Day_4

import (
	"testing"
	"fmt"
	"container/list"
)

func TestPart1_Example1(t *testing.T) {
	input := list.New()
	input.PushFront("aa bb cc dd ee")

	result := CountValidPassphrases(input)

	expected := 1
	if result != expected {
		t.Error("Expected ", expected, " but got ", result)
	}
}

func TestPart1_Example2(t *testing.T) {
	input := list.New()
	input.PushFront("aa bb cc dd aa")

	result := CountValidPassphrases(input)

	expected := 0
	if result != expected {
		t.Error("Expected ", expected, " but got ", result)
	}
}

func TestPart1_Example3(t *testing.T) {
	input := list.New()
	input.PushFront("aa bb cc dd aaa")

	result := CountValidPassphrases(input)

	expected := 1
	if result != expected {
		t.Error("Expected ", expected, " but got ", result)
	}
}

func TestPart1_Solution(t *testing.T) {
	input := ReadFromFile("input.txt")

	result := CountValidPassphrases(input)

	fmt.Printf("Part 1 solution is: %d \n\n", result)
}
