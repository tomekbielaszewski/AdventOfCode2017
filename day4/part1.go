package Day_4

import (
	"container/list"
	"strings"
)

func CountValidPassphrases(input *list.List) int {
	invalid := 0
	for e := input.Front(); e != nil; e = e.Next() {
		passphrase := e.Value.(string)
		words := strings.Fields(passphrase)
		if isInvalid(words) {
			invalid++
		}
	}
	return input.Len() - invalid
}

func isInvalid(words []string) bool {
	isPresent := make(map[string]bool)
	for _, word := range words {
		if isPresent[word] {
			return true
			break
		}
		isPresent[word] = true
	}
	return false
}
