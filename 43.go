package main

import "fmt"

func main() {
	var s Iterable[int] = Slice[int]{0, 2, 4, 6, 8}
	i := 0
	s.Each(func(v int) {
		fmt.Printf("%v: %v\n", i, v)
		i++
	})
}

type Iterable[T any] interface {
	Each(func(T))
}

type Slice[T any] []T

func (s Slice[T]) Each(f func(T)) {
	for _, v := range s {
		f(v)
	}
}
