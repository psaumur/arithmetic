package factorial

// needed to change output to uint because high numbers were returning 0

func Factorial(n int) uint {
	var f uint = 1
	for i := 2; i <= n; i++ {
		f *= uint(i)
	}
	return f
}

func Add(n int) int {
	return n + 1000
}
