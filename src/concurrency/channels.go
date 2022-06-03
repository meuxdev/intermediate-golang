package main

import "fmt"

func main() {
	// make(chan <typeValue>, <bufferSize>)
	c := make(chan int, 3)

	c <- 1
	c <- 2
	c <- 3
	// passing the limit of buffers
	// c <- 4

	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c)
	// passing the limit of buffers
	// fmt.Println(<-c)
}
