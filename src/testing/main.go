package main

func Sum(x, y int) (result int) {
	result = x + y
	return
}

func GetMax(x, y int) (max int) {
	if x > y {
		max = x
	} else {
		max = y
	}

	return
}

type NodeFibo struct {
	val    int
	result int
}

func Fibonacci(n int) (result int) {

	if n <= 1 {
		result = n
	} else {
		result = Fibonacci(n-1) + Fibonacci(n-2)
	}
	return
}

func main() {
}
