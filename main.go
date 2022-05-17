package main

import (
	"fmt"
	"log"
	"strconv"
)

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

	s := []string{"Tacos", "Tortas", "Hamborguesas"}

	s = append(s, "Chesco")

	for index, value := range s {
		fmt.Println(index)
		fmt.Println(value)
	}

}
