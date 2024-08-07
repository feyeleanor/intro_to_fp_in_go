package main

import . "fmt"

func main() {
	s := &Range[int]{0, 2, 5}
	print_values(s)
	print_values(s)
}

func print_values[T any](s Iterable[T]) (i int) {
	s.Each(func(v T) {
		Printf("%v: %v\n", i, v)
		i++
	})
	return i
}

type Iterable[T any] interface {
	Each(func(T))
}

type Ordinal interface {
	int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64
}

type Range[T Ordinal] struct {
	start T
	step  T
	count T
}

func (r Range[T]) Each(f func(T)) {
	for ; r.count > 0; r.count-- {
		f(r.start)
		r.start += r.step
	}
}
