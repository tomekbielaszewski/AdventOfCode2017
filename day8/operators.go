package Day_8

type intOperator func(int, int) int
type boolOperator func(int, int) bool

func inc(a int, b int) int {
	return a+b
}

func dec(a int, b int) int {
	return a-b
}

func gt(a int, b int) bool {
	return a>b
}

func gte(a int, b int) bool {
	return a>=b
}

func lt(a int, b int) bool {
	return a<b
}

func lte(a int, b int) bool {
	return a<=b
}

func eq(a int, b int) bool {
	return a==b
}

func neq(a int, b int) bool {
	return a!=b
}