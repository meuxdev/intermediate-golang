package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// creating the wait group
	var wg sync.WaitGroup
	// creating the channel with the size of the buffer
	// the buffer will limit the amount of go routines exc.
	// in this case we will limit this to 2 at the sync
	c := make(chan int, 2)

	for i := 0; i < 10; i++ {
		// adding the number to the buffer
		c <- 1
		wg.Add(1)
		go doSomething(i, &wg, c)
	}

	wg.Wait()
}

func doSomething(i int, wg *sync.WaitGroup, c chan int) {
	defer wg.Done()

	fmt.Printf("Id %d started...\n", i)
	time.Sleep(time.Second * 4)
	fmt.Printf("Id %d ended...\n", i)
	// release the value so that the new go routines will  be exc.
	<-c

}
