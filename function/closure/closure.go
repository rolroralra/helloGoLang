package closure

func Summation() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func Fibonacci() func() int {
	prev, next := 0, 1
	return func() int {
		result := prev
		prev = next
		next = result + prev
		return result
	}
}
