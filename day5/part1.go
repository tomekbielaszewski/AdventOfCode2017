package Day_5

func CalculateSteps(instructions []int) int {
	index := 0
	steps := 0

	for !IsOutside(index, instructions) {
		el := instructions[index]
		instructions[index]++
		index += el
		steps++
	}

	return steps
}

func IsOutside(index int, instructions []int) bool {
	return index < 0 || index >= len(instructions)
}

