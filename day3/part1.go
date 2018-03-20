package Day_3

func CalculateDistance1(x int) int {
	if x == 1 {
		return 0
	}

	n := 0
	q := 0

	for i := 1; q < x; i++ {
		n = i
		q = valueOfRightLowerBlockOnNthLayer(i)
	}

	q_of_lower_level := valueOfRightLowerBlockOnNthLayer(n - 1)
	s := q - q_of_lower_level
	x_overflow := x - q_of_lower_level
	s_quarter := s / 4
	s_half_quarter := s / 8
	additional_steps := x_overflow % s_quarter - s_half_quarter

	if additional_steps < 0 {
		additional_steps = -additional_steps
	}

	return (n - 1) + additional_steps
}

func valueOfRightLowerBlockOnNthLayer(n int) int {
	if n == 0 {
		return 1
	}
	sum := 1
	for i := 1; i <= n; i++ {
		sum += (i*2 - 2) * 4
	}
	return sum
}
