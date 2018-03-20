package Day_7

import (
	"container/list"
	"os"
	"bufio"
)

func ReadFromFile(filename string) *list.List {
	stringLines := list.New()

	file, err := os.Open(filename)
	check(err)

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		stringLines.PushBack(scanner.Text())
	}

	return stringLines
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
