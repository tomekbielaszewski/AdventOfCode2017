package Day_9

import (
	"container/list"
	"fmt"
)

func ComputeScore(code *list.List) int {
	score := 0
	for el := code.Front(); el != nil; el = el.Next() {
		line := el.Value.(string)
		chars := []rune(line)
		fmt.Println(string(chars))
		score += computeScore(chars)
	}
	return score
}

func computeScore(chars []rune) int {
	score := 0
	charQueue := New()

	for _, char := range chars {
		switch char {
		case '{':
			{
				charQueue.Push(char)
				break
			}
		case '}':
			{
				if charQueue.Len() > 0 {
					score += charQueue.Len()
					charQueue.Pop()
				}
				break
			}
		case '!':
			{
				break
			}
		case ',':
			{
				break
			}
		case '<':
			{
				break
			}
		case '>':
			{
				break
			}
		default:
			{

			}
		}
	}

	return score
}
