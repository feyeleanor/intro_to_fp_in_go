package main

import (
	"fmt"
	"math"
	. "reflect"
)

func main() {
	s := []int{0, 2, 4, 6, 8}
	print_values(s)
	print_values(func(i int) (v int, ok bool) {
		if ok = (-1 < i) && (i < len(s)); ok {
			v = s[i]
		}
		return
	})
}

func print_values(s any) {
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
				fmt.Printf("6>[%v]: %v\n", i, s.Index(i).Interface())
			}
			return
		})
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
