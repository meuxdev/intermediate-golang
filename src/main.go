package main

import (
	"fmt"
	"log"
	"strconv"
	"time"
)

func doSomething(c chan<- int) {
	time.Sleep(3 * time.Second)
	fmt.Println("DONE!!!")
	c <- 1
}

func main() {
	var x int
	x = 8
	y := 7
	fmt.Println(x)
	fmt.Println(y)

	myValue, err := strconv.ParseInt("6", 0, 64)

	if err != nil {
		log.Fatal("Error with the parse process", err)
	} else {
		fmt.Println(myValue)
	}

	m := make(map[string]int)
	var key string = "KEY"
	m[key] = 113

	fmt.Println(m[key])

	s := []string{"Tacos", "Hot-Dogs", "Hamburgers"}

	s = append(s, "Orange Juice")

	for index, value := range s {
		fmt.Println(index)
		fmt.Println(value)
	}

	c := make(chan int)
	go doSomething(c)
	<-c
	g := 25
	fmt.Println(g)
	h := &g // pointer -> reference

	fmt.Println(h)
	fmt.Println(*h) // access to the value -> *

}
