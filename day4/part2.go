package Day_4

import (
	"container/list"
	"strings"
	"sort"
)

func CountValidPassphrases2(input *list.List) int {
	invalid := 0
	for e := input.Front(); e != nil; e = e.Next() {
		passphrase := e.Value.(string)
		words := strings.Fields(passphrase)
		if isInvalid2(words) {
			invalid++
		}
	}
	return input.Len() - invalid
}

func isInvalid2(words []string) bool {
	isPresent := make(map[string]bool)
	for _, word := range words {
		word := sortString(word)
		if isPresent[word] {
			return true
			break
		}
		isPresent[word] = true
	}
	return false
}

func sortString(str string) string {
	runes := []rune(str)
	sort.Sort(sortRunes(runes))
	return string(runes)
}