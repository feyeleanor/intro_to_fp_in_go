package main

import . "fmt"

func main() {
	c := make(chan Entry[int])
	go sequence(c, 0, 2, 4, 6, 8)
	print_channel(c)
}

func print_channel[T any](c chan Entry[T]) (i int) {
	for v := range c {
		Printf("%v: %v\n", v.i, v.v)
		i = v.i
	}
	return
}

type Entry[T any] struct {
	i int
	v T
}

func sequence[T any](c chan Entry[T], s ...T) {
	for i, v := range s {
		c <- Entry[T]{i, v}
	}
	close(c)
}
