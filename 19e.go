package main

import (
	"fmt"
	"iter"
	"math"
)

func main() {
	s := []int{0, 2, 4, 6, 8}
	print_slice(func(i int) int {
		return s[i]
	})
}

func print_slice[T any](f Iterable[T]) {
	for i, v := range f.Each() {
		fmt.Printf("%v: %v\n", i, v)
	}
}

type Iterable[T any] func(int) T

func (f Iterable[T]) Each() iter.Seq2[int, T] {
	return func(g func(int, T) bool) {
		defer func() {
			recover()
		}()
		for i := range math.MaxInt {
			g(i, f(i))
		}
	}
}
