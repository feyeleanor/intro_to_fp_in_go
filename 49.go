package main

import . "fmt"

func main() {
	s := &Range[int]{0, 2, Limit(5)}
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
	Limit
}

func (r Range[T]) Each(f func(T)) {
	if r.step == 1 {
		for i := range int(r.Limit - Limit(r.start)) {
			f(T(i))
		}
	} else {
		r.Limit.Each(func(i int) {
			f(r.start)
			r.start += r.step
		})
	}
}

type Limit int

func (i Limit) Each(f func(int)) {
	for v := range i {
		f(int(v))
	}
}
