package Day_8

import (
	"container/list"
	"strings"
	"strconv"
)

func ComputeLargestRegisterValueEver(code *list.List) int {
	intOperators, boolOperators := getOperators()
	return computeLargestRegisterEver(code, intOperators, boolOperators)
}

func computeLargestRegisterEver(code *list.List, intOperators map[string]intOperator, boolOperators map[string]boolOperator) int {
	registers := make(map[string]*Register)
	maxRegisterEver := MinInt

	for el := code.Front(); el != nil; el = el.Next() {
		line := strings.Fields(el.Value.(string))

		leftRegisterName := line[0]
		rightRegisterName := line[4]

		leftRegisterValue := 0
		if registers[leftRegisterName] != nil {
			leftRegisterValue = registers[leftRegisterName].value
		} else {
			registers[leftRegisterName] = &Register{value: 0}
		}

		rightRegisterValue := 0
		if registers[rightRegisterName] != nil {
			rightRegisterValue = registers[rightRegisterName].value
		} else {
			registers[rightRegisterName] = &Register{value: 0}
		}

		currentIntOperator := intOperators[line[1]]
		currentBoolOperator := boolOperators[line[5]]

		leftInt, convErr1 := strconv.Atoi(line[2])
		rightInt, convErr2 := strconv.Atoi(line[6])
		check(convErr1)
		check(convErr2)

		if currentBoolOperator(rightRegisterValue, rightInt) {
			registers[leftRegisterName].value = currentIntOperator(leftRegisterValue, leftInt)

			maxRegister := findLargestRegister(registers)
			if maxRegister > maxRegisterEver {
				maxRegisterEver = maxRegister
			}
		}
	}

	return maxRegisterEver
}
