package Day_5

import (
	"os"
	"bufio"
	"strconv"
)

func ReadFromFile(filename string) []int {
	integers := make([]int, 0)

	file, err := os.Open(filename)
	check(err)

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		integer, _ := strconv.Atoi(scanner.Text())
		integers = append(integers, integer)
	}

	return integers
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
