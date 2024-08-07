package main

import . "fmt"

func main() {
	c1, g1 := makeGenerator[int](0, 2, 4, 6, 8)
	go g1()
	print_channel(c1)

	c2, g2 := makeGenerator[int](0, 2, 4, 6, 8)
	go g2()
	print_channel(c2)
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

func makeGenerator[T any](s ...T) (chan Entry[T], func()) {
	c := make(chan Entry[T])
	return c, func() {
		for i, v := range s {
			c <- Entry[T]{i, v}
		}
		close(c)
	}
}
