package Day_7

import (
	"testing"
	"fmt"
)

func TestDay7_Example1(t *testing.T) {
	input := "pbga (66)\n" +
	"xhth (57)\n" +
	"ebii (61)\n" +
	"havc (66)\n" +
	"ktlj (57)\n" +
	"fwft (72) -> ktlj, cntj, xhth\n" +
	"qoyq (66)\n" +
	"padx (45) -> pbga, havc, qoyq\n" +
	"tknk (41) -> ugml, padx, fwft\n" +
	"jptl (61)\n" +
	"ugml (68) -> gyxo, ebii, jptl\n" +
	"gyxo (61)\n" +
	"cntj (57)"

	result := GetBottomNodeName(input)

	expected := "tknk"
	if result != expected {
		t.Error("Expected ", expected, " but got ", result)
	}
}

func TestDay7_Solution(t *testing.T) {
	input := ReadFromFile("input.txt")

	result := GetBottomNodeName(input)

	fmt.Printf("Part 1 solution is: %s \n\n", result)
}

