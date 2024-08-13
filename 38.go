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
	case Sequence[T]:
		print_values[T]([]T(s))
	case []T:
		for i, v := range s {
			fmt.Printf("1>%v: %v\n", i, v)
		}
	case Generator[T]:
		print_values[T]((func(int) (T, bool))(s))
	case func(int) (T, bool):
		for i := range math.MaxInt {
			if v, ok := s(i); ok {
				fmt.Printf("2>%v: %v\n", i, v)
			} else {
				return
			}
		}
	default:
		switch s := ValueOf(s); s.Kind() {
		case Func:
			for i := range math.MaxInt {
				p := []Value{ValueOf(i)}
				if r := s.Call(p); r[1].Bool() {
					fmt.Printf("3>%v: %v\n", i, s.Call(p)[0].Interface())
				} else {
					return
				}
			}
		case Slice:
			for i := range s.Len() {
				fmt.Printf("4>%v: %v\n", i, s.Index(i).Interface())
			}
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
