package Day_7

import (
	"os"
	"bufio"
	"bytes"
)

func GetBottomNodeName(input string) string {
	return input
}

func ReadFromFile(filename string) string {
	var buffer bytes.Buffer

	file, e := os.Open(filename)
	check(e)

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		buffer.WriteString(scanner.Text())
	}

	return buffer.String()
}
func check(e error) {
	if e != nil {
		panic(e)
	}
}

