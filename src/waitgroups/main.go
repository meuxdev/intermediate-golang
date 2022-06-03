package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go doSomething(i, &wg)
	}

	wg.Wait()
}

func doSomething(i int, wg *sync.WaitGroup) {
	// defer exec. that at the end of the func.
	defer wg.Done()
	fmt.Printf("Starting the task number %d...\n", i)
	time.Sleep(time.Second * 2)
	fmt.Printf("Ending the task number %d...\n", i)
}
