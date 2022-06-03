package main

import "fmt"

// c <-chan === Reading
// c chan <- === Writting

func main() {
	generator := make(chan int)
	doubles := make(chan int)

	go Generator(generator)
	go Double(generator, doubles)
	PrintChannels(doubles)
}

func Generator(c chan<- int) {
	for i := 1; i <= 10; i++ {
		c <- i
	}
	close(c)
}

func Double(in <-chan int, out chan<- int) {
	for value := range in {
		out <- 2 * value
	}
	close(out)
}

func PrintChannels(c <-chan int) {
	for val := range c {
		fmt.Println(val)
	}
}
