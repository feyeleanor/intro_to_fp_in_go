package main

import . "fmt"

func main() {
	c := make(chan Entry[int])
	go func(c chan Entry[int], limit int) {
		for i := range limit {
			c <- Entry[int]{i, i * 2}
		}
		close(c)
	}(c, 5)
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
