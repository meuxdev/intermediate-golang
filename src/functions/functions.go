package main

import (
	"fmt"
	"time"
)

func main() {
	x := 5
	z := 10

	// anym function
	y := func(number int) int {
		return number * 2
	}

	double := y(x)
	double2 := y(z)

	fmt.Println(double)
	fmt.Println(double2)

	c := make(chan int)

	go func() {
		fmt.Println("Starting the process...")
		time.Sleep(5 * time.Second)
		fmt.Println("Ending the process...")
		c <- 1
	}()
	<-c

}
