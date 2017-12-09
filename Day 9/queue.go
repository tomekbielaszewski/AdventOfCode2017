package Day_9

type Queue struct {
	runes []rune
}

func NewQueue() *Queue {
	return new(Queue).Init()
}

func (q *Queue) Init() *Queue {
	q.runes = make([]rune, 0)
	return q
}

func (q *Queue) Push(char rune) {
	q.runes = append(q.runes, char)
}

func (q *Queue) Pop() rune {
	char := q.Top()
	q.runes = q.runes[1:]
	return char
}

func (q *Queue) Top() rune {
	return q.runes[0]
}

func (q *Queue) IsEmpty() bool {
	return q.Len() == 0
}

func (q *Queue) Len() int {
	return len(q.runes)
}
