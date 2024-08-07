package main

import . "fmt"

func main() {
	s := []int{0, 2, 4, 6, 8}
	print_slice(func(i int) int {
		return s[i]
	})
}

func print_slice[T any](f Iterable[T]) {
	f.Each(func(i int, v T) {
		Printf("%v: %v\n", i, f(i))
	})
}

type Iterable[T any] func(int) T

func (iter Iterable[T]) Each(f func(int, T)) {
	defer func() {
		recover()
	}()
	for i := 0; ; i++ {
		f(i, iter(i))
	}
}
