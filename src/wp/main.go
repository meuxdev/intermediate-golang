package main

import "fmt"

func main() {
	tasks := []int{2, 3, 4, 5, 6, 7, 10, 12, 40, 39, 37}
	nWorkers := 3
	jobs := make(chan int, len(tasks))
	results := make(chan int, len(tasks))

	for i := 0; i < nWorkers; i++ {
		go Worker(i, jobs, results)
	}

	for _, value := range tasks {
		jobs <- value
	}

	close(jobs)

	for r := 0; r < len(tasks); r++ {
		<-results
	}
}

func Worker(id int, jobs <-chan int, results chan<- int) {
	for job := range jobs {
		fmt.Println("*********************************************************")
		fmt.Printf("Worker with id -> %d started...\nFib of %d\n", id, job)
		fib := Fibonacci(job)
		fmt.Println("*********************************************************")
		fmt.Printf("Worker with id ->%d finished...\nFib of %d\nResult of fib %d\n", id, job, fib)
		results <- fib
	}
}

func Fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}
