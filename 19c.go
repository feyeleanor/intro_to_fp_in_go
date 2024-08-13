package main

import "fmt"

func main() {
	s := []int{0, 2, 4, 6, 8}
	print_slice(func(g func(i, v int) bool) {
		for i, v := range s {
			if !g(i, v) {
				return
			}
		}
	})
	print_slice(each(s))
}

func each[T any](s []T) func(func(int, T) bool) {
	return func(g func(int, T) bool) {
		for i, v := range s {
			if !g(i, v) {
				return
			}
		}
	}
}

var lines int

func print_slice[T any](f func(func(int, T) bool)) {
	lines++
	for i, v := range f {
		fmt.Printf("%v>%v: %v\n", lines, i, v)
	}
}
