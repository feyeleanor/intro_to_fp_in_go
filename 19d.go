package main

import (
	"fmt"
	"iter"
)

func main() {
	s := []int{0, 2, 4, 6, 8}
	print_slice[int](func(g func(i, v int) bool) {
		for i, v := range s {
			if !g(i, v) {
				return
			}
		}
	})
	print_slice[int](each(s))
}

func each[T any](s []T) iter.Seq2[int, T] {
	return func(g func(int, T) bool) {
		for i, v := range s {
			if !g(i, v) {
				return
			}
		}
	}
}

var lines int

func print_slice[T any](f iter.Seq2[int, T]) {
	lines++
	for i, v := range f {
		fmt.Printf("%v>%v: %v\n", lines, i, v)
	}
}
