package main

import "testing"

func TestSum(t *testing.T) {
	/*
		total := Sum(5, 5)

		if total != 10 {
			t.Errorf("Sum was incorrected, got %d expected %d", total, 10)
		}
	*/

	tables := []struct {
		a int
		b int
		n int
	}{
		{1, 2, 3},
		{2, 2, 4},
		{25, 26, 51},
	}

	for _, item := range tables {
		total := Sum(item.a, item.b)
		if total != item.n {
			t.Errorf("Sum was incorrected, got %d expected %d", total, item.n)
		}
	}
}

func TestMaxFunc(t *testing.T) {
	tables := []struct {
		a       int
		b       int
		maxNumb int
	}{
		{4, 2, 4},
		{2, 4, 4},
	}

	for _, item := range tables {
		max := GetMax(item.a, item.b)
		if max != item.maxNumb {
			t.Errorf("Get Mas was incorrect, got %d, expected %d", max, item.maxNumb)
		}
	}
}

func TestFib(t *testing.T) {
	tables := []struct {
		query  int
		result int
	}{
		{1, 1},
		{8, 21},
		{11, 89},
		{25, 75025},
	}

	for _, item := range tables {
		fib := Fibonacci(item.query)

		if fib != item.result {
			t.Errorf("Fibonacci was incorrect, expected %d received %d", item.result, fib)
		}
	}
}
