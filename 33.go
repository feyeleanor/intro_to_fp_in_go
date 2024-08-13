package main

import (
	"fmt"
	"math"
)

func main() {
	s := []int{0, 2, 4, 6, 8}
	print_values[int](Sequence[int](s))
	print_values[int](Generator[int](func(i int) (v int, ok bool) {
		if ok = (-1 < i) && (i < len(s)); ok {
			v = s[i]
		}
		return
	}))
	print_values[int](Iterator[int](func(g func(i, v int) bool) {
		for i, v := range s {
			if !g(i, v) {
				return
			}
		}
	}))
}

func print_values[T any](s any) {
	switch s := s.(type) {
	case []T: // Not matched by Sequence[T]
		for i, v := range s {
			fmt.Printf("1>%v: %v\n", i, v)
		}
	case func(int) (T, bool): // Not matched by Generator[T]
		for i := range math.MaxInt {
			if v, ok := s(i); ok {
				fmt.Printf("2>%v: %v\n", i, v)
			} else {
				return
			}
		}
	case func(func(int, T) bool): // Not matched by Iterator[T]
		for i, v := range s {
			fmt.Printf("3>%v: %v\n", i, v)
		}
	}
}

type Sequence[T any] []T
type Generator[T any] func(int) (T, bool)
type Iterator[T any] func(func(int, T) bool)
