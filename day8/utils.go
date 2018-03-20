package Day_8

import (
	"container/list"
	"os"
	"bufio"
	"runtime"
	"reflect"
	"strings"
)

const MaxUint = ^uint(0)
const MaxInt = int(MaxUint >> 1)
const MinInt = -MaxInt - 1

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

func name(i interface{}) string {
	return strings.Split(runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name(), ".")[1]
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}