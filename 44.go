package main

import "fmt"

func main() {
	print_values(Limit(5))
	print_values(Slice[int]{0, 2, 4, 6, 8})
}

func print_values[T any](s Iterable[T]) (i int) {
	s.Each(func(v T) {
		fmt.Printf("%v: %v\n", i, v)
		i++
	})
	return i
}

type Iterable[T any] interface {
	Each(func(T))
}

type Limit int

func (i Limit) Each(f func(int)) {
	for v := range i {
		f(int(v))
	}
}

type Slice[T any] []T

func (s Slice[T]) Each(f func(T)) {
	for _, v := range s {
		f(v)
	}
}
