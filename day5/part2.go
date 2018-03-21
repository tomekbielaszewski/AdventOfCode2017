package Day_5

func CalculateSteps2(instructions []int) int {
	index := 0
	steps := 0

	for !IsOutside(index, instructions) {
		el := instructions[index]
		if el < 3 {
			instructions[index]++
		} else {
			instructions[index]--
		}
		index += el
		steps++
	}

	return steps
}

