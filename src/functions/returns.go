package main

import "fmt"

// ... -> slice
func sum(values ...int) int {
	total := 0
	for _, num := range values {
		total += num
	}
	return total
}

func printNames(names ...string) {
	for _, name := range names {
		fmt.Println(name)
	}
}

func getValues(x int) (double int, triple int, quadra int) {
	double = 2 * x
	triple = 3 * x
	quadra = 4 * x
	return // looks for the values of the vars and returns auto
}

func main() {
	fmt.Println(sum(1, 2, 4, 2, 2, 3, 4, 5, 7, 8, 9))
	fmt.Println(sum(1, 2, 4))
	fmt.Println(sum(1, 2))
	fmt.Println(sum(1))
	fmt.Println(sum())
	printNames("Alicia", "David")
	fmt.Println("*************************")
	printNames("Nacho", "Alex", "Sofia", "Rick")
	fmt.Println(getValues(2))
}
