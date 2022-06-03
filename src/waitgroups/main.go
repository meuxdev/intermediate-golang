package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		// adding a task to the wait group
		wg.Add(1)
		go doSomething(i, &wg)
	}

	// waits until the index of the tasks in the waitgroup is 0.
	wg.Wait()
}

func doSomething(i int, wg *sync.WaitGroup) {
	// defer exec. that at the end of the func.
	// Done substracs 1 to the index of tasks of the wait group
	defer wg.Done()
	fmt.Printf("Starting the task number %d...\n", i)
	time.Sleep(time.Second * 2)
	fmt.Printf("Ending the task number %d...\n", i)
}
