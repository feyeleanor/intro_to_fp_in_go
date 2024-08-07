package main

import (
	"fmt"
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
		for i := 0; ; i++ {
			p := []Value{ValueOf(i)}
			if r := s.Call(p); r[1].Bool() {
				fmt.Printf("5>s(%v): %v\n", i, s.Call(p)[0].Interface())
			} else {
				return
			}
		}
	case Slice:
		defer func() {
			recover()
		}()
		for i := 0; ; i++ {
			fmt.Printf("6>[%v]: %v\n", i, s.Index(i).Interface())
		}
	}
}
