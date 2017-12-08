package Day_8

import (
	"container/list"
	"strings"
	"strconv"
	"fmt"
)

func ComputeLargestRegisterValue(code *list.List) int {
	intOperators, boolOperators := getOperators()
	registers := compute(code, intOperators, boolOperators)
	return findLargestRegister(registers)
}

type Register struct {
	value    int
}

func compute(code *list.List, intOperators map[string]intOperator, boolOperators map[string]boolOperator) map[string]*Register {
	registers := make(map[string]*Register)
	for el := code.Front(); el != nil; el = el.Next() {
		line := strings.Fields(el.Value.(string))

		leftRegisterName := line[0]
		rightRegisterName := line[4]

		leftRegisterValue := 0
		if registers[leftRegisterName] != nil {
			leftRegisterValue = registers[leftRegisterName].value
		} else {
			registers[leftRegisterName] = &Register{value:0}
		}

		rightRegisterValue := 0
		if registers[rightRegisterName] != nil {
			rightRegisterValue = registers[rightRegisterName].value
		} else {
			registers[rightRegisterName] = &Register{value:0}
		}

		currentIntOperator := intOperators[line[1]]
		currentBoolOperator := boolOperators[line[5]]

		leftInt, convErr1 := strconv.Atoi(line[2])
		rightInt, convErr2 := strconv.Atoi(line[6])
		check(convErr1)
		check(convErr2)

		fmt.Println(line)
		fmt.Println(leftRegisterName, " ", name(currentIntOperator), " ", leftInt, " if ", rightRegisterName, " ", name(currentBoolOperator), " ", rightInt)
		fmt.Println()

		if currentBoolOperator(rightRegisterValue, rightInt) {
			registers[leftRegisterName].value = currentIntOperator(leftRegisterValue, leftInt)
		}
	}

	printRegisters(registers)
	return registers
}

func printRegisters(registers map[string]*Register) {
	for key, value := range registers {
		fmt.Printf(" %s[%v]", key, value.value)
	}
	fmt.Println()
}

func findLargestRegister(registers map[string]*Register) int {
	max := MinInt

	for _, value := range registers {
		if value.value > max {
			max = value.value
		}
	}

	return max
}

func getOperators() (map[string]intOperator, map[string]boolOperator) {
	intOperators := make(map[string]intOperator)
	intOperators["inc"] = inc
	intOperators["dec"] = dec

	boolOperators := make(map[string]boolOperator)
	boolOperators[">"] = gt
	boolOperators[">="] = gte
	boolOperators["<"] = lt
	boolOperators["<="] = lte
	boolOperators["=="] = eq
	boolOperators["!="] = neq

	return intOperators, boolOperators
}
