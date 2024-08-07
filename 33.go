package main

import . "fmt"

func main() {
	s := []int{0, 2, 4, 6, 8}
	print_values[int](Sequence[int](s))
	print_values[int](Generator[int](func(i int) (v int, ok bool) {
		if ok = (-1 < i) && (i < len(s)); ok {
			v = s[i]
		}
		return
	}))
}

func print_values[T any](s any) {
	switch s := s.(type) {
	case func(int) (T, bool): // Not matched by Generator[T]
		for i := 0; ; i++ {
			if v, ok := s(i); ok {
				Printf("3>s(%v): %v\n", i, v)
			} else {
				return
			}
		}
	case []T: // Not matched by Sequence[T]
		for i, v := range s {
			Printf("4>s[%v]: %v\n", i, v)
		}
	}
}

type Sequence[T any] []T
type Generator[T any] func(int) (T, bool)
