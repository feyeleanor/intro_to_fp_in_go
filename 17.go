package main

import (
	"fmt"
	"math"
)

func main() {
	s := []int{0, 2, 4, 6, 8}
	print_slice(func(i int) int {
		return s[i]
	})
}

func print_slice[T any](f Iterable[T]) {
	defer func() {
		recover()
	}()
	for i := range math.MaxInt {
		fmt.Printf("%v: %v\n", i, f(i))
	}
}

type Iterable[T any] func(int) T
