package main

import "fmt"

func main() {
	c := make(Pipeline[int])
	go func(p Pipeline[int], s ...int) {
		p.push(s...)
		close(p)
	}(c, 0, 2, 4, 6, 8)
	print_channel(c)
}

func print_channel[T any](p Pipeline[T]) (c int) {
	p.each(func(i int, v T) {
		fmt.Printf("%v: %v\n", i, v)
		c = i
	})
	return
}

type Entry[T any] struct {
	i int
	v T
}

type Pipeline[T any] chan Entry[T]

func (p Pipeline[T]) push(s ...T) {
	for i, v := range s {
		p <- Entry[T]{i, v}
	}
}

func (p Pipeline[T]) each(f func(int, T)) {
	for v := range p {
		f(v.i, v.v)
	}
}
