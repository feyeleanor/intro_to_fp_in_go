package main

import (
	"fmt"
	"math"
)

func main() {
	print_slice([]int{0, 2, 4, 6, 8}, yield)
}

func print_slice(s []int, f func([]int, int) int) {
	defer func() {
		recover()
	}()
	for i := range math.MaxInt {
		fmt.Printf("%v: %v\n", i, f(s, i))
	}
}

func yield(s []int, i int) int {
	return s[i]
}
