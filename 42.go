package main

import "fmt"

func main() {
	var s Iterable = Slice{0, 2, 4, 6, 8}
	i := 0
	s.Each(func(v any) {
		fmt.Printf("%v: %v\n", i, v)
		i++
	})
}

type Iterable interface {
	Each(func(any))
}

type Slice []int

func (s Slice) Each(f func(any)) {
	for _, v := range s {
		f(v)
	}
}
