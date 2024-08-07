package main

import (
	"fmt"
	"math"
	. "reflect"
)

func main() {
	s := []int{0, 2, 4, 6, 8}
	print_values[int](s)
	print_values[int](func(i int) (v int, ok bool) {
		if ok = (-1 < i) && (i < len(s)); ok {
			v = s[i]
		}
		return
	})

	print_values[int](Sequence[int](s))
	print_values[int](Generator[int](func(i int) (v int, ok bool) {
		if ok = (-1 < i) && (i < len(s)); ok {
			v = s[i]
		}
		return
	}))

	print_values[uint](s)
	print_values[uint](func(i int) (v int, ok bool) {
		if ok = (-1 < i) && (i < len(s)); ok {
			v = s[i]
		}
		return
	})
}

func print_values[T any](s any) {
	switch s := s.(type) {
	case Generator[T]:
		for_each(func(i int) bool {
			if v, ok := s(i); ok {
				fmt.Printf("1>s(%v): %v\n", i, v)
				return true
			}
			return false
		})
	case Sequence[T]:
		for i, v := range s {
			fmt.Printf("2>s[%v]: %v\n", i, v)
		}
	case func(int) (T, bool):
		for_each(func(i int) bool {
			if v, ok := s(i); ok {
				fmt.Printf("3>s(%v): %v\n", i, v)
				return true
			}
			return false
		})
	case []T:
		for i, v := range s {
			fmt.Printf("4>s[%v]: %v\n", i, v)
		}
	default:
		switch s := ValueOf(s); s.Kind() {
		case Func:
			for_each(func(i int) (ok bool) {
				p := []Value{ValueOf(i)}
				r := s.Call(p)
				if ok = r[1].Bool(); ok {
					fmt.Printf("5>s(%v): %v\n", i, s.Call(p)[0].Interface())
				}
				return
			})
		case Slice:
			for_each(func(i int) (ok bool) {
				if ok = i < s.Len(); ok {
					fmt.Printf("6>s[%v]: %v\n", i, s.Index(i).Interface())
				}
				return
			})
		}
	}

}

func for_each(f func(int) bool) (i int) {
	for i = range math.MaxInt {
		if !f(i) {
			return
		}
	}
	return
}

type Sequence[T any] []T
type Generator[T any] func(int) (T, bool)
