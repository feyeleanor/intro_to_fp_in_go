package main

import . "fmt"

func main() {
	s := []int{0, 2, 4, 6, 8}
	print_values[int](s)
	print_values[int](func(i int) (v int, ok bool) {
		if ok = (-1 < i) && (i < len(s)); ok {
			v = s[i]
		}
		return
	})
}

func print_values[T any](s any) {
	switch s := s.(type) {
	case Generator[T]: // Not matched by func(int) (int, bool)
		for i := 0; ; i++ {
			if v, ok := s(i); ok {
				Printf("1>s(%v): %v\n", i, v)
			} else {
				return
			}
		}
	case Sequence[T]: // Not matched by []int
		for i, v := range s {
			Printf("2>s[%v]: %v\n", i, v)
		}
	}
}

type Sequence[T any] []T
type Generator[T any] func(int) (T, bool)
