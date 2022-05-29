package factorial

func Factorial(n int) int {
	var f int = 1
	for i := 2; i <= n; i++ {
		f *= i
	}
	return f
}

func Add(n int) int {
	return n + 1000
}
