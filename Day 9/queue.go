package Day_9

import "fmt"

type Queue struct {
	runes []rune
}

func New() *Queue {
	return new(Queue).Init()
}

func (q *Queue) Init() *Queue {
	fmt.Println("new queue initialized")
	q.runes = make([]rune, 0)
	return q
}

func (q *Queue) Push(char rune)  {
	fmt.Print("push: '", string(char))
	q.runes = append(q.runes, char)
	fmt.Println("' runes after push:", string(q.runes))
}

func (q *Queue) Pop() rune {
	char := q.Top()
	fmt.Print("pop: '", string(char))
	q.runes = q.runes[1:]
	fmt.Println("' runes after pop:", string(q.runes))
	return char
}

func (q *Queue) Top() rune {
	return q.runes[0]
}

func (q *Queue) IsEmpty() bool {
	return q.Len() == 0
}

func (q *Queue) Len() int {
	fmt.Println("len: ", len(q.runes))
	return len(q.runes)
}

func printSlice(s []rune) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}