package main

import "fmt"

func main() {
	c := make(chan Entry[int])
	go func(c chan Entry[int], s ...int) {
		for i, v := range s {
			c <- Entry[int]{i, v}
		}
		close(c)
	}(c, 0, 2, 4, 6, 8)
	print_channel(c)
}

func print_channel[T any](c chan Entry[T]) (i int) {
	for v := range c {
		fmt.Printf("%v: %v\n", v.i, v.v)
		i = v.i
	}
	return
}

type Entry[T any] struct {
	i int
	v T
}
