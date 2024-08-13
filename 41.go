package main

import "fmt"

func main() {
	s := Slice[int]{0, 2, 4, 6, 8}
	i := 0
	s.Each(func(v int) {
		fmt.Printf("%v: %v\n", i, v)
		i++
	})
}

type Slice[T any] []T

func (i Slice[T]) Each(f func(T)) {
	for _, v := range s {
		f(v)
	}
}
