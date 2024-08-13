package main

import (
	"fmt"
	"math"
)

func main() {
	print_slice(yield, 0, 2, 4, 6, 8)
}

func print_slice(f func(int, ...int) int, s ...int) {
	defer func() {
		recover()
	}()
	for i := range math.MaxInt {
		fmt.Printf("%v: %v\n", i, f(i, s...))
	}
}

func yield(i int, s ...int) int {
	return s[i]
}
