package main

import . "fmt"

func main() {
	print_slice(yield, 0, 2, 4, 6, 8)
}

func print_slice[T any](f Iterator[T], s ...T) {
	defer func() {
		recover()
	}()
	for i := 0; ; i++ {
		Printf("%v: %v\n", i, f(i, s...))
	}
}

func yield[T any](i int, s ...T) T {
	return s[i]
}

type Iterator[T any] func(int, ...T) T
